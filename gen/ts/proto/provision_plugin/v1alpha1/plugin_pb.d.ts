import * as jspb from 'google-protobuf'



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

export class DeployRequest extends jspb.Message {
  getDataSource(): number;
  setDataSource(value: number): DeployRequest;
  hasDataSource(): boolean;
  clearDataSource(): DeployRequest;

  getKubeCfgFile(): string;
  setKubeCfgFile(value: string): DeployRequest;
  hasKubeCfgFile(): boolean;
  clearKubeCfgFile(): DeployRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeployRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeployRequest): DeployRequest.AsObject;
  static serializeBinaryToWriter(message: DeployRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeployRequest;
  static deserializeBinaryFromReader(message: DeployRequest, reader: jspb.BinaryReader): DeployRequest;
}

export namespace DeployRequest {
  export type AsObject = {
    dataSource?: number,
    kubeCfgFile?: string,
  }

  export enum DataSourceCase { 
    _DATA_SOURCE_NOT_SET = 0,
    DATA_SOURCE = 1,
  }

  export enum KubeCfgFileCase { 
    _KUBE_CFG_FILE_NOT_SET = 0,
    KUBE_CFG_FILE = 2,
  }
}

export class DeployResponse extends jspb.Message {
  getStatus(): Status | undefined;
  setStatus(value?: Status): DeployResponse;
  hasStatus(): boolean;
  clearStatus(): DeployResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeployResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeployResponse): DeployResponse.AsObject;
  static serializeBinaryToWriter(message: DeployResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeployResponse;
  static deserializeBinaryFromReader(message: DeployResponse, reader: jspb.BinaryReader): DeployResponse;
}

export namespace DeployResponse {
  export type AsObject = {
    status?: Status.AsObject,
  }

  export enum StatusCase { 
    _STATUS_NOT_SET = 0,
    STATUS = 1,
  }
}

export class TearDownRequest extends jspb.Message {
  getDataSource(): number;
  setDataSource(value: number): TearDownRequest;
  hasDataSource(): boolean;
  clearDataSource(): TearDownRequest;

  getKubeCfgFile(): string;
  setKubeCfgFile(value: string): TearDownRequest;
  hasKubeCfgFile(): boolean;
  clearKubeCfgFile(): TearDownRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TearDownRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TearDownRequest): TearDownRequest.AsObject;
  static serializeBinaryToWriter(message: TearDownRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TearDownRequest;
  static deserializeBinaryFromReader(message: TearDownRequest, reader: jspb.BinaryReader): TearDownRequest;
}

export namespace TearDownRequest {
  export type AsObject = {
    dataSource?: number,
    kubeCfgFile?: string,
  }

  export enum DataSourceCase { 
    _DATA_SOURCE_NOT_SET = 0,
    DATA_SOURCE = 1,
  }

  export enum KubeCfgFileCase { 
    _KUBE_CFG_FILE_NOT_SET = 0,
    KUBE_CFG_FILE = 2,
  }
}

export class TearDownResponse extends jspb.Message {
  getStatus(): Status | undefined;
  setStatus(value?: Status): TearDownResponse;
  hasStatus(): boolean;
  clearStatus(): TearDownResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TearDownResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TearDownResponse): TearDownResponse.AsObject;
  static serializeBinaryToWriter(message: TearDownResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TearDownResponse;
  static deserializeBinaryFromReader(message: TearDownResponse, reader: jspb.BinaryReader): TearDownResponse;
}

export namespace TearDownResponse {
  export type AsObject = {
    status?: Status.AsObject,
  }

  export enum StatusCase { 
    _STATUS_NOT_SET = 0,
    STATUS = 1,
  }
}

export class Status extends jspb.Message {
  getStage(): string;
  setStage(value: string): Status;
  hasStage(): boolean;
  clearStage(): Status;

  getMessage(): string;
  setMessage(value: string): Status;
  hasMessage(): boolean;
  clearMessage(): Status;

  getDone(): boolean;
  setDone(value: boolean): Status;
  hasDone(): boolean;
  clearDone(): Status;

  getError(): string;
  setError(value: string): Status;
  hasError(): boolean;
  clearError(): Status;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Status.AsObject;
  static toObject(includeInstance: boolean, msg: Status): Status.AsObject;
  static serializeBinaryToWriter(message: Status, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Status;
  static deserializeBinaryFromReader(message: Status, reader: jspb.BinaryReader): Status;
}

export namespace Status {
  export type AsObject = {
    stage?: string,
    message?: string,
    done?: boolean,
    error?: string,
  }

  export enum StageCase { 
    _STAGE_NOT_SET = 0,
    STAGE = 1,
  }

  export enum MessageCase { 
    _MESSAGE_NOT_SET = 0,
    MESSAGE = 2,
  }

  export enum DoneCase { 
    _DONE_NOT_SET = 0,
    DONE = 3,
  }

  export enum ErrorCase { 
    _ERROR_NOT_SET = 0,
    ERROR = 4,
  }
}

