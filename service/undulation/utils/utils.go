package utils

import (
	"fmt"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"math"
	"strconv"
)

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

func RemoveDuplicates(nums []float64) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	// 1  2  2  4  5  5  6
	count := 0               // 统计重复出现的次数
	for i := 1; i < n; i++ { // 循环整个数组
		if nums[i] == nums[i-1] { //数组的前与后的元素进行对比.如果相等则,记录count
			count++
			continue
		}
		// 如果不相等则交换位置,注意这里是使用i-count,count是记录出现重复的次数.
		nums[i-count] = nums[i]

	}
	return n - count //返回整个数组长度减去重复的count次数,则等于不相等的数组长度.
}

// GetUpDown 获取数据的起始涨跌位置
// 1 1 1 2 2 2 3 3 3
func getUpDown(nums []float64) (int, string) {
	n := len(nums)
	for i := 1; i < n; i++ { // 循环整个数组
		if nums[i] == nums[i-1] { // 数组的前与后的元素进行对比.如果相等则,跳出循环
			continue
		}
		// 直到找到下一个进行对比
		if nums[i] > nums[i-1] { // 获取第一个顶部位置
			return i, "top"
		} else {
			return i, "down"
		}
	}
	return 0, ""
}

type Dstruct struct {
	Index int
	Val   float64
}

// GetVarianceData 计算方差数据，返回
func GetVarianceData(nums []float64) (float64, float64) {
	var data = "c_close\n42007.4525\n42007.4525\n42007.4525\n42007.6675\n42003.7425\n42003.72\n42003.555\n42002.8225\n42002.8225\n42005.9075\n42005.9075\n42005.9075\n42011.4325\n42011.4325\n42011.4325\n42011.43\n42010.5825\n42008.9475\n42008.9525\n42009.27\n42009.27\n42009.27\n42008.95\n42008.95\n42008.95\n42010.4475\n42010.435\n42008.9875\n42009.2175\n42009.225\n42009.255\n42009.045\n42010.165\n42016.6525\n42016.6\n42015.855\n42016.6\n42015.97\n42012.3825\n42007.815\n42004.455\n42003.705\n42003.705\n41998.355\n41998.3475\n41997.06\n41997.06\n41996.8325\n41997.0225\n41998.6175\n41998.2975\n41998.2975\n41998.63\n41998.18\n41999.09\n41999.0175\n41998.51\n41998.51\n41999.0925\n41999.5475\n41999.875\n41999.875\n42002.2625\n42002.2625\n42002.2625\n42003.8525\n42002.99\n42002.9875\n42002.2225\n42002.2225\n41998.0575\n42000.905\n42000.6\n42001.095\n42001.6175\n42002.6075\n42003.79\n42003.79\n42003.9625\n42004.6175\n42003.11\n42003.11\n42003.105\n42002.885\n42003.6075\n42003.6075\n42003.6075\n42003.6075\n42003.7725\n42002.0275\n42001.32\n42001.7975\n41997.72\n41996.6325\n41996.5675\n42000.0325\n41997.0575\n41996.3025\n41996.3025\n41996.3025\n41996.3025\n41996.3025\n41996.2975\n41995.5375\n41995.5375\n41996.74\n41996.0375\n41996.5575\n41997.5875\n41997.2\n41998.1025\n41998.1025\n42000.32\n42000.225\n41999.97\n42000.0675\n42000.0525\n41997.71\n41997.0025\n41997.0925\n42001.7725\n42001.7725\n42000.8775\n42000.7675\n41999.7025\n42000.7225\n42000.7225\n42000.345\n42001.91\n42001.91\n42002.3925\n42002.3925\n42000.195\n42003.995\n42004.6625\n42004.6625\n42004.6625\n42004.0725\n42003.8175\n42003.3875\n42002.2425\n42002.815\n42003.4925\n42003.935\n42003.935\n42003.64\n42003.0125\n42003.0625\n42002.0225\n42002.9175\n42002.115\n41999.16\n41999.7525\n41999.7525\n41999.4925\n42000.3\n42000.0925\n42000.0125\n42000.0125\n42000.24\n42001.1\n42001.1\n42000.0375\n42000.0075\n42000.015\n42000.3525\n41999.8175\n41999.82\n41999.89\n41999.1775\n42000.5525\n42000.9275\n42000.9275\n42000.8375\n42000.8375\n42001.03\n41999.88\n41999.4525\n41999.4525\n41999.4475\n41999.3975\n41999.3025\n41999.785\n42000.175\n42000.175\n42000.4425\n42000.4425\n41999.5325\n41999.775\n42000.07\n42000.07\n41999.525\n41997.835\n41993.83\n41993.83\n41995.2375\n41994.7375\n41998.32\n41998.32\n41998.32\n41998.32\n41998.32\n41998.32\n42006.9975\n42010.0925\n42022.3525\n42025.6525\n42031.365\n42032.06\n42031.835\n42031.8375\n42032.7175\n42033.045\n42032.3175\n42032.3175\n42031.03\n42031.03\n42031.3925\n42030.8325\n42031.8125\n42031.73\n42032.02\n42032.02\n42032.8375\n42032.8375\n42032.695\n42032.8575\n42032.8575\n42033.47\n42036.4675\n42039.3925\n42040.3925\n42040.565\n42040.4425\n42039.575\n42040.775\n42040.745\n42040.7025\n42042.4975\n42042.9\n42043.7175\n42043.76\n42043.76\n42041.1925\n42039.1775\n42040.0575\n42040.0575\n42036.7575\n42038.3625\n42038.3625\n42038.2825\n42038.2825\n42038.7875\n42038.7875\n42038.7225\n42037.7275\n42034.5875\n42034.5875\n42035.055\n42035.055\n42031.5425\n42031.5425\n42032.415\n42032.415\n42032.6925\n42034.3925\n42035.33\n42036.915\n42036.915\n42037.6375\n42037.61\n42037.61\n42036.55\n42035.3825\n42036.3925\n42036.3925\n42034.2225\n42034.2225\n42034.885\n42034.835\n42032.9\n42032.9\n42032.26\n42032.26\n42032.195\n42032.195\n42029.78\n42029.7675\n42029.64\n42029.01\n42021.255\n42021.255\n42021.15\n42021.15\n42021.4875\n42021.4875\n"
	nums = helpers.CloseDataConvert(data)
	sum := getSum(nums)
	logger.Info(fmt.Sprint("sum=", sum))
	// 计算平均数
	avg := sum / float64(len(nums))
	var sqdArr []float64
	for i := 0; i < len(nums); i++ {
		// 平方差
		sqdArr = append(sqdArr, math.Pow(nums[i]-avg, 2))
	}
	// 平方差之和
	sqdSum := getSum(sqdArr)
	logger.Info(fmt.Sprint("sqdSum=", sqdSum))
	// 平方差的平均数
	sqdAvg := sqdSum / float64(len(nums))
	logger.Info(fmt.Sprint("sqdAvg=", sqdAvg))
	// 平方差
	Sqrt := math.Sqrt(sqdAvg)
	// 平方差
	logger.Info(fmt.Sprint("平均值=", avg))
	logger.Info(fmt.Sprint("平方差=", Sqrt))
	return avg, Sqrt
}

