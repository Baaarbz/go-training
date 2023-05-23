// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	ad "barbz.dev/marketplace/internal/pkg/application/ad"

	mock "github.com/stretchr/testify/mock"
)

// FindAllAds is an autogenerated mock type for the FindAllAds type
type FindAllAds struct {
	mock.Mock
}

type FindAllAds_Expecter struct {
	mock *mock.Mock
}

func (_m *FindAllAds) EXPECT() *FindAllAds_Expecter {
	return &FindAllAds_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx
func (_m *FindAllAds) Execute(ctx context.Context) ([]ad.GetAdsDto, error) {
	ret := _m.Called(ctx)

	var r0 []ad.GetAdsDto
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]ad.GetAdsDto, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []ad.GetAdsDto); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ad.GetAdsDto)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllAds_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type FindAllAds_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FindAllAds_Expecter) Execute(ctx interface{}) *FindAllAds_Execute_Call {
	return &FindAllAds_Execute_Call{Call: _e.mock.On("Execute", ctx)}
}

func (_c *FindAllAds_Execute_Call) Run(run func(ctx context.Context)) *FindAllAds_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FindAllAds_Execute_Call) Return(_a0 []ad.GetAdsDto, _a1 error) *FindAllAds_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FindAllAds_Execute_Call) RunAndReturn(run func(context.Context) ([]ad.GetAdsDto, error)) *FindAllAds_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewFindAllAds interface {
	mock.TestingT
	Cleanup(func())
}

// NewFindAllAds creates a new instance of FindAllAds. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFindAllAds(t mockConstructorTestingTNewFindAllAds) *FindAllAds {
	mock := &FindAllAds{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
