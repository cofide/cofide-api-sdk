import * as jspb from 'google-protobuf'

import * as proto_ap_binding_v1alpha1_ap_binding_pb from '../../../proto/ap_binding/v1alpha1/ap_binding_pb'; // proto import: "proto/ap_binding/v1alpha1/ap_binding.proto"
import * as proto_attestation_policy_v1alpha1_attestation_policy_pb from '../../../proto/attestation_policy/v1alpha1/attestation_policy_pb'; // proto import: "proto/attestation_policy/v1alpha1/attestation_policy.proto"
import * as proto_cluster_v1alpha1_cluster_pb from '../../../proto/cluster/v1alpha1/cluster_pb'; // proto import: "proto/cluster/v1alpha1/cluster.proto"
import * as proto_federation_v1alpha1_federation_pb from '../../../proto/federation/v1alpha1/federation_pb'; // proto import: "proto/federation/v1alpha1/federation.proto"
import * as proto_trust_zone_v1alpha1_trust_zone_pb from '../../../proto/trust_zone/v1alpha1/trust_zone_pb'; // proto import: "proto/trust_zone/v1alpha1/trust_zone.proto"


export class ValidateRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ValidateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ValidateRequest): ValidateRequest.AsObject;
  static serializeBinaryToWriter(message: ValidateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ValidateRequest;
  static deserializeBinaryFromReader(message: ValidateRequest, reader: jspb.BinaryReader): ValidateRequest;
}

export namespace ValidateRequest {
  export type AsObject = {
  }
}

export class ValidateResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ValidateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ValidateResponse): ValidateResponse.AsObject;
  static serializeBinaryToWriter(message: ValidateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ValidateResponse;
  static deserializeBinaryFromReader(message: ValidateResponse, reader: jspb.BinaryReader): ValidateResponse;
}

export namespace ValidateResponse {
  export type AsObject = {
  }
}

export class AddTrustZoneRequest extends jspb.Message {
  getTrustZone(): proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone | undefined;
  setTrustZone(value?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone): AddTrustZoneRequest;
  hasTrustZone(): boolean;
  clearTrustZone(): AddTrustZoneRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddTrustZoneRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddTrustZoneRequest): AddTrustZoneRequest.AsObject;
  static serializeBinaryToWriter(message: AddTrustZoneRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddTrustZoneRequest;
  static deserializeBinaryFromReader(message: AddTrustZoneRequest, reader: jspb.BinaryReader): AddTrustZoneRequest;
}

export namespace AddTrustZoneRequest {
  export type AsObject = {
    trustZone?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone.AsObject,
  }

  export enum TrustZoneCase { 
    _TRUST_ZONE_NOT_SET = 0,
    TRUST_ZONE = 1,
  }
}

export class AddTrustZoneResponse extends jspb.Message {
  getTrustZone(): proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone | undefined;
  setTrustZone(value?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone): AddTrustZoneResponse;
  hasTrustZone(): boolean;
  clearTrustZone(): AddTrustZoneResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddTrustZoneResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddTrustZoneResponse): AddTrustZoneResponse.AsObject;
  static serializeBinaryToWriter(message: AddTrustZoneResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddTrustZoneResponse;
  static deserializeBinaryFromReader(message: AddTrustZoneResponse, reader: jspb.BinaryReader): AddTrustZoneResponse;
}

export namespace AddTrustZoneResponse {
  export type AsObject = {
    trustZone?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone.AsObject,
  }

  export enum TrustZoneCase { 
    _TRUST_ZONE_NOT_SET = 0,
    TRUST_ZONE = 1,
  }
}

export class GetTrustZoneRequest extends jspb.Message {
  getName(): string;
  setName(value: string): GetTrustZoneRequest;
  hasName(): boolean;
  clearName(): GetTrustZoneRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTrustZoneRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTrustZoneRequest): GetTrustZoneRequest.AsObject;
  static serializeBinaryToWriter(message: GetTrustZoneRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTrustZoneRequest;
  static deserializeBinaryFromReader(message: GetTrustZoneRequest, reader: jspb.BinaryReader): GetTrustZoneRequest;
}

export namespace GetTrustZoneRequest {
  export type AsObject = {
    name?: string,
  }

  export enum NameCase { 
    _NAME_NOT_SET = 0,
    NAME = 1,
  }
}

