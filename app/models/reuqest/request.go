package reuqest

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
)

//func GetJson(c *gin.Context) (map[string]interface{}, error) {
//	jsonstr, _ := io.ReadAll(c.Request.Body)
//	var data map[string]interface{}
//	err := json.Unmarshal(jsonstr, &data)
//	return data, err
//}

func GetJson(ctx *gin.Context) (map[string]interface{}, error) {
	data, err := io.ReadAll(ctx.Request.Body)
	maps := make(map[string]interface{})
	err = json.Unmarshal(data, &maps)
	return maps, err
}

func GetJsonToObj[T any](ctx *gin.Context) (T, error) {
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		var zero T // 返回类型的零值
		return zero, err
	}
	var t T
	err = json.Unmarshal(data, &t)
	if err != nil {
		return t, err
	}
	return t, nil
}
