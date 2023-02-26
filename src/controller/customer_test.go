package controller

import (
	"context"
	"reflect"
	"templatebe/lib/loggers"
	"templatebe/src/config"
	"templatebe/src/model"
	"templatebe/src/repository/mockrepository"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCustomerController_CreateCustomer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cfg := config.New()
	logger := loggers.NewZerolog(cfg)
	mockCustomerRepo := mockrepository.NewMockCustomerRepository(mockCtrl)

	type args struct {
		ctx context.Context
		req model.CreateCustomerRequest
	}
	tests := []struct {
		name    string
		con     *CustomerController
		args    args
		want    *model.CreateCustomerResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			con: &CustomerController{
				logger:             logger,
				customerRepository: mockCustomerRepo,
			},
			args: args{
				ctx: context.Background(),
				req: model.CreateCustomerRequest{
					ID:   1,
					Name: "Boss",
				},
			},
			want: &model.CreateCustomerResponse{
				ID: 1,
			},
			wantErr: false,
			mock: func() {
				mockCustomerRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
			got, err := tt.con.CreateCustomer(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerController.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerController.CreateCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}
