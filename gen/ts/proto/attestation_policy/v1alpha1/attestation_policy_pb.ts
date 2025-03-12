// Copyright 2024 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file proto/attestation_policy/v1alpha1/attestation_policy.proto (package proto.attestation_policy.v1alpha1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Selector } from "../../spire/api/types/selector_pb";
import { file_proto_spire_api_types_selector } from "../../spire/api/types/selector_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file proto/attestation_policy/v1alpha1/attestation_policy.proto.
 */
export const file_proto_attestation_policy_v1alpha1_attestation_policy: GenFile = /*@__PURE__*/
  fileDesc("Cjpwcm90by9hdHRlc3RhdGlvbl9wb2xpY3kvdjFhbHBoYTEvYXR0ZXN0YXRpb25fcG9saWN5LnByb3RvEiFwcm90by5hdHRlc3RhdGlvbl9wb2xpY3kudjFhbHBoYTEi6QEKEUF0dGVzdGF0aW9uUG9saWN5Eg8KAmlkGAQgASgJSAGIAQESDAoEbmFtZRgBIAEoCRITCgZvcmdfaWQYBSABKAlIAogBARJFCgprdWJlcm5ldGVzGAIgASgLMi8ucHJvdG8uYXR0ZXN0YXRpb25fcG9saWN5LnYxYWxwaGExLkFQS3ViZXJuZXRlc0gAEj0KBnN0YXRpYxgDIAEoCzIrLnByb3RvLmF0dGVzdGF0aW9uX3BvbGljeS52MWFscGhhMS5BUFN0YXRpY0gAQggKBnBvbGljeUIFCgNfaWRCCQoHX29yZ19pZCLaAQoMQVBLdWJlcm5ldGVzElMKEm5hbWVzcGFjZV9zZWxlY3RvchgBIAEoCzIyLnByb3RvLmF0dGVzdGF0aW9uX3BvbGljeS52MWFscGhhMS5BUExhYmVsU2VsZWN0b3JIAIgBARJNCgxwb2Rfc2VsZWN0b3IYAiABKAsyMi5wcm90by5hdHRlc3RhdGlvbl9wb2xpY3kudjFhbHBoYTEuQVBMYWJlbFNlbGVjdG9ySAGIAQFCFQoTX25hbWVzcGFjZV9zZWxlY3RvckIPCg1fcG9kX3NlbGVjdG9yIvEBCg9BUExhYmVsU2VsZWN0b3ISWQoMbWF0Y2hfbGFiZWxzGAEgAygLMkMucHJvdG8uYXR0ZXN0YXRpb25fcG9saWN5LnYxYWxwaGExLkFQTGFiZWxTZWxlY3Rvci5NYXRjaExhYmVsc0VudHJ5Ek8KEW1hdGNoX2V4cHJlc3Npb25zGAIgAygLMjQucHJvdG8uYXR0ZXN0YXRpb25fcG9saWN5LnYxYWxwaGExLkFQTWF0Y2hFeHByZXNzaW9uGjIKEE1hdGNoTGFiZWxzRW50cnkSCwoDa2V5GAEgASgJEg0KBXZhbHVlGAIgASgJOgI4ASJCChFBUE1hdGNoRXhwcmVzc2lvbhILCgNrZXkYASABKAkSEAoIb3BlcmF0b3IYAiABKAkSDgoGdmFsdWVzGAMgAygJIl4KCEFQU3RhdGljEhYKCXNwaWZmZV9pZBgBIAEoCUgAiAEBEiwKCXNlbGVjdG9ycxgCIAMoCzIZLnNwaXJlLmFwaS50eXBlcy5TZWxlY3RvckIMCgpfc3BpZmZlX2lkQktaSWdpdGh1Yi5jb20vY29maWRlL2NvZmlkZS1hcGktc2RrL2dlbi9nby9wcm90by9hdHRlc3RhdGlvbl9wb2xpY3kvdjFhbHBoYTFiBnByb3RvMw", [file_proto_spire_api_types_selector]);

/**
 * @generated from message proto.attestation_policy.v1alpha1.AttestationPolicy
 */
