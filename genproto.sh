#!/bin/bash
protodir="./protos"
gendir="./protos/gen"
mkdir -p $gendir
protoc --go_out=$gendir --go-grpc_out=$gendir -I $protodir $protodir/*.proto
