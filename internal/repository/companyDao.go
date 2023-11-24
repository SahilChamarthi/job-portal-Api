package repository

import "project/internal/model"

func (r *Repo) CreateCompany(u model.Company) (model.Company, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return model.Company{}, err
	}
	return u, nil
}

func (r *Repo) GetAllCompany() ([]model.Company, error) {
	var s []model.Company
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *Repo) GetCompany(id int) (model.Company, error) {
	var m model.Company

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return model.Company{}, err
	}
	return m, nil

}