export type AttestationPolicy = Message<"proto.attestation_policy.v1alpha1.AttestationPolicy"> & {
  /**
   * @generated from field: optional string id = 4;
   */
  id?: string;

  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: optional string org_id = 5;
   */
  orgId?: string;

  /**
   * @generated from oneof proto.attestation_policy.v1alpha1.AttestationPolicy.policy
   */
  policy: {
    /**
     * @generated from field: proto.attestation_policy.v1alpha1.APKubernetes kubernetes = 2;
     */
    value: APKubernetes;
    case: "kubernetes";
  } | {
    /**
     * @generated from field: proto.attestation_policy.v1alpha1.APStatic static = 3;
     */
    value: APStatic;
    case: "static";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message proto.attestation_policy.v1alpha1.AttestationPolicy.
 * Use `create(AttestationPolicySchema)` to create a new message.
 */
export const AttestationPolicySchema: GenMessage<AttestationPolicy> = /*@__PURE__*/
  messageDesc(file_proto_attestation_policy_v1alpha1_attestation_policy, 0);

/**
 * @generated from message proto.attestation_policy.v1alpha1.APKubernetes
 */
export type APKubernetes = Message<"proto.attestation_policy.v1alpha1.APKubernetes"> & {
  /**
   * @generated from field: optional proto.attestation_policy.v1alpha1.APLabelSelector namespace_selector = 1;
   */
  namespaceSelector?: APLabelSelector;

  /**
   * @generated from field: optional proto.attestation_policy.v1alpha1.APLabelSelector pod_selector = 2;
   */
  podSelector?: APLabelSelector;
};

/**
 * Describes the message proto.attestation_policy.v1alpha1.APKubernetes.
 * Use `create(APKubernetesSchema)` to create a new message.
 */
export const APKubernetesSchema: GenMessage<APKubernetes> = /*@__PURE__*/
  messageDesc(file_proto_attestation_policy_v1alpha1_attestation_policy, 1);

/**
 * This definition has been adapted from the LabelSelector message in Kubernetes.
 * https://github.com/kubernetes/apimachinery/blob/master/pkg/apis/meta/v1/generated.proto
 *
 * @generated from message proto.attestation_policy.v1alpha1.APLabelSelector
 */
export type APLabelSelector = Message<"proto.attestation_policy.v1alpha1.APLabelSelector"> & {
  /**
   * @generated from field: map<string, string> match_labels = 1;
   */
  matchLabels: { [key: string]: string };

  /**
   * @generated from field: repeated proto.attestation_policy.v1alpha1.APMatchExpression match_expressions = 2;
   */
  matchExpressions: APMatchExpression[];
};

/**
 * Describes the message proto.attestation_policy.v1alpha1.APLabelSelector.
 * Use `create(APLabelSelectorSchema)` to create a new message.
 */
export const APLabelSelectorSchema: GenMessage<APLabelSelector> = /*@__PURE__*/
  messageDesc(file_proto_attestation_policy_v1alpha1_attestation_policy, 2);

/**
 * @generated from message proto.attestation_policy.v1alpha1.APMatchExpression
 */
export type APMatchExpression = Message<"proto.attestation_policy.v1alpha1.APMatchExpression"> & {
  /**
   * @generated from field: string key = 1;
   */
  key: string;

  /**
   * @generated from field: string operator = 2;
   */
  operator: string;

  /**
   * @generated from field: repeated string values = 3;
   */
  values: string[];
};

/**
 * Describes the message proto.attestation_policy.v1alpha1.APMatchExpression.
 * Use `create(APMatchExpressionSchema)` to create a new message.
 */
export const APMatchExpressionSchema: GenMessage<APMatchExpression> = /*@__PURE__*/
  messageDesc(file_proto_attestation_policy_v1alpha1_attestation_policy, 3);

/**
 * APStatic represents a static attestation policy
 *
 * @generated from message proto.attestation_policy.v1alpha1.APStatic
 */
export type APStatic = Message<"proto.attestation_policy.v1alpha1.APStatic"> & {
  /**
   * @generated from field: optional string spiffe_id = 1;
   */
  spiffeId?: string;

  /**
   * @generated from field: repeated spire.api.types.Selector selectors = 2;
   */
  selectors: Selector[];
};

/**
 * Describes the message proto.attestation_policy.v1alpha1.APStatic.
 * Use `create(APStaticSchema)` to create a new message.
 */
export const APStaticSchema: GenMessage<APStatic> = /*@__PURE__*/
  messageDesc(file_proto_attestation_policy_v1alpha1_attestation_policy, 4);

