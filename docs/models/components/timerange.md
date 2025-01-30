# TimeRange

The period of time to return data.  Relative time ranges are relative to the current moment.  Absolute time ranges specify an unchanging period of time.  If not specified, a default relative timeframe of last 24 hours will be chosen.



## Supported Types

### RelativeTimeRange

```go
timeRange := components.CreateTimeRangeRelative(components.RelativeTimeRange{/* values here */})
```

### AbsoluteTimeRange

```go
timeRange := components.CreateTimeRangeAbsolute(components.AbsoluteTimeRange{/* values here */})
```

