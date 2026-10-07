---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-0202323233031101-1323302210302322-1231003212131301-2002332233310112-1203102032311001-1101301313321102-3100320222201300-0312023130123210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- active_service_policies

<a id="canonical-3333102103131111-1312023003033200-1130132002210132-3213012133332111-0223202122021231-3103322130331313-2311232333031131-1333023300112222"></a>

Type: `"single"`. Computed.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Additional upstream details:

List of service policies.

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

- [active_service_policies](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3333102103131111-1312023003033200-1130132002210132-3213012133332111-0223202122021231-3103322130331313-2311232333031131-1333023300112222)
- [no_service_policies](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2123133221100300-1332201322123031-0223020033302100-0012203203012302-3210301020331023-1032022110113221-1110220203031213-2123232211032300)
- [service_policies_from_namespace](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1331300022003002-1003301212201010-2332302332323033-1313300013202011-1100101111101023-0231310011000031-3302000103312122-2320001102213020)

Select alternatives according to the provider validators above.

<a id="canonical-1031310003113110-0021223213311021-2221310203110230-2300331232022310-0121113120030202-1212130013210131-0323232013131003-2130102031232120"></a>

### Direct properties for `active_service_policies`

- [policies](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3022120312121210-3110301332301001-2212033121220221-3133130201121213-0131212331101311-1030023011130331-3321130012013230-3232302223331201): complete subsection reference.

<a id="canonical-3022120312121210-3110301332301001-2212033121220221-3133130201121213-0131212331101311-1030023011130331-3321130012013230-3232302223331201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies.policies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [active_service_policies](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0202323233031101-1323302210302322-1231003212131301-2002332233310112-1203102032311001-1101301313321102-3100320222201300-0312023130123210)
- active_service_policies.policies

<a id="canonical-2331221313302313-2003131303032022-3113021001202011-2132303203020222-0011212001123201-3033320132311311-3000333103323203-0200030000300033"></a>

Type: `"list"`. Computed.

Service Policies is a sequential engine where policies (and rules within the policy) are evaluated
one after the other. It's important to define the correct order (policies evaluated from top to
bottom in the list) for service policies, to GET the intended result. For each request, its
characteristics are evaluated based on the match criteria in each service policy starting at the
top. If there is a match in the current policy, then the policy takes effect, and no more policies
are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none
match, then the request will be denied by default.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0120232320313022-0300332200013311-1222220331002231-2011323231003032-3212022230330002-1133000113032332-3202012002111331-3223323302231121"></a>

### Direct properties for `active_service_policies.policies`

<a id="canonical-3222121122231213-2202302330222223-2120332332220311-3232212102032231-2133302011113310-3300322311000212-1322231131112332-2313010231120300"></a>

#### `active_service_policies.policies.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1110032133303333-1301113123023302-1211310003003200-0111113311321203-0332023023003223-1222111032312030-3123032200310321-2101202111231101"></a>

<a id="canonical-3322023311113220-0013101012303000-1120132103121313-2100332131233220-2202003202312331-1320300132011123-0002323031032103-2021213021101003"></a>

#### `active_service_policies.policies.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2102223313301232-1130101133001133-1002013312121301-0301300203010310-2312012233021103-2021222022200023-3131302231002323-3031031100133033"></a>

<a id="canonical-2220133132023233-2002212122123022-0320112332223011-3123123313311320-1223022300020020-3032003031131233-0032122212121313-1031333132000311"></a>

#### `active_service_policies.policies.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- api_rate_limit

<a id="canonical-1032212300210002-2030201132210230-1102312332230021-2110031033031111-0001110013233200-2123112122011112-3212003022312203-2110000121332012"></a>

Type: `"single"`. Computed.

