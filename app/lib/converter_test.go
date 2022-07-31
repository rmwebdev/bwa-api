package lib

import (
	"testing"

	"github.com/gofiber/fiber/v2/utils"
)

func TestConvertToMD5(t *testing.T) {
	value := 1
	ConvertToMD5(&value)
}

func TestConvertStrToMD5(t *testing.T) {
	value := "development usage"
	gen := ConvertStrToMD5(&value)
	gen2 := ConvertStrToMD5(&value)
	utils.AssertEqual(t, gen2, gen)
}

func TestConvertToSHA1(t *testing.T) {
	value := "development usage"
	ConvertToSHA1(value)
}

func TestConvertToSHA256(t *testing.T) {
	value := "development usage"
	ConvertToSHA256(value)
}

func TestIntToStr(t *testing.T) {
	value := 1
	res := IntToStr(value)
	utils.AssertEqual(t, "1", res)
}

func TestStrToInt(t *testing.T) {
	value := "1"
	res := StrToInt(value)
	utils.AssertEqual(t, 1, res)
}

func TestStrToInt64(t *testing.T) {
	value := "1"
	res := StrToInt64(value)
	utils.AssertEqual(t, int64(1), res)
}

func TestStrToFloat(t *testing.T) {
	value := "1"
	res := StrToFloat(value)
	utils.AssertEqual(t, float64(1), res)
}

func TestStrToBool(t *testing.T) {
	value := "true"
	res := StrToBool(value)
	utils.AssertEqual(t, true, res)
}

func TestFloatToStr(t *testing.T) {
	value := 1.2
	res := FloatToStr(value)
	utils.AssertEqual(t, "1.200000", res)
}

func TestConvertJsonToStr(t *testing.T) {
	value := []interface{}{"first", "second"}
	res := ConvertJsonToStr(value)
	utils.AssertEqual(t, `["first","second"]`, res)
}

func TestConvertStrToObj(t *testing.T) {
	value := `{"index":"value"}`
	res := ConvertStrToObj(value)
	utils.AssertEqual(t, "value", res["index"])
}

func TestConvertStrToJson(t *testing.T) {
	expect := map[string]interface{}{
		"index": "value",
	}
	value := `{"index":"value"}`
	res := ConvertStrToJson(value)
	utils.AssertEqual(t, expect, res)
}

func TestConvertStrToTime(t *testing.T) {
	value := "2021-05-19 11:56:30"
	gen := ConvertStrToTime(value)
	utils.AssertEqual(t, gen, gen)
}
