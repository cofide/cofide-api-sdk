import * as jspb from 'google-protobuf'

import * as proto_cluster_v1alpha1_cluster_pb from '../../../../proto/cluster/v1alpha1/cluster_pb'; // proto import: "proto/cluster/v1alpha1/cluster.proto"
import * as proto_spire_api_types_bundle_pb from '../../../../proto/spire/api/types/bundle_pb'; // proto import: "proto/spire/api/types/bundle.proto"
import * as proto_trust_zone_v1alpha1_trust_zone_pb from '../../../../proto/trust_zone/v1alpha1/trust_zone_pb'; // proto import: "proto/trust_zone/v1alpha1/trust_zone.proto"


export class Agent extends jspb.Message {
  getAgentId(): string;
  setAgentId(value: string): Agent;

  getClusterId(): string;
  setClusterId(value: string): Agent;

  getTrustZoneId(): string;
  setTrustZoneId(value: string): Agent;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Agent.AsObject;
  static toObject(includeInstance: boolean, msg: Agent): Agent.AsObject;
  static serializeBinaryToWriter(message: Agent, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Agent;
  static deserializeBinaryFromReader(message: Agent, reader: jspb.BinaryReader): Agent;
}

export namespace Agent {
  export type AsObject = {
    agentId: string,
    clusterId: string,
    trustZoneId: string,
  }
}

export class CreateTrustZoneRequest extends jspb.Message {
  getTrustZone(): proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone | undefined;
  setTrustZone(value?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone): CreateTrustZoneRequest;
  hasTrustZone(): boolean;
  clearTrustZone(): CreateTrustZoneRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTrustZoneRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTrustZoneRequest): CreateTrustZoneRequest.AsObject;
  static serializeBinaryToWriter(message: CreateTrustZoneRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTrustZoneRequest;
  static deserializeBinaryFromReader(message: CreateTrustZoneRequest, reader: jspb.BinaryReader): CreateTrustZoneRequest;
}

export namespace CreateTrustZoneRequest {
  export type AsObject = {
    trustZone?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone.AsObject,
  }
}

export class CreateTrustZoneResponse extends jspb.Message {
  getTrustZone(): proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone | undefined;
  setTrustZone(value?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone): CreateTrustZoneResponse;
  hasTrustZone(): boolean;
  clearTrustZone(): CreateTrustZoneResponse;

  getSuccess(): boolean;
  setSuccess(value: boolean): CreateTrustZoneResponse;

  getMessage(): string;
  setMessage(value: string): CreateTrustZoneResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTrustZoneResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTrustZoneResponse): CreateTrustZoneResponse.AsObject;
  static serializeBinaryToWriter(message: CreateTrustZoneResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTrustZoneResponse;
  static deserializeBinaryFromReader(message: CreateTrustZoneResponse, reader: jspb.BinaryReader): CreateTrustZoneResponse;
}

export namespace CreateTrustZoneResponse {
  export type AsObject = {
    trustZone?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone.AsObject,
    success: boolean,
    message: string,
  }
}

export class ListTrustZonesRequest extends jspb.Message {
  getFilter(): ListTrustZonesRequest.Filter | undefined;
  setFilter(value?: ListTrustZonesRequest.Filter): ListTrustZonesRequest;
  hasFilter(): boolean;
  clearFilter(): ListTrustZonesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTrustZonesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListTrustZonesRequest): ListTrustZonesRequest.AsObject;
  static serializeBinaryToWriter(message: ListTrustZonesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTrustZonesRequest;
  static deserializeBinaryFromReader(message: ListTrustZonesRequest, reader: jspb.BinaryReader): ListTrustZonesRequest;
}

export namespace ListTrustZonesRequest {
  export type AsObject = {
    filter?: ListTrustZonesRequest.Filter.AsObject,
  }

  export class Filter extends jspb.Message {
    getName(): string;
    setName(value: string): Filter;
    hasName(): boolean;
    clearName(): Filter;

    getOrgId(): string;
    setOrgId(value: string): Filter;
    hasOrgId(): boolean;
    clearOrgId(): Filter;

    getTrustDomain(): string;
    setTrustDomain(value: string): Filter;
    hasTrustDomain(): boolean;
    clearTrustDomain(): Filter;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Filter.AsObject;
    static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
    static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Filter;
    static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
  }

  export namespace Filter {
    export type AsObject = {
      name?: string,
      orgId?: string,
      trustDomain?: string,
    }

    export enum NameCase { 
      _NAME_NOT_SET = 0,
      NAME = 1,
    }

    export enum OrgIdCase { 
      _ORG_ID_NOT_SET = 0,
      ORG_ID = 2,
    }

    export enum TrustDomainCase { 
      _TRUST_DOMAIN_NOT_SET = 0,
      TRUST_DOMAIN = 3,
    }
  }


