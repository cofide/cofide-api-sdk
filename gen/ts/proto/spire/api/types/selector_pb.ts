// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file proto/spire/api/types/selector.proto (package spire.api.types, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file proto/spire/api/types/selector.proto.
 */
export const file_proto_spire_api_types_selector: GenFile = /*@__PURE__*/
  fileDesc("CiRwcm90by9zcGlyZS9hcGkvdHlwZXMvc2VsZWN0b3IucHJvdG8SD3NwaXJlLmFwaS50eXBlcyInCghTZWxlY3RvchIMCgR0eXBlGAEgASgJEg0KBXZhbHVlGAIgASgJIoICCg1TZWxlY3Rvck1hdGNoEiwKCXNlbGVjdG9ycxgBIAMoCzIZLnNwaXJlLmFwaS50eXBlcy5TZWxlY3RvchI7CgVtYXRjaBgCIAEoDjIsLnNwaXJlLmFwaS50eXBlcy5TZWxlY3Rvck1hdGNoLk1hdGNoQmVoYXZpb3IihQEKDU1hdGNoQmVoYXZpb3ISJAogTUFUQ0hfQkVIQVZJT1JfRVhBQ1RfVU5TUEVDSUZJRUQQABIZChVNQVRDSF9CRUhBVklPUl9TVUJTRVQQARIbChdNQVRDSF9CRUhBVklPUl9TVVBFUlNFVBACEhYKEk1BVENIX0JFSEFWSU9SX0FOWRADQjdaNWdpdGh1Yi5jb20vc3BpZmZlL3NwaXJlLWFwaS1zZGsvcHJvdG8vc3BpcmUvYXBpL3R5cGVzYgZwcm90bzM");

/**
 * @generated from message spire.api.types.Selector
 */
export type Selector = Message<"spire.api.types.Selector"> & {
  /**
   * The type of the selector. This is typically the name of the plugin that
   * produces the selector.
   *
   * @generated from field: string type = 1;
   */
  type: string;

  /**
   * The value of the selector.
   *
   * @generated from field: string value = 2;
   */
  value: string;
};

/**
 * Describes the message spire.api.types.Selector.
 * Use `create(SelectorSchema)` to create a new message.
 */
export const SelectorSchema: GenMessage<Selector> = /*@__PURE__*/
  messageDesc(file_proto_spire_api_types_selector, 0);

/**
 * @generated from message spire.api.types.SelectorMatch
 */
export type SelectorMatch = Message<"spire.api.types.SelectorMatch"> & {
  /**
   * The set of selectors to match on.
   *
   * @generated from field: repeated spire.api.types.Selector selectors = 1;
   */
  selectors: Selector[];

  /**
   * How to match the selectors.
   *
   * @generated from field: spire.api.types.SelectorMatch.MatchBehavior match = 2;
   */
  match: SelectorMatch_MatchBehavior;
};

/**
 * Describes the message spire.api.types.SelectorMatch.
 * Use `create(SelectorMatchSchema)` to create a new message.
 */
export const SelectorMatchSchema: GenMessage<SelectorMatch> = /*@__PURE__*/
  messageDesc(file_proto_spire_api_types_selector, 1);

/**
 * @generated from enum spire.api.types.SelectorMatch.MatchBehavior
 */
export enum SelectorMatch_MatchBehavior {
  /**
   * Indicates that the selectors in this match are equal to the
   * candidate selectors, independent of ordering.
   * Example:
   *   Given:
   *     - 'e1 { Selectors: ["a:1", "b:2", "c:3"]}'
   *     - 'e2 { Selectors: ["a:1", "b:2"]}'
   *     - 'e3 { Selectors: ["a:1"]}'
   *   Operation:
   *     - MATCH_EXACT ["a:1", "b:2"]
   *   Entries that match:
   *     - 'e2'
   *
   * @generated from enum value: MATCH_BEHAVIOR_EXACT_UNSPECIFIED = 0;
   */
  EXACT_UNSPECIFIED = 0,

  /**
   * Indicates that all candidates which have a non-empty subset
   * of the provided set of selectors will match.
   * Example:
   *   Given:
   *     - 'e1 { Selectors: ["a:1", "b:2", "c:3"]}'
   *     - 'e2 { Selectors: ["a:1", "b:2"]}'
   *     - 'e3 { Selectors: ["a:1"]}'
   *   Operation:
   *     - MATCH_SUBSET ["a:1"]
   *   Entries that match:
   *     - 'e1'
   *
   * @generated from enum value: MATCH_BEHAVIOR_SUBSET = 1;
   */
  SUBSET = 1,

  /**
   * Indicates that all candidates which are a superset
   * of the provided selectors will match.
   * Example:
   *   Given:
   *     - 'e1 { Selectors: ["a:1", "b:2", "c:3"]}'
   *     - 'e2 { Selectors: ["a:1", "b:2"]}'
   *     - 'e3 { Selectors: ["a:1"]}'
   *   Operation:
   *     - MATCH_SUPERSET ["a:1", "b:2"]
   *   Entries that match:
   *     - 'e1'
   *     - 'e2'
   *
   * @generated from enum value: MATCH_BEHAVIOR_SUPERSET = 2;
   */
  SUPERSET = 2,

  /**
   * Indicates that all candidates which have at least one
   * of the provided set of selectors will match.
   * Example:
   *   Given:
   *     - 'e1 { Selectors: ["a:1", "b:2", "c:3"]}'
   *     - 'e2 { Selectors: ["a:1", "b:2"]}'
   *     - 'e3 { Selectors: ["a:1"]}'
   *   Operation:
   *     - MATCH_ANY ["a:1"]
   *   Entries that match:
   *     - 'e1'
   *     - 'e2'
   *     - 'e3'
   *
   * @generated from enum value: MATCH_BEHAVIOR_ANY = 3;
   */
  ANY = 3,
}

/**
 * Describes the enum spire.api.types.SelectorMatch.MatchBehavior.
 */
export const SelectorMatch_MatchBehaviorSchema: GenEnum<SelectorMatch_MatchBehavior> = /*@__PURE__*/
  enumDesc(file_proto_spire_api_types_selector, 1, 0);

