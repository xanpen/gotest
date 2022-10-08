package learning

import (
	"context"
	iClient "github.com/xanpen/gotest/gin/client"
	"gitlab.youshu.cc/go/protobuf/go/learning/api/plan"
	"gitlab.youshu.cc/go/protobuf/go/learning/api/themebooklist"
	"gitlab.youshu.cc/go/yslib/conf"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"net/http"
	"time"
)

var envEndpointMap = map[string]string{
	"rd":     "http://127.0.0.1:8125",
	"qa":     "http://qa-learning.laidan.com",
	"online": "http://learning.youshu.cc",
	"luban":  "http://learning.laidan.com",
}

var apiInfoMap = map[string]map[string]string{
	"ApiV1PlanGetPlanSimpleInfoByIds": {
		"path":   "/api/v1/Plan/GetPlanSimpleInfoByIds",
		"method": "GET",
	},
	"ApiV1PlanGetPlanInfoByIds": {
		"path":   "/api/v1/Plan/GetPlanInfoByIds",
		"method": "GET",
	},
	"ApiV1ThemeBookListQuitTbl": {
		"path":   "/api/v1/ThemeBookList/QuitTbl",
		"method": "POST",
	},
}

type client struct {
	appName      string
	retryTimes   int
	apiInfos     map[string]map[string]string
	envEndpoints map[string]string
	httpClient   http.Client
}

func (c *client) HttpClient() http.Client {
	return c.httpClient
}

func (c *client) GetAppName() string {
	return c.appName
}

func (c *client) GetRetryTimes() int {
	return c.retryTimes
}

func (c *client) GetEndpoint() string {
	runtime := conf.GetRuntime()
	env := string(runtime.EnvMode)

	if endpoint, ok := envEndpointMap[env]; ok {
		return endpoint
	}
	return ""
}

func (c *client) GetApiInfo(key string) map[string]string {
	if apiInfo, ok := apiInfoMap[key]; ok {
		return apiInfo
	}
	return nil
}

func (c *client) ApiV1PlanGetPlanSimpleInfoByIds(ctx context.Context, r *plan.GetPlanSimpleInfoByIdsRequest) (*plan.GetPlanSimpleInfoByIdsResponseData, error) {
	respData := &plan.GetPlanSimpleInfoByIdsResponseData{}
	anyData, err := iClient.Execute(ctx, c, r)
	if err != nil {
		return respData, err
	}
	e := anypb.UnmarshalTo(anyData, respData, proto.UnmarshalOptions{DiscardUnknown: true})
	if e != nil {
		return respData, iClient.NewError(-1, e.Error())
	}
	return respData, nil
}

func (c *client) ApiV1ThemeBookListQuitTbl(ctx context.Context, r *themebooklist.QuitTblRequest) (*themebooklist.QuitTblResponseData, error) {
	respData := &themebooklist.QuitTblResponseData{}
	anyData, err := iClient.Execute(ctx, c, r)
	if err != nil {
		return respData, err
	}
	e := anypb.UnmarshalTo(anyData, respData, proto.UnmarshalOptions{DiscardUnknown: true})
	if e != nil {
		return respData, iClient.NewError(-1, e.Error())
	}
	return respData, nil
}

var Client *client

func NewDefaultClient() *client {
	if Client == nil {
		Client = &client{
			appName:      "learning",
			retryTimes:   1,
			apiInfos:     apiInfoMap,
			envEndpoints: envEndpointMap,
			httpClient: http.Client{
				Timeout: time.Duration(100) * time.Second,
			},
		}
	}
	return Client
}

func NewClient(conf *iClient.Config) *client {
	if Client == nil {
		Client = &client{
			appName:      "learning",
			retryTimes:   conf.RetryTimes,
			apiInfos:     apiInfoMap,
			envEndpoints: envEndpointMap,
			httpClient: http.Client{
				Timeout: conf.HttpClientTimeout,
			},
		}
	}
	return Client
}
