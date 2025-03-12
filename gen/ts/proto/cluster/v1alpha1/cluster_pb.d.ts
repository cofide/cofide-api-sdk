import * as jspb from 'google-protobuf'

import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb'; // proto import: "google/protobuf/struct.proto"
import * as proto_trust_provider_v1alpha1_trust_provider_pb from '../../../proto/trust_provider/v1alpha1/trust_provider_pb'; // proto import: "proto/trust_provider/v1alpha1/trust_provider.proto"


export class Cluster extends jspb.Message {
  getId(): string;
  setId(value: string): Cluster;
  hasId(): boolean;
  clearId(): Cluster;

  getName(): string;
  setName(value: string): Cluster;
  hasName(): boolean;
  clearName(): Cluster;

  getOrgId(): string;
  setOrgId(value: string): Cluster;
  hasOrgId(): boolean;
  clearOrgId(): Cluster;

  getTrustZone(): string;
  setTrustZone(value: string): Cluster;
  hasTrustZone(): boolean;
  clearTrustZone(): Cluster;

  getTrustZoneId(): string;
  setTrustZoneId(value: string): Cluster;
  hasTrustZoneId(): boolean;
  clearTrustZoneId(): Cluster;

  getKubernetesContext(): string;
  setKubernetesContext(value: string): Cluster;
  hasKubernetesContext(): boolean;
  clearKubernetesContext(): Cluster;

  getTrustProvider(): proto_trust_provider_v1alpha1_trust_provider_pb.TrustProvider | undefined;
  setTrustProvider(value?: proto_trust_provider_v1alpha1_trust_provider_pb.TrustProvider): Cluster;
  hasTrustProvider(): boolean;
  clearTrustProvider(): Cluster;

  getExtraHelmValues(): google_protobuf_struct_pb.Struct | undefined;
  setExtraHelmValues(value?: google_protobuf_struct_pb.Struct): Cluster;
  hasExtraHelmValues(): boolean;
  clearExtraHelmValues(): Cluster;

  getProfile(): string;
  setProfile(value: string): Cluster;
  hasProfile(): boolean;
  clearProfile(): Cluster;

  getExternalServer(): boolean;
  setExternalServer(value: boolean): Cluster;
  hasExternalServer(): boolean;
  clearExternalServer(): Cluster;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Cluster.AsObject;
  static toObject(includeInstance: boolean, msg: Cluster): Cluster.AsObject;
  static serializeBinaryToWriter(message: Cluster, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Cluster;
  static deserializeBinaryFromReader(message: Cluster, reader: jspb.BinaryReader): Cluster;
}

export namespace Cluster {
  export type AsObject = {
    id?: string,
    name?: string,
    orgId?: string,
    trustZone?: string,
    trustZoneId?: string,
    kubernetesContext?: string,
    trustProvider?: proto_trust_provider_v1alpha1_trust_provider_pb.TrustProvider.AsObject,
    extraHelmValues?: google_protobuf_struct_pb.Struct.AsObject,
    profile?: string,
    externalServer?: boolean,
  }

  export enum IdCase { 
    _ID_NOT_SET = 0,
    ID = 8,
  }

  export enum NameCase { 
    _NAME_NOT_SET = 0,
    NAME = 1,
  }

  export enum OrgIdCase { 
    _ORG_ID_NOT_SET = 0,
    ORG_ID = 9,
  }

  export enum TrustZoneCase { 
    _TRUST_ZONE_NOT_SET = 0,
    TRUST_ZONE = 2,
  }

  export enum TrustZoneIdCase { 
    _TRUST_ZONE_ID_NOT_SET = 0,
    TRUST_ZONE_ID = 10,
  }

  export enum KubernetesContextCase { 
    _KUBERNETES_CONTEXT_NOT_SET = 0,
    KUBERNETES_CONTEXT = 3,
  }

  export enum TrustProviderCase { 
    _TRUST_PROVIDER_NOT_SET = 0,
    TRUST_PROVIDER = 4,
  }

  export enum ExtraHelmValuesCase { 
    _EXTRA_HELM_VALUES_NOT_SET = 0,
    EXTRA_HELM_VALUES = 5,
  }

  export enum ProfileCase { 
    _PROFILE_NOT_SET = 0,
    PROFILE = 6,
  }

  export enum ExternalServerCase { 
    _EXTERNAL_SERVER_NOT_SET = 0,
    EXTERNAL_SERVER = 7,
  }
}

