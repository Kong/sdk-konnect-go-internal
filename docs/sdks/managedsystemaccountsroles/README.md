# ManagedSystemAccountsRoles
(*ManagedSystemAccountsRoles*)

## Overview

### Available Operations

* [GetSystemAccountsAssignedRolesInternal](#getsystemaccountsassignedrolesinternal) - Fetch assigned roles for a managed system account.
* [CreateSystemAccountsAssignedRolesInternal](#createsystemaccountsassignedrolesinternal) - Assign a role to a managed System Account.

## GetSystemAccountsAssignedRolesInternal

Lists the roles belonging to a managed system account.  Returns 400 if any filter parameters are invalid.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgointernal "github.com/Kong/sdk-konnect-go-internal"
	"github.com/Kong/sdk-konnect-go-internal/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgointernal.New(
        sdkkonnectgointernal.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgointernal.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ManagedSystemAccountsRoles.GetSystemAccountsAssignedRolesInternal(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ParameterizedAssignedRolesCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                               | Type                                                                                                                                                    | Required                                                                                                                                                | Description                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                   | :heavy_check_mark:                                                                                                                                      | The context to use for the request.                                                                                                                     |
| `accountID`                                                                                                                                             | *string*                                                                                                                                                | :heavy_check_mark:                                                                                                                                      | ID of the system account.                                                                                                                               |
| `filter`                                                                                                                                                | [*operations.GetSystemAccountsAssignedRolesInternalQueryParamFilter](../../models/operations/getsystemaccountsassignedrolesinternalqueryparamfilter.md) | :heavy_minus_sign:                                                                                                                                      | Filter roles returned in the response.                                                                                                                  |
| `opts`                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                | :heavy_minus_sign:                                                                                                                                      | The options for this request.                                                                                                                           |

### Response

**[*operations.GetSystemAccountsAssignedRolesInternalResponse](../../models/operations/getsystemaccountsassignedrolesinternalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateSystemAccountsAssignedRolesInternal

Assigns a role to a managed system account. Returns 409 if role is already assigned.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgointernal "github.com/Kong/sdk-konnect-go-internal"
	"github.com/Kong/sdk-konnect-go-internal/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgointernal.New(
        sdkkonnectgointernal.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgointernal.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ManagedSystemAccountsRoles.CreateSystemAccountsAssignedRolesInternal(ctx, "<id>", &components.AssignParameterizedRole{
        RoleName: sdkkonnectgointernal.String("Connector"),
        EntityID: sdkkonnectgointernal.String("e67490ce-44dc-4cbd-b65e-b52c746fc26a"),
        EntityTypeName: sdkkonnectgointernal.String("Mesh Control Planes"),
        EntityRegion: components.AssignParameterizedRoleEntityRegionEu.ToPointer(),
        Parameters: &components.Parameters{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ParameterizedAssignedRole != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `accountID`                                                                               | *string*                                                                                  | :heavy_check_mark:                                                                        | ID of the system account.                                                                 |
| `assignParameterizedRole`                                                                 | [*components.AssignParameterizedRole](../../models/components/assignparameterizedrole.md) | :heavy_minus_sign:                                                                        | The request schema for assigning a role.                                                  |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |

### Response

**[*operations.CreateSystemAccountsAssignedRolesInternalResponse](../../models/operations/createsystemaccountsassignedrolesinternalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401, 403                    | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |