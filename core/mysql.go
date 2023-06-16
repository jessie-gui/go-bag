package core

import (
	_ "database/sql"
	"fmt"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// MysqlOption 数据库配置项。
type MysqlOption func(o *Mysql)

// Mysql 数据库对象。
type Mysql struct {
	user    string
	pwd     string
	address string
	dbBase  string
	dbGms   string
}

// SetUser 配置用户名。
func SetUser(user string) MysqlOption {
	return func(o *Mysql) {
		o.user = user
	}
}

// SetPwd 配置密码。
func SetPwd(pwd string) MysqlOption {
	return func(o *Mysql) {
		o.pwd = pwd
	}
}

// SetAddress 配置连接地址。
func SetAddress(address string) MysqlOption {
	return func(o *Mysql) {
		o.address = address
	}
}

// SetDbBase 配置数据库。
func SetDbBase(dbBase string) MysqlOption {
	return func(o *Mysql) {
		o.dbBase = dbBase
	}
}

// NewMysql 新建数据库对象。
func NewMysql(opts ...MysqlOption) *Mysql {
	opt := &Mysql{}
	for _, o := range opts {
		o(opt)
	}

	return opt
}

// Init 初始化数据库对象。
func (m *Mysql) Init() *gorm.DB {
	dsnBase := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.user, m.pwd, m.address, m.dbBase)
	dbBase, err := gorm.Open(mysql.Open(dsnBase), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		log.Fatal("connect DbBase failed", err)
	}

	log.Println("MYSQL初始化完成!")

	return dbBase
}
