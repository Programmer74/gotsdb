package tss

import (
	"fmt"
	"github.com/nikita-tomilov/gotsdb/utils"
	"time"
)

func BuildStoragesForTesting() []TimeSeriesStorage {
	inMem := buildInMemStorage()
	qL := buildQlStorage()
	lsm := buildLSMStorage()
	sQ := buildSqliteStorage()
	return toArray(inMem, qL, lsm, sQ)
}

func BuildStoragesForBenchmark() []TimeSeriesStorage {
	inMem := buildInMemStorage()
	qL := buildQlStorage()
	lsm := buildLSMStorageForBenchmark()
	sQ := buildSqliteStorage()
	return toArray(inMem, qL, lsm, sQ)
	//return toArray(lsm, sQ)
	//return toArray(lsm)
}

func toArray(items ...TimeSeriesStorage) []TimeSeriesStorage {
	return items
}

func buildInMemStorage() *InMemTSS {
	s := InMemTSS{periodBetweenWipes: time.Second * 1}
	s.InitStorage()
	return &s
}

func buildQlStorage() *QlBasedPersistentTSS {
	idx += 1
	s := QlBasedPersistentTSS{Path: fmt.Sprintf("/tmp/gotsdb_test/test%d%d", utils.GetNowMillis(), idx), periodBetweenWipes: time.Second * 1}
	s.InitStorage()
	return &s
}

func buildLSMStorage() *LSMTSS {
	idx += 1
	s := LSMTSS{Path: fmt.Sprintf("/tmp/gotsdb_test/test%d%d", utils.GetNowMillis(), idx), CommitlogFlushPeriodSeconds: 1, CommitlogMaxEntries: 10, MemtExpirationPeriodSeconds: 1, MemtMaxEntriesPerTag: 100}
	s.InitStorage()
	return &s
}

func buildLSMStorageForBenchmark() *LSMTSS {
	idx += 1
	s := LSMTSS{Path: fmt.Sprintf("/tmp/gotsdb_test/test%d%d", utils.GetNowMillis(), idx), CommitlogFlushPeriodSeconds: 2, CommitlogMaxEntries: 30, MemtExpirationPeriodSeconds: 5, MemtMaxEntriesPerTag: 100}
	s.InitStorage()
	return &s
}

func buildSqliteStorage() *SqliteTSS {
	idx += 1
	s := SqliteTSS{Path: fmt.Sprintf("/tmp/gotsdb_test/test%d%d", utils.GetNowMillis(), idx), periodBetweenWipes: time.Second * 1}
	s.InitStorage()
	return &s
}

var idx = 0
