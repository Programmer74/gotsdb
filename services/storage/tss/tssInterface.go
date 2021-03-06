package tss

import pb "github.com/nikita-tomilov/gotsdb/proto"

type Void struct{}

type TimeSeriesStorage interface {
	InitStorage()
	CloseStorage()
	Save(dataSource string, data map[string]*pb.TSPoints, expirationMillis uint64)
	SaveBatch(dataSource string, data []*pb.TSPoint, expirationMillis uint64)
	Retrieve(dataSource string, tags []string, fromTimestamp uint64, toTimestamp uint64) map[string]*pb.TSPoints
	Availability(dataSource string, fromTimestamp uint64, toTimestamp uint64) []*pb.TSAvailabilityChunk
	String() string
}
