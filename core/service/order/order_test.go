package order

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

func TestGetOrderByUId(t *testing.T) {
	Convey("GetOrderByUid", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockOrderRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
			orderUid = "1f0d352b-d635-4a09-b7bc-4fb490e1682d"
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error order not found", func() {
			mockRepo.EXPECT().GetOrderByUid(ctx, &orderUid).Return(nil, err).AnyTimes()
			res, err := service.GetOrder(ctx, &orderUid)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})
		Convey("success order found", func() {
			mockRepo.EXPECT().GetOrderByUid(ctx, &orderUid).Return(&entity.Order{}, nil).AnyTimes()
			res, err := service.GetOrder(ctx, &orderUid)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})
	})
}

func TestGetAllOrder(t *testing.T) {
	Convey("GetAllOrder", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockOrderRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error order not found", func() {
			mockRepo.EXPECT().GetAllOrder(ctx, mysql_datasource.QueryOption{Limit: 10}).Return(nil, err).AnyTimes()
			res, err := service.GetAllOrder(ctx, &entity.RequestGetList{Limit: 10})
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})
		Convey("success order found", func() {
			mockRepo.EXPECT().GetAllOrder(ctx, mysql_datasource.QueryOption{Limit: 10}).Return(&entity.ResponseGetAllOrder{}, nil).AnyTimes()
			res, err := service.GetAllOrder(ctx, &entity.RequestGetList{Limit: 10})
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})
	})
}

func TestDeleteOrder(t *testing.T) {
	Convey("DeleteOrder", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockOrderRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
			orderUid = "1f0d352b-d635-4a09-b7bc-4fb490e1682d"
		)

		service := New(Opts{Repo: mockRepo})

		propertyMap := []map[string]interface{}{
			{
				"key":      "uid",
				"operator": "=",
				"value":    orderUid,
			},
		}
		Convey("error order not found", func() {
			mockRepo.EXPECT().DeleteOrder(ctx, &entity.Order{}, propertyMap).Return(err).AnyTimes()
			err := service.DeleteOrder(ctx, &orderUid)
			So(err, ShouldNotBeNil)
		})
		Convey("success order found", func() {
			mockRepo.EXPECT().DeleteOrder(ctx, &entity.Order{}, propertyMap).Return(nil).AnyTimes()
			err := service.DeleteOrder(ctx, &orderUid)
			So(err, ShouldBeNil)
		})
	})
}

func TestCreateOrder(t *testing.T) {
	Convey("CreateOrder", t, func() {

		var (
			ctrl             = gomock.NewController(t)
			mockRepo         = repositories.NewMockOrderRepository(ctrl)
			mockRepoCustomer = repositories.NewMockCustomerRepository(ctrl)
			err              = errors.New("error")
			ctx              = context.Background()
			goodsName        = "test"
			customerUid      = "1f0d352b-d635-4a09-b7bc-4fb490e1682d"
		)

		service := New(Opts{Repo: mockRepo, CustomerRepo: mockRepoCustomer})

		Convey("error order not found", func() {
			mockRepoCustomer.EXPECT().GetCustomerByUid(ctx, &customerUid).Return(nil, err).AnyTimes()
			mockRepo.EXPECT().CreateOrder(ctx, &entity.Order{GoodsName: goodsName}).Return(err).AnyTimes()
			err := service.CreateOrder(ctx, &entity.RequestCreateOrder{CustomerUId: customerUid, GoodsName: goodsName})
			So(err, ShouldNotBeNil)
		})
		Convey("success order found", func() {
			mockRepoCustomer.EXPECT().GetCustomerByUid(ctx, &customerUid).Return(&entity.Customer{}, nil).AnyTimes()
			mockRepo.EXPECT().CreateOrder(ctx, &entity.Order{GoodsName: goodsName}).Return(nil).AnyTimes()
			err := service.CreateOrder(ctx, &entity.RequestCreateOrder{CustomerUId: customerUid, GoodsName: goodsName})
			So(err, ShouldBeNil)
		})
	})
}
func TestUpdateOrder(t *testing.T) {
	Convey("UpdateOrder", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockOrderRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
			orderUid = "1f0d352b-d635-4a09-b7bc-4fb490e1682d"
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error order not found", func() {
			mockRepo.EXPECT().UpdateOrder(ctx, &entity.Order{MetaData: entity.MetaData{UID: uuid.MustParse(orderUid)}}, map[string]interface{}{"goods_name": "", "description": "", "amount": float64(0)}).Return(err).AnyTimes()
			err := service.UpdateOrder(ctx, &orderUid, &entity.RequestUpdateOrder{})
			So(err, ShouldNotBeNil)
		})
		Convey("success order found", func() {
			mockRepo.EXPECT().UpdateOrder(ctx, &entity.Order{MetaData: entity.MetaData{UID: uuid.MustParse(orderUid)}}, map[string]interface{}{"goods_name": "", "description": "", "amount": float64(0)}).Return(nil).AnyTimes()
			err := service.UpdateOrder(ctx, &orderUid, &entity.RequestUpdateOrder{})
			So(err, ShouldBeNil)
		})
	})
}
