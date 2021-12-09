// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	chats "messaging/business/chats"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// ListChat provides a mock function with given fields: id_users
func (_m *Service) ListChat(id_users int) ([]*chats.ChatsList, error) {
	ret := _m.Called(id_users)

	var r0 []*chats.ChatsList
	if rf, ok := ret.Get(0).(func(int) []*chats.ChatsList); ok {
		r0 = rf(id_users)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*chats.ChatsList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id_users)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadChat provides a mock function with given fields: id_users, id_groups
func (_m *Service) ReadChat(id_users int, id_groups int) ([]*chats.ReadList, error) {
	ret := _m.Called(id_users, id_groups)

	var r0 []*chats.ReadList
	if rf, ok := ret.Get(0).(func(int, int) []*chats.ReadList); ok {
		r0 = rf(id_users, id_groups)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*chats.ReadList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(id_users, id_groups)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: chat_specs
func (_m *Service) SendMessage(chat_specs chats.ChatsSpec) (error, interface{}) {
	ret := _m.Called(chat_specs)

	var r0 error
	if rf, ok := ret.Get(0).(func(chats.ChatsSpec) error); ok {
		r0 = rf(chat_specs)
	} else {
		r0 = ret.Error(0)
	}

	var r1 interface{}
	if rf, ok := ret.Get(1).(func(chats.ChatsSpec) interface{}); ok {
		r1 = rf(chat_specs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(interface{})
		}
	}

	return r0, r1
}
