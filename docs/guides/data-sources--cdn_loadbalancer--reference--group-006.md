---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-0103010200011033-2303013300303311-3301011332130110-0003030230213130-3203113203330303-3012012031120030-2102130230232103-3012000130112300"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list`

<a id="canonical-0010322323020310-0001302200322202-1233033130302121-3223132211300210-0123002333232311-3231003113203000-2310230321131231-1000021331001320"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list.invert_match` property

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

<a id="canonical-2232302110010022-3313103201102303-3213001131201132-2011200211123323-1012213133332211-0001030030301023-2010213033211030-3311023233210301"></a>

<a id="canonical-1332010120233221-0102200312133110-2021201030320001-2313311132222202-3132020011022102-1001232332133002-1222330120323200-0232232120201202"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list.ip_prefixes` property

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

<a id="canonical-3122022200133313-0102331311012322-3021212000002223-2000203333200202-1322331230332033-2221130003120212-3032002303233211-2303213201111032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-1012022123311200-2120033023303033-3011323210331001-2203231020330221-1022102320231312-0020013020103201-2301312131021323-3231222230000003"></a>

Type: `"single"`. Computed.

IP Threat Category List Type. List of IP threat categories.

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

<a id="canonical-2010022221323220-0120103101010130-3013133301010330-0130320021030212-0121002032112332-0121230032211023-2131213111113230-0002132222001202"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list`

<a id="canonical-0320202202323311-2002133230212320-2110333110032230-2130031003332221-1233032233200131-2232321112101211-1001333222103112-0203112032231100"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1331332223312102-1132112313212300-1023223213002222-2333213112000033-3131223010112201-3111213233133211-0202302112202003-2021131013002322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-3021132223233021-2333211013310010-0032232021113202-0222320100011201-0001233022332132-3021003332303121-0022033103133130-3121110100122233"></a>

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

<a id="canonical-2302303231110303-0211303113233320-3322310110201123-1210020132032023-3031032020112200-0320213223030322-1110012220320133-3322100210221010"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-0012102330111020-3310233201111112-2332121200331120-1001033001202111-2020303100332321-2333310233301102-1211210222203331-3012103310211103"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.classes` property

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0111021302311313-2331022232331130-0213101223221101-1002133130302203-2201303033200111-0021320031021002-0011210303232221-3110013303301012"></a>

<a id="canonical-2133030312322023-1130033313223232-2101313212111101-2233231103032320-1131132210111010-2020211332033123-3222023111022222-2132330030332103"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

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

<a id="canonical-2313030013121230-1322123313202333-1020223013211201-3322101001223332-0221212211303111-1203320021001200-1120133221300323-3020223021211100"></a>

<a id="canonical-1303200122112312-0100310300222003-0003120220111111-2010321332003030-3320131012113302-0310130101133221-2310201301211020-1320220323030212"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

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

<a id="canonical-3020203131201121-0321103000110021-0000001321301203-1220123001023020-3230213101001103-0331101023220030-3223330201312212-3232321223223320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-0110030131302220-0313212131333122-1233310303300113-3120132122303302-2332102313213003-2132001031302122-2002020220331122-3200312103000031"></a>

Type: `"single"`. Computed.

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

<a id="canonical-0011112321130220-1313221312212313-1000313010021322-2232032203321011-0210113103212000-0301332310203221-1313123333330323-2103231230301202"></a>

### Direct properties for `api_rate_limit.server_url_rules.inline_rate_limiter`

- [ref_user_id](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3231133310000120-3211032302203321-1332331220130321-3313122302003213-3012233131133101-3321013331312301-2103321112331321-2310333123031130): complete subsection reference.

<a id="canonical-0231120203220331-1322303033132222-2213220101211312-0321132002223032-2103120320022030-1103223312003322-1013100023103322-1300230213333113"></a>

