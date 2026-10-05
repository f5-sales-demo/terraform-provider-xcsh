---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0111230211211110-1302302232002121-0103221130121213-2312112112021333-0113201133233023-2111201202103313-1010010031222302-3013323020330310"></a>

## Direct properties — sensitive_data_types_in_response / 331122021311 / 3

- [api_endpoint](resources--http_loadbalancer--reference--group-027.md#canonical-3103211332300101-2111000032002202-2123203011222111-0232030123311232-0101001131020012-0312321031323223-1120100023333102-3201121233122010): complete subsection reference.

- [body](resources--http_loadbalancer--reference--group-027.md#canonical-2323121012223210-0232113003332230-0221101000212202-0011210103332001-1122131032233232-1201112100203231-0302310310112000-1231223202202011): complete subsection reference.

- [mask](resources--http_loadbalancer--reference--group-027.md#canonical-1321030133212020-1032103222103223-3011223320311030-1311111102300013-1323231300132100-1102302210313222-3233222023020003-2110203301303303): complete subsection reference.

- [report](resources--http_loadbalancer--reference--group-027.md#canonical-1031321211211013-2032033023020123-3013031030132211-2120302132030123-3013231032233200-0201320012321231-2300332000133200-3132311003300311): complete subsection reference.

<a id="canonical-1322322203123101-3301212100112323-3131003223201032-3102210232230100-1103022322302200-1211030010032102-0002233022230220-0030013232021302"></a>

## Next pages — sensitive_data_types_in_response / 331122021311 / 4

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint](resources--http_loadbalancer--reference--group-027.md#canonical-3103211332300101-2111000032002202-2123203011222111-0232030123311232-0101001131020012-0312321031323223-1120100023333102-3201121233122010)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.body](resources--http_loadbalancer--reference--group-027.md#canonical-2323121012223210-0232113003332230-0221101000212202-0011210103332001-1122131032233232-1201112100203231-0302310310112000-1231223202202011)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask](resources--http_loadbalancer--reference--group-027.md#canonical-1321030133212020-1032103222103223-3011223320311030-1311111102300013-1323231300132100-1102302210313222-3233222023020003-2110203301303303)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.report](resources--http_loadbalancer--reference--group-027.md#canonical-1031321211211013-2032033023020123-3013031030132211-2120302132030123-3013231032233200-0201320012321231-2300332000133200-3132311003300311)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-3010130000211103-2132323303233130-3123203102231323-2302320310112320-0023100231121123-3223331112120111-3300221200232213-1133321320223332)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3103211332300101-2111000032002202-2123203011222111-0232030123311232-0101001131020012-0312321031323223-1120100023333102-3201121233122010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101003332130122-2321121312013221-3012130021120231-1002003332000332-1033023313233203-1210132213310023-0303101021213331-1312001123303122"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint — api_endpoint / 123003211332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-3010130000211103-2132323303233130-3123203102231323-2302320310112320-0023100231121123-3223331112120111-3300221200232213-1133321320223332)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-2023120330103011-1030003123103102-1102223030033323-3132213302112330-2131203003333230-1012102100122000-3211320013000233-2230021332122030)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint

<a id="canonical-3223122221101301-2310320132012120-1030001310232030-2200322203312000-2313033132020222-2302310123301330-2013123033033132-1213203220333012"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102133102330003-0332020120101323-0223011231030313-2322211201202132-1101331031031222-2232321112110201-0301233232331000-1123000301200022"></a>

## Direct properties — api_endpoint / 123003211332 / 3

<a id="canonical-2311102211012221-2012211122023200-0123122333103133-1011032131122213-3022101010310110-0232312312030120-0312112031102221-0333001033322100"></a>

<a id="canonical-1331232322112023-0103322231031202-0002200213302231-2130020113301120-2133331100022033-2000333203302032-2201321230113011-2103231332203222"></a>

## methods property — api_endpoint / 123003211332 / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2232000131001103-2030032120012030-0001331133022132-0310112332120121-2100113210011232-0113203201030000-2020123003003032-0222100302132210"></a>

<a id="canonical-3001132032031122-0332300210222330-0112122221030102-0213311310001111-1132100331101313-3000321010010101-2330032123113303-0231131011113210"></a>

## path property — api_endpoint / 123003211332 / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-2231122103013101-1102330223122121-2100122222310123-2231012201331233-1131311333133103-0111013022132132-0003313103312113-2230112231112103"></a>

## Next pages — api_endpoint / 123003211332 / 6

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-2023120330103011-1030003123103102-1102223030033323-3132213302112330-2131203003333230-1012102100122000-3211320013000233-2230021332122030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2323121012223210-0232113003332230-0221101000212202-0011210103332001-1122131032233232-1201112100203231-0302310310112000-1231223202202011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033303022300012-1213203122320132-0023223033102003-0221023110300021-0021020133000123-2013320320202213-3222323100220201-2200312313232201"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.body — body / 303312033111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-3010130000211103-2132323303233130-3123203102231323-2302320310112320-0023100231121123-3223331112120111-3300221200232213-1133321320223332)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-2023120330103011-1030003123103102-1102223030033323-3132213302112330-2131203003333230-1012102100122000-3211320013000233-2230021332122030)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.body

<a id="canonical-2231303112010333-2033000013021233-0303012031300230-2111100032030213-3000320220030201-0110102331310123-0013033201332300-0132100313331121"></a>

Type: `"object"`. single nested block, Optional.

Body Section Masking OPTIONS. OPTIONS for HTTP Body Masking.

Upstream description:

OPTIONS for HTTP Body Masking.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("fields")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
body {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113233023321331-1133122201032010-3332002231302231-2321030211032202-1323313011201101-2313213101121002-0101032213200203-1332031321322033"></a>

## Direct properties — body / 303312033111 / 3

<a id="canonical-0001230300022110-1332303022332203-3332101023300103-0220123201011310-0313203300003022-2202211033010101-2013320313133310-3203320303001221"></a>

<a id="canonical-1023021223312221-2222301301000312-2303123212203233-1310121222322132-0302221202300301-2031112312130001-2211120331330102-3030031223113001"></a>

## fields property — body / 303312033111 / 4

Type: `["list", "string"]`. Optional.

List of JSON Path field values. Use square brackets with an underscore \[\_\] to indicate array
elements (e.g., person.emails\[\_\]). To reference JSON keys that contain spaces, enclose the entire
path in double quotes.

Upstream description:

List of JSON Path field values. Use square brackets with an underscore \[\_\] to indicate array
elements (e.g., person.emails\[\_\]). To reference JSON keys that contain spaces, enclose the entire
path in double quotes. For example: "person.first name".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0013033132210101-3330032113222000-3201002302203211-2300030120021231-0012131220231303-1023132203132202-1331323123331312-0300200032000302"></a>

## Next pages — body / 303312033111 / 5

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-2023120330103011-1030003123103102-1102223030033323-3132213302112330-2131203003333230-1012102100122000-3211320013000233-2230021332122030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1321030133212020-1032103222103223-3011223320311030-1311111102300013-1323231300132100-1102302210313222-3233222023020003-2110203301303303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021123020320103-0313202213311221-3201121333123221-1001200203313222-2230103333103020-1220102322213330-1011210013323120-1132131210032102"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask — mask / 000230331322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-3010130000211103-2132323303233130-3123203102231323-2302320310112320-0023100231121123-3223331112120111-3300221200232213-1133321320223332)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-2023120330103011-1030003123103102-1102223030033323-3132213302112330-2131203003333230-1012102100122000-3211320013000233-2230021332122030)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask

<a id="canonical-3312131321313331-3221220122310311-1100021312132110-1320022333211013-3122013103101020-1213233211010121-1211330311003003-2331323303332220"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
mask = {}
```

<a id="canonical-1133110330221101-1300013131103132-3211100333332332-3211311020000201-3221022002212300-2313203202201002-3020212333302032-0022221302003121"></a>

## Direct properties — mask / 000230331322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202303311113113-3301330021030032-1101331301222012-3200110330201222-0002023231332303-3202213332112210-1220301220311200-1033223213222303"></a>

## Next pages — mask / 000230331322 / 4

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-2023120330103011-1030003123103102-1102223030033323-3132213302112330-2131203003333230-1012102100122000-3211320013000233-2230021332122030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1031321211211013-2032033023020123-3013031030132211-2120302132030123-3013231032233200-0201320012321231-2300332000133200-3132311003300311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001230123032112-3222222122333011-0033002301112130-3220322010211012-1210233323211133-3020113323300003-1320133032332312-1201101233300233"></a>

## sensitive_data_disclosure_rules.sensitive_data_types_in_response.report — report / 322100201013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-026.md#canonical-3010130000211103-2132323303233130-3123203102231323-2302320310112320-0023100231121123-3223331112120111-3300221200232213-1133321320223332)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-2023120330103011-1030003123103102-1102223030033323-3132213302112330-2131203003333230-1012102100122000-3211320013000233-2230021332122030)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.report

<a id="canonical-3120320112302000-2200203111013303-0013323100032203-3013021300321020-3131132301231130-2210202133220313-0001330132111130-3010030233233230"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
report = {}
```

<a id="canonical-0323123213122022-2112103121000110-1021323222032323-0121233332123322-3313113030330012-2232302000022200-1023011131321231-0033312131010211"></a>

## Direct properties — report / 322100201013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103333100233113-2313001210333122-1203220011001230-0022202302330321-0311332130202011-1221001310132010-3001212132232222-2303232310331232"></a>

## Next pages — report / 322100201013 / 4

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](resources--http_loadbalancer--reference--group-026.md#canonical-2023120330103011-1030003123103102-1102223030033323-3132213302112330-2131203003333230-1012102100122000-3211320013000233-2230021332122030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3300321023010221-0210333331122211-1013103230213030-3001321133032323-0100212010332121-1322030300320113-3031022123002130-1203331012330020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210110233333021-2031331121012111-3021330203332031-2203202003202132-0003230300312121-3220203033232020-3021203232220011-1321032231220310"></a>

## sensitive_data_policy — sensitive_data_policy / 223120301021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- sensitive_data_policy

<a id="canonical-2331322013011313-3111202121321023-1113002312200000-0132100333303032-2311313300200313-3003330023023213-0133322032323223-3212012023320222"></a>

Type: `"object"`. single nested block, Optional.

Policy configuration for this feature.

Upstream description:

Settings for data type policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
sensitive_data_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023111202111032-2312223233321333-2120120323103212-2202222212322220-2320101221120122-0230111110132303-2113112032120131-3100310203300222"></a>

## Direct properties — sensitive_data_policy / 223120301021 / 3

- [sensitive_data_policy_ref](resources--http_loadbalancer--reference--group-027.md#canonical-1223012002323103-0010121123003320-0012323003132010-1103030021012132-1333100333023200-0132023121322200-2313313332312201-1010131221020113): complete subsection reference.

<a id="canonical-1001230123101130-2333331101311121-1103330132003333-1110301300213120-2120220010213233-2332111322313200-2002311312033121-1030332310112032"></a>

## Next pages — sensitive_data_policy / 223120301021 / 4

- [sensitive_data_policy.sensitive_data_policy_ref](resources--http_loadbalancer--reference--group-027.md#canonical-1223012002323103-0010121123003320-0012323003132010-1103030021012132-1333100333023200-0132023121322200-2313313332312201-1010131221020113)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1223012002323103-0010121123003320-0012323003132010-1103030021012132-1333100333023200-0132023121322200-2313313332312201-1010131221020113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111321023222123-2102310300333002-2112200203030020-1300300323020011-2103203110331110-3231112120010101-2233011202132332-1030200002010203"></a>

## sensitive_data_policy.sensitive_data_policy_ref — sensitive_data_policy_ref / 112232322102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [sensitive_data_policy](resources--http_loadbalancer--reference--group-027.md#canonical-3300321023010221-0210333331122211-1013103230213030-3001321133032323-0100212010332121-1322030300320113-3031022123002130-1203331012330020)
- sensitive_data_policy.sensitive_data_policy_ref

<a id="canonical-0231002003101311-3203010303130301-2212311331023311-1101000010100332-1132110331222232-0210012120132312-3303123210201012-2132310331233201"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
sensitive_data_policy_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100320330323312-2222012211122003-0010323003311323-0203222330111320-2231112032211221-3213111032231110-2022311130031232-0313001232201303"></a>

## Direct properties — sensitive_data_policy_ref / 112232322102 / 3

<a id="canonical-3233321110121131-3222321023323213-3210012132203033-1310211102232313-0103330210303102-2230223001311102-3023303310323312-0232302123312130"></a>

<a id="canonical-2201200030202223-0102331122011013-3203220203210301-3100033121330310-0111020203211210-0221220331113210-2322312333213230-3001132133321123"></a>

## name property — sensitive_data_policy_ref / 112232322102 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2312203213101313-0221332030100223-1211001130030212-2110013212131001-1332223030023011-0332121001302303-0210120103130223-3000233101302321"></a>

<a id="canonical-0233102003220230-0122313333101023-1102200300223312-0330303000230330-1103010200202312-2311302133200321-3111302020033021-2113221303210233"></a>

## namespace property — sensitive_data_policy_ref / 112232322102 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0000322012202312-1030210332120120-2120003231102100-0202112130303321-3211332212012211-1123010233020212-0213333112333302-2111201311032000"></a>

<a id="canonical-3300321002111213-1323330132201212-0022113033002201-3223131231110113-2100222222113002-0122322312121031-1131130021313202-0030100010320103"></a>

## tenant property — sensitive_data_policy_ref / 112232322102 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2013002232331302-3033011020000002-0132030010330211-2233012310210103-3200122211030231-2321003033310211-3100133113033312-1233121102112113"></a>

## Next pages — sensitive_data_policy_ref / 112232322102 / 7

- [sensitive_data_policy](resources--http_loadbalancer--reference--group-027.md#canonical-3300321023010221-0210333331122211-1013103230213030-3001321133032323-0100212010332121-1322030300320113-3031022123002130-1203331012330020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3320321312231220-2203023313122012-3333323031313130-0133112032311321-1013010313122112-3211000010020331-1110330001120132-0323002112221111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122112032002200-0122312212322332-1101122023223101-0222312313131002-1300203321022101-3113000200200112-0103102013102303-0030212213110312"></a>

## service_policies_from_namespace — service_policies_from_namespace / 132011210023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- service_policies_from_namespace

<a id="canonical-3123232231202213-3113013113132233-2130200300221332-2101202232201032-3213020100120021-3201232030111013-1022231010320313-1100333020311221"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
service_policies_from_namespace = {}
```

<a id="canonical-2221031313221101-0301123133222132-2312302200000323-0331120300320310-0222211323202200-2031121312222232-1300301110021122-3212211032022031"></a>

## Direct properties — service_policies_from_namespace / 132011210023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201002130222322-0320003021122001-0023023211030002-1313223002130003-3132112311102223-3110201203201301-0322331001030123-1323313033011313"></a>

## Next pages — service_policies_from_namespace / 132011210023 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321120231033232-1313230330121231-2322000122310100-1331233300121203-0131121233213301-0120220211101303-3303300021032030-3320012123230230"></a>

## single_lb_app — single_lb_app / 100310001120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- single_lb_app

<a id="canonical-1201323130312033-1332212010203222-1112131130000000-3313131233123132-0112201202300002-2122020103032101-3012032203220001-1123233022300133"></a>

Type: `"object"`. single nested block, Optional.

Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_discovery",
    "enable_discovery"),
  validators.ConflictingObjectAttributes("disable_malicious_user_detection",
    "enable_malicious_user_detection")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_choice": "[\"disable_discovery\",\"enable_discovery\"]",
  "x-ves-oneof-field-malicious_user_detection_choice": "[\"disable_malicious_user_detection\",\"enable_malicious_user_detection\"]"
}
```

Terraform syntax:

```terraform
single_lb_app {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122230333313112-3303203220213010-2030022130201322-1012230133213112-3002333202110332-0113010201000102-3323103030102300-3312123033003111"></a>

## Direct properties — single_lb_app / 100310001120 / 3

- [disable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-1002330301111021-3230331100121210-1223300130323020-0213302300120331-0112313302210030-1131331102223311-3223320113133133-2312310122000303): complete subsection reference.

- [disable_malicious_user_detection](resources--http_loadbalancer--reference--group-027.md#canonical-2222103321313112-3112230212133331-1011302132103213-0301012223121231-2232203302310132-1203201001030033-3002121311332322-1220011031322033): complete subsection reference.

- [enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233): complete subsection reference.

- [enable_malicious_user_detection](resources--http_loadbalancer--reference--group-027.md#canonical-0323203132202222-0231330232030303-0301303112213002-0030002232012332-3202313221231110-1231102102000012-1012221103330222-0003130020021302): complete subsection reference.

<a id="canonical-0212201313030311-2031323313203313-1220000102313202-1131230231110230-0323123110100330-1301023100230301-0011322133330030-3221332232330223"></a>

## Next pages — single_lb_app / 100310001120 / 4

- [single_lb_app.disable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-1002330301111021-3230331100121210-1223300130323020-0213302300120331-0112313302210030-1131331102223311-3223320113133133-2312310122000303)
- [single_lb_app.disable_malicious_user_detection](resources--http_loadbalancer--reference--group-027.md#canonical-2222103321313112-3112230212133331-1011302132103213-0301012223121231-2232203302310132-1203201001030033-3002121311332322-1220011031322033)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_malicious_user_detection](resources--http_loadbalancer--reference--group-027.md#canonical-0323203132202222-0231330232030303-0301303112213002-0030002232012332-3202313221231110-1231102102000012-1012221103330222-0003130020021302)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1002330301111021-3230331100121210-1223300130323020-0213302300120331-0112313302210030-1131331102223311-3223320113133133-2312310122000303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210033102310113-0311202132021213-1301303130020110-1012311332200011-3113203002212001-2313202232032203-1012332030200111-3122130322022300"></a>

## single_lb_app.disable_discovery — disable_discovery / 322233220333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- single_lb_app.disable_discovery

<a id="canonical-0012013333023201-1012231212131031-2130303320003023-0023022232133102-3023010122000101-3222103221332310-1032313033021012-3132313231212211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable discovery.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_discovery = {}
```

<a id="canonical-2113101321133001-2211132133132000-0330133311322101-2112030033031302-2302211112233220-3131332103231332-0200310330102302-0322220322330302"></a>

## Direct properties — disable_discovery / 322233220333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303323101132122-2312132030220132-0213011131322110-2220033231332321-1001213031333223-3000123332110122-0201300001213012-1300313222210020"></a>

## Next pages — disable_discovery / 322233220333 / 4

- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2222103321313112-3112230212133331-1011302132103213-0301012223121231-2232203302310132-1203201001030033-3002121311332322-1220011031322033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120312020301202-2013333021302321-1000003332332121-2321001113010033-0212001201312220-0102000112001201-0200213322113031-0022313030311013"></a>

## single_lb_app.disable_malicious_user_detection — disable_malicious_user_detection / 133301230313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- single_lb_app.disable_malicious_user_detection

<a id="canonical-2120323232331223-0123311131200101-3101301122112010-0113132102003111-0120123233231020-3220332321303020-2333002101130011-0110001033112231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable malicious user detection.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_malicious_user_detection = {}
```

<a id="canonical-3100301131121131-2101113130122020-0010000222002200-3232111103032002-2301103221031301-3000130332210230-2321230311203013-0220130031230101"></a>

## Direct properties — disable_malicious_user_detection / 133301230313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200002212132210-2220321021312301-2121102002301232-3101022022302233-3133121012132012-2211333131002032-1123121101101111-2302032330301211"></a>

## Next pages — disable_malicious_user_detection / 133301230313 / 4

- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130211301333331-2322232330013223-2302001013130030-3131232333122111-2001300131233011-1312121030330223-2123033311210032-2020100313103220"></a>

## single_lb_app.enable_discovery — enable_discovery / 332132010132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- single_lb_app.enable_discovery

<a id="canonical-2311201022122322-2033323010301013-2112013312212220-3131103332331002-1201003320303023-3322330320002111-2033223113321133-2203113003300333"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings used for API discovery.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_api_auth_discovery",
    "default_api_auth_discovery"),
  validators.ConflictingObjectAttributes("disable_learn_from_redirect_traffic",
    "enable_learn_from_redirect_traffic")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

Terraform syntax:

```terraform
enable_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131233200130200-3333223133333103-2111201113222101-0221331020112111-0133021213021330-3220100110132032-0300031222122320-3200220231333032"></a>

## Direct properties — enable_discovery / 332132010132 / 3

- [api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012): complete subsection reference.

- [api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-027.md#canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220): complete subsection reference.

- [custom_api_auth_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-0033020300201021-2112213210103002-3302223002101203-1012232012320110-3312220020100332-0310003001132021-2230020302333111-0112132130323101): complete subsection reference.

- [default_api_auth_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-1312311021131010-1232123020112022-3111131320310202-0101232302110333-2132211003300303-0130323303100031-0002131301100313-0230303031012130): complete subsection reference.

- [disable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-027.md#canonical-1131121200123133-1212031302303313-1013301330332331-1101002022012232-2321013232202222-2303001302131230-1130011122113321-2201223210020302): complete subsection reference.

- [discovered_api_settings](resources--http_loadbalancer--reference--group-027.md#canonical-0221012011230012-1010330112321300-3320100021211210-0310132001223101-1033302320113013-1203012103122302-1000022330113011-3012213201002311): complete subsection reference.

- [enable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-027.md#canonical-1103000230303211-3013102200220030-3232211010330301-3323201112330001-3233012203213323-2002131102133010-3312223131112321-2202313010211300): complete subsection reference.

<a id="canonical-0112011332121031-2002200320201232-0122210331201211-2111121000022223-2103031032210313-2002312201330003-1201210003321022-3012101312031203"></a>

## Next pages — enable_discovery / 332132010132 / 4

- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-027.md#canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220)
- [single_lb_app.enable_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-0033020300201021-2112213210103002-3302223002101203-1012232012320110-3312220020100332-0310003001132021-2230020302333111-0112132130323101)
- [single_lb_app.enable_discovery.default_api_auth_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-1312311021131010-1232123020112022-3111131320310202-0101232302110333-2132211003300303-0130323303100031-0002131301100313-0230303031012130)
- [single_lb_app.enable_discovery.disable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-027.md#canonical-1131121200123133-1212031302303313-1013301330332331-1101002022012232-2321013232202222-2303001302131230-1130011122113321-2201223210020302)
- [single_lb_app.enable_discovery.discovered_api_settings](resources--http_loadbalancer--reference--group-027.md#canonical-0221012011230012-1010330112321300-3320100021211210-0310132001223101-1033302320113013-1203012103122302-1000022330113011-3012213201002311)
- [single_lb_app.enable_discovery.enable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-027.md#canonical-1103000230303211-3013102200220030-3232211010330301-3323201112330001-3233012203213323-2002131102133010-3312223131112321-2202313010211300)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123123332323000-2212301213011321-2002302201311220-3323113121020230-1303332212003333-3311120112321203-3031000303331331-3223310222020132"></a>

## single_lb_app.enable_discovery.api_crawler — api_crawler / 221211331230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.api_crawler

<a id="canonical-1133113221302203-2121312312021301-2311031123103033-0233220102113000-2223303212121000-3220031311310232-2120013233021012-1331001232233102"></a>

Type: `"object"`. single nested block, Optional.

API Crawling. API Crawler message.

Upstream description:

API Crawler message.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("api_crawler_config",
    "disable_api_crawler")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

Terraform syntax:

```terraform
api_crawler {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103213022023333-0120011132030300-0100311113311123-3231222233322111-0212203321013311-1122000212031222-1332312023103313-0321203020303203"></a>

## Direct properties — api_crawler / 221211331230 / 3

- [api_crawler_config](resources--http_loadbalancer--reference--group-027.md#canonical-2033123112032130-0113003000233101-0330031031202323-2112223102232220-2002023312033311-2000033123111201-1302130320201230-3320001020130112): complete subsection reference.

- [disable_api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-0202100011213003-3130210101301000-1131320200112003-0103010103133022-1321303322303211-3333110123203332-2331103013222131-3222203113011110): complete subsection reference.

<a id="canonical-3231222311311001-1200133232220323-0211121113301322-2003232000033323-3121301011032301-3321131301021003-2330223123222221-2321112333302212"></a>

## Next pages — api_crawler / 221211331230 / 4

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-027.md#canonical-2033123112032130-0113003000233101-0330031031202323-2112223102232220-2002023312033311-2000033123111201-1302130320201230-3320001020130112)
- [single_lb_app.enable_discovery.api_crawler.disable_api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-0202100011213003-3130210101301000-1131320200112003-0103010103133022-1321303322303211-3333110123203332-2331103013222131-3222203113011110)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2033123112032130-0113003000233101-0330031031202323-2112223102232220-2002023312033311-2000033123111201-1302130320201230-3320001020130112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013233311001301-3133021123320312-1133030011022011-2102202031122301-0300102330123012-3011032133231202-3130002001110102-0112220333102010"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config — api_crawler_config / 132203033311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config

<a id="canonical-2132231220302010-0232100003322120-2132031113123320-2112013200003031-0231322110202103-3010231001112313-0331030332122131-1023222113120211"></a>

Type: `"object"`. single nested block, Optional.

Crawler Configure.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
api_crawler_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113103203002222-3021222012213103-3331333003010003-3230212222022020-1212303102101221-3233211320231100-2011323211100011-1302031223112213"></a>

## Direct properties — api_crawler_config / 132203033311 / 3

- [domains](resources--http_loadbalancer--reference--group-027.md#canonical-2123220030322000-1021313231133233-0031223003213010-0212112312101230-2002311230123033-3012302000131202-2123203120133100-3333100300210003): complete subsection reference.

<a id="canonical-3022312112322131-1012031030103033-0111032310321213-0220022323212130-0100000111302022-3010022322332103-2222303000301120-2232101133321222"></a>

## Next pages — api_crawler_config / 132203033311 / 4

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-027.md#canonical-2123220030322000-1021313231133233-0031223003213010-0212112312101230-2002311230123033-3012302000131202-2123203120133100-3333100300210003)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2123220030322000-1021313231133233-0031223003213010-0212112312101230-2002311230123033-3012302000131202-2123203120133100-3333100300210003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131013202033212-1210320200031311-2111010300211211-2302233303012201-3022333100102312-3233033001102122-2021012221131003-3203311230311131"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains — domains / 012321003102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-027.md#canonical-2033123112032130-0113003000233101-0330031031202323-2112223102232220-2002023312033311-2000033123111201-1302130320201230-3320001020130112)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-0210013311330213-2103213103010113-1212200100213221-3323311012331211-0212110003311321-2211233100112323-0321012211232120-2112132011113021"></a>

Type: `"object"`. list nested block, Optional.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330000101020120-3221203213313031-3212223130033122-2333132102021013-0012203100222202-1301320330110022-0222133003232230-3302131102103221"></a>

## Direct properties — domains / 012321003102 / 3

<a id="canonical-3032132013330323-0323303003310010-3202303223230321-2302333103331020-3013323100100021-2330121332031021-1201301301133301-1001310101000102"></a>

<a id="canonical-1302021033201330-0100021100121311-0211103212003323-1023212100221321-1013010032022300-0312330130222233-0332032233131221-2303010210230002"></a>

## domain property — domains / 012321003102 / 4

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](resources--http_loadbalancer--reference--group-027.md#canonical-2101003003311231-2012222323232001-0131232323112100-0031121001321121-2321110231310231-1102100200223121-0301332121121202-0311301230102310): complete subsection reference.

<a id="canonical-3001032322123212-2200002031020301-3131303202103230-2223323220313131-3313111003302023-1302313323033320-3330311230032003-2223023210123321"></a>

## Next pages — domains / 012321003102 / 5

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-027.md#canonical-2101003003311231-2012222323232001-0131232323112100-0031121001321121-2321110231310231-1102100200223121-0301332121121202-0311301230102310)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-027.md#canonical-2033123112032130-0113003000233101-0330031031202323-2112223102232220-2002023312033311-2000033123111201-1302130320201230-3320001020130112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2101003003311231-2012222323232001-0131232323112100-0031121001321121-2321110231310231-1102100200223121-0301332121121202-0311301230102310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010313211210303-2022020330201220-0201103010230103-2222221202233110-1231010122132320-1112111223213320-0303023101120200-1110032131120013"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login — simple_login / 103122030130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-027.md#canonical-2033123112032130-0113003000233101-0330031031202323-2112223102232220-2002023312033311-2000033123111201-1302130320201230-3320001020130112)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-027.md#canonical-2123220030322000-1021313231133233-0031223003213010-0212112312101230-2002311230123033-3012302000131202-2123203120133100-3333100300210003)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-2233322301223230-2000023031023112-2110132100130120-2211112030212121-0213200133201320-0312031310201111-0122020332201003-3320110112020231"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple login.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
simple_login {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301220312121210-1110223213022120-0013231321111300-2022012233323300-2320220022223232-2123020230221232-3112103300203210-3301030130102012"></a>

## Direct properties — simple_login / 103122030130 / 3

- [password](resources--http_loadbalancer--reference--group-027.md#canonical-3212023010311330-0110001320133130-3222031233120311-1132010133120332-0223233200303331-0131332200020212-0110002020231331-2203200001011110): complete subsection reference.

<a id="canonical-1000103212210223-1211030202211213-1203123321223321-2221311002032222-1111113102032331-2331231213030110-0002111203002130-2230113123201230"></a>

<a id="canonical-0301220210332133-1002003312130011-3203031320213022-0100021121202303-0110200321332121-2313111213203013-0103213210220130-0323322301203003"></a>

## user property — simple_login / 103122030130 / 4

Type: `"string"`. Optional.

Enter the username to assign credentials for the selected domain to crawl.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1113130330232232-2121202122031012-0002321010332012-1132232032211032-1112033200220210-1112100233031323-0201122331021102-0223221111323123"></a>

## Next pages — simple_login / 103122030130 / 5

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-027.md#canonical-3212023010311330-0110001320133130-3222031233120311-1132010133120332-0223233200303331-0131332200020212-0110002020231331-2203200001011110)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-027.md#canonical-2123220030322000-1021313231133233-0031223003213010-0212112312101230-2002311230123033-3012302000131202-2123203120133100-3333100300210003)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3212023010311330-0110001320133130-3222031233120311-1132010133120332-0223233200303331-0131332200020212-0110002020231331-2203200001011110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311110030232133-0003202302110121-3332120102322223-2021132110220030-1210231010133320-1213301302221230-3212230121202132-2031111331203000"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password — password / 030311012321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-027.md#canonical-2033123112032130-0113003000233101-0330031031202323-2112223102232220-2002023312033311-2000033123111201-1302130320201230-3320001020130112)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-027.md#canonical-2123220030322000-1021313231133233-0031223003213010-0212112312101230-2002311230123033-3012302000131202-2123203120133100-3333100300210003)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-027.md#canonical-2101003003311231-2012222323232001-0131232323112100-0031121001321121-2321110231310231-1102100200223121-0301332121121202-0311301230102310)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-3200033310111222-2001230211201221-2313121302123332-3022132222302332-0130233102211113-3331202030131100-1002000301200111-0030230321010333"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231110332322300-3011033122121003-3220220111030132-2232213332112001-1120033201303122-0010233132030201-1033320013101030-2112320112302122"></a>

## Direct properties — password / 030311012321 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-027.md#canonical-0232330022332001-1201010112322313-2203231020030020-1133210220032003-2030112300020113-3111033010003311-2203310322031133-1113211013130301): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-027.md#canonical-3021233013330333-1131330012202311-2023022301001332-3321032112020233-0330323130123211-1310122101233202-1110303102231122-1201130030022310): complete subsection reference.

<a id="canonical-0122230102020312-2320131300023113-0120131303003302-3021103100200332-3203023101131030-1011112312112332-2220230310020020-3310101130221201"></a>

## Next pages — password / 030311012321 / 4

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](resources--http_loadbalancer--reference--group-027.md#canonical-0232330022332001-1201010112322313-2203231020030020-1133210220032003-2030112300020113-3111033010003311-2203310322031133-1113211013130301)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](resources--http_loadbalancer--reference--group-027.md#canonical-3021233013330333-1131330012202311-2023022301001332-3321032112020233-0330323130123211-1310122101233202-1110303102231122-1201130030022310)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-027.md#canonical-2101003003311231-2012222323232001-0131232323112100-0031121001321121-2321110231310231-1102100200223121-0301332121121202-0311301230102310)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0232330022332001-1201010112322313-2203231020030020-1133210220032003-2030112300020113-3111033010003311-2203310322031133-1113211013130301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100331303200222-0223130021223011-1001012113021131-2103210011312303-2112313210231203-0011313310200213-3301102321213310-2321130321101022"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info — blindfold_secret_info / 131122230213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-027.md#canonical-2033123112032130-0113003000233101-0330031031202323-2112223102232220-2002023312033311-2000033123111201-1302130320201230-3320001020130112)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-027.md#canonical-2123220030322000-1021313231133233-0031223003213010-0212112312101230-2002311230123033-3012302000131202-2123203120133100-3333100300210003)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-027.md#canonical-2101003003311231-2012222323232001-0131232323112100-0031121001321121-2321110231310231-1102100200223121-0301332121121202-0311301230102310)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-027.md#canonical-3212023010311330-0110001320133130-3222031233120311-1132010133120332-0223233200303331-0131332200020212-0110002020231331-2203200001011110)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-0223031111130131-0012232130202113-1030231121021113-2121321201323212-1031333130022032-2321320103111110-0202022033103012-1131323303102321"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333002320033122-0031300122112221-0231000311203121-0303222131001020-3103123213300033-1010231200000101-2113323132002223-0223013200101112"></a>

## Direct properties — blindfold_secret_info / 131122230213 / 3

<a id="canonical-0101310311212331-2100330233220102-2312321220102031-3111032123031021-3031120221330000-1330020222110002-1103201023232000-2302010302333002"></a>

<a id="canonical-0131002101230101-3022123300322000-0333331121033130-1030330200321333-2213233301100132-1033120202122221-1200021221002112-1221202311201303"></a>

## decryption_provider property — blindfold_secret_info / 131122230213 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0212223103200101-1211000001012222-0212132130213022-3103320213112212-1210222121223200-2022330212130003-1213223323111011-3020120222012133"></a>

<a id="canonical-2001102310000132-3112122133322131-0232332201022220-0013213221212021-2001213032322330-0123101210010110-3231131303101012-3220133221233200"></a>

## location property — blindfold_secret_info / 131122230213 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3001210100213132-3311010233221110-0012201120101201-0220202120121111-1033313310310202-2130232133213203-0222023201121332-0111113311211101"></a>

<a id="canonical-3230113303002201-1020131311100231-0110033321301233-1211000210332220-1220020011220220-0122232130301101-2130200220320321-0011120023023301"></a>

## store_provider property — blindfold_secret_info / 131122230213 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1233122210002023-0010323101112332-0220200301020323-1012330313202223-0002103121302331-2012200123331322-1201301102100021-2333023012011110"></a>

## Next pages — blindfold_secret_info / 131122230213 / 7

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-027.md#canonical-3212023010311330-0110001320133130-3222031233120311-1132010133120332-0223233200303331-0131332200020212-0110002020231331-2203200001011110)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3021233013330333-1131330012202311-2023022301001332-3321032112020233-0330323130123211-1310122101233202-1110303102231122-1201130030022310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320132321122133-2300121221203033-2021112301230031-0023112011212211-3310231002322002-2013112110003322-2121320000002003-0120323221332013"></a>

## single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info — clear_secret_info / 012303213302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-027.md#canonical-2033123112032130-0113003000233101-0330031031202323-2112223102232220-2002023312033311-2000033123111201-1302130320201230-3320001020130112)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-027.md#canonical-2123220030322000-1021313231133233-0031223003213010-0212112312101230-2002311230123033-3012302000131202-2123203120133100-3333100300210003)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-027.md#canonical-2101003003311231-2012222323232001-0131232323112100-0031121001321121-2321110231310231-1102100200223121-0301332121121202-0311301230102310)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-027.md#canonical-3212023010311330-0110001320133130-3222031233120311-1132010133120332-0223233200303331-0131332200020212-0110002020231331-2203200001011110)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-3032310131033002-1113222122221130-0110312322122131-2003021130020301-2332111211302201-3013323233130320-2322213233131130-3300111031001001"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031102131220103-1130303002232131-1131203100021131-1121202311212201-1111011303310301-3003013220101101-2220021331202001-1031201123330223"></a>

## Direct properties — clear_secret_info / 012303213302 / 3

<a id="canonical-3311203011220130-0200130211103120-0323123233331202-2303110010312321-1032301122103203-0123302001031301-2120031130221223-0121210331303122"></a>

<a id="canonical-0003032001000233-2221223200211111-2123001103023310-3332013230310300-3110233233110221-3310201013131322-2332030130333010-1123131200301313"></a>

## provider_ref property — clear_secret_info / 012303213302 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0013323311331113-2313210311223020-1010312330112213-1302303310300301-1331013001300032-0131032203113121-1312001022202300-2231123033033033"></a>

<a id="canonical-2200312130120302-0132132221201031-3220101020012110-2222113220203210-0010313333022022-1121010113332233-2333033013231100-0010132133013211"></a>

## URL property — clear_secret_info / 012303213302 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1012121301113221-1322103310302113-1010110330100210-2331012331331131-1130122010012212-1331100332122113-1010200101020202-0300203021320132"></a>

## Next pages — clear_secret_info / 012303213302 / 6

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-027.md#canonical-3212023010311330-0110001320133130-3222031233120311-1132010133120332-0223233200303331-0131332200020212-0110002020231331-2203200001011110)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0202100011213003-3130210101301000-1131320200112003-0103010103133022-1321303322303211-3333110123203332-2331103013222131-3222203113011110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313031102302012-1303003103031033-0231032321132312-3301202331121223-1100112011223122-3023323222002310-0123312102200130-1132212213311231"></a>

## single_lb_app.enable_discovery.api_crawler.disable_api_crawler — disable_api_crawler / 330221303122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012)
- single_lb_app.enable_discovery.api_crawler.disable_api_crawler

<a id="canonical-2112300023301203-3311112302112203-1031203120103221-3123211332331010-0031021121202313-0033331233213121-2333320023001200-3333100312203203"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_api_crawler = {}
```

<a id="canonical-2000012203313221-3313301223103332-3122032033022023-3022121112322021-2223111010001023-2223333300002021-3111202110211021-0320010000212330"></a>

## Direct properties — disable_api_crawler / 330221303122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220021123231130-2000121312020302-3203233213122313-0222322323012300-0331013100303332-0200103200321032-3322022103332320-3100200233231221"></a>

## Next pages — disable_api_crawler / 330221303122 / 4

- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--reference--group-027.md#canonical-2331202111331331-1133302120003213-3010013033221332-1220323000132130-3321132103131213-0020230011000311-0022023133100211-1313031123033012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130013200131311-0302110210020200-3220200020231211-1033123120211323-1230310102310120-1023130232221103-0033303131313012-2212112130331100"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan — api_discovery_from_code_scan / 320210121233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.api_discovery_from_code_scan

<a id="canonical-0011122103200213-2303002202230301-0223021011330121-1130100121333300-0011001003121303-1231013131113233-1232132131232013-3102310113103211"></a>

Type: `"object"`. single nested block, Optional.

Select codebase and Repositories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("code_base_integrations")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
api_discovery_from_code_scan {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303302013311230-3011030113321100-0311111233013011-3120002332322311-3303330003322302-3333212002111030-2110212012313130-2222303320301113"></a>

## Direct properties — api_discovery_from_code_scan / 320210121233 / 3

- [code_base_integrations](resources--http_loadbalancer--reference--group-027.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222): complete subsection reference.

<a id="canonical-2332103311221300-0220010101231132-2213203102332120-1033012230222103-2303302312331313-3223002331130301-0122013100020321-2310020210130102"></a>

## Next pages — api_discovery_from_code_scan / 320210121233 / 4

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-027.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230020312133132-0330333310202210-0011101011010331-2310301103023302-3222031212030122-1012212021101111-3203133211301200-1123013300333111"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations — code_base_integrations / 323220210002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-027.md#canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-3213332301002032-3000003203210100-0022201231330031-0113331201323310-2002123122101320-2220233020101131-3123113133201121-3310203112100130"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for codebase integrations.

Upstream description:

Configuration parameter for codebase integrations

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_repos",
    "selected_repos")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
code_base_integrations {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001220130131231-2003122002010022-0113111102210033-1121122223212101-0333201323320332-2233323202112222-3020321333022111-3232202000231211"></a>

## Direct properties — code_base_integrations / 323220210002 / 3

- [all_repos](resources--http_loadbalancer--reference--group-027.md#canonical-2223021303113303-1200233201332222-1302013331020322-3221000122322211-1013301011300330-2303002012013300-0231220113321030-1113121130233313): complete subsection reference.

- [code_base_integration](resources--http_loadbalancer--reference--group-027.md#canonical-0103332111221020-2200131012210000-2200013132110303-2300120313323223-2103212130323213-1323003210322233-1130223200203102-0003221302203232): complete subsection reference.

- [selected_repos](resources--http_loadbalancer--reference--group-027.md#canonical-0120030311332310-3032221103100333-3300120312102000-1320231120202002-2210010032113201-1011110103311301-1303300223303313-2132011301200233): complete subsection reference.

<a id="canonical-2200200320022212-2323303021022000-2203300323312101-1303131332113212-3013220200033303-1312320131313133-3103120133203101-1130023332320102"></a>

## Next pages — code_base_integrations / 323220210002 / 4

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](resources--http_loadbalancer--reference--group-027.md#canonical-2223021303113303-1200233201332222-1302013331020322-3221000122322211-1013301011300330-2303002012013300-0231220113321030-1113121130233313)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](resources--http_loadbalancer--reference--group-027.md#canonical-0103332111221020-2200131012210000-2200013132110303-2300120313323223-2103212130323213-1323003210322233-1130223200203102-0003221302203232)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](resources--http_loadbalancer--reference--group-027.md#canonical-0120030311332310-3032221103100333-3300120312102000-1320231120202002-2210010032113201-1011110103311301-1303300223303313-2132011301200233)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-027.md#canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2223021303113303-1200233201332222-1302013331020322-3221000122322211-1013301011300330-2303002012013300-0231220113321030-1113121130233313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131101322212302-2312013010011133-1210200233303121-0323310312303000-3333320312330333-2331000000313223-2312113321013303-2133131333331203"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos — all_repos / 132203321320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-027.md#canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-027.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-3133322330221112-3330033120230133-0031101233223330-0230203021321023-3130210231021130-0102133032112123-0210323203010201-3303330003033000"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
all_repos = {}
```

<a id="canonical-0003031111303231-2212113231213131-3233023013120321-1201131330313001-1132101031013300-2310323120313001-3100200021012210-3213322023111111"></a>

## Direct properties — all_repos / 132203321320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301132330301132-1102220231221202-3001223113301023-0012120032312330-1332310002232021-0210333313013020-3300021023010133-2301103111031220"></a>

## Next pages — all_repos / 132203321320 / 4

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-027.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0103332111221020-2200131012210000-2200013132110303-2300120313323223-2103212130323213-1323003210322233-1130223200203102-0003221302203232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102320312020002-1000201310230020-3331000311311321-2110211133121311-3100122202013212-3132332212030321-3320201330010312-2103331302121233"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration — code_base_integration / 332133031313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-027.md#canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-027.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-2132332003233123-0003312302232301-0321102001120122-3322111331302320-1113110222020233-2011112033211010-3311110221031123-3103302313130123"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
code_base_integration {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021132110300300-0302000232031201-2021332210323013-1101123300132001-0201120200112232-3013300023301220-0013020232220310-2210010210212220"></a>

## Direct properties — code_base_integration / 332133031313 / 3

<a id="canonical-3213123230210323-2311033112130030-1102230222000001-0212313333311031-1322222322331311-0311223313030211-1100113230230300-1212030323030230"></a>

<a id="canonical-0321313300322310-3302011130032133-0220111220013030-3212020112323322-0320213002020010-0220132321033331-0213033210010110-3301331101213213"></a>

## name property — code_base_integration / 332133031313 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3232300333302113-2100102123230011-2211212021303030-3111330120223011-1130012013102013-1233113211113010-2013222310131213-0003310033203031"></a>

<a id="canonical-1132110312202310-3310022123313333-2331311231001221-0313000211321131-2011223300330313-3121111110021031-0032133103302131-1203101101133003"></a>

## namespace property — code_base_integration / 332133031313 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1300013232311013-2302233202320202-2300321113321101-1222010023133022-1021333032130023-1101201310211312-3101332030213012-2202312010110131"></a>

<a id="canonical-3220320101113113-0130221313310102-3301132000213330-2330013121211331-1102012130312310-0113330311010032-2320202333212212-1211221332212310"></a>

## tenant property — code_base_integration / 332133031313 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0002110202012033-1232132112310022-3222000221123120-0000303130202032-0300131111133100-3323230311332032-3221021003231211-2001223233013202"></a>

## Next pages — code_base_integration / 332133031313 / 7

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-027.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0120030311332310-3032221103100333-3300120312102000-1320231120202002-2210010032113201-1011110103311301-1303300223303313-2132011301200233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102312320223120-2231131122122330-0001003023022302-3033123022010032-0101132230011202-2230323013133011-2322230220133201-0323000303310112"></a>

## single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos — selected_repos / 321122300101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-027.md#canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-027.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-1331101231311200-0121211113322311-1332212230103313-2020020230112022-1201002122300103-0330330023121021-0021132130313202-0333133331022113"></a>

Type: `"object"`. single nested block, Optional.

Select which API repositories represent the LB applications.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("api_code_repo")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
selected_repos {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021320200332103-2031130222130201-3300033310021321-0300210201131220-2123231033212013-0222132232113100-1333003021202202-2122112311133022"></a>

## Direct properties — selected_repos / 321122300101 / 3

<a id="canonical-3102333323000130-3213103201301231-0031212023331023-2232320020322123-2311201131122203-2313131021331032-2230020303101100-1110133301011102"></a>

<a id="canonical-2020000122022033-3311212322111203-0201110222131133-2003133230332103-3101120221303212-2311202223120032-3332203203031121-0221213022303220"></a>

## api_code_repo property — selected_repos / 321122300101 / 4

Type: `["list", "string"]`. Optional.

Code repository which contain API endpoints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1013312011300312-1220212312013223-1131121330012131-1022321021212311-0011323103002103-3110230230102311-3211012320020000-2130003302220300"></a>

## Next pages — selected_repos / 321122300101 / 5

- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-027.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0033020300201021-2112213210103002-3302223002101203-1012232012320110-3312220020100332-0310003001132021-2230020302333111-0112132130323101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302320020012210-2030210332300302-0133210233003233-1220221033211210-3331300212322132-3112230110320220-0132023301313110-2321222000321113"></a>

## single_lb_app.enable_discovery.custom_api_auth_discovery — custom_api_auth_discovery / 020012033301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.custom_api_auth_discovery

<a id="canonical-0333321320210332-0013030320320113-0201130312310220-0110031121302131-3032233230211220-1230002200133303-1101002002202232-0311131011123200"></a>

Type: `"object"`. single nested block, Optional.

API Discovery Advanced Settings. API Discovery Advanced settings.

Upstream description:

API Discovery Advanced settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
custom_api_auth_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303223113213133-2031301303202312-0220002320212021-3210130110330311-1000331303012301-2230000202133001-3022132132130030-0211323013330003"></a>

## Direct properties — custom_api_auth_discovery / 020012033301 / 3

- [api_discovery_ref](resources--http_loadbalancer--reference--group-027.md#canonical-1310002311031003-3222203030001112-2330202232232302-2312330100030212-2010011320331203-2001310010001020-2133103030333320-1321330111213323): complete subsection reference.

<a id="canonical-2011323010331313-2201133212001233-3100031311211203-2203231101012010-2323320310221330-1220103123131101-3030001011302331-2031221021020230"></a>

## Next pages — custom_api_auth_discovery / 020012033301 / 4

- [single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref](resources--http_loadbalancer--reference--group-027.md#canonical-1310002311031003-3222203030001112-2330202232232302-2312330100030212-2010011320331203-2001310010001020-2133103030333320-1321330111213323)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1310002311031003-3222203030001112-2330202232232302-2312330100030212-2010011320331203-2001310010001020-2133103030333320-1321330111213323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231302311323013-2010230002011111-3010303121213113-2003203101212133-1000110101133213-2321223203321010-2331033311021032-0301302212223031"></a>

## single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref — api_discovery_ref / 203233313111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-0033020300201021-2112213210103002-3302223002101203-1012232012320110-3312220020100332-0310003001132021-2230020302333111-0112132130323101)
- single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-0322223301323110-2210100211023321-0132230321301312-2000301313322332-3331010202221113-3301023031101312-1320233002113230-1010130030112010"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
api_discovery_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003200121233303-3102101333230230-0330133220301312-1103211131201330-1131013312121022-1303210012213002-3310322103120220-3101310200230211"></a>

## Direct properties — api_discovery_ref / 203233313111 / 3

<a id="canonical-1121001212212111-3002301322203222-3113123102011333-0311123231310332-2030203103200322-0122131130023000-2231113121333130-1312120010022023"></a>

<a id="canonical-0301320312100133-3220022210210102-2013033000000021-1211123100032110-0033132201103131-3122013003203213-1223110233322233-1111110321223002"></a>

## name property — api_discovery_ref / 203233313111 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2232023303111032-1212313003103030-2113220210133302-3301323213000201-3213333323212033-2202320321011322-0021301121321010-0131103011132020"></a>

<a id="canonical-3020213032303223-1001200101013133-1221013133011021-0021012233332203-0332213203123103-1322003201231210-3300112210030210-2233100323002211"></a>

## namespace property — api_discovery_ref / 203233313111 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0212110031322202-0000011212312201-3110330333321323-2130320303130030-0320013030233031-2202020220230213-1122013223202123-3031310021031300"></a>

<a id="canonical-0231111201222113-2102013311013313-0333130031031233-0030210102120230-0103013022300233-2133101211323002-2320020132120313-0000233220011032"></a>

## tenant property — api_discovery_ref / 203233313111 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0032300220333010-3221111311012220-3301331101203333-0012220311022221-1031033303311323-1330131223232032-3200332202232120-2121132330001122"></a>

## Next pages — api_discovery_ref / 203233313111 / 7

- [single_lb_app.enable_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-0033020300201021-2112213210103002-3302223002101203-1012232012320110-3312220020100332-0310003001132021-2230020302333111-0112132130323101)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1312311021131010-1232123020112022-3111131320310202-0101232302110333-2132211003300303-0130323303100031-0002131301100313-0230303031012130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212212120120020-1200232323101232-3111323002230302-2312320233233313-0032311300203200-0001203031121203-2131210203321220-3330210213032033"></a>

## single_lb_app.enable_discovery.default_api_auth_discovery — default_api_auth_discovery / 202312020001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.default_api_auth_discovery

<a id="canonical-3331030310133231-1233101221112220-0111201030102033-2212120313302003-1322022120211030-1011010333230301-2100132013122312-3123230310133320"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
default_api_auth_discovery = {}
```

<a id="canonical-2201020231333320-0101320213232110-0312232333001312-1203113330130033-2213203022020232-2222203023000112-0312201003330201-2210010121111013"></a>

## Direct properties — default_api_auth_discovery / 202312020001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303332211312333-1200200113020102-3322130300030001-1101112213221301-2033220002320001-3013311030311110-1232121312112012-2222001202002030"></a>

## Next pages — default_api_auth_discovery / 202312020001 / 4

- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1131121200123133-1212031302303313-1013301330332331-1101002022012232-2321013232202222-2303001302131230-1130011122113321-2201223210020302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202003322013311-3102301321012110-3222120120200021-2110332132202100-1220232221223333-1001232003011322-2301213021023300-2223033200211011"></a>

## single_lb_app.enable_discovery.disable_learn_from_redirect_traffic — disable_learn_from_redirect_traffic / 101203220222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.disable_learn_from_redirect_traffic

<a id="canonical-3000103101111331-1221030123211303-2333230223222013-3213210022120132-2331200010130202-0110310122131211-3021332232110023-1301103212022110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learn from redirect traffic.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_learn_from_redirect_traffic = {}
```

<a id="canonical-0101323100201300-2221321000300010-0312010123210103-1023020310310201-2310333030333122-1222321323211302-2203101030313220-0331332021112230"></a>

## Direct properties — disable_learn_from_redirect_traffic / 101203220222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330223110310203-0222210221331212-2300230332322103-0302013132131313-1123320203222333-1301313100300312-1133113030313132-0100311233210230"></a>

## Next pages — disable_learn_from_redirect_traffic / 101203220222 / 4

- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0221012011230012-1010330112321300-3320100021211210-0310132001223101-1033302320113013-1203012103122302-1000022330113011-3012213201002311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120100110302300-2103021030103201-0221111130211223-3103321312111030-0003202111122220-0123113321300031-0110200220123321-0103303233323311"></a>

## single_lb_app.enable_discovery.discovered_api_settings — discovered_api_settings / 121113322303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.discovered_api_settings

<a id="canonical-0331303102032213-2110332322121012-1212111102130321-1230112302323030-3210330213100023-0313011323232303-1113113323023023-0330111013322020"></a>

Type: `"object"`. single nested block, Optional.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("purge_duration_for_inactive_discovered_apis")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
discovered_api_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133113300203010-3332232112213030-0012001230111222-3210323120202233-3210200202122122-2201323302232203-2102230001332220-0301123001332322"></a>

## Direct properties — discovered_api_settings / 121113322303 / 3

<a id="canonical-0230333102322131-1000222010223211-2222211313132201-0030111012330322-3312220001212133-1220031011220203-3302212223020203-0312332133033220"></a>

<a id="canonical-0131332130310122-0311321113323200-2120102233222332-3103233210002103-2111030131220331-3302123312032230-0020301111100320-2010223221213102"></a>

## purge_duration_for_inactive_discovered_apis property — discovered_api_settings / 121113322303 / 4

Type: `"number"`. Optional.

Inactive discovered API will be deleted after configured duration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-3230313313030020-1310131003220223-1232013201321310-1112310030023122-1002320030211313-2233212130102201-0321011002321231-0311312331332130"></a>

## Next pages — discovered_api_settings / 121113322303 / 5

- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1103000230303211-3013102200220030-3232211010330301-3323201112330001-3233012203213323-2002131102133010-3312223131112321-2202313010211300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203232100202301-1312321232301111-3113200002231002-0022031220122200-2230311132121013-2120312120303020-0012031311333211-1133030120221233"></a>

## single_lb_app.enable_discovery.enable_learn_from_redirect_traffic — enable_learn_from_redirect_traffic / 133010002021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.enable_learn_from_redirect_traffic

<a id="canonical-1120032233233120-1033130000101212-3222302220223310-0300230102312311-2133231313132030-3300100023231122-0021210120032233-1121002103332101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learn from redirect traffic.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
enable_learn_from_redirect_traffic = {}
```

<a id="canonical-0220110111133330-0231123203220222-3231310123302300-3021113101123132-2000100111310202-0011100130011310-3201212133221202-3100002331320113"></a>

## Direct properties — enable_learn_from_redirect_traffic / 133010002021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330101220213200-0010001002000313-2112323031022022-1223002321311302-2131303322111123-2113220000122121-0301310033113131-3110303002200213"></a>

## Next pages — enable_learn_from_redirect_traffic / 133010002021 / 4

- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0323203132202222-0231330232030303-0301303112213002-0030002232012332-3202313221231110-1231102102000012-1012221103330222-0003130020021302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021203323000233-3332132332322031-3103331003233022-0030122300231233-2310120230333011-3201223033010001-1033002311100202-1201022321300111"></a>

## single_lb_app.enable_malicious_user_detection — enable_malicious_user_detection / 122210303030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- single_lb_app.enable_malicious_user_detection

<a id="canonical-1330312121112131-2300221301232022-3132111230031311-1010132310122030-3101330022302222-1232210323232321-2220202023212022-0020111300333201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable malicious user detection.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
enable_malicious_user_detection = {}
```

<a id="canonical-2210103030101002-0100230213132102-3321002312022123-1323313010200313-1010332111210223-2000030121212001-2300321302310203-2130223222002331"></a>

## Direct properties — enable_malicious_user_detection / 122210303030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033230231203332-0000032122210112-0313313111233111-1113131120121133-3100233113020333-1312230310030203-0013100311233011-1212202220121122"></a>

## Next pages — enable_malicious_user_detection / 122210303030 / 4

- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3333311120100023-2230031010023220-2013013300321213-0302220321233313-0222113033133022-2131111320010323-2303320333133122-3002301330000322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301322122211210-3321122110231313-1211221132220222-2033002312011332-3113021021110002-3103120333001002-2103222222001331-3313002323130100"></a>

## slow_ddos_mitigation — slow_ddos_mitigation / 121101311112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- slow_ddos_mitigation

<a id="canonical-3321130210030111-2101320330200320-3223131102102320-2111002221120321-3012211013203212-2010231131011301-0312012001201331-3232322133122101"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("request_headers_timeout"),
  validators.ConflictingObjectAttributes("disable_request_timeout",
    "request_timeout")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

OneOf alternatives in this subsection:

- [slow_ddos_mitigation](resources--http_loadbalancer--reference--group-027.md#canonical-3321130210030111-2101320330200320-3223131102102320-2111002221120321-3012211013203212-2010231131011301-0312012001201331-3232322133122101)
- [system_default_timeouts](resources--http_loadbalancer--reference--group-027.md#canonical-0330212210310002-3223011012010302-3213000212311310-1111301100131233-0001230210102033-3202311013232013-0011111000002231-0320131031001100)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
slow_ddos_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320221012302020-0333231101012013-3320313102331320-2003200221003310-2321121100200131-2020220022121120-2132100011233002-0001023132111120"></a>

## Direct properties — slow_ddos_mitigation / 121101311112 / 3

- [disable_request_timeout](resources--http_loadbalancer--reference--group-027.md#canonical-1333223010311031-1223211010120203-2231021132203220-2012312213221002-1001311321000223-1010102333310300-2231212011132223-0110031230110023): complete subsection reference.

<a id="canonical-2332323132331313-2301011303102131-3122131232223101-0023003230013212-0122201330203221-3122011302000123-1333133112232020-3231221322233121"></a>

<a id="canonical-3121320133300331-1202311021031110-2122333330313013-3013233122032132-2303022032012121-2312000213220122-2332312023121313-1303111100302313"></a>

## request_headers_timeout property — slow_ddos_mitigation / 121101311112 / 4

Type: `"number"`. Optional.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(2000, 30000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-2123132102020221-3213012233113223-1313101031103031-2123333301100223-1002121220302013-2112132002133001-1230213032133122-1310103233322231"></a>

<a id="canonical-3002131221103110-0200201312121233-2103323103221101-2301120100131202-2320222330020202-2322022231122303-2213002302312322-3232011133212201"></a>

## request_timeout property — slow_ddos_mitigation / 121101311112 / 5

Type: `"number"`. Optional.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(2000, 300000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-2331102233210201-0111212310320022-3113122323021330-3110111103202032-1230021330312320-0030121331322020-1220331312021312-0013103223333133"></a>

## Next pages — slow_ddos_mitigation / 121101311112 / 6

- [slow_ddos_mitigation.disable_request_timeout](resources--http_loadbalancer--reference--group-027.md#canonical-1333223010311031-1223211010120203-2231021132203220-2012312213221002-1001311321000223-1010102333310300-2231212011132223-0110031230110023)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1333223010311031-1223211010120203-2231021132203220-2012312213221002-1001311321000223-1010102333310300-2231212011132223-0110031230110023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031223202123123-0011102303313002-0312012131031112-1112310332203323-3201321122123110-0220133222002313-2210021120103121-1220201231301012"></a>

## slow_ddos_mitigation.disable_request_timeout — disable_request_timeout / 300210213013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [slow_ddos_mitigation](resources--http_loadbalancer--reference--group-027.md#canonical-3333311120100023-2230031010023220-2013013300321213-0302220321233313-0222113033133022-2131111320010323-2303320333133122-3002301330000322)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-1101111231221203-1120230003312300-3132000101201213-2122000000111201-2222312231110333-3323010112121333-3200220210210310-0101120303120212"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable request timeout.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_request_timeout = {}
```

<a id="canonical-3101023013013110-2113200231110010-0130013301000000-3030232031102320-0223133203120022-3031223011323012-0120320112202233-3003121232202102"></a>

## Direct properties — disable_request_timeout / 300210213013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203312220233123-2201231013022023-3220132001101323-0230010032213210-2233212133321322-3321321131130111-2230113302110312-3132221223132210"></a>

## Next pages — disable_request_timeout / 300210213013 / 4

- [slow_ddos_mitigation](resources--http_loadbalancer--reference--group-027.md#canonical-3333311120100023-2230031010023220-2013013300321213-0302220321233313-0222113033133022-2131111320010323-2303320333133122-3002301330000322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0120111122231002-2333331020233201-3020032131233000-0232201223221311-1012212322212322-1113023212320102-3111221221223322-0223221231203000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022003032210231-0002213112213103-3210301300021330-0213022010203321-0312301010002130-2213030231310212-1210213303210202-0211013020303303"></a>

## source_ip_stickiness — source_ip_stickiness / 301133113223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- source_ip_stickiness

<a id="canonical-0202231310310202-3102203321100120-0022003303331033-1312113312302011-2331102102230311-0003123123020233-2303210033023103-2322033033310111"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
source_ip_stickiness = {}
```

<a id="canonical-2122030112021320-1112200200113122-1031111021321221-2100221220302112-1011333122333321-0301211002201033-1022233200130133-1213100101130220"></a>

## Direct properties — source_ip_stickiness / 301133113223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011123330332231-1031132212310312-3320313102333013-0001330222203221-1320131203332101-0330211113202013-1301333010100310-2123033203023322"></a>

## Next pages — source_ip_stickiness / 301133113223 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3111101023101332-0113013332131030-0031310130033201-2323032100231211-0300131012212002-1223021312230103-2033033310310331-1121322330011111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102011020033302-3020300123011230-1031012231222021-2131102112001303-0132100032300112-0012233032230231-3230113120230320-3013323020132011"></a>

## system_default_timeouts — system_default_timeouts / 112120132322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- system_default_timeouts

<a id="canonical-0330212210310002-3223011012010302-3213000212311310-1111301100131233-0001230210102033-3202311013232013-0011111000002231-0320131031001100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for system default timeouts.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
system_default_timeouts = {}
```

<a id="canonical-1030022203223213-3332010131113021-0133123023010030-2312220332232222-0311330332333312-1331003132231321-3101030130020032-3131113133302330"></a>

## Direct properties — system_default_timeouts / 112120132322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111030222301200-1220233213202330-1123320213011101-0312002101323322-3323333010011233-1221232201003133-2330132033121020-1202313210222223"></a>

## Next pages — system_default_timeouts / 112120132322 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0031000301301132-0002232123000132-3130011021001231-3113311013210310-0122300301203211-0213222010232211-0313220210221220-2120201133330322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013013123312311-3121031012202211-2230023031213113-2131331013201120-2203310300300332-0131210023112331-1023333211120121-2202001131101032"></a>

## timeouts — timeouts / 011332010202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- timeouts

<a id="canonical-3221131003030223-3133021020333000-3010231230322102-0101101302101223-0203322021001000-2003133323221213-1313313210212103-1213330213302210"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322030001231022-0203111310131032-1303332202100112-0301200021121303-3110002030033111-2013000331131323-1301322231232220-1021213121023130"></a>

## Direct properties — timeouts / 011332010202 / 3

<a id="canonical-1332203320130332-0231203032112023-3301120321320013-1010100230012123-3201311200302220-1311031011012131-1013132200020001-3332033302320330"></a>

<a id="canonical-0132103323232221-3323101030003201-1300003303223301-3023233310002023-1133021213131230-0103103213001010-1300023300012100-0221023323030111"></a>

## create property — timeouts / 011332010202 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1121333232030121-3033032203200111-3113230100002100-2230132212133120-0113312011022300-0100013303111121-1302333013212322-2232203033310130"></a>

<a id="canonical-2100302223210101-1213212113010132-1131011213301001-0113232313311023-1112203302301132-2110313131021300-0310031121310033-0033132320200133"></a>

## delete property — timeouts / 011332010202 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2321131303121121-1001333213122331-1301312312013203-3120122000313320-1232102131222323-3110303310102123-0113102332322322-0222320301231022"></a>

<a id="canonical-2012300201120003-3001132331023322-1100032330220320-2102023231332033-1130100033311231-2111203221012213-2312210003132310-0322200022101331"></a>

## read property — timeouts / 011332010202 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2123013211302232-0200232332112232-3323213233023001-2203212000321322-2020212100021000-2112120103222120-2200331123330000-2100310120030113"></a>

<a id="canonical-2223122221001020-1310123013221201-3211130130031211-2022012303212122-3132013320131021-0001102213122231-1012212220100030-0103133001202102"></a>

## update property — timeouts / 011332010202 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2231013022220120-0133302330333021-1220312330331103-1103111002323033-3313310212213023-2300311311212232-2112330310333300-2223320312013030"></a>

## Next pages — timeouts / 011332010202 / 8

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030220021312223-1223200321322010-3232012311021321-0031033032333101-3321033130003313-2123031123321002-1311111002302202-2132231111323133"></a>

## trusted_clients — trusted_clients / 012210111111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- trusted_clients

<a id="canonical-0133130032112133-1100010030010311-2103222010210003-2011202010331100-2122300201100231-3112021222021103-0323312032332332-2223111331101233"></a>

Type: `"object"`. list nested block, Optional.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc.

Upstream description:

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("actions"),
  validators.ConflictingListObjectAttributes("as_number",
    "http_header"),
  validators.ConflictingListObjectAttributes("as_number",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "skip_processing"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("http_header",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ipv6_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("skip_processing",
    "waf_skip_processing")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
trusted_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110002102023301-1103133303232221-3203333113031301-1232303103122320-0201323003102230-1211311021200131-3033032013110312-3121322100223031"></a>

## Direct properties — trusted_clients / 012210111111 / 3

<a id="canonical-0020211212331332-0001212123012033-1102123321023121-1012121121222022-1333323303102323-2322230211230310-3133032323232122-3012001301021202"></a>

<a id="canonical-0020313110231301-0232303331101123-3101022211103221-3211233303323231-2223321302220201-0331030220312012-3313332122211310-1131032121112131"></a>

## actions property — trusted_clients / 012210111111 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Upstream description:

Actions that should be taken when client identifier matches the rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(10),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2011023331310301-3131222211031201-2200311100233212-0111030031020211-2300320002220201-1323022011300130-1332023212133212-3210002221032012"></a>

<a id="canonical-3003303222123011-3230033031321301-2111012120321302-1020331032230110-3130013231222022-3111211131211321-2200202233113100-1133033313132202"></a>

## as_number property — trusted_clients / 012210111111 / 5

Type: `"number"`. Optional.

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Upstream description:

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 401308),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](resources--http_loadbalancer--reference--group-027.md#canonical-1002300133313303-0010300302103203-3100123023033331-0200331101103131-1113001202202300-0202002021320031-1100312103103123-1232103311233301): complete subsection reference.

<a id="canonical-2303203232120000-0103130330122113-1102113131033333-3313200130110112-1312113122323022-3132013202323112-0230130122000013-2203212030320111"></a>

<a id="canonical-2001321031321331-1310031320101332-2010232222301323-0303120123232101-1201123023031031-0012030132221130-0310130213110130-3031311230223113"></a>

## expiration_timestamp property — trusted_clients / 012210111111 / 6

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](resources--http_loadbalancer--reference--group-027.md#canonical-0130212200202101-2221112101301312-2221330231323300-2033221313303030-0332100303130330-3000322033023322-2211001320111312-0312213323202112): complete subsection reference.

<a id="canonical-0230231100332211-3101212220321002-3332133200203123-2111113200301010-1333223100233032-2202112002332333-1323333121220003-2213102002232210"></a>

<a id="canonical-3300221302000030-1230322122322312-3303122213023010-1133301323332002-0022003222123132-1102103101221122-1331132111321003-2131122210000200"></a>

## ip_prefix property — trusted_clients / 012210111111 / 7

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header IPv6\_prefix user\_identifier\] IPv4 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header IPv6\_prefix user\_identifier\] IPv4 prefix string.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2320331030201033-0012031012121333-2330212233220010-2011302302112000-0122320301302201-2022311113012122-0210320012103203-0122331331201333"></a>

<a id="canonical-1132120031300303-2021300212100030-0221103033222032-1132310320223102-2321203333033223-0102221111101133-1033300032000022-1132330232311201"></a>

## ipv6_prefix property — trusted_clients / 012210111111 / 8

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-027.md#canonical-3323212001322300-3233202033131233-0023032330121023-1010133000323213-0130121100033002-1111332233212002-2200301302130230-1021031221023213): complete subsection reference.

- [skip_processing](resources--http_loadbalancer--reference--group-027.md#canonical-3213233031322121-3133002030130220-0121202211301220-3233302310133200-3013022330013131-0033013222122233-0301221231302101-3333021203323021): complete subsection reference.

<a id="canonical-2001202032132210-1011311211101023-2030131000302121-2131331331111311-1322101112123120-0321123013323023-0000101330233003-1312132232312220"></a>

<a id="canonical-2032200301311132-3100201033200222-3202101312010223-0311111212200003-2122311023320011-2200301233003321-0123323023302300-3000021122332330"></a>

## user_identifier property — trusted_clients / 012210111111 / 9

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](resources--http_loadbalancer--reference--group-027.md#canonical-0230200100110103-0330023321103022-2112322010330320-0212120122121323-3130132223211233-0010201332230200-3102203100220300-1231303111321002): complete subsection reference.

<a id="canonical-1332201333101100-2211011003113111-2131103321002102-3120122133123001-3123122032300120-2211101223123021-3011030022232132-1211013121003113"></a>

## Next pages — trusted_clients / 012210111111 / 10

- [trusted_clients.bot_skip_processing](resources--http_loadbalancer--reference--group-027.md#canonical-1002300133313303-0010300302103203-3100123023033331-0200331101103131-1113001202202300-0202002021320031-1100312103103123-1232103311233301)
- [trusted_clients.http_header](resources--http_loadbalancer--reference--group-027.md#canonical-0130212200202101-2221112101301312-2221330231323300-2033221313303030-0332100303130330-3000322033023322-2211001320111312-0312213323202112)
- [trusted_clients.metadata](resources--http_loadbalancer--reference--group-027.md#canonical-3323212001322300-3233202033131233-0023032330121023-1010133000323213-0130121100033002-1111332233212002-2200301302130230-1021031221023213)
- [trusted_clients.skip_processing](resources--http_loadbalancer--reference--group-027.md#canonical-3213233031322121-3133002030130220-0121202211301220-3233302310133200-3013022330013131-0033013222122233-0301221231302101-3333021203323021)
- [trusted_clients.waf_skip_processing](resources--http_loadbalancer--reference--group-027.md#canonical-0230200100110103-0330023321103022-2112322010330320-0212120122121323-3130132223211233-0010201332230200-3102203100220300-1231303111321002)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1002300133313303-0010300302103203-3100123023033331-0200331101103131-1113001202202300-0202002021320031-1100312103103123-1232103311233301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113123202111122-1131222332031110-1102120102122013-3113002102212110-1122011013202231-2213110212201212-0332101110203330-1101310000023010"></a>

## trusted_clients.bot_skip_processing — bot_skip_processing / 202311123210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- trusted_clients.bot_skip_processing

<a id="canonical-3031013201211122-1020133023133202-2221032003033203-0123302210320321-1332223113130321-2011310210021333-3332332023001312-1000300321213133"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
bot_skip_processing = {}
```

<a id="canonical-0010123112033012-2010112301003232-1312302013313130-3023102311331102-0101013303130231-0212200003212102-1332031300002132-0022013003013002"></a>

## Direct properties — bot_skip_processing / 202311123210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320133113330012-1303131012213211-1122310011133102-0301113011222313-3123121222320010-0220100010332033-2013011011310103-2111300322330330"></a>

## Next pages — bot_skip_processing / 202311123210 / 4

- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0130212200202101-2221112101301312-2221330231323300-2033221313303030-0332100303130330-3000322033023322-2211001320111312-0312213323202112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312210233201331-0133323212320001-2000003300002211-3230011011313313-1322233031233033-3223111211221302-0103222120022111-2102120013101220"></a>

## trusted_clients.http_header — http_header / 221220113003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- trusted_clients.http_header

<a id="canonical-1231110113303031-3023112300000001-3210300020332330-2130013022123330-1123122302101101-2031120032122222-0021100220000020-1120021301120000"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113211220311201-2322022022200120-2121212000232200-2223032223003330-0101333001312100-2212003330133330-3321231202201222-1211210121100130"></a>

## Direct properties — http_header / 221220113003 / 3

- [headers](resources--http_loadbalancer--reference--group-027.md#canonical-2213310213001202-2300303212123322-3020023121302202-3313131031202312-3321001301213121-2221133112122120-1313021233202331-1130223200231021): complete subsection reference.

<a id="canonical-1132233121230203-3310000032012011-2020203002330330-2212223203030331-2123100111022003-0100333101321330-3033113132023302-1021211331212210"></a>

## Next pages — http_header / 221220113003 / 4

- [trusted_clients.http_header.headers](resources--http_loadbalancer--reference--group-027.md#canonical-2213310213001202-2300303212123322-3020023121302202-3313131031202312-3321001301213121-2221133112122120-1313021233202331-1130223200231021)
- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2213310213001202-2300303212123322-3020023121302202-3313131031202312-3321001301213121-2221133112122120-1313021233202331-1130223200231021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002302003333012-2121111031022122-1101033223310330-1103323322310302-2223200031221133-3330323302031032-0110000203212220-2023111230233130"></a>

## trusted_clients.http_header.headers — headers / 113012020201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- [trusted_clients.http_header](resources--http_loadbalancer--reference--group-027.md#canonical-0130212200202101-2221112101301312-2221330231323300-2033221313303030-0332100303130330-3000322033023322-2211001320111312-0312213323202112)
- trusted_clients.http_header.headers

<a id="canonical-2122010131223230-0300030103331221-1333011221210123-3322322100000223-3132200310020010-3230030200020123-2123230322301033-0300311123230212"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020231303013011-2223213001021021-1120301212212032-1013330033121302-2213330331302302-3302030022033210-1132332211031012-0100302211002131"></a>

## Direct properties — headers / 113012020201 / 3

<a id="canonical-2330212113321002-2210022133333131-1101131301012321-0113103333303201-0133013110133033-3133121310210020-0302111230031012-2322122110101030"></a>

<a id="canonical-2311303233133013-2320112030331122-2302222011202010-2033020300213103-2131132102122023-1212201130211333-0320200313100023-1133300213211302"></a>

## exact property — headers / 113012020201 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regular expression\] Header value to match exactly.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3022311211221300-1003313211013122-1311320210033111-3110111233010013-2032021312022002-0231311320232023-1033210221031331-2012121202322001"></a>

<a id="canonical-1313132213103300-1220021010032001-3021103310321321-1222233231201333-1230112210330322-0022323000011303-2312033210320010-0123131223031232"></a>

## invert_match property — headers / 113012020201 / 5

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3133123211022010-0000101223010313-0331300322002003-1012131101300202-0131011103333001-3232113320231303-0012001222000311-2132223301100331"></a>

<a id="canonical-1020322330102033-3122100231023213-0130133232031110-1003122323231323-2212032111020322-2101023313001101-3101200100003110-2223032003203232"></a>

## name property — headers / 113012020201 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2010220121021123-2201031010100222-1111303313132232-2301121131231301-0310230123010331-0002113212302322-3133003303133313-0210301313031210"></a>

<a id="canonical-0020330321010033-0102213033232132-2321003132131213-3312121132110311-0031003331003323-0330011202313310-1300313310232022-1110010222012222"></a>

## presence property — headers / 113012020201 / 7

Type: `"bool"`. Optional.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regular expression\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0003121223122333-0200232030333201-2012031230203012-0222320200323330-3213100212132131-3222331202232321-1223211001223231-2020212111113003"></a>

<a id="canonical-1112322001032103-3323031323230302-3210120020003000-2230020222111201-0033111132123122-1300000033020002-0032002303201210-3310322211331132"></a>

## regular expression property — headers / 113012020201 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3301113003112000-2330233002331302-1001301001203132-3011130222333223-3033011222211023-1030122331013222-1022023001121220-2110313012013320"></a>

## Next pages — headers / 113012020201 / 9

- [trusted_clients.http_header](resources--http_loadbalancer--reference--group-027.md#canonical-0130212200202101-2221112101301312-2221330231323300-2033221313303030-0332100303130330-3000322033023322-2211001320111312-0312213323202112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3323212001322300-3233202033131233-0023032330121023-1010133000323213-0130121100033002-1111332233212002-2200301302130230-1021031221023213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120221311232122-3212230022100102-1033121130312031-2223031131121211-2303211010101201-1000332012211020-0203221101102022-0313311131220002"></a>

## trusted_clients.metadata — metadata / 232202232333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- trusted_clients.metadata

<a id="canonical-2221031221101212-1202121220300331-2002011023112131-3033032000322300-2003323200302310-0122221331031300-3022223120333100-3023300311332211"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310300131233001-3033110120200320-0323103100002203-3123332311320231-0320003303311213-2000122220131121-1102133121131103-0221232101103103"></a>

## Direct properties — metadata / 232202232333 / 3

<a id="canonical-2231131022002100-3213000102012201-3100201301200221-2000221112233031-2020031333133103-0211333333013332-3121023230223102-2123213233210100"></a>

<a id="canonical-3311023131203330-3020111132231221-1120132102101221-1330001311223303-0300031002011031-2033320313330033-3031321001021233-0322323001222103"></a>

## description_spec property — metadata / 232202232333 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2011110131120202-3011301122233232-3323002212233210-0302113003032220-2302301311123100-2323011022311213-2302031003013033-0030201210020033"></a>

<a id="canonical-2033032113233301-1202100031102232-1132320330013303-3101223001332332-1030223223222021-1122321102201003-1300211022202123-2012031100200200"></a>

## name property — metadata / 232202232333 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-3131233323101231-2110203233133020-1201113301030330-2233301010232012-0231222221112323-2311002210011231-2032132010102122-3031320013321231"></a>

## Next pages — metadata / 232202232333 / 6

- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3213233031322121-3133002030130220-0121202211301220-3233302310133200-3013022330013131-0033013222122233-0301221231302101-3333021203323021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032213001303200-0300001323113020-0310111230010232-3111113221102211-1312121333303320-0220013331011302-2122313311331100-3121123202021312"></a>

## trusted_clients.skip_processing — skip_processing / 213212002130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- trusted_clients.skip_processing

<a id="canonical-0213321111333300-2101220101311003-1321323230120103-2302130113131021-1110111233012020-0232213123221323-3021100201031223-2322103000000130"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
skip_processing = {}
```

<a id="canonical-2121230001332203-2302113311221003-3330202230110011-1311300212032011-0111030111112210-3221022310131221-1313231012221300-2302010331001300"></a>

## Direct properties — skip_processing / 213212002130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330221320323101-3210032223303111-1120320333030211-0110310012223210-2002131232100213-1113120131301133-2323031201022300-3013130300213223"></a>

## Next pages — skip_processing / 213212002130 / 4

- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0230200100110103-0330023321103022-2112322010330320-0212120122121323-3130132223211233-0010201332230200-3102203100220300-1231303111321002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230022113001323-0033303020101210-1013132000211213-0130300111001203-2311222300013313-0002121111312233-1311213012221213-1233112213201122"></a>

## trusted_clients.waf_skip_processing — waf_skip_processing / 113123110333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- trusted_clients.waf_skip_processing

<a id="canonical-2131003200000233-0110223232333210-3221001033333233-0100210222332310-0211313021202222-0233112032003203-3213000030222300-2110031213300020"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
waf_skip_processing = {}
```

<a id="canonical-0332130222102200-1233113222111002-1311200200113231-1032103331132232-3303020102221233-3330100013331310-3012120110301023-2120310121333330"></a>

## Direct properties — waf_skip_processing / 113123110333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132101132031132-3322300232231221-2302130112220000-0032202200111101-0223311231011333-0301213011330330-1110002311110113-1100230020033000"></a>

## Next pages — waf_skip_processing / 113123110333 / 4

- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3332013311013213-1233202210102101-2202301332332033-2130302002311222-1231331220201121-2021122303232113-1311030310321233-1201033022300000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202201033211023-3230101012332200-3132310110012021-3000032213111021-1232330003013002-3331321032022000-2100120210333132-3332002212032033"></a>

## user_id_client_ip — user_id_client_ip / 122222300230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- user_id_client_ip

<a id="canonical-0022033130131210-3032332201221021-3230233002311132-0130000033332332-3330030231331101-0022000213122321-0301012211103113-2011213130320010"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: user\_id\_client\_ip, user\_identification\] Enable this option. Defaults to \`map\[\]\`.
Server applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

OneOf alternatives in this subsection:

- [user_id_client_ip](resources--http_loadbalancer--reference--group-027.md#canonical-0022033130131210-3032332201221021-3230233002311132-0130000033332332-3330030231331101-0022000213122321-0301012211103113-2011213130320010)
- [user_identification](resources--http_loadbalancer--reference--group-027.md#canonical-1333212101102101-1201312203331203-0313313322223030-1231103321303311-1303322331300113-2010133022030221-3223013213222031-2231132130202030)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
user_id_client_ip = {}
```

<a id="canonical-1111303021300201-3221132132110312-1001202220021030-2321012312113212-3110121321203312-0103033131231130-0320000301312113-1223000313231103"></a>

## Direct properties — user_id_client_ip / 122222300230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310111003301101-0333002113012033-1201101313111310-2321002013323133-3320301000221313-3123331223031313-0031112313232332-1333103123013003"></a>

## Next pages — user_id_client_ip / 122222300230 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0032001213002301-0331120033322333-0030100320321033-2022130323111321-1033121323230311-3230223120103201-0000013032330321-1020332302220223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202323210303323-2030223003111121-1013131033220122-2323131000113022-2023222201320031-2232200122310312-1100121110220021-3101100133210123"></a>

## user_identification — user_identification / 331333222023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- user_identification

<a id="canonical-1333212101102101-1201312203331203-0313313322223030-1231103321303311-1303322331300113-2010133022030221-3223013213222031-2231132130202030"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
user_identification {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212232323203133-1320333022121110-3210023110322111-2311303310222113-0320323020330120-0123323033003210-2300132031120331-3330331130113201"></a>

## Direct properties — user_identification / 331333222023 / 3

<a id="canonical-1120231302022003-3200330131332122-2003121020012202-3301013213330122-1031231313131102-1312020023320311-0032030022012213-3020000112320313"></a>

<a id="canonical-3023212032113022-1112120032211033-1323332221322101-1122113201121201-3223211310321312-0303101003123221-1021300121313330-1130300100003102"></a>

## name property — user_identification / 331333222023 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3000331120030103-1033030030120303-1230203231331010-1311001233121222-1011000203011101-1300311130031123-3320302130130220-3323111312310103"></a>

<a id="canonical-1332232000332020-1332311303223200-1223230033313220-2302122133230212-2210311132303321-0032332300302000-0222212213130202-3233212121110000"></a>

## namespace property — user_identification / 331333222023 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0312011300203320-1211333023232132-0012332000101133-1022012332000100-1200221223321321-3231032223311311-1011122102203130-3023013001132330"></a>

<a id="canonical-3122023323311132-1201332133022100-0333031010221022-2230023201202222-2013332332202333-0200023130122323-0013110333020200-3312112201110332"></a>

## tenant property — user_identification / 331333222023 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3113210202210223-2130211132100012-0310232012310310-2313320002032221-2323011331100013-0203002123212200-1223012033020233-3312331111103022"></a>

## Next pages — user_identification / 331333222023 / 7

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320100302310201-1033012213310333-1002300002010320-3201200133002100-1030333121130013-2320311203222010-2011322121103331-0132321232310121"></a>

## waf_exclusion — waf_exclusion / 200100230212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- waf_exclusion

<a id="canonical-3030320320020232-2321302200221220-3233200121302202-2302000023301023-2230300231002210-3331111223323213-3001101212233312-2230201002203010"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for waf exclusion.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("waf_exclusion_inline_rules",
    "waf_exclusion_policy")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

Terraform syntax:

```terraform
waf_exclusion {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133010003020112-1322232110031003-2203131303310133-3303100132022110-0123022032021220-1103010330321020-0211303023322001-3230232120100322"></a>

## Direct properties — waf_exclusion / 200100230212 / 3

- [waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-027.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210): complete subsection reference.

- [waf_exclusion_policy](resources--http_loadbalancer--reference--group-028.md#canonical-2300231203030031-1332023310203311-2011202020030201-3003033103130323-3301122311203303-1021210131330310-0012111123120220-1030113003313033): complete subsection reference.

<a id="canonical-2003012220210121-3310311103132022-2113021031213121-3200121313201202-0213102231130013-1131202120203122-2100123300102331-3332203001211001"></a>

## Next pages — waf_exclusion / 200100230212 / 4

- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-027.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_policy](resources--http_loadbalancer--reference--group-028.md#canonical-2300231203030031-1332023310203311-2011202020030201-3003033103130323-3301122311203303-1021210131330310-0012111123120220-1030113003313033)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221233320213301-0201010232232010-2103333111322313-0013133100330102-3201203102313023-1222301203133102-3010120113313313-3020221302101303"></a>

## waf_exclusion.waf_exclusion_inline_rules — waf_exclusion_inline_rules / 302112003313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-027.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- waf_exclusion.waf_exclusion_inline_rules

<a id="canonical-0011011011112311-3302132133230202-0303221210330032-3123333313220001-0002232101233332-3233223233201001-1221202201313213-1002023021110031"></a>

Type: `"object"`. single nested block, Optional.

List of WAF exclusion rules that will be applied inline.

Upstream description:

A list of WAF exclusion rules that will be applied inline.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
waf_exclusion_inline_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110201312312313-1120210223201121-2002120333011210-2131222330113300-2101100012233002-1103322003233333-1100102131320311-2130202100112212"></a>

## Direct properties — waf_exclusion_inline_rules / 302112003313 / 3

- [rules](resources--http_loadbalancer--reference--group-027.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102): complete subsection reference.

<a id="canonical-1112032130200012-0312111103321111-0103220330232133-3010100312101113-0132323130220300-1001303111100322-2102233132133030-2102303232003111"></a>

## Next pages — waf_exclusion_inline_rules / 302112003313 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-027.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- [waf_exclusion](resources--http_loadbalancer--reference--group-027.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320202020230112-3113331300323122-1120023133202003-1222132103311311-3030211010223120-3223000300303002-1122002033100011-3203233130310300"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules — rules / 223103021000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-027.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-027.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="canonical-1102113331130131-2023231231001111-0002223222222100-1112210131212010-3122031212021203-0032031100221221-1011022122301233-2302320121123031"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of WAF Exclusions specific to this Load Balancer.

Upstream description:

An ordered list of WAF Exclusions specific to this Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex"),
  validators.ConflictingListObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_prefix",
    "path_regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203210322312323-1221101012021112-1110312201330130-1212321333132302-0333123012000022-1100002231001222-0110211330320230-1032221310000311"></a>

## Direct properties — rules / 223103021000 / 3

- [any_domain](resources--http_loadbalancer--reference--group-027.md#canonical-1233330223321302-1130120332310310-0203100300123123-2130331311030110-0132230321210230-2132221103000313-1003202111301303-2033023223131322): complete subsection reference.

- [any_path](resources--http_loadbalancer--reference--group-027.md#canonical-0310132112001312-0121211110302132-3301100330233002-0032202211303222-2023031010201303-0213032213310021-2210333223313222-3110023101111231): complete subsection reference.

- [app_firewall_detection_control](resources--http_loadbalancer--reference--group-027.md#canonical-0203003232313033-0223230331033131-3221113013002333-3200303010013212-3202033211033320-3210102132113031-2230233013221023-0302302302002122): complete subsection reference.

<a id="canonical-1321111013100100-2122003310012003-3100010203122031-1021322011133013-2333322102330331-0213021131312100-0112120103212213-1032301311032130"></a>

<a id="canonical-0003010001000230-2300032003022212-1012132133012033-2303231020221301-2301203210221211-2002031022001113-1113032012003220-2120210221032332"></a>

## exact_value property — rules / 223103021000 / 4

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3113021303211310-2333021311003010-0133013013113132-1012210011323130-1330121012312102-0030220203123001-0030103322010220-1103101313322123"></a>

<a id="canonical-3230110323211131-0331000103210230-0222120321012301-2323231103120123-2203210013031223-1330013220223212-3223011202000313-2101131322311333"></a>

## expiration_timestamp property — rules / 223103021000 / 5

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-028.md#canonical-3120302131000211-1331213131113022-3323221113300010-2300303303032223-3201231223203103-1003211220131202-2033111223021311-0032101121231310): complete subsection reference.

<a id="canonical-0033221133033211-1110103031332331-1211131223320320-2212123102122212-2122300330033321-2112000333010122-3332002111011233-2231313232201310"></a>

<a id="canonical-2123013023333231-1202301003331220-1102231100201232-3322300100032031-1021231122002211-2130100100213103-3013203321021322-2132022121013030"></a>

## methods property — rules / 223103021000 / 6

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1223010022210132-2101321313011111-0332222010033011-1203032320130331-0300112011023012-2030330311020113-3313301101223320-1103102231210033"></a>

<a id="canonical-1213313311330020-2312210112020023-2233001011102003-0231003003231032-2132003000122133-3102122330101023-1313120331230310-3113012230031002"></a>

## path_prefix property — rules / 223103021000 / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0202230033212332-3302232121130333-2302131313300130-0313233322112220-0032101233221101-1212020103103013-3203230312220022-3203120113002023"></a>

<a id="canonical-2203000222300111-1122232001131211-2001330331103202-1202300221200210-2233013031001110-2312233033031213-2130013303222022-3002113303123233"></a>

## path_regex property — rules / 223103021000 / 8

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0320110112032110-0101110310133113-2222203231022113-1013102212020002-1231121322330221-2122031001203003-0030231231102200-3312110133100200"></a>

<a id="canonical-1123013123303203-3032232012120322-0023123220113002-2332212100002011-2000233113201210-2311300102000003-2210030300212031-3032001303220123"></a>

## suffix_value property — rules / 223103021000 / 9

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [waf_skip_processing](resources--http_loadbalancer--reference--group-028.md#canonical-1202001010122130-1232112232211102-0102102132023332-3020023330320231-3120312230233312-0111230122132322-0202011203322330-2223113000302022): complete subsection reference.

<a id="canonical-0333012113333233-2000221230212022-0130013003231120-2113023130322232-2321301332202210-2210102212132002-1311100003033010-0101313201330020"></a>

## Next pages — rules / 223103021000 / 10

- [waf_exclusion.waf_exclusion_inline_rules.rules.any_domain](resources--http_loadbalancer--reference--group-027.md#canonical-1233330223321302-1130120332310310-0203100300123123-2130331311030110-0132230321210230-2132221103000313-1003202111301303-2033023223131322)
- [waf_exclusion.waf_exclusion_inline_rules.rules.any_path](resources--http_loadbalancer--reference--group-027.md#canonical-0310132112001312-0121211110302132-3301100330233002-0032202211303222-2023031010201303-0213032213310021-2210333223313222-3110023101111231)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--http_loadbalancer--reference--group-027.md#canonical-0203003232313033-0223230331033131-3221113013002333-3200303010013212-3202033211033320-3210102132113031-2230233013221023-0302302302002122)
- [waf_exclusion.waf_exclusion_inline_rules.rules.metadata](resources--http_loadbalancer--reference--group-028.md#canonical-3120302131000211-1331213131113022-3323221113300010-2300303303032223-3201231223203103-1003211220131202-2033111223021311-0032101121231310)
- [waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing](resources--http_loadbalancer--reference--group-028.md#canonical-1202001010122130-1232112232211102-0102102132023332-3020023330320231-3120312230233312-0111230122132322-0202011203322330-2223113000302022)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-027.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1233330223321302-1130120332310310-0203100300123123-2130331311030110-0132230321210230-2132221103000313-1003202111301303-2033023223131322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310003323200310-2103012210133030-1012332233230122-1330230131030231-2031021223331212-2021212030332301-0130330303013321-3212012023023133"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_domain — any_domain / 301031311220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-027.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-027.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-027.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_domain

<a id="canonical-0311122320020033-1323120031123203-3303311123100230-0133210123310132-3022332103101210-0023100321310201-3300233122201100-2030213233031300"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
any_domain = {}
```

<a id="canonical-3132122222310221-1020203000012112-2302001012113132-1332120030320231-3011100103301211-0321020221313322-3001000301122111-2101010201201013"></a>

## Direct properties — any_domain / 301031311220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313112220312133-3220112213231330-0103003203201210-1012312133211301-3300332211311112-1030002230310230-0123302102213001-2020023133023123"></a>

## Next pages — any_domain / 301031311220 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-027.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0310132112001312-0121211110302132-3301100330233002-0032202211303222-2023031010201303-0213032213310021-2210333223313222-3110023101111231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103320301222331-1321121300232121-1201231333313121-0212231301322031-1230131121313222-1302031210033230-1013233303212031-0120303212022333"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_path — any_path / 310300102013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-027.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-027.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-027.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_path

<a id="canonical-1311002003222122-2323213130230332-2312201320110000-2320003110301032-3211113120223112-1121221122200331-3202100322031112-2321220202231032"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
any_path = {}
```

<a id="canonical-3232332100313030-1111300223202032-2212220012233102-0313232222202312-3223220030310321-1220132313313213-3121102020121022-2120212322112113"></a>

## Direct properties — any_path / 310300102013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222022300132322-2312102001233221-3233223113302001-2302001001101232-0312220320222130-2322231010011221-0031300202111313-2321221023331023"></a>

## Next pages — any_path / 310300102013 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-027.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0203003232313033-0223230331033131-3221113013002333-3200303010013212-3202033211033320-3210102132113031-2230233013221023-0302302302002122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123202110103233-2311000013322210-3322213002120110-0213313211101033-1023013331201233-1300212102231031-0210220313013232-0321310220301003"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control — app_firewall_detection_control / 120021013233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-027.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-027.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-027.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="canonical-1003101013202000-2231202332301122-3310310210122320-0320123211212310-3021100022212103-0221232013311300-1123230222220001-0021112003001003"></a>

Type: `"object"`. single nested block, Optional.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313233112312020-0033301023231302-1122001322112300-0301132323333120-0330112013313131-2131213301332011-1213322233301210-1100020122112333"></a>

## Direct properties — app_firewall_detection_control / 120021013233 / 3

- [exclude_attack_type_contexts](resources--http_loadbalancer--reference--group-027.md#canonical-3201121111311030-2123122000311120-1220032201333203-1111001200333213-0231022113010030-2103210133201311-0312222131220013-3312320332301320): complete subsection reference.

- [exclude_bot_name_contexts](resources--http_loadbalancer--reference--group-028.md#canonical-0013000213203121-1302233212102023-3311321210110330-3303021000002211-3132312122212322-1032232001223211-2123112213300201-3010301331100133): complete subsection reference.

- [exclude_signature_contexts](resources--http_loadbalancer--reference--group-028.md#canonical-0333233003310312-0330131131103102-0000033020223132-0322311302000323-1330000332103012-0223131021220312-2000301102133300-1220301113220030): complete subsection reference.

- [exclude_violation_contexts](resources--http_loadbalancer--reference--group-028.md#canonical-2223300030211001-0132110221002313-3233013321223313-2012010332231223-3120133023113023-2011221200123212-1220022320210111-1133313122302030): complete subsection reference.

<a id="canonical-2333210312132321-2301023103212103-3113133111033033-0101001200310110-2130211110312230-0312023102120112-1221132110131003-2320202202311301"></a>

## Next pages — app_firewall_detection_control / 120021013233 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts](resources--http_loadbalancer--reference--group-027.md#canonical-3201121111311030-2123122000311120-1220032201333203-1111001200333213-0231022113010030-2103210133201311-0312222131220013-3312320332301320)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts](resources--http_loadbalancer--reference--group-028.md#canonical-0013000213203121-1302233212102023-3311321210110330-3303021000002211-3132312122212322-1032232001223211-2123112213300201-3010301331100133)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts](resources--http_loadbalancer--reference--group-028.md#canonical-0333233003310312-0330131131103102-0000033020223132-0322311302000323-1330000332103012-0223131021220312-2000301102133300-1220301113220030)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts](resources--http_loadbalancer--reference--group-028.md#canonical-2223300030211001-0132110221002313-3233013321223313-2012010332231223-3120133023113023-2011221200123212-1220022320210111-1133313122302030)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-027.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3201121111311030-2123122000311120-1220032201333203-1111001200333213-0231022113010030-2103210133201311-0312222131220013-3312320332301320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
