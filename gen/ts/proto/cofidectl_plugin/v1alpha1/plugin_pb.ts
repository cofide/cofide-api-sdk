// Copyright 2024 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file proto/cofidectl_plugin/v1alpha1/plugin.proto (package proto.cofidectl_plugin.v1alpha1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { APBinding } from "../../ap_binding/v1alpha1/ap_binding_pb";
import { file_proto_ap_binding_v1alpha1_ap_binding } from "../../ap_binding/v1alpha1/ap_binding_pb";
import type { AttestationPolicy } from "../../attestation_policy/v1alpha1/attestation_policy_pb";
import { file_proto_attestation_policy_v1alpha1_attestation_policy } from "../../attestation_policy/v1alpha1/attestation_policy_pb";
import type { Cluster } from "../../cluster/v1alpha1/cluster_pb";
import { file_proto_cluster_v1alpha1_cluster } from "../../cluster/v1alpha1/cluster_pb";
import type { Federation } from "../../federation/v1alpha1/federation_pb";
import { file_proto_federation_v1alpha1_federation } from "../../federation/v1alpha1/federation_pb";
import type { TrustZone } from "../../trust_zone/v1alpha1/trust_zone_pb";
import { file_proto_trust_zone_v1alpha1_trust_zone } from "../../trust_zone/v1alpha1/trust_zone_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file proto/cofidectl_plugin/v1alpha1/plugin.proto.
 */