<a id="canonical-1123302101300210-1323021023201103-2133211321311112-3220222113322331-2123230200310111-2031033300011320-3320332330112011-0230103210230311"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.threshold` property

Type: `"number"`. Computed.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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

<a id="canonical-1301210211110023-2301300221030333-0213233131201333-1310021033032211-3332100030022300-1321001302121131-2323323200001132-3101011110312333"></a>

<a id="canonical-1121013211123320-1312302011320122-0301130110301311-0322333020330033-0233330131132003-2032113102220112-1002013300123003-1003322000221303"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.unit` property

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

- [use_http_lb_user_id](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3303003110311231-1300100313220132-1033031032233330-1213322120310203-3331231301333213-3121000011121200-2121221232332002-0010232122012112): complete subsection reference.

<a id="canonical-3231133310000120-3211032302203321-1332331220130321-3313122302003213-3012233131133101-3321013331312301-2103321112331321-2310333123031130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3020203131201121-0321103000110021-0000001321301203-1220123001023020-3230213101001103-0331101023220030-3223330201312212-3232321223223320)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-1331020331311213-2020310122021002-3312101201013121-2232130130023321-0321133210320320-3101303202013122-3233311113212121-3223011221021203"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2321231023311331-1002112220233010-2100320033223312-1233003021102232-3331333213131333-2001033230220200-2103001010013230-2121031230231132"></a>

### Direct properties for `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id`

<a id="canonical-3003310123112230-0232203300133331-0232201131101233-3133022111200231-0203211012002210-3032101001130000-0331022001100323-2312102311002122"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.name` property

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

<a id="canonical-3222122132111213-3000321232320301-3303300320010301-2321320103302211-1202030222303133-3100320230232023-1121320002223200-3131122110331001"></a>

<a id="canonical-2323222223031111-3131333302230031-2031301123021321-1233101003332230-3230121101302302-0332022220032200-1123122132130213-1113211133201222"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.namespace` property

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

<a id="canonical-0120223123133001-0313030033202331-3110301222320212-1332302032131130-0111300110323322-3211333213323101-3013221121233010-1322131101000122"></a>

<a id="canonical-2301331111222011-1131133220130313-3021302100231033-3001302011003200-2001030033023302-0032211122300022-0330333011100213-3003201121222213"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.tenant` property

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

<a id="canonical-3303003110311231-1300100313220132-1033031032233330-1213322120310203-3331231301333213-3121000011121200-2121221232332002-0010232122012112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3020203131201121-0321103000110021-0000001321301203-1220123001023020-3230213101001103-0331101023220030-3223330201312212-3232321223223320)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-0333322331300213-2011213312211130-2321023211112232-3012310323131123-0203112013321021-0101333023000103-1011113133313012-2133022112121121"></a>

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

<a id="canonical-1121003122321003-1030113023222001-1120323002113122-1103033302313332-0220011200302103-2120112030110322-1320212132033201-2023201113201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.ref_rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-3012030002310130-2100121211230112-0100232010333101-0311103121231301-3202133021203023-0302221322021110-0011130013301213-0221303232010312"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Additional upstream details:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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

<a id="canonical-2231123221100320-0202210120020220-2322023113313112-2030012202121223-2310132201201011-2122032111003312-0110202222132310-1032301322330133"></a>

### Direct properties for `api_rate_limit.server_url_rules.ref_rate_limiter`

<a id="canonical-2333203222330323-2332021223301222-1023202003103332-1002023312201120-3112201212103001-3230013221020000-2130212020211003-2331100220020121"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.name` property

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

<a id="canonical-1013021132013231-3210212131112012-0101000013013101-2303103312231330-0222213012213001-1200203321312101-0122222230233120-0210201010301112"></a>

<a id="canonical-0303120333023123-2031321033112233-3110320223323033-3330133122303113-1312022332321120-2122221233221222-0012023023003202-1201310020003112"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.namespace` property

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

<a id="canonical-2300111310112011-1210102131202130-2013003102102221-3202333100201112-0200231113331212-2302113331130330-3312333331210132-2211022003130333"></a>

<a id="canonical-2032033102332130-0322023001333103-2330321212221133-0202132300112221-3120312031312122-2320210231132202-3020212131211132-0122102322110022"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.tenant` property

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

