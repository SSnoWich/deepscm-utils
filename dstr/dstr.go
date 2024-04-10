// 描述:
// 作者: SNOWICH
// 日期: 2024/4/8 00:03
// 版权: 山东深链智能科技有限公司 @Since 2024

package dstr

import (
	"bytes"
	"strings"
)

// Splicing 拼接字符串
func Splicing(str ...string) string {
	switch len(str) {
	case 0:
		return ""
	case 1:
		return str[0]
	default:
		var b bytes.Buffer
		for _, v := range str {
			b.WriteString(v)
		}
		return b.String()
	}
}

// StringJoin 拼接字符串带间隔符
func StringJoin(sep string, str ...string) string {
	switch len(str) {
	case 0:
		return ""
	case 1:
		return str[0]
	default:
		return strings.Join(str, sep)
	}

}
