import { PrismaClient } from "@prisma/client";
//gRPC imports
import path from "path";
import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import Stripe from "stripe";

//This is the overall Typescript type derived from the proto file, applied when loading package definition
import { ProtoGrpcType } from "../proto/entpb";

export const stripe: Stripe = require("stripe")(process.env.STRIPE_SECRET_KEY!);
export const prisma = new PrismaClient({
    log: ["query", "info", "warn", "error"],
});
// This is your Stripe CLI webhook secret for testing your endpoint locally.
export const endpointSecret = process.env.STRIPE_WEBHOOK_SECRET!;
export const ENVIRONMENT = process.env.ENVIRONMENT;

const PROTO_FILE = "../../proto/entpb.proto";
const ticketingPORT = 5003;

let grpcURL;
if (ENVIRONMENT === "compose") {
    grpcURL = `ticketing:${ticketingPORT}`;
} else if (ENVIRONMENT === "kubernetes") {
    grpcURL = `ticketing-service:${ticketingPORT}`;
} else {
    grpcURL = `0.0.0.0:${ticketingPORT}`;
}

// Now you can use grpcURL in your application

const packageDef = protoLoader.loadSync(path.resolve(__dirname, PROTO_FILE));
const grpcObj = grpc.loadPackageDefinition(
    packageDef
) as unknown as ProtoGrpcType;
export const ticketClient = new grpcObj.entpb.TicketService(
    grpcURL,
    grpc.credentials.createInsecure()
);
export const sectionClient = new grpcObj.entpb.SectionService(
    grpcURL,
    grpc.credentials.createInsecure()
);

export const MAXIMUM_TICKETS_PER_USER = 4;