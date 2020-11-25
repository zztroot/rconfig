# rconfig  
**用于读取配置文件，现可以读取json、ini、conf等后缀文件。**  
  
## Installation  
Just ```pull zztroot/rconfig``` from github using ```go get```:
```
#download the code
go get github.com/zztroot/rconfig

#import the code 
import "github.com/zztroot/rconfig"
```
  
## 例子  
**1.读取ini文件**  
  
ini：  
```  
[server]
run = debug
ip = 0.0.0.0
port = 8080

[mysql]
db_host = 127.0.0.1
db_port = 3306
db_user = root
db_pwd = 123456
db_name = xxx

[redis]
db_host=127.0.0.1
db_port = 6379
```  
  
code:  
```golang
func main() { 
	file, _ := rconfig.OpenConfig("config.ini") 
	data := file.Get("server.port")  //8080
	run := file.Get("server.run")  //debug
	fmt.Println(data, run)  
}
```  
  
results:  
```
8080 debug
```  
  
**2.读取json文件**  
  
json:  
```
{
  "test": [
    {
      "name": "ArticleQuery",
      "desc": "left",
      "params": [
        {
          "name": "key",
          "desc": "xx"
        },
        {
          "name": "id",
          "desc": "123",
          "note": "xx"
        }
      ]
    },
    {
      "name": "TypeQuery",
      "desc": "xx",
      "params": [
        {
          "name": "key",
          "desc": "333"
        },
        {
          "name": "id",
          "desc": "222",
          "note": "rrr"
        }
      ]
    }
  ]
}
```  
  
code:  
```golang
func main() {  
	files, _ := rconfig.OpenJson("test.json")  
	name := files.Get("test.1.params.0.name")  //key
	desc := files.Get("test.1.params.0.desc")  //333
	fmt.Println(name, desc)  
}
```  
  
results:  
```
key 333
```  
  
**你还可以**
code:  
```golang
func main() {  
	files, _ := rconfig.OpenJson("test.json")  
	name := files.GetString("test.1.params.0.name")  //key
	desc := files.GetInt("test.1.params.0.desc")  //333
	fmt.Println(name, desc)  
}
```





