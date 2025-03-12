import * as jspb from 'google-protobuf'

import * as proto_agent_v1alpha1_agent_pb from '../../../../proto/agent/v1alpha1/agent_pb'; // proto import: "proto/agent/v1alpha1/agent.proto"
import * as proto_federated_service_v1alpha1_federated_service_pb from '../../../../proto/federated_service/v1alpha1/federated_service_pb'; // proto import: "proto/federated_service/v1alpha1/federated_service.proto"
import * as proto_spire_api_types_bundle_pb from '../../../../proto/spire/api/types/bundle_pb'; // proto import: "proto/spire/api/types/bundle.proto"


export class CreateAgentJoinTokenRequest extends jspb.Message {
  getTrustZoneId(): string;
  setTrustZoneId(value: string): CreateAgentJoinTokenRequest;
  hasTrustZoneId(): boolean;
  clearTrustZoneId(): CreateAgentJoinTokenRequest;

  getClusterId(): string;
  setClusterId(value: string): CreateAgentJoinTokenRequest;
  hasClusterId(): boolean;
  clearClusterId(): CreateAgentJoinTokenRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAgentJoinTokenRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAgentJoinTokenRequest): CreateAgentJoinTokenRequest.AsObject;
  static serializeBinaryToWriter(message: CreateAgentJoinTokenRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAgentJoinTokenRequest;
  static deserializeBinaryFromReader(message: CreateAgentJoinTokenRequest, reader: jspb.BinaryReader): CreateAgentJoinTokenRequest;
}

export namespace CreateAgentJoinTokenRequest {
  export type AsObject = {
    trustZoneId?: string,
    clusterId?: string,
  }

  export enum TrustZoneIdCase { 
    _TRUST_ZONE_ID_NOT_SET = 0,
    TRUST_ZONE_ID = 1,
  }

  export enum ClusterIdCase { 
    _CLUSTER_ID_NOT_SET = 0,
    CLUSTER_ID = 2,
  }
}

export class CreateAgentJoinTokenResponse extends jspb.Message {
  getAgentToken(): string;
  setAgentToken(value: string): CreateAgentJoinTokenResponse;
  hasAgentToken(): boolean;
  clearAgentToken(): CreateAgentJoinTokenResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAgentJoinTokenResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAgentJoinTokenResponse): CreateAgentJoinTokenResponse.AsObject;
  static serializeBinaryToWriter(message: CreateAgentJoinTokenResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAgentJoinTokenResponse;
  static deserializeBinaryFromReader(message: CreateAgentJoinTokenResponse, reader: jspb.BinaryReader): CreateAgentJoinTokenResponse;
}

export namespace CreateAgentJoinTokenResponse {
  export type AsObject = {
    agentToken?: string,
  }

  export enum AgentTokenCase { 
    _AGENT_TOKEN_NOT_SET = 0,
    AGENT_TOKEN = 1,
  }
}

export class UpdateTrustZoneBundleRequest extends jspb.Message {
  getBundle(): proto_spire_api_types_bundle_pb.Bundle | undefined;
  setBundle(value?: proto_spire_api_types_bundle_pb.Bundle): UpdateTrustZoneBundleRequest;
  hasBundle(): boolean;
  clearBundle(): UpdateTrustZoneBundleRequest;

  getTrustZoneId(): string;
  setTrustZoneId(value: string): UpdateTrustZoneBundleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateTrustZoneBundleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateTrustZoneBundleRequest): UpdateTrustZoneBundleRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateTrustZoneBundleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateTrustZoneBundleRequest;
  static deserializeBinaryFromReader(message: UpdateTrustZoneBundleRequest, reader: jspb.BinaryReader): UpdateTrustZoneBundleRequest;
}

export namespace UpdateTrustZoneBundleRequest {
  export type AsObject = {
    bundle?: proto_spire_api_types_bundle_pb.Bundle.AsObject,
    trustZoneId: string,
  }
}

export class UpdateTrustZoneBundleResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): UpdateTrustZoneBundleResponse;

  getMessage(): string;
  setMessage(value: string): UpdateTrustZoneBundleResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateTrustZoneBundleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateTrustZoneBundleResponse): UpdateTrustZoneBundleResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateTrustZoneBundleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateTrustZoneBundleResponse;
  static deserializeBinaryFromReader(message: UpdateTrustZoneBundleResponse, reader: jspb.BinaryReader): UpdateTrustZoneBundleResponse;
}

export namespace UpdateTrustZoneBundleResponse {
  export type AsObject = {
    success: boolean,
    message: string,
  }
}

