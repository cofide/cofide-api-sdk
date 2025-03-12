import * as jspb from 'google-protobuf'



export class APBinding extends jspb.Message {
  getId(): string;
  setId(value: string): APBinding;
  hasId(): boolean;
  clearId(): APBinding;

  getOrgId(): string;
  setOrgId(value: string): APBinding;
  hasOrgId(): boolean;
  clearOrgId(): APBinding;

  getTrustZone(): string;
  setTrustZone(value: string): APBinding;

  getTrustZoneId(): string;
  setTrustZoneId(value: string): APBinding;
  hasTrustZoneId(): boolean;
  clearTrustZoneId(): APBinding;

  getPolicy(): string;
  setPolicy(value: string): APBinding;

  getPolicyId(): string;
  setPolicyId(value: string): APBinding;
  hasPolicyId(): boolean;
  clearPolicyId(): APBinding;

  getFederatesWithList(): Array<string>;
  setFederatesWithList(value: Array<string>): APBinding;
  clearFederatesWithList(): APBinding;
  addFederatesWith(value: string, index?: number): APBinding;

  getFederationsList(): Array<APBindingFederation>;
  setFederationsList(value: Array<APBindingFederation>): APBinding;
  clearFederationsList(): APBinding;
  addFederations(value?: APBindingFederation, index?: number): APBindingFederation;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): APBinding.AsObject;
  static toObject(includeInstance: boolean, msg: APBinding): APBinding.AsObject;
  static serializeBinaryToWriter(message: APBinding, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): APBinding;
  static deserializeBinaryFromReader(message: APBinding, reader: jspb.BinaryReader): APBinding;
}

export namespace APBinding {
  export type AsObject = {
    id?: string,
    orgId?: string,
    trustZone: string,
    trustZoneId?: string,
    policy: string,
    policyId?: string,
    federatesWithList: Array<string>,
    federationsList: Array<APBindingFederation.AsObject>,
  }

  export enum IdCase { 
    _ID_NOT_SET = 0,
    ID = 4,
  }

  export enum OrgIdCase { 
    _ORG_ID_NOT_SET = 0,
    ORG_ID = 5,
  }

  export enum TrustZoneIdCase { 
    _TRUST_ZONE_ID_NOT_SET = 0,
    TRUST_ZONE_ID = 6,
  }

  export enum PolicyIdCase { 
    _POLICY_ID_NOT_SET = 0,
    POLICY_ID = 7,
  }
}

export class APBindingFederation extends jspb.Message {
  getTrustZoneId(): string;
  setTrustZoneId(value: string): APBindingFederation;
  hasTrustZoneId(): boolean;
  clearTrustZoneId(): APBindingFederation;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): APBindingFederation.AsObject;
  static toObject(includeInstance: boolean, msg: APBindingFederation): APBindingFederation.AsObject;
  static serializeBinaryToWriter(message: APBindingFederation, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): APBindingFederation;
  static deserializeBinaryFromReader(message: APBindingFederation, reader: jspb.BinaryReader): APBindingFederation;
}

export namespace APBindingFederation {
  export type AsObject = {
    trustZoneId?: string,
  }

  export enum TrustZoneIdCase { 
    _TRUST_ZONE_ID_NOT_SET = 0,
    TRUST_ZONE_ID = 1,
  }
}

