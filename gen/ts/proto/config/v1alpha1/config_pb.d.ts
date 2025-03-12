import * as jspb from 'google-protobuf'

import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb'; // proto import: "google/protobuf/struct.proto"
import * as proto_attestation_policy_v1alpha1_attestation_policy_pb from '../../../proto/attestation_policy/v1alpha1/attestation_policy_pb'; // proto import: "proto/attestation_policy/v1alpha1/attestation_policy.proto"
import * as proto_cluster_v1alpha1_cluster_pb from '../../../proto/cluster/v1alpha1/cluster_pb'; // proto import: "proto/cluster/v1alpha1/cluster.proto"
import * as proto_plugins_v1alpha1_plugins_pb from '../../../proto/plugins/v1alpha1/plugins_pb'; // proto import: "proto/plugins/v1alpha1/plugins.proto"
import * as proto_trust_zone_v1alpha1_trust_zone_pb from '../../../proto/trust_zone/v1alpha1/trust_zone_pb'; // proto import: "proto/trust_zone/v1alpha1/trust_zone.proto"


export class Config extends jspb.Message {
  getTrustZonesList(): Array<proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone>;
  setTrustZonesList(value: Array<proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone>): Config;
  clearTrustZonesList(): Config;
  addTrustZones(value?: proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone, index?: number): proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone;

  getClustersList(): Array<proto_cluster_v1alpha1_cluster_pb.Cluster>;
  setClustersList(value: Array<proto_cluster_v1alpha1_cluster_pb.Cluster>): Config;
  clearClustersList(): Config;
  addClusters(value?: proto_cluster_v1alpha1_cluster_pb.Cluster, index?: number): proto_cluster_v1alpha1_cluster_pb.Cluster;

  getAttestationPoliciesList(): Array<proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy>;
  setAttestationPoliciesList(value: Array<proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy>): Config;
  clearAttestationPoliciesList(): Config;
  addAttestationPolicies(value?: proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy, index?: number): proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy;

  getPluginConfigMap(): jspb.Map<string, google_protobuf_struct_pb.Struct>;
  clearPluginConfigMap(): Config;

  getPlugins(): proto_plugins_v1alpha1_plugins_pb.Plugins | undefined;
  setPlugins(value?: proto_plugins_v1alpha1_plugins_pb.Plugins): Config;
  hasPlugins(): boolean;
  clearPlugins(): Config;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Config.AsObject;
  static toObject(includeInstance: boolean, msg: Config): Config.AsObject;
  static serializeBinaryToWriter(message: Config, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Config;
  static deserializeBinaryFromReader(message: Config, reader: jspb.BinaryReader): Config;
}

export namespace Config {
  export type AsObject = {
    trustZonesList: Array<proto_trust_zone_v1alpha1_trust_zone_pb.TrustZone.AsObject>,
    clustersList: Array<proto_cluster_v1alpha1_cluster_pb.Cluster.AsObject>,
    attestationPoliciesList: Array<proto_attestation_policy_v1alpha1_attestation_policy_pb.AttestationPolicy.AsObject>,
    pluginConfigMap: Array<[string, google_protobuf_struct_pb.Struct.AsObject]>,
    plugins?: proto_plugins_v1alpha1_plugins_pb.Plugins.AsObject,
  }

  export enum PluginsCase { 
    _PLUGINS_NOT_SET = 0,
    PLUGINS = 5,
  }
}

