---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1013101020301211-0300113323013020-1012211223002133-1221021010320123-0201111110312321-3122000023212320-3102123312313202-3203312200103233"></a>

## `default_pool.origin_servers.consul_service.site_locator.site.tenant` property

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

<a id="canonical-0133302031203232-3121110331002233-0133300101033002-1311320120221330-2330211033311312-3111021032210110-3212310010101333-1232332132302231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-2020023111021302-1200212320111301-0102000200230210-2010021200130100-2000133311311320-2220000303333020-0002102320120101-3200312012201003)
- [default_pool.origin_servers.consul_service.site_locator](resources--http_loadbalancer--reference--group-015.md#canonical-1323111100223303-1312011332301001-3103112212130120-1123302201120002-2201130210332220-3203331113133221-3111311301322111-1101213123011332)
- default_pool.origin_servers.consul_service.site_locator.virtual_site

<a id="canonical-1010330112120130-2120210010302032-2230021300032203-0200331032320112-0320110120010231-1323303003110231-3231021131012020-1002132310000010"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033233033310033-3012023302020132-2120112331232131-2133000111120103-0023302131301132-0101211011332131-3132112000331100-2203332033223003"></a>

### Direct properties for `default_pool.origin_servers.consul_service.site_locator.virtual_site`

<a id="canonical-1221300210130230-0322202303012020-0021033311032303-2111211001020002-0212103320130101-3023223203102213-2103323202131000-0312313232211021"></a>

#### `default_pool.origin_servers.consul_service.site_locator.virtual_site.name` property

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

<a id="canonical-2312220231000012-1103233323120132-2131332130213223-2003010001231032-2333100322100123-3103212111132222-2120100210022103-3221021203202333"></a>

<a id="canonical-1221103021123023-2112210220020202-0103220332200102-3112011100200100-3023211102321013-0232213131333312-0200230110022113-1203302102102020"></a>

#### `default_pool.origin_servers.consul_service.site_locator.virtual_site.namespace` property

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

<a id="canonical-3100313222210030-3011102231111132-0312013102122011-2211101002330110-2013202222001023-2033013203232112-2222203103132313-0201223210102231"></a>

<a id="canonical-2232010003332321-2021011022020132-0022033310232330-0311233002132120-0310323231032303-0110111103101120-3031230220002233-2221201233002311"></a>

#### `default_pool.origin_servers.consul_service.site_locator.virtual_site.tenant` property

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

<a id="canonical-1203200311313231-3123130203231211-0212312001001320-2000333120210300-0003011112231121-2210223300312313-2213133112232123-2212023301012001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-2020023111021302-1200212320111301-0102000200230210-2010021200130100-2000133311311320-2220000303333020-0002102320120101-3200312012201003)
- default_pool.origin_servers.consul_service.snat_pool

<a id="canonical-1130002121200030-1210331321012030-1101033102130212-2302323320332020-2113310120230120-0221122133223322-2102130321110331-0001202023212012"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-0233311120100011-2102223111032200-2232021321010011-2112001313103120-2223330123223000-2201033313023233-0010023121211031-1320230033030313"></a>

### Direct properties for `default_pool.origin_servers.consul_service.snat_pool`

- [no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-2301003111221220-1013111013322223-1111022020003311-0333002213101120-0020010010223201-1231300022231210-3333012213310323-2023133312210021): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-0333102310112313-2200020111021211-3323021323313103-0203103011312223-3303300232010221-1333120302212132-2110213221332113-3013312020122112): complete subsection reference.

<a id="canonical-2301003111221220-1013111013322223-1111022020003311-0333002213101120-0020010010223201-1231300022231210-3333012213310323-2023133312210021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-2020023111021302-1200212320111301-0102000200230210-2010021200130100-2000133311311320-2220000303333020-0002102320120101-3200312012201003)
- [default_pool.origin_servers.consul_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-1203200311313231-3123130203231211-0212312001001320-2000333120210300-0003011112231121-2210223300312313-2213133112232123-2212023301012001)
- default_pool.origin_servers.consul_service.snat_pool.no_snat_pool

<a id="canonical-2021223113030031-1110221031332332-3001122220211012-1013103123013332-3313111221213020-1110103102002310-1121000130201202-3330032032023020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333102310112313-2200020111021211-3323021323313103-0203103011312223-3303300232010221-1333120302212132-2110213221332113-3013312020122112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.consul_service.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-2020023111021302-1200212320111301-0102000200230210-2010021200130100-2000133311311320-2220000303333020-0002102320120101-3200312012201003)
- [default_pool.origin_servers.consul_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-1203200311313231-3123130203231211-0212312001001320-2000333120210300-0003011112231121-2210223300312313-2213133112232123-2212023301012001)
- default_pool.origin_servers.consul_service.snat_pool.snat_pool

<a id="canonical-0023100133110100-1030210002100220-1132133221023000-0012233223130233-1221111111010322-0132032330301322-0322100210333323-1301110301102312"></a>

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203231132230122-0000300303033020-0312210032123201-1303122332112220-1220121102113203-1113230203310111-0010232033302222-0202120210223133"></a>

### Direct properties for `default_pool.origin_servers.consul_service.snat_pool.snat_pool`

<a id="canonical-3001223031222320-2032103220313332-2203203011212113-0032131000002323-1101022020231201-1022333323222101-1210330201233230-2331301033302233"></a>

#### `default_pool.origin_servers.consul_service.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-3320101200101333-0101332020212110-3023110121121203-0211302010232331-0301203130200233-3221233212020201-1322313330211100-1110201120130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.custom_endpoint_object` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.custom_endpoint_object

<a id="canonical-2331023211122210-1200120121113323-2003331303312332-0010022011110030-0012031320013223-0113001010323101-3133031330102311-3110223022212210"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with a reference to endpoint object.

Receipt-pinned upstream constraints:

```json
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
custom_endpoint_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002033201123220-2330202023021221-3030003201330131-1222233220022320-1331010020110223-3101133203303321-0220121003332132-0012023132213313"></a>

### Direct properties for `default_pool.origin_servers.custom_endpoint_object`

- [endpoint](resources--http_loadbalancer--reference--group-016.md#canonical-0001211020103113-3111202302021333-0000212031000100-0120322133302003-0122131200013113-1331113023220013-0220023121220103-1312230312211202): complete subsection reference.

<a id="canonical-0001211020103113-3111202302021333-0000212031000100-0120322133302003-0122131200013113-1331113023220013-0220023121220103-1312230312211202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.custom_endpoint_object.endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.custom_endpoint_object](resources--http_loadbalancer--reference--group-016.md#canonical-3320101200101333-0101332020212110-3023110121121203-0211302010232331-0301203130200233-3221233212020201-1322313330211100-1110201120130123)
- default_pool.origin_servers.custom_endpoint_object.endpoint

<a id="canonical-2222301201123110-3200302321130321-3000030310102132-2212320311202121-3022312131333123-2330201320030220-0130200302200232-3220022121121012"></a>

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
endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232122323303201-3213023123320011-3323122312131120-1002130320121311-1110120020032030-2330311131311133-1220323210312111-1122200120102113"></a>

### Direct properties for `default_pool.origin_servers.custom_endpoint_object.endpoint`

<a id="canonical-2033010031031323-0030032211113321-3112133211121301-2220211010011222-2013130210321112-3120212300332211-3302230222112212-3332102322323333"></a>

#### `default_pool.origin_servers.custom_endpoint_object.endpoint.name` property

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

<a id="canonical-0132211113130202-0023222101211330-0332030312103201-3211231323030111-3303012001011233-3203233123001012-1010322132101333-0200223300203333"></a>

<a id="canonical-1310302032230122-1100221223003113-0103000112133003-3213122221013201-2033130100001001-1213000021212223-3001210132321013-2200121013231202"></a>

#### `default_pool.origin_servers.custom_endpoint_object.endpoint.namespace` property

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

<a id="canonical-2210331020220320-2310231120233021-2101301032230031-0320223313103221-3213120332213300-0131021122311031-3113220001130100-0202103212012003"></a>

<a id="canonical-0001203113330212-0132122133200121-1101011102321223-2100311301131130-2033030021312302-3101001010032120-1222301002030311-2213121303113112"></a>

#### `default_pool.origin_servers.custom_endpoint_object.endpoint.tenant` property

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

<a id="canonical-0100230210321023-0023121030121111-1331103102003302-1011001233230211-3322100123012322-1333212023202112-1222312332312123-3210332033123213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.k8s_service

<a id="canonical-3113330000322021-1200112103200332-0212003213200313-3012203311302123-3303212231121333-3332202211323223-2001131020110031-2201331033003201"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with K8s service name and site information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "vk8s_networks"),
  validators.ConflictingObjectAttributes("outside_network",
    "vk8s_networks")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

Terraform syntax:

```terraform
k8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011230000300123-1012110123013212-2031211101112110-2013332120303202-2311023230210130-3132120121333233-0333110332331000-3310212213022320"></a>

