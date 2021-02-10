/**
 * @fileoverview gRPC-Web generated client stub for helmgate
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('./google/api/annotations_pb.js')

var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var google_protobuf_any_pb = require('google-protobuf/google/protobuf/any_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var github_com_mwitkow_go$proto$validators_validator_pb = require('./github.com/mwitkow/go-proto-validators/validator_pb.js')
const proto = {};
proto.helmgate = require('./schema_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.helmgate.HelmProxyServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.helmgate.HelmProxyServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmgate.AppRef,
 *   !proto.helmgate.App>}
 */
const methodDescriptor_HelmProxyService_GetApp = new grpc.web.MethodDescriptor(
  '/helmgate.HelmProxyService/GetApp',
  grpc.web.MethodType.UNARY,
  proto.helmgate.AppRef,
  proto.helmgate.App,
  /**
   * @param {!proto.helmgate.AppRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.App.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmgate.AppRef,
 *   !proto.helmgate.App>}
 */
const methodInfo_HelmProxyService_GetApp = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmgate.App,
  /**
   * @param {!proto.helmgate.AppRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.App.deserializeBinary
);


/**
 * @param {!proto.helmgate.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmgate.App)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmgate.App>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmgate.HelmProxyServiceClient.prototype.getApp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmgate.HelmProxyService/GetApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_GetApp,
      callback);
};


/**
 * @param {!proto.helmgate.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmgate.App>}
 *     Promise that resolves to the response
 */
proto.helmgate.HelmProxyServicePromiseClient.prototype.getApp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmgate.HelmProxyService/GetApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_GetApp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmgate.HistoryFilter,
 *   !proto.helmgate.Apps>}
 */
const methodDescriptor_HelmProxyService_GetHistory = new grpc.web.MethodDescriptor(
  '/helmgate.HelmProxyService/GetHistory',
  grpc.web.MethodType.UNARY,
  proto.helmgate.HistoryFilter,
  proto.helmgate.Apps,
  /**
   * @param {!proto.helmgate.HistoryFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.Apps.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmgate.HistoryFilter,
 *   !proto.helmgate.Apps>}
 */
const methodInfo_HelmProxyService_GetHistory = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmgate.Apps,
  /**
   * @param {!proto.helmgate.HistoryFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.Apps.deserializeBinary
);


/**
 * @param {!proto.helmgate.HistoryFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmgate.Apps)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmgate.Apps>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmgate.HelmProxyServiceClient.prototype.getHistory =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmgate.HelmProxyService/GetHistory',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_GetHistory,
      callback);
};


/**
 * @param {!proto.helmgate.HistoryFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmgate.Apps>}
 *     Promise that resolves to the response
 */
proto.helmgate.HelmProxyServicePromiseClient.prototype.getHistory =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmgate.HelmProxyService/GetHistory',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_GetHistory);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmgate.AppFilter,
 *   !proto.helmgate.Apps>}
 */
const methodDescriptor_HelmProxyService_SearchApps = new grpc.web.MethodDescriptor(
  '/helmgate.HelmProxyService/SearchApps',
  grpc.web.MethodType.UNARY,
  proto.helmgate.AppFilter,
  proto.helmgate.Apps,
  /**
   * @param {!proto.helmgate.AppFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.Apps.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmgate.AppFilter,
 *   !proto.helmgate.Apps>}
 */
const methodInfo_HelmProxyService_SearchApps = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmgate.Apps,
  /**
   * @param {!proto.helmgate.AppFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.Apps.deserializeBinary
);


/**
 * @param {!proto.helmgate.AppFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmgate.Apps)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmgate.Apps>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmgate.HelmProxyServiceClient.prototype.searchApps =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmgate.HelmProxyService/SearchApps',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_SearchApps,
      callback);
};


/**
 * @param {!proto.helmgate.AppFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmgate.Apps>}
 *     Promise that resolves to the response
 */
proto.helmgate.HelmProxyServicePromiseClient.prototype.searchApps =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmgate.HelmProxyService/SearchApps',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_SearchApps);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmgate.AppRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_HelmProxyService_UninstallApp = new grpc.web.MethodDescriptor(
  '/helmgate.HelmProxyService/UninstallApp',
  grpc.web.MethodType.UNARY,
  proto.helmgate.AppRef,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.helmgate.AppRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmgate.AppRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_HelmProxyService_UninstallApp = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.helmgate.AppRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.helmgate.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmgate.HelmProxyServiceClient.prototype.uninstallApp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmgate.HelmProxyService/UninstallApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_UninstallApp,
      callback);
};


/**
 * @param {!proto.helmgate.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.helmgate.HelmProxyServicePromiseClient.prototype.uninstallApp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmgate.HelmProxyService/UninstallApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_UninstallApp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmgate.AppRef,
 *   !proto.helmgate.App>}
 */
