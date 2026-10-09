---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0323011220101201-2112220022310331-3302030110032002-2002132123233002-0233310301113111-3010133021232011-2332220213011233-3121021331210132"></a>

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1122231112133130-1003001300221001-2030232220002331-3022221121332101-2010211020000313-2132230300310022-0233230310113333-0010223020330003"></a>

<a id="canonical-3112010330210302-0313221023132223-0130122033101200-1023202120002003-3001200102023302-1020001013312130-2212130000210300-0321030333030230"></a>

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2032322312011002-1201210300012222-2302230020210011-3033313312000131-0010100331233002-1001003331203032-0000020000202013-0033111002122122"></a>

<a id="canonical-3221011223013323-1021220001211231-0102301031002233-0210032201211301-0331110100310321-1312132311301131-3321222211310013-0120212310310121"></a>

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1102001221022131-0132213131000330-2213121230102311-3003220222030020-2122112010010301-3101211133133231-0230303101321030-2211232002330200"></a>

<a id="canonical-0103023121313130-1013013023203100-1223303103303020-1102320332313233-1310110232112222-2111202310232122-0212221200220132-1032003323001130"></a>

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1200113333332102-3331022332002031-1120103010102301-2121331130312211-2212212010121333-2200210300321022-1003200210000330-3113200103222123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list

<a id="canonical-0202312213312000-3203203322023010-0011230313132331-3112023010323212-2113132222003010-1320121310320300-1030031333320102-1303123120302212"></a>

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

<a id="canonical-1103313100230223-3102111002100312-1310003033221001-1233101010022222-0210223103322010-0000303200201320-2111113321233321-0321231112332023"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list`

<a id="canonical-2202103202101002-3320000021333022-2213333103201221-3121120321212102-2303312203113213-3000322320103323-3231223121021133-3020303201332212"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list.invert_match` property

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

<a id="canonical-2213330313001113-0111112013131331-1111020311100030-0031030013011202-3321333333210302-2220222121121211-2330122131303212-0111201011311203"></a>

<a id="canonical-2330222020103032-0002120022221100-0312033323032200-2320321010120221-3223322121033301-1330031313003222-2111220233112220-1323121013313222"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3312120300100220-2330131103032312-0202010330213232-3020322331303310-0212132301011221-1033003031230310-3321101123123011-1122312120111220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list

<a id="canonical-0302332011311321-0221033331302301-0221021130101110-0033201112323320-3023013313331103-0333333322032211-3022120332322131-2232331331212000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000223310121233-1213232012311322-1101323223120323-2233310223102301-2222002313123210-0122133303312030-3320232102201202-1231331103201311"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list`

<a id="canonical-1133220333300300-1012001201111020-2000203021233221-3313322102331220-0221021303111203-2200331323010032-0201132211231201-2230121010020211"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3031203222211302-2303123322301120-3302130021100031-1000101233131220-3101201020012210-1301302133100302-2200220030010101-2113133013012220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-0031120021232233-0233123101022221-2112113211301033-3201302022332302-1121122132102213-0113031020213320-2001333022103103-0303230330013103"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330210233110210-1330020321223101-0100100312100112-1230203030022200-1212110001100133-3300132320133032-1120022033011222-2131312311213222"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-0320213132131201-3131223031311303-3112212312100200-3122111213112330-3231331003231131-2102213231132323-3002000223020203-0011230233121120"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1021101222131230-3302130312002103-0332220303232302-3000331102323210-1203301333212300-2012121110002130-1030023011120132-1230323230100013"></a>

<a id="canonical-1301321001010223-3030211210030312-3220321322310322-1102110231123023-0020302202303020-3120113011133310-0011322000030012-0232023330021000"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3022201232220002-1131320312131203-3022011332132331-0003213332110130-2121230200223320-1033211003011222-0100300132220302-3201333201323100"></a>

<a id="canonical-1221320033033021-0331210312201213-0300313033223323-3130330231201321-2030102320133210-0332022321103103-0330201210333023-0320131311121001"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher

<a id="canonical-1332010230112030-1013220102023001-2131300200110331-2020002313312011-1022021101231311-1232101230032222-0203330321332031-1331230121003211"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203232312002020-0320310221331103-1010212312221211-0313130112222123-1110133222003130-2120103231333123-2311010002200133-0202110001023112"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`

- [cookie_matchers](resources--http_loadbalancer--reference--group-009.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-009.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-009.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-009.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011): complete subsection reference.