export class GetTrustZoneResponse extends jspb.Message {
  getTrustZone(): proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone | undefined;
  setTrustZone(value?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone): GetTrustZoneResponse;
  hasTrustZone(): boolean;
  clearTrustZone(): GetTrustZoneResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTrustZoneResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTrustZoneResponse): GetTrustZoneResponse.AsObject;
  static serializeBinaryToWriter(message: GetTrustZoneResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTrustZoneResponse;
  static deserializeBinaryFromReader(message: GetTrustZoneResponse, reader: jspb.BinaryReader): GetTrustZoneResponse;
}

export namespace GetTrustZoneResponse {
  export type AsObject = {
    trustZone?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone.AsObject,
  }

  export enum TrustZoneCase { 
    _TRUST_ZONE_NOT_SET = 0,
    TRUST_ZONE = 1,
  }
}

export class ListTrustZonesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTrustZonesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListTrustZonesRequest): ListTrustZonesRequest.AsObject;
  static serializeBinaryToWriter(message: ListTrustZonesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTrustZonesRequest;
  static deserializeBinaryFromReader(message: ListTrustZonesRequest, reader: jspb.BinaryReader): ListTrustZonesRequest;
}

export namespace ListTrustZonesRequest {
  export type AsObject = {
  }
}

export class ListTrustZonesResponse extends jspb.Message {
  getTrustZonesList(): Array<proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone>;
  setTrustZonesList(value: Array<proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone>): ListTrustZonesResponse;
  clearTrustZonesList(): ListTrustZonesResponse;
  addTrustZones(value?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone, index?: number): proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTrustZonesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListTrustZonesResponse): ListTrustZonesResponse.AsObject;
  static serializeBinaryToWriter(message: ListTrustZonesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTrustZonesResponse;
  static deserializeBinaryFromReader(message: ListTrustZonesResponse, reader: jspb.BinaryReader): ListTrustZonesResponse;
}

export namespace ListTrustZonesResponse {
  export type AsObject = {
    trustZonesList: Array<proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone.AsObject>,
  }
}

export class UpdateTrustZoneRequest extends jspb.Message {
  getTrustZone(): proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone | undefined;
  setTrustZone(value?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone): UpdateTrustZoneRequest;
  hasTrustZone(): boolean;
  clearTrustZone(): UpdateTrustZoneRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateTrustZoneRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateTrustZoneRequest): UpdateTrustZoneRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateTrustZoneRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateTrustZoneRequest;
  static deserializeBinaryFromReader(message: UpdateTrustZoneRequest, reader: jspb.BinaryReader): UpdateTrustZoneRequest;
}

export namespace UpdateTrustZoneRequest {
  export type AsObject = {
    trustZone?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone.AsObject,
  }

  export enum TrustZoneCase { 
    _TRUST_ZONE_NOT_SET = 0,
    TRUST_ZONE = 1,
  }
}

export class UpdateTrustZoneResponse extends jspb.Message {
  getTrustZone(): proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone | undefined;
  setTrustZone(value?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone): UpdateTrustZoneResponse;
  hasTrustZone(): boolean;
  clearTrustZone(): UpdateTrustZoneResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateTrustZoneResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateTrustZoneResponse): UpdateTrustZoneResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateTrustZoneResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateTrustZoneResponse;
  static deserializeBinaryFromReader(message: UpdateTrustZoneResponse, reader: jspb.BinaryReader): UpdateTrustZoneResponse;
}

export namespace UpdateTrustZoneResponse {
  export type AsObject = {
    trustZone?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone.AsObject,
  }

  export enum TrustZoneCase { 
    _TRUST_ZONE_NOT_SET = 0,
    TRUST_ZONE = 1,
  }
}

export class AddClusterRequest extends jspb.Message {
  getCluster(): proto_cluster_v1alpha1_cluster_pb.Cluster | undefined;
  setCluster(value?: proto_cluster_v1alpha1_cluster_pb.Cluster): AddClusterRequest;
  hasCluster(): boolean;
  clearCluster(): AddClusterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddClusterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddClusterRequest): AddClusterRequest.AsObject;
  static serializeBinaryToWriter(message: AddClusterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddClusterRequest;
  static deserializeBinaryFromReader(message: AddClusterRequest, reader: jspb.BinaryReader): AddClusterRequest;
}

export namespace AddClusterRequest {
  export type AsObject = {
    cluster?: proto_cluster_v1alpha1_cluster_pb.Cluster.AsObject,
  }

