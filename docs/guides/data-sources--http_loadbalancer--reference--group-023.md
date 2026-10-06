---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0130320232310300-1003033303212310-1030202000201111-3022320032203012-1110200010232110-0320022322130201-0102110002122103-1202210113103203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge

<a id="canonical-0020100031330330-1332120301020103-2000332231021221-1133231322203020-3300103301201321-2320300121001303-2230113030101202-1212333110213021"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="canonical-1311012110101103-2022102003212133-2331122223022030-3031333303023030-2000311212003233-0321221133123001-0122233030012310-2112130211211113"></a>

Type: `"list"`. Computed.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1302023203010330-1202031123121312-2201103210320021-0120321331322000-3022320200211300-3232011122112020-1220001031111232-2310100230113012"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.headers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-023.md#canonical-3323311011222320-2030013323322120-3010020022223023-1033012230302202-2123232010313312-2303321011002130-3033302032033200-1001311122032221): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-023.md#canonical-2210223333011301-2122103132330213-2100323122032322-0313113112102202-1213302023032001-2133100231213332-2202203102130320-1301120032202120): complete subsection reference.

<a id="canonical-2100333321121233-2223111320233021-3020201312131233-2300022010302030-2221330113223011-1030312013033133-3220102333311323-0301221101221020"></a>

<a id="canonical-1330222102010120-1012200331013002-0321020333211210-0200013213100020-0322312222031013-2200102123230223-2212002013320010-2012221110221333"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-023.md#canonical-0012330233033310-1033320213101320-3310023313021031-0020110031211302-1210131020003323-2203002032013100-0023220123332332-3121103332101233): complete subsection reference.

<a id="canonical-0123233031211021-1303113320011121-0303331133310131-2313030023003301-3231001322013003-2111300323131031-0300033333202023-0030321330012023"></a>

<a id="canonical-0331301012232201-0301120103201331-3113201112033312-0303123013100100-3003233302113313-0312222010121130-3210232331131101-2320333211332222"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3323311011222320-2030013323322120-3010020022223023-1033012230302202-2123232010313312-2303321011002130-3033302032033200-1001311122032221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-023.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101)
- policy_based_challenge.rule_list.rules.spec.headers.check_not_present

<a id="canonical-2100321020233210-0111101310220031-0220302021231012-0330003213130013-3020231020111122-0211112323213222-0222221131002101-2201313111232231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210223333011301-2122103132330213-2100323122032322-0313113112102202-1213302023032001-2133100231213332-2202203102130320-1301120032202120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-023.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101)
- policy_based_challenge.rule_list.rules.spec.headers.check_present

<a id="canonical-3120112032333021-2103301331013303-0032133231220022-0133103202232031-1313301031323221-3011112212301211-2020221333313310-3132131013101322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012330233033310-1033320213101320-3310023313021031-0020110031211302-1210131020003323-2203002032013100-0023220123332332-3121103332101233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--reference--group-023.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101)
- policy_based_challenge.rule_list.rules.spec.headers.item

<a id="canonical-2311111320212101-3103013302120022-3211121113212023-2311102222330121-2212002030102023-3000233012112112-2103032011022110-1122332202020121"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-0200013033230033-1212332203213332-2021003302123202-1212233202212213-0231033130330100-2102322013202000-3103230320100233-1300033001330000"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.headers.item`

<a id="canonical-2312300113013031-2021333123103211-0003011013102210-2210322011000002-1332023210203220-2023211302221333-3223322033122000-1023032311202222"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0210032212131333-1302231330220321-0323213001101113-0331000200001012-0203131102020033-0202103123203203-1200222211100333-0323103103201002"></a>

<a id="canonical-3303321021311030-0303033231012302-0303212233011321-2201120213132321-1012110323211131-0011331102310010-0122203121110111-3012011013201030"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2033021301212310-3133200203222313-0030310101113123-3020012311102211-2331033030031323-0131322313000102-2002300130132333-1321021333213121"></a>

<a id="canonical-1232013113133201-2233032022203121-1033223032332332-1201230232130121-0300011322232231-0020130230311020-3001223103112200-3212330300000102"></a>

#### `policy_based_challenge.rule_list.rules.spec.headers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3020310301331122-1212000113011100-1023213212330331-3122103113210321-2002122203320320-2120121220311201-1333220030000032-3120130221023203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.http_method` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.http_method

<a id="canonical-2232002321022103-2312101333311212-1330302022001232-3032110202111131-3001113301031203-0202212102210301-1230232231221000-3123110313301303"></a>

Type: `"single"`. Computed.

An HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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

