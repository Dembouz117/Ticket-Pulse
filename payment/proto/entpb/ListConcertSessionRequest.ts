// Original file: ../ticketing/ent/proto/entpb/entpb.proto


// Original file: ../ticketing/ent/proto/entpb/entpb.proto

export const _entpb_ListConcertSessionRequest_View = {
  VIEW_UNSPECIFIED: 0,
  BASIC: 1,
  WITH_EDGE_IDS: 2,
} as const;

export type _entpb_ListConcertSessionRequest_View =
  | 'VIEW_UNSPECIFIED'
  | 0
  | 'BASIC'
  | 1
  | 'WITH_EDGE_IDS'
  | 2

export type _entpb_ListConcertSessionRequest_View__Output = typeof _entpb_ListConcertSessionRequest_View[keyof typeof _entpb_ListConcertSessionRequest_View]

export interface ListConcertSessionRequest {
  'pageSize'?: (number);
  'pageToken'?: (string);
  'view'?: (_entpb_ListConcertSessionRequest_View);
}

export interface ListConcertSessionRequest__Output {
  'pageSize'?: (number);
  'pageToken'?: (string);
  'view'?: (_entpb_ListConcertSessionRequest_View__Output);
}
