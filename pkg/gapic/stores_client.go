// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.


package gapic

import (
	"context"
	"fmt"
	"math"
	"net/url"

	storespb "github.com/bobadojo/go/pkg/stores/v1/storespb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
)

var newStoresClientHook clientHook

// StoresCallOptions contains the retry settings for each method of StoresClient.
type StoresCallOptions struct {
	FindStores []gax.CallOption
	GetStore []gax.CallOption
}

func defaultStoresGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint(":443"),
		internaloption.WithDefaultEndpointTemplate(":443"),
		internaloption.WithDefaultMTLSEndpoint(":443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultStoresCallOptions() *StoresCallOptions {
	return &StoresCallOptions{
		FindStores: []gax.CallOption{
		},
		GetStore: []gax.CallOption{
		},
	}
}

// internalStoresClient is an interface that defines the methods available from .
type internalStoresClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	FindStores(context.Context, *storespb.FindStoresRequest, ...gax.CallOption) (*storespb.FindStoresResponse, error)
	GetStore(context.Context, *storespb.GetStoreRequest, ...gax.CallOption) (*storespb.Store, error)
}

// StoresClient is a client for interacting with .
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A store locator API
type StoresClient struct {
	// The internal transport-dependent client.
	internalClient internalStoresClient

	// The call options for this service.
	CallOptions *StoresCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *StoresClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *StoresClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *StoresClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// FindStores returns a list of all stores in the store.
func (c *StoresClient) FindStores(ctx context.Context, req *storespb.FindStoresRequest, opts ...gax.CallOption) (*storespb.FindStoresResponse, error) {
	return c.internalClient.FindStores(ctx, req, opts...)
}

// GetStore returns a specific store.
func (c *StoresClient) GetStore(ctx context.Context, req *storespb.GetStoreRequest, opts ...gax.CallOption) (*storespb.Store, error) {
	return c.internalClient.GetStore(ctx, req, opts...)
}

// storesGRPCClient is a client for interacting with  over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type storesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing StoresClient
	CallOptions **StoresCallOptions

	// The gRPC API client.
	storesClient storespb.StoresClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewStoresClient creates a new stores client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A store locator API
func NewStoresClient(ctx context.Context, opts ...option.ClientOption) (*StoresClient, error) {
	clientOpts := defaultStoresGRPCClientOptions()
	if newStoresClientHook != nil {
		hookOpts, err := newStoresClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := StoresClient{CallOptions: defaultStoresCallOptions()}

	c := &storesGRPCClient{
		connPool:    connPool,
		storesClient: storespb.NewStoresClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *storesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *storesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *storesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *storesGRPCClient) FindStores(ctx context.Context, req *storespb.FindStoresRequest, opts ...gax.CallOption) (*storespb.FindStoresResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).FindStores[0:len((*c.CallOptions).FindStores):len((*c.CallOptions).FindStores)], opts...)
	var resp *storespb.FindStoresResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.storesClient.FindStores(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *storesGRPCClient) GetStore(ctx context.Context, req *storespb.GetStoreRequest, opts ...gax.CallOption) (*storespb.Store, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetStore[0:len((*c.CallOptions).GetStore):len((*c.CallOptions).GetStore)], opts...)
	var resp *storespb.Store
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.storesClient.GetStore(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}