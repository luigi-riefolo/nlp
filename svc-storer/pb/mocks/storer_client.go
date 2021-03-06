// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import empty "github.com/golang/protobuf/ptypes/empty"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"
import pb "github.com/luigi-riefolo/nlp/svc-storer/pb"

// StorerClient is an autogenerated mock type for the StorerClient type
type StorerClient struct {
	mock.Mock
}

// StoreEntryHandler provides a mock function with given fields: ctx, in, opts
func (_m *StorerClient) StoreEntryHandler(ctx context.Context, in *pb.StoreEntryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *pb.StoreEntryRequest, ...grpc.CallOption) *empty.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.StoreEntryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
