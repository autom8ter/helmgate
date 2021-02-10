// GENERATED CODE -- DO NOT EDIT!

// package: helmgate
// file: schema.proto

import * as schema_pb from "./schema_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "grpc";

interface IHelmProxyServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  getApp: grpc.MethodDefinition<schema_pb.AppRef, schema_pb.App>;
  getHistory: grpc.MethodDefinition<schema_pb.HistoryFilter, schema_pb.Apps>;
  searchApps: grpc.MethodDefinition<schema_pb.AppFilter, schema_pb.Apps>;
  uninstallApp: grpc.MethodDefinition<schema_pb.AppRef, google_protobuf_empty_pb.Empty>;
  rollbackApp: grpc.MethodDefinition<schema_pb.AppRef, schema_pb.App>;
  installApp: grpc.MethodDefinition<schema_pb.AppInput, schema_pb.App>;
  updateApp: grpc.MethodDefinition<schema_pb.AppInput, schema_pb.App>;
  searchCharts: grpc.MethodDefinition<schema_pb.ChartFilter, schema_pb.Charts>;
}

export const HelmProxyServiceService: IHelmProxyServiceService;

export class HelmProxyServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  getApp(argument: schema_pb.AppRef, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  getApp(argument: schema_pb.AppRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  getApp(argument: schema_pb.AppRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  getHistory(argument: schema_pb.HistoryFilter, callback: grpc.requestCallback<schema_pb.Apps>): grpc.ClientUnaryCall;
  getHistory(argument: schema_pb.HistoryFilter, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Apps>): grpc.ClientUnaryCall;
  getHistory(argument: schema_pb.HistoryFilter, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Apps>): grpc.ClientUnaryCall;
  searchApps(argument: schema_pb.AppFilter, callback: grpc.requestCallback<schema_pb.Apps>): grpc.ClientUnaryCall;
  searchApps(argument: schema_pb.AppFilter, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Apps>): grpc.ClientUnaryCall;
  searchApps(argument: schema_pb.AppFilter, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Apps>): grpc.ClientUnaryCall;
  uninstallApp(argument: schema_pb.AppRef, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  uninstallApp(argument: schema_pb.AppRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  uninstallApp(argument: schema_pb.AppRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  rollbackApp(argument: schema_pb.AppRef, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  rollbackApp(argument: schema_pb.AppRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  rollbackApp(argument: schema_pb.AppRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  installApp(argument: schema_pb.AppInput, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  installApp(argument: schema_pb.AppInput, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  installApp(argument: schema_pb.AppInput, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  updateApp(argument: schema_pb.AppInput, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  updateApp(argument: schema_pb.AppInput, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  updateApp(argument: schema_pb.AppInput, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.App>): grpc.ClientUnaryCall;
  searchCharts(argument: schema_pb.ChartFilter, callback: grpc.requestCallback<schema_pb.Charts>): grpc.ClientUnaryCall;
  searchCharts(argument: schema_pb.ChartFilter, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Charts>): grpc.ClientUnaryCall;
  searchCharts(argument: schema_pb.ChartFilter, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Charts>): grpc.ClientUnaryCall;
}
