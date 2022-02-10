package database

//
//import (
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jinzhu/gorm"
//)
//
//type SQLHandler struct {
//	Conn *gorm.DB
//}
//
//func NewSQLHandler(connInfo string) (*SQLHandler, error) {
//	db, err := gorm.Open("mysql", connInfo)
//	if err != nil {
//		return nil, err
//	}
//	db.LogMode(true)
//	return &SQLHandler{
//		Conn: db,
//	}, nil
//}
//
//func (h *SQLHandler) Close() error {
//	return h.Conn.Close()
//}
