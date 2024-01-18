package infrustructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func Init() SqlHandler {
	dsn := "root:root@tcp(127.0.0.1:306)/goarch?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB Error")
	}
	sh := new(SqlHandler)
	sh.db = db
	return *sh
}

func (obj *SqlHandler) create(model interface{}) {
	obj.db.Create(&model)
}

func (obj *SqlHandler) delete(model interface{}, id int) {
	obj.db.Delete(&model, id)
}

func (obj *SqlHandler) find(query interface{}) {
	obj.db.Find(query)
}
