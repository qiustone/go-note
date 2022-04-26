### 什么是ORM ?
----
ORM = Object Relational Mapping (对象关系映射)  

<br>

### 在Golang中的对应关系
----
* 数据表 <---> 结构体   
* 数据行 <---> 结构体实例  
* 字段   <---> 结构体字段  

<br>

### ORM优缺点
----
优点：  
* 提高开发效率

缺点：
* 牺牲执行性能
* 牺牲灵活性
* 弱化SQL能力

<br>

### GORM 特性
---
* 全功能 ORM
* 关联（Has One, Has Many, Belongs To, Many To Many, 多态, 单表继承）
* Hooks (Before/After Create/Save/Update/Delete/Find)
* 支持 Preload、Joins 的预加载
* 事务, 嵌套事务, Save Point, Rollback To Save Point
* Context、预编译模式、DryRun 模式
* 批量插入, FindInBatches, Find/Create with Map, 使用SQL表达式、Context Valuer 进行 CRUD
* SQL构建器, Upsert, 数据库锁, Optimizer/Index/Comment Hint, 命名参数，子查询
* 复合主键, 索引, 约束
* Auto Migration
* 自定义Logger
* 灵活的可扩展插件API: Database Resolver (多数据库, 读写分离)、Prometheus

### GORM 安装
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```