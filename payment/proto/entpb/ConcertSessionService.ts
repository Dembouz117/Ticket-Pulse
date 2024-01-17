// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { BatchCreateConcertSessionsRequest as _entpb_BatchCreateConcertSessionsRequest, BatchCreateConcertSessionsRequest__Output as _entpb_BatchCreateConcertSessionsRequest__Output } from '../entpb/BatchCreateConcertSessionsRequest';
import type { BatchCreateConcertSessionsResponse as _entpb_BatchCreateConcertSessionsResponse, BatchCreateConcertSessionsResponse__Output as _entpb_BatchCreateConcertSessionsResponse__Output } from '../entpb/BatchCreateConcertSessionsResponse';
import type { ConcertSession as _entpb_ConcertSession, ConcertSession__Output as _entpb_ConcertSession__Output } from '../entpb/ConcertSession';
import type { CreateConcertSessionRequest as _entpb_CreateConcertSessionRequest, CreateConcertSessionRequest__Output as _entpb_CreateConcertSessionRequest__Output } from '../entpb/CreateConcertSessionRequest';
import type { DeleteConcertSessionRequest as _entpb_DeleteConcertSessionRequest, DeleteConcertSessionRequest__Output as _entpb_DeleteConcertSessionRequest__Output } from '../entpb/DeleteConcertSessionRequest';
import type { Empty as _google_protobuf_Empty, Empty__Output as _google_protobuf_Empty__Output } from '../google/protobuf/Empty';
import type { GetConcertSessionRequest as _entpb_GetConcertSessionRequest, GetConcertSessionRequest__Output as _entpb_GetConcertSessionRequest__Output } from '../entpb/GetConcertSessionRequest';
import type { ListConcertSessionRequest as _entpb_ListConcertSessionRequest, ListConcertSessionRequest__Output as _entpb_ListConcertSessionRequest__Output } from '../entpb/ListConcertSessionRequest';
import type { ListConcertSessionResponse as _entpb_ListConcertSessionResponse, ListConcertSessionResponse__Output as _entpb_ListConcertSessionResponse__Output } from '../entpb/ListConcertSessionResponse';
import type { UpdateConcertSessionRequest as _entpb_UpdateConcertSessionRequest, UpdateConcertSessionRequest__Output as _entpb_UpdateConcertSessionRequest__Output } from '../entpb/UpdateConcertSessionRequest';

export interface ConcertSessionServiceClient extends grpc.Client {
  BatchCreate(argument: _entpb_BatchCreateConcertSessionsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateConcertSessionsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateConcertSessionsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_BatchCreateConcertSessionsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateConcertSessionsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateConcertSessionsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateConcertSessionsRequest, callback: grpc.requestCallback<_entpb_BatchCreateConcertSessionsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateConcertSessionsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateConcertSessionsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateConcertSessionsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_BatchCreateConcertSessionsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateConcertSessionsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateConcertSessionsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateConcertSessionsRequest, callback: grpc.requestCallback<_entpb_BatchCreateConcertSessionsResponse__Output>): grpc.ClientUnaryCall;
  
  Create(argument: _entpb_CreateConcertSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateConcertSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateConcertSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateConcertSessionRequest, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateConcertSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateConcertSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateConcertSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateConcertSessionRequest, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  
  Delete(argument: _entpb_DeleteConcertSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteConcertSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteConcertSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteConcertSessionRequest, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteConcertSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteConcertSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteConcertSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteConcertSessionRequest, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  Get(argument: _entpb_GetConcertSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetConcertSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetConcertSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetConcertSessionRequest, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetConcertSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetConcertSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetConcertSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetConcertSessionRequest, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  
  List(argument: _entpb_ListConcertSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListConcertSessionResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListConcertSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ListConcertSessionResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListConcertSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListConcertSessionResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListConcertSessionRequest, callback: grpc.requestCallback<_entpb_ListConcertSessionResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListConcertSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListConcertSessionResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListConcertSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ListConcertSessionResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListConcertSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListConcertSessionResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListConcertSessionRequest, callback: grpc.requestCallback<_entpb_ListConcertSessionResponse__Output>): grpc.ClientUnaryCall;
  
  Update(argument: _entpb_UpdateConcertSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateConcertSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateConcertSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateConcertSessionRequest, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateConcertSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateConcertSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateConcertSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateConcertSessionRequest, callback: grpc.requestCallback<_entpb_ConcertSession__Output>): grpc.ClientUnaryCall;
  
}

export interface ConcertSessionServiceHandlers extends grpc.UntypedServiceImplementation {
  BatchCreate: grpc.handleUnaryCall<_entpb_BatchCreateConcertSessionsRequest__Output, _entpb_BatchCreateConcertSessionsResponse>;
  
  Create: grpc.handleUnaryCall<_entpb_CreateConcertSessionRequest__Output, _entpb_ConcertSession>;
  
  Delete: grpc.handleUnaryCall<_entpb_DeleteConcertSessionRequest__Output, _google_protobuf_Empty>;
  
  Get: grpc.handleUnaryCall<_entpb_GetConcertSessionRequest__Output, _entpb_ConcertSession>;
  
  List: grpc.handleUnaryCall<_entpb_ListConcertSessionRequest__Output, _entpb_ListConcertSessionResponse>;
  
  Update: grpc.handleUnaryCall<_entpb_UpdateConcertSessionRequest__Output, _entpb_ConcertSession>;
  
}

export interface ConcertSessionServiceDefinition extends grpc.ServiceDefinition {
  BatchCreate: MethodDefinition<_entpb_BatchCreateConcertSessionsRequest, _entpb_BatchCreateConcertSessionsResponse, _entpb_BatchCreateConcertSessionsRequest__Output, _entpb_BatchCreateConcertSessionsResponse__Output>
  Create: MethodDefinition<_entpb_CreateConcertSessionRequest, _entpb_ConcertSession, _entpb_CreateConcertSessionRequest__Output, _entpb_ConcertSession__Output>
  Delete: MethodDefinition<_entpb_DeleteConcertSessionRequest, _google_protobuf_Empty, _entpb_DeleteConcertSessionRequest__Output, _google_protobuf_Empty__Output>
  Get: MethodDefinition<_entpb_GetConcertSessionRequest, _entpb_ConcertSession, _entpb_GetConcertSessionRequest__Output, _entpb_ConcertSession__Output>
  List: MethodDefinition<_entpb_ListConcertSessionRequest, _entpb_ListConcertSessionResponse, _entpb_ListConcertSessionRequest__Output, _entpb_ListConcertSessionResponse__Output>
  Update: MethodDefinition<_entpb_UpdateConcertSessionRequest, _entpb_ConcertSession, _entpb_UpdateConcertSessionRequest__Output, _entpb_ConcertSession__Output>
}
