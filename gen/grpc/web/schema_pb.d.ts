import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from './google/api/annotations_pb';
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_any_pb from 'google-protobuf/google/protobuf/any_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as github_com_mwitkow_go$proto$validators_validator_pb from './github.com/mwitkow/go-proto-validators/validator_pb';


export class Dependency extends jspb.Message {
  getChart(): string;
  setChart(value: string): Dependency;

  getVersion(): string;
  setVersion(value: string): Dependency;

  getRepository(): string;
  setRepository(value: string): Dependency;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Dependency.AsObject;
  static toObject(includeInstance: boolean, msg: Dependency): Dependency.AsObject;
  static serializeBinaryToWriter(message: Dependency, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Dependency;
  static deserializeBinaryFromReader(message: Dependency, reader: jspb.BinaryReader): Dependency;
}

export namespace Dependency {
  export type AsObject = {
    chart: string,
    version: string,
    repository: string,
  }
}

export class Maintainer extends jspb.Message {
  getName(): string;
  setName(value: string): Maintainer;

  getEmail(): string;
  setEmail(value: string): Maintainer;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Maintainer.AsObject;
  static toObject(includeInstance: boolean, msg: Maintainer): Maintainer.AsObject;
  static serializeBinaryToWriter(message: Maintainer, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Maintainer;
  static deserializeBinaryFromReader(message: Maintainer, reader: jspb.BinaryReader): Maintainer;
}

export namespace Maintainer {
  export type AsObject = {
    name: string,
    email: string,
  }
}

export class ChartFilter extends jspb.Message {
  getTerm(): string;
  setTerm(value: string): ChartFilter;

  getRegex(): boolean;
  setRegex(value: boolean): ChartFilter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChartFilter.AsObject;
  static toObject(includeInstance: boolean, msg: ChartFilter): ChartFilter.AsObject;
  static serializeBinaryToWriter(message: ChartFilter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChartFilter;
  static deserializeBinaryFromReader(message: ChartFilter, reader: jspb.BinaryReader): ChartFilter;
}

export namespace ChartFilter {
  export type AsObject = {
    term: string,
    regex: boolean,
  }
}

export class Chart extends jspb.Message {
  getName(): string;
  setName(value: string): Chart;

  getHome(): string;
  setHome(value: string): Chart;

  getDescription(): string;
  setDescription(value: string): Chart;

  getVersion(): string;
  setVersion(value: string): Chart;

  getSourcesList(): Array<string>;
  setSourcesList(value: Array<string>): Chart;
  clearSourcesList(): Chart;
  addSources(value: string, index?: number): Chart;

  getKeywordsList(): Array<string>;
  setKeywordsList(value: Array<string>): Chart;
  clearKeywordsList(): Chart;
  addKeywords(value: string, index?: number): Chart;

  getIcon(): string;
  setIcon(value: string): Chart;

  getDeprecated(): boolean;
  setDeprecated(value: boolean): Chart;

  getDependenciesList(): Array<Dependency>;
  setDependenciesList(value: Array<Dependency>): Chart;
  clearDependenciesList(): Chart;
  addDependencies(value?: Dependency, index?: number): Dependency;

  getMaintainersList(): Array<Maintainer>;
  setMaintainersList(value: Array<Maintainer>): Chart;
  clearMaintainersList(): Chart;
  addMaintainers(value?: Maintainer, index?: number): Maintainer;

  getMetadataMap(): jspb.Map<string, string>;
  clearMetadataMap(): Chart;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Chart.AsObject;
  static toObject(includeInstance: boolean, msg: Chart): Chart.AsObject;
  static serializeBinaryToWriter(message: Chart, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Chart;
  static deserializeBinaryFromReader(message: Chart, reader: jspb.BinaryReader): Chart;
}

export namespace Chart {
  export type AsObject = {
    name: string,
    home: string,
    description: string,
    version: string,
    sourcesList: Array<string>,
    keywordsList: Array<string>,
    icon: string,
    deprecated: boolean,
    dependenciesList: Array<Dependency.AsObject>,
    maintainersList: Array<Maintainer.AsObject>,
    metadataMap: Array<[string, string]>,
  }
}

export class Charts extends jspb.Message {
  getChartsList(): Array<Chart>;
  setChartsList(value: Array<Chart>): Charts;
  clearChartsList(): Charts;
  addCharts(value?: Chart, index?: number): Chart;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Charts.AsObject;
  static toObject(includeInstance: boolean, msg: Charts): Charts.AsObject;
  static serializeBinaryToWriter(message: Charts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Charts;
  static deserializeBinaryFromReader(message: Charts, reader: jspb.BinaryReader): Charts;
}

export namespace Charts {
  export type AsObject = {
    chartsList: Array<Chart.AsObject>,
  }
}

export class App extends jspb.Message {
  getName(): string;
  setName(value: string): App;

