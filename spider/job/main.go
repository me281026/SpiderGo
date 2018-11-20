package main

import (
	"./execute"
	"./pipline"
	"log"
	"sync"
)

//定义变量
var (
	//设定抓取语言
	lang = []string{
		"golang",
	}

	//抓群的城市
	citys = []string{
		"北京",
		"上海",
		"深圳",
		"杭州",
		"广州",
		"武汉",
		"南京",
		"成都",
	}

	//入参
	InitParam = []execute.InitParam{}

	//出参
	ResultParam = []execute.ResultParam{}

	//初始化接入数据
	JobInfo = pipline.InitJobInfo()

	//协程计数
	wg sync.WaitGroup
)

func main() {

	//依次传入参数
	for _, kd := range lang {
		for _, city := range citys {
			//传入一个执行协程
			wg.Add(1)
			go func(kd string, city string) {
				//进入协程后取出
				defer wg.Done()
				//初始化参数并传入
				params, err := execute.Initexe(city, kd, 1)
				if err != nil {
					log.Fatalln(err)
				}
				InitParam = append(InitParam, params...)
				ResultParam = append(ResultParam, execute.ResultJobs())
			}(city, kd)

		}
	}
	//wait
	wg.Wait()
	JobInfo.Push()
	//log记录打印
	log.Printf("InitParam: %v", InitParam)
	log.Printf("ResultParam: %v", ResultParam)
}
