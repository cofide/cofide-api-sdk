import * as jspb from 'google-protobuf'



export class Plugins extends jspb.Message {
  getDataSource(): string;
  setDataSource(value: string): Plugins;
  hasDataSource(): boolean;
  clearDataSource(): Plugins;

  getProvision(): string;
  setProvision(value: string): Plugins;
  hasProvision(): boolean;
  clearProvision(): Plugins;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Plugins.AsObject;
  static toObject(includeInstance: boolean, msg: Plugins): Plugins.AsObject;
  static serializeBinaryToWriter(message: Plugins, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Plugins;
  static deserializeBinaryFromReader(message: Plugins, reader: jspb.BinaryReader): Plugins;
}

export namespace Plugins {
  export type AsObject = {
    dataSource?: string,
    provision?: string,
  }

  export enum DataSourceCase { 
    _DATA_SOURCE_NOT_SET = 0,
    DATA_SOURCE = 1,
  }

  export enum ProvisionCase { 
    _PROVISION_NOT_SET = 0,
    PROVISION = 2,
  }
}

