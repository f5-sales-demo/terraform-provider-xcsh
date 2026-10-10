---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3131233300321233-0212301102322121-1202111033301033-2113200310232222-0003112222313301-0222301020302212-3020100103323101-2212320023210010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_v6_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-3233101001322201-3010101330010210-1230122302103210-1020231313123311-1123022210033012-0001211313213020-2123001023213012-0132232031212310)
- advertise_v6_on_public.public_ip

<a id="canonical-0120232222213030-0322113230232220-0122302022331323-0133100120000200-3012023110230021-3221033221331202-2310300030233032-0233212300121311"></a>

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

<a id="canonical-1033203233221001-3220330331333130-0333330321212211-0022123231213302-2302303012110132-2102311311202122-1212012112301222-3123103012201123"></a>

### Direct properties for `advertise_v6_on_public.public_ip`

<a id="canonical-1123333233032233-3310221210321012-3323300220332231-3220000103230113-3122121101122111-3233222113200320-0330200221030022-1010230333231300"></a>

#### `advertise_v6_on_public.public_ip.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0012211132102221-2320012103330101-3002000133233010-3113122221310231-2221313212023123-1102120221110010-2330010131002131-0232223103321112"></a>

<a id="canonical-2323101123110231-2332312232301202-0310321231333320-2333323110132300-3322310212312011-0131020020023103-0102201131312133-1122123111202313"></a>

#### `advertise_v6_on_public.public_ip.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0001223020133321-0111133213200012-3323101113220030-3023222331030013-0010213231310011-1233232000322003-1302111133223113-2113320110022111"></a>

<a id="canonical-2130201312010023-0030312203330101-1220032202122122-1102011022132032-3123131303023133-3013022221021301-1201002313023103-3223223003011200"></a>

#### `advertise_v6_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- api_protection_rules

<a id="canonical-1002001301021001-0310310133011000-0203211211232121-0011202230320113-3202313012002012-1103332112003211-3333231331302212-2022203301022021"></a>

Type: `"single"`. Computed.

API Protection Rules. API Protection Rules.

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

<a id="canonical-3030333020131010-0012113113010301-3213132331111213-3322220300111012-2221232210003010-3201321021310311-2121112102300101-2313110200113031"></a>

### Direct properties for `api_protection_rules`

- [api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100): complete subsection reference.

- [api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313): complete subsection reference.

<a id="canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- api_protection_rules.api_endpoint_rules

<a id="canonical-1331300022030313-1033111212213122-1012212303212200-0200130231021033-0031021110032020-1320130330011112-2000333033132021-0333013103210220"></a>

Type: `"list"`. Computed.

This category defines specific rules per API endpoints. If request matches any of these rules,
skipping second category rules.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

<a id="canonical-2202330100222203-1211233320212012-1333313220131213-3121110033000123-1332303312013003-1022130010312133-2111000221103031-2002321311321113"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules`

- [action](data-sources--http_loadbalancer--reference--group-005.md#canonical-0101112111332320-2210011322232033-2103230203130113-0132012311021110-0203002212031131-0211113333122200-3200122202123212-3033221000020213): complete subsection reference.

- [any_domain](data-sources--http_loadbalancer--reference--group-005.md#canonical-3003123000111312-1031323322333132-0100212110031220-0230112211200003-2222221232111312-1132133312230110-1023201223012213-1322211020102322): complete subsection reference.

- [api_endpoint_method](data-sources--http_loadbalancer--reference--group-005.md#canonical-2302200311230100-0330113230211232-0122003112233310-0122212331031231-3031003110312032-0001032212002311-0132203113232133-3011330130223200): complete subsection reference.

<a id="canonical-2023023312232222-2111000120131301-3230123331021013-2123322023031131-2133202220331120-0110303221002223-0300121123020122-0203023320200131"></a>

<a id="canonical-3101202030132213-0310033111002221-0022010010322202-3020030230200001-0211312100021111-3301201133022110-2101311222232011-2112033010320303"></a>

#### `api_protection_rules.api_endpoint_rules.api_endpoint_path` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-005.md#canonical-2032313010100003-1133311220101322-0021122232323122-1133033002303220-1232033022102302-2302100213001011-2221303001331022-0232220203020313): complete subsection reference.

- [request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200): complete subsection reference.

<a id="canonical-0023210132133012-3003201221002230-2130100111112203-3112021030331100-1321212000010201-3233121213212110-1033112213031001-1212330233212232"></a>

<a id="canonical-1103303031313103-0212211032000300-0201312233002000-2023013211303122-1331331230020021-3010112303220002-0020102102120010-3131312213310121"></a>

#### `api_protection_rules.api_endpoint_rules.specific_domain` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0101112111332320-2210011322232033-2103230203130113-0132012311021110-0203002212031131-0211113333122200-3200122202123212-3033221000020213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.action` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- api_protection_rules.api_endpoint_rules.action

