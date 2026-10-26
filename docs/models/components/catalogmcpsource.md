# CatalogMCPSource

An MCP source resource.


## Supported Types

### CatalogMCPGatewaySource

```go
catalogMCPSource := components.CreateCatalogMCPSourceCatalogMCPGatewaySource(components.CatalogMCPGatewaySource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch catalogMCPSource.Type {
	case components.CatalogMCPSourceTypeCatalogMCPGatewaySource:
		// catalogMCPSource.CatalogMCPGatewaySource is populated
}
```
