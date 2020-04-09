package db

import "getwiki/model"

func (con *stroage) CreateArticle(artcile *model.Article) (err error) {
	con.db.Debug().Create(artcile)
	if errs := con.db.GetErrors(); len(errs) > 0 {
		err := errs[0]
		return err
	}
	return nil
}
