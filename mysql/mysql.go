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

// dataSourceName returns dsn (data source name).
func dataSourceName() string {
	username := "root"
	password := "123"
	address := "127.0.0.1:3306"
	dbname := "test"
	timeout := "10s"
	// dsn: username:password@protocol(address)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, address, dbname, timeout)
	return dsn
}

func Insert(table string, id int32, name string, price float32, typeId int32, createTime int64) {
	// connect mysql
	dsn := dataSourceName()
	var db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer db.Close()

	food := &Food{
		id,
		name,
		price,
		typeId,
		createTime,
	}
	db.Table(table).Create(food)
}

func Delete(table string, id int32, name string, price float32, typeId int32, createTime int64) {
	// connect mysql
	dsn := dataSourceName()
	var db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer db.Close()

	food := &Food{
		id,
		name,
		price,
		typeId,
		createTime,
	}
	db.Table(table).Delete(&food)
}

func Update(table string, id int32, name string, price float32, typeId int32, createTime int64) {
	// connect mysql
	dsn := dataSourceName()
	var db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer db.Close()

	food := &Food{Id:id}
	db.Model(&food).Update(Food{Name:name, Price:price, TypeId:typeId, CreateTime:createTime})
}

// Select returns query results
func Select(table string, columns string, condition string) string {
	// connect mysql
	dsn := dataSourceName()
	var db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer db.Close()

	var foods []Food
	var sql = "SELECT " + columns + " FROM " + table
	if condition != "" {
		sql += " WHERE " + condition
	}

	db.Raw(sql).Scan(&foods)
	response := fmt.Sprintf("%v",foods)
	return response
}

// ExecSql executes SQL statement
func ExecSql(sql string) {
	dsn := dataSourceName()
	db ,err:= gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer db.Close()

	db.Exec(sql)
}