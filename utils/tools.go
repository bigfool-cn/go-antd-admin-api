package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

// 获取当前时间
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// MD5
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 求交集
func Intersect(slice1 interface{}, slice2 interface{}) (interface{}, error){
	switch slice := slice1.(type) {
	case []string:
		base := make(map[string]int)
		result := make([]string, 0)
		for _, v := range slice1.([]string) {
			base[v]++
		}

		for _, v := range slice2.([]string) {
			times, _ := base[v]
			if times == 1 {
				result = append(result, v)
			}
		}
		return result, nil
	case []int64:
		base := make(map[int64]int)
		result := make([]int64, 0)
		for _, v := range slice1.([]int64) {
			base[v]++
		}
		for _, v := range slice2.([]int64) {
			times, _ := base[v]
			if times == 1 {
				result = append(result, v)
			}
		}
		return result, nil
	default:
		err := fmt.Errorf("Unknown type: %T", slice)
		return nil, err
	}
}

// 求差集
func Difference(slice1 interface{}, slice2 interface{}) (interface{}, error) {
	switch slice := slice1.(type) {
	case []string:
		base := make(map[string]int)
		result := make([]string, 0)
		inter, err := Intersect(slice1, slice2)
		if err != nil {
			return nil, err
		}
		for _, v := range inter.([]string) {
			base[v]++
		}

		for _, value := range slice1.([]string) {
			times, _ := base[value]
			if times == 0 {
				result = append(result, value)
			}
		}
		return result, nil
	case []int64:
		base := make(map[int64]int)
		result := make([]int64, 0)
		inter, err := Intersect(slice1, slice2)
		if err != nil {
			return nil, err
		}
		for _, v := range inter.([]int64) {
			base[v]++
		}

		for _, value := range slice1.([]int64) {
			times, _ := base[value]
			if times == 0 {
				result = append(result, value)
			}
		}
		return result, nil
	default:
		err := fmt.Errorf("Unknown type: %T", slice)
		return nil, err
	}
}