<a id="canonical-1230200031013233-3232122210303320-0113112233111300-2223220113332030-1022100320000231-0133322322320330-1010021132323303-2011210031020003"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.http_method`

<a id="canonical-0321020010113201-1320101301030102-1332310122010213-1223223013110221-3313330202011213-3230020111130301-1203020122200232-0203321013303320"></a>

#### `policy_based_challenge.rule_list.rules.spec.http_method.invert_matcher` property

Type: `"bool"`. Computed.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-2313113103303130-2011212201302310-2102301300200023-0332221031032021-1330210033112213-1123300313311021-1312022103032313-2210212013302012"></a>

<a id="canonical-2302311133301000-3332021132102003-0322230202003303-0130201321120233-0003132330313003-0230121103001030-0011113330120320-3301101122101212"></a>

#### `policy_based_challenge.rule_list.rules.spec.http_method.methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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

<a id="canonical-3022123003013200-3203320321330332-0121233201103200-2331212133333301-0201213111233230-2211313020321111-1111213132220212-0330232000202000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="canonical-1100121103030130-0022012303133233-1002302123220211-3023000232011010-1333012003102303-1022102312100313-3121233300012311-2103123311313012"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="canonical-1022323020012211-2122333311313001-3200021300322213-2022321013032113-0102333311312013-1201231123022030-0222031133222031-0020000100312230"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_matcher`

<a id="canonical-3210203211033230-2030003223202013-3013113013123012-3012011321130211-1100301312223031-1210032032313133-0331303202133302-2121300210332313"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.invert_matcher` property

Type: `"bool"`. Computed.

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-023.md#canonical-3321330223011200-2113221111001023-1010132303333213-0213102211233022-0213123033123312-0102321111200103-3310223201232203-2011311300020120): complete subsection reference.

<a id="canonical-3321330223011200-2113221111001023-1010132303333213-0213102211233022-0213123033123312-0102321111200103-3310223201232203-2011311300020120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--http_loadbalancer--reference--group-023.md#canonical-3022123003013200-3203320321330332-0121233201103200-2331212133333301-0201213111233230-2211313020321111-1111213132220212-0330232000202000)
- policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-3323203121223011-0131011032032203-0122000333013321-2000223133312001-1032033231020330-2223033203323330-1210130032022322-2231130110322100"></a>

Type: `"list"`. Computed.

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

<a id="canonical-0120323211022302-2001220203312232-3222031303121100-1312320300313300-0202131231310123-3003031233110000-3122011223102320-0122331213333203"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets`

<a id="canonical-3003213221022123-0132320132231202-0121312003011011-2312021200100323-0212113112213023-2013133321032210-0221310211110000-2003023113001230"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.kind` property

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

<a id="canonical-3312331133201211-3322303113210320-0023111031132331-2123110312020230-2223220120322212-1130303103030020-0330120321003002-3121002100301002"></a>

<a id="canonical-1321320000121002-3023322313310031-1120120321111010-2303313223022333-2230331203231022-0330312031323131-1230012110012231-1323120310023332"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.name` property

Type: `"string"`. Computed.

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

<a id="canonical-0222023311032201-2111013111303203-0001322000203112-1001203030220123-3013222130111301-1220203002300230-2010023000103131-2310101320112202"></a>

<a id="canonical-2011211201001130-0121012030123003-1230200000330320-2313222102110311-0132232112210231-3003320120201302-2013221223020010-2321121203332133"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0030312002230022-2003302300112323-2220313220301121-1002120030003300-3333111000301310-2303233320132220-3302320213102221-2322223313212121"></a>

<a id="canonical-3130031223211331-0212322133133211-2233313112023113-3311130133131323-1231300310230323-3311003010020233-2011231033020131-1003131022021001"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.tenant` property

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

<a id="canonical-2021033211331230-2301103233302111-3221112321332320-1333232330112120-2303132100202301-3301311323331300-1222110032003101-2231031222003113"></a>

<a id="canonical-2130012232321031-2100002033021113-2323303201231311-1331003020320322-0102301100310103-3033210012203120-1012012022002110-1303013103013112"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets.uid` property

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

<a id="canonical-2102212001013000-1331112303230010-2332311101103003-1203312030331231-1110322311233210-0133212103001332-1212233031202211-1222020223230121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="canonical-2100320012213103-1022210223101301-3120202322033231-0022120223112323-1011212000002002-1131001002312002-1303211202323123-1123312212201102"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0223230221100311-3220003201212231-0201001220031302-0123200133002211-0102131231120313-1122120001230012-3121023132220312-1033233311310111"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.ip_prefix_list`

<a id="canonical-2003321100321130-3120202030220211-1232100012030231-3132201323300032-3230010110033300-0211233130020102-2103003133032103-1210132213213310"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_prefix_list.invert_match` property

Type: `"bool"`. Computed.

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

<a id="canonical-3001233132022020-1302112203132130-2003001321032003-3011103322011300-3132013131001021-3230023303101021-1300003223030303-2301331203001200"></a>

<a id="canonical-0212302232130231-3321303200222113-2230303211202311-0321330302131020-3002020110133121-2120012111110231-2223213113232122-3201321202100230"></a>

#### `policy_based_challenge.rule_list.rules.spec.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

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

<a id="canonical-1130112213311222-2212120101303321-2312021201330112-1010000001222113-3331020212132323-1321320012001100-3330022120222112-2310110330301333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.path