  export enum ClusterCase { 
    _CLUSTER_NOT_SET = 0,
    CLUSTER = 1,
  }
}

export class AddClusterResponse extends jspb.Message {
  getCluster(): proto_cluster_v1alpha1_cluster_pb.Cluster | undefined;
  setCluster(value?: proto_cluster_v1alpha1_cluster_pb.Cluster): AddClusterResponse;
  hasCluster(): boolean;
  clearCluster(): AddClusterResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddClusterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddClusterResponse): AddClusterResponse.AsObject;
  static serializeBinaryToWriter(message: AddClusterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddClusterResponse;
  static deserializeBinaryFromReader(message: AddClusterResponse, reader: jspb.BinaryReader): AddClusterResponse;
}

export namespace AddClusterResponse {
  export type AsObject = {
    cluster?: proto_cluster_v1alpha1_cluster_pb.Cluster.AsObject,
  }

  export enum ClusterCase { 
    _CLUSTER_NOT_SET = 0,
    CLUSTER = 1,
  }
}

export class GetClusterRequest extends jspb.Message {
  getName(): string;
  setName(value: string): GetClusterRequest;
  hasName(): boolean;
  clearName(): GetClusterRequest;

  getTrustZone(): string;
  setTrustZone(value: string): GetClusterRequest;
  hasTrustZone(): boolean;
  clearTrustZone(): GetClusterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClusterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetClusterRequest): GetClusterRequest.AsObject;
  static serializeBinaryToWriter(message: GetClusterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClusterRequest;
  static deserializeBinaryFromReader(message: GetClusterRequest, reader: jspb.BinaryReader): GetClusterRequest;
}

export namespace GetClusterRequest {
  export type AsObject = {
    name?: string,
    trustZone?: string,
  }

  export enum NameCase { 
    _NAME_NOT_SET = 0,
    NAME = 1,
  }

  export enum TrustZoneCase { 
    _TRUST_ZONE_NOT_SET = 0,
    TRUST_ZONE = 2,
  }
}

export class GetClusterResponse extends jspb.Message {
  getCluster(): proto_cluster_v1alpha1_cluster_pb.Cluster | undefined;
  setCluster(value?: proto_cluster_v1alpha1_cluster_pb.Cluster): GetClusterResponse;
  hasCluster(): boolean;
  clearCluster(): GetClusterResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClusterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetClusterResponse): GetClusterResponse.AsObject;
  static serializeBinaryToWriter(message: GetClusterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClusterResponse;
  static deserializeBinaryFromReader(message: GetClusterResponse, reader: jspb.BinaryReader): GetClusterResponse;
}

export namespace GetClusterResponse {
  export type AsObject = {
    cluster?: proto_cluster_v1alpha1_cluster_pb.Cluster.AsObject,
  }

  export enum ClusterCase { 
    _CLUSTER_NOT_SET = 0,
    CLUSTER = 1,
  }
}

export class ListClustersRequest extends jspb.Message {
  getTrustZone(): string;
  setTrustZone(value: string): ListClustersRequest;
  hasTrustZone(): boolean;
  clearTrustZone(): ListClustersRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListClustersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListClustersRequest): ListClustersRequest.AsObject;
  static serializeBinaryToWriter(message: ListClustersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListClustersRequest;
  static deserializeBinaryFromReader(message: ListClustersRequest, reader: jspb.BinaryReader): ListClustersRequest;
}

export namespace ListClustersRequest {
  export type AsObject = {
    trustZone?: string,
  }

  export enum TrustZoneCase { 
    _TRUST_ZONE_NOT_SET = 0,
    TRUST_ZONE = 1,
  }
}

export class ListClustersResponse extends jspb.Message {
  getClustersList(): Array<proto_cluster_v1alpha1_cluster_pb.Cluster>;
  setClustersList(value: Array<proto_cluster_v1alpha1_cluster_pb.Cluster>): ListClustersResponse;
  clearClustersList(): ListClustersResponse;
  addClusters(value?: proto_cluster_v1alpha1_cluster_pb.Cluster, index?: number): proto_cluster_v1alpha1_cluster_pb.Cluster;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListClustersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListClustersResponse): ListClustersResponse.AsObject;
  static serializeBinaryToWriter(message: ListClustersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListClustersResponse;
  static deserializeBinaryFromReader(message: ListClustersResponse, reader: jspb.BinaryReader): ListClustersResponse;
}

