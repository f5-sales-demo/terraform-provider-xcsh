---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-1121321233302210-2201233030311302-0233012110112111-3301022303130020-3330230211332201-1122013113002133-3100021113102003-1302122332233003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.ip_prefix_list

<a id="canonical-2331230113331202-0322001220233032-1100221123320311-3021001302112131-1203113033200222-2012211303003130-2020101132130300-0311012312220230"></a>

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

<a id="canonical-1120110333210222-0332210123320230-3313311113311032-2310330332130220-1121201310212101-0201003221222331-0210123120002311-3323133312111321"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list`

<a id="canonical-3113331313220101-2331330033023320-3010103032031011-1332002231302203-3011330030323102-3003020013232031-2031300130020310-0001130101320101"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list.invert_match` property

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

<a id="canonical-3232112202210200-1100032001303302-2232320012101323-2212231002032123-1221132211210020-3013022110320323-2210323312203221-3122230230033311"></a>

<a id="canonical-1020221202303220-1232322313331320-1130110001130121-3321323203120203-0013212023020333-2010303322223102-3302122103313230-1000133213020202"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list.ip_prefixes` property

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

<a id="canonical-3000100122201321-1031000322331033-0320100200312313-0310123303011001-0231212022301023-2030131301002101-0231000210030100-2010003030323113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-3100330322300200-1331102120020231-0220330232301003-0313300300130023-2212200103120302-3001021301231221-2212230110112332-0310333103203003"></a>

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

<a id="canonical-3220130222122332-2103321002303133-3102111102203210-3213322010330120-1023021230111320-0201130121022232-0101312231010310-1302303333332233"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list`

<a id="canonical-2100202112303322-1022020213112022-1003103301132313-0331310102130231-2202133110110223-3001310110321032-2132011331211220-1010303103330232"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

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

<a id="canonical-0130003213232033-2202300221200002-2331102232130333-0210030322312120-3203313330213022-0310130203033112-1123331112033221-0112233032020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1310201021312320-0113202323013002-3130321313012030-3301211231221331-2003311330120133-3222222113203213-3321210102021302-3301011103013002"></a>

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

<a id="canonical-2213212211312012-0020312302200302-1030311201220011-0011230213300200-1101222223133120-3223233321232132-3000310103332010-1102212223002230"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-3203012101201101-3311213222002223-2332220001020010-1303122223120133-2033032101020032-0223030310131232-3113030310021210-0323000112311123"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.classes` property

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

<a id="canonical-3323010032322312-2211331320121303-0033312111301333-2310230003123130-1211103200202021-1210202210102323-1332101003001111-0311330023133011"></a>

<a id="canonical-0003022033320311-2113331221103201-2131012003021012-2122130011132211-3103001231201220-1313201033111021-0320303130011032-0100330131210120"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

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

<a id="canonical-0013121000303010-3220021200023022-1230331032221011-0201222311221102-0321030230130302-0310201303201321-0003012323102023-3133101202302132"></a>

<a id="canonical-0333133110210031-3032123002333010-3211202122132032-0231313230300200-2101002110332300-0010131022300310-0021231111203133-0220002030102101"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

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

<a id="canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-2112001120300200-0213110213211102-2232121213220101-3130111122212112-3021133133212111-0123323133032120-0120122200022121-0321200311032112"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222320320100022-0220311000321130-0312213233102103-0212200210000111-3223030221120202-1121301101013222-0133221023213210-2220322323110001"></a>

### Direct properties for `api_rate_limit.server_url_rules.inline_rate_limiter`