export class UpdateAgentStatusRequest extends jspb.Message {
  getStatus(): proto_agent_v1alpha1_agent_pb.AgentStatus | undefined;
  setStatus(value?: proto_agent_v1alpha1_agent_pb.AgentStatus): UpdateAgentStatusRequest;
  hasStatus(): boolean;
  clearStatus(): UpdateAgentStatusRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAgentStatusRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAgentStatusRequest): UpdateAgentStatusRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateAgentStatusRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAgentStatusRequest;
  static deserializeBinaryFromReader(message: UpdateAgentStatusRequest, reader: jspb.BinaryReader): UpdateAgentStatusRequest;
}

export namespace UpdateAgentStatusRequest {
  export type AsObject = {
    status?: proto_agent_v1alpha1_agent_pb.AgentStatus.AsObject,
  }
}

export class UpdateAgentStatusResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): UpdateAgentStatusResponse;

  getMessage(): string;
  setMessage(value: string): UpdateAgentStatusResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAgentStatusResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAgentStatusResponse): UpdateAgentStatusResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateAgentStatusResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAgentStatusResponse;
  static deserializeBinaryFromReader(message: UpdateAgentStatusResponse, reader: jspb.BinaryReader): UpdateAgentStatusResponse;
}

export namespace UpdateAgentStatusResponse {
  export type AsObject = {
    success: boolean,
    message: string,
  }
}

export class RegisterFederatedServiceRequest extends jspb.Message {
  getService(): proto_federated_service_v1alpha1_federated_service_pb.FederatedService | undefined;
  setService(value?: proto_federated_service_v1alpha1_federated_service_pb.FederatedService): RegisterFederatedServiceRequest;
  hasService(): boolean;
  clearService(): RegisterFederatedServiceRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterFederatedServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterFederatedServiceRequest): RegisterFederatedServiceRequest.AsObject;
  static serializeBinaryToWriter(message: RegisterFederatedServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterFederatedServiceRequest;
  static deserializeBinaryFromReader(message: RegisterFederatedServiceRequest, reader: jspb.BinaryReader): RegisterFederatedServiceRequest;
}

export namespace RegisterFederatedServiceRequest {
  export type AsObject = {
    service?: proto_federated_service_v1alpha1_federated_service_pb.FederatedService.AsObject,
  }
}

export class RegisterFederatedServiceResponse extends jspb.Message {
  getServiceId(): string;
  setServiceId(value: string): RegisterFederatedServiceResponse;

  getSuccess(): boolean;
  setSuccess(value: boolean): RegisterFederatedServiceResponse;

  getMessage(): string;
  setMessage(value: string): RegisterFederatedServiceResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterFederatedServiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterFederatedServiceResponse): RegisterFederatedServiceResponse.AsObject;
  static serializeBinaryToWriter(message: RegisterFederatedServiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterFederatedServiceResponse;
  static deserializeBinaryFromReader(message: RegisterFederatedServiceResponse, reader: jspb.BinaryReader): RegisterFederatedServiceResponse;
}

export namespace RegisterFederatedServiceResponse {
  export type AsObject = {
    serviceId: string,
    success: boolean,
    message: string,
  }
}

export class DeregisterFederatedServiceRequest extends jspb.Message {
  getServiceId(): string;
  setServiceId(value: string): DeregisterFederatedServiceRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeregisterFederatedServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeregisterFederatedServiceRequest): DeregisterFederatedServiceRequest.AsObject;
  static serializeBinaryToWriter(message: DeregisterFederatedServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeregisterFederatedServiceRequest;
  static deserializeBinaryFromReader(message: DeregisterFederatedServiceRequest, reader: jspb.BinaryReader): DeregisterFederatedServiceRequest;
}

export namespace DeregisterFederatedServiceRequest {
  export type AsObject = {
    serviceId: string,
  }
}

export class DeregisterFederatedServiceResponse extends jspb.Message {
  getServiceId(): string;
  setServiceId(value: string): DeregisterFederatedServiceResponse;

  getSuccess(): boolean;
  setSuccess(value: boolean): DeregisterFederatedServiceResponse;

  getMessage(): string;
  setMessage(value: string): DeregisterFederatedServiceResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeregisterFederatedServiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeregisterFederatedServiceResponse): DeregisterFederatedServiceResponse.AsObject;
  static serializeBinaryToWriter(message: DeregisterFederatedServiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeregisterFederatedServiceResponse;
  static deserializeBinaryFromReader(message: DeregisterFederatedServiceResponse, reader: jspb.BinaryReader): DeregisterFederatedServiceResponse;
}

export namespace DeregisterFederatedServiceResponse {
  export type AsObject = {
    serviceId: string,
    success: boolean,
    message: string,
  }
}

