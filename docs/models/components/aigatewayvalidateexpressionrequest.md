# AIGatewayValidateExpressionRequest


## Supported Types

### AIGatewayValidateExpressionValueRequest

```go
aiGatewayValidateExpressionRequest := components.CreateAIGatewayValidateExpressionRequestAIGatewayValidateExpressionValueRequest(components.AIGatewayValidateExpressionValueRequest{/* values here */})
```

### AIGatewayValidateConditionRequest

```go
aiGatewayValidateExpressionRequest := components.CreateAIGatewayValidateExpressionRequestAIGatewayValidateConditionRequest(components.AIGatewayValidateConditionRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayValidateExpressionRequest.Type {
	case components.AIGatewayValidateExpressionRequestTypeAIGatewayValidateExpressionValueRequest:
		// aiGatewayValidateExpressionRequest.AIGatewayValidateExpressionValueRequest is populated
	case components.AIGatewayValidateExpressionRequestTypeAIGatewayValidateConditionRequest:
		// aiGatewayValidateExpressionRequest.AIGatewayValidateConditionRequest is populated
}
```
