import * as jspb from 'google-protobuf'

import * as proto_federation_v1alpha1_federation_pb from '../../../../proto/federation/v1alpha1/federation_pb'; // proto import: "proto/federation/v1alpha1/federation.proto"


export class CreateFederationRequest extends jspb.Message {
  getFederation(): proto_federation_v1alpha1_federation_pb.Federation | undefined;
  setFederation(value?: proto_federation_v1alpha1_federation_pb.Federation): CreateFederationRequest;
  hasFederation(): boolean;
  clearFederation(): CreateFederationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateFederationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateFederationRequest): CreateFederationRequest.AsObject;
  static serializeBinaryToWriter(message: CreateFederationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateFederationRequest;
  static deserializeBinaryFromReader(message: CreateFederationRequest, reader: jspb.BinaryReader): CreateFederationRequest;
}

export namespace CreateFederationRequest {
  export type AsObject = {
    federation?: proto_federation_v1alpha1_federation_pb.Federation.AsObject,
  }
}

export class CreateFederationResponse extends jspb.Message {
  getFederation(): proto_federation_v1alpha1_federation_pb.Federation | undefined;
  setFederation(value?: proto_federation_v1alpha1_federation_pb.Federation): CreateFederationResponse;
  hasFederation(): boolean;
  clearFederation(): CreateFederationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateFederationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateFederationResponse): CreateFederationResponse.AsObject;
  static serializeBinaryToWriter(message: CreateFederationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateFederationResponse;
  static deserializeBinaryFromReader(message: CreateFederationResponse, reader: jspb.BinaryReader): CreateFederationResponse;
}

export namespace CreateFederationResponse {
  export type AsObject = {
    federation?: proto_federation_v1alpha1_federation_pb.Federation.AsObject,
  }
}

export class DestroyFederationRequest extends jspb.Message {
  getFederationId(): string;
  setFederationId(value: string): DestroyFederationRequest;
  hasFederationId(): boolean;
  clearFederationId(): DestroyFederationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestroyFederationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DestroyFederationRequest): DestroyFederationRequest.AsObject;
  static serializeBinaryToWriter(message: DestroyFederationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestroyFederationRequest;
  static deserializeBinaryFromReader(message: DestroyFederationRequest, reader: jspb.BinaryReader): DestroyFederationRequest;
}

export namespace DestroyFederationRequest {
  export type AsObject = {
    federationId?: string,
  }

  export enum FederationIdCase { 
    _FEDERATION_ID_NOT_SET = 0,
    FEDERATION_ID = 1,
  }
}

export class DestroyFederationResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestroyFederationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DestroyFederationResponse): DestroyFederationResponse.AsObject;
  static serializeBinaryToWriter(message: DestroyFederationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestroyFederationResponse;
  static deserializeBinaryFromReader(message: DestroyFederationResponse, reader: jspb.BinaryReader): DestroyFederationResponse;
}

export namespace DestroyFederationResponse {
  export type AsObject = {
  }
}

export class ListFederationsRequest extends jspb.Message {
  getFilter(): ListFederationsRequest.Filter | undefined;
  setFilter(value?: ListFederationsRequest.Filter): ListFederationsRequest;
  hasFilter(): boolean;
  clearFilter(): ListFederationsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFederationsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListFederationsRequest): ListFederationsRequest.AsObject;
  static serializeBinaryToWriter(message: ListFederationsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFederationsRequest;
  static deserializeBinaryFromReader(message: ListFederationsRequest, reader: jspb.BinaryReader): ListFederationsRequest;
}

export namespace ListFederationsRequest {
  export type AsObject = {
    filter?: ListFederationsRequest.Filter.AsObject,
  }

  export class Filter extends jspb.Message {
    getOrgId(): string;
    setOrgId(value: string): Filter;
    hasOrgId(): boolean;
    clearOrgId(): Filter;

    getTrustZoneId(): string;
    setTrustZoneId(value: string): Filter;
    hasTrustZoneId(): boolean;
    clearTrustZoneId(): Filter;

    getRemoteTrustZoneId(): string;
    setRemoteTrustZoneId(value: string): Filter;
    hasRemoteTrustZoneId(): boolean;
    clearRemoteTrustZoneId(): Filter;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Filter.AsObject;
    static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
    static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Filter;
    static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
  }

  export namespace Filter {
    export type AsObject = {
      orgId?: string,
      trustZoneId?: string,
      remoteTrustZoneId?: string,
    }

    export enum OrgIdCase { 
      _ORG_ID_NOT_SET = 0,
      ORG_ID = 1,
    }

    export enum TrustZoneIdCase { 
      _TRUST_ZONE_ID_NOT_SET = 0,
      TRUST_ZONE_ID = 2,
    }

    export enum RemoteTrustZoneIdCase { 
      _REMOTE_TRUST_ZONE_ID_NOT_SET = 0,
      REMOTE_TRUST_ZONE_ID = 3,
    }
  }


  export enum FilterCase { 
    _FILTER_NOT_SET = 0,
    FILTER = 1,
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

