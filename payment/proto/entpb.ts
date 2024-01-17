import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { ConcertServiceClient as _entpb_ConcertServiceClient, ConcertServiceDefinition as _entpb_ConcertServiceDefinition } from './entpb/ConcertService';
import type { ConcertSessionServiceClient as _entpb_ConcertSessionServiceClient, ConcertSessionServiceDefinition as _entpb_ConcertSessionServiceDefinition } from './entpb/ConcertSessionService';
import type { SectionServiceClient as _entpb_SectionServiceClient, SectionServiceDefinition as _entpb_SectionServiceDefinition } from './entpb/SectionService';
import type { TicketServiceClient as _entpb_TicketServiceClient, TicketServiceDefinition as _entpb_TicketServiceDefinition } from './entpb/TicketService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  entpb: {
    BatchCreateConcertSessionsRequest: MessageTypeDefinition
    BatchCreateConcertSessionsResponse: MessageTypeDefinition
    BatchCreateConcertsRequest: MessageTypeDefinition
    BatchCreateConcertsResponse: MessageTypeDefinition
    BatchCreateSectionsRequest: MessageTypeDefinition
    BatchCreateSectionsResponse: MessageTypeDefinition
    BatchCreateTicketsRequest: MessageTypeDefinition
    BatchCreateTicketsResponse: MessageTypeDefinition
    Concert: MessageTypeDefinition
    ConcertService: SubtypeConstructor<typeof grpc.Client, _entpb_ConcertServiceClient> & { service: _entpb_ConcertServiceDefinition }
    ConcertSession: MessageTypeDefinition
    ConcertSessionService: SubtypeConstructor<typeof grpc.Client, _entpb_ConcertSessionServiceClient> & { service: _entpb_ConcertSessionServiceDefinition }
    CreateConcertRequest: MessageTypeDefinition
    CreateConcertSessionRequest: MessageTypeDefinition
    CreateSectionRequest: MessageTypeDefinition
    CreateTicketRequest: MessageTypeDefinition
    DeleteConcertRequest: MessageTypeDefinition
    DeleteConcertSessionRequest: MessageTypeDefinition
    DeleteSectionRequest: MessageTypeDefinition
    DeleteTicketRequest: MessageTypeDefinition
    GetConcertRequest: MessageTypeDefinition
    GetConcertSessionRequest: MessageTypeDefinition
    GetSectionRequest: MessageTypeDefinition
    GetTicketRequest: MessageTypeDefinition
    ListConcertRequest: MessageTypeDefinition
    ListConcertResponse: MessageTypeDefinition
    ListConcertSessionRequest: MessageTypeDefinition
    ListConcertSessionResponse: MessageTypeDefinition
    ListSectionRequest: MessageTypeDefinition
    ListSectionResponse: MessageTypeDefinition
    ListTicketRequest: MessageTypeDefinition
    ListTicketResponse: MessageTypeDefinition
    Section: MessageTypeDefinition
    SectionService: SubtypeConstructor<typeof grpc.Client, _entpb_SectionServiceClient> & { service: _entpb_SectionServiceDefinition }
    Ticket: MessageTypeDefinition
    TicketService: SubtypeConstructor<typeof grpc.Client, _entpb_TicketServiceClient> & { service: _entpb_TicketServiceDefinition }
    UpdateConcertRequest: MessageTypeDefinition
    UpdateConcertSessionRequest: MessageTypeDefinition
    UpdateSectionRequest: MessageTypeDefinition
    UpdateTicketRequest: MessageTypeDefinition
  }
  google: {
    protobuf: {
      BoolValue: MessageTypeDefinition
      BytesValue: MessageTypeDefinition
      DoubleValue: MessageTypeDefinition
      Empty: MessageTypeDefinition
      FloatValue: MessageTypeDefinition
      Int32Value: MessageTypeDefinition
      Int64Value: MessageTypeDefinition
      StringValue: MessageTypeDefinition
      UInt32Value: MessageTypeDefinition
      UInt64Value: MessageTypeDefinition
    }
  }
}

