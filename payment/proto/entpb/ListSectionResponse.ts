// Original file: ../ticketing/ent/proto/entpb/entpb.proto

import type { Section as _entpb_Section, Section__Output as _entpb_Section__Output } from '../entpb/Section';

export interface ListSectionResponse {
  'sectionList'?: (_entpb_Section)[];
  'nextPageToken'?: (string);
}

export interface ListSectionResponse__Output {
  'sectionList'?: (_entpb_Section__Output)[];
  'nextPageToken'?: (string);
}
