// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/snow/engine/common (interfaces: Sender)

// Package common is a generated GoMock package.
package common

import (
        snow "github.com/ava-labs/avalanchego/snow"
        set "github.com/ava-labs/avalanchego/utils/set"
        gomock "github.com/golang/mock/gomock"
        reflect "reflect"
        context "context"
        ids "github.com/ava-labs/avalanchego/ids"
)

// MockSender is a mock of Sender interface.
type MockSender struct {
        ctrl     *gomock.Controller
        recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender.
type MockSenderMockRecorder struct {
        mock *MockSender
}

// NewMockSender creates a new mock instance.
func NewMockSender(ctrl *gomock.Controller) *MockSender {
        mock := &MockSender{ctrl: ctrl}
        mock.recorder = &MockSenderMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
        return m.recorder
}

// Accept mocks base method.
func (m *MockSender) Accept(arg0 *snow.ConsensusContext, arg1 ids.ID, arg2 []byte) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Accept", arg0, arg1, arg2)
        ret0, _ := ret[0].(error)
        return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockSenderMockRecorder) Accept(arg0, arg1, arg2 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockSender)(nil).Accept), arg0, arg1, arg2)
}

// SendAccepted mocks base method.
func (m *MockSender) SendAccepted(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []ids.ID) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendAccepted", arg0, arg1, arg2, arg3)
}

// SendAccepted indicates an expected call of SendAccepted.
func (mr *MockSenderMockRecorder) SendAccepted(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAccepted", reflect.TypeOf((*MockSender)(nil).SendAccepted), arg0, arg1, arg2, arg3)
}

// SendAcceptedFrontier mocks base method.
func (m *MockSender) SendAcceptedFrontier(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []ids.ID) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendAcceptedFrontier", arg0, arg1, arg2, arg3)
}

// SendAcceptedFrontier indicates an expected call of SendAcceptedFrontier.
func (mr *MockSenderMockRecorder) SendAcceptedFrontier(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAcceptedFrontier", reflect.TypeOf((*MockSender)(nil).SendAcceptedFrontier), arg0, arg1, arg2, arg3)
}

// SendAcceptedStateSummary mocks base method.
func (m *MockSender) SendAcceptedStateSummary(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []ids.ID) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendAcceptedStateSummary", arg0, arg1, arg2, arg3)
}

// SendAcceptedStateSummary indicates an expected call of SendAcceptedStateSummary.
func (mr *MockSenderMockRecorder) SendAcceptedStateSummary(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAcceptedStateSummary", reflect.TypeOf((*MockSender)(nil).SendAcceptedStateSummary), arg0, arg1, arg2, arg3)
}

// SendAncestors mocks base method.
func (m *MockSender) SendAncestors(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 [][]byte) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendAncestors", arg0, arg1, arg2, arg3)
}

// SendAncestors indicates an expected call of SendAncestors.
func (mr *MockSenderMockRecorder) SendAncestors(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAncestors", reflect.TypeOf((*MockSender)(nil).SendAncestors), arg0, arg1, arg2, arg3)
}

// SendAppGossip mocks base method.
func (m *MockSender) SendAppGossip(arg0 context.Context, arg1 []byte) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendAppGossip", arg0, arg1)
        ret0, _ := ret[0].(error)
        return ret0
}

// SendAppGossip indicates an expected call of SendAppGossip.
func (mr *MockSenderMockRecorder) SendAppGossip(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppGossip", reflect.TypeOf((*MockSender)(nil).SendAppGossip), arg0, arg1)
}

// SendAppGossipSpecific mocks base method.
func (m *MockSender) SendAppGossipSpecific(arg0 context.Context, arg1 set.Set[ids.NodeID], arg2 []byte) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendAppGossipSpecific", arg0, arg1, arg2)
        ret0, _ := ret[0].(error)
        return ret0
}

// SendAppGossipSpecific indicates an expected call of SendAppGossipSpecific.
func (mr *MockSenderMockRecorder) SendAppGossipSpecific(arg0, arg1, arg2 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppGossipSpecific", reflect.TypeOf((*MockSender)(nil).SendAppGossipSpecific), arg0, arg1, arg2)
}