export const file_proto_cofidectl_plugin_v1alpha1_plugin: GenFile = /*@__PURE__*/
  fileDesc("Cixwcm90by9jb2ZpZGVjdGxfcGx1Z2luL3YxYWxwaGExL3BsdWdpbi5wcm90bxIfcHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMSIRCg9WYWxpZGF0ZVJlcXVlc3QiEgoQVmFsaWRhdGVSZXNwb25zZSJjChNBZGRUcnVzdFpvbmVSZXF1ZXN0Ej0KCnRydXN0X3pvbmUYASABKAsyJC5wcm90by50cnVzdF96b25lLnYxYWxwaGExLlRydXN0Wm9uZUgAiAEBQg0KC190cnVzdF96b25lImQKFEFkZFRydXN0Wm9uZVJlc3BvbnNlEj0KCnRydXN0X3pvbmUYASABKAsyJC5wcm90by50cnVzdF96b25lLnYxYWxwaGExLlRydXN0Wm9uZUgAiAEBQg0KC190cnVzdF96b25lIjEKE0dldFRydXN0Wm9uZVJlcXVlc3QSEQoEbmFtZRgBIAEoCUgAiAEBQgcKBV9uYW1lImQKFEdldFRydXN0Wm9uZVJlc3BvbnNlEj0KCnRydXN0X3pvbmUYASABKAsyJC5wcm90by50cnVzdF96b25lLnYxYWxwaGExLlRydXN0Wm9uZUgAiAEBQg0KC190cnVzdF96b25lIhcKFUxpc3RUcnVzdFpvbmVzUmVxdWVzdCJTChZMaXN0VHJ1c3Rab25lc1Jlc3BvbnNlEjkKC3RydXN0X3pvbmVzGAEgAygLMiQucHJvdG8udHJ1c3Rfem9uZS52MWFscGhhMS5UcnVzdFpvbmUiZgoWVXBkYXRlVHJ1c3Rab25lUmVxdWVzdBI9Cgp0cnVzdF96b25lGAEgASgLMiQucHJvdG8udHJ1c3Rfem9uZS52MWFscGhhMS5UcnVzdFpvbmVIAIgBAUINCgtfdHJ1c3Rfem9uZSJnChdVcGRhdGVUcnVzdFpvbmVSZXNwb25zZRI9Cgp0cnVzdF96b25lGAEgASgLMiQucHJvdG8udHJ1c3Rfem9uZS52MWFscGhhMS5UcnVzdFpvbmVIAIgBAUINCgtfdHJ1c3Rfem9uZSJWChFBZGRDbHVzdGVyUmVxdWVzdBI1CgdjbHVzdGVyGAEgASgLMh8ucHJvdG8uY2x1c3Rlci52MWFscGhhMS5DbHVzdGVySACIAQFCCgoIX2NsdXN0ZXIiVwoSQWRkQ2x1c3RlclJlc3BvbnNlEjUKB2NsdXN0ZXIYASABKAsyHy5wcm90by5jbHVzdGVyLnYxYWxwaGExLkNsdXN0ZXJIAIgBAUIKCghfY2x1c3RlciJXChFHZXRDbHVzdGVyUmVxdWVzdBIRCgRuYW1lGAEgASgJSACIAQESFwoKdHJ1c3Rfem9uZRgCIAEoCUgBiAEBQgcKBV9uYW1lQg0KC190cnVzdF96b25lIlcKEkdldENsdXN0ZXJSZXNwb25zZRI1CgdjbHVzdGVyGAEgASgLMh8ucHJvdG8uY2x1c3Rlci52MWFscGhhMS5DbHVzdGVySACIAQFCCgoIX2NsdXN0ZXIiPQoTTGlzdENsdXN0ZXJzUmVxdWVzdBIXCgp0cnVzdF96b25lGAEgASgJSACIAQFCDQoLX3RydXN0X3pvbmUiSQoUTGlzdENsdXN0ZXJzUmVzcG9uc2USMQoIY2x1c3RlcnMYASADKAsyHy5wcm90by5jbHVzdGVyLnYxYWxwaGExLkNsdXN0ZXIiWQoUVXBkYXRlQ2x1c3RlclJlcXVlc3QSNQoHY2x1c3RlchgBIAEoCzIfLnByb3RvLmNsdXN0ZXIudjFhbHBoYTEuQ2x1c3RlckgAiAEBQgoKCF9jbHVzdGVyIloKFVVwZGF0ZUNsdXN0ZXJSZXNwb25zZRI1CgdjbHVzdGVyGAEgASgLMh8ucHJvdG8uY2x1c3Rlci52MWFscGhhMS5DbHVzdGVySACIAQFCCgoIX2NsdXN0ZXIicwobQWRkQXR0ZXN0YXRpb25Qb2xpY3lSZXF1ZXN0EkkKBnBvbGljeRgBIAEoCzI0LnByb3RvLmF0dGVzdGF0aW9uX3BvbGljeS52MWFscGhhMS5BdHRlc3RhdGlvblBvbGljeUgAiAEBQgkKB19wb2xpY3kidAocQWRkQXR0ZXN0YXRpb25Qb2xpY3lSZXNwb25zZRJJCgZwb2xpY3kYASABKAsyNC5wcm90by5hdHRlc3RhdGlvbl9wb2xpY3kudjFhbHBoYTEuQXR0ZXN0YXRpb25Qb2xpY3lIAIgBAUIJCgdfcG9saWN5IjkKG0dldEF0dGVzdGF0aW9uUG9saWN5UmVxdWVzdBIRCgRuYW1lGAEgASgJSACIAQFCBwoFX25hbWUiZAocR2V0QXR0ZXN0YXRpb25Qb2xpY3lSZXNwb25zZRJECgZwb2xpY3kYASABKAsyNC5wcm90by5hdHRlc3RhdGlvbl9wb2xpY3kudjFhbHBoYTEuQXR0ZXN0YXRpb25Qb2xpY3kiIAoeTGlzdEF0dGVzdGF0aW9uUG9saWNpZXNSZXF1ZXN0ImkKH0xpc3RBdHRlc3RhdGlvblBvbGljaWVzUmVzcG9uc2USRgoIcG9saWNpZXMYASADKAsyNC5wcm90by5hdHRlc3RhdGlvbl9wb2xpY3kudjFhbHBoYTEuQXR0ZXN0YXRpb25Qb2xpY3kiXQoTQWRkQVBCaW5kaW5nUmVxdWVzdBI6CgdiaW5kaW5nGAEgASgLMiQucHJvdG8uYXBfYmluZGluZy52MWFscGhhMS5BUEJpbmRpbmdIAIgBAUIKCghfYmluZGluZyJeChRBZGRBUEJpbmRpbmdSZXNwb25zZRI6CgdiaW5kaW5nGAEgASgLMiQucHJvdG8uYXBfYmluZGluZy52MWFscGhhMS5BUEJpbmRpbmdIAIgBAUIKCghfYmluZGluZyJhChdEZXN0cm95QVBCaW5kaW5nUmVxdWVzdBI6CgdiaW5kaW5nGAEgASgLMiQucHJvdG8uYXBfYmluZGluZy52MWFscGhhMS5BUEJpbmRpbmdIAIgBAUIKCghfYmluZGluZyIaChhEZXN0cm95QVBCaW5kaW5nUmVzcG9uc2Ui3AEKFUxpc3RBUEJpbmRpbmdzUmVxdWVzdBJSCgZmaWx0ZXIYASABKAsyPS5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkxpc3RBUEJpbmRpbmdzUmVxdWVzdC5GaWx0ZXJIAIgBARpkCgZGaWx0ZXISHAoPdHJ1c3Rfem9uZV9uYW1lGAEgASgJSACIAQESGAoLcG9saWN5X25hbWUYAiABKAlIAYgBAUISChBfdHJ1c3Rfem9uZV9uYW1lQg4KDF9wb2xpY3lfbmFtZUIJCgdfZmlsdGVyIlAKFkxpc3RBUEJpbmRpbmdzUmVzcG9uc2USNgoIYmluZGluZ3MYASADKAsyJC5wcm90by5hcF9iaW5kaW5nLnYxYWxwaGExLkFQQmluZGluZyJlChRBZGRGZWRlcmF0aW9uUmVxdWVzdBI+CgpmZWRlcmF0aW9uGAEgASgLMiUucHJvdG8uZmVkZXJhdGlvbi52MWFscGhhMS5GZWRlcmF0aW9uSACIAQFCDQoLX2ZlZGVyYXRpb24iZgoVQWRkRmVkZXJhdGlvblJlc3BvbnNlEj4KCmZlZGVyYXRpb24YASABKAsyJS5wcm90by5mZWRlcmF0aW9uLnYxYWxwaGExLkZlZGVyYXRpb25IAIgBAUINCgtfZmVkZXJhdGlvbiIYChZMaXN0RmVkZXJhdGlvbnNSZXF1ZXN0IlUKF0xpc3RGZWRlcmF0aW9uc1Jlc3BvbnNlEjoKC2ZlZGVyYXRpb25zGAEgAygLMiUucHJvdG8uZmVkZXJhdGlvbi52MWFscGhhMS5GZWRlcmF0aW9uIlUKIUxpc3RGZWRlcmF0aW9uc0J5VHJ1c3Rab25lUmVxdWVzdBIcCg90cnVzdF96b25lX25hbWUYASABKAlIAIgBAUISChBfdHJ1c3Rfem9uZV9uYW1lImAKIkxpc3RGZWRlcmF0aW9uc0J5VHJ1c3Rab25lUmVzcG9uc2USOgoLZmVkZXJhdGlvbnMYASADKAsyJS5wcm90by5mZWRlcmF0aW9uLnYxYWxwaGExLkZlZGVyYXRpb24y/xIKF0RhdGFTb3VyY2VQbHVnaW5TZXJ2aWNlEm8KCFZhbGlkYXRlEjAucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5WYWxpZGF0ZVJlcXVlc3QaMS5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLlZhbGlkYXRlUmVzcG9uc2USewoMQWRkVHJ1c3Rab25lEjQucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5BZGRUcnVzdFpvbmVSZXF1ZXN0GjUucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5BZGRUcnVzdFpvbmVSZXNwb25zZRJ7CgxHZXRUcnVzdFpvbmUSNC5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkdldFRydXN0Wm9uZVJlcXVlc3QaNS5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkdldFRydXN0Wm9uZVJlc3BvbnNlEoEBCg5MaXN0VHJ1c3Rab25lcxI2LnByb3RvLmNvZmlkZWN0bF9wbHVnaW4udjFhbHBoYTEuTGlzdFRydXN0Wm9uZXNSZXF1ZXN0GjcucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5MaXN0VHJ1c3Rab25lc1Jlc3BvbnNlEoQBCg9VcGRhdGVUcnVzdFpvbmUSNy5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLlVwZGF0ZVRydXN0Wm9uZVJlcXVlc3QaOC5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLlVwZGF0ZVRydXN0Wm9uZVJlc3BvbnNlEnUKCkFkZENsdXN0ZXISMi5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkFkZENsdXN0ZXJSZXF1ZXN0GjMucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5BZGRDbHVzdGVyUmVzcG9uc2USdQoKR2V0Q2x1c3RlchIyLnByb3RvLmNvZmlkZWN0bF9wbHVnaW4udjFhbHBoYTEuR2V0Q2x1c3RlclJlcXVlc3QaMy5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkdldENsdXN0ZXJSZXNwb25zZRJ7CgxMaXN0Q2x1c3RlcnMSNC5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkxpc3RDbHVzdGVyc1JlcXVlc3QaNS5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkxpc3RDbHVzdGVyc1Jlc3BvbnNlEn4KDVVwZGF0ZUNsdXN0ZXISNS5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLlVwZGF0ZUNsdXN0ZXJSZXF1ZXN0GjYucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5VcGRhdGVDbHVzdGVyUmVzcG9uc2USkwEKFEFkZEF0dGVzdGF0aW9uUG9saWN5EjwucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5BZGRBdHRlc3RhdGlvblBvbGljeVJlcXVlc3QaPS5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkFkZEF0dGVzdGF0aW9uUG9saWN5UmVzcG9uc2USkwEKFEdldEF0dGVzdGF0aW9uUG9saWN5EjwucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5HZXRBdHRlc3RhdGlvblBvbGljeVJlcXVlc3QaPS5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkdldEF0dGVzdGF0aW9uUG9saWN5UmVzcG9uc2USnAEKF0xpc3RBdHRlc3RhdGlvblBvbGljaWVzEj8ucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5MaXN0QXR0ZXN0YXRpb25Qb2xpY2llc1JlcXVlc3QaQC5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkxpc3RBdHRlc3RhdGlvblBvbGljaWVzUmVzcG9uc2USewoMQWRkQVBCaW5kaW5nEjQucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5BZGRBUEJpbmRpbmdSZXF1ZXN0GjUucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5BZGRBUEJpbmRpbmdSZXNwb25zZRKHAQoQRGVzdHJveUFQQmluZGluZxI4LnByb3RvLmNvZmlkZWN0bF9wbHVnaW4udjFhbHBoYTEuRGVzdHJveUFQQmluZGluZ1JlcXVlc3QaOS5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkRlc3Ryb3lBUEJpbmRpbmdSZXNwb25zZRKBAQoOTGlzdEFQQmluZGluZ3MSNi5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkxpc3RBUEJpbmRpbmdzUmVxdWVzdBo3LnByb3RvLmNvZmlkZWN0bF9wbHVnaW4udjFhbHBoYTEuTGlzdEFQQmluZGluZ3NSZXNwb25zZRJ+Cg1BZGRGZWRlcmF0aW9uEjUucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5BZGRGZWRlcmF0aW9uUmVxdWVzdBo2LnByb3RvLmNvZmlkZWN0bF9wbHVnaW4udjFhbHBoYTEuQWRkRmVkZXJhdGlvblJlc3BvbnNlEoQBCg9MaXN0RmVkZXJhdGlvbnMSNy5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkxpc3RGZWRlcmF0aW9uc1JlcXVlc3QaOC5wcm90by5jb2ZpZGVjdGxfcGx1Z2luLnYxYWxwaGExLkxpc3RGZWRlcmF0aW9uc1Jlc3BvbnNlEqUBChpMaXN0RmVkZXJhdGlvbnNCeVRydXN0Wm9uZRJCLnByb3RvLmNvZmlkZWN0bF9wbHVnaW4udjFhbHBoYTEuTGlzdEZlZGVyYXRpb25zQnlUcnVzdFpvbmVSZXF1ZXN0GkMucHJvdG8uY29maWRlY3RsX3BsdWdpbi52MWFscGhhMS5MaXN0RmVkZXJhdGlvbnNCeVRydXN0Wm9uZVJlc3BvbnNlQklaR2dpdGh1Yi5jb20vY29maWRlL2NvZmlkZS1hcGktc2RrL2dlbi9nby9wcm90by9jb2ZpZGVjdGxfcGx1Z2luL3YxYWxwaGExYgZwcm90bzM", [file_proto_ap_binding_v1alpha1_ap_binding, file_proto_attestation_policy_v1alpha1_attestation_policy, file_proto_cluster_v1alpha1_cluster, file_proto_federation_v1alpha1_federation, file_proto_trust_zone_v1alpha1_trust_zone]);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ValidateRequest
 */
