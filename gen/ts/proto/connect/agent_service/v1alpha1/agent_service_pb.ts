// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file proto/connect/agent_service/v1alpha1/agent_service.proto (package proto.connect.agent_service.v1alpha1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { AgentStatus } from "../../../agent/v1alpha1/agent_pb";
import { file_proto_agent_v1alpha1_agent } from "../../../agent/v1alpha1/agent_pb";
import type { FederatedService } from "../../../federated_service/v1alpha1/federated_service_pb";
import { file_proto_federated_service_v1alpha1_federated_service } from "../../../federated_service/v1alpha1/federated_service_pb";
import type { Bundle } from "../../../spire/api/types/bundle_pb";
import { file_proto_spire_api_types_bundle } from "../../../spire/api/types/bundle_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file proto/connect/agent_service/v1alpha1/agent_service.proto.
 */
export const file_proto_connect_agent_service_v1alpha1_agent_service: GenFile = /*@__PURE__*/
  fileDesc("Cjhwcm90by9jb25uZWN0L2FnZW50X3NlcnZpY2UvdjFhbHBoYTEvYWdlbnRfc2VydmljZS5wcm90bxIkcHJvdG8uY29ubmVjdC5hZ2VudF9zZXJ2aWNlLnYxYWxwaGExInMKG0NyZWF0ZUFnZW50Sm9pblRva2VuUmVxdWVzdBIaCg10cnVzdF96b25lX2lkGAEgASgJSACIAQESFwoKY2x1c3Rlcl9pZBgCIAEoCUgBiAEBQhAKDl90cnVzdF96b25lX2lkQg0KC19jbHVzdGVyX2lkIkgKHENyZWF0ZUFnZW50Sm9pblRva2VuUmVzcG9uc2USGAoLYWdlbnRfdG9rZW4YASABKAlIAIgBAUIOCgxfYWdlbnRfdG9rZW4iXgocVXBkYXRlVHJ1c3Rab25lQnVuZGxlUmVxdWVzdBInCgZidW5kbGUYASABKAsyFy5zcGlyZS5hcGkudHlwZXMuQnVuZGxlEhUKDXRydXN0X3pvbmVfaWQYAiABKAkiHwodVXBkYXRlVHJ1c3Rab25lQnVuZGxlUmVzcG9uc2UiTQoYVXBkYXRlQWdlbnRTdGF0dXNSZXF1ZXN0EjEKBnN0YXR1cxgBIAEoCzIhLnByb3RvLmFnZW50LnYxYWxwaGExLkFnZW50U3RhdHVzIhsKGVVwZGF0ZUFnZW50U3RhdHVzUmVzcG9uc2UiZgofUmVnaXN0ZXJGZWRlcmF0ZWRTZXJ2aWNlUmVxdWVzdBJDCgdzZXJ2aWNlGAEgASgLMjIucHJvdG8uZmVkZXJhdGVkX3NlcnZpY2UudjFhbHBoYTEuRmVkZXJhdGVkU2VydmljZSI2CiBSZWdpc3RlckZlZGVyYXRlZFNlcnZpY2VSZXNwb25zZRISCgpzZXJ2aWNlX2lkGAEgASgJIjcKIURlcmVnaXN0ZXJGZWRlcmF0ZWRTZXJ2aWNlUmVxdWVzdBISCgpzZXJ2aWNlX2lkGAEgASgJIjgKIkRlcmVnaXN0ZXJGZWRlcmF0ZWRTZXJ2aWNlUmVzcG9uc2USEgoKc2VydmljZV9pZBgBIAEoCSJkCh1VcGRhdGVGZWRlcmF0ZWRTZXJ2aWNlUmVxdWVzdBJDCgdzZXJ2aWNlGAEgASgLMjIucHJvdG8uZmVkZXJhdGVkX3NlcnZpY2UudjFhbHBoYTEuRmVkZXJhdGVkU2VydmljZSI0Ch5VcGRhdGVGZWRlcmF0ZWRTZXJ2aWNlUmVzcG9uc2USEgoKc2VydmljZV9pZBgBIAEoCSIwChpHZXRGZWRlcmF0ZWRTZXJ2aWNlUmVxdWVzdBISCgpzZXJ2aWNlX2lkGAEgASgJImIKG0dldEZlZGVyYXRlZFNlcnZpY2VSZXNwb25zZRJDCgdzZXJ2aWNlGAEgASgLMjIucHJvdG8uZmVkZXJhdGVkX3NlcnZpY2UudjFhbHBoYTEuRmVkZXJhdGVkU2VydmljZSIuChxMaXN0RmVkZXJhdGVkU2VydmljZXNSZXF1ZXN0Eg4KBm9yZ19pZBgBIAEoCSJlCh1MaXN0RmVkZXJhdGVkU2VydmljZXNSZXNwb25zZRJECghzZXJ2aWNlcxgBIAMoCzIyLnByb3RvLmZlZGVyYXRlZF9zZXJ2aWNlLnYxYWxwaGExLkZlZGVyYXRlZFNlcnZpY2UyvAoKDEFnZW50U2VydmljZRKfAQoUQ3JlYXRlQWdlbnRKb2luVG9rZW4SQS5wcm90by5jb25uZWN0LmFnZW50X3NlcnZpY2UudjFhbHBoYTEuQ3JlYXRlQWdlbnRKb2luVG9rZW5SZXF1ZXN0GkIucHJvdG8uY29ubmVjdC5hZ2VudF9zZXJ2aWNlLnYxYWxwaGExLkNyZWF0ZUFnZW50Sm9pblRva2VuUmVzcG9uc2UiABKiAQoVVXBkYXRlVHJ1c3Rab25lQnVuZGxlEkIucHJvdG8uY29ubmVjdC5hZ2VudF9zZXJ2aWNlLnYxYWxwaGExLlVwZGF0ZVRydXN0Wm9uZUJ1bmRsZVJlcXVlc3QaQy5wcm90by5jb25uZWN0LmFnZW50X3NlcnZpY2UudjFhbHBoYTEuVXBkYXRlVHJ1c3Rab25lQnVuZGxlUmVzcG9uc2UiABKWAQoRVXBkYXRlQWdlbnRTdGF0dXMSPi5wcm90by5jb25uZWN0LmFnZW50X3NlcnZpY2UudjFhbHBoYTEuVXBkYXRlQWdlbnRTdGF0dXNSZXF1ZXN0Gj8ucHJvdG8uY29ubmVjdC5hZ2VudF9zZXJ2aWNlLnYxYWxwaGExLlVwZGF0ZUFnZW50U3RhdHVzUmVzcG9uc2UiABKrAQoYUmVnaXN0ZXJGZWRlcmF0ZWRTZXJ2aWNlEkUucHJvdG8uY29ubmVjdC5hZ2VudF9zZXJ2aWNlLnYxYWxwaGExLlJlZ2lzdGVyRmVkZXJhdGVkU2VydmljZVJlcXVlc3QaRi5wcm90by5jb25uZWN0LmFnZW50X3NlcnZpY2UudjFhbHBoYTEuUmVnaXN0ZXJGZWRlcmF0ZWRTZXJ2aWNlUmVzcG9uc2UiABKxAQoaRGVyZWdpc3RlckZlZGVyYXRlZFNlcnZpY2USRy5wcm90by5jb25uZWN0LmFnZW50X3NlcnZpY2UudjFhbHBoYTEuRGVyZWdpc3RlckZlZGVyYXRlZFNlcnZpY2VSZXF1ZXN0GkgucHJvdG8uY29ubmVjdC5hZ2VudF9zZXJ2aWNlLnYxYWxwaGExLkRlcmVnaXN0ZXJGZWRlcmF0ZWRTZXJ2aWNlUmVzcG9uc2UiABKlAQoWVXBkYXRlRmVkZXJhdGVkU2VydmljZRJDLnByb3RvLmNvbm5lY3QuYWdlbnRfc2VydmljZS52MWFscGhhMS5VcGRhdGVGZWRlcmF0ZWRTZXJ2aWNlUmVxdWVzdBpELnByb3RvLmNvbm5lY3QuYWdlbnRfc2VydmljZS52MWFscGhhMS5VcGRhdGVGZWRlcmF0ZWRTZXJ2aWNlUmVzcG9uc2UiABKcAQoTR2V0RmVkZXJhdGVkU2VydmljZRJALnByb3RvLmNvbm5lY3QuYWdlbnRfc2VydmljZS52MWFscGhhMS5HZXRGZWRlcmF0ZWRTZXJ2aWNlUmVxdWVzdBpBLnByb3RvLmNvbm5lY3QuYWdlbnRfc2VydmljZS52MWFscGhhMS5HZXRGZWRlcmF0ZWRTZXJ2aWNlUmVzcG9uc2UiABKiAQoVTGlzdEZlZGVyYXRlZFNlcnZpY2VzEkIucHJvdG8uY29ubmVjdC5hZ2VudF9zZXJ2aWNlLnYxYWxwaGExLkxpc3RGZWRlcmF0ZWRTZXJ2aWNlc1JlcXVlc3QaQy5wcm90by5jb25uZWN0LmFnZW50X3NlcnZpY2UudjFhbHBoYTEuTGlzdEZlZGVyYXRlZFNlcnZpY2VzUmVzcG9uc2UiAEJOWkxnaXRodWIuY29tL2NvZmlkZS9jb2ZpZGUtYXBpLXNkay9nZW4vZ28vcHJvdG8vY29ubmVjdC9hZ2VudF9zZXJ2aWNlL3YxYWxwaGExYgZwcm90bzM", [file_proto_agent_v1alpha1_agent, file_proto_federated_service_v1alpha1_federated_service, file_proto_spire_api_types_bundle]);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.CreateAgentJoinTokenRequest
 */
