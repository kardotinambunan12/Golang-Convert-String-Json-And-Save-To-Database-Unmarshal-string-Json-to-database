package controllers

import (
	"time"
	"crypto/md5"
	"strconv"
	"encoding/hex"
	"github.com/labstack/echo"	
	logger    "../../customlogger"
)

// ## Define Config Variable Global
var logs 		  			= logger.GetInstance("SYSTEMS -")
// ## End

// ## Define Type Global
type response_json map[string]interface{}
type DB struct {
    Value        interface{}
    Error        error
    RowsAffected int64
}
func AppIndex(c echo.Context) error{
	return c.JSON(200, "High performance, minimalist Go web framework running")
}

func ConvertToMD5(value string) string{
	var str string = value
	hasher := md5.New()
	hasher.Write([]byte(str))
	converId := hex.EncodeToString(hasher.Sum(nil))

	return converId
}

func ConvertStringToInt(value string) int{
	value_int, _  	:= strconv.Atoi(value)
	return value_int
}

func ConvertStringToFloat(value string) float64{
	value_float, _ 	:= strconv.ParseFloat(value, 8)
	return value_float
}

func current_time(format string) string{
	t 			 := time.Now()
	current_time := t.Format(format)
	return current_time
}
// #