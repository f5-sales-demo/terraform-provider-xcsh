---
page_title: "xcsh_device_intelligence_risk_score_distribution reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_risk_score_distribution reference."
---

# xcsh_device_intelligence_risk_score_distribution reference

<a id="canonical-1f6bece68929c1bc8f8d705b9a282c5d577a38a8bc7f44e3d818748cb10109ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db0d89a6a0acd58e13f3ee163e12706c215d7d45b595d2620482559582904488"></a>

## Property reference — Property reference / 535c761e3fbc / 2

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)
- Property reference

<a id="canonical-ed5bf2b497344ae44f0e3b70eed73fa703b06567d56fc2244dd3082f8eccd62d"></a>

## Direct properties — Property reference / 535c761e3fbc / 3

<a id="canonical-8417f450e020a9fab10d828cc385124d68ecaf524f632f9998a30df05783aa3a"></a>

<a id="canonical-f3227b032b650513789b04e5347ae4e787835131640ba1db21a8902e8e537bce"></a>

## end_time property — Property reference / 535c761e3fbc / 4

Type: `"string"`. Optional.

End time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
end\_time will be evaluated to start\_time+10m If start\_time is not specified, then the end\_time
will be evaluated to &lt;current time&gt;.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

- [filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-d0f786fd2f7f7ebf1e908f0af412d5201e5d86eabd2c6e48cdcc2f8677a5ad25): complete subsection reference.

<a id="canonical-d017dd923a45db72bfc1d6c1053b82ffe90ee9b6942afda388a1cd633dd0a3aa"></a>

<a id="canonical-ee954fea65e0e21540931388059d522a7db8f9c968650adc52833f116fa10118"></a>

## namespace property — Property reference / 535c761e3fbc / 5

Type: `"string"`. Required.

Namespace. Namespace name.

- [risk_score_distribution](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-c71e60fa9aadf919727367c93c798e534f50d3cb69f02d0c951ad28437dcde0e): complete subsection reference.

<a id="canonical-5b5c5b98ad424238c09343c0472bbeeed8112c887117c64f8ab80818d4cb0fed"></a>

<a id="canonical-763389ecccf8e8e9fcc8df52b4ada85c42a5d6ed3253bc5ec0276fc2b1c1af3e"></a>

## start_time property — Property reference / 535c761e3fbc / 6

Type: `"string"`. Optional.

Start time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
start\_time will be evaluated to end\_time-10m If end\_time is not specified, then the start\_time
will be evaluated to &lt;current time&gt;-10m.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="canonical-94380adc465bcba395e9d7b702e03fc30201286d7033892978e2be5ac6948b95"></a>

## All schema paths — Property reference / 535c761e3fbc / 7

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `end_time` | [end_time](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-8417f450e020a9fab10d828cc385124d68ecaf524f632f9998a30df05783aa3a) |
| `filters` | [filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-03c8c72adee32b9905997438267cfadf473020129479952fc6c7b7d0d5f8daba) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-6ee93e45e3bb0745be49b52bc98e05cb51aca3d12deaa8f06b69d5d1a7d3978b) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-f0840bcb5e4243f417dff1a4536142281a56ff4de4aed83adfd6af42348cc875) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-cd3d5d0f0a2597b26931392c527be15310dd3ddf65cb68615fd305f99727ea09) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-020e04067a7441815b29d662ae343d54c1108c9dc7508538c70f74adf1d61835) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-bdf110d98236e4bf602113641365a265d25e165894de37ca3d3cc553a20106da) |
| `namespace` | [namespace](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-d017dd923a45db72bfc1d6c1053b82ffe90ee9b6942afda388a1cd633dd0a3aa) |
| `risk_score_distribution` | [risk_score_distribution](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-f4e47c1ef1aedb08e2d87d52b2fae6913364f37f33f7a6e148c4b331f3ae9c0d) |
| `risk_score_distribution.device_count` | [risk_score_distribution.device_count](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-425218f30408024f203a6867a712cd2e7b219f196c2e164d2e5bac3c9db03b20) |
| `risk_score_distribution.risk_rank` | [risk_score_distribution.risk_rank](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-03035bb72663b5feb5bf0456e6b55c759dbc7d0a4dbbb6528fa22638d2407824) |
| `start_time` | [start_time](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-5b5c5b98ad424238c09343c0472bbeeed8112c887117c64f8ab80818d4cb0fed) |

<a id="canonical-54317f5926d5cf950e47aefc69102d0f683bc6b9fbb1aca371bc9e0faf91b95d"></a>

## Next pages — Property reference / 535c761e3fbc / 8

