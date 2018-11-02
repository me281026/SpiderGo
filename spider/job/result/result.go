package result

import (
	"../util/uuid"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//访问链接
var HttpUrl = "https://www.lagou.com/jobs/positionAjax.json?city=%s&needAddtionalResult=false"

//结果
type ListResult struct {
	Code    int
	Success bool
	Msg     string
	Content Content
}

//content对象
type Content struct {
	PositionResult PositionResult
	PageNum        int
	PageSize       int
}

//PositionResult对象
type PositionResult struct {
	DataResult []DataResult
	TotalCount int
}

//DataReult对象
type DataResult struct {
	City              string
	BusinessZones     []string
	CompanyFullName   string
	CompanyLabelList  []string
	CompanyShortName  string
	CompanySize       string
	CreateTime        string
	District          string
	Education         string
	FinanceStage      string
	FirstType         string
	IndustryField     string
	IndustryLables    []string
	JobNature         string
	Latitude          string
	Longitude         string
	PositionAdvantage string
	PositionId        int32
	PositionLables    []string
	PositionName      string
	Salary            string
	SecondType        string
	Stationname       string
	Subwayline        string
	Linestaion        string
	WorkYear          string
}

//工作服务
type JobSevice struct {
	City string
}

//创建JobSevice对象
func NewJobSevice(city string) *JobSevice {
	return &JobSevice{City: city}
}

//创建JobSevice的获取Job的信息
func (jobSevice *JobSevice) GetJobSeviceInfo(pn int, kd string) (*ListResult, error) {
	//创建http client抓群信息
	client := http.Client{}
	//创建POST提交请求
	PostReader := strings.NewReader(fmt.Sprintf("first=false&pn=%d&kd=%s", pn, kd))
	//创建request请求
	request, err := http.NewRequest("POST", jobSevice.GetUrl(), PostReader)
	if err != nil {
		log.Printf("ttp.NewRequest err: %v", err)
	}
	//设置请求的参数头
	request.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	request.Header.Add("Accept-Languag", "zh-CN,zh;q=0.9")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Content-Length", "25")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("Cookie", "_ga=GA1.2.161331334.1522592243; "+
		"user_trace_token=20180401221723-"+uuid.GetUUID()+"; "+
		"LGUID=20180401221723-"+uuid.GetUUID()+"; "+
		"index_location_city=%E6%B7%B1%E5%9C%B3; "+
		"JSESSIONID="+uuid.GetUUID()+"; "+
		"_gid=GA1.2.1140631185.1523090450; "+
		"Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1522592243,1523090450; "+
		"TG-TRACK-CODE=index_search; _gat=1; "+
		"LGSID=20180407221340-"+uuid.GetUUID()+"; "+
		"PRE_UTM=; PRE_HOST=; PRE_SITE=https%3A%2F%2Fwww.lagou.com%2F; "+
		"PRE_LAND=https%3A%2F%2Fwww.lagou.com%2Fjobs%2Flist_golang%3FlabelWords%3D%26fromSearch%3Dtrue%26suginput%3D; "+
		"Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1523110425; "+
		"LGRID=20180407221344-"+uuid.GetUUID()+"; "+
		"SEARCH_ID="+uuid.GetUUID()+"")
	request.Header.Add("Host", "www.lagou.com")
	request.Header.Add("Origin", "https://www.lagou.com")
	request.Header.Add("Referer", "https://www.lagou.com/jobs/list_golang?labelWords=&fromSearch=true&suginput=")

	//client调用request
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	//ioutil读取
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	//创建ResultList对象
	var listResult ListResult

	//JSON解析判断
	err = json.Unmarshal([]byte(bytes), &listResult)
	if err != nil {
		return nil, err
	}
	return &listResult, nil

}

//JobSevice的获取URL的方法
func (jobSevice *JobSevice) GetUrl() string {
	//url匹配格式化
	sprintf := fmt.Sprintf(HttpUrl, jobSevice.City)
	parse, err := url.Parse(sprintf)
	if err != nil {
		log.Printf(" JobSevice.GetUrl.err: %v", err)
	}
	//解析解码
	parse.RawQuery = parse.Query().Encode()
	return parse.String()

}