\[OneOf: api\_rate\_limit, disable\_rate\_limit, rate\_limit; Default: disable\_rate\_limit\]
APIRateLimit.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"bypass_rate_limiting_rules\",\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]"
}
```

OneOf alternatives in this subsection:

- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1032212300210002-2030201132210230-1102312332230021-2110031033031111-0001110013233200-2123112122011112-3212003022312203-2110000121332012)
- [disable_rate_limit](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0230200132300230-0033310301032010-3311301312300002-1010120312120321-2303002232103233-3230310021030203-1302300133013231-0102323300011210)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3103123023131201-0322000101103313-0121302000032313-3323033313013312-3121010201201001-2221102102113023-0202100323302113-2113001301020212)

Select alternatives according to the provider validators above.

<a id="canonical-1012032101221110-3021103101333101-2113230310310030-3012332313020201-0323333232113321-3232233002131123-3203112013310322-0113302003303221"></a>

### Direct properties for `api_rate_limit`

- [api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222): complete subsection reference.

- [bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123): complete subsection reference.

- [custom_ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1010130033000023-0003001221210201-0012121132223103-0120030303211233-1033132223023030-1321021102203331-2100133012021110-2303033132131131): complete subsection reference.

- [ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2303033302312111-0231012102313300-0030302012321311-0331312000002132-3021100202110103-2120133133300232-0230222110332003-1102203220130300): complete subsection reference.

- [no_ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0231102032303012-3001232002310032-0303213312332003-1321010302001031-0232222203033013-0102231202310231-2103133110112032-0131120102211320): complete subsection reference.

- [server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011): complete subsection reference.

<a id="canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- api_rate_limit.api_endpoint_rules

<a id="canonical-1220110203202202-1303322210231133-1010032221231010-0012303120303011-3201322320033011-0233301310133021-0311110203330321-1133131223230111"></a>

Type: `"list"`. Computed.

Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate\_limiter\_choice:
inline\_rate\_limiter or ref\_rate\_limiter.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1000230331033213-1221101211303222-2123202310010202-3001321222321233-1300230221201101-1312200001310231-2120031031000210-1002230302300320"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules`

- [any_domain](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1320023211030301-2122103330213231-2222123321233330-3222333303100313-1312031213211322-0220302323311201-0011033101011003-3001313012323320): complete subsection reference.

- [api_endpoint_method](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2010311300003302-2232122120331200-2313233000213333-3031011201222130-1232201103100123-1002023301203131-2333032011002121-3010003223120000): complete subsection reference.

<a id="canonical-1321120311031020-0122300223303123-2230213220133000-3120003012300321-0020123122210102-0323312333123302-2132322132213300-0132112222223323"></a>

<a id="canonical-3303231202112121-3212210112120003-1030200223022021-0310332313123113-2122310120301320-0011312012200230-2011130002100030-1332320312322000"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_path` property

Type: `"string"`. Computed.

API Endpoint. The endpoint (path) of the request.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120): complete subsection reference.

- [inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3012131113032013-2011102123010231-0200013130010331-3303113012200221-2032121332302003-3312123233133323-1022101001123233-2300022322113120): complete subsection reference.

- [ref_rate_limiter](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0203131121031203-0231022021312213-3030031313101002-3103010030313212-2112331113110000-1132102002132201-1110221111132330-3312022221021011): complete subsection reference.

- [request_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1222020303302222-2231310311313220-1200223312213020-0331213020120100-2203201022122013-1021321123100301-3222310220022301-1310123112333330): complete subsection reference.

<a id="canonical-1003010303313333-2132321310310232-0021013323010221-2300100231001130-3011323101011102-2102001122113321-1103231102330203-3220330003132322"></a>

<a id="canonical-1300230230212130-0310113113210010-1233222010002221-3112110101222222-1222311012133311-2330210333102030-1110102323001220-0111311111001213"></a>

#### `api_rate_limit.api_endpoint_rules.specific_domain` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1320023211030301-2122103330213231-2222123321233330-3222333303100313-1312031213211322-0220302323311201-0011033101011003-3001313012323320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- api_rate_limit.api_endpoint_rules.any_domain

<a id="canonical-1201201111100222-3131202102222212-3300312132003303-1110101020213310-3320001123203212-2101032030312212-1132103213103310-1111211322021002"></a>

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

<a id="canonical-2010311300003302-2232122120331200-2313233000213333-3031011201222130-1232201103100123-1002023301203131-2333032011002121-3010003223120000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.api_endpoint_method` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- api_rate_limit.api_endpoint_rules.api_endpoint_method

<a id="canonical-1230123333023203-1203201331113220-2113302300213220-1230310021112333-1100031213010203-1112032033130320-0312230303231231-3233200310323100"></a>

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

<a id="canonical-1110001020303201-0133000130111022-3303021220031121-2032012321110310-3130110230330031-3003012113110123-2320323003110332-3021200212223323"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.api_endpoint_method`

<a id="canonical-3323233203322221-3022213130323013-0312330311010121-3031313131333231-1032213201112010-0131331211213332-3121012212102002-3132333200201020"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_method.invert_matcher` property

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

