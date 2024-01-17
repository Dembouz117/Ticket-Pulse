grpc_tools_node_protoc \
    --proto_path=../ticketing/ent/proto/entpb/ \
    --js_out=import_style=commonjs,binary:proto \
    --grpc_out=grpc_js:proto \
    --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` \
    --generate_package_definition \
    entpb.proto
