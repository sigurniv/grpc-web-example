/**
 * @fileoverview gRPC-Web generated client stub for game
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.game = require('./game_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.game.GameServiceClient =
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

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.game.GameServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.game.GameServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.game.GameServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.game.SearchGameRequest,
 *   !proto.game.SearchGameResponse>}
 */
const methodInfo_GameService_SearchGame = new grpc.web.AbstractClientBase.MethodInfo(
  proto.game.SearchGameResponse,
  /** @param {!proto.game.SearchGameRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.game.SearchGameResponse.deserializeBinary
);


/**
 * @param {!proto.game.SearchGameRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.game.SearchGameResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.game.SearchGameResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.game.GameServiceClient.prototype.searchGame =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/game.GameService/SearchGame',
      request,
      metadata,
      methodInfo_GameService_SearchGame,
      callback);
};


/**
 * @param {!proto.game.SearchGameRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.game.SearchGameResponse>}
 *     The XHR Node Readable Stream
 */
proto.game.GameServicePromiseClient.prototype.searchGame =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.searchGame(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.game.GameDetailsRequest,
 *   !proto.game.GameDetailsResponse>}
 */
const methodInfo_GameService_GameDetails = new grpc.web.AbstractClientBase.MethodInfo(
  proto.game.GameDetailsResponse,
  /** @param {!proto.game.GameDetailsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.game.GameDetailsResponse.deserializeBinary
);


/**
 * @param {!proto.game.GameDetailsRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.game.GameDetailsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.game.GameDetailsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.game.GameServiceClient.prototype.gameDetails =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/game.GameService/GameDetails',
      request,
      metadata,
      methodInfo_GameService_GameDetails,
      callback);
};


/**
 * @param {!proto.game.GameDetailsRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.game.GameDetailsResponse>}
 *     The XHR Node Readable Stream
 */
proto.game.GameServicePromiseClient.prototype.gameDetails =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.gameDetails(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.game;

