package rconfig

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type configStruct struct {
	data string
}

func (c *configStruct) loading(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	c.data = string(data)
	return nil
}

func OpenConfig(path string) (*configStruct, error) {
	j := configStruct{}
	err := j.loading(path)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &j, nil
}

func (c *configStruct) Get(str string) interface{} {
	se := strings.Split(str, ".")
	var s1 string
	var s2 string
	data := strings.Split(c.data, "[")
	for _, v := range data {
		if strings.Contains(v, se[0]) {
			s1 = v
			break
		}
	}
	if s1 == "" {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, se[0]))
		return nil
	}
	re := regexp.MustCompile(`(.*?)]`)
	res := re.FindStringSubmatch(s1)
	if res[1] != se[0] {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, res[1]))
		return nil
	}
	line := strings.Split(s1, "\n")
	for _, v := range line {
		if strings.Contains(v, se[1]) {
			s2 = v
			if strings.Contains(s2, "#") {
				continue
			}
			break
		}
	}
	if s2 == "" {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, se[1]))
		return nil
	}
	var newS2 string
	if strings.Contains(s2, " ") {
		newS2 = strings.Join(strings.Fields(s2), "")
	}
	var results interface{}
	if newS2 != "" {
		re = regexp.MustCompile(`(.*?)=`)
		res = re.FindStringSubmatch(newS2)
		if res[1] != se[1] {
			log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, res[1]))
			return nil
		}
		field := strings.Split(newS2, "=")
		results = strings.TrimRight(field[1], "\r")
		return results
	}
	re = regexp.MustCompile(`(.*?)=`)
	res = re.FindStringSubmatch(s2)
	if res[1] != se[1] {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, res[1]))
		return nil
	}
	field := strings.Split(s2, "=")
	results = strings.TrimRight(field[1], "\r")
	return results
}

func (c *configStruct) GetString(str string) string {
	se := strings.Split(str, ".")
	var s1 string
	var s2 string
	data := strings.Split(c.data, "[")
	for _, v := range data {
		if strings.Contains(v, se[0]) {
			s1 = v
			break
		}
	}
	if s1 == "" {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, se[0]))
		return ""
	}
	re := regexp.MustCompile(`(.*?)]`)
	res := re.FindStringSubmatch(s1)
	if res[1] != se[0] {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, res[1]))
		return ""
	}
	line := strings.Split(s1, "\n")
	for _, v := range line {
		if strings.Contains(v, se[1]) {
			s2 = v
			if strings.Contains(s2, "#") {
				continue
			}
			break
		}
	}
	if s2 == "" {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, se[1]))
		return ""
	}
	var newS2 string
	if strings.Contains(s2, " ") {
		newS2 = strings.Join(strings.Fields(s2), "")
	}
	if newS2 != "" {
		re = regexp.MustCompile(`(.*?)=`)
		res = re.FindStringSubmatch(newS2)
		if res[1] != se[1] {
			log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, res[1]))
			return ""
		}
		field := strings.Split(newS2, "=")
		results := strings.TrimRight(field[1], "\r")
		return results
	}
	re = regexp.MustCompile(`(.*?)=`)
	res = re.FindStringSubmatch(s2)
	if res[1] != se[1] {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, res[1]))
		return ""
	}
	field := strings.Split(s2, "=")
	results := strings.TrimRight(field[1], "\r")
	return results
}

func (c *configStruct) GetInt(str string) int {
	se := strings.Split(str, ".")
	var s1 string
	var s2 string
	data := strings.Split(c.data, "[")
	for _, v := range data {
		if strings.Contains(v, se[0]) {
			s1 = v
			break
		}
	}
	if s1 == "" {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, se[0]))
		return 0
	}
	re := regexp.MustCompile(`(.*?)]`)
	res := re.FindStringSubmatch(s1)
	if res[1] != se[0] {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, res[1]))
		return 0
	}
	line := strings.Split(s1, "\n")
	for _, v := range line {
		if strings.Contains(v, se[1]) {
			s2 = v
			if strings.Contains(s2, "#") {
				continue
			}
			break
		}
	}
	if s2 == "" {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, se[1]))
		return 0
	}
	var newS2 string
	if strings.Contains(s2, " ") {
		newS2 = strings.Join(strings.Fields(s2), "")
	}
	if newS2 != "" {
		re = regexp.MustCompile(`(.*?)=`)
		res = re.FindStringSubmatch(newS2)
		if res[1] != se[1] {
			log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, res[1]))
			return 0
		}
		field := strings.Split(newS2, "=")
		results := strings.TrimRight(field[1], "\r")
		m, _ := strconv.Atoi(results)
		return m
	}
	re = regexp.MustCompile(`(.*?)=`)
	res = re.FindStringSubmatch(s2)
	if res[1] != se[1] {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, res[1]))
		return 0
	}
	field := strings.Split(s2, "=")
	results := strings.TrimRight(field[1], "\r")
	m, _ := strconv.Atoi(results)
	return m
}