- [ref_user_id](resources--cdn_loadbalancer--reference--group-006.md#canonical-0013311012330202-0121212300112232-0023221120001330-0102000030031120-3321233120230310-2003230003222132-1020133112230313-2220300233121113): complete subsection reference.

<a id="canonical-3101110210003223-0233211032303021-0013321113303300-1202210220023000-2000302331313302-2332331311032323-2003211210113012-1103230101010110"></a>

<a id="canonical-2222103100102322-2103010311230321-2110311203131231-0111310213030313-3232013030202223-2102032102023310-1230212232121023-1233303103232023"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.threshold` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1131032333200223-2101000311023221-3121312300232221-3313113021113000-2323030201312213-0113231210233033-0233201203113010-0001301300000013"></a>

<a id="canonical-0310300201010210-1022130203000310-0021131030231313-2231013210232331-3031023320033201-0033313023332003-3322212113332112-1011320320322021"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.unit` property

Type: `"string"`. Optional.

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

- [use_http_lb_user_id](resources--cdn_loadbalancer--reference--group-006.md#canonical-2230212032311201-2231312302320020-3212020200322021-0003131301030311-3323121300123202-3022102231011330-1031030320220021-1212332030021031): complete subsection reference.

<a id="canonical-0013311012330202-0121212300112232-0023221120001330-0102000030031120-3321233120230310-2003230003222132-1020133112230313-2220300233121113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-006.md#canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-2221100320102101-2332302121002013-2031011233230131-0210222330203303-3102102103311033-2010321020322101-3311332322002220-0020101222220000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ref_user_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212332111133030-3021110311130222-0331031200303010-3320032111213320-3003210220321123-1313113120133210-2333133201021033-1222030203033223"></a>

### Direct properties for `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id`

<a id="canonical-1331313131200122-3120223221230121-3023303230012233-3233113332331322-2120032113321311-3313232331031220-0221110330231133-0233131011301210"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.name` property

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

<a id="canonical-1313212020221323-0300033320200012-0123031323123133-1303013311321231-0331101130033123-0113002201320323-0333031200010132-1302032323222002"></a>

<a id="canonical-1123333123011200-3200011301110112-1033002002333103-2201212221220313-2112112200322320-0102231231013123-3132103313310122-2300330203321122"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.namespace` property

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

<a id="canonical-0232132330300231-0111032320011321-2000013212132010-2230100000011323-1132130320111112-2012023310230110-0222120231231020-1312333013112012"></a>

<a id="canonical-3320323122221231-2313112122010231-1300220332033003-2022212210101223-3030000130310331-3313220012103231-1111123131221322-1131302003103233"></a>

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

<a id="canonical-2230212032311201-2231312302320020-3212020200322021-0003131301030311-3323121300123202-3022102231011330-1031030320220021-1212332030021031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-006.md#canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-1123112033221002-0231000020012233-1321001311200030-1312102331121120-0211303221233112-0310011012131132-1000022312010311-3031223201010331"></a>

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
use_http_lb_user_id = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210111123202013-2212101212113221-1330130203301223-2331201130033000-0332303011120301-2031002033003132-3203032332303101-2213233012110233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.ref_rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-0133331221033120-3233110330132200-3212312233133233-0010301332311201-2202031302133023-3321323221000000-3303321133030130-2312121000012211"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ref_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101322130203111-1310010102312003-0210003033101212-0201020011231100-0020133301232101-0330201321312012-0231200313203301-0200110011320311"></a>

### Direct properties for `api_rate_limit.server_url_rules.ref_rate_limiter`

<a id="canonical-3220303000310130-2113300121023031-3021011010133132-1201002210112230-2331033030131022-3331103110210123-3110002333003320-1322112103123133"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.name` property

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

<a id="canonical-0022000301213133-3131320033000011-0133100101123313-2322022132032131-2223230101103302-1020300310023313-2310030230210002-0000010131323221"></a>

<a id="canonical-2332101110033203-1300120231112230-0031201031030212-1213033110121001-2100201201103231-3202111100332123-2013222200232321-0230311102032120"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.namespace` property

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

<a id="canonical-3131101233001030-3321310132330203-2123120313210230-2020011113223323-2123100211101112-3221203221011230-3231022120313012-1302111112033023"></a>

<a id="canonical-3312001211131301-1303011132203331-1323231023321303-2332230303103333-1230233322021033-2232301030212121-1202022133222311-3001032113322022"></a>

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

<a id="canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-0011033133021301-2313223002222133-2222211103133300-2201212331231002-0020103000321310-1221311022111220-0310300313120013-2000012312013102"></a>

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

<a id="canonical-3212010010100023-0332312230211111-0332132021313201-1102330202122113-0210233132233203-3011103220223213-3231301212010300-3131130003222313"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher`

- [cookie_matchers](resources--cdn_loadbalancer--reference--group-006.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030): complete subsection reference.

- [headers](resources--cdn_loadbalancer--reference--group-006.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323): complete subsection reference.

- [jwt_claims](resources--cdn_loadbalancer--reference--group-006.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320): complete subsection reference.

- [query_params](resources--cdn_loadbalancer--reference--group-006.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231): complete subsection reference.

<a id="canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-0101000130230322-2011232033211002-0003000213202020-2312130230100211-2213301330013323-2123301023312131-3202301332100011-0033010020003330"></a>

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

<a id="canonical-1030222002331123-1221330333200122-2302331232132210-0032032213222000-2210323332023203-3211300100212003-2311021302113210-0323303203123102"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.cookie_matchers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-0022023013102010-2123110121321310-0012323310123032-1133003210323313-2002012113210013-1103120130131013-1022022322311101-0313121302122002): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-1120300322323132-1200130313310131-0002203230003120-3003321321113332-2222331133131013-3230333133133113-0122033123232211-2031231022331323): complete subsection reference.

<a id="canonical-1112200200312101-1110130302231220-2222323232222101-2303032302221021-3333033210122133-2022320001122201-1210121032122131-2113201002222102"></a>

<a id="canonical-1223322231020201-1100023132300203-1321122322032230-0003120131320333-2002022210102303-1122312020302311-2002301320012302-0123102312232102"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.invert_matcher` property

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

- [item](resources--cdn_loadbalancer--reference--group-006.md#canonical-2200003022203033-3232100113133101-3330303230031122-3310313132032300-1302323220000332-2100013130213202-1020331333022300-3221312111301310): complete subsection reference.

<a id="canonical-3333011123232033-1113031111122313-2000022130321320-0332313303321110-0302003002211211-3112221331132202-2110122330210303-1230221203022102"></a>

<a id="canonical-2012321100033122-2213011003232213-2231203111013222-1022130331322231-3133112020031121-2031300021030222-0202302202222010-0230220120122320"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.name` property

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

<a id="canonical-0022023013102010-2123110121321310-0012323310123032-1133003210323313-2002012113210013-1103120130131013-1022022322311101-0313121302122002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-006.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2221210220332301-0132213302302221-0321121032111120-0323121203103301-2202313022220120-0033203221102333-1210021233111330-1101332333301310"></a>

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

<a id="canonical-1120300322323132-1200130313310131-0002203230003120-3003321321113332-2222331133131013-3230333133133113-0122033123232211-2031231022331323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-006.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-1212323132202333-3100231123220010-3221133222020002-1113231032331223-0212001003010231-3300301102120313-1033101201303211-3331112210001302"></a>

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

<a id="canonical-2200003022203033-3232100113133101-3330303230031122-3310313132032300-1302323220000332-2100013130213202-1020331333022300-3221312111301310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-006.md#canonical-3033211112003230-2302031233322223-3011210231211232-3130203102313021-2123313232211002-0033322022102113-0323310000003111-2330301133031030)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-1221033222211231-2200220211030101-2131030103002200-1222101223133313-0110101121231303-3321310110032300-1121122211110310-3002002210312003"></a>

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

<a id="canonical-0032200321210012-0112030300313000-1210211020212012-3230203230311122-0123103231002201-0132023103130202-2032102010100032-0200121222122023"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item`

<a id="canonical-3332322123201122-1212103112233210-0111323302033203-0233023302313332-3013022230330011-1303020320103202-1223302110120010-2313101032231122"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.exact_values` property

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

<a id="canonical-0202203113130203-0132320302012012-0230210120023011-3211103000002021-3111122303301101-1201330233013020-2021321331212130-0032200221011133"></a>

<a id="canonical-0220330120232233-3120000033013203-0312113320333201-3111012130213232-2221012000121303-2323220231322032-1110002022121332-2313030021313320"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.regex_values` property

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

<a id="canonical-3212000331131221-2222003010021132-3130133212133013-2021110201321023-0103320222311121-2222211002213101-1233031003302010-3203310132201031"></a>

<a id="canonical-2020111230310000-3311112213320130-0233313101223022-2133311122323301-2212133032020110-2320203133100320-0022330120010203-0003331122303011"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.transformers` property

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

<a id="canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-1331322133020001-2023222003233133-2323020002012031-2213132122033232-2012032322021212-2232231101003330-0302031011300010-3333121120202033"></a>

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

<a id="canonical-3130320231032101-3103032122001331-2001023133131030-0131232032303321-1312213230312130-2130332023321310-2232133202231302-0223311222212011"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.headers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-3313133023110133-3231332201323201-0232132312101311-3230130032312001-1212102113222112-3202002113133022-0032222013230302-2122333100023121): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-1331001203220231-1210232321130113-0002023333310220-1120311111210203-3123123120321320-3330133133100323-0233320030311210-1120022231031332): complete subsection reference.

<a id="canonical-0223302231320013-3222103102122030-2211310321022223-1001201220022012-1110323220030231-3011201301100002-0223103011310102-0112331210113201"></a>

<a id="canonical-2322120222100330-0330222002023001-1110330101102323-0323020320221032-1103323221233312-1213312113312020-2013332113021123-3100310123332220"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.invert_matcher` property

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

- [item](resources--cdn_loadbalancer--reference--group-006.md#canonical-0221231223321303-3120300201222020-1120200011202210-3211331311120130-0220332033231000-1011003020230110-0311133003003031-3230012101233210): complete subsection reference.

<a id="canonical-0032203303231223-0303121210231230-1132011333221301-0002123111132023-1330200221002132-0323002200132303-0131132022232113-1132201032230313"></a>

<a id="canonical-0120201012200232-1103021132010202-2202232002320120-0302322212001333-1301000330012021-1300110102220333-1130212330321012-0102132000000220"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.name` property

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

<a id="canonical-3313133023110133-3231332201323201-0232132312101311-3230130032312001-1212102113222112-3202002113133022-0032222013230302-2122333100023121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-006.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-3130030012231213-1303323332300120-2300122330301002-2010300230202330-0310323311223123-1302320002221012-0232112022200321-0303112323320321"></a>

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

<a id="canonical-1331001203220231-1210232321130113-0002023333310220-1120311111210203-3123123120321320-3330133133100323-0233320030311210-1120022231031332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-006.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-1333313223232333-2220000213213003-3230020333320332-0032312120101311-2333112202110321-0311332131210013-1312221033113332-2301030203231102"></a>

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

<a id="canonical-0221231223321303-3120300201222020-1120200011202210-3211331311120130-0220332033231000-1011003020230110-0311133003003031-3230012101233210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-006.md#canonical-0002122103212131-2020122213023330-0233012102211313-2223031022231311-3110202023102222-1303322023020132-3300313112001132-0021201011102323)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-3112002101313321-2112203230130322-1211232321021201-1113113220120331-0233132133031200-2220302302232000-3022111032013232-1313310033221222"></a>

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

<a id="canonical-3122033113202203-0322100003133121-0220110113301312-3223201111133220-3002302212210233-2022333122233222-3001321202001021-1310331300000120"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.headers.item`

<a id="canonical-0021120303301111-3321230013312221-3333002330031320-3000111201110121-1321313123333132-1010123112202003-0330332222111211-2322202310112123"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.exact_values` property

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

<a id="canonical-2232332231032321-3330222201301303-2113202300231301-1222003030321100-1032010130003001-1132112010201302-2021003301010100-1001313231231330"></a>

<a id="canonical-3313232003311221-0013101301201200-2112302011331201-3121011200102030-2002032212100201-2210320303330011-1132230121031222-0123322322032333"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.regex_values` property

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

<a id="canonical-0203010333012201-1320211020310103-3131012111122033-0202301103113023-1333313220310320-2232312033331313-3122030300301222-1132203021323330"></a>

<a id="canonical-2320120213002102-0123122013010011-1333220032332112-3132103220303003-3202121200022301-1221103102022221-0100023213222033-3113023210213333"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.transformers` property

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

<a id="canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-3212312113023033-2033123110232322-3113330222122021-2133230113312131-2013113100331130-2130132033012311-2111322322322132-3232120322310311"></a>

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

<a id="canonical-3021101030032231-2031103133210100-1300332321211321-0132203022133230-1121320330223332-0203323110212103-1031331302122311-1233133212130031"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.jwt_claims`

- [check_not_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-1110323000122121-3303301302312310-1320200003111323-3032011220203120-2131033201113102-2203313202211313-1132323001133122-1101331012201031): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-2322331120020121-0002113312012132-0111202000311011-2100002203231213-3200312100012310-3222010320113103-3313302201132000-3321102322211131): complete subsection reference.

<a id="canonical-0030311220322001-2002010011010112-2030220020200032-3322230202000003-1332331110213131-0313201230132032-0000301110202012-0121010112232301"></a>

<a id="canonical-2033212100030022-3102222302123031-3113102201101122-1012002223332231-0320202103102313-1201131333312020-1120300331322210-0303231331023220"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.invert_matcher` property

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

- [item](resources--cdn_loadbalancer--reference--group-006.md#canonical-0003022123333033-0123121313130321-2121132222111112-0130011033022321-2232120020210000-0221331031033030-0233013312030201-2011300220130001): complete subsection reference.

<a id="canonical-3002133203310330-3023102321102112-1113123113233013-0000003012322331-1110032031332201-0303223133111203-3131232101130202-0232121133001222"></a>

<a id="canonical-1201320312302021-0130232310200301-0023111001121030-1302033010201033-1303112002131111-0013311133231022-3112131331002000-3012330121311123"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.name` property

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

<a id="canonical-1110323000122121-3303301302312310-1320200003111323-3032011220203120-2131033201113102-2203313202211313-1132323001133122-1101331012201031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-006.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-1211101130033310-1313033332132001-3202330020130010-3210101123230132-0033313020012002-1032011112301230-0310121200201102-1132102321131003"></a>

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

<a id="canonical-2322331120020121-0002113312012132-0111202000311011-2100002203231213-3200312100012310-3222010320113103-3313302201132000-3321102322211131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-006.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2232303200223332-1232123112231210-1222323330102001-0212113332320200-3032302231131122-0011211030023233-2111003302020331-1120010132131022"></a>

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

<a id="canonical-0003022123333033-0123121313130321-2121132222111112-0130011033022321-2232120020210000-0221331031033030-0233013312030201-2011300220130001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-006.md#canonical-2120002012323320-0210320131230122-1101202032031212-0321112122211211-2002131210323022-3122021323013001-1311113321020303-2332232002032320)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.item

<a id="canonical-2121231321210132-2203000303010210-2020300212022333-1002122211333233-0300123232010313-0311311030013211-3113210023010323-2210111100003300"></a>

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

<a id="canonical-2223011333321302-1030203122121202-3113102111010211-1030113233201221-0222033321111122-1011101031020210-3200212302132033-2113002101202000"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item`

<a id="canonical-2033301130312200-2332222203233331-1211131212120323-0203021013202221-2222231221100000-3202012200222020-2210101301200303-1022331312110123"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.exact_values` property

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

<a id="canonical-3032112131111232-3000012023033001-1020001232102223-2011010130111322-2031020212102131-1200220221201302-1200000132233102-3310232132002323"></a>

<a id="canonical-2012000023310003-1003232330021212-3013210310310320-3302011211033211-3013232323122023-3111112033012133-1301301321210303-3110200210133032"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.regex_values` property

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

<a id="canonical-2103020130003303-1310020221031202-3120010313003320-2320323331100123-0120331121203011-3333013312210020-1220202222323323-0031102323033213"></a>

<a id="canonical-2303203011300333-0232111101332020-1122233020331032-1021133232200222-0312233120310132-2203120031031123-2211311131013200-3132333333323133"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.transformers` property

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

<a id="canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="canonical-1231321222222110-0033113223323002-1010111300220300-2311321231203323-0011103330120103-2003220003003202-0132332113321123-0102212330123003"></a>

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

<a id="canonical-1003121030032031-3132021122021303-0323020332113121-2331222033101303-3332012031023103-3100300223322322-0112333201123221-2222121323301003"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.query_params`

- [check_not_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-0031211030333021-3121133300032202-2222032100113212-0222322311221311-3200013103011122-2032001331333233-1121130331332110-2130300203113110): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-006.md#canonical-3312033211310313-3303303032031201-0313212311020003-1002300120123310-3102100123213103-2210322300013230-0023122312323202-2033033022000301): complete subsection reference.

<a id="canonical-3001121131031030-1210110313021110-0001101202111212-2021303200301221-3203132033032103-0122003112113130-3310200121002003-0313310210311130"></a>

<a id="canonical-2312300133230322-2333121021300202-3032123023213122-3121010013112112-0021313012022110-2001012213313121-3200121021323310-1222012100123013"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.invert_matcher` property

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

- [item](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022303020223232-1111323302302211-3023001310232001-1301133223032302-2031221220101030-3011311302103021-3213330300230233-0302101230233121): complete subsection reference.

<a id="canonical-0202211210031202-0333122030200020-1113311312023111-1200320211101313-2220002033232233-3132003303203231-0301102310033003-3302031030100020"></a>

<a id="canonical-3103120121132101-2120033123312133-0123222011133223-1013031233220313-1300113003212002-3301211113132232-0233033021222013-1330222202321122"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.key` property

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

<a id="canonical-0031211030333021-3121133300032202-2222032100113212-0222322311221311-3200013103011122-2032001331333233-1121130331332110-2130300203113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-006.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="canonical-0012121133222020-1030131122232111-1233302033330021-3230323321023021-0200132320213100-0211123201223001-0330111223232103-0110023212113231"></a>

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

<a id="canonical-3312033211310313-3303303032031201-0313212311020003-1002300120123310-3102100123213103-2210322300013230-0023122312323202-2033033022000301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-006.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_present

<a id="canonical-0201200311312110-1212232202311321-0212020031330121-3202010300333132-2332021223320111-1022232222013110-3333232210023002-0310332111210302"></a>

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

<a id="canonical-3022303020223232-1111323302302211-3023001310232001-1301133223032302-2031221220101030-3011311302103021-3213330300230233-0302101230233121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-006.md#canonical-0122313023112111-2300122102223032-0012011121332013-0133313131303331-3101132233230020-2203113102113133-3002130201213302-3001033321123231)
- api_rate_limit.server_url_rules.request_matcher.query_params.item

<a id="canonical-1211321221121011-3111113022112222-3321233101000030-2131103233321302-1333200102022230-0203323112310220-2031210012123232-2202132322222232"></a>

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

<a id="canonical-3032203303020112-3032310001220322-2223310303110131-2131033011300111-2031332311033320-2222100123211033-2211033201023102-3210100103230113"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.query_params.item`

<a id="canonical-3301123011303113-3021030010102310-2222131101101210-3100001111313100-3301023303130030-3222332201130332-1223022020133320-2112103000233033"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.exact_values` property

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

<a id="canonical-2121221113022030-1101131012330010-0023200330200332-3132312310000023-1221221313113002-0001310132113000-1212111031001031-0120112101012203"></a>

<a id="canonical-3032313103030201-0213233331103302-2200233300320100-3021132312102122-3003111021031122-0032300103111203-2323031032131023-3132211202130112"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.regex_values` property

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

<a id="canonical-2213203210333323-2130113302001000-0212231201232203-1023011312330321-2130033032002021-0312102332230211-0123131033011121-1120303000103310"></a>

<a id="canonical-1303303232312323-0123032322311001-2133121121201320-3002220203032222-3003102203010310-1021223130020012-1121123221013221-3233130110022320"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.transformers` property

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

<a id="canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- api_specification

<a id="canonical-0102031021120331-0211222232131320-1022002322130010-0221000333310320-3012323211302110-0213323320200103-0202010323332222-0330001323100020"></a>

Type: `"object"`. single nested block, Optional.

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

- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-0102031021120331-0211222232131320-1022002322130010-0221000333310320-3012323211302110-0213323320200103-0202010323332222-0330001323100020)
- [disable_api_definition](resources--cdn_loadbalancer--reference--group-010.md#canonical-1013231301331321-3201321033031021-0303223333210301-2112110233330302-0302100120123301-0132331033332233-0312121102001302-1222030220330231)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_specification {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221000001313202-1033302122333012-2301030213020202-1123103023112120-3030003201212031-2002112201322021-3212231122031111-3211021332111113"></a>

### Direct properties for `api_specification`

- [api_definition](resources--cdn_loadbalancer--reference--group-006.md#canonical-3332331233113011-1302210123002122-1102002302230210-0222311320311221-0202010322311020-0233111223111000-2110301333002222-0313013021210020): complete subsection reference.

- [validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333): complete subsection reference.

- [validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133): complete subsection reference.

- [validation_disabled](resources--cdn_loadbalancer--reference--group-007.md#canonical-3201130202332300-2123131032222110-0303233101102201-2330112312130010-0032100122031021-0111232221032121-0030201333002210-1203121111323232): complete subsection reference.

<a id="canonical-3332331233113011-1302210123002122-1102002302230210-0222311320311221-0202010322311020-0233111223111000-2110301333002222-0313013021210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.api_definition` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- api_specification.api_definition

<a id="canonical-0132023233102013-3301203333023110-1023321032301230-3033120000333332-1212220201310133-0330322333220231-3223221302131210-1230320020313321"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
api_definition {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333303212002302-2321203231022011-1103210012011322-1031220202313130-0323031022010000-3133330311121302-2313130300213303-0231021301121222"></a>

### Direct properties for `api_specification.api_definition`

<a id="canonical-3320110100101323-1333110203233132-2023010323210303-3332210011021132-2023002101212022-3203321321002312-0033101122123203-0222231330103201"></a>

#### `api_specification.api_definition.name` property

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

<a id="canonical-0112011330312200-0233130300332031-3100032113100212-2332031003303023-1133330300331033-2210100231002000-3120213310303122-2103213313111303"></a>

<a id="canonical-0121123211310200-0032110233021313-3231230203210233-2232210300121110-3012033312031210-0220233312031211-3102333210110320-3122230101011302"></a>

#### `api_specification.api_definition.namespace` property

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

<a id="canonical-0121331331012103-3133202301030032-2332002201012132-0221311210311330-0301020223230311-2112310202032322-1222203222003110-2112312010201300"></a>

<a id="canonical-2300333111002022-3031012012121213-1103011223203023-3302232323001223-0012301031223120-1311101022113030-3113223330321330-3121211322233332"></a>

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

<a id="canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- api_specification.validation_all_spec_endpoints

<a id="canonical-2202223010323020-2322320100113123-0021130200300331-0112323201310023-2230200100311001-2221311213022211-1222302010231032-0130202221030121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
validation_all_spec_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203233100311121-2013310310323311-2023332231011021-1012231120021033-0122031131310332-3130200323332222-0301203322231332-1313112012121122"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints`

- [fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332): complete subsection reference.

- [settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310): complete subsection reference.

- [validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131): complete subsection reference.

<a id="canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="canonical-3010321102320013-1120320322131003-3023111131303010-1333122101011123-3220112213113121-3032131133021302-1333310031023031-1110233200122231"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
fall_through_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101232011322332-3001313010030122-2202002233012131-1112023231112123-1221231233001313-1100001120020311-0331110330112321-3231213002323312"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode`

- [fall_through_mode_allow](resources--cdn_loadbalancer--reference--group-006.md#canonical-1112311230100333-2313002332302102-0031001300211321-2232211332010232-3221220002131020-2023103121122032-0200131013302101-2233320323300012): complete subsection reference.

- [fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013): complete subsection reference.

<a id="canonical-1112311230100333-2313002332302102-0031001300211321-2232211332010232-3221220002131020-2023103121122032-0200131013302101-2233320323300012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow

<a id="canonical-3110330022103213-1000221320230102-0221232120133033-1301213221021012-1032301223230231-2200321021032033-0100303102200230-1102032022123302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
fall_through_mode_allow = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="canonical-1100000102332202-0023301233232201-1012001213320022-2030103333231302-0203130122320031-3310110303021301-3123121010222210-0012133013332111"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
fall_through_mode_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013121210321003-2201031333030003-0020322211232310-3333220010012012-0101223111321210-3012322132300331-3130232221330330-3122111100232121"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom`

- [open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100): complete subsection reference.
