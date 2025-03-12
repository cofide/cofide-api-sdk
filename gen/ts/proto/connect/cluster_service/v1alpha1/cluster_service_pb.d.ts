import * as jspb from 'google-protobuf'

import * as proto_cluster_v1alpha1_cluster_pb from '../../../../proto/cluster/v1alpha1/cluster_pb'; // proto import: "proto/cluster/v1alpha1/cluster.proto"


export class CreateClusterRequest extends jspb.Message {
  getCluster(): proto_cluster_v1alpha1_cluster_pb.Cluster | undefined;
  setCluster(value?: proto_cluster_v1alpha1_cluster_pb.Cluster): CreateClusterRequest;
  hasCluster(): boolean;
  clearCluster(): CreateClusterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateClusterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateClusterRequest): CreateClusterRequest.AsObject;
  static serializeBinaryToWriter(message: CreateClusterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateClusterRequest;
  static deserializeBinaryFromReader(message: CreateClusterRequest, reader: jspb.BinaryReader): CreateClusterRequest;
}

export namespace CreateClusterRequest {
  export type AsObject = {
    cluster?: proto_cluster_v1alpha1_cluster_pb.Cluster.AsObject,
  }

  export enum ClusterCase { 
    _CLUSTER_NOT_SET = 0,
    CLUSTER = 1,
  }
}

export class CreateClusterResponse extends jspb.Message {
  getCluster(): proto_cluster_v1alpha1_cluster_pb.Cluster | undefined;
  setCluster(value?: proto_cluster_v1alpha1_cluster_pb.Cluster): CreateClusterResponse;
  hasCluster(): boolean;
  clearCluster(): CreateClusterResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateClusterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateClusterResponse): CreateClusterResponse.AsObject;
  static serializeBinaryToWriter(message: CreateClusterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateClusterResponse;
  static deserializeBinaryFromReader(message: CreateClusterResponse, reader: jspb.BinaryReader): CreateClusterResponse;
}

export namespace CreateClusterResponse {
  export type AsObject = {
    cluster?: proto_cluster_v1alpha1_cluster_pb.Cluster.AsObject,
  }

  export enum ClusterCase { 
    _CLUSTER_NOT_SET = 0,
    CLUSTER = 1,
  }
}

export class DestroyClusterRequest extends jspb.Message {
  getClusterId(): string;
  setClusterId(value: string): DestroyClusterRequest;
  hasClusterId(): boolean;
  clearClusterId(): DestroyClusterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestroyClusterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DestroyClusterRequest): DestroyClusterRequest.AsObject;
  static serializeBinaryToWriter(message: DestroyClusterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestroyClusterRequest;
  static deserializeBinaryFromReader(message: DestroyClusterRequest, reader: jspb.BinaryReader): DestroyClusterRequest;
}

export namespace DestroyClusterRequest {
  export type AsObject = {
    clusterId?: string,
  }

  export enum ClusterIdCase { 
    _CLUSTER_ID_NOT_SET = 0,
    CLUSTER_ID = 1,
  }
}

export class DestroyClusterResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DestroyClusterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DestroyClusterResponse): DestroyClusterResponse.AsObject;
  static serializeBinaryToWriter(message: DestroyClusterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DestroyClusterResponse;
  static deserializeBinaryFromReader(message: DestroyClusterResponse, reader: jspb.BinaryReader): DestroyClusterResponse;
}

export namespace DestroyClusterResponse {
  export type AsObject = {
  }
}

export class GetClusterRequest extends jspb.Message {
  getClusterId(): string;
  setClusterId(value: string): GetClusterRequest;
  hasClusterId(): boolean;
  clearClusterId(): GetClusterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetClusterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetClusterRequest): GetClusterRequest.AsObject;
  static serializeBinaryToWriter(message: GetClusterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetClusterRequest;
  static deserializeBinaryFromReader(message: GetClusterRequest, reader: jspb.BinaryReader): GetClusterRequest;
}

export namespace GetClusterRequest {
  export type AsObject = {
    clusterId?: string,
  }

  export enum ClusterIdCase { 
    _CLUSTER_ID_NOT_SET = 0,
    CLUSTER_ID = 1,
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
  getFilter(): ListClustersRequest.Filter | undefined;
  setFilter(value?: ListClustersRequest.Filter): ListClustersRequest;
  hasFilter(): boolean;
  clearFilter(): ListClustersRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListClustersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListClustersRequest): ListClustersRequest.AsObject;
  static serializeBinaryToWriter(message: ListClustersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListClustersRequest;
  static deserializeBinaryFromReader(message: ListClustersRequest, reader: jspb.BinaryReader): ListClustersRequest;
}

export namespace ListClustersRequest {
  export type AsObject = {
    filter?: ListClustersRequest.Filter.AsObject,
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

    getTrustZoneId(): string;
    setTrustZoneId(value: string): Filter;
    hasTrustZoneId(): boolean;
    clearTrustZoneId(): Filter;

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
      trustZoneId?: string,
    }

    export enum NameCase { 
      _NAME_NOT_SET = 0,
      NAME = 1,
    }

    export enum OrgIdCase { 
      _ORG_ID_NOT_SET = 0,
      ORG_ID = 2,
    }

    export enum TrustZoneIdCase { 
      _TRUST_ZONE_ID_NOT_SET = 0,
      TRUST_ZONE_ID = 3,
    }
  }


  export enum FilterCase { 
    _FILTER_NOT_SET = 0,
    FILTER = 1,
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

