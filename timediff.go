package timedifference

import (
	//"fmt"
	"strconv"
	"time"
)

func TimeDifference(postdate string)string{
loc,_:=time.LoadLocation("UTC")
dt:=time.Now().In(loc)
ct:=dt.Format("2006-01-02 3:04:05")
x,_:=time.Parse("2006-01-02 3:04:05",postdate)
y,_:=time.Parse("2006-01-02 3:04:05",ct)
diff:=y.Sub(x)
//fmt.Println(int(diff.Hours()))
timediff:=int(diff.Hours())
if timediff<=24{
	if timediff<1{
		min:=diff.Hours()*60
		if int(min)>=1{
		minutes:=strconv.Itoa(int(min))
		return "posted "+minutes+" minutes ago"
		}
		if int(min)<1{
			return "posted a few seconds ago"
		}
	}
	if timediff>=1{
		hours:=strconv.Itoa(timediff)
		return "posted "+hours+" hour ago"
	}
}
if timediff>24{
	days:=timediff/24
	if days<=30{
	nofdays:=strconv.Itoa(days)
	return "posted "+nofdays+" days ago"
}
if days>30{
	months:=days/30
	if months<12{
	nofmonths:=strconv.Itoa(months)
	return "posted "+nofmonths+" months ago"
	}
	if months>=12{
		years:=months/12
		nofyears:=strconv.Itoa(years)
		if years==1{
			
		return "posted "+nofyears+" year ago"
		}
		if years>1{
			return "posted "+nofyears+" years ago"
		}
	}
}
}
return "error"
}
