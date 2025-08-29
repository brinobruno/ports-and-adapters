package cli_test

import (
	"fmt"
	"testing"

	"github.com/brinobruno/ports-and-adapters/adapters/cli"
	mock_application "github.com/brinobruno/ports-and-adapters/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product 1"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()

	expectedResult := fmt.Sprintf(
		"Product ID %s with the name %s has been created with the price %f and status %s",
		productId, productName, productPrice, productStatus,
	)
	result, err := cli.Run(service, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		"Product ID %s with the name %s has been enabled with the price %f and status %s",
		productId, productName, productPrice, productStatus,
	)
	result, err = cli.Run(service, "enable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		"Product ID %s has been disabled",
		productId,
	)
	result, err = cli.Run(service, "disable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		"Product ID %s with the name %s has the price %f and status %s",
		productId, productName, productPrice, productStatus,
	)
	result, err = cli.Run(service, "something", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