// SendAppRequest mocks base method.
func (m *MockSender) SendAppRequest(arg0 context.Context, arg1 set.Set[ids.NodeID], arg2 uint32, arg3 []byte) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendAppRequest", arg0, arg1, arg2, arg3)
        ret0, _ := ret[0].(error)
        return ret0
}

// SendAppRequest indicates an expected call of SendAppRequest.
func (mr *MockSenderMockRecorder) SendAppRequest(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppRequest", reflect.TypeOf((*MockSender)(nil).SendAppRequest), arg0, arg1, arg2, arg3)
}

// SendAppResponse mocks base method.
func (m *MockSender) SendAppResponse(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []byte) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendAppResponse", arg0, arg1, arg2, arg3)
        ret0, _ := ret[0].(error)
        return ret0
}

// SendAppResponse indicates an expected call of SendAppResponse.
func (mr *MockSenderMockRecorder) SendAppResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppResponse", reflect.TypeOf((*MockSender)(nil).SendAppResponse), arg0, arg1, arg2, arg3)
}

// SendChits mocks base method.
func (m *MockSender) SendChits(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []ids.ID) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendChits", arg0, arg1, arg2, arg3)
}

// SendChits indicates an expected call of SendChits.
func (mr *MockSenderMockRecorder) SendChits(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendChits", reflect.TypeOf((*MockSender)(nil).SendChits), arg0, arg1, arg2, arg3)
}

// SendCrossChainAppRequest mocks base method.
func (m *MockSender) SendCrossChainAppRequest(arg0 context.Context, arg1 ids.ID, arg2 uint32, arg3 []byte) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendCrossChainAppRequest", arg0, arg1, arg2, arg3)
        ret0, _ := ret[0].(error)
        return ret0
}

// SendCrossChainAppRequest indicates an expected call of SendCrossChainAppRequest.
func (mr *MockSenderMockRecorder) SendCrossChainAppRequest(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCrossChainAppRequest", reflect.TypeOf((*MockSender)(nil).SendCrossChainAppRequest), arg0, arg1, arg2, arg3)
}

// SendCrossChainAppResponse mocks base method.
func (m *MockSender) SendCrossChainAppResponse(arg0 context.Context, arg1 ids.ID, arg2 uint32, arg3 []byte) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendCrossChainAppResponse", arg0, arg1, arg2, arg3)
        ret0, _ := ret[0].(error)
        return ret0
}

// SendCrossChainAppResponse indicates an expected call of SendCrossChainAppResponse.
func (mr *MockSenderMockRecorder) SendCrossChainAppResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCrossChainAppResponse", reflect.TypeOf((*MockSender)(nil).SendCrossChainAppResponse), arg0, arg1, arg2, arg3)
}

// SendGet mocks base method.
func (m *MockSender) SendGet(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 ids.ID) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendGet", arg0, arg1, arg2, arg3)
}

// SendGet indicates an expected call of SendGet.
func (mr *MockSenderMockRecorder) SendGet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGet", reflect.TypeOf((*MockSender)(nil).SendGet), arg0, arg1, arg2, arg3)
}

// SendGetAccepted mocks base method.
func (m *MockSender) SendGetAccepted(arg0 context.Context, arg1 set.Set[ids.NodeID], arg2 uint32, arg3 []ids.ID) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendGetAccepted", arg0, arg1, arg2, arg3)
}

// SendGetAccepted indicates an expected call of SendGetAccepted.
func (mr *MockSenderMockRecorder) SendGetAccepted(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAccepted", reflect.TypeOf((*MockSender)(nil).SendGetAccepted), arg0, arg1, arg2, arg3)
}

// SendGetAcceptedFrontier mocks base method.
func (m *MockSender) SendGetAcceptedFrontier(arg0 context.Context, arg1 set.Set[ids.NodeID], arg2 uint32) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendGetAcceptedFrontier", arg0, arg1, arg2)
}

// SendGetAcceptedFrontier indicates an expected call of SendGetAcceptedFrontier.
func (mr *MockSenderMockRecorder) SendGetAcceptedFrontier(arg0, arg1, arg2 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAcceptedFrontier", reflect.TypeOf((*MockSender)(nil).SendGetAcceptedFrontier), arg0, arg1, arg2)
}

