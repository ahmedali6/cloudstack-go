//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/RegistrationService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRegistrationServiceIface is a mock of RegistrationServiceIface interface.
type MockRegistrationServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrationServiceIfaceMockRecorder
}

// MockRegistrationServiceIfaceMockRecorder is the mock recorder for MockRegistrationServiceIface.
type MockRegistrationServiceIfaceMockRecorder struct {
	mock *MockRegistrationServiceIface
}

// NewMockRegistrationServiceIface creates a new mock instance.
func NewMockRegistrationServiceIface(ctrl *gomock.Controller) *MockRegistrationServiceIface {
	mock := &MockRegistrationServiceIface{ctrl: ctrl}
	mock.recorder = &MockRegistrationServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistrationServiceIface) EXPECT() *MockRegistrationServiceIfaceMockRecorder {
	return m.recorder
}

// NewRegisterOauthProviderParams mocks base method.
func (m *MockRegistrationServiceIface) NewRegisterOauthProviderParams(clientid, description, provider, redirecturi, secretkey string) *RegisterOauthProviderParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRegisterOauthProviderParams", clientid, description, provider, redirecturi, secretkey)
	ret0, _ := ret[0].(*RegisterOauthProviderParams)
	return ret0
}

// NewRegisterOauthProviderParams indicates an expected call of NewRegisterOauthProviderParams.
func (mr *MockRegistrationServiceIfaceMockRecorder) NewRegisterOauthProviderParams(clientid, description, provider, redirecturi, secretkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRegisterOauthProviderParams", reflect.TypeOf((*MockRegistrationServiceIface)(nil).NewRegisterOauthProviderParams), clientid, description, provider, redirecturi, secretkey)
}

// RegisterOauthProvider mocks base method.
func (m *MockRegistrationServiceIface) RegisterOauthProvider(p *RegisterOauthProviderParams) (*RegisterOauthProviderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterOauthProvider", p)
	ret0, _ := ret[0].(*RegisterOauthProviderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterOauthProvider indicates an expected call of RegisterOauthProvider.
func (mr *MockRegistrationServiceIfaceMockRecorder) RegisterOauthProvider(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterOauthProvider", reflect.TypeOf((*MockRegistrationServiceIface)(nil).RegisterOauthProvider), p)
}