<a id="canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-0003303123012330-2303130220210200-2311013220322133-3330120302100130-3133022220211221-1223011031223113-0123330232332230-3020203202310103"></a>

Type: `"single"`. Computed.

Configuration parameter for request matcher.

Additional upstream details:

Request conditions for matching a rule.

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

<a id="canonical-1232021223100333-1322010032102030-1013112102010130-2101020223101001-2030321313333033-3311312320003100-3302021022100220-0002121210211033"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher`

- [cookie_matchers](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311): complete subsection reference.

- [headers](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102): complete subsection reference.

- [jwt_claims](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132): complete subsection reference.

- [query_params](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010): complete subsection reference.

<a id="canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-2321312320331132-3001003321112331-0010210132103201-0303133101112303-2002100311230302-2321112022310102-2000110021100120-3123213201210222"></a>

Type: `"list"`. Computed.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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

<a id="canonical-1100331302033103-2120123332121111-0002032132110333-1013021332011002-3012213211312311-1100310010033202-3202303112312121-2130123133103030"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.cookie_matchers`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0320322103103020-2313313232322031-0000321023013022-0223111311311020-3122022310110320-3230111122313001-0120203130212333-1032002100332332): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-2022310110121332-1001203013102033-3031021310111131-0203233032131331-3322211210121222-0023111130010123-1012012100333230-0201332112220201): complete subsection reference.

<a id="canonical-3130313330330220-2201300321203003-3002003201013121-2003310232311002-3201010001112331-1033211301203122-1210020031132122-3310031100033332"></a>

<a id="canonical-1101020302023022-0301120230221331-0210101123001303-2303212130122122-2012203111302220-1113332032331121-0121200322212323-0122022002232330"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.invert_matcher` property

Type: `"bool"`. Computed.

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

- [item](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3010001213003122-0330233130333320-2101133003200200-2303102332223321-1033332013330213-3220230323322301-3323023021023003-3231030302330311): complete subsection reference.

<a id="canonical-3333302000133133-2231003101100211-3311200230223033-3110020031211002-0031000203101121-0131311011132130-1331013333233101-3231120002031213"></a>

<a id="canonical-2110301110001020-3013223100332302-3102233322121030-0012203131132130-0121011123003303-2321031030322202-2222232331121033-0323211303330032"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.name` property

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0320322103103020-2313313232322031-0000321023013022-0223111311311020-3122022310110320-3230111122313001-0120203130212333-1032002100332332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-3301233301232200-1033222231223322-3021332123020112-1313222012221332-0212300110120321-3323210111020230-0232000330012202-1311302232023000"></a>

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

<a id="canonical-2022310110121332-1001203013102033-3031021310111131-0203233032131331-3322211210121222-0023111130010123-1012012100333230-0201332112220201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3311030303330321-2223222301222030-1201111223201101-0002031321231213-0021333132123120-2211310133010313-1122330101222211-1323002221100101"></a>

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

<a id="canonical-3010001213003122-0330233130333320-2101133003200200-2303102332223321-1033332013330213-3220230323322301-3323023021023003-3231030302330311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-2222033032003020-2000032133110232-3330123303122131-2122101313120202-0130301121310031-3220333112012131-3010012300001102-0012122311203332"></a>

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

<a id="canonical-3101030120310000-2113301223203112-1200200031111122-1303310023013020-3030120003012123-3131031030112322-0111333103010332-0213100023133113"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item`

<a id="canonical-1311220230222101-3013032320131100-3011133230013030-3032101311111332-3220131310233132-3213231110211301-3101311303133022-0300211102201310"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.exact_values` property

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

<a id="canonical-3020230330202131-1023131111201031-1123011021322003-1120122010032202-1332222213111112-2311132033233233-2112113030121322-3310303202313213"></a>

