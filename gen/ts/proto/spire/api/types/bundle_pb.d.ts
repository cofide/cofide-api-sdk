import * as jspb from 'google-protobuf'



export class Bundle extends jspb.Message {
  getTrustDomain(): string;
  setTrustDomain(value: string): Bundle;

  getX509AuthoritiesList(): Array<X509Certificate>;
  setX509AuthoritiesList(value: Array<X509Certificate>): Bundle;
  clearX509AuthoritiesList(): Bundle;
  addX509Authorities(value?: X509Certificate, index?: number): X509Certificate;

  getJwtAuthoritiesList(): Array<JWTKey>;
  setJwtAuthoritiesList(value: Array<JWTKey>): Bundle;
  clearJwtAuthoritiesList(): Bundle;
  addJwtAuthorities(value?: JWTKey, index?: number): JWTKey;

  getRefreshHint(): number;
  setRefreshHint(value: number): Bundle;

  getSequenceNumber(): number;
  setSequenceNumber(value: number): Bundle;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Bundle.AsObject;
  static toObject(includeInstance: boolean, msg: Bundle): Bundle.AsObject;
  static serializeBinaryToWriter(message: Bundle, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Bundle;
  static deserializeBinaryFromReader(message: Bundle, reader: jspb.BinaryReader): Bundle;
}

export namespace Bundle {
  export type AsObject = {
    trustDomain: string,
    x509AuthoritiesList: Array<X509Certificate.AsObject>,
    jwtAuthoritiesList: Array<JWTKey.AsObject>,
    refreshHint: number,
    sequenceNumber: number,
  }
}

export class X509Certificate extends jspb.Message {
  getAsn1(): Uint8Array | string;
  getAsn1_asU8(): Uint8Array;
  getAsn1_asB64(): string;
  setAsn1(value: Uint8Array | string): X509Certificate;

  getTainted(): boolean;
  setTainted(value: boolean): X509Certificate;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): X509Certificate.AsObject;
  static toObject(includeInstance: boolean, msg: X509Certificate): X509Certificate.AsObject;
  static serializeBinaryToWriter(message: X509Certificate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): X509Certificate;
  static deserializeBinaryFromReader(message: X509Certificate, reader: jspb.BinaryReader): X509Certificate;
}

export namespace X509Certificate {
  export type AsObject = {
    asn1: Uint8Array | string,
    tainted: boolean,
  }
}

export class JWTKey extends jspb.Message {
  getPublicKey(): Uint8Array | string;
  getPublicKey_asU8(): Uint8Array;
  getPublicKey_asB64(): string;
  setPublicKey(value: Uint8Array | string): JWTKey;

  getKeyId(): string;
  setKeyId(value: string): JWTKey;

  getExpiresAt(): number;
  setExpiresAt(value: number): JWTKey;

  getTainted(): boolean;
  setTainted(value: boolean): JWTKey;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JWTKey.AsObject;
  static toObject(includeInstance: boolean, msg: JWTKey): JWTKey.AsObject;
  static serializeBinaryToWriter(message: JWTKey, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JWTKey;
  static deserializeBinaryFromReader(message: JWTKey, reader: jspb.BinaryReader): JWTKey;
}

export namespace JWTKey {
  export type AsObject = {
    publicKey: Uint8Array | string,
    keyId: string,
    expiresAt: number,
    tainted: boolean,
  }
}

export class BundleMask extends jspb.Message {
  getX509Authorities(): boolean;
  setX509Authorities(value: boolean): BundleMask;

  getJwtAuthorities(): boolean;
  setJwtAuthorities(value: boolean): BundleMask;

  getRefreshHint(): boolean;
  setRefreshHint(value: boolean): BundleMask;

  getSequenceNumber(): boolean;
  setSequenceNumber(value: boolean): BundleMask;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BundleMask.AsObject;
  static toObject(includeInstance: boolean, msg: BundleMask): BundleMask.AsObject;
  static serializeBinaryToWriter(message: BundleMask, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BundleMask;
  static deserializeBinaryFromReader(message: BundleMask, reader: jspb.BinaryReader): BundleMask;
}

export namespace BundleMask {
  export type AsObject = {
    x509Authorities: boolean,
    jwtAuthorities: boolean,
    refreshHint: boolean,
    sequenceNumber: boolean,
  }
}