export type ValidateRequest = Message<"proto.cofidectl_plugin.v1alpha1.ValidateRequest"> & {
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ValidateRequest.
 * Use `create(ValidateRequestSchema)` to create a new message.
 */
export const ValidateRequestSchema: GenMessage<ValidateRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 0);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ValidateResponse
 */
export type ValidateResponse = Message<"proto.cofidectl_plugin.v1alpha1.ValidateResponse"> & {
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ValidateResponse.
 * Use `create(ValidateResponseSchema)` to create a new message.
 */
export const ValidateResponseSchema: GenMessage<ValidateResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 1);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.AddTrustZoneRequest
 */
export type AddTrustZoneRequest = Message<"proto.cofidectl_plugin.v1alpha1.AddTrustZoneRequest"> & {
  /**
   * @generated from field: optional proto.trust_zone.v1alpha1.TrustZone trust_zone = 1;
   */
  trustZone?: TrustZone;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.AddTrustZoneRequest.
 * Use `create(AddTrustZoneRequestSchema)` to create a new message.
 */
export const AddTrustZoneRequestSchema: GenMessage<AddTrustZoneRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 2);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.AddTrustZoneResponse
 */
export type AddTrustZoneResponse = Message<"proto.cofidectl_plugin.v1alpha1.AddTrustZoneResponse"> & {
  /**
   * @generated from field: optional proto.trust_zone.v1alpha1.TrustZone trust_zone = 1;
   */
  trustZone?: TrustZone;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.AddTrustZoneResponse.
 * Use `create(AddTrustZoneResponseSchema)` to create a new message.
 */
export const AddTrustZoneResponseSchema: GenMessage<AddTrustZoneResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 3);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.GetTrustZoneRequest
 */
