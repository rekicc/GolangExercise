package exercise

import (
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// 测试1.1-1.3-echoCmdParameter
func TestEchoCmdParameter(t *testing.T) {
	/* type cmdParameter struct {
		cmd    []string
		toggle bool
		want   string
	} */
	echoCmdParameter(os.Args, true)
	if s := echoCmdParameter(os.Args, false); s != os.Args[0] {
		t.Error("false while test os.Args[0]")
	}
}

// 测试1.4-dup
func TestDup(t *testing.T) {
	f1, err := os.Open("/Users/reki/Program/Go/src/pratice/exercise/chapter1_0.text")
	if err != nil {
		log.Println(err)
	}
	defer f1.Close()
	f2, err := os.Open("/Users/reki/Program/Go/src/pratice/exercise/chapter1_1.text")
	if err != nil {
		log.Println(err)
	}
	defer f2.Close()
	f3, err := os.Open("/Users/reki/Program/Go/src/pratice/exercise/chapter1_2.text")
	if err != nil {
		log.Println(err)
	}
	defer f3.Close()
	var f []*os.File
	f = append(f, f1)
	f = append(f, f2)
	f = append(f, f3)
	res := dup(f)
	for k, v := range res{
		n := strings.Join(v, " ")
		fmt.Printf("total has %s in:%s\n", k, n)
	}
}

func TestServe1(t *testing.T) {
	req := httptest.NewRequest("GET", "http://127.0.0.1/?Cycles=10", nil)
	w := httptest.NewRecorder()

	serve1(w, req)
	res := w.Result()

	fmt.Println(res.Status)
	fmt.Println(res.Header.Get("Content-Type"))
	f, err := os.Create("/Users/reki/Program/Go/src/pratice/exercise/chapter1_web_lisajous_result.gif")
	if err != nil {
		fmt.Println(err)
	}
	if _, err := f.ReadFrom(w.Body); err != nil {
		fmt.Println(err)
	}
}