<a id="canonical-2023330133031203-3223222013202102-0132333033320322-0003032131121200-3330201303032010-2213232332200321-1121312332030332-0323232130200112"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.regex_values` property

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

<a id="canonical-3013122031202331-3033111301210112-2122300102102331-2321030230312010-3330012300012312-2200320303103030-0012000223123120-1220220201333333"></a>

<a id="canonical-3213303031020323-3002213111133031-1302100231312022-0122030301231301-1311130130101003-0200002133201032-1201220022222302-1103032022132203"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.transformers` property

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-0002320110110031-2130232230111222-1011030100322310-2031312120211013-2201031100213311-0223220232320133-3123132330102311-2103031020330032"></a>

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-3313333111123020-0012212310101111-3211321312120100-0120323230210311-1230311313101021-1011331023322000-0111130013311322-0311131101030201"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.headers`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0010103220103033-0132101131333333-1311001203302210-1021232313303321-2303203110232331-1230322110222202-0032302113102102-3112010220031000): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-2023231022000000-1131113332332132-2131233330010221-1013000120011020-2102220131033021-3320023310320102-0103311333021202-0232033103120221): complete subsection reference.

<a id="canonical-0202001211010110-2301231320311323-0311031230201023-1312213112320023-3300001001222011-3011231021010322-3320331133100123-2300000133011123"></a>

<a id="canonical-3213011033213111-3032030331220132-0313301232220103-0331202202210020-0022303203111222-3011231301201013-1031321212123201-1110220103312313"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.invert_matcher` property

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

- [item](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3022310101300212-1322011313030311-0231333020333233-0011201200100310-2123320120310103-1213320233012230-3230112002012200-3123103010320233): complete subsection reference.

<a id="canonical-0001032012023003-1330103010000102-3223331332213202-3100110001321233-1302022120021023-3133203010012111-0102022301223203-0030212111323112"></a>

<a id="canonical-1302132013022321-3213310032013332-3322232003203102-0310330201303212-1323032111310213-2033323303302130-2020111000222303-0121022001033032"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.name` property

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0010103220103033-0132101131333333-1311001203302210-1021232313303321-2303203110232331-1230322110222202-0032302113102102-3112010220031000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-3032100301212001-3011123110301203-2312201002311102-0111130011013331-2200031202230010-3233301222013211-3011120012231312-0232332203301132"></a>

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

<a id="canonical-2023231022000000-1131113332332132-2131233330010221-1013000120011020-2102220131033021-3320023310320102-0103311333021202-0232033103120221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-2212232202201331-0330100202300311-2320013001230222-2101022021310321-0232031220231323-1220221332031323-0021022112002211-2002012212321323"></a>

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

<a id="canonical-3022310101300212-1322011313030311-0231333020333233-0011201200100310-2123320120310103-1213320233012230-3230112002012200-3123103010320233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-0032312021212130-0221113013320233-3333312131223102-0101102231313120-2233322212100003-0112131023203012-3312321012132013-1030233332321031"></a>

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

<a id="canonical-1130103303023130-2311332022121332-1100312132321212-2101122032102112-1033220331030303-3020003220023310-1133121122122230-2133230122121313"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.headers.item`

<a id="canonical-3020300303131202-0311222232121200-0100021333101012-1330312211331211-1122132133101230-2330200103302203-1123200111210200-2002222321322020"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.exact_values` property

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

<a id="canonical-2323312020111021-3210030131330332-3113322320330122-3332012331000200-2233223312310131-2111232032332311-1113132023230311-2200302103110132"></a>

<a id="canonical-3212011102332031-3120002023033213-2203032100321201-3302211203211120-1320333031021302-0322103233203032-2303023103111110-3012200212231210"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.regex_values` property

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

<a id="canonical-2231132202331201-3323013210321110-0132310321221032-3000310103000230-3132323020133022-2303120000210313-3231302203320030-2322123331231312"></a>

<a id="canonical-1033123332111311-3123321123011303-1030012131013110-0021111023112023-3111030120111203-2202323223010021-3133222301013031-1311330103201300"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.transformers` property

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-0311232131121030-3211231300310002-1213211133131233-3302100221201122-1131323131221210-3200303231321210-3332123330103233-2113232232011022"></a>

Type: `"list"`. Computed.

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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