export type GetTrustZoneRequest = Message<"proto.cofidectl_plugin.v1alpha1.GetTrustZoneRequest"> & {
  /**
   * @generated from field: optional string name = 1;
   */
  name?: string;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.GetTrustZoneRequest.
 * Use `create(GetTrustZoneRequestSchema)` to create a new message.
 */
export const GetTrustZoneRequestSchema: GenMessage<GetTrustZoneRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 4);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.GetTrustZoneResponse
 */
export type GetTrustZoneResponse = Message<"proto.cofidectl_plugin.v1alpha1.GetTrustZoneResponse"> & {
  /**
   * @generated from field: optional proto.trust_zone.v1alpha1.TrustZone trust_zone = 1;
   */
  trustZone?: TrustZone;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.GetTrustZoneResponse.
 * Use `create(GetTrustZoneResponseSchema)` to create a new message.
 */
export const GetTrustZoneResponseSchema: GenMessage<GetTrustZoneResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 5);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListTrustZonesRequest
 */
export type ListTrustZonesRequest = Message<"proto.cofidectl_plugin.v1alpha1.ListTrustZonesRequest"> & {
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListTrustZonesRequest.
 * Use `create(ListTrustZonesRequestSchema)` to create a new message.
 */
export const ListTrustZonesRequestSchema: GenMessage<ListTrustZonesRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 6);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListTrustZonesResponse
 */