export namespace ListClustersResponse {
  export type AsObject = {
    clustersList: Array<proto_cluster_v1alpha1_cluster_pb.Cluster.AsObject>,
  }
}

export class UpdateClusterRequest extends jspb.Message {
  getCluster(): proto_cluster_v1alpha1_cluster_pb.Cluster | undefined;
  setCluster(value?: proto_cluster_v1alpha1_cluster_pb.Cluster): UpdateClusterRequest;
  hasCluster(): boolean;
  clearCluster(): UpdateClusterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateClusterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateClusterRequest): UpdateClusterRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateClusterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateClusterRequest;
  static deserializeBinaryFromReader(message: UpdateClusterRequest, reader: jspb.BinaryReader): UpdateClusterRequest;
}

export namespace UpdateClusterRequest {
  export type AsObject = {
    cluster?: proto_cluster_v1alpha1_cluster_pb.Cluster.AsObject,
  }

  export enum ClusterCase { 
    _CLUSTER_NOT_SET = 0,
    CLUSTER = 1,
  }
}

export class UpdateClusterResponse extends jspb.Message {
  getCluster(): proto_cluster_v1alpha1_cluster_pb.Cluster | undefined;
  setCluster(value?: proto_cluster_v1alpha1_cluster_pb.Cluster): UpdateClusterResponse;
  hasCluster(): boolean;
  clearCluster(): UpdateClusterResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateClusterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateClusterResponse): UpdateClusterResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateClusterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateClusterResponse;
  static deserializeBinaryFromReader(message: UpdateClusterResponse, reader: jspb.BinaryReader): UpdateClusterResponse;
}

export namespace UpdateClusterResponse {
  export type AsObject = {
    cluster?: proto_cluster_v1alpha1_cluster_pb.Cluster.AsObject,
  }

  export enum ClusterCase { 
    _CLUSTER_NOT_SET = 0,
    CLUSTER = 1,
  }
}

export class AddAttestationPolicyRequest extends jspb.Message {
  getPolicy(): proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy | undefined;
  setPolicy(value?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy): AddAttestationPolicyRequest;
  hasPolicy(): boolean;
  clearPolicy(): AddAttestationPolicyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddAttestationPolicyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddAttestationPolicyRequest): AddAttestationPolicyRequest.AsObject;
  static serializeBinaryToWriter(message: AddAttestationPolicyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddAttestationPolicyRequest;
  static deserializeBinaryFromReader(message: AddAttestationPolicyRequest, reader: jspb.BinaryReader): AddAttestationPolicyRequest;
}

export namespace AddAttestationPolicyRequest {
  export type AsObject = {
    policy?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy.AsObject,
  }

  export enum PolicyCase { 
    _POLICY_NOT_SET = 0,
    POLICY = 1,
  }
}

export class AddAttestationPolicyResponse extends jspb.Message {
  getPolicy(): proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy | undefined;
  setPolicy(value?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy): AddAttestationPolicyResponse;
  hasPolicy(): boolean;
  clearPolicy(): AddAttestationPolicyResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddAttestationPolicyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddAttestationPolicyResponse): AddAttestationPolicyResponse.AsObject;
  static serializeBinaryToWriter(message: AddAttestationPolicyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddAttestationPolicyResponse;
  static deserializeBinaryFromReader(message: AddAttestationPolicyResponse, reader: jspb.BinaryReader): AddAttestationPolicyResponse;
}

export namespace AddAttestationPolicyResponse {
  export type AsObject = {
    policy?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy.AsObject,
  }

  export enum PolicyCase { 
    _POLICY_NOT_SET = 0,
    POLICY = 1,
  }
}

export class GetAttestationPolicyRequest extends jspb.Message {
  getName(): string;
  setName(value: string): GetAttestationPolicyRequest;
  hasName(): boolean;
  clearName(): GetAttestationPolicyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAttestationPolicyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAttestationPolicyRequest): GetAttestationPolicyRequest.AsObject;
  static serializeBinaryToWriter(message: GetAttestationPolicyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAttestationPolicyRequest;
  static deserializeBinaryFromReader(message: GetAttestationPolicyRequest, reader: jspb.BinaryReader): GetAttestationPolicyRequest;
}

export namespace GetAttestationPolicyRequest {
  export type AsObject = {
    name?: string,
  }

