/**
 * @fileoverview gRPC-Web generated client stub for helmProxy
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
proto.helmProxy = require('./schema_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.helmProxy.HelmProxyServiceClient =
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
proto.helmProxy.HelmProxyServicePromiseClient =
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
 *   !proto.helmProxy.AppRef,
 *   !proto.helmProxy.App>}
 */
const methodDescriptor_HelmProxyService_GetApp = new grpc.web.MethodDescriptor(
  '/helmProxy.HelmProxyService/GetApp',
  grpc.web.MethodType.UNARY,
  proto.helmProxy.AppRef,
  proto.helmProxy.App,
  /**
   * @param {!proto.helmProxy.AppRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.App.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmProxy.AppRef,
 *   !proto.helmProxy.App>}
 */
const methodInfo_HelmProxyService_GetApp = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmProxy.App,
  /**
   * @param {!proto.helmProxy.AppRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.App.deserializeBinary
);


/**
 * @param {!proto.helmProxy.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmProxy.App)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmProxy.App>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmProxy.HelmProxyServiceClient.prototype.getApp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmProxy.HelmProxyService/GetApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_GetApp,
      callback);
};


/**
 * @param {!proto.helmProxy.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmProxy.App>}
 *     Promise that resolves to the response
 */
proto.helmProxy.HelmProxyServicePromiseClient.prototype.getApp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmProxy.HelmProxyService/GetApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_GetApp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmProxy.HistoryFilter,
 *   !proto.helmProxy.Apps>}
 */
const methodDescriptor_HelmProxyService_GetHistory = new grpc.web.MethodDescriptor(
  '/helmProxy.HelmProxyService/GetHistory',
  grpc.web.MethodType.UNARY,
  proto.helmProxy.HistoryFilter,
  proto.helmProxy.Apps,
  /**
   * @param {!proto.helmProxy.HistoryFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.Apps.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmProxy.HistoryFilter,
 *   !proto.helmProxy.Apps>}
 */
const methodInfo_HelmProxyService_GetHistory = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmProxy.Apps,
  /**
   * @param {!proto.helmProxy.HistoryFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.Apps.deserializeBinary
);


/**
 * @param {!proto.helmProxy.HistoryFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmProxy.Apps)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmProxy.Apps>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmProxy.HelmProxyServiceClient.prototype.getHistory =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmProxy.HelmProxyService/GetHistory',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_GetHistory,
      callback);
};


/**
 * @param {!proto.helmProxy.HistoryFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmProxy.Apps>}
 *     Promise that resolves to the response
 */
proto.helmProxy.HelmProxyServicePromiseClient.prototype.getHistory =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmProxy.HelmProxyService/GetHistory',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_GetHistory);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmProxy.AppFilter,
 *   !proto.helmProxy.Apps>}
 */
const methodDescriptor_HelmProxyService_SearchApps = new grpc.web.MethodDescriptor(
  '/helmProxy.HelmProxyService/SearchApps',
  grpc.web.MethodType.UNARY,
  proto.helmProxy.AppFilter,
  proto.helmProxy.Apps,
  /**
   * @param {!proto.helmProxy.AppFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.Apps.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmProxy.AppFilter,
 *   !proto.helmProxy.Apps>}
 */
const methodInfo_HelmProxyService_SearchApps = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmProxy.Apps,
  /**
   * @param {!proto.helmProxy.AppFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.Apps.deserializeBinary
);


/**
 * @param {!proto.helmProxy.AppFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmProxy.Apps)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmProxy.Apps>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmProxy.HelmProxyServiceClient.prototype.searchApps =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmProxy.HelmProxyService/SearchApps',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_SearchApps,
      callback);
};


/**
 * @param {!proto.helmProxy.AppFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmProxy.Apps>}
 *     Promise that resolves to the response
 */
proto.helmProxy.HelmProxyServicePromiseClient.prototype.searchApps =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmProxy.HelmProxyService/SearchApps',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_SearchApps);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmProxy.AppRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_HelmProxyService_UninstallApp = new grpc.web.MethodDescriptor(
  '/helmProxy.HelmProxyService/UninstallApp',
  grpc.web.MethodType.UNARY,
  proto.helmProxy.AppRef,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.helmProxy.AppRef} request
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
 *   !proto.helmProxy.AppRef,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_HelmProxyService_UninstallApp = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.helmProxy.AppRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.helmProxy.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmProxy.HelmProxyServiceClient.prototype.uninstallApp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmProxy.HelmProxyService/UninstallApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_UninstallApp,
      callback);
};


/**
 * @param {!proto.helmProxy.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.helmProxy.HelmProxyServicePromiseClient.prototype.uninstallApp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmProxy.HelmProxyService/UninstallApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_UninstallApp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmProxy.AppRef,
 *   !proto.helmProxy.App>}
 */