export type ListTrustZonesResponse = Message<"proto.cofidectl_plugin.v1alpha1.ListTrustZonesResponse"> & {
  /**
   * @generated from field: repeated proto.trust_zone.v1alpha1.TrustZone trust_zones = 1;
   */
  trustZones: TrustZone[];
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListTrustZonesResponse.
 * Use `create(ListTrustZonesResponseSchema)` to create a new message.
 */
export const ListTrustZonesResponseSchema: GenMessage<ListTrustZonesResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 7);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.UpdateTrustZoneRequest
 */
export type UpdateTrustZoneRequest = Message<"proto.cofidectl_plugin.v1alpha1.UpdateTrustZoneRequest"> & {
  /**
   * @generated from field: optional proto.trust_zone.v1alpha1.TrustZone trust_zone = 1;
   */
  trustZone?: TrustZone;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.UpdateTrustZoneRequest.
 * Use `create(UpdateTrustZoneRequestSchema)` to create a new message.
 */
export const UpdateTrustZoneRequestSchema: GenMessage<UpdateTrustZoneRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 8);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.UpdateTrustZoneResponse
 */
export type UpdateTrustZoneResponse = Message<"proto.cofidectl_plugin.v1alpha1.UpdateTrustZoneResponse"> & {
  /**
   * @generated from field: optional proto.trust_zone.v1alpha1.TrustZone trust_zone = 1;
   */
  trustZone?: TrustZone;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.UpdateTrustZoneResponse.
 * Use `create(UpdateTrustZoneResponseSchema)` to create a new message.
 */
export const UpdateTrustZoneResponseSchema: GenMessage<UpdateTrustZoneResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 9);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.AddClusterRequest
 */
export type AddClusterRequest = Message<"proto.cofidectl_plugin.v1alpha1.AddClusterRequest"> & {
  /**
   * @generated from field: optional proto.cluster.v1alpha1.Cluster cluster = 1;
   */
  cluster?: Cluster;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.AddClusterRequest.
 * Use `create(AddClusterRequestSchema)` to create a new message.
 */
export const AddClusterRequestSchema: GenMessage<AddClusterRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 10);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.AddClusterResponse
 */
export type AddClusterResponse = Message<"proto.cofidectl_plugin.v1alpha1.AddClusterResponse"> & {
  /**
   * @generated from field: optional proto.cluster.v1alpha1.Cluster cluster = 1;
   */
  cluster?: Cluster;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.AddClusterResponse.
 * Use `create(AddClusterResponseSchema)` to create a new message.
 */
export const AddClusterResponseSchema: GenMessage<AddClusterResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 11);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.GetClusterRequest
 */
export type GetClusterRequest = Message<"proto.cofidectl_plugin.v1alpha1.GetClusterRequest"> & {
  /**
   * @generated from field: optional string name = 1;
   */
  name?: string;

  /**
   * @generated from field: optional string trust_zone = 2;
   */
  trustZone?: string;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.GetClusterRequest.
 * Use `create(GetClusterRequestSchema)` to create a new message.
 */
export const GetClusterRequestSchema: GenMessage<GetClusterRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 12);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.GetClusterResponse
 */
export type GetClusterResponse = Message<"proto.cofidectl_plugin.v1alpha1.GetClusterResponse"> & {
  /**
   * @generated from field: optional proto.cluster.v1alpha1.Cluster cluster = 1;
   */
  cluster?: Cluster;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.GetClusterResponse.
 * Use `create(GetClusterResponseSchema)` to create a new message.
 */
export const GetClusterResponseSchema: GenMessage<GetClusterResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 13);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListClustersRequest
 */
export type ListClustersRequest = Message<"proto.cofidectl_plugin.v1alpha1.ListClustersRequest"> & {
  /**
   * @generated from field: optional string trust_zone = 1;
   */
  trustZone?: string;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListClustersRequest.
 * Use `create(ListClustersRequestSchema)` to create a new message.
 */
export const ListClustersRequestSchema: GenMessage<ListClustersRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 14);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListClustersResponse
 */
export type ListClustersResponse = Message<"proto.cofidectl_plugin.v1alpha1.ListClustersResponse"> & {
  /**
   * @generated from field: repeated proto.cluster.v1alpha1.Cluster clusters = 1;
   */
  clusters: Cluster[];
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListClustersResponse.
 * Use `create(ListClustersResponseSchema)` to create a new message.
 */
export const ListClustersResponseSchema: GenMessage<ListClustersResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 15);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.UpdateClusterRequest
 */
