// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type { BytesValue as _google_protobuf_BytesValue, BytesValue__Output as _google_protobuf_BytesValue__Output } from '../google/protobuf/BytesValue';
import type { Section as _entpb_Section, Section__Output as _entpb_Section__Output } from '../entpb/Section';
import type { Long } from '@grpc/proto-loader';

// Original file: ../ticketing/ent/proto/entpb/entpb.proto

export const _entpb_Ticket_Status = {
  STATUS_UNSPECIFIED: 0,
  STATUS_AVAILABLE: 1,
  STATUS_BOUGHT: 2,
  STATUS_RESERVED: 3,
} as const;

export type _entpb_Ticket_Status =
  | 'STATUS_UNSPECIFIED'
  | 0
  | 'STATUS_AVAILABLE'
  | 1
  | 'STATUS_BOUGHT'
  | 2
  | 'STATUS_RESERVED'
  | 3

export type _entpb_Ticket_Status__Output = typeof _entpb_Ticket_Status[keyof typeof _entpb_Ticket_Status]

export interface Ticket {
  'id'?: (Buffer | Uint8Array | string);
  'seatNumber'?: (number | string | Long);
  'status'?: (_entpb_Ticket_Status);
  'userId'?: (_google_protobuf_BytesValue | null);
  'withinSection'?: (_entpb_Section | null);
}

export interface Ticket__Output {
  'id'?: (Buffer);
  'seatNumber'?: (Long);
  'status'?: (_entpb_Ticket_Status__Output);
  'userId'?: (_google_protobuf_BytesValue__Output);
  'withinSection'?: (_entpb_Section__Output);
}
