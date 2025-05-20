# ReplaceAPIAttributesRequest


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `APIID`                                                     | *string*                                                    | :heavy_check_mark:                                          | The UUID API identifier                                     | 9f5061ce-78f6-4452-9108-ad7c02821fd5                        |
| `RequestBody`                                               | map[string][]*string*                                       | :heavy_check_mark:                                          | N/A                                                         | {<br/>"env": [<br/>"production"<br/>],<br/>"domains": [<br/>"web",<br/>"mobile"<br/>]<br/>} |