export type CreateAgentJoinTokenRequest = Message<"proto.connect.agent_service.v1alpha1.CreateAgentJoinTokenRequest"> & {
  /**
   * @generated from field: optional string trust_zone_id = 1;
   */
  trustZoneId?: string;

  /**
   * @generated from field: optional string cluster_id = 2;
   */
  clusterId?: string;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.CreateAgentJoinTokenRequest.
 * Use `create(CreateAgentJoinTokenRequestSchema)` to create a new message.
 */
export const CreateAgentJoinTokenRequestSchema: GenMessage<CreateAgentJoinTokenRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 0);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.CreateAgentJoinTokenResponse
 */
export type CreateAgentJoinTokenResponse = Message<"proto.connect.agent_service.v1alpha1.CreateAgentJoinTokenResponse"> & {
  /**
   * @generated from field: optional string agent_token = 1;
   */
  agentToken?: string;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.CreateAgentJoinTokenResponse.
 * Use `create(CreateAgentJoinTokenResponseSchema)` to create a new message.
 */
export const CreateAgentJoinTokenResponseSchema: GenMessage<CreateAgentJoinTokenResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 1);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.UpdateTrustZoneBundleRequest
 */
export type UpdateTrustZoneBundleRequest = Message<"proto.connect.agent_service.v1alpha1.UpdateTrustZoneBundleRequest"> & {
  /**
   * @generated from field: spire.api.types.Bundle bundle = 1;
   */
  bundle?: Bundle;

  /**
   * @generated from field: string trust_zone_id = 2;
   */
  trustZoneId: string;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.UpdateTrustZoneBundleRequest.
 * Use `create(UpdateTrustZoneBundleRequestSchema)` to create a new message.
 */
export const UpdateTrustZoneBundleRequestSchema: GenMessage<UpdateTrustZoneBundleRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 2);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.UpdateTrustZoneBundleResponse
 */
