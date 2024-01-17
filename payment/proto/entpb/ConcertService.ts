// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { BatchCreateConcertsRequest as _entpb_BatchCreateConcertsRequest, BatchCreateConcertsRequest__Output as _entpb_BatchCreateConcertsRequest__Output } from '../entpb/BatchCreateConcertsRequest';
import type { BatchCreateConcertsResponse as _entpb_BatchCreateConcertsResponse, BatchCreateConcertsResponse__Output as _entpb_BatchCreateConcertsResponse__Output } from '../entpb/BatchCreateConcertsResponse';
import type { Concert as _entpb_Concert, Concert__Output as _entpb_Concert__Output } from '../entpb/Concert';
import type { CreateConcertRequest as _entpb_CreateConcertRequest, CreateConcertRequest__Output as _entpb_CreateConcertRequest__Output } from '../entpb/CreateConcertRequest';
import type { DeleteConcertRequest as _entpb_DeleteConcertRequest, DeleteConcertRequest__Output as _entpb_DeleteConcertRequest__Output } from '../entpb/DeleteConcertRequest';
import type { Empty as _google_protobuf_Empty, Empty__Output as _google_protobuf_Empty__Output } from '../google/protobuf/Empty';
import type { GetConcertRequest as _entpb_GetConcertRequest, GetConcertRequest__Output as _entpb_GetConcertRequest__Output } from '../entpb/GetConcertRequest';
import type { ListConcertRequest as _entpb_ListConcertRequest, ListConcertRequest__Output as _entpb_ListConcertRequest__Output } from '../entpb/ListConcertRequest';
import type { ListConcertResponse as _entpb_ListConcertResponse, ListConcertResponse__Output as _entpb_ListConcertResponse__Output } from '../entpb/ListConcertResponse';
import type { UpdateConcertRequest as _entpb_UpdateConcertRequest, UpdateConcertRequest__Output as _entpb_UpdateConcertRequest__Output } from '../entpb/UpdateConcertRequest';

export interface ConcertServiceClient extends grpc.Client {
  BatchCreate(argument: _entpb_BatchCreateConcertsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateConcertsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateConcertsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_BatchCreateConcertsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateConcertsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateConcertsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateConcertsRequest, callback: grpc.requestCallback<_entpb_BatchCreateConcertsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateConcertsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateConcertsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateConcertsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_BatchCreateConcertsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateConcertsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateConcertsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateConcertsRequest, callback: grpc.requestCallback<_entpb_BatchCreateConcertsResponse__Output>): grpc.ClientUnaryCall;
  
  Create(argument: _entpb_CreateConcertRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateConcertRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateConcertRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateConcertRequest, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateConcertRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateConcertRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateConcertRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateConcertRequest, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  
  Delete(argument: _entpb_DeleteConcertRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteConcertRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteConcertRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteConcertRequest, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteConcertRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteConcertRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteConcertRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteConcertRequest, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  Get(argument: _entpb_GetConcertRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetConcertRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetConcertRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetConcertRequest, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetConcertRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetConcertRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetConcertRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetConcertRequest, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  
  List(argument: _entpb_ListConcertRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListConcertResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListConcertRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ListConcertResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListConcertRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListConcertResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListConcertRequest, callback: grpc.requestCallback<_entpb_ListConcertResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListConcertRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListConcertResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListConcertRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ListConcertResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListConcertRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListConcertResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListConcertRequest, callback: grpc.requestCallback<_entpb_ListConcertResponse__Output>): grpc.ClientUnaryCall;
  
  Update(argument: _entpb_UpdateConcertRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateConcertRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateConcertRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateConcertRequest, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateConcertRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateConcertRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateConcertRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateConcertRequest, callback: grpc.requestCallback<_entpb_Concert__Output>): grpc.ClientUnaryCall;
  
}

export interface ConcertServiceHandlers extends grpc.UntypedServiceImplementation {
  BatchCreate: grpc.handleUnaryCall<_entpb_BatchCreateConcertsRequest__Output, _entpb_BatchCreateConcertsResponse>;
  
  Create: grpc.handleUnaryCall<_entpb_CreateConcertRequest__Output, _entpb_Concert>;
  
  Delete: grpc.handleUnaryCall<_entpb_DeleteConcertRequest__Output, _google_protobuf_Empty>;
  
  Get: grpc.handleUnaryCall<_entpb_GetConcertRequest__Output, _entpb_Concert>;
  
  List: grpc.handleUnaryCall<_entpb_ListConcertRequest__Output, _entpb_ListConcertResponse>;
  
  Update: grpc.handleUnaryCall<_entpb_UpdateConcertRequest__Output, _entpb_Concert>;
  
}

export interface ConcertServiceDefinition extends grpc.ServiceDefinition {
  BatchCreate: MethodDefinition<_entpb_BatchCreateConcertsRequest, _entpb_BatchCreateConcertsResponse, _entpb_BatchCreateConcertsRequest__Output, _entpb_BatchCreateConcertsResponse__Output>
  Create: MethodDefinition<_entpb_CreateConcertRequest, _entpb_Concert, _entpb_CreateConcertRequest__Output, _entpb_Concert__Output>
  Delete: MethodDefinition<_entpb_DeleteConcertRequest, _google_protobuf_Empty, _entpb_DeleteConcertRequest__Output, _google_protobuf_Empty__Output>
  Get: MethodDefinition<_entpb_GetConcertRequest, _entpb_Concert, _entpb_GetConcertRequest__Output, _entpb_Concert__Output>
  List: MethodDefinition<_entpb_ListConcertRequest, _entpb_ListConcertResponse, _entpb_ListConcertRequest__Output, _entpb_ListConcertResponse__Output>
  Update: MethodDefinition<_entpb_UpdateConcertRequest, _entpb_Concert, _entpb_UpdateConcertRequest__Output, _entpb_Concert__Output>
}
