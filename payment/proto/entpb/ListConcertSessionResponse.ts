// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type { ConcertSession as _entpb_ConcertSession, ConcertSession__Output as _entpb_ConcertSession__Output } from '../entpb/ConcertSession';

export interface ListConcertSessionResponse {
  'concertSessionList'?: (_entpb_ConcertSession)[];
  'nextPageToken'?: (string);
}

export interface ListConcertSessionResponse__Output {
  'concertSessionList'?: (_entpb_ConcertSession__Output)[];
  'nextPageToken'?: (string);
}
