// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { BatchCreateSectionsRequest as _entpb_BatchCreateSectionsRequest, BatchCreateSectionsRequest__Output as _entpb_BatchCreateSectionsRequest__Output } from '../entpb/BatchCreateSectionsRequest';
import type { BatchCreateSectionsResponse as _entpb_BatchCreateSectionsResponse, BatchCreateSectionsResponse__Output as _entpb_BatchCreateSectionsResponse__Output } from '../entpb/BatchCreateSectionsResponse';
import type { CreateSectionRequest as _entpb_CreateSectionRequest, CreateSectionRequest__Output as _entpb_CreateSectionRequest__Output } from '../entpb/CreateSectionRequest';
import type { DeleteSectionRequest as _entpb_DeleteSectionRequest, DeleteSectionRequest__Output as _entpb_DeleteSectionRequest__Output } from '../entpb/DeleteSectionRequest';
import type { Empty as _google_protobuf_Empty, Empty__Output as _google_protobuf_Empty__Output } from '../google/protobuf/Empty';
import type { GetSectionRequest as _entpb_GetSectionRequest, GetSectionRequest__Output as _entpb_GetSectionRequest__Output } from '../entpb/GetSectionRequest';
import type { ListSectionRequest as _entpb_ListSectionRequest, ListSectionRequest__Output as _entpb_ListSectionRequest__Output } from '../entpb/ListSectionRequest';
import type { ListSectionResponse as _entpb_ListSectionResponse, ListSectionResponse__Output as _entpb_ListSectionResponse__Output } from '../entpb/ListSectionResponse';
import type { Section as _entpb_Section, Section__Output as _entpb_Section__Output } from '../entpb/Section';
import type { UpdateSectionRequest as _entpb_UpdateSectionRequest, UpdateSectionRequest__Output as _entpb_UpdateSectionRequest__Output } from '../entpb/UpdateSectionRequest';

export interface SectionServiceClient extends grpc.Client {
  BatchCreate(argument: _entpb_BatchCreateSectionsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateSectionsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateSectionsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_BatchCreateSectionsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateSectionsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateSectionsResponse__Output>): grpc.ClientUnaryCall;
  BatchCreate(argument: _entpb_BatchCreateSectionsRequest, callback: grpc.requestCallback<_entpb_BatchCreateSectionsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateSectionsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateSectionsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateSectionsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_BatchCreateSectionsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateSectionsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_BatchCreateSectionsResponse__Output>): grpc.ClientUnaryCall;
  batchCreate(argument: _entpb_BatchCreateSectionsRequest, callback: grpc.requestCallback<_entpb_BatchCreateSectionsResponse__Output>): grpc.ClientUnaryCall;
  
  Create(argument: _entpb_CreateSectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateSectionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateSectionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  Create(argument: _entpb_CreateSectionRequest, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateSectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateSectionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateSectionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  create(argument: _entpb_CreateSectionRequest, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  
  Delete(argument: _entpb_DeleteSectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteSectionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteSectionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  Delete(argument: _entpb_DeleteSectionRequest, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteSectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteSectionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteSectionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  delete(argument: _entpb_DeleteSectionRequest, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  Get(argument: _entpb_GetSectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetSectionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetSectionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  Get(argument: _entpb_GetSectionRequest, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetSectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetSectionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetSectionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  get(argument: _entpb_GetSectionRequest, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  
  List(argument: _entpb_ListSectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListSectionResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListSectionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ListSectionResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListSectionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListSectionResponse__Output>): grpc.ClientUnaryCall;
  List(argument: _entpb_ListSectionRequest, callback: grpc.requestCallback<_entpb_ListSectionResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListSectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListSectionResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListSectionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_ListSectionResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListSectionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_ListSectionResponse__Output>): grpc.ClientUnaryCall;
  list(argument: _entpb_ListSectionRequest, callback: grpc.requestCallback<_entpb_ListSectionResponse__Output>): grpc.ClientUnaryCall;
  
  Update(argument: _entpb_UpdateSectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateSectionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateSectionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  Update(argument: _entpb_UpdateSectionRequest, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateSectionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateSectionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateSectionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  update(argument: _entpb_UpdateSectionRequest, callback: grpc.requestCallback<_entpb_Section__Output>): grpc.ClientUnaryCall;
  
}

export interface SectionServiceHandlers extends grpc.UntypedServiceImplementation {
  BatchCreate: grpc.handleUnaryCall<_entpb_BatchCreateSectionsRequest__Output, _entpb_BatchCreateSectionsResponse>;
  
  Create: grpc.handleUnaryCall<_entpb_CreateSectionRequest__Output, _entpb_Section>;
  
  Delete: grpc.handleUnaryCall<_entpb_DeleteSectionRequest__Output, _google_protobuf_Empty>;
  
  Get: grpc.handleUnaryCall<_entpb_GetSectionRequest__Output, _entpb_Section>;
  
  List: grpc.handleUnaryCall<_entpb_ListSectionRequest__Output, _entpb_ListSectionResponse>;
  
  Update: grpc.handleUnaryCall<_entpb_UpdateSectionRequest__Output, _entpb_Section>;
  
}

export interface SectionServiceDefinition extends grpc.ServiceDefinition {
  BatchCreate: MethodDefinition<_entpb_BatchCreateSectionsRequest, _entpb_BatchCreateSectionsResponse, _entpb_BatchCreateSectionsRequest__Output, _entpb_BatchCreateSectionsResponse__Output>
  Create: MethodDefinition<_entpb_CreateSectionRequest, _entpb_Section, _entpb_CreateSectionRequest__Output, _entpb_Section__Output>
  Delete: MethodDefinition<_entpb_DeleteSectionRequest, _google_protobuf_Empty, _entpb_DeleteSectionRequest__Output, _google_protobuf_Empty__Output>
  Get: MethodDefinition<_entpb_GetSectionRequest, _entpb_Section, _entpb_GetSectionRequest__Output, _entpb_Section__Output>
  List: MethodDefinition<_entpb_ListSectionRequest, _entpb_ListSectionResponse, _entpb_ListSectionRequest__Output, _entpb_ListSectionResponse__Output>
  Update: MethodDefinition<_entpb_UpdateSectionRequest, _entpb_Section, _entpb_UpdateSectionRequest__Output, _entpb_Section__Output>
}
