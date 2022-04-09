package utils

// GetValue 获取值
func GetValue(line []float64, key int) float64 {
	return line[key]
}

// GetFirst 获取第一个值
func GetFirst(line []float64) float64 {
	return line[0]
}

// GetLast 获取最后一个值
func GetLast(line []float64) float64 {
	return line[len(line)-1]
}
