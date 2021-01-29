import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as schema_pb from './schema_pb';


export class HelmProxyServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getApp(
    request: schema_pb.AppRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.App) => void
  ): grpcWeb.ClientReadableStream<schema_pb.App>;

  getHistory(
    request: schema_pb.HistoryFilter,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Apps) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Apps>;

  searchApps(
    request: schema_pb.AppFilter,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Apps) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Apps>;

  uninstallApp(
    request: schema_pb.AppRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  rollbackApp(
    request: schema_pb.AppRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.App) => void
  ): grpcWeb.ClientReadableStream<schema_pb.App>;

  installApp(
    request: schema_pb.AppInput,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.App) => void
  ): grpcWeb.ClientReadableStream<schema_pb.App>;

  updateApp(
    request: schema_pb.AppInput,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.App) => void
  ): grpcWeb.ClientReadableStream<schema_pb.App>;

  searchCharts(
    request: schema_pb.ChartFilter,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Charts) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Charts>;

}

export class HelmProxyServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getApp(
    request: schema_pb.AppRef,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.App>;

  getHistory(
    request: schema_pb.HistoryFilter,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Apps>;

  searchApps(
    request: schema_pb.AppFilter,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Apps>;

  uninstallApp(
    request: schema_pb.AppRef,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  rollbackApp(
    request: schema_pb.AppRef,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.App>;

  installApp(
    request: schema_pb.AppInput,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.App>;

  updateApp(
    request: schema_pb.AppInput,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.App>;

  searchCharts(
    request: schema_pb.ChartFilter,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Charts>;

}

