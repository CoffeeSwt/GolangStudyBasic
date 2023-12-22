## Grom操作数据库

Grom是golang中的ORM框架，用于将数据库表结构和golang中的结构体（struct）关联起来

### 结构体（struct）定义

```go
type User struct {
	ID       uint      `gorm:"comment:用户ID"`
	UserName string    `gorm:"comment:用户名"`
	Gender   string    `gorm:"comment:性别"`
	Birthday time.Time `gorm:"comment:出生日期"`
}
```

例如上面这一段定义，映射到数据库中就会生成一个名为**user**的表，其中包含了id,user_name,gender,birthday这几个字段，同时如果struct中有id的话，gorm会默认设置其为主键

定义配置可以参考文档
[模型定义 | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/zh_CN/docs/models.html#%E5%AD%97%E6%AE%B5%E6%A0%87%E7%AD%BE)

### 连接数据库