<a id="canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers

<a id="canonical-3131223312031330-2131301312303203-2103203112011103-0103130213121212-3020203110122103-2210323330120212-0023211113003303-0201233200102301"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022221320003013-2311013202102213-1230300000220033-2101213022033001-0211220233030103-2302002220312123-1300221123010313-3213122101100020"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers`

- [check_not_present](resources--http_loadbalancer--reference--group-009.md#canonical-0332001211220202-2222302130003203-0200000303223110-2320212201333033-2003221000321000-1111332320123003-2223120102320032-2021301300331321): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-009.md#canonical-1013311121031001-0323310112313032-0200311210032213-2003222203010331-2122132302232200-3003203030221220-2310303333330010-2333120002303232): complete subsection reference.

<a id="canonical-0330123202222313-1212123103203130-3032023303022212-2101001233132022-2201130322210202-1012311100003211-0112211011310330-0320121322230202"></a>

<a id="canonical-2122230213320030-3002112313031232-1231110332211330-0010021213211301-1112313012122330-0210032020011320-2033203102312301-1003112003222012"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-009.md#canonical-0310011312213322-1120200011320002-1111312033130310-3102020121101100-3002323002130121-1130010313020013-3130121101032131-0020313110132231): complete subsection reference.

<a id="canonical-3022320002033011-1210210311211011-0220012000200201-1011302021021110-0231202230001111-2133313211011232-1213222033020031-0003331330010222"></a>

<a id="canonical-1133110203110023-1002333300222303-0000103323103321-2303133120102200-0032020333330300-2013111120331003-2123302032030133-0210112110223230"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0332001211220202-2222302130003203-0200000303223110-2320212201333033-2003221000321000-1111332320123003-2223120102320032-2021301300331321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-009.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2031121112020331-3302223131021322-0023321220220321-2100131103122313-3013232300101001-3122332122022333-1303103223202323-1113221312333232"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013311121031001-0323310112313032-0200311210032213-2003222203010331-2122132302232200-3003203030221220-2310303333330010-2333120002303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-009.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3101031320010310-3103310000233320-3120022022311012-0033322331002211-0203230132132112-2013333021123021-3123203000102100-1112000231102220"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310011312213322-1120200011320002-1111312033130310-3102020121101100-3002323002130121-1130010313020013-3130121101032131-0020313110132231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-009.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item

<a id="canonical-0322112323232031-0311132313001003-3031222022213013-0003021002023003-1002211230002032-0210320122023312-0000133323000221-3213200012022310"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233101322133310-2201300001132012-0322100202112002-2211202303303121-2102033002101013-1023203212211320-1013122023221032-0002203211232333"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item`

<a id="canonical-0211131030230313-2002320012220223-0130113020030010-1131003320022312-2123221021030233-2201212122331212-0232013302132022-0311202303102220"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2132310332120021-1202210100203312-0330302021221322-2213013303312323-1313012130213022-1333122012213122-2103023023313003-1123101311123303"></a>

<a id="canonical-3301211220230310-2312302100001333-2201233122310231-0331232331211030-0002032002312013-3211220030322233-2303303032010131-0301133302112202"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3023203023101113-3133111313211022-1201101232023301-2321233120232012-1113010020233013-2012310001020120-3012120021033032-2031001221300332"></a>

<a id="canonical-0232123220112311-2202331321332013-3133200130322200-3210012103310131-0231212200202302-0033001200311101-3030233011333110-3211001010013303"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers

<a id="canonical-0100022210022131-3130131232331323-2322121320013101-0303232103132330-3102201322120011-2303302320202210-3221223231202122-1230303101203233"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101132012220211-2101003013321211-1202230210012120-1011131013103322-1220331000200311-2021203330003232-3030120003011132-1322121030103121"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-009.md#canonical-1122130132031003-3223101202211302-2201131001300001-1122323213110303-0030221023012002-3321022031222123-1213320211233011-1013123122303033): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-009.md#canonical-0320220010023331-2302202111322120-0120310011031302-0232113023120201-3001011012322033-0301313233320222-3103310230023331-0323010203001222): complete subsection reference.

<a id="canonical-0212000221221212-3311231333313022-3123321331310120-3203013101333330-2020123320123010-3001223333222022-0200010220010121-3211222000003312"></a>

<a id="canonical-2100211323031123-3013100003111123-3000123033331033-1010202102133310-2021202200113231-1022013133203231-0321233220023320-2211300332121201"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-009.md#canonical-2312310310233132-3312120113120301-2102131130101312-0113220232131123-0233110022113003-2010102233201332-3031202331220220-3013210133102313): complete subsection reference.

<a id="canonical-0121100023210323-2122023002111211-0103102233031111-0213000222000332-0123232333031111-0022222223023222-3023010201032102-0301102000313003"></a>

<a id="canonical-2130223112330010-2001132020200133-1113120030101011-3213301301300300-0301211032032330-0210213213122303-3032202312130300-3331302303113212"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1122130132031003-3223101202211302-2201131001300001-1122323213110303-0030221023012002-3321022031222123-1213320211233011-1013123122303033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-009.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present

<a id="canonical-2110023000333300-2111303303020220-2012003311332112-0100323131021013-1233303112002022-0011312123233222-1312021310221330-1030302232110312"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320220010023331-2302202111322120-0120310011031302-0232113023120201-3001011012322033-0301313233320222-3103310230023331-0323010203001222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-009.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present

<a id="canonical-0230101211001311-1020011021210311-1201223130201220-0212020100010331-1312212131210120-1231233200303203-3322032220330012-0231101212023210"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2312310310233132-3312120113120301-2102131130101312-0113220232131123-0233110022113003-2010102233201332-3031202331220220-3013210133102313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-009.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item

<a id="canonical-2232311113030211-2030310331202002-0203131102122220-2030002002320103-3021122002102223-1301202233212122-1201132221203200-2023020310223133"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301120021012012-1020313000003012-1123203210313221-1322230031210222-1111012123310020-3223010332201220-0212332222112100-3133231233032301"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item`

<a id="canonical-2101222210121330-3321320211222010-0301012111212121-0323102121212032-0100032232223200-3312212032310210-3303013310213121-0030021301321013"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3131002031111121-3023100021131011-3012133331202130-2210211001202010-3311031111123023-1303022333203021-2230222003313201-2031323012323120"></a>

<a id="canonical-2211211310332120-2321110301321203-1311230133021333-3002130210210013-1023321310310103-2113213220210230-0200211103000113-1110013101022030"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0022220002020230-1003010030131333-3303110321122313-2213133011022103-2003221333303101-2020202103302021-0022330130301031-1203130102230330"></a>

<a id="canonical-3102333232120331-1231301131302132-0123331230102323-3223331332323213-2211310203110102-2211220023323210-1301131123031103-1001111311210032"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

<a id="canonical-2131222033203211-3130123231030020-0222211332333330-0010122200133101-1331332022230023-2333001203131011-3033303223200320-1110210131130032"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330032300200001-2002231213322231-0030021030102230-2112000130023222-0013100201203103-1131221031131203-0130212103111212-1033111102230313"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims`

- [check_not_present](resources--http_loadbalancer--reference--group-009.md#canonical-2030312121020031-0233023120302331-1311101330233222-2010021032222120-2101010233312112-3122310011100011-0101001332310321-2111122212212302): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-009.md#canonical-3310131023313222-0323311302301330-0330121123220032-1103002110303210-1132022103030201-1222113010022210-2003300020022001-0030231132332212): complete subsection reference.

<a id="canonical-2033320033213331-2132313032212122-0011231123030200-3122010210022023-2221203322113032-0211322231302033-0132022102300020-2133133311232320"></a>

<a id="canonical-0322120111321202-2121302011200123-0123131332021231-1200233133101300-0130011111110222-1102113021300133-0033302031012311-1112303322220013"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-009.md#canonical-0312023210333011-3220122310300130-1332132131103031-3011312321333103-3130103313320121-3331322220220322-3332331211132303-1113203221103111): complete subsection reference.

<a id="canonical-2020323023320221-1300312210331021-2122123102102202-3132220030330233-1232223303211210-3223000310130101-0030201010223033-2320313003112130"></a>

<a id="canonical-1031212221000130-0310123030200210-2333103331002002-3210010222021202-2123023023022213-1010221121111033-3302012320133131-0322020132323300"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2030312121020031-0233023120302331-1311101330233222-2010021032222120-2101010233312112-3122310011100011-0101001332310321-2111122212212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-009.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-0201321003202122-3122222320310313-1111310222320002-0031310012211013-1013321002032321-1122222133300300-0001113130031013-3011120231100022"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310131023313222-0323311302301330-0330121123220032-1103002110303210-1132022103030201-1222113010022210-2003300020022001-0030231132332212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-009.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2312301311303133-2211312233203211-3010313022023001-3121302021100211-2323033122003020-2012101303322103-2320011133203121-0100102020103101"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312023210333011-3220122310300130-1332132131103031-3011312321333103-3130103313320121-3331322220220322-3332331211132303-1113203221103111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-009.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item

<a id="canonical-3121021123020123-3210011130010302-1122120113033220-1312301031022020-2332111201220030-2121223000331212-1012233213313300-0132121301130312"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033313212210320-3203333020320132-2030010122300033-1213110333302203-0023203122322321-0031130001302001-2010212201310330-2032031300003000"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item`

<a id="canonical-2013301233100303-0321202212312013-1202223313200212-1222111221231233-3310110013213021-3331213101013320-0021311120332032-1133303022312011"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2212032231020022-2303223201233120-1031310000201211-2213322002033031-2022113313123000-1003213232222112-3010133311023011-1011203311131303"></a>

<a id="canonical-2230033021230313-0220033013321113-1030230031020103-3013020331031013-1312301130332113-3230313203221012-2020032303210032-0003331113313121"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3010001313020230-0221212322202202-3010122100322121-1023112131312231-2311003203201013-2101231232310303-3000102301030302-3330021103030330"></a>

<a id="canonical-0332010231333211-0222222232323201-0230211110323133-1033203020300233-3033330021321221-0001300001230310-1011101031022131-0121113021123213"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params

<a id="canonical-0230100302023013-2001113023300110-1232120030000011-2001231222132202-0233333320233231-2023000211212022-0333210033200211-1221213002030203"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310010123120233-3023330101330132-0111331213223032-3311202203113023-2000300003202100-2231031203103010-3330313123203300-2311132311210120"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params`

- [check_not_present](resources--http_loadbalancer--reference--group-009.md#canonical-2133331220212010-1131020000300212-0000202201222330-2102033122223221-3002330212312012-0110323033001201-2323121022032323-2121221021033131): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-009.md#canonical-1030312120322111-0033230321102320-1211023133031131-2303322232000221-3210112012331103-2210022021220122-0302033212333022-2010323221130123): complete subsection reference.

<a id="canonical-3203003323010322-1020002033023023-2321322202113113-3132100022200211-1321211123002000-3101111110133110-1111222031022300-1211310133320010"></a>

<a id="canonical-1332231333212033-2001302303103330-2020030131130021-1000322021113331-0213223123013001-0303311133113323-0203033000112011-2330033110300222"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-009.md#canonical-3202300301312031-0020030211222211-0313121300121311-2031212131321320-1102131113201100-3312312301133303-3130013111111322-1113322333210333): complete subsection reference.

<a id="canonical-2032101200220213-2122331222030110-1123103312122312-1333223010320302-3211102230312002-2212102212312202-3211033301232031-2133220300000023"></a>

<a id="canonical-0121300013213013-0011113230031110-0100301211121223-0020323232002300-0320123003002232-1211213120331231-2002331012023103-2220033021021313"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.key` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2133331220212010-1131020000300212-0000202201222330-2102033122223221-3002330212312012-0110323033001201-2323121022032323-2121221021033131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-009.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present

<a id="canonical-3302111023300033-0212112200122102-2331123213320131-1211020210333302-1022122103313112-3230133020301312-2130201310233312-2002010033201122"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030312120322111-0033230321102320-1211023133031131-2303322232000221-3210112012331103-2210022021220122-0302033212333022-2010323221130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-009.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present

<a id="canonical-3111020013202112-2320301232313300-1200333333023323-3022000210133201-0110212233122203-2010113333211221-2110333211103223-2232330120300320"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202300301312031-0020030211222211-0313121300121311-2031212131321320-1102131113201100-3312312301133303-3130013111111322-1113322333210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-009.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item

<a id="canonical-0222223133123003-1130123233020020-0000023113032012-3200313130032333-0111201113110210-2122212011132210-1121311003111010-1310032001200111"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3121031322321202-2020022110230321-3323120220132322-2300003030213010-3303330121203212-3233223002010012-1320101221101203-3132333211200310"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item`

<a id="canonical-0230121300110213-3132030212223321-1303011002120332-1000321232312021-0013102122030120-1231031100020223-1012222021033202-0303212031012103"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0203301102003303-2222233313330110-2121020303100303-3022031231132300-2233222131101121-3333320232033210-0302001130321323-1013210112202012"></a>

<a id="canonical-2110311333311323-0102033303323233-3200332323202023-3010020003332231-1013110000223032-1101213303023103-1032331021232033-0300200012333030"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3123010223133230-1300122323133003-0013230020002003-0023233312300103-2031102121030031-1320122032010100-1201131211333310-0112332100120122"></a>

<a id="canonical-0301101032323210-1320003021232211-0113232110321330-0302010333322201-1201213001210220-2321003331313230-2000232223010012-0102322123133201"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2313230220001103-3212301300010212-3013123133132033-1130203100112210-1121313032010300-0313200120132330-2230131221332112-1011333232233322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.custom_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.custom_ip_allowed_list

<a id="canonical-0223321212212223-3323033301230012-1310121200013311-0223303033003103-0110303301020330-1213220233123013-2211330312003211-1321022110131220"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200320100220220-3000020002333123-0132311322310213-3023310102223002-1020112202001332-2023201203201221-3302322111021111-2203222133120231"></a>

### Direct properties for `api_rate_limit.custom_ip_allowed_list`

- [rate_limiter_allowed_prefixes](resources--http_loadbalancer--reference--group-009.md#canonical-1010211010113121-0331310121113033-2030031032022330-1021322001300232-3231010303030310-2212033322222311-3023322011123320-0021202201302200): complete subsection reference.

<a id="canonical-1010211010113121-0331310121113033-2030031032022330-1021322001300232-3231010303030310-2212033322222311-3023322011123320-0021202201302200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-009.md#canonical-2313230220001103-3212301300010212-3013123133132033-1130203100112210-1121313032010300-0313200120132330-2230131221332112-1011333232233322)
- api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-3312110313030210-2131230011132231-1010202033133332-2100210220332002-1030303211221312-3131213103012332-1332122223133213-3111011131200210"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303103133331003-2133200100121031-0002330000023200-2112110331233230-2011012221202320-3123311201102021-1223101003010222-2003332032332033"></a>

### Direct properties for `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes`

<a id="canonical-2322113020233302-1013001312301010-2113201112300003-3203222031231200-1322011123103303-2131223202002312-1023332313300013-3121021301010222"></a>

#### `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2010002123333030-2112221213223230-3301202303322002-3202012203231101-0301230203330010-1130110122130113-1133213201211232-0100321010101033"></a>

<a id="canonical-1221012113222303-3201020122310122-3333102002321300-2303303323120221-1031311000013001-3120033023031112-2031132100013322-2200032030320031"></a>

#### `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2003321330223331-1021132032302121-0212011123323201-0301312210112130-2223032022102233-0103120320131011-1233100223022130-0021131220121023"></a>

<a id="canonical-3122103133201110-3210122001202300-2321200032112321-2022211323123200-1330210023330131-1200300213122123-0232222220102110-0203330122211312"></a>

#### `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1201210302313203-2211300232121320-0323321131002323-0331122322021300-2230032123313000-3330020132002300-3030111322130233-3130310012323300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.ip_allowed_list

<a id="canonical-3012333312322230-1212020221132021-3111130330021031-2202011012313020-0113002210323312-0332123301131333-3223100320121322-0320222232023001"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031302332013223-3020113112310220-2112211011333322-2031221020230233-1200212113131013-0102011211101103-0021130332223120-0312011101300302"></a>

### Direct properties for `api_rate_limit.ip_allowed_list`

<a id="canonical-3132130030310123-2122120100331231-0012220001232330-2003022220321111-2010222123112213-1211203331333302-0232313102301213-1210200030330213"></a>

#### `api_rate_limit.ip_allowed_list.prefixes` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1030302110202203-3022230330030230-2123232230310122-2013030321120121-3102121110303121-1120312201020213-1322321131223223-1130233033033321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.no_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.no_ip_allowed_list

<a id="canonical-2001212101100010-1031331233322320-3300321322323011-1313221031220113-1111030210022001-0302201211302001-3330310202332003-0222110203332211"></a>

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
no_ip_allowed_list = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.server_url_rules

<a id="canonical-1022322002012112-1333230230322312-3130222323122203-1230312210223231-0311012301032012-0233333131213112-2003111123210302-2110230030012100"></a>

Type: `"object"`. list nested block, Optional.

Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one
rate\_limiter\_choice: inline\_rate\_limiter or ref\_rate\_limiter.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
server_url_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112131322133112-0311110030232200-1021300200112010-2330333122100033-3031312222202220-3022030033223130-2132011301000303-1220331020023103"></a>

### Direct properties for `api_rate_limit.server_url_rules`

- [any_domain](resources--http_loadbalancer--reference--group-009.md#canonical-0212023132210021-3231013310023330-3332013202121221-0002003101003132-3121122302111200-3213110110213330-0033110231020100-3113033032103031): complete subsection reference.

<a id="canonical-1020332131331132-2130131300220002-1101010231213131-3310130032300300-0323220100032300-1313130213020322-1323302210203012-0031102311221132"></a>

<a id="canonical-2102221301233332-3122111213333202-1301333233031233-2111301223312233-2300023220300223-0112103100313132-3331333130233303-1010031203023010"></a>

#### `api_rate_limit.server_url_rules.api_group` property

Type: `"string"`. Optional.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Additional upstream details:

Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3313310122333011-3123110332133223-2120011013302313-0322101133033012-3311101223103210-0232223202233321-0003213132330233-1233102131312033"></a>

<a id="canonical-0310203201000021-1213211033312020-3301322020001121-2012013213032102-3221101113112011-1332102311331001-2033233132301313-3301200032002312"></a>

#### `api_rate_limit.server_url_rules.base_path` property

Type: `"string"`. Optional.

Base Path. Prefix of the request path.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311): complete subsection reference.

- [inline_rate_limiter](resources--http_loadbalancer--reference--group-010.md#canonical-3230011100300201-3021322232013132-2120130113220102-3000312330300303-2132212200020332-2301101031230211-2123121301333223-1211301231000212): complete subsection reference.

- [ref_rate_limiter](resources--http_loadbalancer--reference--group-010.md#canonical-1003023331222033-2221012320323333-0303300312100321-0001000033001020-3133220003001112-2131311032211230-1003011113023110-0101320100203112): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031): complete subsection reference.

<a id="canonical-3123333301133020-2200200331133310-1311231301003320-2131122103021110-3333210222022322-1122103200201113-1303003120101101-1203120330230103"></a>

<a id="canonical-2233202221011322-2102200013200112-0323012231111112-2010031032003213-2011023111300112-1013111223110233-3323322232010310-2310321333231130"></a>

#### `api_rate_limit.server_url_rules.specific_domain` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-0212023132210021-3231013310023330-3332013202121221-0002003101003132-3121122302111200-3213110110213330-0033110231020100-3113033032103031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- api_rate_limit.server_url_rules.any_domain

<a id="canonical-3203321320012003-0011220012131130-0320133200230300-2002312213101131-2231110302021100-0213113302100120-1311033322302232-2120300033223120"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- api_rate_limit.server_url_rules.client_matcher

<a id="canonical-1213112212110323-3222300213131031-0221222122121033-0112211103103323-0320120313200300-3330311131333331-1211102100003010-3002213121331023"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312332203323103-0332211132333313-3332102231123213-3201300332223031-3210202311012210-0103132102010221-1331333000301313-2031100203122203"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher`

- [any_client](resources--http_loadbalancer--reference--group-010.md#canonical-3201313120222312-3232003221302110-0311130020311312-0301122020312122-3101232211013300-2031210131210003-3200321013030133-0011010102222102): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-010.md#canonical-0022011202103021-0333123310221021-2132010333112323-3312330330102330-3013211101220331-1222231010222202-2200120320010011-0210121103221302): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-010.md#canonical-1133033031101131-3302230123123022-1122323011022101-0320331123222100-0301010130230202-3013302033301123-3330213033221301-0203302213331121): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-2000002131130203-0313323223312023-2132220302130203-3322322103201031-1223203023121310-2021010311321311-0130130232033020-2200211201222213): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-010.md#canonical-3221232222331013-2013333101232303-1211121213320301-0100113132313033-2331203203133003-2031021202220201-1321001103020332-3123212023212212): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-1201312112023220-3110102133111300-3200210011301312-1102221323231012-3201330200312020-0221113222213323-0212313312200010-1101310120100232): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-010.md#canonical-2203121131301233-2203201100322000-2220212301120232-2023030221110012-2001211200121321-1311112302313220-1333303302223032-3033032303332202): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-010.md#canonical-1231330211322320-1122320030111011-1330301312012202-2310122321121110-0212323111010212-3313202200230031-0222223021333031-2130231112220022): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-2111102223312220-3103020330310103-3100313013131110-0322313121022222-0030312201003130-0023132031031331-1201122310213220-3210012002132112): complete subsection reference.
