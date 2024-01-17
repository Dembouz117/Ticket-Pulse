//Seats page and seat mapping
export interface sectionSVGInterface{
    dataName:string,
    pathShape:string,
    pathCircle:string,
    pathLetter:string
}

export interface sectionMappingInterface extends sectionSVGInterface{
    sectionUuid: string,
    price:number
}

//for backend Section
export interface sectionInterface{
    capacity: number,
    category: string,
    edges: {},
    id: string,
    name: string,
    price: number
}


export interface cartUIType{
    section: String,
    quantity: number
}

export interface userSchema{
    userId: string
}

export interface sectionsSchema{
    id: string,
    name:  string,
    capacity: number,
    edges:sectionEdge
}

export interface sectionEdge{
    id: string,
    sessionDateTime: string,
    edges:{}
}



export interface sectionAlgo{
    id: number,
    section: string,
    seats: number,
    category: string,
    seatsAvailable: number
}

export interface TicketInterface {
    id: string;
    status: string;
  }
  
export interface TicketDataInterface {
    [sectionKey: string]: TicketInterface[];
  }
  

export interface cartEntryType{
    sectionName: String,
    quantity: number,
    price: number,
    sectionId: String
}

export interface showStripeCheckout{
    visible: boolean,
    amount: number
}

export interface seatCategoryProps {
    CAT1: boolean,
    CAT2: boolean,
    CAT3: boolean
}

//ima merge this with the showCartAtom later
export interface cartCoordinateInterface{
    x: number,
    y: number
}

export interface mockTicketsInterface{
    ticketId: string,
}

export interface ticketResponseInterface{
    ticketId: string,
    seatNumber: number,
    sectionId: string,
    sectionCategory: string,
    sectionName: string,
}

export interface ticketSchemaInterface{
    id: string,
    seatNumber: number,
    status: string,
    userId: string,
    edges: any
}

export interface concertInterface{
    id: string,
    artist: string,
    title: string,
    imageUrl: string,
    edges: any
}