---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1213221123121332-0130210312020112-0030010203112202-1320321321320032-1222212130023330-2011211132030330-2231013002131030-0013313003020331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_service_policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- no_service_policies

<a id="canonical-1101022221112332-2030130211022031-1133103101333213-1101003230013311-3103302223310030-3213132312012201-0000201122030113-1223213310313002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no service policies.

Additional upstream details:

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
no_service_policies = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- origin_server_subset_rule_list

<a id="canonical-3011031302203312-1320300230222133-2210031101333001-3012120333010320-1021301101102203-3232213320210222-0233230312223211-3020200330331313"></a>

Type: `"object"`. single nested block, Optional.

Origin Server Subset Rule List Type. List of Origin Pools.

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
origin_server_subset_rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320132321221023-1012033003330010-1002032213320131-1311333132022102-3032103112322320-0123110222201011-3332201302333111-1202100021121321"></a>

### Direct properties for `origin_server_subset_rule_list`

- [origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210): complete subsection reference.

<a id="canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- origin_server_subset_rule_list.origin_server_subset_rules

<a id="canonical-1213300313322030-1003111033010333-3131211020312030-3221030020022012-3233021021232302-2230010311030323-2112103131133320-1010110322311213"></a>

Type: `"object"`. list nested block, Optional.

Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN,
Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin
Server Subset is a sequential engine where rules are evaluated one after the other. It's important
to define the correct order for Origin Server Subset to GET the intended result, rules are evaluated
from top to bottom in the list. When an Origin server subset rule is matched, then this selection
rule takes effect and no more rules are evaluated.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("origin_server_subsets_action"),
  validators.ConflictingListObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingListObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingListObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingListObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingListObjectAttributes("client_selector",
    "none"),
  validators.ConflictingListObjectAttributes("ip_matcher",
    "ip_prefix_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
origin_server_subset_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212302032011001-1323310322311331-1222133113021222-3102313012223113-2001220112021220-0001213121120103-3332021120300022-0333010133220331"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules`

- [any_asn](resources--http_loadbalancer--reference--group-022.md#canonical-1302111220302030-3001330021000231-3112112223302013-2220311032232231-2020101312020033-2310233211031132-0131223203311131-1222001331101321): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-022.md#canonical-2210123021022011-3220321020322313-1232200022032120-0002100200013131-1202331212220302-3210312031113231-0320130320222102-0120112330131132): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-022.md#canonical-3222323023321133-3012102203110231-3232333031113211-1203011322120130-2013302110020210-2230233122312101-0311121100023013-1331111002121103): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-2131232131220022-3231312032001311-1001123321021221-1010012233233102-3032020321321013-2003031121101100-2013021031132131-2032111122221213): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-022.md#canonical-3212132133222002-2131213233133333-0121323232210133-0023102330231222-3110220221030033-1321101302100320-2212002031310232-2021123112120111): complete subsection reference.

<a id="canonical-2211011233311303-1332330221112122-3021130201122320-0102131002101030-2333322100030033-3320121100111310-1210211221300030-2020103311213320"></a>

<a id="canonical-3020033113003231-3101231201031230-2010312113232221-2301001223130102-3333110120223203-2330330323032010-1130233002131231-2023122220123133"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.country_codes` property

Type: `["list", "string"]`. Optional.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Country Codes List. List of Country Codes. Possible values are \`COUNTRY\_NONE\`, \`COUNTRY\_AD\`,
\`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`, \`COUNTRY\_AI\`, \`COUNTRY\_AL\`,
\`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`, \`COUNTRY\_AQ\`, \`COUNTRY\_AR\`,
\`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`, \`COUNTRY\_AW\`, \`COUNTRY\_AX\`,
\`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`, \`COUNTRY\_BD\`, \`COUNTRY\_BE\`,
\`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`, \`COUNTRY\_BI\`, \`COUNTRY\_BJ\`,
\`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`, \`COUNTRY\_BO\`, \`COUNTRY\_BQ\`,
\`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`, \`COUNTRY\_BV\`, \`COUNTRY\_BW\`,
\`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`, \`COUNTRY\_CC\`, \`COUNTRY\_CD\`,
\`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`, \`COUNTRY\_CI\`, \`COUNTRY\_CK\`,
\`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`, \`COUNTRY\_CO\`, \`COUNTRY\_CR\`,
\`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`, \`COUNTRY\_CW\`, \`COUNTRY\_CX\`,
\`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`, \`COUNTRY\_DJ\`, \`COUNTRY\_DK\`,
\`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`, \`COUNTRY\_EC\`, \`COUNTRY\_EE\`,
\`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`, \`COUNTRY\_ES\`, \`COUNTRY\_ET\`,
\`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`, \`COUNTRY\_FM\`, \`COUNTRY\_FO\`,
\`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`, \`COUNTRY\_GD\`, \`COUNTRY\_GE\`,
\`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`, \`COUNTRY\_GI\`, \`COUNTRY\_GL\`,
\`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`, \`COUNTRY\_GQ\`, \`COUNTRY\_GR\`,
\`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`, \`COUNTRY\_GW\`, \`COUNTRY\_GY\`,
\`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`, \`COUNTRY\_HR\`, \`COUNTRY\_HT\`,
\`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`, \`COUNTRY\_IL\`, \`COUNTRY\_IM\`,
\`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`, \`COUNTRY\_IR\`, \`COUNTRY\_IS\`,
\`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`, \`COUNTRY\_JO\`, \`COUNTRY\_JP\`,
\`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`, \`COUNTRY\_KI\`, \`COUNTRY\_KM\`,
\`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`, \`COUNTRY\_KW\`, \`COUNTRY\_KY\`,
\`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`, \`COUNTRY\_LC\`, \`COUNTRY\_LI\`,
\`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`, \`COUNTRY\_LT\`, \`COUNTRY\_LU\`,
\`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`, \`COUNTRY\_MC\`, \`COUNTRY\_MD\`,
\`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`, \`COUNTRY\_MH\`, \`COUNTRY\_MK\`,
\`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`, \`COUNTRY\_MO\`, \`COUNTRY\_MP\`,
\`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`, \`COUNTRY\_MT\`, \`COUNTRY\_MU\`,
\`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`, \`COUNTRY\_MY\`, \`COUNTRY\_MZ\`,
\`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`, \`COUNTRY\_NF\`, \`COUNTRY\_NG\`,
\`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`, \`COUNTRY\_NP\`, \`COUNTRY\_NR\`,
\`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`, \`COUNTRY\_PA\`, \`COUNTRY\_PE\`,
\`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`, \`COUNTRY\_PK\`, \`COUNTRY\_PL\`,
\`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`, \`COUNTRY\_PS\`, \`COUNTRY\_PT\`,
\`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`, \`COUNTRY\_RE\`, \`COUNTRY\_RO\`,
\`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`, \`COUNTRY\_SA\`, \`COUNTRY\_SB\`,
\`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`, \`COUNTRY\_SG\`, \`COUNTRY\_SH\`,
\`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`, \`COUNTRY\_SL\`, \`COUNTRY\_SM\`,
\`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`, \`COUNTRY\_SS\`, \`COUNTRY\_ST\`,
\`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`, \`COUNTRY\_SZ\`, \`COUNTRY\_TC\`,
\`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`, \`COUNTRY\_TH\`, \`COUNTRY\_TJ\`,
\`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`, \`COUNTRY\_TN\`, \`COUNTRY\_TO\`,
\`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`, \`COUNTRY\_TW\`, \`COUNTRY\_TZ\`,
\`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`, \`COUNTRY\_US\`, \`COUNTRY\_UY\`,
\`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`, \`COUNTRY\_VE\`, \`COUNTRY\_VG\`,
\`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`, \`COUNTRY\_WF\`, \`COUNTRY\_WS\`,
\`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`, \`COUNTRY\_YT\`, \`COUNTRY\_ZA\`,
\`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ip_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-0203121202130031-2203220200231131-2001322323233203-0133212303033123-1123320021332113-2112022122112233-2331321110200131-3131031333310133): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-022.md#canonical-0111102320002302-0210101031332300-3113003023002203-3122211120031033-0021332113133230-0120303321101031-3133132212123213-3011103223302301): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-022.md#canonical-0020222103110101-2120130130101001-2303101303303210-2203023310303011-0101200212201213-1310123333103230-3133121312302332-2300113211122131): complete subsection reference.

- [none](resources--http_loadbalancer--reference--group-022.md#canonical-2211230323332300-3330211101102100-2023012322001003-1112202131233213-2321022211100031-0100321221002031-1333310000231113-0102002232313223): complete subsection reference.

<a id="canonical-1331010112333200-0320033011322030-2202133222113201-1112211202033001-0031031022113003-1213332232330312-0021111131300112-0332222212232031"></a>

<a id="canonical-2110112033203022-0032210333102222-1220130030122203-0322021201333101-2110122103101020-1010102200232000-1021101021103212-1223301201000101"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.origin_server_subsets_action` property

Type: `["map", "string"]`. Optional.

Add labels to select one or more origin servers. Note: The pre-requisite settings to be configured
in the origin pool are: &#8203;1. Add labels to origin servers &#8203;2. Enable subset load
balancing in the Origin Server Subsets section and configure keys in origin server subsets classes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3320220033301301-0000332222110111-3023100203013113-1111020331203121-3130020233100300-3022330130023113-1011312010210033-2000223331211331"></a>

<a id="canonical-2322200001110333-0213120121000021-2030232201323302-0230302321300201-0030102131203020-0100313213211023-1123300031132330-1102133123302213"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.re_name_list` property

Type: `["list", "string"]`. Optional.

RE Names. List of RE names for match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1302111220302030-3001330021000231-3112112223302013-2220311032232231-2020101312020033-2310233211031132-0131223203311131-1222001331101321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.any_asn` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.any_asn

<a id="canonical-3211310331033002-1022231030332230-3023111101120210-0030031322010331-3010013212100011-3211010123033013-0133020310031010-0013201323010112"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
any_asn = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210123021022011-3220321020322313-1232200022032120-0002100200013131-1202331212220302-3210312031113231-0320130320222102-0120112330131132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.any_ip

<a id="canonical-3113211132333302-1212020311323010-0001213121233100-2031202201112121-3212331323010321-0111222131031033-0330100332220202-1001020322331333"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
any_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222323023321133-3012102203110231-3232333031113211-1203011322120130-2013302110020210-2230233122312101-0311121100023013-1331111002121103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_list

<a id="canonical-3111011320213333-2033230120210312-2311021333320203-3223220222320300-0201202323113212-3132011231120223-2003133021303123-3030111103033211"></a>

Type: `"object"`. single nested block, Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202123123021111-2301332013032001-0222101302002103-2203303002021320-1322310030331003-0230012002002030-0232010332303121-3221303210110223"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.asn_list`

<a id="canonical-2332221012013123-3232312030103201-1011331003013103-1301103010320320-0121021021112303-3223320013322010-0322323111012212-2223130102123222"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2131232131220022-3231312032001311-1001123321021221-1010012233233102-3032020321321013-2003031121101100-2013021031132131-2032111122221213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher

<a id="canonical-1122010312123323-0220323031202020-2032021011320120-0132333012232321-3312223000122321-2112010100210331-0030122000210132-3232230012222232"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130122302102213-0213321300003013-1012221310012100-2321002033021133-3310003112323101-2332133223312330-3033323210110322-1022022211221333"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher`

- [asn_sets](resources--http_loadbalancer--reference--group-022.md#canonical-2133331202111302-1223330200330231-1303003211331110-0302111201302110-3220213222231011-1023012032011322-2311313300133130-1011321011233023): complete subsection reference.

<a id="canonical-2133331202111302-1223330200330231-1303003211331110-0302111201302110-3220213222231011-1023012032011322-2311313300133130-1011321011233023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-2131232131220022-3231312032001311-1001123321021221-1010012233233102-3032020321321013-2003031121101100-2013021031132131-2032111122221213)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets

<a id="canonical-0031332100011201-1210220132310232-0131021330110001-3112210331002003-2110200200121133-2333121300020131-0011003210032333-3131303231232310"></a>

Type: `"object"`. list nested block, Optional.

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031220202303023-2033011020213322-2120310112022011-3203013302111030-1311310201213110-1130033230132102-1022101031100013-2322201223112221"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets`

<a id="canonical-0021001123330013-2133102310013111-3321302321023010-3223012033332123-1321001321023313-2011221111302130-3313101222133111-1210302100310132"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2033201210130112-0122231311332221-1230323110032330-0003210331002110-2112131310311313-0133222233002200-3113201013323331-1331002300300121"></a>

<a id="canonical-3212033212101113-1212331003321133-2003233021203031-2220211230301222-2011020200031210-0320110023013203-3131221023111213-2123233131323030"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3121133031202300-0033331000331130-0200213201010013-2110112300102023-2013320223333332-2110012103300120-2202132011011312-1110031233220012"></a>

<a id="canonical-1200330020301022-1030332220023203-2002113323121031-2130111113102012-3110031212311301-1220303202330033-1002031331310303-0313301020313312"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
  }
}
```

<a id="canonical-0201113112311211-2100011012111222-2103301000222023-0312103312002212-2332122301103002-0100032002222120-2022013331112103-1211333300321232"></a>

<a id="canonical-2001002031123211-3002112101102003-3323130321223330-1111100010023322-0210301320223132-0333313321322312-0303020102301012-0232103312210003"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0121033002101020-0020111221211331-2013123232013003-0022002233201230-0212023133110020-2332232303321302-0332012103202200-3030131021031312"></a>

<a id="canonical-2102030033012332-2202132233101310-3221311213220000-0321032211232123-1303220313231322-2000121101032300-3201010223120311-3201111033203321"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3212132133222002-2131213233133333-0121323232210133-0023102330231222-3110220221030033-1321101302100320-2212002031310232-2021123112120111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.client_selector

<a id="canonical-3312322003112231-3332221132212021-1100221223100220-2231012323101203-0323101112003001-1111012112013203-0021300202210111-1222230212000221"></a>

Type: `"object"`. single nested block, Optional.

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133000201210300-1013302030222023-0201022332230011-2020302021210133-2001301113011123-1010220110133220-3002300322123000-1330100311122100"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.client_selector`

<a id="canonical-0003103120121013-1203102003121331-0120102031013121-1323103113110222-0323231123130010-3310302120023100-1012323123123331-0210002302212312"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.client_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0203121202130031-2203220200231131-2001322323233203-0133212303033123-1123320021332113-2112022122112233-2331321110200131-3131031333310133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher

<a id="canonical-0111302330233122-0031221011331101-0331003300113023-2003100302003102-2210130202223001-3022130203330232-0002332313103022-0230312231331303"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212101232021120-2333122132132322-1012103212230002-1112222203200333-2220000130231313-0220023020203011-3111212102032030-0210321233212122"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher`

<a id="canonical-1331310031213312-3130221021100003-1202200000313121-3021123311313023-2000333103312112-1220000211121232-0031013023112100-0312222033221103"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.invert_matcher` property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](resources--http_loadbalancer--reference--group-022.md#canonical-1002210121103330-0232111131223301-2202012213202122-1212001301000302-3312010302120311-1322203332231203-2032221300121331-2332312222132221): complete subsection reference.

<a id="canonical-1002210121103330-0232111131223301-2202012213202122-1212001301000302-3312010302120311-1322203332231203-2032221300121331-2332312222132221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-0203121202130031-2203220200231131-2001322323233203-0133212303033123-1123320021332113-2112022122112233-2331321110200131-3131031333310133)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets

<a id="canonical-3000103012213031-0000013032122121-1030202112101223-1012030102320210-2001232230210223-1202232203101330-1002031200320301-3200033103310302"></a>

Type: `"object"`. list nested block, Optional.

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2031321222231000-1301212131131132-3220232302011113-2323101221220302-3113223233132220-1312023132230233-3213332320103230-3323000101111131"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets`

<a id="canonical-0001210321320320-0231220212332121-2232222103003131-3033330302331111-2033010200230331-3100233031220002-1300123103313120-2001220201122031"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1130321112020323-3212312011233033-2300102232001311-3100333131313000-2332211022311311-3332031102101113-2320011102031130-0023022231022302"></a>

<a id="canonical-0300220331100330-1222333332202302-0100031003112023-3030312210333112-1331321102233230-2331110332211313-1312011020010303-0132323112011011"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1213331223320121-3302133002020132-1102333100233211-0030131000003011-2321133212120210-1010101333222000-1203323023021103-1211230011331033"></a>

<a id="canonical-2200311223221003-1301331131133212-0030132112222211-2201013101312210-0020333021003112-3131021233121112-2032011230213011-3303330322021203"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
  }
}
```

<a id="canonical-1322201323013313-2213320300322300-2301100202123102-1123301302200220-0133222133213321-0111302300112302-2233001320013322-3303310112003201"></a>

<a id="canonical-2033320323220033-1203230203233331-0030110221121212-3331220311131223-1133100213132332-0201302030231301-1132021313232003-0320132220000102"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2110100310223230-0121220011233012-0131303012231222-2002101302333201-2220230123132303-3330232113033302-2103122303321111-0212112112220201"></a>

<a id="canonical-1303132233322000-1332102311211320-1201103332300003-3231000023212103-3132002301000130-3210020221001220-3020222102310030-0033131231231313"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0111102320002302-0210101031332300-3113003023002203-3122211120031033-0021332113133230-0120303321101031-3133132212123213-3011103223302301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list

<a id="canonical-2322031131330030-2011201131110022-2233311200212320-3001202323231202-2322101130330202-0220211332202301-3022233001312301-0303223110213111"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003322003231333-2003332010031331-0011330033120322-3101111031222332-3232013002103011-3321332221022120-1221311230123102-3201231201232310"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list`