<a id="canonical-0202302120110111-0121133021220313-0332132103003021-1301221011203021-1321322202332102-2033101122102121-2032001012123031-1031123102132311"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.jwt_claims`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0333101330332230-2332030123201200-1303333113012133-2023231322030100-1201011231002223-0331323011102222-1212303110132330-3030321311110300): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0200313122001322-1313112000002023-2020202313101330-3223130102323210-1203332201311103-2031031020113210-2231211121301111-1213011330030223): complete subsection reference.

<a id="canonical-1213301222231210-3320022313322121-3230223321330320-2232022222032101-0203030021233112-3213101300020121-3333311030312202-3323333032222220"></a>

<a id="canonical-1102121112101230-3330302333121311-2001223332330030-3132013103223312-2300211122003313-1031111310331133-1220200210010212-2003121130010131"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

- [item](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-2321232222300023-0132311022213301-2202132022301222-0322011310133320-3301313311210300-0202300020110100-2300303000223323-0101103132332333): complete subsection reference.

<a id="canonical-2321012033222321-1231311113000032-3132302210012233-0320003102102101-0103221032103320-3330213311221213-3010112001333231-3333012203100333"></a>

<a id="canonical-0030202121311030-3032331232131002-1211221201332311-1002201231330202-3022123021212130-0022110123001310-1102021013331102-1101112033011231"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.name` property

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0333101330332230-2332030123201200-1303333113012133-2023231322030100-1201011231002223-0331323011102222-1212303110132330-3030321311110300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-3331010132013021-0333320112023002-1121331123203313-1220131011002200-1131331123132021-0020232203110200-1321201230121002-2132032132221000"></a>

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

<a id="canonical-0200313122001322-1313112000002023-2020202313101330-3223130102323210-1203332201311103-2031031020113210-2231211121301111-1213011330030223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2210120313000013-0333231223002211-0301121033200230-3120300122330023-1012231331211303-3123120121312120-3123222311120223-3302002120103111"></a>

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

<a id="canonical-2321232222300023-0132311022213301-2202132022301222-0322011310133320-3301313311210300-0202300020110100-2300303000223323-0101103132332333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.item

<a id="canonical-0031121031022221-0110032202112033-0112331332132300-1311331323131221-2113120203012302-2302030321111102-1100201012013200-1230122002102232"></a>

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

<a id="canonical-3133233010233333-0112220200012111-0312323102103321-3322313133221023-3033111131320200-2130121332011132-3302321002103311-3200123100330312"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item`

<a id="canonical-3302333201102030-2111322331302110-2121320100333132-0211301330312033-3332230302202322-1113110121301233-0321030320220012-1302220131132303"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.exact_values` property

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

<a id="canonical-1031323221223213-0302020331222222-3101012200112123-0312231232332121-0211022031122211-0031311011233320-2132223203020302-2333102023232231"></a>

<a id="canonical-0001301222220231-2100202333103033-3211011102030213-3030221203133212-3030100023002121-3232023312211302-1023033021322111-1102213023223001"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.regex_values` property

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

<a id="canonical-1212303320033110-0331100333120101-2023233113212030-2020103030233032-0112320301113001-3201331212030322-0330220032130331-3020231031013113"></a>

<a id="canonical-3201220313313121-0002220203220312-0312033032321223-2123300100221233-1032230032323023-3210031302111121-3100300110010012-3202323132121330"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.transformers` property

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="canonical-2120332013323121-0001231332031020-2113321023313332-2300013032231102-0232233011110030-3020311002303111-3303200213021131-3110203332132001"></a>

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2230100312120333-2323320222312131-0011200230121230-2121233232301120-2331220323320020-2122132121010002-1202110032101232-2010320300030003"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.query_params`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3321231032133332-2332011032203202-1012330112120330-0122123202131132-0121022230003033-0312212130223213-0111002303001212-0310022131300332): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1311122023131330-2222100033212110-0313000012330101-0032333320201000-0023230122102111-3330130030323001-3023020010320130-3312121303331001): complete subsection reference.

<a id="canonical-0123031031032131-2123331122200230-0100031113201020-0323333120300001-0332133012111123-2120132213123200-2031211200132033-3211010123213110"></a>

