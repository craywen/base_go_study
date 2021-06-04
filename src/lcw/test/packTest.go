package test

import (
	"fmt"
	"math"
	"math/big"
	"mygo/lcw/aaabbb"
	pack "mygo/lcw/pack"
	"regexp"
	"strconv"
	"sync"
)

type Info struct {
	mu sync.Mutex //sync.Mutex 是一个互斥锁，它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区。
}

const LINUX_REBOOT_MAGIC1 uintptr = 0xfee1dead
const LINUX_REBOOT_MAGIC2 uintptr = 672274793
const LINUX_REBOOT_CMD_RESTART uintptr = 0x1234567

func PackTest() {

	//正则包
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	match, err := regexp.Match(pat, []byte(searchIn))
	if match {
		fmt.Println("Match Found! %s %s", err, match)
	}
	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	//参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

	//锁和sync 包

	//精密计算的big包
	bigOne := big.NewInt(math.MaxInt64)
	bigTwo := bigOne
	bigThree := big.NewInt(1937)
	bigFour := big.NewInt(1)
	div := bigFour.Mul(bigOne, bigTwo).Add(bigFour, bigOne).Div(bigFour, bigThree)
	fmt.Printf("Big Int: %v\n", div)
	//
	rm := big.NewRat(math.MaxInt64, 1956)
	fmt.Printf("Big Rat: %v\n", rm)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)

	//使用自定义包 千辛万苦
	//1:要看环境变量(项目是否在GOPATH的src 目录下)
	//2:idea环境变量 设置GOPATH
	//3:是否是特殊的包名如果是,取一个别名  方法首字母大写表示public 小写只有内部可以调用
	aaabbb.Re()
	rStr := pack.ReturnStr()
	fmt.Printf("packTest rStr: %v\n", rStr)

	//
}

func Update(info *Info) {
	info.mu.Lock()
	// critical section:
	fmt.Errorf("锁逻辑块")
	//代码逻辑
	info.mu.Unlock()
}
