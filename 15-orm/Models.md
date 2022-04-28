### 模型定义 (Declaring Models)
---
模型是标准的 Struct, 由Go的基本数据类型、实现了 `Scanner` 和 `Valuer` 接口的自定义类型及其指针或别名组成

```go
type User struct  {
    ID           uint
    Name         string
    Email        *string
    Age          uint8
    Birthday     *time.Time
    MemberNumber sql.NullString
    ActivatedAt  sql.NullTime
    CreatedAt    time.Time
    UpateAt      time.Time
}
```

<br>

### 约定 (Conventions)
---
GORM 更倾向于约定，而不是配置。  
默认情况下，GORM 使用 `ID` 作为主键,  
使用结构体名的 `snake_cases` 形式作为表名,  
`snake_case` 作为列名,  
并使用 `CreatedAt`, `UpdatedAt` 去追踪创建和更新时间

遵循 GORM 已有的约定，可以减少您的配置和代码量。  
如果约定不符合您的需求，GORM 允许您自定义配置它们  

<br>

### gorm.Model
---
GORM 内置了一个 `gorm.Model` 结构体， 其中包括字段`ID`,`CreatedAt`,`UpdatedAt`,`DeletedAt`
```go
type Model struct {
    ID       uint         `gorm:"primaryKey"`
    CreateAt time.Time
    UpdateAt time.Time
    DeleteAt gorm.DeleteAt `gorm:"index"`
}
```
您可以将它嵌入到您的结构体中，以包含这几个字段

<br>

### 字段级权限(Field-Level Permission)
---
可导出的字段在使用GORM进行CRUD时拥有全部的权限，  
此外，GORM允许您使用`标签`控制字段级别的权限。  
这样您就可以让一个字段的权限是`只读`、`只写`、`只创建`、`只更新`或者被`忽略`  

> 注意:使用 GORM Migrator 创建表时, 不会创建被忽略的字段

```go
type User struct {
    Name string `gorm:"<-:create"`   // allow read and create
    Name string `gorm:"<-:update"`   // allow read and update
    Name string `gorm:"<-"`          // allow read and write (create and upate)
    Name string `gorm:"<-:false"`    // allow read, disable write permission
    Name string `gorm:"->"`          // readonly (disable write permission)
    Name string `gorm:"->;<-create"` // allow read and create
    Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
    Name string `gorm:"-"`           // ignore this field when write and read with struct
    Name string `gorm:"-:all"`       // ignore this field when write, read and migrate with struct
    Name string `gorm:"-:migration"` // ignore this field when migrate with struct
}
```

<br>

### 创建/更新时间追踪（纳秒、毫秒、秒、Time）
----
`GORM` 约定使用 `CreatedAt`、`UpdatedAt` 追踪创建/更新时间。如果您定义了这种字段，`GORM` 在创建、更新时会自动填充当前时间

要使用不同名称的字段，您可以配置 `autoCreateTime`、`autoUpdateTime` 标签  

如果您想要保存 UNIX（毫/纳）秒时间戳，而不是 time，您只需简单地将 time.Time 修改为 int 即可
```go
type User struct {
    CreatedAt time.Time  // 在创建时，如果该字段值为零值, 则使用当前时间填充
    UpdatedAt int        // 在创建时该字段值为零值或者再更新时，使用当前时间戳秒数填充
    Updated   int64 `gorm:"autoUpdateTime:nano"`  // 使用时间错纳秒数填充更新时间
    Updated   int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
    Created   int64 `gorm:"autoCreateTime"`       // 使用时间戳秒数填充创建时间
}
```

<br>

### 嵌入结构体
---
对于匿名字段, GORM 会将其字段包含在父结构体中, 例如:
```go
type User struct {
    gorm.Model
    Name string
}

//等效于

type User struct {
    ID        uint           `gorm:"primaryKey"` 
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
    Name      string
}
```
对于正常的结构体字段，你也可以通过标签 embedded 将其嵌入, 例如:
```go
type Author struct {
    Name  string
    Email string
}

type Blog struct {
    ID      int
    Author  Author `gorm:"embedded"`
    Upovtes int32
}

//等效于

type Blog struct {
    ID      int
    Name    string
    Email   string
    Upvotes int32
}
```
并且，您可以使用标签 `embeddedPrefix` 来为`db`中的字段名添加前缀, 例如:
```go
type Blog struct {
    ID      int
    Author  Author `gorm:"embedded;embeddedPrefix:author_"`
    Upvotes int32
}

//等效于

type Blog struct {
    ID          int
    AuthorName  string
    AuthorEmail string
    Upvotes     int32
}
```

<br>

### 字段标签
---
声明model时，tag 是可选的, GORM 支持以下 tag: tag名大小写不敏感，但建议使用camelCase风格

| 标签名 | 说明 |
| :-----| :---- |
| column | 指定 db 列名 |
| type | 列数据类型，推荐使用兼容性好的通用类型，例如: 所有数据库都支持`bool`、`int`、`uint`、`float`、`string`、`time`、`bytes` 并且可以和其他标签一起使用, 例如: `not null`、`size`、`autoIncrement`... 像`varbinary(8)`这样指定数据库数据类型也是支持的。在使用指定数据库类型时，它需要是完整的数据库数据类型，如: `MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT` |
| size | 指定列大小，例如：`size:256` |
| primaryKey | 指定列为主键 |
| unique | 指定列为唯一 |
| default | 指定列的默认值 |
| precision | 指定列的精度 |
| not null | 指定列为 NOT NULL |
| autoIncrement | 指定列为自动增长 |
| autoIncrementIncrement | 自动步长，控制连续记录之间的间隔 |
| embedded | 嵌套字段 |
| embeddedPrefix | 嵌入字段的列名前缀 |
| autoCreateTime | 创建时追踪当前时间，对于 int 字段，它会追踪秒级时间戳，您可以使用nano/milli 来追踪纳秒毫秒时间戳，例如: autoCreateTime:"nano" |
|  |  |
|  |  |
|  |  |