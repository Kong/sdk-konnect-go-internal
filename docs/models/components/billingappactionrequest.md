# BillingAppActionRequest

Request to execute an operator action on an installed app.


## Supported Types

### BillingAppReconcileWebhookEventsActionRequest

```go
billingAppActionRequest := components.CreateBillingAppActionRequestReconcileWebhookEvents(components.BillingAppReconcileWebhookEventsActionRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingAppActionRequest.Type {
	case components.BillingAppActionRequestTypeReconcileWebhookEvents:
		// billingAppActionRequest.BillingAppReconcileWebhookEventsActionRequest is populated
}
```