- [filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-d0f786fd2f7f7ebf1e908f0af412d5201e5d86eabd2c6e48cdcc2f8677a5ad25)
- [risk_score_distribution](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-c71e60fa9aadf919727367c93c798e534f50d3cb69f02d0c951ad28437dcde0e)
- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)

<a id="canonical-d0f786fd2f7f7ebf1e908f0af412d5201e5d86eabd2c6e48cdcc2f8677a5ad25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-acd599c88364dcbbf99aae6162a680e790d76526e3612ecced83093fca7e2d75"></a>

## filters — filters / 5448b18b73b6 / 2

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)
- [Property reference](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-1f6bece68929c1bc8f8d705b9a282c5d577a38a8bc7f44e3d818748cb10109ed)
- filters

<a id="canonical-03c8c72adee32b9905997438267cfadf473020129479952fc6c7b7d0d5f8daba"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

<a id="canonical-4e5559712c20d5f7b534b89c5abe3359797053ae6d77523ff0de3ddaf2a4c3a6"></a>

## Direct properties — filters / 5448b18b73b6 / 3

- [global_filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-f724b0d8d43a3e6fb3ba748017ab69c6fb539648666462e315c6c503902547dd): complete subsection reference.

<a id="canonical-bdf110d98236e4bf602113641365a265d25e165894de37ca3d3cc553a20106da"></a>

<a id="canonical-9dfd0323bda169eac2831ec7d15f0e925491448e6759bb4e653e21282e878afb"></a>

## region_filter property — filters / 5448b18b73b6 / 4

Type: `"string"`. Optional.

\[Enum: US|EU|ASIA|CA\] Defines a selection for Bot Defense region - US: US United States of America
&#8203;- EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are \`US\`, \`EU\`,
\`ASIA\`, \`CA\`. Defaults to \`US\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA",
    "CA"),
}
```

<a id="canonical-bdc1cd91f057a2bacc1950e390349761d6621c567c60f05705b704205d78ab56"></a>

## Next pages — filters / 5448b18b73b6 / 5

- [filters.global_filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-f724b0d8d43a3e6fb3ba748017ab69c6fb539648666462e315c6c503902547dd)
- [Property reference](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-1f6bece68929c1bc8f8d705b9a282c5d577a38a8bc7f44e3d818748cb10109ed)
- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)

<a id="canonical-f724b0d8d43a3e6fb3ba748017ab69c6fb539648666462e315c6c503902547dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7261c6befdaaea4e9e0d4a255b2b70bea68f5b3dd469cbd85828e9c1bc23a4a3"></a>

## filters.global_filters — filters.global_filters / 3b80f397ac5f / 2

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)
- [Property reference](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-1f6bece68929c1bc8f8d705b9a282c5d577a38a8bc7f44e3d818748cb10109ed)
- [filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-d0f786fd2f7f7ebf1e908f0af412d5201e5d86eabd2c6e48cdcc2f8677a5ad25)
- filters.global_filters

<a id="canonical-6ee93e45e3bb0745be49b52bc98e05cb51aca3d12deaa8f06b69d5d1a7d3978b"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

<a id="canonical-6d1871c903310c74d0e2973c7ce0a22722c4fc8763e8c30053f2e0981d7bafc4"></a>

## Direct properties — filters.global_filters / 3b80f397ac5f / 3

<a id="canonical-f0840bcb5e4243f417dff1a4536142281a56ff4de4aed83adfd6af42348cc875"></a>

<a id="canonical-6f4e4aa83b16e634d9e43fb1b2381f39a1d9b127979363541636384590de25f8"></a>

## key property — filters.global_filters / 3b80f397ac5f / 4

Type: `"string"`. Optional.

