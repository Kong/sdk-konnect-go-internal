# ReplacePortalAssetFaviconResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ContentType`                                           | *string*                                                | :heavy_check_mark:                                      | HTTP response content type for this operation           |                                                         |
| `StatusCode`                                            | *int*                                                   | :heavy_check_mark:                                      | HTTP response status code for this operation            |                                                         |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_check_mark:                                      | Raw HTTP response; suitable for custom response parsing |                                                         |
| `TwoHundredImagePngPortalImageAsset`                    | *io.ReadCloser*                                         | :heavy_minus_sign:                                      | Logo of the portal. Can be either png, jpeg or svg      | data:image/png;base64,bmljZV9sb29raW5nX3BpY3R1cmU=      |
| `TwoHundredImageJpegPortalImageAsset`                   | *io.ReadCloser*                                         | :heavy_minus_sign:                                      | Logo of the portal. Can be either png, jpeg or svg      | data:image/png;base64,bmljZV9sb29raW5nX3BpY3R1cmU=      |
| `TwoHundredImageSvgPlusXMLPortalImageAsset`             | *io.ReadCloser*                                         | :heavy_minus_sign:                                      | Logo of the portal. Can be either png, jpeg or svg      | data:image/png;base64,bmljZV9sb29raW5nX3BpY3R1cmU=      |