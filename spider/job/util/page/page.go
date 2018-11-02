package page

import "math"

//获取页码
func CalculatePage(totalCount, pageSize float64) int {
	f := float64(totalCount) / float64(pageSize)
	//总条数除以每页大小,结果向上取整数,获取页数
	return int(math.Ceil(f))

}
