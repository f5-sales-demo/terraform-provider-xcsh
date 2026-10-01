---
page_title: "xcsh_device_intelligence_device_summary reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_device_summary reference."
---

# xcsh_device_intelligence_device_summary reference

<a id="canonical-93f159bc76d0b51ddc452c0d614270f715e8fbdad2fe323e8255e24d25ca5d9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4bc5edad990badc0526b6e5c91134a38beba6d8bce27ac5cb1dbf103cdbb685"></a>

## Property reference — Property reference / 136c06d1959b / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-f8e0600c589deef40fd9b13736eefcb6123534aaac445791790673b2e06e3989)
- Property reference

<a id="canonical-24d770a2f796446389b78a19bab4cdf7789c650ae0336dc02938662c69d657fc"></a>

## Direct properties — Property reference / 136c06d1959b / 3

<a id="canonical-dd4d560899cb69cf574e443d9e3872cdeee24129b287f141da3ab15e5a97421a"></a>

<a id="canonical-db2c39b2c287bf5e71f2ea468c3d2f46b7c444e36caae7b4429a4b8250759dbb"></a>

## action_taken property — Property reference / 136c06d1959b / 4

Type: `"string"`. Computed.

Action taken or recommended based on the risk assessment.

<a id="canonical-94196b8cfe0f1a315a450b5bc3ad92b87039e534dcda8f69ef4c29c404d9e444"></a>

<a id="canonical-eed6adc404a53569ca29e6ae3d1bd96c37303af1441e9068734fe6830b16275b"></a>

## channel property — Property reference / 136c06d1959b / 5

Type: `"string"`. Computed.

Channel or platform used by the device (e.g., Web, Mobile).

<a id="canonical-af2074b06d700f939f32edfed92e8f05b84bb92f179469fc39a22ee80d9b9136"></a>

<a id="canonical-ea68da9164052a54e9c0891ff266f0ab8ebfe6e739d19f37168aaf11b7611a8d"></a>

## confidence property — Property reference / 136c06d1959b / 6

Type: `"number"`. Computed.

Confidence level of the risk assessment (0–100).

<a id="canonical-c6c98cb7a131fffec9aef5351679052d069bc357b8c5e05bf822c6b5feb662cf"></a>

<a id="canonical-391dc367f64e7633d3fc0cf51b1b33efc20c36d57215bfe1e51c40531de29790"></a>

## detected_signals property — Property reference / 136c06d1959b / 7

Type: `["list", "string"]`. Computed.

List of risk or integrity signals detected for this device.

<a id="canonical-98d75d320f8510ce0589b873d6e8328b04374d52f77bc7f168b94dcdb2b48f64"></a>

<a id="canonical-1194929b15c06d6dae3d64a35bc3b7fade97d6799ffc5fca780bf37c833d9dc8"></a>

## device_id property — Property reference / 136c06d1959b / 8

Type: `"string"`. Required.

DeviceID. Device identifier.

<a id="canonical-4dc02e246a83ea10f5b1db1c91aae46160f33598173bb4165e44a999fe9bc2d5"></a>

<a id="canonical-148daa6f853d129c0e1cf110ebfb50ce0b3bc0cf0cdc9e8ce804e849fa1f19f5"></a>

## device_risk_score property — Property reference / 136c06d1959b / 9

Type: `"number"`. Computed.

Overall risk score for the device (0–100).

<a id="canonical-42cc7efc207b9d8548599253481d232718a9fab9bf7303f582ac2620014b4434"></a>

<a id="canonical-b418989bdb6ff12b556a36b991d26c5ca0494117d9a3e225a547d78938401e75"></a>

## end_time property — Property reference / 136c06d1959b / 10

Type: `"string"`. Optional.

End Time. End time of the query period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

- [filters](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-547e9333ad37605b8f929e5c900eea6a57fc8daf0478ab407f267b9923dd7b6d): complete subsection reference.

<a id="canonical-afc426cebea6524288fe1824c97553d321a751d7e4eea26fe5ebcefd984424e4"></a>

