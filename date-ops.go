// Date : 7/3/17 AM9:26
// Copyright: TradeShift.com
// Author = 'liming'
// Email = 'lim@cn.tradeshift.com'
package common

import "time"

func GetDateStr(t time.Time, f string) (d string) {
	// f: format   datetime, day, min, month, hour, default is datetime
	//if t == nil {
	//	t = time.Now()
	//}
	switch f {
	case "day":
		d = t.Format("2006-01-02")
	case "min":
		d = t.Format("2006-01-02-15-04")
	case "hour":
		d = t.Format("2006-01-02-15")
	default:
		d = t.Format("2006-01-02 15:04:05") // must be "2006-01-02 15:04:05"
	}

	return d
}

func GetNowStr(f string) (d string) {
	// f: format   datetime, day, min, month, hour, default is datetime
	//if t == nil {
	//	t = time.Now()
	//}
	t := time.Now()
	switch f {
	case "day":
		d = t.Format("2006-01-02")
	case "min":
		d = t.Format("2006-01-02-15-04")
	case "hour":
		d = t.Format("2006-01-02-15")
	default:
		d = t.Format("2006-01-02 15:04:05") // must be "2006-01-02 15:04:05"
	}

	return d
}