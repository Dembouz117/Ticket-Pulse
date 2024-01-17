// Original file: ../ticketing/ent/proto/entpb/entpb.proto


// Original file: ../ticketing/ent/proto/entpb/entpb.proto

export const _entpb_GetConcertRequest_View = {
  VIEW_UNSPECIFIED: 0,
  BASIC: 1,
  WITH_EDGE_IDS: 2,
} as const;

export type _entpb_GetConcertRequest_View =
  | 'VIEW_UNSPECIFIED'
  | 0
  | 'BASIC'
  | 1
  | 'WITH_EDGE_IDS'
  | 2

export type _entpb_GetConcertRequest_View__Output = typeof _entpb_GetConcertRequest_View[keyof typeof _entpb_GetConcertRequest_View]

export interface GetConcertRequest {
  'id'?: (Buffer | Uint8Array | string);
  'view'?: (_entpb_GetConcertRequest_View);
}

export interface GetConcertRequest__Output {
  'id'?: (Buffer);
  'view'?: (_entpb_GetConcertRequest_View__Output);
}
