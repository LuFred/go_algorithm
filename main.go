package main

import (
	"log"
	"reflect"
	"strings"
)

func main() {
	m := ModelEntity{
		Ename: "sss",
		Eage:  "123",
	}
	 d := ModelDto{}
	Mapper(m, d)

	return
}

type ModelEntity struct {
	Ename string `amp:"Iname",json:"e_name"`
	Eage  string `amp:"Iage",json:"e_age"`
}

type ModelDto struct {
	Iname string
	Iage  string
}
type MapperInfo map[string]string

var info = make(MapperInfo)

func init() {

	ent := reflect.ValueOf(ModelEntity{})
	dto := reflect.ValueOf(ModelDto{})
	ent_f := ""
	for i := 0; i < ent.NumField(); i++ {
		ent_f += ent.Type().Field(i).Name + "_" + dto.Type().Field(i).Name + "_"

	}
	tn := ent.Type().Name() + "_" + dto.Type().Name()
	info[tn] = ent_f
	//log.Println(info)

}

var res = ModelDto{}

func Mapper(o interface{}, d interface{}) {
	ent := reflect.ValueOf(o)
	dto := reflect.ValueOf(d)

	mkey := ent.Type().Name() + "_" + dto.Type().Name()
	if info[mkey] != "" {
		keys := info[mkey]
		kl := strings.Split(keys, "_")
		eElem := reflect.ValueOf(&res).Elem()
		for i := 0; i < len(kl)-1; i += 2 {
			log.Println(eElem.FieldByName(kl[i+1]).CanSet())
			eElem.FieldByName(kl[i+1]).SetString("s")
		}
		log.Println("ttt", res)
	} else {

	}

}
