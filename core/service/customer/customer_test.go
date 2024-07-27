package customer

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"order-service/core/entity"
	"order-service/data_source/mysql_datasource"
	"order-service/shared/mock/repositories"
	"testing"
)

func TestGetCustomerByUId(t *testing.T) {
	Convey("GetCustomerByUid", t, func() {

		var (
			ctrl        = gomock.NewController(t)
			mockRepo    = repositories.NewMockCustomerRepository(ctrl)
			err         = errors.New("error")
			ctx         = context.Background()
			customerUid = "1f0d352b-d635-4a09-b7bc-4fb490e1682d"
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error customer not found", func() {
			mockRepo.EXPECT().GetCustomerByUid(ctx, &customerUid).Return(nil, err).AnyTimes()
			res, err := service.GetCustomer(ctx, &customerUid)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})
		Convey("success customer found", func() {
			mockRepo.EXPECT().GetCustomerByUid(ctx, &customerUid).Return(&entity.Customer{}, nil).AnyTimes()
			res, err := service.GetCustomer(ctx, &customerUid)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})
	})
}

func TestGetAllCustomer(t *testing.T) {
	Convey("GetAllCustomer", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockCustomerRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error customer not found", func() {
			mockRepo.EXPECT().GetAllCustomer(ctx, mysql_datasource.QueryOption{Limit: 10}).Return(nil, err).AnyTimes()
			res, err := service.GetAllCustomer(ctx, &entity.RequestGetList{Limit: 10})
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})
		Convey("success customer found", func() {
			mockRepo.EXPECT().GetAllCustomer(ctx, mysql_datasource.QueryOption{Limit: 10}).Return(&entity.ResponseGetAllCustomer{}, nil).AnyTimes()
			res, err := service.GetAllCustomer(ctx, &entity.RequestGetList{Limit: 10})
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})
	})
}

func TestDeleteCustomer(t *testing.T) {
	Convey("DeleteCustomer", t, func() {

		var (
			ctrl        = gomock.NewController(t)
			mockRepo    = repositories.NewMockCustomerRepository(ctrl)
			err         = errors.New("error")
			ctx         = context.Background()
			customerUid = "1f0d352b-d635-4a09-b7bc-4fb490e1682d"
		)

		service := New(Opts{Repo: mockRepo})

		propertyMap := []map[string]interface{}{
			{
				"key":      "uid",
				"operator": "=",
				"value":    customerUid,
			},
		}
		Convey("error customer not found", func() {
			mockRepo.EXPECT().DeleteCustomer(ctx, &entity.Customer{}, propertyMap).Return(err).AnyTimes()
			err := service.DeleteCustomer(ctx, &customerUid)
			So(err, ShouldNotBeNil)
		})
		Convey("success customer found", func() {
			mockRepo.EXPECT().DeleteCustomer(ctx, &entity.Customer{}, propertyMap).Return(nil).AnyTimes()
			err := service.DeleteCustomer(ctx, &customerUid)
			So(err, ShouldBeNil)
		})
	})
}

func TestCreateCustomer(t *testing.T) {
	Convey("CreateCustomer", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockCustomerRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
			email    = "test@gmail.com"
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error customer not found", func() {
			mockRepo.EXPECT().GetCustomerByEmail(ctx, &email).Return(nil, err).AnyTimes()
			mockRepo.EXPECT().CreateCustomer(ctx, &entity.Customer{Email: email}).Return(err).AnyTimes()
			err := service.CreateCustomer(ctx, &entity.RequestCreateCustomer{Email: email})
			So(err, ShouldNotBeNil)
		})
		Convey("success customer found", func() {
			mockRepo.EXPECT().GetCustomerByEmail(ctx, &email).Return(nil, err).AnyTimes()
			mockRepo.EXPECT().CreateCustomer(ctx, &entity.Customer{Email: email}).Return(nil).AnyTimes()
			err := service.CreateCustomer(ctx, &entity.RequestCreateCustomer{Email: email})
			So(err, ShouldBeNil)
		})
	})
}

func TestUpdateCustomer(t *testing.T) {
	Convey("UpdateCustomer", t, func() {

		var (
			ctrl        = gomock.NewController(t)
			mockRepo    = repositories.NewMockCustomerRepository(ctrl)
			err         = errors.New("error")
			ctx         = context.Background()
			customerUid = "1f0d352b-d635-4a09-b7bc-4fb490e1682d"
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error customer not found", func() {
			mockRepo.EXPECT().UpdateCustomer(ctx, &entity.Customer{MetaData: entity.MetaData{UID: uuid.MustParse(customerUid)}}, map[string]interface{}{"email": "", "name": "", "address": "", "mobile": ""}).Return(err).AnyTimes()
			err := service.UpdateCustomer(ctx, &customerUid, &entity.RequestUpdateCustomer{})
			So(err, ShouldNotBeNil)
		})
		Convey("success customer found", func() {
			mockRepo.EXPECT().UpdateCustomer(ctx, &entity.Customer{MetaData: entity.MetaData{UID: uuid.MustParse(customerUid)}}, map[string]interface{}{"email": "", "name": "", "address": "", "mobile": ""}).Return(nil).AnyTimes()
			err := service.UpdateCustomer(ctx, &customerUid, &entity.RequestUpdateCustomer{})
			So(err, ShouldBeNil)
		})
	})
}
