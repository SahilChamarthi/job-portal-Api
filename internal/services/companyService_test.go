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

func TestServices_CompanyCreate(t *testing.T) {
	type args struct {
		nc model.CreateCompany
	}
	tests := []struct {
		name string
		//ser              *Services
		args             args
		want             model.Company
		wantErr          bool
		mockRepoResponse func() (model.Company, error)
	}{
		// TODO: Add test cases.

		{
			name: "error in creation",
			args: args{
				nc: model.CreateCompany{},
			},
			want:    model.Company{},
			wantErr: true,
			mockRepoResponse: func() (model.Company, error) {
				return model.Company{}, errors.New("error in accessing data from db")

			},
		},
		{
			name: "success",
			args: args{
				nc: model.CreateCompany{CompanyName: "accenture", Adress: "banglore ecospace", Domain: "finance"},
			},
			want:    model.Company{CompanyName: "accenture", Adress: "banglore ecospace", Domain: "finance"},
			wantErr: false,
			mockRepoResponse: func() (model.Company, error) {
				return model.Company{
					CompanyName: "accenture", Adress: "banglore ecospace", Domain: "finance",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)
			mockCache := redispack.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateCompany(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo, mockCache)
			//got, err := tt.s.CompanyCreate(tt.args.nc)
			got, err := s.CompanyCreate(tt.args.nc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.CompanyCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.CompanyCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_GetAllCompanies(t *testing.T) {
	tests := []struct {
		name             string
		want             []model.Company
		wantErr          bool
		mockRepoResponse func() ([]model.Company, error)
	}{
		// TODO: Add test cases.
		{
			name: "sucess",
			want: []model.Company{
				{
					CompanyName: "njfiodjf",
					Adress:      "hnufdshiufh",
					Domain:      "dhbsfudshfui",
				},
				{
					CompanyName: "dsfiuhdiufh",
					Adress:      "dfhniuhdiu",
					Domain:      "jhdsfiudsyf9u",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]model.Company, error) {
				return []model.Company{
					{
						CompanyName: "njfiodjf",
						Adress:      "hnufdshiufh",
						Domain:      "dhbsfudshfui",
					},
					{
						CompanyName: "dsfiuhdiufh",
						Adress:      "dfhniuhdiu",
						Domain:      "jhdsfiudsyf9u",
					},
				}, nil
			},
		},
		{
			name:    "failure",
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]model.Company, error) {
				return nil, errors.New("company are not there")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockAllInRepo(mc)
			mockCache := redispack.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetAllCompany().Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewServices(mockRepo, mockCache)
			got, err := s.GetAllCompanies()
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.GetAllCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.GetAllCompanies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_GetCompanyById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name             string
		args             args
		want             model.Company
		wantErr          bool
		mockRepoResponse func() (model.Company, error)
	}{
		{
			name:    "success",
			args:    args{id: 1},
			want:    model.Company{CompanyName: "asjgc", Adress: "wgcu", Domain: "jdjdjh"},
			wantErr: false,
			mockRepoResponse: func() (model.Company, error) {
				return model.Company{
					CompanyName: "asjgc", Adress: "wgcu", Domain: "jdjdjh",
				}, nil
			},
		},
		{
			name:    "failure",
			args:    args{id: 1},
			want:    model.Company{},
			wantErr: true,
			mockRepoResponse: func() (model.Company, error) {
				return model.Company{}, errors.New("company not found")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepo := repository.NewMockAllInRepo(mc)
			mockCache := redispack.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetCompany(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewServices(mockRepo, mockCache)
			got, err := s.GetCompanyById(tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Services.GetCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.GetCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}