\[Enum:
TIMESTAMP|USERNAME|CLIENT\_TOKEN|IP\_ADDRESS|ASN|AS\_ORGANIZATION|COUNTRY|METHOD|HOST|PATH|URL|REFERER|TRAFFIC\_CHANNEL|IS\_ATTACK|BOT\_REASON|TRAFFIC\_TYPE|THREAT\_TYPE|SDK\_VERSION|ACTION\_TAKEN|COOKIE\_AGE|BOT\_COOKIE|USER\_AGENT|USER\_AGENT\_OS\_FAMILY|USER\_AGENT\_FAMILY|BROWSER\_FINGERPRINT|USER\_FINGERPRINT|HEADER\_FINGERPRINT|DEVICE\_ID|FLOW|AGENT|APPLICATION\_NAME|PROTECTED\_APPLICATION|RESPONSE\_CODE|SERVER\_RESPONSE\_CODE|TRANSACTION\_RESULT|MOBILE\_TRANSACTION\_INSIGHT|WEB\_TRANSACTION\_INSIGHT|TRIGGERED\_RULE|FLOW\_CATEGORY|FLOW\_LABEL|ENDPOINT\_NAME|ENDPOINT\_LABEL|BOT\_ENDPOINT\_POLICY|KNOWN\_BOT\_NAME|KNOWN\_BOT\_CATEGORY|KNOWN\_BOT\_PROVIDER|KNOWN\_BOT\_CATEGORY\_TYPE|KNOWN\_BOT\_MITIGATION|ABSOLUTE|PERCENTAGE|TREND|ENDPOINT\_POLICY\]
Key for query filter - TIMESTAMP: Timestamp Filter Key Use Timestamp as key to query. Possible
values are \`TIMESTAMP\`, \`USERNAME\`, \`CLIENT\_TOKEN\`, \`IP\_ADDRESS\`, \`ASN\`,
\`AS\_ORGANIZATION\`, \`COUNTRY\`, \`METHOD\`, \`HOST\`, \`PATH\`, \`URL\`, \`REFERER\`,
\`TRAFFIC\_CHANNEL\`, \`IS\_ATTACK\`, \`BOT\_REASON\`, \`TRAFFIC\_TYPE\`, \`THREAT\_TYPE\`,
\`SDK\_VERSION\`, \`ACTION\_TAKEN\`, \`COOKIE\_AGE\`, \`BOT\_COOKIE\`, \`USER\_AGENT\`,
\`USER\_AGENT\_OS\_FAMILY\`, \`USER\_AGENT\_FAMILY\`, \`BROWSER\_FINGERPRINT\`,
\`USER\_FINGERPRINT\`, \`HEADER\_FINGERPRINT\`, \`DEVICE\_ID\`, \`FLOW\`, \`AGENT\`,
\`APPLICATION\_NAME\`, \`PROTECTED\_APPLICATION\`, \`RESPONSE\_CODE\`, \`SERVER\_RESPONSE\_CODE\`,
\`TRANSACTION\_RESULT\`, \`MOBILE\_TRANSACTION\_INSIGHT\`, \`WEB\_TRANSACTION\_INSIGHT\`,
\`TRIGGERED\_RULE\`, \`FLOW\_CATEGORY\`, \`FLOW\_LABEL\`, \`ENDPOINT\_NAME\`, \`ENDPOINT\_LABEL\`,
\`BOT\_ENDPOINT\_POLICY\`, \`KNOWN\_BOT\_NAME\`, \`KNOWN\_BOT\_CATEGORY\`, \`KNOWN\_BOT\_PROVIDER\`,
\`KNOWN\_BOT\_CATEGORY\_TYPE\`, \`KNOWN\_BOT\_MITIGATION\`, \`ABSOLUTE\`, \`PERCENTAGE\`, \`TREND\`,
\`ENDPOINT\_POLICY\`. Defaults to \`TIMESTAMP\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TIMESTAMP",
    "USERNAME",
    "CLIENT_TOKEN",
    "IP_ADDRESS",
    "ASN",
    "AS_ORGANIZATION",
    "COUNTRY",
    "METHOD",
    "HOST",
    "PATH",
    "URL",
    "REFERER",
    "TRAFFIC_CHANNEL",
    "IS_ATTACK",
    "BOT_REASON",
    "TRAFFIC_TYPE",
    "THREAT_TYPE",
    "SDK_VERSION",
    "ACTION_TAKEN",
    "COOKIE_AGE",
    "BOT_COOKIE",
    "USER_AGENT",
    "USER_AGENT_OS_FAMILY",
    "USER_AGENT_FAMILY",
    "BROWSER_FINGERPRINT",
    "USER_FINGERPRINT",
    "HEADER_FINGERPRINT",
    "DEVICE_ID",
    "FLOW",
    "AGENT",
    "APPLICATION_NAME",
    "PROTECTED_APPLICATION",
    "RESPONSE_CODE",
    "SERVER_RESPONSE_CODE",
    "TRANSACTION_RESULT",
    "MOBILE_TRANSACTION_INSIGHT",
    "WEB_TRANSACTION_INSIGHT",
    "TRIGGERED_RULE",
    "FLOW_CATEGORY",
    "FLOW_LABEL",
    "ENDPOINT_NAME",
    "ENDPOINT_LABEL",
    "BOT_ENDPOINT_POLICY",
    "KNOWN_BOT_NAME",
    "KNOWN_BOT_CATEGORY",
    "KNOWN_BOT_PROVIDER",
    "KNOWN_BOT_CATEGORY_TYPE",
    "KNOWN_BOT_MITIGATION",
    "ABSOLUTE",
    "PERCENTAGE",
    "TREND",
    "ENDPOINT_POLICY"),
}
```

