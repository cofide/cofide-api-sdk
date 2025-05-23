// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file proto/connect/trust_zone_service/v1alpha1/trust_zone_service.proto (package proto.connect.trust_zone_service.v1alpha1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { Bundle } from "../../../spire/api/types/bundle_pb";
import { file_proto_spire_api_types_bundle } from "../../../spire/api/types/bundle_pb";
import type { TrustZone } from "../../../trust_zone/v1alpha1/trust_zone_pb";
import { file_proto_trust_zone_v1alpha1_trust_zone } from "../../../trust_zone/v1alpha1/trust_zone_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file proto/connect/trust_zone_service/v1alpha1/trust_zone_service.proto.
 */
export const file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service: GenFile = /*@__PURE__*/
  fileDesc("CkJwcm90by9jb25uZWN0L3RydXN0X3pvbmVfc2VydmljZS92MWFscGhhMS90cnVzdF96b25lX3NlcnZpY2UucHJvdG8SKXByb3RvLmNvbm5lY3QudHJ1c3Rfem9uZV9zZXJ2aWNlLnYxYWxwaGExIkQKBUFnZW50EhAKCGFnZW50X2lkGAEgASgJEhIKCmNsdXN0ZXJfaWQYAiABKAkSFQoNdHJ1c3Rfem9uZV9pZBgDIAEoCSJSChZDcmVhdGVUcnVzdFpvbmVSZXF1ZXN0EjgKCnRydXN0X3pvbmUYASABKAsyJC5wcm90by50cnVzdF96b25lLnYxYWxwaGExLlRydXN0Wm9uZSJTChdDcmVhdGVUcnVzdFpvbmVSZXNwb25zZRI4Cgp0cnVzdF96b25lGAEgASgLMiQucHJvdG8udHJ1c3Rfem9uZS52MWFscGhhMS5UcnVzdFpvbmUiRwoXRGVzdHJveVRydXN0Wm9uZVJlcXVlc3QSGgoNdHJ1c3Rfem9uZV9pZBgBIAEoCUgAiAEBQhAKDl90cnVzdF96b25lX2lkIhoKGERlc3Ryb3lUcnVzdFpvbmVSZXNwb25zZSJDChNHZXRUcnVzdFpvbmVSZXF1ZXN0EhoKDXRydXN0X3pvbmVfaWQYASABKAlIAIgBAUIQCg5fdHJ1c3Rfem9uZV9pZCJkChRHZXRUcnVzdFpvbmVSZXNwb25zZRI9Cgp0cnVzdF96b25lGAEgASgLMiQucHJvdG8udHJ1c3Rfem9uZS52MWFscGhhMS5UcnVzdFpvbmVIAIgBAUINCgtfdHJ1c3Rfem9uZSLyAQoVTGlzdFRydXN0Wm9uZXNSZXF1ZXN0ElwKBmZpbHRlchgBIAEoCzJHLnByb3RvLmNvbm5lY3QudHJ1c3Rfem9uZV9zZXJ2aWNlLnYxYWxwaGExLkxpc3RUcnVzdFpvbmVzUmVxdWVzdC5GaWx0ZXJIAIgBARpwCgZGaWx0ZXISEQoEbmFtZRgBIAEoCUgAiAEBEhMKBm9yZ19pZBgCIAEoCUgBiAEBEhkKDHRydXN0X2RvbWFpbhgDIAEoCUgCiAEBQgcKBV9uYW1lQgkKB19vcmdfaWRCDwoNX3RydXN0X2RvbWFpbkIJCgdfZmlsdGVyIlMKFkxpc3RUcnVzdFpvbmVzUmVzcG9uc2USOQoLdHJ1c3Rfem9uZXMYASADKAsyJC5wcm90by50cnVzdF96b25lLnYxYWxwaGExLlRydXN0Wm9uZSJSChZVcGRhdGVUcnVzdFpvbmVSZXF1ZXN0EjgKCnRydXN0X3pvbmUYASABKAsyJC5wcm90by50cnVzdF96b25lLnYxYWxwaGExLlRydXN0Wm9uZSJTChdVcGRhdGVUcnVzdFpvbmVSZXNwb25zZRI4Cgp0cnVzdF96b25lGAEgASgLMiQucHJvdG8udHJ1c3Rfem9uZS52MWFscGhhMS5UcnVzdFpvbmUilQEKFFJlZ2lzdGVyQWdlbnRSZXF1ZXN0Ej8KBWFnZW50GAEgASgLMjAucHJvdG8uY29ubmVjdC50cnVzdF96b25lX3NlcnZpY2UudjFhbHBoYTEuQWdlbnQSEwoLYWdlbnRfdG9rZW4YAiABKAkSJwoGYnVuZGxlGAMgASgLMhcuc3BpcmUuYXBpLnR5cGVzLkJ1bmRsZSIpChVSZWdpc3RlckFnZW50UmVzcG9uc2USEAoIYWdlbnRfaWQYAyABKAkysQcKEFRydXN0Wm9uZVNlcnZpY2USmgEKD0NyZWF0ZVRydXN0Wm9uZRJBLnByb3RvLmNvbm5lY3QudHJ1c3Rfem9uZV9zZXJ2aWNlLnYxYWxwaGExLkNyZWF0ZVRydXN0Wm9uZVJlcXVlc3QaQi5wcm90by5jb25uZWN0LnRydXN0X3pvbmVfc2VydmljZS52MWFscGhhMS5DcmVhdGVUcnVzdFpvbmVSZXNwb25zZSIAEp0BChBEZXN0cm95VHJ1c3Rab25lEkIucHJvdG8uY29ubmVjdC50cnVzdF96b25lX3NlcnZpY2UudjFhbHBoYTEuRGVzdHJveVRydXN0Wm9uZVJlcXVlc3QaQy5wcm90by5jb25uZWN0LnRydXN0X3pvbmVfc2VydmljZS52MWFscGhhMS5EZXN0cm95VHJ1c3Rab25lUmVzcG9uc2UiABKRAQoMR2V0VHJ1c3Rab25lEj4ucHJvdG8uY29ubmVjdC50cnVzdF96b25lX3NlcnZpY2UudjFhbHBoYTEuR2V0VHJ1c3Rab25lUmVxdWVzdBo/LnByb3RvLmNvbm5lY3QudHJ1c3Rfem9uZV9zZXJ2aWNlLnYxYWxwaGExLkdldFRydXN0Wm9uZVJlc3BvbnNlIgASlwEKDkxpc3RUcnVzdFpvbmVzEkAucHJvdG8uY29ubmVjdC50cnVzdF96b25lX3NlcnZpY2UudjFhbHBoYTEuTGlzdFRydXN0Wm9uZXNSZXF1ZXN0GkEucHJvdG8uY29ubmVjdC50cnVzdF96b25lX3NlcnZpY2UudjFhbHBoYTEuTGlzdFRydXN0Wm9uZXNSZXNwb25zZSIAEpoBCg9VcGRhdGVUcnVzdFpvbmUSQS5wcm90by5jb25uZWN0LnRydXN0X3pvbmVfc2VydmljZS52MWFscGhhMS5VcGRhdGVUcnVzdFpvbmVSZXF1ZXN0GkIucHJvdG8uY29ubmVjdC50cnVzdF96b25lX3NlcnZpY2UudjFhbHBoYTEuVXBkYXRlVHJ1c3Rab25lUmVzcG9uc2UiABKUAQoNUmVnaXN0ZXJBZ2VudBI/LnByb3RvLmNvbm5lY3QudHJ1c3Rfem9uZV9zZXJ2aWNlLnYxYWxwaGExLlJlZ2lzdGVyQWdlbnRSZXF1ZXN0GkAucHJvdG8uY29ubmVjdC50cnVzdF96b25lX3NlcnZpY2UudjFhbHBoYTEuUmVnaXN0ZXJBZ2VudFJlc3BvbnNlIgBCU1pRZ2l0aHViLmNvbS9jb2ZpZGUvY29maWRlLWFwaS1zZGsvZ2VuL2dvL3Byb3RvL2Nvbm5lY3QvdHJ1c3Rfem9uZV9zZXJ2aWNlL3YxYWxwaGExYgZwcm90bzM", [file_proto_spire_api_types_bundle, file_proto_trust_zone_v1alpha1_trust_zone]);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.Agent
 */
