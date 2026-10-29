# BillingNotificationRule

A notification rule selects the type of event to generate, the conditions
specific to that type, and the channels to deliver the events to.


## Supported Types

### BillingNotificationRuleBalanceThreshold

```go
billingNotificationRule := components.CreateBillingNotificationRuleEntitlementsBalanceThreshold(components.BillingNotificationRuleBalanceThreshold{/* values here */})
```

### BillingNotificationRuleEntitlementReset

```go
billingNotificationRule := components.CreateBillingNotificationRuleEntitlementsReset(components.BillingNotificationRuleEntitlementReset{/* values here */})
```

### BillingNotificationRuleInvoiceCreated

```go
billingNotificationRule := components.CreateBillingNotificationRuleInvoiceCreated(components.BillingNotificationRuleInvoiceCreated{/* values here */})
```

### BillingNotificationRuleInvoiceUpdated

```go
billingNotificationRule := components.CreateBillingNotificationRuleInvoiceUpdated(components.BillingNotificationRuleInvoiceUpdated{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingNotificationRule.Type {
	case components.BillingNotificationRuleTypeEntitlementsBalanceThreshold:
		// billingNotificationRule.BillingNotificationRuleBalanceThreshold is populated
	case components.BillingNotificationRuleTypeEntitlementsReset:
		// billingNotificationRule.BillingNotificationRuleEntitlementReset is populated
	case components.BillingNotificationRuleTypeInvoiceCreated:
		// billingNotificationRule.BillingNotificationRuleInvoiceCreated is populated
	case components.BillingNotificationRuleTypeInvoiceUpdated:
		// billingNotificationRule.BillingNotificationRuleInvoiceUpdated is populated
}
```
