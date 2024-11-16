// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/controller-runtime/pkg/client (interfaces: Patch,Reader,Writer,StatusClient,StatusWriter,Client,WithWatch,FieldIndexer)
//
// Generated by this command:
//
//	mockgen sigs.k8s.io/controller-runtime/pkg/client Patch,Reader,Writer,StatusClient,StatusWriter,Client,WithWatch,FieldIndexer
//

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	meta "k8s.io/apimachinery/pkg/api/meta"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockPatch is a mock of Patch interface.
type MockPatch struct {
	ctrl     *gomock.Controller
	recorder *MockPatchMockRecorder
	isgomock struct{}
}

// MockPatchMockRecorder is the mock recorder for MockPatch.
type MockPatchMockRecorder struct {
	mock *MockPatch
}

// NewMockPatch creates a new mock instance.
func NewMockPatch(ctrl *gomock.Controller) *MockPatch {
	mock := &MockPatch{ctrl: ctrl}
	mock.recorder = &MockPatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPatch) EXPECT() *MockPatchMockRecorder {
	return m.recorder
}

// Data mocks base method.
func (m *MockPatch) Data(obj client.Object) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data", obj)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Data indicates an expected call of Data.
func (mr *MockPatchMockRecorder) Data(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockPatch)(nil).Data), obj)
}

// Type mocks base method.
func (m *MockPatch) Type() types.PatchType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(types.PatchType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockPatchMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockPatch)(nil).Type))
}

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
	isgomock struct{}
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockReader) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockReaderMockRecorder) Get(ctx, key, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReader)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockReader) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, list}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockReaderMockRecorder) List(ctx, list any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, list}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReader)(nil).List), varargs...)
}

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
	isgomock struct{}
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWriter) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockWriterMockRecorder) Create(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWriter)(nil).Create), varargs...)
}

// Delete mocks base method.
func (m *MockWriter) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWriterMockRecorder) Delete(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWriter)(nil).Delete), varargs...)
}

// DeleteAllOf mocks base method.
func (m *MockWriter) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOf", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOf indicates an expected call of DeleteAllOf.
func (mr *MockWriterMockRecorder) DeleteAllOf(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOf", reflect.TypeOf((*MockWriter)(nil).DeleteAllOf), varargs...)
}

// Patch mocks base method.
func (m *MockWriter) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockWriterMockRecorder) Patch(ctx, obj, patch any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockWriter)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockWriter) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWriterMockRecorder) Update(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWriter)(nil).Update), varargs...)
}

// MockStatusClient is a mock of StatusClient interface.
type MockStatusClient struct {
	ctrl     *gomock.Controller
	recorder *MockStatusClientMockRecorder
	isgomock struct{}
}

// MockStatusClientMockRecorder is the mock recorder for MockStatusClient.
type MockStatusClientMockRecorder struct {
	mock *MockStatusClient
}

// NewMockStatusClient creates a new mock instance.
func NewMockStatusClient(ctrl *gomock.Controller) *MockStatusClient {
	mock := &MockStatusClient{ctrl: ctrl}
	mock.recorder = &MockStatusClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusClient) EXPECT() *MockStatusClientMockRecorder {
	return m.recorder
}

// Status mocks base method.
func (m *MockStatusClient) Status() client.SubResourceWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(client.SubResourceWriter)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockStatusClientMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockStatusClient)(nil).Status))
}

// MockStatusWriter is a mock of StatusWriter interface.
type MockStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockStatusWriterMockRecorder
	isgomock struct{}
}

// MockStatusWriterMockRecorder is the mock recorder for MockStatusWriter.
type MockStatusWriterMockRecorder struct {
	mock *MockStatusWriter
}

// NewMockStatusWriter creates a new mock instance.
func NewMockStatusWriter(ctrl *gomock.Controller) *MockStatusWriter {
	mock := &MockStatusWriter{ctrl: ctrl}
	mock.recorder = &MockStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusWriter) EXPECT() *MockStatusWriterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStatusWriter) Create(ctx context.Context, obj, subResource client.Object, opts ...client.SubResourceCreateOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj, subResource}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockStatusWriterMockRecorder) Create(ctx, obj, subResource any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj, subResource}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStatusWriter)(nil).Create), varargs...)
}

// Patch mocks base method.
func (m *MockStatusWriter) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockStatusWriterMockRecorder) Patch(ctx, obj, patch any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockStatusWriter)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockStatusWriter) Update(ctx context.Context, obj client.Object, opts ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStatusWriterMockRecorder) Update(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStatusWriter)(nil).Update), varargs...)
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockClientMockRecorder) Create(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClient)(nil).Create), varargs...)
}

