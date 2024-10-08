package user

import (
	"context"
	"database/sql"
	"launchpad-go-rest/internal/lib/errors"
	mock_utils "launchpad-go-rest/internal/lib/utils/mock"
	"launchpad-go-rest/internal/repository/mock/mock_user"
	"launchpad-go-rest/pkg/types/user"
	"testing"

	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/mock/gomock"
)

func Test_User_FindByID(t *testing.T) {
	Convey("Test_User_FindByID", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		deps := struct {
			mockUser  *mock_user.MockRepository
			mockUtils *mock_utils.MockUtils
		}{
			mockUser:  mock_user.NewMockRepository(ctrl),
			mockUtils: mock_utils.NewMockUtils(ctrl),
		}

		logger := echo.New().Logger
		svc := New(deps.mockUser, logger, deps.mockUtils)

		errMock := errors.New("mock error")

		testCases := []struct {
			desc         string
			req          user.FindByIDRequest
			mockCalls    func()
			errExpected  bool
			expectedResp user.FindByIDResponse
		}{
			{
				desc:         "Error: Validation error",
				req:          user.FindByIDRequest{},
				mockCalls:    func() {},
				errExpected:  true,
				expectedResp: user.FindByIDResponse{},
			},
			{
				desc: "Error: User not found",
				req: user.FindByIDRequest{
					ID: "id",
				},
				mockCalls: func() {
					deps.mockUser.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(user.User{}, sql.ErrNoRows)
				},
				errExpected:  true,
				expectedResp: user.FindByIDResponse{},
			},
			{
				desc: "Error: Find by id",
				req: user.FindByIDRequest{
					ID: "id",
				},
				mockCalls: func() {
					deps.mockUser.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(user.User{}, errMock)
				},
				errExpected:  true,
				expectedResp: user.FindByIDResponse{},
			},
			{
				desc: "Success",
				req: user.FindByIDRequest{
					ID: "id",
				},
				mockCalls: func() {
					deps.mockUser.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(user.User{
						ID:       "id",
						Email:    "test@email.com",
						Password: "hash123",
					}, nil)
				},
				errExpected: false,
				expectedResp: user.FindByIDResponse{
					ID:    "id",
					Email: "test@email.com",
				},
			},
		}

		for _, testCase := range testCases {
			Convey(testCase.desc, func() {
				testCase.mockCalls()

				resp, err := svc.FindByID(context.Background(), testCase.req)
				So(resp, ShouldResemble, testCase.expectedResp)
				if testCase.errExpected {
					So(err, ShouldNotBeNil)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}
