package repository

import (
	"project/internal/model"
)

func (r *Repo) CreateJob(j model.Job) (model.Job, error) {
	err := r.db.Create(&j).Error
	if err != nil {
		return model.Job{}, err
	}
	return j, nil
}

func (r *Repo) GetJobs(id int) ([]model.Job, error) {
	var m []model.Job

	tx := r.db.Where("uid = ?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil

}

func (r *Repo) GetAllJobs() ([]model.Job, error) {
	var s []model.Job
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *Repo) GetJobId(id uint64) (model.Job, error) {
	var j model.Job
	tx := r.db.Preload("Qualifications").
		Preload("Shifts").
		Preload("JobTypes").
		Preload("JobLocations").
		Preload("Technology_stack").
		Preload("WorkMode").Where("ID=?", id)
	err := tx.Find(&j).Error
	if err != nil {
		return model.Job{}, err
	}

	return j, nil
}

func (r *Repo) ApplyJob_Repository(id uint64) (model.Job, error) {

	var job model.Job
	tx := r.db.
		Preload("JobLocations").
		Preload("TechnologyStack").
		Preload("WorkModes").
		Preload("Qualifications").
		Preload("Shifts").
		Preload("JobTypes").
		Where("ID = ?", id)

	err := tx.First(&job).Error

	if err != nil {
		return model.Job{}, err
	}

	return job, nil

}
