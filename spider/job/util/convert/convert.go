package convert

import (
	"../../pipline"
	"../../result"
	"../etime"
	"strconv"
	"strings"
	"time"
)

func ToPiplineJob(result []result.DataResult) []pipline.JobInfo {
	var resultJobinfos []pipline.JobInfo
	for _, v := range result {
		//类型转换
		longitude, _ := strconv.ParseFloat(v.Longitude, 64)
		latitude, _ := strconv.ParseFloat(v.Latitude, 64)
		resultJobinfos = append(resultJobinfos, pipline.JobInfo{
			// 城市名称
			City: v.City,
			// 地区
			District: v.District,

			// 公司简称
			CompanyShortName: v.CompanyShortName,
			// 公司全称
			CompanyFullName: v.CompanyFullName,
			// 公司标签
			CompanyLabelList: strings.Join(v.CompanyLabelList, ","),
			// 公司规模
			CompanySize: v.CompanySize,
			// 融资阶段
			FinanceStage: v.FinanceStage,

			// 行业领域
			IndustryField: v.IndustryField,
			// 行业标签
			IndustryLables: strings.Join(v.IndustryLables, ","),

			// 职位名称
			PositionName: v.PositionName,
			// 职位标签
			PositionLables: strings.Join(v.PositionLables, ","),
			// 职位诱惑
			PositionAdvantage: v.PositionAdvantage,
			// 工作年限
			WorkYear: v.WorkYear,
			// 学历要求
			Education: v.Education,
			// 薪资范畴
			Salary: v.Salary,

			// 经度
			Longitude: longitude,
			// 纬度
			Latitude: latitude,
			// 附近的地铁
			Linestaion: v.Linestaion,

			// 发布时间
			CreateTime: etime.MustDateToUnix(v.CreateTime),
			// 新增时间
			AddTime: time.Now().Unix(),
		})

	}
	return resultJobinfos

}
