package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
//DBConn Is Used To Connect To Databases 
var DBConn *gorm.DB