<a id="canonical-2321223112212323-3312122322110121-2000220202223020-2210121323000131-1330122212000211-2330311021020022-0012331313300012-2233201310022002"></a>

<a id="canonical-3221330133212000-2321103000202033-3320232310023213-3221333131221232-1231222110201111-0202011030320023-3310232113113021-0220022121023033"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_method.methods` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- api_rate_limit.api_endpoint_rules.client_matcher

<a id="canonical-1331033013220213-3330333220132100-3031123102212033-0103111312211312-1003033102011213-1300121221030333-0212010201031112-1211002323003332"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2032210103331100-2120013113120121-3333322230330001-0120010330030132-3133211032122033-0321023303103320-3303322221011112-3130120210311211"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher`

- [any_client](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1322020131003000-3231032323013332-1101211112223000-2313131212323033-2320330201013111-0121220310230013-0101323032011112-0233233001332221): complete subsection reference.

- [any_ip](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1112212202203213-3133303101002333-1133000321121220-0200311031121020-3322011202211320-1131120000310221-1111223031101130-1020131101201302): complete subsection reference.

- [asn_list](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3120113023203232-3212023300300232-2101100120120103-2030213310333230-3031130320313010-0023021233020300-1013223100003301-0212013220030011): complete subsection reference.

- [asn_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0020001330021010-2211130231300323-0122200312322012-2110223200311233-3032113033302312-3121033320321300-0002013031200320-1030300231233020): complete subsection reference.

- [client_selector](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1303212033123031-1130213010203321-2123133123113012-0121131021302020-0223222323001322-2001030221320022-1202231233233132-1212322203302103): complete subsection reference.

- [ip_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1013331332023310-1222120322223012-1012320211212233-3001101013303122-0123223330001301-2010100220303231-3201302132233122-3301123103333002): complete subsection reference.

- [ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3131003120131223-1000031102213012-1133011013010000-3202331131200231-0211101313132301-1033031131001021-3231122021220303-2112123200231322): complete subsection reference.

- [ip_threat_category_list](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0321312321003321-0201200000330112-0333013111211201-1201320010022320-2233003200303322-3011012313210323-0031221212020220-2221100201232332): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1231210133121021-1000021202200130-2302111220130231-2312210000110102-1312022031323300-1111022020020033-0111021131310022-3300232022202212): complete subsection reference.

<a id="canonical-1322020131003000-3231032323013332-1101211112223000-2313131212323033-2320330201013111-0121220310230013-0101323032011112-0233233001332221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120)
- api_rate_limit.api_endpoint_rules.client_matcher.any_client

<a id="canonical-2300330233122100-3013121031303221-3013221300201230-1202022130121303-2123222322330123-1013000323232101-1112022231100212-3003220011003310"></a>

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

<a id="canonical-1112212202203213-3133303101002333-1133000321121220-0200311031121020-3322011202211320-1131120000310221-1111223031101130-1020131101201302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120)
- api_rate_limit.api_endpoint_rules.client_matcher.any_ip

<a id="canonical-3002131201031320-3320113321120230-1311121331232123-1132131100100231-2231112002233200-3203111122310010-1032023132003031-0231213002133320"></a>

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

<a id="canonical-3120113023203232-3212023300300232-2101100120120103-2030213310333230-3031130320313010-0023021233020300-1013223100003301-0212013220030011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-1213333002200211-3312022020333123-3322231101101332-1223033031321032-0120311202203022-2112133210211322-0232313203310322-2000032120113331"></a>

Type: `"single"`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-2103223233113021-0301123010103123-1331000133301213-2032223322313110-3313221323010333-3321113310300011-1121201210203002-0032331133023211"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_list`

<a id="canonical-2121231011330221-3330330213301010-3021001321003203-2303123133213303-0120333313322223-1021220002322303-1023100311311131-2212033202133120"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_list.as_numbers` property

Type: `["list", "number"]`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0020001330021010-2211130231300323-0122200312322012-2110223200311233-3032113033302312-3121033320321300-0002013031200320-1030300231233020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-2130123202223110-1231230032233213-0210100123231100-3102331010213033-2000132211230100-1102313002101313-1102220100033323-1210111103332001"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-3031222212200110-2202112330100200-0211032223020123-1202223132303100-2312033010310321-2303123123012320-0301233113123220-2222130330303003"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher`

