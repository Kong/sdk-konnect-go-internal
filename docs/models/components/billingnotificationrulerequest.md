# BillingNotificationRuleRequest

Request body for creating or updating a notification rule. Updates replace the
rule's mutable state: omitting `disabled`, `labels`, or `features` resets them
to their defaults. The `type` must match the existing rule on update.


## Supported Types

### BillingNotificationRuleBalanceThresholdRequest

```go
billingNotificationRuleRequest := components.CreateBillingNotificationRuleRequestEntitlementsBalanceThreshold(components.BillingNotificationRuleBalanceThresholdRequest{/* values here */})
```

### BillingNotificationRuleEntitlementResetRequest

```go
billingNotificationRuleRequest := components.CreateBillingNotificationRuleRequestEntitlementsReset(components.BillingNotificationRuleEntitlementResetRequest{/* values here */})
```

### BillingNotificationRuleInvoiceCreatedRequest

```go
billingNotificationRuleRequest := components.CreateBillingNotificationRuleRequestInvoiceCreated(components.BillingNotificationRuleInvoiceCreatedRequest{/* values here */})
```

### BillingNotificationRuleInvoiceUpdatedRequest

```go
billingNotificationRuleRequest := components.CreateBillingNotificationRuleRequestInvoiceUpdated(components.BillingNotificationRuleInvoiceUpdatedRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingNotificationRuleRequest.Type {
	case components.BillingNotificationRuleRequestTypeEntitlementsBalanceThreshold:
		// billingNotificationRuleRequest.BillingNotificationRuleBalanceThresholdRequest is populated
	case components.BillingNotificationRuleRequestTypeEntitlementsReset:
		// billingNotificationRuleRequest.BillingNotificationRuleEntitlementResetRequest is populated
	case components.BillingNotificationRuleRequestTypeInvoiceCreated:
		// billingNotificationRuleRequest.BillingNotificationRuleInvoiceCreatedRequest is populated
	case components.BillingNotificationRuleRequestTypeInvoiceUpdated:
		// billingNotificationRuleRequest.BillingNotificationRuleInvoiceUpdatedRequest is populated
}
```
