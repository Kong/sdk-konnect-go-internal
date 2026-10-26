# ValidatePluginConditionRequestBody


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Condition`                                                   | `string`                                                      | :heavy_check_mark:                                            | The CEL condition to validate. It must evaluate to a boolean. | http.method == "GET"                                          |