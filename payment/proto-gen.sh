#!/bin/bash
pnpm proto-loader-gen-types --grpcLib=@grpc/grpc-js --outDir=proto/ ../ticketing/ent/proto/entpb/*.proto