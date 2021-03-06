package servers

import (
	log "github.com/jeanphorn/log4go"
	pb "github.com/nikita-tomilov/gotsdb/proto"
	"github.com/nikita-tomilov/gotsdb/services/storage"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

type GrpcUserServer struct {
	StorageManager *interface{} `summer:"StorageManager"`
	ListenAddress           string       `summer.property:"grpc.listenAddress|:5300"`
	pb.UnimplementedGoTSDBServer
}

func (s *GrpcUserServer) getStorageManager() *storage.Manager {
	sm := *s.StorageManager
	sm2 := (sm).(storage.Manager)
	return &sm2
}

func (s *GrpcUserServer) BeginListening() {
	log.Warn("Starting to listen at '%s'", s.ListenAddress)
	listener, err := net.Listen("tcp", s.ListenAddress)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterGoTSDBServer(grpcServer, &server{storageManager: *s.getStorageManager()})
	grpcServer.Serve(listener)
}

//TODO: pass errors from storageManager level to grpc level?

type server struct {
	storageManager storage.Manager
	pb.UnimplementedGoTSDBServer
}

func (s *server) KvsSave(c context.Context, req *pb.KvsStoreRequest) (*pb.KvsStoreResponse, error) {
	return s.storageManager.KvsSave(c, req)
}

func (s *server) KvsKeyExists(c context.Context, req *pb.KvsKeyExistsRequest) (*pb.KvsKeyExistsResponse, error) {
	return s.storageManager.KvsKeyExists(c, req)
}

func (s *server) KvsRetrieve(c context.Context, req *pb.KvsRetrieveRequest) (*pb.KvsRetrieveResponse, error) {
	return s.storageManager.KvsRetrieve(c, req)
}

func (s *server) KvsDelete(c context.Context, req *pb.KvsDeleteRequest) (*pb.KvsDeleteResponse, error) {
	return s.storageManager.KvsDelete(c, req)
}

func (s *server) KvsGetKeys(c context.Context, req *pb.KvsAllKeysRequest) (*pb.KvsAllKeysResponse, error) {
	return s.storageManager.KvsGetKeys(c, req)
}

func (s *server) TSSave(c context.Context, req *pb.TSStoreRequest) (*pb.TSStoreResponse, error) {
	return s.storageManager.TSSave(c, req)
}

func (s *server) TSSaveBatch(c context.Context, req *pb.TSStoreBatchRequest) (*pb.TSStoreResponse, error) {
	return s.storageManager.TSSaveBatch(c, req)
}

func (s *server) TSRetrieve(c context.Context, req *pb.TSRetrieveRequest) (*pb.TSRetrieveResponse, error) {
	return s.storageManager.TSRetrieve(c, req)
}

func (s *server) TSAvailability(c context.Context, req *pb.TSAvailabilityRequest) (*pb.TSAvailabilityResponse, error) {
	return s.storageManager.TSAvailability(c, req)
}
