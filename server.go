package main

import (
	log "github.com/jeanphorn/log4go"
	"github.com/nikita-tomilov/gotsdb/services"
	"github.com/nikita-tomilov/gotsdb/services/cluster"
	"github.com/nikita-tomilov/gotsdb/services/servers"
	"github.com/nikita-tomilov/gotsdb/services/storage"
	"github.com/nikita-tomilov/gotsdb/services/storage/kvs"
	"github.com/nikita-tomilov/gotsdb/services/storage/tss"
	"github.com/nikita-tomilov/summer/summer"
	"os"
)

const defaultLogSettingsFile = "./config/log4go.json"
const defaultPropertiesFile = "./config/app.properties"

const propertiesOverrideEnvironmentVariable = "GOTSDB_PROPERTY_FILE"
const logSettingsOverrideEnvironmentVariable = "GOTSDB_LOG_SETTINGS_FILE"

const kvsStorageBeanName = "KeyValueStorage"
const kvsStorageBeanType = "*kvs.KeyValueStorage"
const tssStorageBeanName = "TimeSeriesStorage"
const tssStorageBeanType = "*tss.TimeSeriesStorage"
const applicationBeanName = "Application"
const grpcUserServerBeanName = "GrpcUserServer"
const grpcClusterServerBeanName = "GrpcClusterServer"
const clusterManagerBeanName = "ClusterManager"
const storageManagerBeanName = "StorageManager"

const kvsEnginePropertyKey = "kvs.engine"
const kvsEnginePropertyFileValue = "file"
const kvsEnginePropertyInMemValue = "inmem"

const tssEnginePropertyKey = "tss.engine"
const tssEnginePropertyCSVValue = "csv"
const tssEnginePropertyBCSVValue = "bcsv"
const tssEnginePropertyInMemValue = "inmem"
const tssEnginePropertyLSMValue = "lsm"
const tssEnginePropertySQLiteValue = "sqlite"
const tssEnginePropertyQLValue = "ql"
const tssEnginePropertyBboltValue = "bbolt"

const serverModePropertyKey = "server.mode"
const serverModePropertyClusterValue = "cluster"

func setupDI() {
	propertiesOverride, propertiesOverridePresent:= os.LookupEnv(propertiesOverrideEnvironmentVariable)
	if propertiesOverridePresent {
		summer.ParseProperties(propertiesOverride)
	} else {
		summer.ParseProperties(defaultPropertiesFile)
	}

	summer.RegisterBean(applicationBeanName, services.Application{})
	summer.RegisterBean(grpcUserServerBeanName, servers.GrpcUserServer{})

	mode, _ := summer.GetPropertyValue(serverModePropertyKey)
	if mode == serverModePropertyClusterValue {
		summer.RegisterBean(grpcClusterServerBeanName, cluster.GrpcClusterServer{})
		summer.RegisterBean(clusterManagerBeanName, cluster.Manager{})
		summer.RegisterBean(storageManagerBeanName, cluster.ClusteredStorageManager{})
	} else {
		summer.RegisterBean(storageManagerBeanName, storage.SingleNodeStorageManager{})
	}

	kvsEngine, _ := summer.GetPropertyValue(kvsEnginePropertyKey)
	switch kvsEngine {
	case kvsEnginePropertyFileValue:
		summer.RegisterBeanWithTypeAlias(kvsStorageBeanName, kvs.FileKVS{}, kvsStorageBeanType)
	case kvsEnginePropertyInMemValue:
		summer.RegisterBeanWithTypeAlias(kvsStorageBeanName, kvs.InMemKVS{}, kvsStorageBeanType)
	default:
		panic("unrecognized kv storage engine '" + kvsEngine + "'; exiting")
	}

	tssEngine, _ := summer.GetPropertyValue(tssEnginePropertyKey)
	switch tssEngine {
	case tssEnginePropertyInMemValue:
		summer.RegisterBeanWithTypeAlias(tssStorageBeanName, tss.InMemTSS{}, tssStorageBeanType)
	case tssEnginePropertyCSVValue:
		summer.RegisterBeanWithTypeAlias(tssStorageBeanName, tss.CSVTSS{}, tssStorageBeanType)
	case tssEnginePropertyBCSVValue:
		summer.RegisterBeanWithTypeAlias(tssStorageBeanName, tss.BCSVTSS{}, tssStorageBeanType)
	case tssEnginePropertyLSMValue:
		summer.RegisterBeanWithTypeAlias(tssStorageBeanName, tss.LSMTSS{}, tssStorageBeanType)
	case tssEnginePropertySQLiteValue:
		summer.RegisterBeanWithTypeAlias(tssStorageBeanName, tss.SqliteTSS{}, tssStorageBeanType)
	case tssEnginePropertyQLValue:
		summer.RegisterBeanWithTypeAlias(tssStorageBeanName, tss.QlBasedPersistentTSS{}, tssStorageBeanType)
	case tssEnginePropertyBboltValue:
		summer.RegisterBeanWithTypeAlias(tssStorageBeanName, tss.BboltTSS{}, tssStorageBeanType)
	default:
		panic("unrecognized ts storage engine '" + tssEngine + "'; exiting")
	}

	summer.PerformDependencyInjection()

	summer.PrintDependencyGraphVertex()
}

func main() {
	logSettingsFile := defaultLogSettingsFile
	logSettingsOverride, logSettingsOverridePresent:= os.LookupEnv(logSettingsOverrideEnvironmentVariable)
	if logSettingsOverridePresent {
		logSettingsFile = logSettingsOverride
	}

	log.LoadConfiguration(logSettingsFile)
	defer log.Close()

	setupDI()

	app := (*summer.GetBean("Application")).(*services.Application)
	app.Startup()
}