<a id="canonical-1110133321100212-2313321020102220-1023231221000012-1101223310211010-3022103100202332-0332132010113121-1132231223003202-1002323012302102"></a>

Type: `"single"`. Computed.

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

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

<a id="canonical-1120222330130302-0010321023000100-2113023110211011-1300103003310233-0331223111020332-1133000103202120-3022333002133311-1132333331313230"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.path`

<a id="canonical-2311201301010031-3131230123301300-0130002300010122-1311100331122011-3110023203001301-2002012310000303-0002210113312123-2300113232230322"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.encoded_path_matcher` property

Type: `"bool"`. Computed.

Match against the encoded, escaped path.

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

<a id="canonical-1030033001131333-2120023013302212-0223232230003110-2100021131100101-3130030201033001-2220120211212212-3321321223302303-3303120310000200"></a>

<a id="canonical-3002332220330020-2303000131110111-1200130033223000-0110212010002230-0013303312301100-3311133301000310-0023311021011010-2031332030213123"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact path values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2022232323003023-3120200330122032-0122021201101021-0312110232202312-3102311211203113-2002330132121113-2331001021223232-1132310223010010"></a>

<a id="canonical-1332022222123322-0130230233320200-0111100213321332-0131320200323220-1102332133223320-2212330230130223-1210321211133033-1301113311112121"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.invert_matcher` property

Type: `"bool"`. Computed.

Invert Path Matcher. Invert the match result.

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

<a id="canonical-0122102103133303-0122003133113201-3112203033033100-3022312130032212-1210232302200202-3103210022313022-1203121022102112-3230221012112010"></a>

<a id="canonical-3122220103231000-3011032310211101-0333030022210013-2023021230212330-3110130130322112-2020331320011033-3130331102021321-3211202313023113"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.prefix_values` property

Type: `["list", "string"]`. Computed.

A list of path prefix values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1303010311222332-2312333120111100-0213012223000212-1121011210031010-2111333300301321-1320120100023232-3202320321203013-2212112021000200"></a>

<a id="canonical-0112001023203013-1111332211131332-2123123332020333-3020031021333030-3033300333032103-1110001202133312-2003210120310200-3303313030311021"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0322033301131212-0100121220122233-2020211310220300-1323030330100222-2000112112212220-3223133232320200-3021001012231202-2132302332231001"></a>

<a id="canonical-0211313110300101-2131322233203102-3320112200320101-2213023020310123-1232132000122022-0310323022001302-1113002112323213-1223013301321330"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.suffix_values` property

Type: `["list", "string"]`. Computed.

A list of path suffix values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2111213330330132-0221322121333031-3132023330030202-3033032121132322-2331131222231210-2111320102001013-2030020311333233-3201230221231102"></a>

<a id="canonical-3131223032211033-0100022112032203-2113213020212221-2303023201100222-0120211211321012-0222122122301210-1201002103003213-2323123113213321"></a>

#### `policy_based_challenge.rule_list.rules.spec.path.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="canonical-1212322002201322-2110330103203321-2133302003022303-2102021332002111-0210003110120023-3223120321130013-3013111110202233-2013100003133003"></a>

Type: `"list"`. Computed.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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

<a id="canonical-2220033323233112-0333301322301133-3222112003013001-0112131033012322-1300213032002222-0100122311201212-1001000213230203-1333321130311013"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.query_params`

- [check_not_present](data-sources--http_loadbalancer--reference--group-023.md#canonical-3301330120000030-1310320133003031-1323322323331212-2013121021013122-0012103223220123-1230300003331333-0131331320110330-2000320031133221): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-023.md#canonical-2101113120020302-3300300103223201-3002302132220303-3202211103002112-3100000203200320-3310123021313333-0300232230301000-3220231220001332): complete subsection reference.

<a id="canonical-2300102022001000-2131310232002203-2120203000222023-2112011010002331-0213111330230010-1022010301301212-3212202333032321-1331101231222123"></a>

<a id="canonical-2111033121221221-3123232020232302-1330212110130202-3301023312302103-2213023230010213-2203010221312332-1102000233300131-3001313231331033"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.invert_matcher` property

Type: `"bool"`. Computed.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-023.md#canonical-1302012230310122-2012322122203331-1330220123231131-0311210132332223-0100310120103122-2311031213033211-2321133023211223-2103103230011023): complete subsection reference.

<a id="canonical-3202331131010032-1230011022331331-1011013021012101-3300102011303212-0233123010032213-3102221020310022-2313000331003321-1103300020010022"></a>

