// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { BatchCreateTicketsRequest as _entpb_BatchCreateTicketsRequest, BatchCreateTicketsRequest__Output as _entpb_BatchCreateTicketsRequest__Output } from '../entpb/BatchCreateTicketsRequest';
import type { BatchCreateTicketsResponse as _entpb_BatchCreateTicketsResponse, BatchCreateTicketsResponse__Output as _entpb_BatchCreateTicketsResponse__Output } from '../entpb/BatchCreateTicketsResponse';
import type { CreateTicketRequest as _entpb_CreateTicketRequest, CreateTicketRequest__Output as _entpb_CreateTicketRequest__Output } from '../entpb/CreateTicketRequest';
import type { DeleteTicketRequest as _entpb_DeleteTicketRequest, DeleteTicketRequest__Output as _entpb_DeleteTicketRequest__Output } from '../entpb/DeleteTicketRequest';
import type { Empty as _google_protobuf_Empty, Empty__Output as _google_protobuf_Empty__Output } from '../google/protobuf/Empty';
import type { GetTicketRequest as _entpb_GetTicketRequest, GetTicketRequest__Output as _entpb_GetTicketRequest__Output } from '../entpb/GetTicketRequest';
import type { ListTicketRequest as _entpb_ListTicketRequest, ListTicketRequest__Output as _entpb_ListTicketRequest__Output } from '../entpb/ListTicketRequest';
import type { ListTicketResponse as _entpb_ListTicketResponse, ListTicketResponse__Output as _entpb_ListTicketResponse__Output } from '../entpb/ListTicketResponse';
import type { Ticket as _entpb_Ticket, Ticket__Output as _entpb_Ticket__Output } from '../entpb/Ticket';
import type { UpdateTicketRequest as _entpb_UpdateTicketRequest, UpdateTicketRequest__Output as _entpb_UpdateTicketRequest__Output } from '../entpb/UpdateTicketRequest';

export interface TicketServiceClient extends grpc.Client {
  BatchCreate(argument: _entpb_BatchCreateTicketsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateTicketsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateTicketsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_BatchCreateTicketsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateTicketsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateTicketsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateTicketsRequest, callback: grpc.requestCallback<_entpb_BatchCreateTicketsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateTicketsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateTicketsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateTicketsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_BatchCreateTicketsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateTicketsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateTicketsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateTicketsRequest, callback: grpc.requestCallback<_entpb_BatchCreateTicketsResponse__Output>): grpc.ClientUnaryCall;
  
  Create(argument: _entpb_CreateTicketRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateTicketRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateTicketRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateTicketRequest, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateTicketRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateTicketRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateTicketRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateTicketRequest, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  
  Delete(argument: _entpb_DeleteTicketRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteTicketRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteTicketRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteTicketRequest, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteTicketRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteTicketRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteTicketRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteTicketRequest, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  Get(argument: _entpb_GetTicketRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetTicketRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetTicketRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetTicketRequest, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetTicketRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetTicketRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetTicketRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetTicketRequest, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  
  List(argument: _entpb_ListTicketRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListTicketResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListTicketRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ListTicketResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListTicketRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListTicketResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListTicketRequest, callback: grpc.requestCallback<_entpb_ListTicketResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListTicketRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListTicketResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListTicketRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ListTicketResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListTicketRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListTicketResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListTicketRequest, callback: grpc.requestCallback<_entpb_ListTicketResponse__Output>): grpc.ClientUnaryCall;
  
  Update(argument: _entpb_UpdateTicketRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateTicketRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateTicketRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateTicketRequest, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateTicketRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateTicketRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateTicketRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateTicketRequest, callback: grpc.requestCallback<_entpb_Ticket__Output>): grpc.ClientUnaryCall;
  
}

export interface TicketServiceHandlers extends grpc.UntypedServiceImplementation {
  BatchCreate: grpc.handleUnaryCall<_entpb_BatchCreateTicketsRequest__Output, _entpb_BatchCreateTicketsResponse>;
  
  Create: grpc.handleUnaryCall<_entpb_CreateTicketRequest__Output, _entpb_Ticket>;
  
  Delete: grpc.handleUnaryCall<_entpb_DeleteTicketRequest__Output, _google_protobuf_Empty>;
  
  Get: grpc.handleUnaryCall<_entpb_GetTicketRequest__Output, _entpb_Ticket>;
  
  List: grpc.handleUnaryCall<_entpb_ListTicketRequest__Output, _entpb_ListTicketResponse>;
  
  Update: grpc.handleUnaryCall<_entpb_UpdateTicketRequest__Output, _entpb_Ticket>;
  
}

export interface TicketServiceDefinition extends grpc.ServiceDefinition {
  BatchCreate: MethodDefinition<_entpb_BatchCreateTicketsRequest, _entpb_BatchCreateTicketsResponse, _entpb_BatchCreateTicketsRequest__Output, _entpb_BatchCreateTicketsResponse__Output>
  Create: MethodDefinition<_entpb_CreateTicketRequest, _entpb_Ticket, _entpb_CreateTicketRequest__Output, _entpb_Ticket__Output>
  Delete: MethodDefinition<_entpb_DeleteTicketRequest, _google_protobuf_Empty, _entpb_DeleteTicketRequest__Output, _google_protobuf_Empty__Output>
  Get: MethodDefinition<_entpb_GetTicketRequest, _entpb_Ticket, _entpb_GetTicketRequest__Output, _entpb_Ticket__Output>
  List: MethodDefinition<_entpb_ListTicketRequest, _entpb_ListTicketResponse, _entpb_ListTicketRequest__Output, _entpb_ListTicketResponse__Output>
  Update: MethodDefinition<_entpb_UpdateTicketRequest, _entpb_Ticket, _entpb_UpdateTicketRequest__Output, _entpb_Ticket__Output>
}
