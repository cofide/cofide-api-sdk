import * as jspb from 'google-protobuf'



export class FederatedService extends jspb.Message {
  getId(): string;
  setId(value: string): FederatedService;

  getName(): string;
  setName(value: string): FederatedService;

  getNamespace(): string;
  setNamespace(value: string): FederatedService;

  getClusterName(): string;
  setClusterName(value: string): FederatedService;

  getTrustDomain(): string;
  setTrustDomain(value: string): FederatedService;

  getGatewayIp(): string;
  setGatewayIp(value: string): FederatedService;

  getWorkloadLabelsMap(): jspb.Map<string, string>;
  clearWorkloadLabelsMap(): FederatedService;

  getExportedTrustDomainsList(): Array<string>;
  setExportedTrustDomainsList(value: Array<string>): FederatedService;
  clearExportedTrustDomainsList(): FederatedService;
  addExportedTrustDomains(value: string, index?: number): FederatedService;

  getPort(): number;
  setPort(value: number): FederatedService;

  getGatewayEntriesList(): Array<GatewayEntry>;
  setGatewayEntriesList(value: Array<GatewayEntry>): FederatedService;
  clearGatewayEntriesList(): FederatedService;
  addGatewayEntries(value?: GatewayEntry, index?: number): GatewayEntry;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FederatedService.AsObject;
  static toObject(includeInstance: boolean, msg: FederatedService): FederatedService.AsObject;
  static serializeBinaryToWriter(message: FederatedService, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FederatedService;
  static deserializeBinaryFromReader(message: FederatedService, reader: jspb.BinaryReader): FederatedService;
}

export namespace FederatedService {
  export type AsObject = {
    id: string,
    name: string,
    namespace: string,
    clusterName: string,
    trustDomain: string,
    gatewayIp: string,
    workloadLabelsMap: Array<[string, string]>,
    exportedTrustDomainsList: Array<string>,
    port: number,
    gatewayEntriesList: Array<GatewayEntry.AsObject>,
  }
}

export class GatewayEntry extends jspb.Message {
  getHostname(): string;
  setHostname(value: string): GatewayEntry;

  getType(): string;
  setType(value: string): GatewayEntry;

  getIp(): string;
  setIp(value: string): GatewayEntry;

  getPort(): number;
  setPort(value: number): GatewayEntry;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GatewayEntry.AsObject;
  static toObject(includeInstance: boolean, msg: GatewayEntry): GatewayEntry.AsObject;
  static serializeBinaryToWriter(message: GatewayEntry, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GatewayEntry;
  static deserializeBinaryFromReader(message: GatewayEntry, reader: jspb.BinaryReader): GatewayEntry;
}

export namespace GatewayEntry {
  export type AsObject = {
    hostname: string,
    type: string,
    ip: string,
    port: number,
  }
}