<a id="canonical-0201211323322112-0233032030103111-0221003320202303-2003103110233030-0020001320223112-0111103203321211-3030030102201120-3020232010222113"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.key` property

Type: `"string"`. Computed.

A case-sensitive HTTP query parameter name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3301330120000030-1310320133003031-1323322323331212-2013121021013122-0012103223220123-1230300003331333-0131331320110330-2000320031133221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-023.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332)
- policy_based_challenge.rule_list.rules.spec.query_params.check_not_present

<a id="canonical-0232230022001023-3131300232303200-0223002003012131-1032102022110312-2223302212110001-2202232333223012-3021023212303201-0312020221233003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101113120020302-3300300103223201-3002302132220303-3202211103002112-3100000203200320-3310123021313333-0300232230301000-3220231220001332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-023.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

<a id="canonical-0323233300032223-0113012322221010-2120302333111233-2133020220031223-0010112002201203-2123103310222022-2323300113101023-2100303131233323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302012230310122-2012322122203331-1330220123231131-0311210132332223-0100310120103122-2311031213033211-2321133023211223-2103103230011023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--reference--group-023.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="canonical-2232020031231011-0202320232223332-0020210213111212-2302200221132330-0320331013120122-3103020302203213-2331200013221112-1200112333211111"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-3021330022103111-1122010313212212-2133230031211220-2201201212233312-0320213230311131-3223123112102132-2220323210213113-0110032122201201"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.query_params.item`

<a id="canonical-1002301300202010-3221001102302100-0311123333020311-3133131201320131-2200121323223002-3133133210133311-1122011131020023-3311230003200203"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3011021323113331-2221120130013331-1100213003030111-0131301000023322-0023032031302220-1112121030000213-1122221010331110-1120013120133202"></a>

<a id="canonical-1022111202133200-0001000033203310-3333012220033020-2220221012313122-0130223022112333-0322022322011012-0132013221112003-0231333330122013"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1123331013023120-3312210220322131-2310221312012301-1031202110023010-2131122233120332-2132302203031122-0131021130033021-2101231203030233"></a>

<a id="canonical-0311121302203121-0202230322331123-2320323221001310-2210232122023313-1320311203031022-1103203000112222-0210303220222313-1231302023203010"></a>

#### `policy_based_challenge.rule_list.rules.spec.query_params.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0022113021303313-2123000230233031-0231021111110113-2111313001202330-1220202120123310-1300312201020212-0013033111222322-0130112033123023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-2032323200230031-2230223130230021-2010212122120201-0002213122002302-0010312312210032-3202122230132331-2201010133221113-3332233311003001"></a>

Type: `"single"`. Computed.

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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

