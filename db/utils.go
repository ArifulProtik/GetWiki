package db

import (
	"getwiki/model"
)

func (con *stroage) Migrator() {
	con.db.Debug().AutoMigrate(&model.Article{})
}