// SendGetAcceptedStateSummary mocks base method.
func (m *MockSender) SendGetAcceptedStateSummary(arg0 context.Context, arg1 set.Set[ids.NodeID], arg2 uint32, arg3 []uint64) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendGetAcceptedStateSummary", arg0, arg1, arg2, arg3)
}

// SendGetAcceptedStateSummary indicates an expected call of SendGetAcceptedStateSummary.
func (mr *MockSenderMockRecorder) SendGetAcceptedStateSummary(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAcceptedStateSummary", reflect.TypeOf((*MockSender)(nil).SendGetAcceptedStateSummary), arg0, arg1, arg2, arg3)
}

// SendGetAncestors mocks base method.
func (m *MockSender) SendGetAncestors(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 ids.ID) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendGetAncestors", arg0, arg1, arg2, arg3)
}

// SendGetAncestors indicates an expected call of SendGetAncestors.
func (mr *MockSenderMockRecorder) SendGetAncestors(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAncestors", reflect.TypeOf((*MockSender)(nil).SendGetAncestors), arg0, arg1, arg2, arg3)
}

// SendGetStateSummaryFrontier mocks base method.
func (m *MockSender) SendGetStateSummaryFrontier(arg0 context.Context, arg1 set.Set[ids.NodeID], arg2 uint32) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendGetStateSummaryFrontier", arg0, arg1, arg2)
}

// SendGetStateSummaryFrontier indicates an expected call of SendGetStateSummaryFrontier.
func (mr *MockSenderMockRecorder) SendGetStateSummaryFrontier(arg0, arg1, arg2 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetStateSummaryFrontier", reflect.TypeOf((*MockSender)(nil).SendGetStateSummaryFrontier), arg0, arg1, arg2)
}

// SendGossip mocks base method.
func (m *MockSender) SendGossip(arg0 context.Context, arg1 []byte) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendGossip", arg0, arg1)
}

// SendGossip indicates an expected call of SendGossip.
func (mr *MockSenderMockRecorder) SendGossip(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGossip", reflect.TypeOf((*MockSender)(nil).SendGossip), arg0, arg1)
}

// SendPullQuery mocks base method.
func (m *MockSender) SendPullQuery(arg0 context.Context, arg1 set.Set[ids.NodeID], arg2 uint32, arg3 ids.ID) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendPullQuery", arg0, arg1, arg2, arg3)
}

// SendPullQuery indicates an expected call of SendPullQuery.
func (mr *MockSenderMockRecorder) SendPullQuery(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPullQuery", reflect.TypeOf((*MockSender)(nil).SendPullQuery), arg0, arg1, arg2, arg3)
}

// SendPushQuery mocks base method.
func (m *MockSender) SendPushQuery(arg0 context.Context, arg1 set.Set[ids.NodeID], arg2 uint32, arg3 []byte) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendPushQuery", arg0, arg1, arg2, arg3)
}

// SendPushQuery indicates an expected call of SendPushQuery.
func (mr *MockSenderMockRecorder) SendPushQuery(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPushQuery", reflect.TypeOf((*MockSender)(nil).SendPushQuery), arg0, arg1, arg2, arg3)
}

// SendPut mocks base method.
func (m *MockSender) SendPut(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []byte) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendPut", arg0, arg1, arg2, arg3)
}

// SendPut indicates an expected call of SendPut.
func (mr *MockSenderMockRecorder) SendPut(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPut", reflect.TypeOf((*MockSender)(nil).SendPut), arg0, arg1, arg2, arg3)
}

// SendStateSummaryFrontier mocks base method.
func (m *MockSender) SendStateSummaryFrontier(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []byte) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "SendStateSummaryFrontier", arg0, arg1, arg2, arg3)
}

// SendStateSummaryFrontier indicates an expected call of SendStateSummaryFrontier.
func (mr *MockSenderMockRecorder) SendStateSummaryFrontier(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendStateSummaryFrontier", reflect.TypeOf((*MockSender)(nil).SendStateSummaryFrontier), arg0, arg1, arg2, arg3)
}