  export enum NameCase { 
    _NAME_NOT_SET = 0,
    NAME = 1,
  }
}

export class GetAttestationPolicyResponse extends jspb.Message {
  getPolicy(): proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy | undefined;
  setPolicy(value?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy): GetAttestationPolicyResponse;
  hasPolicy(): boolean;
  clearPolicy(): GetAttestationPolicyResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAttestationPolicyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAttestationPolicyResponse): GetAttestationPolicyResponse.AsObject;
  static serializeBinaryToWriter(message: GetAttestationPolicyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAttestationPolicyResponse;
  static deserializeBinaryFromReader(message: GetAttestationPolicyResponse, reader: jspb.BinaryReader): GetAttestationPolicyResponse;
}

export namespace GetAttestationPolicyResponse {
  export type AsObject = {
    policy?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy.AsObject,
  }
}

export class ListAttestationPoliciesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAttestationPoliciesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListAttestationPoliciesRequest): ListAttestationPoliciesRequest.AsObject;
  static serializeBinaryToWriter(message: ListAttestationPoliciesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAttestationPoliciesRequest;
  static deserializeBinaryFromReader(message: ListAttestationPoliciesRequest, reader: jspb.BinaryReader): ListAttestationPoliciesRequest;
}

export namespace ListAttestationPoliciesRequest {
  export type AsObject = {
  }
}

export class ListAttestationPoliciesResponse extends jspb.Message {
  getPoliciesList(): Array<proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy>;
  setPoliciesList(value: Array<proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy>): ListAttestationPoliciesResponse;
  clearPoliciesList(): ListAttestationPoliciesResponse;
  addPolicies(value?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy, index?: number): proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAttestationPoliciesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListAttestationPoliciesResponse): ListAttestationPoliciesResponse.AsObject;
  static serializeBinaryToWriter(message: ListAttestationPoliciesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAttestationPoliciesResponse;
  static deserializeBinaryFromReader(message: ListAttestationPoliciesResponse, reader: jspb.BinaryReader): ListAttestationPoliciesResponse;
}

export namespace ListAttestationPoliciesResponse {
  export type AsObject = {
    policiesList: Array<proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy.AsObject>,
  }
}

export class AddAPBindingRequest extends jspb.Message {
  getBinding(): proto_ap_binding_v1alpha1_ap_binding_pb.APBinding | undefined;
  setBinding(value?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding): AddAPBindingRequest;
  hasBinding(): boolean;
  clearBinding(): AddAPBindingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddAPBindingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddAPBindingRequest): AddAPBindingRequest.AsObject;
  static serializeBinaryToWriter(message: AddAPBindingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddAPBindingRequest;
  static deserializeBinaryFromReader(message: AddAPBindingRequest, reader: jspb.BinaryReader): AddAPBindingRequest;
}

export namespace AddAPBindingRequest {
  export type AsObject = {
    binding?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding.AsObject,
  }

  export enum BindingCase { 
    _BINDING_NOT_SET = 0,
    BINDING = 1,
  }
}

export class AddAPBindingResponse extends jspb.Message {
  getBinding(): proto_ap_binding_v1alpha1_ap_binding_pb.APBinding | undefined;
  setBinding(value?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding): AddAPBindingResponse;
  hasBinding(): boolean;
  clearBinding(): AddAPBindingResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddAPBindingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddAPBindingResponse): AddAPBindingResponse.AsObject;
  static serializeBinaryToWriter(message: AddAPBindingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddAPBindingResponse;
  static deserializeBinaryFromReader(message: AddAPBindingResponse, reader: jspb.BinaryReader): AddAPBindingResponse;
}

export namespace AddAPBindingResponse {
  export type AsObject = {
    binding?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding.AsObject,
  }

  export enum BindingCase { 
    _BINDING_NOT_SET = 0,
    BINDING = 1,
  }
}

export class DestroyAPBindingRequest extends jspb.Message {
  getBinding(): proto_ap_binding_v1alpha1_ap_binding_pb.APBinding | undefined;
  setBinding(value?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding): DestroyAPBindingRequest;
  hasBinding(): boolean;
  clearBinding(): DestroyAPBindingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestroyAPBindingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DestroyAPBindingRequest): DestroyAPBindingRequest.AsObject;
  static serializeBinaryToWriter(message: DestroyAPBindingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestroyAPBindingRequest;
  static deserializeBinaryFromReader(message: DestroyAPBindingRequest, reader: jspb.BinaryReader): DestroyAPBindingRequest;
}