<a id="canonical-2200201133221020-2311232130100333-3321323312120322-3033212100202100-2103000330300313-0200101033000121-1311001330131301-3201203101332131"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher`

<a id="canonical-3023111302321211-0313202303300311-2302112222000332-0013130211323210-2232032321201120-1013023223132013-0311232302030303-1301333031212301"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1212132003131222-3313321310030230-1302000012023213-0033010122123022-3230322303302130-3122312102023131-3321133000232133-1102213011123103"></a>

<a id="canonical-0233111113222030-0200231212201222-0123233113320032-1202202300302212-2211223102313010-1011103223012300-0022123121120131-3320223203220130"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3231121311030333-3213202031113011-0120021212113130-2011333232221211-1202132300020211-0013030023131230-1320111332221203-1011131233300303"></a>

<a id="canonical-3023332102313121-3211020112102212-3021133222110100-0203202111301101-2111201222211221-3120203003312313-3333231130012032-2031031212310303"></a>

#### `policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Computed.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3113023112302332-0333303302133233-0010212202113330-1313232030220213-3320210011012203-0101333020233000-0132210002222132-1003010103121113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.temporary_user_blocking` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-0123013121122200-3132220222010300-2103133032302133-3121000310102231-3002010121022012-2311132030130002-0303020232223021-3130203130013210"></a>

Type: `"single"`. Computed.

Specifies configuration for temporary user blocking resulting from user behavior analysis.

When Malicious User Mitigation is enabled from service policy rules, users' accessing the
application will be analyzed for malicious activity and the configured mitigation actions will be
taken on identified malicious users. These mitigation actions include setting up temporary blocking
on that user. This configuration specifies settings on how that blocking should be done by the
loadbalancer.

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

<a id="canonical-2003023112221111-3031211312012130-2121321111010111-1020230031322323-3113301010021113-3011032221332013-0221132302103203-2312013022123003"></a>

### Direct properties for `policy_based_challenge.temporary_user_blocking`

<a id="canonical-2002310133223303-2213332220030321-2030101103210323-0111013001321300-1320231023222223-3110211222023212-3330320101203001-2202110201130101"></a>

#### `policy_based_challenge.temporary_user_blocking.custom_page` property

Type: `"string"`. Computed.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Additional upstream details:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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

<a id="canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- protected_cookies

<a id="canonical-2020032312223312-0212212232031133-3330123312220112-3203300332221121-2002323120320111-0021333131321000-0113123012110322-0012332102112313"></a>

Type: `"list"`. Computed.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0030232211130121-3313011233323231-3301233330121022-2010323033323323-0222220201013323-3032233310110131-0213110222313031-1012200230312322"></a>

### Direct properties for `protected_cookies`

- [add_httponly](data-sources--http_loadbalancer--reference--group-023.md#canonical-3332312320111301-0221020103030022-3213101102123100-3210233313330331-1310222212231123-2130301330232322-2000013002022210-1131111301213003): complete subsection reference.

- [add_secure](data-sources--http_loadbalancer--reference--group-023.md#canonical-3103032303213002-1002003012103201-0231331000200303-3120331023011112-3031121001221103-0030023221130123-1133133303122222-1202021030000322): complete subsection reference.

- [disable_tampering_protection](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332111131001023-1121113212020001-2010113210213021-0010233213310012-3310330322032331-0300301310211211-1013232233231221-3222301221030203): complete subsection reference.

- [enable_tampering_protection](data-sources--http_loadbalancer--reference--group-023.md#canonical-2221011233302121-2002100113113302-0332000120300001-2213012233003032-1131133301221133-0132311310330303-2233112232013301-3032200323023221): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-023.md#canonical-3101113101300232-3030302222032113-2131003021320212-2203110210220320-3231020233100000-1012221032320110-1133331311232222-2110101132310203): complete subsection reference.

- [ignore_max_age](data-sources--http_loadbalancer--reference--group-023.md#canonical-3022011330020213-3203323031123113-2023123000131131-2202020303330303-2110333033302003-1223222312311102-0032101313211133-2233130113031233): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-023.md#canonical-1112303213003220-2233321130311001-1131213330200111-0123102102203121-0333223133122321-0022130233013103-2013120201121200-3000221310120323): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-023.md#canonical-2031220320331011-3213000230222112-2312333020302112-3012222132002322-0323121300222233-1101221101230121-3022010232320321-2233030112030211): complete subsection reference.

<a id="canonical-0220111010322223-1113300331111002-1210001113332021-3102130222002301-3123312331311231-1132222222010201-1021322102201003-1001322121212310"></a>

<a id="canonical-2311130201301132-1000103121330300-2331211321030012-2033021310311003-0232211120112001-0000023002331303-0203212300322120-1112202331210201"></a>

#### `protected_cookies.max_age_value` property

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-0322120232121001-2103022031003211-0333031023003312-2213012200020321-2120022311033332-2321313033323020-2002222212120330-2122132010110101"></a>

<a id="canonical-0333200130120313-0202203021012230-2112022213323231-3012211133002200-3110102003022031-3310012320231010-3233101020332210-1331002103333333"></a>

#### `protected_cookies.name` property

Type: `"string"`. Computed.

Cookie Name. Name of the Cookie.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [samesite_lax](data-sources--http_loadbalancer--reference--group-023.md#canonical-1123212322233202-0233321102200330-0002211133323011-2322000210121213-0322310000031221-3013130133232022-2232302100211102-1032223302003300): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-023.md#canonical-2120310030321121-3330023020102312-0032111100030010-3333301223020213-2231232123133030-1103110321010323-3113033123003312-3030122230310220): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-023.md#canonical-3023033032321212-1120133122002230-0202212130200011-1003023012130132-1011103132132210-2110023030320011-0233033111103022-3331323121301222): complete subsection reference.

<a id="canonical-3332312320111301-0221020103030022-3213101102123100-3210233313330331-1310222212231123-2130301330232322-2000013002022210-1131111301213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.add_httponly

<a id="canonical-0230021303020312-1002110112121301-0231133133023301-3203033011010100-2013000112111200-1302002101131100-0131202232010010-0321003301010330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103032303213002-1002003012103201-0231331000200303-3120331023011112-3031121001221103-0030023221130123-1133133303122222-1202021030000322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.add_secure

<a id="canonical-2132121311203201-2003200010122201-1010321020200002-1313131002210032-0102312312111100-0112303020012112-0102120131302130-2200133012210113"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332111131001023-1121113212020001-2010113210213021-0010233213310012-3310330322032331-0300301310211211-1013232233231221-3222301221030203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.disable_tampering_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.disable_tampering_protection

<a id="canonical-3310130223000132-1110002231113313-1213212213301320-2203013132112323-3213200233233222-2120030022031223-1220231231011010-1313232320021211"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable tampering protection.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221011233302121-2002100113113302-0332000120300001-2213012233003032-1131133301221133-0132311310330303-2233112232013301-3032200323023221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.enable_tampering_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.enable_tampering_protection

<a id="canonical-0310030122012322-1100113010023333-3110220102103212-2113301110000221-0223130332132333-0320033020312021-1212221200320230-2110210012131003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable tampering protection.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101113101300232-3030302222032113-2131003021320212-2203110210220320-3231020233100000-1012221032320110-1133331311232222-2110101132310203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.ignore_httponly

<a id="canonical-3002032230032132-2331021121200000-3032211301012330-0323211220233102-1110011001023332-1211232333232332-3132030110203030-3001320201203023"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022011330020213-3203323031123113-2023123000131131-2202020303330303-2110333033302003-1223222312311102-0032101313211133-2233130113031233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_max_age` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.ignore_max_age