<a id="canonical-d7cbdd0d44dd28b454bff42ab486e8f5a0a98a1630fb23cedca6fccff24a3828"></a>

## high_risk_txn_count property — Property reference / 136c06d1959b / 11

Type: `"string"`. Computed.

Number of high-risk transactions linked to the device.

<a id="canonical-00825c9da1a27c89b04ffa85b86cd05d3f99ea1da937469729fe471f400d9b45"></a>

<a id="canonical-13238905fa1cc56a949b3e84d9017ba5c7d4b371d8c7ab648b9a3abc26163e7e"></a>

## hosting property — Property reference / 136c06d1959b / 12

Type: `"string"`. Computed.

Type of hosting or network environment detected for the device.

<a id="canonical-49c55d9b1ac9e487a429e02137212e469fd17eef3133785a805ee7bbb715010c"></a>

<a id="canonical-3daecdd0197f8934cdd046145d9f8ec5ec305debd82cfe6ca3f831eb2dcff90f"></a>

## ip_address property — Property reference / 136c06d1959b / 13

Type: `"string"`. Computed.

IP Address. IP address observed for the device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`),
    ""),
}
```

<a id="canonical-c5d4e1eb9b5dfe15d5271dd540d3ef2e0fc466f338bd6a1c6af00c19a06c3b31"></a>

<a id="canonical-56205cbd66c64698e59f710f63aa74005b136a34e22bd3ee486457e1b8ab96b8"></a>

## location property — Property reference / 136c06d1959b / 14

Type: `["list", "number"]`. Computed.

Geographic coordinates associated with the device.

<a id="canonical-9acd290cdd1921e7507284a74dec8ab7a28d4dad003795601423625b71d6e754"></a>

<a id="canonical-819ff05f516f950cfd3da562ab93fc13a53831b97ebe56690f8a951bc1f3acfe"></a>

## namespace property — Property reference / 136c06d1959b / 15

Type: `"string"`. Required.

Namespace. Namespace name.

<a id="canonical-6a2d1f619124a8ba896ac2a18bb113f69d833918a6499055675e8635baf514ac"></a>

<a id="canonical-48c881285b8f70dfd01554b144433f73046f8c4823d46ebd24b6064bb7d8e443"></a>

## new_or_returning_device property — Property reference / 136c06d1959b / 16

Type: `"string"`. Computed.

Indicates whether the device is new or returning.

<a id="canonical-9c85365cf86f7ba97c10f3737ffdc689714544bbe1f859442c87cf88194dbaf1"></a>

<a id="canonical-a9907f02fece4d0883163c05928bd20b2e5c53ca9fcd5cf051a7a482a2c4d090"></a>

## os property — Property reference / 136c06d1959b / 17

Type: `"string"`. Computed.

Operating system and version reported for the device.

<a id="canonical-3bd19876f8b91cb46de96f06f07df72fa94d47cf0f73fa4276ea3d91dc0ed04e"></a>

<a id="canonical-2a8bb8ddbd9c39ec0a2bd3d307dfdef393842322b64dd1ff39ed2d75d512aff0"></a>

## start_time property — Property reference / 136c06d1959b / 18

Type: `"string"`. Optional.

Start Time. Start time of the query period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="canonical-0cab226f3291d2aa1eebaa516c05b1448adf0cd56997247921e1d346bbd7f9c4"></a>

<a id="canonical-545188f8a1118130f520a9e47acca17cf0a24b761b967de249c58f1a491e1e64"></a>

## user_agent property — Property reference / 136c06d1959b / 19

Type: `"string"`. Computed.

User agent string reported by the device/browser.

<a id="canonical-708b062d08dd0333554c4ba14415a88ff8fdc399b505f7399da4482ca7d7f743"></a>

