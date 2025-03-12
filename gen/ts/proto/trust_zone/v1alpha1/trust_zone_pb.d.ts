import * as jspb from 'google-protobuf'

import * as proto_ap_binding_v1alpha1_ap_binding_pb from '../../../proto/ap_binding/v1alpha1/ap_binding_pb'; // proto import: "proto/ap_binding/v1alpha1/ap_binding.proto"
import * as proto_federation_v1alpha1_federation_pb from '../../../proto/federation/v1alpha1/federation_pb'; // proto import: "proto/federation/v1alpha1/federation.proto"
import * as proto_spire_api_types_bundle_pb from '../../../proto/spire/api/types/bundle_pb'; // proto import: "proto/spire/api/types/bundle.proto"


export class TrustZone extends jspb.Message {
  getName(): string;
  setName(value: string): TrustZone;

  getTrustDomain(): string;
  setTrustDomain(value: string): TrustZone;

  getBundleEndpointUrl(): string;
  setBundleEndpointUrl(value: string): TrustZone;
  hasBundleEndpointUrl(): boolean;
  clearBundleEndpointUrl(): TrustZone;

  getBundle(): proto_spire_api_types_bundle_pb.Bundle | undefined;
  setBundle(value?: proto_spire_api_types_bundle_pb.Bundle): TrustZone;
  hasBundle(): boolean;
  clearBundle(): TrustZone;

  getFederationsList(): Array<proto_federation_v1alpha1_federation_pb.Federation>;
  setFederationsList(value: Array<proto_federation_v1alpha1_federation_pb.Federation>): TrustZone;
  clearFederationsList(): TrustZone;
  addFederations(value?: proto_federation_v1alpha1_federation_pb.Federation, index?: number): proto_federation_v1alpha1_federation_pb.Federation;

  getAttestationPoliciesList(): Array<proto_ap_binding_v1alpha1_ap_binding_pb.APBinding>;
  setAttestationPoliciesList(value: Array<proto_ap_binding_v1alpha1_ap_binding_pb.APBinding>): TrustZone;
  clearAttestationPoliciesList(): TrustZone;
  addAttestationPolicies(value?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding, index?: number): proto_ap_binding_v1alpha1_ap_binding_pb.APBinding;

  getJwtIssuer(): string;
  setJwtIssuer(value: string): TrustZone;
  hasJwtIssuer(): boolean;
  clearJwtIssuer(): TrustZone;

  getBundleEndpointProfile(): BundleEndpointProfile;
  setBundleEndpointProfile(value: BundleEndpointProfile): TrustZone;
  hasBundleEndpointProfile(): boolean;
  clearBundleEndpointProfile(): TrustZone;

  getId(): string;
  setId(value: string): TrustZone;
  hasId(): boolean;
  clearId(): TrustZone;

  getIsManagementZone(): boolean;
  setIsManagementZone(value: boolean): TrustZone;

  getOrgId(): string;
  setOrgId(value: string): TrustZone;
  hasOrgId(): boolean;
  clearOrgId(): TrustZone;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TrustZone.AsObject;
  static toObject(includeInstance: boolean, msg: TrustZone): TrustZone.AsObject;
  static serializeBinaryToWriter(message: TrustZone, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TrustZone;
  static deserializeBinaryFromReader(message: TrustZone, reader: jspb.BinaryReader): TrustZone;
}

export namespace TrustZone {
  export type AsObject = {
    name: string,
    trustDomain: string,
    bundleEndpointUrl?: string,
    bundle?: proto_spire_api_types_bundle_pb.Bundle.AsObject,
    federationsList: Array<proto_federation_v1alpha1_federation_pb.Federation.AsObject>,
    attestationPoliciesList: Array<proto_ap_binding_v1alpha1_ap_binding_pb.APBinding.AsObject>,
    jwtIssuer?: string,
    bundleEndpointProfile?: BundleEndpointProfile,
    id?: string,
    isManagementZone: boolean,
    orgId?: string,
  }

  export enum BundleEndpointUrlCase { 
    _BUNDLE_ENDPOINT_URL_NOT_SET = 0,
    BUNDLE_ENDPOINT_URL = 3,
  }

  export enum BundleCase { 
    _BUNDLE_NOT_SET = 0,
    BUNDLE = 4,
  }

  export enum JwtIssuerCase { 
    _JWT_ISSUER_NOT_SET = 0,
    JWT_ISSUER = 7,
  }

  export enum BundleEndpointProfileCase { 
    _BUNDLE_ENDPOINT_PROFILE_NOT_SET = 0,
    BUNDLE_ENDPOINT_PROFILE = 8,
  }

  export enum IdCase { 
    _ID_NOT_SET = 0,
    ID = 9,
  }

  export enum OrgIdCase { 
    _ORG_ID_NOT_SET = 0,
    ORG_ID = 11,
  }
}

export enum BundleEndpointProfile { 
  BUNDLE_ENDPOINT_PROFILE_UNSPECIFIED = 0,
  BUNDLE_ENDPOINT_PROFILE_HTTPS_SPIFFE = 1,
  BUNDLE_ENDPOINT_PROFILE_HTTPS_WEB = 2,
}
