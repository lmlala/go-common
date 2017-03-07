// Date : 4/1/17 PM3:15
// Copyright: TradeShift.com
// Author = 'liming'
// Email = 'lim@cn.tradeshift.com'
package common

import (
	"text/template"
	"bytes"
	"time"
	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
)

//type string_template string

//func (tmpl *string_template) StringFormat (data interface{}) (r string) {
func StringFormat (tmpl string, data interface{}) (r string, err error) {
	// todo: different format of data

	switch data.(type) {
	//case string:
	//    	return  "", Err("Format: Not a string data")
	default:
	    	t, _ := template.New("test").Parse(tmpl)

		var buff bytes.Buffer

		err = t.Execute(&buff, data)

		if err != nil {
			r = tmpl
			return r, Err("Format Error: " + tmpl)
		} else {
			r = buff.String()
			return r, nil
		}
	}

}


//func JsonToStruct (data []byte, model struct{}) (s interface{}, err error) {
//
//	switch model.(type) {
//	default:
//		var error_model = ""
//		return error_model, Err("Format: Not a string data")
//	case Nomad:
//		err = json.Unmarshal(data, &model)
//		if err != nil {
//			return nil, Err("Json load faild: " + string(data))
//		}
//		return model, nil
//	}
//
//}

func Nano2Date(ns int64) (date string) {
	//f := "2006-01-02 15:04:05"
	//t := time.Now().UnixNano()
	//fmt.Println(t)
	//tt := t/1e9
	//fmt.Println(tt)
	//ttt := time.Unix(tt, 0)
	//fmt.Println(ttt)
	//fmt.Println(ttt.String())
	//tttt := ttt.Format(f)
	//fmt.Println(tttt)

	f := "2006-01-02 15:04:05"   // must be golang born time :)
	date = time.Unix(ns / 1e9, 0).Format(f)
	return
}

func YamlLoad(path string, d interface{}) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	// d already be a point
	err = yaml.Unmarshal(fd, d)
	if err != nil {
		return err
	}

	return nil
}

func YamlLoadWithoutStruct(path string) (d map[string]interface{}, err error) {
	f, err := os.Open(path)
	if err != nil {
		return d, err
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return d, err
	}

	err = yaml.Unmarshal(fd, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}
