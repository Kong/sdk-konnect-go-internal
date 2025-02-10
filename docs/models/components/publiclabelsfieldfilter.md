# PublicLabelsFieldFilter

Filters on the given string field value by either exact or fuzzy match.


## Supported Types

### StringFieldEqualsFilter

```go
publicLabelsFieldFilter := components.CreatePublicLabelsFieldFilterStringFieldEqualsFilter(components.StringFieldEqualsFilter{/* values here */})
```

### StringFieldContainsFilter

```go
publicLabelsFieldFilter := components.CreatePublicLabelsFieldFilterStringFieldContainsFilter(components.StringFieldContainsFilter{/* values here */})
```

### StringFieldOContainsFilter

```go
publicLabelsFieldFilter := components.CreatePublicLabelsFieldFilterStringFieldOContainsFilter(components.StringFieldOContainsFilter{/* values here */})
```

### StringFieldOEQFilter

```go
publicLabelsFieldFilter := components.CreatePublicLabelsFieldFilterStringFieldOEQFilter(components.StringFieldOEQFilter{/* values here */})
```

### StringFieldNEQFilter

```go
publicLabelsFieldFilter := components.CreatePublicLabelsFieldFilterStringFieldNEQFilter(components.StringFieldNEQFilter{/* values here */})
```