## All schema paths — Property reference / 136c06d1959b / 20

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action_taken` | [action_taken](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-dd4d560899cb69cf574e443d9e3872cdeee24129b287f141da3ab15e5a97421a) |
| `channel` | [channel](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-94196b8cfe0f1a315a450b5bc3ad92b87039e534dcda8f69ef4c29c404d9e444) |
| `confidence` | [confidence](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-af2074b06d700f939f32edfed92e8f05b84bb92f179469fc39a22ee80d9b9136) |
| `detected_signals` | [detected_signals](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-c6c98cb7a131fffec9aef5351679052d069bc357b8c5e05bf822c6b5feb662cf) |
| `device_id` | [device_id](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-98d75d320f8510ce0589b873d6e8328b04374d52f77bc7f168b94dcdb2b48f64) |
| `device_risk_score` | [device_risk_score](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-4dc02e246a83ea10f5b1db1c91aae46160f33598173bb4165e44a999fe9bc2d5) |
| `end_time` | [end_time](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-42cc7efc207b9d8548599253481d232718a9fab9bf7303f582ac2620014b4434) |
| `filters` | [filters](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-8d224c4b03676a98d3f7781b6d4da5122e7c201d3b9bfad5325174accea5e23a) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-44b8ca9234e0eab987a71c6705a98daac4d4acadd659c2ce4d1e76514f0f645d) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-d32ece6b80a383cf1f287fc5ca34d8241f95038b687f9908423057357c5f8d45) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-0bfc6defedccfa8cfed28586f0b0403a4391a9270016def2a17d8371decb9584) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-d9aa0f16ef6599ec1b67bb6b1394d0171477b699f0625ac842011a7e5a22a99b) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-471b8178c9f46b5513aa9585f36fff6578eb95ec3e5b4fc3b12bd119232813ae) |
| `high_risk_txn_count` | [high_risk_txn_count](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-afc426cebea6524288fe1824c97553d321a751d7e4eea26fe5ebcefd984424e4) |
| `hosting` | [hosting](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-00825c9da1a27c89b04ffa85b86cd05d3f99ea1da937469729fe471f400d9b45) |
| `ip_address` | [ip_address](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-49c55d9b1ac9e487a429e02137212e469fd17eef3133785a805ee7bbb715010c) |
| `location` | [location](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-c5d4e1eb9b5dfe15d5271dd540d3ef2e0fc466f338bd6a1c6af00c19a06c3b31) |
| `namespace` | [namespace](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-9acd290cdd1921e7507284a74dec8ab7a28d4dad003795601423625b71d6e754) |
| `new_or_returning_device` | [new_or_returning_device](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-6a2d1f619124a8ba896ac2a18bb113f69d833918a6499055675e8635baf514ac) |
| `os` | [os](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-9c85365cf86f7ba97c10f3737ffdc689714544bbe1f859442c87cf88194dbaf1) |
| `start_time` | [start_time](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-3bd19876f8b91cb46de96f06f07df72fa94d47cf0f73fa4276ea3d91dc0ed04e) |
| `user_agent` | [user_agent](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-0cab226f3291d2aa1eebaa516c05b1448adf0cd56997247921e1d346bbd7f9c4) |

<a id="canonical-a0abef5f8d6d0e4066218d6fe51bc28f6b1718a61a50122fcbaad909ab9ac0dc"></a>

## Next pages — Property reference / 136c06d1959b / 21

- [filters](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-547e9333ad37605b8f929e5c900eea6a57fc8daf0478ab407f267b9923dd7b6d)
- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-f8e0600c589deef40fd9b13736eefcb6123534aaac445791790673b2e06e3989)

<a id="canonical-547e9333ad37605b8f929e5c900eea6a57fc8daf0478ab407f267b9923dd7b6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df8f9ee5befd8c324d6c205a023eb5312ff70362baec247316136f941384d49d"></a>

## filters — filters / 067ae5eaca16 / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-f8e0600c589deef40fd9b13736eefcb6123534aaac445791790673b2e06e3989)
- [Property reference](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-93f159bc76d0b51ddc452c0d614270f715e8fbdad2fe323e8255e24d25ca5d9a)
- filters

<a id="canonical-8d224c4b03676a98d3f7781b6d4da5122e7c201d3b9bfad5325174accea5e23a"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

<a id="canonical-fa140893d9ab181d691daa50edb83402f2c6a094889211d932f6d884ddb2a7c7"></a>

## Direct properties — filters / 067ae5eaca16 / 3

- [global_filters](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-0a0639c6c3691ec73f0ece13e0393873b264b65d6f231cdf179a88708113023c): complete subsection reference.

<a id="canonical-471b8178c9f46b5513aa9585f36fff6578eb95ec3e5b4fc3b12bd119232813ae"></a>

<a id="canonical-ec9fc8adb9c9a6c856eef1104c4f4a8ef7ce52c6e35b60e535b2b59f89a98402"></a>

## region_filter property — filters / 067ae5eaca16 / 4

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

<a id="canonical-7d2b011badefa628653466ec0c0c6cda3dea6176bb282b951cf1771cb4a531b2"></a>

## Next pages — filters / 067ae5eaca16 / 5

- [filters.global_filters](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-0a0639c6c3691ec73f0ece13e0393873b264b65d6f231cdf179a88708113023c)
- [Property reference](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-93f159bc76d0b51ddc452c0d614270f715e8fbdad2fe323e8255e24d25ca5d9a)
- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-f8e0600c589deef40fd9b13736eefcb6123534aaac445791790673b2e06e3989)

<a id="canonical-0a0639c6c3691ec73f0ece13e0393873b264b65d6f231cdf179a88708113023c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25ed0304508d727e773d472ac803d52886f93865a689145e0ed5bb2a3c955d04"></a>

## filters.global_filters — filters.global_filters / b77022fc10d1 / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-f8e0600c589deef40fd9b13736eefcb6123534aaac445791790673b2e06e3989)
- [Property reference](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-93f159bc76d0b51ddc452c0d614270f715e8fbdad2fe323e8255e24d25ca5d9a)
- [filters](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-547e9333ad37605b8f929e5c900eea6a57fc8daf0478ab407f267b9923dd7b6d)
- filters.global_filters

<a id="canonical-44b8ca9234e0eab987a71c6705a98daac4d4acadd659c2ce4d1e76514f0f645d"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

<a id="canonical-bf4a5fbe792ddddd4181145ddd251b1a56eef9d916fb43ea3b6d46acccb99774"></a>

## Direct properties — filters.global_filters / b77022fc10d1 / 3

<a id="canonical-d32ece6b80a383cf1f287fc5ca34d8241f95038b687f9908423057357c5f8d45"></a>

<a id="canonical-85622236c9a36b3852576bd957ef048fd6f710788d3a4416266b9e4b31b7cfe9"></a>

## key property — filters.global_filters / b77022fc10d1 / 4

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

<a id="canonical-0bfc6defedccfa8cfed28586f0b0403a4391a9270016def2a17d8371decb9584"></a>

<a id="canonical-3a2717fd8457cf175c117206a67ac6048570d000f8da33327082111e20b907a2"></a>

## op property — filters.global_filters / b77022fc10d1 / 5

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

<a id="canonical-d9aa0f16ef6599ec1b67bb6b1394d0171477b699f0625ac842011a7e5a22a99b"></a>

<a id="canonical-36f8cf93aee966541ad578e704d5b521ad56151a841bbd451d16d20c3245817a"></a>

## values property — filters.global_filters / b77022fc10d1 / 6

Type: `["list", "string"]`. Optional.

Values. An unordered list of filter strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

<a id="canonical-fc1da1da66ddd8c33cdcc423cb8dfcf8217a0faac8efbb6890e4c67d68002687"></a>

## Next pages — filters.global_filters / b77022fc10d1 / 7

- [filters](data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-547e9333ad37605b8f929e5c900eea6a57fc8daf0478ab407f267b9923dd7b6d)
- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-f8e0600c589deef40fd9b13736eefcb6123534aaac445791790673b2e06e3989)
