# rconfig  
**用于读取配置文件，现可以读取json、ini、conf等后缀文件。**  

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
	data := file.Get("server.port")
	run := file.Get("server.run")
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
	name := files.Get("test.1.params.0.name")
	desc := files.Get("test.1.params.0.desc")
	fmt.Println(name, desc)
}
```  
  
results:  
```
key 333
```  
  




