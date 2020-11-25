package rconfig

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type ConfigStruct struct {
	Data string
}

func (c *ConfigStruct) loading(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	c.Data = string(data)
	return nil
}

func OpenConfig(path string) (*ConfigStruct, error) {
	j := ConfigStruct{}
	err := j.loading(path)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &j, nil
}

func (c *ConfigStruct) Get(str string) interface{} {
	se := strings.Split(str, ".")
	var s1 string
	var s2 string
	data := strings.Split(c.Data, "[")
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
	if newS2 != "" {
		re = regexp.MustCompile(`(.*?)=`)
		res = re.FindStringSubmatch(newS2)
		if res[1] != se[1] {
			log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, res[1]))
			return nil
		}
		field := strings.Split(newS2, "=")
		results := strings.TrimRight(field[1], "\r")
		return results
	}
	re = regexp.MustCompile(`(.*?)=`)
	res = re.FindStringSubmatch(s2)
	if res[1] != se[1] {
		log.Fatalln(fmt.Sprintf(`This field does not exist: %s`, res[1]))
		return nil
	}
	field := strings.Split(s2, "=")
	results := strings.TrimRight(field[1], "\r")
	return results
}
