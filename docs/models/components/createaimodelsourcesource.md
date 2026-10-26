# CreateAiModelSourceSource


## Supported Types

### AiModelGatewaySourceBlock

```go
createAiModelSourceSource := components.CreateCreateAiModelSourceSourceAiModelGatewaySourceBlock(components.AiModelGatewaySourceBlock{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAiModelSourceSource.Type {
	case components.CreateAiModelSourceSourceTypeAiModelGatewaySourceBlock:
		// createAiModelSourceSource.AiModelGatewaySourceBlock is populated
}
```
