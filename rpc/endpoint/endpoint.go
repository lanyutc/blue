package endpoint

import (
	"strings"
)

type Endpoint struct {
	Set      string
	Name     string
	Addr     string
	IsActive bool
}

//Set匹配规则：
//Set被"."分为三个部分，可以理解成SET[名字.区域.组]，必须由数字或者字母组成，唯一允许特殊字符"*"，表示任意匹配
//我们按照从[名字.区域.组]递进匹配的规则，若主调和被调有一方为"*"则匹配，否则他们必须相等才算匹配
//例如：1.1.1->1.1.1 match
//      1.1.1->1.1.2 mismatch
//      1.1.*->1.1.2 match
//      1.2.*->1.1.2 mismatch
//      1.1.9->1.1.* match
func (e Endpoint) IsSetMatch(set string) bool {
	setA := strings.Split(e.Set, ".") //被调用
	setB := strings.Split(set, ".")
	if len(setA) == 3 && len(setB) == 3 {
		if setA[0] == "*" || setB[0] == "*" || (setA[0] != "*" && setA[0] == setB[0]) {
			if setA[1] == "*" || setB[1] == "*" || (setA[1] != "*" && setA[1] == setB[1]) {
				if setA[2] == "*" || setB[2] == "*" || (setA[2] != "*" && setA[2] == setB[2]) {
					return true
				}
			}
		}
	}

	return false
}
