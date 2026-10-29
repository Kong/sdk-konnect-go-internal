# ListNotificationRulesParamsFilterULIDFieldFilter

Filter by an assigned channel. Matches rules that deliver to at least one of the
given channels. Only `eq` and `oeq` are supported.


## Supported Types

### 

```go
listNotificationRulesParamsFilterULIDFieldFilter := components.CreateListNotificationRulesParamsFilterULIDFieldFilterStr(string{/* values here */})
```

### ListNotificationRulesParamsFilterULIDFieldFilter2

```go
listNotificationRulesParamsFilterULIDFieldFilter := components.CreateListNotificationRulesParamsFilterULIDFieldFilterListNotificationRulesParamsFilterULIDFieldFilter2(components.ListNotificationRulesParamsFilterULIDFieldFilter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listNotificationRulesParamsFilterULIDFieldFilter.Type {
	case components.ListNotificationRulesParamsFilterULIDFieldFilterTypeStr:
		// listNotificationRulesParamsFilterULIDFieldFilter.Str is populated
	case components.ListNotificationRulesParamsFilterULIDFieldFilterTypeListNotificationRulesParamsFilterULIDFieldFilter2:
		// listNotificationRulesParamsFilterULIDFieldFilter.ListNotificationRulesParamsFilterULIDFieldFilter2 is populated
}
```