- [asn_sets](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2103323012021211-0311130110320103-3021011220230213-0330132333113200-3133032213021220-0020210332032211-0330013323030222-1331001301301133): complete subsection reference.

<a id="canonical-2103323012021211-0311130110320103-3021011220230213-0330132333113200-3133032213021220-0020210332032211-0330013323030222-1331001301301133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0020001330021010-2211130231300323-0122200312322012-2110223200311233-3032113033302312-3121033320321300-0002013031200320-1030300231233020)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-3032133102103212-1322001131110221-3203112202300103-1222003121022001-2212332121012223-1100001221300111-2103313221131300-3301213013122221"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3121223211002123-2203031030023322-3300313331313112-3013003213323012-2012123113322210-2210203211300323-1333203110212202-2233201313110110"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-0120003323033203-3100012031131311-2303113102001021-3220020123233112-3113330332302103-0310212323130330-2122201231102201-2030330313123122"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0202130033001330-2223001000020210-0003300221220000-0312031330013112-1031302033322211-3233121100122310-0311132132030212-0003233233203111"></a>

<a id="canonical-0110130201233330-0023232221332230-3113030333201003-3113213123011000-1331311232003132-0202223332130333-0010300001330002-1203030233310232"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1203030020203133-2230011211022310-0323132222230330-3012033132110321-2011321103333212-2331332303000020-2012212033113320-1121303131002323"></a>

<a id="canonical-0110310123202210-0131303130031112-1200203232021212-1031032311332211-0230023321100012-1321201331121202-1121203320302133-2113221231322023"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2102323010311123-0200011030133121-3212200111002111-0203332301203010-2123123331223212-3111123322321321-2102233123001120-1303023112330120"></a>

<a id="canonical-1220033022331132-3012133212323020-0122333021320310-3212233321313103-3012010030303202-0112111030030312-3112300033133211-0113223322030121"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2012002132233132-1201132133233332-0123220323020031-3113101203123203-3323313221230101-0311332132120003-3211122123233200-2002001321331212"></a>

<a id="canonical-2331030220000130-2111333302333220-1233020321131221-0332221230302123-3311201033021100-2332131012331220-1120213330233023-2203120300212303"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1303212033123031-1130213010203321-2123133123113012-0121131021302020-0223222323001322-2001030221320022-1202231233233132-1212322203302103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120)
- api_rate_limit.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-1000030212221033-0101121111333112-2113113303323013-0031120311223110-1021130313310132-3121010322222120-3010330301332332-0112230313032003"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0222230100011002-3132003133303321-0112212023113310-1102102212311210-0033122333221223-0013231220031111-0312010303111203-0300103003203110"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.client_selector`

<a id="canonical-1133323311130313-0022100001103120-3212332301310012-3132330133111200-0123123212011320-1221002030031010-2300300112101300-3213120200323003"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.client_selector.expressions` property

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1013331332023310-1222120322223012-1012320211212233-3001101013303122-0123223330001301-2010100220303231-3201302132233122-3301123103333002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-3232020222031030-2002020101210200-0003112103101100-1332021133021111-3031113332330330-1130110020012230-2311222132010233-1000000321323011"></a>

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

<a id="canonical-1021313220200113-1312010122032103-1020303212200113-1311320303130303-2313332012213110-1233212320220033-3123101130232212-3210030001032030"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher`

<a id="canonical-0013201211013023-0120231120030313-3100030020100110-0331220003232303-3020031200310200-2310333332330233-1103310312130000-1121100100020223"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.invert_matcher` property

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

- [prefix_sets](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0202310003233112-2022200120120013-3323001032000130-0020000312210001-0013312233211303-1101203001112122-3231120323230103-0001213233030313): complete subsection reference.

<a id="canonical-0202310003233112-2022200120120013-3323001032000130-0020000312210001-0013312233211303-1101203001112122-3231120323230103-0001213233030313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1013331332023310-1222120322223012-1012320211212233-3001101013303122-0123223330001301-2010100220303231-3201302132233122-3301123103333002)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-2021222302100120-0032300003233111-3223122023110211-3110101322032321-2203110113021030-1210222003313212-3112122231102220-3331203202000201"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0330021131303131-1303233113011120-3121301313222032-3231333300303032-1011232232103231-1230321103201332-1000200231330101-1121232223023101"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-0133231222102020-0120212000000321-2332113003223201-3313331012033021-3222223001110000-1330210302133020-0223133300123223-2303320301033100"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2032201111220101-2130231021110303-3122110310330010-0013011010231123-3130301322221020-2133211300320101-0231111331013133-1303002330111232"></a>

