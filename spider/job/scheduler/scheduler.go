package scheduler

import "sync"

var (
	mutex      sync.Mutex
	jobInfoArr = []JobInfo{}
)

//调度的JobInfo参数
type JobInfo struct {
	City string
	Pn   int
	Kd   string
}

//创建JobInfo对象
func NewJobScheduler() *JobInfo {
	return &JobInfo{}

}

//append方法
func (jobInfo *JobInfo) Append(city, kd string, pn int) {
	//使用append的时候先上锁
	mutex.Lock()

	jobInfoArr = append(jobInfoArr, JobInfo{
		City: city,
		Pn:   pn,
		Kd:   kd,
	})

	mutex.Unlock()

}

//count计数方法
func (jobInfo *JobInfo) Count() int {
	return len(jobInfoArr)
}

//Pop方法
func (jobInfo *JobInfo) Pop() *JobInfo {
	mutex.Lock()
	length := len(jobInfoArr)
	if length < 1 {
		return nil
	}
	//取值
	job := jobInfoArr[length-1]
	jobInfoArr = jobInfoArr[:length-1]

	//解锁
	mutex.Unlock()
	return &job
}