  export enum FilterCase { 
    _FILTER_NOT_SET = 0,
    FILTER = 1,
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

export class GetTrustZoneDetailsRequest extends jspb.Message {
  getTrustZoneId(): string;
  setTrustZoneId(value: string): GetTrustZoneDetailsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTrustZoneDetailsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTrustZoneDetailsRequest): GetTrustZoneDetailsRequest.AsObject;
  static serializeBinaryToWriter(message: GetTrustZoneDetailsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTrustZoneDetailsRequest;
  static deserializeBinaryFromReader(message: GetTrustZoneDetailsRequest, reader: jspb.BinaryReader): GetTrustZoneDetailsRequest;
}

export namespace GetTrustZoneDetailsRequest {
  export type AsObject = {
    trustZoneId: string,
  }
}

export class GetTrustZoneDetailsResponse extends jspb.Message {
  getTrustZone(): proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone | undefined;
  setTrustZone(value?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone): GetTrustZoneDetailsResponse;
  hasTrustZone(): boolean;
  clearTrustZone(): GetTrustZoneDetailsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTrustZoneDetailsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTrustZoneDetailsResponse): GetTrustZoneDetailsResponse.AsObject;
  static serializeBinaryToWriter(message: GetTrustZoneDetailsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTrustZoneDetailsResponse;
  static deserializeBinaryFromReader(message: GetTrustZoneDetailsResponse, reader: jspb.BinaryReader): GetTrustZoneDetailsResponse;
}

export namespace GetTrustZoneDetailsResponse {
  export type AsObject = {
    trustZone?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone.AsObject,
  }
}

export class RegisterClusterRequest extends jspb.Message {
  getTrustZoneId(): string;
  setTrustZoneId(value: string): RegisterClusterRequest;

  getCluster(): proto_cluster_v1alpha1_cluster_pb.Cluster | undefined;
  setCluster(value?: proto_cluster_v1alpha1_cluster_pb.Cluster): RegisterClusterRequest;
  hasCluster(): boolean;
  clearCluster(): RegisterClusterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterClusterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterClusterRequest): RegisterClusterRequest.AsObject;
  static serializeBinaryToWriter(message: RegisterClusterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterClusterRequest;
  static deserializeBinaryFromReader(message: RegisterClusterRequest, reader: jspb.BinaryReader): RegisterClusterRequest;
}

export namespace RegisterClusterRequest {
  export type AsObject = {
    trustZoneId: string,
    cluster?: proto_cluster_v1alpha1_cluster_pb.Cluster.AsObject,
  }
}

export class RegisterClusterResponse extends jspb.Message {
  getAgentToken(): string;
  setAgentToken(value: string): RegisterClusterResponse;

  getSuccess(): boolean;
  setSuccess(value: boolean): RegisterClusterResponse;

  getMessage(): string;
  setMessage(value: string): RegisterClusterResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterClusterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterClusterResponse): RegisterClusterResponse.AsObject;
  static serializeBinaryToWriter(message: RegisterClusterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterClusterResponse;
  static deserializeBinaryFromReader(message: RegisterClusterResponse, reader: jspb.BinaryReader): RegisterClusterResponse;
}

export namespace RegisterClusterResponse {
  export type AsObject = {
    agentToken: string,
    success: boolean,
    message: string,
  }
}

export class RegisterAgentRequest extends jspb.Message {
  getAgent(): Agent | undefined;
  setAgent(value?: Agent): RegisterAgentRequest;
  hasAgent(): boolean;
  clearAgent(): RegisterAgentRequest;

  getAgentToken(): string;
  setAgentToken(value: string): RegisterAgentRequest;

  getBundle(): proto_spire_api_types_bundle_pb.Bundle | undefined;
  setBundle(value?: proto_spire_api_types_bundle_pb.Bundle): RegisterAgentRequest;
  hasBundle(): boolean;
  clearBundle(): RegisterAgentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterAgentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterAgentRequest): RegisterAgentRequest.AsObject;
  static serializeBinaryToWriter(message: RegisterAgentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterAgentRequest;
  static deserializeBinaryFromReader(message: RegisterAgentRequest, reader: jspb.BinaryReader): RegisterAgentRequest;
}

export namespace RegisterAgentRequest {
  export type AsObject = {
    agent?: Agent.AsObject,
    agentToken: string,
    bundle?: proto_spire_api_types_bundle_pb.Bundle.AsObject,
  }
}

export class RegisterAgentResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): RegisterAgentResponse;

  getMessage(): string;
  setMessage(value: string): RegisterAgentResponse;

  getAgentId(): string;
  setAgentId(value: string): RegisterAgentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterAgentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterAgentResponse): RegisterAgentResponse.AsObject;
  static serializeBinaryToWriter(message: RegisterAgentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterAgentResponse;
  static deserializeBinaryFromReader(message: RegisterAgentResponse, reader: jspb.BinaryReader): RegisterAgentResponse;
}

export namespace RegisterAgentResponse {
  export type AsObject = {
    success: boolean,
    message: string,
    agentId: string,
  }
}

