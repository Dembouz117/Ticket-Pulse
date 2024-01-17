// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type { ConcertSession as _entpb_ConcertSession, ConcertSession__Output as _entpb_ConcertSession__Output } from '../entpb/ConcertSession';

export interface Concert {
  'id'?: (Buffer | Uint8Array | string);
  'title'?: (string);
  'artist'?: (string);
  'imageUrl'?: (string);
  'hasConcertSessions'?: (_entpb_ConcertSession)[];
}

export interface Concert__Output {
  'id'?: (Buffer);
  'title'?: (string);
  'artist'?: (string);
  'imageUrl'?: (string);
  'hasConcertSessions'?: (_entpb_ConcertSession__Output)[];
}
