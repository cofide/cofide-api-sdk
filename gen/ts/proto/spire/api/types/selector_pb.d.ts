import * as jspb from 'google-protobuf'



export class Selector extends jspb.Message {
  getType(): string;
  setType(value: string): Selector;

  getValue(): string;
  setValue(value: string): Selector;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Selector.AsObject;
  static toObject(includeInstance: boolean, msg: Selector): Selector.AsObject;
  static serializeBinaryToWriter(message: Selector, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Selector;
  static deserializeBinaryFromReader(message: Selector, reader: jspb.BinaryReader): Selector;
}

export namespace Selector {
  export type AsObject = {
    type: string,
    value: string,
  }
}

export class SelectorMatch extends jspb.Message {
  getSelectorsList(): Array<Selector>;
  setSelectorsList(value: Array<Selector>): SelectorMatch;
  clearSelectorsList(): SelectorMatch;
  addSelectors(value?: Selector, index?: number): Selector;

  getMatch(): SelectorMatch.MatchBehavior;
  setMatch(value: SelectorMatch.MatchBehavior): SelectorMatch;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SelectorMatch.AsObject;
  static toObject(includeInstance: boolean, msg: SelectorMatch): SelectorMatch.AsObject;
  static serializeBinaryToWriter(message: SelectorMatch, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SelectorMatch;
  static deserializeBinaryFromReader(message: SelectorMatch, reader: jspb.BinaryReader): SelectorMatch;
}

export namespace SelectorMatch {
  export type AsObject = {
    selectorsList: Array<Selector.AsObject>,
    match: SelectorMatch.MatchBehavior,
  }

  export enum MatchBehavior { 
    MATCH_BEHAVIOR_EXACT_UNSPECIFIED = 0,
    MATCH_BEHAVIOR_SUBSET = 1,
    MATCH_BEHAVIOR_SUPERSET = 2,
    MATCH_BEHAVIOR_ANY = 3,
  }
}