<a id="canonical-cd3d5d0f0a2597b26931392c527be15310dd3ddf65cb68615fd305f99727ea09"></a>

<a id="canonical-993bc93b8ed5c13ceb40153c61405d80ade85b4f9aa319596beda29384db0951"></a>

## op property — filters.global_filters / 3b80f397ac5f / 5

Type: `"string"`. Optional.

\[Enum:
IN|NOT\_IN|MATCHES\_REGEX|DOES\_NOT\_MATCH\_REGEX|INCLUDES|DOES\_NOT\_INCLUDE|STARTS\_WITH|ENDS\_WITH\]
Operator for query filter - IN: Filter Operator Specifies that query result includes filter values -
NOT\_IN: Filter Operator Specifies that query result excludes filter values - MATCHES\_REGEX: Filter
Operator Specifies that query result matches filter regex - DOES\_NOT\_MATCH\_REGEX: Filter..
Possible values are \`IN\`, \`NOT\_IN\`, \`MATCHES\_REGEX\`, \`DOES\_NOT\_MATCH\_REGEX\`,
\`INCLUDES\`, \`DOES\_NOT\_INCLUDE\`, \`STARTS\_WITH\`, \`ENDS\_WITH\`. Defaults to \`IN\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("IN",
    "NOT_IN",
    "MATCHES_REGEX",
    "DOES_NOT_MATCH_REGEX",
    "INCLUDES",
    "DOES_NOT_INCLUDE",
    "STARTS_WITH",
    "ENDS_WITH"),
}
```

<a id="canonical-020e04067a7441815b29d662ae343d54c1108c9dc7508538c70f74adf1d61835"></a>

<a id="canonical-42224297f1fa312c631409575de8eeb74efc2b77d689960582ed277d1f7a0466"></a>

## values property — filters.global_filters / 3b80f397ac5f / 6

Type: `["list", "string"]`. Optional.

Values. An unordered list of filter strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

<a id="canonical-d8334463d0cdd37b658fc722fc44380712042257b8e614283904865036f6a119"></a>

## Next pages — filters.global_filters / 3b80f397ac5f / 7

- [filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-d0f786fd2f7f7ebf1e908f0af412d5201e5d86eabd2c6e48cdcc2f8677a5ad25)
- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)

<a id="canonical-c71e60fa9aadf919727367c93c798e534f50d3cb69f02d0c951ad28437dcde0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce6b5bae8352358de55490b8d4b4e2aceff09b15b07c9fe5d63716a66b99fd66"></a>

## risk_score_distribution — risk_score_distribution / efff6d396ecb / 2

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)
- [Property reference](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-1f6bece68929c1bc8f8d705b9a282c5d577a38a8bc7f44e3d818748cb10109ed)
- risk_score_distribution

<a id="canonical-f4e47c1ef1aedb08e2d87d52b2fae6913364f37f33f7a6e148c4b331f3ae9c0d"></a>

Type: `"list"`. Computed.

Distribution of devices across risk ranks.

<a id="canonical-e026d09daf98f1da89f39296c3eba569c019ac56d18ac8283d60c766c4441969"></a>

## Direct properties — risk_score_distribution / efff6d396ecb / 3

<a id="canonical-425218f30408024f203a6867a712cd2e7b219f196c2e164d2e5bac3c9db03b20"></a>

<a id="canonical-fb90d29b3e6e7c318b59714b9a98aee5ae06ce4782058208df1a2f83075e9256"></a>

## device_count property — risk_score_distribution / efff6d396ecb / 4

Type: `"string"`. Computed.

Device Count. Number of devices in this risk rank.

<a id="canonical-03035bb72663b5feb5bf0456e6b55c759dbc7d0a4dbbb6528fa22638d2407824"></a>

<a id="canonical-e3207bffdc6fc42d34f92a030a121db761f781492eeda2f12152b56e666d6d70"></a>

## risk_rank property — risk_score_distribution / efff6d396ecb / 5

Type: `"string"`. Computed.

Risk rank label (e.g., low, medium, high).

<a id="canonical-18cba55318907d504cb20d34595da9e9f4a3a947569f0d583a1ebc974a2bfeb6"></a>

## Next pages — risk_score_distribution / efff6d396ecb / 6

- [Property reference](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-1f6bece68929c1bc8f8d705b9a282c5d577a38a8bc7f44e3d818748cb10109ed)
- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-28a66890f008fdac1b26cd785850ece6401f38637446a0123c4e80cf02ad050e)
