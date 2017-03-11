// Date : 7/3/17 AM9:25
// Copyright: TradeShift.com
// Author = 'liming'
// Email = 'lim@cn.tradeshift.com'
package common

import (
	"time"
)

type TimeStr struct {
	Stamp	time.Time
	Str	string
}


func Now() (t TimeStr) {

	t.Stamp = time.Now()
	t.Str = t.Stamp.Format("2006-01-02 15:04:05")
	return t
}

func ThisTime(this time.Time) (t TimeStr){
	t.Stamp = this
	t.Str = this.Format("2006-01-02")
	return t
}

func ThisTimeStr(this string) (err error, t TimeStr){
	t.Str = this
	t.Stamp, err = time.Parse("2006-01-02 15:04:05", this)
	return err, t
}

func (t TimeStr) DateTime() (string){
	return t.Stamp.Format("2006-01-02 15:04:05") // must be "2006-01-02 15:04:05"
}

func (t TimeStr) Month() (string){
	return t.Stamp.Format("2006-01")
}

func (t TimeStr) MonthlyTable() (string){
	return t.Stamp.Format("2006_01")
}

func (t TimeStr) Day() (string){
	return t.Stamp.Format("2006-01-02")
}

func (t TimeStr) Hour() (string){
	return t.Stamp.Format("2006-01-02 15")
}

func (t TimeStr) Minute() (string){
	return t.Stamp.Format("2006-01-02 15:04")
}
