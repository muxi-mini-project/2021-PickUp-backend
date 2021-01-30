package model

import (
	"strconv"
	"strings"
)

func TypeChange(spot string) int {
	switch spot {
	case "管理学院":
		return 1
	case "佑铭体育馆":
		return 2
	case "桂香园":
		return 3
	case "yy楼":
		return 4
	case "桂园广场":
		return 5
	case "3号楼":
		return 6
	case "8号楼":
		return 7
	case "利群书社":
		return 8
	case "图书馆":
		return 9
	case "校医院":
		return 10
	case "9号楼":
		return 11
	case "华师东区广场":
		return 12
	case "5号楼":
		return 13
	case "杜鹃广场":
		return 14
	case "数学与统计学院":
		return 15
	case "美术学院":
		return 16
	case "化学学院":
		return 17
	case "城市与环境管理学院":
		return 18
	case "大学生活动中心":
		return 19
	case "生命与科学学院":
		return 20
	case "网球场":
		return 21
	case "学子餐厅":
		return 22
	case "东二食堂":
		return 23
	case "东区小树林":
		return 24
	case "东一食堂":
		return 25
	case "音乐楼":
		return 26
	case "7号楼":
		return 27
	case "行政楼":
		return 28
	case "恽代英广场":
		return 29
	case "露天电影院":
		return 30
	}
	return 0
}

type Time1 struct {
	Year  int
	Month int
	Day   int
}

type Time2 struct {
	Hour int
	Min  int
}

//注意地点是用","间隔开
func MatchAandE(s string, start string, end string) int {
	var i, ok, differ int
	ss := ""
	spots := strings.SplitN(s, ",", 30)
	for key, value := range spots {
		i = key
		ss += value
	}
	ok = strings.Index(ss, start)
	if ok == -1 {
		return 200
	}
	ok = strings.Index(ss, end)
	if ok == -1 {
		return 40
	}
	differ = i - ok
	//fmt.Println(i, ok, differ)
	return differ
}

//注意发时间的形式是"xxx年xxx月xxx日"
func Time1Change(time string) Time1 {
	year := strings.SplitN(time, "年", 2)
	Tyear, _ := strconv.Atoi(year[0])
	//fmt.Println(year[0])
	month := strings.SplitN(year[1], "月", 2)
	Tmonth, _ := strconv.Atoi(month[0])
	//fmt.Println(month[0])
	day := strings.SplitN(month[1], "日", 2)
	Tday, _ := strconv.Atoi(day[0])
	//fmt.Println(day[0])
	mytime := Time1{
		Year:  Tyear,
		Month: Tmonth,
		Day:   Tday,
	}
	return mytime
}

//注意时间的形式是"xx:xx"
func Time2Change(time string) Time2 {
	hour := strings.SplitN(time, ":", 2)
	Thour, _ := strconv.Atoi(hour[0])
	min := hour[1]
	Tmin, _ := strconv.Atoi(min)
	mytime := Time2{
		Hour: Thour,
		Min:  Tmin,
	}
	return mytime
}

func PercentChange(percent int) string {
	s := strconv.Itoa(percent) + `%`
	return s
}

func MatchDegree(tmpP RequirePassenger, tmpD RequireDriver) int {
	Ptime := Time1Change(tmpP.Ymd)
	Dtime := Time1Change(tmpD.Ymd)
	PStime := Time2Change(tmpP.StartTime)
	PEtime := Time2Change(tmpP.EndTime)
	DStime := Time2Change(tmpD.StartTime)
	DEtime := Time2Change(tmpD.EndTime)
	percent := 0
	if Ptime.Year != Dtime.Year {
		return percent
	}

	if Ptime.Month != Dtime.Month {
		return percent
	}

	if Ptime.Day != Dtime.Day {
		return percent
	}

	if PStime.Hour > DEtime.Hour || DStime.Hour > PEtime.Hour {
		return percent
	}

	if tmpP.Urgent == 2 {
		percent += 10
	}

	if PStime.Min > DEtime.Min || DStime.Min > PEtime.Min {
		return percent
	}

	if PStime.Hour == DStime.Hour && PEtime.Hour == DEtime.Hour {
		if PStime.Min == DStime.Min && PEtime.Min == DEtime.Min {
			percent += 20
		} else {
			percent += 15
		}
	} else {
		percent += 10
	}

	s := tmpD.StartSpot + "," + tmpD.PassingSpots
	num := MatchAandE(s, tmpP.StartSpot, tmpP.EndSpot)
	if num == 200 {
		return percent
	}

	if num == 40 {
		return percent + 40
	}

	percent = percent + 80 - num*10
	if percent < 60 {
		return 60
	}

	if percent > 100 {
		return 100
	}

	return percent
}
