package rjson

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type JsonStruct struct {
	Data map[string]interface{}
}

func (j *JsonStruct)loading(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	j.Data = m
	return nil
}

func OpenJson(path string) (*JsonStruct, error) {
	j := JsonStruct{}
	err := j.loading(path)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &j, nil

}

func (j *JsonStruct) Get(str string) (interface{}, error){
	m := j.Data
	g := strings.Split(str, ".")
	var result interface{}
	for _, v := range g {
		if result == nil {
			result = m[v].(interface{})
			continue
		}else if _, err := strconv.Atoi(v); err == nil{
			index, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result = result.([]interface{})[index]
			continue
		}else {
			result = result.(map[string]interface{})[v]
			continue
		}
	}
	return result, nil
}

