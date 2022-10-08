package client

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.youshu.cc/go/protobuf/go/common"
	gin2 "gitlab.youshu.cc/go/yslib/gin"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type IClient interface {
	GetAppName() string
	GetRetryTimes() int
	GetEndpoint() string
	GetApiInfo(key string) map[string]string
	HttpClient() http.Client
}

type Config struct {
	RetryTimes        int
	HttpClientTimeout time.Duration
}

// Error 用户描述common.response中的errno、errmsg错误信息，业务端需要关心
type Error struct {
	errNo  int32
	errMsg string
}

func NewError(code int32, msg string) *Error {
	return &Error{
		errNo:  code,
		errMsg: msg,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("errno:%d errmsg:%s", e.errNo, e.errMsg)
}

// GetCallerName 获取上游调用者的方法名称
func GetCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

// Execute 解析、执行请求，返回响应
func Execute(ctx context.Context, c IClient, request interface{}) (*anypb.Any, error) {
	var u url.URL
	var err error
	query := u.Query()   // query 参数
	form := url.Values{} // form 参数

	// 获取当前方法的调用者名称
	callerNameTokens := strings.Split(GetCallerName(), ".")
	apiInfoKey := callerNameTokens[len(callerNameTokens)-1]
	// 根据调用者方法名获取api接口信息
	apiInfo := c.GetApiInfo(apiInfoKey)
	if apiInfo == nil {
		return nil, NewError(-1, "api not found")
	}
	if c.GetEndpoint() == "" {
		return nil, NewError(-1, "invalid endpoint")
	}

	path := apiInfo["path"]
	method := strings.ToUpper(apiInfo["method"])
	uri := c.GetEndpoint() + path

	// 通过反射获取pb的请求类型的值的信息，并构造http的请求
	// 注：request为接口类型，它包裹的值类型须为指针类型。
	refType := reflect.TypeOf(request).Elem()
	refVal := reflect.ValueOf(request).Elem()
	for i := 0; i < refType.NumField(); i++ {
		refFieldType := refType.Field(i)
		refFieldVal := refVal.Field(i)
		jsonTag := refFieldType.Tag.Get("json")
		if refFieldType.IsExported() && jsonTag != "" {
			jsonTags := strings.Split(jsonTag, ",")
			refFieldValStr := ""
			switch refFieldVal.Kind() {
			case reflect.Int8, reflect.Uint8, reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32, reflect.Int, reflect.Uint, reflect.Int64, reflect.Uint64:
				refFieldValStr = strconv.FormatInt(refFieldVal.Int(), 10)
			case reflect.Float32, reflect.Float64:
				refFieldValStr = strconv.FormatFloat(refFieldVal.Float(), 'g', 6, 64)
			case reflect.String:
				refFieldValStr = refFieldVal.String()
			}

			query.Add(jsonTags[0], refFieldValStr)
			form.Add(jsonTags[0], refFieldValStr)
		}
	}

	// 设置公共header
	var headers http.Header
	ginCtx, ok := ctx.(*gin.Context)
	if ok {
		headers = gin2.GetXHeaders(ginCtx)
	}

	var req *http.Request
	httpClient := c.HttpClient()
	if method == http.MethodGet {
		query.Add("output", "pb")
		queryStr := query.Encode()
		uri += "?" + queryStr
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
		req.Header = headers
		if err != nil {
			return nil, NewError(-1, err.Error())
		}
	} else if method == http.MethodPost {
		req, err = http.NewRequestWithContext(ctx, http.MethodPost, uri+"?output=pb", strings.NewReader(form.Encode()))
		headers.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header = headers
		if err != nil {
			return nil, NewError(-1, err.Error())
		}
	} else {
		return nil, NewError(-1, "not supported method")
	}

	// 重试发送
	var e *Error
	var any *anypb.Any
	for i := 0; i < c.GetRetryTimes(); i++ {
		any, e = executeSend(httpClient, req)
		if e == nil {
			return any, nil
		}
	}

	retryErr := NewError(e.errNo, fmt.Sprintf("retry %d times, still fails; %s", c.GetRetryTimes(), e.errMsg))
	return nil, retryErr
}

//
func executeSend(httpClient http.Client, req *http.Request) (*anypb.Any, *Error) {
	httpClientResponse, err := httpClient.Do(req)
	if err != nil {
		return nil, NewError(-1, err.Error())
	}
	defer func() {
		_ = httpClientResponse.Body.Close()
	}()
	bodyBytes, err := io.ReadAll(httpClientResponse.Body)
	if err != nil {
		return nil, NewError(-1, err.Error())
	}

	// 请求响应处理
	response := &common.Response{}
	err = proto.Unmarshal(bodyBytes, response)
	if err != nil {
		return nil, NewError(-1, err.Error())
	}
	// 响应的错误码和错误信息可以从error中检测和获取
	if response.GetErrno() != 0 {
		return nil, NewError(response.GetErrno(), response.GetErrmsg())
	}
	// 返回相应的data部分，这里是一个pb的any类型，sdk会根据相应pb的responseData类型反序列化，最终返回给sdk调用者具体的responseData对象。
	return response.GetData(), nil
}
