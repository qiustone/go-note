### 连接到数据库
GORM 官方支持的数据库类型有: `Mysql`,`PostgreSQL`,`SQLite`,`SQL Server`

<br/>

### MySQL
```go
import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
```
> 注意: 想要正确处理 `time.Time`, 您需要带上 parseTime 参数，要支持完整的UTF-8编码，您需要将 charset=utf8 更改为 `charset=utf8mb4`

MySQL 驱动程序提供了一些高级配置可以在初始化过程中使用，例如：
```go
db, err := gorm.Open(mysql.New(mysql.Config{
    DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
    DefaultStringSize: 256, // string 类型字段的默认长度
    DisableDatetimePrecision: true, // 禁用datetime精度, MySQL5.6之前的数据库不支持
    DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式, MySQL5.7之前的数据库和 MariaDB 不支持重命名索引
    DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
    SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
}), &gorm.Config{})
```