<a id="canonical-3331131212301000-0311021310113101-3030010032333210-2213131202221133-1012001112032213-0123200333013200-3023010201210101-0100311120333103"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2301220100023310-3230032203101211-0033200203020021-2121313021213032-3210002001200031-3302000331112222-1200232010222013-1121020000100313"></a>

<a id="canonical-3033121012300200-1131311221131113-2131131321200130-2033221000312230-2313003201322213-0112330330310121-0030110131030113-3312032121001011"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2131003003230101-0331301202100311-3222303330220221-2010000221310102-3121133121213102-2001101333303303-0313101013331202-3223010302130310"></a>

<a id="canonical-1021133220030131-3033123311100312-0011223133020302-2300213020030301-1200132221013331-0333132011221033-2001200312213111-2013121233123022"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3310233031110131-1212002121113332-2002202010011230-2321203012222211-3022321102212310-2302023000113321-3232003332322103-3333131010301223"></a>

<a id="canonical-3102223330013101-0132031221023020-3213212313000113-2100222221301313-0330010233212111-3332001310023300-3000033030022122-3321203200103310"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3131003120131223-1000031102213012-1133011013010000-3202331131200231-0211101313132301-1033031131001021-3231122021220303-2112123200231322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-2220301312323113-2012031131213011-3233203332303310-3101233313200301-3223330223230010-0103012030320000-0323103020331121-1120301303001002"></a>

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

<a id="canonical-3131132022202100-0131320132021123-0133320021110331-0210230101313021-3230133121012111-1233021021322213-1101202103030100-1012112131011232"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list`

<a id="canonical-2031021333212231-3101222233030231-1312122022231100-0021231030130313-2120331020320333-1020132200002301-0210300113222333-0100300011223123"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list.invert_match` property

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

<a id="canonical-3010130302230132-1232121330002211-1222312002331132-2031100212331333-2221212333201030-1323212013031333-2103101111313031-3031310030120210"></a>

<a id="canonical-1310003231102322-1012013033133202-1101031032020132-1202100101020100-2223203112020003-3031331213321232-1000210003012121-1210120001032113"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list.ip_prefixes` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0321312321003321-0201200000330112-0333013111211201-1201320010022320-2233003200303322-3011012313210323-0031221212020220-2221100201232332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-1122002212302200-3311031320232310-1133310222223023-2312033203022133-2333233221220013-2003222030100233-3012120021033011-3020233020233323"></a>

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

<a id="canonical-3133200332203123-3200010230131033-2223211321012301-3302133321311011-2003222132101223-0001312003102302-2300021101032013-1133133210330121"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list`

<a id="canonical-3310120113003101-0011213302311320-1103211233322302-1210110233121212-1133323131111332-1111230312102002-0022112321230313-2030313021202220"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1231210133121021-1000021202200130-2302111220130231-2312210000110102-1312022031323300-1111022020020033-0111021131310022-3300232022202212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2333000313200333-1002311130232303-1003320210120320-3303132032032002-0222032111022122-0213320132220311-3231101100200213-1132222033312120)
- api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1131010113010022-0231013003301220-3013203232302323-3120102133301303-3220311111122303-2030113210303332-0101131000022030-1033230330332002"></a>

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

<a id="canonical-1122300232100002-1031112021003331-3130011012300110-0113120022020123-2033311100302211-2030312033311012-0012010233331302-0232320121301000"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-1100233302322301-0212202130132233-2222112230222231-3102221321023301-0313120330001202-1221320210110230-1013021301330203-1132200033131023"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.classes` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2231203001033230-1311110121101221-2120112332203002-0131013321232201-0220320302201002-3010230102130033-0202230220221113-2100202011022211"></a>

<a id="canonical-3110130331202003-2103002032320303-0011112132030031-2013210011302331-2221231300303120-1003013312213221-1120331131201110-3133101313311200"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0333213211231322-0132002031022211-0232202022012210-0112130111003203-1300122130231220-2210101312020310-0023022220022333-3032112312233000"></a>

