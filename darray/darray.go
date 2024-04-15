// 描述:
// 作者: SNOWICH
// 日期: 2024/4/15 23:49
// 版权: 山东深链智能科技有限公司 @Since 2024

package darray

import "github.com/gogf/gf/v2/database/gdb"

func GormArrayToIds(array []gdb.Value) []uint {
	ids := make([]uint, 0)
	for _, v := range array {
		ids = append(ids, v.Uint())
	}
	return ids
}

// UintReverse 数组倒序
func UintReverse(arr *[]uint) {
	var temp uint
	length := len(*arr)
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}
