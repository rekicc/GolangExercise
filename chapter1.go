package exercise

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// 练习1.1, 1.2, 1.3
func echoCmdParameter(s []string, a bool) string {
	if a {
		//练习1.2, 逐行输出命令行参数的索引及其值
		for n, c := range s {
			fmt.Printf("cmd's index is %d, cmd's value is %s\n", n, c)
		}

		//练习1.3, 使用range遍历命令行参数, 然后将其连接起来并输出、打印执行时间
		t1 := time.Now()
		sep := " "
		str1 := ""
		for _, c := range s {
			str1 += c + sep
		}
		fmt.Printf("complete cmd is %s\n", str1)
		fmt.Printf("process spend %dns\n", time.Since(t1))

		//练习1.3, 使用strings.Join函数连接命令行参数, 然后输出, 并执行打印时间
		t2 := time.Now()
		str2 := strings.Join(s, " ")
		fmt.Printf("complete cmd is %s\n", str2)
		fmt.Printf("spend time is %dns\n", time.Since(t2))
	}
	//练习1.1, 输出命令行命令的名字
	fmt.Printf("cmd is %s\n", s[0])
	return s[0]
}

// 练习1.4
// fileInfo用于记录文件分析完后的信息, 其中name存储文件名, content的key存储每行的内容,value存储该内容在文件中的行号
type fileInfo struct {
	name    string
	content map[string][]int
}

// 解析文件的内容
func paraseFileStruct(f []*os.File) (r []fileInfo) {
	for _, f1 := range f {
		//必须在每次循环开始的时候都初始化一次fi, 这个fi存储每个文件的分析结果, 在分析完后会填充到返回值r中, 如果在循环外初始化会导致对每个文件的分析结果会被下一个结果覆盖
		var fi = new(fileInfo)
		//必须先初始化fileInfo中的map, 否则会导致在nil值的map上填充数据的错误
		fi.content = make(map[string][]int)
		fi.name = f1.Name()
		f1scan := bufio.NewScanner(f1)
		//n用于记录行号
		var n int = 1
		for f1scan.Scan() {
			fi.content[f1scan.Text()] = append(fi.content[f1scan.Text()], n)
			n++
		}
		r = append(r, *fi)
	}
	return
}

// 将文件的绝对路径切分, 只留文件名
func splitFilename(n string) string {
	_, filename := path.Split(n)
	return filename
}

// 接收paraseFileStruct解析后的内容, 分析其中是否有重复
func dup(f []*os.File) map[string][]string {
	res := paraseFileStruct(f)
	resLen := len(res)
	str := make(map[string][]string)
	//record记录当前内容是否被遍历过
	var record = make(map[string]bool)
	for i, f1 := range res {
		for k, v := range f1.content {
			if record[k] {
				continue
			}
			//flag1用于标识当前内容是否在其他文件中有重复
			var flag1 bool = false
			for x := i + 1; x <= resLen-1; x++ {
				v1, ok := res[x].content[k]
				if ok {
					res[x].name = splitFilename(res[x].name)
					str[k] = append(str[k], fmt.Sprintf("%s:%v\t", res[x].name, v1))
					flag1 = true
				}
			}
			if flag1 || len(v) > 1 {
				f1.name = splitFilename(f1.name)
				str[k] = append(str[k], fmt.Sprintf("%s:%v\t", f1.name, v))
			}
			record[k] = true
		}
	}
	return str
}

// 练习1.5, 1.6只是改几个参数,没做, 这里是原版的lissajous源码, 唯一的问题是lissajous具体怎么画出来gif图的
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
	if err != nil {
		log.Println(err)
	}
}

// 练习1.7, 1.9, 将获取到的内容输出到os.stdout, 并且打印状态码
func fetch(urlstring string) error {
	resp, err := http.Get(urlstring)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	s := os.Stdout
	_, err = io.Copy(s, resp.Body)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println(resp.Status)
	return nil
}

// 练习1.8, 如果URL没有http://前缀, 则为该URL添加这个前缀
func fetchAddPrefix(urlstring string) error {
	if strings.HasPrefix(urlstring, "http://") {
		return fetch(urlstring)
	}
	urlstring = "http://" + urlstring
	return fetch(urlstring)
}

// 练习1.10
func fetchAll(urlstring string, ch chan<- string) {
	startTime := time.Now()
	resp, err := http.Get(urlstring)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}
	defer resp.Body.Close()
	n := rand.Int31()
	filename := "/users/reki/Program/Go/src/pratice/exercise/resource/" + strconv.Itoa(int(n))
	f, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}
	count, err := io.Copy(f, resp.Body)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}
	spendTime := time.Since(startTime).Seconds()
	res := fmt.Sprintf("%s has %d byte, spend %fs", urlstring, count, spendTime)
	ch <- res
}

func serve1(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for x, y := range r.Form {
		if x == "Cycles" {
			for _, z := range y {
				z, err := strconv.Atoi(z)
				if err != nil {
					log.Println(err)
				}
				lissajous(w, z)
			}
		}
	}
}