export type UpdateTrustZoneBundleResponse = Message<"proto.connect.agent_service.v1alpha1.UpdateTrustZoneBundleResponse"> & {
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.UpdateTrustZoneBundleResponse.
 * Use `create(UpdateTrustZoneBundleResponseSchema)` to create a new message.
 */
export const UpdateTrustZoneBundleResponseSchema: GenMessage<UpdateTrustZoneBundleResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 3);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.UpdateAgentStatusRequest
 */
export type UpdateAgentStatusRequest = Message<"proto.connect.agent_service.v1alpha1.UpdateAgentStatusRequest"> & {
  /**
   * @generated from field: proto.agent.v1alpha1.AgentStatus status = 1;
   */
  status?: AgentStatus;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.UpdateAgentStatusRequest.
 * Use `create(UpdateAgentStatusRequestSchema)` to create a new message.
 */
export const UpdateAgentStatusRequestSchema: GenMessage<UpdateAgentStatusRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 4);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.UpdateAgentStatusResponse
 */
export type UpdateAgentStatusResponse = Message<"proto.connect.agent_service.v1alpha1.UpdateAgentStatusResponse"> & {
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.UpdateAgentStatusResponse.
 * Use `create(UpdateAgentStatusResponseSchema)` to create a new message.
 */
export const UpdateAgentStatusResponseSchema: GenMessage<UpdateAgentStatusResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 5);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.RegisterFederatedServiceRequest
 */