export type Agent = Message<"proto.connect.trust_zone_service.v1alpha1.Agent"> & {
  /**
   * @generated from field: string agent_id = 1;
   */
  agentId: string;

  /**
   * @generated from field: string cluster_id = 2;
   */
  clusterId: string;

  /**
   * @generated from field: string trust_zone_id = 3;
   */
  trustZoneId: string;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.Agent.
 * Use `create(AgentSchema)` to create a new message.
 */
export const AgentSchema: GenMessage<Agent> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 0);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.CreateTrustZoneRequest
 */
export type CreateTrustZoneRequest = Message<"proto.connect.trust_zone_service.v1alpha1.CreateTrustZoneRequest"> & {
  /**
   * @generated from field: proto.trust_zone.v1alpha1.TrustZone trust_zone = 1;
   */
  trustZone?: TrustZone;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.CreateTrustZoneRequest.
 * Use `create(CreateTrustZoneRequestSchema)` to create a new message.
 */
export const CreateTrustZoneRequestSchema: GenMessage<CreateTrustZoneRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 1);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.CreateTrustZoneResponse
 */
export type CreateTrustZoneResponse = Message<"proto.connect.trust_zone_service.v1alpha1.CreateTrustZoneResponse"> & {
  /**
   * @generated from field: proto.trust_zone.v1alpha1.TrustZone trust_zone = 1;
   */
  trustZone?: TrustZone;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.CreateTrustZoneResponse.
 * Use `create(CreateTrustZoneResponseSchema)` to create a new message.
 */
export const CreateTrustZoneResponseSchema: GenMessage<CreateTrustZoneResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 2);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.DestroyTrustZoneRequest
 */