// 求和
func getSum(data []float64) float64 {
	s := 0.0
	l := len(data)
	for i := 0; i < l; i++ {
		s += data[i]
	}
	return s
}

// 计算波动率

func GetWaveRate(arr []Dstruct) {
	var totalWave float64 = 0
	// 计算下降趋势波动率
	if len(arr) > 1 {
		var waveRateArr []float64
		for i := 1; i < len(arr); i++ {
			// 时间间隔
			interval := arr[i].Index - arr[i-1].Index
			// 顶部/底部 的差值
			wave := arr[i].Val - arr[i-1].Val
			waveRate := Decimal(wave / float64(interval))
			// 收集波动率
			waveRateArr = append(waveRateArr, waveRate)
			// 波动率之和
			totalWave = totalWave + waveRate
		}
		// 平均波动率
		avg := totalWave / float64(len(arr)-1)
		fmt.Println("平均波动率", avg)
		fmt.Println("波动值之和", totalWave)
		fmt.Println("波动率数组", waveRateArr)
		// 计算每次波动的值于平均波动值对比，大于1的认为为大波动，
		var bigWaveCount = 0
		for i := 0; i < len(waveRateArr); i++ {
			if math.Abs(math.Abs(waveRateArr[i])-avg) > 1 {
				bigWaveCount++
			}
		}
		// 大波动占比
		bigWaveRate := float64(bigWaveCount) / float64(len(arr)-1)
		// 大波动占比55%认为波动大
		if bigWaveRate > 0.55 {

		}
		fmt.Println(bigWaveCount, len(arr)-1, bigWaveRate)
	}
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value
}
