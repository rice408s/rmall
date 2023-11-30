package utils

import "strconv"

// ConvertStringArrayToIntArray 将字符串数组转换为整数数组
func ConvertStringArrayToIntArray(strArray []string) ([]int, error) {
	var intArray []int

	for _, str := range strArray {
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, err // 如果有一个字符串无法成功转换为整数，返回错误
		}
		intArray = append(intArray, i)
	}

	return intArray, nil
}
