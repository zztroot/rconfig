package main

import (
	"fmt"
	"v1/rconfig"
)

func main() {
	file, _ := rconfig.OpenConfig("config.ini")
	data := file.GetInt("server.port") //8080
	run := file.Get("server.run") //debug
	//fmt.Println(data, run)
	fmt.Printf(`%t,%t`, data, run)

	files, _ := rconfig.OpenJson("api.json")
	name := files.Get("test.1.params.0.name") //key
	desc := files.GetInt("test.1.params.0.desc") //333
	fmt.Printf(`%t,%t`, name, desc)

	//get map
	//file, err := rconfig.OpenJson("api.json")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//m := file.GetMap()
	//fmt.Printf("%t", m)
}
