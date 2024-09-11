package user

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
	mock_utils "launchpad-go-rest/internal/lib/utils/mock"
	"launchpad-go-rest/internal/repository/mock/mock_user"
	"launchpad-go-rest/pkg/types/user"
	"testing"

	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/mock/gomock"
)

func Test_User_Find(t *testing.T) {
	Convey("Test_User_Find", t, func() {
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
			mockCalls    func()
			errExpected  bool
			expectedResp []user.FindResponse
		}{
			{
				desc: "Error: Find",
				mockCalls: func() {
					deps.mockUser.EXPECT().Find(gomock.Any()).Return([]user.User{}, errMock)
				},
				errExpected:  true,
				expectedResp: []user.FindResponse{},
			},
			{
				desc: "Success",
				mockCalls: func() {
					deps.mockUser.EXPECT().Find(gomock.Any()).Return([]user.User{
						{
							ID:       "id",
							Email:    "test@email.com",
							Password: "hash123",
						},
					}, nil)
				},
				errExpected: false,
				expectedResp: []user.FindResponse{
					{
						ID:    "id",
						Email: "test@email.com",
					},
				},
			},
		}

		for _, testCase := range testCases {
			Convey(testCase.desc, func() {
				testCase.mockCalls()

				resp, err := svc.Find(context.Background())
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
