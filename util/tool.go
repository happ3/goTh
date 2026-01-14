package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"os"
	"path"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// 获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// md5加密
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 把字符串解析成html
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}

// 表示把string转换成int
func Int(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// 表示把string转换成Float64
func Float(str string) (float64, error) {
	n, err := strconv.ParseFloat(str, 64)
	return n, err
}

// 表示把int转换成string
func String(n int) string {
	str := strconv.Itoa(n)
	return str
}

// Substr截取字符串
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	rl := len(rs)
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = 0
	}

	if end < 0 {
		end = rl
	}
	if end > rl {
		end = rl
	}
	if start > end {
		start, end = end, start
	}

	return string(rs[start:end])

}

func Sub(a int, b int) int {
	return a - b
}

func Mul(price float64, num int) float64 {
	return price * float64(num)
}

// 上传图片到本地
func LocalUploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err := c.FormFile(picName)

	// fmt.Println(file)
	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3、创建图片保存目录  static/upload/20210624

	day := GetDay()
	dir := "./static/upload/" + day

	err1 := os.MkdirAll(dir, 0666)
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil

}

// 生成随机数
func GetRandomNum() string {
	var str string
	for i := 0; i < 4; i++ {
		current := rand.Intn(10)
		str += String(current)
	}
	return str
}

// GetOrderId

func GetOrderId() string {
	// 2022020312233
	template := "20060102150405"
	return time.Now().Format(template) + GetRandomNum()
}

func GetKeys(obj interface{}) []string {
	result := make(map[string]interface{})
	var keys []string

	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	// 如果是指针，获取指向的元素
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// 直接使用字段名（反射获取）
		fieldName := fieldType.Name

		fieldValue := field.Interface()

		// 检查是否为零值（包含时间处理）
		if !isZeroValue(fieldValue) {
			result[fieldName] = fieldValue
			keys = append(keys, fieldName)
		}
	}
	return keys
}

// 零值判断函数（包含时间处理）
func isZeroValue(value interface{}) bool {
	if value == nil {
		return true
	}

	rv := reflect.ValueOf(value)

	switch rv.Kind() {
	case reflect.String:
		return rv.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Ptr:
		return rv.IsNil()
	case reflect.Interface:
		return rv.Interface() == nil
	case reflect.Slice, reflect.Array, reflect.Map:
		return rv.Len() == 0
	case reflect.Struct:
		// 特殊处理 time.Time
		if rv.Type().String() == "time.Time" {
			t, ok := value.(time.Time)
			if !ok {
				return true
			}
			return t.IsZero()
		}
		return rv.IsZero()
	default:
		return rv.IsZero()
	}
}
