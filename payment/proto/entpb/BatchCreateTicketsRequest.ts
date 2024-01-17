// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type { CreateTicketRequest as _entpb_CreateTicketRequest, CreateTicketRequest__Output as _entpb_CreateTicketRequest__Output } from '../entpb/CreateTicketRequest';

export interface BatchCreateTicketsRequest {
  'requests'?: (_entpb_CreateTicketRequest)[];
}

export interface BatchCreateTicketsRequest__Output {
  'requests'?: (_entpb_CreateTicketRequest__Output)[];
}
