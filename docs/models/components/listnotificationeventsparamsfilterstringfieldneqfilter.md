# ListNotificationEventsParamsFilterStringFieldNEQFilter

Filter by delivery state. Matches events where delivery to at least one channel
is in one of the given states. Only `eq` and `oeq` are supported.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Eq`               | `*string`          | :heavy_minus_sign: | N/A                |
| `Oeq`              | `string`           | :heavy_check_mark: | N/A                |
| `Neq`              | `string`           | :heavy_check_mark: | N/A                |