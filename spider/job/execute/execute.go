package execute

import (
	"../pipline"
	"../result"
	"../scheduler"
	"../util/page"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

var (
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