  getNamespace(): string;
  setNamespace(value: string): App;

  getRelease(): Release | undefined;
  setRelease(value?: Release): App;
  hasRelease(): boolean;
  clearRelease(): App;

  getChart(): Chart | undefined;
  setChart(value?: Chart): App;
  hasChart(): boolean;
  clearChart(): App;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): App.AsObject;
  static toObject(includeInstance: boolean, msg: App): App.AsObject;
  static serializeBinaryToWriter(message: App, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): App;
  static deserializeBinaryFromReader(message: App, reader: jspb.BinaryReader): App;
}

export namespace App {
  export type AsObject = {
    name: string,
    namespace: string,
    release?: Release.AsObject,
    chart?: Chart.AsObject,
  }
}

export class Apps extends jspb.Message {
  getAppsList(): Array<App>;
  setAppsList(value: Array<App>): Apps;
  clearAppsList(): Apps;
  addApps(value?: App, index?: number): App;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Apps.AsObject;
  static toObject(includeInstance: boolean, msg: Apps): Apps.AsObject;
  static serializeBinaryToWriter(message: Apps, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Apps;
  static deserializeBinaryFromReader(message: Apps, reader: jspb.BinaryReader): Apps;
}

export namespace Apps {
  export type AsObject = {
    appsList: Array<App.AsObject>,
  }
}

export class AppFilter extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): AppFilter;

  getSelector(): string;
  setSelector(value: string): AppFilter;

  getLimit(): number;
  setLimit(value: number): AppFilter;

  getOffset(): number;
  setOffset(value: number): AppFilter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AppFilter.AsObject;
  static toObject(includeInstance: boolean, msg: AppFilter): AppFilter.AsObject;
  static serializeBinaryToWriter(message: AppFilter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AppFilter;
  static deserializeBinaryFromReader(message: AppFilter, reader: jspb.BinaryReader): AppFilter;
}

export namespace AppFilter {
  export type AsObject = {
    namespace: string,
    selector: string,
    limit: number,
    offset: number,
  }
}

export class Release extends jspb.Message {
  getVersion(): number;
  setVersion(value: number): Release;

  getConfig(): google_protobuf_struct_pb.Struct | undefined;
  setConfig(value?: google_protobuf_struct_pb.Struct): Release;
  hasConfig(): boolean;
  clearConfig(): Release;

  getNotes(): string;
  setNotes(value: string): Release;

  getDescription(): string;
  setDescription(value: string): Release;

  getStatus(): string;
  setStatus(value: string): Release;

