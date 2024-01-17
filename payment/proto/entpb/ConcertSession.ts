// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type { Concert as _entpb_Concert, Concert__Output as _entpb_Concert__Output } from '../entpb/Concert';
import type { Section as _entpb_Section, Section__Output as _entpb_Section__Output } from '../entpb/Section';
import type { Long } from '@grpc/proto-loader';

export interface ConcertSession {
  'id'?: (Buffer | Uint8Array | string);
  'sessionDateTime'?: (number | string | Long);
  'ofConcert'?: (_entpb_Concert)[];
  'hasSections'?: (_entpb_Section)[];
}

export interface ConcertSession__Output {
  'id'?: (Buffer);
  'sessionDateTime'?: (Long);
  'ofConcert'?: (_entpb_Concert__Output)[];
  'hasSections'?: (_entpb_Section__Output)[];
}
