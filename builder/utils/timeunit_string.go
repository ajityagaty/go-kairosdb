// Code generated by "stringer -type=TimeUnit"; DO NOT EDIT

package utils

import "fmt"

const _TimeUnit_name = "MILLISECONDSSECONDSMINUTESHOURSDAYSWEEKSMONTHSYEARS"

var _TimeUnit_index = [...]uint8{0, 12, 19, 26, 31, 35, 40, 46, 51}

func (i TimeUnit) String() string {
	if i < 0 || i >= TimeUnit(len(_TimeUnit_index)-1) {
		return fmt.Sprintf("TimeUnit(%d)", i)
	}
	return _TimeUnit_name[_TimeUnit_index[i]:_TimeUnit_index[i+1]]
}