<a id="canonical-0110131220100301-1010102320023010-3023311313231330-0220301230300131-1111332101021110-3223030302221020-0200111313301011-0213332123330030"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.invert_matcher` property

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

- [item](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-2032112230003320-1101112311321321-0231130232220100-0310213022230312-0120022103302100-1032320313031302-0113333120322133-0311121311123213): complete subsection reference.

<a id="canonical-3031022111130221-3300100001330000-2132300121203021-0132230201301200-2112332311111320-2330312020031303-3110023123331310-1333023323030122"></a>

<a id="canonical-2330102203011132-1222023211120110-3011331332201001-3211231021322033-0112022113012300-2001113223002321-2310323121312030-2213103222020310"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.key` property

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3321231032133332-2332011032203202-1012330112120330-0122123202131132-0121022230003033-0312212130223213-0111002303001212-0310022131300332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="canonical-2320000232203100-1223203022203202-1312221012333333-0023130233012101-0023321001111311-1122330021120132-0221222010112233-3231210000103302"></a>

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

<a id="canonical-1311122023131330-2222100033212110-0313000012330101-0032333320201000-0023230122102111-3330130030323001-3023020010320130-3312121303331001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_present

<a id="canonical-2211000221001310-2101013132123312-2022302001032010-0210103133103031-2002231110332303-0031003011312011-0100220010331213-1131113132001101"></a>

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

<a id="canonical-2032112230003320-1101112311321321-0231130232220100-0310213022230312-0120022103302100-1032320313031302-0113333120322133-0311121311123213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010)
- api_rate_limit.server_url_rules.request_matcher.query_params.item

<a id="canonical-2032131303311331-0101023203120001-3213001101113033-1320220003001220-0211233031111220-2001101231001110-1303103020110102-0310310131322230"></a>

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

<a id="canonical-1021020302032110-2002002012323323-3020332202231200-1202210331122013-3301220320331032-0312020030102330-2002131322232002-1022031212112203"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.query_params.item`

<a id="canonical-2201200112321223-0200222321303101-2311112023301103-0130033031032320-1233333030111031-2120132121231031-0030331312101133-3211213133013013"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.exact_values` property

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

<a id="canonical-0311121313001023-3220232303133032-1203023301100131-2110100211321021-2211120010213230-0021032112131013-3000110013302120-3303030320131101"></a>

<a id="canonical-1002323312200020-1323210031322013-2210130022221202-1221103210321003-2222111331012112-1310213122310020-3201202220120003-0201201112300322"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.regex_values` property

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

<a id="canonical-3233031002022320-0213001221202202-2011233313312201-1322020313100113-0033321031103102-3201103020011110-3023131222002131-0100003011133330"></a>

<a id="canonical-0021232212112323-2030023100233102-0312022120210201-3301222231100231-2132101323120030-1013233011123133-0332302203110333-3323111210001310"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.transformers` property

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- api_specification

<a id="canonical-0103332123202233-1321002331003032-0001313121330030-0110313123121212-2312131130012102-2302113002030233-2202330001303202-3131131301300132"></a>

Type: `"single"`. Computed.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_target_choice": "[\"validation_all_spec_endpoints\",\"validation_custom_list\",\"validation_disabled\"]"
}
```

OneOf alternatives in this subsection:

- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0103332123202233-1321002331003032-0001313121330030-0110313123121212-2312131130012102-2302113002030233-2202330001303202-3131131301300132)
- [disable_api_definition](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0123123210010032-2000110022103021-1311000120001310-2022020021332330-3303230013303001-1302033001330103-1133323112122130-2022201130112011)

Select alternatives according to the provider validators above.

<a id="canonical-0022010321232222-2122020123131112-0202230323021112-3100033112033232-1102311110231003-1033123112212023-1100333010210130-0220132223111232"></a>

### Direct properties for `api_specification`

- [api_definition](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0221212311312323-3110130002001100-0321312322230103-1130303111331232-2120030223003322-0011020222121321-0113022030133120-0320312233133320): complete subsection reference.

- [validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001): complete subsection reference.

- [validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000): complete subsection reference.

- [validation_disabled](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2200322100312233-1331332231131222-2201003130302321-0312322000110013-0131003103300102-0103311103220223-2320201221023111-2120100012312210): complete subsection reference.

<a id="canonical-0221212311312323-3110130002001100-0321312322230103-1130303111331232-2120030223003322-0011020222121321-0113022030133120-0320312233133320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.api_definition` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- api_specification.api_definition

