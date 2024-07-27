package user

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"order-service/core/entity"
	"order-service/data_source/mysql_datasource"
	"order-service/shared/mock/repositories"
	"order-service/shared/util"

	"testing"
)

func TestGetUserByUId(t *testing.T) {
	Convey("GetUserByUid", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockUserRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
			userUid  = "1f0d352b-d635-4a09-b7bc-4fb490e1682d"
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error user not found", func() {
			mockRepo.EXPECT().GetUserByUid(ctx, &userUid).Return(nil, err).AnyTimes()
			res, err := service.GetUser(ctx, &userUid)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})
		Convey("success user found", func() {
			mockRepo.EXPECT().GetUserByUid(ctx, &userUid).Return(&entity.User{}, nil).AnyTimes()
			res, err := service.GetUser(ctx, &userUid)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})
	})
}

func TestGetAllUser(t *testing.T) {
	Convey("GetAllUser", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockUserRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error user not found", func() {
			mockRepo.EXPECT().GetAllUser(ctx, mysql_datasource.QueryOption{Limit: 10}).Return(nil, err).AnyTimes()
			res, err := service.GetAllUser(ctx, &entity.RequestGetList{Limit: 10})
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})
		Convey("success user found", func() {
			mockRepo.EXPECT().GetAllUser(ctx, mysql_datasource.QueryOption{Limit: 10}).Return(&entity.ResponseGetAllUser{}, nil).AnyTimes()
			res, err := service.GetAllUser(ctx, &entity.RequestGetList{Limit: 10})
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})
	})
}

func TestDeleteUser(t *testing.T) {
	Convey("DeleteUser", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockUserRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
			userUid  = "1f0d352b-d635-4a09-b7bc-4fb490e1682d"
		)

		service := New(Opts{Repo: mockRepo})

		propertyMap := []map[string]interface{}{
			{
				"key":      "uid",
				"operator": "=",
				"value":    userUid,
			},
		}
		Convey("error user not found", func() {
			mockRepo.EXPECT().DeleteUser(ctx, &entity.User{}, propertyMap).Return(err).AnyTimes()
			err := service.DeleteUser(ctx, &userUid)
			So(err, ShouldNotBeNil)
		})
		Convey("success user found", func() {
			mockRepo.EXPECT().DeleteUser(ctx, &entity.User{}, propertyMap).Return(nil).AnyTimes()
			err := service.DeleteUser(ctx, &userUid)
			So(err, ShouldBeNil)
		})
	})
}

func TestCreateUser(t *testing.T) {
	Convey("CreateUser", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockUserRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
			email    = "test@gmail.com"
			pass     = "test"
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error user not found", func() {
			mockRepo.EXPECT().GetUserByEmail(ctx, &email).Return(nil, err).AnyTimes()
			mockRepo.EXPECT().CreateUser(ctx, &entity.User{Email: email, Password: util.HashPassword(pass)}).Return(err).AnyTimes()
			err := service.CreateUser(ctx, &entity.RequestCreateUser{Email: email, Password: pass})
			So(err, ShouldNotBeNil)
		})
		Convey("success user found", func() {
			mockRepo.EXPECT().GetUserByEmail(ctx, &email).Return(nil, err).AnyTimes()
			mockRepo.EXPECT().CreateUser(ctx, &entity.User{Email: email, Password: util.HashPassword(pass)}).Return(nil).AnyTimes()
			err := service.CreateUser(ctx, &entity.RequestCreateUser{Email: email, Password: pass})
			So(err, ShouldBeNil)
		})
	})
}

func TestUpdateUser(t *testing.T) {
	Convey("UpdateUser", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockUserRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
			userUid  = "1f0d352b-d635-4a09-b7bc-4fb490e1682d"
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error user not found", func() {
			mockRepo.EXPECT().UpdateUser(ctx, &entity.User{MetaData: entity.MetaData{UID: uuid.MustParse(userUid)}}, map[string]interface{}{"email": "", "password": "", "username": ""}).Return(err).AnyTimes()
			err := service.UpdateUser(ctx, &userUid, &entity.RequestUpdateUser{})
			So(err, ShouldNotBeNil)
		})
		Convey("success user found", func() {
			mockRepo.EXPECT().UpdateUser(ctx, &entity.User{MetaData: entity.MetaData{UID: uuid.MustParse(userUid)}}, map[string]interface{}{"email": "", "password": "", "username": ""}).Return(nil).AnyTimes()
			err := service.UpdateUser(ctx, &userUid, &entity.RequestUpdateUser{})
			So(err, ShouldBeNil)
		})
	})
}

func TestLoginUser(t *testing.T) {
	Convey("LoginUser", t, func() {

		var (
			ctrl     = gomock.NewController(t)
			mockRepo = repositories.NewMockUserRepository(ctrl)
			err      = errors.New("error")
			ctx      = context.Background()
			email    = "test@gmail.com"
			pass     = "test"
		)

		service := New(Opts{Repo: mockRepo})

		Convey("error user not found", func() {
			mockRepo.EXPECT().GetUserByEmail(ctx, &email).Return(nil, err).AnyTimes()
			res, err := service.LoginUser(ctx, &entity.RequestLoginUser{Email: email, Password: pass})
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})
		Convey("success user found", func() {
			mockRepo.EXPECT().GetUserByEmail(ctx, &email).Return(&entity.User{
				Email:    email,
				Password: util.HashPassword(pass),
			}, nil).AnyTimes()
			res, err := service.LoginUser(ctx, &entity.RequestLoginUser{Email: email, Password: pass})
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})
	})
}
