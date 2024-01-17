// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type { Ticket as _entpb_Ticket, Ticket__Output as _entpb_Ticket__Output } from '../entpb/Ticket';

export interface ListTicketResponse {
  'ticketList'?: (_entpb_Ticket)[];
  'nextPageToken'?: (string);
}

export interface ListTicketResponse__Output {
  'ticketList'?: (_entpb_Ticket__Output)[];
  'nextPageToken'?: (string);
}
