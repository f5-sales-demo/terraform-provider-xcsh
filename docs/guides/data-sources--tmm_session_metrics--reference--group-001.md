---
page_title: "xcsh_tmm_session_metrics reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tmm_session_metrics reference."
---

# xcsh_tmm_session_metrics reference

<a id="canonical-527d0176eb1fcf0d1ab47bc9f224535c861a92810150e2e9a25333b1b6c3fe84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-447f725647e10a4bc0ab529243adbee3f84bc7a6d2a0485a901f23e819c95fd7"></a>

## Property reference — Property reference / 93cb9827ed5a / 2

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)
- Property reference

<a id="canonical-6331918f1a3407569fb03c218b82b271761f7296c90b5fdfc1dd417f1939bfb1"></a>

## Direct properties — Property reference / 93cb9827ed5a / 3

- [data](data-sources--tmm_session_metrics--reference--group-001.md#canonical-6071d66977f819cd88722bc95ed4f96c00fe7f1e44cb22459f7e9c5bab25a344): complete subsection reference.

<a id="canonical-9c72fcd61e192c759e5faa3053b12e0fccf22593151fcd4aaaf332a2d2f7c3bb"></a>

<a id="canonical-04a7862b1880e3af7956a04cd23b4cc8c04a4b6863c9481a12999816b0bb37b3"></a>

## end_time property — Property reference / 93cb9827ed5a / 4

Type: `"string"`. Optional.

End time of metric collection from which data will be considered to build dashboard. Format:
unix\_timestamp|RFC 3339 Optional: If not specified, then the end\_time will be evaluated to
start\_time+10m If start\_time is not specified, then the end\_time will be evaluated to &lt;current
time&gt;.

<a id="canonical-9214bfa30326b1646a5dcc0c02c3f0e8f4eb93e2dd123c17f295bfd1ef687544"></a>

<a id="canonical-67bb283c659097a0fca5b2d6952967febf3d2d31d8dfb090e04a48161eac48b1"></a>

## field_selector property — Property reference / 93cb9827ed5a / 5

Type: `["list", "string"]`. Optional.

\[Enum:
METRIC\_TYPE\_NONE|METRIC\_TYPE\_ACTIVE|METRIC\_TYPE\_ALLOWED|METRIC\_TYPE\_DENIED|METRIC\_TYPE\_TOTAL|METRIC\_TYPE\_LOGOUT|METRIC\_TYPE\_ESTABLISHED\_TIMEOUT|METRIC\_TYPE\_EVALUATION\_TIMEOUT|METRIC\_TYPE\_ADMIN\_TERMINATED\]
Select fields to be returned in the response. Field\_selector is used to specify the fields to be
returned in the response, thereby limiting the amount of data returned in the response. Possible
values are \`METRIC\_TYPE\_NONE\`, \`METRIC\_TYPE\_ACTIVE\`, \`METRIC\_TYPE\_ALLOWED\`,
\`METRIC\_TYPE\_DENIED\`, \`METRIC\_TYPE\_TOTAL\`, \`METRIC\_TYPE\_LOGOUT\`,
\`METRIC\_TYPE\_ESTABLISHED\_TIMEOUT\`, \`METRIC\_TYPE\_EVALUATION\_TIMEOUT\`,
\`METRIC\_TYPE\_ADMIN\_TERMINATED\`. Defaults to \`METRIC\_TYPE\_NONE\`.

<a id="canonical-843f0b52915f2ff6161cac1da81972db9858ff5d1fa59121a9a16949418b7591"></a>

<a id="canonical-94483d0dedad04745f74059ea8a9de9378e4e631fe540fabf23d7e7ef91181b4"></a>

## group_by property — Property reference / 93cb9827ed5a / 6

Type: `["list", "string"]`. Optional.

\[Enum: METRIC\_LABEL\_NONE|METRIC\_LABEL\_NAMESPACE|METRIC\_LABEL\_VIRTUAL\_SERVER\] Aggregate data
by any or all of namespace, site and virtual\_server. Optional: If not specified, then the data is
aggregated/grouped by namespace and service. Possible values are \`METRIC\_LABEL\_NONE\`,
\`METRIC\_LABEL\_NAMESPACE\`, \`METRIC\_LABEL\_VIRTUAL\_SERVER\`. Defaults to
\`METRIC\_LABEL\_NONE\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

<a id="canonical-cf1498bb5997d9f010ec44eba76803b8760205d3430e227c208e7ef5eece593b"></a>

<a id="canonical-92ef9e192979d266cd9e288b89be41a9142ec17f6dd00ea4cde574cea170ae2e"></a>

## is_trend_request property — Property reference / 93cb9827ed5a / 7

Type: `"bool"`. Optional.

Trend value computation requested by the user Optional:. Defaults to \`false\`.

- [label_filter](data-sources--tmm_session_metrics--reference--group-001.md#canonical-4401ddc79888b54eaf42c7ca558018746a8ccaecd10f4578cdd4f5e25e000c76): complete subsection reference.

<a id="canonical-f6808e34ec2d3b3458e4a71f7f138bbc54812a8a28d3f981c9dc70e2dde6c8e9"></a>

<a id="canonical-33b342d7dd1cccf971a5a596a748470096778fb9165f8edb5a1d50eebbf34243"></a>

## namespace property — Property reference / 93cb9827ed5a / 8

Type: `"string"`. Required.

Namespace namespace is used to scope session metrics. Only virtual server in given namespace will be
considered.

<a id="canonical-3aefaf5452479d1f6991dfc641af2ed818b2ef318ff7cca99746d3d63a4e6a60"></a>

<a id="canonical-174b2d8724880a71f927a94af199a9d309bc0546ab72abb1d8e7a1c0029a0d20"></a>

## start_time property — Property reference / 93cb9827ed5a / 9

Type: `"string"`. Optional.

Start time of metric collection from which data will be considered to build metrics dashboard.
Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the start\_time will be evaluated
to end\_time-10m If end\_time is not specified, then the start\_time will be evaluated to
&lt;current time&gt;-10m.

<a id="canonical-67fde01179d0b5535da33c8b80a753d5fc2bbde80c286195faecf970f83b84a1"></a>

<a id="canonical-411babb8a8ba6f2abfcdb7eeb8590f5f091f0af04e1d0879a98da99824c76ebd"></a>

## step property — Property reference / 93cb9827ed5a / 10

Type: `"string"`. Optional.

Step is the resolution width, which determines the number of the data points \[x-axis (time)\] to be
returned in the response. The timestamps in the response will be t1=start\_time, t2=t1+step, ..
Tn=tn-1+step, where tn &lt;= end\_time.

<a id="canonical-3a4f5613b13ec766f754a915a45124788d1cccbf9dd680e7dc860836789bce15"></a>

## All schema paths — Property reference / 93cb9827ed5a / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `data` | [data](data-sources--tmm_session_metrics--reference--group-001.md#canonical-0d80eb0701c813f8f5346b1178a07a7aaee56f1e81798993a03b302021a95ff4) |
| `data.metric` | [data.metric](data-sources--tmm_session_metrics--reference--group-001.md#canonical-11489bf91bf897c0c394093762607d843828dd21c3907271c748d8ddfd7a2add) |
| `data.metric.key` | [data.metric.key](data-sources--tmm_session_metrics--reference--group-001.md#canonical-946cf6f2734ec0e8ab9010a5aab9160ffe17e6ce46453a08875bb2b342fb3dee) |
| `data.metric.value` | [data.metric.value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-e1203cc2706d5e0221fe801d59f0f9d9d818b19dec1d2f327eb7afcf25965b8b) |
| `data.metric.value.timestamp` | [data.metric.value.timestamp](data-sources--tmm_session_metrics--reference--group-001.md#canonical-dbe169cc1791d08f5abcb479fdb5876e19ea1cecef7477b88e8c23e89698777e) |
| `data.metric.value.trend_value` | [data.metric.value.trend_value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-90dfb62a6bcdf8ead41e0a0c4ee783c611c1f93e278fb23bcb123b29cb7f4a1b) |
| `data.metric.value.trend_value.description_spec` | [data.metric.value.trend_value.description_spec](data-sources--tmm_session_metrics--reference--group-001.md#canonical-ce4b3d264c465b4290810a960b8ea83edb93a2caf405cbace0e54110019cb20c) |
| `data.metric.value.trend_value.previous_value` | [data.metric.value.trend_value.previous_value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-93fd207092664ec499ae16ea4fc4c70bc6a931ca398d436b8734253e5ddb3a27) |
| `data.metric.value.trend_value.sentiment` | [data.metric.value.trend_value.sentiment](data-sources--tmm_session_metrics--reference--group-001.md#canonical-b8e6f1b9cd6817d55e0c9f5caf8932305b3b235d305e834637675d9aed7903b7) |
| `data.metric.value.trend_value.value` | [data.metric.value.trend_value.value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-c6e83257f43f640a191c068287b68837682559d898ac9bd92b776e8bbfbaabaa) |
| `data.metric.value.value` | [data.metric.value.value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-1fbb39d0e5ca98baa69e783a4efb07d9455f8c61d17e5b1c44d7058788fa47f1) |
| `data.type` | [data.type](data-sources--tmm_session_metrics--reference--group-001.md#canonical-18d00233513d513e137f4881c2f1333340475788443cd9175d6cc81f75d3a7ec) |
| `data.unit` | [data.unit](data-sources--tmm_session_metrics--reference--group-001.md#canonical-5d3b27d4f3352ca3682fe8c7f7d16ec3a11ec17afcd2d74ac07d0d41e3d87551) |
| `end_time` | [end_time](data-sources--tmm_session_metrics--reference--group-001.md#canonical-9c72fcd61e192c759e5faa3053b12e0fccf22593151fcd4aaaf332a2d2f7c3bb) |
| `field_selector` | [field_selector](data-sources--tmm_session_metrics--reference--group-001.md#canonical-9214bfa30326b1646a5dcc0c02c3f0e8f4eb93e2dd123c17f295bfd1ef687544) |
| `group_by` | [group_by](data-sources--tmm_session_metrics--reference--group-001.md#canonical-843f0b52915f2ff6161cac1da81972db9858ff5d1fa59121a9a16949418b7591) |
| `is_trend_request` | [is_trend_request](data-sources--tmm_session_metrics--reference--group-001.md#canonical-cf1498bb5997d9f010ec44eba76803b8760205d3430e227c208e7ef5eece593b) |
| `label_filter` | [label_filter](data-sources--tmm_session_metrics--reference--group-001.md#canonical-6c8afc2bac643256e9e8f0a2f01d7c5b72bab654c17c40b46055ca8b6ea37877) |
| `label_filter.label` | [label_filter.label](data-sources--tmm_session_metrics--reference--group-001.md#canonical-28952a15a4dd2b39c6d521f55417bbc5792e9f53d97bd57eeaecf3c8ddf14cae) |
| `label_filter.op` | [label_filter.op](data-sources--tmm_session_metrics--reference--group-001.md#canonical-180cec6c5cadbe04944e427acd6f006f3e04b674b0ac740b612e2f694cfbe881) |
| `label_filter.value` | [label_filter.value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-34ad5e62dc66796d741ef1e6e4680ff9f4d733ef338db14e654f2ee2abad302b) |
| `namespace` | [namespace](data-sources--tmm_session_metrics--reference--group-001.md#canonical-f6808e34ec2d3b3458e4a71f7f138bbc54812a8a28d3f981c9dc70e2dde6c8e9) |
| `start_time` | [start_time](data-sources--tmm_session_metrics--reference--group-001.md#canonical-3aefaf5452479d1f6991dfc641af2ed818b2ef318ff7cca99746d3d63a4e6a60) |
| `step` | [step](data-sources--tmm_session_metrics--reference--group-001.md#canonical-67fde01179d0b5535da33c8b80a753d5fc2bbde80c286195faecf970f83b84a1) |

<a id="canonical-535da31188531792613e9713863e3598f3a2bbbcd2be76569c15c8fc24a7ff8b"></a>

## Next pages — Property reference / 93cb9827ed5a / 12

- [data](data-sources--tmm_session_metrics--reference--group-001.md#canonical-6071d66977f819cd88722bc95ed4f96c00fe7f1e44cb22459f7e9c5bab25a344)
- [label_filter](data-sources--tmm_session_metrics--reference--group-001.md#canonical-4401ddc79888b54eaf42c7ca558018746a8ccaecd10f4578cdd4f5e25e000c76)
- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)

<a id="canonical-6071d66977f819cd88722bc95ed4f96c00fe7f1e44cb22459f7e9c5bab25a344"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58e329b3eb69cad335f40fa600ba669ba2d28f5e332a64160c90541bf7a35fa2"></a>

## data — data / 92026aefaecc / 2

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)
- [Property reference](data-sources--tmm_session_metrics--reference--group-001.md#canonical-527d0176eb1fcf0d1ab47bc9f224535c861a92810150e2e9a25333b1b6c3fe84)
- data

<a id="canonical-0d80eb0701c813f8f5346b1178a07a7aaee56f1e81798993a03b302021a95ff4"></a>

Type: `"list"`. Computed.

Data contains time-series TMM Session data.

<a id="canonical-acd606a3ca59de80ea4b5f39a283974f3a5f2288adb1924bd39fb0ef30508a42"></a>

## Direct properties — data / 92026aefaecc / 3

- [metric](data-sources--tmm_session_metrics--reference--group-001.md#canonical-8a6cc9aceb7e3819db97af404bd3dac0aa475effbc80dc18eeac764f713fcfd0): complete subsection reference.

<a id="canonical-18d00233513d513e137f4881c2f1333340475788443cd9175d6cc81f75d3a7ec"></a>

<a id="canonical-55452aa6fc90f117d4b07c65d401bc768b2415f3130c7d2623e3b3ea430abdf5"></a>

## type property — data / 92026aefaecc / 4

Type: `"string"`. Computed.

\[Enum:
METRIC\_TYPE\_NONE|METRIC\_TYPE\_ACTIVE|METRIC\_TYPE\_ALLOWED|METRIC\_TYPE\_DENIED|METRIC\_TYPE\_TOTAL|METRIC\_TYPE\_LOGOUT|METRIC\_TYPE\_ESTABLISHED\_TIMEOUT|METRIC\_TYPE\_EVALUATION\_TIMEOUT|METRIC\_TYPE\_ADMIN\_TERMINATED\]
X-displayName: TMM Session Metric Type' FieldSelector specifies the metrics that can be queried for
virtual servers. Indicates field not being set x-unit: 'count' Total number of active sessions
x-unit: 'count' Total number of allowed sessions x-unit: 'count' Total number of denied Session..
Possible values are \`METRIC\_TYPE\_NONE\`, \`METRIC\_TYPE\_ACTIVE\`, \`METRIC\_TYPE\_ALLOWED\`,
\`METRIC\_TYPE\_DENIED\`, \`METRIC\_TYPE\_TOTAL\`, \`METRIC\_TYPE\_LOGOUT\`,
\`METRIC\_TYPE\_ESTABLISHED\_TIMEOUT\`, \`METRIC\_TYPE\_EVALUATION\_TIMEOUT\`,
\`METRIC\_TYPE\_ADMIN\_TERMINATED\`. Defaults to \`METRIC\_TYPE\_NONE\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("METRIC_TYPE_NONE",
    "METRIC_TYPE_ACTIVE",
    "METRIC_TYPE_ALLOWED",
    "METRIC_TYPE_DENIED",
    "METRIC_TYPE_TOTAL",
    "METRIC_TYPE_LOGOUT",
    "METRIC_TYPE_ESTABLISHED_TIMEOUT",
    "METRIC_TYPE_EVALUATION_TIMEOUT",
    "METRIC_TYPE_ADMIN_TERMINATED"),
}
```

<a id="canonical-5d3b27d4f3352ca3682fe8c7f7d16ec3a11ec17afcd2d74ac07d0d41e3d87551"></a>

<a id="canonical-1fb7aa4890e84bfae2c45ef3be64b875a96729191177357e694b9864eca40520"></a>

## unit property — data / 92026aefaecc / 5

Type: `"string"`. Computed.

\[Enum:
UNIT\_MILLISECONDS|UNIT\_SECONDS|UNIT\_MINUTES|UNIT\_HOURS|UNIT\_DAYS|UNIT\_BYTES|UNIT\_KBYTES|UNIT\_MBYTES|UNIT\_GBYTES|UNIT\_TBYTES|UNIT\_KIBIBYTES|UNIT\_MIBIBYTES|UNIT\_GIBIBYTES|UNIT\_TEBIBYTES|UNIT\_BITS\_PER\_SECOND|UNIT\_BYTES\_PER\_SECOND|UNIT\_KBITS\_PER\_SECOND|UNIT\_KBYTES\_PER\_SECOND|UNIT\_MBITS\_PER\_SECOND|UNIT\_MBYTES\_PER\_SECOND|UNIT\_CONNECTIONS\_PER\_SECOND|UNIT\_ERRORS\_PER\_SECOND|UNIT\_PACKETS\_PER\_SECOND|UNIT\_REQUESTS\_PER\_SECOND|UNIT\_PACKETS|UNIT\_PERCENTAGE|UNIT\_COUNT\]
UnitType is enumeration of units for scalar fields. Possible values are \`UNIT\_MILLISECONDS\`,
\`UNIT\_SECONDS\`, \`UNIT\_MINUTES\`, \`UNIT\_HOURS\`, \`UNIT\_DAYS\`, \`UNIT\_BYTES\`,
\`UNIT\_KBYTES\`, \`UNIT\_MBYTES\`, \`UNIT\_GBYTES\`, \`UNIT\_TBYTES\`, \`UNIT\_KIBIBYTES\`,
\`UNIT\_MIBIBYTES\`, \`UNIT\_GIBIBYTES\`, \`UNIT\_TEBIBYTES\`, \`UNIT\_BITS\_PER\_SECOND\`,
\`UNIT\_BYTES\_PER\_SECOND\`, \`UNIT\_KBITS\_PER\_SECOND\`, \`UNIT\_KBYTES\_PER\_SECOND\`,
\`UNIT\_MBITS\_PER\_SECOND\`, \`UNIT\_MBYTES\_PER\_SECOND\`, \`UNIT\_CONNECTIONS\_PER\_SECOND\`,
\`UNIT\_ERRORS\_PER\_SECOND\`, \`UNIT\_PACKETS\_PER\_SECOND\`, \`UNIT\_REQUESTS\_PER\_SECOND\`,
\`UNIT\_PACKETS\`, \`UNIT\_PERCENTAGE\`, \`UNIT\_COUNT\`. Defaults to \`UNIT\_MILLISECONDS\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNIT_MILLISECONDS",
    "UNIT_SECONDS",
    "UNIT_MINUTES",
    "UNIT_HOURS",
    "UNIT_DAYS",
    "UNIT_BYTES",
    "UNIT_KBYTES",
    "UNIT_MBYTES",
    "UNIT_GBYTES",
    "UNIT_TBYTES",
    "UNIT_KIBIBYTES",
    "UNIT_MIBIBYTES",
    "UNIT_GIBIBYTES",
    "UNIT_TEBIBYTES",
    "UNIT_BITS_PER_SECOND",
    "UNIT_BYTES_PER_SECOND",
    "UNIT_KBITS_PER_SECOND",
    "UNIT_KBYTES_PER_SECOND",
    "UNIT_MBITS_PER_SECOND",
    "UNIT_MBYTES_PER_SECOND",
    "UNIT_CONNECTIONS_PER_SECOND",
    "UNIT_ERRORS_PER_SECOND",
    "UNIT_PACKETS_PER_SECOND",
    "UNIT_REQUESTS_PER_SECOND",
    "UNIT_PACKETS",
    "UNIT_PERCENTAGE",
    "UNIT_COUNT"),
}
```

<a id="canonical-8eec9214b0a3ea7a739b53f0e9c9c21f9bbaeade79e6925b7cd2ad0add5cb07d"></a>

## Next pages — data / 92026aefaecc / 6

- [data.metric](data-sources--tmm_session_metrics--reference--group-001.md#canonical-8a6cc9aceb7e3819db97af404bd3dac0aa475effbc80dc18eeac764f713fcfd0)
- [Property reference](data-sources--tmm_session_metrics--reference--group-001.md#canonical-527d0176eb1fcf0d1ab47bc9f224535c861a92810150e2e9a25333b1b6c3fe84)
- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)

<a id="canonical-8a6cc9aceb7e3819db97af404bd3dac0aa475effbc80dc18eeac764f713fcfd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-605a29b202526b8bcb783f3a05d5e913845802414bfe6b26e6491ee9761b443f"></a>

## data.metric — data.metric / 83e4e916879b / 2

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)
- [Property reference](data-sources--tmm_session_metrics--reference--group-001.md#canonical-527d0176eb1fcf0d1ab47bc9f224535c861a92810150e2e9a25333b1b6c3fe84)
- [data](data-sources--tmm_session_metrics--reference--group-001.md#canonical-6071d66977f819cd88722bc95ed4f96c00fe7f1e44cb22459f7e9c5bab25a344)
- data.metric

<a id="canonical-11489bf91bf897c0c394093762607d843828dd21c3907271c748d8ddfd7a2add"></a>

Type: `"list"`. Computed.

Metric. List of metrics.

<a id="canonical-0bf1e13971c4a0e9d91b7fa85713fd17bed345a850f76b953f72270c54e86682"></a>

## Direct properties — data.metric / 83e4e916879b / 3

- [key](data-sources--tmm_session_metrics--reference--group-001.md#canonical-e4d05a37f4d83ea26e4acc3284be1484d9222b6000946976aba64b320accc2eb): complete subsection reference.

- [value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-7ad82d2e1b9a6a9e6cadbf28d029c4513dbdd9278670f13126af7f039e73188f): complete subsection reference.

<a id="canonical-9d9d8b9cbf44a41e92331b44e545f49f0aca251c07c8ceef2ff902017048a681"></a>

## Next pages — data.metric / 83e4e916879b / 4

- [data.metric.key](data-sources--tmm_session_metrics--reference--group-001.md#canonical-e4d05a37f4d83ea26e4acc3284be1484d9222b6000946976aba64b320accc2eb)
- [data.metric.value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-7ad82d2e1b9a6a9e6cadbf28d029c4513dbdd9278670f13126af7f039e73188f)
- [data](data-sources--tmm_session_metrics--reference--group-001.md#canonical-6071d66977f819cd88722bc95ed4f96c00fe7f1e44cb22459f7e9c5bab25a344)
- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)

<a id="canonical-e4d05a37f4d83ea26e4acc3284be1484d9222b6000946976aba64b320accc2eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c3631071efd1113aa2b18b86c79ebb4ec0be9789984562ea1d869d024972f50"></a>

## data.metric.key — data.metric.key / 83d97f18d562 / 2

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)
- [Property reference](data-sources--tmm_session_metrics--reference--group-001.md#canonical-527d0176eb1fcf0d1ab47bc9f224535c861a92810150e2e9a25333b1b6c3fe84)
- [data](data-sources--tmm_session_metrics--reference--group-001.md#canonical-6071d66977f819cd88722bc95ed4f96c00fe7f1e44cb22459f7e9c5bab25a344)
- [data.metric](data-sources--tmm_session_metrics--reference--group-001.md#canonical-8a6cc9aceb7e3819db97af404bd3dac0aa475effbc80dc18eeac764f713fcfd0)
- data.metric.key

<a id="canonical-946cf6f2734ec0e8ab9010a5aab9160ffe17e6ce46453a08875bb2b342fb3dee"></a>

Type: `"single"`. Computed.

Key contains the name/value pair. 'name' is the label name defined in 'MetricLabel'.

<a id="canonical-2aca0480c4928dbafcab43fcf4925db8e94e124ae2ccec7227f0b0cbb57716c6"></a>

## Direct properties — data.metric.key / 83d97f18d562 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e565a778116d017eaba5ef37586915c3c67b6083251ee123306c34f07084346d"></a>

## Next pages — data.metric.key / 83d97f18d562 / 4

- [data.metric](data-sources--tmm_session_metrics--reference--group-001.md#canonical-8a6cc9aceb7e3819db97af404bd3dac0aa475effbc80dc18eeac764f713fcfd0)
- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)

<a id="canonical-7ad82d2e1b9a6a9e6cadbf28d029c4513dbdd9278670f13126af7f039e73188f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85569d1a0cb6cd43368afb266e986a02de65feca21c75ccbe813da1d1017597a"></a>

## data.metric.value — data.metric.value / 810eca911cff / 2

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)
- [Property reference](data-sources--tmm_session_metrics--reference--group-001.md#canonical-527d0176eb1fcf0d1ab47bc9f224535c861a92810150e2e9a25333b1b6c3fe84)
- [data](data-sources--tmm_session_metrics--reference--group-001.md#canonical-6071d66977f819cd88722bc95ed4f96c00fe7f1e44cb22459f7e9c5bab25a344)
- [data.metric](data-sources--tmm_session_metrics--reference--group-001.md#canonical-8a6cc9aceb7e3819db97af404bd3dac0aa475effbc80dc18eeac764f713fcfd0)
- data.metric.value

<a id="canonical-e1203cc2706d5e0221fe801d59f0f9d9d818b19dec1d2f327eb7afcf25965b8b"></a>

Type: `"list"`. Computed.

Value. List of metric values.

<a id="canonical-398af823ce51c55e97ed8c41b4d805c95fceb0797a1eb3a4fc2b1803cef6ce84"></a>

## Direct properties — data.metric.value / 810eca911cff / 3

<a id="canonical-dbe169cc1791d08f5abcb479fdb5876e19ea1cecef7477b88e8c23e89698777e"></a>

<a id="canonical-7b26617254201c12c214686219adcbfd38958877ace6e18d0b2c21139192cae2"></a>

## timestamp property — data.metric.value / 810eca911cff / 4

Type: `"number"`. Computed.

Timestamp. Timestamp

- [trend_value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-a5f3143dd9224006057c6a9b38292f562db420d49f79a98e7c3ff54afe99a009): complete subsection reference.

<a id="canonical-1fbb39d0e5ca98baa69e783a4efb07d9455f8c61d17e5b1c44d7058788fa47f1"></a>

<a id="canonical-629ce7bb07abfd4b5b5766a05f9d9c4aca062b3d54d1fa4fa99034d4e38a41e1"></a>

## value property — data.metric.value / 810eca911cff / 5

Type: `"string"`. Computed.

Value. Configuration parameter for value

<a id="canonical-190bc04f7236f10fa9611a55c36200d7046fe60288017c76e3749c5193d867dd"></a>

## Next pages — data.metric.value / 810eca911cff / 6

- [data.metric.value.trend_value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-a5f3143dd9224006057c6a9b38292f562db420d49f79a98e7c3ff54afe99a009)
- [data.metric](data-sources--tmm_session_metrics--reference--group-001.md#canonical-8a6cc9aceb7e3819db97af404bd3dac0aa475effbc80dc18eeac764f713fcfd0)
- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)

<a id="canonical-a5f3143dd9224006057c6a9b38292f562db420d49f79a98e7c3ff54afe99a009"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e45922073076fba48445e78a86f66ca0574573de39d5caccdf47edddb7ad536"></a>

## data.metric.value.trend_value — data.metric.value.trend_value / 41fc8dade696 / 2

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)
- [Property reference](data-sources--tmm_session_metrics--reference--group-001.md#canonical-527d0176eb1fcf0d1ab47bc9f224535c861a92810150e2e9a25333b1b6c3fe84)
- [data](data-sources--tmm_session_metrics--reference--group-001.md#canonical-6071d66977f819cd88722bc95ed4f96c00fe7f1e44cb22459f7e9c5bab25a344)
- [data.metric](data-sources--tmm_session_metrics--reference--group-001.md#canonical-8a6cc9aceb7e3819db97af404bd3dac0aa475effbc80dc18eeac764f713fcfd0)
- [data.metric.value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-7ad82d2e1b9a6a9e6cadbf28d029c4513dbdd9278670f13126af7f039e73188f)
- data.metric.value.trend_value

<a id="canonical-90dfb62a6bcdf8ead41e0a0c4ee783c611c1f93e278fb23bcb123b29cb7f4a1b"></a>

Type: `"single"`. Computed.

Trend value contains trend value, trend sentiment and trend calculation description and window size.

<a id="canonical-400353f4765e6a884ed1ec76f65571da873160cedf8dd528227f923924fa2e0a"></a>

## Direct properties — data.metric.value.trend_value / 41fc8dade696 / 3

<a id="canonical-ce4b3d264c465b4290810a960b8ea83edb93a2caf405cbace0e54110019cb20c"></a>

<a id="canonical-d558429f230fefca60de63264bbe1b42f11a69461ea47feb2295d143e59e7364"></a>

## description_spec property — data.metric.value.trend_value / 41fc8dade696 / 4

Type: `"string"`. Computed.

Description of the method used to calculate trend.

<a id="canonical-93fd207092664ec499ae16ea4fc4c70bc6a931ca398d436b8734253e5ddb3a27"></a>

<a id="canonical-dfce5041ce430359747aeb644a893ad2aca5378dfc64004038563a84eb979230"></a>

## previous_value property — data.metric.value.trend_value / 41fc8dade696 / 5

Type: `"string"`. Computed.

Configuration parameter for previous value.

<a id="canonical-b8e6f1b9cd6817d55e0c9f5caf8932305b3b235d305e834637675d9aed7903b7"></a>

<a id="canonical-32a2bf060a3cc7d4ecc2fc4199835f7ad46d09291a0bb648aabb4a2a80865dd5"></a>

## sentiment property — data.metric.value.trend_value / 41fc8dade696 / 6

Type: `"string"`. Computed.

\[Enum: TREND\_SENTIMENT\_NONE|TREND\_SENTIMENT\_POSITIVE|TREND\_SENTIMENT\_NEGATIVE\] Trend
sentiment Indicates trend sentiment is positive Indicates trend sentiment is negative. Possible
values are \`TREND\_SENTIMENT\_NONE\`, \`TREND\_SENTIMENT\_POSITIVE\`,
\`TREND\_SENTIMENT\_NEGATIVE\`. Defaults to \`TREND\_SENTIMENT\_NONE\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TREND_SENTIMENT_NONE",
    "TREND_SENTIMENT_POSITIVE",
    "TREND_SENTIMENT_NEGATIVE"),
}
```

<a id="canonical-c6e83257f43f640a191c068287b68837682559d898ac9bd92b776e8bbfbaabaa"></a>

<a id="canonical-5bc0246db46e023251226ae247840e2880781ec127c687849f35532ee77787be"></a>

## value property — data.metric.value.trend_value / 41fc8dade696 / 7

Type: `"string"`. Computed.

Value. Configuration parameter for value

<a id="canonical-b365e5aa10aca1783b2616d856e7c988843bed1fdd0c4b2e13ca93100cb75bb4"></a>

## Next pages — data.metric.value.trend_value / 41fc8dade696 / 8

- [data.metric.value](data-sources--tmm_session_metrics--reference--group-001.md#canonical-7ad82d2e1b9a6a9e6cadbf28d029c4513dbdd9278670f13126af7f039e73188f)
- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)

<a id="canonical-4401ddc79888b54eaf42c7ca558018746a8ccaecd10f4578cdd4f5e25e000c76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2287091c9d026c1a3e1fe483b13e2872498e1fb664013e0f5be24bccccbed35"></a>

## label_filter — label_filter / 4f924663d6e7 / 2

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)
- [Property reference](data-sources--tmm_session_metrics--reference--group-001.md#canonical-527d0176eb1fcf0d1ab47bc9f224535c861a92810150e2e9a25333b1b6c3fe84)
- label_filter

<a id="canonical-6c8afc2bac643256e9e8f0a2f01d7c5b72bab654c17c40b46055ca8b6ea37877"></a>

Type: `"list"`. Optional.

List of label filter expressions of the form 'label key' QueryOp 'value'. Response will only contain
data that matches all the conditions specified in the label\_filter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

<a id="canonical-ffa6d7ce4a7608e6c7aaacb273a6df1f0a4b49d7b2863b4ba130dd0bf56ccc95"></a>

## Direct properties — label_filter / 4f924663d6e7 / 3

<a id="canonical-28952a15a4dd2b39c6d521f55417bbc5792e9f53d97bd57eeaecf3c8ddf14cae"></a>

<a id="canonical-7ad79184047d7f4095860a464f1f0820b81b42825b82a258974e5de423706dcd"></a>

## label property — label_filter / 4f924663d6e7 / 4

Type: `"string"`. Optional.

\[Enum: METRIC\_LABEL\_NONE|METRIC\_LABEL\_NAMESPACE|METRIC\_LABEL\_VIRTUAL\_SERVER\] Metrics used
to construct the session metrics are tagged with these labels and therefore the metrics can be
sliced and diced based on one or more of these labels. Indicates the field not being set Identifies
the workspace where the service is deployed Identifies the virtual server. Possible values are
\`METRIC\_LABEL\_NONE\`, \`METRIC\_LABEL\_NAMESPACE\`, \`METRIC\_LABEL\_VIRTUAL\_SERVER\`. Defaults
to \`METRIC\_LABEL\_NONE\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("METRIC_LABEL_NONE",
    "METRIC_LABEL_NAMESPACE",
    "METRIC_LABEL_VIRTUAL_SERVER"),
}
```

<a id="canonical-180cec6c5cadbe04944e427acd6f006f3e04b674b0ac740b612e2f694cfbe881"></a>

<a id="canonical-6c6e6d1d0702b013f0a94540f2ca1c667706da6d343bd9c68db96bfa46d15677"></a>

## op property — label_filter / 4f924663d6e7 / 5

Type: `"string"`. Optional.

\[Enum: EQ|NEQ\] The operator to use when filtering metrics based on label values. Possible values
are \`EQ\`, \`NEQ\`. Defaults to \`EQ\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EQ",
    "NEQ"),
}
```

<a id="canonical-34ad5e62dc66796d741ef1e6e4680ff9f4d733ef338db14e654f2ee2abad302b"></a>

<a id="canonical-9b05924fa0305cdb8185609003c3de7d0f2e3949a65d5a7dddaf0166a7edcc2f"></a>

## value property — label_filter / 4f924663d6e7 / 6

Type: `"string"`. Optional.

Value. Value of the label.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-8c61e04d55c120b62220c88c4afd71d1308694193c5a56e65c2433861d69e5af"></a>

## Next pages — label_filter / 4f924663d6e7 / 7

- [Property reference](data-sources--tmm_session_metrics--reference--group-001.md#canonical-527d0176eb1fcf0d1ab47bc9f224535c861a92810150e2e9a25333b1b6c3fe84)
- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)
