// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/etcinit/phabulous/app/interfaces (interfaces: Bot,Message,HandlerTuple)

package mocks

import (
	gonduit "github.com/etcinit/gonduit"
	interfaces "github.com/etcinit/phabulous/app/interfaces"
	messages "github.com/etcinit/phabulous/app/messages"
	gomock "github.com/golang/mock/gomock"
	confer "github.com/jacobstr/confer"
	regexp "regexp"
)

// Mock of Bot interface
type MockBot struct {
	ctrl     *gomock.Controller
	recorder *_MockBotRecorder
}

// Recorder for MockBot (not exported)
type _MockBotRecorder struct {
	mock *MockBot
}

func NewMockBot(ctrl *gomock.Controller) *MockBot {
	mock := &MockBot{ctrl: ctrl}
	mock.recorder = &_MockBotRecorder{mock}
	return mock
}

func (_m *MockBot) EXPECT() *_MockBotRecorder {
	return _m.recorder
}

func (_m *MockBot) Excuse(_param0 interfaces.Message, _param1 error) {
	_m.ctrl.Call(_m, "Excuse", _param0, _param1)
}

func (_mr *_MockBotRecorder) Excuse(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Excuse", arg0, arg1)
}

func (_m *MockBot) GetConfig() *confer.Config {
	ret := _m.ctrl.Call(_m, "GetConfig")
	ret0, _ := ret[0].(*confer.Config)
	return ret0
}

func (_mr *_MockBotRecorder) GetConfig() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetConfig")
}

func (_m *MockBot) GetGonduit() (*gonduit.Conn, error) {
	ret := _m.ctrl.Call(_m, "GetGonduit")
	ret0, _ := ret[0].(*gonduit.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBotRecorder) GetGonduit() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetGonduit")
}

func (_m *MockBot) GetHandlers() []interfaces.HandlerTuple {
	ret := _m.ctrl.Call(_m, "GetHandlers")
	ret0, _ := ret[0].([]interfaces.HandlerTuple)
	return ret0
}

func (_mr *_MockBotRecorder) GetHandlers() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetHandlers")
}

func (_m *MockBot) GetIMHandlers() []interfaces.HandlerTuple {
	ret := _m.ctrl.Call(_m, "GetIMHandlers")
	ret0, _ := ret[0].([]interfaces.HandlerTuple)
	return ret0
}

func (_mr *_MockBotRecorder) GetIMHandlers() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetIMHandlers")
}

func (_m *MockBot) GetModules() []interfaces.Module {
	ret := _m.ctrl.Call(_m, "GetModules")
	ret0, _ := ret[0].([]interfaces.Module)
	return ret0
}

func (_mr *_MockBotRecorder) GetModules() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetModules")
}

func (_m *MockBot) GetUsageHandler() interfaces.Handler {
	ret := _m.ctrl.Call(_m, "GetUsageHandler")
	ret0, _ := ret[0].(interfaces.Handler)
	return ret0
}

func (_mr *_MockBotRecorder) GetUsageHandler() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUsageHandler")
}

func (_m *MockBot) GetUsername(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetUsername", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBotRecorder) GetUsername(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUsername", arg0)
}

func (_m *MockBot) Post(_param0 string, _param1 string, _param2 messages.Icon, _param3 bool) {
	_m.ctrl.Call(_m, "Post", _param0, _param1, _param2, _param3)
}

func (_mr *_MockBotRecorder) Post(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Post", arg0, arg1, arg2, arg3)
}

func (_m *MockBot) PostImage(_param0 string, _param1 string, _param2 string, _param3 messages.Icon, _param4 bool) {
	_m.ctrl.Call(_m, "PostImage", _param0, _param1, _param2, _param3, _param4)
}

func (_mr *_MockBotRecorder) PostImage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PostImage", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockBot) PostOnFeed(_param0 string) error {
	ret := _m.ctrl.Call(_m, "PostOnFeed", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockBotRecorder) PostOnFeed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PostOnFeed", arg0)
}

func (_m *MockBot) StartTyping(_param0 string) {
	_m.ctrl.Call(_m, "StartTyping", _param0)
}

func (_mr *_MockBotRecorder) StartTyping(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartTyping", arg0)
}

// Mock of Message interface
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *_MockMessageRecorder
}

// Recorder for MockMessage (not exported)
type _MockMessageRecorder struct {
	mock *MockMessage
}

func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &_MockMessageRecorder{mock}
	return mock
}

func (_m *MockMessage) EXPECT() *_MockMessageRecorder {
	return _m.recorder
}

func (_m *MockMessage) GetChannel() string {
	ret := _m.ctrl.Call(_m, "GetChannel")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockMessageRecorder) GetChannel() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetChannel")
}

func (_m *MockMessage) GetContent() string {
	ret := _m.ctrl.Call(_m, "GetContent")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockMessageRecorder) GetContent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetContent")
}

func (_m *MockMessage) GetProviderName() string {
	ret := _m.ctrl.Call(_m, "GetProviderName")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockMessageRecorder) GetProviderName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetProviderName")
}

func (_m *MockMessage) GetUserID() string {
	ret := _m.ctrl.Call(_m, "GetUserID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockMessageRecorder) GetUserID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserID")
}

func (_m *MockMessage) IsIM() bool {
	ret := _m.ctrl.Call(_m, "IsIM")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockMessageRecorder) IsIM() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsIM")
}

func (_m *MockMessage) IsSelf() bool {
	ret := _m.ctrl.Call(_m, "IsSelf")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockMessageRecorder) IsSelf() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsSelf")
}

// Mock of HandlerTuple interface
type MockHandlerTuple struct {
	ctrl     *gomock.Controller
	recorder *_MockHandlerTupleRecorder
}

// Recorder for MockHandlerTuple (not exported)
type _MockHandlerTupleRecorder struct {
	mock *MockHandlerTuple
}

func NewMockHandlerTuple(ctrl *gomock.Controller) *MockHandlerTuple {
	mock := &MockHandlerTuple{ctrl: ctrl}
	mock.recorder = &_MockHandlerTupleRecorder{mock}
	return mock
}

func (_m *MockHandlerTuple) EXPECT() *_MockHandlerTupleRecorder {
	return _m.recorder
}

func (_m *MockHandlerTuple) GetHandler() interfaces.Handler {
	ret := _m.ctrl.Call(_m, "GetHandler")
	ret0, _ := ret[0].(interfaces.Handler)
	return ret0
}

func (_mr *_MockHandlerTupleRecorder) GetHandler() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetHandler")
}

func (_m *MockHandlerTuple) GetPattern() *regexp.Regexp {
	ret := _m.ctrl.Call(_m, "GetPattern")
	ret0, _ := ret[0].(*regexp.Regexp)
	return ret0
}

func (_mr *_MockHandlerTupleRecorder) GetPattern() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPattern")
}