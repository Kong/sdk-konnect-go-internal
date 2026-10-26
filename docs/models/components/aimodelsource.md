# AiModelSource

An AI Model source resource.


## Supported Types

### AiModelGatewaySource

```go
aiModelSource := components.CreateAiModelSourceAiModelGatewaySource(components.AiModelGatewaySource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiModelSource.Type {
	case components.AiModelSourceTypeAiModelGatewaySource:
		// aiModelSource.AiModelGatewaySource is populated
}
```