<a id="canonical-2210302210310001-1212112313200131-3012330201333320-0100310110301230-3033301332320232-3203330132020010-0111231210022311-2223332212320110"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-1130302312233210-2222020023013332-0033310112200032-2200310220233300-0032303113120102-3230100112230321-2201330001112321-3120012330230023"></a>

### Direct properties for `api_specification.api_definition`

<a id="canonical-2330002002200110-1322022313322010-2121123101221102-3122020103212230-1033221021130223-2233133211123032-0010203320103012-3012102330013313"></a>

#### `api_specification.api_definition.name` property

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

<a id="canonical-2233222100310132-0221212023202002-0032031200121222-0210111120200201-2313000033233012-1302231000210320-2010300230221200-3210231300212023"></a>

<a id="canonical-3213121330222113-2120023121233200-3223321320320033-0131313221200101-2201230303123231-3100113303203030-0031003211313232-0312103021323320"></a>

#### `api_specification.api_definition.namespace` property

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

<a id="canonical-0032300300312233-0301021302333131-2220330102223131-3001132211212012-3113220030030213-0000211110133312-2332112321021122-0211332002231022"></a>

<a id="canonical-3330302002212212-2101311023332110-1001011311001131-3101012233320301-0010102200021100-2120213220102013-2032030233022323-3132123012312100"></a>

#### `api_specification.api_definition.tenant` property

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

<a id="canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- api_specification.validation_all_spec_endpoints

<a id="canonical-3001220010203320-1231332202120020-0133002130332030-2100121111302310-3210311313223123-3103313130233203-2132111013100133-0332121113013322"></a>

Type: `"single"`. Computed.

API Inventory. Settings for API Inventory validation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

<a id="canonical-2032011210011200-0022122113203030-3232322001231000-2010301323000222-1122223303303220-0232100102103312-1230201033213111-0012111300102032"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints`

- [fall_through_mode](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120): complete subsection reference.

- [settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112): complete subsection reference.

- [validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221): complete subsection reference.

<a id="canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="canonical-1301121333222130-3330221211230021-2211213103210000-2031031001130132-2301201331132023-3101133333301231-0020100220022031-2300122203033313"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

<a id="canonical-3110010302033232-0301021221111112-3022012212032111-3330322113321100-3012102000132210-3133011131032331-3322200322011200-0112023232031211"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode`

- [fall_through_mode_allow](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1202011201033211-3121333123003020-3201332032202001-2321101021322012-3312111111323101-2002003022332203-1202202103321031-0301113233233222): complete subsection reference.

- [fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332): complete subsection reference.

<a id="canonical-1202011201033211-3121333123003020-3201332032202001-2321101021322012-3312111111323101-2002003022332203-1202202103321031-0301113233233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow

<a id="canonical-0013132022213022-2222233103210031-2330133311111003-2233320212211013-3333020313133211-0033000033000330-1010030012101002-3221232200111130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for fall through mode allow.

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

<a id="canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="canonical-1333202002310022-0033330331321002-3200100021133303-2300022300200001-1211333132212001-2332003120112103-2000011033302321-0123311023213103"></a>

Type: `"single"`. Computed.

Configuration parameter for fall through mode custom.

Additional upstream details:

Define the fall through settings.

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

<a id="canonical-3031113032112321-0132220332021101-0322130001020313-2203223312302030-2300000220320030-2300203221201301-0113003123101123-1213030121233333"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom`

- [open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210): complete subsection reference.

<a id="canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-3000032313023320-3200302331002010-0201103232102211-3010311211020222-0213030310123200-1131031301013221-0212312202021223-2302022120011230"></a>

Type: `"list"`. Computed.

Custom Fall Through Rule List. Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```