export type UpdateClusterRequest = Message<"proto.cofidectl_plugin.v1alpha1.UpdateClusterRequest"> & {
  /**
   * @generated from field: optional proto.cluster.v1alpha1.Cluster cluster = 1;
   */
  cluster?: Cluster;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.UpdateClusterRequest.
 * Use `create(UpdateClusterRequestSchema)` to create a new message.
 */
export const UpdateClusterRequestSchema: GenMessage<UpdateClusterRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 16);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.UpdateClusterResponse
 */
export type UpdateClusterResponse = Message<"proto.cofidectl_plugin.v1alpha1.UpdateClusterResponse"> & {
  /**
   * @generated from field: optional proto.cluster.v1alpha1.Cluster cluster = 1;
   */
  cluster?: Cluster;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.UpdateClusterResponse.
 * Use `create(UpdateClusterResponseSchema)` to create a new message.
 */
export const UpdateClusterResponseSchema: GenMessage<UpdateClusterResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 17);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.AddAttestationPolicyRequest
 */
export type AddAttestationPolicyRequest = Message<"proto.cofidectl_plugin.v1alpha1.AddAttestationPolicyRequest"> & {
  /**
   * @generated from field: optional proto.attestation_policy.v1alpha1.AttestationPolicy policy = 1;
   */
  policy?: AttestationPolicy;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.AddAttestationPolicyRequest.
 * Use `create(AddAttestationPolicyRequestSchema)` to create a new message.
 */
export const AddAttestationPolicyRequestSchema: GenMessage<AddAttestationPolicyRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 18);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.AddAttestationPolicyResponse
 */
export type AddAttestationPolicyResponse = Message<"proto.cofidectl_plugin.v1alpha1.AddAttestationPolicyResponse"> & {
  /**
   * @generated from field: optional proto.attestation_policy.v1alpha1.AttestationPolicy policy = 1;
   */
  policy?: AttestationPolicy;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.AddAttestationPolicyResponse.
 * Use `create(AddAttestationPolicyResponseSchema)` to create a new message.
 */
export const AddAttestationPolicyResponseSchema: GenMessage<AddAttestationPolicyResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 19);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.GetAttestationPolicyRequest
 */
export type GetAttestationPolicyRequest = Message<"proto.cofidectl_plugin.v1alpha1.GetAttestationPolicyRequest"> & {
  /**
   * @generated from field: optional string name = 1;
   */
  name?: string;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.GetAttestationPolicyRequest.
 * Use `create(GetAttestationPolicyRequestSchema)` to create a new message.
 */
export const GetAttestationPolicyRequestSchema: GenMessage<GetAttestationPolicyRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 20);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.GetAttestationPolicyResponse
 */
export type GetAttestationPolicyResponse = Message<"proto.cofidectl_plugin.v1alpha1.GetAttestationPolicyResponse"> & {
  /**
   * @generated from field: proto.attestation_policy.v1alpha1.AttestationPolicy policy = 1;
   */
  policy?: AttestationPolicy;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.GetAttestationPolicyResponse.
 * Use `create(GetAttestationPolicyResponseSchema)` to create a new message.
 */
export const GetAttestationPolicyResponseSchema: GenMessage<GetAttestationPolicyResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 21);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListAttestationPoliciesRequest
 */
export type ListAttestationPoliciesRequest = Message<"proto.cofidectl_plugin.v1alpha1.ListAttestationPoliciesRequest"> & {
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListAttestationPoliciesRequest.
 * Use `create(ListAttestationPoliciesRequestSchema)` to create a new message.
 */
export const ListAttestationPoliciesRequestSchema: GenMessage<ListAttestationPoliciesRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 22);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListAttestationPoliciesResponse
 */
export type ListAttestationPoliciesResponse = Message<"proto.cofidectl_plugin.v1alpha1.ListAttestationPoliciesResponse"> & {
  /**
   * @generated from field: repeated proto.attestation_policy.v1alpha1.AttestationPolicy policies = 1;
   */
  policies: AttestationPolicy[];
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListAttestationPoliciesResponse.
 * Use `create(ListAttestationPoliciesResponseSchema)` to create a new message.
 */
export const ListAttestationPoliciesResponseSchema: GenMessage<ListAttestationPoliciesResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 23);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.AddAPBindingRequest
 */
export type AddAPBindingRequest = Message<"proto.cofidectl_plugin.v1alpha1.AddAPBindingRequest"> & {
  /**
   * @generated from field: optional proto.ap_binding.v1alpha1.APBinding binding = 1;
   */
  binding?: APBinding;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.AddAPBindingRequest.
 * Use `create(AddAPBindingRequestSchema)` to create a new message.
 */
export const AddAPBindingRequestSchema: GenMessage<AddAPBindingRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 24);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.AddAPBindingResponse
 */
