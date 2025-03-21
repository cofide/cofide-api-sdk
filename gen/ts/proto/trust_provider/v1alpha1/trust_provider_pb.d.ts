import * as jspb from 'google-protobuf'



export class TrustProvider extends jspb.Message {
  getKind(): string;
  setKind(value: string): TrustProvider;
  hasKind(): boolean;
  clearKind(): TrustProvider;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TrustProvider.AsObject;
  static toObject(includeInstance: boolean, msg: TrustProvider): TrustProvider.AsObject;
  static serializeBinaryToWriter(message: TrustProvider, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TrustProvider;
  static deserializeBinaryFromReader(message: TrustProvider, reader: jspb.BinaryReader): TrustProvider;
}

export namespace TrustProvider {
  export type AsObject = {
    kind?: string,
  }

  export enum KindCase { 
    _KIND_NOT_SET = 0,
    KIND = 1,
  }
}

export enum TrustProviderKind { 
  TRUST_PROVIDER_KIND_UNSPECIFIED = 0,
  TRUST_PROVIDER_KIND_KUBERNETES = 1,
}
