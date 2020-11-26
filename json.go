package rconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type JsonStruct struct {
	Data map[string]interface{}
	Path string
}

func (j *JsonStruct) loading(path string) (*map[string]interface{}, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	j.Data = m
	j.Path = path
	return &m, nil
}

func OpenJson(path string) (*JsonStruct, error) {
	j := JsonStruct{}
	_, err := j.loading(path)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &j, nil

}

func (j *JsonStruct) Get(str string) interface{} {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("Field error:field does not exist")
		}
	}()
	m := j.Data
	g := strings.Split(str, ".")
	var result interface{}
	for _, v := range g {
		if result == nil {
			result = m[v].(interface{})
			continue
		} else if _, err := strconv.Atoi(v); err == nil {
			index, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
				return nil
			}
			result = result.([]interface{})[index]
			continue
		} else {
			result = result.(map[string]interface{})[v]
			continue
		}
	}
	return result
}

func (j *JsonStruct) GetString(str string) string {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("Field error:field does not exist")
		}
	}()
	m := j.Data
	g := strings.Split(str, ".")
	var result interface{}
	for _, v := range g {
		if result == nil {
			result = m[v].(interface{})
			continue
		} else if _, err := strconv.Atoi(v); err == nil {
			index, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
				return ""
			}
			result = result.([]interface{})[index]
			continue
		} else {
			result = result.(map[string]interface{})[v]
			continue
		}
	}
	return result.(string)
}

func (j *JsonStruct) GetInt(str string) int {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(fmt.Sprintf(`Field error or %v erro`, err))
		}
	}()
	m := j.Data
	g := strings.Split(str, ".")
	var result interface{}
	for _, v := range g {
		if result == nil {
			result = m[v].(interface{})
			continue
		} else if _, err := strconv.Atoi(v); err == nil {
			index, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
				return 0
			}
			result = result.([]interface{})[index]
			continue
		} else {
			result = result.(map[string]interface{})[v]
			continue
		}
	}
	ms, _ := strconv.Atoi(result.(string))
	return ms
}

func (j *JsonStruct) GetMap() map[string]interface{} {
	 m, err := j.loading(j.Path)
	 if err != nil {
	 	log.Fatalln(err)
	 }
	 return *m
}
