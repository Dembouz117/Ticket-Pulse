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
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	regexp "regexp"
	strings "strings"
	ent "ticketing/ent"
	section "ticketing/ent/section"
	ticket "ticketing/ent/ticket"
)

// TicketService implements TicketServiceServer
type TicketService struct {
	client *ent.Client
	UnimplementedTicketServiceServer
}

// NewTicketService returns a new TicketService
func NewTicketService(client *ent.Client) *TicketService {
	return &TicketService{
		client: client,
	}
}

var protoIdentNormalizeRegexpTicket_Status = regexp.MustCompile(`[^a-zA-Z0-9_]+`)

func protoIdentNormalizeTicket_Status(e string) string {
	return protoIdentNormalizeRegexpTicket_Status.ReplaceAllString(e, "_")
}

func toProtoTicket_Status(e ticket.Status) Ticket_Status {
	if v, ok := Ticket_Status_value[strings.ToUpper("STATUS_"+protoIdentNormalizeTicket_Status(string(e)))]; ok {
		return Ticket_Status(v)
	}
	return Ticket_Status(0)
}

func toEntTicket_Status(e Ticket_Status) ticket.Status {
	if v, ok := Ticket_Status_name[int32(e)]; ok {
		entVal := map[string]string{
			"STATUS_AVAILABLE": "AVAILABLE",
			"STATUS_BOUGHT":    "BOUGHT",
			"STATUS_RESERVED":  "RESERVED",
		}[v]
		return ticket.Status(entVal)
	}
	return ""
}

// toProtoTicket transforms the ent type to the pb type
func toProtoTicket(e *ent.Ticket) (*Ticket, error) {
	v := &Ticket{}
	id, err := e.ID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.Id = id
	seatNumber := int64(e.SeatNumber)
	v.SeatNumber = seatNumber
	status := toProtoTicket_Status(e.Status)
	v.Status = status
	byteSlice := e.UserId[:]
	v.UserId = &wrapperspb.BytesValue{Value: byteSlice}
	if edg := e.Edges.WithinSection; edg != nil {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.WithinSection = &Section{
			Id: id,
		}
	}
	return v, nil
}

// toProtoTicketList transforms a list of ent type to a list of pb type
func toProtoTicketList(e []*ent.Ticket) ([]*Ticket, error) {
	var pbList []*Ticket
	for _, entEntity := range e {
		pbEntity, err := toProtoTicket(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements TicketServiceServer.Create
func (svc *TicketService) Create(ctx context.Context, req *CreateTicketRequest) (*Ticket, error) {
	ticket := req.GetTicket()
	m, err := svc.createBuilder(ticket)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoTicket(res)
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

// Get implements TicketServiceServer.Get
func (svc *TicketService) Get(ctx context.Context, req *GetTicketRequest) (*Ticket, error) {
	var (
		err error
		get *ent.Ticket
	)
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	switch req.GetView() {
	case GetTicketRequest_VIEW_UNSPECIFIED, GetTicketRequest_BASIC:
		get, err = svc.client.Ticket.Get(ctx, id)
	case GetTicketRequest_WITH_EDGE_IDS:
		get, err = svc.client.Ticket.Query().
			Where(ticket.ID(id)).
			WithWithinSection(func(query *ent.SectionQuery) {
				query.Select(section.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoTicket(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements TicketServiceServer.Update
func (svc *TicketService) Update(ctx context.Context, req *UpdateTicketRequest) (*Ticket, error) {
	ticket := req.GetTicket()
	var ticketID uuid.UUID
	if err := (&ticketID).UnmarshalBinary(ticket.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.Ticket.UpdateOneID(ticketID)
	ticketSeatNumber := int(ticket.GetSeatNumber())
	m.SetSeatNumber(ticketSeatNumber)
	ticketStatus := toEntTicket_Status(ticket.GetStatus())
	m.SetStatus(ticketStatus)
	if ticket.GetUserId() != nil {
		var ticketUserId uuid.UUID
		if err := (&ticketUserId).UnmarshalBinary(ticket.GetUserId().GetValue()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetUserId(ticketUserId)
	}
	if ticket.GetWithinSection() != nil {
		var ticketWithinSection uuid.UUID
		if err := (&ticketWithinSection).UnmarshalBinary(ticket.GetWithinSection().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetWithinSectionID(ticketWithinSection)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoTicket(res)
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

// Delete implements TicketServiceServer.Delete
func (svc *TicketService) Delete(ctx context.Context, req *DeleteTicketRequest) (*emptypb.Empty, error) {
	var err error
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err = svc.client.Ticket.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements TicketServiceServer.List
func (svc *TicketService) List(ctx context.Context, req *ListTicketRequest) (*ListTicketResponse, error) {
	var (
		err      error
		entList  []*ent.Ticket
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Ticket.Query().
		Order(ent.Desc(ticket.FieldID)).
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
			Where(ticket.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListTicketRequest_VIEW_UNSPECIFIED, ListTicketRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListTicketRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithWithinSection(func(query *ent.SectionQuery) {
				query.Select(section.FieldID)
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
		protoList, err := toProtoTicketList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListTicketResponse{
			TicketList:    protoList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements TicketServiceServer.BatchCreate
func (svc *TicketService) BatchCreate(ctx context.Context, req *BatchCreateTicketsRequest) (*BatchCreateTicketsResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.TicketCreate, len(requests))
	for i, req := range requests {
		ticket := req.GetTicket()
		var err error
		bulk[i], err = svc.createBuilder(ticket)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Ticket.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoTicketList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateTicketsResponse{
			Tickets: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *TicketService) createBuilder(ticket *Ticket) (*ent.TicketCreate, error) {
	m := svc.client.Ticket.Create()
	ticketSeatNumber := int(ticket.GetSeatNumber())
	m.SetSeatNumber(ticketSeatNumber)
	ticketStatus := toEntTicket_Status(ticket.GetStatus())
	m.SetStatus(ticketStatus)
	if ticket.GetUserId() != nil {
		var ticketUserId uuid.UUID
		if err := (&ticketUserId).UnmarshalBinary(ticket.GetUserId().GetValue()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetUserId(ticketUserId)
	}
	if ticket.GetWithinSection() != nil {
		var ticketWithinSection uuid.UUID
		if err := (&ticketWithinSection).UnmarshalBinary(ticket.GetWithinSection().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetWithinSectionID(ticketWithinSection)
	}
	return m, nil
}