export namespace DestroyAPBindingRequest {
  export type AsObject = {
    binding?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding.AsObject,
  }

  export enum BindingCase { 
    _BINDING_NOT_SET = 0,
    BINDING = 1,
  }
}

export class DestroyAPBindingResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestroyAPBindingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DestroyAPBindingResponse): DestroyAPBindingResponse.AsObject;
  static serializeBinaryToWriter(message: DestroyAPBindingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestroyAPBindingResponse;
  static deserializeBinaryFromReader(message: DestroyAPBindingResponse, reader: jspb.BinaryReader): DestroyAPBindingResponse;
}

export namespace DestroyAPBindingResponse {
  export type AsObject = {
  }
}

export class ListAPBindingsRequest extends jspb.Message {
  getFilter(): ListAPBindingsRequest.Filter | undefined;
  setFilter(value?: ListAPBindingsRequest.Filter): ListAPBindingsRequest;
  hasFilter(): boolean;
  clearFilter(): ListAPBindingsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAPBindingsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListAPBindingsRequest): ListAPBindingsRequest.AsObject;
  static serializeBinaryToWriter(message: ListAPBindingsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAPBindingsRequest;
  static deserializeBinaryFromReader(message: ListAPBindingsRequest, reader: jspb.BinaryReader): ListAPBindingsRequest;
}

export namespace ListAPBindingsRequest {
  export type AsObject = {
    filter?: ListAPBindingsRequest.Filter.AsObject,
  }

  export class Filter extends jspb.Message {
    getTrustZoneName(): string;
    setTrustZoneName(value: string): Filter;
    hasTrustZoneName(): boolean;
    clearTrustZoneName(): Filter;

    getPolicyName(): string;
    setPolicyName(value: string): Filter;
    hasPolicyName(): boolean;
    clearPolicyName(): Filter;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Filter.AsObject;
    static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
    static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Filter;
    static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
  }

  export namespace Filter {
    export type AsObject = {
      trustZoneName?: string,
      policyName?: string,
    }

    export enum TrustZoneNameCase { 
      _TRUST_ZONE_NAME_NOT_SET = 0,
      TRUST_ZONE_NAME = 1,
    }

    export enum PolicyNameCase { 
      _POLICY_NAME_NOT_SET = 0,
      POLICY_NAME = 2,
    }
  }


  export enum FilterCase { 
    _FILTER_NOT_SET = 0,
    FILTER = 1,
  }
}

export class ListAPBindingsResponse extends jspb.Message {
  getBindingsList(): Array<proto_ap_binding_v1alpha1_ap_binding_pb.APBinding>;
  setBindingsList(value: Array<proto_ap_binding_v1alpha1_ap_binding_pb.APBinding>): ListAPBindingsResponse;
  clearBindingsList(): ListAPBindingsResponse;
  addBindings(value?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding, index?: number): proto_ap_binding_v1alpha1_ap_binding_pb.APBinding;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAPBindingsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListAPBindingsResponse): ListAPBindingsResponse.AsObject;
  static serializeBinaryToWriter(message: ListAPBindingsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAPBindingsResponse;
  static deserializeBinaryFromReader(message: ListAPBindingsResponse, reader: jspb.BinaryReader): ListAPBindingsResponse;
}

export namespace ListAPBindingsResponse {
  export type AsObject = {
    bindingsList: Array<proto_ap_binding_v1alpha1_ap_binding_pb.APBinding.AsObject>,
  }
}

export class AddFederationRequest extends jspb.Message {
  getFederation(): proto_federation_v1alpha1_federation_pb.Federation | undefined;
  setFederation(value?: proto_federation_v1alpha1_federation_pb.Federation): AddFederationRequest;
  hasFederation(): boolean;
  clearFederation(): AddFederationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddFederationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddFederationRequest): AddFederationRequest.AsObject;
  static serializeBinaryToWriter(message: AddFederationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddFederationRequest;
  static deserializeBinaryFromReader(message: AddFederationRequest, reader: jspb.BinaryReader): AddFederationRequest;
}

export namespace AddFederationRequest {
  export type AsObject = {
    federation?: proto_federation_v1alpha1_federation_pb.Federation.AsObject,
  }

  export enum FederationCase { 
    _FEDERATION_NOT_SET = 0,
    FEDERATION = 1,
  }
}