<a id="canonical-0311010311111232-1012333130130112-3301120111000321-1323032322312202-3002030333220130-2202231232032023-2012312132010021-3210333133302031"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list.invert_match` property

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

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

<a id="canonical-0101230020031101-0300001020321222-3212123100310312-1110100333310302-0112223000231303-3321223320020000-2320111010203101-1332022313012202"></a>

<a id="canonical-0003223122232131-3022223233113113-2231232331002211-3011111000013212-0032101310032210-3202213310012203-0031033031002231-3222102230113203"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0020222103110101-2120130130101001-2303101303303210-2203023310303011-0101200212201213-1310123333103230-3133121312302332-2300113211122131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.metadata

<a id="canonical-3310232202230200-1301223221221232-0032130002323321-0121200212302101-1202300333111131-2222111323201213-2222310030023011-1200300112132133"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1210222033220032-3201213012222011-0132020311302300-1113311132222122-1232023102300200-3122002322322330-0332103122231012-0031022011031221"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.metadata`

<a id="canonical-3033123210121212-2320010113303230-0131103010322330-3131223313133121-1031211232133321-1133310022230001-3332101233203133-0303220022223031"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3031123200211102-0130302002033110-3201210122013103-0313120321201310-2323000220110202-0201133030320232-1000302302300113-2113103202030033"></a>