<a id="canonical-0112022102213210-0201200020002333-1110333313102330-3302032332300203-1103231232121200-1031321301333233-0112122230023033-1123023122113303"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3012131113032013-2011102123010231-0200013130010331-3303113012200221-2032121332302003-3312123233133323-1022101001123233-2300022322113120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter

<a id="canonical-0013321321120221-1123210222233102-0200001102302230-2310012000002203-3132333333033003-0221012312101112-0233301212001301-0220301332112233"></a>

Type: `"single"`. Computed.

Configuration parameter for inline rate limiter.

Additional upstream details:

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

<a id="canonical-2302322123121021-3320100102322302-0200000313010232-0001333211233120-0321221233212032-2333132332102323-1332211233322101-2203020222103312"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.inline_rate_limiter`

- [ref_user_id](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3233122001002012-1313022303133223-0123203321303201-2212311332312310-1132211012223201-0000232011002021-2311030323011203-3322001030232220): complete subsection reference.

<a id="canonical-1112202030023032-1320120120113220-1023210130001112-2032210230322232-1203022120002331-3100012031330202-3323101110101012-2333212133033003"></a>

<a id="canonical-0311333013331311-2300213303102202-2111311111030012-0000320111023123-3301121301220133-2210012311203303-3331113201220203-0013322330221212"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.threshold` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1221320103133110-0023333023201113-1223012133223002-0310210321223001-1103110010130311-3202103032013213-1020001010323030-3102100012000230"></a>

<a id="canonical-1031201131301130-3311131310132320-1111310103221231-0201333323233003-0012210321233322-0310233330311100-1231201130133320-0303021301302320"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.unit` property

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

- [use_http_lb_user_id](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1033231212123013-2203110322231113-1331013112223133-1323301331010011-2323122111312233-3133012100030021-0013110200111313-0121013110221030): complete subsection reference.

<a id="canonical-3233122001002012-1313022303133223-0123203321303201-2212311332312310-1132211012223201-0000232011002021-2311030323011203-3322001030232220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3012131113032013-2011102123010231-0200013130010331-3303113012200221-2032121332302003-3312123233133323-1022101001123233-2300022322113120)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id

<a id="canonical-3011000203113101-1202213332013110-0331310020220131-0310232121021313-2312010320103003-3033300230230020-3000003333101011-0000032033010031"></a>

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

<a id="canonical-3001233221121220-3302132013223323-1320201323003123-3230010233132311-3021100103002021-0232032102001012-0003022002322302-2002321221012120"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id`

<a id="canonical-3222112310020210-1212131030021221-2023110020020000-1211001333113101-1220201121220311-2212203203201030-3310222213033113-2020020232211001"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3302101001331101-0223323132201100-3030020012103333-0000233210012203-3231012033121012-0000002020203112-2130003113013212-0131313210110113"></a>

<a id="canonical-1100003300223201-2330300121330010-3010312332012110-3220023303210323-1333131133311021-1122210322312323-3131301202301233-1323210011130300"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2221331022122030-2230210000123123-1001311112113000-3023302201331023-0212323203013120-3112223212110311-1300123000202331-1232221300023120"></a>

<a id="canonical-1103211123221021-0021100022301301-3001121122301100-0201211130031212-2211132122202321-3022311232231022-1231212332213023-0102200311103301"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1033231212123013-2203110322231113-1331013112223133-1323301331010011-2323122111312233-3133012100030021-0013110200111313-0121013110221030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3012131113032013-2011102123010231-0200013130010331-3303113012200221-2032121332302003-3312123233133323-1022101001123233-2300022322113120)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-1210213301132100-3013333320233231-1333213303111012-2302321320223330-0221210212113213-3023221312122113-0023232113231122-3013312110112020"></a>

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

<a id="canonical-0203131121031203-0231022021312213-3030031313101002-3103010030313212-2112331113110000-1132102002132201-1110221111132330-3312022221021011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.ref_rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- api_rate_limit.api_endpoint_rules.ref_rate_limiter

<a id="canonical-2322323121102322-3333120022331130-0123212233202313-2301111221123332-1302103103110020-3232211012303231-3203022032103212-0022311111013032"></a>

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

