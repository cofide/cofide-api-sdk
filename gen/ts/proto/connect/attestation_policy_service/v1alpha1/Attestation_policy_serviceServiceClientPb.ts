/**
 * @fileoverview gRPC-Web generated client stub for proto.connect.attestation_policy_service.v1alpha1
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v0.0.0
// source: proto/connect/attestation_policy_service/v1alpha1/attestation_policy_service.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb from '../../../../proto/connect/attestation_policy_service/v1alpha1/attestation_policy_service_pb'; // proto import: "proto/connect/attestation_policy_service/v1alpha1/attestation_policy_service.proto"


export class AttestationPolicyServiceClient {
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

  methodDescriptorCreateAttestationPolicy = new grpcWeb.MethodDescriptor(
    '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/CreateAttestationPolicy',
    grpcWeb.MethodType.UNARY,
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.CreateAttestationPolicyRequest,
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.CreateAttestationPolicyResponse,
    (request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.CreateAttestationPolicyRequest) => {
      return request.serializeBinary();
    },
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.CreateAttestationPolicyResponse.deserializeBinary
  );

  createAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.CreateAttestationPolicyRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.CreateAttestationPolicyResponse>;

  createAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.CreateAttestationPolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.CreateAttestationPolicyResponse) => void): grpcWeb.ClientReadableStream<proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.CreateAttestationPolicyResponse>;

  createAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.CreateAttestationPolicyRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.CreateAttestationPolicyResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/CreateAttestationPolicy',
        request,
        metadata || {},
        this.methodDescriptorCreateAttestationPolicy,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/CreateAttestationPolicy',
    request,
    metadata || {},
    this.methodDescriptorCreateAttestationPolicy);
  }

  methodDescriptorDestroyAttestationPolicy = new grpcWeb.MethodDescriptor(
    '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/DestroyAttestationPolicy',
    grpcWeb.MethodType.UNARY,
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.DestroyAttestationPolicyRequest,
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.DestroyAttestationPolicyResponse,
    (request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.DestroyAttestationPolicyRequest) => {
      return request.serializeBinary();
    },
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.DestroyAttestationPolicyResponse.deserializeBinary
  );

  destroyAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.DestroyAttestationPolicyRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.DestroyAttestationPolicyResponse>;

  destroyAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.DestroyAttestationPolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.DestroyAttestationPolicyResponse) => void): grpcWeb.ClientReadableStream<proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.DestroyAttestationPolicyResponse>;

  destroyAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.DestroyAttestationPolicyRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.DestroyAttestationPolicyResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/DestroyAttestationPolicy',
        request,
        metadata || {},
        this.methodDescriptorDestroyAttestationPolicy,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/DestroyAttestationPolicy',
    request,
    metadata || {},
    this.methodDescriptorDestroyAttestationPolicy);
  }

  methodDescriptorGetAttestationPolicy = new grpcWeb.MethodDescriptor(
    '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/GetAttestationPolicy',
    grpcWeb.MethodType.UNARY,
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.GetAttestationPolicyRequest,
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.GetAttestationPolicyResponse,
    (request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.GetAttestationPolicyRequest) => {
      return request.serializeBinary();
    },
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.GetAttestationPolicyResponse.deserializeBinary
  );

  getAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.GetAttestationPolicyRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.GetAttestationPolicyResponse>;

  getAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.GetAttestationPolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.GetAttestationPolicyResponse) => void): grpcWeb.ClientReadableStream<proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.GetAttestationPolicyResponse>;

  getAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.GetAttestationPolicyRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.GetAttestationPolicyResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/GetAttestationPolicy',
        request,
        metadata || {},
        this.methodDescriptorGetAttestationPolicy,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/GetAttestationPolicy',
    request,
    metadata || {},
    this.methodDescriptorGetAttestationPolicy);
  }

  methodDescriptorListAttestationPolicies = new grpcWeb.MethodDescriptor(
    '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/ListAttestationPolicies',
    grpcWeb.MethodType.UNARY,
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.ListAttestationPoliciesRequest,
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.ListAttestationPoliciesResponse,
    (request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.ListAttestationPoliciesRequest) => {
      return request.serializeBinary();
    },
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.ListAttestationPoliciesResponse.deserializeBinary
  );

  listAttestationPolicies(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.ListAttestationPoliciesRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.ListAttestationPoliciesResponse>;

  listAttestationPolicies(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.ListAttestationPoliciesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.ListAttestationPoliciesResponse) => void): grpcWeb.ClientReadableStream<proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.ListAttestationPoliciesResponse>;

  listAttestationPolicies(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.ListAttestationPoliciesRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.ListAttestationPoliciesResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/ListAttestationPolicies',
        request,
        metadata || {},
        this.methodDescriptorListAttestationPolicies,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/ListAttestationPolicies',
    request,
    metadata || {},
    this.methodDescriptorListAttestationPolicies);
  }

  methodDescriptorUpdateAttestationPolicy = new grpcWeb.MethodDescriptor(
    '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/UpdateAttestationPolicy',
    grpcWeb.MethodType.UNARY,
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.UpdateAttestationPolicyRequest,
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.UpdateAttestationPolicyResponse,
    (request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.UpdateAttestationPolicyRequest) => {
      return request.serializeBinary();
    },
    proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.UpdateAttestationPolicyResponse.deserializeBinary
  );

  updateAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.UpdateAttestationPolicyRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.UpdateAttestationPolicyResponse>;

  updateAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.UpdateAttestationPolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.UpdateAttestationPolicyResponse) => void): grpcWeb.ClientReadableStream<proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.UpdateAttestationPolicyResponse>;

  updateAttestationPolicy(
    request: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.UpdateAttestationPolicyRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_connect_attestation_policy_service_v1alpha1_attestation_policy_service_pb.UpdateAttestationPolicyResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/UpdateAttestationPolicy',
        request,
        metadata || {},
        this.methodDescriptorUpdateAttestationPolicy,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.connect.attestation_policy_service.v1alpha1.AttestationPolicyService/UpdateAttestationPolicy',
    request,
    metadata || {},
    this.methodDescriptorUpdateAttestationPolicy);
  }

}