### Direct properties for `default_pool.origin_servers.k8s_service`

- [inside_network](resources--http_loadbalancer--reference--group-016.md#canonical-0130001232102003-0310113003003323-3003301123011231-2020132332310021-2013312121322130-0112012213332223-2213213130122302-2313330001330301): complete subsection reference.

- [outside_network](resources--http_loadbalancer--reference--group-016.md#canonical-3033012011323300-2132203000313001-0331203322112033-1221331212320120-0010222120332212-0322213021202321-3113113320031203-3132333111231011): complete subsection reference.

<a id="canonical-1233202230320011-1233220100213012-2233331100113311-3203222321210321-3001202300003301-1200102201103203-0330233221101332-3302310111211021"></a>

<a id="canonical-2012131022131103-0320233122220002-1003022311123123-0331320001202120-3120111323020112-2212211223012321-1322002103223312-1110230121323331"></a>

#### `default_pool.origin_servers.k8s_service.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["PROTOCOL_TCP","PROTOCOL_UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0033232113012230-0112033302331023-2112030231322033-3332013103022212-1031103233111200-1031031231021220-0330331023230023-2102021230133120"></a>

<a id="canonical-0213311300120003-2002313202001033-3012300120301101-1222010103000100-2232120231322323-1220122121332331-2111132213120123-1101230120113202"></a>

#### `default_pool.origin_servers.k8s_service.service_name` property

Type: `"string"`. Optional.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Additional upstream details:

For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-2303320310331022-1013103300011103-1202111113213321-3233310010033022-0313121123312132-3022310103213022-1202130201022220-2200033313323332): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-2301203302113232-3211031310211322-3111320331112100-1303312113222221-0231203303112033-0220333300001100-2202001011012031-0130213321022303): complete subsection reference.

