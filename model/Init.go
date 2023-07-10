// 这里是我们使用sqlite的部分

package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

//type ImageTable struct {
//	gorm.Model
//	ID       uint64 `gorm:"primarykey"`
//	Filename string `gorm:"type:varchar(100)"`
//	Data     []byte `gorm:"type:blob"`
//	Length   int64
//}

var (
	db  *gorm.DB
	err error
)

func InitDao() {
	// this should do for init ?
	os.Mkdir("data", os.ModePerm)
	db, err = gorm.Open(sqlite.Open(filepath.Join("data", "image.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Choice{})
	// db.Raw("create index index_length on image_tables (id,length);")
	// db.Raw("create index index_filename on image_tables (id,filename);")
}
