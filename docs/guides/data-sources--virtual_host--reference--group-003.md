---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-3300211200203031-0102030220211032-2101332013321022-2020233013011313-3111012122302233-3321322111201122-3301232102311203-1002222023222302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_type.app_firewall.app_firewall` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [waf_type](data-sources--virtual_host--reference--group-002.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- [waf_type.app_firewall](data-sources--virtual_host--reference--group-002.md#canonical-1333300322323113-1232131101000031-2233210011203232-3322023300230033-1233133023310221-1233001022120331-2003112011233330-2313020200000013)
- waf_type.app_firewall.app_firewall

<a id="canonical-3031320302232103-0313103102301121-3310310322030220-1331222233331322-3323312112300120-0132330303223100-3230033103103121-1302130133000100"></a>

Type: `"list"`. Computed.

References to an Application Firewall configuration object.

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
    "ves.io.schema.rules.repeated.num_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  }
}
```

<a id="canonical-2313312001123103-0212312030231133-1133321222322032-2301313200203331-0000131313331022-2223130110311103-1300011232323033-2123133310103110"></a>

### Direct properties for `waf_type.app_firewall.app_firewall`

<a id="canonical-2122301313000122-3213122132202331-0303320313121310-1013233322231122-2033012022001310-1012311223231011-1122320130002211-0023012113203300"></a>

#### `waf_type.app_firewall.app_firewall.kind` property

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

<a id="canonical-2210021022020201-1101003131010303-1033011311003033-3313112311133020-2212001121211021-1100213001032311-3233011301030120-1112313221311113"></a>

<a id="canonical-2103012331212031-1333130101300122-1012310113022322-0201221003301311-3011000332212230-1001131232323223-3033101231302332-1212232213201300"></a>

#### `waf_type.app_firewall.app_firewall.name` property

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

<a id="canonical-3033002313132221-2232211031303012-3110113302303333-3021112021230110-3321323113210103-0223221211001323-0302100133213212-1021312212203133"></a>

<a id="canonical-2330213213302112-2223011003232111-1321012213023020-2212203113200002-2122002320332113-1202123203023122-2132300003313031-0331011320132130"></a>

#### `waf_type.app_firewall.app_firewall.namespace` property

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
  }
}
```

<a id="canonical-3331131032302210-1323121002333201-1230311011223103-0121133132223223-0130311111330201-0301032302331131-3303320103113200-0311132221113133"></a>

<a id="canonical-1310130333321133-0221032102000123-1332032020232033-2313123112111013-2231232323313122-3322303213023202-2213323330322002-3233230210012101"></a>

#### `waf_type.app_firewall.app_firewall.tenant` property

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

<a id="canonical-1013211020013021-0011321100123322-0231222213313331-2021120130333111-3211303102212321-2133230011012132-0202012022103100-1012310300010120"></a>

<a id="canonical-0113012131012321-2023001120121203-1110300100120311-2213303031113133-3020123100232210-0022031112013202-1213212001230303-2220332231020130"></a>

#### `waf_type.app_firewall.app_firewall.uid` property

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

<a id="canonical-2230010330302310-1312132323300100-3023320210221023-2030200010011000-1310013232033111-0133103323302222-1021312002012220-2103031000003013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_type.disable_waf` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [waf_type](data-sources--virtual_host--reference--group-002.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- waf_type.disable_waf

<a id="canonical-0013120130201131-1002011002120203-2223320033320023-3202121202131112-3200222111030331-3131000003002201-2333113123200223-2321301120031311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable waf.

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

<a id="canonical-3111320320023322-3033330033011220-2031221222111000-3211122202020221-2303103121031102-0213033333303222-0321322313130302-2022131333210103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_type.inherit_waf` properties

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [waf_type](data-sources--virtual_host--reference--group-002.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- waf_type.inherit_waf

<a id="canonical-3010301221131302-2113032200130303-1223220212210300-0101211011121200-2020201320122112-3013320110313203-2213221213333033-2120022000130002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherit waf.

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
