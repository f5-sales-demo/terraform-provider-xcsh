---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-0001033333001312-1110000223233102-1301013231111010-2202132121010201-0001211300033022-1122123130321202-1323202300302333-1310032301322011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.inside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-001.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.inside_network

<a id="canonical-1023300302332022-1002212333203103-0200330012333210-0302332010031111-1320013012311210-0100021132200333-1022300201013233-2313233023220331"></a>

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

<a id="canonical-2032033302122022-1103111020001001-2303301333213001-3102322130322113-3113231020002223-3033131021221332-2021231222100202-3011311100002022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.outside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-001.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-001.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.site_locator

<a id="canonical-0210110101223301-1332211132103003-2122101201000300-3232202220003030-0323132200322021-1031030230301100-3013303132331112-1301002231102332"></a>

Type: `"object"`. single nested block, Optional.

This message defines a reference to a site or virtual site object.

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

- [site](resources--origin_pool--reference--group-002.md#canonical-3100133211320001-3312100030232320-3031301010230301-3330030203200230-0033031210223320-2010200033333010-3223011122110301-3320211111203031): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-002.md#canonical-0022302012110203-2130330132230330-3101130123013031-2332313300310132-0313113133213110-3100010223101033-2233301002233221-1303300221201021): complete subsection reference.

<a id="canonical-3100133211320001-3312100030232320-3031301010230301-3330030203200230-0033031210223320-2010200033333010-3223011122110301-3320211111203031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.site_locator.site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-001.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320)
- origin_servers.k8s_service.site_locator.site

<a id="canonical-3132212002001320-1011333321223332-1320123123221213-3111323223031202-3012133121203111-0202303002033130-3131332300331001-2333322330123022"></a>

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

<a id="canonical-2000013301321012-2001230112000330-0201112100203003-1113213330132032-1300121001130102-0122300001311330-0002333212332130-0011100301221220"></a>

<a id="canonical-1031231100223010-3212110013333133-1112132300222031-0120333331310013-0202310313223323-2202313303203222-1300322110222311-2212211112031102"></a>

#### `origin_servers.k8s_service.site_locator.site.namespace` property

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

<a id="canonical-2210010001330323-0311122232312312-2111301223332213-2200102000133300-0032322003322331-3211200212121033-2333002310233103-2320202311320212"></a>

<a id="canonical-0022022201110031-2031023123203111-0031131233022312-3011211221302322-2000023112122003-2321301030302012-2201202130320033-3301213022200303"></a>

#### `origin_servers.k8s_service.site_locator.site.tenant` property

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

<a id="canonical-0022302012110203-2130330132230330-3101130123013031-2332313300310132-0313113133213110-3100010223101033-2233301002233221-1303300221201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-001.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320)
- origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-2202122132021231-3321020200323200-1310202030313310-1330230003113200-2023120102022020-3101323010320321-0320133022010001-3312132023330131"></a>

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

<a id="canonical-0312030111303223-0203213101223030-1001010012130330-3010221020030021-0031302012311123-2301331321021112-1202202012301222-0111103021310012"></a>

<a id="canonical-0021011021222132-3033301003213320-3032020013223203-3302230230003023-1203220311101010-2130312120233212-1023122122323110-0031222113010101"></a>

#### `origin_servers.k8s_service.site_locator.virtual_site.namespace` property

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

<a id="canonical-2113321233311132-1131200212200302-2320031233131202-1111221313012210-0131110123231321-3201303030032220-0202212002031331-1203222023210300"></a>

<a id="canonical-3301033033133201-0112002310011101-0031211101201223-0022321200021130-2032022202030000-0303101002113032-3300233321203331-0103131323210303"></a>

#### `origin_servers.k8s_service.site_locator.virtual_site.tenant` property

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

<a id="canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-001.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.snat_pool

<a id="canonical-1313332303200303-0131030312220311-3230121111020130-1203211302210213-3301211032330300-0012033312120030-3301312322223121-1032312222131033"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

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

- [no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-3120030123300002-3302220202331111-2130210203032133-1210003121120212-1331220113313303-3023331330303323-2302302121321021-1131001030133123): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-0203100021110200-1222213130002321-1303223321201100-1222003111122102-3031102313132102-2212033122301233-3100012230012301-3221313203112130): complete subsection reference.