const methodDescriptor_HelmProxyService_RollbackApp = new grpc.web.MethodDescriptor(
  '/helmProxy.HelmProxyService/RollbackApp',
  grpc.web.MethodType.UNARY,
  proto.helmProxy.AppRef,
  proto.helmProxy.App,
  /**
   * @param {!proto.helmProxy.AppRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.App.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmProxy.AppRef,
 *   !proto.helmProxy.App>}
 */
const methodInfo_HelmProxyService_RollbackApp = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmProxy.App,
  /**
   * @param {!proto.helmProxy.AppRef} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.App.deserializeBinary
);


/**
 * @param {!proto.helmProxy.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmProxy.App)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmProxy.App>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmProxy.HelmProxyServiceClient.prototype.rollbackApp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmProxy.HelmProxyService/RollbackApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_RollbackApp,
      callback);
};


/**
 * @param {!proto.helmProxy.AppRef} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmProxy.App>}
 *     Promise that resolves to the response
 */
proto.helmProxy.HelmProxyServicePromiseClient.prototype.rollbackApp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmProxy.HelmProxyService/RollbackApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_RollbackApp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmProxy.AppInput,
 *   !proto.helmProxy.App>}
 */
const methodDescriptor_HelmProxyService_InstallApp = new grpc.web.MethodDescriptor(
  '/helmProxy.HelmProxyService/InstallApp',
  grpc.web.MethodType.UNARY,
  proto.helmProxy.AppInput,
  proto.helmProxy.App,
  /**
   * @param {!proto.helmProxy.AppInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.App.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmProxy.AppInput,
 *   !proto.helmProxy.App>}
 */
const methodInfo_HelmProxyService_InstallApp = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmProxy.App,
  /**
   * @param {!proto.helmProxy.AppInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.App.deserializeBinary
);


/**
 * @param {!proto.helmProxy.AppInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmProxy.App)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmProxy.App>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmProxy.HelmProxyServiceClient.prototype.installApp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmProxy.HelmProxyService/InstallApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_InstallApp,
      callback);
};


/**
 * @param {!proto.helmProxy.AppInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmProxy.App>}
 *     Promise that resolves to the response
 */
proto.helmProxy.HelmProxyServicePromiseClient.prototype.installApp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmProxy.HelmProxyService/InstallApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_InstallApp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmProxy.AppInput,
 *   !proto.helmProxy.App>}
 */
const methodDescriptor_HelmProxyService_UpdateApp = new grpc.web.MethodDescriptor(
  '/helmProxy.HelmProxyService/UpdateApp',
  grpc.web.MethodType.UNARY,
  proto.helmProxy.AppInput,
  proto.helmProxy.App,
  /**
   * @param {!proto.helmProxy.AppInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.App.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmProxy.AppInput,
 *   !proto.helmProxy.App>}
 */
const methodInfo_HelmProxyService_UpdateApp = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmProxy.App,
  /**
   * @param {!proto.helmProxy.AppInput} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.App.deserializeBinary
);


/**
 * @param {!proto.helmProxy.AppInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmProxy.App)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmProxy.App>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmProxy.HelmProxyServiceClient.prototype.updateApp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmProxy.HelmProxyService/UpdateApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_UpdateApp,
      callback);
};


/**
 * @param {!proto.helmProxy.AppInput} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmProxy.App>}
 *     Promise that resolves to the response
 */
proto.helmProxy.HelmProxyServicePromiseClient.prototype.updateApp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmProxy.HelmProxyService/UpdateApp',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_UpdateApp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.helmProxy.ChartFilter,
 *   !proto.helmProxy.Charts>}
 */
const methodDescriptor_HelmProxyService_SearchCharts = new grpc.web.MethodDescriptor(
  '/helmProxy.HelmProxyService/SearchCharts',
  grpc.web.MethodType.UNARY,
  proto.helmProxy.ChartFilter,
  proto.helmProxy.Charts,
  /**
   * @param {!proto.helmProxy.ChartFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.Charts.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.helmProxy.ChartFilter,
 *   !proto.helmProxy.Charts>}
 */
const methodInfo_HelmProxyService_SearchCharts = new grpc.web.AbstractClientBase.MethodInfo(
  proto.helmProxy.Charts,
  /**
   * @param {!proto.helmProxy.ChartFilter} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.helmProxy.Charts.deserializeBinary
);


/**
 * @param {!proto.helmProxy.ChartFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.helmProxy.Charts)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.helmProxy.Charts>|undefined}
 *     The XHR Node Readable Stream
 */
proto.helmProxy.HelmProxyServiceClient.prototype.searchCharts =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/helmProxy.HelmProxyService/SearchCharts',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_SearchCharts,
      callback);
};


/**
 * @param {!proto.helmProxy.ChartFilter} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.helmProxy.Charts>}
 *     Promise that resolves to the response
 */
proto.helmProxy.HelmProxyServicePromiseClient.prototype.searchCharts =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/helmProxy.HelmProxyService/SearchCharts',
      request,
      metadata || {},
      methodDescriptor_HelmProxyService_SearchCharts);
};


module.exports = proto.helmProxy;

