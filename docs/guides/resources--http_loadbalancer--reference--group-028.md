---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0202100011213003-3130210101301000-1131320200112003-0103010103133022-1321303322303211-3333110123203332-2331103013222131-3222203113011110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_crawler.disable_api_crawler` properties

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
disable_api_crawler = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_discovery_from_code_scan` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.api_discovery_from_code_scan

<a id="canonical-0011122103200213-2303002202230301-0223021011330121-1130100121333300-0011001003121303-1231013131113233-1232132131232013-3102310113103211"></a>

Type: `"object"`. single nested block, Optional.

Select codebase and Repositories.

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

<a id="canonical-1130013200131311-0302110210020200-3220200020231211-1033123120211323-1230310102310120-1023130232221103-0033303131313012-2212112130331100"></a>

### Direct properties for `single_lb_app.enable_discovery.api_discovery_from_code_scan`

- [code_base_integrations](resources--http_loadbalancer--reference--group-028.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222): complete subsection reference.

<a id="canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-028.md#canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-3213332301002032-3000003203210100-0022201231330031-0113331201323310-2002123122101320-2220233020101131-3123113133201121-3310203112100130"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for codebase integrations.

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

<a id="canonical-1230020312133132-0330333310202210-0011101011010331-2310301103023302-3222031212030122-1012212021101111-3203133211301200-1123013300333111"></a>

### Direct properties for `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations`

