/*

Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package server

import (
	"strings"

	"golang.org/x/net/context"

	pb "github.com/google/pubkeystore/api"
)

// Server is GRPC server.
type Server struct {
	staticKeys []*pb.SSHKey
}

// NewServer creates a new server instance.
func NewServer(keys []string) Server {
	parsedKeys := make([]*pb.SSHKey, len(keys))
	for i, k := range keys {
		parsedKeys[i] = parseKey(k)
	}
	return Server{staticKeys: parsedKeys}
}

func parseKey(key string) *pb.SSHKey {
	parts := strings.SplitN(key, " ", 3)
	return &pb.SSHKey{
		Name:    "static-key",
		Algo:    parts[0],
		Pubkey:  parts[1],
		Comment: parts[2],
	}
}

// GetKeys is GRPC handler for GetKeys API.
func (s Server) GetKeys(context.Context, *pb.GetKeysRequest) (*pb.GetKeysReply, error) {
	repl := &pb.GetKeysReply{}
	repl.Keys = s.staticKeys

	return repl, nil
}
