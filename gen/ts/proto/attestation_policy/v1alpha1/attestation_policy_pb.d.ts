import * as jspb from 'google-protobuf'

import * as proto_spire_api_types_selector_pb from '../../../proto/spire/api/types/selector_pb'; // proto import: "proto/spire/api/types/selector.proto"


export class AttestationPolicy extends jspb.Message {
  getId(): string;
  setId(value: string): AttestationPolicy;
  hasId(): boolean;
  clearId(): AttestationPolicy;

  getName(): string;
  setName(value: string): AttestationPolicy;

  getOrgId(): string;
  setOrgId(value: string): AttestationPolicy;
  hasOrgId(): boolean;
  clearOrgId(): AttestationPolicy;

  getKubernetes(): APKubernetes | undefined;
  setKubernetes(value?: APKubernetes): AttestationPolicy;
  hasKubernetes(): boolean;
  clearKubernetes(): AttestationPolicy;

  getStatic(): APStatic | undefined;
  setStatic(value?: APStatic): AttestationPolicy;
  hasStatic(): boolean;
  clearStatic(): AttestationPolicy;

  getPolicyCase(): AttestationPolicy.PolicyCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AttestationPolicy.AsObject;
  static toObject(includeInstance: boolean, msg: AttestationPolicy): AttestationPolicy.AsObject;
  static serializeBinaryToWriter(message: AttestationPolicy, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AttestationPolicy;
  static deserializeBinaryFromReader(message: AttestationPolicy, reader: jspb.BinaryReader): AttestationPolicy;
}

export namespace AttestationPolicy {
  export type AsObject = {
    id?: string,
    name: string,
    orgId?: string,
    kubernetes?: APKubernetes.AsObject,
    pb_static?: APStatic.AsObject,
  }

  export enum PolicyCase { 
    POLICY_NOT_SET = 0,
    KUBERNETES = 2,
    STATIC = 3,
  }

  export enum IdCase { 
    _ID_NOT_SET = 0,
    ID = 4,
  }

  export enum OrgIdCase { 
    _ORG_ID_NOT_SET = 0,
    ORG_ID = 5,
  }
}

export class APKubernetes extends jspb.Message {
  getNamespaceSelector(): APLabelSelector | undefined;
  setNamespaceSelector(value?: APLabelSelector): APKubernetes;
  hasNamespaceSelector(): boolean;
  clearNamespaceSelector(): APKubernetes;

  getPodSelector(): APLabelSelector | undefined;
  setPodSelector(value?: APLabelSelector): APKubernetes;
  hasPodSelector(): boolean;
  clearPodSelector(): APKubernetes;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): APKubernetes.AsObject;
  static toObject(includeInstance: boolean, msg: APKubernetes): APKubernetes.AsObject;
  static serializeBinaryToWriter(message: APKubernetes, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): APKubernetes;
  static deserializeBinaryFromReader(message: APKubernetes, reader: jspb.BinaryReader): APKubernetes;
}

export namespace APKubernetes {
  export type AsObject = {
    namespaceSelector?: APLabelSelector.AsObject,
    podSelector?: APLabelSelector.AsObject,
  }

  export enum NamespaceSelectorCase { 
    _NAMESPACE_SELECTOR_NOT_SET = 0,
    NAMESPACE_SELECTOR = 1,
  }

  export enum PodSelectorCase { 
    _POD_SELECTOR_NOT_SET = 0,
    POD_SELECTOR = 2,
  }
}

export class APLabelSelector extends jspb.Message {
  getMatchLabelsMap(): jspb.Map<string, string>;
  clearMatchLabelsMap(): APLabelSelector;

  getMatchExpressionsList(): Array<APMatchExpression>;
  setMatchExpressionsList(value: Array<APMatchExpression>): APLabelSelector;
  clearMatchExpressionsList(): APLabelSelector;
  addMatchExpressions(value?: APMatchExpression, index?: number): APMatchExpression;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): APLabelSelector.AsObject;
  static toObject(includeInstance: boolean, msg: APLabelSelector): APLabelSelector.AsObject;
  static serializeBinaryToWriter(message: APLabelSelector, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): APLabelSelector;
  static deserializeBinaryFromReader(message: APLabelSelector, reader: jspb.BinaryReader): APLabelSelector;
}

export namespace APLabelSelector {
  export type AsObject = {
    matchLabelsMap: Array<[string, string]>,
    matchExpressionsList: Array<APMatchExpression.AsObject>,
  }
}

export class APMatchExpression extends jspb.Message {
  getKey(): string;
  setKey(value: string): APMatchExpression;

  getOperator(): string;
  setOperator(value: string): APMatchExpression;

  getValuesList(): Array<string>;
  setValuesList(value: Array<string>): APMatchExpression;
  clearValuesList(): APMatchExpression;
  addValues(value: string, index?: number): APMatchExpression;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): APMatchExpression.AsObject;
  static toObject(includeInstance: boolean, msg: APMatchExpression): APMatchExpression.AsObject;
  static serializeBinaryToWriter(message: APMatchExpression, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): APMatchExpression;
  static deserializeBinaryFromReader(message: APMatchExpression, reader: jspb.BinaryReader): APMatchExpression;
}

export namespace APMatchExpression {
  export type AsObject = {
    key: string,
    operator: string,
    valuesList: Array<string>,
  }
}

export class APStatic extends jspb.Message {
  getSpiffeId(): string;
  setSpiffeId(value: string): APStatic;
  hasSpiffeId(): boolean;
  clearSpiffeId(): APStatic;

  getSelectorsList(): Array<proto_spire_api_types_selector_pb.Selector>;
  setSelectorsList(value: Array<proto_spire_api_types_selector_pb.Selector>): APStatic;
  clearSelectorsList(): APStatic;
  addSelectors(value?: proto_spire_api_types_selector_pb.Selector, index?: number): proto_spire_api_types_selector_pb.Selector;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): APStatic.AsObject;
  static toObject(includeInstance: boolean, msg: APStatic): APStatic.AsObject;
  static serializeBinaryToWriter(message: APStatic, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): APStatic;
  static deserializeBinaryFromReader(message: APStatic, reader: jspb.BinaryReader): APStatic;
}

export namespace APStatic {
  export type AsObject = {
    spiffeId?: string,
    selectorsList: Array<proto_spire_api_types_selector_pb.Selector.AsObject>,
  }

  export enum SpiffeIdCase { 
    _SPIFFE_ID_NOT_SET = 0,
    SPIFFE_ID = 1,
  }
}

