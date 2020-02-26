package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// table foods
type Food struct {
	Id         int32
	Name       string
	Price      float32
	TypeId     int32
	CreateTime int64 `gorm:"column:createtime"`
}

// connMysql connects Mysql with dsn
func connMysql() *gorm.DB{
	// data source name: username:password@protocol(address)/dbname?param=value
	username := "root"
	password := "123"
	address := "127.0.0.1:3306"
	dbname := "test"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, address, dbname, timeout)

	// connect mysql
	var db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	return db
}

// InsDelUpd operates on row-level data
func InsDelUpd(op string, id int32, name string, price float32, typeId int32, createTime int64) {
	// connect mysql
	db := connMysql()
	defer db.Close()

	switch op{
	case "insert":
		food := &Food{
			id,
			name,
			price,
			typeId,
			createTime,
		}
		db.Create(food)
	case "delete":
		food := &Food{
			id,
			name,
			price,
			typeId,
			createTime,
		}
		db.Delete(&food)
	case "update":
		food := &Food{Id:id}
		db.Model(&food).Update(Food{Name:name, Price:price, TypeId:typeId, CreateTime:createTime})
	}
}

// Select returns query results
func Select(table string, columns string, condition string) string {
	// connect mysql
	db := connMysql()
	defer db.Close()

	var foods []Food
	db.Where(condition).Select(columns).Find(&foods)
	response := fmt.Sprintf("%v", foods)
	return response
}

// ExecSql executes SQL statement
func ExecSql(sql string) {
	db := connMysql()
	defer db.Close()

	db.Exec(sql)
}