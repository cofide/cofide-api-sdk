import * as jspb from 'google-protobuf'



export class Federation extends jspb.Message {
  getId(): string;
  setId(value: string): Federation;
  hasId(): boolean;
  clearId(): Federation;

  getOrgId(): string;
  setOrgId(value: string): Federation;
  hasOrgId(): boolean;
  clearOrgId(): Federation;

  getFrom(): string;
  setFrom(value: string): Federation;

  getTo(): string;
  setTo(value: string): Federation;

  getTrustZoneId(): string;
  setTrustZoneId(value: string): Federation;
  hasTrustZoneId(): boolean;
  clearTrustZoneId(): Federation;

  getRemoteTrustZoneId(): string;
  setRemoteTrustZoneId(value: string): Federation;
  hasRemoteTrustZoneId(): boolean;
  clearRemoteTrustZoneId(): Federation;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Federation.AsObject;
  static toObject(includeInstance: boolean, msg: Federation): Federation.AsObject;
  static serializeBinaryToWriter(message: Federation, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Federation;
  static deserializeBinaryFromReader(message: Federation, reader: jspb.BinaryReader): Federation;
}

export namespace Federation {
  export type AsObject = {
    id?: string,
    orgId?: string,
    from: string,
    to: string,
    trustZoneId?: string,
    remoteTrustZoneId?: string,
  }

  export enum IdCase { 
    _ID_NOT_SET = 0,
    ID = 6,
  }

  export enum OrgIdCase { 
    _ORG_ID_NOT_SET = 0,
    ORG_ID = 5,
  }

  export enum TrustZoneIdCase { 
    _TRUST_ZONE_ID_NOT_SET = 0,
    TRUST_ZONE_ID = 3,
  }

  export enum RemoteTrustZoneIdCase { 
    _REMOTE_TRUST_ZONE_ID_NOT_SET = 0,
    REMOTE_TRUST_ZONE_ID = 4,
  }
}

