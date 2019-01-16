package utils

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

var (
	timeLayout  = "2006-01-02 15:04:05"
	baseLayout  = "2018-01-01 00:00:00"
	baseTime, _ = time.Parse(timeLayout, baseLayout)
)

func setField(obj interface{}, name string, value interface{}) error {
	structData := reflect.ValueOf(obj).Elem()
	fieldValue := structData.FieldByName(name)

	if !fieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj ", name)
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value ", name)
	}

	fieldType := fieldValue.Type()
	val := reflect.ValueOf(value)

	valTypeStr := val.Type().String()
	fieldTypeStr := fieldType.String()
	if valTypeStr == "float64" && fieldTypeStr == "int" {
		val = val.Convert(fieldType)
	} else if fieldType != val.Type() {
		return errors.New("Provided value type " + valTypeStr + " didn't match obj field type " + fieldTypeStr)
	}
	fieldValue.Set(val)
	return nil
}

// SetStructByJSON 由json对象生成 struct
func SetStructByJSON(obj interface{}, mapData map[string]interface{}) error {
	for key, value := range mapData {
		if err := setField(obj, key, value); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

//ParseToUTC 将时间转换为utc时间
func ParseToUTC(timeString string) (time.Time, error) {
	var res time.Time
	var err error
	timeLayout := "2006-01-02T15:04:05Z07:00"
	if timeString == "" {
		return res, errors.New("时间不能为空")
	}
	res, err = time.Parse(timeLayout, timeString)
	if err != nil {
		return res, err
	}
	return res.UTC(), nil
}

//MakeUTCToIndex 将utc时间转为时间片
func MakeUTCToIndex(timeUTC time.Time) int {
	minutes := timeUTC.Sub(baseTime).Minutes()
	index := int(minutes)/5 + 1
	return index
}

//MakeIndexToUTC 将时间片转为utc时间
func MakeIndexToUTC(index int) time.Time {
	minutes := (index - 1) * 5
	utcTime := baseTime.Add(time.Minute * time.Duration(minutes))
	return utcTime
}

//MakeUTCToString 将时间转为字符串
func MakeUTCToString(utcTime time.Time) string {
	return utcTime.Format(timeLayout)
}

//CompareIntSlice 比较数组
func CompareIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
