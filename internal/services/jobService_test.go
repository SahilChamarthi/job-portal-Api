package services

import (
	"errors"
	"project/internal/model"
	redispack "project/internal/redisPack"
	"project/internal/repository"
	"reflect"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func TestServices_JobCreate(t *testing.T) {
	type args struct {
		nj model.CreateJob
		id uint64
	}
	tests := []struct {
		name             string
		args             args
		want             model.Job
		wantErr          bool
		mockRepoResponse func() (model.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{nj: model.CreateJob{
				JobTitle:  "gwyug",
				JobSalary: 514212,
			},
				id: 15,
			},
			want: model.Job{
				JobTitle:  "gwyug",
				JobSalary: 514212,

				CompanyID: 15,
			},
			wantErr: false,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{
					JobTitle:  "gwyug",
					JobSalary: 514212,

					CompanyID: 15,
				}, nil

			},
		},

		{
			name: "failure",
			args: args{
				nj: model.CreateJob{},
				id: 15,
			},
			want:    model.Job{},
			wantErr: true,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{}, errors.New("failed to crerate job")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)
			mockCache := redispack.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateJob(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo, mockCache)
			//got, err := tt.s.CompanyCreate(tt.args.nc)
			got, err := s.JobCreate(tt.args.nj, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.JobCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.JobCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_GetJobsByCompanyId(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name             string
		args             args
		want             []model.Job
		wantErr          bool
		mockRepoResponse func() ([]model.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 12,
			},
			want: []model.Job{
				{
					JobTitle:  "hcwc",
					JobSalary: 25778,
					CompanyID: 12,
				},
				{
					JobTitle:  "sajchbuc",
					JobSalary: 25778,
					CompanyID: 13,
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]model.Job, error) {
				return []model.Job{
					{
						JobTitle:  "hcwc",
						JobSalary: 25778,
						CompanyID: 12,
					},
					{
						JobTitle:  "ygyudsgf",
						JobSalary: 25778,
						CompanyID: 13,
					},
				}, nil
			},
		},
		{
			name: "failure",
			args: args{
				id: 12,
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]model.Job, error) {
				return nil, errors.New("job not found")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)
			mockCache := redispack.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetJobs(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo, mockCache)
			//got, err := tt.s.CompanyCreate(tt.args.nc)
			got, err := s.GetJobsByCompanyId(tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Services.GetJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.GetJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_FetchAllJobs(t *testing.T) {
	tests := []struct {
		name             string
		want             []model.Job
		wantErr          bool
		mockRepoResponse func() ([]model.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "sucesss",
			want: []model.Job{
				{
					JobTitle:  "hcwc",
					JobSalary: 25778,
					CompanyID: 12,
				},
				{
					JobTitle:  "ygyudsgf",
					JobSalary: 25778,
					CompanyID: 13,
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]model.Job, error) {
				return []model.Job{
					{
						JobTitle:  "hcwc",
						JobSalary: 25778,
						CompanyID: 12,
					},
					{
						JobTitle:  "ygyudsgf",
						JobSalary: 25778,
						CompanyID: 13,
					},
				}, nil
			},
		},
		{
			name:    "failure",
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]model.Job, error) {
				return nil, errors.New("fail to found jobs")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)
			mockCache := redispack.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetAllJobs().Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo, mockCache)
			got, err := s.FetchAllJobs()
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.FetchAllJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.FetchAllJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_Getjobid(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name             string
		args             args
		want             model.Job
		wantErr          bool
		mockRepoResponse func() (model.Job, error)
	}{
		// TODO: Add test cases.

		{
			name: "success",
			args: args{
				id: 15,
			},
			want: model.Job{

				JobTitle:  "eyhe",
				JobSalary: 162522,
				CompanyID: 17,
			},
			wantErr: false,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{

					JobTitle:  "eyhe",
					JobSalary: 162522,
					CompanyID: 17,
				}, nil
			},
		},
		{
			name: "failure",
			args: args{
				id: 12,
			},
			want:    model.Job{},
			wantErr: true,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{}, errors.New("job not found")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)
			mockCache := redispack.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetJobId(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo, mockCache)
			got, err := s.Getjobid(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.Getjobid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.Getjobid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_ApplyJob_Service(t *testing.T) {
	type args struct {
		ja []model.JobApplication
	}
	tests := []struct {
		name              string
		args              args
		want              []model.ApprovedApplication
		wantErr           bool
		mockRepoResponse  func() (model.Job, error)
		mockCacheResponse func() (model.Job, error)
	}{

		{
			name: "successs",
			args: args{
				ja: []model.JobApplication{
					{
						JobId:          1,
						Name:           "bumesh",
						Gmail:          "bumesh@gmail.com",
						Age:            23,
						Phone:          9018373973,
						JobTitle:       "software testing",
						ExpectedSalary: 26000,
						NoticePeriod:   30,
						Experience:     3,

						Qualifications:   []uint{1, 2},
						Shift:            []uint{1},
						JobType:          []uint{1},
						JobLocations:     []uint{1},
						Technology_stack: []uint{1, 2},
						WorkMode:         []uint{1, 2},
					},
					{
						JobId:          1,
						Name:           "seenu",
						Gmail:          "seenu@gmail.com",
						Age:            24,
						Phone:          9018373979,
						JobTitle:       "software testing",
						ExpectedSalary: 27000,
						NoticePeriod:   0,
						Experience:     3,

						Qualifications:   []uint{1, 2},
						Shift:            []uint{1},
						JobType:          []uint{1},
						JobLocations:     []uint{1},
						Technology_stack: []uint{1, 2},
						WorkMode:         []uint{1, 2},
					},
				},
			},

			want: []model.ApprovedApplication{
				{
					Name:  "bumesh",
					Gmail: "bumesh@gmail.com",
					Phone: 9018373973,
				},
				{
					Name:  "seenu",
					Gmail: "seenu@gmail.com",
					Phone: 9018373979,
				},
			},
			wantErr: false,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{
					ID:                 1,
					JobTitle:           "Java Developper",
					Description:        "oppartuinity on java develop",
					CompanyID:          3,
					Min_NoticePeriod:   0,
					Max_NoticePeriod:   60,
					Budget:             60000,
					Minimum_Experience: 1,
					Maximum_Experience: 3,
					JobLocations:       []model.JobLocation{{ID: 1, Name: "banglore"}, {ID: 2, Name: "hyderabad"}},
					TechnologyStack:    []model.Technology{{ID: 1, Name: "java"}, {ID: 2, Name: "sql"}},
					WorkModes:          []model.WorkMode{{ID: 1, Name: "remote"}, {ID: 2, Name: "work from office"}},
					Qualifications:     []model.Qualification{{ID: 1, Name: "B.Tech"}, {ID: 2, Name: "M.Tech"}},
					Shifts:             []model.Shift{{ID: 1, Name: "day"}, {ID: 2, Name: "night"}},
					JobTypes:           []model.JobType{{ID: 1, Name: "full time"}, {ID: 2, Name: "contract"}},
				}, nil
			},
			mockCacheResponse: func() (model.Job, error) {
				return model.Job{
					ID:                 1,
					JobTitle:           "Java Developper",
					Description:        "oppartuinity on java develop",
					CompanyID:          3,
					Min_NoticePeriod:   0,
					Max_NoticePeriod:   60,
					Budget:             60000,
					Minimum_Experience: 1,
					Maximum_Experience: 3,
					JobLocations:       []model.JobLocation{{ID: 1, Name: "banglore"}, {ID: 2, Name: "hyderabad"}},
					TechnologyStack:    []model.Technology{{ID: 1, Name: "java"}, {ID: 2, Name: "sql"}},
					WorkModes:          []model.WorkMode{{ID: 1, Name: "remote"}, {ID: 2, Name: "work from office"}},
					Qualifications:     []model.Qualification{{ID: 1, Name: "B.Tech"}, {ID: 2, Name: "M.Tech"}},
					Shifts:             []model.Shift{{ID: 1, Name: "day"}, {ID: 2, Name: "night"}},
					JobTypes:           []model.JobType{{ID: 1, Name: "full time"}, {ID: 2, Name: "contract"}},
				}, nil

			},
		},
		// {

		// 	name: "error in db",
		// 	args: args{
		// 		ja: []model.JobApplication{
		// 			{
		// 				JobId:          1,
		// 				Name:           "bumesh",
		// 				Gmail:          "bumesh@gmail.com",
		// 				Age:            23,
		// 				Phone:          9018373973,
		// 				JobTitle:       "software testing",
		// 				ExpectedSalary: 26000,
		// 				NoticePeriod:   30,
		// 				Experience:     3,

		// 				Qualifications:   []uint{1, 2},
		// 				Shift:            []uint{1},
		// 				JobType:          []uint{1},
		// 				JobLocations:     []uint{1},
		// 				Technology_stack: []uint{1, 2},
		// 				WorkMode:         []uint{1, 2},
		// 			},
		// 			{
		// 				JobId:          1,
		// 				Name:           "seenu",
		// 				Gmail:          "seenu@gmail.com",
		// 				Age:            24,
		// 				Phone:          9018373979,
		// 				JobTitle:       "software testing",
		// 				ExpectedSalary: 27000,
		// 				NoticePeriod:   0,
		// 				Experience:     3,

		// 				Qualifications:   []uint{1, 2},
		// 				Shift:            []uint{1},
		// 				JobType:          []uint{1},
		// 				JobLocations:     []uint{1},
		// 				Technology_stack: []uint{1, 2},
		// 				WorkMode:         []uint{1, 2},
		// 			},
		// 		},
		// 	},

		// 	want:    []model.ApprovedApplication{},
		// 	wantErr: true,
		// 	mockRepoResponse: func() (model.Job, error) {
		// 		return model.Job{}, errors.New("job not found raa")
		// 	},
		// },
		// {
		// 	name: "failing in experience",
		// 	args: args{
		// 		ja: []model.JobApplication{
		// 			{
		// 				JobId:          1,
		// 				Name:           "bumesh",
		// 				Gmail:          "bumesh@gmail.com",
		// 				Age:            23,
		// 				Phone:          9018373973,
		// 				JobTitle:       "software testing",
		// 				ExpectedSalary: 26000,
		// 				NoticePeriod:   30,
		// 				Experience:     3,

		// 				Qualifications:   []uint{1, 2},
		// 				Shift:            []uint{1},
		// 				JobType:          []uint{1},
		// 				JobLocations:     []uint{1},
		// 				Technology_stack: []uint{1, 2},
		// 				WorkMode:         []uint{1, 2},
		// 			},
		// 			{
		// 				JobId:          1,
		// 				Name:           "seenu",
		// 				Gmail:          "seenu@gmail.com",
		// 				Age:            24,
		// 				Phone:          9018373979,
		// 				JobTitle:       "software testing",
		// 				ExpectedSalary: 27000,
		// 				NoticePeriod:   0,
		// 				Experience:     5,

		// 				Qualifications:   []uint{1, 2},
		// 				Shift:            []uint{1},
		// 				JobType:          []uint{1},
		// 				JobLocations:     []uint{1},
		// 				Technology_stack: []uint{1, 2},
		// 				WorkMode:         []uint{1, 2},
		// 			},
		// 		},
		// 	},

		// 	want: []model.ApprovedApplication{
		// 		{
		// 			Name:  "seenu",
		// 			Gmail: "seenu@gmail.com",
		// 			Phone: 9018373979,
		// 		},
		// 	},
		// 	wantErr: false,
		// 	mockRepoResponse: func() (model.Job, error) {
		// 		return model.Job{
		// 			ID:                 1,
		// 			JobTitle:           "Java Developper",
		// 			Description:        "oppartuinity on java develop",
		// 			CompanyID:          3,
		// 			Min_NoticePeriod:   0,
		// 			Max_NoticePeriod:   60,
		// 			Budget:             60000,
		// 			Minimum_Experience: 5,
		// 			Maximum_Experience: 8,
		// 			JobLocations:       []model.JobLocation{{ID: 1, Name: "banglore"}, {ID: 2, Name: "hyderabad"}},
		// 			TechnologyStack:    []model.Technology{{ID: 1, Name: "java"}, {ID: 2, Name: "sql"}},
		// 			WorkModes:          []model.WorkMode{{ID: 1, Name: "remote"}, {ID: 2, Name: "work from office"}},
		// 			Qualifications:     []model.Qualification{{ID: 1, Name: "B.Tech"}, {ID: 2, Name: "M.Tech"}},
		// 			Shifts:             []model.Shift{{ID: 1, Name: "day"}, {ID: 2, Name: "night"}},
		// 			JobTypes:           []model.JobType{{ID: 1, Name: "full time"}, {ID: 2, Name: "contract"}},
		// 		}, nil
		// 	},
		// },
		// {
		// 	name: "failing in budget",
		// 	args: args{
		// 		ja: []model.JobApplication{
		// 			{
		// 				JobId:          1,
		// 				Name:           "bumesh",
		// 				Gmail:          "bumesh@gmail.com",
		// 				Age:            23,
		// 				Phone:          9018373973,
		// 				JobTitle:       "software testing",
		// 				ExpectedSalary: 80000,
		// 				NoticePeriod:   30,
		// 				Experience:     3,

		// 				Qualifications:   []uint{1, 2},
		// 				Shift:            []uint{1},
		// 				JobType:          []uint{1},
		// 				JobLocations:     []uint{1},
		// 				Technology_stack: []uint{1, 2},
		// 				WorkMode:         []uint{1, 2},
		// 			},
		// 			{
		// 				JobId:          1,
		// 				Name:           "seenu",
		// 				Gmail:          "seenu@gmail.com",
		// 				Age:            24,
		// 				Phone:          9018373979,
		// 				JobTitle:       "software testing",
		// 				ExpectedSalary: 90000,
		// 				NoticePeriod:   0,
		// 				Experience:     3,

		// 				Qualifications:   []uint{1, 2},
		// 				Shift:            []uint{1},
		// 				JobType:          []uint{1},
		// 				JobLocations:     []uint{1},
		// 				Technology_stack: []uint{1, 2},
		// 				WorkMode:         []uint{1, 2},
		// 			},
		// 		},
		// 	},

		// 	want: []model.ApprovedApplication{
		// 		{
		// 			Name:  "bumesh",
		// 			Gmail: "bumesh@gmail.com",
		// 			Phone: 9018373973,
		// 		},
		// 	},
		// 	wantErr: false,
		// 	mockRepoResponse: func() (model.Job, error) {
		// 		return model.Job{
		// 			ID:                 1,
		// 			JobTitle:           "Java Developper",
		// 			Description:        "oppartuinity on java develop",
		// 			CompanyID:          3,
		// 			Min_NoticePeriod:   0,
		// 			Max_NoticePeriod:   60,
		// 			Budget:             80000,
		// 			Minimum_Experience: 1,
		// 			Maximum_Experience: 3,
		// 			JobLocations:       []model.JobLocation{{ID: 1, Name: "banglore"}, {ID: 2, Name: "hyderabad"}},
		// 			TechnologyStack:    []model.Technology{{ID: 1, Name: "java"}, {ID: 2, Name: "sql"}},
		// 			WorkModes:          []model.WorkMode{{ID: 1, Name: "remote"}, {ID: 2, Name: "work from office"}},
		// 			Qualifications:     []model.Qualification{{ID: 1, Name: "B.Tech"}, {ID: 2, Name: "M.Tech"}},
		// 			Shifts:             []model.Shift{{ID: 1, Name: "day"}, {ID: 2, Name: "night"}},
		// 			JobTypes:           []model.JobType{{ID: 1, Name: "full time"}, {ID: 2, Name: "contract"}},
		// 		}, nil
		// 	},
		// },
		// {
		// 	name: "failing in notice period",
		// 	args: args{
		// 		ja: []model.JobApplication{
		// 			{
		// 				JobId:          1,
		// 				Name:           "bumesh",
		// 				Gmail:          "bumesh@gmail.com",
		// 				Age:            23,
		// 				Phone:          9018373973,
		// 				JobTitle:       "software testing",
		// 				ExpectedSalary: 26000,
		// 				NoticePeriod:   80,
		// 				Experience:     3,

		// 				Qualifications:   []uint{1, 2},
		// 				Shift:            []uint{1},
		// 				JobType:          []uint{1},
		// 				JobLocations:     []uint{1},
		// 				Technology_stack: []uint{1, 2},
		// 				WorkMode:         []uint{1, 2},
		// 			},
		// 			{
		// 				JobId:          1,
		// 				Name:           "seenu",
		// 				Gmail:          "seenu@gmail.com",
		// 				Age:            24,
		// 				Phone:          9018373979,
		// 				JobTitle:       "software testing",
		// 				ExpectedSalary: 27000,
		// 				NoticePeriod:   0,
		// 				Experience:     3,

		// 				Qualifications:   []uint{1, 2},
		// 				Shift:            []uint{1},
		// 				JobType:          []uint{1},
		// 				JobLocations:     []uint{1},
		// 				Technology_stack: []uint{1, 2},
		// 				WorkMode:         []uint{1, 2},
		// 			},
		// 		},
		// 	},

		// 	want: []model.ApprovedApplication{
		// 		{
		// 			Name:  "seenu",
		// 			Gmail: "seenu@gmail.com",
		// 			Phone: 9018373979,
		// 		},
		// 	},
		// 	wantErr: false,
		// 	mockRepoResponse: func() (model.Job, error) {
		// 		return model.Job{
		// 			ID:                 1,
		// 			JobTitle:           "Java Developper",
		// 			Description:        "oppartuinity on java develop",
		// 			CompanyID:          3,
		// 			Min_NoticePeriod:   0,
		// 			Max_NoticePeriod:   60,
		// 			Budget:             60000,
		// 			Minimum_Experience: 1,
		// 			Maximum_Experience: 3,
		// 			JobLocations:       []model.JobLocation{{ID: 1, Name: "banglore"}, {ID: 2, Name: "hyderabad"}},
		// 			TechnologyStack:    []model.Technology{{ID: 1, Name: "java"}, {ID: 2, Name: "sql"}},
		// 			WorkModes:          []model.WorkMode{{ID: 1, Name: "remote"}, {ID: 2, Name: "work from office"}},
		// 			Qualifications:     []model.Qualification{{ID: 1, Name: "B.Tech"}, {ID: 2, Name: "M.Tech"}},
		// 			Shifts:             []model.Shift{{ID: 1, Name: "day"}, {ID: 2, Name: "night"}},
		// 			JobTypes:           []model.JobType{{ID: 1, Name: "full time"}, {ID: 2, Name: "contract"}},
		// 		}, nil
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)
			mockCache := redispack.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ApplyJob_Repository(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			if tt.mockCacheResponse != nil {
				mockCache.EXPECT().CheckRedisKey(gomock.Any()).Return(tt.mockCacheResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo, mockCache)
			got, err := s.ApplyJob_Service(tt.args.ja)

			if (err != nil) != tt.wantErr {
				t.Errorf("Services.ApplyJob_Service() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.ApplyJob_Service() = %v, want %v", got, tt.want)
			}
		})
	}
}