<a id="canonical-2110010212122033-3120111300031133-3133023222031333-1220003323323021-3100330213101322-2311310112210011-3332013100230321-2231222021212310"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2211230323332300-3330211101102100-2023012322001003-1112202131233213-2321022211100031-0100321221002031-1333310000231113-0102002232313223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.none

<a id="canonical-0230220011112222-0331131123310103-1111230301221011-0010010003111023-1333101100011200-3311211023012032-3303310310320330-3002322011200331"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- policy_based_challenge

<a id="canonical-0321223103102130-0121132313031223-1220210002011320-2133231330312121-2131132110220002-1313302332130023-1110033030130013-0300230120302232"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings for policy rule based challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "always_enable_js_challenge"),
  validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("always_enable_js_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("captcha_challenge_parameters",
    "default_captcha_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_js_challenge_parameters",
    "js_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_mitigation_settings",
    "malicious_user_mitigation"),
  validators.ConflictingObjectAttributes("default_temporary_blocking_parameters",
    "temporary_user_blocking")}
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
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-challenge_choice": "[\"always_enable_captcha_challenge\",\"always_enable_js_challenge\",\"no_challenge\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]",
  "x-ves-oneof-field-temporary_blocking_parameters_choice": "[\"default_temporary_blocking_parameters\",\"temporary_user_blocking\"]"
}
```

Terraform syntax:

```terraform
policy_based_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113100030000322-1000020320313012-3123001301230221-2310022300233321-2003320223202212-3211200100311000-1021110011223310-3111111112301232"></a>

