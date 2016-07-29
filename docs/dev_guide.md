### DEV GUIDE

model层
----

定义所有可能services层用到的方法, 方法简单, 只做一件事情, 避免在service层里做db操作, 这样方便底层更新database接口, 如果某一db库用得不爽, 只需要改model层, 而services层不用动. 已经吃过这个亏了, 从pg换到mysql, 又从go-pg/pg换到gorm, 又想从gorm换sqlx, 只要保证services层里用到model操作方法, 在model层里都定义好, 那么这样的改动, 就不会头疼了.

不需要defer操作, 直接返回error, defer操作放到services里Do方法里用于纪录错误, 不需要对error有太多自定义, 直接纪录原生错误

service层
----

所有的业务写到此, service层可组合成新的业务, 所有的controller层只能调业务层, 不能调用model层的任何方法(除了初始化db, InitModels()在server启动时可以被调用外), services层可调用model层的方法.


controller层
----

controller层借用form层用于绑定数据, 以及做业务层前的检测, 通过middlerware层用于简单业务逻辑共享. 禁止调用model层的方法