export type RegisterFederatedServiceRequest = Message<"proto.connect.agent_service.v1alpha1.RegisterFederatedServiceRequest"> & {
  /**
   * @generated from field: proto.federated_service.v1alpha1.FederatedService service = 1;
   */
  service?: FederatedService;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.RegisterFederatedServiceRequest.
 * Use `create(RegisterFederatedServiceRequestSchema)` to create a new message.
 */
export const RegisterFederatedServiceRequestSchema: GenMessage<RegisterFederatedServiceRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 6);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.RegisterFederatedServiceResponse
 */
export type RegisterFederatedServiceResponse = Message<"proto.connect.agent_service.v1alpha1.RegisterFederatedServiceResponse"> & {
  /**
   * @generated from field: string service_id = 1;
   */
  serviceId: string;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.RegisterFederatedServiceResponse.
 * Use `create(RegisterFederatedServiceResponseSchema)` to create a new message.
 */
export const RegisterFederatedServiceResponseSchema: GenMessage<RegisterFederatedServiceResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 7);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.DeregisterFederatedServiceRequest
 */
export type DeregisterFederatedServiceRequest = Message<"proto.connect.agent_service.v1alpha1.DeregisterFederatedServiceRequest"> & {
  /**
   * @generated from field: string service_id = 1;
   */
  serviceId: string;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.DeregisterFederatedServiceRequest.
 * Use `create(DeregisterFederatedServiceRequestSchema)` to create a new message.
 */
export const DeregisterFederatedServiceRequestSchema: GenMessage<DeregisterFederatedServiceRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 8);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.DeregisterFederatedServiceResponse
 */
export type DeregisterFederatedServiceResponse = Message<"proto.connect.agent_service.v1alpha1.DeregisterFederatedServiceResponse"> & {
  /**
   * @generated from field: string service_id = 1;
   */
  serviceId: string;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.DeregisterFederatedServiceResponse.
 * Use `create(DeregisterFederatedServiceResponseSchema)` to create a new message.
 */
export const DeregisterFederatedServiceResponseSchema: GenMessage<DeregisterFederatedServiceResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 9);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.UpdateFederatedServiceRequest
 */
export type UpdateFederatedServiceRequest = Message<"proto.connect.agent_service.v1alpha1.UpdateFederatedServiceRequest"> & {
  /**
   * @generated from field: proto.federated_service.v1alpha1.FederatedService service = 1;
   */
  service?: FederatedService;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.UpdateFederatedServiceRequest.
 * Use `create(UpdateFederatedServiceRequestSchema)` to create a new message.
 */
export const UpdateFederatedServiceRequestSchema: GenMessage<UpdateFederatedServiceRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 10);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.UpdateFederatedServiceResponse
 */
export type UpdateFederatedServiceResponse = Message<"proto.connect.agent_service.v1alpha1.UpdateFederatedServiceResponse"> & {
  /**
   * @generated from field: string service_id = 1;
   */
  serviceId: string;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.UpdateFederatedServiceResponse.
 * Use `create(UpdateFederatedServiceResponseSchema)` to create a new message.
 */
export const UpdateFederatedServiceResponseSchema: GenMessage<UpdateFederatedServiceResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 11);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.GetFederatedServiceRequest
 */
export type GetFederatedServiceRequest = Message<"proto.connect.agent_service.v1alpha1.GetFederatedServiceRequest"> & {
  /**
   * @generated from field: string service_id = 1;
   */
  serviceId: string;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.GetFederatedServiceRequest.
 * Use `create(GetFederatedServiceRequestSchema)` to create a new message.
 */
export const GetFederatedServiceRequestSchema: GenMessage<GetFederatedServiceRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 12);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.GetFederatedServiceResponse
 */
