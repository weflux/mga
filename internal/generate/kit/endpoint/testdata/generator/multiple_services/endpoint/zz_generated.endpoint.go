//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Acme Inc.
// All rights reserved.
//
// Licensed under "Only for testing purposes" license.

// Code generated by mga tool. DO NOT EDIT.

package pkgdriver

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
	"sagikazarmark.dev/mga/internal/generate/kit/endpoint/testdata/generator/multiple_services"
)

// endpointError identifies an error that should be returned as an endpoint error.
type endpointError interface {
	EndpointError() bool
}

// serviceError identifies an error that should be returned as a service error.
type serviceError interface {
	ServiceError() bool
}

// Endpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateTodo endpoint.Endpoint
}

// MakeEndpoints returns a(n) Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeEndpoints(service multiple_services.Service, middleware ...endpoint.Middleware) Endpoints {
	mw := kitxendpoint.Combine(middleware...)

	return Endpoints{CreateTodo: kitxendpoint.OperationNameMiddleware("multiple_services.CreateTodo")(mw(MakeCreateTodoEndpoint(service)))}
}

// TraceEndpoints returns a(n) Endpoints struct where each endpoint is wrapped with a tracing middleware.
func TraceEndpoints(endpoints Endpoints) Endpoints {
	return Endpoints{CreateTodo: kitoc.TraceEndpoint("multiple_services.CreateTodo")(endpoints.CreateTodo)}
}

// CreateTodoRequest is a request struct for CreateTodo endpoint.
type CreateTodoRequest struct {
	Text string
}

// CreateTodoResponse is a response struct for CreateTodo endpoint.
type CreateTodoResponse struct {
	Id  string
	Err error
}

func (r CreateTodoResponse) Failed() error {
	return r.Err
}

// MakeCreateTodoEndpoint returns an endpoint for the matching method of the underlying service.
func MakeCreateTodoEndpoint(service multiple_services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTodoRequest)

		id, err := service.CreateTodo(ctx, req.Text)

		if err != nil {
			if endpointErr := endpointError(nil); errors.As(err, &endpointErr) && endpointErr.EndpointError() {
				return CreateTodoResponse{
					Err: err,
					Id:  id,
				}, err
			}

			return CreateTodoResponse{
				Err: err,
				Id:  id,
			}, nil
		}

		return CreateTodoResponse{Id: id}, nil
	}
}

// OtherEndpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type OtherEndpoints struct {
	CreateTodo endpoint.Endpoint
}

// MakeOtherEndpoints returns a(n) OtherEndpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeOtherEndpoints(service multiple_services.OtherService, middleware ...endpoint.Middleware) OtherEndpoints {
	mw := kitxendpoint.Combine(middleware...)

	return OtherEndpoints{CreateTodo: kitxendpoint.OperationNameMiddleware("multiple_services.Other.CreateTodo")(mw(MakeCreateTodoOtherEndpoint(service)))}
}

// TraceOtherEndpoints returns a(n) OtherEndpoints struct where each endpoint is wrapped with a tracing middleware.
func TraceOtherEndpoints(endpoints OtherEndpoints) OtherEndpoints {
	return OtherEndpoints{CreateTodo: kitoc.TraceEndpoint("multiple_services.Other.CreateTodo")(endpoints.CreateTodo)}
}

// CreateTodoOtherRequest is a request struct for CreateTodo endpoint.
type CreateTodoOtherRequest struct{}

// CreateTodoOtherResponse is a response struct for CreateTodo endpoint.
type CreateTodoOtherResponse struct {
	Err error
}

func (r CreateTodoOtherResponse) Failed() error {
	return r.Err
}

// MakeCreateTodoOtherEndpoint returns an endpoint for the matching method of the underlying service.
func MakeCreateTodoOtherEndpoint(service multiple_services.OtherService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := service.CreateTodo(ctx)

		if err != nil {
			if endpointErr := endpointError(nil); errors.As(err, &endpointErr) && endpointErr.EndpointError() {
				return CreateTodoOtherResponse{Err: err}, err
			}

			return CreateTodoOtherResponse{Err: err}, nil
		}

		return CreateTodoOtherResponse{}, nil
	}
}

// AnotherEndpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type AnotherEndpoints struct {
	CreateTodo endpoint.Endpoint
}

// MakeAnotherEndpoints returns a(n) AnotherEndpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeAnotherEndpoints(service multiple_services.Another, middleware ...endpoint.Middleware) AnotherEndpoints {
	mw := kitxendpoint.Combine(middleware...)

	return AnotherEndpoints{CreateTodo: kitxendpoint.OperationNameMiddleware("multiple_services.Another.CreateTodo")(mw(MakeCreateTodoAnotherEndpoint(service)))}
}

// TraceAnotherEndpoints returns a(n) AnotherEndpoints struct where each endpoint is wrapped with a tracing middleware.
func TraceAnotherEndpoints(endpoints AnotherEndpoints) AnotherEndpoints {
	return AnotherEndpoints{CreateTodo: kitoc.TraceEndpoint("multiple_services.Another.CreateTodo")(endpoints.CreateTodo)}
}

// CreateTodoAnotherRequest is a request struct for CreateTodo endpoint.
type CreateTodoAnotherRequest struct{}

// CreateTodoAnotherResponse is a response struct for CreateTodo endpoint.
type CreateTodoAnotherResponse struct {
	Err error
}

func (r CreateTodoAnotherResponse) Failed() error {
	return r.Err
}

// MakeCreateTodoAnotherEndpoint returns an endpoint for the matching method of the underlying service.
func MakeCreateTodoAnotherEndpoint(service multiple_services.Another) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := service.CreateTodo(ctx)

		if err != nil {
			if endpointErr := endpointError(nil); errors.As(err, &endpointErr) && endpointErr.EndpointError() {
				return CreateTodoAnotherResponse{Err: err}, err
			}

			return CreateTodoAnotherResponse{Err: err}, nil
		}

		return CreateTodoAnotherResponse{}, nil
	}
}
