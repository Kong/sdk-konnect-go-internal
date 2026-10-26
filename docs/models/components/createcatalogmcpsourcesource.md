# CreateCatalogMCPSourceSource


## Supported Types

### CatalogMCPGatewaySourceBlock

```go
createCatalogMCPSourceSource := components.CreateCreateCatalogMCPSourceSourceCatalogMCPGatewaySourceBlock(components.CatalogMCPGatewaySourceBlock{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createCatalogMCPSourceSource.Type {
	case components.CreateCatalogMCPSourceSourceTypeCatalogMCPGatewaySourceBlock:
		// createCatalogMCPSourceSource.CatalogMCPGatewaySourceBlock is populated
}
```
