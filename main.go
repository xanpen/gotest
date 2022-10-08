package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xanpen/gotest/gin/client/learning"
	"gitlab.youshu.cc/go/protobuf/go/learning/api/themebooklist"
	"gitlab.youshu.cc/go/yslib/conf"
	"net/http"
	"net/url"
)

//curl -L -X POST '127.0.0.1:8125/api/v1/ThemeBookList/QuitTbl' \
//-H 'Content-Type: application/json' \
//--data-raw '{
//    "user_id":90002490,
//    "theme_book_list_id":1118
//}'
func main() {
	conf.SetRuntime(&conf.Runtime{
		AppName: "readwith",
		EnvMode: "rd",
	})
	ctx := &gin.Context{Request: &http.Request{
		URL: &url.URL{
			Scheme:     "",
			Opaque:     "",
			User:       nil,
			Host:       "test.com",
			Path:       "/test/test",
			RawPath:    "",
			ForceQuery: false,
			RawQuery:   "name=wong&age=33",
		},
		Header:     nil,
		Host:       "test.com",
		RemoteAddr: "111.111.111.111",
		RequestURI: "/test/test",
	}}
	/*r := &plan.GetPlanSimpleInfoByIdsRequest{
		PlanIds:      "1760,1758",
		PlanStatuses: "100",
		ShowClients:  "app",
		StudyClients: "app,tiny_course",
		PlanTags:     "",
	}
	respData, err := learning.NewDefaultClient().ApiV1PlanGetPlanSimpleInfoByIds(ctx, r)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(respData)*/

	r2 := &themebooklist.QuitTblRequest{
		UserId:          90002490,
		ThemeBookListId: 1118,
	}
	r2Data, err := learning.NewDefaultClient().ApiV1ThemeBookListQuitTbl(ctx, r2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(r2Data)
}
