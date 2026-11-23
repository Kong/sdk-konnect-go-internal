# AIGatewayModelBalancerIntelligentClassifierAuth

The credentials sent to the classifier service with each request.

**Requires a minimum runtime version of `2.5`**.


## Supported Types

### AIGatewayModelBalancerIntelligentClassifierAuthBasic

```go
aiGatewayModelBalancerIntelligentClassifierAuth := components.CreateAIGatewayModelBalancerIntelligentClassifierAuthBasic(components.AIGatewayModelBalancerIntelligentClassifierAuthBasic{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelBalancerIntelligentClassifierAuth.Type {
	case components.AIGatewayModelBalancerIntelligentClassifierAuthTypeBasic:
		// aiGatewayModelBalancerIntelligentClassifierAuth.AIGatewayModelBalancerIntelligentClassifierAuthBasic is populated
}
```
