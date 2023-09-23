package theGoProgrammingLanguage

import (
	"flag"
	"fmt"
	"runtime"
	"testing"
)

var tempFlag = flag.String("t", " ", "指定温度单位，C为摄氏度，F为华氏度，K为开尔文温度")
var lengFlag = flag.String("l", " ", "指定长度单位，m为米，in为英寸，ft为英尺，mi为英里，nmi为海里")

// 测试2.2
func TestUnitConv(t *testing.T) {
	unitConv(tempFlag, lengFlag)
}
func BenchmarkPopCount1(b *testing.B) {
	num := uint64(9889348999)
	for i := 0; i < b.N; i++ {
		_ = popCount1(num)
	}
}
func BenchmarkPopCount2(b *testing.B) {
	num := uint64(9889348999)
	for i := 0; i < b.N; i++ {
		_ = popCount2(num)
	}
}
func BenchmarkPopCount3(b *testing.B) {
	runtime.GOMAXPROCS(8)
	num := uint64(9889348999)
	for i := 0; i < b.N; i++ {
		_ = popCount3(num)
	}
}
func BenchmarkPopCount4(b *testing.B) {
	num := uint64(9889348999)
	for i := 0; i < b.N; i++ {
		_ = popCount4(num)
	}
}

func TestPopCount1(t *testing.T) {
	var i uint64 = 9889348999
	fmt.Println(popCount1(i))
}
func TestPopCount2(t *testing.T) {
	var i uint64 = 9889348999
	fmt.Println(popCount2(i))
}
func TestPopCount3(t *testing.T) {
	var i uint64 = 9889348999
	fmt.Println(popCount3(i))
}
func TestPopCount4(t *testing.T) {
	var i uint64 = 9889348999
	fmt.Println(popCount4(i))
}
