# CustomReports
(*CustomReports*)

## Overview

### Available Operations

* [GetReports](#getreports) - List Custom Reports
* [CreateReport](#createreport) - New Custom Report
* [GetReport](#getreport) - Fetch Custom Report
* [UpdateReport](#updatereport) - Update Report
* [DeleteReport](#deletereport) - Delete Custom Report
* [ExportCsv](#exportcsv) - Export as CSV

## GetReports

Returns a list of custom reports.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CustomReports.GetReports(ctx, sdkkonnectgo.Int64(10), sdkkonnectgo.Int64(2))
    if err != nil {
        log.Fatal(err)
    }
    if res.ReportCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `pageSize`                                                                                              | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | The maximum number of items to include per page. The last page of a collection may include fewer items. | 10                                                                                                      |
| `pageNumber`                                                                                            | **int64*                                                                                                | :heavy_minus_sign:                                                                                      | The specific page number in the collection results.                                                     | 2                                                                                                       |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.GetReportsResponse](../../models/operations/getreportsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.PermissionDenied | 403                        | application/problem+json   |
| sdkerrors.TooManyRequests  | 429                        | application/problem+json   |
| sdkerrors.SDKError         | 4XX, 5XX                   | \*/\*                      |

## CreateReport

Creates a new custom report.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CustomReports.CreateReport(ctx, &components.ReportInput{
        Description: sdkkonnectgo.String("A vertical bar chart grouping data first by API product, then by status code, for the past 24 hours."),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Report != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.ReportInput](../../models/components/reportinput.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.CreateReportResponse](../../models/operations/createreportresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.BadRequestError  | 400                        | application/problem+json   |
| sdkerrors.PermissionDenied | 403                        | application/problem+json   |
| sdkerrors.TooManyRequests  | 429                        | application/problem+json   |
| sdkerrors.SDKError         | 4XX, 5XX                   | \*/\*                      |

## GetReport

Returns a single report.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CustomReports.GetReport(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.Report != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `reportID`                                               | *string*                                                 | :heavy_check_mark:                                       | The report's ID                                          | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetReportResponse](../../models/operations/getreportresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.BadRequestError  | 400                        | application/problem+json   |
| sdkerrors.PermissionDenied | 403                        | application/problem+json   |
| sdkerrors.NotFoundError    | 404                        | application/problem+json   |
| sdkerrors.TooManyRequests  | 429                        | application/problem+json   |
| sdkerrors.SDKError         | 4XX, 5XX                   | \*/\*                      |

## UpdateReport

Updates a custom report.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CustomReports.UpdateReport(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", &components.ReportInput{
        Description: sdkkonnectgo.String("A vertical bar chart grouping data first by API product, then by status code, for the past 24 hours."),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Report != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                     | Type                                                                                                                                                                                                                                                                                                                                                                                                                                          | Required                                                                                                                                                                                                                                                                                                                                                                                                                                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                   | Example                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                            | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `reportID`                                                                                                                                                                                                                                                                                                                                                                                                                                    | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                            | The report's ID                                                                                                                                                                                                                                                                                                                                                                                                                               | d32d905a-ed33-46a3-a093-d8f536af9a8a                                                                                                                                                                                                                                                                                                                                                                                                          |
| `report`                                                                                                                                                                                                                                                                                                                                                                                                                                      | [*components.ReportInput](../../models/components/reportinput.md)                                                                                                                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                            | The request schema for the update report request.                                                                                                                                                                                                                                                                                                                                                                                             | {<br/>"summary": "Bar chart of status codes by API product for the past 24 hours",<br/>"description": "A vertical bar chart grouping data first by API product, then by status code, for the past 24 hours.",<br/>"value": {<br/>"name": "Status codes by service",<br/>"chart_type": "VERTICAL_BAR",<br/>"time_range": {<br/>"type": "RELATIVE",<br/>"time_period": "24H"<br/>},<br/>"query": {<br/>"metrics": [<br/>"REQUEST_COUNT"<br/>],<br/>"dimensions": [<br/>"API_PRODUCT",<br/>"STATUS_CODE"<br/>]<br/>}<br/>}<br/>} |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                            | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                 |                                                                                                                                                                                                                                                                                                                                                                                                                                               |

### Response

**[*operations.UpdateReportResponse](../../models/operations/updatereportresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.BadRequestError  | 400                        | application/problem+json   |
| sdkerrors.PermissionDenied | 403                        | application/problem+json   |
| sdkerrors.NotFoundError    | 404                        | application/problem+json   |
| sdkerrors.TooManyRequests  | 429                        | application/problem+json   |
| sdkerrors.SDKError         | 4XX, 5XX                   | \*/\*                      |

## DeleteReport

Deletes a custom report. Returns 404 if the report is not found.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CustomReports.DeleteReport(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `reportID`                                               | *string*                                                 | :heavy_check_mark:                                       | The report's ID                                          | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteReportResponse](../../models/operations/deletereportresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.BadRequestError  | 400                        | application/problem+json   |
| sdkerrors.PermissionDenied | 403                        | application/problem+json   |
| sdkerrors.NotFoundError    | 404                        | application/problem+json   |
| sdkerrors.TooManyRequests  | 429                        | application/problem+json   |
| sdkerrors.SDKError         | 4XX, 5XX                   | \*/\*                      |

## ExportCsv

Exports report data in CSV format.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.CustomReports.ExportCsv(ctx, operations.ExportCsvRequest{
        ReportID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        Tz: sdkkonnectgo.String("Etc/UTC"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ExportCsvRequest](../../models/operations/exportcsvrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.ExportCsvResponse](../../models/operations/exportcsvresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.BadRequestError  | 400                        | application/problem+json   |
| sdkerrors.PermissionDenied | 403                        | application/problem+json   |
| sdkerrors.NotFoundError    | 404                        | application/problem+json   |
| sdkerrors.TooManyRequests  | 429                        | application/problem+json   |
| sdkerrors.SDKError         | 4XX, 5XX                   | \*/\*                      |