package helper

import (
	"boilerplate/config"
	"boilerplate/exception"
	"boilerplate/helper/gofast"
	"boilerplate/model"
	"fmt"
	"github.com/goccy/go-json"
	"strings"

	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type DefaultExternalRequest struct {
	Key string `json:"key"`
}

func InsertRedis(payload model.SetDataRedis) {
	var redis = config.RedisConnection()
	jsonResult, err := json.Marshal(payload.Data)
	exception.PanicIfNeeded(err)

	parsedKey := strings.Replace(payload.Key, "appv3", "pupv1", -1)
	err = redis.Set(parsedKey, jsonResult, payload.Exp).Err()
	exception.PanicIfNeeded(err)

	if !payload.IsExternalDel {
		go DeleteExternalRedis(payload.Key)
	}
}

func GetRedis[T any](key string) (cek bool, raw T) {
	var redis = config.RedisConnection()

	parsedKey := strings.Replace(key, "appv3", "pupv1", -1)
	value, _ := redis.Get(parsedKey).Result()
	if value == "" {
		return false, raw
	}

	_ = json.Unmarshal([]byte(value), &raw)

	return true, raw
}

func DelRedis(key string) {
	var redis = config.RedisConnection()
	parsedKey := strings.Replace(key, "appv3", "pupv1", -1)
	redis.Eval("for i, name in ipairs(redis.call('KEYS', '"+parsedKey+"')) do redis.call('expire', name, 0); end", []string{"*"})
	go DeleteExternalRedis(key)
}

func DeleteExternalRedis(key string) {
	fast := gofast.New()
	var out interface{}
	body := DefaultExternalRequest{
		Key: key,
	}

	baseUrl := os.Getenv("TOPINDO_BASE_URL")
	uri := fmt.Sprintf("%s/h2h/redis/delete", baseUrl)
	if err := fast.Post(uri, &body, &out, nil); err != nil {
		fmt.Println("Failed to fetch", uri, "\n", err)
		return
	}
	return
}

func CreateToken(request model.JwtPayload) *model.TokenDetails {
	accessExpired, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_MINUTE"))

	td := &model.TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * time.Duration(accessExpired)).Unix()

	keyAccess, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("JWT_ACCESS_PRIVATE_KEY")))
	exception.PanicIfNeeded(err)

	now := time.Now().UTC()

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = request.UserId
	atClaims["username"] = request.Username
	atClaims["exp"] = td.AtExpires
	atClaims["iat"] = now.Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	at.Header["topindopay"] = "jwt"
	td.AccessToken, err = at.SignedString(keyAccess)

	if err != nil {
		exception.PanicIfNeeded(errors.New(model.AUTHENTICATION_FAILURE_ERR_TYPE))
	}

	return td

}

func CreateAuth(request model.JwtPayload, td *model.TokenDetails) {
	at := time.Unix(td.AtExpires, 0)
	now := time.Now()

	InsertRedis(model.SetDataRedis{
		Key:           td.AccessToken,
		Data:          request,
		Exp:           at.Sub(now),
		IsExternalDel: true,
	})
}

func ValidateFileExtByFilename(filename string) (err error) {
	arrFilename := strings.Split(filename, ".")
	ext := arrFilename[len(arrFilename)-1]
	if ext != "jpg" && ext != "jpeg" && ext != "png" {
		return exception.EXTENSION_NOT_ALLOWED
	}
	return err
}

func ValidateIconExtByFilename(filename string) (err error) {
	arrFilename := strings.Split(filename, ".")
	ext := arrFilename[len(arrFilename)-1]
	if ext != "svg" {
		return exception.EXTENSION_NOT_ALLOWED
	}
	return err
}

func ValidateExcelExtByFilename(filename string) (err error) {
	arrFilename := strings.Split(filename, ".")
	ext := arrFilename[len(arrFilename)-1]
	if ext != "xlsx" {
		return exception.EXTENSION_NOT_ALLOWED
	}
	return err
}

func FirstLastDayMonth() (firstDay, lastDay time.Time) {
	firstDay = time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)
	lastDay = time.Date(time.Now().Year(), time.Now().Month()+1, 1, 0, 0, 0, 0, time.Local).Add(-time.Nanosecond)
	return firstDay, lastDay
}

func TimeToFormatted(t time.Time) string {
	year, month, day := t.Date()

	if t.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d %s %d | %s WIB", day, ReadableMonthMap[int(month)], year, t.Format("15:04"))
}

func TimeToFormatted2(t time.Time) string {
	year, month, day := t.Date()

	if t.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d %s %d", day, ReadableMonthMap[int(month)], year)
}

func FormatTime(t time.Time, format string) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(format)
}
