package theGoProgrammingLanguage

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type Inch float64
type Foot float64
type Mile float64
type Nauticalmile float64
type Metre float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// 练习2.1, 开尔文温度, 华氏温度, 摄氏温度三者互相转换
func (f Fahrenheit) FToC() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (f Fahrenheit) FToK() Kelvin {
	return Kelvin((f-32)*5/9 + 273.5)
}

func (c Celsius) CToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (c Celsius) CToK() Kelvin {
	return Kelvin(c + 273.15)
}

func (k Kelvin) KToC() Celsius {
	return Celsius(k - 273.15)
}

func (k Kelvin) KToF() Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32)
}

// 三种温度各自的String()方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

// 练习2.2, 英尺、英寸、英里、海里、米之间的转换与各自的String()方法
func (i Inch) String() string {
	return fmt.Sprintf("%gin", i)
}

func (f Foot) String() string {
	return fmt.Sprintf("%gft", f)
}

func (m Mile) String() string {
	return fmt.Sprintf("%gmi", m)
}

func (n Nauticalmile) String() string {
	return fmt.Sprintf("%gnmi", n)
}

func (m Metre) String() string {
	return fmt.Sprintf("%gm", m)
}

func InchToMetre(i Inch) Metre {
	return Metre(i * 0.0254)
}

func MetreToInch(m Metre) Inch {
	return Inch(m / 0.0254)
}

func FootToMetre(f Foot) Metre {
	return Metre(f * 0.3048)
}

func MetreToFoot(m Metre) Foot {
	return Foot(m / 0.3048)
}

func MileToMetre(m Mile) Metre {
	return Metre(m * 1609.344)
}

func MetreToMile(m Metre) Mile {
	return Mile(m / 1609.344)
}

func NauticalmileToMetre(n Nauticalmile) Metre {
	return Metre(n * 1852)
}

func MetreToNauticalmile(m Metre) Nauticalmile {
	return Nauticalmile(m / 1852)
}

func unitConv(tempFlag, lengFlag *string) {
	fmt.Println(*tempFlag, *lengFlag)
	switch {
	case *tempFlag == "C":
		for _, arg := range flag.Args() {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert false: %v\n", err)
			}
			fmt.Printf("%s = %s = %s\n", Celsius(num), Celsius(num).CToF(), Celsius(num).CToK())
		}
	case *tempFlag == "F":
		for _, arg := range flag.Args() {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert false: %v\n", err)
			}
			fmt.Printf("%s = %s = %s\n", Fahrenheit(num), Fahrenheit(num).FToC(), Fahrenheit(num).FToK())
		}
	case *tempFlag == "K":
		for _, arg := range flag.Args() {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert false: %v\n", err)
			}
			fmt.Printf("%s = %s = %s\n", Kelvin(num), Kelvin(num).KToC(), Kelvin(num).KToF())
		}
	case *lengFlag == "m":
		for _, arg := range flag.Args() {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert false: %v\n", err)
			}
			fmt.Printf("%s = %s = %s\n", Metre(num), MetreToInch(Metre(num)), MetreToFoot(Metre(num)))
		}
	case *lengFlag == "in":
		for _, arg := range flag.Args() {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert false: %v\n", err)
			}
			fmt.Printf("%s = %s\n", Inch(num), InchToMetre(Inch(num)))
		}
	case *lengFlag == "mi":
		for _, arg := range flag.Args() {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert false: %v\n", err)
			}
			fmt.Printf("%s = %s\n", Mile(num), MileToMetre(Mile(num)))
		}
	case *lengFlag == "nmi":
		for _, arg := range flag.Args() {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "convert false: %v\n", err)
			}
			fmt.Printf("%s = %s\n", Nauticalmile(num), NauticalmileToMetre(Nauticalmile(num)))
		}
	default:
		fmt.Println("no input")
	}
}