  getTimestamps(): Timestamps | undefined;
  setTimestamps(value?: Timestamps): Release;
  hasTimestamps(): boolean;
  clearTimestamps(): Release;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Release.AsObject;
  static toObject(includeInstance: boolean, msg: Release): Release.AsObject;
  static serializeBinaryToWriter(message: Release, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Release;
  static deserializeBinaryFromReader(message: Release, reader: jspb.BinaryReader): Release;
}

export namespace Release {
  export type AsObject = {
    version: number,
    config?: google_protobuf_struct_pb.Struct.AsObject,
    notes: string,
    description: string,
    status: string,
    timestamps?: Timestamps.AsObject,
  }
}

export class Timestamps extends jspb.Message {
  getCreated(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreated(value?: google_protobuf_timestamp_pb.Timestamp): Timestamps;
  hasCreated(): boolean;
  clearCreated(): Timestamps;

  getUpdated(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdated(value?: google_protobuf_timestamp_pb.Timestamp): Timestamps;
  hasUpdated(): boolean;
  clearUpdated(): Timestamps;

  getDeleted(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeleted(value?: google_protobuf_timestamp_pb.Timestamp): Timestamps;
  hasDeleted(): boolean;
  clearDeleted(): Timestamps;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Timestamps.AsObject;
  static toObject(includeInstance: boolean, msg: Timestamps): Timestamps.AsObject;
  static serializeBinaryToWriter(message: Timestamps, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Timestamps;
  static deserializeBinaryFromReader(message: Timestamps, reader: jspb.BinaryReader): Timestamps;
}

export namespace Timestamps {
  export type AsObject = {
    created?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updated?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deleted?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class AppRef extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): AppRef;

  getName(): string;
  setName(value: string): AppRef;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AppRef.AsObject;
  static toObject(includeInstance: boolean, msg: AppRef): AppRef.AsObject;
  static serializeBinaryToWriter(message: AppRef, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AppRef;
  static deserializeBinaryFromReader(message: AppRef, reader: jspb.BinaryReader): AppRef;
}

export namespace AppRef {
  export type AsObject = {
    namespace: string,
    name: string,
  }
}

export class AppInput extends jspb.Message {
  getNamespace(): string;
  setNamespace(value: string): AppInput;

  getChart(): string;
  setChart(value: string): AppInput;

  getAppName(): string;
  setAppName(value: string): AppInput;

  getConfigMap(): jspb.Map<string, string>;
  clearConfigMap(): AppInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AppInput.AsObject;
  static toObject(includeInstance: boolean, msg: AppInput): AppInput.AsObject;
  static serializeBinaryToWriter(message: AppInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AppInput;
  static deserializeBinaryFromReader(message: AppInput, reader: jspb.BinaryReader): AppInput;
}

export namespace AppInput {
  export type AsObject = {
    namespace: string,
    chart: string,
    appName: string,
    configMap: Array<[string, string]>,
  }
}

export class NamespaceRef extends jspb.Message {
  getName(): string;
  setName(value: string): NamespaceRef;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NamespaceRef.AsObject;
  static toObject(includeInstance: boolean, msg: NamespaceRef): NamespaceRef.AsObject;
  static serializeBinaryToWriter(message: NamespaceRef, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NamespaceRef;
  static deserializeBinaryFromReader(message: NamespaceRef, reader: jspb.BinaryReader): NamespaceRef;
}

export namespace NamespaceRef {
  export type AsObject = {
    name: string,
  }
}

export class NamespaceRefs extends jspb.Message {
  getNamespacesList(): Array<NamespaceRef>;
  setNamespacesList(value: Array<NamespaceRef>): NamespaceRefs;
  clearNamespacesList(): NamespaceRefs;
  addNamespaces(value?: NamespaceRef, index?: number): NamespaceRef;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NamespaceRefs.AsObject;
  static toObject(includeInstance: boolean, msg: NamespaceRefs): NamespaceRefs.AsObject;
  static serializeBinaryToWriter(message: NamespaceRefs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NamespaceRefs;
  static deserializeBinaryFromReader(message: NamespaceRefs, reader: jspb.BinaryReader): NamespaceRefs;
}

export namespace NamespaceRefs {
  export type AsObject = {
    namespacesList: Array<NamespaceRef.AsObject>,
  }
}

export class HistoryFilter extends jspb.Message {
  getRef(): AppRef | undefined;
  setRef(value?: AppRef): HistoryFilter;
  hasRef(): boolean;
  clearRef(): HistoryFilter;

  getLimit(): number;
  setLimit(value: number): HistoryFilter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HistoryFilter.AsObject;
  static toObject(includeInstance: boolean, msg: HistoryFilter): HistoryFilter.AsObject;
  static serializeBinaryToWriter(message: HistoryFilter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HistoryFilter;
  static deserializeBinaryFromReader(message: HistoryFilter, reader: jspb.BinaryReader): HistoryFilter;
}

export namespace HistoryFilter {
  export type AsObject = {
    ref?: AppRef.AsObject,
    limit: number,
  }
}