<a id="canonical-0200331221323110-2030232021000302-2212333112103101-2032302033130233-0133203012113223-2003101100111300-0120121323010331-3131233211330000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore max age.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112303213003220-2233321130311001-1131213330200111-0123102102203121-0333223133122321-0022130233013103-2013120201121200-3000221310120323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.ignore_samesite

<a id="canonical-1112322323222131-0010003100010112-2103321310231223-3113033321233311-3131232333130121-3131003211123130-1220212223333031-1201110112100232"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031220320331011-3213000230222112-2312333020302112-3012222132002322-0323121300222233-1101221101230121-3022010232320321-2233030112030211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.ignore_secure

<a id="canonical-1032333033020331-0220212101303033-1321132032311301-2333000300132321-2303101011311230-1010222312210001-1010302300021231-1031310302213032"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123212322233202-0233321102200330-0002211133323011-2322000210121213-0322310000031221-3013130133232022-2232302100211102-1032223302003300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.samesite_lax

<a id="canonical-1302020010200300-3201032001030223-3130312022112022-3333103233110031-0321110101000321-3022311020231022-1110001001023200-3011322203120230"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120310030321121-3330023020102312-0032111100030010-3333301223020213-2231232123133030-1103110321010323-3113033123003312-3030122230310220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.samesite_none

<a id="canonical-1103100123120220-0033023322333200-2021011212022231-1010322203011130-3002122112203331-3233330232110300-1133021033300031-2003303002101020"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023033032321212-1120133122002230-0202212130200011-1003023012130132-1011103132132210-2110023030320011-0233033111103022-3331323121301222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `protected_cookies.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [protected_cookies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2332103023230011-3100322302220110-1030000021023233-0230033122113003-3202333211321333-2313233021311312-1001320211331020-3331221301201111)
- protected_cookies.samesite_strict

<a id="canonical-2101312301022022-2113112203322233-1222102113301100-2311230201320103-2000201332110023-0111131020202022-1232123202013121-2120113122032323"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211233301202123-3021122113200023-3121032321113322-3231020200122333-1021331133012303-3330210102330302-1000033231000302-3302120223110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `random` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- random

<a id="canonical-2123033220232332-2220220302011202-0000233120322210-2323011320220213-3231120212211230-3023111203101113-0102303132032332-3003021103131300"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- rate_limit

<a id="canonical-0111113122010303-2301222300220321-1211003332003113-1321232202233231-0110133202230113-1313110222123003-1231230132033110-0121210312001113"></a>

Type: `"single"`. Computed.

Load-balancer-wide per-client rate limiting. The counter applies across every path; use
api\_rate\_limit rules when only selected paths such as /login should be limited.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]",
  "x-ves-oneof-field-policy_choice": "[\"no_policies\",\"policies\"]"
}
```

<a id="canonical-1012131100322032-2302020202111131-3011300003211022-3303331203202112-2223101130222222-1333331313023200-3131022110113302-1101122022333012"></a>

### Direct properties for `rate_limit`