### Direct properties for `policy_based_challenge`

- [always_enable_captcha_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1020103212230210-0302013301230200-1330100003322012-0303133122222112-2120233032330102-2023332131210232-2001103103013232-0223230323101222): complete subsection reference.

- [always_enable_js_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-2020101203121101-3122131030121332-0001233010301221-2210011313322031-2230000010103111-2322322320133211-3211101333203323-0113330131100123): complete subsection reference.

- [captcha_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-0303231223000210-3231100031100331-2121330111002123-2223112222233123-0023323312332202-1030113002130111-1203331300230133-2320320211010312): complete subsection reference.

- [default_captcha_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-3100203011131011-3211013133233133-2213000122113301-1021032002123120-2312002312023020-3010133101311113-1122232312121323-1030100121231011): complete subsection reference.

- [default_js_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-0103303331213011-2202323031323123-0102330313313310-0232020323231013-3323001032002010-3302121300202230-1123012033300110-2102313203013002): complete subsection reference.

- [default_mitigation_settings](resources--http_loadbalancer--reference--group-022.md#canonical-2032331232010130-3032232001332021-3210130210012020-0000320322331212-0210133323311100-3200011130023303-2331221023110031-0222332120002102): complete subsection reference.

- [default_temporary_blocking_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-2021133201102323-2012112211133233-2020010222002310-1013203132233103-3332132002103023-0011012112312331-3021312321312110-3300200102202133): complete subsection reference.

- [js_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-2231013300332210-2031112112313030-2221032321102312-0132223230313021-1321303100322023-2033103133323020-3122021120230210-1013101123002233): complete subsection reference.

- [malicious_user_mitigation](resources--http_loadbalancer--reference--group-022.md#canonical-1322110130002031-0223022033013121-0332110131013120-0000132123321212-2123232030112001-0231012210300222-3202331330011102-2112222303312322): complete subsection reference.

- [no_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-2232210000001111-3302330222322313-1131322202232312-3100230033303233-2301233011220321-1313012120021222-3122202003333100-2121101032211221): complete subsection reference.

- [rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033): complete subsection reference.

- [temporary_user_blocking](resources--http_loadbalancer--reference--group-024.md#canonical-0332322110110101-3123223013302023-0321200302320201-2312111021001220-0032301331120201-2310003212230313-2113130213200312-2022222232311001): complete subsection reference.

<a id="canonical-1020103212230210-0302013301230200-1330100003322012-0303133122222112-2120233032330102-2023332131210232-2001103103013232-0223230323101222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.always_enable_captcha_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.always_enable_captcha_challenge

<a id="canonical-3202301202020322-3102312023130303-3211312103313203-0003110221223113-3021101100231301-3323131103010221-3103300012013030-0032203212110233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable captcha challenge.

Additional upstream details:

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
always_enable_captcha_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020101203121101-3122131030121332-0001233010301221-2210011313322031-2230000010103111-2322322320133211-3211101333203323-0113330131100123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.always_enable_js_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.always_enable_js_challenge

<a id="canonical-3202221220113021-3213310202320012-0011232033203303-2203031310130130-2303032300301001-2013003312032103-0101103220320320-0013020231122131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable js challenge.

Additional upstream details:

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
always_enable_js_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303231223000210-3231100031100331-2121330111002123-2223112222233123-0023323312332202-1030113002130111-1203331300230133-2320320211010312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-3021003323123112-0111312010132002-3002120230301131-1122111133331303-0301133202010330-3101121321210232-3232020101213213-2330022321033203"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
captcha_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210022332010012-1223231322022233-2330302312023110-2310213222031321-0000011131303313-0232303022221322-3321010000212012-1231133320322120"></a>

### Direct properties for `policy_based_challenge.captcha_challenge_parameters`

<a id="canonical-3030300121113111-0121113012013123-1021213032033200-3120011212311213-0130223233100211-2311120132130100-2222233031011221-0300120310102031"></a>

#### `policy_based_challenge.captcha_challenge_parameters.cookie_expiry` property

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-2332013001311022-1112323121222103-1333223000230032-3111223021300203-0201200203302311-0201112302220120-2202033222033112-1303313212030321"></a>

<a id="canonical-1122230201301320-1120000131311031-3113322321131032-3321003113023310-2212300311020101-2023203033030210-3122201211102333-3131103111233202"></a>

#### `policy_based_challenge.captcha_challenge_parameters.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3100203011131011-3211013133233133-2213000122113301-1021032002123120-2312002312023020-3010133101311113-1122232312121323-1030100121231011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-1022223320333010-1323310210203121-1333211223311011-2132023113222113-1132301212133331-0023200020201331-3222211223310021-2231321033110030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default captcha challenge parameters.

Additional upstream details:

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
default_captcha_challenge_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103303331213011-2202323031323123-0102330313313310-0232020323231013-3323001032002010-3302121300202230-1123012033300110-2102313203013002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-3323223032321022-0210001331213103-3103333023100313-0021213333232301-1021022111201313-1111012011221013-2332303231133232-2221203213201033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default js challenge parameters.

Additional upstream details:

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
default_js_challenge_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032331232010130-3032232001332021-3210130210012020-0000320322331212-0210133323311100-3200011130023303-2331221023110031-0222332120002102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_mitigation_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-0200111120313200-0302001130002103-2230233333331321-0002223101111020-0013010133212303-0020320313223030-2112220331212101-1033233330122203"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
default_mitigation_settings = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021133201102323-2012112211133233-2020010222002310-1013203132233103-3332132002103023-0011012112312331-3021312321312110-3300200102202133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_temporary_blocking_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-0312312320020110-2333130330210331-0330023302013312-0300000100311213-1311310300110010-1321200303103113-0230001112203122-2213333012312202"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
default_temporary_blocking_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231013300332210-2031112112313030-2221032321102312-0132223230313021-1321303100322023-2033103133323020-3122021120230210-1013101123002233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-1302101310203202-1200232210223303-0013013232101202-0333332022300033-2233030321023030-3023001201302102-3322011222000021-3021331211030030"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
js_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200300132320222-0303213032030020-2002222303301232-1303001321113231-1030330031130033-3111223123220313-2003112323033112-0222222120103231"></a>

### Direct properties for `policy_based_challenge.js_challenge_parameters`

<a id="canonical-2111211113300310-2120133101120122-3121213202210021-2322132020103301-3211122222300013-3101132130331311-3130221212122021-3033012021203122"></a>

#### `policy_based_challenge.js_challenge_parameters.cookie_expiry` property

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-3203122313231121-1320222320123202-1202230200030233-3130120123110010-3112031331321311-2020222301002012-1103000020233310-2232320021201023"></a>

<a id="canonical-2020021021201333-0121112020300301-3320031103112310-0003111300132123-1220322300001111-1323313012230231-1012013031232331-2111212320231302"></a>

#### `policy_based_challenge.js_challenge_parameters.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3020023123302223-3003323111203300-2132133020030023-1121212222322313-0320201021233132-0013030000112001-3320212300332002-2213213231222212"></a>

<a id="canonical-3223030110100022-3112031201023230-0203002010220022-1221203231010221-2021212221301011-3231121333303220-0330223120202012-3331132231010101"></a>

#### `policy_based_challenge.js_challenge_parameters.js_script_delay` property

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-1322110130002031-0223022033013121-0332110131013120-0000132123321212-2123232030112001-0231012210300222-3202331330011102-2112222303312322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.malicious_user_mitigation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-0110113330020201-1122230200121211-2230213021322131-2323211120330222-3000331303300013-2213000031021030-0122123023101210-2100132323321312"></a>

Type: `"object"`. single nested block, Optional.

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
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323220220321130-1321222022231302-2211022320233200-3030202223220112-2313320310232133-3230210001232130-1303101011213032-2021213221120012"></a>

### Direct properties for `policy_based_challenge.malicious_user_mitigation`

<a id="canonical-2030120120023123-1233011320121100-1021313030013320-1123212303113020-1222323131211132-1002031102301200-0230000200212233-2111003011101323"></a>

#### `policy_based_challenge.malicious_user_mitigation.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2203202101203100-2101010302211213-3213302231331312-1323311202012111-1212100132312101-0212213233031230-0303202032013011-1201102011112002"></a>

<a id="canonical-1212230111330312-3021223133031101-1302301023133232-0310020120310213-1323231320110300-3301131002021200-0302013102033231-1012111331233002"></a>

#### `policy_based_challenge.malicious_user_mitigation.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1222022003331221-3120132031233233-2032202013321323-0310202032132013-1313312000330012-0313032223231012-0311122122031331-1123131221231101"></a>

<a id="canonical-3012320221330113-1311130220020232-0323133313101101-1102133311332220-3303110232132203-0133023022121123-2010203032303230-3230033023123133"></a>

#### `policy_based_challenge.malicious_user_mitigation.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2232210000001111-3302330222322313-1131322202232312-3100230033303233-2301233011220321-1313012120021222-3122202003333100-2121101032211221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.no_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.no_challenge

<a id="canonical-3322020102033022-2300220312302330-0131122233313010-0330000030030122-2032122320103010-2201203300001313-0212133230003321-3302111031222110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no challenge.

Additional upstream details:

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
no_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.rule_list

<a id="canonical-1021200113120231-1032213101120203-2022312102303122-3132200101003110-2332320123033011-2202000211133003-0320221133100202-3200333001230010"></a>

Type: `"object"`. single nested block, Optional.

List of challenge rules to be used in policy based challenge.

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
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231013011211332-0212210130201212-2203010011313103-2322103100220231-2310221220231230-1233201330332231-1233332032322333-0021101132100231"></a>

### Direct properties for `policy_based_challenge.rule_list`

- [rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030): complete subsection reference.

<a id="canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- policy_based_challenge.rule_list.rules

<a id="canonical-1021200210032232-3303203123232232-3103201232322213-3200130123021131-2311200233232120-3010200302131222-0300013310133222-2202033022332232"></a>

Type: `"object"`. list nested block, Optional.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
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

<a id="canonical-2102301313310312-0100211212030221-2112133012321300-2201110031212201-2101133022130033-3322131211010222-0202233112313320-1321002010310102"></a>

### Direct properties for `policy_based_challenge.rule_list.rules`

- [metadata](resources--http_loadbalancer--reference--group-022.md#canonical-0302122012200302-0210113312111030-0212231000333011-2200033300033333-2220311000021123-2332322330021310-0013231001133031-0020320013212302): complete subsection reference.

- [spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212): complete subsection reference.

<a id="canonical-0302122012200302-0210113312111030-0212231000333011-2200033300033333-2220311000021123-2332322330021310-0013231001133031-0020320013212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- policy_based_challenge.rule_list.rules.metadata

<a id="canonical-3131100303011033-2200033232132011-3031231230200102-2322030133132213-3013301131031102-3231021120332121-0032032322222000-2303000233220113"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1023210210010310-1002032020010001-1310312302310111-0110231003011311-0312100111001130-2331030321331033-3213132332102212-3001311003212111"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.metadata`

<a id="canonical-1320232312110130-2102312300220031-1311103222322030-1233121000231001-1013202331212211-2312112201221212-1013322221222332-1003133210031102"></a>

#### `policy_based_challenge.rule_list.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3332332130010230-1123002320221000-3232120212013032-3031300131300120-2023213100010210-0202111223011012-3320023221210312-2210010311300233"></a>

<a id="canonical-3213323013100323-2102113102111320-0122030203122113-2321230122120130-3001032112312131-2113213221302312-2201313133001222-0023011012002330"></a>

#### `policy_based_challenge.rule_list.rules.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-2022303111012211-1220213132123231-2022321212300020-1131132033230111-1321111301000001-2331030223131302-3011023300232321-0221120102012232"></a>

Type: `"object"`. single nested block, Optional.

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_captcha_challenge"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_javascript_challenge"),
  validators.ConflictingObjectAttributes("enable_captcha_challenge",
    "enable_javascript_challenge"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121113212202332-3022111213032320-3000332021000223-0022001032121331-2330200023312130-1301221132320311-0230130203210210-1131321111230210"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec`

- [any_asn](resources--http_loadbalancer--reference--group-022.md#canonical-0210300211122310-2122130331131112-2202100211131200-2203112212032321-2320330111111003-1022233132221233-0131330231132000-2210223331000133): complete subsection reference.

- [any_client](resources--http_loadbalancer--reference--group-022.md#canonical-3001013311103320-0133221211231322-0231122010331231-2020231333201210-0112210302022123-3232230023220113-1021012221130132-0031320122021332): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-022.md#canonical-3002222112023221-1013111330033200-2033000003103010-1313132112333032-3103213131123020-0130332002312010-2013133022111211-0133020111203033): complete subsection reference.

- [arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-1101313011001300-1031201312333332-2301022020220121-1201202331311323-3022013032203330-1000001113112111-2300012330010303-3001011233020023): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-023.md#canonical-0003301213131101-1312033131333213-2101303001113021-1023212221120031-1331213111320113-3131102321210212-0323000011233200-3231112332221032): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-2330232220313013-2120000211302223-3012212202222313-1113232301013021-3231202321121233-0332211110312023-1013001211321310-2012322132110120): complete subsection reference.

- [body_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-1322031002303310-1111010110113010-3010000120211223-0112031202323033-3213001011000120-0032032230230101-3220223203111233-0023030020301031): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-023.md#canonical-3301001033202231-2133133323213110-2321233202103301-1310201321100023-1302201121201002-3200232221023023-0110031303001223-0212313222223133): complete subsection reference.

- [cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030): complete subsection reference.

- [disable_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-0202302323030031-0332332132022130-1301100023003233-2302201010011300-1300320313311232-3332303233323013-2021332220011213-0032313121021322): complete subsection reference.

- [domain_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-2231322321300130-2211003200213110-0230120220232310-0330031202332122-1202330333221203-2022232030123301-0132023201301310-1122331113003030): complete subsection reference.

- [enable_captcha_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-2331103200120201-0301001021002322-0132233020221032-2203102002320022-3001110313102332-0311012113002330-0202002031220211-1023110131120121): complete subsection reference.

- [enable_javascript_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-2221112113201102-3022103132211101-1320220303332222-2303330112030301-1020203320020311-2100330302222210-1012001232303023-3130312322023023): complete subsection reference.

<a id="canonical-3030123003333323-3321200033223023-0321023003322300-2310122221113131-3223202212223313-0122123320110331-1332103101001331-3212311310012200"></a>

<a id="canonical-2110030122020113-1111323220212023-0322032000231323-2000123212032103-3012122201031231-0211123033103312-0321011300002331-2032222322032230"></a>

#### `policy_based_challenge.rule_list.rules.spec.expiration_timestamp` property

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111): complete subsection reference.

- [http_method](resources--http_loadbalancer--reference--group-023.md#canonical-2330011102111203-3200210100131120-2112200113232323-3120021212132232-3203312330132100-2321330223303101-1023321200013301-2303010303011323): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-3131232312332313-2331122113110111-0131010211100102-0310213132121112-3111200310110021-1010121212033013-3232311000121330-1220231122112110): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-023.md#canonical-1102303013213221-2101121232233030-3032103211220121-2031323223113213-2210001011111012-1001212110123010-0120200232130210-3120021232133212): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-023.md#canonical-2110130302132222-1333323202323202-2101132211132333-3120123331020002-2321313212221101-0010000002320121-1033331331113211-1113310312332201): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-024.md#canonical-1131033020202020-3130333330012030-3002001123121030-1323223321202032-1221131132031001-3323021113022011-1203131303033130-3010031021310313): complete subsection reference.

<a id="canonical-0210300211122310-2122130331131112-2202100211131200-2203112212032321-2320330111111003-1022233132221233-0131330231132000-2210223331000133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_asn` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-2013222223203003-1010002210001202-0221101112111322-3011022010201213-3203222223030333-2212232131313102-3301211023323030-2220330030202203"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
any_asn = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001013311103320-0133221211231322-0231122010331231-2020231333201210-0112210302022123-3232230023220113-1021012221130132-0031320122021332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-3000222333000233-0223102010013120-1301003330013123-1010033303310201-3212103301133331-0200332233032123-0200333010322013-0201120201211232"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
any_client = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002222112023221-1013111330033200-2033000003103010-1313132112333032-3103213131123020-0130332002312010-2013133022111211-0133020111203033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-0123320101312210-3323010200330110-3120331031000321-1101301302221312-1102131310230303-3002031220301321-0321113031203132-3311110011202033"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
any_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101313011001300-1031201312333332-2301022020220121-1201202331311323-3022013032203330-1000001113112111-2300012330010303-3001011233020023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-1001313021020303-2233232132332232-2201323103221012-3133111210202121-3020221223111300-3033323232101202-0203203312332002-2121003213212002"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
arg_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013023210232130-1232121023233100-3102111131110020-0303131122003332-1230030110323230-0320102211132321-3322111123023201-0110311203122030"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.arg_matchers`

- [check_not_present](resources--http_loadbalancer--reference--group-023.md#canonical-2303020231320132-3322122201120122-3210131111212200-3120303012212232-2322233112010021-1311130030331303-0012320213231121-1203220020213211): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-023.md#canonical-1131222013031202-3211232223001312-0013120330133031-1110232213332332-2102200023123203-2111103310301333-2323122300323123-0020231033321300): complete subsection reference.

<a id="canonical-1130221332110311-2333233103121133-1233013311131210-1212123111203221-2233022131100120-2032031021023331-3013133121220100-0101121100320032"></a>

<a id="canonical-1132023312223130-1212131310031303-3000222200310012-3323321023100322-2023312031133232-3222123210210300-3003331200221003-1103020300030203"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

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

- [item](resources--http_loadbalancer--reference--group-023.md#canonical-1103010012000013-2110202002230020-2032130233221320-1233032213333113-3202021112010211-0033113003100012-2012222100210311-3211210200221222): complete subsection reference.

<a id="canonical-2231310202311021-0222332222232011-2023002233000023-1211213110302201-1313330001200102-0003230032011032-0303031012110313-2203002020020020"></a>

<a id="canonical-1033321322003302-1232011331032211-1221300323332331-2012231331331011-0233320012203310-1302201101223311-3100123313300232-3302130123030332"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.name` property

Type: `"string"`. Optional.

A case-sensitive JSON path in the HTTP request body.

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
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```
