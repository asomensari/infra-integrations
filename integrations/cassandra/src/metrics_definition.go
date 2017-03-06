package main

import "github.com/newrelic/infra-integrations-sdk/metric"

var commonDefinition = map[string][]interface{}{
	"software.version":   {"org.apache.cassandra.db:type=StorageService,attr=ReleaseVersion", metric.ATTRIBUTE},
	"cluster.name":       {"org.apache.cassandra.db:type=StorageService,attr=ClusterName", metric.ATTRIBUTE},
	"cluster.datacenter": {"org.apache.cassandra.db:type=EndpointSnitchInfo,attr=Datacenter", metric.ATTRIBUTE},
	"cluster.rack":       {"org.apache.cassandra.db:type=EndpointSnitchInfo,attr=Rack", metric.ATTRIBUTE},
}

// All metrics we want to provide for the cassandra integration
var metricsDefinition = map[string][]interface{}{
	"query.viewWriteRequestsPerSecond":              {"org.apache.cassandra.metrics:type=ClientRequest,scope=ViewWrite,name=Latency,attr=OneMinuteRate", metric.GAUGE},
	"query.rangeSliceRequestsPerSecond":             {"org.apache.cassandra.metrics:type=ClientRequest,scope=RangeSlice,name=Latency,attr=OneMinuteRate", metric.GAUGE},
	"query.CASWriteRequestsPerSecond":               {"org.apache.cassandra.metrics:type=ClientRequest,scope=CASWrite,name=Latency,attr=OneMinuteRate", metric.GAUGE},
	"query.CASReadRequestsPerSecond":                {"org.apache.cassandra.metrics:type=ClientRequest,scope=CASRead,name=Latency,attr=OneMinuteRate", metric.GAUGE},
	"query.readRequestsPerSecond":                   {"org.apache.cassandra.metrics:type=ClientRequest,scope=Read,name=Latency,attr=OneMinuteRate", metric.GAUGE},
	"query.writeRequestsPerSecond":                  {"org.apache.cassandra.metrics:type=ClientRequest,scope=Write,name=Latency,attr=OneMinuteRate", metric.GAUGE},
	"query.writeLatency98thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=ClientRequest,scope=Write,name=Latency,attr=98thPercentile", metric.GAUGE},
	"query.writeLatency99thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=ClientRequest,scope=Write,name=Latency,attr=99thPercentile", metric.GAUGE},
	"query.writeLatency999thPercentileMilliseconds": {"org.apache.cassandra.metrics:type=ClientRequest,scope=Write,name=Latency,attr=999thPercentile", metric.GAUGE},
	"query.writeLatency50thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=ClientRequest,scope=Write,name=Latency,attr=50thPercentile", metric.GAUGE},
	"query.writeLatency75thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=ClientRequest,scope=Write,name=Latency,attr=75thPercentile", metric.GAUGE},
	"query.writeLatency95thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=ClientRequest,scope=Write,name=Latency,attr=95thPercentile", metric.GAUGE},
	"query.readLatency98thPercentileMilliseconds":   {"org.apache.cassandra.metrics:type=ClientRequest,scope=Read,name=Latency,attr=98thPercentile", metric.GAUGE},
	"query.readLatency99thPercentileMilliseconds":   {"org.apache.cassandra.metrics:type=ClientRequest,scope=Read,name=Latency,attr=99thPercentile", metric.GAUGE},
	"query.readLatency999thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=ClientRequest,scope=Read,name=Latency,attr=999thPercentile", metric.GAUGE},
	"query.readLatency50thPercentileMilliseconds":   {"org.apache.cassandra.metrics:type=ClientRequest,scope=Read,name=Latency,attr=50thPercentile", metric.GAUGE},
	"query.readLatency75thPercentileMilliseconds":   {"org.apache.cassandra.metrics:type=ClientRequest,scope=Read,name=Latency,attr=75thPercentile", metric.GAUGE},
	"query.readLatency95thPercentileMilliseconds":   {"org.apache.cassandra.metrics:type=ClientRequest,scope=Read,name=Latency,attr=95thPercentile", metric.GAUGE},

	"query.readTimeoutsPerSecond":           {"org.apache.cassandra.metrics:type=ClientRequest,scope=Read,name=Timeouts,attr=OneMinuteRate", metric.GAUGE},
	"query.readUnavailablesPerSecond":       {"org.apache.cassandra.metrics:type=ClientRequest,scope=Read,name=Unavailables,attr=OneMinuteRate", metric.GAUGE},
	"query.writeTimeoutsPerSecond":          {"org.apache.cassandra.metrics:type=ClientRequest,scope=Write,name=Timeouts,attr=OneMinuteRate", metric.RATE},
	"query.writeUnavailablesPerSecond":      {"org.apache.cassandra.metrics:type=ClientRequest,scope=Write,name=Unavailables,attr=OneMinuteRate", metric.GAUGE},
	"query.rangeSliceTimeoutsPerSecond":     {"org.apache.cassandra.metrics:type=ClientRequest,scope=RangeSlice,name=Timeouts,attr=OneMinuteRate", metric.GAUGE},
	"query.rangeSliceUnavailablesPerSecond": {"org.apache.cassandra.metrics:type=ClientRequest,scope=RangeSlice,name=Unavailables,attr=OneMinuteRate", metric.GAUGE},

	"db.threadpool.requestCounterMutationStageActiveTasks":       {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=CounterMutationStage,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.requestCounterMutationStagePendingTasks":      {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=CounterMutationStage,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.requestMutationStageActiveTasks":              {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=MutationStage,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.requestMutationStagePendingTasks":             {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=MutationStage,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.requestReadRepairStageActiveTasks":            {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=ReadRepairStage,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.requestReadRepairStagePendingTasks":           {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=ReadRepairStage,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.requestReadStageActiveTasks":                  {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=ReadStage,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.requestReadStagePendingTasks":                 {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=ReadStage,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.requestRequestResponseStageActiveTasks":       {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=RequestResponseStage,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.requestRequestResponseStagePendingTasks":      {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=RequestResponseStage,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.requestViewMutationStageActiveTasks":          {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=ViewMutationStage,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.requestViewMutationStagePendingTasks":         {"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=ViewMutationStage,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalAntiEntropyStageActiveTasks":          {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=AntiEntropyStage,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalAntiEntropyStagePendingTasks":         {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=AntiEntropyStage,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalCacheCleanupExecutorActiveTasks":      {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=CacheCleanupExecutor,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalCacheCleanupExecutorPendingTasks":     {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=CacheCleanupExecutor,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalCompactionExecutorActiveTasks":        {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=CompactionExecutor,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalCompactionExecutorPendingTasks":       {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=CompactionExecutor,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalGossipStageActiveTasks":               {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=GossipStage,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalGossipStagePendingTasks":              {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=GossipStage,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalHintsDispatcherActiveTasks":           {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=HintsDispatcher,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalHintsDispatcherPendingTasks":          {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=HintsDispatcher,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalInternalResponseStageActiveTasks":     {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=InternalResponseStage,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalInternalResponseStagePendingTasks":    {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=InternalResponseStage,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalMemtableFlushWriterActiveTasks":       {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=MemtableFlushWriter,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalMemtableFlushWriterPendingTasks":      {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=MemtableFlushWriter,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalMemtablePostFlushActiveTasks":         {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=MemtablePostFlush,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalMemtablePostFlushPendingTasks":        {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=MemtablePostFlush,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalMemtableReclaimMemoryActiveTasks":     {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=MemtableReclaimMemory,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalMemtableReclaimMemoryPendingTasks":    {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=MemtableReclaimMemory,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalMigrationStageActiveTasks":            {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=MigrationStage,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalMigrationStagePendingTasks":           {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=MigrationStage,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalMiscStageActiveTasks":                 {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=MiscStage,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalMiscStagePendingTasks":                {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=MiscStage,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalPendingRangeCalculatorActiveTasks":    {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=PendingRangeCalculator,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalPendingRangeCalculatorPendingTasks":   {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=PendingRangeCalculator,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalSamplerActiveTasks":                   {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=Sampler,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalSamplerPendingTasks":                  {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=Sampler,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalSecondaryIndexManagementActiveTasks":  {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=SecondaryIndexManagement,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalSecondaryIndexManagementPendingTasks": {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=SecondaryIndexManagement,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalValidationExecutorActiveTasks":        {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=ValidationExecutor,name=ActiveTasks,attr=Value", metric.GAUGE},
	"db.threadpool.internalValidationExecutorPendingTasks":       {"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=ValidationExecutor,name=PendingTasks,attr=Value", metric.GAUGE},

	"db.droppedBatchRemoveMessagesPerSecond":     {"org.apache.cassandra.metrics:type=DroppedMessage,scope=BATCH_REMOVE,name=Dropped,attr=Count", metric.RATE},
	"db.droppedBatchStoreMessagesPerSecond":      {"org.apache.cassandra.metrics:type=DroppedMessage,scope=BATCH_STORE,name=Dropped,attr=Count", metric.RATE},
	"db.droppedCounterMutationMessagesPerSecond": {"org.apache.cassandra.metrics:type=DroppedMessage,scope=COUNTER_MUTATION,name=Dropped,attr=Count", metric.RATE},
	"db.droppedHintMessagesPerSecond":            {"org.apache.cassandra.metrics:type=DroppedMessage,scope=HINT,name=Dropped,attr=Count", metric.RATE},
	"db.droppedMutationMessagesPerSecond":        {"org.apache.cassandra.metrics:type=DroppedMessage,scope=MUTATION,name=Dropped,attr=Count", metric.RATE},
	"db.droppedPagedRangeMessagesPerSecond":      {"org.apache.cassandra.metrics:type=DroppedMessage,scope=PAGED_RANGE,name=Dropped,attr=Count", metric.RATE},
	"db.droppedRangeSliceMessagesPerSecond":      {"org.apache.cassandra.metrics:type=DroppedMessage,scope=RANGE_SLICE,name=Dropped,attr=Count", metric.RATE},
	"db.droppedReadMessagesPerSecond":            {"org.apache.cassandra.metrics:type=DroppedMessage,scope=READ,name=Dropped,attr=Count", metric.RATE},
	"db.droppedReadRepairMessagesPerSecond":      {"org.apache.cassandra.metrics:type=DroppedMessage,scope=READ_REPAIR,name=Dropped,attr=Count", metric.RATE},
	"db.droppedRequestResponseMessagesPerSecond": {"org.apache.cassandra.metrics:type=DroppedMessage,scope=REQUEST_RESPONSE,name=Dropped,attr=Count", metric.RATE},
	"db.droppedTraceMessagesPerSecond":           {"org.apache.cassandra.metrics:type=DroppedMessage,scope=_TRACE,name=Dropped,attr=Count", metric.RATE},

	"db.liveSSTableCount":             {"org.apache.cassandra.metrics:type=Table,name=LiveSSTableCount,attr=Value", metric.GAUGE},
	"db.allMemtablesOnHeapSizeBytes":  {"org.apache.cassandra.metrics:type=Table,name=AllMemtablesHeapSize,attr=Value", metric.GAUGE},
	"db.allMemtablesOffHeapSizeBytes": {"org.apache.cassandra.metrics:type=Table,name=AllMemtablesOffHeapSize,attr=Value", metric.GAUGE},

	"db.loadBytes":            {"org.apache.cassandra.metrics:type=Storage,name=Load,attr=Count", metric.GAUGE},
	"db.totalHintsPerSecond":  {"org.apache.cassandra.metrics:type=Storage,name=TotalHints,attr=Count", metric.RATE},
	"db.totalHintsInProgress": {"org.apache.cassandra.metrics:type=Storage,name=TotalHintsInProgress,attr=Count", metric.GAUGE},

	"db.keyCacheCapacityBytes":     {"org.apache.cassandra.metrics:type=Cache,scope=KeyCache,name=Capacity,attr=Value", metric.GAUGE},
	"db.keyCacheHitsPerSecond":     {"org.apache.cassandra.metrics:type=Cache,scope=KeyCache,name=Hits,attr=OneMinuteRate", metric.GAUGE},
	"db.keyCacheHitRate":           {"org.apache.cassandra.metrics:type=Cache,scope=KeyCache,name=OneMinuteHitRate,attr=Value", metric.GAUGE},
	"db.keyCacheRequestsPerSecond": {"org.apache.cassandra.metrics:type=Cache,scope=KeyCache,name=Requests,attr=OneMinuteRate", metric.GAUGE},
	"db.keyCacheSizeBytes":         {"org.apache.cassandra.metrics:type=Cache,scope=KeyCache,name=Size,attr=Value", metric.GAUGE},
	"db.rowCacheCapacityBytes":     {"org.apache.cassandra.metrics:type=Cache,scope=RowCache,name=Capacity,attr=Value", metric.GAUGE},
	"db.rowCacheHitsPerSecond":     {"org.apache.cassandra.metrics:type=Cache,scope=RowCache,name=Hits,attr=OneMinuteRate", metric.GAUGE},
	"db.rowCacheHitRate":           {"org.apache.cassandra.metrics:type=Cache,scope=RowCache,name=OneMinuteHitRate,attr=Value", metric.GAUGE},
	"db.rowCacheRequestsPerSecond": {"org.apache.cassandra.metrics:type=Cache,scope=RowCache,name=Requests,attr=OneMinuteRate", metric.GAUGE},
	"db.rowCacheSizeBytes":         {"org.apache.cassandra.metrics:type=Cache,scope=RowCache,name=Size,attr=Value", metric.GAUGE},

	"db.commitLogCompletedTasksPerSecond": {"org.apache.cassandra.metrics:type=CommitLog,name=CompletedTasks,attr=Value", metric.RATE},
	"db.commitLogPendindTasks":            {"org.apache.cassandra.metrics:type=CommitLog,name=PendingTasks,attr=Value", metric.GAUGE},
	"db.commitLogTotalSizeBytes":          {"org.apache.cassandra.metrics:type=CommitLog,name=TotalCommitLogSize,attr=Value", metric.GAUGE},
}

var columnFamilyDefinition = map[string][]interface{}{
	"db.liveSSTableCount":                           {"org.apache.cassandra.metrics:type=Table,name=LiveSSTableCount,attr=Value", metric.GAUGE},
	"db.pendingCompactions":                         {"org.apache.cassandra.metrics:type=Table,name=PendingCompactions,attr=Value", metric.GAUGE},
	"db.liveDiskSpaceUsedBytes":                     {"org.apache.cassandra.metrics:type=Table,name=LiveDiskSpaceUsed,attr=Count", metric.GAUGE},
	"db.SSTablesPerRead50thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=Table,name=SSTablesPerReadHistogram,attr=50thPercentile", metric.GAUGE},
	"db.SSTablesPerRead75thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=Table,name=SSTablesPerReadHistogram,attr=75thPercentile", metric.GAUGE},
	"db.SSTablesPerRead95thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=Table,name=SSTablesPerReadHistogram,attr=95thPercentile", metric.GAUGE},
	"db.SSTablesPerRead98thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=Table,name=SSTablesPerReadHistogram,attr=98thPercentile", metric.GAUGE},
	"db.SSTablesPerRead99thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=Table,name=SSTablesPerReadHistogram,attr=99thPercentile", metric.GAUGE},
	"db.SSTablesPerRead999thPercentileMilliseconds": {"org.apache.cassandra.metrics:type=Table,name=SSTablesPerReadHistogram,attr=999thPercentile", metric.GAUGE},
	"query.writeRequestsPerSecond":                  {"org.apache.cassandra.metrics:type=Table,name=WriteLatency,attr=OneMinuteRate", metric.GAUGE},
	"query.writeLatency50thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=Table,name=WriteLatency,attr=50thPercentile", metric.GAUGE},
	"query.writeLatency75thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=Table,name=WriteLatency,attr=75thPercentile", metric.GAUGE},
	"query.writeLatency95thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=Table,name=WriteLatency,attr=95thPercentile", metric.GAUGE},
	"query.writeLatency98thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=Table,name=WriteLatency,attr=98thPercentile", metric.GAUGE},
	"query.writeLatency99thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=Table,name=WriteLatency,attr=99thPercentile", metric.GAUGE},
	"query.writeLatency999thPercentileMilliseconds": {"org.apache.cassandra.metrics:type=Table,name=WriteLatency,attr=999thPercentile", metric.GAUGE},
	"query.readRequestsPerSecond":                   {"org.apache.cassandra.metrics:type=Table,name=ReadLatency,attr=OneMinuteRate", metric.GAUGE},
	"query.readLatency50thPercentileMilliseconds":   {"org.apache.cassandra.metrics:type=Table,name=ReadLatency,attr=50thPercentile", metric.GAUGE},
	"query.readLatency75thPercentileMilliseconds":   {"org.apache.cassandra.metrics:type=Table,name=ReadLatency,attr=75thPercentile", metric.GAUGE},
	"query.readLatency95thPercentileMilliseconds":   {"org.apache.cassandra.metrics:type=Table,name=ReadLatency,attr=95thPercentile", metric.GAUGE},
	"query.readLatency98thPercentileMilliseconds":   {"org.apache.cassandra.metrics:type=Table,name=ReadLatency,attr=98thPercentile", metric.GAUGE},
	"query.readLatency99thPercentileMilliseconds":   {"org.apache.cassandra.metrics:type=Table,name=ReadLatency,attr=99thPercentile", metric.GAUGE},
	"query.readLatency999thPercentileMilliseconds":  {"org.apache.cassandra.metrics:type=Table,name=ReadLatency,attr=999thPercentile", metric.GAUGE},
	"db.allMemtablesOnHeapSizeBytes":                {"org.apache.cassandra.metrics:type=Table,name=AllMemtablesHeapSize,attr=Value", metric.GAUGE},
	"db.allMemtablesOffHeapSizeBytes":               {"org.apache.cassandra.metrics:type=Table,name=AllMemtablesOffHeapSize,attr=Value", metric.GAUGE},

	"db.keyspace":                {"keyspace", metric.ATTRIBUTE},
	"db.columnFamily":            {"columnFamily", metric.ATTRIBUTE},
	"db.keyspaceAndColumnFamily": {"keyspaceAndColumnFamily", metric.ATTRIBUTE},
}

// The patterns used to get all the beans needed for the metrics defined above
var jmxPatterns = []string{
	"org.apache.cassandra.metrics:type=ClientRequest,scope=*,name=Latency",
	"org.apache.cassandra.metrics:type=ClientRequest,scope=*,name=Timeouts",
	"org.apache.cassandra.metrics:type=ClientRequest,scope=*,name=Unavailables",
	"org.apache.cassandra.metrics:type=Table,name=LiveSSTableCount",
	"org.apache.cassandra.metrics:type=Table,name=AllMemtablesHeapSize",
	"org.apache.cassandra.metrics:type=Table,name=AllMemtablesOffHeapSize",
	"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=*,name=ActiveTasks",
	"org.apache.cassandra.metrics:type=ThreadPools,path=request,scope=*,name=PendingTasks",
	"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=*,name=ActiveTasks",
	"org.apache.cassandra.metrics:type=ThreadPools,path=internal,scope=*,name=PendingTasks",
	"org.apache.cassandra.metrics:type=DroppedMessage,scope=*,name=Dropped",
	"org.apache.cassandra.metrics:type=Storage,name=TotalHints",
	"org.apache.cassandra.metrics:type=Storage,name=TotalHintsInProgress",
	"org.apache.cassandra.metrics:type=Cache,scope=*,name=*",
	"org.apache.cassandra.metrics:type=CommitLog,name=*",
	// Cluster and software information
	"org.apache.cassandra.db:type=StorageService",
	"org.apache.cassandra.db:type=EndpointSnitchInfo",
	// ColumnFamily metrics
	"org.apache.cassandra.metrics:type=Table,keyspace=*,scope=*,name=LiveSSTableCount",
	"org.apache.cassandra.metrics:type=Table,keyspace=*,scope=*,name=SSTablesPerReadHistogram",
	"org.apache.cassandra.metrics:type=Table,keyspace=*,scope=*,name=LiveDiskSpaceUsed",
	"org.apache.cassandra.metrics:type=Table,keyspace=*,scope=*,name=ReadLatency",
	"org.apache.cassandra.metrics:type=Table,keyspace=*,scope=*,name=WriteLatency",
	"org.apache.cassandra.metrics:type=Table,keyspace=*,scope=*,name=LiveSSTableCount",
	"org.apache.cassandra.metrics:type=Table,keyspace=*,scope=*,name=PendingCompactions",
	"org.apache.cassandra.metrics:type=Table,keyspace=*,scope=*,name=AllMemtablesHeapSize",
	"org.apache.cassandra.metrics:type=Table,keyspace=*,scope=*,name=AllMemtablesOffHeapSize",
}