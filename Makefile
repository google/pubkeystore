api/pubkeystore.pb.go: proto/pubkeystore.proto
	mkdir -p api
	cd proto && protoc -I/usr/local/include -I. \
	 -I${GOPATH}/src \
	 -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	 --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:../api \
	 pubkeystore.proto

clean:
	rm -rf api/*.pb.go

.PHONY: clean
