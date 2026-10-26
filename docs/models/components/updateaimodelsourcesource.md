# UpdateAiModelSourceSource


## Supported Types

### AiModelGatewaySourceBlock

```go
updateAiModelSourceSource := components.CreateUpdateAiModelSourceSourceAiModelGatewaySourceBlock(components.AiModelGatewaySourceBlock{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateAiModelSourceSource.Type {
	case components.UpdateAiModelSourceSourceTypeAiModelGatewaySourceBlock:
		// updateAiModelSourceSource.AiModelGatewaySourceBlock is populated
}
```
