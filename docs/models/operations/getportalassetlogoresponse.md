# GetPortalAssetLogoResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ContentType`                                           | *string*                                                | :heavy_check_mark:                                      | HTTP response content type for this operation           |                                                         |
| `StatusCode`                                            | *int*                                                   | :heavy_check_mark:                                      | HTTP response status code for this operation            |                                                         |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_check_mark:                                      | Raw HTTP response; suitable for custom response parsing |                                                         |
| `TwoHundredImagePngPortalImageAssetV3`                  | *io.ReadCloser*                                         | :heavy_minus_sign:                                      | Logo of the portal. Can be either png or jpeg           | data:image/png;base64,bmljZV9sb29raW5nX3BpY3R1cmU=      |
| `TwoHundredImageJpegPortalImageAssetV3`                 | *io.ReadCloser*                                         | :heavy_minus_sign:                                      | Logo of the portal. Can be either png or jpeg           | data:image/png;base64,bmljZV9sb29raW5nX3BpY3R1cmU=      |