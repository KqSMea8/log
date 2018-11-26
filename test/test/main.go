package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ftoa.go
//------------------------------------------------------------
// FormatFloat 将浮点数 f 转换为字符串值
// f：要转换的浮点数
// fmt：格式标记（b、e、E、f、g、G）
// prec：精度（数字部分的长度，不包括指数部分）
// bitSize：指定浮点类型（32:float32、64:float64）
//// 格式标记：
// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)
//// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
//func FormatFloat(f float64, fmt byte, prec, bitSize int) string
func main() {
	//fmt.Println(time.Now().Unix())
	// 计算下一个零点

	//t := time.NewTimer()
	//<-t.C
	//fmt.Printf("定时结算Boottime表数据，结算完成: %v\n",time.Now().Unix())
	//t2 := time.Unix(1543204161, 0)
	//fmt.Println(t2)
	//
	//t2 = time.Unix(1543204161+200, 0)
	//fmt.Println(t2)

	var sig = make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM)

	go test()

	for {
		select {
		case <-sig:
			fmt.Println("Signal received.  Shutting down...\n")
			return
			//case <-ticker.C:
			//	fmt.Println("break heart")
		}
	}
}

func test() {
	time1 := time.Now().Unix()
	fmt.Println("time1:", time.Unix(time1, 0))
	time2 := time1 + 5
	fmt.Println("time2:", time.Unix(time2, 0))
	t := time.NewTicker(time.Duration(time2 - time1)*time.Second)


	//t := time.NewTimer(time.Second * 5)
	for {
		select {
		case <-t.C:
			t.Stop()
			fmt.Println("hehe \n")
			return
			//case <-ticker.C:
			//	fmt.Println("break heart")
		}
	}
}

//定时结算Boottime表数据
func BoottimeTimingSettlement() {
	for {
		now := time.Now()
		// 计算下一个零点
		next := now.Add(time.Second * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		fmt.Printf("定时结算Boottime表数据，结算完成: %v\n", time.Now())
		//以下为定时执行的操作
	}
}