const methodDescriptor_HelmProxyService_RollbackApp = new grpc.web.MethodDescriptor(
  '/helmgate.HelmProxyService/RollbackApp',
  grpc.web.MethodType.UNARY,
  proto.helmgate.AppRef,
  proto.helmgate.App,
  /**
   * @param {!proto.helmgate.AppRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.App.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmgate.AppRef,
 *   !proto.helmgate.App>}
 */
const methodInfo_HelmProxyService_RollbackApp = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmgate.App,
  /**
   * @param {!proto.helmgate.AppRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.App.deserializeBinary
);


/**
 * @param {!proto.helmgate.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmgate.App)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmgate.App>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmgate.HelmProxyServiceClient.prototype.rollbackApp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmgate.HelmProxyService/RollbackApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_RollbackApp,
      callback);
};


/**
 * @param {!proto.helmgate.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmgate.App>}
 *     Promise that resolves to the response
 */
proto.helmgate.HelmProxyServicePromiseClient.prototype.rollbackApp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmgate.HelmProxyService/RollbackApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_RollbackApp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmgate.AppInput,
 *   !proto.helmgate.App>}
 */
const methodDescriptor_HelmProxyService_InstallApp = new grpc.web.MethodDescriptor(
  '/helmgate.HelmProxyService/InstallApp',
  grpc.web.MethodType.UNARY,
  proto.helmgate.AppInput,
  proto.helmgate.App,
  /**
   * @param {!proto.helmgate.AppInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.App.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmgate.AppInput,
 *   !proto.helmgate.App>}
 */
const methodInfo_HelmProxyService_InstallApp = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmgate.App,
  /**
   * @param {!proto.helmgate.AppInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.App.deserializeBinary
);


/**
 * @param {!proto.helmgate.AppInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmgate.App)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmgate.App>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmgate.HelmProxyServiceClient.prototype.installApp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmgate.HelmProxyService/InstallApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_InstallApp,
      callback);
};


/**
 * @param {!proto.helmgate.AppInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmgate.App>}
 *     Promise that resolves to the response
 */
proto.helmgate.HelmProxyServicePromiseClient.prototype.installApp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmgate.HelmProxyService/InstallApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_InstallApp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmgate.AppInput,
 *   !proto.helmgate.App>}
 */
const methodDescriptor_HelmProxyService_UpdateApp = new grpc.web.MethodDescriptor(
  '/helmgate.HelmProxyService/UpdateApp',
  grpc.web.MethodType.UNARY,
  proto.helmgate.AppInput,
  proto.helmgate.App,
  /**
   * @param {!proto.helmgate.AppInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.App.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmgate.AppInput,
 *   !proto.helmgate.App>}
 */
const methodInfo_HelmProxyService_UpdateApp = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmgate.App,
  /**
   * @param {!proto.helmgate.AppInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.App.deserializeBinary
);


/**
 * @param {!proto.helmgate.AppInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmgate.App)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmgate.App>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmgate.HelmProxyServiceClient.prototype.updateApp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmgate.HelmProxyService/UpdateApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_UpdateApp,
      callback);
};


/**
 * @param {!proto.helmgate.AppInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmgate.App>}
 *     Promise that resolves to the response
 */
proto.helmgate.HelmProxyServicePromiseClient.prototype.updateApp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmgate.HelmProxyService/UpdateApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_UpdateApp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmgate.ChartFilter,
 *   !proto.helmgate.Charts>}
 */
const methodDescriptor_HelmProxyService_SearchCharts = new grpc.web.MethodDescriptor(
  '/helmgate.HelmProxyService/SearchCharts',
  grpc.web.MethodType.UNARY,
  proto.helmgate.ChartFilter,
  proto.helmgate.Charts,
  /**
   * @param {!proto.helmgate.ChartFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.Charts.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmgate.ChartFilter,
 *   !proto.helmgate.Charts>}
 */
const methodInfo_HelmProxyService_SearchCharts = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmgate.Charts,
  /**
   * @param {!proto.helmgate.ChartFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmgate.Charts.deserializeBinary
);


/**
 * @param {!proto.helmgate.ChartFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmgate.Charts)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmgate.Charts>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmgate.HelmProxyServiceClient.prototype.searchCharts =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmgate.HelmProxyService/SearchCharts',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_SearchCharts,
      callback);
};


/**
 * @param {!proto.helmgate.ChartFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmgate.Charts>}
 *     Promise that resolves to the response
 */
proto.helmgate.HelmProxyServicePromiseClient.prototype.searchCharts =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmgate.HelmProxyService/SearchCharts',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_SearchCharts);
};


module.exports = proto.helmgate;

