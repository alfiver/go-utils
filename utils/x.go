package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const digitBytes = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}
func Contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// 将数组平均分成N份
func SplitSlice[T any](sl []T, num int) [][]T {
	slLen := len(sl)
	if slLen < 1 {
		return nil
	}
	if num > slLen {
		num = slLen
	} else if num < 1 {
		num = 1
	}
	step := slLen / num
	if slLen%num > 0 {
		step += 1
	}
	newSl := make([][]T, step)
	for i := 0; i < step; i++ {
		end := (i + 1) * num
		if end > slLen {
			newSl[i] = sl[i*num:]
		} else {
			newSl[i] = sl[i*num : end]
		}
	}
	return newSl
}
func RandomString(n int) string {
	if n < 1 {
		return ""
	}
	b := make([]byte, n)
	l := len(letterBytes)
	for i := range b {
		b[i] = letterBytes[rand.Intn(l)]
	}
	return string(b)
}
func RandomDigits(n int) string {
	if n < 1 {
		return ""
	}
	b := make([]byte, n)
	l := len(digitBytes)
	for i := range b {
		b[i] = digitBytes[rand.Intn(l)]
	}
	return string(b)
}

// 计算小数点精度 1.2323 返回4，  1.0 返回1,  33 返回0
func CalcPrecision(v string) int {
	if !strings.Contains(v, ".") {
		return 0
	}
	arr := strings.Split(v, ".")
	return len(arr[1])
}
func SysUptime(now, launchTime int64) string {
	sec := now - launchTime
	uptime := make([]string, 0, 3)
	var day, hour, min int64
	if sec >= 86400 {
		day = sec / 86400
		if day > 0 {
			uptime = append(uptime, fmt.Sprintf("%dd", day))
		}
		sec = sec % 86400
	}
	if sec >= 3600 {
		hour = sec / 3600
		if hour > 0 {
			uptime = append(uptime, fmt.Sprintf("%dh", hour))
		}
		sec = sec % 3600
	}
	if sec >= 60 {
		min = sec / 60
		if min > 0 {
			uptime = append(uptime, fmt.Sprintf("%dm", min))
		}
		sec = sec % 60
	}
	if sec > 0 {
		uptime = append(uptime, fmt.Sprintf("%ds", sec))
	}
	return strings.Join(uptime, ".")
}
