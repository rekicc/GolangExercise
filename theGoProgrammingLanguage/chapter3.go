// 练习3.1, 3.2, 3.3, 3.6, 3.7, 3.8因无法理解题目以及对应的算法没有实现
package theGoProgrammingLanguage

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
	"strings"
)

// 练习3.4
const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

// 因为这些量会在后续被http请求传来的参数更改, 所以将它们声明为变量
var (
	sin30, cos30  = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
	width, height = 600., 320.
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
)

func surfacePlot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		if k == "height" {
			res, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Println(err)
			}
			if res > 0 {
				height = res
			}
		}
		if k == "width" {
			res, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Println(err)
			}
			if res > 0 {
				width = res
			}
		}
	}
	xyscale = width / 2 / xyrange
	zscale = height * 0.4

	var b bytes.Buffer
	fmt.Fprintf(&b, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", int64(width), int64(height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(&b, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(&b, "</svg>")
	if _, err := io.Copy(w, &b); err != nil {
		fmt.Println(err)
	}
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// 练习3.9, 遗留问题是没有实现zoom, 不知道zoom是干嘛的
func mandelbrotServe(w http.ResponseWriter, r *http.Request) {
	var (
		xmin, ymin, xmax, ymax float64 = -2, -2, +2, +2
		width, height          float64 = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		if k == "x" {
			res, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Println(err)
			}
			xmin = -res
			xmax = +res
		}
		if k == "y" {
			res, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Println(err)
			}
			ymin = -res
			ymax = +res
		}
	}
	for py := 0; py < int(height); py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < int(width); px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	err := png.Encode(w, img)
	if err != nil {
		log.Println(err)
	}
}

// 练习3.5
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			n := 255 - contrast*n
			//返回彩色
			return color.YCbCr{n, 250, 240}
		}
	}
	return color.Black
}

// 分离正负号, 形参s是原始字符串; 第一个返回值是分离完符号后剩余字符串, 第二个返回值是分离出来的符号, 如果有的话
func splitSign(s string) (string, string) {
	var sy string
	if s[0] == '+' || s[0] == '-' {
		sy = string(s[0])
		return string(s[1:]), sy
	}
	return s, ""
}
func commaInt1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	s = commaInt1(s[:n-3]) + "," + s[n-3:]
	return s
}

// 练习3.10, 但是功能不完善，只能处理特定数值字符串，无法处理超出int64表示范围的数值
func commaInt(s string) string {
	w, _ := strconv.Atoi(s)
	var m string
	if w <= 999 {
		return strconv.Itoa(w)
	}
	for {
		x := w % 1000
		m = "," + strconv.Itoa(x) + m
		w = w / 1000
		if w < 1000 {
			m = strconv.Itoa(w%1000) + m
			break
		}
	}
	return m
}

// 练习3.10, 使用bytes.Buffer以及循环来处理整数值
func commaInt2(s string) string {
	var temp bytes.Buffer
	var x, count int = len(s), 0
	//s整体长度正好是3n
	if (x % 3) == 0 {
		for n, i := range s {
			// n==0代表的是s的第一位
			// n%3 != 0 代表s中非3的倍数的位置,
			// count用于统计循环了多少次, 当count == x/3-1时, 代表已经到了s最后的三位
			if n != 0 && n%3 == 0 && count != x/3-1 {
				fmt.Fprintf(&temp, ",%s", string(i))
				count++
				continue
			}
			fmt.Fprintf(&temp, "%s", string(i))
		}
	}
	//s整体长度是3n+1
	if (x % 3) == 1 {
		for n, i := range s {
			//因为s长度是3n+1, 所以n==1时需要加comma
			if count != x/3 && n == 1 || (n-1)%3 == 0 {
				fmt.Fprintf(&temp, ",%s", string(i))
				count++
				continue
			}
			fmt.Fprintf(&temp, "%s", string(i))
		}
	}
	//s整体长度是3n+2
	if (x % 3) == 2 {
		for n, i := range s {
			if count != x/3 && n == 2 || (n-2)%3 == 0 {
				fmt.Fprintf(&temp, ",%s", string(i))
				count++
				continue
			}
			fmt.Fprintf(&temp, "%s", string(i))
		}
	}
	return temp.String()
}

// 练习3.10 从网上获得的灵感，依次遍历字符串，下标模3的结果与字符串长度模3的结果相等时，该下标前面就是要插入逗号的位置，唯一的特例是当字符串长度模3等于0时要跳过下标0
func commaInt3(s string) string {
	var temp bytes.Buffer
	for n, i := range s {
		if len(s)%3 != 0 && n%3 == len(s)%3 {
			fmt.Fprintf(&temp, ",%s", string(i))
			continue
		}
		fmt.Fprintf(&temp, "%s", string(i))
	}
	return temp.String()
}

// 练习3.11, 为浮点数添加comma
func commaFloat(s string) string {
	var inte, deci string
	/*	for n, i := range s {
		if i == '.' {
			inte = s[:n]
			deci = s[n+1:]
		}
	}*/
	//不再使用遍历，而是直接用strings.Split来分割整数和小数部分
	inte = strings.Split(s, ".")[0]
	deci = strings.Split(s, ".")[1]

	inte = commaInt3(inte)
	return (inte + "." + deci)

	// 小数部分无千位分隔符
	/*	var temp bytes.Buffer
		for n, i := range deci {
			if n != 0 && n%3 == 0 {
				fmt.Fprintf(&temp, ",%s", string(i))
				continue
			}
			fmt.Fprintf(&temp, "%s", string(i))
		}
		return (inte + "." + temp.String())
	*/
}

// 练习3.12, 判断两个string是否是同字母异序词
func compareString(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var tempMapA, tempMapB = make(map[string]int, len(a)), make(map[string]int, len(a))
	for _, i := range a {
		tempMapA[string(i)]++
	}

	for _, i := range b {
		tempMapB[string(i)]++
	}

	for k, r := range tempMapA {
		if tempMapB[k] != r {
			return false
		}
	}
	return true
}

func str() {
	x := -123
	fmt.Println(strconv.FormatUint(uint64(x), 2))
	y := "123888"
	fmt.Println(strconv.ParseInt(y, 10, 8))
}

// 练习3.13
const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)