- [vk8s_networks](resources--http_loadbalancer--reference--group-016.md#canonical-3221311010131322-1132330333333330-3223002022000132-1131330101022231-1102233303110203-2331313002202223-3213323333032020-2223312030303333): complete subsection reference.

<a id="canonical-0130001232102003-0310113003003323-3003301123011231-2020132332310021-2013312121322130-0112012213332223-2213213130122302-2313330001330301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.inside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-0100230210321023-0023121030121111-1331103102003302-1011001233230211-3322100123012322-1333212023202112-1222312332312123-3210332033123213)
- default_pool.origin_servers.k8s_service.inside_network

<a id="canonical-0200133120021300-1313031100203010-1121121113003301-0102032313101010-3002122131312300-2302133033313331-1323020303110300-0132310323322202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033012011323300-2132203000313001-0331203322112033-1221331212320120-0010222120332212-0322213021202321-3113113320031203-3132333111231011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.outside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-0100230210321023-0023121030121111-1331103102003302-1011001233230211-3322100123012322-1333212023202112-1222312332312123-3210332033123213)
- default_pool.origin_servers.k8s_service.outside_network

<a id="canonical-2110133013003230-0031023012321303-2211122112332331-1313102312200002-3011012230013122-1231120211013111-2231121301200002-2300311010200223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303320310331022-1013103300011103-1202111113213321-3233310010033022-0313121123312132-3022310103213022-1202130201022220-2200033313323332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.site_locator` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-0100230210321023-0023121030121111-1331103102003302-1011001233230211-3322100123012322-1333212023202112-1222312332312123-3210332033123213)
- default_pool.origin_servers.k8s_service.site_locator

<a id="canonical-0232131233210112-2120303211221021-3220121222023212-1202012213223100-1222302022122010-2002320021101320-3331222231320012-1022003322303031"></a>

Type: `"object"`. single nested block, Optional.

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000320310110213-2211313333002223-1111311033012320-2202003112122300-3302002102323221-2022000331321110-1031121132020000-0231101322203322"></a>

### Direct properties for `default_pool.origin_servers.k8s_service.site_locator`

- [site](resources--http_loadbalancer--reference--group-016.md#canonical-3201300232021022-0230213322313133-2013110112203132-1132320121300123-1313110310332000-2330220202211013-3303211232122321-1320011110120201): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--reference--group-016.md#canonical-2122302300111132-1311021312031020-3330300323233212-0123121210020200-0201011203122100-3012323131110311-0121101103023321-1010131221211100): complete subsection reference.

<a id="canonical-3201300232021022-0230213322313133-2013110112203132-1132320121300123-1313110310332000-2330220202211013-3303211232122321-1320011110120201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.site_locator.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-0100230210321023-0023121030121111-1331103102003302-1011001233230211-3322100123012322-1333212023202112-1222312332312123-3210332033123213)
- [default_pool.origin_servers.k8s_service.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-2303320310331022-1013103300011103-1202111113213321-3233310010033022-0313121123312132-3022310103213022-1202130201022220-2200033313323332)
- default_pool.origin_servers.k8s_service.site_locator.site

<a id="canonical-2210203300103213-2000310303011201-2220032322212011-1133212213212110-1320213310310310-1311322200323203-2331120222312010-1010103120203131"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300222011320110-3232222021211022-1203003121212321-3110310313113120-0103023012002030-1011230122330113-0300023210103112-0231331102323320"></a>

### Direct properties for `default_pool.origin_servers.k8s_service.site_locator.site`

<a id="canonical-1032333232221100-2121330131231002-3222001123320031-1312203320000210-1110030321033223-0230102331310113-0000021323003112-0102310211021103"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.site.name` property

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

<a id="canonical-0333021001023233-3201001312300300-2322010213232321-0333132120031013-0101322323320110-2230232203210233-1023332133001320-2010221301202032"></a>

<a id="canonical-3000111030300230-1331101011112010-2200003112230220-2000332133100311-1003032112010213-0330300221321303-3011311033122031-0223032023103323"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.site.namespace` property

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

<a id="canonical-1312030311112100-2230331020223321-3312021131120022-2020121203003233-0332313022012233-1301322110012222-0302122011033101-3312130130313301"></a>

<a id="canonical-1210313132002002-0313313112221212-0133201203003222-1122112000111212-1121301023222222-2312302033220213-0032132023310221-0122102302122122"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.site.tenant` property

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

<a id="canonical-2122302300111132-1311021312031020-3330300323233212-0123121210020200-0201011203122100-3012323131110311-0121101103023321-1010131221211100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-0100230210321023-0023121030121111-1331103102003302-1011001233230211-3322100123012322-1333212023202112-1222312332312123-3210332033123213)
- [default_pool.origin_servers.k8s_service.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-2303320310331022-1013103300011103-1202111113213321-3233310010033022-0313121123312132-3022310103213022-1202130201022220-2200033313323332)
- default_pool.origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-2021021311113332-3302332230000012-0112310331101231-1303233112222202-3331120120213323-0332310113101323-0030311100222330-3131220233103313"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033001233213011-0231321132231312-3331103110130033-1003201102111322-1220113022303220-3111000110120020-0331312033321100-3311102232211023"></a>

### Direct properties for `default_pool.origin_servers.k8s_service.site_locator.virtual_site`

<a id="canonical-3002030002222301-3233032011102311-3323112110233322-2310310010201211-1303312023211110-1113102110103330-0322100303132020-3101222210012333"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.virtual_site.name` property

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

<a id="canonical-0220301310021213-0113301230021231-0032020303100103-3100031032200022-1322203111010230-3200010001321022-0301110111230032-1323213222003120"></a>

<a id="canonical-0101330020020211-3022233210233303-2321031222020221-3330323020131321-1003012323231131-1011131110000003-0000301201130021-0130120101110221"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.virtual_site.namespace` property

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

<a id="canonical-0212123110323310-0123211232223302-1110112213200233-1121012231103320-0313300202330113-1333013221122102-1331200132221313-1132322123303302"></a>

<a id="canonical-1220120021330232-1210230301123323-2312232221122033-0100131013310112-3031212202102313-2201031300322133-3012323202210300-0123102012200310"></a>

#### `default_pool.origin_servers.k8s_service.site_locator.virtual_site.tenant` property

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

<a id="canonical-2301203302113232-3211031310211322-3111320331112100-1303312113222221-0231203303112033-0220333300001100-2202001011012031-0130213321022303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-0100230210321023-0023121030121111-1331103102003302-1011001233230211-3322100123012322-1333212023202112-1222312332312123-3210332033123213)
- default_pool.origin_servers.k8s_service.snat_pool

<a id="canonical-3221201100111020-3120030221200121-2101122033313301-1232021313320122-3313130130021003-2222021221313123-3003200021033010-0231012131200331"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320311102202003-2210301313023203-0330213031031101-3333312010000303-3200221213033211-2023103301321021-0312330231300220-2013231120323220"></a>

### Direct properties for `default_pool.origin_servers.k8s_service.snat_pool`

- [no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-1111320223302200-0330212323101202-1220220210100210-1223130010122030-0200001313233000-2033303122203313-0111212110200310-0121133212221311): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3010123020001201-0001102220223211-1133102011213203-3203233323320013-3331330203011230-2023303333110113-1111103013111220-2112023122331113): complete subsection reference.

<a id="canonical-1111320223302200-0330212323101202-1220220210100210-1223130010122030-0200001313233000-2033303122203313-0111212110200310-0121133212221311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-0100230210321023-0023121030121111-1331103102003302-1011001233230211-3322100123012322-1333212023202112-1222312332312123-3210332033123213)
- [default_pool.origin_servers.k8s_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-2301203302113232-3211031310211322-3111320331112100-1303312113222221-0231203303112033-0220333300001100-2202001011012031-0130213321022303)
- default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-1010030301311001-2221331010220221-1222020213210333-0221212212120020-0002332313221013-3013311223302231-0312120100330100-0032102322233223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010123020001201-0001102220223211-1133102011213203-3203233323320013-3331330203011230-2023303333110113-1111103013111220-2112023122331113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-0100230210321023-0023121030121111-1331103102003302-1011001233230211-3322100123012322-1333212023202112-1222312332312123-3210332033123213)
- [default_pool.origin_servers.k8s_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-2301203302113232-3211031310211322-3111320331112100-1303312113222221-0231203303112033-0220333300001100-2202001011012031-0130213321022303)
- default_pool.origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-3301033103120032-0330210311113211-3030300033130133-0111103302121032-0231303022321311-3122001020220321-0023121010300320-1300101311300032"></a>

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320110202003000-3012133102322003-0332023122122311-1313303331021221-0333010201332003-0000223322331210-0321030231313113-1220001013313222"></a>

### Direct properties for `default_pool.origin_servers.k8s_service.snat_pool.snat_pool`

<a id="canonical-1331300231021330-1332122130322132-1023200012331023-3013103222120133-3032232330122000-3223011322003210-0310122303010033-0300033020021321"></a>

#### `default_pool.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-3221311010131322-1132330333333330-3223002022000132-1131330101022231-1102233303110203-2331313002202223-3213323333032020-2223312030303333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.k8s_service.vk8s_networks` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-0100230210321023-0023121030121111-1331103102003302-1011001233230211-3322100123012322-1333212023202112-1222312332312123-3210332033123213)
- default_pool.origin_servers.k8s_service.vk8s_networks

<a id="canonical-0312202330101001-0332100311030222-2100231111023223-2010220012112323-0121203320200121-1002023022022333-1030331021223303-0113011301120223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vk8s networks.

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
vk8s_networks = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.private_ip

<a id="canonical-3100111233021211-3332113012310333-3012322000202123-0013120232031030-2220301230013132-1113313103210001-2113132331231202-0212220300220032"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public IP address and site information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "segment"),
  validators.ConflictingObjectAttributes("outside_network",
    "segment")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
private_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231223202102131-1003133003312033-3000222210132213-2222111301020331-2202023223110310-2230020033011033-3021023000220103-2222212222313032"></a>

### Direct properties for `default_pool.origin_servers.private_ip`

- [inside_network](resources--http_loadbalancer--reference--group-016.md#canonical-0233331230120202-0332123101000310-3003213100121032-3203321121103230-1212010332020032-1231200033220313-0111303123113022-3102311312211102): complete subsection reference.

<a id="canonical-1222323011021030-0211001233323200-1300120030320301-1011133333133301-0312223013031120-3111202100330320-0111331230230222-1213022123112212"></a>

<a id="canonical-0300102233022301-3010212221112023-3212333313003213-3113232130303311-2133332320110313-1100312111310213-3213020210203110-1011121330101023"></a>

#### `default_pool.origin_servers.private_ip.ip` property

Type: `"string"`. Optional.

IP. Exclusive with \[\] Private IPv4 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [outside_network](resources--http_loadbalancer--reference--group-016.md#canonical-1212201011023212-2333030020231320-2103223210113133-0310102230200030-2223122111313323-3020322132220113-1211032123301230-1312031131100230): complete subsection reference.

- [segment](resources--http_loadbalancer--reference--group-016.md#canonical-3311030313320110-3330101201111301-3231222303232123-3000112002131230-0001000023123101-2022133313232212-0201223031302200-2302101213203200): complete subsection reference.

- [site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-2230102320211223-1131200223233012-0310211203012233-2111111320221032-0103211330121021-0322101013013221-3221112322101322-1130322202321103): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-1002330332212233-3220333132112323-2111220000122132-3011130211310220-2321201310000030-2321112100223321-3020030021202023-0222333100302101): complete subsection reference.

<a id="canonical-0233331230120202-0332123101000310-3003213100121032-3203321121103230-1212010332020032-1231200033220313-0111303123113022-3102311312211102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.inside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021)
- default_pool.origin_servers.private_ip.inside_network

<a id="canonical-0220232322310131-1122001301322013-2333122222231301-2111310002012033-2301002021033002-3232212210222112-2200101210230212-1213311312323121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212201011023212-2333030020231320-2103223210113133-0310102230200030-2223122111313323-3020322132220113-1211032123301230-1312031131100230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.outside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021)
- default_pool.origin_servers.private_ip.outside_network

<a id="canonical-3213220032211032-2010031012222311-3203000031320103-2231102111301322-1102032231231123-0222030222313232-0232012020013221-0201222231020301"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311030313320110-3330101201111301-3231222303232123-3000112002131230-0001000023123101-2022133313232212-0201223031302200-2302101213203200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.segment` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021)
- default_pool.origin_servers.private_ip.segment

<a id="canonical-0121112232022321-3333330221322001-3123300322312301-1200023232210021-3120330231113112-1323221210000222-0323022133101222-3333210001003103"></a>

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231102010310310-0012322130220231-3300002310302131-0313003303221002-1131323122102011-3203303312230333-0333100123031013-3203023002032222"></a>

### Direct properties for `default_pool.origin_servers.private_ip.segment`

<a id="canonical-3333133331212012-0130200133102112-0323223323102000-1101121032110211-3020203330231210-1000332001323131-0221101032211312-1203223322301021"></a>

#### `default_pool.origin_servers.private_ip.segment.name` property

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

<a id="canonical-2201023032111300-2201221200123333-1130010201022122-3021130030210132-2113201210101003-1030101000100221-3102223213121223-1033103121112033"></a>

<a id="canonical-1332223210031212-0223000220013311-2302033210310320-1113210321121000-1221322013332212-2021233112202203-3113230221011020-0321311202220133"></a>

#### `default_pool.origin_servers.private_ip.segment.namespace` property

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

<a id="canonical-3312230123031113-3003010300011322-1000301123100302-3323331133213301-1123122102331330-1201223211200232-2013012301130103-0031020323033021"></a>

<a id="canonical-3212312110122302-2100201133010313-2233011333121303-2100013233133221-0133133113322332-2021220230202232-1112200000110002-3223102032023322"></a>

#### `default_pool.origin_servers.private_ip.segment.tenant` property

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

<a id="canonical-2230102320211223-1131200223233012-0310211203012233-2111111320221032-0103211330121021-0322101013013221-3221112322101322-1130322202321103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.site_locator` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021)
- default_pool.origin_servers.private_ip.site_locator

<a id="canonical-0211001322313200-2301112112312220-0222221211302012-1123133100210011-0131210302223002-2033330210222301-0331121230133302-3322332120220011"></a>

Type: `"object"`. single nested block, Optional.

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230013333200132-1203002002333332-3331320112022212-3111232103012132-2333202113221022-2300223103331232-0112300020020003-1001021331032233"></a>

### Direct properties for `default_pool.origin_servers.private_ip.site_locator`

- [site](resources--http_loadbalancer--reference--group-016.md#canonical-1120311030001030-1210032200011022-3012100032231301-0012013002101303-3011022023322030-3021330101131321-0131200013233320-2300120023112221): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--reference--group-016.md#canonical-3201230301120210-2333332123332130-0300303013203102-1100132220331033-1321003221132003-2013122120313111-3312113303132010-1322103101011311): complete subsection reference.

<a id="canonical-1120311030001030-1210032200011022-3012100032231301-0012013002101303-3011022023322030-3021330101131321-0131200013233320-2300120023112221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.site_locator.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021)
- [default_pool.origin_servers.private_ip.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-2230102320211223-1131200223233012-0310211203012233-2111111320221032-0103211330121021-0322101013013221-3221112322101322-1130322202321103)
- default_pool.origin_servers.private_ip.site_locator.site

<a id="canonical-0302222003313203-3302122212010221-0332230122120203-2222003313210010-0303320211132031-3103112200310001-2002331033233000-3322302033010011"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200200103321301-1023203013330000-3210232313032203-1300233310322310-3201311330102302-3233130131003102-3113221030213202-2001011202211031"></a>

### Direct properties for `default_pool.origin_servers.private_ip.site_locator.site`

<a id="canonical-1221332001032323-1113010111202030-3023113322000312-2112022213222212-3011101003321313-2122130010021032-0013222111103300-2000130112122001"></a>

#### `default_pool.origin_servers.private_ip.site_locator.site.name` property

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

<a id="canonical-3322100100031200-1310013001221021-2210013333210122-1301132031120122-3230233303022032-2023021110330121-1023111010330200-3000101012311201"></a>

<a id="canonical-0002020232302010-0212002330112121-2101131302213320-0230122311020311-3110222111200032-3101211223131301-2022032331212312-2320103012132300"></a>

#### `default_pool.origin_servers.private_ip.site_locator.site.namespace` property

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

<a id="canonical-3212103331230031-3222021022313113-2023123021102021-0320021320121322-2030200121310233-1201132313330300-0303123113331201-0031022132103010"></a>

<a id="canonical-2020312311301023-3233012101222313-1110312300232101-2221320031132000-1101330332311003-0220233213123320-3002302123010121-0230210213110223"></a>

#### `default_pool.origin_servers.private_ip.site_locator.site.tenant` property

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

<a id="canonical-3201230301120210-2333332123332130-0300303013203102-1100132220331033-1321003221132003-2013122120313111-3312113303132010-1322103101011311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021)
- [default_pool.origin_servers.private_ip.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-2230102320211223-1131200223233012-0310211203012233-2111111320221032-0103211330121021-0322101013013221-3221112322101322-1130322202321103)
- default_pool.origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-0212313311131100-3001230111021223-3332033003220201-0121323010022021-2302311111200312-1120310221301220-0131301110130101-3300201022323310"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321313132223203-0002122322220133-0121121131323033-1131002211022010-0321113211310133-0321331013033001-3101123030101222-3221203203313120"></a>

### Direct properties for `default_pool.origin_servers.private_ip.site_locator.virtual_site`

<a id="canonical-1013003003001322-3102221033322330-0120211200031122-1201333210120130-2330233102000120-0103103023123131-2210330010301303-3203101112310310"></a>

#### `default_pool.origin_servers.private_ip.site_locator.virtual_site.name` property

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

<a id="canonical-1021202132222033-0321222212202233-3220120131210222-0330221002222313-0223323033002011-3222033033123330-0112032222020123-0311313222133123"></a>

<a id="canonical-2330032031022000-0022323211211131-0202021213301320-0222233020012003-1321201330203301-1223101123322200-2122131001220221-2002323120030223"></a>

#### `default_pool.origin_servers.private_ip.site_locator.virtual_site.namespace` property

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

<a id="canonical-3303030330002032-2321331012211310-3133330023110030-3002003020022233-1033121130033303-0303213020102301-0020332112311022-1223213301003031"></a>

<a id="canonical-0202000021123000-2313101102100201-1211220030023311-3220231311021320-2321313202031111-1320222030110121-0222111213132211-2230030102312303"></a>

#### `default_pool.origin_servers.private_ip.site_locator.virtual_site.tenant` property

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

<a id="canonical-1002330332212233-3220333132112323-2111220000122132-3011130211310220-2321201310000030-2321112100223321-3020030021202023-0222333100302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021)
- default_pool.origin_servers.private_ip.snat_pool

<a id="canonical-1320221323311221-3203111132320012-3132100212232130-2312133121312110-2033310103202102-0213201233333203-3033023332321033-0020101102210011"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132113333000123-1033012130013011-2022130013331031-0210031221110131-0011203003302122-3132233131303133-0120202011102123-2122202103331100"></a>

### Direct properties for `default_pool.origin_servers.private_ip.snat_pool`

- [no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-0033311100123131-1310130032201010-2332031210323120-2203333222221022-0012031300120212-2323003111221233-3022211101102000-3002310123000201): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-2201312320111122-0322302231032321-2233131321021020-3222020100021013-0302312131020101-2020323002213310-1210133320301301-3313321123131122): complete subsection reference.

<a id="canonical-0033311100123131-1310130032201010-2332031210323120-2203333222221022-0012031300120212-2323003111221233-3022211101102000-3002310123000201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021)
- [default_pool.origin_servers.private_ip.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-1002330332212233-3220333132112323-2111220000122132-3011130211310220-2321201310000030-2321112100223321-3020030021202023-0222333100302101)
- default_pool.origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-2232221020000021-1300203010010022-0103232032302201-1321212100303122-1311102310223002-1303133310103213-2121002023232212-1122021112023120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201312320111122-0322302231032321-2233131321021020-3222020100021013-0302312131020101-2020323002213310-1210133320301301-3313321123131122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_ip.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021)
- [default_pool.origin_servers.private_ip.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-1002330332212233-3220333132112323-2111220000122132-3011130211310220-2321201310000030-2321112100223321-3020030021202023-0222333100302101)
- default_pool.origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-2232232113212331-3031222012203013-2001331232020330-2301130031300332-0330332112133223-0231011220300113-1222222013110122-1202103213110012"></a>

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202202323220213-3312222332030022-3311130033123332-3210323230100332-3110312321002312-0332233331200100-0123320112002323-2013301002231012"></a>

