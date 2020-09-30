package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
	有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
	例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。
	示例 1：

	输入：s = "25525511135"
	输出：["255.255.11.135","255.255.111.35"]
	示例 2：

	输入：s = "0000"
	输出：["0.0.0.0"]
	示例 3：

	输入：s = "1111"
	输出：["1.1.1.1"]
	示例 4：

	输入：s = "010010"
	输出：["0.10.0.10","0.100.1.0"]
	示例 5：

	输入：s = "101023"
	输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
	 
	提示：
	0 <= s.length <= 3000
	s 仅由数字组成

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/restore-ip-addresses
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func restoreIpAddresses(s string) []string {

	var res []string
	var list []string
	resHelper(s, list, &res, len(s))

	return res
}

func resHelper(s string, list []string, res *[]string, gs int){

	if len(list) == 4 && len(list[0])+len(list[1])+len(list[2])+len(list[3])==gs{
		ans := make([]string, len(list))
		copy(ans, list)
		*res = append(*res, strings.Join(ans[:], "."))
	}

	for i:=1; i<=len(s); i++{

		// 如果到了4层，返回
		if len(list)>4{return}

		r, err := strconv.Atoi(s[0:i])
		if err != nil {
			fmt.Println(err)
		}
		if (len(s[0:i])>=2 && string(s[0:i][0])=="0") || r>255{return}

		list = append(list, string(s[0:i]))
		resHelper(s[i:], list, res, gs)
		list = list[0:len(list)-1]
	}
}

func main(){
	fmt.Println(restoreIpAddresses("101023"))
}
