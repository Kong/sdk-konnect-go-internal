# EnableConnectedAppsPayload

Connected Apps configuration (enable or disable)


## Supported Types

### EnableConnectedAppsKonnect

```go
enableConnectedAppsPayload := components.CreateEnableConnectedAppsPayloadKonnect(components.EnableConnectedAppsKonnect{/* values here */})
```

### EnableConnectedAppsPortal

```go
enableConnectedAppsPayload := components.CreateEnableConnectedAppsPayloadPortal(components.EnableConnectedAppsPortal{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch enableConnectedAppsPayload.Type {
	case components.EnableConnectedAppsPayloadTypeKonnect:
		// enableConnectedAppsPayload.EnableConnectedAppsKonnect is populated
	case components.EnableConnectedAppsPayloadTypePortal:
		// enableConnectedAppsPayload.EnableConnectedAppsPortal is populated
}
```
