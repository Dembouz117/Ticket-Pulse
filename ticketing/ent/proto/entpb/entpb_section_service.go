// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	uuid "github.com/google/uuid"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	regexp "regexp"
	strings "strings"
	ent "ticketing/ent"
	concertsession "ticketing/ent/concertsession"
	section "ticketing/ent/section"
	ticket "ticketing/ent/ticket"
)

// SectionService implements SectionServiceServer
type SectionService struct {
	client *ent.Client
	UnimplementedSectionServiceServer
}

// NewSectionService returns a new SectionService
func NewSectionService(client *ent.Client) *SectionService {
	return &SectionService{
		client: client,
	}
}

var protoIdentNormalizeRegexpSection_Category = regexp.MustCompile(`[^a-zA-Z0-9_]+`)

func protoIdentNormalizeSection_Category(e string) string {
	return protoIdentNormalizeRegexpSection_Category.ReplaceAllString(e, "_")
}

func toProtoSection_Category(e section.Category) Section_Category {
	if v, ok := Section_Category_value[strings.ToUpper("CATEGORY_"+protoIdentNormalizeSection_Category(string(e)))]; ok {
		return Section_Category(v)
	}
	return Section_Category(0)
}

func toEntSection_Category(e Section_Category) section.Category {
	if v, ok := Section_Category_name[int32(e)]; ok {
		entVal := map[string]string{
			"CATEGORY_CAT1": "CAT1",
			"CATEGORY_CAT2": "CAT2",
			"CATEGORY_CAT3": "CAT3",
			"CATEGORY_CAT4": "CAT4",
			"CATEGORY_CAT5": "CAT5",
		}[v]
		return section.Category(entVal)
	}
	return ""
}

// toProtoSection transforms the ent type to the pb type
func toProtoSection(e *ent.Section) (*Section, error) {
	v := &Section{}
	bought := int64(e.Bought)
	v.Bought = bought
	capacity := int64(e.Capacity)
	v.Capacity = capacity
	category := toProtoSection_Category(e.Category)
	v.Category = category
	id, err := e.ID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.Id = id
	name := e.Name
	v.Name = name
	price := int64(e.Price)
	v.Price = price
	reserved := int64(e.Reserved)
	v.Reserved = reserved
	if edg := e.Edges.AtConcertSession; edg != nil {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.AtConcertSession = &ConcertSession{
			Id: id,
		}
	}
	for _, edg := range e.Edges.HasTickets {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.HasTickets = append(v.HasTickets, &Ticket{
			Id: id,
		})
	}
	return v, nil
}

// toProtoSectionList transforms a list of ent type to a list of pb type
func toProtoSectionList(e []*ent.Section) ([]*Section, error) {
	var pbList []*Section
	for _, entEntity := range e {
		pbEntity, err := toProtoSection(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements SectionServiceServer.Create
func (svc *SectionService) Create(ctx context.Context, req *CreateSectionRequest) (*Section, error) {
	section := req.GetSection()
	m, err := svc.createBuilder(section)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoSection(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements SectionServiceServer.Get
func (svc *SectionService) Get(ctx context.Context, req *GetSectionRequest) (*Section, error) {
	var (
		err error
		get *ent.Section
	)
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	switch req.GetView() {
	case GetSectionRequest_VIEW_UNSPECIFIED, GetSectionRequest_BASIC:
		get, err = svc.client.Section.Get(ctx, id)
	case GetSectionRequest_WITH_EDGE_IDS:
		get, err = svc.client.Section.Query().
			Where(section.ID(id)).
			WithAtConcertSession(func(query *ent.ConcertSessionQuery) {
				query.Select(concertsession.FieldID)
			}).
			WithHasTickets(func(query *ent.TicketQuery) {
				query.Select(ticket.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoSection(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements SectionServiceServer.Update
func (svc *SectionService) Update(ctx context.Context, req *UpdateSectionRequest) (*Section, error) {
	section := req.GetSection()
	var sectionID uuid.UUID
	if err := (&sectionID).UnmarshalBinary(section.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.Section.UpdateOneID(sectionID)
	sectionBought := int(section.GetBought())
	m.SetBought(sectionBought)
	sectionCapacity := int(section.GetCapacity())
	m.SetCapacity(sectionCapacity)
	sectionCategory := toEntSection_Category(section.GetCategory())
	m.SetCategory(sectionCategory)
	sectionName := section.GetName()
	m.SetName(sectionName)
	sectionPrice := int(section.GetPrice())
	m.SetPrice(sectionPrice)
	sectionReserved := int(section.GetReserved())
	m.SetReserved(sectionReserved)
	if section.GetAtConcertSession() != nil {
		var sectionAtConcertSession uuid.UUID
		if err := (&sectionAtConcertSession).UnmarshalBinary(section.GetAtConcertSession().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetAtConcertSessionID(sectionAtConcertSession)
	}
	for _, item := range section.GetHasTickets() {
		var hastickets uuid.UUID
		if err := (&hastickets).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddHasTicketIDs(hastickets)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoSection(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements SectionServiceServer.Delete
func (svc *SectionService) Delete(ctx context.Context, req *DeleteSectionRequest) (*emptypb.Empty, error) {
	var err error
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err = svc.client.Section.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements SectionServiceServer.List
func (svc *SectionService) List(ctx context.Context, req *ListSectionRequest) (*ListSectionResponse, error) {
	var (
		err      error
		entList  []*ent.Section
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Section.Query().
		Order(ent.Desc(section.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken, err := uuid.ParseBytes(bytes)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		listQuery = listQuery.
			Where(section.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListSectionRequest_VIEW_UNSPECIFIED, ListSectionRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListSectionRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithAtConcertSession(func(query *ent.ConcertSessionQuery) {
				query.Select(concertsession.FieldID)
			}).
			WithHasTickets(func(query *ent.TicketQuery) {
				query.Select(ticket.FieldID)
			}).
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		protoList, err := toProtoSectionList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListSectionResponse{
			SectionList:   protoList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements SectionServiceServer.BatchCreate
func (svc *SectionService) BatchCreate(ctx context.Context, req *BatchCreateSectionsRequest) (*BatchCreateSectionsResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.SectionCreate, len(requests))
	for i, req := range requests {
		section := req.GetSection()
		var err error
		bulk[i], err = svc.createBuilder(section)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Section.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoSectionList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateSectionsResponse{
			Sections: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *SectionService) createBuilder(section *Section) (*ent.SectionCreate, error) {
	m := svc.client.Section.Create()
	sectionBought := int(section.GetBought())
	m.SetBought(sectionBought)
	sectionCapacity := int(section.GetCapacity())
	m.SetCapacity(sectionCapacity)
	sectionCategory := toEntSection_Category(section.GetCategory())
	m.SetCategory(sectionCategory)
	sectionName := section.GetName()
	m.SetName(sectionName)
	sectionPrice := int(section.GetPrice())
	m.SetPrice(sectionPrice)
	sectionReserved := int(section.GetReserved())
	m.SetReserved(sectionReserved)
	if section.GetAtConcertSession() != nil {
		var sectionAtConcertSession uuid.UUID
		if err := (&sectionAtConcertSession).UnmarshalBinary(section.GetAtConcertSession().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetAtConcertSessionID(sectionAtConcertSession)
	}
	for _, item := range section.GetHasTickets() {
		var hastickets uuid.UUID
		if err := (&hastickets).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddHasTicketIDs(hastickets)
	}
	return m, nil
}
