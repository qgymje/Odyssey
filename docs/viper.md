Viper作为Config工具
===

[godoc](https://godoc.org/github.com/spf13/viper)

初始化viper
```
config = viper.New()
config.AddConfigPath("/path/to/config/directory")
config.SetConfigName("/config.yml")
if err := config.ReadInConfig(); err != nil {
  panic("read config error")
}
```

通过viper获取值
```
config.GetString("app.http_port")
//获取一个key, 得到一个map
dbconfig := config.GetStringMapString("database")
dbconfig["database"]
```