export type DestroyTrustZoneRequest = Message<"proto.connect.trust_zone_service.v1alpha1.DestroyTrustZoneRequest"> & {
  /**
   * @generated from field: optional string trust_zone_id = 1;
   */
  trustZoneId?: string;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.DestroyTrustZoneRequest.
 * Use `create(DestroyTrustZoneRequestSchema)` to create a new message.
 */
export const DestroyTrustZoneRequestSchema: GenMessage<DestroyTrustZoneRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 3);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.DestroyTrustZoneResponse
 */
export type DestroyTrustZoneResponse = Message<"proto.connect.trust_zone_service.v1alpha1.DestroyTrustZoneResponse"> & {
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.DestroyTrustZoneResponse.
 * Use `create(DestroyTrustZoneResponseSchema)` to create a new message.
 */
export const DestroyTrustZoneResponseSchema: GenMessage<DestroyTrustZoneResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 4);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.GetTrustZoneRequest
 */
export type GetTrustZoneRequest = Message<"proto.connect.trust_zone_service.v1alpha1.GetTrustZoneRequest"> & {
  /**
   * @generated from field: optional string trust_zone_id = 1;
   */
  trustZoneId?: string;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.GetTrustZoneRequest.
 * Use `create(GetTrustZoneRequestSchema)` to create a new message.
 */
export const GetTrustZoneRequestSchema: GenMessage<GetTrustZoneRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 5);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.GetTrustZoneResponse
 */
export type GetTrustZoneResponse = Message<"proto.connect.trust_zone_service.v1alpha1.GetTrustZoneResponse"> & {
  /**
   * @generated from field: optional proto.trust_zone.v1alpha1.TrustZone trust_zone = 1;
   */
  trustZone?: TrustZone;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.GetTrustZoneResponse.
 * Use `create(GetTrustZoneResponseSchema)` to create a new message.
 */
export const GetTrustZoneResponseSchema: GenMessage<GetTrustZoneResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 6);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.ListTrustZonesRequest
 */
export type ListTrustZonesRequest = Message<"proto.connect.trust_zone_service.v1alpha1.ListTrustZonesRequest"> & {
  /**
   * @generated from field: optional proto.connect.trust_zone_service.v1alpha1.ListTrustZonesRequest.Filter filter = 1;
   */
  filter?: ListTrustZonesRequest_Filter;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.ListTrustZonesRequest.
 * Use `create(ListTrustZonesRequestSchema)` to create a new message.
 */
export const ListTrustZonesRequestSchema: GenMessage<ListTrustZonesRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 7);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.ListTrustZonesRequest.Filter
 */
export type ListTrustZonesRequest_Filter = Message<"proto.connect.trust_zone_service.v1alpha1.ListTrustZonesRequest.Filter"> & {
  /**
   * @generated from field: optional string name = 1;
   */
  name?: string;

  /**
   * @generated from field: optional string org_id = 2;
   */
  orgId?: string;

  /**
   * @generated from field: optional string trust_domain = 3;
   */
  trustDomain?: string;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.ListTrustZonesRequest.Filter.
 * Use `create(ListTrustZonesRequest_FilterSchema)` to create a new message.
 */
export const ListTrustZonesRequest_FilterSchema: GenMessage<ListTrustZonesRequest_Filter> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 7, 0);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.ListTrustZonesResponse
 */
export type ListTrustZonesResponse = Message<"proto.connect.trust_zone_service.v1alpha1.ListTrustZonesResponse"> & {
  /**
   * @generated from field: repeated proto.trust_zone.v1alpha1.TrustZone trust_zones = 1;
   */
  trustZones: TrustZone[];
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.ListTrustZonesResponse.
 * Use `create(ListTrustZonesResponseSchema)` to create a new message.
 */
export const ListTrustZonesResponseSchema: GenMessage<ListTrustZonesResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 8);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.UpdateTrustZoneRequest
 */