export type AddAPBindingResponse = Message<"proto.cofidectl_plugin.v1alpha1.AddAPBindingResponse"> & {
  /**
   * @generated from field: optional proto.ap_binding.v1alpha1.APBinding binding = 1;
   */
  binding?: APBinding;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.AddAPBindingResponse.
 * Use `create(AddAPBindingResponseSchema)` to create a new message.
 */
export const AddAPBindingResponseSchema: GenMessage<AddAPBindingResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 25);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.DestroyAPBindingRequest
 */
export type DestroyAPBindingRequest = Message<"proto.cofidectl_plugin.v1alpha1.DestroyAPBindingRequest"> & {
  /**
   * @generated from field: optional proto.ap_binding.v1alpha1.APBinding binding = 1;
   */
  binding?: APBinding;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.DestroyAPBindingRequest.
 * Use `create(DestroyAPBindingRequestSchema)` to create a new message.
 */
export const DestroyAPBindingRequestSchema: GenMessage<DestroyAPBindingRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 26);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.DestroyAPBindingResponse
 */
export type DestroyAPBindingResponse = Message<"proto.cofidectl_plugin.v1alpha1.DestroyAPBindingResponse"> & {
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.DestroyAPBindingResponse.
 * Use `create(DestroyAPBindingResponseSchema)` to create a new message.
 */
export const DestroyAPBindingResponseSchema: GenMessage<DestroyAPBindingResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 27);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListAPBindingsRequest
 */
export type ListAPBindingsRequest = Message<"proto.cofidectl_plugin.v1alpha1.ListAPBindingsRequest"> & {
  /**
   * @generated from field: optional proto.cofidectl_plugin.v1alpha1.ListAPBindingsRequest.Filter filter = 1;
   */
  filter?: ListAPBindingsRequest_Filter;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListAPBindingsRequest.
 * Use `create(ListAPBindingsRequestSchema)` to create a new message.
 */
export const ListAPBindingsRequestSchema: GenMessage<ListAPBindingsRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 28);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListAPBindingsRequest.Filter
 */
export type ListAPBindingsRequest_Filter = Message<"proto.cofidectl_plugin.v1alpha1.ListAPBindingsRequest.Filter"> & {
  /**
   * @generated from field: optional string trust_zone_name = 1;
   */
  trustZoneName?: string;

  /**
   * @generated from field: optional string policy_name = 2;
   */
  policyName?: string;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListAPBindingsRequest.Filter.
 * Use `create(ListAPBindingsRequest_FilterSchema)` to create a new message.
 */
export const ListAPBindingsRequest_FilterSchema: GenMessage<ListAPBindingsRequest_Filter> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 28, 0);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListAPBindingsResponse
 */
export type ListAPBindingsResponse = Message<"proto.cofidectl_plugin.v1alpha1.ListAPBindingsResponse"> & {
  /**
   * @generated from field: repeated proto.ap_binding.v1alpha1.APBinding bindings = 1;
   */
  bindings: APBinding[];
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListAPBindingsResponse.
 * Use `create(ListAPBindingsResponseSchema)` to create a new message.
 */
export const ListAPBindingsResponseSchema: GenMessage<ListAPBindingsResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 29);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.AddFederationRequest
 */
export type AddFederationRequest = Message<"proto.cofidectl_plugin.v1alpha1.AddFederationRequest"> & {
  /**
   * @generated from field: optional proto.federation.v1alpha1.Federation federation = 1;
   */
  federation?: Federation;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.AddFederationRequest.
 * Use `create(AddFederationRequestSchema)` to create a new message.
 */
export const AddFederationRequestSchema: GenMessage<AddFederationRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 30);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.AddFederationResponse
 */
export type AddFederationResponse = Message<"proto.cofidectl_plugin.v1alpha1.AddFederationResponse"> & {
  /**
   * @generated from field: optional proto.federation.v1alpha1.Federation federation = 1;
   */
  federation?: Federation;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.AddFederationResponse.
 * Use `create(AddFederationResponseSchema)` to create a new message.
 */
export const AddFederationResponseSchema: GenMessage<AddFederationResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 31);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListFederationsRequest
 */
export type ListFederationsRequest = Message<"proto.cofidectl_plugin.v1alpha1.ListFederationsRequest"> & {
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListFederationsRequest.
 * Use `create(ListFederationsRequestSchema)` to create a new message.
 */
export const ListFederationsRequestSchema: GenMessage<ListFederationsRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 32);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListFederationsResponse
 */
export type ListFederationsResponse = Message<"proto.cofidectl_plugin.v1alpha1.ListFederationsResponse"> & {
  /**
   * @generated from field: repeated proto.federation.v1alpha1.Federation federations = 1;
   */
  federations: Federation[];
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListFederationsResponse.
 * Use `create(ListFederationsResponseSchema)` to create a new message.
 */
export const ListFederationsResponseSchema: GenMessage<ListFederationsResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 33);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListFederationsByTrustZoneRequest
 */
