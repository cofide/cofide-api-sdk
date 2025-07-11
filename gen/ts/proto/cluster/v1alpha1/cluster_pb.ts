// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file proto/cluster/v1alpha1/cluster.proto (package proto.cluster.v1alpha1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_struct } from "@bufbuild/protobuf/wkt";
import type { TrustProvider } from "../../trust_provider/v1alpha1/trust_provider_pb";
import { file_proto_trust_provider_v1alpha1_trust_provider } from "../../trust_provider/v1alpha1/trust_provider_pb";
import type { JsonObject, Message } from "@bufbuild/protobuf";

/**
 * Describes the file proto/cluster/v1alpha1/cluster.proto.
 */
export const file_proto_cluster_v1alpha1_cluster: GenFile = /*@__PURE__*/
  fileDesc("CiRwcm90by9jbHVzdGVyL3YxYWxwaGExL2NsdXN0ZXIucHJvdG8SFnByb3RvLmNsdXN0ZXIudjFhbHBoYTEingQKB0NsdXN0ZXISDwoCaWQYCCABKAlIAIgBARIRCgRuYW1lGAEgASgJSAGIAQESEwoGb3JnX2lkGAkgASgJSAKIAQESFwoKdHJ1c3Rfem9uZRgCIAEoCUgDiAEBEhoKDXRydXN0X3pvbmVfaWQYCiABKAlIBIgBARIfChJrdWJlcm5ldGVzX2NvbnRleHQYAyABKAlIBYgBARJJCg50cnVzdF9wcm92aWRlchgEIAEoCzIsLnByb3RvLnRydXN0X3Byb3ZpZGVyLnYxYWxwaGExLlRydXN0UHJvdmlkZXJIBogBARI3ChFleHRyYV9oZWxtX3ZhbHVlcxgFIAEoCzIXLmdvb2dsZS5wcm90b2J1Zi5TdHJ1Y3RIB4gBARIUCgdwcm9maWxlGAYgASgJSAiIAQESHAoPZXh0ZXJuYWxfc2VydmVyGAcgASgISAmIAQESHAoPb2lkY19pc3N1ZXJfdXJsGAsgASgJSAqIAQFCBQoDX2lkQgcKBV9uYW1lQgkKB19vcmdfaWRCDQoLX3RydXN0X3pvbmVCEAoOX3RydXN0X3pvbmVfaWRCFQoTX2t1YmVybmV0ZXNfY29udGV4dEIRCg9fdHJ1c3RfcHJvdmlkZXJCFAoSX2V4dHJhX2hlbG1fdmFsdWVzQgoKCF9wcm9maWxlQhIKEF9leHRlcm5hbF9zZXJ2ZXJCEgoQX29pZGNfaXNzdWVyX3VybEJAWj5naXRodWIuY29tL2NvZmlkZS9jb2ZpZGUtYXBpLXNkay9nZW4vZ28vcHJvdG8vY2x1c3Rlci92MWFscGhhMWIGcHJvdG8z", [file_google_protobuf_struct, file_proto_trust_provider_v1alpha1_trust_provider]);

/**
 * @generated from message proto.cluster.v1alpha1.Cluster
 */
export type Cluster = Message<"proto.cluster.v1alpha1.Cluster"> & {
  /**
   * @generated from field: optional string id = 8;
   */
  id?: string;

  /**
   * @generated from field: optional string name = 1;
   */
  name?: string;

  /**
   * @generated from field: optional string org_id = 9;
   */
  orgId?: string;

  /**
   * DEPRECATED: replaced by trust_zone_id.
   *
   * @generated from field: optional string trust_zone = 2;
   */
  trustZone?: string;

  /**
   * @generated from field: optional string trust_zone_id = 10;
   */
  trustZoneId?: string;

  /**
   * @generated from field: optional string kubernetes_context = 3;
   */
  kubernetesContext?: string;

  /**
   * @generated from field: optional proto.trust_provider.v1alpha1.TrustProvider trust_provider = 4;
   */
  trustProvider?: TrustProvider;

  /**
   * @generated from field: optional google.protobuf.Struct extra_helm_values = 5;
   */
  extraHelmValues?: JsonObject;

  /**
   * @generated from field: optional string profile = 6;
   */
  profile?: string;

  /**
   * @generated from field: optional bool external_server = 7;
   */
  externalServer?: boolean;

  /**
   * @generated from field: optional string oidc_issuer_url = 11;
   */
  oidcIssuerUrl?: string;
};

/**
 * Describes the message proto.cluster.v1alpha1.Cluster.
 * Use `create(ClusterSchema)` to create a new message.
 */
export const ClusterSchema: GenMessage<Cluster> = /*@__PURE__*/
  messageDesc(file_proto_cluster_v1alpha1_cluster, 0);

