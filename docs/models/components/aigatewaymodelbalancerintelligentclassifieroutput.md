# AIGatewayModelBalancerIntelligentClassifierOutput

The classifier model the intelligent balancer asks what kind of work a conversation is, before it
selects a target. If the classifier does not answer in time or its answer cannot be used, the
balancer selects a target for general work.

**Requires a minimum runtime version of `2.5`**.


## Supported Types

### AIGatewayModelBalancerIntelligentClassifierCustomOutput

```go
aiGatewayModelBalancerIntelligentClassifierOutput := components.CreateAIGatewayModelBalancerIntelligentClassifierOutputCustom(components.AIGatewayModelBalancerIntelligentClassifierCustomOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelBalancerIntelligentClassifierOutput.Type {
	case components.AIGatewayModelBalancerIntelligentClassifierOutputTypeCustom:
		// aiGatewayModelBalancerIntelligentClassifierOutput.AIGatewayModelBalancerIntelligentClassifierCustomOutput is populated
}
```