// Delete mocks base method.
func (m *MockClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClientMockRecorder) Delete(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), varargs...)
}

// DeleteAllOf mocks base method.
func (m *MockClient) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOf", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOf indicates an expected call of DeleteAllOf.
func (mr *MockClientMockRecorder) DeleteAllOf(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOf", reflect.TypeOf((*MockClient)(nil).DeleteAllOf), varargs...)
}

// Get mocks base method.
func (m *MockClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockClientMockRecorder) Get(ctx, key, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), varargs...)
}

// GroupVersionKindFor mocks base method.
func (m *MockClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupVersionKindFor", obj)
	ret0, _ := ret[0].(schema.GroupVersionKind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupVersionKindFor indicates an expected call of GroupVersionKindFor.
func (mr *MockClientMockRecorder) GroupVersionKindFor(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupVersionKindFor", reflect.TypeOf((*MockClient)(nil).GroupVersionKindFor), obj)
}

// IsObjectNamespaced mocks base method.
func (m *MockClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsObjectNamespaced", obj)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsObjectNamespaced indicates an expected call of IsObjectNamespaced.
func (mr *MockClientMockRecorder) IsObjectNamespaced(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsObjectNamespaced", reflect.TypeOf((*MockClient)(nil).IsObjectNamespaced), obj)
}

// List mocks base method.
func (m *MockClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, list}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockClientMockRecorder) List(ctx, list any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, list}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClient)(nil).List), varargs...)
}

// Patch mocks base method.
func (m *MockClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockClientMockRecorder) Patch(ctx, obj, patch any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockClient)(nil).Patch), varargs...)
}

// RESTMapper mocks base method.
func (m *MockClient) RESTMapper() meta.RESTMapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTMapper")
	ret0, _ := ret[0].(meta.RESTMapper)
	return ret0
}

// RESTMapper indicates an expected call of RESTMapper.
func (mr *MockClientMockRecorder) RESTMapper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTMapper", reflect.TypeOf((*MockClient)(nil).RESTMapper))
}

// Scheme mocks base method.
func (m *MockClient) Scheme() *runtime.Scheme {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scheme")
	ret0, _ := ret[0].(*runtime.Scheme)
	return ret0
}

// Scheme indicates an expected call of Scheme.
func (mr *MockClientMockRecorder) Scheme() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scheme", reflect.TypeOf((*MockClient)(nil).Scheme))
}

// Status mocks base method.
func (m *MockClient) Status() client.SubResourceWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(client.SubResourceWriter)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockClientMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockClient)(nil).Status))
}

// SubResource mocks base method.
func (m *MockClient) SubResource(subResource string) client.SubResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubResource", subResource)
	ret0, _ := ret[0].(client.SubResourceClient)
	return ret0
}

// SubResource indicates an expected call of SubResource.
func (mr *MockClientMockRecorder) SubResource(subResource any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubResource", reflect.TypeOf((*MockClient)(nil).SubResource), subResource)
}

// Update mocks base method.
func (m *MockClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockClientMockRecorder) Update(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClient)(nil).Update), varargs...)
}

// MockWithWatch is a mock of WithWatch interface.
type MockWithWatch struct {
	ctrl     *gomock.Controller
	recorder *MockWithWatchMockRecorder
	isgomock struct{}
}

// MockWithWatchMockRecorder is the mock recorder for MockWithWatch.
type MockWithWatchMockRecorder struct {
	mock *MockWithWatch
}

// NewMockWithWatch creates a new mock instance.
func NewMockWithWatch(ctrl *gomock.Controller) *MockWithWatch {
	mock := &MockWithWatch{ctrl: ctrl}
	mock.recorder = &MockWithWatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWithWatch) EXPECT() *MockWithWatchMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWithWatch) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockWithWatchMockRecorder) Create(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWithWatch)(nil).Create), varargs...)
}

// Delete mocks base method.
func (m *MockWithWatch) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWithWatchMockRecorder) Delete(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWithWatch)(nil).Delete), varargs...)
}

// DeleteAllOf mocks base method.
func (m *MockWithWatch) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOf", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOf indicates an expected call of DeleteAllOf.
func (mr *MockWithWatchMockRecorder) DeleteAllOf(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOf", reflect.TypeOf((*MockWithWatch)(nil).DeleteAllOf), varargs...)
}