- [custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-023.md#canonical-3012011302113031-2233310123210100-0101123331122212-0023001231213021-0130013312023002-2200200330001100-0332121300030230-1330013132000333): complete subsection reference.

- [ip_allowed_list](data-sources--http_loadbalancer--reference--group-023.md#canonical-1123102130330220-1301033010013123-2232330102200012-0000133230320131-1223310232213332-0121230212030202-3122103201121103-2323303003121102): complete subsection reference.

- [no_ip_allowed_list](data-sources--http_loadbalancer--reference--group-023.md#canonical-1000001213220102-2203102001203203-1101002130331122-1102132331031001-1111321303002132-2103233202021033-1310303011110031-1321121123131310): complete subsection reference.

- [no_policies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2311212103213011-1023003113233321-0332030111302223-3110120213230333-0231322130011200-1110332332213233-1120111133010023-1131000212200201): complete subsection reference.

- [policies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2323120302221121-3022003003212333-1121333003012301-0311300013202023-0211323012212332-2000000132001202-0320312020001031-2100313002132012): complete subsection reference.

- [rate_limiter](data-sources--http_loadbalancer--reference--group-023.md#canonical-2001202332221232-3333020122211202-2233223223033313-1230201322322310-0013102301231100-3030011221013313-3003132022332331-1003223103230131): complete subsection reference.

<a id="canonical-3012011302113031-2233310123210100-0101123331122212-0023001231213021-0130013312023002-2200200330001100-0332121300030230-1330013132000333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.custom_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- rate_limit.custom_ip_allowed_list

<a id="canonical-0210230310202221-1110302111102323-0131120121013102-1222030203110000-1322113023001210-3220310002111110-1303231010330233-0101013012133002"></a>

Type: `"single"`. Computed.

IP Allowed list using existing ip\_prefix\_set objects.

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

<a id="canonical-3231133230021113-0320120231001321-0211311222300122-0301200022112030-1120033101022110-3221330020133032-0220133122001222-2302023122120113"></a>

### Direct properties for `rate_limit.custom_ip_allowed_list`

- [rate_limiter_allowed_prefixes](data-sources--http_loadbalancer--reference--group-023.md#canonical-3111211333031032-1001113230113112-3110212300133001-2332112121110323-1213203112022100-3200031210000010-3113333112131333-2123133332332301): complete subsection reference.

<a id="canonical-3111211333031032-1001113230113112-3110212300133001-2332112121110323-1213203112022100-3200031210000010-3113333112131333-2123133332332301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- [rate_limit.custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-023.md#canonical-3012011302113031-2233310123210100-0101123331122212-0023001231213021-0130013312023002-2200200330001100-0332121300030230-1330013132000333)
- rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-2022032010233302-3322102030230110-1022322211200231-3121112311302010-3002211333030332-2321130021222211-2110111331323123-2112211310032323"></a>

Type: `"list"`. Computed.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3020313021321020-0103203012331033-1223321113223232-3303101313301121-0233023301233203-1100220322101112-0032121320100323-3001210330303002"></a>

### Direct properties for `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes`

<a id="canonical-2123323122311000-1322312002231323-1102303310012232-1321200130203100-1023022103022111-2111211011303330-1321202312213020-3300023201211231"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3011301123130023-1130210011002101-0322121122022011-1221322001030220-3003002213330122-0211001130120013-2100333010300130-3123113122210002"></a>

<a id="canonical-0333103030120032-1023131221313101-3301100322322121-0313100211000201-2323003021003111-1112022310330110-2030303223331331-2131010302000311"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0303212332220212-1310220330211001-1133110000320003-1121321220130003-1121222323210313-3211210112212122-3023010313220321-1011210213133112"></a>

<a id="canonical-0320200313333132-3113302120202232-0032330012300222-3111233132032222-1330012020220300-1323213333210221-2300022323033131-0210121310130303"></a>

#### `rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-1123102130330220-1301033010013123-2232330102200012-0000133230320131-1223310232213332-0121230212030202-3122103201121103-2323303003121102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- rate_limit.ip_allowed_list

<a id="canonical-0303133110131232-2322231003112212-3311201201030033-0130003122113023-2320300220322313-0303130100113013-0330032212013321-1012312313121012"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-2113330303010032-2200320332310021-3012030300012102-3230300231322000-1311111230102311-3101131122301201-0200112102101331-3012312212103231"></a>

### Direct properties for `rate_limit.ip_allowed_list`

<a id="canonical-1211223133130002-1233211110230130-1132302031003012-0123312231021220-2120123032011311-3233110303133323-1332221223233111-0122322200230200"></a>

#### `rate_limit.ip_allowed_list.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1000001213220102-2203102001203203-1101002130331122-1102132331031001-1111321303002132-2103233202021033-1310303011110031-1321121123131310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.no_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- rate_limit.no_ip_allowed_list

<a id="canonical-3230101232332123-0210023122011233-0311131020030202-1310002011232111-3112310222023013-3311313210022121-3033313122210103-1210121120131032"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311212103213011-1023003113233321-0332030111302223-3110120213230333-0231322130011200-1110332332213233-1120111133010023-1131000212200201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.no_policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- rate_limit.no_policies

<a id="canonical-2011200303201210-1113332323021313-0033330110303323-1230300201321110-1211021031132210-3321310222220331-1213301002111010-3231302210231312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no policies. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323120302221121-3022003003212333-1121333003012301-0311300013202023-0211323012212332-2000000132001202-0320312020001031-2100313002132012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- rate_limit.policies

<a id="canonical-2130221312311120-1201112213003103-1023011021100113-2003203233300020-2200321013110102-2311133100020021-2223131303331331-0101311200133231"></a>

Type: `"single"`. Computed.

List of rate limiter policies to be applied.

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

<a id="canonical-0203131030300330-3102332132330012-1211211223321210-3032023003210010-0102333301102120-3110003200011200-0011122101112223-0232111213311131"></a>

### Direct properties for `rate_limit.policies`

- [policies](data-sources--http_loadbalancer--reference--group-023.md#canonical-1013012201311110-3001231300321101-0130222332210101-1311022022003033-3033110031133203-1310000331022112-0121023130333121-2121222310110231): complete subsection reference.

<a id="canonical-1013012201311110-3001231300321101-0130222332210101-1311022022003033-3033110031133203-1310000331022112-0121023130333121-2121222310110231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.policies.policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- [rate_limit.policies](data-sources--http_loadbalancer--reference--group-023.md#canonical-2323120302221121-3022003003212333-1121333003012301-0311300013202023-0211323012212332-2000000132001202-0320312020001031-2100313002132012)
- rate_limit.policies.policies

<a id="canonical-0130221000200120-2323301000330222-3031333221300000-3312323000211313-0120323312311022-0231201131011230-0032232121010232-1022012131102111"></a>

Type: `"list"`. Computed.

Rate Limiter Policies. Ordered list of rate limiter policies.

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

<a id="canonical-2020300313020132-1122220020112201-2101222213232320-2102133300210133-0032301110001133-1331322333010320-0310033231121333-3000120011212003"></a>

### Direct properties for `rate_limit.policies.policies`

<a id="canonical-2212031223002230-2121012110023230-3032313102232110-0311130031123023-0132031021221313-1013220223211212-2032011023332210-3311301132100220"></a>

#### `rate_limit.policies.policies.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2012133221032123-0210132221333301-0330221102321211-2013210301200231-1032302023031302-1222303220112300-0133201113013312-1301313132101220"></a>

<a id="canonical-3102320323001213-3103100321212201-0332231103330320-2322320033111122-3112211101232320-2303303231221230-0220021310201013-3102312313311320"></a>

#### `rate_limit.policies.policies.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2011302303123212-3303211030323021-3033131123101302-0303322020320220-2223200130212330-3023022222310210-1311303100033001-0101330021223111"></a>

<a id="canonical-3103311331230211-2230132223003201-3223222212332203-1321223203121313-3321233111100010-3002121302002211-1131030122330313-2322222231012202"></a>

#### `rate_limit.policies.policies.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2001202332221232-3333020122211202-2233223223033313-1230201322322310-0013102301231100-3030011221013313-3003132022332331-1003223103230131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rate_limit.rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-2202131303231211-2311213323110101-0233201202330113-1123131100201021-0201222220031000-3133313321222112-1313230220231100-0232331032130321)
- rate_limit.rate_limiter

<a id="canonical-2030233131311012-2200023001220203-1210122310002022-0132221212210012-2132031023100013-3232021301133320-2002211310310131-1132321200012020"></a>

Type: `"single"`. Computed.

A tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"action_block\",\"disabled\"]",
  "x-ves-oneof-field-algorithm": "[\"leaky_bucket\",\"token_bucket\"]"
}
```

<a id="canonical-3132023320313111-2211303201211133-2033311023102123-0203310312330323-1331100302230311-3003111332113120-3000210302211201-1202232000100333"></a>

### Direct properties for `rate_limit.rate_limiter`

- [action_block](data-sources--http_loadbalancer--reference--group-024.md#canonical-1200211123003010-1233100203111302-0210123300133003-2122300031101021-1201313133330132-2231222222013231-2311202021020223-0023311303233311): complete subsection reference.

<a id="canonical-3022023003332303-1312332322303100-1013231131213213-1110202111110120-2310133232313203-1032201222313333-3033311000323101-2313200102021330"></a>

<a id="canonical-3103311220110323-1203112333211023-0333332010233201-3310312120032312-0130122312313311-2002132022020232-2222112313322003-3020313313110230"></a>

#### `rate_limit.rate_limiter.burst_multiplier` property

Type: `"number"`. Computed.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](data-sources--http_loadbalancer--reference--group-024.md#canonical-3000320012303032-0010200332232323-1213021202222230-2302102123202102-0111031331020030-0022220102102200-1000221213100020-2102210110132210): complete subsection reference.

- [leaky_bucket](data-sources--http_loadbalancer--reference--group-024.md#canonical-3102303111103122-1313202011123133-1200301113011002-1220133132121012-0003013002301132-1010100032120132-3002323202122132-3030102232323133): complete subsection reference.

<a id="canonical-2010323030031321-0300033102130111-0313323312102323-1122232100013020-3003021023132323-2222203233013222-3232311300202231-1231031122203001"></a>

<a id="canonical-3213301133231113-0011231230230122-2302220111301302-1332332022313012-3230202121003213-2113330201323213-3001203113200100-2232123002300231"></a>

#### `rate_limit.rate_limiter.period_multiplier` property

Type: `"number"`. Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

Additional upstream details:

This setting, combined with Per Period units, provides a duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

- [token_bucket](data-sources--http_loadbalancer--reference--group-024.md#canonical-3303231201200320-2100220323123301-3112332201232200-1311113100210010-3131102021100010-1233103202133031-1031230200300002-0010311031033232): complete subsection reference.

<a id="canonical-0301220013333001-1221233211131200-0022220312223102-3233313020322322-0200011301302012-1210023202102320-2010233032312020-0233211230021110"></a>

<a id="canonical-2000230022022230-1000012021211013-1113302133011313-2113330030131202-3200220332002203-1200313220220032-2322211000311231-0223132130123132"></a>

#### `rate_limit.rate_limiter.total_number` property

Type: `"number"`. Computed.

The total number of allowed requests per rate-limiting period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-2112033101330320-0320320200012312-0200012303030031-2001011101103001-1032331101003301-0223213013102313-2301232223233220-3331333322310120"></a>

<a id="canonical-2330020200022112-3013020321032303-3222322222003013-1120220213223330-0203000300002121-2120200322123301-3300201012032221-2121012213110313"></a>

#### `rate_limit.rate_limiter.unit` property

Type: `"string"`. Computed.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
