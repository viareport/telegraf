# GetEventStore input plugin

The [event store](https://geteventstore.com/) plugin queries endpoints to obtain node stats

### Configuration:
```
[[inputs.geteventstore]]
  ## specify GetEventStore server
  url = "http://127.0.0.1:2113"

  ## specify Login user
  login = "admin"
  
  ## specify Password user
  password = "changeit"
 
  ## specify stream stats 
  streamname = "$stats-0.0.0.0:2113"
``` 
### Measurements & Fields:

field queue usage:
- es-queue
  - avgItemsPerSecond
  - avgProcessingTime
  - currentIdleTime
  - currentItemProcessingTime
  - idleTimePercent
  - inProgressMessage
  - lengthCurrentTryPeak
  - length
  - lengthLifetimePeak
  - totalItemsProcessed
  
