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

### gorm.Model