- [all_repos](resources--http_loadbalancer--reference--group-028.md#canonical-2223021303113303-1200233201332222-1302013331020322-3221000122322211-1013301011300330-2303002012013300-0231220113321030-1113121130233313): complete subsection reference.

- [code_base_integration](resources--http_loadbalancer--reference--group-028.md#canonical-0103332111221020-2200131012210000-2200013132110303-2300120313323223-2103212130323213-1323003210322233-1130223200203102-0003221302203232): complete subsection reference.

- [selected_repos](resources--http_loadbalancer--reference--group-028.md#canonical-0120030311332310-3032221103100333-3300120312102000-1320231120202002-2210010032113201-1011110103311301-1303300223303313-2132011301200233): complete subsection reference.

<a id="canonical-2223021303113303-1200233201332222-1302013331020322-3221000122322211-1013301011300330-2303002012013300-0231220113321030-1113121130233313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-028.md#canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-028.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-3133322330221112-3330033120230133-0031101233223330-0230203021321023-3130210231021130-0102133032112123-0210323203010201-3303330003033000"></a>

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
all_repos = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103332111221020-2200131012210000-2200013132110303-2300120313323223-2103212130323213-1323003210322233-1130223200203102-0003221302203232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-028.md#canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-028.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-2132332003233123-0003312302232301-0321102001120122-3322111331302320-1113110222020233-2011112033211010-3311110221031123-3103302313130123"></a>

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
code_base_integration {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102320312020002-1000201310230020-3331000311311321-2110211133121311-3100122202013212-3132332212030321-3320201330010312-2103331302121233"></a>

### Direct properties for `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration`

<a id="canonical-3213123230210323-2311033112130030-1102230222000001-0212313333311031-1322222322331311-0311223313030211-1100113230230300-1212030323030230"></a>

#### `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.name` property

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

<a id="canonical-3232300333302113-2100102123230011-2211212021303030-3111330120223011-1130012013102013-1233113211113010-2013222310131213-0003310033203031"></a>

<a id="canonical-2021132110300300-0302000232031201-2021332210323013-1101123300132001-0201120200112232-3013300023301220-0013020232220310-2210010210212220"></a>

#### `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.namespace` property

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

<a id="canonical-1300013232311013-2302233202320202-2300321113321101-1222010023133022-1021333032130023-1101201310211312-3101332030213012-2202312010110131"></a>

<a id="canonical-0321313300322310-3302011130032133-0220111220013030-3212020112323322-0320213002020010-0220132321033331-0213033210010110-3301331101213213"></a>

#### `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.tenant` property

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

<a id="canonical-0120030311332310-3032221103100333-3300120312102000-1320231120202002-2210010032113201-1011110103311301-1303300223303313-2132011301200233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-028.md#canonical-0221003101323022-3312012302112023-3000330212232213-1221002030031002-3002003222012013-3323312302011333-3333203113103313-1211103303000220)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](resources--http_loadbalancer--reference--group-028.md#canonical-0322031320002333-1231133321101311-2221031200011131-2131230301003113-2013002331012111-1010333112122313-3100101131311232-2332212003132222)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-1331101231311200-0121211113322311-1332212230103313-2020020230112022-1201002122300103-0330330023121021-0021132130313202-0333133331022113"></a>

Type: `"object"`. single nested block, Optional.

Select which API repositories represent the LB applications.

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

<a id="canonical-2102312320223120-2231131122122330-0001003023022302-3033123022010032-0101132230011202-2230323013133011-2322230220133201-0323000303310112"></a>

### Direct properties for `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos`

<a id="canonical-3102333323000130-3213103201301231-0031212023331023-2232320020322123-2311201131122203-2313131021331032-2230020303101100-1110133301011102"></a>

#### `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos.api_code_repo` property

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

<a id="canonical-0033020300201021-2112213210103002-3302223002101203-1012232012320110-3312220020100332-0310003001132021-2230020302333111-0112132130323101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.custom_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.custom_api_auth_discovery

<a id="canonical-0333321320210332-0013030320320113-0201130312310220-0110031121302131-3032233230211220-1230002200133303-1101002002202232-0311131011123200"></a>

Type: `"object"`. single nested block, Optional.

API Discovery Advanced Settings. API Discovery Advanced settings.

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

<a id="canonical-2302320020012210-2030210332300302-0133210233003233-1220221033211210-3331300212322132-3112230110320220-0132023301313110-2321222000321113"></a>

### Direct properties for `single_lb_app.enable_discovery.custom_api_auth_discovery`

- [api_discovery_ref](resources--http_loadbalancer--reference--group-028.md#canonical-1310002311031003-3222203030001112-2330202232232302-2312330100030212-2010011320331203-2001310010001020-2133103030333320-1321330111213323): complete subsection reference.

<a id="canonical-1310002311031003-3222203030001112-2330202232232302-2312330100030212-2010011320331203-2001310010001020-2133103030333320-1321330111213323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- [single_lb_app.enable_discovery.custom_api_auth_discovery](resources--http_loadbalancer--reference--group-028.md#canonical-0033020300201021-2112213210103002-3302223002101203-1012232012320110-3312220020100332-0310003001132021-2230020302333111-0112132130323101)
- single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-0322223301323110-2210100211023321-0132230321301312-2000301313322332-3331010202221113-3301023031101312-1320233002113230-1010130030112010"></a>

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
api_discovery_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231302311323013-2010230002011111-3010303121213113-2003203101212133-1000110101133213-2321223203321010-2331033311021032-0301302212223031"></a>

### Direct properties for `single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref`

<a id="canonical-1121001212212111-3002301322203222-3113123102011333-0311123231310332-2030203103200322-0122131130023000-2231113121333130-1312120010022023"></a>

#### `single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref.name` property

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

<a id="canonical-2232023303111032-1212313003103030-2113220210133302-3301323213000201-3213333323212033-2202320321011322-0021301121321010-0131103011132020"></a>

<a id="canonical-1003200121233303-3102101333230230-0330133220301312-1103211131201330-1131013312121022-1303210012213002-3310322103120220-3101310200230211"></a>

#### `single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref.namespace` property

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

<a id="canonical-0212110031322202-0000011212312201-3110330333321323-2130320303130030-0320013030233031-2202020220230213-1122013223202123-3031310021031300"></a>

<a id="canonical-0301320312100133-3220022210210102-2013033000000021-1211123100032110-0033132201103131-3122013003203213-1223110233322233-1111110321223002"></a>

#### `single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref.tenant` property

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

<a id="canonical-1312311021131010-1232123020112022-3111131320310202-0101232302110333-2132211003300303-0130323303100031-0002131301100313-0230303031012130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.default_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.default_api_auth_discovery

<a id="canonical-3331030310133231-1233101221112220-0111201030102033-2212120313302003-1322022120211030-1011010333230301-2100132013122312-3123230310133320"></a>

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
default_api_auth_discovery = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131121200123133-1212031302303313-1013301330332331-1101002022012232-2321013232202222-2303001302131230-1130011122113321-2201223210020302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.disable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.disable_learn_from_redirect_traffic

<a id="canonical-3000103101111331-1221030123211303-2333230223222013-3213210022120132-2331200010130202-0110310122131211-3021332232110023-1301103212022110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learn from redirect traffic.

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
disable_learn_from_redirect_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221012011230012-1010330112321300-3320100021211210-0310132001223101-1033302320113013-1203012103122302-1000022330113011-3012213201002311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.discovered_api_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.discovered_api_settings

<a id="canonical-0331303102032213-2110332322121012-1212111102130321-1230112302323030-3210330213100023-0313011323232303-1113113323023023-0330111013322020"></a>

Type: `"object"`. single nested block, Optional.

Discovered API Settings. Configure Discovered API Settings.

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

<a id="canonical-3120100110302300-2103021030103201-0221111130211223-3103321312111030-0003202111122220-0123113321300031-0110200220123321-0103303233323311"></a>

### Direct properties for `single_lb_app.enable_discovery.discovered_api_settings`

<a id="canonical-0230333102322131-1000222010223211-2222211313132201-0030111012330322-3312220001212133-1220031011220203-3302212223020203-0312332133033220"></a>

#### `single_lb_app.enable_discovery.discovered_api_settings.purge_duration_for_inactive_discovered_apis` property

Type: `"number"`. Optional.

Inactive discovered API will be deleted after configured duration.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1103000230303211-3013102200220030-3232211010330301-3323201112330001-3233012203213323-2002131102133010-3312223131112321-2202313010211300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.enable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--reference--group-027.md#canonical-2103133222323230-1231231312012301-0313301022302012-0132222321102312-2202322120122030-1310011200232322-2331313010001230-3301022113231233)
- single_lb_app.enable_discovery.enable_learn_from_redirect_traffic

<a id="canonical-1120032233233120-1033130000101212-3222302220223310-0300230102312311-2133231313132030-3300100023231122-0021210120032233-1121002103332101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learn from redirect traffic.

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
enable_learn_from_redirect_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323203132202222-0231330232030303-0301303112213002-0030002232012332-3202313221231110-1231102102000012-1012221103330222-0003130020021302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203)
- single_lb_app.enable_malicious_user_detection

<a id="canonical-1330312121112131-2300221301232022-3132111230031311-1010132310122030-3101330022302222-1232210323232321-2220202023212022-0020111300333201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable malicious user detection.

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
enable_malicious_user_detection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333311120100023-2230031010023220-2013013300321213-0302220321233313-0222113033133022-2131111320010323-2303320333133122-3002301330000322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- slow_ddos_mitigation

<a id="canonical-3321130210030111-2101320330200320-3223131102102320-2111002221120321-3012211013203212-2010231131011301-0312012001201331-3232322133122101"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Additional upstream details:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

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

- [slow_ddos_mitigation](resources--http_loadbalancer--reference--group-028.md#canonical-3321130210030111-2101320330200320-3223131102102320-2111002221120321-3012211013203212-2010231131011301-0312012001201331-3232322133122101)
- [system_default_timeouts](resources--http_loadbalancer--reference--group-028.md#canonical-0330212210310002-3223011012010302-3213000212311310-1111301100131233-0001230210102033-3202311013232013-0011111000002231-0320131031001100)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
slow_ddos_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301322122211210-3321122110231313-1211221132220222-2033002312011332-3113021021110002-3103120333001002-2103222222001331-3313002323130100"></a>

### Direct properties for `slow_ddos_mitigation`

- [disable_request_timeout](resources--http_loadbalancer--reference--group-028.md#canonical-1333223010311031-1223211010120203-2231021132203220-2012312213221002-1001311321000223-1010102333310300-2231212011132223-0110031230110023): complete subsection reference.

<a id="canonical-2332323132331313-2301011303102131-3122131232223101-0023003230013212-0122201330203221-3122011302000123-1333133112232020-3231221322233121"></a>

<a id="canonical-0320221012302020-0333231101012013-3320313102331320-2003200221003310-2321121100200131-2020220022121120-2132100011233002-0001023132111120"></a>

#### `slow_ddos_mitigation.request_headers_timeout` property

Type: `"number"`. Optional.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Additional upstream details:

The default value is 10000 milliseconds.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3121320133300331-1202311021031110-2122333330313013-3013233122032132-2303022032012121-2312000213220122-2332312023121313-1303111100302313"></a>

#### `slow_ddos_mitigation.request_timeout` property

Type: `"number"`. Optional.

Exclusive with \[disable\_request\_timeout\].

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1333223010311031-1223211010120203-2231021132203220-2012312213221002-1001311321000223-1010102333310300-2231212011132223-0110031230110023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation.disable_request_timeout` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [slow_ddos_mitigation](resources--http_loadbalancer--reference--group-028.md#canonical-3333311120100023-2230031010023220-2013013300321213-0302220321233313-0222113033133022-2131111320010323-2303320333133122-3002301330000322)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-1101111231221203-1120230003312300-3132000101201213-2122000000111201-2222312231110333-3323010112121333-3200220210210310-0101120303120212"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable request timeout.

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
disable_request_timeout = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120111122231002-2333331020233201-3020032131233000-0232201223221311-1012212322212322-1113023212320102-3111221221223322-0223221231203000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `source_ip_stickiness` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- source_ip_stickiness

<a id="canonical-0202231310310202-3102203321100120-0022003303331033-1312113312302011-2331102102230311-0003123123020233-2303210033023103-2322033033310111"></a>

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
source_ip_stickiness = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111101023101332-0113013332131030-0031310130033201-2323032100231211-0300131012212002-1223021312230103-2033033310310331-1121322330011111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `system_default_timeouts` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- system_default_timeouts

<a id="canonical-0330212210310002-3223011012010302-3213000212311310-1111301100131233-0001230210102033-3202311013232013-0011111000002231-0320131031001100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for system default timeouts.

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
system_default_timeouts = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031000301301132-0002232123000132-3130011021001231-3113311013210310-0122300301203211-0213222010232211-0313220210221220-2120201133330322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

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

<a id="canonical-0013013123312311-3121031012202211-2230023031213113-2131331013201120-2203310300300332-0131210023112331-1023333211120121-2202001131101032"></a>

### Direct properties for `timeouts`

<a id="canonical-1332203320130332-0231203032112023-3301120321320013-1010100230012123-3201311200302220-1311031011012131-1013132200020001-3332033302320330"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1121333232030121-3033032203200111-3113230100002100-2230132212133120-0113312011022300-0100013303111121-1302333013212322-2232203033310130"></a>

<a id="canonical-1322030001231022-0203111310131032-1303332202100112-0301200021121303-3110002030033111-2013000331131323-1301322231232220-1021213121023130"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2321131303121121-1001333213122331-1301312312013203-3120122000313320-1232102131222323-3110303310102123-0113102332322322-0222320301231022"></a>

<a id="canonical-0132103323232221-3323101030003201-1300003303223301-3023233310002023-1133021213131230-0103103213001010-1300023300012100-0221023323030111"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2123013211302232-0200232332112232-3323213233023001-2203212000321322-2020212100021000-2112120103222120-2200331123330000-2100310120030113"></a>

<a id="canonical-2100302223210101-1213212113010132-1131011213301001-0113232313311023-1112203302301132-2110313131021300-0310031121310033-0033132320200133"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- trusted_clients

<a id="canonical-0133130032112133-1100010030010311-2103222010210003-2011202010331100-2122300201100231-3112021222021103-0323312032332332-2223111331101233"></a>

Type: `"object"`. list nested block, Optional.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

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

<a id="canonical-2030220021312223-1223200321322010-3232012311021321-0031033032333101-3321033130003313-2123031123321002-1311111002302202-2132231111323133"></a>

### Direct properties for `trusted_clients`

<a id="canonical-0020211212331332-0001212123012033-1102123321023121-1012121121222022-1333323303102323-2322230211230310-3133032323232122-3012001301021202"></a>

#### `trusted_clients.actions` property

Type: `["list", "string"]`. Optional.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

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

<a id="canonical-1110002102023301-1103133303232221-3203333113031301-1232303103122320-0201323003102230-1211311021200131-3033032013110312-3121322100223031"></a>

#### `trusted_clients.as_number` property

Type: `"number"`. Optional.

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [bot_skip_processing](resources--http_loadbalancer--reference--group-028.md#canonical-1002300133313303-0010300302103203-3100123023033331-0200331101103131-1113001202202300-0202002021320031-1100312103103123-1232103311233301): complete subsection reference.

<a id="canonical-2303203232120000-0103130330122113-1102113131033333-3313200130110112-1312113122323022-3132013202323112-0230130122000013-2203212030320111"></a>

<a id="canonical-0020313110231301-0232303331101123-3101022211103221-3211233303323231-2223321302220201-0331030220312012-3313332122211310-1131032121112131"></a>

#### `trusted_clients.expiration_timestamp` property

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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](resources--http_loadbalancer--reference--group-028.md#canonical-0130212200202101-2221112101301312-2221330231323300-2033221313303030-0332100303130330-3000322033023322-2211001320111312-0312213323202112): complete subsection reference.

<a id="canonical-0230231100332211-3101212220321002-3332133200203123-2111113200301010-1333223100233032-2202112002332333-1323333121220003-2213102002232210"></a>

<a id="canonical-3003303222123011-3230033031321301-2111012120321302-1020331032230110-3130013231222022-3111211131211321-2200202233113100-1133033313132202"></a>

#### `trusted_clients.ip_prefix` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2320331030201033-0012031012121333-2330212233220010-2011302302112000-0122320301302201-2022311113012122-0210320012103203-0122331331201333"></a>

<a id="canonical-2001321031321331-1310031320101332-2010232222301323-0303120123232101-1201123023031031-0012030132221130-0310130213110130-3031311230223113"></a>

#### `trusted_clients.ipv6_prefix` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-028.md#canonical-3323212001322300-3233202033131233-0023032330121023-1010133000323213-0130121100033002-1111332233212002-2200301302130230-1021031221023213): complete subsection reference.

- [skip_processing](resources--http_loadbalancer--reference--group-028.md#canonical-3213233031322121-3133002030130220-0121202211301220-3233302310133200-3013022330013131-0033013222122233-0301221231302101-3333021203323021): complete subsection reference.

<a id="canonical-2001202032132210-1011311211101023-2030131000302121-2131331331111311-1322101112123120-0321123013323023-0000101330233003-1312132232312220"></a>

<a id="canonical-3300221302000030-1230322122322312-3303122213023010-1133301323332002-0022003222123132-1102103101221122-1331132111321003-2131122210000200"></a>

#### `trusted_clients.user_identifier` property

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](resources--http_loadbalancer--reference--group-028.md#canonical-0230200100110103-0330023321103022-2112322010330320-0212120122121323-3130132223211233-0010201332230200-3102203100220300-1231303111321002): complete subsection reference.

<a id="canonical-1002300133313303-0010300302103203-3100123023033331-0200331101103131-1113001202202300-0202002021320031-1100312103103123-1232103311233301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.bot_skip_processing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-028.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- trusted_clients.bot_skip_processing

<a id="canonical-3031013201211122-1020133023133202-2221032003033203-0123302210320321-1332223113130321-2011310210021333-3332332023001312-1000300321213133"></a>

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
bot_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130212200202101-2221112101301312-2221330231323300-2033221313303030-0332100303130330-3000322033023322-2211001320111312-0312213323202112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.http_header` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-028.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- trusted_clients.http_header

<a id="canonical-1231110113303031-3023112300000001-3210300020332330-2130013022123330-1123122302101101-2031120032122222-0021100220000020-1120021301120000"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Additional upstream details:

Request header name and value pairs.

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

<a id="canonical-2312210233201331-0133323212320001-2000003300002211-3230011011313313-1322233031233033-3223111211221302-0103222120022111-2102120013101220"></a>

### Direct properties for `trusted_clients.http_header`

- [headers](resources--http_loadbalancer--reference--group-028.md#canonical-2213310213001202-2300303212123322-3020023121302202-3313131031202312-3321001301213121-2221133112122120-1313021233202331-1130223200231021): complete subsection reference.

<a id="canonical-2213310213001202-2300303212123322-3020023121302202-3313131031202312-3321001301213121-2221133112122120-1313021233202331-1130223200231021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.http_header.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-028.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- [trusted_clients.http_header](resources--http_loadbalancer--reference--group-028.md#canonical-0130212200202101-2221112101301312-2221330231323300-2033221313303030-0332100303130330-3000322033023322-2211001320111312-0312213323202112)
- trusted_clients.http_header.headers

<a id="canonical-2122010131223230-0300030103331221-1333011221210123-3322322100000223-3132200310020010-3230030200020123-2123230322301033-0300311123230212"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

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

<a id="canonical-3002302003333012-2121111031022122-1101033223310330-1103323322310302-2223200031221133-3330323302031032-0110000203212220-2023111230233130"></a>

### Direct properties for `trusted_clients.http_header.headers`

<a id="canonical-2330212113321002-2210022133333131-1101131301012321-0113103333303201-0133013110133033-3133121310210020-0302111230031012-2322122110101030"></a>

#### `trusted_clients.http_header.headers.exact` property

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

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

<a id="canonical-2020231303013011-2223213001021021-1120301212212032-1013330033121302-2213330331302302-3302030022033210-1132332211031012-0100302211002131"></a>

#### `trusted_clients.http_header.headers.invert_match` property

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

<a id="canonical-2311303233133013-2320112030331122-2302222011202010-2033020300213103-2131132102122023-1212201130211333-0320200313100023-1133300213211302"></a>

#### `trusted_clients.http_header.headers.name` property

Type: `"string"`. Optional.

Name. Name of the header.

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

<a id="canonical-1313132213103300-1220021010032001-3021103310321321-1222233231201333-1230112210330322-0022323000011303-2312033210320010-0123131223031232"></a>

#### `trusted_clients.http_header.headers.presence` property

Type: `"bool"`. Optional.

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

<a id="canonical-1020322330102033-3122100231023213-0130133232031110-1003122323231323-2212032111020322-2101023313001101-3101200100003110-2223032003203232"></a>

#### `trusted_clients.http_header.headers.regex` property

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

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

<a id="canonical-3323212001322300-3233202033131233-0023032330121023-1010133000323213-0130121100033002-1111332233212002-2200301302130230-1021031221023213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-028.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- trusted_clients.metadata

<a id="canonical-2221031221101212-1202121220300331-2002011023112131-3033032000322300-2003323200302310-0122221331031300-3022223120333100-3023300311332211"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120221311232122-3212230022100102-1033121130312031-2223031131121211-2303211010101201-1000332012211020-0203221101102022-0313311131220002"></a>

### Direct properties for `trusted_clients.metadata`

<a id="canonical-2231131022002100-3213000102012201-3100201301200221-2000221112233031-2020031333133103-0211333333013332-3121023230223102-2123213233210100"></a>

#### `trusted_clients.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-2011110131120202-3011301122233232-3323002212233210-0302113003032220-2302301311123100-2323011022311213-2302031003013033-0030201210020033"></a>

<a id="canonical-1310300131233001-3033110120200320-0323103100002203-3123332311320231-0320003303311213-2000122220131121-1102133121131103-0221232101103103"></a>

#### `trusted_clients.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3213233031322121-3133002030130220-0121202211301220-3233302310133200-3013022330013131-0033013222122233-0301221231302101-3333021203323021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.skip_processing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-028.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- trusted_clients.skip_processing

<a id="canonical-0213321111333300-2101220101311003-1321323230120103-2302130113131021-1110111233012020-0232213123221323-3021100201031223-2322103000000130"></a>

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
skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230200100110103-0330023321103022-2112322010330320-0212120122121323-3130132223211233-0010201332230200-3102203100220300-1231303111321002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [trusted_clients](resources--http_loadbalancer--reference--group-028.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021)
- trusted_clients.waf_skip_processing

<a id="canonical-2131003200000233-0110223232333210-3221001033333233-0100210222332310-0211313021202222-0233112032003203-3213000030222300-2110031213300020"></a>

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
waf_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332013311013213-1233202210102101-2202301332332033-2130302002311222-1231331220201121-2021122303232113-1311030310321233-1201033022300000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_id_client_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- user_id_client_ip

<a id="canonical-0022033130131210-3032332201221021-3230233002311132-0130000033332332-3330030231331101-0022000213122321-0301012211103113-2011213130320010"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: user\_id\_client\_ip, user\_identification\] Enable this option. Defaults to \`map\[\]\`.
Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [user_id_client_ip](resources--http_loadbalancer--reference--group-028.md#canonical-0022033130131210-3032332201221021-3230233002311132-0130000033332332-3330030231331101-0022000213122321-0301012211103113-2011213130320010)
- [user_identification](resources--http_loadbalancer--reference--group-028.md#canonical-1333212101102101-1201312203331203-0313313322223030-1231103321303311-1303322331300113-2010133022030221-3223013213222031-2231132130202030)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
user_id_client_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032001213002301-0331120033322333-0030100320321033-2022130323111321-1033121323230311-3230223120103201-0000013032330321-1020332302220223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_identification` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- user_identification

<a id="canonical-1333212101102101-1201312203331203-0313313322223030-1231103321303311-1303322331300113-2010133022030221-3223013213222031-2231132130202030"></a>

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
user_identification {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202323210303323-2030223003111121-1013131033220122-2323131000113022-2023222201320031-2232200122310312-1100121110220021-3101100133210123"></a>

### Direct properties for `user_identification`

<a id="canonical-1120231302022003-3200330131332122-2003121020012202-3301013213330122-1031231313131102-1312020023320311-0032030022012213-3020000112320313"></a>

#### `user_identification.name` property

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

<a id="canonical-3000331120030103-1033030030120303-1230203231331010-1311001233121222-1011000203011101-1300311130031123-3320302130130220-3323111312310103"></a>

<a id="canonical-0212232323203133-1320333022121110-3210023110322111-2311303310222113-0320323020330120-0123323033003210-2300132031120331-3330331130113201"></a>

#### `user_identification.namespace` property

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

<a id="canonical-0312011300203320-1211333023232132-0012332000101133-1022012332000100-1200221223321321-3231032223311311-1011122102203130-3023013001132330"></a>

<a id="canonical-3023212032113022-1112120032211033-1323332221322101-1122113201121201-3223211310321312-0303101003123221-1021300121313330-1130300100003102"></a>

#### `user_identification.tenant` property

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

<a id="canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- waf_exclusion

<a id="canonical-3030320320020232-2321302200221220-3233200121302202-2302000023301023-2230300231002210-3331111223323213-3001101212233312-2230201002203010"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for waf exclusion.

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

<a id="canonical-1320100302310201-1033012213310333-1002300002010320-3201200133002100-1030333121130013-2320311203222010-2011322121103331-0132321232310121"></a>

### Direct properties for `waf_exclusion`

- [waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-028.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210): complete subsection reference.

- [waf_exclusion_policy](resources--http_loadbalancer--reference--group-028.md#canonical-2300231203030031-1332023310203311-2011202020030201-3003033103130323-3301122311203303-1021210131330310-0012111123120220-1030113003313033): complete subsection reference.

<a id="canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- waf_exclusion.waf_exclusion_inline_rules

<a id="canonical-0011011011112311-3302132133230202-0303221210330032-3123333313220001-0002232101233332-3233223233201001-1221202201313213-1002023021110031"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3221233320213301-0201010232232010-2103333111322313-0013133100330102-3201203102313023-1222301203133102-3010120113313313-3020221302101303"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules`

- [rules](resources--http_loadbalancer--reference--group-028.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102): complete subsection reference.

<a id="canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-028.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="canonical-1102113331130131-2023231231001111-0002223222222100-1112210131212010-3122031212021203-0032031100221221-1011022122301233-2302320121123031"></a>

Type: `"object"`. list nested block, Optional.

An ordered list of WAF Exclusions specific to this Load Balancer.

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

<a id="canonical-3320202020230112-3113331300323122-1120023133202003-1222132103311311-3030211010223120-3223000300303002-1122002033100011-3203233130310300"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules`

- [any_domain](resources--http_loadbalancer--reference--group-028.md#canonical-1233330223321302-1130120332310310-0203100300123123-2130331311030110-0132230321210230-2132221103000313-1003202111301303-2033023223131322): complete subsection reference.

- [any_path](resources--http_loadbalancer--reference--group-028.md#canonical-0310132112001312-0121211110302132-3301100330233002-0032202211303222-2023031010201303-0213032213310021-2210333223313222-3110023101111231): complete subsection reference.

- [app_firewall_detection_control](resources--http_loadbalancer--reference--group-028.md#canonical-0203003232313033-0223230331033131-3221113013002333-3200303010013212-3202033211033320-3210102132113031-2230233013221023-0302302302002122): complete subsection reference.

<a id="canonical-1321111013100100-2122003310012003-3100010203122031-1021322011133013-2333322102330331-0213021131312100-0112120103212213-1032301311032130"></a>

<a id="canonical-1203210322312323-1221101012021112-1110312201330130-1212321333132302-0333123012000022-1100002231001222-0110211330320230-1032221310000311"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0003010001000230-2300032003022212-1012132133012033-2303231020221301-2301203210221211-2002031022001113-1113032012003220-2120210221032332"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.expiration_timestamp` property

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

- [metadata](resources--http_loadbalancer--reference--group-028.md#canonical-3120302131000211-1331213131113022-3323221113300010-2300303303032223-3201231223203103-1003211220131202-2033111223021311-0032101121231310): complete subsection reference.

<a id="canonical-0033221133033211-1110103031332331-1211131223320320-2212123102122212-2122300330033321-2112000333010122-3332002111011233-2231313232201310"></a>

<a id="canonical-3230110323211131-0331000103210230-0222120321012301-2323231103120123-2203210013031223-1330013220223212-3223011202000313-2101131322311333"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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

<a id="canonical-2123013023333231-1202301003331220-1102231100201232-3322300100032031-1021231122002211-2130100100213103-3013203321021322-2132022121013030"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.path_prefix` property

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

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

<a id="canonical-1213313311330020-2312210112020023-2233001011102003-0231003003231032-2132003000122133-3102122330101023-1313120331230310-3113012230031002"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.path_regex` property

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

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

<a id="canonical-2203000222300111-1122232001131211-2001330331103202-1202300221200210-2233013031001110-2312233033031213-2130013303222022-3002113303123233"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1233330223321302-1130120332310310-0203100300123123-2130331311030110-0132230321210230-2132221103000313-1003202111301303-2033023223131322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-028.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-028.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_domain

<a id="canonical-0311122320020033-1323120031123203-3303311123100230-0133210123310132-3022332103101210-0023100321310201-3300233122201100-2030213233031300"></a>

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

<a id="canonical-0310132112001312-0121211110302132-3301100330233002-0032202211303222-2023031010201303-0213032213310021-2210333223313222-3110023101111231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.any_path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-028.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-028.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_path

<a id="canonical-1311002003222122-2323213130230332-2312201320110000-2320003110301032-3211113120223112-1121221122200331-3202100322031112-2321220202231032"></a>

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
any_path = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203003232313033-0223230331033131-3221113013002333-3200303010013212-3202033211033320-3210102132113031-2230233013221023-0302302302002122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-028.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-028.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
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

<a id="canonical-3123202110103233-2311000013322210-3322213002120110-0213313211101033-1023013331201233-1300212102231031-0210220313013232-0321310220301003"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`

- [exclude_attack_type_contexts](resources--http_loadbalancer--reference--group-028.md#canonical-3201121111311030-2123122000311120-1220032201333203-1111001200333213-0231022113010030-2103210133201311-0312222131220013-3312320332301320): complete subsection reference.

- [exclude_bot_name_contexts](resources--http_loadbalancer--reference--group-028.md#canonical-0013000213203121-1302233212102023-3311321210110330-3303021000002211-3132312122212322-1032232001223211-2123112213300201-3010301331100133): complete subsection reference.

- [exclude_signature_contexts](resources--http_loadbalancer--reference--group-028.md#canonical-0333233003310312-0330131131103102-0000033020223132-0322311302000323-1330000332103012-0223131021220312-2000301102133300-1220301113220030): complete subsection reference.

- [exclude_violation_contexts](resources--http_loadbalancer--reference--group-028.md#canonical-2223300030211001-0132110221002313-3233013321223313-2012010332231223-3120133023113023-2011221200123212-1220022320210111-1133313122302030): complete subsection reference.

<a id="canonical-3201121111311030-2123122000311120-1220032201333203-1111001200333213-0231022113010030-2103210133201311-0312222131220013-3312320332301320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-028.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-028.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--http_loadbalancer--reference--group-028.md#canonical-0203003232313033-0223230331033131-3221113013002333-3200303010013212-3202033211033320-3210102132113031-2230233013221023-0302302302002122)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-2320221103120003-1211103102320113-0331000213210111-0102131210320110-3012000121221303-3203303111111123-1301312210122302-3032030113311121"></a>

Type: `"object"`. list nested block, Optional.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_attack_type_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210002013332211-3013302032203332-0120031111132131-1212312123120002-0022231120100103-1103210311222300-3010030113122323-1030320131032131"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts`

<a id="canonical-0120103133020223-2310132312013232-3231203130221132-1211203012011222-3122302003212030-0102001130222312-3333031220102302-2210131202122113"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.context` property

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2013231013203130-1023020313213202-3121122223301323-0123033221220133-0033233311122022-3131201010312002-0232133101222220-2312123122330133"></a>

<a id="canonical-2012210023122222-0213120030030020-3213310022112121-1331132303201133-2033321122012212-3213100311333230-0203112312230313-1311110303323020"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name` property

Type: `"string"`. Optional.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0222222133012021-1000021000230201-2031101332312310-0231320232211323-2122010320001233-0103031313231222-3120031323320101-3223210001000321"></a>

<a id="canonical-3230022021022201-3120200203323210-1303030123323012-2013231312111333-0320213313022302-3230123212300002-2333133300123201-0200333210222020"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` property

Type: `"string"`. Optional.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Additional upstream details:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0013000213203121-1302233212102023-3311321210110330-3303021000002211-3132312122212322-1032232001223211-2123112213300201-3010301331100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-028.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-028.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--http_loadbalancer--reference--group-028.md#canonical-0203003232313033-0223230331033131-3221113013002333-3200303010013212-3202033211033320-3210102132113031-2230233013221023-0302302302002122)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-1011031200323033-3331322102200313-0100220230013302-1033023231101023-3033312011012022-2122011223012102-3200131021120102-1012111010202311"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_bot_name_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333030132222022-3222202333232312-1231023012320302-0030231023022200-0000030321012203-1123122313110101-3323212300332103-3023300033100032"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts`

<a id="canonical-1003211131311221-1230321121122000-2132031110233223-0202102011320213-2012232231300031-1220233123331013-1312000001322103-1132321010231101"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` property

Type: `"string"`. Optional.

Bot Name. Human-readable name for the resource

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0333233003310312-0330131131103102-0000033020223132-0322311302000323-1330000332103012-0223131021220312-2000301102133300-1220301113220030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-028.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-028.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--http_loadbalancer--reference--group-028.md#canonical-0203003232313033-0223230331033131-3221113013002333-3200303010013212-3202033211033320-3210102132113031-2230233013221023-0302302302002122)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-3023120302030310-1333331101113132-2120201331232111-1123221131122221-1320303303332233-2111210330202331-2101002020200223-0210201300110313"></a>

Type: `"object"`. list nested block, Optional.

Signature IDs to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_signature_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312033303230233-1100020101203133-0113313012000203-3112330021331223-2220011320221033-3201331312112031-3323200020000231-1303232012230102"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts`

<a id="canonical-1311211130033002-0221000030002220-1013311013312133-3331133300020123-1031103103001120-3200301200030100-1021131300022030-0230021113110032"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.context` property

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1200033113013023-3103332021121312-1033102201213313-1202202033200123-1020010210311310-0130213012223322-1120032232321001-3022320001223212"></a>

<a id="canonical-0122001103111202-0333030210222332-2103023301322333-3302130222202332-1110021222001211-1002010102303301-3330101100132200-2202000213221021"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.context_name` property

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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

<a id="canonical-0200131333210133-1020212323111301-3130301302100002-0001203311323113-1031000010232310-3020200033031223-1031120123331221-2023021332200112"></a>

<a id="canonical-2130132020030023-2311222333332221-1320011213132100-3102300110120203-0222311010311301-3233130232001310-0232001112303033-0302020032310302"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.signature_id` property

Type: `"number"`. Optional.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-2223300030211001-0132110221002313-3233013321223313-2012010332231223-3120133023113023-2011221200123212-1220022320210111-1133313122302030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-028.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-028.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--http_loadbalancer--reference--group-028.md#canonical-0203003232313033-0223230331033131-3221113013002333-3200303010013212-3202033211033320-3210102132113031-2230233013221023-0302302302002122)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-1032013130333122-0212323023231111-2332230103033230-0010212220010132-0011212300323313-2010321022221220-0010122211013030-1131232231332302"></a>

Type: `"object"`. list nested block, Optional.

Violations to be excluded for the defined match criteria.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_violation_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122312032110101-3310012010233222-3121030212301201-3112100331133001-2101102122223301-2232013303101202-0301002022222331-0133000201132122"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts`

<a id="canonical-3101131121213210-3300130113313023-3120200030011112-1203101232112020-2201323032010212-3233210322033321-2022332323330333-0303222210212100"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts.context` property

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1323112213131130-0210301100211113-2312230210110110-1221220001011312-0331223000210031-0302312131100103-0201123332103222-0020123300200233"></a>

<a id="canonical-0330303121323312-3110123323320231-1313202003132111-1233010133001302-2030013120233013-0232131312113030-1123322333102233-1122112033313230"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts.context_name` property

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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

<a id="canonical-0213013011312313-3221213232113032-2222120131313110-2000010021322000-2130220023233303-0123003032110022-3331003300201130-3232031000303022"></a>

<a id="canonical-0303012003312112-2031232123001203-0113202201312021-0333002013020030-2022120323222220-0212103333110020-3303210003220210-1201331212230221"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` property

Type: `"string"`. Optional.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Additional upstream details:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3120302131000211-1331213131113022-3323221113300010-2300303303032223-3201231223203103-1003211220131202-2033111223021311-0032101121231310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-028.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-028.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- waf_exclusion.waf_exclusion_inline_rules.rules.metadata

<a id="canonical-2220103033003212-1332130212310321-1012312231323102-3332313332321130-0223133101002321-2032123001213112-1332230010320123-3323103302300133"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123313112221301-2023011233321223-2123020213133323-0321223001202310-1111010133201220-3112112103121201-3333232231200321-1311300332121231"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.metadata`

<a id="canonical-2333032213213100-2333300223233331-1013213302233120-1113313012200130-3300030133302312-1030322011121231-3212331231023232-3310231333122302"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-0033313322313021-2303231221130001-3002130103123132-2000302032213001-3303032030301122-1021100010112031-2232021213213233-2222123302233222"></a>

<a id="canonical-1130313133022332-3031220331103323-1231003202120131-1100102103030002-1221211303200312-3223123112000022-0033032111032302-0102133102333320"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1202001010122130-1232112232211102-0102102132023332-3020023330320231-3120312230233312-0111230122132322-0202011203322330-2223113000302022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- [waf_exclusion.waf_exclusion_inline_rules](resources--http_loadbalancer--reference--group-028.md#canonical-2101010121302102-0101122133300022-1331321313303211-1133022130232322-3010003003012310-3011321330022331-0312321031220123-3031300100331210)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--http_loadbalancer--reference--group-028.md#canonical-1301221002011203-1101100220032312-0102211011011102-0031331031113320-3322200223031003-1310300212102323-3331021112201113-3101233011231102)
- waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing

<a id="canonical-2103331220022032-1303010102302121-3012001201200032-2322133213002220-2200112100332331-2021001210131233-1332322132013133-2303021032100233"></a>

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
waf_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300231203030031-1332023310203311-2011202020030201-3003033103130323-3301122311203303-1021210131330310-0012111123120220-1030113003313033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101)
- waf_exclusion.waf_exclusion_policy

<a id="canonical-3103213001110000-2300333222120102-0301100113022213-2010111332313102-3302200112113110-0102131201033031-0201110021302021-3213232131103113"></a>

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
waf_exclusion_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023000310333012-2010130020333100-3002330112122013-2200030113130103-3320101013021102-3300020011230321-2013203113300100-1333111212013210"></a>

### Direct properties for `waf_exclusion.waf_exclusion_policy`

<a id="canonical-1133001222213233-1131031202331200-2300212112020303-2031003200020131-3121332133120320-2212310300032100-0320112030023301-3301310031231201"></a>

#### `waf_exclusion.waf_exclusion_policy.name` property

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

<a id="canonical-0213300212012230-0331013120103222-0323220202202312-1001303031103013-2103313330223233-2103112223032203-1200301223212102-3123320300100023"></a>

<a id="canonical-0222113023321331-3121200212231031-3102111323221312-1133203120303323-1131320132202033-2021312233222231-1001300301022313-0202101033112103"></a>

#### `waf_exclusion.waf_exclusion_policy.namespace` property

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

<a id="canonical-1013311010201121-2020010103011113-1330323222321131-0123320302232200-2111310111130311-1303033210000310-0110120303003203-3130000201031310"></a>

<a id="canonical-3320100100021012-0001323013312133-3330233303332122-0300302333320131-2011230300231310-2202331333332133-1332201123333123-2120333210103302"></a>

#### `waf_exclusion.waf_exclusion_policy.tenant` property

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
