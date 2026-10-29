# Payload

The payload delivered to the channels.


## Supported Types

### BillingNotificationEventBalanceThresholdPayload

```go
payload := components.CreatePayloadEntitlementsBalanceThreshold(components.BillingNotificationEventBalanceThresholdPayload{/* values here */})
```

### BillingNotificationEventResetPayload

```go
payload := components.CreatePayloadEntitlementsReset(components.BillingNotificationEventResetPayload{/* values here */})
```

### BillingNotificationEventInvoiceCreatedPayload

```go
payload := components.CreatePayloadInvoiceCreated(components.BillingNotificationEventInvoiceCreatedPayload{/* values here */})
```

### BillingNotificationEventInvoiceUpdatedPayload

```go
payload := components.CreatePayloadInvoiceUpdated(components.BillingNotificationEventInvoiceUpdatedPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch payload.Type {
	case components.PayloadUnionTypeEntitlementsBalanceThreshold:
		// payload.BillingNotificationEventBalanceThresholdPayload is populated
	case components.PayloadUnionTypeEntitlementsReset:
		// payload.BillingNotificationEventResetPayload is populated
	case components.PayloadUnionTypeInvoiceCreated:
		// payload.BillingNotificationEventInvoiceCreatedPayload is populated
	case components.PayloadUnionTypeInvoiceUpdated:
		// payload.BillingNotificationEventInvoiceUpdatedPayload is populated
}
```
