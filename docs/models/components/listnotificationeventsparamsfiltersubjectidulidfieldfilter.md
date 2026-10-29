# ListNotificationEventsParamsFilterSubjectIDULIDFieldFilter

Filter by the id of the subject the event refers to. Events without a subject
never match, not even with `neq`.


## Supported Types

### 

```go
listNotificationEventsParamsFilterSubjectIDULIDFieldFilter := components.CreateListNotificationEventsParamsFilterSubjectIDULIDFieldFilterStr(string{/* values here */})
```

### ListNotificationEventsParamsFilterULIDFieldFilterSubjectID2

```go
listNotificationEventsParamsFilterSubjectIDULIDFieldFilter := components.CreateListNotificationEventsParamsFilterSubjectIDULIDFieldFilterListNotificationEventsParamsFilterULIDFieldFilterSubjectID2(components.ListNotificationEventsParamsFilterULIDFieldFilterSubjectID2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listNotificationEventsParamsFilterSubjectIDULIDFieldFilter.Type {
	case components.ListNotificationEventsParamsFilterSubjectIDULIDFieldFilterTypeStr:
		// listNotificationEventsParamsFilterSubjectIDULIDFieldFilter.Str is populated
	case components.ListNotificationEventsParamsFilterSubjectIDULIDFieldFilterTypeListNotificationEventsParamsFilterULIDFieldFilterSubjectID2:
		// listNotificationEventsParamsFilterSubjectIDULIDFieldFilter.ListNotificationEventsParamsFilterULIDFieldFilterSubjectID2 is populated
}
```
