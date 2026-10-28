# UpdateCatalogMCPSourceSource


## Supported Types

### CatalogMCPGatewaySourceBlock

```go
updateCatalogMCPSourceSource := components.CreateUpdateCatalogMCPSourceSourceCatalogMCPGatewaySourceBlock(components.CatalogMCPGatewaySourceBlock{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateCatalogMCPSourceSource.Type {
	case components.UpdateCatalogMCPSourceSourceTypeCatalogMCPGatewaySourceBlock:
		// updateCatalogMCPSourceSource.CatalogMCPGatewaySourceBlock is populated
}
```
