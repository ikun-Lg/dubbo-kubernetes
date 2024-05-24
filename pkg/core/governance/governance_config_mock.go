/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apache/dubbo-admin/pkg/admin/config (interfaces: GovernanceConfig)

// Package config is a generated GoMock package.
package governance

import (
	reflect "reflect"
)

import (
	common "dubbo.apache.org/dubbo-go/v3/common"

	gomock "github.com/golang/mock/gomock"
)

// MockGovernanceConfig is a mock of GovernanceConfig interface.
type MockGovernanceConfig struct {
	ctrl     *gomock.Controller
	recorder *MockGovernanceConfigMockRecorder
}

func (m *MockGovernanceConfig) GetList() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MockGovernanceConfigMockRecorder is the mock recorder for MockGovernanceConfig.
type MockGovernanceConfigMockRecorder struct {
	mock *MockGovernanceConfig
}

// NewMockGovernanceConfig creates a new mock instance.
func NewMockGovernanceConfig(ctrl *gomock.Controller) *MockGovernanceConfig {
	mock := &MockGovernanceConfig{ctrl: ctrl}
	mock.recorder = &MockGovernanceConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGovernanceConfig) EXPECT() *MockGovernanceConfigMockRecorder {
	return m.recorder
}

// DeleteConfig mocks base method.
func (m *MockGovernanceConfig) DeleteConfig(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConfig indicates an expected call of DeleteConfig.
func (mr *MockGovernanceConfigMockRecorder) DeleteConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConfig", reflect.TypeOf((*MockGovernanceConfig)(nil).DeleteConfig), arg0)
}

// DeleteConfigWithGroup mocks base method.
func (m *MockGovernanceConfig) DeleteConfigWithGroup(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConfigWithGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConfigWithGroup indicates an expected call of DeleteConfigWithGroup.
func (mr *MockGovernanceConfigMockRecorder) DeleteConfigWithGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConfigWithGroup", reflect.TypeOf((*MockGovernanceConfig)(nil).DeleteConfigWithGroup), arg0, arg1)
}

// GetConfig mocks base method.
func (m *MockGovernanceConfig) GetConfig(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockGovernanceConfigMockRecorder) GetConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockGovernanceConfig)(nil).GetConfig), arg0)
}

// GetConfigWithGroup mocks base method.
func (m *MockGovernanceConfig) GetConfigWithGroup(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigWithGroup", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigWithGroup indicates an expected call of GetConfigWithGroup.
func (mr *MockGovernanceConfigMockRecorder) GetConfigWithGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigWithGroup", reflect.TypeOf((*MockGovernanceConfig)(nil).GetConfigWithGroup), arg0, arg1)
}

// Register mocks base method.
func (m *MockGovernanceConfig) Register(arg0 *common.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockGovernanceConfigMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockGovernanceConfig)(nil).Register), arg0)
}

// SetConfig mocks base method.
func (m *MockGovernanceConfig) SetConfig(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfig indicates an expected call of SetConfig.
func (mr *MockGovernanceConfigMockRecorder) SetConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockGovernanceConfig)(nil).SetConfig), arg0, arg1)
}

// SetConfigWithGroup mocks base method.
func (m *MockGovernanceConfig) SetConfigWithGroup(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfigWithGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfigWithGroup indicates an expected call of SetConfigWithGroup.
func (mr *MockGovernanceConfigMockRecorder) SetConfigWithGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfigWithGroup", reflect.TypeOf((*MockGovernanceConfig)(nil).SetConfigWithGroup), arg0, arg1, arg2)
}

// UnRegister mocks base method.
func (m *MockGovernanceConfig) UnRegister(arg0 *common.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnRegister", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnRegister indicates an expected call of UnRegister.
func (mr *MockGovernanceConfigMockRecorder) UnRegister(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnRegister", reflect.TypeOf((*MockGovernanceConfig)(nil).UnRegister), arg0)
}
