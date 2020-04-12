package db

import "getwiki/model"

func (con *stroage) CreateArticle(artcile *model.Article) (err error) {
	con.db.Debug().Model(model.Article{}).Create(&artcile)
	if errs := con.db.GetErrors(); len(errs) > 0 {
		err := errs[0]
		return err
	}
	return nil
}
func (con *stroage) SearchArticle(searchtext string) ([]model.Article, error) {
	var artciles []model.Article
	con.db.Debug().Where("title ilike ?", "%"+searchtext+"%").Find(&artciles)
	if errs := con.db.GetErrors(); len(errs) > 0 {
		err := errs[0]
		return artciles, err
	}
	return artciles, nil
}
func (con *stroage) Getall() ([]model.Article, error) {
	var allartcl []model.Article
	err := con.db.Debug().Model(&model.Article{}).Find(&allartcl).Error
	if err != nil {
		return allartcl, err

	}
	return allartcl, nil
}
func (con *stroage) Getbyid(id uint64) (model.Article, error) {
	var single model.Article
	err := con.db.Debug().Model(&model.Article{}).Where("id = ?", id).Take(&single).Error
	if err != nil {
		return single, err
	}
	return single, nil
}