// Get mocks base method.
func (m *MockWithWatch) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockWithWatchMockRecorder) Get(ctx, key, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockWithWatch)(nil).Get), varargs...)
}

// GroupVersionKindFor mocks base method.
func (m *MockWithWatch) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupVersionKindFor", obj)
	ret0, _ := ret[0].(schema.GroupVersionKind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupVersionKindFor indicates an expected call of GroupVersionKindFor.
func (mr *MockWithWatchMockRecorder) GroupVersionKindFor(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupVersionKindFor", reflect.TypeOf((*MockWithWatch)(nil).GroupVersionKindFor), obj)
}

// IsObjectNamespaced mocks base method.
func (m *MockWithWatch) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsObjectNamespaced", obj)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsObjectNamespaced indicates an expected call of IsObjectNamespaced.
func (mr *MockWithWatchMockRecorder) IsObjectNamespaced(obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsObjectNamespaced", reflect.TypeOf((*MockWithWatch)(nil).IsObjectNamespaced), obj)
}

// List mocks base method.
func (m *MockWithWatch) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, list}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockWithWatchMockRecorder) List(ctx, list any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, list}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWithWatch)(nil).List), varargs...)
}

// Patch mocks base method.
func (m *MockWithWatch) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockWithWatchMockRecorder) Patch(ctx, obj, patch any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockWithWatch)(nil).Patch), varargs...)
}

// RESTMapper mocks base method.
func (m *MockWithWatch) RESTMapper() meta.RESTMapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTMapper")
	ret0, _ := ret[0].(meta.RESTMapper)
	return ret0
}

// RESTMapper indicates an expected call of RESTMapper.
func (mr *MockWithWatchMockRecorder) RESTMapper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTMapper", reflect.TypeOf((*MockWithWatch)(nil).RESTMapper))
}

// Scheme mocks base method.
func (m *MockWithWatch) Scheme() *runtime.Scheme {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scheme")
	ret0, _ := ret[0].(*runtime.Scheme)
	return ret0
}

// Scheme indicates an expected call of Scheme.
func (mr *MockWithWatchMockRecorder) Scheme() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scheme", reflect.TypeOf((*MockWithWatch)(nil).Scheme))
}

// Status mocks base method.
func (m *MockWithWatch) Status() client.SubResourceWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(client.SubResourceWriter)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockWithWatchMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockWithWatch)(nil).Status))
}

// SubResource mocks base method.
func (m *MockWithWatch) SubResource(subResource string) client.SubResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubResource", subResource)
	ret0, _ := ret[0].(client.SubResourceClient)
	return ret0
}

// SubResource indicates an expected call of SubResource.
func (mr *MockWithWatchMockRecorder) SubResource(subResource any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubResource", reflect.TypeOf((*MockWithWatch)(nil).SubResource), subResource)
}

// Update mocks base method.
func (m *MockWithWatch) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWithWatchMockRecorder) Update(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWithWatch)(nil).Update), varargs...)
}

// Watch mocks base method.
func (m *MockWithWatch) Watch(ctx context.Context, obj client.ObjectList, opts ...client.ListOption) (watch.Interface, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockWithWatchMockRecorder) Watch(ctx, obj any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockWithWatch)(nil).Watch), varargs...)
}

// MockFieldIndexer is a mock of FieldIndexer interface.
type MockFieldIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockFieldIndexerMockRecorder
	isgomock struct{}
}

// MockFieldIndexerMockRecorder is the mock recorder for MockFieldIndexer.
type MockFieldIndexerMockRecorder struct {
	mock *MockFieldIndexer
}

// NewMockFieldIndexer creates a new mock instance.
func NewMockFieldIndexer(ctrl *gomock.Controller) *MockFieldIndexer {
	mock := &MockFieldIndexer{ctrl: ctrl}
	mock.recorder = &MockFieldIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFieldIndexer) EXPECT() *MockFieldIndexerMockRecorder {
	return m.recorder
}

// IndexField mocks base method.
func (m *MockFieldIndexer) IndexField(ctx context.Context, obj client.Object, field string, extractValue client.IndexerFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexField", ctx, obj, field, extractValue)
	ret0, _ := ret[0].(error)
	return ret0
}

// IndexField indicates an expected call of IndexField.
func (mr *MockFieldIndexerMockRecorder) IndexField(ctx, obj, field, extractValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexField", reflect.TypeOf((*MockFieldIndexer)(nil).IndexField), ctx, obj, field, extractValue)
}
