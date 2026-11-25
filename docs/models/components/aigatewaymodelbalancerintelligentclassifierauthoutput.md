# AIGatewayModelBalancerIntelligentClassifierAuthOutput

The credentials sent to the classifier service with each request.

**Requires a minimum runtime version of `2.5`**.


## Supported Types

### AIGatewayModelBalancerIntelligentClassifierAuthBasicOutput

```go
aiGatewayModelBalancerIntelligentClassifierAuthOutput := components.CreateAIGatewayModelBalancerIntelligentClassifierAuthOutputBasic(components.AIGatewayModelBalancerIntelligentClassifierAuthBasicOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelBalancerIntelligentClassifierAuthOutput.Type {
	case components.AIGatewayModelBalancerIntelligentClassifierAuthOutputTypeBasic:
		// aiGatewayModelBalancerIntelligentClassifierAuthOutput.AIGatewayModelBalancerIntelligentClassifierAuthBasicOutput is populated
}
```
