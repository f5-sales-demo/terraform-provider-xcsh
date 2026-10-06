---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-2032033302122022-1103111020001001-2303301333213001-3102322130322113-3113231020002223-3033131021221332-2021231222100202-3011311100002022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.outside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.outside_network

<a id="canonical-0200213003113132-2023222231331310-2230113323301202-2211222031132011-0033332200031130-2013120323322321-1330102032233232-3101213331030010"></a>

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

<a id="canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.site_locator` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.site_locator

<a id="canonical-0210110101223301-1332211132103003-2122101201000300-3232202220003030-0323132200322021-1031030230301100-3013303132331112-1301002231102332"></a>

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

<a id="canonical-3012002323010230-1113330103013311-1023222321223120-3132013013101210-3300030200123321-1310012002202012-0330220131322123-3202030322120222"></a>

### Direct properties for `origin_servers.k8s_service.site_locator`

- [site](resources--origin_pool--reference--group-003.md#canonical-3100133211320001-3312100030232320-3031301010230301-3330030203200230-0033031210223320-2010200033333010-3223011122110301-3320211111203031): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-003.md#canonical-0022302012110203-2130330132230330-3101130123013031-2332313300310132-0313113133213110-3100010223101033-2233301002233221-1303300221201021): complete subsection reference.

<a id="canonical-3100133211320001-3312100030232320-3031301010230301-3330030203200230-0033031210223320-2010200033333010-3223011122110301-3320211111203031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.site_locator.site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-003.md#canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320)
- origin_servers.k8s_service.site_locator.site

<a id="canonical-3132212002001320-1011333321223332-1320123123221213-3111323223031202-3012133121203111-0202303002033130-3131332300331001-2333322330123022"></a>

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

<a id="canonical-3122132202132123-3212130020033233-0121113011310113-1330111311000332-2300203322111001-2201310310202101-2032212023212212-0221011120221211"></a>

### Direct properties for `origin_servers.k8s_service.site_locator.site`

<a id="canonical-1231302310302222-1101203001232232-1130302320321033-0000320130103230-0001002231012230-3311130121303031-3032110323331223-1113013012233012"></a>

#### `origin_servers.k8s_service.site_locator.site.name` property

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

<a id="canonical-2000013301321012-2001230112000330-0201112100203003-1113213330132032-1300121001130102-0122300001311330-0002333212332130-0011100301221220"></a>

<a id="canonical-1031231100223010-3212110013333133-1112132300222031-0120333331310013-0202310313223323-2202313303203222-1300322110222311-2212211112031102"></a>

#### `origin_servers.k8s_service.site_locator.site.namespace` property

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

<a id="canonical-2210010001330323-0311122232312312-2111301223332213-2200102000133300-0032322003322331-3211200212121033-2333002310233103-2320202311320212"></a>

<a id="canonical-0022022201110031-2031023123203111-0031131233022312-3011211221302322-2000023112122003-2321301030302012-2201202130320033-3301213022200303"></a>

#### `origin_servers.k8s_service.site_locator.site.tenant` property

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

<a id="canonical-0022302012110203-2130330132230330-3101130123013031-2332313300310132-0313113133213110-3100010223101033-2233301002233221-1303300221201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-003.md#canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320)
- origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-2202122132021231-3321020200323200-1310202030313310-1330230003113200-2023120102022020-3101323010320321-0320133022010001-3312132023330131"></a>

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

<a id="canonical-2121303313030031-1113033223011330-2320120211223113-2220233312212013-3102100111201330-0023312113132323-1023210020233233-1321112221100130"></a>

### Direct properties for `origin_servers.k8s_service.site_locator.virtual_site`

<a id="canonical-1111333110031032-3213113231122321-3132230132020123-1123302333233320-2201113113032032-3301013313120111-0211131021101103-3011311232300232"></a>

#### `origin_servers.k8s_service.site_locator.virtual_site.name` property

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

<a id="canonical-0312030111303223-0203213101223030-1001010012130330-3010221020030021-0031302012311123-2301331321021112-1202202012301222-0111103021310012"></a>

<a id="canonical-0021011021222132-3033301003213320-3032020013223203-3302230230003023-1203220311101010-2130312120233212-1023122122323110-0031222113010101"></a>

#### `origin_servers.k8s_service.site_locator.virtual_site.namespace` property

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

<a id="canonical-2113321233311132-1131200212200302-2320031233131202-1111221313012210-0131110123231321-3201303030032220-0202212002031331-1203222023210300"></a>

<a id="canonical-3301033033133201-0112002310011101-0031211101201223-0022321200021130-2032022202030000-0303101002113032-3300233321203331-0103131323210303"></a>

#### `origin_servers.k8s_service.site_locator.virtual_site.tenant` property

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

<a id="canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.snat_pool

<a id="canonical-1313332303200303-0131030312220311-3230121111020130-1203211302210213-3301211032330300-0012033312120030-3301312322223121-1032312222131033"></a>

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

<a id="canonical-3121111310130013-3111221031010102-1232112020230330-0133123000130321-2310002201202211-1211001202202030-0303330123000311-2220033130302002"></a>

### Direct properties for `origin_servers.k8s_service.snat_pool`

- [no_snat_pool](resources--origin_pool--reference--group-003.md#canonical-3120030123300002-3302220202331111-2130210203032133-1210003121120212-1331220113313303-3023331330303323-2302302121321021-1131001030133123): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-003.md#canonical-0203100021110200-1222213130002321-1303223321201100-1222003111122102-3031102313132102-2212033122301233-3100012230012301-3221313203112130): complete subsection reference.

<a id="canonical-3120030123300002-3302220202331111-2130210203032133-1210003121120212-1331220113313303-3023331330303323-2302302121321021-1131001030133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-003.md#canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102)
- origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-3210010303322003-3021021200000033-3323212120110303-1332000331023212-1133233102021230-1322212231333012-2220131303130310-3120202011301203"></a>

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

<a id="canonical-0203100021110200-1222213130002321-1303223321201100-1222003111122102-3031102313132102-2212033122301233-3100012230012301-3221313203112130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-003.md#canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102)
- origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-2301002222010202-0121131311200100-3203331322232323-3313020223203210-2101213201012213-2113030120331002-0013111231101320-3202312001002031"></a>

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

<a id="canonical-3023301020131321-2031332013013221-1111331211230221-2212321110001203-1310123212220100-3301203131220312-0333113020013013-1130011331100132"></a>

### Direct properties for `origin_servers.k8s_service.snat_pool.snat_pool`

<a id="canonical-0232203132122110-0010012323011322-1033300302032200-2222120301013310-1310102210222222-0033302213333223-0023221332213222-1123212130320123"></a>

#### `origin_servers.k8s_service.snat_pool.snat_pool.prefixes` property

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

<a id="canonical-2010112100222333-2130310003111231-1311320001103301-0331012012231200-3121203332212212-1103022100102313-2310201311120330-2120010221201230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.vk8s_networks` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.vk8s_networks

<a id="canonical-3232303302233102-2031023011300221-3110021223111212-0022120200222200-0311213000033220-2201201100032011-0100022002301201-2221320203122020"></a>

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

<a id="canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.private_ip

<a id="canonical-0132122003323203-1113202003201132-3321320022031212-2230321230311222-2210101323121301-3303110010030232-3303110003233000-0303302222121003"></a>

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

<a id="canonical-1113211030313311-0110022301202300-1223320100321111-3311020100113322-1112333213133101-1233311310011130-1222030112123021-1001202012322201"></a>

### Direct properties for `origin_servers.private_ip`

- [inside_network](resources--origin_pool--reference--group-003.md#canonical-2333033312003333-2020012201210130-2333321300102023-2223301201133333-1131320133122101-3322011331100320-1203013102230121-3220223120201200): complete subsection reference.

<a id="canonical-3131321310122213-2013013222233311-2013100113320312-1111123221322122-3013030033122032-2032133200310211-2223132120100033-2301131030132321"></a>

<a id="canonical-2212033312111111-2020120120121321-1022023233122332-1113113003102210-0311201113130113-1322332231212120-0212023123333331-1133111123233020"></a>

#### `origin_servers.private_ip.ip` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [outside_network](resources--origin_pool--reference--group-003.md#canonical-2332312302210220-3001332333131122-1100220121131321-2233230013312312-2320320220221301-1130102233012220-2120000130122303-2120113222003130): complete subsection reference.

- [segment](resources--origin_pool--reference--group-003.md#canonical-3330110221310112-1110312123133100-2123130300111320-2011312003131331-2323030230133212-2300303013130222-3200300203112331-0000030201112330): complete subsection reference.

- [site_locator](resources--origin_pool--reference--group-003.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-003.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012): complete subsection reference.

<a id="canonical-2333033312003333-2020012201210130-2333321300102023-2223301201133333-1131320133122101-3322011331100320-1203013102230121-3220223120201200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.inside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-003.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.inside_network

<a id="canonical-1010312132210302-2122330213320320-3002121000121223-2321111221333320-3233220102120023-3111220020033121-0022020003220201-2003211121111023"></a>

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

<a id="canonical-2332312302210220-3001332333131122-1100220121131321-2233230013312312-2320320220221301-1130102233012220-2120000130122303-2120113222003130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.outside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-003.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.outside_network

<a id="canonical-3020123121020013-1233223013301111-3001332131013133-2303021113203001-3120321223301031-2112101011011222-0221013312230300-3030013213102321"></a>

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

<a id="canonical-3330110221310112-1110312123133100-2123130300111320-2011312003131331-2323030230133212-2300303013130222-3200300203112331-0000030201112330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.segment` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-003.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.segment

<a id="canonical-0122201332003012-3133322332232022-1210211032033221-1113102110020313-3132031203113300-1100300320102320-1110112012323103-2100131212133310"></a>

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

<a id="canonical-3101010032323132-1332120322210102-2231300020301131-3330003023310313-2120323131312300-2301021222011130-2030012232022330-1130110101231302"></a>

### Direct properties for `origin_servers.private_ip.segment`

<a id="canonical-1223002113120231-2231221222212323-3200032101100123-0100301000333301-1310301122013232-0130201212230011-1121333303122332-1211330012331332"></a>

#### `origin_servers.private_ip.segment.name` property

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

<a id="canonical-0002023023112303-3000002202301301-2122112011220210-2300210311301112-1310303122333301-3023310112222330-3112020232133100-3013231300223301"></a>

<a id="canonical-1121210302211112-2102331332022031-1213001222330111-2220210300030101-2302202110201300-2231023332321331-3000101101022232-0123032010313203"></a>

#### `origin_servers.private_ip.segment.namespace` property

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

<a id="canonical-2211231312012130-3330102313111132-0010213002332311-3333322201120203-3210311203321223-2202121223213200-0210312203332122-3122031211132310"></a>

<a id="canonical-2320011113321101-3300131001001003-0112231333023222-1300023323013010-1100031331100210-0210231002210023-3103023222331030-3230220333120011"></a>

#### `origin_servers.private_ip.segment.tenant` property

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

<a id="canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.site_locator` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-003.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.site_locator

<a id="canonical-2300231310100223-1133011131133232-1022201302332233-0201232012331201-2310030012112201-1312110221032323-1012213021333220-1123101310031303"></a>

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

<a id="canonical-0033120212203030-0023120232002213-2303013201112122-3001121202333312-3201111012312210-2332310320311202-0110132023310310-3312322210202212"></a>

### Direct properties for `origin_servers.private_ip.site_locator`

- [site](resources--origin_pool--reference--group-003.md#canonical-0033003220302032-1233233113310201-1221132212023222-0023132332022003-3233200113230112-0210031031003020-1120331220022312-1302013132002100): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-003.md#canonical-2333130022331103-1112130211201220-0203022112321012-1323001330210201-2223232033131220-1223103200121120-2323132201200131-1232331312223321): complete subsection reference.

<a id="canonical-0033003220302032-1233233113310201-1221132212023222-0023132332022003-3233200113230112-0210031031003020-1120331220022312-1302013132002100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.site_locator.site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-003.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-003.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200)
- origin_servers.private_ip.site_locator.site

<a id="canonical-2331020232213023-3031110031113100-3210011130022031-1310130031320000-1301021211133300-0320112000033032-2013320111122100-1133023031220313"></a>

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

<a id="canonical-0032230231000332-0110101002103330-3232031222020030-3213301313013121-3203311110102220-3130123303111100-1231201211210032-3012020000010113"></a>

### Direct properties for `origin_servers.private_ip.site_locator.site`

<a id="canonical-0213220102003313-2321021101003033-0112323000222111-3312212011301311-1330003232101003-1220312113223211-1331121211323313-3121012013301312"></a>

#### `origin_servers.private_ip.site_locator.site.name` property

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

<a id="canonical-0310332310121302-0012001110320332-2133332133310001-3302331031221212-1122023203031010-3121012220302020-0011011002012001-3331221202100202"></a>

<a id="canonical-0310330121031203-0021000020111322-3321313103101321-2232002000031133-2023031121002111-3102112200012321-0111010130230213-0311230122021210"></a>

#### `origin_servers.private_ip.site_locator.site.namespace` property

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

<a id="canonical-3333130012100110-3133010102330332-2230012130101130-2032111033113120-2101210230323002-3121102032110030-0031220330012002-1022302023022111"></a>

<a id="canonical-3013333010012010-1201221213132121-3111100121312122-0001203202102100-0301312010022102-1333302302000202-0332302310033331-3232112102230310"></a>

#### `origin_servers.private_ip.site_locator.site.tenant` property

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

<a id="canonical-2333130022331103-1112130211201220-0203022112321012-1323001330210201-2223232033131220-1223103200121120-2323132201200131-1232331312223321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-003.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-003.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200)
- origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-2222230231021133-1310332222203132-2213203023213220-0203330233023131-0122101031012301-0203303032321110-2230021233220330-2121021123221222"></a>

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

<a id="canonical-3112323023131331-2102100322102021-0313301033002032-2232333001132202-3121313231100302-0111310323122131-3002302322221300-0331312021323230"></a>

### Direct properties for `origin_servers.private_ip.site_locator.virtual_site`

<a id="canonical-1013210020122120-3113310033130023-2203233010131300-1232120113110230-1012120320130103-2210000321001023-2232103311023030-1331321022012011"></a>

#### `origin_servers.private_ip.site_locator.virtual_site.name` property

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

<a id="canonical-2200220020123203-3033213310312002-0113221220033300-1103023110203201-3032000302331022-2001131023203233-0200213310311000-0133101003022121"></a>

<a id="canonical-0011202022201311-0112210121132020-1320100000311033-1112222013000101-0103203202213112-1100300212013021-0232302310302223-0330022313111301"></a>

#### `origin_servers.private_ip.site_locator.virtual_site.namespace` property

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

<a id="canonical-0323020000231022-2321300331100223-1221131022231231-3322133121100023-3202320023222230-0110102231120112-0202310122123211-0220310023123201"></a>

<a id="canonical-3132110211120020-3323320110100210-1212200221121321-0123011121211200-0002322013033021-2202021123233010-0020201203202211-2303001020031211"></a>

#### `origin_servers.private_ip.site_locator.virtual_site.tenant` property

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

<a id="canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-003.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.snat_pool

<a id="canonical-1230022110130033-0122300323220323-2233303320012120-2123001031323003-0331000322013032-3031212033032211-0210003203200313-3110000013232222"></a>

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

<a id="canonical-3120001111101032-2131212220033110-2321013103311030-0332332220122320-0011310022101011-0333222323023003-2323002003103230-2003330233101130"></a>

### Direct properties for `origin_servers.private_ip.snat_pool`

- [no_snat_pool](resources--origin_pool--reference--group-003.md#canonical-2223123321310311-0022101202202202-2313330102010022-2220011122020023-1030013111210103-1122131100131311-1000123233211222-1311201133113313): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-003.md#canonical-0223213203101123-1313311203312233-1030000202323301-1000222111110203-0212020020331110-0120321020133203-1022120113221100-3331221220332032): complete subsection reference.

<a id="canonical-2223123321310311-0022101202202202-2313330102010022-2220011122020023-1030013111210103-1122131100131311-1000123233211222-1311201133113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-003.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-003.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012)
- origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-1213000033011202-1103310310032130-1003133221100132-2020230011311220-1103111231331110-0213111131321230-2330010102223102-1132132222213302"></a>

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

<a id="canonical-0223213203101123-1313311203312233-1030000202323301-1000222111110203-0212020020331110-0120321020133203-1022120113221100-3331221220332032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-003.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-003.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012)
- origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-0033010020000203-0223123013022223-0200013323230300-3310100102021313-3100202201211321-3000122310222333-3233331313312012-0111022110330203"></a>

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

<a id="canonical-1311331302231020-3003331322301210-3021122011010110-1312211003213010-2111133100223030-2033003202010302-3013121223313322-3013330023031033"></a>

### Direct properties for `origin_servers.private_ip.snat_pool.snat_pool`

<a id="canonical-2310332121303123-3333230133231031-0211120300333113-1202323232220302-0032213011303110-1313113201221230-2313031222100132-2103101033132003"></a>

#### `origin_servers.private_ip.snat_pool.snat_pool.prefixes` property

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

<a id="canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.private_name

<a id="canonical-3022300021030132-2220222012003231-1032303111232231-2212212013120222-2323130212232113-0112320023321113-2313230210023201-3013310312203012"></a>

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

<a id="canonical-2111331002200033-2001132132101202-1331331202320200-3033231230031331-0003330310233110-0302320001210220-1033113332233201-0311323132210013"></a>

### Direct properties for `origin_servers.private_name`

<a id="canonical-0030320020312221-2112213011032231-1010223120203032-0030320313030110-2330203033200010-2000333313312010-1000133300301332-0301233013232223"></a>

#### `origin_servers.private_name.dns_name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [inside_network](resources--origin_pool--reference--group-003.md#canonical-2203123102332202-2332013300310111-3022033222313002-2213133113121032-1122013013203332-2023020212333100-3023233133301122-0311232310323331): complete subsection reference.

- [outside_network](resources--origin_pool--reference--group-003.md#canonical-3120201313033031-2111011302310031-3230023222102013-3032031102130021-1212010010233033-0030201131221331-0112102022203121-1231311201021120): complete subsection reference.

<a id="canonical-1001321013320301-3231130133303213-1012133100211202-0202310003112322-0320013300303023-2210013102000202-3300202320012331-0013231020013230"></a>

<a id="canonical-0032311010112203-2211032022002131-0123203110023121-1122101210231021-0021230110323230-2033021302111032-2110211230123011-0213121132331010"></a>

#### `origin_servers.private_name.refresh_interval` property

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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

- [segment](resources--origin_pool--reference--group-003.md#canonical-1200322213030001-3020120330301111-1323013131020133-1113202333233202-1000222302203230-1223201032001001-3323222331313231-1020213223310111): complete subsection reference.

- [site_locator](resources--origin_pool--reference--group-003.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-003.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203): complete subsection reference.

<a id="canonical-2203123102332202-2332013300310111-3022033222313002-2213133113121032-1122013013203332-2023020212333100-3023233133301122-0311232310323331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.inside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-003.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.inside_network

<a id="canonical-2322312020100021-0201033330302321-0022320311110130-2203302111310130-2112111212210332-0030012002320200-2020220333231022-0100020021213131"></a>

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

<a id="canonical-3120201313033031-2111011302310031-3230023222102013-3032031102130021-1212010010233033-0030201131221331-0112102022203121-1231311201021120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.outside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-003.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.outside_network

<a id="canonical-0013231003130331-3120123023103021-2320131211031233-3223331122101021-2103123112122130-2020233001233000-0013323130033302-2011301100101231"></a>

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

<a id="canonical-1200322213030001-3020120330301111-1323013131020133-1113202333233202-1000222302203230-1223201032001001-3323222331313231-1020213223310111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.segment` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-003.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.segment

<a id="canonical-1133101023013231-2111002010302323-2132111310121133-1101003132011232-1310123201023212-3333020132201122-2000001333102003-0200320101132013"></a>

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

<a id="canonical-2201010011012121-1012130020321211-1230100301130220-3313032311113221-3013112033111220-3220030010021122-1013330020102123-3212202121112010"></a>

### Direct properties for `origin_servers.private_name.segment`

<a id="canonical-3130300300300122-2021121312321213-2203003013113330-1312322123211312-2132330301132333-0003220131030032-1010311023201132-0303113001021320"></a>

#### `origin_servers.private_name.segment.name` property

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

<a id="canonical-2213202322130032-1302102012000111-0130011113033110-0332003132300102-3013320111020321-0123202013202103-2200010222222002-1313303132130320"></a>

<a id="canonical-3011031011021330-2012122131230112-2132333303032132-3320032030331330-0002011321111231-3312231023101113-0020202231221222-0110223213211112"></a>

#### `origin_servers.private_name.segment.namespace` property

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

<a id="canonical-0020132303013213-0000202322312013-0130223103003201-2320022211103132-1011010001103111-3210213313122021-0323000330112332-0230310211022122"></a>

<a id="canonical-0301033310011303-3313222012210101-2001331301122303-1131202330020220-1103001331113203-1120211330211321-0000321333011330-1233231202220031"></a>

#### `origin_servers.private_name.segment.tenant` property

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

<a id="canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.site_locator` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-003.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.site_locator

<a id="canonical-1020001220021113-3201030222113010-3233123302013110-2033130012033221-0330122021331100-1113302312232110-2312031322302123-3113203103030302"></a>

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

<a id="canonical-1322231131300002-1110001131112130-0203301221221232-3010103202112330-0131033102200100-2332101322333231-3321102330312301-0330230021303001"></a>

### Direct properties for `origin_servers.private_name.site_locator`

- [site](resources--origin_pool--reference--group-003.md#canonical-3202010311011020-1202031221320113-3122030101221230-2200102322211100-2123330202232233-3001123100320133-3112032000130220-0233002310102130): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-003.md#canonical-0310120101132210-0322321121303212-2123011032321233-0312231230321133-2231302322012302-2001330223222303-0223221321130131-3023110131010102): complete subsection reference.

<a id="canonical-3202010311011020-1202031221320113-3122030101221230-2200102322211100-2123330202232233-3001123100320133-3112032000130220-0233002310102130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.site_locator.site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-003.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-003.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222)
- origin_servers.private_name.site_locator.site

<a id="canonical-2310110312030231-0323110012101103-1231313132110123-3322003102003131-1320220002300020-0210320122211021-3233202103231300-0312122000232233"></a>

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

<a id="canonical-1003330323212123-2312011010210021-2200100231122312-2130002123113110-1013110200213022-1303010121210103-3101003203320000-3013212012103132"></a>

### Direct properties for `origin_servers.private_name.site_locator.site`

<a id="canonical-1233332330010203-2103101120011230-3302022311211002-3131021321120101-1030011110033031-3201111031230123-3322103220002102-0023103113220323"></a>

#### `origin_servers.private_name.site_locator.site.name` property

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

<a id="canonical-2030110230122011-0121312231332013-0103303202223333-0202112032330303-0111300300121002-0020021300130230-3302131213012013-1302221311122312"></a>

<a id="canonical-1122331112312113-2332233321220222-0031330211120013-1112103201313213-1202333120030333-3332011232200231-3031121031121230-3131121232311210"></a>

#### `origin_servers.private_name.site_locator.site.namespace` property

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

<a id="canonical-3122311300323320-0211222202200211-3102330002320302-3333310303332213-1212202113332222-3222331232300233-2232323000323223-1033321013103023"></a>

<a id="canonical-1131323032213200-2211331220220101-2030321331132230-1323023210112331-1121322111120230-2323000302323102-3320011121031111-2213220303110013"></a>

#### `origin_servers.private_name.site_locator.site.tenant` property

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

<a id="canonical-0310120101132210-0322321121303212-2123011032321233-0312231230321133-2231302322012302-2001330223222303-0223221321130131-3023110131010102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-003.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-003.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222)
- origin_servers.private_name.site_locator.virtual_site

<a id="canonical-0310202213011000-3111221033301021-0000311230003001-3133110320333332-2332303200200020-1133002033223110-1113111032333320-0303223221312122"></a>

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

<a id="canonical-3112031122031333-0302033201200002-3210233133231200-3012011222303332-0331210010122033-1002321233200030-0310331123201233-0120321303300213"></a>

### Direct properties for `origin_servers.private_name.site_locator.virtual_site`

<a id="canonical-3230020311221133-3320203223320221-3003211320333331-0100110201013203-0013132002113213-1033002220010200-0032321303320130-3003202000212202"></a>

#### `origin_servers.private_name.site_locator.virtual_site.name` property

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

<a id="canonical-1211202113232100-2103013212012101-3223031333022230-2020222132122330-1013333210103312-2133112122232022-0231130221203310-1301032203320222"></a>

<a id="canonical-3013012201102123-1033102123030303-1020323132133130-1100013131103302-3131130232010301-3000123221010203-3023233211113221-0122210201130012"></a>

#### `origin_servers.private_name.site_locator.virtual_site.namespace` property

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

<a id="canonical-3333120331021200-0122122122112221-1033232311130311-2001001031121012-1323101121200221-3031223233321322-1020232202213333-0032031311223013"></a>

<a id="canonical-0310111030121122-1023013000311201-1221021332003310-0132102323220013-1312022323323322-3201311033012302-3102231330021130-1002321213220010"></a>

#### `origin_servers.private_name.site_locator.virtual_site.tenant` property

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

<a id="canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-003.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.snat_pool

<a id="canonical-1012322211101010-0203023132110100-0321333121033313-1121311033211030-0012132301213012-1133130003230120-0020130231013323-3113120111010133"></a>

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

<a id="canonical-0102113211320023-2000222013122112-2332023220301331-3023103222231232-3131211313122102-0200033213301101-1203021113122130-2323333331223002"></a>

### Direct properties for `origin_servers.private_name.snat_pool`

- [no_snat_pool](resources--origin_pool--reference--group-003.md#canonical-3130023221023210-3020131101000132-2331203133232103-1230023310301102-2223220110200130-0220200112120302-3310132133311233-3200112023312112): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-003.md#canonical-2221131331021133-1131233031222100-2123110022013001-1003330121330001-2001002310023110-1111101220230032-3332220113200220-3133303312002021): complete subsection reference.

<a id="canonical-3130023221023210-3020131101000132-2331203133232103-1230023310301102-2223220110200130-0220200112120302-3310132133311233-3200112023312112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-003.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-003.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203)
- origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-0113203312310213-2110203230230023-2212211321322013-2102102302112231-3013121131100003-2120303003130323-0233330323032322-0223131100133121"></a>

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

<a id="canonical-2221131331021133-1131233031222100-2123110022013001-1003330121330001-2001002310023110-1111101220230032-3332220113200220-3133303312002021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-003.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-003.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203)
- origin_servers.private_name.snat_pool.snat_pool

<a id="canonical-1012022110102030-0322310000331011-1121232011213022-0223112103323312-2103220000323222-3113313200023112-3120001311231113-0031212113330001"></a>

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

<a id="canonical-0013002003210123-0322131311021100-1302011131302222-2033301121130200-1001233010331301-3230111132010130-1131202232033131-0323121322302222"></a>

### Direct properties for `origin_servers.private_name.snat_pool.snat_pool`

<a id="canonical-0022211003233101-1123312210101003-0022330320112012-0103323033222030-1222232031131222-1020220031211132-0033330032101220-2332200002003220"></a>

#### `origin_servers.private_name.snat_pool.snat_pool.prefixes` property

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

<a id="canonical-2012202102220023-2022233311031000-2131130202320231-1200001300020311-2032322312323113-2110322301123032-3132132030003331-0200031213133012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.public_ip` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.public_ip

<a id="canonical-2330102103323312-2311122332333212-2111322303321021-1202231123132011-2230321120110033-1323002102021311-3322132202133231-3130033013112100"></a>

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

<a id="canonical-0100100002113111-2332111002331120-1003002202021103-2112321021213200-0132202020010323-0021200210201032-1301113330230321-0032223233022302"></a>

### Direct properties for `origin_servers.public_ip`

<a id="canonical-0301113230322103-0122131333211003-2200221332132002-0002323200320120-0322023121113332-2030330012133033-0131102333011030-1221302101232132"></a>

#### `origin_servers.public_ip.ip` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0101123032132221-2211032101001201-1000311221111203-2133111213011322-0333020102301313-3100023301322230-0212222321323322-2203130310321323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.public_name` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.public_name

<a id="canonical-0222302133210320-2001221203003010-3022300302110110-3000333100030103-2313301012110103-2212210120131013-1012130222122313-3202312023112330"></a>

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

<a id="canonical-1131321220010200-1302000320021330-1102220110202212-1030222100120002-2023031300111111-1211322301122203-1122212331312221-0111133023101132"></a>

### Direct properties for `origin_servers.public_name`

<a id="canonical-2111122303201022-2233210133031300-0320303100012233-1103103233312021-2032231230310321-1120300132031101-2002033321232112-1233310112133232"></a>

#### `origin_servers.public_name.dns_name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3021102323312312-2101333023321220-1102102211203121-3133300301203323-3213311103001301-2121213022101223-1023111301321210-1011010022020222"></a>

<a id="canonical-0303200232132000-2120313211022310-0203102030330113-2131303031213111-1111221120000203-1022202100212201-3231313020200012-1202232113321200"></a>

#### `origin_servers.public_name.refresh_interval` property

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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-2130122232222011-2300321030110011-1031031023222200-1001031200312013-2120002123011312-1310133101323000-3001103213310123-2023322321313313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.vn_private_ip` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.vn_private_ip

<a id="canonical-1212301103220103-0222133311031102-2310030020003310-3011032130010130-0333113121203022-1121200301130110-1121333133231213-2122303322012222"></a>

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

<a id="canonical-3221012132202023-1311312101133310-3110223313302012-0121313002310333-1101331020220201-1231201332110332-2001020022021033-3221223100313211"></a>

### Direct properties for `origin_servers.vn_private_ip`

<a id="canonical-0303031013321122-3312002010330312-3300200032023111-1022212011213113-2232202312122132-0131302301220001-3101202011123200-0312301013023023"></a>

#### `origin_servers.vn_private_ip.ip` property

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](resources--origin_pool--reference--group-003.md#canonical-3122100030131322-0113010211030321-3203333130200133-1111310203002032-0020302020200012-0010322030112332-3310313213332211-3231002332012330): complete subsection reference.

<a id="canonical-3122100030131322-0113010211030321-3203333130200133-1111310203002032-0020302020200012-0010322030112332-3310313213332211-3231002332012330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.vn_private_ip.virtual_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.vn_private_ip](resources--origin_pool--reference--group-003.md#canonical-2130122232222011-2300321030110011-1031031023222200-1001031200312013-2120002123011312-1310133101323000-3001103213310123-2023322321313313)
- origin_servers.vn_private_ip.virtual_network

<a id="canonical-0003212300322202-3222000200122020-3203010013001101-1311323200003200-2102323331103133-0230313302203300-3302211300312303-3202332311301130"></a>

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

<a id="canonical-3022211213221120-3202311013310330-3122310013202011-0132002112233211-3132302330032021-2010301123033221-3121033322232201-2023012331211132"></a>

### Direct properties for `origin_servers.vn_private_ip.virtual_network`

<a id="canonical-0000202302112102-1112012101000233-3302013312100000-0023221022201010-0201113330200201-2003201103322210-1330030120230101-2002210032322331"></a>

#### `origin_servers.vn_private_ip.virtual_network.name` property

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

<a id="canonical-0223223101122021-2213133211011111-1110303311033212-2202212310201012-1232122321032311-0022200133002320-1203111330100322-3332031233320120"></a>

<a id="canonical-2211231211330221-2032031113021110-1213213100021003-1033233311120332-1320210323301002-2232301111113131-1332221203312030-1100003133312110"></a>

#### `origin_servers.vn_private_ip.virtual_network.namespace` property

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

<a id="canonical-0111322301320300-0322201111323203-1212030301001200-0211310021223012-3123201121032322-2030010332303230-2231323310021331-0120323133322133"></a>

<a id="canonical-0310322312201233-2000332330303222-2312030031212312-1011333103200111-0132332212221311-0211213000102210-0221100201203231-1022220113220320"></a>

#### `origin_servers.vn_private_ip.virtual_network.tenant` property

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

<a id="canonical-1013210030120200-0300013133323110-0211222301220003-2111231123223232-2130003221211000-0010213100230320-3322112211323121-2102120111321112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.vn_private_name` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.vn_private_name

<a id="canonical-1133331232233003-1320111130012311-1123222301130121-1020122210302320-1322110011220312-3110310013210300-1130202300310203-1332113133000131"></a>

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

<a id="canonical-3003101232323020-3001020300300133-1023230033123020-3212311012112300-2020230032322001-0120100110103023-2320320020331110-1213301032320100"></a>

### Direct properties for `origin_servers.vn_private_name`

<a id="canonical-3002030022032132-3311301011121132-2213132122130322-0202002203100100-0201311301120102-1020000023121010-2020131020121000-2302221200332000"></a>

#### `origin_servers.vn_private_name.dns_name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [private_network](resources--origin_pool--reference--group-003.md#canonical-3313202330121310-0332121101021311-0020212111023330-2010022013100332-0131010011102330-1300031011320320-3223132122001110-3220213212120000): complete subsection reference.

<a id="canonical-3313202330121310-0332121101021311-0020212111023330-2010022013100332-0131010011102330-1300031011320320-3223132122001110-3220213212120000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.vn_private_name.private_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-002.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.vn_private_name](resources--origin_pool--reference--group-003.md#canonical-1013210030120200-0300013133323110-0211222301220003-2111231123223232-2130003221211000-0010213100230320-3322112211323121-2102120111321112)
- origin_servers.vn_private_name.private_network

<a id="canonical-1133303321023212-2001300031121022-1011301211231302-0001021121033021-2323331011202103-0022221131021002-1223211000123212-0002221220103230"></a>

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

<a id="canonical-3210130011123132-1221333223113212-3203103010121211-0030031233023212-0322030303332112-2021002113021323-2101221221003002-3030013211200300"></a>

### Direct properties for `origin_servers.vn_private_name.private_network`

<a id="canonical-0211012010133233-1230101313033231-2203330313202001-0230213001201123-0023222333323120-1021301221110101-3130101312333010-0022323203212030"></a>

#### `origin_servers.vn_private_name.private_network.name` property

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

<a id="canonical-3001031121122201-2021001100001030-1223130321201021-2103201022330133-0031030202301102-2003332030311030-1013301002031230-0333123312320300"></a>

<a id="canonical-0012301331302022-1303130000001202-3200103020002222-0320223322203303-0320201320021013-2222221200112120-2200011131132120-2020211101323323"></a>

#### `origin_servers.vn_private_name.private_network.namespace` property

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

<a id="canonical-3203203021110231-0111012212301333-0231202321023031-0200023321203132-2122323200121130-1202232210232212-2112323113223002-3222230212232320"></a>

<a id="canonical-0303321311102333-1230110211213001-3123232212213210-2333121021122233-1223223203203332-0330001220230130-3311303133033201-0013012223113021"></a>

#### `origin_servers.vn_private_name.private_network.tenant` property

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

<a id="canonical-3110220132232202-3223311123032232-3333333030113010-2133200310011011-3010332223210221-3113000210333321-1221001103212033-0210001201112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `same_as_endpoint_port` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- same_as_endpoint_port

<a id="canonical-2003120102323102-0320322222233123-0110310222321212-1303023300232011-2231113032010222-0322332011312201-0332033220131211-2002332200121120"></a>

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

<a id="canonical-1322321330320022-2021311033121221-0210313231232031-3331330210233122-2333300130103032-3230223313322102-2023330330022001-1110122123303032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- timeouts

<a id="canonical-0000231121303330-2003202121011221-0020300020231001-1223302100322112-1221333321301112-3331212200332003-2111332002333011-3211101330111331"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222033010120310-0112013213211130-1311130211110030-1001312102332212-1210033102013201-1102332001012213-1211213131200033-0211013102003032"></a>

### Direct properties for `timeouts`

<a id="canonical-0322021132111321-3112303003031330-0130110311120000-2323112211313121-2012101131003231-2310330113310310-0320302320102011-2221132032313132"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1123132211211221-0212322301210311-3001010203232213-1212101120321000-2222310231130132-1231023213030003-0031133032232203-3212022030111323"></a>

<a id="canonical-1033323101210003-0330031130031231-1220133032102130-3330202302223232-3011223311212223-3323012303020112-2301212100130213-3123313001310010"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0112102231203010-1002213233010322-3013322001331311-1200201112200033-3200131003022301-2112112000322322-2232133101312103-0233312112020123"></a>

<a id="canonical-0200310130121112-1303133311330033-3120301030230323-3123012032013021-0202232101001212-3021321100201000-3200101222010201-2003022222232221"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0110313103330301-1331101223223003-0022300330300231-0323203311011100-3322001332200002-3310012203101213-2201112330320131-0311333030233120"></a>

<a id="canonical-0102012122130023-2202211032033233-3011000210113033-3111223021201001-3121123211130323-3132211320121101-1231121030320331-1322313120233131"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- upstream_conn_pool_reuse_type

<a id="canonical-2020113323212003-1003032002212201-3230111332303132-1132100131321213-1302302110200330-2111203102330131-3033100023112000-2101313201120202"></a>

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

<a id="canonical-3230212333110220-1133211332032130-3210330102110023-2003103031220011-3022221302312033-1112032310001333-3332212313032332-1003322103001123"></a>

### Direct properties for `upstream_conn_pool_reuse_type`

- [disable_conn_pool_reuse](resources--origin_pool--reference--group-003.md#canonical-0333220323011211-3002123102021320-3000100303102013-1201002102221100-3121001022112320-0330121330301200-1020213332211223-0312133001222323): complete subsection reference.

- [enable_conn_pool_reuse](resources--origin_pool--reference--group-003.md#canonical-3131222111103001-1022113210301001-2003220322321311-2101301230133010-3001313222333012-1033323302303222-3113132123323330-0121133313120301): complete subsection reference.

<a id="canonical-0333220323011211-3002123102021320-3000100303102013-1201002102221100-3121001022112320-0330121330301200-1020213332211223-0312133001222323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-003.md#canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-1302332230212100-0332010233013101-3300102301101123-0002201113311001-1130033132122002-0201100203010021-1301330033303021-3000221202311030"></a>

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

<a id="canonical-3131222111103001-1022113210301001-2003220322321311-2101301230133010-3001313222333012-1033323302303222-3113132123323330-0121133313120301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [upstream_conn_pool_reuse_type](resources--origin_pool--reference--group-003.md#canonical-2110330210320132-2322030312200020-2112323000003130-3200322302113022-0330212202312000-2122312023332311-1230210101310023-3020210011330320)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-2101101012002111-0212122133332312-2211023230011231-2323313110333203-3111100301202012-3302031232001111-1000203100112302-0033120332203101"></a>

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

<a id="canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- use_tls

<a id="canonical-3122122211032132-3132313203100001-1013000203202031-2132002223311123-2211000023310210-0233300322102210-0033100013333133-1022321012001012"></a>

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

<a id="canonical-0033111310100321-1010320213112001-2222323331120000-1211223131302000-3330300130013021-3120200201102310-0013031101223222-1203111013033300"></a>

### Direct properties for `use_tls`

- [default_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-3123111322223130-2113231002312322-2021111203021000-2221013312223310-1330300113321002-0023103113202103-2200320023203012-2100211113110032): complete subsection reference.

- [disable_session_key_caching](resources--origin_pool--reference--group-003.md#canonical-0231110120222102-3131002013331211-2033120103001302-0200303313101013-0110101211131112-2230120011012022-0333010030130311-3003202112111112): complete subsection reference.

- [disable_sni](resources--origin_pool--reference--group-003.md#canonical-2311231331030322-0130021030000103-0232213212020223-2002202301113003-2232100021031011-2301222221032322-1112230231312120-2003313001133103): complete subsection reference.

<a id="canonical-1010332301301203-1302100023030012-0012312120112131-0223010111101231-0301233110003032-1101022330100330-2331001011130013-2200321301310220"></a>

<a id="canonical-0103333122330111-0011013233212321-1121102300232223-2223312321230013-1100230310222320-0300130333300200-0100122203121232-0323211000211331"></a>

#### `use_tls.max_session_keys` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [no_mtls](resources--origin_pool--reference--group-003.md#canonical-1020313213111330-2202000101202102-0120200210030112-2323310203330012-1213220013320322-3032000023223012-0120133301120131-2122220311110221): complete subsection reference.

- [skip_server_verification](resources--origin_pool--reference--group-003.md#canonical-0102332020023212-1113031013000102-1232320111330022-1303001202020211-2121331221103003-2023022213211102-2323031210232300-1021020203331201): complete subsection reference.

<a id="canonical-1220200011310322-1313110000203111-0131320231230110-1312300003202311-2021312133301200-1203030202220211-2233131120002033-1310221211332200"></a>

<a id="canonical-3222033320102121-1120201033110323-1222123210103102-3013032023300211-2022323103210310-0301301132031001-1003022220110021-0233133302000103"></a>

#### `use_tls.sni` property

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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131): complete subsection reference.

- [use_host_header_as_sni](resources--origin_pool--reference--group-003.md#canonical-3320212303200311-1302210200012223-1223330320011323-0032131002221211-3123001022122222-2330223333111000-2211103232202111-1032322131022022): complete subsection reference.

- [use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011): complete subsection reference.

- [use_mtls_obj](resources--origin_pool--reference--group-003.md#canonical-1100000211323212-3202110010310330-3033130000112230-3222111010321321-1333210130131112-1111210130020333-2222310120300310-1221330010012110): complete subsection reference.

- [use_server_verification](resources--origin_pool--reference--group-003.md#canonical-3000330133311333-2312222030122331-0220231200123233-1320011200021220-0002310301112001-0231020200121203-2330012211110321-3000301310100002): complete subsection reference.

- [volterra_trusted_ca](resources--origin_pool--reference--group-003.md#canonical-0303020022031131-2122331331231202-2331332210010010-1201301230030123-0022302320133221-2300002131320312-3100230312312211-2111020303131110): complete subsection reference.

<a id="canonical-3123111322223130-2113231002312322-2021111203021000-2221013312223310-1330300113321002-0023103113202103-2200320023203012-2100211113110032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.default_session_key_caching` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.default_session_key_caching

<a id="canonical-3030232203023330-2211203000103303-0030123222002110-1210003113321012-3223303023211312-0300132233212333-0312103033022102-0100002311022323"></a>

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

<a id="canonical-0231110120222102-3131002013331211-2033120103001302-0200303313101013-0110101211131112-2230120011012022-0333010030130311-3003202112111112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.disable_session_key_caching` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.disable_session_key_caching

<a id="canonical-0133330110333312-2101233020121221-1212201013003111-0213300012112330-3332101120100012-2122111011331212-2102213012002122-3212032202130203"></a>

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

<a id="canonical-2311231331030322-0130021030000103-0232213212020223-2002202301113003-2232100021031011-2301222221032322-1112230231312120-2003313001133103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.disable_sni` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.disable_sni

<a id="canonical-3121221222210030-0011023312311000-0213231013120333-3321330130232201-1122111133021030-0322220021233033-1301012022233032-0331111232122201"></a>

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

<a id="canonical-1020313213111330-2202000101202102-0120200210030112-2323310203330012-1213220013320322-3032000023223012-0120133301120131-2122220311110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.no_mtls` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.no_mtls

<a id="canonical-3011203232233111-0210021203020023-3032203220101031-1110110203331323-3133231000132032-0101010211121311-1223020302132201-1013232312121222"></a>

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

<a id="canonical-0102332020023212-1113031013000102-1232320111330022-1303001202020211-2121331221103003-2023022213211102-2323031210232300-1021020203331201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.skip_server_verification` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.skip_server_verification

<a id="canonical-2222233003100111-0302003010100213-2011310100011121-3002311212312332-2311313311313130-2211231120332133-1223313220232221-0113201012312301"></a>

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

<a id="canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.tls_config

<a id="canonical-1020312000221113-0332112223333302-3122020211012330-2001011231100112-2100330310322121-0023311223022300-1202133123232313-2202220123133300"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0210312202132303-3201213002011322-0221132123010203-1300003333102100-0312002311021132-1223330120310303-2112230201210000-1313332011102131"></a>

### Direct properties for `use_tls.tls_config`

- [custom_security](resources--origin_pool--reference--group-003.md#canonical-3321220300323330-2011010022013312-3223331223303132-1321323020132210-3213103013213101-1222330223012211-3320101022222330-3122312021321231): complete subsection reference.

- [default_security](resources--origin_pool--reference--group-003.md#canonical-0010312130132201-0000033312203031-1330213232320210-3223020302013121-1122301103331002-1313101011110111-1312203122120132-1000132012313312): complete subsection reference.

- [low_security](resources--origin_pool--reference--group-003.md#canonical-2112232222000230-2202322132200020-1311102031112203-1313302300230002-2212023020201330-1020012331320010-1022320203311321-1211131311200110): complete subsection reference.

- [medium_security](resources--origin_pool--reference--group-003.md#canonical-3313302220222212-1100021000102303-2330332310012332-1332212023201003-3011302302000011-0332220203323323-0312122331230232-0210023210120221): complete subsection reference.

<a id="canonical-3321220300323330-2011010022013312-3223331223303132-1321323020132210-3213103013213101-1222330223012211-3320101022222330-3122312021321231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.custom_security

<a id="canonical-0203303001311310-2322002212132332-1010120123231232-1113000203111122-2201232211310022-1032120102202233-0332023301120102-1323302313333132"></a>

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

<a id="canonical-1223322230311021-0323220302223212-2010213221232002-3320313211230223-1211030230323310-3011020031333211-0310013322203110-2010010010300332"></a>

### Direct properties for `use_tls.tls_config.custom_security`

<a id="canonical-3203013203033320-2103302313120031-2110301221221120-1030331013001031-0321112320310212-1020103003323321-2002032100001303-0312103212130330"></a>

#### `use_tls.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0122213110200311-0101230001332320-0130020201122132-3113320102110003-3223310321312132-3232133210333113-2323320131111230-3101002032132012"></a>

<a id="canonical-3112121020131121-1103112213320331-2012030100300110-2001312110111221-1323010301231011-1200211231023222-0010011002310033-3333330301213221"></a>

#### `use_tls.tls_config.custom_security.max_version` property

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

<a id="canonical-1233311302332033-3103001123311321-3313220110300313-3012312233103323-1211312223130202-0123001303100212-2233321301012023-3333203132331202"></a>

<a id="canonical-0210110200022301-2231302301030220-3121302330312103-3321102202131323-0330123322233013-0312030233301310-0210220321023121-3030303130021012"></a>

#### `use_tls.tls_config.custom_security.min_version` property

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

<a id="canonical-0010312130132201-0000033312203031-1330213232320210-3223020302013121-1122301103331002-1313101011110111-1312203122120132-1000132012313312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.default_security

<a id="canonical-2200230203011232-1012310231101221-1302222003121322-0000213033303310-1321100332211312-3202202332303220-3101323313203001-1331301111100332"></a>

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

<a id="canonical-2112232222000230-2202322132200020-1311102031112203-1313302300230002-2212023020201330-1020012331320010-1022320203311321-1211131311200110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.low_security

<a id="canonical-2300211011100131-2312131300032321-3133020322230003-3030323201100300-0020003002220020-1031212102213310-2202030002123013-0032322311122213"></a>

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

<a id="canonical-3313302220222212-1100021000102303-2330332310012332-1332212023201003-3011302302000011-0332220203323323-0312122331230232-0210023210120221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.tls_config](resources--origin_pool--reference--group-003.md#canonical-1112222132222202-3113101010112012-2301033031100332-1302203012210011-1202110221111220-2012033212332201-1202112331303003-3201200232221131)
- use_tls.tls_config.medium_security

<a id="canonical-3220202323330311-0003102032032322-0223033101210321-2013332030011010-2213021230323322-0302133022011001-2330131022102303-1020303313021110"></a>

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

<a id="canonical-3320212303200311-1302210200012223-1223330320011323-0032131002221211-3123001022122222-2330223333111000-2211103232202111-1032322131022022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_host_header_as_sni` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_host_header_as_sni

<a id="canonical-3010233110232013-0320031332313011-3002333211230103-3232321131201122-1230310031110203-0133221320122110-3131120320021100-2303232111101310"></a>

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

<a id="canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_mtls

<a id="canonical-1011223301223132-2013212102312122-0221212033123111-3331012332113120-1212201210230330-1203130033210303-0303021021321123-0333022103120211"></a>

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

<a id="canonical-0030303000323103-0230230023020220-0001103110131003-2302200033012211-2212231002301303-3332222302113211-0113312032310333-1231231312322230"></a>

### Direct properties for `use_tls.use_mtls`

- [tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320): complete subsection reference.

<a id="canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- use_tls.use_mtls.tls_certificates

<a id="canonical-3322232030000201-2023133321112003-1230111113231113-0103010032120300-0011131231103233-2012100003111330-1021213301220103-2120010120303223"></a>

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

<a id="canonical-1322013210011121-1132221301321101-3012130223222221-0113333300031111-3321102112131331-3003300020320223-2013013223212231-0313012100200300"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates`

<a id="canonical-1020333333103300-3203001211023011-0331100312033032-1331323202120103-1003032303313233-2322333023311310-0003202103221112-1101012322013010"></a>

#### `use_tls.use_mtls.tls_certificates.certificate_url` property

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

- [custom_hash_algorithms](resources--origin_pool--reference--group-003.md#canonical-2202201331300211-2223113011101021-2230232031101313-3320210222212322-2320120330232023-3311301020001300-0001123200002210-3133121001210210): complete subsection reference.

<a id="canonical-2313010232001222-1321021230103222-1000002003103000-3111022130302323-2211203311020331-1320202010132122-1013002033320032-3122310031333311"></a>

<a id="canonical-1030313333323100-1303221130011233-3312333020120031-2223233312133101-0011002231012113-0300321313012332-1131103023310132-2113001133031302"></a>

#### `use_tls.use_mtls.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--origin_pool--reference--group-003.md#canonical-2001300233233210-3100202223023232-0211131310311312-1300020331133133-2010310112330200-2222202113022230-2213023222332303-3320223201331222): complete subsection reference.

- [private_key](resources--origin_pool--reference--group-003.md#canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312): complete subsection reference.

- [use_system_defaults](resources--origin_pool--reference--group-003.md#canonical-1132223131001013-2000112132212113-2112311231033010-0202003212322031-3101112131010220-0031313121102311-2023003320113022-1211330122030113): complete subsection reference.

<a id="canonical-2202201331300211-2223113011101021-2230232031101313-3320210222212322-2320120330232023-3311301020001300-0001123200002210-3133121001210210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-2213223211110223-3021301202303321-3000101311331301-0110113113100000-0031030230003220-2102010111213120-0332200330001012-2310011311303112"></a>

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

<a id="canonical-2330221122302013-2110222132001213-2010031001221220-1100100202312111-1011212332031120-0013111321133300-1000332231002313-2213303202132010"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.custom_hash_algorithms`

<a id="canonical-0321110130030203-3123231300301322-0100330000011001-3203112331003102-3021331301133322-0300331222030123-2231130120300120-3121221330333033"></a>

#### `use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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

<a id="canonical-2001300233233210-3100202223023232-0211131310311312-1300020331133133-2010310112330200-2222202113022230-2213023222332303-3320223201331222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-3333300201211233-1032301121302000-2022223332122122-2223111331102203-1323103332103020-3302230030122232-1212130120321021-1111203300330332"></a>

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

<a id="canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-2312122300111210-2320203233333123-0110010220000333-2133221212322120-3332123321213233-2112300310000012-2032231020312032-0001321213022110"></a>

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

<a id="canonical-1300101123330201-0030121031103211-3332333212100113-0020330033031222-0102310102302330-1321330222232120-0002102301022031-1320233013223302"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.private_key`

- [blindfold_secret_info](resources--origin_pool--reference--group-003.md#canonical-2232321113220211-3021203013022313-0103003330222203-1320101212320233-2102003030032200-2112310121230223-0321133201110130-3112221230223000): complete subsection reference.

- [clear_secret_info](resources--origin_pool--reference--group-003.md#canonical-1212011202233101-3332323031211311-1202332113313113-2110231000330301-3003121233113330-1010203302100333-0003231332212333-2130000310102000): complete subsection reference.

<a id="canonical-2232321113220211-3021203013022313-0103003330222203-1320101212320233-2102003030032200-2112310121230223-0321133201110130-3112221230223000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312)
- use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1321030203013330-2010233100103102-0122331220200012-1133022023030330-0113333100231223-3131102101133033-2303111220320013-3013030233123210"></a>

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

<a id="canonical-1203211223020122-1111122122021333-1132001031023320-2222130031212230-2110322311132102-3013130233102120-3223220012210100-3221102320312013"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-2013331301120103-2210112033020212-0312000203011300-3120003131103213-3203013231030210-3330031202200330-3233002212013021-0301223113122001"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3102221000300000-1010023013101110-0132000101123120-1131113230111302-2213302333033232-2030201122210021-2213233022022031-3030001212001030"></a>

<a id="canonical-0021033101313312-0121320300300032-2123221032023323-1120202112033012-0301210200103122-3102003200312133-2123022233211323-1103232203333112"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0033311012223112-1130122113201000-2123313110033211-0330223023223333-2212033331030011-3203322312323213-2232033003201121-3013302131310202"></a>

<a id="canonical-2303230222000201-3213001300103022-2221120031101221-3012222203232132-2202222313002201-3101302201013233-1123202022230333-2210022001220231"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-1212011202233101-3332323031211311-1202332113313113-2110231000330301-3003121233113330-1010203302100333-0003231332212333-2130000310102000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--reference--group-003.md#canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312)
- use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-2111132313310011-2312201030120001-1123021332221032-2101010113212013-3212022320002303-2313331301003330-2331210133313200-3023131301312032"></a>

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

<a id="canonical-1033112132010113-1110233011330131-2302021002011121-1122030112123103-0032031221112000-0002012100223310-3202221001320311-3311213103301011"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0033032221002213-2122323020200123-2232322102333232-0323201332200320-3122021310323232-0230012231030103-3132210321021211-2012200313310033"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1211332112310200-1232002021223003-3120011200102233-1332111032220302-2322202210213020-0200202211320031-2010220010302203-0221300031211030"></a>

<a id="canonical-1212301102231203-1310103301311220-2110330031120100-1233300201222213-3133103130003301-1212310323312133-2330033230001020-2103312030020302"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1132223131001013-2000112132212113-2112311231033010-0202003212322031-3101112131010220-0031313121102311-2023003320113022-1211330122030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_mtls](resources--origin_pool--reference--group-003.md#canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011)
- [use_tls.use_mtls.tls_certificates](resources--origin_pool--reference--group-003.md#canonical-3110323203232332-1133011302100203-3022100131002020-2323322113213222-2100030010302202-0001013131230231-0323112011203133-0313303000312320)
- use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-0301312001311001-3013220001311121-1010021310302213-1333300031230210-0110113323013013-3201201200133221-3332230303220222-3303022020031233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100000211323212-3202110010310330-3033130000112230-3222111010321321-1333210130131112-1111210130020333-2222310120300310-1221330010012110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls_obj` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_mtls_obj

<a id="canonical-0311020032113212-1203103100333300-3120011330330310-2133030131113310-1200133213303001-0103330110212320-0311211222012323-1000301320200020"></a>

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
use_mtls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301322212311213-3230032230200211-3230012231133220-0331102331011031-0122201231032000-0110312310013223-1323213113000022-0212031323002112"></a>

### Direct properties for `use_tls.use_mtls_obj`

<a id="canonical-0231211130302032-2202100003223131-3133021022010120-1031222032311213-3310111011332112-3103221020013201-2111031200302330-0323030313333001"></a>

#### `use_tls.use_mtls_obj.name` property

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

<a id="canonical-3210200223021301-1323211003102302-3230323233302133-3221102000111102-3223120130322211-2001232023212122-0233312112011310-1233230002321120"></a>

<a id="canonical-2301311312303232-3213212013330200-0320100033230113-3023201320322112-1212112220100310-0211013032230302-2032321130030302-3300201311221222"></a>

#### `use_tls.use_mtls_obj.namespace` property

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

<a id="canonical-3200110220221112-3132202102011101-3023302001331222-0320032203131012-2132103012300313-0101322313011012-1233011231000111-1333011302033102"></a>

<a id="canonical-2013313221011112-2230020132313131-0031211202202220-0302123022302021-3010220203022332-1103101220332000-1032202102312211-0112100133333130"></a>

#### `use_tls.use_mtls_obj.tenant` property

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

<a id="canonical-3000330133311333-2312222030122331-0220231200123233-1320011200021220-0002310301112001-0231020200121203-2330012211110321-3000301310100002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_server_verification` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.use_server_verification

<a id="canonical-1202022113131322-3223333122031311-2012231003203110-2200120033103313-0303123113012001-1021131200222210-3302220021012302-2121023331223031"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Additional upstream details:

Upstream TLS Validation Context.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230130300220332-2220112133333010-2330133300102111-2303113300000313-2302211333123123-0223002301021111-1211003112223033-3120311222001311"></a>

### Direct properties for `use_tls.use_server_verification`

- [trusted_ca](resources--origin_pool--reference--group-003.md#canonical-0120132103213330-1221300320110330-2321301033111133-2013000231030120-3331212310220231-0131231033221131-2002231222231110-2310012232111211): complete subsection reference.

<a id="canonical-1220133023112011-0133103323321013-2230232111003000-1311033011331123-2002332230231132-0210313333022312-1133232021122222-0311333013311320"></a>

<a id="canonical-2223100100003123-1232231022032302-3300010033210300-0213101020200111-1221002032320301-0303223221331132-0331022123113131-3112211221003130"></a>

#### `use_tls.use_server_verification.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-0120132103213330-1221300320110330-2321301033111133-2013000231030120-3331212310220231-0131231033221131-2002231222231110-2310012232111211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_server_verification.trusted_ca` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- [use_tls.use_server_verification](resources--origin_pool--reference--group-003.md#canonical-3000330133311333-2312222030122331-0220231200123233-1320011200021220-0002310301112001-0231020200121203-2330012211110321-3000301310100002)
- use_tls.use_server_verification.trusted_ca

<a id="canonical-3211120110011201-0121003012003331-3200121102030103-3100212002221031-3100111133302033-0212211223332131-0003023220323223-0303330120232302"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310322322300221-1013011011301211-0131221202310311-2130033322231100-3233111220133121-1201333021223121-0210002100210021-3333031121020132"></a>

### Direct properties for `use_tls.use_server_verification.trusted_ca`

<a id="canonical-2300000132123033-2103011001320010-0003201212202311-2131202021300121-0100301332031130-3100031203003111-3213000321133131-3220001111032321"></a>

#### `use_tls.use_server_verification.trusted_ca.name` property

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

<a id="canonical-0121211112011311-3022002120201003-1200133210002120-3103310112300321-1331332330013221-1322110233131121-2221323113132132-0330103213300331"></a>

<a id="canonical-1230303213202211-3310231021222022-3132223300113103-0211030113012310-2321102002112223-2212220323331023-3202003113131111-0321032200210332"></a>

#### `use_tls.use_server_verification.trusted_ca.namespace` property

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

<a id="canonical-3323101123113223-3112300031313201-0323233213330222-1122012021031001-1213311231311210-1321012323131120-1220113301313301-2232122310332010"></a>

<a id="canonical-1320222031003013-1333202031030121-0110103021102000-0112102233101131-1313202221310320-1212213301130112-2313202002233303-0111131101013213"></a>

#### `use_tls.use_server_verification.trusted_ca.tenant` property

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

<a id="canonical-0303020022031131-2122331331231202-2331332210010010-1201301230030123-0022302320133221-2300002131320312-3100230312312211-2111020303131110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [use_tls](resources--origin_pool--reference--group-003.md#canonical-0120300213231231-1220111330332003-0132322111012203-1003133000222320-0301223123302330-2321023020333113-0032332321003313-1233232210333133)
- use_tls.volterra_trusted_ca

<a id="canonical-1002302310301002-1302021231211331-0023032120010312-1013223211301323-1323121220130010-2023101130223113-3023221113032221-2012303103232033"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
volterra_trusted_ca = {}
```

This is an empty object or choice marker. It has no direct properties.
