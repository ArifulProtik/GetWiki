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
func (con *stroage) SearchArticle(searchtext string) ([]model.Article, error) {
	var artciles []model.Article
	con.db.Debug().Where("title like ?", "%"+searchtext+"%").Find(&artciles)
	if errs := con.db.GetErrors(); len(errs) > 0 {
		err := errs[0]
		return artciles, err
	}
	return artciles, nil
}
