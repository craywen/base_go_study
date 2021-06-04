package test

import (
	"fmt"
	"time"
)

func TimeTest() {
	//当前时间
	nowTime :=	time.Now();
	fmt.Printf("%v\n",nowTime)
	fmt.Printf("%02d.%02d.%4d\n", nowTime.Year(), nowTime.Month(), nowTime.Year())
	//转成普通的 yyyy-MM-dd 普通时间
	nowTimeStr :=nowTime.Format("2006-01-02 15:04:05")
	fmt.Printf("%v\n",nowTimeStr)
	//转成时间戳 然后新增一天
	addTime :=nowTime.Unix()+3600*24;
	//将时间戳转化为 Time
	tm := time.Unix(addTime, 0)
	fmt.Printf("addTime: %v\n",time.Unix(addTime, 0).Format("2006-01-02 15:04:05"))
	// 判断两个时间的大小
    isTrue :=	tm.After(nowTime)
	fmt.Printf("%v\n",isTrue)
	//计算两个时间的时间差
	subTime :=tm.Sub(nowTime)
	fmt.Printf("时间差毫秒 %v ,秒 %v\n",subTime.Milliseconds(),subTime.Seconds())
	// 10分钟前
	m, _ := time.ParseDuration("-10m") // -1h
	nowTimeMin :=nowTime.Add(m)
	fmt.Printf("nowTimeMin %v\n ",nowTimeMin)
	//10分钟后
	tenAfter, _ :=time.ParseDuration("10m");
	nowTimetenAfter :=nowTime.Add(tenAfter)
	fmt.Printf("nowTimetenAfter %v\n ",nowTimetenAfter)

}
