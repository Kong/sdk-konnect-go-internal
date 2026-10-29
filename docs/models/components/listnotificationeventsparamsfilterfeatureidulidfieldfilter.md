# ListNotificationEventsParamsFilterFeatureIDULIDFieldFilter

Filter by the id of the feature the event refers to. Events without a feature
never match, not even with `neq`.


## Supported Types

###

```go
listNotificationEventsParamsFilterFeatureIDULIDFieldFilter := components.CreateListNotificationEventsParamsFilterFeatureIDULIDFieldFilterStr(string{/* values here */})
```

### ListNotificationEventsParamsFilterULIDFieldFilterFeatureID2

```go
listNotificationEventsParamsFilterFeatureIDULIDFieldFilter := components.CreateListNotificationEventsParamsFilterFeatureIDULIDFieldFilterListNotificationEventsParamsFilterULIDFieldFilterFeatureID2(components.ListNotificationEventsParamsFilterULIDFieldFilterFeatureID2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listNotificationEventsParamsFilterFeatureIDULIDFieldFilter.Type {
	case components.ListNotificationEventsParamsFilterFeatureIDULIDFieldFilterTypeStr:
		// listNotificationEventsParamsFilterFeatureIDULIDFieldFilter.Str is populated
	case components.ListNotificationEventsParamsFilterFeatureIDULIDFieldFilterTypeListNotificationEventsParamsFilterULIDFieldFilterFeatureID2:
		// listNotificationEventsParamsFilterFeatureIDULIDFieldFilter.ListNotificationEventsParamsFilterULIDFieldFilterFeatureID2 is populated
}
```