<a id="canonical-3021202022223133-2200212301323031-3210103002121233-2232113220021130-0003202112212110-2302331002202010-0100000020323000-0000322312330022"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.ref_rate_limiter`

<a id="canonical-2231321233100013-0321101001031011-0322221311101212-3032012223001311-1002321223210223-2130333232202313-0231301303022311-3201122300200123"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3123221023111202-0203221032120222-1322330322312320-1113002320131231-3032112112212110-0212003333230030-2032002311133321-2121121333212200"></a>

<a id="canonical-0231212300200313-2333112231232220-1321301320210112-1120233233203233-1121303200123310-3011013110203101-0200003000122322-1120202031231133"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1310120303100303-2103203230101023-2030013321133212-0130013122201020-2102212002333322-1202333033310011-1312123032100012-1232002033303213"></a>

<a id="canonical-1101020132031232-2312320103220320-3111313002001212-1303020331012113-2320032211210210-3012102123311100-1003000202033032-2020320211031020"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1222020303302222-2231310311313220-1200223312213020-0331213020120100-2203201022122013-1021321123100301-3222310220022301-1310123112333330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- api_rate_limit.api_endpoint_rules.request_matcher

<a id="canonical-0100223101011013-3222033121221030-2032123113220301-1320310202322213-0221222210223122-3233122311113012-2201003113031223-1312000210021331"></a>

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

<a id="canonical-0021113003123131-2211312232112012-1300231133031100-0321230201230000-2222233003311020-1322220132331120-1221333123213312-2213002113132330"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher`

- [cookie_matchers](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0011001201020201-0330122223121033-0131100103132132-3101331032132122-1201330130210120-3100112220110033-3311311210332013-2130111030120303): complete subsection reference.

- [headers](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0110230002233220-3000033102212131-0201000101113333-3301002021132020-1310131101111201-0003330302201213-1101301122130000-0102020322020310): complete subsection reference.

- [jwt_claims](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2010012232111323-1311123013131021-2012222311303113-2021022101332033-2111022213031022-3103023222211322-2303223113113000-2300122320211032): complete subsection reference.

- [query_params](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3300213323010133-2233323002213103-3122021020212002-0100310031001130-1222313130020222-1303213100101110-2001223103220131-0023123113110321): complete subsection reference.

<a id="canonical-0011001201020201-0330122223121033-0131100103132132-3101331032132122-1201330130210120-3100112220110033-3311311210332013-2130111030120303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1222020303302222-2231310311313220-1200223312213020-0331213020120100-2203201022122013-1021321123100301-3222310220022301-1310123112333330)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-1221001111313112-2011012020322111-2022231221030021-1131313232112320-1210123331113113-2111013300112213-3202112200311013-2110011121211301"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2200132100220120-2102113000202103-1112303310102003-2213112221012210-1302203122110110-2302221310110120-0303020113001222-0021121112132301"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2100320133032332-3122330202133223-1222210233021121-2130301331210222-1132233321130310-0010000321330330-2223301211022101-1230212000020301): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0301012032013110-0332013031011031-0120303020333333-2020100023110110-3121002122302330-1230013333103232-0110002132130331-0323300332131031): complete subsection reference.

<a id="canonical-1132221230112131-1212031121232130-3221131001331313-0232330233012211-0333131011000223-1211101000231023-3322011230011003-1023010231230033"></a>

