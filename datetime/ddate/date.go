// ============================================================
// 描述:
// 作者: SnoWich
// 日期: 2023/3/29 16:10
// 版权: 山东深链智能科技有限公司 @Since 2023
// ============================================================

package ddate

/*
输入一个月份 来计算天数
GetDaysByYearMouth 查询指定年份指定月份有多少天
@params year int 制定年份
@params month int 制定月份
*/

func GetDaysByYearMouth(year int, month int) int {
	// 有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	if day31[month] == true {
		return 31
	}
	// 有30天的月份

	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] == true {
		return 30
	}
	// 计算平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出二月天数
		return 29
	}
	// 得出平年二月天数
	return 28
}
