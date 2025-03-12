/**
 * @fileoverview gRPC-Web generated client stub for proto.connect.federation_service.v1alpha1
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v0.0.0
// source: proto/connect/federation_service/v1alpha1/federation_service.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as proto_connect_federation_service_v1alpha1_federation_service_pb from '../../../../proto/connect/federation_service/v1alpha1/federation_service_pb'; // proto import: "proto/connect/federation_service/v1alpha1/federation_service.proto"


export class FederationServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorCreateFederation = new grpcWeb.MethodDescriptor(
    '/proto.connect.federation_service.v1alpha1.FederationService/CreateFederation',
    grpcWeb.MethodType.UNARY,
    proto_connect_federation_service_v1alpha1_federation_service_pb.CreateFederationRequest,
    proto_connect_federation_service_v1alpha1_federation_service_pb.CreateFederationResponse,
    (request: proto_connect_federation_service_v1alpha1_federation_service_pb.CreateFederationRequest) => {
      return request.serializeBinary();
    },
    proto_connect_federation_service_v1alpha1_federation_service_pb.CreateFederationResponse.deserializeBinary
  );

  createFederation(
    request: proto_connect_federation_service_v1alpha1_federation_service_pb.CreateFederationRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_connect_federation_service_v1alpha1_federation_service_pb.CreateFederationResponse>;

  createFederation(
    request: proto_connect_federation_service_v1alpha1_federation_service_pb.CreateFederationRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_connect_federation_service_v1alpha1_federation_service_pb.CreateFederationResponse) => void): grpcWeb.ClientReadableStream<proto_connect_federation_service_v1alpha1_federation_service_pb.CreateFederationResponse>;

  createFederation(
    request: proto_connect_federation_service_v1alpha1_federation_service_pb.CreateFederationRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_connect_federation_service_v1alpha1_federation_service_pb.CreateFederationResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.connect.federation_service.v1alpha1.FederationService/CreateFederation',
        request,
        metadata || {},
        this.methodDescriptorCreateFederation,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.connect.federation_service.v1alpha1.FederationService/CreateFederation',
    request,
    metadata || {},
    this.methodDescriptorCreateFederation);
  }

  methodDescriptorDestroyFederation = new grpcWeb.MethodDescriptor(
    '/proto.connect.federation_service.v1alpha1.FederationService/DestroyFederation',
    grpcWeb.MethodType.UNARY,
    proto_connect_federation_service_v1alpha1_federation_service_pb.DestroyFederationRequest,
    proto_connect_federation_service_v1alpha1_federation_service_pb.DestroyFederationResponse,
    (request: proto_connect_federation_service_v1alpha1_federation_service_pb.DestroyFederationRequest) => {
      return request.serializeBinary();
    },
    proto_connect_federation_service_v1alpha1_federation_service_pb.DestroyFederationResponse.deserializeBinary
  );

  destroyFederation(
    request: proto_connect_federation_service_v1alpha1_federation_service_pb.DestroyFederationRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_connect_federation_service_v1alpha1_federation_service_pb.DestroyFederationResponse>;

  destroyFederation(
    request: proto_connect_federation_service_v1alpha1_federation_service_pb.DestroyFederationRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_connect_federation_service_v1alpha1_federation_service_pb.DestroyFederationResponse) => void): grpcWeb.ClientReadableStream<proto_connect_federation_service_v1alpha1_federation_service_pb.DestroyFederationResponse>;

  destroyFederation(
    request: proto_connect_federation_service_v1alpha1_federation_service_pb.DestroyFederationRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_connect_federation_service_v1alpha1_federation_service_pb.DestroyFederationResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.connect.federation_service.v1alpha1.FederationService/DestroyFederation',
        request,
        metadata || {},
        this.methodDescriptorDestroyFederation,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.connect.federation_service.v1alpha1.FederationService/DestroyFederation',
    request,
    metadata || {},
    this.methodDescriptorDestroyFederation);
  }

  methodDescriptorListFederations = new grpcWeb.MethodDescriptor(
    '/proto.connect.federation_service.v1alpha1.FederationService/ListFederations',
    grpcWeb.MethodType.UNARY,
    proto_connect_federation_service_v1alpha1_federation_service_pb.ListFederationsRequest,
    proto_connect_federation_service_v1alpha1_federation_service_pb.ListFederationsResponse,
    (request: proto_connect_federation_service_v1alpha1_federation_service_pb.ListFederationsRequest) => {
      return request.serializeBinary();
    },
    proto_connect_federation_service_v1alpha1_federation_service_pb.ListFederationsResponse.deserializeBinary
  );

  listFederations(
    request: proto_connect_federation_service_v1alpha1_federation_service_pb.ListFederationsRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_connect_federation_service_v1alpha1_federation_service_pb.ListFederationsResponse>;

  listFederations(
    request: proto_connect_federation_service_v1alpha1_federation_service_pb.ListFederationsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_connect_federation_service_v1alpha1_federation_service_pb.ListFederationsResponse) => void): grpcWeb.ClientReadableStream<proto_connect_federation_service_v1alpha1_federation_service_pb.ListFederationsResponse>;

  listFederations(
    request: proto_connect_federation_service_v1alpha1_federation_service_pb.ListFederationsRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_connect_federation_service_v1alpha1_federation_service_pb.ListFederationsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.connect.federation_service.v1alpha1.FederationService/ListFederations',
        request,
        metadata || {},
        this.methodDescriptorListFederations,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.connect.federation_service.v1alpha1.FederationService/ListFederations',
    request,
    metadata || {},
    this.methodDescriptorListFederations);
  }

}

