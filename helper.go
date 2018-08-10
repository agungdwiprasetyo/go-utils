package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/agungdwiprasetyo/go-utils/parser"
)

type KeyValue struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func GetDateNow() string {
	return strings.Split((time.Now().String()), " ")[0]
}
func GetDateString(date time.Time) string {
	var day, month = fmt.Sprint(date.Day()), fmt.Sprint(int(date.Month()))
	if date.Day() < 10 {
		day = fmt.Sprintf("0%s", day)
	}
	if date.Month() < 10 {
		month = fmt.Sprintf("0%s", month)
	}
	return fmt.Sprintf("%s-%s-%d", day, month, date.Year())
}

func ParseDateString(date string) string {
	dt, _ := parser.ParseTime(date)
	return fmt.Sprintf("%d %s %d", dt.Day(), dt.Month().String(), dt.Year())
}

// NormalizeInt untuk konversi tipe data pointer integer ke integer biasa
func NormalizeInt(x *int) (y int) {
	if x == nil {
		y = 0
	} else {
		y = *x
	}
	return
}

// NormalizeString untuk konversi tipe data pointer string ke string biasa
func NormalizeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// NormalizePointerString untuk mengambil alamat dari tipe data string
func NormalizePointerString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
func NormalizeDate(dt string) string {
	if dt == "" {
		return "1900-01-01"
	}
	return dt
}

// ParseInt untuk konversi tipe data string ke tipe data integer (menggunakan package strconv)
func ParseInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
func ConvertToString64(val int64) string {
	res := strconv.Itoa(int(val))
	return res
}
func FloatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', 0, 64)
}
func IsNilObject(object interface{}) bool {
	return reflect.DeepEqual(object, reflect.Zero(reflect.TypeOf(object)).Interface())
}

// ConvertObjectToArray is method for convert object to array of key-value
func ConvertObjectToArray(object map[string]string) interface{} {
	var result []KeyValue

	for key, val := range object {
		result = append(result, KeyValue{Key: key, Value: val})
	}
	return result
}

func MD5(str string) string {
	hash1 := md5.New()
	hash1.Write([]byte(str))
	return hex.EncodeToString(hash1.Sum(nil))
}

func SHA1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// ComputeHmac256 is ...
func ComputeHmac256(str, salt string) string {
	key := []byte(salt)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func HashFilename() string {
	filename := time.Now().Unix()
	hs := md5.New()
	hs.Write([]byte(fmt.Sprint(filename)))
	return hex.EncodeToString(hs.Sum(nil))
}

// BinarySearch is algorithm for search value in slice/array, with complexity O(log n)
func BinarySearch(val int, arr []int) int {
	var has int
	n := len(arr)
	first, last := 0, n-1
	for first <= last {
		mid := (first + last) / 2
		if val > arr[mid] {
			first = mid + 1
		} else if val < arr[mid] {
			last = mid - 1
		} else {
			has = mid
			break
		}
	}
	return has
}

// STRING HELPER
// RandomString is method for generate user password
func RandomString(n int) string {
	alphaNum := `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	var letterRunes = []rune(alphaNum)

	b := make([]rune, n)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandomInt(n int) string {
	rand.Seed(time.Now().UnixNano())
	alphaNum := `0123456789`
	var letterRunes = []rune(alphaNum)
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetAbbrFromTitle(s string) string {
	arr := strings.Split(s, " ")
	var res []byte
	for _, val := range arr {
		ch := val[0]
		if ch >= 65 && ch <= 90 {
			res = append(res, ch)
		}
	}
	if len(res) < 3 {
		res = nil
		for _, val := range arr {
			if val[0] >= 65 && val[0] <= 90 {
				ch := []byte(val[0:3])
				res = append(res, ch...)
			}
		}
	}

	return string(res)
}

// Slice Helper
func ShuffleSliceInt(sl []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range sl {
		j := rand.Intn(i + 1)
		sl[i], sl[j] = sl[j], sl[i]
	}
}

func ConvertSliceToMap(data interface{}) map[string]string {
	res := make(map[string]string)
	if data == nil {
		return res
	}
	ref := reflect.ValueOf(data)
	if ref.Kind() != reflect.Slice {
		return res
	}

	for i := 0; i < ref.Len(); i++ {
		obj := ref.Index(i)
		typeOfData := obj.Type()
		var key, value string
		for j := 0; j < obj.NumField(); j++ {
			name := typeOfData.Field(j).Name
			if val, ok := obj.Field(j).Interface().(string); ok {
				if name == "Key" {
					key = val
				} else if name == "Value" {
					value = val
				}
			}
			res[key] = value
		}
	}
	return res
}

func SaveFileToLocal(file multipart.File, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}
	return nil
}

func ReadFileLocal(path string) (*os.File, error) {
	fileRes, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return fileRes, nil
}

func RemoveFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func DecodeImage(bs64, dest string) error {
	unbased, err := base64.StdEncoding.DecodeString(bs64)
	if err != nil {
		return err
	}
	r := bytes.NewReader(unbased)
	im, err := jpeg.Decode(r)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	png.Encode(f, im)
	return nil
}

func RemoveEmptyString(src []string) (res []string) {
	for _, s := range src {
		if !IsNilObject(s) {
			res = append(res, s)
		}
	}
	return
}
