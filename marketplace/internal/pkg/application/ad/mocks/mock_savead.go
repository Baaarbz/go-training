// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	ad "barbz.dev/marketplace/internal/pkg/application/ad"

	mock "github.com/stretchr/testify/mock"
)

// SaveAd is an autogenerated mock type for the SaveAd type
type SaveAd struct {
	mock.Mock
}

type SaveAd_Expecter struct {
	mock *mock.Mock
}

func (_m *SaveAd) EXPECT() *SaveAd_Expecter {
	return &SaveAd_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, request
func (_m *SaveAd) Execute(ctx context.Context, request ad.SaveAdRequest) (ad.SaveAdResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 ad.SaveAdResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ad.SaveAdRequest) (ad.SaveAdResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ad.SaveAdRequest) ad.SaveAdResponse); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(ad.SaveAdResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ad.SaveAdRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAd_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type SaveAd_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - request ad.SaveAdRequest
func (_e *SaveAd_Expecter) Execute(ctx interface{}, request interface{}) *SaveAd_Execute_Call {
	return &SaveAd_Execute_Call{Call: _e.mock.On("Execute", ctx, request)}
}

func (_c *SaveAd_Execute_Call) Run(run func(ctx context.Context, request ad.SaveAdRequest)) *SaveAd_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ad.SaveAdRequest))
	})
	return _c
}

func (_c *SaveAd_Execute_Call) Return(_a0 ad.SaveAdResponse, _a1 error) *SaveAd_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SaveAd_Execute_Call) RunAndReturn(run func(context.Context, ad.SaveAdRequest) (ad.SaveAdResponse, error)) *SaveAd_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSaveAd interface {
	mock.TestingT
	Cleanup(func())
}

// NewSaveAd creates a new instance of SaveAd. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSaveAd(t mockConstructorTestingTNewSaveAd) *SaveAd {
	mock := &SaveAd{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
