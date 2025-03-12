import * as jspb from 'google-protobuf'



export class Agent extends jspb.Message {
  getId(): string;
  setId(value: string): Agent;
  hasId(): boolean;
  clearId(): Agent;

  getClusterId(): string;
  setClusterId(value: string): Agent;
  hasClusterId(): boolean;
  clearClusterId(): Agent;

  getTrustZoneId(): string;
  setTrustZoneId(value: string): Agent;
  hasTrustZoneId(): boolean;
  clearTrustZoneId(): Agent;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Agent.AsObject;
  static toObject(includeInstance: boolean, msg: Agent): Agent.AsObject;
  static serializeBinaryToWriter(message: Agent, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Agent;
  static deserializeBinaryFromReader(message: Agent, reader: jspb.BinaryReader): Agent;
}

export namespace Agent {
  export type AsObject = {
    id?: string,
    clusterId?: string,
    trustZoneId?: string,
  }

  export enum IdCase { 
    _ID_NOT_SET = 0,
    ID = 1,
  }

  export enum ClusterIdCase { 
    _CLUSTER_ID_NOT_SET = 0,
    CLUSTER_ID = 2,
  }

  export enum TrustZoneIdCase { 
    _TRUST_ZONE_ID_NOT_SET = 0,
    TRUST_ZONE_ID = 3,
  }
}

export class AgentStatus extends jspb.Message {
  getStatus(): AgentStatusCode;
  setStatus(value: AgentStatusCode): AgentStatus;
  hasStatus(): boolean;
  clearStatus(): AgentStatus;

  getStatusMessage(): string;
  setStatusMessage(value: string): AgentStatus;
  hasStatusMessage(): boolean;
  clearStatusMessage(): AgentStatus;

  getLastUpdated(): number;
  setLastUpdated(value: number): AgentStatus;
  hasLastUpdated(): boolean;
  clearLastUpdated(): AgentStatus;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AgentStatus.AsObject;
  static toObject(includeInstance: boolean, msg: AgentStatus): AgentStatus.AsObject;
  static serializeBinaryToWriter(message: AgentStatus, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AgentStatus;
  static deserializeBinaryFromReader(message: AgentStatus, reader: jspb.BinaryReader): AgentStatus;
}

export namespace AgentStatus {
  export type AsObject = {
    status?: AgentStatusCode,
    statusMessage?: string,
    lastUpdated?: number,
  }

  export enum StatusCase { 
    _STATUS_NOT_SET = 0,
    STATUS = 1,
  }

  export enum StatusMessageCase { 
    _STATUS_MESSAGE_NOT_SET = 0,
    STATUS_MESSAGE = 2,
  }

  export enum LastUpdatedCase { 
    _LAST_UPDATED_NOT_SET = 0,
    LAST_UPDATED = 3,
  }
}

export enum AgentStatusCode { 
  AGENT_STATUS_CODE_UNSPECIFIED = 0,
  AGENT_STATUS_CODE_RUNNING = 1,
  AGENT_STATUS_CODE_STOPPED = 2,
  AGENT_STATUS_CODE_ERROR = 3,
}
