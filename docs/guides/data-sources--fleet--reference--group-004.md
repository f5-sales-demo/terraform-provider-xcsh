---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-3221013101022303-0031020003333031-1030012032322102-2001322333132322-0323001032121312-3112313131203323-0220202103032320-2230222202110303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.subnets` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- storage_static_routes.storage_routes.subnets

<a id="canonical-0121220231232033-2201120120322120-2020231221003131-1112312103331211-1033023301303323-0003123032310111-0003201333213233-1113303033221230"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-2232232012202220-2021312303232001-1331312121030033-0303233312222010-3023110013210120-0113130213212031-0312202100202202-1020231030301301"></a>

### Direct properties for `storage_static_routes.storage_routes.subnets`

- [IPv4](data-sources--fleet--reference--group-004.md#canonical-1132211300313231-3232222201233203-0221010020302300-3323323100303012-2223321212111321-1130121030332303-1000311123133212-3333130120132012): complete subsection reference.

- [IPv6](data-sources--fleet--reference--group-004.md#canonical-1311003212203132-3330033022230231-0100323000203331-0023132223330130-2131030223221112-1223012010100231-2203010313311022-2010130333101313): complete subsection reference.

<a id="canonical-1132211300313231-3232222201233203-0221010020302300-3323323100303012-2223321212111321-1130121030332303-1000311123133212-3333130120132012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.subnets.ipv4` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-3221013101022303-0031020003333031-1030012032322102-2001322333132322-0323001032121312-3112313131203323-0220202103032320-2230222202110303)
- storage_static_routes.storage_routes.subnets.IPv4

<a id="canonical-3111021123013120-1113020100122333-0331321321332011-1130023020312203-3021313031121130-0120232323121310-2323100221032020-3213122112301233"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301122333000303-1221320230203233-1032121123333123-0301213332102210-1232213101133101-1223021302212100-0112120203023231-3133103310022203"></a>

### Direct properties for `storage_static_routes.storage_routes.subnets.ipv4`

<a id="canonical-1010201232103233-3212022220013201-1002123030101122-0302001220102323-0102322310321333-1021012131133233-2330301101201121-1301020122301222"></a>

#### `storage_static_routes.storage_routes.subnets.ipv4.plen` property

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3330122023022133-3021202132103230-0202020213021100-1330111313110213-1222233312330011-2021023332121002-0232120230103223-3222203120111201"></a>

<a id="canonical-0322322131303101-2213020203312311-0120031020200100-0312330221012220-3301312212211301-3221203231313221-2310311313320232-0013210301033303"></a>

#### `storage_static_routes.storage_routes.subnets.ipv4.prefix` property

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-1311003212203132-3330033022230231-0100323000203331-0023132223330130-2131030223221112-1223012010100231-2203010313311022-2010130333101313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes.subnets.ipv6` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-003.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-3221013101022303-0031020003333031-1030012032322102-2001322333132322-0323001032121312-3112313131203323-0220202103032320-2230222202110303)
- storage_static_routes.storage_routes.subnets.IPv6

<a id="canonical-1330332231122313-3220022201313212-3332030102322033-3021113101203313-2003130320102221-0320322202003210-3203013113103010-3333033111020013"></a>

Type: `"single"`. Computed.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1132312323222133-3233303311302212-1220111221023310-0233300231002012-3221210303030210-3020331011330133-0130232013210022-0100220222203330"></a>

### Direct properties for `storage_static_routes.storage_routes.subnets.ipv6`

<a id="canonical-2232103210312231-1221103313021222-0123100131001300-1200212030320313-1001231022002031-0320022201220110-2131100210201001-1313203333010000"></a>

#### `storage_static_routes.storage_routes.subnets.ipv6.plen` property

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-1112320312302203-0021021011033332-0032012322211232-2012212033000321-3202031300201103-0032002230032012-1022200231333022-2021310011201221"></a>

<a id="canonical-3322302111302111-1120301010233233-3311300100030030-3310330310301213-1220031320112120-0211101310202201-3331330131020110-2221313221320222"></a>

#### `storage_static_routes.storage_routes.subnets.ipv6.prefix` property

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Additional upstream details:

IPv6 address must be specified as hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0"
The address can be compacted by suppressing zeros e.g. "2001:db8::2::"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2132300120030302-2100113132002202-2230322111111011-2301031030022032-0303102003030130-2111000030223010-1011003201131323-2230112300300113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `usb_policy` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- usb_policy

<a id="canonical-0031113210110330-1320123122203030-3102000002211020-0300002012223023-3100323131030300-1303202220232201-2302030110112130-2113203111101330"></a>

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

<a id="canonical-0000313021232330-1323201102333031-1221303131021210-3301331031301211-0232332022122311-0023110111033333-0212033301333321-1002100222121102"></a>

### Direct properties for `usb_policy`

<a id="canonical-0301221010202211-0132101021111021-1312222323232010-2212232001211122-3211332100232100-1323120210013011-3020222032301300-3132112021201001"></a>

#### `usb_policy.name` property

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

<a id="canonical-0003220322203031-2113233001023022-2013102230031303-0031013310022321-0122032333010211-1010010131111121-0333333203300100-2320010023301322"></a>

<a id="canonical-3210132013011232-1312130111102132-0133113321233211-2320032220310202-2312122301133120-3201012020121112-1302121231001022-2101131213220002"></a>

#### `usb_policy.namespace` property

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

<a id="canonical-2011103303103012-0131103303222023-1331321033102010-0030230212112233-0301102001302101-3113023320012123-0322223321300332-1232203200023021"></a>

<a id="canonical-2123210022331012-0320211112323311-0113030201312333-0121112302200203-2203322203231022-3013210321203222-1113320102311313-3002132002311121"></a>

#### `usb_policy.tenant` property

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
