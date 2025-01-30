# Granularity

`granularity` is only valid for queries that include a time dimension, and it specifies the time buckets for the returned data.  For example, `MINUTELY` granularity will return datapoints for every minute.  Not all granularities are available for all time ranges: for example, custom timeframes only have `DAILY` granularity.

If unspecified, a default value for the given time range will be chosen according to the following table:

- `FIFTEEN_MIN`: `MINUTELY`
- `ONE_HOUR`: `MINUTELY`
- `SIX_HOUR`: `HOURLY`
- `TWELVE_HOUR`: `HOURLY`
- `ONE_DAY`: `HOURLY`
- `SEVEN_DAY`: `DAILY`
- `THIRTY_DAY`: `DAILY`
- `CURRENT_WEEK`: `DAILY`
- `CURRENT_MONTH`: `DAILY`
- `PREVIOUS_WEEK`: `DAILY`
- `PREVIOUS_MONTH`: `DAILY`



## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `GranularityMinutely` | MINUTELY              |
| `GranularityHourly`   | HOURLY                |
| `GranularityDaily`    | DAILY                 |
| `GranularityWeekly`   | WEEKLY                |