<a id="canonical-1333230313303103-0002022012332300-1020311302300332-1300103032222213-2102102131310331-3323111201223033-1133022023212123-0211101011031103"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.invert_matcher` property

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

- [item](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-2302110303320102-2311132331100003-1131012101332010-2002011332231003-1002313103033330-3220321023303221-2130302313131330-3232131313131130): complete subsection reference.

<a id="canonical-2213223330101210-3100110230000221-0131210131111133-3311132001311232-2230231100312120-0032312212322330-1130111133230222-0020310220333001"></a>

<a id="canonical-0213112110030021-0031330222002120-2000301203113033-0301330122011110-0203313110112313-0332200203000002-1232222321110000-0300111331321131"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2100320133032332-3122330202133223-1222210233021121-2130301331210222-1132233321130310-0010000321330330-2223301211022101-1230212000020301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1222020303302222-2231310311313220-1200223312213020-0331213020120100-2203201022122013-1021321123100301-3222310220022301-1310123112333330)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0011001201020201-0330122223121033-0131100103132132-3101331032132122-1201330130210120-3100112220110033-3311311210332013-2130111030120303)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2030200231002012-3022102331201300-3101110011230322-2212023321031003-1221211020201313-3012021231210111-3132001000203003-1233132013320100"></a>

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

<a id="canonical-0301012032013110-0332013031011031-0120303020333333-2020100023110110-3121002122302330-1230013333103232-0110002132130331-0323300332131031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1222020303302222-2231310311313220-1200223312213020-0331213020120100-2203201022122013-1021321123100301-3222310220022301-1310123112333330)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0011001201020201-0330122223121033-0131100103132132-3101331032132122-1201330130210120-3100112220110033-3311311210332013-2130111030120303)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3010311303301332-0203013101011002-3003211310312101-1013000322300022-3331011213311101-0122211020311111-3103130223100130-1111223320312311"></a>

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

<a id="canonical-2302110303320102-2311132331100003-1131012101332010-2002011332231003-1002313103033330-3220321023303221-2130302313131330-3232131313131130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1222020303302222-2231310311313220-1200223312213020-0331213020120100-2203201022122013-1021321123100301-3222310220022301-1310123112333330)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0011001201020201-0330122223121033-0131100103132132-3101331032132122-1201330130210120-3100112220110033-3311311210332013-2130111030120303)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-0221201023311030-3100232210003003-3112332102102011-2313011313131322-0312001232312010-0210231101313101-3002311003002030-2200222033233232"></a>

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

<a id="canonical-1002321230112111-2010211122222112-0323213231100003-1121120322121212-3300022201231221-2202320010132020-0001022013223132-3010221213231011"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item`

<a id="canonical-0202333021010233-0300211122320210-2202233212003212-1220131120211233-0220303031010113-3002100231303020-2300232020021000-0212032020000131"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0321231001211313-2332022212320331-3300302230331330-3221121200303331-2322330312110020-1203330031120332-3032321310200233-1120133112201323"></a>

<a id="canonical-0002332320233221-2301220133011131-2300001322013221-1203332132210123-3221011320301013-2310100111133312-3030230202032120-1031221132330022"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.regex_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2001002103012202-2011330133203333-0101031020003021-0130232212200123-3310300212120211-1022112222130313-0101232230130130-2133120131101022"></a>

<a id="canonical-1110020310120023-0120311103101002-0202031113012000-1002133112100023-3002200012301313-3212330222203002-3002302320023001-2003101112232112"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.transformers` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0110230002233220-3000033102212131-0201000101113333-3301002021132020-1310131101111201-0003330302201213-1101301122130000-0102020322020310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1222020303302222-2231310311313220-1200223312213020-0331213020120100-2203201022122013-1021321123100301-3222310220022301-1310123112333330)
- api_rate_limit.api_endpoint_rules.request_matcher.headers

<a id="canonical-3220112320012213-2030003331302301-2232101320002232-1200303301100020-2112223321312103-2011210212101231-1322200231032011-2101100001033222"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3211303023111322-1221322102203020-2233132122231102-3323103201102003-0133110100301000-1023223313100022-1131032131103002-2322211311033223"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.headers`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3323111222233001-1013312331010221-0331312300030212-0103201230201201-0003031033033222-1132231023222010-1313130331302312-2121013110223303): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0300320313321231-1230121320131032-1323003123203103-0233002211022031-3201232011301021-1121132311102222-3222112222011200-1002200101020331): complete subsection reference.

<a id="canonical-3022101322031212-0232320212231101-0231123200303233-3010200202123130-2123331023300220-3312101202210323-3001303102030010-3012102020300102"></a>

<a id="canonical-3223130210121012-3322231210000012-0113121210021132-3133223012213020-3123022231213013-0110100212203320-2321120300201120-0311220132201133"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.invert_matcher` property

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

- [item](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1233002300231013-3310012231300000-2322121211111203-3010000023203011-1130302210123232-1030211011130302-2220310133213122-2010320123012230): complete subsection reference.

<a id="canonical-1010123110202302-1210200233020023-3013332312210011-1211201203122122-3201212132011210-3002212233223132-2013200333210133-3302320111103021"></a>

<a id="canonical-2321110332023003-1022122021100230-3113331210231120-1131333132321021-3031310331332021-1122210033030003-1202130213331103-0130200132021200"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