export type ListFederationsByTrustZoneRequest = Message<"proto.cofidectl_plugin.v1alpha1.ListFederationsByTrustZoneRequest"> & {
  /**
   * @generated from field: optional string trust_zone_name = 1;
   */
  trustZoneName?: string;
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListFederationsByTrustZoneRequest.
 * Use `create(ListFederationsByTrustZoneRequestSchema)` to create a new message.
 */
export const ListFederationsByTrustZoneRequestSchema: GenMessage<ListFederationsByTrustZoneRequest> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 34);

/**
 * @generated from message proto.cofidectl_plugin.v1alpha1.ListFederationsByTrustZoneResponse
 */
export type ListFederationsByTrustZoneResponse = Message<"proto.cofidectl_plugin.v1alpha1.ListFederationsByTrustZoneResponse"> & {
  /**
   * @generated from field: repeated proto.federation.v1alpha1.Federation federations = 1;
   */
  federations: Federation[];
};

/**
 * Describes the message proto.cofidectl_plugin.v1alpha1.ListFederationsByTrustZoneResponse.
 * Use `create(ListFederationsByTrustZoneResponseSchema)` to create a new message.
 */
export const ListFederationsByTrustZoneResponseSchema: GenMessage<ListFederationsByTrustZoneResponse> = /*@__PURE__*/
  messageDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 35);

/**
 * @generated from service proto.cofidectl_plugin.v1alpha1.DataSourcePluginService
 */
export const DataSourcePluginService: GenService<{
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.Validate
   */
  validate: {
    methodKind: "unary";
    input: typeof ValidateRequestSchema;
    output: typeof ValidateResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.AddTrustZone
   */
  addTrustZone: {
    methodKind: "unary";
    input: typeof AddTrustZoneRequestSchema;
    output: typeof AddTrustZoneResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.GetTrustZone
   */
  getTrustZone: {
    methodKind: "unary";
    input: typeof GetTrustZoneRequestSchema;
    output: typeof GetTrustZoneResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.ListTrustZones
   */
  listTrustZones: {
    methodKind: "unary";
    input: typeof ListTrustZonesRequestSchema;
    output: typeof ListTrustZonesResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.UpdateTrustZone
   */
  updateTrustZone: {
    methodKind: "unary";
    input: typeof UpdateTrustZoneRequestSchema;
    output: typeof UpdateTrustZoneResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.AddCluster
   */
  addCluster: {
    methodKind: "unary";
    input: typeof AddClusterRequestSchema;
    output: typeof AddClusterResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.GetCluster
   */
  getCluster: {
    methodKind: "unary";
    input: typeof GetClusterRequestSchema;
    output: typeof GetClusterResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.ListClusters
   */
  listClusters: {
    methodKind: "unary";
    input: typeof ListClustersRequestSchema;
    output: typeof ListClustersResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.UpdateCluster
   */
  updateCluster: {
    methodKind: "unary";
    input: typeof UpdateClusterRequestSchema;
    output: typeof UpdateClusterResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.AddAttestationPolicy
   */
  addAttestationPolicy: {
    methodKind: "unary";
    input: typeof AddAttestationPolicyRequestSchema;
    output: typeof AddAttestationPolicyResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.GetAttestationPolicy
   */
  getAttestationPolicy: {
    methodKind: "unary";
    input: typeof GetAttestationPolicyRequestSchema;
    output: typeof GetAttestationPolicyResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.ListAttestationPolicies
   */
  listAttestationPolicies: {
    methodKind: "unary";
    input: typeof ListAttestationPoliciesRequestSchema;
    output: typeof ListAttestationPoliciesResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.AddAPBinding
   */
  addAPBinding: {
    methodKind: "unary";
    input: typeof AddAPBindingRequestSchema;
    output: typeof AddAPBindingResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.DestroyAPBinding
   */
  destroyAPBinding: {
    methodKind: "unary";
    input: typeof DestroyAPBindingRequestSchema;
    output: typeof DestroyAPBindingResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.ListAPBindings
   */
  listAPBindings: {
    methodKind: "unary";
    input: typeof ListAPBindingsRequestSchema;
    output: typeof ListAPBindingsResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.AddFederation
   */
  addFederation: {
    methodKind: "unary";
    input: typeof AddFederationRequestSchema;
    output: typeof AddFederationResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.ListFederations
   */
  listFederations: {
    methodKind: "unary";
    input: typeof ListFederationsRequestSchema;
    output: typeof ListFederationsResponseSchema;
  },
  /**
   * @generated from rpc proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.ListFederationsByTrustZone
   */
  listFederationsByTrustZone: {
    methodKind: "unary";
    input: typeof ListFederationsByTrustZoneRequestSchema;
    output: typeof ListFederationsByTrustZoneResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_proto_cofidectl_plugin_v1alpha1_plugin, 0);

