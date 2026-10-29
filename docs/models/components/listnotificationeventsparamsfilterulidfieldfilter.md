# ListNotificationEventsParamsFilterULIDFieldFilter

Filter by a channel of the generating rule. Matches events whose rule targets at
least one of the given channels. Only `eq` and `oeq` are supported.


## Supported Types

### 

```go
listNotificationEventsParamsFilterULIDFieldFilter := components.CreateListNotificationEventsParamsFilterULIDFieldFilterStr(string{/* values here */})
```

### ListNotificationEventsParamsFilterULIDFieldFilter2

```go
listNotificationEventsParamsFilterULIDFieldFilter := components.CreateListNotificationEventsParamsFilterULIDFieldFilterListNotificationEventsParamsFilterULIDFieldFilter2(components.ListNotificationEventsParamsFilterULIDFieldFilter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listNotificationEventsParamsFilterULIDFieldFilter.Type {
	case components.ListNotificationEventsParamsFilterULIDFieldFilterTypeStr:
		// listNotificationEventsParamsFilterULIDFieldFilter.Str is populated
	case components.ListNotificationEventsParamsFilterULIDFieldFilterTypeListNotificationEventsParamsFilterULIDFieldFilter2:
		// listNotificationEventsParamsFilterULIDFieldFilter.ListNotificationEventsParamsFilterULIDFieldFilter2 is populated
}
```
