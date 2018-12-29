package main

import (
	"fmt"
	"github.com/gogo/protobuf/io"
)

// *celsiusFlag 满足 flag.Value 接口
type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)// 无须检查错误
	switch unit {
	case "C", "`C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "`F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
	io.Writer()
}