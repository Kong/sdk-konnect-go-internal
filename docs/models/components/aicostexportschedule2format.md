# AICostExportSchedule2Format

The file format of each export's data files.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AICostExportSchedule2FormatParquet

// Open enum: custom values can be created with a direct type cast
custom := components.AICostExportSchedule2Format("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AICostExportSchedule2FormatParquet` | parquet                              |
| `AICostExportSchedule2FormatCsv`     | csv                                  |