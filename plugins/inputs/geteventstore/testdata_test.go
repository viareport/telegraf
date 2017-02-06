package geteventstore

const lastEvents = `
<?xml version="1.0" encoding="utf-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
	<title>Event stream '$stats-0.0.0.0:2113'</title>
	<id>http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113</id>
	<updated>2017-02-06T08:49:42.680925Z</updated>
	<author>
		<name>EventStore</name>
	</author>
	<link href="http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113" rel="self" />
	<link href="http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113/head/backward/1" rel="first" />
	<link href="http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113/0/forward/1" rel="last" />
	<link href="http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113/16569/backward/1" rel="next" />
	<link href="http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113/16571/forward/1" rel="previous" />
	<link href="http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113/metadata" rel="metadata" />
	<entry>
		<title>16570@$stats-0.0.0.0:2113</title>
		<id>http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113/16570</id>
		<updated>2017-02-06T08:49:42.680925Z</updated>
		<author>
			<name>EventStore</name>
		</author>
		<summary>$statsCollected</summary>
		<link href="http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113/16570" rel="edit" />
		<link href="http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113/16570" rel="alternate" />
	</entry>
</feed>
`

const statsResponse = `
{
  "title": "16570@$stats-0.0.0.0:2113",
  "id": "http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113/16570",
  "updated": "2017-02-06T08:49:42.680925Z",
  "author": {
    "name": "EventStore"
  },
  "summary": "$statsCollected",
  "content": {
    "eventStreamId": "$stats-0.0.0.0:2113",
    "eventNumber": 16570,
    "eventType": "$statsCollected",
    "eventId": "fce5e5ec-1a9f-48bf-8b4f-4a177b553abd",
    "data": {
      "proc-startTime": "2017-01-31T14:28:30.0000000Z",
      "proc-id": 1,
      "proc-mem": 372736000,
      "proc-cpu": 0.0,
      "proc-cpuScaled": 0.0,
      "proc-threadsCount": 0,
      "proc-contentionsRate": 0.0,
      "proc-thrownExceptionsRate": 0.0,
      "sys-cpu": 6.901122,
      "sys-freeMem": 6277042176,
      "proc-diskIo-readBytes": 1245052928,
      "proc-diskIo-writtenBytes": 78434304,
      "proc-diskIo-readOps": 5099748,
      "proc-diskIo-writeOps": 55411,
      "proc-tcp-connections": 0,
      "proc-tcp-receivingSpeed": 0.0,
      "proc-tcp-sendingSpeed": 0.0,
      "proc-tcp-inSend": 0,
      "proc-tcp-measureTime": "00:00:30.0622670",
      "proc-tcp-pendingReceived": 0,
      "proc-tcp-pendingSend": 0,
      "proc-tcp-receivedBytesSinceLastRun": 0,
      "proc-tcp-receivedBytesTotal": 33054035,
      "proc-tcp-sentBytesSinceLastRun": 0,
      "proc-tcp-sentBytesTotal": 1692062,
      "proc-gc-allocationSpeed": 0.0,
      "proc-gc-gen0ItemsCount": 0,
      "proc-gc-gen0Size": 0,
      "proc-gc-gen1ItemsCount": 0,
      "proc-gc-gen1Size": 0,
      "proc-gc-gen2ItemsCount": 0,
      "proc-gc-gen2Size": 0,
      "proc-gc-largeHeapSize": 0,
      "proc-gc-timeInGc": 0.0,
      "proc-gc-totalBytesInHeaps": 0,
      "es-checksum": 956167012,
      "es-checksumNonFlushed": 956167012,
      "es-queue-MainQueue-queueName": "MainQueue",
      "es-queue-MainQueue-groupName": "",
      "es-queue-MainQueue-avgItemsPerSecond": 12,
      "es-queue-MainQueue-avgProcessingTime": 0.0192825641025641,
      "es-queue-MainQueue-currentIdleTime": "0:00:00:00.2087878",
      "es-queue-MainQueue-currentItemProcessingTime": null,
      "es-queue-MainQueue-idleTimePercent": 99.974977960683432,
      "es-queue-MainQueue-length": 0,
      "es-queue-MainQueue-lengthCurrentTryPeak": 3,
      "es-queue-MainQueue-lengthLifetimePeak": 10858,
      "es-queue-MainQueue-totalItemsProcessed": 16039587,
      "es-queue-MainQueue-inProgressMessage": "<none>",
      "es-queue-MainQueue-lastProcessedMessage": "Schedule",
      "es-queue-MonitoringQueue-queueName": "MonitoringQueue",
      "es-queue-MonitoringQueue-groupName": "",
      "es-queue-MonitoringQueue-avgItemsPerSecond": 0,
      "es-queue-MonitoringQueue-avgProcessingTime": 0.0,
      "es-queue-MonitoringQueue-currentIdleTime": "4:16:31:30.5307446",
      "es-queue-MonitoringQueue-currentItemProcessingTime": null,
      "es-queue-MonitoringQueue-idleTimePercent": 100.0,
      "es-queue-MonitoringQueue-length": 0,
      "es-queue-MonitoringQueue-lengthCurrentTryPeak": 0,
      "es-queue-MonitoringQueue-lengthLifetimePeak": 2,
      "es-queue-MonitoringQueue-totalItemsProcessed": 83,
      "es-queue-MonitoringQueue-inProgressMessage": "<none>",
      "es-queue-MonitoringQueue-lastProcessedMessage": "GetFreshTcpConnectionStats",
      "es-queue-Projection Core #0-queueName": "Projection Core #0",
      "es-queue-Projection Core #0-groupName": "Projection Core",
      "es-queue-Projection Core #0-avgItemsPerSecond": 0,
      "es-queue-Projection Core #0-avgProcessingTime": 0.041514285714285719,
      "es-queue-Projection Core #0-currentIdleTime": "0:00:00:30.0597386",
      "es-queue-Projection Core #0-currentItemProcessingTime": null,
      "es-queue-Projection Core #0-idleTimePercent": 99.999033008032953,
      "es-queue-Projection Core #0-length": 0,
      "es-queue-Projection Core #0-lengthCurrentTryPeak": 2,
      "es-queue-Projection Core #0-lengthLifetimePeak": 1403,
      "es-queue-Projection Core #0-totalItemsProcessed": 6113818,
      "es-queue-Projection Core #0-inProgressMessage": "<none>",
      "es-queue-Projection Core #0-lastProcessedMessage": "EventReaderIdle",
      "es-queue-Projection Core #1-queueName": "Projection Core #1",
      "es-queue-Projection Core #1-groupName": "Projection Core",
      "es-queue-Projection Core #1-avgItemsPerSecond": 0,
      "es-queue-Projection Core #1-avgProcessingTime": 0.042328571428571431,
      "es-queue-Projection Core #1-currentIdleTime": "0:00:00:30.0597525",
      "es-queue-Projection Core #1-currentItemProcessingTime": null,
      "es-queue-Projection Core #1-idleTimePercent": 99.9990147126913,
      "es-queue-Projection Core #1-length": 0,
      "es-queue-Projection Core #1-lengthCurrentTryPeak": 2,
      "es-queue-Projection Core #1-lengthLifetimePeak": 1532,
      "es-queue-Projection Core #1-totalItemsProcessed": 5796733,
      "es-queue-Projection Core #1-inProgressMessage": "<none>",
      "es-queue-Projection Core #1-lastProcessedMessage": "EventReaderIdle",
      "es-queue-Projection Core #2-queueName": "Projection Core #2",
      "es-queue-Projection Core #2-groupName": "Projection Core",
      "es-queue-Projection Core #2-avgItemsPerSecond": 0,
      "es-queue-Projection Core #2-avgProcessingTime": 0.0324,
      "es-queue-Projection Core #2-currentIdleTime": "0:00:00:30.0598167",
      "es-queue-Projection Core #2-currentItemProcessingTime": null,
      "es-queue-Projection Core #2-idleTimePercent": 99.999245566628673,
      "es-queue-Projection Core #2-length": 0,
      "es-queue-Projection Core #2-lengthCurrentTryPeak": 2,
      "es-queue-Projection Core #2-lengthLifetimePeak": 883,
      "es-queue-Projection Core #2-totalItemsProcessed": 3682245,
      "es-queue-Projection Core #2-inProgressMessage": "<none>",
      "es-queue-Projection Core #2-lastProcessedMessage": "EventReaderIdle",
      "es-queue-Projections Master-queueName": "Projections Master",
      "es-queue-Projections Master-groupName": "",
      "es-queue-Projections Master-avgItemsPerSecond": 19,
      "es-queue-Projections Master-avgProcessingTime": 0.0088816666666666662,
      "es-queue-Projections Master-currentIdleTime": "0:00:00:00.0833222",
      "es-queue-Projections Master-currentItemProcessingTime": null,
      "es-queue-Projections Master-idleTimePercent": 99.9821976344063,
      "es-queue-Projections Master-length": 0,
      "es-queue-Projections Master-lengthCurrentTryPeak": 0,
      "es-queue-Projections Master-lengthLifetimePeak": 7,
      "es-queue-Projections Master-totalItemsProcessed": 10080451,
      "es-queue-Projections Master-inProgressMessage": "<none>",
      "es-queue-Projections Master-lastProcessedMessage": "Schedule",
      "es-queue-Storage Chaser-queueName": "Storage Chaser",
      "es-queue-Storage Chaser-groupName": "",
      "es-queue-Storage Chaser-avgItemsPerSecond": 99,
      "es-queue-Storage Chaser-avgProcessingTime": 0.0060370592149035264,
      "es-queue-Storage Chaser-currentIdleTime": "0:00:00:00.0070188",
      "es-queue-Storage Chaser-currentItemProcessingTime": null,
      "es-queue-Storage Chaser-idleTimePercent": 99.9396852474037,
      "es-queue-Storage Chaser-length": 0,
      "es-queue-Storage Chaser-lengthCurrentTryPeak": 0,
      "es-queue-Storage Chaser-lengthLifetimePeak": 0,
      "es-queue-Storage Chaser-totalItemsProcessed": 51528120,
      "es-queue-Storage Chaser-inProgressMessage": "<none>",
      "es-queue-Storage Chaser-lastProcessedMessage": "ChaserCheckpointFlush",
      "es-queue-StorageReaderQueue #1-queueName": "StorageReaderQueue #1",
      "es-queue-StorageReaderQueue #1-groupName": "StorageReaderQueue",
      "es-queue-StorageReaderQueue #1-avgItemsPerSecond": 0,
      "es-queue-StorageReaderQueue #1-avgProcessingTime": 0.067688888888888882,
      "es-queue-StorageReaderQueue #1-currentIdleTime": "0:00:00:02.5468934",
      "es-queue-StorageReaderQueue #1-currentItemProcessingTime": null,
      "es-queue-StorageReaderQueue #1-idleTimePercent": 99.997971212925137,
      "es-queue-StorageReaderQueue #1-length": 0,
      "es-queue-StorageReaderQueue #1-lengthCurrentTryPeak": 0,
      "es-queue-StorageReaderQueue #1-lengthLifetimePeak": 859,
      "es-queue-StorageReaderQueue #1-totalItemsProcessed": 852873,
      "es-queue-StorageReaderQueue #1-inProgressMessage": "<none>",
      "es-queue-StorageReaderQueue #1-lastProcessedMessage": "ReadStreamEventsForward",
      "es-queue-StorageReaderQueue #2-queueName": "StorageReaderQueue #2",
      "es-queue-StorageReaderQueue #2-groupName": "StorageReaderQueue",
      "es-queue-StorageReaderQueue #2-avgItemsPerSecond": 0,
      "es-queue-StorageReaderQueue #2-avgProcessingTime": 0.3997,
      "es-queue-StorageReaderQueue #2-currentIdleTime": "0:00:00:02.2408833",
      "es-queue-StorageReaderQueue #2-currentItemProcessingTime": null,
      "es-queue-StorageReaderQueue #2-idleTimePercent": 99.988031852775265,
      "es-queue-StorageReaderQueue #2-length": 0,
      "es-queue-StorageReaderQueue #2-lengthCurrentTryPeak": 0,
      "es-queue-StorageReaderQueue #2-lengthLifetimePeak": 679,
      "es-queue-StorageReaderQueue #2-totalItemsProcessed": 891592,
      "es-queue-StorageReaderQueue #2-inProgressMessage": "<none>",
      "es-queue-StorageReaderQueue #2-lastProcessedMessage": "ReadStreamEventsBackward",
      "es-queue-StorageReaderQueue #3-queueName": "StorageReaderQueue #3",
      "es-queue-StorageReaderQueue #3-groupName": "StorageReaderQueue",
      "es-queue-StorageReaderQueue #3-avgItemsPerSecond": 0,
      "es-queue-StorageReaderQueue #3-avgProcessingTime": 0.073119999999999991,
      "es-queue-StorageReaderQueue #3-currentIdleTime": "0:00:00:01.5458169",
      "es-queue-StorageReaderQueue #3-currentItemProcessingTime": null,
      "es-queue-StorageReaderQueue #3-idleTimePercent": 99.997568050130354,
      "es-queue-StorageReaderQueue #3-length": 0,
      "es-queue-StorageReaderQueue #3-lengthCurrentTryPeak": 0,
      "es-queue-StorageReaderQueue #3-lengthLifetimePeak": 612,
      "es-queue-StorageReaderQueue #3-totalItemsProcessed": 852776,
      "es-queue-StorageReaderQueue #3-inProgressMessage": "<none>",
      "es-queue-StorageReaderQueue #3-lastProcessedMessage": "ReadStreamEventsForward",
      "es-queue-StorageReaderQueue #4-queueName": "StorageReaderQueue #4",
      "es-queue-StorageReaderQueue #4-groupName": "StorageReaderQueue",
      "es-queue-StorageReaderQueue #4-avgItemsPerSecond": 0,
      "es-queue-StorageReaderQueue #4-avgProcessingTime": 0.071218181818181819,
      "es-queue-StorageReaderQueue #4-currentIdleTime": "0:00:00:00.5451610",
      "es-queue-StorageReaderQueue #4-currentItemProcessingTime": null,
      "es-queue-StorageReaderQueue #4-idleTimePercent": 99.997358816219574,
      "es-queue-StorageReaderQueue #4-length": 0,
      "es-queue-StorageReaderQueue #4-lengthCurrentTryPeak": 0,
      "es-queue-StorageReaderQueue #4-lengthLifetimePeak": 510,
      "es-queue-StorageReaderQueue #4-totalItemsProcessed": 867808,
      "es-queue-StorageReaderQueue #4-inProgressMessage": "<none>",
      "es-queue-StorageReaderQueue #4-lastProcessedMessage": "ReadStreamEventsForward",
      "es-queue-StorageWriterQueue-queueName": "StorageWriterQueue",
      "es-queue-StorageWriterQueue-groupName": "",
      "es-queue-StorageWriterQueue-avgItemsPerSecond": 0,
      "es-queue-StorageWriterQueue-avgProcessingTime": 0.2168,
      "es-queue-StorageWriterQueue-currentIdleTime": "0:00:00:30.0603688",
      "es-queue-StorageWriterQueue-currentItemProcessingTime": null,
      "es-queue-StorageWriterQueue-idleTimePercent": 99.999278830399689,
      "es-queue-StorageWriterQueue-length": 0,
      "es-queue-StorageWriterQueue-lengthCurrentTryPeak": 0,
      "es-queue-StorageWriterQueue-lengthLifetimePeak": 19801,
      "es-queue-StorageWriterQueue-totalItemsProcessed": 1033206,
      "es-queue-StorageWriterQueue-inProgressMessage": "<none>",
      "es-queue-StorageWriterQueue-lastProcessedMessage": "WritePrepares",
      "es-queue-Subscriptions-queueName": "Subscriptions",
      "es-queue-Subscriptions-groupName": "",
      "es-queue-Subscriptions-avgItemsPerSecond": 1,
      "es-queue-Subscriptions-avgProcessingTime": 0.023806451612903224,
      "es-queue-Subscriptions-currentIdleTime": "0:00:00:00.3261471",
      "es-queue-Subscriptions-currentItemProcessingTime": null,
      "es-queue-Subscriptions-idleTimePercent": 99.9975317903244,
      "es-queue-Subscriptions-length": 0,
      "es-queue-Subscriptions-lengthCurrentTryPeak": 0,
      "es-queue-Subscriptions-lengthLifetimePeak": 891,
      "es-queue-Subscriptions-totalItemsProcessed": 2060136,
      "es-queue-Subscriptions-inProgressMessage": "<none>",
      "es-queue-Subscriptions-lastProcessedMessage": "CheckPollTimeout",
      "es-queue-Timer-queueName": "Timer",
      "es-queue-Timer-groupName": "",
      "es-queue-Timer-avgItemsPerSecond": 32,
      "es-queue-Timer-avgProcessingTime": 0.10219372427983539,
      "es-queue-Timer-currentIdleTime": "0:00:00:00.0001143",
      "es-queue-Timer-currentItemProcessingTime": null,
      "es-queue-Timer-idleTimePercent": 99.663947923675451,
      "es-queue-Timer-length": 9,
      "es-queue-Timer-lengthCurrentTryPeak": 9,
      "es-queue-Timer-lengthLifetimePeak": 15,
      "es-queue-Timer-totalItemsProcessed": 16124691,
      "es-queue-Timer-inProgressMessage": "<none>",
      "es-queue-Timer-lastProcessedMessage": "ExecuteScheduledTasks",
      "es-queue-Worker #1-queueName": "Worker #1",
      "es-queue-Worker #1-groupName": "Workers",
      "es-queue-Worker #1-avgItemsPerSecond": 1,
      "es-queue-Worker #1-avgProcessingTime": 0.020327906976744185,
      "es-queue-Worker #1-currentIdleTime": "0:00:00:00.4645613",
      "es-queue-Worker #1-currentItemProcessingTime": null,
      "es-queue-Worker #1-idleTimePercent": 99.99709569550906,
      "es-queue-Worker #1-length": 0,
      "es-queue-Worker #1-lengthCurrentTryPeak": 0,
      "es-queue-Worker #1-lengthLifetimePeak": 1,
      "es-queue-Worker #1-totalItemsProcessed": 1560844,
      "es-queue-Worker #1-inProgressMessage": "<none>",
      "es-queue-Worker #1-lastProcessedMessage": "PurgeTimedOutRequests",
      "es-queue-Worker #2-queueName": "Worker #2",
      "es-queue-Worker #2-groupName": "Workers",
      "es-queue-Worker #2-avgItemsPerSecond": 1,
      "es-queue-Worker #2-avgProcessingTime": 0.010634883720930232,
      "es-queue-Worker #2-currentIdleTime": "0:00:00:00.4645612",
      "es-queue-Worker #2-currentItemProcessingTime": null,
      "es-queue-Worker #2-idleTimePercent": 99.998468179772857,
      "es-queue-Worker #2-length": 0,
      "es-queue-Worker #2-lengthCurrentTryPeak": 0,
      "es-queue-Worker #2-lengthLifetimePeak": 1,
      "es-queue-Worker #2-totalItemsProcessed": 1568720,
      "es-queue-Worker #2-inProgressMessage": "<none>",
      "es-queue-Worker #2-lastProcessedMessage": "PurgeTimedOutRequests",
      "es-queue-Worker #3-queueName": "Worker #3",
      "es-queue-Worker #3-groupName": "Workers",
      "es-queue-Worker #3-avgItemsPerSecond": 1,
      "es-queue-Worker #3-avgProcessingTime": 0.035368181818181819,
      "es-queue-Worker #3-currentIdleTime": "0:00:00:00.4645584",
      "es-queue-Worker #3-currentItemProcessingTime": null,
      "es-queue-Worker #3-idleTimePercent": 99.994821416318885,
      "es-queue-Worker #3-length": 0,
      "es-queue-Worker #3-lengthCurrentTryPeak": 0,
      "es-queue-Worker #3-lengthLifetimePeak": 1768,
      "es-queue-Worker #3-totalItemsProcessed": 2059420,
      "es-queue-Worker #3-inProgressMessage": "<none>",
      "es-queue-Worker #3-lastProcessedMessage": "PurgeTimedOutRequests",
      "es-queue-Worker #4-queueName": "Worker #4",
      "es-queue-Worker #4-groupName": "Workers",
      "es-queue-Worker #4-avgItemsPerSecond": 1,
      "es-queue-Worker #4-avgProcessingTime": 0.010337209302325581,
      "es-queue-Worker #4-currentIdleTime": "0:00:00:00.4645594",
      "es-queue-Worker #4-currentItemProcessingTime": null,
      "es-queue-Worker #4-idleTimePercent": 99.998502774597171,
      "es-queue-Worker #4-length": 0,
      "es-queue-Worker #4-lengthCurrentTryPeak": 0,
      "es-queue-Worker #4-lengthLifetimePeak": 2,
      "es-queue-Worker #4-totalItemsProcessed": 1572525,
      "es-queue-Worker #4-inProgressMessage": "<none>",
      "es-queue-Worker #4-lastProcessedMessage": "PurgeTimedOutRequests",
      "es-queue-Worker #5-queueName": "Worker #5",
      "es-queue-Worker #5-groupName": "Workers",
      "es-queue-Worker #5-avgItemsPerSecond": 1,
      "es-queue-Worker #5-avgProcessingTime": 0.010506976744186046,
      "es-queue-Worker #5-currentIdleTime": "0:00:00:00.4645569",
      "es-queue-Worker #5-currentItemProcessingTime": null,
      "es-queue-Worker #5-idleTimePercent": 99.9984957890626,
      "es-queue-Worker #5-length": 0,
      "es-queue-Worker #5-lengthCurrentTryPeak": 0,
      "es-queue-Worker #5-lengthLifetimePeak": 3,
      "es-queue-Worker #5-totalItemsProcessed": 1564078,
      "es-queue-Worker #5-inProgressMessage": "<none>",
      "es-queue-Worker #5-lastProcessedMessage": "PurgeTimedOutRequests",
      "es-writer-lastFlushSize": 16174,
      "es-writer-lastFlushDelayMs": 0.0087,
      "es-writer-meanFlushSize": 16173,
      "es-writer-meanFlushDelayMs": 0.009936328125,
      "es-writer-maxFlushSize": 278784,
      "es-writer-maxFlushDelayMs": 50.1815,
      "es-writer-queuedFlushMessages": 0,
      "es-readIndex-cachedRecord": 14930701,
      "es-readIndex-notCachedRecord": 0,
      "es-readIndex-cachedStreamInfo": 17640455,
      "es-readIndex-notCachedStreamInfo": 11165,
      "es-readIndex-cachedTransInfo": 0,
      "es-readIndex-notCachedTransInfo": 0
    },
    "metadata": ""
  },
  "links": [
    {
      "uri": "http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113/16570",
      "relation": "edit"
    },
    {
      "uri": "http://192.168.1.30:2113/streams/%24stats-0.0.0.0%3A2113/16570",
      "relation": "alternate"
    }
  ]
}
`

var esQueueExpected = map[string]interface{}{
	"avgItemsPerSecond":         float64(12),
	"avgProcessingTime":         float64(0.0192825641025641),
	"currentIdleTime":           "0:00:00:00.2087878",
	"currentItemProcessingTime": nil,
	"groupName":                 "",
	"idleTimePercent":           float64(99.97497796068343),
	"inProgressMessage":         "<none>",
	"lengthCurrentTryPeak":      float64(3),
	"length":                    float64(0),
	"lengthLifetimePeak":        float64(10858),
	"totalItemsProcessed":       float64(1.6039587e+07),
}