<a id="canonical-0221212101313020-1320031020232002-0103223123130030-1123003312121122-0300121030203201-0120121212123322-3220231230230231-1233303111133012"></a>

Type: `"single"`. Computed.

The action to take if the input request matches the rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"allow\",\"deny\"]"
}
```

<a id="canonical-3212113212330013-1012131313003203-2231203120121110-3210231111032120-1203232323031013-1121120022002033-0211332311220032-2312130011221230"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.action`

- [allow](data-sources--http_loadbalancer--reference--group-005.md#canonical-1130212130100210-2132112221213100-1232002013023220-2021033003120120-3232121110330102-3011333232101002-1103112200323321-0203031323101133): complete subsection reference.

- [deny](data-sources--http_loadbalancer--reference--group-005.md#canonical-0222220030010011-2020212310003312-2020111223231031-1033231220311011-0033030100012102-2222103001210131-0212310232010310-3121330232321221): complete subsection reference.

<a id="canonical-1130212130100210-2132112221213100-1232002013023220-2021033003120120-3232121110330102-3011333232101002-1103112200323321-0203031323101133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.action.allow` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.action](data-sources--http_loadbalancer--reference--group-005.md#canonical-0101112111332320-2210011322232033-2103230203130113-0132012311021110-0203002212031131-0211113333122200-3200122202123212-3033221000020213)
- api_protection_rules.api_endpoint_rules.action.allow

<a id="canonical-3320000003333202-0122331130030231-2013102210201311-0233301120021123-1021221323111200-1012302103133222-2000212233012210-1102122200011330"></a>

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

<a id="canonical-0222220030010011-2020212310003312-2020111223231031-1033231220311011-0033030100012102-2222103001210131-0212310232010310-3121330232321221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.action.deny` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.action](data-sources--http_loadbalancer--reference--group-005.md#canonical-0101112111332320-2210011322232033-2103230203130113-0132012311021110-0203002212031131-0211113333122200-3200122202123212-3033221000020213)
- api_protection_rules.api_endpoint_rules.action.deny

<a id="canonical-2110120110022132-3302302101031323-1001012103230303-3120310323020022-2012030222101020-3231322332313010-2202023132213023-0230232230201333"></a>

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

<a id="canonical-3003123000111312-1031323322333132-0100212110031220-0230112211200003-2222221232111312-1132133312230110-1023201223012213-1322211020102322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- api_protection_rules.api_endpoint_rules.any_domain

<a id="canonical-1112131033001010-3112200033122321-1313021220113333-3120223031303111-1313311023200210-1032033103312022-0323313130131322-3020011021121110"></a>

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

<a id="canonical-2302200311230100-0330113230211232-0122003112233310-0122212331031231-3031003110312032-0001032212002311-0132203113232133-3011330130223200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.api_endpoint_method` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- api_protection_rules.api_endpoint_rules.api_endpoint_method

<a id="canonical-0100021110112323-3233323311010203-3222303102032301-0322010131302210-0122310330112122-0103330220102201-1011333331213221-0320123301220311"></a>

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

<a id="canonical-2022233131331223-0030111011303033-0202303302030102-3100012020013200-1303101303003011-2102022211103210-2101230223322312-3220111230311110"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.api_endpoint_method`

<a id="canonical-1231313003211322-3011300103131222-2311131313213222-2102313221013033-3202031111212123-1322312200133231-1322002132103121-0201202202033322"></a>

#### `api_protection_rules.api_endpoint_rules.api_endpoint_method.invert_matcher` property

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

<a id="canonical-0300102121210122-2223321103033232-3122013311331023-0102202130121202-0130211003030313-1323320133101311-1332011201300102-1021312023111121"></a>

<a id="canonical-3111331103023202-1212001130120333-0003013222100211-0123210011002111-2110303201101221-3231331222200020-1231032303303322-1021300231320320"></a>

#### `api_protection_rules.api_endpoint_rules.api_endpoint_method.methods` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- api_protection_rules.api_endpoint_rules.client_matcher

<a id="canonical-3330123221110001-1101000201212301-3111310032131010-2321323222210003-3333320000220102-0210233032022313-1220213023320012-3200012112021200"></a>

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

<a id="canonical-0001302200201023-0322021033132331-3101121031312331-3212311203002132-1303220111001002-1212111213111100-1023123332123300-1333300202313121"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher`

- [any_client](data-sources--http_loadbalancer--reference--group-005.md#canonical-2332122312112113-0032202200321031-0231131010020301-1102123322020003-1012230002123003-2221213201021002-2013002210020333-3102300121323023): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-005.md#canonical-1021102310232111-3103313221100013-1131332130132123-2200310212302121-2203111210022033-0122110003320330-0201302300330031-0331031310332020): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-005.md#canonical-1002113013302030-1121201010132131-2011023110333110-2232032022233221-2312311001210112-0001121220010220-2021112120103002-0330033331111320): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3103111133100112-3221211011323101-1112031302003123-3001301223211223-2332203321322200-2321200031031000-1331013201321102-2331210220231331): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011333100122230-1112323122220323-0121130103121033-2323131000020321-1202103012031302-0312100132221220-3303133111110101-1101220122231103): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0001203103110201-0203123332322212-1001211320003302-3112310013311330-1001023103331200-1322103001200211-3010232221103311-2200233121031211): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-005.md#canonical-0021011232313002-1330200211123212-2221021101012330-0212323120030222-2011320300302031-2321023301212132-3312020131230221-0232332011220032): complete subsection reference.

- [ip_threat_category_list](data-sources--http_loadbalancer--reference--group-005.md#canonical-0311133213310011-0331131310322011-3213312120230221-3001120210102003-1013012212110132-0121123102001333-3210223011212230-3213330131202200): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3101031310101220-1023332230123232-1000313023000233-3220320010021231-0130333130130201-2103230222013323-3231321120020211-2031201003211001): complete subsection reference.

<a id="canonical-2332122312112113-0032202200321031-0231131010020301-1102123322020003-1012230002123003-2221213201021002-2013002210020333-3102300121323023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313)
- api_protection_rules.api_endpoint_rules.client_matcher.any_client

<a id="canonical-1031111312123002-3320213203202012-2113213220303220-1310331123313302-2303230012223333-1233002102222331-1021333103300022-1023101332000330"></a>

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

<a id="canonical-1021102310232111-3103313221100013-1131332130132123-2200310212302121-2203111210022033-0122110003320330-0201302300330031-0331031310332020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313)
- api_protection_rules.api_endpoint_rules.client_matcher.any_ip

<a id="canonical-3332122002131211-0111102101100133-2222033112133120-3103130130323031-2300302023310201-3231110222101222-3003223300010330-0331301112222012"></a>

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

<a id="canonical-1002113013302030-1121201010132131-2011023110333110-2232032022233221-2312311001210112-0001121220010220-2021112120103002-0330033331111320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313)
- api_protection_rules.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-2201122312020220-3321213220110202-0233230001100330-0231201233303120-1011222103311222-0003101233212201-0122312330211211-2001132002130020"></a>

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

<a id="canonical-0013113202010102-3100332123300013-1032221332000013-3121212131213011-3103302033023213-3010013032030022-1030200011313300-2323020123223132"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.asn_list`

<a id="canonical-0103310012310001-1100002002122222-1110322102102102-0300202001223322-0113021033110110-2111032312012131-0113301212201121-1310001331233132"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_list.as_numbers` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3103111133100112-3221211011323101-1112031302003123-3001301223211223-2332203321322200-2321200031031000-1331013201321102-2331210220231331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313)
- api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-2012322213101130-2032012322003113-0002303232231102-0322130100100310-3021332311212023-2230201000301330-2000232311132133-3300221312121033"></a>

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

<a id="canonical-2301312303012023-2213130023201213-3133210201123021-2033312113002223-0122012023312123-3111310213202133-2022321120001202-1121132202312303"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher`

- [asn_sets](data-sources--http_loadbalancer--reference--group-005.md#canonical-1313303322111222-2313012002301322-0030203310213220-2131010332330310-1021211222112302-3233331023231100-0233032131321103-2131211200013011): complete subsection reference.

<a id="canonical-1313303322111222-2313012002301322-0030203310213220-2131010332330310-1021211222112302-3233331023231100-0233032131321103-2131211200013011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313)
- [api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3103111133100112-3221211011323101-1112031302003123-3001301223211223-2332203321322200-2321200031031000-1331013201321102-2331210220231331)
- api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-3220201130032322-2312231113102321-3203332312031302-1301132022022020-1002303322030233-2003321130030210-3301121023100100-2320112220200103"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3201001002110213-3020010031223111-1110333232300323-1110000011300030-0023123223021113-0031333113001213-2203223121331322-0330030022123333"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-1103221022312111-2300023332313103-2302313322020131-0210112132311200-3132003332103123-2020200110300102-0010131211002123-3201333330210300"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.kind` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2020302300003200-0330210320301120-2332221022321020-3232310333013323-0030332121330023-2032322120002322-1303200103030333-3031232313111011"></a>

<a id="canonical-1011022210312220-2302333211221312-2020210132213302-0012011002103123-0320220221322221-1232311300000123-2201102323101132-1302121221000313"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0003301331203213-3022013322111103-1322110330121131-2301000020323310-3330320100123001-0230230231331322-0132323213101213-0032323221303323"></a>

<a id="canonical-2031302312032030-3123201311323203-2122003030000322-0023202200123003-0010131332133103-2212310222330033-1313220233003212-3021232211312333"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3203231212230013-0002030301330021-1323323003110002-3032131202003233-1120213131001223-0302132322311002-1301310131120003-2310122032123300"></a>

<a id="canonical-2130121220213021-1221220320213303-2111020232332330-2322313233013033-3013023030120122-1321211121201111-0122201003032010-2133012223122030"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3330302003230300-1323122330010100-1213131222220012-3102111100103212-1331301011232013-3311100310201233-2102333033222022-2131231112310211"></a>

<a id="canonical-3211101323220313-0121303131031032-1301130023211301-2203211000130010-1330322222032133-2100323131010322-3211120033120103-1021230322331231"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.uid` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1011333100122230-1112323122220323-0121130103121033-2323131000020321-1202103012031302-0312100132221220-3303133111110101-1101220122231103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313)
- api_protection_rules.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-1021320001211020-3032102121312030-0302130130302320-1012232210321221-0322302213011320-0303010210002001-1211200012322212-3203313300122102"></a>

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

<a id="canonical-1211021222331132-1021102222121321-0113321332220033-0100320203322313-0121000331333311-2111213201131220-0103110112320222-3303033133110011"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.client_selector`

<a id="canonical-1113032310231110-2321333332010333-1331002321300210-1321011303023113-3112330200232203-3131133313121212-3300110131221031-0033200003323110"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.client_selector.expressions` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0001203103110201-0203123332322212-1001211320003302-3112310013311330-1001023103331200-1322103001200211-3010232221103311-2200233121031211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-2232321333001213-2123110031031132-0333331030301120-1101203131213101-3023203030000203-2331332321231312-0223022212311113-1212003121102223"></a>

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

<a id="canonical-3011311012313112-0030120011312310-0011312011112302-2112303111030231-2231022230223032-2211300133103003-2102003001230123-1013200330011001"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher`

<a id="canonical-0011302030100030-1311332320223231-3231202312232003-3020221100222101-2101131233110302-3230333003011102-0111031110102232-3021033321100121"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.invert_matcher` property

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-005.md#canonical-3132210133121112-1133011231222331-1222233302200302-0132013002213011-3103112211130100-1322121103032123-1030133313321212-3311301311322322): complete subsection reference.

<a id="canonical-3132210133121112-1133011231222331-1222233302200302-0132013002213011-3103112211130100-1322121103032123-1030133313321212-3311301311322322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313)
- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0001203103110201-0203123332322212-1001211320003302-3112310013311330-1001023103331200-1322103001200211-3010232221103311-2200233121031211)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-1331130133120100-1330011301210103-0011023102010301-1332313110222021-1001210211233311-1332022012130020-1033211300313323-1021030030101212"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2230021232132223-3322311120130200-0110032313001123-2202201200231003-3100203300103122-0033201133011123-0220133012020201-1230310310313230"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-3103121022031210-0122001302310030-1310111030330023-2001130001033000-1032322111100331-1102123120201313-3000002113300100-2313331331202302"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.kind` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3012212023202132-1133100213231233-1101331002120203-1010333332133203-1230111311330103-3200301100120213-3323002230330232-3201323200320002"></a>

<a id="canonical-2111311230201330-0303011200020303-0131110011322120-2100222321323310-1002302222201122-0233300012101321-0321310313013132-2303301331011121"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1110000233311123-2201123213122121-3020213020032023-1232332233033311-2003302021112032-1220111100331203-3101131103001003-3222210031323030"></a>

<a id="canonical-3322031303012310-1012030012201223-1122002110233320-3103120102233302-3023330202003331-0111322313021320-0120021320123201-1322310103303100"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0001213011233203-3002322223120323-1201233013030111-3013032130300232-1032102110010232-1030332221321321-3022000102333112-0131330130023120"></a>

<a id="canonical-0021132331122201-0022023211313201-0031012003133300-2110203231210130-2300001213323130-1110130233122233-2013022033320000-0120222120320112"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2202223332230333-3200003313033211-3231330020020223-2130213210313212-3203302313300231-2231210003320331-1103223123212312-3302222222203332"></a>

<a id="canonical-2113210222312320-2330023130233312-3113201331231321-2330230032202002-0223321203232303-0102020120323301-3221202122023133-0002220301311230"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.uid` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0021011232313002-1330200211123212-2221021101012330-0212323120030222-2011320300302031-2321023301212132-3312020131230221-0232332011220032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-2221121020133101-2203333001001012-3210311213233233-1021331300110110-2123002313221302-1011321133321231-0120002231100111-1033112320203212"></a>

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

<a id="canonical-3210112121023213-0113102231231212-0033102100010310-3301231121231220-2323203213013133-3012023233102330-2210112331022121-1112111120310223"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list`

<a id="canonical-2302003122322310-0213130102131222-1313233223221313-1213123113012110-3330300022220330-1131301131312211-0032030123010122-3000211323331131"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list.invert_match` property

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

<a id="canonical-0323002133111010-0230003101013212-1022102231333000-1200031012320113-2102330313031022-0120303112130312-3031000110020233-2321201231200011"></a>

<a id="canonical-1111220310113101-0020022111133220-1002232102220132-3313001300303230-0131212223311220-0220121133120321-3113020112200030-1323223301231333"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list.ip_prefixes` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0311133213310011-0331131310322011-3213312120230221-3001120210102003-1013012212110132-0121123102001333-3210223011212230-3213330131202200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-1312022121222113-0033210333311112-2012112110010330-1223212211330232-1110021310212203-0033203303122201-1020131231301023-0121020201201030"></a>

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

<a id="canonical-0311210310202102-3011103032231032-1300130312031111-3132023112223333-1303113032101332-1130002232210002-0213120010100010-0120112230013233"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list`

<a id="canonical-3231100132322000-1023213231320222-2002301103203131-1113220113130222-1002021123001023-0022212000230100-3200303121103301-1302133033221031"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3101031310101220-1023332230123232-1000313023000233-3220320010021231-0130333130130201-2103230222013323-3231321120020211-2031201003211001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011232012120030-1131010320010202-2021010202232213-3323002102323123-0133030100133020-0221012213101102-3010221200000000-3231213201002313)
- api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-0111231302020211-0212022211012021-2233130103122031-0023213203320132-1303221022203103-0121231232121132-1201332033120120-1002001111123212"></a>

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

<a id="canonical-2113323033232230-0331122131300101-2011212322113023-2101333023002311-1021111313002030-0331103010031130-3330313223132011-2313013313120330"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-3023012000231022-1110102110031322-2022121013221023-0133311123020210-1020110022120031-1331322220200202-1201003101113322-1310030333331212"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.classes` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1010222020222012-0130211303321323-3212210322113323-3303011233313010-3022313120203001-1011010212022312-3231003310330102-3301011023330203"></a>

<a id="canonical-1023133222103103-2101123303223230-0131101101211323-0330122010030101-3100201312111222-2230023100012033-2132001022230100-0103223331321013"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2321020033023220-2012123323013220-2131223023002213-2111023123013220-1000100223220303-2202023301120021-0011111031120200-1121222002123333"></a>

<a id="canonical-1232211132202111-2121111133220111-2221100312001223-2132232231023122-0320131132213310-2201012113233130-1333131032021033-2302313200031330"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2032313010100003-1133311220101322-0021122232323122-1133033002303220-1232033022102302-2302100213001011-2221303001331022-0232220203020313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- api_protection_rules.api_endpoint_rules.metadata

<a id="canonical-1330130020201213-2001211021032112-1011322232120332-1321203023131020-1102021023001302-1203103101031233-1331103100021301-0101210330210212"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-0330003233302332-0310101211001200-2010211213313331-3023032310210222-0121310330323301-0302300033013230-3132201001213031-2123202121023121"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.metadata`

<a id="canonical-3033211021112212-2022211002303231-3313212112121323-2103222010301332-0201102031000003-3201030331003301-0021103003122210-2112301312002323"></a>

#### `api_protection_rules.api_endpoint_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1232320221010212-0111110202213102-0220231301123311-0301200330010000-1103233223221102-0020122111031120-1330010133012313-2320033301000131"></a>

<a id="canonical-0202010133321113-1110023232013103-1221210332323302-2210103112223032-3212221321030021-3110012101331310-0022001230221332-2001032122230333"></a>

#### `api_protection_rules.api_endpoint_rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- api_protection_rules.api_endpoint_rules.request_matcher

<a id="canonical-1111231000030230-3002000320110333-1332112223200330-0210130201311022-2321333220112101-1002313230102122-0321313032130311-2322222211331203"></a>

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

<a id="canonical-2200001100013020-3333230020310122-2230230010203323-3111231320033312-0111033212313212-3232321131303000-1330113303330133-1322130311323231"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher`

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-0322110311330020-3323112230332113-3310302010223331-0222120332303012-3310103122023201-3122103321202031-3320303133010022-3032013030212033): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-2231022310320233-0200010223220000-3322322122021331-0311013010121202-3013123213203013-3031101200112003-3020232032321302-0332103133011011): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-2011102013200313-1002102013312111-0120321310212113-0202112331222210-3321220313232213-2030300312202101-2203330121322322-3313322323303322): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-3310122132030102-1203120120323111-2132201013301221-0100331310020010-0001300021101210-2203001130113222-2032311310131010-0110121221112121): complete subsection reference.

<a id="canonical-0322110311330020-3323112230332113-3310302010223331-0222120332303012-3310103122023201-3122103321202031-3320303133010022-3032013030212033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-3021321202002013-1001012002230022-1012320301212201-1123221320100200-3001211312223233-2021201133010022-0232031221303312-0322301021322013"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1311201211101322-1233022333323131-2233213333022000-3131001302033111-2021001033021310-3123011210321302-3133002231231133-2201101320213203"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-0121031033312110-0110213203110321-1322233302310120-2223031110320221-3213101103030212-3201211012231132-0101203102123220-3001311101112032): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-1012331030303233-2330211233302201-1300303013302113-0222302001232310-2133022103330030-1322212010213211-3223000111021033-0031322213331230): complete subsection reference.

<a id="canonical-1232301213013233-1123123332031000-2332323021112031-1133013002301031-3031122032012230-0200013231222123-0001023332212231-1132312330021021"></a>

<a id="canonical-1012230322122133-2011123300020123-0013130333322022-2012103202222223-2232132022203220-0300132202232311-3120201332103311-0323200100222301"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.invert_matcher` property

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

- [item](data-sources--http_loadbalancer--reference--group-005.md#canonical-1011020021123333-3202300023210103-0202022002303101-1221113123321031-1222302311322301-0113022201331001-1300323310133020-1112112031313332): complete subsection reference.

<a id="canonical-3323222332210012-1030231302121203-0103011011033022-1303033021210321-0201211111230000-0301131123101301-3223302210222113-0132121100231322"></a>

<a id="canonical-3232220103330030-2223130211331012-1101302030000330-2203132311233112-1101320132313333-0023200113132202-1200003101121310-2311103232203320"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0121031033312110-0110213203110321-1322233302310120-2223031110320221-3213101103030212-3201211012231132-0101203102123220-3001311101112032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-0322110311330020-3323112230332113-3310302010223331-0222120332303012-3310103122023201-3122103321202031-3320303133010022-3032013030212033)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-0332021032323231-0310302123333111-2021222132232031-3113313220203302-1232113010200002-2222113030212302-2020201003010223-2310222320131302"></a>

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

<a id="canonical-1012331030303233-2330211233302201-1300303013302113-0222302001232310-2133022103330030-1322212010213211-3223000111021033-0031322213331230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-0322110311330020-3323112230332113-3310302010223331-0222120332303012-3310103122023201-3122103321202031-3320303133010022-3032013030212033)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3131223320030232-2202220111232302-3333022212133320-2330210133021030-0332012013100110-0133102210323221-3002101231013332-3001332300112301"></a>

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

<a id="canonical-1011020021123333-3202300023210103-0202022002303101-1221113123321031-1222302311322301-0113022201331001-1300323310133020-1112112031313332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-005.md#canonical-0322110311330020-3323112230332113-3310302010223331-0222120332303012-3310103122023201-3122103321202031-3320303133010022-3032013030212033)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-1000032303202002-3202330123302331-2132101112013312-1130323002211303-3312033311203312-3320301121212310-2320221130320333-3130112133031102"></a>

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

<a id="canonical-0123112103222320-2312301221100231-2303121333120313-1333220311032233-2321031213201031-2030332222213112-0321031201101003-2012120230033221"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item`

<a id="canonical-0003000013300302-2321230032031121-0033332121012102-2102301132022103-1110030132101322-2311021121231120-3302220221223210-1121132320120013"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item.exact_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1100211110123303-3000321100101201-2120313213101201-1200133310000012-3210023123000100-2230322021021232-3221332101310303-0121232232123222"></a>

<a id="canonical-2210300011131123-2333131011133010-0310220102203330-2233113102211300-0330013211321121-0303111320123130-1020131130102111-3312103111211330"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item.regex_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2320222021312013-0001212331021100-1300211012330221-0201023021320202-1120200113320303-3221303201220233-3310313310301031-2302200233133230"></a>

<a id="canonical-2102310121203100-0033023302220102-3201031020100200-1233332131113103-1212213321101232-0123200210221111-2123320112130201-1021333302112311"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item.transformers` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2231022310320233-0200010223220000-3322322122021331-0311013010121202-3013123213203013-3031101200112003-3020232032321302-0332103133011011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- api_protection_rules.api_endpoint_rules.request_matcher.headers

<a id="canonical-1223031011120213-2303101220211132-1321323331222033-3210312012120212-0032303111010020-0001203302210313-3010021302312331-1132221113130311"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0230300311113023-3120301003202203-0121131021111330-1230231210232023-1202233103332101-1203231101110222-2131121001313100-3220122213112030"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.headers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-1231112310230001-2123132030200212-1021111120212021-3320002313000220-1323221033130323-3211130330211102-2211121103230213-3213013003222011): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-2120231033223112-2013010102310133-3210333211100231-3133112200001220-3231331231333223-0111213121023201-0221101021031030-1021020200212033): complete subsection reference.

<a id="canonical-2132021120121330-2311100000130031-2320333223002123-1000323030303330-2010210003100023-2313322221221213-2112101320130210-0121003311131332"></a>

<a id="canonical-3232232210212132-2110032120000023-1021001310102020-2032323231022220-2120020200011331-0032330102130103-2023311123312033-1003221223212013"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.headers.invert_matcher` property

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

- [item](data-sources--http_loadbalancer--reference--group-005.md#canonical-1123033002300013-0220212130203130-3210123223033000-2201312100003011-3011121120132101-2101232212231032-1213230102111213-3202313301311012): complete subsection reference.

<a id="canonical-3122221210101221-1012013320132002-2202132031312300-1231310113320233-1032313210203312-0031012111220102-0120311212133220-3100101300312120"></a>

<a id="canonical-1100022300003100-3321122222012223-2321232110102211-1022221103110200-3303012332231100-0121000301301200-3223203223221322-3303011103012110"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.headers.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1231112310230001-2123132030200212-1021111120212021-3320002313000220-1323221033130323-3211130330211102-2211121103230213-3213013003222011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-2231022310320233-0200010223220000-3322322122021331-0311013010121202-3013123213203013-3031101200112003-3020232032321302-0332103133011011)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present

<a id="canonical-0120010320011321-2031132202333221-2330220200311203-0231001001310112-2030103120300030-0023013323232310-0223210023310313-3032201221300122"></a>

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

<a id="canonical-2120231033223112-2013010102310133-3210333211100231-3133112200001220-3231331231333223-0111213121023201-0221101021031030-1021020200212033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-2231022310320233-0200010223220000-3322322122021331-0311013010121202-3013123213203013-3031101200112003-3020232032321302-0332103133011011)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present

<a id="canonical-3301213230231231-0331303223232131-1211222132333311-3011023200330232-2012313300022310-0002130032221202-2111320331002212-3211203022201320"></a>

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

<a id="canonical-1123033002300013-0220212130203130-3210123223033000-2201312100003011-3011121120132101-2101232212231032-1213230102111213-3202313301311012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-005.md#canonical-2231022310320233-0200010223220000-3322322122021331-0311013010121202-3013123213203013-3031101200112003-3020232032321302-0332103133011011)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.item

<a id="canonical-1022002232310131-0102231223021221-1222220213220032-0320200001203223-2113132332312310-0211012023212311-1002313232320001-1122100331312111"></a>

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

<a id="canonical-0312230131110000-2300230321200002-3313323110222101-2120230332011311-0230111131022001-0110321022320330-3003323113222102-2300120321213031"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.headers.item`

<a id="canonical-3323313003203223-2331003332111113-3001202021110121-2033132010131233-0203231033202130-0003033222310021-2130211331101233-1200233001131032"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.headers.item.exact_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3202122123002130-2323012110102232-1102311031131033-1011030001321322-3020200310233320-2102223213313120-2223310312030110-2000320010110023"></a>

<a id="canonical-0023310101300010-3331202320122011-1123120332021332-0102033100021031-1011103211121223-1021313103102303-1333320311331200-0013313213200221"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.headers.item.regex_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2002302002031200-2101213130121232-1310320132133213-2313012212030032-0331033320330332-1333002102232311-0220211303310131-3021230323303312"></a>

<a id="canonical-1013111133110002-3000100103300202-2311101213313020-3022222120031213-2132200010310022-0330202300022233-0311130331220120-3132212202202302"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.headers.item.transformers` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2011102013200313-1002102013312111-0120321310212113-0202112331222210-3321220313232213-2030300312202101-2203330121322322-3313322323303322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims

<a id="canonical-3022222000030122-3231001302122223-0213230122203130-1213022110020001-3010331312312033-0102103112303022-3322001202331112-3003111202121222"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0321201210023102-2132100123313311-0103001200011000-3301210111213011-3223102300221202-0102132232122131-1330002033120132-2220300202030321"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims`

- [check_not_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-3132211022000123-3121200101133130-0322022112300103-0322021113330111-3322032001202031-1122303302312330-1223201130310323-3101311322320230): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-005.md#canonical-0123010211330020-0212112022232021-3013013213223030-3311111211133332-2121003021323000-1023032031301222-3133131132311301-1310112332321012): complete subsection reference.

<a id="canonical-0112320223303133-3202303331100222-2321131210312230-3320112123001011-0202311130322113-2320211100223302-0221300320020233-0012212303331300"></a>

<a id="canonical-0123213122022220-2230130311221133-1032312322111312-2020023122023122-2110020133213301-2031320330233330-2323132301012303-1100232103130302"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.invert_matcher` property

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

- [item](data-sources--http_loadbalancer--reference--group-005.md#canonical-1103320110103003-2133200222131032-2333010301221303-3111133302031203-1121000031220103-2322233011110212-0123133230022103-0001101110231120): complete subsection reference.

<a id="canonical-2303230023303110-2011300222311013-1222322213132210-0023221222202312-2123132121233200-3123311110303120-1110210111300220-0123220000300232"></a>

<a id="canonical-1332021321112100-3222303310330101-3022320013133223-3323332323030333-0233303303221022-1300301003310230-3332201003032112-1200203313312111"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3132211022000123-3121200101133130-0322022112300103-0322021113330111-3322032001202031-1122303302312330-1223201130310323-3101311322320230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-2011102013200313-1002102013312111-0120321310212113-0202112331222210-3321220313232213-2030300312202101-2203330121322322-3313322323303322)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-0132311203200323-3331133330022110-1321230123110133-3103000031230110-3323221030310112-3202103213101031-1021031022101100-3030030122132113"></a>

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

<a id="canonical-0123010211330020-0212112022232021-3013013213223030-3311111211133332-2121003021323000-1023032031301222-3133131132311301-1310112332321012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-2011102013200313-1002102013312111-0120321310212113-0202112331222210-3321220313232213-2030300312202101-2203330121322322-3313322323303322)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present

<a id="canonical-1323211311312111-0230010010031323-0112200203113221-0112102131321301-3320023111112331-3300203010002000-0222333331000313-0113102122102313"></a>

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

<a id="canonical-1103320110103003-2133200222131032-2333010301221303-3111133302031203-1121000031220103-2322233011110212-0123133230022103-0001101110231120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-005.md#canonical-2011102013200313-1002102013312111-0120321310212113-0202112331222210-3321220313232213-2030300312202101-2203330121322322-3313322323303322)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item

<a id="canonical-3221012320220311-3011022313112321-0210121030212230-3202002303031001-1201000303230111-0211120223112121-1223102032132032-1011131311330210"></a>

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

<a id="canonical-1033203222300021-1200021022033223-0310113203000110-0203103003211303-0123202232112330-3001313122003331-3213212230001011-1332010331020233"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item`

<a id="canonical-3230230321303212-1102032211120201-1333023222101312-2203120320010103-1113331113210202-2311031333223100-1002313022331131-3230330232203121"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item.exact_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1013233000222303-0202133231111103-2202002202210001-0032130222232102-3023133012103231-1212031313123201-3303232033021222-1131001102310002"></a>

<a id="canonical-1331232212132213-2031013322202231-3011322212223112-1131231203232021-2221003033210330-0323332212232203-3313321310221210-2320300011113123"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item.regex_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2310303023223111-2023002023020100-2103233121103120-1211010111220333-0020203201000113-3202321212001102-3311122032033031-2012300132131230"></a>

<a id="canonical-2210332233101212-3123033031002210-3302111123100123-0303102020232110-3222200301202122-1322031230101321-2131333222233013-3120312110320132"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item.transformers` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