### Direct properties for `default_pool.origin_servers.private_ip.snat_pool.snat_pool`

<a id="canonical-1121302222312322-3123321122330332-0133331233023330-0312023210232212-1111203120211133-0312233101102302-3322212110100211-0321323123330021"></a>

#### `default_pool.origin_servers.private_ip.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.private_name

<a id="canonical-2221103320010002-1220031212130203-0030030200133000-1320012313121131-1112212331303210-1132311112120121-3012311302102221-1203203021133112"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public DNS name and site information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "segment"),
  validators.ConflictingObjectAttributes("outside_network",
    "segment")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]"
}
```

Terraform syntax:

```terraform
private_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120023221213120-2123303131313120-2330120122131213-0300330020133120-0212112313233313-0103101002230133-3220031310222302-1320121220000230"></a>

### Direct properties for `default_pool.origin_servers.private_name`

<a id="canonical-2321002022221332-1113231110230323-3330033033002333-3203312303313032-0312000101312312-2023012301203123-0212021222200030-0312110100322311"></a>

#### `default_pool.origin_servers.private_name.dns_name` property

Type: `"string"`. Optional.

DNS Name. DNS Name

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [inside_network](resources--http_loadbalancer--reference--group-016.md#canonical-1012330102231300-3300211120313123-1033000300003311-2033220022003310-3110013230300333-3111000023130032-1032303200312332-0223310111113132): complete subsection reference.

- [outside_network](resources--http_loadbalancer--reference--group-016.md#canonical-0311301122230130-3202322330221100-3010223302023213-3000100231023333-1000112032103002-1132120213013212-3131232033220233-0012222212112021): complete subsection reference.

<a id="canonical-1221000301200322-3330011231323302-0231220220131132-2213333022101012-1301322203313333-2333200033123022-1032031303110323-3101231123320210"></a>

<a id="canonical-1011133100002120-2332011130303320-0322122322130331-2101003201203023-2001223023223120-2220220023031333-0023103203030321-2230303203231000"></a>

#### `default_pool.origin_servers.private_name.refresh_interval` property

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

- [segment](resources--http_loadbalancer--reference--group-016.md#canonical-3322321310112123-3331203210113301-2130130132133122-2032103130113112-2102312033030203-0213230232130021-0013232233102032-3321020322030233): complete subsection reference.

- [site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330): complete subsection reference.

<a id="canonical-1012330102231300-3300211120313123-1033000300003311-2033220022003310-3110013230300333-3111000023130032-1032303200312332-0223310111113132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.inside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.inside_network

<a id="canonical-0201300111332021-1302021030321323-3333021213202233-1200021003223030-1022212320212310-3313023211131022-0133111023222130-3313012200323102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311301122230130-3202322330221100-3010223302023213-3000100231023333-1000112032103002-1132120213013212-3131232033220233-0012222212112021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.outside_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.outside_network

<a id="canonical-1303232330031210-1303120031000032-3121101321212101-3011030103331201-0300021021232131-1323122223113303-0302002200012013-2130332303030222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322321310112123-3331203210113301-2130130132133122-2032103130113112-2102312033030203-0213230232130021-0013232233102032-3321020322030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.segment` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.segment

<a id="canonical-3022021303003031-2123233030302002-2011220010333131-3022300113220132-0221301023133002-3320202022222133-2032023213100133-1133213011213000"></a>

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211333233120030-2211100212213312-3332210233112313-3131322331213002-0003303012312310-0020121310121332-3013013332323330-1223120222220121"></a>

### Direct properties for `default_pool.origin_servers.private_name.segment`

<a id="canonical-3133220022030122-2001112013033120-0301332133201113-0221312333132323-3003011133211330-2201313321201230-0030200011220110-2201102232303312"></a>

#### `default_pool.origin_servers.private_name.segment.name` property

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

<a id="canonical-3131032023101000-1101313122023232-3202002233300101-3212221023331011-0323111231002113-0301232010020211-2223302303302233-2022212302102102"></a>

<a id="canonical-1101302301003120-2223131202302022-1321121221232021-1313100303320312-3003202101303221-0012112131330203-1113100133030032-1031311033033121"></a>

#### `default_pool.origin_servers.private_name.segment.namespace` property

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

<a id="canonical-3113201210120312-0120312220002321-3103203002233112-0010313123233212-3321122300201121-2031201120003100-2303300222102011-1233211012301031"></a>

<a id="canonical-0210030033230132-0320212110023020-3231131002210033-0220211331011113-2031223312333322-3321031323030230-3103131233001213-2003131130303302"></a>

#### `default_pool.origin_servers.private_name.segment.tenant` property

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

<a id="canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.site_locator` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.site_locator

<a id="canonical-0302202221122101-1020001003021233-0032023122222011-1221020112122103-2132010320021102-2120111200010321-0300002230031221-2230323332022001"></a>

Type: `"object"`. single nested block, Optional.

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121100023202223-3112023213031201-1100112121113300-2230320101001233-0032230233231231-2203221123030003-1222202121303021-0213021322321011"></a>

### Direct properties for `default_pool.origin_servers.private_name.site_locator`

- [site](resources--http_loadbalancer--reference--group-016.md#canonical-1332313011120202-1010002020031311-2303320320310102-1231133210103131-2100011110331232-2120331112022210-3101220212213003-1023322200303033): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--reference--group-016.md#canonical-3130200003012212-2200212123310111-3112101032311330-2012320313331221-2010131132322310-0221123031111332-1110002013020130-0321020223010200): complete subsection reference.

<a id="canonical-1332313011120202-1010002020031311-2303320320310102-1231133210103131-2100011110331232-2120331112022210-3101220212213003-1023322200303033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.site_locator.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330)
- default_pool.origin_servers.private_name.site_locator.site

<a id="canonical-3011232212001132-0012120012320321-3203323001312230-1332203313101030-0122122310023210-1113102111312203-1013001220102211-1123030312100101"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212303123122001-1023332133231033-2300333001022031-3111000211232303-1200231012031331-1130102201120013-2230132022303333-3223201012001201"></a>

### Direct properties for `default_pool.origin_servers.private_name.site_locator.site`

<a id="canonical-1112010021112123-3201102202302321-0220032131233133-0310202121331023-0100030103201033-0203230023013110-0122301130111133-2332022121003203"></a>

#### `default_pool.origin_servers.private_name.site_locator.site.name` property

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

<a id="canonical-1202300301222100-2012022221333013-0130102030301210-3211110031202011-1112130231303231-3301233000303100-3021130001013120-0101123122323123"></a>

<a id="canonical-1212202213030213-3313100131131111-1122202323000010-3210121332000221-1300331302200220-1100100300300230-0333002012222133-3201203231003103"></a>

#### `default_pool.origin_servers.private_name.site_locator.site.namespace` property

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

<a id="canonical-3321233211033221-1112031222131010-0111133332112121-1020200133103212-0010332012112320-0010120322322131-1312102121302022-0321131212103032"></a>

<a id="canonical-1101211011012322-1311333112031322-0102123032221300-2203333012000131-1002230111211332-1300103213202231-2030303320210301-1301102311101322"></a>

#### `default_pool.origin_servers.private_name.site_locator.site.tenant` property

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

<a id="canonical-3130200003012212-2200212123310111-3112101032311330-2012320313331221-2010131132322310-0221123031111332-1110002013020130-0321020223010200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330)
- default_pool.origin_servers.private_name.site_locator.virtual_site

<a id="canonical-0302210233131221-3232120302303310-1133033103111201-3300011212221102-0101010023112321-2321222002132013-0223002103201020-2213012313131001"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223221322210212-3111322032302000-2121113112333322-0132022222200010-0213301012301120-2113011120303210-3302201210311012-3310102023110102"></a>

### Direct properties for `default_pool.origin_servers.private_name.site_locator.virtual_site`

<a id="canonical-3032210012121211-0020022121323101-3002112200130113-3321113021123233-0032313032222222-1311120011003331-2323100232302320-3122123332213031"></a>

#### `default_pool.origin_servers.private_name.site_locator.virtual_site.name` property

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

<a id="canonical-3200302122000310-3220202232022232-0021023220310201-3133232302030220-1132131001032020-1221023113200231-2133233212113203-2322133101011120"></a>

<a id="canonical-3101303321313233-1332132103232211-2020231332130212-3222100133131020-1123001013303132-1003210112323001-3120322212112013-0123321000123321"></a>

#### `default_pool.origin_servers.private_name.site_locator.virtual_site.namespace` property

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

<a id="canonical-2112200130122133-3333330303023123-2323223000030303-3321121222100313-3331232330003120-1300230212232230-1103221010002100-3031003233322001"></a>

<a id="canonical-0023221331103231-0221002200302031-3112321310313101-1332310112001310-0333200331123113-3132133033100020-3113313120332223-2300031200310101"></a>

#### `default_pool.origin_servers.private_name.site_locator.virtual_site.tenant` property

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

<a id="canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.snat_pool

<a id="canonical-2320301120230101-1321101213212110-0231221300332001-3023123200110103-1221133312100032-3322323102022310-0230013131131311-3113211013332331"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2201320020130103-2220032303002212-3303002110102321-1112300021103311-2112232130102210-3300020310113320-3222323021311201-1233023333320310"></a>

### Direct properties for `default_pool.origin_servers.private_name.snat_pool`

- [no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-1221311033021110-0022231221001322-3333312321200100-2220310323023132-2323323222132122-3101233013033220-0133113303032131-3110200300020302): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-0023201330313033-3012311220001023-3020110330231121-1123002133300333-0230303300110332-2203000201021020-1110113003100210-2301230331311013): complete subsection reference.

<a id="canonical-1221311033021110-0022231221001322-3333312321200100-2220310323023132-2323323222132122-3101233013033220-0133113303032131-3110200300020302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330)
- default_pool.origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-3010011311112200-0330331122233210-0203123023122111-3231232220230010-1333302213032132-3013220010023111-1320001113233330-1301233110031332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023201330313033-3012311220001023-3020110330231121-1123002133300333-0230303300110332-2203000201021020-1110113003100210-2301230331311013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.private_name.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330)
- default_pool.origin_servers.private_name.snat_pool.snat_pool

<a id="canonical-3103122222102210-1012332212123003-1330010023111300-1221023101111210-0011332012303200-1113323033323333-0303222033220103-2302113302130333"></a>

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232102210211131-0030103333320023-2110231230311201-1300310202020011-0330120131132113-3302131020301023-3220233333011021-2220002123003111"></a>

### Direct properties for `default_pool.origin_servers.private_name.snat_pool.snat_pool`

<a id="canonical-1012311200321331-1320031233130221-2302010222201100-0020311210310333-3002011300322112-0010103000322223-0323312202323330-0123310223012013"></a>

#### `default_pool.origin_servers.private_name.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-1123323200112132-2011301210333320-0230333321233220-3120200223212323-0221012302113321-1102011330332121-0221103113230111-0011130112303131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.public_ip

<a id="canonical-2233020311003121-3121333321021131-3001122003132113-1301222300022020-1120113331012303-0312322233233120-0102201110013213-2020312300030233"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320121003222032-3310311231032211-2212031120013103-1010112113121012-0202231230232310-1231023111303210-2212313121302320-3313002301322323"></a>

### Direct properties for `default_pool.origin_servers.public_ip`

<a id="canonical-1000013300233110-0012120220113212-2010132312303303-3231031102201032-2131001003003211-1012030320130230-1011300130213021-1200312020010213"></a>

#### `default_pool.origin_servers.public_ip.ip` property

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3302320302223102-1023033333230010-2222202133303121-2101111202221230-2113030322031221-1332101010000231-1202031111300230-1003313332002300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.public_name` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.public_name

<a id="canonical-1001132011011000-0202223032003200-2201310223322201-2301022130002302-2310300302010222-3101213112032003-2323032031122031-2013322101233301"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302123012033012-3230112200032102-1222323103130102-1002102230210231-1301002131301222-2033203301323333-1103323003201233-1100223202233321"></a>

### Direct properties for `default_pool.origin_servers.public_name`

<a id="canonical-0333302200003211-1210112202300202-0220113030021222-2000123111031031-3110333222010230-3000302321313102-2001003200333210-3202302130002112"></a>

#### `default_pool.origin_servers.public_name.dns_name` property

Type: `"string"`. Optional.

DNS Name. DNS Name

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
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0023123303332312-3122011222213000-0211231320220230-2331322222311001-1232100123220010-3223120300221100-2221002333213110-0212123331231301"></a>

<a id="canonical-1320331003032222-1030022331322123-1013130302320223-3130122222212223-3232002013020003-0311202330031201-1110201231220202-3032103030211312"></a>

#### `default_pool.origin_servers.public_name.refresh_interval` property

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-0022310013130030-2111303220210213-2020201222032233-1120103131300012-1111203033323012-1021113330333233-3121232112020030-0200301203202122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.vn_private_ip

<a id="canonical-2210121123312130-3103200230220021-0013001213201023-0331001313021131-0203202312200302-2012112310012022-1230221032133132-3123200303232133"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with IP on Virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_network_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
vn_private_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203231103120212-1232312001313221-2223333132133133-2303213133032222-0300130320322121-0221221001112112-2220123220301330-2221120101331320"></a>

### Direct properties for `default_pool.origin_servers.vn_private_ip`

<a id="canonical-1101003013130331-3001011200313212-0300323212213213-2311110210321110-3100023320302121-0303100212220002-3112330322123123-0310113311002322"></a>

#### `default_pool.origin_servers.vn_private_ip.ip` property

Type: `"string"`. Optional.

IPv4. Exclusive with \[\] IPv4 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](resources--http_loadbalancer--reference--group-016.md#canonical-0022111032132322-0332003132203223-2021223302101220-2331130303203021-3023123313013130-3330101201022111-1223020213011322-1133103323103121): complete subsection reference.

<a id="canonical-0022111032132322-0332003132203223-2021223302101220-2331130303203021-3023123313013130-3330101201022111-1223020213011322-1133103323103121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_ip.virtual_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.vn_private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-0022310013130030-2111303220210213-2020201222032233-1120103131300012-1111203033323012-1021113330333233-3121232112020030-0200301203202122)
- default_pool.origin_servers.vn_private_ip.virtual_network

<a id="canonical-2222011203123233-1231300123133010-1302300210111000-2103002203232133-2031123112002200-3321113323303113-0321013120333312-0303331013203100"></a>

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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031020010100311-3021333303030010-0211302300122020-0020232100111200-3233111202131312-0102002210003132-3132231212300130-2313212300331323"></a>

### Direct properties for `default_pool.origin_servers.vn_private_ip.virtual_network`

<a id="canonical-0111132101010223-0301022232122322-1311111222032122-1031021203103310-1302233303013320-3100120111011302-1323033103102332-0030131231201332"></a>

#### `default_pool.origin_servers.vn_private_ip.virtual_network.name` property

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

<a id="canonical-3321123122101032-3011112312021012-3121100022223300-3000220113012023-2223003010233133-0012211113000110-2101303132201031-2012321022013003"></a>

<a id="canonical-2231330321103330-3222021321030102-1031333231321010-0213132030313201-0012322010000322-3213133123030021-2100331330221210-0333212213212213"></a>

#### `default_pool.origin_servers.vn_private_ip.virtual_network.namespace` property

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

<a id="canonical-3113310013103302-1211100331111300-2013123333022101-0202332000302211-3333112003322321-0131111103312020-3021011210211103-3130003323100213"></a>

<a id="canonical-1022211103320111-1101210012120010-2013130133302110-3323121231031320-3131300111221203-2213220101023020-2322023222200012-3312312101032232"></a>

#### `default_pool.origin_servers.vn_private_ip.virtual_network.tenant` property

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

<a id="canonical-2101033313001230-3100221301111002-2011003233203011-1100130312221121-1002231322001133-0002023210303033-3312032213131120-2332031232122312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_name` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.vn_private_name

<a id="canonical-2220132331000203-3110331213121213-1230131221203200-3213321231111321-1213102122203002-2020001233110101-1102020200321030-2220133333332113"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with DNS name on Virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
vn_private_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223011300100011-2120100033322020-0001033103022102-0201100331322113-2113030033113200-2212021322300120-2332020021230023-2131100310313012"></a>

### Direct properties for `default_pool.origin_servers.vn_private_name`

<a id="canonical-1311303101122130-3300123300122221-3301210122222203-1012132302312301-2112320111001300-2230213032023133-0132233212221123-1131023230213030"></a>

#### `default_pool.origin_servers.vn_private_name.dns_name` property

Type: `"string"`. Optional.

DNS Name. DNS Name

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [private_network](resources--http_loadbalancer--reference--group-016.md#canonical-0203020320000112-0102231033232010-1223102323000102-2100120021102020-0320320120202210-2013203021122133-2011000000301030-2011300203301111): complete subsection reference.

<a id="canonical-0203020320000112-0102231033232010-1223102323000102-2100120021102020-0320320120202210-2013203021122133-2011000000301030-2011300203301111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.origin_servers.vn_private_name.private_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.vn_private_name](resources--http_loadbalancer--reference--group-016.md#canonical-2101033313001230-3100221301111002-2011003233203011-1100130312221121-1002231322001133-0002023210303033-3312032213131120-2332031232122312)
- default_pool.origin_servers.vn_private_name.private_network

<a id="canonical-2331213320121203-1220232222333230-1322222102013220-1112010332020110-1020313232110022-0133002120120112-0112031123001302-2113033022000322"></a>

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
private_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023200301022111-0203023133002333-2011021330101000-2021131020301130-1112032312131203-3013023122103020-1012320322113233-1231122120110233"></a>

### Direct properties for `default_pool.origin_servers.vn_private_name.private_network`

<a id="canonical-0011022111002210-1100033313003213-0132120333103120-2203231211201231-3233131011300303-1022121012131113-3101000020303213-0111213023203223"></a>

#### `default_pool.origin_servers.vn_private_name.private_network.name` property

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

<a id="canonical-1001300112001202-2021132333113100-0200010200333013-3113203220203232-3013330013000200-2221122112103213-3100220033332133-3312211111203200"></a>

<a id="canonical-1123033330210120-1101010110111212-1021210130110330-0302210113130130-0003312031222022-1101021300312200-3221222121213200-2130003330022120"></a>

#### `default_pool.origin_servers.vn_private_name.private_network.namespace` property

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

<a id="canonical-3223331230003012-0130313330232130-0230233311212200-3022210100300310-1320323311020300-0120030112030103-2030321023330030-1033123110000302"></a>

<a id="canonical-3311033330330213-0100132332301200-2322122221133201-3320022301101332-3022210133122233-2101033233323100-3022102030201123-2332322302000102"></a>

#### `default_pool.origin_servers.vn_private_name.private_network.tenant` property

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

<a id="canonical-1232100333033333-0221010121032103-2010033203301010-1223212021203320-2001102032013210-0333301222330003-1231200013113223-2010211132110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.same_as_endpoint_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.same_as_endpoint_port

<a id="canonical-2033213221230001-2133023101322313-1230200120100220-1012323122020200-2222200231023032-1012202202131121-3102330121032032-1030231322320113"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
same_as_endpoint_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.upstream_conn_pool_reuse_type` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.upstream_conn_pool_reuse_type

<a id="canonical-0123322202101000-1321100330232001-2300223133032122-2100000323122112-0020213022331320-3213210102203003-0120120302223030-2210302332110313"></a>

Type: `"object"`. single nested block, Optional.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_conn_pool_reuse",
    "enable_conn_pool_reuse")}
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
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302310033303032-3323330200011311-3000120333200003-0010033311321113-1221202201023033-0110230231300033-0121123233023220-3122102133100002"></a>

### Direct properties for `default_pool.upstream_conn_pool_reuse_type`

- [disable_conn_pool_reuse](resources--http_loadbalancer--reference--group-016.md#canonical-3022003332112203-3313121000200322-1020012101323212-1011010321303313-0220213231201133-3123213123103002-0213302001100112-0013020212002000): complete subsection reference.

- [enable_conn_pool_reuse](resources--http_loadbalancer--reference--group-016.md#canonical-0122233311301230-2033102130212202-1303001033113312-2303130331113333-0121303320101230-1213010322331113-0123121210000310-3310332303220133): complete subsection reference.

<a id="canonical-3022003332112203-3313121000200322-1020012101323212-1011010321303313-0220213231201133-3123213123103002-0213302001100112-0013020212002000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-016.md#canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230)
- default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-1033210212033123-0322132311231131-2003001201100033-1222031301130021-3020120221133033-0133021221203001-1300033210001320-3322032323211021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable conn pool reuse.

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
disable_conn_pool_reuse = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122233311301230-2033102130212202-1303001033113312-2303130331113333-0121303320101230-1213010322331113-0123121210000310-3310332303220133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-016.md#canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230)
- default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-3113313020022133-0330300231330231-3032212110201123-0231101332023132-3121312031113003-3020230103323212-3303122210032002-3213330300132320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable conn pool reuse.

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
enable_conn_pool_reuse = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.use_tls

<a id="canonical-2332213023031332-1013323323031210-0100123112302011-2330013230103010-0102132311330320-2010003302003033-2313120303330331-0312102122123221"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "use_server_verification"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("use_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("use_server_verification",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103122312033133-0232020122221200-3322322232313322-3012123330332032-0000121120031111-1021333002233122-0311211200123330-1203031110312330"></a>

### Direct properties for `default_pool.use_tls`

- [default_session_key_caching](resources--http_loadbalancer--reference--group-016.md#canonical-2202223033321103-1023313022220032-1100110331023310-2000233222211323-0101221310312021-3102332013202020-0101210122012110-2103210131213301): complete subsection reference.

- [disable_session_key_caching](resources--http_loadbalancer--reference--group-016.md#canonical-1013331020231333-0202203320300021-1232330310100320-2300031020020203-1123330312131323-3001113200312303-0122030002201222-3233033231132121): complete subsection reference.

- [disable_sni](resources--http_loadbalancer--reference--group-016.md#canonical-2300213333200132-0123123213203223-0200321201003332-1133131212130322-0120022102301322-3132311102333002-1200211200102323-3233332031012120): complete subsection reference.

<a id="canonical-1213330013021023-0213323002023222-2310113003121231-0010330232301332-1211121220311200-0331022332003023-0330221202302201-1302013110331311"></a>

<a id="canonical-2333032221123011-1312320102032030-0222130012112100-0300313100213112-3310300033310112-3300102003222320-0201120203223113-2323001120313230"></a>

#### `default_pool.use_tls.max_session_keys` property

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](resources--http_loadbalancer--reference--group-016.md#canonical-0303331032002210-0312010001012120-1330112001013310-0010320010222322-2132130013000332-1231131311012023-3021102123223331-2330222230010032): complete subsection reference.

- [skip_server_verification](resources--http_loadbalancer--reference--group-016.md#canonical-0303332132110311-1223303331210202-2030002213310231-2101021031103002-3222212230031331-2213230130330002-3303131031031310-2010310111301030): complete subsection reference.

<a id="canonical-3203103222000233-2001032322130211-2232220032202333-3332301033030323-2013113331033102-3033013031113121-0111203103331320-1122023110231123"></a>

<a id="canonical-0001021102223102-1002303213321120-2321132020021113-3001122012301201-3312032111232212-2323012221313122-2233031230223221-1213112022113001"></a>

#### `default_pool.use_tls.sni` property

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](resources--http_loadbalancer--reference--group-016.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112): complete subsection reference.

- [use_host_header_as_sni](resources--http_loadbalancer--reference--group-016.md#canonical-1322020121333230-3003002212033100-2332300001021121-1102021010223121-3131123312322320-1020132011323230-0202022322031312-3313322321233013): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-016.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031): complete subsection reference.

- [use_mtls_obj](resources--http_loadbalancer--reference--group-017.md#canonical-2011310303210021-1222202221103110-1130033232102230-1131030100322101-0032220031011132-0202012301223101-3013101110032002-0102022212323012): complete subsection reference.

- [use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-3302122303122133-0323233203310323-3012112030002300-2003211111010032-2031330012231010-1011303111010020-0113222311123003-1002223100220030): complete subsection reference.

- [volterra_trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-3321010323221300-3221232231012333-1000122333111322-3013010233100302-2311021323120212-1121301113331022-3222031002320022-3101011201003023): complete subsection reference.

<a id="canonical-2202223033321103-1023313022220032-1100110331023310-2000233222211323-0101221310312021-3102332013202020-0101210122012110-2103210131213301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.default_session_key_caching` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.default_session_key_caching

<a id="canonical-0121200021230013-1210110220011333-2010002013003130-0332331033121331-3210220321103133-0003331123200131-0031332131222032-3333001233022131"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
default_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013331020231333-0202203320300021-1232330310100320-2300031020020203-1123330312131323-3001113200312303-0122030002201222-3233033231132121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.disable_session_key_caching` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.disable_session_key_caching

<a id="canonical-3133021332013103-1320013333102132-1221130111322301-2103212230210211-0003102122023023-0211011210211020-1221123121131000-3233100211221103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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
disable_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300213333200132-0123123213203223-0200321201003332-1133131212130322-0120022102301322-3132311102333002-1200211200102323-3233332031012120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.disable_sni` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.disable_sni

<a id="canonical-2220212132201313-2030112331313111-2010033223000013-0212133231032022-3210020311230301-3211012122013333-3113202200333123-2122133012213030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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
disable_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303331032002210-0312010001012120-1330112001013310-0010320010222322-2132130013000332-1231131311012023-3021102123223331-2330222230010032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.no_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.no_mtls

<a id="canonical-1332101130102002-2013003030331203-2321313313211121-1313313112103031-3132102123231203-1113032031122220-1011112032010133-3230233321301131"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303332132110311-1223303331210202-2030002213310231-2101021031103002-3222212230031331-2213230130330002-3303131031031310-2010310111301030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.skip_server_verification` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.skip_server_verification

<a id="canonical-1212002212011012-3301133312303113-1013031220110323-3303223000111220-3012313232202033-2132132223120222-0322311011312202-0002032201302312"></a>

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
skip_server_verification = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.tls_config

<a id="canonical-3200322132320010-0012131232103001-1322321033121031-0031033101302232-1313330021103112-3102301211222010-1001032231012121-2320123131013313"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231013032332023-2323012321201032-0113233032310002-0121100300001020-2201301333031211-0032310301110012-0233111121030003-2131203211113222"></a>

### Direct properties for `default_pool.use_tls.tls_config`

- [custom_security](resources--http_loadbalancer--reference--group-016.md#canonical-2200333031323122-3012011201002322-3130330223013331-1233310103131031-2030012001303132-0303212231120001-1201102011200321-1010332223231233): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-016.md#canonical-0100203223333331-3222113310331031-3023321131313312-3310033221020010-1313130223033303-1002200231311002-3020333203131231-3220021002103110): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-016.md#canonical-0303100332303222-3233213233330122-3131223311022000-1321203230112312-0031221232100122-2031312301300013-2212123210100030-0103022302311131): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-016.md#canonical-0222121100303332-0023323120332332-0330001011121000-1210020031030130-2220031231123120-0120333002033001-2221322033313322-0131003300131300): complete subsection reference.

<a id="canonical-2200333031323122-3012011201002322-3130330223013331-1233310103131031-2030012001303132-0303212231120001-1201102011200321-1010332223231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-016.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.custom_security

<a id="canonical-1220133112231103-3133202012022230-0331123031020012-3112010330311112-0202101003110102-1213221301210002-0300333021103032-1121322030113003"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102220012202123-3321223213022213-3301212212010123-0333200120320230-1033011213022010-3321101213333321-1332021300021123-1310232010223020"></a>

### Direct properties for `default_pool.use_tls.tls_config.custom_security`

<a id="canonical-2302131311321310-2313111213002113-0100103000220313-2320303332300231-2303022003032033-3132133303200303-2222210003100101-0321130323100122"></a>

#### `default_pool.use_tls.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3021112011223320-0021132033313301-1023310221110301-3311031233211110-1301132310331000-2010122332103110-2133232211213330-1203020030100311"></a>

<a id="canonical-3112003212202120-3021113011313011-0233233221011133-1011020231013310-1231200131020013-2220221123130033-0303332321101210-0210131310221023"></a>

#### `default_pool.use_tls.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301111312203121-3301313013110313-3101201213000210-3230133310320030-0031223213103012-1130123011223330-1010032130303022-3332012300313323"></a>

<a id="canonical-2002032131210120-1031100130223013-3020001330000232-0222311122301100-0323212221310101-3002323113221332-3102230320032000-2210233221322203"></a>

#### `default_pool.use_tls.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0100203223333331-3222113310331031-3023321131313312-3310033221020010-1313130223033303-1002200231311002-3020333203131231-3220021002103110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-016.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.default_security

<a id="canonical-2001121231113323-2202302022312013-1331012232331221-2323103111030031-1130222031133131-0113313212323213-0300232212322333-3320000103001121"></a>

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303100332303222-3233213233330122-3131223311022000-1321203230112312-0031221232100122-2031312301300013-2212123210100030-0103022302311131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-016.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.low_security

<a id="canonical-0130332222331002-0011221102013332-0301322031321302-3022120020233011-2302221311111002-3302121031031122-3230332233003002-0003201330112311"></a>

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
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222121100303332-0023323120332332-0330001011121000-1210020031030130-2220031231123120-0120333002033001-2221322033313322-0131003300131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-016.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.medium_security

<a id="canonical-1000302303231022-3322010013221311-0123300200102313-2202220003120212-3111311222313132-2222122030223012-1312020130221101-3112232313213203"></a>

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
medium_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322020121333230-3003002212033100-2332300001021121-1102021010223121-3131123312322320-1020132011323230-0202022322031312-3313322321233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_host_header_as_sni` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_host_header_as_sni

<a id="canonical-0313031012231211-1002132313113203-0121133132332010-3323033303133001-2320223110111303-0231130011133120-1311031200103022-3301103201002302"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
use_host_header_as_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_mtls

<a id="canonical-0012033332103102-2000112211213221-2202123011211022-1112211212300213-3323011221300011-1203333113303113-3102003221010000-1301333232022310"></a>

Type: `"object"`. single nested block, Optional.

MTLS Certificate. MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates")}
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
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302122203211220-2121332323211221-0102313012221101-0033133321201112-1003001100120130-1000100220210213-2313103302202011-2133103030320110"></a>

### Direct properties for `default_pool.use_tls.use_mtls`

- [tls_certificates](resources--http_loadbalancer--reference--group-016.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023): complete subsection reference.

<a id="canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-016.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- default_pool.use_tls.use_mtls.tls_certificates

<a id="canonical-2212133130313131-0212312110312103-1001211323113313-0122031120232010-2021312001202331-3111320320020322-2211331321031122-1112301000111222"></a>

Type: `"object"`. list nested block, Optional.

MTLS Client Certificate. MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322332012030213-2101321010323011-0323230010300033-0311312032020030-1220103001011021-3023011030213022-2013210133320030-3112131223100023"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates`

<a id="canonical-0020200123311232-3102103202333331-0002123021121203-3031120233301302-2100323232320100-2020032022301301-3100213212010200-3211120101203113"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.certificate_url` property

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--http_loadbalancer--reference--group-016.md#canonical-2020213113300220-2333110233301203-3300222130030231-2130303232320030-2032111112212201-2300222323013202-0110222032000202-1320111112111031): complete subsection reference.

<a id="canonical-2221103223010020-3100323022012230-3032130320011322-2113230031211331-1222220132311223-0211100112030330-0302212101133010-0201131020300123"></a>

<a id="canonical-0103022201121202-1302132303010331-0023323221331221-0322223330213331-3011130023133202-0113120122130133-0233013302312332-3333021001221030"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--http_loadbalancer--reference--group-016.md#canonical-1322333013020002-3320101320322330-1230010223210130-0332001200121133-0020320123030210-0012032023300221-0031101222021121-2210211232313232): complete subsection reference.

- [private_key](resources--http_loadbalancer--reference--group-016.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223): complete subsection reference.

- [use_system_defaults](resources--http_loadbalancer--reference--group-017.md#canonical-1112030331010001-1223303203031330-0333010110213321-3202130323302212-0212221311313102-1030203333133001-1312232202322201-0223303001130330): complete subsection reference.

<a id="canonical-2020213113300220-2333110233301203-3300222130030231-2130303232320030-2032111112212201-2300222323013202-0110222032000202-1320111112111031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-016.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-016.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-1030021013012233-1132023103020003-0322031310212302-1102323132200031-1330231033232303-1101231202333203-0200121200223332-0323333130022223"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223031031211033-0320311121020102-1310011323203221-2211230320301032-1332022133211022-3133010133211020-1321002202133110-2002301230121103"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms`

<a id="canonical-3331002031032101-2001130333223101-0302021313121000-3033323313301110-2233313222022113-0310102131332223-3010311101032113-2133332103300131"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1322333013020002-3320101320322330-1230010223210130-0332001200121133-0020320123030210-0012032023300221-0031101222021121-2210211232313232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-016.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-016.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-1300032021202110-0011031032313012-0300212012032212-3333302112232001-0331202031331003-3330010131203303-1201000113210300-3211001103103331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-016.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-016.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-2203210231332111-3031132312132123-1130302233133303-0110113131033132-2310332200311320-0330113002033032-3000202300132332-3111130320320313"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000020210331122-1132010321001303-2131303013231130-2312030102130212-2113301020003331-2331220300303232-2023111000331011-1122212223201201"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.private_key`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-016.md#canonical-3202131212220022-3310331222003333-0132103303101313-3122213000000300-2332013301033012-3213003113213002-3322023211002302-2102301302313033): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-016.md#canonical-3303001112302110-2223022023002212-0233233102211101-2011031233223201-2222213211201220-2121132300103222-3002120121011222-3323020013000102): complete subsection reference.

<a id="canonical-3202131212220022-3310331222003333-0132103303101313-3122213000000300-2332013301033012-3213003113213002-3322023211002302-2102301302313033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-016.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-016.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-016.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0022023210023202-2110310333003320-1012230102222210-2233113231311001-0133022031032111-3201010222100011-1032010111103100-3001101003123221"></a>

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

<a id="canonical-3022123013132032-1211200012123101-3203133100311113-1330210213232022-1131312032133021-3230003113201313-2023110023022120-1131222231323102"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-2333112213321311-2001302311302120-3231303003201322-3022311002323320-3030210223320032-1001332012312303-0112023031001221-0011200221322313"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2111132120203332-1311130030030123-3213013212012201-2031333230020311-3233301023312330-3113330012020303-2132001123033300-2113231333020013"></a>

<a id="canonical-3231112100121132-2030200201021300-2003030123133322-2003200330323323-2003312223220220-0103112113321223-2312121112000310-0333120312220303"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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

<a id="canonical-0200303010213013-1132000102112230-2003002331021311-3213033131202221-3310031302201130-2213111100122120-0022023220201031-2111201322113232"></a>

<a id="canonical-0032031131221303-0132210003012202-3333101212203313-0300302322331020-0130222203322333-0130102020222133-2303100032333132-1101022210223210"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-3303001112302110-2223022023002212-0233233102211101-2011031233223201-2222213211201220-2121132300103222-3002120121011222-3323020013000102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-016.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-016.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-016.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-3300000121331033-2030010003221332-0033231020032330-2023223021123323-1001021121030230-1033001102331233-2321102003300023-0012201032120330"></a>

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

<a id="canonical-1120230331230010-1032221231322231-2230320202013233-3123322312010201-1003002323310101-1310011332310131-1303020203311212-2003033002230001"></a>

### Direct properties for `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0200323133202323-3011021311113011-1303110212331210-1112101223331200-2031320312033333-0221222200323121-1013130021123101-0100323201211333"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3202120300223332-3022201000103233-1010230210303111-2022212302032133-3133120013120033-1032031221023220-1302011331020303-1232333130333211"></a>

<a id="canonical-2032120031030031-3103133010001123-2131200301121033-1203103231032202-2003133303003033-0232030202032332-3332331321300132-3022221221202222"></a>

#### `default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
