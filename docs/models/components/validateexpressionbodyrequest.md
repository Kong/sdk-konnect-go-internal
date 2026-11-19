# ValidateExpressionBodyRequest


## Supported Types

### ValidateExpressionValueRequest

```go
validateExpressionBodyRequest := components.CreateValidateExpressionBodyRequestValidateExpressionValueRequest(components.ValidateExpressionValueRequest{/* values here */})
```

### ValidateConditionRequest

```go
validateExpressionBodyRequest := components.CreateValidateExpressionBodyRequestValidateConditionRequest(components.ValidateConditionRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch validateExpressionBodyRequest.Type {
	case components.ValidateExpressionBodyRequestTypeValidateExpressionValueRequest:
		// validateExpressionBodyRequest.ValidateExpressionValueRequest is populated
	case components.ValidateExpressionBodyRequestTypeValidateConditionRequest:
		// validateExpressionBodyRequest.ValidateConditionRequest is populated
}
```
