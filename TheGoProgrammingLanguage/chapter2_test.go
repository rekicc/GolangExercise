package theGoProgrammingLanguage

import (
	"flag"
	"testing"
)

var tempFlag = flag.String("t", " ", "指定温度单位，C为摄氏度，F为华氏度，K为开尔文温度")
var lengFlag = flag.String("l", " ", "指定长度单位，m为米，in为英寸，ft为英尺，mi为英里，nmi为海里")

//测试2.2
func TestUnitConv(t *testing.T) {
	unitConv(tempFlag, lengFlag)
}
