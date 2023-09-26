package theGoProgrammingLanguage

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

// 测试练习3.4
func TestSurfacePolt(t *testing.T) {
	req := httptest.NewRequest("GET", "http://127.0.0.1/?height=1000&width=100", nil)
	w := httptest.NewRecorder()

	surfacePlot(w, req)
}

func TestMandelbrotServe(t *testing.T) {
	req := httptest.NewRequest("GET", "http://127.0.0.1/?x=7&y=5", nil)
	w := httptest.NewRecorder()

	mandelbrotServe(w, req)

	fmt.Print(w.Body)
}

func TestComma(t *testing.T) {
	s := "+12344556990090009839483498349384934893489348394"
	/*	if strings.HasPrefix(s, "+") {
			s = s[1:]
			s = commaInt2(s)
			s = "+" + s
		} else if strings.HasPrefix(s, "-") {
			s = s[1:]
			s = commaInt2(s)
			s = "-" + s
		}*/

	s, sy := splitSign(s)
	fmt.Println("commaInt1")
	if sy == "" {
		fmt.Println(commaInt1(s))
	} else {
		fmt.Println(sy + commaInt1(s))
	}

	fmt.Println("commaInt2:")
	if sy == "" {
		fmt.Println(commaInt2(s))
	}else{
		fmt.Println(sy + commaInt2(s))
	}
		
	fmt.Println("commaInt3:")
	if sy == "" {
		fmt.Println(commaInt2(s))
	}else{
		fmt.Println(sy + commaInt2(s))
	}

	w := "1232341234985080.20392039837043710509"

	fmt.Println("commaFloat:")
	r, sy := splitSign(w)
	if sy == "" {
		fmt.Println(commaFloat(r))
	}else{
		fmt.Println(sy + commaFloat(r))
	}

	fmt.Println("compareString:")
	f := "wasol"
	g := "awslo"
	fmt.Println(compareString(f, g))

	m := "11123232039203920392039209129888484884848"
	fmt.Println(commaInt3(m))
	fmt.Println(commaInt1(m))
}

func TestStr(t *testing.T){
	str()
}
