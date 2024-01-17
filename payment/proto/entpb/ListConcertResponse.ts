// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type { Concert as _entpb_Concert, Concert__Output as _entpb_Concert__Output } from '../entpb/Concert';

export interface ListConcertResponse {
  'concertList'?: (_entpb_Concert)[];
  'nextPageToken'?: (string);
}

export interface ListConcertResponse__Output {
  'concertList'?: (_entpb_Concert__Output)[];
  'nextPageToken'?: (string);
}
