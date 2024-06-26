// Code generated by MockGen. DO NOT EDIT.
// Source: metrics.go
//
// Generated by this command:
//
//	mockgen -source=metrics.go -destination=mock_metrics.go -package=container
//

// Package container is a generated GoMock package.
package container

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// DeltaUpDownCounter mocks base method.
func (m *MockMetrics) DeltaUpDownCounter(ctx context.Context, name string, value float64, labels ...string) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, name, value}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "DeltaUpDownCounter", varargs...)
}

// DeltaUpDownCounter indicates an expected call of DeltaUpDownCounter.
func (mr *MockMetricsMockRecorder) DeltaUpDownCounter(ctx, name, value any, labels ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, name, value}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeltaUpDownCounter", reflect.TypeOf((*MockMetrics)(nil).DeltaUpDownCounter), varargs...)
}

// IncrementCounter mocks base method.
func (m *MockMetrics) IncrementCounter(ctx context.Context, name string, labels ...string) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, name}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "IncrementCounter", varargs...)
}

// IncrementCounter indicates an expected call of IncrementCounter.
func (mr *MockMetricsMockRecorder) IncrementCounter(ctx, name any, labels ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, name}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementCounter", reflect.TypeOf((*MockMetrics)(nil).IncrementCounter), varargs...)
}

// NewCounter mocks base method.
func (m *MockMetrics) NewCounter(name, desc string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewCounter", name, desc)
}

// NewCounter indicates an expected call of NewCounter.
func (mr *MockMetricsMockRecorder) NewCounter(name, desc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCounter", reflect.TypeOf((*MockMetrics)(nil).NewCounter), name, desc)
}

// NewGauge mocks base method.
func (m *MockMetrics) NewGauge(name, desc string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewGauge", name, desc)
}

// NewGauge indicates an expected call of NewGauge.
func (mr *MockMetricsMockRecorder) NewGauge(name, desc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGauge", reflect.TypeOf((*MockMetrics)(nil).NewGauge), name, desc)
}

// NewHistogram mocks base method.
func (m *MockMetrics) NewHistogram(name, desc string, buckets ...float64) {
	m.ctrl.T.Helper()
	varargs := []any{name, desc}
	for _, a := range buckets {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "NewHistogram", varargs...)
}

// NewHistogram indicates an expected call of NewHistogram.
func (mr *MockMetricsMockRecorder) NewHistogram(name, desc any, buckets ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, desc}, buckets...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHistogram", reflect.TypeOf((*MockMetrics)(nil).NewHistogram), varargs...)
}

// NewUpDownCounter mocks base method.
func (m *MockMetrics) NewUpDownCounter(name, desc string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewUpDownCounter", name, desc)
}

// NewUpDownCounter indicates an expected call of NewUpDownCounter.
func (mr *MockMetricsMockRecorder) NewUpDownCounter(name, desc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpDownCounter", reflect.TypeOf((*MockMetrics)(nil).NewUpDownCounter), name, desc)
}

// RecordHistogram mocks base method.
func (m *MockMetrics) RecordHistogram(ctx context.Context, name string, value float64, labels ...string) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, name, value}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "RecordHistogram", varargs...)
}

// RecordHistogram indicates an expected call of RecordHistogram.
func (mr *MockMetricsMockRecorder) RecordHistogram(ctx, name, value any, labels ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, name, value}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordHistogram", reflect.TypeOf((*MockMetrics)(nil).RecordHistogram), varargs...)
}

// SetGauge mocks base method.
func (m *MockMetrics) SetGauge(name string, value float64, labels ...string) {
	m.ctrl.T.Helper()
	varargs := []any{name, value}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetGauge", varargs...)
}

// SetGauge indicates an expected call of SetGauge.
func (mr *MockMetricsMockRecorder) SetGauge(name, value any, labels ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, value}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGauge", reflect.TypeOf((*MockMetrics)(nil).SetGauge), varargs...)
}