export type UpdateTrustZoneRequest = Message<"proto.connect.trust_zone_service.v1alpha1.UpdateTrustZoneRequest"> & {
  /**
   * @generated from field: proto.trust_zone.v1alpha1.TrustZone trust_zone = 1;
   */
  trustZone?: TrustZone;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.UpdateTrustZoneRequest.
 * Use `create(UpdateTrustZoneRequestSchema)` to create a new message.
 */
export const UpdateTrustZoneRequestSchema: GenMessage<UpdateTrustZoneRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 9);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.UpdateTrustZoneResponse
 */
export type UpdateTrustZoneResponse = Message<"proto.connect.trust_zone_service.v1alpha1.UpdateTrustZoneResponse"> & {
  /**
   * @generated from field: proto.trust_zone.v1alpha1.TrustZone trust_zone = 1;
   */
  trustZone?: TrustZone;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.UpdateTrustZoneResponse.
 * Use `create(UpdateTrustZoneResponseSchema)` to create a new message.
 */
export const UpdateTrustZoneResponseSchema: GenMessage<UpdateTrustZoneResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 10);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.RegisterAgentRequest
 */
export type RegisterAgentRequest = Message<"proto.connect.trust_zone_service.v1alpha1.RegisterAgentRequest"> & {
  /**
   * @generated from field: proto.connect.trust_zone_service.v1alpha1.Agent agent = 1;
   */
  agent?: Agent;

  /**
   * @generated from field: string agent_token = 2;
   */
  agentToken: string;

  /**
   * @generated from field: spire.api.types.Bundle bundle = 3;
   */
  bundle?: Bundle;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.RegisterAgentRequest.
 * Use `create(RegisterAgentRequestSchema)` to create a new message.
 */
export const RegisterAgentRequestSchema: GenMessage<RegisterAgentRequest> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 11);

/**
 * @generated from message proto.connect.trust_zone_service.v1alpha1.RegisterAgentResponse
 */
export type RegisterAgentResponse = Message<"proto.connect.trust_zone_service.v1alpha1.RegisterAgentResponse"> & {
  /**
   * @generated from field: string agent_id = 3;
   */
  agentId: string;
};

/**
 * Describes the message proto.connect.trust_zone_service.v1alpha1.RegisterAgentResponse.
 * Use `create(RegisterAgentResponseSchema)` to create a new message.
 */
export const RegisterAgentResponseSchema: GenMessage<RegisterAgentResponse> = /*@__PURE__*/
  messageDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 12);

/**
 * @generated from service proto.connect.trust_zone_service.v1alpha1.TrustZoneService
 */
export const TrustZoneService: GenService<{
  /**
   * @generated from rpc proto.connect.trust_zone_service.v1alpha1.TrustZoneService.CreateTrustZone
   */
  createTrustZone: {
    methodKind: "unary";
    input: typeof CreateTrustZoneRequestSchema;
    output: typeof CreateTrustZoneResponseSchema;
  },
  /**
   * @generated from rpc proto.connect.trust_zone_service.v1alpha1.TrustZoneService.DestroyTrustZone
   */
  destroyTrustZone: {
    methodKind: "unary";
    input: typeof DestroyTrustZoneRequestSchema;
    output: typeof DestroyTrustZoneResponseSchema;
  },
  /**
   * @generated from rpc proto.connect.trust_zone_service.v1alpha1.TrustZoneService.GetTrustZone
   */
  getTrustZone: {
    methodKind: "unary";
    input: typeof GetTrustZoneRequestSchema;
    output: typeof GetTrustZoneResponseSchema;
  },
  /**
   * @generated from rpc proto.connect.trust_zone_service.v1alpha1.TrustZoneService.ListTrustZones
   */
  listTrustZones: {
    methodKind: "unary";
    input: typeof ListTrustZonesRequestSchema;
    output: typeof ListTrustZonesResponseSchema;
  },
  /**
   * @generated from rpc proto.connect.trust_zone_service.v1alpha1.TrustZoneService.UpdateTrustZone
   */
  updateTrustZone: {
    methodKind: "unary";
    input: typeof UpdateTrustZoneRequestSchema;
    output: typeof UpdateTrustZoneResponseSchema;
  },
  /**
   * @generated from rpc proto.connect.trust_zone_service.v1alpha1.TrustZoneService.RegisterAgent
   */
  registerAgent: {
    methodKind: "unary";
    input: typeof RegisterAgentRequestSchema;
    output: typeof RegisterAgentResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_proto_connect_trust_zone_service_v1alpha1_trust_zone_service, 0);

