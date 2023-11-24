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

func TestServices_UserSignup(t *testing.T) {
	type args struct {
		nu model.UserSignup
	}
	tests := []struct {
		name             string
		args             args
		want             model.User
		wantErr          bool
		mockRepoResponse func() (model.User, error)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				nu: model.UserSignup{
					UserName: "hdhhdhu",
					Email:    "shjswjjs@gmail,com",
					Password: "12345",
				},
			},
			want: model.User{
				UserName:     "hdhhdhu",
				Email:        "shjswjjs@gmail,com",
				PasswordHash: "hashed passs",
			},
			wantErr: false,
			mockRepoResponse: func() (model.User, error) {
				return model.User{
					UserName:     "hdhhdhu",
					Email:        "shjswjjs@gmail,com",
					PasswordHash: "hashed passs",
				}, nil
			},
		},

		{
			name: "failure",
			args: args{
				nu: model.UserSignup{
					UserName: "hdhhdhu",
					Email:    "shjswjjs@gmail,com",
					Password: "12345",
				},
			},
			want:    model.User{},
			wantErr: true,
			mockRepoResponse: func() (model.User, error) {
				return model.User{}, errors.New("signin failed")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)
			mockCache := redispack.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo, mockCache)
			got, err := s.UserSignup(tt.args.nu)
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}
