- [ ] todo 登陆页面verify.js
> https://blog.csdn.net/qq_28073073/article/details/85333819

- [ ] 定时任务 
> 上传文件定时扫描详情入库

- [ ] 前端模板统一化 

> Bind ShouldBind区别：Bind返回错误会加入400错误码到http头

> gorm所支持的回调方法：
BeforeSave、BeforeCreate、AfterCreate、AfterSave
更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
删除：BeforeDelete、AfterDelete
查询：AfterFind

```
// User 模型，改变 description 字段的数据类型为 `text`
db.Model(&User{}).ModifyColumn("description", "text")

// User 模型，删除  description 字段
db.Model(&User{}).DropColumn("description")

// 为 `name` 字段建立一个名叫 `idx_user_name` 的索引
db.Model(&User{}).AddIndex("idx_user_name", "name")

// 为 `name`, `age` 字段建立一个名叫 `idx_user_name_age` 的索引
db.Model(&User{}).AddIndex("idx_user_name_age", "name", "age")

// 添加一条唯一索引
db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")

// 为多个字段添加唯一索引
db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "age")

// 移除索引
db.Model(&User{}).RemoveIndex("idx_user_name")

// 添加主键
// 第一个参数 : 主键的字段
// 第二个参数 : 目标表的 ID
// 第三个参数 : ONDELETE
// 第四个参数 : ONUPDATE
db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")

db.Model(&User{}).RemoveForeignKey("city_id", "cities(id)")
```