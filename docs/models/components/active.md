# Active


## Supported Types

###

```go
active := components.CreateActiveBoolean(bool{/* values here */})
```

### StringFieldOEQFilter

```go
active := components.CreateActiveStringFieldOEQFilter(components.StringFieldOEQFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch active.Type {
	case components.ActiveTypeBoolean:
		// active.Boolean is populated
	case components.ActiveTypeStringFieldOEQFilter:
		// active.StringFieldOEQFilter is populated
}
```
