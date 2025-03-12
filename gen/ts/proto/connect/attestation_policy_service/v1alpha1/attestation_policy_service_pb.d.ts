import * as jspb from 'google-protobuf'

import * as proto_attestation_policy_v1alpha1_attestation_policy_pb from '../../../../proto/attestation_policy/v1alpha1/attestation_policy_pb'; // proto import: "proto/attestation_policy/v1alpha1/attestation_policy.proto"


export class CreateAttestationPolicyRequest extends jspb.Message {
  getPolicy(): proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy | undefined;
  setPolicy(value?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy): CreateAttestationPolicyRequest;
  hasPolicy(): boolean;
  clearPolicy(): CreateAttestationPolicyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAttestationPolicyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAttestationPolicyRequest): CreateAttestationPolicyRequest.AsObject;
  static serializeBinaryToWriter(message: CreateAttestationPolicyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAttestationPolicyRequest;
  static deserializeBinaryFromReader(message: CreateAttestationPolicyRequest, reader: jspb.BinaryReader): CreateAttestationPolicyRequest;
}

export namespace CreateAttestationPolicyRequest {
  export type AsObject = {
    policy?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy.AsObject,
  }

  export enum PolicyCase { 
    _POLICY_NOT_SET = 0,
    POLICY = 1,
  }
}

export class CreateAttestationPolicyResponse extends jspb.Message {
  getPolicy(): proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy | undefined;
  setPolicy(value?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy): CreateAttestationPolicyResponse;
  hasPolicy(): boolean;
  clearPolicy(): CreateAttestationPolicyResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAttestationPolicyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAttestationPolicyResponse): CreateAttestationPolicyResponse.AsObject;
  static serializeBinaryToWriter(message: CreateAttestationPolicyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAttestationPolicyResponse;
  static deserializeBinaryFromReader(message: CreateAttestationPolicyResponse, reader: jspb.BinaryReader): CreateAttestationPolicyResponse;
}

export namespace CreateAttestationPolicyResponse {
  export type AsObject = {
    policy?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy.AsObject,
  }

  export enum PolicyCase { 
    _POLICY_NOT_SET = 0,
    POLICY = 1,
  }
}

export class DestroyAttestationPolicyRequest extends jspb.Message {
  getPolicyId(): string;
  setPolicyId(value: string): DestroyAttestationPolicyRequest;
  hasPolicyId(): boolean;
  clearPolicyId(): DestroyAttestationPolicyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestroyAttestationPolicyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DestroyAttestationPolicyRequest): DestroyAttestationPolicyRequest.AsObject;
  static serializeBinaryToWriter(message: DestroyAttestationPolicyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestroyAttestationPolicyRequest;
  static deserializeBinaryFromReader(message: DestroyAttestationPolicyRequest, reader: jspb.BinaryReader): DestroyAttestationPolicyRequest;
}

export namespace DestroyAttestationPolicyRequest {
  export type AsObject = {
    policyId?: string,
  }

  export enum PolicyIdCase { 
    _POLICY_ID_NOT_SET = 0,
    POLICY_ID = 1,
  }
}

export class DestroyAttestationPolicyResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestroyAttestationPolicyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DestroyAttestationPolicyResponse): DestroyAttestationPolicyResponse.AsObject;
  static serializeBinaryToWriter(message: DestroyAttestationPolicyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestroyAttestationPolicyResponse;
  static deserializeBinaryFromReader(message: DestroyAttestationPolicyResponse, reader: jspb.BinaryReader): DestroyAttestationPolicyResponse;
}

export namespace DestroyAttestationPolicyResponse {
  export type AsObject = {
  }
}

export class GetAttestationPolicyRequest extends jspb.Message {
  getPolicyId(): string;
  setPolicyId(value: string): GetAttestationPolicyRequest;
  hasPolicyId(): boolean;
  clearPolicyId(): GetAttestationPolicyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAttestationPolicyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAttestationPolicyRequest): GetAttestationPolicyRequest.AsObject;
  static serializeBinaryToWriter(message: GetAttestationPolicyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAttestationPolicyRequest;
  static deserializeBinaryFromReader(message: GetAttestationPolicyRequest, reader: jspb.BinaryReader): GetAttestationPolicyRequest;
}

export namespace GetAttestationPolicyRequest {
  export type AsObject = {
    policyId?: string,
  }

  export enum PolicyIdCase { 
    _POLICY_ID_NOT_SET = 0,
    POLICY_ID = 1,
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

  export enum PolicyCase { 
    _POLICY_NOT_SET = 0,
    POLICY = 1,
  }
}

export class ListAttestationPoliciesRequest extends jspb.Message {
  getFilter(): ListAttestationPoliciesRequest.Filter | undefined;
  setFilter(value?: ListAttestationPoliciesRequest.Filter): ListAttestationPoliciesRequest;
  hasFilter(): boolean;
  clearFilter(): ListAttestationPoliciesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAttestationPoliciesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListAttestationPoliciesRequest): ListAttestationPoliciesRequest.AsObject;
  static serializeBinaryToWriter(message: ListAttestationPoliciesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAttestationPoliciesRequest;
  static deserializeBinaryFromReader(message: ListAttestationPoliciesRequest, reader: jspb.BinaryReader): ListAttestationPoliciesRequest;
}

export namespace ListAttestationPoliciesRequest {
  export type AsObject = {
    filter?: ListAttestationPoliciesRequest.Filter.AsObject,
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
    }

    export enum NameCase { 
      _NAME_NOT_SET = 0,
      NAME = 1,
    }

    export enum OrgIdCase { 
      _ORG_ID_NOT_SET = 0,
      ORG_ID = 2,
    }
  }


  export enum FilterCase { 
    _FILTER_NOT_SET = 0,
    FILTER = 1,
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

export class UpdateAttestationPolicyRequest extends jspb.Message {
  getPolicy(): proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy | undefined;
  setPolicy(value?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy): UpdateAttestationPolicyRequest;
  hasPolicy(): boolean;
  clearPolicy(): UpdateAttestationPolicyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAttestationPolicyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAttestationPolicyRequest): UpdateAttestationPolicyRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateAttestationPolicyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAttestationPolicyRequest;
  static deserializeBinaryFromReader(message: UpdateAttestationPolicyRequest, reader: jspb.BinaryReader): UpdateAttestationPolicyRequest;
}

export namespace UpdateAttestationPolicyRequest {
  export type AsObject = {
    policy?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy.AsObject,
  }

  export enum PolicyCase { 
    _POLICY_NOT_SET = 0,
    POLICY = 1,
  }
}

export class UpdateAttestationPolicyResponse extends jspb.Message {
  getPolicy(): proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy | undefined;
  setPolicy(value?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy): UpdateAttestationPolicyResponse;
  hasPolicy(): boolean;
  clearPolicy(): UpdateAttestationPolicyResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAttestationPolicyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAttestationPolicyResponse): UpdateAttestationPolicyResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateAttestationPolicyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAttestationPolicyResponse;
  static deserializeBinaryFromReader(message: UpdateAttestationPolicyResponse, reader: jspb.BinaryReader): UpdateAttestationPolicyResponse;
}

export namespace UpdateAttestationPolicyResponse {
  export type AsObject = {
    policy?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy.AsObject,
  }

  export enum PolicyCase { 
    _POLICY_NOT_SET = 0,
    POLICY = 1,
  }
}

