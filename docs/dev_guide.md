### DEV GUIDE

此服务经过长期考虑以及实践, 将其转变成一个基于gRPC+NSQ+Docker的微服务架构的系统, 主要分为以下主要服务:
accountCenter:
    1. 提供用户注册, 登录等操作
relationCenter:
    1. 用户关注
groupCenter:
    1. 跑团管理
runDataCenter:
    1. 跑步数据
runAnalysis:
    1. 分析跑步数据
commentCenter:
    1. 用于评论处理
geoCenter:
    1. 地理位置分析服务

其它:
    推送服务, 短信发送服务, 日志服务等

model层
----

定义所有可能services层用到的方法, 方法简单, 只做一件事情, 避免在service层里做db操作, 这样方便底层更新database接口, 如果某一db库用得不爽, 只需要改model层, 而services层不用动. 已经吃过这个亏了, 从pg换到mysql, 又从go-pg/pg换到gorm, 又想从gorm换sqlx, 只要保证services层里用到model操作方法, 在model层里都定义好, 那么这样的改动, 就不会头疼了.

不需要defer操作, 直接返回error, defer操作放到services里Do方法里用于纪录错误, 不需要对error有太多自定义, 直接纪录原生错误

service层
----

所有的业务写到此, service层可组合成新的业务, 所有的controller层只能调业务层, 不能跨层调用, services层可调用model层的方法.

service层不必关心请求来自于http还是cli,或者worker, 只关心数据, 这样让service层可以工作更多

解除service层对binging层的依赖

下层不能知道上层的任何东西, 所以目前service层使用controller层的对象是错误的, service层不能知道上层的任何对象, 如果services层需要一个对象, 则定义一个接口, 让上层去实现此接口, 从而解除对上层的依赖


controller层
----

controller层借用binging层用于绑定数据, 以及做业务层前的检测, 通过middlerware层用于简单业务逻辑共享. 禁止调用model层的方法

binding层用于解析http.Request里绑定的数据, 然后传给service层作为DataConfig层调用

middlewares 用于通用的HTTP操作, 比如CSRF控制 CORS请求, Gzip response等


worker层
----

worker层主要用于一些后台处理的job, 比如作为RabbitMQ的worker, 或者发送sms的worker, 或者发送推送的服务等, 可以开启多个workers
worker层只能调用service层, 不能直接调用model层


console层
----

console层主要用于一些命令行工作, 比如dump sql schema, generate docs等
