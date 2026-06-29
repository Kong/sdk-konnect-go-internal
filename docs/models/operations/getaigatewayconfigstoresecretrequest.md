# GetAiGatewayConfigStoreSecretRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `GatewayID`                                           | `string`                                              | :heavy_check_mark:                                    | The unique ID of the AI Gateway.                      | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                  |
| `ConfigStoreID`                                       | `string`                                              | :heavy_check_mark:                                    | The unique ID or name of the AI Gateway Config Store. | bf138ba2-c9b1-4229-b268-04d9d8a6410b                  |
| `Key`                                                 | `string`                                              | :heavy_check_mark:                                    | Config Store Secret key.                              | my-secret-key                                         |