// Code generated by mockery (devel). DO NOT EDIT.

package mocks

import (
	model "github.com/enesusta/tercuman/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// SqliteClient is an autogenerated mock type for the SqliteClient type
type SqliteClient struct {
	mock.Mock
}

// RetrieveTranslation provides a mock function with given fields: word
func (_m *SqliteClient) RetrieveTranslation(word string) model.Translation {
	ret := _m.Called(word)

	var r0 model.Translation
	if rf, ok := ret.Get(0).(func(string) model.Translation); ok {
		r0 = rf(word)
	} else {
		r0 = ret.Get(0).(model.Translation)
	}

	return r0
}

// RetrieveTranslations provides a mock function with given fields:
func (_m *SqliteClient) RetrieveTranslations() []model.Translation {
	ret := _m.Called()

	var r0 []model.Translation
	if rf, ok := ret.Get(0).(func() []model.Translation); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Translation)
		}
	}

	return r0
}