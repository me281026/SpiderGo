package execute

import (
	"../pipline"
	"../result"
	"../scheduler"
	"../util/convert"
	"../util/page"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"time"
)

var (
	delayTime    = time.Tick(time.Millisecond * 500)
	jobScheduler = scheduler.NewJobScheduler()
	JobPipline   = pipline.InitJobInfo()
)

//入参
type InitParam struct {
	City       string
	Kd         string
	TotalPage  int
	TotalCount int
}

//出参
type ResultParam struct {
	Success int
	Error   int
	Empty   int
	Errors  []string
}

//初始化
func Initexe(city, kd string, pn int) ([]InitParam, error) {
	//初始化变量
	var (
		jobs       []result.DataResult
		totalPage  int
		totalCount int
		err        error

		init []InitParam
	)
	//获取Job信息
	jobs, totalPage, totalCount, err = GetJob(city, kd, pn)
	if err != nil {
		return nil, err
	}

	//初始化入参
	init = append(init, InitParam{
		City:       city,
		Kd:         kd,
		TotalCount: totalCount,
		TotalPage:  totalPage,
	})
	//通过循环执行抓群页面大于2的
	for i := 2; i <= totalPage; i++ {
		jobScheduler.Append(city, kd, pn)
	}
	//pipline添加
	JobPipline.Append(convert.ToPiplineJob(jobs))
	return init, nil
}

//获取Job信息
func GetJob(city, kd string, pn int) ([]result.DataResult, int, int, error) {
	//从起始页码开始为0
	totalPage := 0
	//初始化JobSevice对象
	jobSevice := result.NewJobSevice(city)
	listResult, err := jobSevice.GetJobSeviceInfo(pn, kd)
	if err != nil {
		return nil, 0, 0, err
	}
	//日志记录
	log.Printf("GetJobs Code: %d, GetJobs City: %s, Pn: %d, Kd: %s", listResult.Code, city, pn, kd)
	//result处理判断
	if listResult.Code == 0 && listResult.Success == true {
		content := listResult.Content
		//判断页码和条数
		if content.PositionResult.TotalCount > 0 && content.PageSize > 0 {
			totalPage = page.CalculatePage(float64(content.PositionResult.TotalCount), float64(content.PageSize))
		}
	} else {
		return nil, 0, 0, errors.New(fmt.Sprintf("GetJobs City: %s, Pn: %d, Kd: %s, Result: %v", city, pn, kd, listResult))
	}
	//返回结果
	return listResult.Content.PositionResult.DataResult, totalPage, listResult.Content.PositionResult.TotalCount, nil
}

//循环信息抓取
func ResultJobs() ResultParam {
	//定义数值和Chan
	var (
		rp    ResultParam
		count = jobScheduler.Count()

		paramChan = make(chan []result.DataResult)
	)
	//循环调用执行器执行
	for i := 0; i < count; i++ {
		//设置延迟
		<-delayTime
		//开启协程
		go func() {
			if job := jobScheduler.Pop(); job != nil {
				results, _, _, err := GetJob(job.City, job.Kd, job.Pn)
				//判断err,并记录,如果无err,则放入chan
				if err != nil {
					rp.Error++
					rp.Errors = append(rp.Errors, err.Error())
				} else {
					paramChan <- results
				}
			} else {
				rp.Empty++
			}

		}()
	}

	//处理chan,设置跳出标记
JUMP:
	for {
		select {
		case param := <-paramChan:
			rp.Success++
			JobPipline.Append(convert.ToPiplineJob(param))
		default:
			if (rp.Success + rp.Empty + rp.Error) >= count {
				log.Printf("Break..........")
				break JUMP
			}

		}
	}
	return rp

}