export class UpdateFederatedServiceRequest extends jspb.Message {
  getService(): proto_federated_service_v1alpha1_federated_service_pb.FederatedService | undefined;
  setService(value?: proto_federated_service_v1alpha1_federated_service_pb.FederatedService): UpdateFederatedServiceRequest;
  hasService(): boolean;
  clearService(): UpdateFederatedServiceRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateFederatedServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateFederatedServiceRequest): UpdateFederatedServiceRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateFederatedServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateFederatedServiceRequest;
  static deserializeBinaryFromReader(message: UpdateFederatedServiceRequest, reader: jspb.BinaryReader): UpdateFederatedServiceRequest;
}

export namespace UpdateFederatedServiceRequest {
  export type AsObject = {
    service?: proto_federated_service_v1alpha1_federated_service_pb.FederatedService.AsObject,
  }
}

export class UpdateFederatedServiceResponse extends jspb.Message {
  getServiceId(): string;
  setServiceId(value: string): UpdateFederatedServiceResponse;

  getSuccess(): boolean;
  setSuccess(value: boolean): UpdateFederatedServiceResponse;

  getMessage(): string;
  setMessage(value: string): UpdateFederatedServiceResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateFederatedServiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateFederatedServiceResponse): UpdateFederatedServiceResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateFederatedServiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateFederatedServiceResponse;
  static deserializeBinaryFromReader(message: UpdateFederatedServiceResponse, reader: jspb.BinaryReader): UpdateFederatedServiceResponse;
}

export namespace UpdateFederatedServiceResponse {
  export type AsObject = {
    serviceId: string,
    success: boolean,
    message: string,
  }
}

export class GetFederatedServiceRequest extends jspb.Message {
  getServiceId(): string;
  setServiceId(value: string): GetFederatedServiceRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFederatedServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetFederatedServiceRequest): GetFederatedServiceRequest.AsObject;
  static serializeBinaryToWriter(message: GetFederatedServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFederatedServiceRequest;
  static deserializeBinaryFromReader(message: GetFederatedServiceRequest, reader: jspb.BinaryReader): GetFederatedServiceRequest;
}

export namespace GetFederatedServiceRequest {
  export type AsObject = {
    serviceId: string,
  }
}

export class GetFederatedServiceResponse extends jspb.Message {
  getService(): proto_federated_service_v1alpha1_federated_service_pb.FederatedService | undefined;
  setService(value?: proto_federated_service_v1alpha1_federated_service_pb.FederatedService): GetFederatedServiceResponse;
  hasService(): boolean;
  clearService(): GetFederatedServiceResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFederatedServiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetFederatedServiceResponse): GetFederatedServiceResponse.AsObject;
  static serializeBinaryToWriter(message: GetFederatedServiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFederatedServiceResponse;
  static deserializeBinaryFromReader(message: GetFederatedServiceResponse, reader: jspb.BinaryReader): GetFederatedServiceResponse;
}

export namespace GetFederatedServiceResponse {
  export type AsObject = {
    service?: proto_federated_service_v1alpha1_federated_service_pb.FederatedService.AsObject,
  }
}

export class ListFederatedServicesRequest extends jspb.Message {
  getOrgId(): string;
  setOrgId(value: string): ListFederatedServicesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFederatedServicesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListFederatedServicesRequest): ListFederatedServicesRequest.AsObject;
  static serializeBinaryToWriter(message: ListFederatedServicesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFederatedServicesRequest;
  static deserializeBinaryFromReader(message: ListFederatedServicesRequest, reader: jspb.BinaryReader): ListFederatedServicesRequest;
}

export namespace ListFederatedServicesRequest {
  export type AsObject = {
    orgId: string,
  }
}

export class ListFederatedServicesResponse extends jspb.Message {
  getServicesList(): Array<proto_federated_service_v1alpha1_federated_service_pb.FederatedService>;
  setServicesList(value: Array<proto_federated_service_v1alpha1_federated_service_pb.FederatedService>): ListFederatedServicesResponse;
  clearServicesList(): ListFederatedServicesResponse;
  addServices(value?: proto_federated_service_v1alpha1_federated_service_pb.FederatedService, index?: number): proto_federated_service_v1alpha1_federated_service_pb.FederatedService;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFederatedServicesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListFederatedServicesResponse): ListFederatedServicesResponse.AsObject;
  static serializeBinaryToWriter(message: ListFederatedServicesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFederatedServicesResponse;
  static deserializeBinaryFromReader(message: ListFederatedServicesResponse, reader: jspb.BinaryReader): ListFederatedServicesResponse;
}

export namespace ListFederatedServicesResponse {
  export type AsObject = {
    servicesList: Array<proto_federated_service_v1alpha1_federated_service_pb.FederatedService.AsObject>,
  }
}