<a id="canonical-3120030123300002-3302220202331111-2130210203032133-1210003121120212-1331220113313303-3023331330303323-2302302121321021-1131001030133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-001.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-001.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-001.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.private_ip

<a id="canonical-0132122003323203-1113202003201132-3321320022031212-2230321230311222-2210101323121301-3303110010030232-3303110003233000-0303302222121003"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public IP address and site information.

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

- [inside_network](resources--origin_pool--reference--group-002.md#canonical-2333033312003333-2020012201210130-2333321300102023-2223301201133333-1131320133122101-3322011331100320-1203013102230121-3220223120201200): complete subsection reference.

<a id="canonical-3131321310122213-2013013222233311-2013100113320312-1111123221322122-3013030033122032-2032133200310211-2223132120100033-2301131030132321"></a>

<a id="canonical-2212033312111111-2020120120121321-1022023233122332-1113113003102210-0311201113130113-1322332231212120-0212023123333331-1133111123233020"></a>

#### `origin_servers.private_ip.ip` property

Type: `"string"`. Optional.

IP. Exclusive with \[\] Private IPv4 address.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [outside_network](resources--origin_pool--reference--group-002.md#canonical-2332312302210220-3001332333131122-1100220121131321-2233230013312312-2320320220221301-1130102233012220-2120000130122303-2120113222003130): complete subsection reference.

- [segment](resources--origin_pool--reference--group-002.md#canonical-3330110221310112-1110312123133100-2123130300111320-2011312003131331-2323030230133212-2300303013130222-3200300203112331-0000030201112330): complete subsection reference.

- [site_locator](resources--origin_pool--reference--group-002.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012): complete subsection reference.

<a id="canonical-2333033312003333-2020012201210130-2333321300102023-2223301201133333-1131320133122101-3322011331100320-1203013102230121-3220223120201200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.inside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.segment

<a id="canonical-0122201332003012-3133322332232022-1210211032033221-1113102110020313-3132031203113300-1100300320102320-1110112012323103-2100131212133310"></a>

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

<a id="canonical-0002023023112303-3000002202301301-2122112011220210-2300210311301112-1310303122333301-3023310112222330-3112020232133100-3013231300223301"></a>

<a id="canonical-1121210302211112-2102331332022031-1213001222330111-2220210300030101-2302202110201300-2231023332321331-3000101101022232-0123032010313203"></a>

#### `origin_servers.private_ip.segment.namespace` property

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

<a id="canonical-2211231312012130-3330102313111132-0010213002332311-3333322201120203-3210311203321223-2202121223213200-0210312203332122-3122031211132310"></a>

<a id="canonical-2320011113321101-3300131001001003-0112231333023222-1300023323013010-1100031331100210-0210231002210023-3103023222331030-3230220333120011"></a>

#### `origin_servers.private_ip.segment.tenant` property

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

<a id="canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.site_locator` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.site_locator

<a id="canonical-2300231310100223-1133011131133232-1022201302332233-0201232012331201-2310030012112201-1312110221032323-1012213021333220-1123101310031303"></a>

Type: `"object"`. single nested block, Optional.

This message defines a reference to a site or virtual site object.

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

- [site](resources--origin_pool--reference--group-002.md#canonical-0033003220302032-1233233113310201-1221132212023222-0023132332022003-3233200113230112-0210031031003020-1120331220022312-1302013132002100): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-002.md#canonical-2333130022331103-1112130211201220-0203022112321012-1323001330210201-2223232033131220-1223103200121120-2323132201200131-1232331312223321): complete subsection reference.

<a id="canonical-0033003220302032-1233233113310201-1221132212023222-0023132332022003-3233200113230112-0210031031003020-1120331220022312-1302013132002100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.site_locator.site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200)
- origin_servers.private_ip.site_locator.site

<a id="canonical-2331020232213023-3031110031113100-3210011130022031-1310130031320000-1301021211133300-0320112000033032-2013320111122100-1133023031220313"></a>

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

<a id="canonical-0310332310121302-0012001110320332-2133332133310001-3302331031221212-1122023203031010-3121012220302020-0011011002012001-3331221202100202"></a>

<a id="canonical-0310330121031203-0021000020111322-3321313103101321-2232002000031133-2023031121002111-3102112200012321-0111010130230213-0311230122021210"></a>

#### `origin_servers.private_ip.site_locator.site.namespace` property

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

<a id="canonical-3333130012100110-3133010102330332-2230012130101130-2032111033113120-2101210230323002-3121102032110030-0031220330012002-1022302023022111"></a>

<a id="canonical-3013333010012010-1201221213132121-3111100121312122-0001203202102100-0301312010022102-1333302302000202-0332302310033331-3232112102230310"></a>

#### `origin_servers.private_ip.site_locator.site.tenant` property

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

<a id="canonical-2333130022331103-1112130211201220-0203022112321012-1323001330210201-2223232033131220-1223103200121120-2323132201200131-1232331312223321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200)
- origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-2222230231021133-1310332222203132-2213203023213220-0203330233023131-0122101031012301-0203303032321110-2230021233220330-2121021123221222"></a>

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

<a id="canonical-2200220020123203-3033213310312002-0113221220033300-1103023110203201-3032000302331022-2001131023203233-0200213310311000-0133101003022121"></a>

<a id="canonical-0011202022201311-0112210121132020-1320100000311033-1112222013000101-0103203202213112-1100300212013021-0232302310302223-0330022313111301"></a>

#### `origin_servers.private_ip.site_locator.virtual_site.namespace` property

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

<a id="canonical-0323020000231022-2321300331100223-1221131022231231-3322133121100023-3202320023222230-0110102231120112-0202310122123211-0220310023123201"></a>

<a id="canonical-3132110211120020-3323320110100210-1212200221121321-0123011121211200-0002322013033021-2202021123233010-0020201203202211-2303001020031211"></a>

#### `origin_servers.private_ip.site_locator.virtual_site.tenant` property

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

<a id="canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.snat_pool

<a id="canonical-1230022110130033-0122300323220323-2233303320012120-2123001031323003-0331000322013032-3031212033032211-0210003203200313-3110000013232222"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

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

- [no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-2223123321310311-0022101202202202-2313330102010022-2220011122020023-1030013111210103-1122131100131311-1000123233211222-1311201133113313): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-0223213203101123-1313311203312233-1030000202323301-1000222111110203-0212020020331110-0120321020133203-1022120113221100-3331221220332032): complete subsection reference.

<a id="canonical-2223123321310311-0022101202202202-2313330102010022-2220011122020023-1030013111210103-1122131100131311-1000123233211222-1311201133113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.private_name

<a id="canonical-3022300021030132-2220222012003231-1032303111232231-2212212013120222-2323130212232113-0112320023321113-2313230210023201-3013310312203012"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public DNS name and site information.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [inside_network](resources--origin_pool--reference--group-002.md#canonical-2203123102332202-2332013300310111-3022033222313002-2213133113121032-1122013013203332-2023020212333100-3023233133301122-0311232310323331): complete subsection reference.

- [outside_network](resources--origin_pool--reference--group-002.md#canonical-3120201313033031-2111011302310031-3230023222102013-3032031102130021-1212010010233033-0030201131221331-0112102022203121-1231311201021120): complete subsection reference.

<a id="canonical-1001321013320301-3231130133303213-1012133100211202-0202310003112322-0320013300303023-2210013102000202-3300202320012331-0013231020013230"></a>

<a id="canonical-0032311010112203-2211032022002131-0123203110023121-1122101210231021-0021230110323230-2033021302111032-2110211230123011-0213121132331010"></a>

#### `origin_servers.private_name.refresh_interval` property

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

- [segment](resources--origin_pool--reference--group-002.md#canonical-1200322213030001-3020120330301111-1323013131020133-1113202333233202-1000222302203230-1223201032001001-3323222331313231-1020213223310111): complete subsection reference.

- [site_locator](resources--origin_pool--reference--group-002.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203): complete subsection reference.

<a id="canonical-2203123102332202-2332013300310111-3022033222313002-2213133113121032-1122013013203332-2023020212333100-3023233133301122-0311232310323331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.inside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.segment

<a id="canonical-1133101023013231-2111002010302323-2132111310121133-1101003132011232-1310123201023212-3333020132201122-2000001333102003-0200320101132013"></a>

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

<a id="canonical-2213202322130032-1302102012000111-0130011113033110-0332003132300102-3013320111020321-0123202013202103-2200010222222002-1313303132130320"></a>

<a id="canonical-3011031011021330-2012122131230112-2132333303032132-3320032030331330-0002011321111231-3312231023101113-0020202231221222-0110223213211112"></a>

#### `origin_servers.private_name.segment.namespace` property

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

<a id="canonical-0020132303013213-0000202322312013-0130223103003201-2320022211103132-1011010001103111-3210213313122021-0323000330112332-0230310211022122"></a>

<a id="canonical-0301033310011303-3313222012210101-2001331301122303-1131202330020220-1103001331113203-1120211330211321-0000321333011330-1233231202220031"></a>

#### `origin_servers.private_name.segment.tenant` property

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

<a id="canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.site_locator` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.site_locator

<a id="canonical-1020001220021113-3201030222113010-3233123302013110-2033130012033221-0330122021331100-1113302312232110-2312031322302123-3113203103030302"></a>

Type: `"object"`. single nested block, Optional.

This message defines a reference to a site or virtual site object.

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

- [site](resources--origin_pool--reference--group-002.md#canonical-3202010311011020-1202031221320113-3122030101221230-2200102322211100-2123330202232233-3001123100320133-3112032000130220-0233002310102130): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-002.md#canonical-0310120101132210-0322321121303212-2123011032321233-0312231230321133-2231302322012302-2001330223222303-0223221321130131-3023110131010102): complete subsection reference.

<a id="canonical-3202010311011020-1202031221320113-3122030101221230-2200102322211100-2123330202232233-3001123100320133-3112032000130220-0233002310102130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.site_locator.site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222)
- origin_servers.private_name.site_locator.site

<a id="canonical-2310110312030231-0323110012101103-1231313132110123-3322003102003131-1320220002300020-0210320122211021-3233202103231300-0312122000232233"></a>

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

<a id="canonical-2030110230122011-0121312231332013-0103303202223333-0202112032330303-0111300300121002-0020021300130230-3302131213012013-1302221311122312"></a>

<a id="canonical-1122331112312113-2332233321220222-0031330211120013-1112103201313213-1202333120030333-3332011232200231-3031121031121230-3131121232311210"></a>

#### `origin_servers.private_name.site_locator.site.namespace` property

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

<a id="canonical-3122311300323320-0211222202200211-3102330002320302-3333310303332213-1212202113332222-3222331232300233-2232323000323223-1033321013103023"></a>

<a id="canonical-1131323032213200-2211331220220101-2030321331132230-1323023210112331-1121322111120230-2323000302323102-3320011121031111-2213220303110013"></a>

#### `origin_servers.private_name.site_locator.site.tenant` property

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

<a id="canonical-0310120101132210-0322321121303212-2123011032321233-0312231230321133-2231302322012302-2001330223222303-0223221321130131-3023110131010102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222)
- origin_servers.private_name.site_locator.virtual_site

<a id="canonical-0310202213011000-3111221033301021-0000311230003001-3133110320333332-2332303200200020-1133002033223110-1113111032333320-0303223221312122"></a>

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

<a id="canonical-1211202113232100-2103013212012101-3223031333022230-2020222132122330-1013333210103312-2133112122232022-0231130221203310-1301032203320222"></a>

<a id="canonical-3013012201102123-1033102123030303-1020323132133130-1100013131103302-3131130232010301-3000123221010203-3023233211113221-0122210201130012"></a>

#### `origin_servers.private_name.site_locator.virtual_site.namespace` property

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

<a id="canonical-3333120331021200-0122122122112221-1033232311130311-2001001031121012-1323101121200221-3031223233321322-1020232202213333-0032031311223013"></a>

<a id="canonical-0310111030121122-1023013000311201-1221021332003310-0132102323220013-1312022323323322-3201311033012302-3102231330021130-1002321213220010"></a>

#### `origin_servers.private_name.site_locator.virtual_site.tenant` property

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

<a id="canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.snat_pool

<a id="canonical-1012322211101010-0203023132110100-0321333121033313-1121311033211030-0012132301213012-1133130003230120-0020130231013323-3113120111010133"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

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

- [no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-3130023221023210-3020131101000132-2331203133232103-1230023310301102-2223220110200130-0220200112120302-3310132133311233-3200112023312112): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-2221131331021133-1131233031222100-2123110022013001-1003330121330001-2001002310023110-1111101220230032-3332220113200220-3133303312002021): complete subsection reference.

<a id="canonical-3130023221023210-3020131101000132-2331203133232103-1230023310301102-2223220110200130-0220200112120302-3310132133311233-3200112023312112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.public_name

<a id="canonical-0222302133210320-2001221203003010-3022300302110110-3000333100030103-2313301012110103-2212210120131013-1012130222122313-3202312023112330"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Receipt-pinned upstream constraints:

```json
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](resources--origin_pool--reference--group-002.md#canonical-3122100030131322-0113010211030321-3203333130200133-1111310203002032-0020302020200012-0010322030112332-3310313213332211-3231002332012330): complete subsection reference.

<a id="canonical-3122100030131322-0113010211030321-3203333130200133-1111310203002032-0020302020200012-0010322030112332-3310313213332211-3231002332012330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.vn_private_ip.virtual_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-2130122232222011-2300321030110011-1031031023222200-1001031200312013-2120002123011312-1310133101323000-3001103213310123-2023322321313313)
- origin_servers.vn_private_ip.virtual_network

<a id="canonical-0003212300322202-3222000200122020-3203010013001101-1311323200003200-2102323331103133-0230313302203300-3302211300312303-3202332311301130"></a>

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

<a id="canonical-0223223101122021-2213133211011111-1110303311033212-2202212310201012-1232122321032311-0022200133002320-1203111330100322-3332031233320120"></a>

<a id="canonical-2211231211330221-2032031113021110-1213213100021003-1033233311120332-1320210323301002-2232301111113131-1332221203312030-1100003133312110"></a>

#### `origin_servers.vn_private_ip.virtual_network.namespace` property

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

<a id="canonical-0111322301320300-0322201111323203-1212030301001200-0211310021223012-3123201121032322-2030010332303230-2231323310021331-0120323133322133"></a>

<a id="canonical-0310322312201233-2000332330303222-2312030031212312-1011333103200111-0132332212221311-0211213000102210-0221100201203231-1022220113220320"></a>

#### `origin_servers.vn_private_ip.virtual_network.tenant` property

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

<a id="canonical-1013210030120200-0300013133323110-0211222301220003-2111231123223232-2130003221211000-0010213100230320-3322112211323121-2102120111321112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.vn_private_name` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.vn_private_name

<a id="canonical-1133331232233003-1320111130012311-1123222301130121-1020122210302320-1322110011220312-3110310013210300-1130202300310203-1332113133000131"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with DNS name on Virtual Network.

Receipt-pinned upstream constraints:

```json
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [private_network](resources--origin_pool--reference--group-002.md#canonical-3313202330121310-0332121101021311-0020212111023330-2010022013100332-0131010011102330-1300031011320320-3223132122001110-3220213212120000): complete subsection reference.

<a id="canonical-3313202330121310-0332121101021311-0020212111023330-2010022013100332-0131010011102330-1300031011320320-3223132122001110-3220213212120000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.vn_private_name.private_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.vn_private_name](resources--origin_pool--reference--group-002.md#canonical-1013210030120200-0300013133323110-0211222301220003-2111231123223232-2130003221211000-0010213100230320-3322112211323121-2102120111321112)
- origin_servers.vn_private_name.private_network

<a id="canonical-1133303321023212-2001300031121022-1011301211231302-0001021121033021-2323331011202103-0022221131021002-1223211000123212-0002221220103230"></a>

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

<a id="canonical-3001031121122201-2021001100001030-1223130321201021-2103201022330133-0031030202301102-2003332030311030-1013301002031230-0333123312320300"></a>
