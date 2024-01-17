// Original file: ../ticketing/ent/proto/entpb/entpb.proto


// Original file: ../ticketing/ent/proto/entpb/entpb.proto

export const _entpb_ListSectionRequest_View = {
  VIEW_UNSPECIFIED: 0,
  BASIC: 1,
  WITH_EDGE_IDS: 2,
} as const;

export type _entpb_ListSectionRequest_View =
  | 'VIEW_UNSPECIFIED'
  | 0
  | 'BASIC'
  | 1
  | 'WITH_EDGE_IDS'
  | 2

export type _entpb_ListSectionRequest_View__Output = typeof _entpb_ListSectionRequest_View[keyof typeof _entpb_ListSectionRequest_View]

export interface ListSectionRequest {
  'pageSize'?: (number);
  'pageToken'?: (string);
  'view'?: (_entpb_ListSectionRequest_View);
}

export interface ListSectionRequest__Output {
  'pageSize'?: (number);
  'pageToken'?: (string);
  'view'?: (_entpb_ListSectionRequest_View__Output);
}
