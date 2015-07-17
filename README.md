# go-conf

简单易用的Go语言ini文件解析工具

### 安装

	go get github.com/lisijie/go-conf
	
### 使用

	c, err := goconf.NewConfig("example.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test_str:", c.GetString("test_str"))
	
	c.SetString("foo", "bar")
	c.Save()