export class AddFederationResponse extends jspb.Message {
  getFederation(): proto_federation_v1alpha1_federation_pb.Federation | undefined;
  setFederation(value?: proto_federation_v1alpha1_federation_pb.Federation): AddFederationResponse;
  hasFederation(): boolean;
  clearFederation(): AddFederationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddFederationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddFederationResponse): AddFederationResponse.AsObject;
  static serializeBinaryToWriter(message: AddFederationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddFederationResponse;
  static deserializeBinaryFromReader(message: AddFederationResponse, reader: jspb.BinaryReader): AddFederationResponse;
}

export namespace AddFederationResponse {
  export type AsObject = {
    federation?: proto_federation_v1alpha1_federation_pb.Federation.AsObject,
  }

  export enum FederationCase { 
    _FEDERATION_NOT_SET = 0,
    FEDERATION = 1,
  }
}

export class ListFederationsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFederationsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListFederationsRequest): ListFederationsRequest.AsObject;
  static serializeBinaryToWriter(message: ListFederationsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFederationsRequest;
  static deserializeBinaryFromReader(message: ListFederationsRequest, reader: jspb.BinaryReader): ListFederationsRequest;
}

export namespace ListFederationsRequest {
  export type AsObject = {
  }
}

export class ListFederationsResponse extends jspb.Message {
  getFederationsList(): Array<proto_federation_v1alpha1_federation_pb.Federation>;
  setFederationsList(value: Array<proto_federation_v1alpha1_federation_pb.Federation>): ListFederationsResponse;
  clearFederationsList(): ListFederationsResponse;
  addFederations(value?: proto_federation_v1alpha1_federation_pb.Federation, index?: number): proto_federation_v1alpha1_federation_pb.Federation;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFederationsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListFederationsResponse): ListFederationsResponse.AsObject;
  static serializeBinaryToWriter(message: ListFederationsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFederationsResponse;
  static deserializeBinaryFromReader(message: ListFederationsResponse, reader: jspb.BinaryReader): ListFederationsResponse;
}

export namespace ListFederationsResponse {
  export type AsObject = {
    federationsList: Array<proto_federation_v1alpha1_federation_pb.Federation.AsObject>,
  }
}

export class ListFederationsByTrustZoneRequest extends jspb.Message {
  getTrustZoneName(): string;
  setTrustZoneName(value: string): ListFederationsByTrustZoneRequest;
  hasTrustZoneName(): boolean;
  clearTrustZoneName(): ListFederationsByTrustZoneRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFederationsByTrustZoneRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListFederationsByTrustZoneRequest): ListFederationsByTrustZoneRequest.AsObject;
  static serializeBinaryToWriter(message: ListFederationsByTrustZoneRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFederationsByTrustZoneRequest;
  static deserializeBinaryFromReader(message: ListFederationsByTrustZoneRequest, reader: jspb.BinaryReader): ListFederationsByTrustZoneRequest;
}

export namespace ListFederationsByTrustZoneRequest {
  export type AsObject = {
    trustZoneName?: string,
  }

  export enum TrustZoneNameCase { 
    _TRUST_ZONE_NAME_NOT_SET = 0,
    TRUST_ZONE_NAME = 1,
  }
}

export class ListFederationsByTrustZoneResponse extends jspb.Message {
  getFederationsList(): Array<proto_federation_v1alpha1_federation_pb.Federation>;
  setFederationsList(value: Array<proto_federation_v1alpha1_federation_pb.Federation>): ListFederationsByTrustZoneResponse;
  clearFederationsList(): ListFederationsByTrustZoneResponse;
  addFederations(value?: proto_federation_v1alpha1_federation_pb.Federation, index?: number): proto_federation_v1alpha1_federation_pb.Federation;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFederationsByTrustZoneResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListFederationsByTrustZoneResponse): ListFederationsByTrustZoneResponse.AsObject;
  static serializeBinaryToWriter(message: ListFederationsByTrustZoneResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFederationsByTrustZoneResponse;
  static deserializeBinaryFromReader(message: ListFederationsByTrustZoneResponse, reader: jspb.BinaryReader): ListFederationsByTrustZoneResponse;
}

export namespace ListFederationsByTrustZoneResponse {
  export type AsObject = {
    federationsList: Array<proto_federation_v1alpha1_federation_pb.Federation.AsObject>,
  }
}

