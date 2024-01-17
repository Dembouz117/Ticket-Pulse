// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type { Ticket as _entpb_Ticket, Ticket__Output as _entpb_Ticket__Output } from '../entpb/Ticket';
import type { ConcertSession as _entpb_ConcertSession, ConcertSession__Output as _entpb_ConcertSession__Output } from '../entpb/ConcertSession';
import type { Long } from '@grpc/proto-loader';

// Original file: ../ticketing/ent/proto/entpb/entpb.proto

export const _entpb_Section_Category = {
  CATEGORY_UNSPECIFIED: 0,
  CATEGORY_CAT1: 1,
  CATEGORY_CAT2: 2,
  CATEGORY_CAT3: 3,
  CATEGORY_CAT4: 4,
  CATEGORY_CAT5: 5,
} as const;

export type _entpb_Section_Category =
  | 'CATEGORY_UNSPECIFIED'
  | 0
  | 'CATEGORY_CAT1'
  | 1
  | 'CATEGORY_CAT2'
  | 2
  | 'CATEGORY_CAT3'
  | 3
  | 'CATEGORY_CAT4'
  | 4
  | 'CATEGORY_CAT5'
  | 5

export type _entpb_Section_Category__Output = typeof _entpb_Section_Category[keyof typeof _entpb_Section_Category]

export interface Section {
  'id'?: (Buffer | Uint8Array | string);
  'name'?: (string);
  'capacity'?: (number | string | Long);
  'reserved'?: (number | string | Long);
  'bought'?: (number | string | Long);
  'category'?: (_entpb_Section_Category);
  'price'?: (number | string | Long);
  'hasTickets'?: (_entpb_Ticket)[];
  'atConcertSession'?: (_entpb_ConcertSession | null);
}

export interface Section__Output {
  'id'?: (Buffer);
  'name'?: (string);
  'capacity'?: (Long);
  'reserved'?: (Long);
  'bought'?: (Long);
  'category'?: (_entpb_Section_Category__Output);
  'price'?: (Long);
  'hasTickets'?: (_entpb_Ticket__Output)[];
  'atConcertSession'?: (_entpb_ConcertSession__Output);
}