export type GetFederatedServiceResponse = Message<"proto.connect.agent_service.v1alpha1.GetFederatedServiceResponse"> & {
  /**
   * @generated from field: proto.federated_service.v1alpha1.FederatedService service = 1;
   */
  service?: FederatedService;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.GetFederatedServiceResponse.
 * Use `create(GetFederatedServiceResponseSchema)` to create a new message.
 */
export const GetFederatedServiceResponseSchema: GenMessage<GetFederatedServiceResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 13);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.ListFederatedServicesRequest
 */
export type ListFederatedServicesRequest = Message<"proto.connect.agent_service.v1alpha1.ListFederatedServicesRequest"> & {
  /**
   * @generated from field: string org_id = 1;
   */
  orgId: string;
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.ListFederatedServicesRequest.
 * Use `create(ListFederatedServicesRequestSchema)` to create a new message.
 */
export const ListFederatedServicesRequestSchema: GenMessage<ListFederatedServicesRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 14);

/**
 * @generated from message proto.connect.agent_service.v1alpha1.ListFederatedServicesResponse
 */
export type ListFederatedServicesResponse = Message<"proto.connect.agent_service.v1alpha1.ListFederatedServicesResponse"> & {
  /**
   * @generated from field: repeated proto.federated_service.v1alpha1.FederatedService services = 1;
   */
  services: FederatedService[];
};

/**
 * Describes the message proto.connect.agent_service.v1alpha1.ListFederatedServicesResponse.
 * Use `create(ListFederatedServicesResponseSchema)` to create a new message.
 */
export const ListFederatedServicesResponseSchema: GenMessage<ListFederatedServicesResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 15);

/**
 * @generated from service proto.connect.agent_service.v1alpha1.AgentService
 */
export const AgentService: GenService<{
  /**
   * @generated from rpc proto.connect.agent_service.v1alpha1.AgentService.CreateAgentJoinToken
   */
  createAgentJoinToken: {
    methodKind: "unary";
    input: typeof CreateAgentJoinTokenRequestSchema;
    output: typeof CreateAgentJoinTokenResponseSchema;
  },
  /**
   * @generated from rpc proto.connect.agent_service.v1alpha1.AgentService.UpdateTrustZoneBundle
   */
  updateTrustZoneBundle: {
    methodKind: "unary";
    input: typeof UpdateTrustZoneBundleRequestSchema;
    output: typeof UpdateTrustZoneBundleResponseSchema;
  },
  /**
   * @generated from rpc proto.connect.agent_service.v1alpha1.AgentService.UpdateAgentStatus
   */
  updateAgentStatus: {
    methodKind: "unary";
    input: typeof UpdateAgentStatusRequestSchema;
    output: typeof UpdateAgentStatusResponseSchema;
  },
  /**
   * DEPRECATED: Federated service RPCs will move to a separate service.
   *
   * @generated from rpc proto.connect.agent_service.v1alpha1.AgentService.RegisterFederatedService
   */
  registerFederatedService: {
    methodKind: "unary";
    input: typeof RegisterFederatedServiceRequestSchema;
    output: typeof RegisterFederatedServiceResponseSchema;
  },
  /**
   * @generated from rpc proto.connect.agent_service.v1alpha1.AgentService.DeregisterFederatedService
   */
  deregisterFederatedService: {
    methodKind: "unary";
    input: typeof DeregisterFederatedServiceRequestSchema;
    output: typeof DeregisterFederatedServiceResponseSchema;
  },
  /**
   * @generated from rpc proto.connect.agent_service.v1alpha1.AgentService.UpdateFederatedService
   */
  updateFederatedService: {
    methodKind: "unary";
    input: typeof UpdateFederatedServiceRequestSchema;
    output: typeof UpdateFederatedServiceResponseSchema;
  },
  /**
   * @generated from rpc proto.connect.agent_service.v1alpha1.AgentService.GetFederatedService
   */
  getFederatedService: {
    methodKind: "unary";
    input: typeof GetFederatedServiceRequestSchema;
    output: typeof GetFederatedServiceResponseSchema;
  },
  /**
   * @generated from rpc proto.connect.agent_service.v1alpha1.AgentService.ListFederatedServices
   */
  listFederatedServices: {
    methodKind: "unary";
    input: typeof ListFederatedServicesRequestSchema;
    output: typeof ListFederatedServicesResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_proto_connect_agent_service_v1alpha1_agent_service, 0);

