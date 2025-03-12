import * as jspb from 'google-protobuf'

import * as proto_ap_binding_v1alpha1_ap_binding_pb from '../../../../proto/ap_binding/v1alpha1/ap_binding_pb'; // proto import: "proto/ap_binding/v1alpha1/ap_binding.proto"


export class CreateAPBindingRequest extends jspb.Message {
  getBinding(): proto_ap_binding_v1alpha1_ap_binding_pb.APBinding | undefined;
  setBinding(value?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding): CreateAPBindingRequest;
  hasBinding(): boolean;
  clearBinding(): CreateAPBindingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAPBindingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAPBindingRequest): CreateAPBindingRequest.AsObject;
  static serializeBinaryToWriter(message: CreateAPBindingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAPBindingRequest;
  static deserializeBinaryFromReader(message: CreateAPBindingRequest, reader: jspb.BinaryReader): CreateAPBindingRequest;
}

export namespace CreateAPBindingRequest {
  export type AsObject = {
    binding?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding.AsObject,
  }

  export enum BindingCase { 
    _BINDING_NOT_SET = 0,
    BINDING = 1,
  }
}

export class CreateAPBindingResponse extends jspb.Message {
  getBinding(): proto_ap_binding_v1alpha1_ap_binding_pb.APBinding | undefined;
  setBinding(value?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding): CreateAPBindingResponse;
  hasBinding(): boolean;
  clearBinding(): CreateAPBindingResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAPBindingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAPBindingResponse): CreateAPBindingResponse.AsObject;
  static serializeBinaryToWriter(message: CreateAPBindingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAPBindingResponse;
  static deserializeBinaryFromReader(message: CreateAPBindingResponse, reader: jspb.BinaryReader): CreateAPBindingResponse;
}

export namespace CreateAPBindingResponse {
  export type AsObject = {
    binding?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding.AsObject,
  }

  export enum BindingCase { 
    _BINDING_NOT_SET = 0,
    BINDING = 1,
  }
}

export class DestroyAPBindingRequest extends jspb.Message {
  getBindingId(): string;
  setBindingId(value: string): DestroyAPBindingRequest;
  hasBindingId(): boolean;
  clearBindingId(): DestroyAPBindingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestroyAPBindingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DestroyAPBindingRequest): DestroyAPBindingRequest.AsObject;
  static serializeBinaryToWriter(message: DestroyAPBindingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestroyAPBindingRequest;
  static deserializeBinaryFromReader(message: DestroyAPBindingRequest, reader: jspb.BinaryReader): DestroyAPBindingRequest;
}

export namespace DestroyAPBindingRequest {
  export type AsObject = {
    bindingId?: string,
  }

  export enum BindingIdCase { 
    _BINDING_ID_NOT_SET = 0,
    BINDING_ID = 1,
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

export class GetAPBindingRequest extends jspb.Message {
  getBindingId(): string;
  setBindingId(value: string): GetAPBindingRequest;
  hasBindingId(): boolean;
  clearBindingId(): GetAPBindingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAPBindingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAPBindingRequest): GetAPBindingRequest.AsObject;
  static serializeBinaryToWriter(message: GetAPBindingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAPBindingRequest;
  static deserializeBinaryFromReader(message: GetAPBindingRequest, reader: jspb.BinaryReader): GetAPBindingRequest;
}

export namespace GetAPBindingRequest {
  export type AsObject = {
    bindingId?: string,
  }

  export enum BindingIdCase { 
    _BINDING_ID_NOT_SET = 0,
    BINDING_ID = 1,
  }
}

export class GetAPBindingResponse extends jspb.Message {
  getBinding(): proto_ap_binding_v1alpha1_ap_binding_pb.APBinding | undefined;
  setBinding(value?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding): GetAPBindingResponse;
  hasBinding(): boolean;
  clearBinding(): GetAPBindingResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAPBindingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAPBindingResponse): GetAPBindingResponse.AsObject;
  static serializeBinaryToWriter(message: GetAPBindingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAPBindingResponse;
  static deserializeBinaryFromReader(message: GetAPBindingResponse, reader: jspb.BinaryReader): GetAPBindingResponse;
}

export namespace GetAPBindingResponse {
  export type AsObject = {
    binding?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding.AsObject,
  }

  export enum BindingCase { 
    _BINDING_NOT_SET = 0,
    BINDING = 1,
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
    getOrgId(): string;
    setOrgId(value: string): Filter;
    hasOrgId(): boolean;
    clearOrgId(): Filter;

    getTrustZoneId(): string;
    setTrustZoneId(value: string): Filter;
    hasTrustZoneId(): boolean;
    clearTrustZoneId(): Filter;

    getPolicyId(): string;
    setPolicyId(value: string): Filter;
    hasPolicyId(): boolean;
    clearPolicyId(): Filter;

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
      policyId?: string,
    }

    export enum OrgIdCase { 
      _ORG_ID_NOT_SET = 0,
      ORG_ID = 1,
    }

    export enum TrustZoneIdCase { 
      _TRUST_ZONE_ID_NOT_SET = 0,
      TRUST_ZONE_ID = 2,
    }

    export enum PolicyIdCase { 
      _POLICY_ID_NOT_SET = 0,
      POLICY_ID = 3,
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

export class UpdateAPBindingRequest extends jspb.Message {
  getBinding(): proto_ap_binding_v1alpha1_ap_binding_pb.APBinding | undefined;
  setBinding(value?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding): UpdateAPBindingRequest;
  hasBinding(): boolean;
  clearBinding(): UpdateAPBindingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAPBindingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAPBindingRequest): UpdateAPBindingRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateAPBindingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAPBindingRequest;
  static deserializeBinaryFromReader(message: UpdateAPBindingRequest, reader: jspb.BinaryReader): UpdateAPBindingRequest;
}

export namespace UpdateAPBindingRequest {
  export type AsObject = {
    binding?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding.AsObject,
  }

  export enum BindingCase { 
    _BINDING_NOT_SET = 0,
    BINDING = 1,
  }
}

export class UpdateAPBindingResponse extends jspb.Message {
  getBinding(): proto_ap_binding_v1alpha1_ap_binding_pb.APBinding | undefined;
  setBinding(value?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding): UpdateAPBindingResponse;
  hasBinding(): boolean;
  clearBinding(): UpdateAPBindingResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAPBindingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAPBindingResponse): UpdateAPBindingResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateAPBindingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAPBindingResponse;
  static deserializeBinaryFromReader(message: UpdateAPBindingResponse, reader: jspb.BinaryReader): UpdateAPBindingResponse;
}

export namespace UpdateAPBindingResponse {
  export type AsObject = {
    binding?: proto_ap_binding_v1alpha1_ap_binding_pb.APBinding.AsObject,
  }

  export enum BindingCase { 
    _BINDING_NOT_SET = 0,
    BINDING = 1,
  }
}

