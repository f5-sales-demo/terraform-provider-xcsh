---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-1120130030212213-1201322300012332-2323302012303013-0101211222220130-3222232201332103-3301012130300012-1320232031331211-1023200303311102"></a>

## Direct properties for `bond_device_list.bond_devices.lacp`

<a id="canonical-1323023103302121-0032331200203312-2332003031022310-0330120232202200-3201112223202103-1312211203131312-3011130130010020-1000302021202302"></a>

### `bond_device_list.bond_devices.lacp.rate` property

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-2112312321330302-1211021232313121-2331211202132011-1132132332222033-2303221021220203-0002022311322203-1311030023321330-2133303232003230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dc_cluster_group` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- dc_cluster_group

<a id="canonical-1323102031301102-2302023121333302-3003033313301001-3200222232333112-3221033022320300-2213112303112112-3111100331223113-3022013100211010"></a>

Type: `"single"`. Computed.

\[OneOf: dc\_cluster\_group, dc\_cluster\_group\_inside, no\_dc\_cluster\_group; Default:
no\_dc\_cluster\_group\] Type establishes a direct reference from one object(the referrer) to
another(the referred). Such a reference is in form of tenant/namespace/name.

Additional upstream details:

This type establishes a direct reference from one object(the referrer) to another(the referred).

Receipt-pinned upstream constraints:

```json
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

- [dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-1323102031301102-2302023121333302-3003033313301001-3200222232333112-3221033022320300-2213112303112112-3111100331223113-3022013100211010)
- [dc_cluster_group_inside](data-sources--fleet--reference--group-002.md#canonical-0320130313200120-1222323131231112-2121322303301022-1322103111131230-1212313323223121-0132201113112130-0230132121302013-3023123020233011)
- [no_dc_cluster_group](data-sources--fleet--reference--group-002.md#canonical-1100023233033321-1031202011312123-3220000313212323-3201320320113320-0301111211231330-3010123121312200-2311202132022322-1301230111210222)

Select alternatives according to the provider validators above.

<a id="canonical-3330003313132102-2201221210110322-1133312102021313-2323003300311221-1301332313233003-0110323033301322-2001110101121302-0000000200033113"></a>

### Direct properties for `dc_cluster_group`

<a id="canonical-3101200231322220-0030110100012122-1122030302023033-2100213222203322-3013110111233233-0030013331212201-2013220320021112-2211001233110202"></a>

#### `dc_cluster_group.name` property

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

<a id="canonical-1111002223033100-3010002312233113-1101011133233312-3311322232030220-1230333231031011-3103100222030202-3221230220300321-0133122311311121"></a>

<a id="canonical-2100120000331133-0223232303023122-0332211033013121-2202132311333300-2222020232021320-0121210330011001-0030312231031123-3201120110112000"></a>

#### `dc_cluster_group.namespace` property

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

<a id="canonical-3200332321310101-1113002021233023-3220132123323222-2110110130020113-3221211221011221-1011132331321031-1212013120021113-3033333033223311"></a>

<a id="canonical-2221032010303311-3123232201210120-1113131303310230-3320210210001021-0003333201322123-0112310203121213-0001000101300012-1331321311232111"></a>

#### `dc_cluster_group.tenant` property

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

<a id="canonical-1011213321333020-0122213232122233-3330203123313132-1102003220000100-0310311103303310-1021233230221210-2033031202202022-0212001013131023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dc_cluster_group_inside` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- dc_cluster_group_inside

<a id="canonical-0320130313200120-1222323131231112-2121322303301022-1322103111131230-1212313323223121-0132201113112130-0230132121302013-3023123020233011"></a>

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

<a id="canonical-2021000313331130-0110201312112323-3103213001302123-3232212101321103-1102001012330212-0230302103211112-0322103030202021-3123111310123123"></a>

### Direct properties for `dc_cluster_group_inside`

<a id="canonical-3220102313322131-2231121010232020-2131232311230230-1221120000212310-3001322301212300-2100222002022212-3022230203131020-3021120101133023"></a>

#### `dc_cluster_group_inside.name` property

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

<a id="canonical-0110332101221300-3103210103011130-0202021120021002-2122213331232223-0001312201021210-2001032102103311-0222231100133011-3032320301103030"></a>

<a id="canonical-1331212303223200-1101012132121323-2322300021112000-3012303033111102-0010223212022032-3113013332302202-3211003311003311-2313223311102210"></a>

#### `dc_cluster_group_inside.namespace` property

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

<a id="canonical-0131233301311000-1023230333031322-2020212213332313-2131013121300000-1312222200223120-2300213320213113-0310300100232020-0200102313012301"></a>

<a id="canonical-2332302022023213-0112233101300311-0131120103332301-3301213230002120-3332013313213113-2201113312030211-0103030330220023-1223332230230320"></a>

#### `dc_cluster_group_inside.tenant` property

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

<a id="canonical-1232232122223230-0300310232023033-2100332200122130-0211200101321332-0000311133222122-0230121123110011-0201012012212022-1221302112223212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_config` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- default_config

<a id="canonical-3123101121002222-0023112111233223-2020310133003302-2100022113000013-2020323002023330-0112103301320002-1313303103032320-1010021113303113"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_config, device\_list, interface\_list; Default: default\_config\] Enable this
option

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

- [default_config](data-sources--fleet--reference--group-002.md#canonical-3123101121002222-0023112111233223-2020310133003302-2100022113000013-2020323002023330-0112103301320002-1313303103032320-1010021113303113)
- [device_list](data-sources--fleet--reference--group-002.md#canonical-2011300210203112-1212221211332200-3213022223032230-0000021313221110-2102020102331131-0130001302313130-3211210123310210-1201013002123013)
- [interface_list](data-sources--fleet--reference--group-002.md#canonical-1121100011030011-0112201021211330-0232123013200213-0013132120303311-0211002302121000-0323003013300122-0112213121000233-1113331010200030)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033001023033030-3300110122022001-2102101121002010-0102011111321302-3323110310023332-2112010011212302-1110320322001132-3102320123230202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_sriov_interface` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- default_sriov_interface

<a id="canonical-1331130211111001-0122132322111332-0020302032000110-2211333211111331-3220132310201330-0330113111110013-2203120013110032-0101011122133233"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_sriov\_interface, sriov\_interfaces; Default: default\_sriov\_interface\]
Configuration parameter for default sriov interface.

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

- [default_sriov_interface](data-sources--fleet--reference--group-002.md#canonical-1331130211111001-0122132322111332-0020302032000110-2211333211111331-3220132310201330-0330113111110013-2203120013110032-0101011122133233)
- [sriov_interfaces](data-sources--fleet--reference--group-002.md#canonical-2123333323013300-0110133102200320-1232210230303220-3011003312322200-3022100322000212-0311121301321002-2220103203200120-2223000320220210)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130302203111131-3121312201232200-1002223322032001-3203202102210303-3010121002122122-3101131223332203-1300222311301312-0311213113232021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_storage_class` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- default_storage_class

<a id="canonical-1230313022311013-2033130033232303-3113323201222203-3110331212320020-2102211123032221-1303110110332232-2222211212001213-0322312320120321"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_storage\_class, storage\_class\_list; Default: default\_storage\_class\]
Configuration parameter for default storage class.

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

- [default_storage_class](data-sources--fleet--reference--group-002.md#canonical-1230313022311013-2033130033232303-3113323201222203-3110331212320020-2102211123032221-1303110110332232-2222211212001213-0322312320120321)
- [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-1012302103312022-3331132220002110-1120220121220023-0000311103003122-0111002031331211-1000002000021111-3101030221022312-0130130013123023)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022132120103110-3011033030320022-2121100111221111-1030322032011211-2033202230113022-0112303030121313-2001001200312221-3320303123030211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_all_usb` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- deny_all_usb

<a id="canonical-0323201230021333-3213210223102010-1222323310230210-1003010300303302-3103023120303323-3331021111202320-1202222300110102-3301322010230232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for deny all usb.

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

<a id="canonical-0322000310332200-0133003311233211-2200301200221003-0000110213122330-0302212202302221-1101312120312330-2230321020131013-1021211013301323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `device_list` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- device_list

<a id="canonical-2011300210203112-1212221211332200-3213022223032230-0000021313221110-2102020102331131-0130001302313130-3211210123310210-1201013002123013"></a>

Type: `"single"`. Computed.

Add device for all interfaces belonging to this fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3012202021311323-0212021031032211-2213131322113222-0132020313103013-2320121011020113-0202202133221133-1331023332120211-0231122112323030"></a>

### Direct properties for `device_list`

- [devices](data-sources--fleet--reference--group-002.md#canonical-1133210223332000-2033333100133222-2032211012031033-1121222303321200-2020320210223013-3113313013313321-1123023331003311-0012002233233132): complete subsection reference.

<a id="canonical-1133210223332000-2033333100133222-2032211012031033-1121222303321200-2020320210223013-3113313013313321-1123023331003311-0012002233233132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `device_list.devices` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [device_list](data-sources--fleet--reference--group-002.md#canonical-0322000310332200-0133003311233211-2200301200221003-0000110213122330-0302212202302221-1101312120312330-2230321020131013-1021211013301323)
- device_list.devices

<a id="canonical-2202032110211321-2300103112313321-1023023022323021-0301213113021332-0022001330301103-2100130231330300-3312322202100211-0020321311112312"></a>

Type: `"list"`. Computed.

Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras,
scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet
and has an corresponding interface/device. The mapping from device configured in fleet with
interface/device in VER node depends on the type of device and is documented in device instance
specific sections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2122322103110310-0011101320320002-2320001121301321-0132322331221303-1212212130000112-1222201320321111-3131100311010010-0033033111223021"></a>

### Direct properties for `device_list.devices`

<a id="canonical-3200123031113200-1001001132222210-0211213012310112-0212012312130010-2311023332120120-0303103022313223-0231121021232122-1331313332020010"></a>

#### `device_list.devices.name` property

Type: `"string"`. Computed.

Name of the device including the unit number (e.g. Eth0 or disk1). The name must match name of
device in host-OS of node.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

- [network_device](data-sources--fleet--reference--group-002.md#canonical-2131311101003331-0110112131012232-1121310012222213-1000103201021121-0000131232300330-1023322300032223-3331100203312202-0321013001002303): complete subsection reference.

<a id="canonical-1323233211030332-3223103001111232-0110100113331220-2213013202023302-1032330232322113-1123332010132313-1312333202021213-1213021001330020"></a>

<a id="canonical-2112200223222312-2233013233103032-3332301131301202-2220203330302121-3031221320102301-2021220031230121-0320013303001001-2112123021023333"></a>

#### `device_list.devices.owner` property

Type: `"string"`. Computed.

\[Enum:
DEVICE\_OWNER\_INVALID|DEVICE\_OWNER\_VER|DEVICE\_OWNER\_VK8S\_WORK\_LOAD|DEVICE\_OWNER\_HOST\]
Defines ownership for a device. Device owner is invalid Device is owned by VER pod. Usually it will
be network interface device or accelerator like crypto engine. Possible values are
\`DEVICE\_OWNER\_INVALID\`, \`DEVICE\_OWNER\_VER\`, \`DEVICE\_OWNER\_VK8S\_WORK\_LOAD\`,
\`DEVICE\_OWNER\_HOST\`. Defaults to \`DEVICE\_OWNER\_INVALID\`.

Additional upstream details:

Defines ownership for a device. Device is available to be owned by vK8s workload on the site, like
camera GPU etc. Device is not available to be owned by vK8s or VER. Can be exposed via some other
service. Like TPM.

Receipt-pinned upstream constraints:

```json
{
  "default": "DEVICE_OWNER_INVALID",
  "enum": [
    "DEVICE_OWNER_INVALID",
    "DEVICE_OWNER_VER",
    "DEVICE_OWNER_VK8S_WORK_LOAD",
    "DEVICE_OWNER_HOST"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2131311101003331-0110112131012232-1121310012222213-1000103201021121-0000131232300330-1023322300032223-3331100203312202-0321013001002303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `device_list.devices.network_device` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [device_list](data-sources--fleet--reference--group-002.md#canonical-0322000310332200-0133003311233211-2200301200221003-0000110213122330-0302212202302221-1101312120312330-2230321020131013-1021211013301323)
- [device_list.devices](data-sources--fleet--reference--group-002.md#canonical-1133210223332000-2033333100133222-2032211012031033-1121222303321200-2020320210223013-3113313013313321-1123023331003311-0012002233233132)
- device_list.devices.network_device

<a id="canonical-0320300202003032-0302232030302230-3203121322110123-3122030332100130-1222123031130121-0013121130321213-2102120313103002-1102203310210112"></a>

Type: `"single"`. Computed.

Represents physical network interface. The 'interface' reference points to a Network Interface
object. Attributes such as Labels, MTU from Network Interface must be applied to the device.

Device mapping to nodes

A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An
interface in node inherits configuration from a device by matching, &#8203;- device\_name in Network
Interface for the device &#8203;- device name for physical-interface in the node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0120321332113202-2233113202032021-2213011020201123-0133301221223312-2301221221032113-3230300311122202-1130001123323012-0211030331310232"></a>

### Direct properties for `device_list.devices.network_device`

- [interface](data-sources--fleet--reference--group-002.md#canonical-1312312011333320-1223003020211321-2312013101231302-1313020202003122-1310102323123131-3130331232001013-0321213321002211-2033031203300131): complete subsection reference.

<a id="canonical-0030022301003213-2210012012300032-0012010111313200-3113330333311301-3321230222312121-2121122130132112-1011330210133233-3120100233121313"></a>

<a id="canonical-0230122103321113-0222000220021132-3103130032123132-1223120033102123-0101010002123130-3312132233321330-0321321210132220-0230123322122323"></a>

#### `device_list.devices.network_device.use` property

Type: `"string"`. Computed.

\[Enum:
NETWORK\_INTERFACE\_USE\_REGULAR|NETWORK\_INTERFACE\_USE\_OUTSIDE|NETWORK\_INTERFACE\_USE\_INSIDE\]
Defines how the device is used If networking device is owned by VER, it is available for users to
configure as required If networking device is owned by VER, it is included in bootstrap config and
member of outside network. If networking device is owned by VER, it is included in bootstrap
config.. Possible values are \`NETWORK\_INTERFACE\_USE\_REGULAR\`,
\`NETWORK\_INTERFACE\_USE\_OUTSIDE\`, \`NETWORK\_INTERFACE\_USE\_INSIDE\`. Defaults to
\`NETWORK\_INTERFACE\_USE\_REGULAR\`.

Additional upstream details:

Defines how the device is used

If networking device is owned by VER, it is available for users to configure as required If
networking device is owned by VER, it is included in bootstrap config and member of outside network.
If networking device is owned by VER, it is included in bootstrap config and member of inside
network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NETWORK_INTERFACE_USE_REGULAR",
  "enum": [
    "NETWORK_INTERFACE_USE_REGULAR",
    "NETWORK_INTERFACE_USE_OUTSIDE",
    "NETWORK_INTERFACE_USE_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1312312011333320-1223003020211321-2312013101231302-1313020202003122-1310102323123131-3130331232001013-0321213321002211-2033031203300131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `device_list.devices.network_device.interface` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [device_list](data-sources--fleet--reference--group-002.md#canonical-0322000310332200-0133003311233211-2200301200221003-0000110213122330-0302212202302221-1101312120312330-2230321020131013-1021211013301323)
- [device_list.devices](data-sources--fleet--reference--group-002.md#canonical-1133210223332000-2033333100133222-2032211012031033-1121222303321200-2020320210223013-3113313013313321-1123023331003311-0012002233233132)
- [device_list.devices.network_device](data-sources--fleet--reference--group-002.md#canonical-2131311101003331-0110112131012232-1121310012222213-1000103201021121-0000131232300330-1023322300032223-3331100203312202-0321013001002303)
- device_list.devices.network_device.interface

<a id="canonical-2112111333200221-1223010122303313-3013000020032101-2023030310011002-0231012111120132-0100001010232313-1130022103330030-0002203321303121"></a>

Type: `"list"`. Computed.

Network Interface attributes for the device. User network interface configuration for this network
device. Attributes like labels, MTU from the 'interface' are applied to corresponding interface in
VER node If network interface refers to a virtual-network, the virtual-netowrk type must be
consistent with use attribute given below If use is NETWORK\_INTERFACE\_USE\_REGULAR, the
virtual-network must be of type VIRTUAL\_NETWORK\_SITE\_LOCAL or
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE if use is NETWORK\_INTERFACE\_USE\_OUTSIDE, the
virtual-network must of type VIRTUAL\_NETWORK\_SITE\_LOCAL if use is
NETWORK\_INTERFACE\_USE\_INSIDE, the virtual-network must of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2123020120131212-3110121002333330-2232202031021132-2231310322110210-1013200330230132-1032003111223023-2003011110323331-1020322211122311"></a>

### Direct properties for `device_list.devices.network_device.interface`

<a id="canonical-1303321303201323-2333033320230121-1211220330021113-3023101221212013-0201022033202120-2011033333212200-3121220230132011-2032132132030230"></a>

#### `device_list.devices.network_device.interface.kind` property

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

<a id="canonical-2303302201213301-3121130133112232-1322311333103233-2011303323001203-2112201313102110-1302020301001021-0220021101110131-2123221223023121"></a>

<a id="canonical-0203031320230222-3023303300103230-0320310001002313-1231320002203323-0000032102321210-3232301232111033-3312022230112231-0020211213321113"></a>

#### `device_list.devices.network_device.interface.name` property

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

<a id="canonical-1000202303113302-3223233313101003-2303320020211113-0310322322000130-3332222310331303-2220320200300112-2202003102213332-0312300303001003"></a>

<a id="canonical-0110313303013100-1111311322120212-3310213100012103-0133030310222020-1211302213013113-2122201311203111-0220332333023022-1202122101323013"></a>

#### `device_list.devices.network_device.interface.namespace` property

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

<a id="canonical-0310133222223303-2113312030313113-0333011003031031-1331331131312313-0332112123023002-0333122231210112-2121331002310322-1300112301202001"></a>

<a id="canonical-0221110333100010-2032131020202313-2000310232000001-2111221131333000-2121021023132323-3122211231021312-3002001323122020-2123010322021031"></a>

#### `device_list.devices.network_device.interface.tenant` property

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

<a id="canonical-3130200202230302-2310030222231212-3010331322120120-1330111022303231-0302300012113112-0032220230211112-1223130000330032-1033001311111200"></a>

<a id="canonical-1232023331120021-0100132301300330-3021121233030312-3302111321122031-0001101333123030-1211302230302112-1101231331022301-3321213031110303"></a>

#### `device_list.devices.network_device.interface.uid` property

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

<a id="canonical-3130311131301002-1132033103331133-1202332201233213-1132213322011332-1011320123020311-1201320333323002-1310223001230020-3202103331012201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_gpu` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- disable_gpu

<a id="canonical-0321031311123031-0000123131131113-0112033002330000-3030303003030323-0131123113120032-0201023322312023-2320333330101233-1031103000101001"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_gpu, enable\_gpu, enable\_vgpu; Default: disable\_gpu\] Configuration parameter
for disable GPU.

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

- [disable_gpu](data-sources--fleet--reference--group-002.md#canonical-0321031311123031-0000123131131113-0112033002330000-3030303003030323-0131123113120032-0201023322312023-2320333330101233-1031103000101001)
- [enable_gpu](data-sources--fleet--reference--group-002.md#canonical-1123320121201021-1123212000311132-2002023122101311-1000030320110023-2002330003123200-3000013212011110-2032112333331201-0333011323201321)
- [enable_vgpu](data-sources--fleet--reference--group-002.md#canonical-2300303031121201-1301011321012331-1113323000210020-2311010232021132-3020133331232123-0001321233221331-1023303300013233-3111322220301110)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010223011012221-3311031211012022-3113130032121323-1100233121200201-0022032322131333-2221122020212113-3102311212332322-0230212232301232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_log_anonymization` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- disable_log_anonymization

<a id="canonical-2101033032331320-3200303222032210-3030311223222310-2131323332233010-0212331102200023-1133032132322221-0222012132120323-3213021210122012"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_log\_anonymization, enable\_log\_anonymization; Default:
disable\_log\_anonymization\] Configuration parameter for disable log anonymization.

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

- [disable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-2101033032331320-3200303222032210-3030311223222310-2131323332233010-0212331102200023-1133032132322221-0222012132120323-3213021210122012)
- [enable_log_anonymization](data-sources--fleet--reference--group-002.md#canonical-1321321032230021-2202033011023221-0301122312020021-0212301302110210-0312012333032201-3112020103331300-3103200131203232-3322213303200103)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330100011232133-1230010220110222-3001003032121133-1003303230211323-1100022022303001-2010203310213123-3103022122231203-2212333211200002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_vm` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- disable_vm

<a id="canonical-2012021110001110-0133232210322213-1323223201321211-1113310113003003-1110012031013001-0213012103310230-3003010202013132-0220222112202110"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_vm, enable\_vm; Default: disable\_vm\] Enable this option

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

- [disable_vm](data-sources--fleet--reference--group-002.md#canonical-2012021110001110-0133232210322213-1323223201321211-1113310113003003-1110012031013001-0213012103310230-3003010202013132-0220222112202110)
- [enable_vm](data-sources--fleet--reference--group-002.md#canonical-2101023113213303-3312233120303332-3230223210222010-2023322002211301-1021123102120221-2121123331321201-0202011100031023-1123001012210033)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310032200331023-2331221302202203-1333300101031103-0203100200300220-1121123020003110-1003110222120021-1121322210111332-3303112011011320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_gpu` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- enable_gpu

<a id="canonical-1123320121201021-1123212000311132-2002023122101311-1000030320110023-2002330003123200-3000013212011110-2032112333331201-0333011323201321"></a>

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

<a id="canonical-3203031033330322-1321112310121322-1323220113000022-0203131210113200-3313302013102112-3331202022212031-1303313212001010-2303130111320221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_log_anonymization` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- enable_log_anonymization

<a id="canonical-1321321032230021-2202033011023221-0301122312020021-0212301302110210-0312012333032201-3112020103331300-3103200131203232-3322213303200103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable log anonymization.

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

<a id="canonical-1031101321301121-1232330032312323-3013110201310103-1201113000102302-0300000320123233-2021130001313132-3302020232203232-0020323013122323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_vgpu` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- enable_vgpu

<a id="canonical-2300303031121201-1301011321012331-1113323000210020-2311010232021132-3020133331232123-0001321233221331-1023303300013233-3111322220301110"></a>

Type: `"single"`. Computed.

Licensing configuration for NVIDIA vGPU.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0301212311112123-2201133001132133-2220302331333232-2312333102223013-2323320023223111-0100130003022323-1101011123301213-1202201113101212"></a>

### Direct properties for `enable_vgpu`

<a id="canonical-0231112300102121-0311233222122022-1223300130111333-2132213131203203-3203201222032132-2113120013013333-2313332110100213-1112331322321231"></a>

#### `enable_vgpu.feature_type` property

Type: `"string"`. Computed.

\[Enum: UNLICENSED|VGPU|VWS|VCS\] Set feature to be enabled Operate with a degraded vGPU performance
Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation Enable NVIDIA Virtual Compute Server.
Possible values are \`UNLICENSED\`, \`VGPU\`, \`VWS\`, \`VCS\`. Defaults to \`UNLICENSED\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNLICENSED",
  "enum": [
    "UNLICENSED",
    "VGPU",
    "VWS",
    "VCS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3002211211002022-1033300022020022-0313030130120011-2201100332233330-0333330100130222-2120201113002313-1031321231322212-3102213301223111"></a>

<a id="canonical-0112012230202020-1021311122103013-2222301012000121-0210001333332332-0232331003211312-1110300132210020-0302233133203113-0031202133122202"></a>

#### `enable_vgpu.server_address` property

Type: `"string"`. Computed.

License Server Address. Set License Server Address.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  }
}
```

<a id="canonical-0210323130301000-1233222212021120-1300101102030012-0020221311020213-2203332223013120-3021300121330112-3113210022102001-0030021331003331"></a>

<a id="canonical-2231313011101301-3132010202231311-1231010221303000-2321231033102101-2021103232202013-0123031113132212-3120231333310212-2110312130223111"></a>

#### `enable_vgpu.server_port` property

Type: `"number"`. Computed.

License Server Port Number. Set License Server port number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1113110323023002-3202020020033032-2102002213011000-1322112033032223-2201001310100103-1213132303120023-2331011300313031-1330120012301000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_vm` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- enable_vm

<a id="canonical-2101023113213303-3312233120303332-3230223210222010-2023322002211301-1021123102120221-2121123331321201-0202011100031023-1123001012210033"></a>

Type: `["object", {}]`. Computed.

VM Configuration. VMs support configuration.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3110033202301221-0230011123210012-2230320123132312-3130002012323300-0300203022111202-3130003001220313-3021313002211102-2012002112213210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `inside_virtual_network` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- inside_virtual_network

<a id="canonical-0301321331200333-2232120232101122-0033233022103113-1330120121333023-2132303111331303-1320001332021323-3102030220223203-1333303103001223"></a>

Type: `"list"`. Computed.

Default inside (site local) virtual network for the fleet.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1030100103230123-0030311310332302-0103103320020310-2120201311010122-3230331203031220-0112020001200013-2002223131322332-1133300133031212"></a>

### Direct properties for `inside_virtual_network`

<a id="canonical-3210320300100322-2120300210100322-2300331211122332-3101300112302323-0010220022133000-1312022322002231-1121000330312312-3223102003221103"></a>

#### `inside_virtual_network.kind` property

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

<a id="canonical-0100211013233113-1120313312223202-3003331131100333-3102231202123001-2033000112332121-2113202002200010-1031200211020323-3230033321202001"></a>

<a id="canonical-2303110321233133-2110331320221303-3021013011230102-2122323022112131-1333003200132200-0201122003102223-0130210232001223-1221103112001112"></a>

#### `inside_virtual_network.name` property

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

<a id="canonical-2310313213303020-0111032110001131-3111232323102123-2212301130130103-1201313223303113-2100133003131133-3231223312210232-1012112313033210"></a>

<a id="canonical-0111210122000212-2330320312110311-2123132320101233-0301232311112202-1013212102103013-2303110213213321-1231000011220303-2230020233110301"></a>

#### `inside_virtual_network.namespace` property

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

<a id="canonical-1002031031213131-1002131221033011-1100313133320323-3231312303203300-0202312120113201-1213101011103300-2001000220303121-0300023133123020"></a>

<a id="canonical-2022213133313302-1000211032331123-3232332010020120-0332311213313201-2203010230202032-3031333111123333-3100231102112330-0330213001112123"></a>

#### `inside_virtual_network.tenant` property

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

<a id="canonical-3213233233302220-3131003223333020-2300100231100332-0200312212021030-3122112331032011-1010233302123310-3331331220301312-3223031322103202"></a>

<a id="canonical-0200311100212232-0022323220201233-1002020001013020-1130331300030030-3311000100210131-2113113001111322-3001010203203132-0222321212332102"></a>

#### `inside_virtual_network.uid` property

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

<a id="canonical-2331132030231322-2323132333102123-2223312302310133-2010010233221210-0131002013222232-2213320220213201-2023231030021312-1122032020022121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `interface_list` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- interface_list

<a id="canonical-1121100011030011-0112201021211330-0232123013200213-0013132120303311-0211002302121000-0323003013300122-0112213121000233-1113331010200030"></a>

Type: `"single"`. Computed.

Add all interfaces belonging to this fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0032131202210200-2320322202102223-3132203113010203-0001311131303110-3200232300013212-0310303320023301-1233300301312333-2201033310311033"></a>

### Direct properties for `interface_list`

- [interfaces](data-sources--fleet--reference--group-002.md#canonical-1102130202010121-1133013312301120-1020100311013010-3303033030002301-0223003020301100-2331122023032010-1213111320302132-0310023231112130): complete subsection reference.

<a id="canonical-1102130202010121-1133013312301120-1020100311013010-3303033030002301-0223003020301100-2331122023032010-1213111320302132-0310023231112130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `interface_list.interfaces` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [interface_list](data-sources--fleet--reference--group-002.md#canonical-2331132030231322-2323132333102123-2223312302310133-2010010233221210-0131002013222232-2213320220213201-2023231030021312-1122032020022121)
- interface_list.interfaces

<a id="canonical-0233013330202230-2103203221230201-0120300120203130-0222113132330103-1123010103030010-2032032002331020-1212112301010123-0023101023102203"></a>

Type: `"list"`. Computed.

Add all interfaces belonging to this fleet.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2113320210320230-3101030220222202-2203113301221002-0100003200212112-3022003030331030-3213302303111203-3302111030232332-3000320023313111"></a>

### Direct properties for `interface_list.interfaces`

<a id="canonical-2000303113030331-3223221011030333-2230011013131211-2032121023202333-1010003303031311-0202132112031013-1302110001031021-2232312220223203"></a>

#### `interface_list.interfaces.name` property

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

<a id="canonical-0103020003230333-0223332322130110-2131322032112230-3120221313332333-3331323013121001-2023011303210102-2021332031310021-3331210223131120"></a>

<a id="canonical-0211303220102113-0102010300112130-2130122320212012-3010202322131212-3200000010012000-3220212321133313-0301301323222102-1130232212313303"></a>

#### `interface_list.interfaces.namespace` property

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

<a id="canonical-2301320300313203-0113301131323200-0002100010131113-1301100022031323-0222013311113202-2112110112313231-1101032222202331-3303312330010330"></a>

<a id="canonical-1232231220132321-1113202001203033-2322300031303131-2032000101233132-2200200230111010-2100323222003132-1031230322003311-0001301223012210"></a>

#### `interface_list.interfaces.tenant` property

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

<a id="canonical-2102203213212232-2121010023212310-1000010030032213-0112022300111120-1031323232020023-0321132223313110-0333211122013003-3212333331332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- kubernetes_upgrade_drain

<a id="canonical-0033222103311202-1301311123132020-2120212211121213-0103232230223330-3021301001000230-1233213322031220-1031203003011122-3203230212301313"></a>

Type: `"single"`. Computed.

Specify how worker nodes within a site will be upgraded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

<a id="canonical-2323020220130230-1021333002330301-3011012213323302-3133133233101033-3222032113110111-2323320113210111-0021203332231100-1110002210303113"></a>

### Direct properties for `kubernetes_upgrade_drain`

- [disable_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-1323110113130032-0020020232201033-1303133103030111-1030103202303012-1022102103331011-2223011222332032-2111301313110333-3010002101322321): complete subsection reference.

- [enable_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-2020202322122121-0301213322112023-1200231220013230-1213110032212022-0020110313021203-0301321002011320-0213122310211123-3312311322330122): complete subsection reference.

<a id="canonical-1323110113130032-0020020232201033-1303133103030111-1030103202303012-1022102103331011-2223011222332032-2111301313110333-3010002101322321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.disable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [kubernetes_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-2102203213212232-2121010023212310-1000010030032213-0112022300111120-1031323232020023-0321132223313110-0333211122013003-3212333331332202)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-3312032011301300-0230332031201323-2320000112331313-3130302010202030-2010101333230101-2300010300130223-0110011123120031-0031100030211230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable upgrade drain.

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

<a id="canonical-2020202322122121-0301213322112023-1200231220013230-1213110032212022-0020110313021203-0301321002011320-0213122310211123-3312311322330122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [kubernetes_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-2102203213212232-2121010023212310-1000010030032213-0112022300111120-1031323232020023-0321132223313110-0333211122013003-3212333331332202)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-2211203311302020-0000022320122023-0330103110112000-0100122032122002-3100111212112231-3210301321320311-0320202333021122-2210231333322332"></a>

Type: `"single"`. Computed.

Specify batch upgrade settings for worker nodes within a site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

<a id="canonical-1301110020323220-2313331203203012-2321000110333321-0013202232031200-3132131013000120-1333100230220000-3132331023233002-3022212033013001"></a>

### Direct properties for `kubernetes_upgrade_drain.enable_upgrade_drain`

- [disable_vega_upgrade_mode](data-sources--fleet--reference--group-002.md#canonical-2031120212133131-1003322331000313-3212232023220022-1310221031321202-1022023000313320-3101022122032203-1221303222232233-1320131031230322): complete subsection reference.

<a id="canonical-0312010020122010-0010003100311002-3201232121200311-3303210130131210-3132012233300011-2003202133301232-1232010120330223-1130103330032332"></a>

<a id="canonical-0233203000311303-2101122302023222-3201223120331220-0211300003013120-3132321133113123-0033301222312033-2033222300120232-2233111211323220"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` property

Type: `"number"`. Computed.

Node Batch Size Count. Exclusive with \[\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-1133010322313123-3330111012330232-3033212213102002-2002302113202110-2112032311212130-2023110033102013-0010000020110013-3322300331301311"></a>

<a id="canonical-3133230103300202-2312221220331102-3303022133300323-0131022320031002-2130232213222232-0313331110333122-1220311001022202-2311302013032020"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` property

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-1213301301020331-3222231122211030-0122111132232101-2030103123233033-2130303021133333-2101002210220020-3022000012111023-1303321201330130"></a>

<a id="canonical-2333220203122232-0131333221331120-1100131132110223-1233302323012202-1323231220101312-3131020200003200-2232302311201323-0122320101121100"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` property

Type: `"number"`. Computed.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](data-sources--fleet--reference--group-002.md#canonical-1003100330021331-2120313030331331-0101220331302323-3000113021030003-1200220322032332-0030120003032320-1222213113121022-2223200210333210): complete subsection reference.

<a id="canonical-2031120212133131-1003322331000313-3212232023220022-1310221031321202-1022023000313320-3101022122032203-1221303222232233-1320131031230322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [kubernetes_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-2102203213212232-2121010023212310-1000010030032213-0112022300111120-1031323232020023-0321132223313110-0333211122013003-3212333331332202)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-2020202322122121-0301213322112023-1200231220013230-1213110032212022-0020110313021203-0301321002011320-0213122310211123-3312311322330122)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-2200201312002301-1033333102111213-0122100123013331-0032203211132232-1301021021323112-3022020031312123-3210020010012313-2011001012102201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable vega upgrade mode.

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

<a id="canonical-1003100330021331-2120313030331331-0101220331302323-3000113021030003-1200220322032332-0030120003032320-1222213113121022-2223200210333210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [kubernetes_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-2102203213212232-2121010023212310-1000010030032213-0112022300111120-1031323232020023-0321132223313110-0333211122013003-3212333331332202)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--fleet--reference--group-002.md#canonical-2020202322122121-0301213322112023-1200231220013230-1213110032212022-0020110313021203-0301321002011320-0213122310211123-3312311322330122)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-0200103012003232-3132222012002223-3010003302331002-1032132233120232-1302331011331100-2231132030003130-1323133202000022-1233002201110033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable vega upgrade mode.

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

<a id="canonical-1201301010102220-2100130021022310-3332321023231220-2333030020220310-0033013130102022-0331111133203220-2213203023002022-0313201111201132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- log_receiver

<a id="canonical-3312133232332200-3131012330233221-3302321110021211-1233220200313031-3223312202102213-1303320202103233-0322132222132031-2331020112313010"></a>

Type: `"single"`. Computed.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

Additional upstream details:

This type establishes a direct reference from one object(the referrer) to another(the referred).

Receipt-pinned upstream constraints:

```json
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

- [log_receiver](data-sources--fleet--reference--group-002.md#canonical-3312133232332200-3131012330233221-3302321110021211-1233220200313031-3223312202102213-1303320202103233-0322132222132031-2331020112313010)
- [logs_streaming_disabled](data-sources--fleet--reference--group-002.md#canonical-0110203111011030-1320123113232221-3030201100112202-2203312312332211-2322203331001002-3133002000011332-2313302313333121-0303133230302231)

Select alternatives according to the provider validators above.

<a id="canonical-3002001310323302-2133020303230001-2233000003322322-2210211320001200-3330111203033330-0232011030011302-0023211032121321-3023002103121011"></a>

### Direct properties for `log_receiver`

<a id="canonical-1032222211330222-3011012102002122-0220023303120323-0310133301230113-2130223032113330-1002113130312112-2132223003313000-0321313032301300"></a>

#### `log_receiver.name` property

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

<a id="canonical-0113331100013321-0013133013311333-0001120223133120-1231032230323033-1100101322202023-0132010200332332-0120121311032111-3333310110130101"></a>

<a id="canonical-1012032033003331-1013003123121312-0233022200313120-0103233000033012-3132232023300022-1232232213232131-0101321220013100-1322202023201332"></a>

#### `log_receiver.namespace` property

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

<a id="canonical-0133310222310121-0313010303230000-2000322302313310-0302333002232031-1300103031021021-3210100031130213-1213233220102020-0223102303113101"></a>

<a id="canonical-0212203100120330-2021122030121122-2013021022210203-0332100030100023-2202010310231213-1231313112100122-1332010033103032-1020010010131232"></a>

#### `log_receiver.tenant` property

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

<a id="canonical-2021310223321313-2033131211212111-0113103022232333-2010100020021011-3113131010131303-2321233323330001-0200120122003231-3232132212231033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `logs_streaming_disabled` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- logs_streaming_disabled

<a id="canonical-0110203111011030-1320123113232221-3030201100112202-2203312312332211-2322203331001002-3133002000011332-2313302313333121-0303133230302231"></a>

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

<a id="canonical-3123211022312221-3200100111011110-1201100232001010-1230100331220302-0013301221331213-2220233203231132-0121230201000122-1132010113102123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_connectors` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- network_connectors

<a id="canonical-3203020222222130-2112202333112110-3013222300132220-3200332010130130-1010201232331313-1223033303322210-2133003223201003-1311200030303102"></a>

Type: `"list"`. Computed.

Network Connector defines connection between two virtual networks in a given site. Fleet defines one
or more such network connectors. The network connectors configuration is applied on all sites that
are member of the fleet.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-0230322211130213-1223331113000012-2202221233111223-3020102301023302-1312203311212113-2020330030020200-1332021311101333-1220121313212223"></a>

### Direct properties for `network_connectors`

<a id="canonical-1000031223220020-2220033010023220-2031011032110322-2302313301122202-1131132030003301-0132220002010203-2112312300211031-2302003211202000"></a>

#### `network_connectors.kind` property

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

<a id="canonical-1010221313033123-0321133003001020-0020330131213311-0212010131101312-1202131230203023-0023131223123212-3120032320321033-0333100332113213"></a>

<a id="canonical-2113301000330332-2322333222012201-3101010202322312-3100011330221100-0203223033132212-1012212021203023-1112011201112203-3302033303330120"></a>

#### `network_connectors.name` property

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

<a id="canonical-2100312301231223-1201211313112103-2201022131133230-3112230311032222-2332301300232200-3111333200120312-1310000130122011-2011113201113102"></a>

<a id="canonical-3200131032112222-0330110011001031-3112002032220212-1132021021110132-2333300120322100-2211311201322133-3131022220322131-2131122230023212"></a>

#### `network_connectors.namespace` property

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

<a id="canonical-1110201230233201-2101302110023223-2003000102220321-3010031331010002-2022222123232112-2112323201311220-3100313321322100-2032002311210131"></a>

<a id="canonical-2212033221001320-2121323111003223-2312003030011112-1312021122121201-0202033031103122-0023110001020300-3213103200221331-2111222122021131"></a>

#### `network_connectors.tenant` property

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

<a id="canonical-0222101321313132-0003203002010120-1231023011023113-2321320133021003-1110301212002212-3323303330000033-3113131012000311-1000302322133112"></a>

<a id="canonical-1330320131201211-1113321033312211-3230132110010211-1101011302021322-2011120132022320-0333312232230222-0013011021023332-1212302110303302"></a>

#### `network_connectors.uid` property

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

<a id="canonical-3321202220211332-1233100100233320-2030101201021220-1310201200230102-1011003120311131-2130232230230210-1022110013030131-2000202132233333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_firewall` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- network_firewall

<a id="canonical-0323002023133100-3213113122023212-2210101100101232-2330032000112331-0032023223132312-0212032011021100-0302010331303301-3123210322012132"></a>

Type: `"list"`. Computed.

Network Firewall defines firewall to be applied for the virtual networks in the fleet. The network
firewall configuration is applied on all sites that are member of the fleet. Constraints The Network
Firewall is applied on Virtual Networks of type site local network and site local inside network.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3202033333023113-3233000221030032-0021130101133111-2310022003303300-2001322221130110-2130001323011210-0321101300131011-3313101323311210"></a>

### Direct properties for `network_firewall`

<a id="canonical-2132312121120231-2003032200221021-3122303302112322-0223321330030132-2131001131133330-1003103320221020-2110233311222023-1103323023012131"></a>

#### `network_firewall.kind` property

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

<a id="canonical-3101321320100333-2301303130233210-3000310230232100-2202333232210111-3000231233120233-3101322012123131-2032331123000222-1211133123203001"></a>

<a id="canonical-3121203220311133-1300233002120130-1222100232013230-3022032101231130-0222110301222211-2112210033203020-0030020102113322-2000232312100310"></a>

#### `network_firewall.name` property

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

<a id="canonical-0303202021223113-3301222201331203-1102210023002322-0223121300133220-2233013020023231-2313210102032323-2221000312230331-1222102312133033"></a>

<a id="canonical-3013102220232012-1023231030002323-1320210001310303-3011313301122220-0121113003312201-2011220303232201-1222103101010023-0100310232101031"></a>

#### `network_firewall.namespace` property

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

<a id="canonical-1212112012320012-0031021310112321-1220303010200232-0332330331212221-3321121012113121-0313322001110221-3023302322122122-1033101113220003"></a>

<a id="canonical-0030212232010133-1301323313132113-3211122201230100-3321100020230021-3010131102130122-1121022123331023-3301232112130022-2213221012332221"></a>

#### `network_firewall.tenant` property

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

<a id="canonical-2333222113120100-1003003203333221-3100333133000002-2123022033013222-0310330333202131-1131001120210013-1133323132210230-2322100121330002"></a>

<a id="canonical-3331022302212121-0220201021122011-1013033022022303-3023122332202002-0221202120212312-0010023331023213-3122200330202102-2033202030310030"></a>

#### `network_firewall.uid` property

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

<a id="canonical-3231300102332313-2132132030313103-2122021022212031-2022211230213122-2030000022032311-3110302202102103-0002101133113121-0021030121221122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_bond_devices` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- no_bond_devices

<a id="canonical-3023231313233222-2123102110130331-0131021311233301-0232210010022131-2100101001012022-1230233223123212-1231312301132111-2110320121201112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no bond devices.

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

<a id="canonical-0221310033011230-1200213332232320-3131023300031210-3031133012130331-3132110111231213-0210320332210001-3122322022310032-1303302321010011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_dc_cluster_group` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- no_dc_cluster_group

<a id="canonical-1100023233033321-1031202011312123-3220000313212323-3201320320113320-0301111211231330-3010123121312200-2311202132022322-1301230111210222"></a>

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

<a id="canonical-1232113232310321-1222130032123323-3333333031011322-2101300132000023-0210131112003320-0023322310102202-2231301313303232-3221202133012133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_storage_device` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- no_storage_device

<a id="canonical-0212022003202000-2210221312322100-1023231001131111-2220322301211232-3220222003032201-0333231023102212-1121310020200021-1313031112120012"></a>

Type: `["object", {}]`. Computed.

\[OneOf: no\_storage\_device, storage\_device\_list; Default: no\_storage\_device\] Configuration
parameter for no storage device.

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

- [no_storage_device](data-sources--fleet--reference--group-002.md#canonical-0212022003202000-2210221312322100-1023231001131111-2220322301211232-3220222003032201-0333231023102212-1121310020200021-1313031112120012)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-0101313010021002-1132021231301111-1333230303120031-2300003010103003-0000212332000030-2023011303230331-3130120302303021-1310301012210220)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301230210330122-2303332032010312-0200323323023201-0033133101310013-0102213033032231-1113332131230030-1221133003003300-1310300203130100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_storage_interfaces` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- no_storage_interfaces

<a id="canonical-0322332222120213-3002300233031103-1211310310112200-1322030123230220-2013300132123111-3031333113133123-2303321030030013-3323013312300231"></a>

Type: `["object", {}]`. Computed.

\[OneOf: no\_storage\_interfaces, storage\_interface\_list; Default: no\_storage\_interfaces\]
Configuration parameter for no storage interfaces.

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

- [no_storage_interfaces](data-sources--fleet--reference--group-002.md#canonical-0322332222120213-3002300233031103-1211310310112200-1322030123230220-2013300132123111-3031333113133123-2303321030030013-3323013312300231)
- [storage_interface_list](data-sources--fleet--reference--group-003.md#canonical-2121301312201313-2003303021323200-1113203311322120-1012103002103230-2013220110320003-0220331131210223-0321010211132303-2320301002023322)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330102102110003-0233200221011033-0302123300312033-1211200102221301-1102322220013301-0112121232210100-2232012022332013-3230220321322303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_storage_static_routes` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- no_storage_static_routes

<a id="canonical-1111022332203132-2200022301101201-3201120013303101-1330033232011220-1301021003311330-0023231001020021-0133211032131311-3130330322013231"></a>

Type: `["object", {}]`. Computed.

\[OneOf: no\_storage\_static\_routes, storage\_static\_routes; Default:
no\_storage\_static\_routes\] Configuration parameter for no storage static routes.

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

- [no_storage_static_routes](data-sources--fleet--reference--group-002.md#canonical-1111022332203132-2200022301101201-3201120013303101-1330033232011220-1301021003311330-0023231001020021-0133211032131311-3130330322013231)
- [storage_static_routes](data-sources--fleet--reference--group-003.md#canonical-3012331233311111-0123111321131231-3213032311022211-3000231010022111-1001031021200012-0323022200021110-3333332210321120-3330100233233210)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322003033013013-1301210121021213-1130031120131113-3200101322001310-3222130102023200-3232303313121331-2133321132323111-0011301020200212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `outside_virtual_network` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- outside_virtual_network

<a id="canonical-1100332213023013-3301211321203301-2132230213220031-1210123100313003-1020033101303102-3213330320103023-1113011213230332-2002002330323102"></a>

Type: `"list"`. Computed.

Default outside (site local) virtual network for the fleet.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0312123132310313-3021322302003021-1203001303211202-2213312231220202-0310333313313012-3222111110300103-3003223133132012-2103103203301333"></a>

### Direct properties for `outside_virtual_network`

<a id="canonical-1032333030013023-2023301001010302-1130122311112233-2303013200030202-3002002301213323-1131330102202212-3223130311012321-1012111120300013"></a>

#### `outside_virtual_network.kind` property

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

<a id="canonical-2133201200110023-2330213000202333-0001101301233021-3213021033103112-2200012321030010-3112102101013103-0000120033030131-2130213021231222"></a>

<a id="canonical-0333333230102131-0111311221312311-2212020113111201-0002003223101231-3331201023222132-1122233222312223-1031321220300202-2002032311022200"></a>

#### `outside_virtual_network.name` property

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

<a id="canonical-3223323110132233-1101221021131211-2011033303201020-2311221133020300-0311203332300020-2101301211031222-3301302323320323-0331032211000123"></a>

<a id="canonical-3021133112212101-1003203212220212-2032230113332103-3310113300010130-2231200032103122-3320030122121213-1213020022231011-0210220213232210"></a>

#### `outside_virtual_network.namespace` property

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

<a id="canonical-3203332013333021-3020301220131033-2320133132321200-2300020002033210-0112313021111323-0020233202333230-3320320331313330-2102103122123222"></a>

<a id="canonical-3202030101120013-1112302120311231-0001110010131312-0030203322232303-2003211320003223-2100200302332212-0210012311030123-1323102102112133"></a>

#### `outside_virtual_network.tenant` property

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

<a id="canonical-3233120331332313-2113223113332212-2333023301231001-2212023012301000-3122123212033001-1212313332030210-1331023330212133-3330133230223032"></a>

<a id="canonical-0320120030230013-1211210023132302-0222000030033202-0023001311112330-2123031113223302-3232023013003212-2132021021201213-3031100012202332"></a>

#### `outside_virtual_network.uid` property

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

<a id="canonical-3133211021322020-2332213213103301-1020331030301223-1210200321112330-0230022112123011-2002100233113211-0213202312111002-0031301033032211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- performance_enhancement_mode

<a id="canonical-2111111130321311-2300101203011331-3220311012120031-1221113012031202-2312313021300030-2122102313312112-1002310021312323-2201201301223020"></a>

Type: `"single"`. Computed.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

<a id="canonical-0302332320311213-1301000001022110-2123321222220010-3231100220011302-1121221213303112-0111212202130002-1332220300322210-0201302310121132"></a>

### Direct properties for `performance_enhancement_mode`

- [perf_mode_l3_enhanced](data-sources--fleet--reference--group-002.md#canonical-3021110002200101-1213011132122310-0300001332022133-2233232313023203-0101212120331023-0313122023033000-0333103212002110-2032300111001303): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--fleet--reference--group-002.md#canonical-2100321330111130-0122101110332312-1203131320322323-1321032201332323-0312310300023333-3200302022110021-0130113231022310-0103223230031112): complete subsection reference.

<a id="canonical-3021110002200101-1213011132122310-0300001332022133-2233232313023203-0101212120331023-0313122023033000-0333103212002110-2032300111001303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [performance_enhancement_mode](data-sources--fleet--reference--group-002.md#canonical-3133211021322020-2332213213103301-1020331030301223-1210200321112330-0230022112123011-2002100233113211-0213202312111002-0031301033032211)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-1210123332001023-0011032303002121-1033032132012021-1330022010331011-3300130002101100-1122130331103132-2123012133310233-1031320311233200"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l3 enhanced.

Additional upstream details:

L3 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

<a id="canonical-2012112210300103-1321233310202101-3320111122030101-2110233013233222-0103210331201220-1132310230203312-3211130202003231-1011231203331222"></a>

### Direct properties for `performance_enhancement_mode.perf_mode_l3_enhanced`

- [jumbo](data-sources--fleet--reference--group-002.md#canonical-0120102200122013-3333103212201200-1200031023301313-0003111022010123-0111132323331012-1033310002203010-1031123223203111-3210122303013212): complete subsection reference.

- [no_jumbo](data-sources--fleet--reference--group-002.md#canonical-2130103323333220-3332320311023023-3102132032213223-2131301020010131-0313323220213302-1132233332022130-3021213203333123-3231110110223123): complete subsection reference.

<a id="canonical-0120102200122013-3333103212201200-1200031023301313-0003111022010123-0111132323331012-1033310002203010-1031123223203111-3210122303013212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [performance_enhancement_mode](data-sources--fleet--reference--group-002.md#canonical-3133211021322020-2332213213103301-1020331030301223-1210200321112330-0230022112123011-2002100233113211-0213202312111002-0031301033032211)
- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--fleet--reference--group-002.md#canonical-3021110002200101-1213011132122310-0300001332022133-2233232313023203-0101212120331023-0313122023033000-0333103212002110-2032300111001303)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-2130031300102213-2233313012101201-1213102112212103-2201122033300030-1033223103010320-2222110003021331-3301033002103112-0133122013310030"></a>

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

<a id="canonical-2130103323333220-3332320311023023-3102132032213223-2131301020010131-0313323220213302-1132233332022130-3021213203333123-3231110110223123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [performance_enhancement_mode](data-sources--fleet--reference--group-002.md#canonical-3133211021322020-2332213213103301-1020331030301223-1210200321112330-0230022112123011-2002100233113211-0213202312111002-0031301033032211)
- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--fleet--reference--group-002.md#canonical-3021110002200101-1213011132122310-0300001332022133-2233232313023203-0101212120331023-0313122023033000-0333103212002110-2032300111001303)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-3303003031332303-1112210012311132-0113313322002101-0101123232303302-0101203110332222-1320230323220320-0112002132223230-1003322323303101"></a>

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

<a id="canonical-2100321330111130-0122101110332312-1203131320322323-1321032201332323-0312310300023333-3200302022110021-0130113231022310-0103223230031112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [performance_enhancement_mode](data-sources--fleet--reference--group-002.md#canonical-3133211021322020-2332213213103301-1020331030301223-1210200321112330-0230022112123011-2002100233113211-0213202312111002-0031301033032211)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-0031303022113010-0022213303333213-0311232230023023-2233021113131110-1330032322002131-1133102221212133-1310230212013221-1303010131120213"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l7 enhanced.

Additional upstream details:

L7 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

<a id="canonical-0103320213103310-2012102001021132-3200131310302030-3021202111132000-0213111330213213-3332200131313310-1311100200101233-2101213031023110"></a>

### Direct properties for `performance_enhancement_mode.perf_mode_l7_enhanced`

- [jumbo_disabled](data-sources--fleet--reference--group-002.md#canonical-2331113020033010-3211213002011303-0210110012000333-3301121123221220-2122233303232000-1333131131333223-2211311010310000-3202203203112230): complete subsection reference.

- [jumbo_enabled](data-sources--fleet--reference--group-002.md#canonical-2123322333012003-0011121010223221-0330003103003322-2230231212210222-0123102200301202-0231111132210000-0112303223201321-1212013103033213): complete subsection reference.

<a id="canonical-2331113020033010-3211213002011303-0210110012000333-3301121123221220-2122233303232000-1333131131333223-2211311010310000-3202203203112230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [performance_enhancement_mode](data-sources--fleet--reference--group-002.md#canonical-3133211021322020-2332213213103301-1020331030301223-1210200321112330-0230022112123011-2002100233113211-0213202312111002-0031301033032211)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--fleet--reference--group-002.md#canonical-2100321330111130-0122101110332312-1203131320322323-1321032201332323-0312310300023333-3200302022110021-0130113231022310-0103223230031112)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-0130110200111003-2302002032101102-2113310203202112-1300200331103323-2222202030022323-2033213021013003-3311231013102231-0132332303011202"></a>

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

<a id="canonical-2123322333012003-0011121010223221-0330003103003322-2230231212210222-0123102200301202-0231111132210000-0112303223201321-1212013103033213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [performance_enhancement_mode](data-sources--fleet--reference--group-002.md#canonical-3133211021322020-2332213213103301-1020331030301223-1210200321112330-0230022112123011-2002100233113211-0213202312111002-0031301033032211)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--fleet--reference--group-002.md#canonical-2100321330111130-0122101110332312-1203131320322323-1321032201332323-0312310300023333-3200302022110021-0130113231022310-0103223230031112)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-1230221021223321-1330010222210131-2211103001211121-2331130110302010-3131331200122303-2113033211011232-1022122101230230-3132020211011201"></a>

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

<a id="canonical-1132110120003111-0323103313100123-2113310212200301-0133232320111330-0030220003213233-1113311012021102-1212011223322232-3112201313103010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sriov_interfaces` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- sriov_interfaces

<a id="canonical-2123333323013300-0110133102200320-1232210230303220-3011003312322200-3022100322000212-0311121301321002-2220103203200120-2223000320220210"></a>

Type: `"single"`. Computed.

List of all custom SR-IOV interfaces configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1133112132222313-3321320103313232-2310333001121221-1332032112023213-1020003132222001-1102300022330213-0032100230211300-2322101303323013"></a>

### Direct properties for `sriov_interfaces`

- [sriov_interface](data-sources--fleet--reference--group-002.md#canonical-2223101330132323-2330122331131311-3303222222011201-3323323003030300-0301222013303320-1221121032013123-3311232310000023-0203302130002122): complete subsection reference.

<a id="canonical-2223101330132323-2330122331131311-3303222222011201-3323323003030300-0301222013303320-1221121032013123-3311232310000023-0203302130002122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sriov_interfaces.sriov_interface` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [sriov_interfaces](data-sources--fleet--reference--group-002.md#canonical-1132110120003111-0323103313100123-2113310212200301-0133232320111330-0030220003213233-1113311012021102-1212011223322232-3112201313103010)
- sriov_interfaces.sriov_interface

<a id="canonical-0112222100133120-2030101303200012-2113333303331123-3112021310103000-1112323121222122-1102121021232010-0233313133101031-1121312122100333"></a>

Type: `"list"`. Computed.

Use custom SR-IOV interfaces Configuration.

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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2030030101221330-2212223120233113-0100231200101212-3121132233332332-2310030323320033-2032321321221100-2123332322332210-2121222013121202"></a>

### Direct properties for `sriov_interfaces.sriov_interface`

<a id="canonical-0310221020023223-1121323131312002-3311112323221210-3130133123211021-3112333011011210-0313222012201213-3303001232112130-3221320232313130"></a>

#### `sriov_interfaces.sriov_interface.interface_name` property

Type: `"string"`. Computed.

Name of physical interface. Name of SR-IOV physical interface.

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

<a id="canonical-2111000223300022-3331212201322303-3232313303030113-1302111011210002-3103210130231311-2232323120132302-3120220233300012-2310200223332313"></a>

<a id="canonical-3312303111003330-2230002030202113-3010011331133203-0303220230230112-2033121110123221-1233113333100311-1321032022100333-2130112312303312"></a>

#### `sriov_interfaces.sriov_interface.number_of_vfio_vfs` property

Type: `"number"`. Computed.

Number of virtual functions reserved for VNFs and DPDK-based CNFs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1023102302001031-3003122132310132-2120330033332332-3123122033313330-0300310112333000-2201330202213321-0303311013100022-3312303223020120"></a>

<a id="canonical-2330211320011233-2123211011212222-2303212212220011-3302302222020031-1103221023122201-0112203002222300-2231201033133113-3331031013302121"></a>

#### `sriov_interfaces.sriov_interface.number_of_vfs` property

Type: `"number"`. Computed.

Total number of virtual functions. Total number of virtual functions.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2023110232112222-3321330312213101-2201202331303111-2221322320001110-1303303110032013-3100010023303012-0321310033103202-3032223021313121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- storage_class_list

<a id="canonical-1012302103312022-3331132220002110-1120220121220023-0000311103003122-0111002031331211-1000002000021111-3101030221022312-0130130013123023"></a>

Type: `"single"`. Computed.

Add additional custom storage classes in Kubernetes for this fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0013231331200020-3300303003203010-3220133223120000-3330012111210210-2321322131331133-3303333303333112-3010101311210013-1133200232212233"></a>

### Direct properties for `storage_class_list`

- [storage_classes](data-sources--fleet--reference--group-002.md#canonical-3332030213120331-3133123310112000-0030211332231133-1313131033100011-1002001323111021-1322321023132033-3130112202202101-3033013311323132): complete subsection reference.

<a id="canonical-3332030213120331-3133123310112000-0030211332231133-1313131033100011-1002001323111021-1322321023132033-3130112202202101-3033013311323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-2023110232112222-3321330312213101-2201202331303111-2221322320001110-1303303110032013-3100010023303012-0321310033103202-3032223021313121)
- storage_class_list.storage_classes

<a id="canonical-1123301131200103-0012020101121230-0233123223010203-3011313030110331-1032003223301131-3321210232002001-0133010001120330-0012211311302133"></a>

Type: `"list"`. Computed.

List of Storage Classes. List of custom storage classes.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3100202132310312-1302112110201321-1102022120322202-2023211312000300-2333001120213230-1102213321223312-0310300312230312-3010311002100003"></a>

### Direct properties for `storage_class_list.storage_classes`

<a id="canonical-0032120033220023-2223220123322220-3302210001003130-1312220332100020-3301132223111033-1021321113123013-0310321111211010-2220220311231313"></a>

#### `storage_class_list.storage_classes.advanced_storage_parameters` property

Type: `["map", "string"]`. Computed.

Advanced Parameters. Map of parameter name and string value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2131112102201023-3231101321131121-3033021331102121-3131011130322303-3230012330033123-2000231202333303-1231302020001013-2210110331100020"></a>

<a id="canonical-3321021020333131-2221032230221113-2013020201000000-3313233033110111-2212131111032331-3213012230330323-0312123110121332-0110110231221112"></a>

#### `storage_class_list.storage_classes.allow_volume_expansion` property

Type: `"bool"`. Computed.

Allow Volume Expansion. Allow volume expansion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [custom_storage](data-sources--fleet--reference--group-002.md#canonical-2103321033310323-1310131123321333-2300302113222311-2310303100322231-3113121032232312-3231121210000222-0332213021113221-2333213221121311): complete subsection reference.

<a id="canonical-0302113031200312-2220302113220101-3001101120212003-1112323122001300-2132032011202032-2223002201303332-1202022330312011-3132222011032312"></a>

<a id="canonical-1332231113021331-3031303100322321-2320211110110233-2003112033331120-0311233131030302-0010132310322300-1103312232013103-3300301010302100"></a>

#### `storage_class_list.storage_classes.default_storage_class` property

Type: `"bool"`. Computed.

Make this storage class default storage class for the K8s cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3121203321320310-3313020133321013-2100133002223210-1202023211310331-2200013013320132-2030231301021110-0131330331320300-3120323030030220"></a>

<a id="canonical-0002030311231023-0101031231133111-0131013112221313-0122203232322030-3202331032203311-1110030120320301-2233110033123233-2322201202302333"></a>

#### `storage_class_list.storage_classes.description_spec` property

Type: `"string"`. Computed.

Storage Class Description. Description for this storage class.

- [hpe_storage](data-sources--fleet--reference--group-002.md#canonical-3333203230031011-0233113121012323-1123220010320122-2133332333202030-1232013102011013-2321333201202303-1132300001011033-0000212113113222): complete subsection reference.

- [netapp_trident](data-sources--fleet--reference--group-002.md#canonical-2203313120110100-3230202300332001-2221002111100231-0300101000123331-3101321212203303-2333320001101321-0131102322311203-0002311222210022): complete subsection reference.

- [pure_service_orchestrator](data-sources--fleet--reference--group-002.md#canonical-3302302030332001-0200211312303013-3212010111223010-3000023122000101-1130200321102230-3233233220320311-0023000311222222-3320302013313202): complete subsection reference.

<a id="canonical-0332233100132100-3313211102031202-1112201310330102-0022222000222310-0232210221321102-1213002013033322-0131230211001303-0323303222221311"></a>

<a id="canonical-3010011103200101-0133030331003000-2210012310032212-0033110330001312-3301300312122312-0112101011321103-2300110001120210-0132321212002120"></a>

#### `storage_class_list.storage_classes.reclaim_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Reclaim Policy.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="canonical-0113021022011012-2301301010330021-2010030022323311-3203122001132300-3211310030221033-1212102113021233-0302130322323220-0120130330200020"></a>

<a id="canonical-1221001131113130-2011313210313000-1021103311113002-0131202102233200-2013322021013202-0022013210112010-3211332210000213-1112333313201212"></a>

#### `storage_class_list.storage_classes.storage_class_name` property

Type: `"string"`. Computed.

Name of the storage class as it will appear in K8s.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0222132210301222-1001230113113223-2111213320101320-0333330123022132-2200112322132221-0212301120313203-1121120010312233-0332231212202112"></a>

<a id="canonical-3111212223001021-2111202301110203-1311320133301002-0130232232023021-1112131113023221-0223311200323033-1100112303211311-2332020211211012"></a>

#### `storage_class_list.storage_classes.storage_device` property

Type: `"string"`. Computed.

Storage device that this class will use. The Device name defined at previous step.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2103321033310323-1310131123321333-2300302113222311-2310303100322231-3113121032232312-3231121210000222-0332213021113221-2333213221121311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes.custom_storage` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-2023110232112222-3321330312213101-2201202331303111-2221322320001110-1303303110032013-3100010023303012-0321310033103202-3032223021313121)
- [storage_class_list.storage_classes](data-sources--fleet--reference--group-002.md#canonical-3332030213120331-3133123310112000-0030211332231133-1313131033100011-1002001323111021-1322321023132033-3130112202202101-3033013311323132)
- storage_class_list.storage_classes.custom_storage

<a id="canonical-1011032211012230-3003330123101032-2221332003020331-1332102020103102-0212032031133333-2311211120122330-3302123002003123-2322311321023313"></a>

Type: `"single"`. Computed.

Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into
given site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3200133201001103-3101100301011301-2333030323333012-1101213032131212-0002322131032222-3021103221010111-3310330313310231-0233310231132220"></a>

### Direct properties for `storage_class_list.storage_classes.custom_storage`

<a id="canonical-3223313220020302-1233232332211031-1100101110320331-0112332103110001-0233322200311130-1120002312212310-1333130301013122-3320203321331030"></a>

#### `storage_class_list.storage_classes.custom_storage.yaml` property

Type: `"string"`. Computed.

Storage Class YAML. K8s YAML for StorageClass.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3333203230031011-0233113121012323-1123220010320122-2133332333202030-1232013102011013-2321333201202303-1132300001011033-0000212113113222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes.hpe_storage` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-2023110232112222-3321330312213101-2201202331303111-2221322320001110-1303303110032013-3100010023303012-0321310033103202-3032223021313121)
- [storage_class_list.storage_classes](data-sources--fleet--reference--group-002.md#canonical-3332030213120331-3133123310112000-0030211332231133-1313131033100011-1002001323111021-1322321023132033-3130112202202101-3033013311323132)
- storage_class_list.storage_classes.hpe_storage

<a id="canonical-1033320212320320-1323231013103122-2211011120321203-2102220331113233-2131132020021312-2312101132100223-2013022220203132-2123100130121120"></a>

Type: `"single"`. Computed.

Storage class Device configuration for HPE Storage.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1221212020032012-1231120323210320-0320012122111301-0112123000211233-0213303131033133-3030132231200330-2210221212230310-1311221300331010"></a>

### Direct properties for `storage_class_list.storage_classes.hpe_storage`

<a id="canonical-1022130100122133-0032132302222132-3120210100110132-0302201100331031-2312031130120011-2200210223320320-0101231000322002-0011103211302210"></a>

#### `storage_class_list.storage_classes.hpe_storage.allow_mutations` property

Type: `"string"`. Computed.

Mutation can override specified parameters.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0110300302201122-0130112011330220-2113312010001321-3132311102023123-2121101121222330-2102320031130011-1333231122301321-3012112200330322"></a>

<a id="canonical-2321302302003020-0022001300331331-2202033012003111-1211332011320113-1300322123211303-2222203112220332-1123221013001000-0111333131130121"></a>

#### `storage_class_list.storage_classes.hpe_storage.allow_overrides` property

Type: `"string"`. Computed.

AllowOverrides. PVC can override specified parameters.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1310011022323021-0020011200310200-2012103122030303-0203223113111300-0211123331112302-1133020133200211-3221003122203212-2230122310301223"></a>

<a id="canonical-3020012030202301-3023220111012002-0321320221330111-0021202132003231-3130320113220201-2200312133302013-1313010231310302-3211310332331201"></a>

#### `storage_class_list.storage_classes.hpe_storage.dedupe_enabled` property

Type: `"bool"`. Computed.

Indicates that the volume should enable deduplication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3202300032023110-0203313012000301-3321013002023003-0301003200013130-0333032013233212-2102322130302300-0111210213231131-3132132102103110"></a>

<a id="canonical-2001312331213213-1002311100021223-3201113003111111-1103020022231320-3112310020110021-3112021233221212-3312121131010303-0232100011032012"></a>

#### `storage_class_list.storage_classes.hpe_storage.description_spec` property

Type: `"string"`. Computed.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

<a id="canonical-1230300330313101-1030221330022300-2222113311332131-0332023311223133-0030022310020212-2131202330333121-0101031100112011-3222033311020020"></a>

<a id="canonical-2020212202030031-2200303120130100-0023101220331222-2212332203033023-2113213131001213-3321010200331103-0123020102223230-3102100113111113"></a>

#### `storage_class_list.storage_classes.hpe_storage.destroy_on_delete` property

Type: `"bool"`. Computed.

Indicates the backing Nimble volume (including snapshots) should be destroyed when the PVC is
deleted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0121200001130232-2101032303201123-3033231030320023-2001313112320222-2021103203102033-1332221001230022-1321211110011033-2103330022121311"></a>

<a id="canonical-3313230132133013-0331102232310112-3130333322012301-2310031112023030-3100332020223030-0323301300100322-3211111122001311-0021323021331100"></a>

#### `storage_class_list.storage_classes.hpe_storage.encrypted` property

Type: `"bool"`. Computed.

Indicates that the volume should be encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2000321213313102-0132330302330331-3110112323111121-3203232233011012-0100321113002030-2022232132000230-3231022022303022-0331332223022012"></a>

<a id="canonical-3011220333031300-2213321323222200-0023330021220010-3232131020203321-1123322200003200-2110212133331033-0131212102001130-1222012131310233"></a>

#### `storage_class_list.storage_classes.hpe_storage.folder` property

Type: `"string"`. Computed.

The name of the folder in which to place the volume.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0131130210231232-1011332332223221-0011031101133103-2011031323332322-2133321211020321-0132312110033123-0320213212311213-2033230030322310"></a>

<a id="canonical-0033112332012131-2300330311223102-2230212333022100-1302121012121233-1203330212311332-0010233032000112-1200100210131113-0202120330011323"></a>

#### `storage_class_list.storage_classes.hpe_storage.limit_iops` property

Type: `"string"`. Computed.

LimitIops. The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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

<a id="canonical-3333003020323302-0131031012023331-0113112203003121-3122021031003211-3311132312122113-1103320332023001-3012332300032012-3002103001332330"></a>

<a id="canonical-3021302312022033-3032221031122320-2023330200230322-1333033201232032-2310303320002123-1112300202310330-2031222323201331-1302013230322131"></a>

#### `storage_class_list.storage_classes.hpe_storage.limit_mbps` property

Type: `"string"`. Computed.

LimitMbps. The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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

<a id="canonical-1100202222021302-1303032030010312-1331020310120022-0120322123010200-3213230132201010-2100233210120322-3102122033302233-3322023310233232"></a>

<a id="canonical-2012121311023022-2330222120311333-2111323033302131-2323201321222022-0121232221220112-3100330300220031-0302121310013131-0331012020113010"></a>

#### `storage_class_list.storage_classes.hpe_storage.performance_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

The name of the performance policy to assign to the volume.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2023021101233313-0132221303122120-2300313131333113-1121323112310201-2233012022222120-1123200000032320-0203012221321030-2123223211031112"></a>

<a id="canonical-3210333003012013-3123032302121323-1021120010110221-3200310003002213-1213130331213131-0212233320022001-3020312111310321-2322212233132333"></a>

#### `storage_class_list.storage_classes.hpe_storage.pool` property

Type: `"string"`. Computed.

The name of the pool in which to place the volume.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2232230020221112-3131101330212320-3232312300002022-0131201202311211-3312202201001203-0223232213313330-1131233020101003-3321311211201211"></a>

<a id="canonical-0203120022013130-1302303110010013-0201121000011211-0310131303231010-2110320121023220-3232222103323000-0303213123300020-3233201112021212"></a>

#### `storage_class_list.storage_classes.hpe_storage.protection_template` property

Type: `"string"`. Computed.

The name of the performance policy to assign to the volume.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1202220133122332-0210233223020133-2330120123212203-2012202312202311-1210231013000032-1110003121021013-1123113311220321-3223102330223022"></a>

<a id="canonical-2303013231022330-3103213031301201-2203312022320111-0311323203312310-1133300333320323-2312322332132212-1120303022020220-1100100032113312"></a>

#### `storage_class_list.storage_classes.hpe_storage.secret_name` property

Type: `"string"`. Computed.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1120210030131033-0202303200113323-3202301031003212-0221322202223100-2331202122111331-0323033310123210-3301332323021302-2131203021013000"></a>

<a id="canonical-0220030200021122-0232122311330130-2101021000003310-2033200202312132-0313321010220213-3021200112130322-0213122131213022-0233221310000322"></a>

#### `storage_class_list.storage_classes.hpe_storage.secret_namespace` property

Type: `"string"`. Computed.

The SecretNamespace parameter is used to identify name of namespace where secret resides.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1003133213102210-2000323123130030-0123001101220312-0132230032320012-1333130110133330-1102013323113201-2300103312002130-0111112122221203"></a>

<a id="canonical-2001122323023220-1010101002302113-0313100121021020-1100220003132233-2131103002112320-2132112012201202-3211210130321122-1110110312000303"></a>

#### `storage_class_list.storage_classes.hpe_storage.sync_on_detach` property

Type: `"bool"`. Computed.

Indicates that a snapshot of the volume should be synced to the replication partner each time it is
detached from a node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2320322113303121-0323223311100012-2032012003302203-3202201200313030-1031102211001200-3030022012310010-1110233131212001-1003213332321012"></a>

<a id="canonical-2231111133123311-0231132113232312-0321102320022130-1133320331133311-1303231101332133-1220110201301221-1002200202011212-0220000103120220"></a>

#### `storage_class_list.storage_classes.hpe_storage.thick` property

Type: `"bool"`. Computed.

Indicates that the volume should be thick provisioned.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2203313120110100-3230202300332001-2221002111100231-0300101000123331-3101321212203303-2333320001101321-0131102322311203-0002311222210022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes.netapp_trident` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-2023110232112222-3321330312213101-2201202331303111-2221322320001110-1303303110032013-3100010023303012-0321310033103202-3032223021313121)
- [storage_class_list.storage_classes](data-sources--fleet--reference--group-002.md#canonical-3332030213120331-3133123310112000-0030211332231133-1313131033100011-1002001323111021-1322321023132033-3130112202202101-3033013311323132)
- storage_class_list.storage_classes.netapp_trident

<a id="canonical-1010132203332021-2231322223023212-2021203320331211-0123330220103120-1232011232011221-3233123112011200-1221311123000212-3220301033330112"></a>

Type: `"single"`. Computed.

Storage class Device configuration for NetApp Trident.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1212201010123030-2012221213330003-3330023010003000-0213203132021221-0312221121313102-1211321000323010-0100012310023230-0110312333212001"></a>

### Direct properties for `storage_class_list.storage_classes.netapp_trident`

- [selector](data-sources--fleet--reference--group-002.md#canonical-2220110333231123-0331223110210132-1332010323112022-2230113220023112-0222100120211312-0303011321223312-1232321003030230-2010103200321031): complete subsection reference.

<a id="canonical-0103300203220011-0203210310021323-3230323320322211-1230032233330212-2010331133132313-2130123231330201-1320113010122310-3123212301210220"></a>

<a id="canonical-0130213311100223-1110000332110103-3200032113001313-2120030132121012-2030110312010100-2012122322102220-1101021332010013-0222233031300013"></a>

#### `storage_class_list.storage_classes.netapp_trident.storage_pools` property

Type: `"string"`. Computed.

The storagePools parameter is used to further restrict the set of pools that match any specified
attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-2220110333231123-0331223110210132-1332010323112022-2230113220023112-0222100120211312-0303011321223312-1232321003030230-2010103200321031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes.netapp_trident.selector` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-2023110232112222-3321330312213101-2201202331303111-2221322320001110-1303303110032013-3100010023303012-0321310033103202-3032223021313121)
- [storage_class_list.storage_classes](data-sources--fleet--reference--group-002.md#canonical-3332030213120331-3133123310112000-0030211332231133-1313131033100011-1002001323111021-1322321023132033-3130112202202101-3033013311323132)
- [storage_class_list.storage_classes.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-2203313120110100-3230202300332001-2221002111100231-0300101000123331-3101321212203303-2333320001101321-0131102322311203-0002311222210022)
- storage_class_list.storage_classes.netapp_trident.selector

<a id="canonical-1210032121301031-0233321201120230-1200222111312013-0030030332100322-3212012233222300-1032111130223131-2130223031221031-1011332021232121"></a>

Type: `"single"`. Computed.

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3302302030332001-0200211312303013-3212010111223010-3000023122000101-1130200321102230-3233233220320311-0023000311222222-3320302013313202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_class_list.storage_classes.pure_service_orchestrator` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_class_list](data-sources--fleet--reference--group-002.md#canonical-2023110232112222-3321330312213101-2201202331303111-2221322320001110-1303303110032013-3100010023303012-0321310033103202-3032223021313121)
- [storage_class_list.storage_classes](data-sources--fleet--reference--group-002.md#canonical-3332030213120331-3133123310112000-0030211332231133-1313131033100011-1002001323111021-1322321023132033-3130112202202101-3033013311323132)
- storage_class_list.storage_classes.pure_service_orchestrator

<a id="canonical-0212332120132020-0223223111102131-0311301300333231-0303221111021331-3311132212301322-3120311012031201-0313132021133101-2102112311313100"></a>

Type: `"single"`. Computed.

Storage class Device configuration for Pure Service Orchestrator.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1021211322013303-3202321000321313-0311123100302312-3001201013221333-3002221103212323-0032100212231213-2312332112333022-0103310333020003"></a>

### Direct properties for `storage_class_list.storage_classes.pure_service_orchestrator`

<a id="canonical-1300022322110132-0311231320020103-2010121223000322-0011232201320102-1023333110001320-1302321311123300-0010322022111133-2201310230331332"></a>

#### `storage_class_list.storage_classes.pure_service_orchestrator.backend` property

Type: `"string"`. Computed.

\[Enum: block|file\] Defines type of Pure storage backend block or file. The volume will have the
aspects defined in the chosen virtual pool. Possible values are \`block\`, \`file\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "block",
    "file"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  }
}
```

<a id="canonical-3300130132110032-0011122313100011-2121123010300322-1320330301101122-0100000210130311-1311200120302133-1030213111110310-2303001222133021"></a>

<a id="canonical-3111302203003201-3002110221133320-2123331212303211-3033110223121003-2101122333101031-3012301312003203-0103213231332323-2121131133332222"></a>

#### `storage_class_list.storage_classes.pure_service_orchestrator.bandwidth_limit` property

Type: `"string"`. Computed.

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 12,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 12,
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
    "ves.io.schema.rules.string.max_len": "12"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "12"
  }
}
```

<a id="canonical-3303230333030000-3310303011032122-3123130102333210-1121202013302231-1000320111123111-0021300033002031-1331031320000300-0002323332233011"></a>

<a id="canonical-0212020101223133-3303032100130233-1110013331110132-3321011212300133-2302022001030213-0020303102320112-2231133221131101-1102010312012302"></a>

#### `storage_class_list.storage_classes.pure_service_orchestrator.iops_limit` property

Type: `"number"`. Computed.

Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not
defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100000000,
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
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  }
}
```

<a id="canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- storage_device_list

<a id="canonical-0101313010021002-1132021231301111-1333230303120031-2300003010103003-0000212332000030-2023011303230331-3130120302303021-1310301012210220"></a>

Type: `"single"`. Computed.

Add additional custom storage classes in Kubernetes for this fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1303122312302311-1222112201310313-1103233002211000-3132233101132001-2211203302301321-2300310212233123-2310302121233331-2121000001313300"></a>

### Direct properties for `storage_device_list`

- [storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211): complete subsection reference.

<a id="canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- storage_device_list.storage_devices

<a id="canonical-2332132030000132-3131331222133020-1110320023312230-0200222110002313-2113121002302103-0330133100123100-2100103222312003-1002113031210330"></a>

Type: `"list"`. Computed.

List of Storage Devices. List of custom storage devices.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2131031231032130-3111232321331001-0210210320300321-2023321030012010-1120310313323200-1211321101120121-1010323103113231-0211030322221121"></a>

### Direct properties for `storage_device_list.storage_devices`

<a id="canonical-0103123113203113-3301021330112121-0221100230210030-1002312300202303-1113233003300321-3213123122130032-0333103000031031-0333103312021010"></a>

#### `storage_device_list.storage_devices.advanced_advanced_parameters` property

Type: `["map", "string"]`. Computed.

Advanced Parameters. Map of parameter name and string value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [custom_storage](data-sources--fleet--reference--group-002.md#canonical-0310213111231111-3222002011210000-0212333331311103-1200300310333020-0103023300323110-2031031010200112-0211222020002331-3130330102011000): complete subsection reference.

- [hpe_storage](data-sources--fleet--reference--group-002.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111): complete subsection reference.

- [netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321): complete subsection reference.

- [pure_service_orchestrator](data-sources--fleet--reference--group-003.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022): complete subsection reference.

<a id="canonical-1133223211122022-1310133110233323-0100211113310010-3130303310132221-2130102103200102-2000122010113011-1030010012210101-2103031201022020"></a>

<a id="canonical-2030112013210320-0021001102210001-2113203023113312-2102300330113210-0230301130202323-3133223222322313-1302032312110132-2210111122013311"></a>

#### `storage_device_list.storage_devices.storage_device` property

Type: `"string"`. Computed.

Storage Device. Storage device and device unit.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0310213111231111-3222002011210000-0212333331311103-1200300310333020-0103023300323110-2031031010200112-0211222020002331-3130330102011000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.custom_storage` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- storage_device_list.storage_devices.custom_storage

<a id="canonical-2033333111100031-0322201230313331-2103013012013113-1333033103312210-1123011002330201-0221030302130102-3120322000132033-3300302230021112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for custom storage.

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

<a id="canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- storage_device_list.storage_devices.hpe_storage

<a id="canonical-1212112201320032-3131303222211210-1233131001130100-1110001310233102-2232031303300033-2321102313222101-0112011000331022-3321320102323221"></a>

Type: `"single"`. Computed.

Configuration parameter for hpe storage.

Additional upstream details:

Device configuration for HPE Storage.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2322132202132110-1010122231212331-1131103033010303-1000221031133030-3110010100020211-0100002232332013-1131120113003202-1030013312200212"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage`

<a id="canonical-2113032022331123-3121132312013300-2100213212300022-1200013130113330-0031302223110322-2230213220201123-0303122230312220-1201101323111213"></a>

#### `storage_device_list.storage_devices.hpe_storage.api_server_port` property

Type: `"number"`. Computed.

Storage server Port. Enter Storage Server Port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [iscsi_chap_password](data-sources--fleet--reference--group-002.md#canonical-2011133210003233-3131330213022333-1301302221301320-0331302020103111-1130013021301311-0223101011022203-2333310310111200-1313002103033111): complete subsection reference.

<a id="canonical-0103200223000220-0331022103311121-3123010030033321-1020223100323331-3232132232112030-3301130310202332-0130200213032122-1200102022230233"></a>

<a id="canonical-2103332220313023-2322132101133111-2221230321113232-1100233231232212-1201311320103013-0020211021201222-2032322232130020-3013031033131220"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_user` property

Type: `"string"`. Computed.

Chap Username to connect to the HPE storage.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [password](data-sources--fleet--reference--group-002.md#canonical-1121120011223221-2200323302012231-1222012300231133-3120111033031022-0210303130321210-2320103330323021-3103002202230203-3131013302202130): complete subsection reference.

<a id="canonical-2330312111310331-0223302321131011-1023231020313012-1133002323300110-1122101312331200-3223210202101120-2021030233003013-2132001013313112"></a>

<a id="canonical-3201230333131332-1000311332111230-3212111131333033-1000021223011031-1312123011202210-3120303231031321-3002313230302331-3233322220032031"></a>

#### `storage_device_list.storage_devices.hpe_storage.storage_server_ip_address` property

Type: `"string"`. Computed.

Storage Server IP address. Enter storage server IP address.

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

<a id="canonical-1220002021111231-0031223333310000-1200032100200232-2300331102021131-3301131102001101-0111101202113003-0133222033113130-1101322020101033"></a>

<a id="canonical-3303001311033311-1303311321313130-0213313030123203-2131031011233012-0200212020120212-0120001113010333-1013230300321223-1323010133012103"></a>

#### `storage_device_list.storage_devices.hpe_storage.storage_server_name` property

Type: `"string"`. Computed.

Storage Server Name. Enter storage server Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-0210031030002121-2031200113300102-2032121230132330-0112121111031321-1000322110223032-1333200011232300-2100113303200201-1120101233131121"></a>

<a id="canonical-3200032013232312-3011012113103221-1010010221333013-2021022211121002-3210031221312113-0303012312012113-3230310121312113-0311103212230200"></a>

#### `storage_device_list.storage_devices.hpe_storage.username` property

Type: `"string"`. Computed.

Username to connect to the HPE storage management IP.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2011133210003233-3131330213022333-1301302221301320-0331302020103111-1130013021301311-0223101011022203-2333310310111200-1313002103033111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

<a id="canonical-3002300233200011-1312213130100100-0232313112203300-2322011203133130-0211323013020113-0310103001321022-1112000003032201-2013223333322232"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-1033000033302103-3012231120313231-2311222130103333-3121210202220102-1333101033031223-0320230013011233-0221201132030230-0213000133010021"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password`

- [blindfold_secret_info](data-sources--fleet--reference--group-002.md#canonical-3300021311030112-0210120120101232-2333200121311000-2023123300221321-1110331011131132-1020130220212120-0120200321200012-2011022331022022): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-002.md#canonical-0131300100122032-2130120010221130-3212110320101213-2000021321321023-0222111012211232-2031123113300133-2132113313131112-0232202213321332): complete subsection reference.

<a id="canonical-3300021311030112-0210120120101232-2333200121311000-2023123300221321-1110331011131132-1020130220212120-0120200321200012-2011022331022022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-002.md#canonical-2011133210003233-3131330213022333-1301302221301320-0331302020103111-1130013021301311-0223101011022203-2333310310111200-1313002103033111)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info

<a id="canonical-2200102010012102-3010112120320011-2301020213233100-2001031132131231-0231033312202200-0301323321020001-0102031012133103-0112303331001112"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1302103301013121-0233022032212223-3032330223033030-0020013301313331-3001122031022300-0132121221113021-2313223020302300-0111232120102301"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info`

<a id="canonical-2230321312010112-1123113310200301-0321232033231221-3210000023212000-2200210103300302-3202303131301120-2012122202312322-0103031120220210"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-3311021013020013-0012130313112002-1303311121022223-1213021130111112-3012310201233333-2133330320202120-0132311300012033-2332310002001011"></a>

<a id="canonical-1320311202130130-3000022022133003-0303313001322222-2023200113110033-3120021310123211-3003112231102323-0123311022130031-3101232330301002"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0330322222223333-3003221111202133-1122230220231133-3002202031102032-0113310113001003-2300013021332132-1221033310311322-0300221203332030"></a>

<a id="canonical-1332331310121102-0313032300233013-0011213202221013-0200023023200300-1201122232000323-0113100331300103-3011331233001231-2021002033033321"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0131300100122032-2130120010221130-3212110320101213-2000021321321023-0222111012211232-2031123113300133-2132113313131112-0232202213321332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-002.md#canonical-2011133210003233-3131330213022333-1301302221301320-0331302020103111-1130013021301311-0223101011022203-2333310310111200-1313002103033111)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info

<a id="canonical-3211331323113323-1330100002010213-0332211113233321-1223230302031003-1202033221022000-2313033120020202-2020110103330323-3031120321211021"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0203001101202111-2313232333331230-2320301101312232-2303320313120311-2211013211323320-2323001320021011-1013300203001221-2310223030221233"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info`

<a id="canonical-0210231030032132-0223203201003202-2233213313222302-2222023011331312-1023032100113211-0310233123311023-1232212111010010-0033303303321203"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2301121123110031-2113211011232000-3111330203211001-0123301300122330-0003310012112112-0110111010223303-3133200232132302-2030232200101323"></a>

<a id="canonical-2201331011013131-3123023113300012-0312231000112111-2021312101111003-3320202131100223-2122200320230112-0132323301020032-0002101022220122"></a>

#### `storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1121120011223221-2200323302012231-1222012300231133-3120111033031022-0210303130321210-2320103330323021-3103002202230203-3131013302202130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.password` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- storage_device_list.storage_devices.hpe_storage.password

<a id="canonical-3022311202213013-0111211232111011-0020331030232013-2120323021100020-2021201033230011-3021113112111000-2021020310302203-0301123201132133"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-1112002023113111-1333011001330321-3320020132200322-2120211203110100-2211233331113020-3230323033200301-0113110103001123-0201312301023213"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.password`

- [blindfold_secret_info](data-sources--fleet--reference--group-002.md#canonical-0232130221121111-1223103201032122-3332221222133123-3012113112033200-0211033023311001-0211332330102000-0201210202331011-3220222021110002): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-002.md#canonical-0130012033220110-0203211303002303-1313210113223103-1201210110022023-1002131232200332-2233130102010010-3132130111120002-0111021202233113): complete subsection reference.

<a id="canonical-0232130221121111-1223103201032122-3332221222133123-3012113112033200-0211033023311001-0211332330102000-0201210202331011-3220222021110002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-002.md#canonical-1121120011223221-2200323302012231-1222012300231133-3120111033031022-0210303130321210-2320103330323021-3103002202230203-3131013302202130)
- storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info

<a id="canonical-3320020032311103-3213233232313221-3330123202331211-1031232003012013-3032112103330110-1231002223302100-2323103131223302-1223023011021300"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1002010210232032-1302201133011133-2131303110331221-0010301302213031-3022033330130232-1311012000011320-0302002023100010-1230330323210101"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info`

<a id="canonical-1121212011333300-3332211131221132-2120320030021303-0230321023023200-3230101310131212-1331032231112012-1023022301123030-3200132100122322"></a>

#### `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-2122101103100010-3011123211303121-2200132303232310-1032123211213331-3201303312002320-3330132112332320-0001320301103321-1213311022110200"></a>

<a id="canonical-1020323120100123-2021300102130220-2233121303001230-0233320331233330-0103230102311010-0012210322021203-3003211210301031-2020121301133021"></a>

#### `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3302011003230232-2331023023303110-3221322311211002-0300130120011320-1231321330100130-1203203113210011-3133121012201012-1311332302303130"></a>

<a id="canonical-2133011232211030-0000102133322010-2232302221113323-1123202300222323-2330110331130310-0013312333303302-3300303200133032-2303101213112102"></a>

#### `storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0130012033220110-0203211303002303-1313210113223103-1201210110022023-1002131232200332-2233130102010010-3132130111120002-0111021202233113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-002.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-002.md#canonical-1121120011223221-2200323302012231-1222012300231133-3120111033031022-0210303130321210-2320103330323021-3103002202230203-3131013302202130)
- storage_device_list.storage_devices.hpe_storage.password.clear_secret_info

<a id="canonical-3000312121032023-3210130231331133-0033213103303301-0011323122303113-3003021012300000-2031213023013232-0320113132120331-0111213313003323"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2032131030231321-1101323201302203-2012132002120301-2322011312120210-2112130030032312-3322303233122202-1010221211211233-1000002310031310"></a>

### Direct properties for `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info`

<a id="canonical-1020302012322333-3121201332233132-0001323120030131-3002203233323322-2100110313121102-1100010213123330-3031011030230121-0231122001110120"></a>

#### `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3132233123232231-2020322101010020-1103312120011203-3322330330220202-1331321322322012-1001232302322100-2212310211123330-1233232000101011"></a>

<a id="canonical-3033331002123323-2031010212122311-3332321313013301-1331021331323331-0320220031030230-1301210000002032-0120223233102000-2203131132123222"></a>

#### `storage_device_list.storage_devices.hpe_storage.password.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- storage_device_list.storage_devices.netapp_trident

<a id="canonical-3212322031313002-0000103303323330-0201110200213003-0303311233302032-2033023220010002-1102111001222103-1310010210112302-1100231333201300"></a>

Type: `"single"`. Computed.

Device configuration for NetApp Trident Storage.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-backend_choice": "[\"netapp_backend_ontap_nas\",\"netapp_backend_ontap_san\"]"
}
```

<a id="canonical-2213100112130102-2203203113213102-1332112333112032-0331322222010132-3311112113203011-3111022322103013-0131120033310010-0202211103132132"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident`

- [netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133): complete subsection reference.

- [netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330): complete subsection reference.

<a id="canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="canonical-3233032001002013-3021100023130033-2100003322311201-2203310113201001-3033030102202132-3321311003001223-3130331230000323-2312111022010011"></a>

Type: `"single"`. Computed.

Configuration of storage backend for NetApp ONTAP NAS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

<a id="canonical-2332030332211201-3023203000310230-3300101323020222-2200022311000033-3021312220123121-0302303031223210-2222013212021310-0002302333131330"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas`

- [auto_export_cidrs](data-sources--fleet--reference--group-002.md#canonical-3113120030020313-1132100122323130-1130312112201311-3233031212003201-2311312103330310-1001203133331302-1223332101303331-3202012131011033): complete subsection reference.

<a id="canonical-3311021333011023-1313302203121110-0231211011221321-2202320120020030-3322121330110331-3000203122312210-1312112302000000-1301333003313022"></a>

<a id="canonical-3123121220110301-3033031023033211-0330012112333003-3121023313113302-2110222210333000-1013322330301112-0020032100320002-2020323323313232"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_policy` property

Type: `"bool"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Enable automatic export policy creation and updating.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2101013322032001-2111233131211122-2122232203003020-3310111113112320-1323100332111232-2300112231213101-2211103221110302-1232313020120213"></a>

<a id="canonical-1013111133012210-3213013130100120-2221213320303020-2200122322032100-3120023322022230-1311132320212112-0330200013311002-0201223221221213"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.backend_name` property

Type: `"string"`. Computed.

Configuration of Backend Name. Driver is name + '\_' + dataLIF.

Additional upstream details:

Driver is name + "\_" + dataLIF.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 50,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2103232230102223-3323030322231200-1221320222113322-2023003022323121-0123030120020203-2233331121302110-2310102011310211-0213021311311103"></a>

<a id="canonical-3330321302223221-2230313023110131-1111011002032201-0333112213332033-3133131322331201-1301020111012110-3103201202011110-2223213031023220"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_certificate` property

Type: `"string"`. Computed.

Please Enter base64-encoded value of client certificate. Used for certificate-based auth.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](data-sources--fleet--reference--group-002.md#canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213): complete subsection reference.

<a id="canonical-3201310203302233-1310321113133113-3220230230131001-1323133101132221-1101223120032303-3232223132320012-1133311011203120-2021021131200121"></a>

<a id="canonical-0133312202033221-1222210100311322-1210201131313300-2023122111021220-2223222012013210-3303220301001223-3231331123213210-0332133333102012"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_dns_name` property

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1300031232220123-2121120221113122-0100131301001201-3233203200203222-2132300030031322-1110101212200211-2112001303112210-3201111100211210"></a>

<a id="canonical-0123000220103210-0310130331312310-3313110221210232-0202100032022220-1333213211331333-2301101203132120-3113322332322330-3111111203033103"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.data_lif_ip` property

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3112233010300131-3311231210010230-3321020232322120-3002100133233330-0010130122031233-2132333033322020-0203333103333101-0110330022112000"></a>

<a id="canonical-0010022002131102-2202303033103031-1203102122311130-2212001220332223-3212030131231031-1030033221033011-3120210003012213-3012123211301012"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.labels` property

Type: `["map", "string"]`. Computed.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-1200312103123212-2013113231211003-1032211132231000-2223232100023021-1310200200111021-3011230032231100-1223313202000131-1301113100110131"></a>

<a id="canonical-1021302333001120-3120123210213121-1001320101232031-0113130213223323-1003210100221021-1333022320030312-2220203102103311-3302033221010102"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_aggregate_usage` property

Type: `"string"`. Computed.

Fail provisioning if usage is above this percentage. Not enforced by default.

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

<a id="canonical-2002322111030132-0312100202112100-0333302021010111-3332103310123103-2320201302201233-0021032131100032-2212331323112120-2121010030133123"></a>

<a id="canonical-2200011031022030-0311203312020122-3332100010333112-3020311121330012-3111121100210123-1003321011310000-2022200002200321-3330010331103210"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.limit_volume_size` property

Type: `"string"`. Computed.

Fail provisioning if requested volume size is above this value. Not enforced by default.

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

<a id="canonical-1331031333022103-2130301132322311-0022011222311023-0013203312223030-1133302231010320-2212001202321221-1200032222220021-2012223323332023"></a>

<a id="canonical-3201000230213230-2110302101310132-0213020202122122-1221032020032020-0333110000321203-0010010213133000-2223211231012323-3131103110113020"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_dns_name` property

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1011201032102120-3333111222212223-2113030202133313-3310311310212011-3323201333203103-0103320133001002-3002031203003321-1030333131012233"></a>

<a id="canonical-3212100330132133-3231303123032130-1111231133022331-2323011123011131-1203320212301112-2220333312211332-2111121103223010-2130113312301330"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.management_lif_ip` property

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0233011322112312-1200231111020212-3313213022310210-1203203122002003-3302213232333303-3330132312330322-0031310221312321-3223201033130311"></a>

<a id="canonical-0122023203110012-0101112220212322-1203033131232033-2013330110312330-2030012233001113-0323201312221003-1331110201313333-3132100113121323"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.nfs_mount_options` property

Type: `"string"`. Computed.

Comma-separated list of NFS mount OPTIONS. Not enforced by default.

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

- [password](data-sources--fleet--reference--group-002.md#canonical-2111022122333231-1020223333303033-3133230013103111-0100210323332213-1002022333132012-2031131203001201-3222003120130130-2112332322021200): complete subsection reference.

<a id="canonical-0303213303210031-0212122212133211-1201332213210201-0201203231300110-3203231311012210-2102330313322000-1200130200322313-1130111223323011"></a>

<a id="canonical-2130300321330323-2310312301331101-3231230230323213-1233233110020332-3100320112320031-0321323223233113-3332230102310110-0333301012103132"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.region` property

Type: `"string"`. Computed.

Backend Region. Virtual Pool Region.

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

- [storage](data-sources--fleet--reference--group-002.md#canonical-2110120020013122-0302220310310103-1003013222131120-1301111323321323-2112112320103320-0131001012111321-0012200230303231-1111121213230133): complete subsection reference.

<a id="canonical-3033301213032113-2330323113102132-1000203032311032-2323202020233120-1010203222202022-2021103123120212-2100323331303323-0021201011210132"></a>

<a id="canonical-3121322322131231-3131230111002223-2113332121122203-0020023213013121-2022222002032013-3201103123103011-1112311133110003-2232013222221230"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_driver_name` property

Type: `"string"`. Computed.

\[Enum: ontap-nas|ontap-nas-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-nas\`, \`ontap-nas-economy\`, \`ontap-nas-flexgroup\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-3020113222210123-0110200213132011-2223211130012322-2012122101330310-3113333132213000-2010230202230113-2132002220032311-3311122111210300"></a>

<a id="canonical-3030313202110132-2211320002322123-1203112133331032-2032022012303321-3011002201010103-2101130120333330-3130320033000201-2330231121222110"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage_prefix` property

Type: `"string"`. Computed.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

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

<a id="canonical-0013110023120213-1130221032313331-2120200112300133-0133222300011102-0102233332313311-3011303212201110-2133230222313002-1101000330221321"></a>

<a id="canonical-3020221233303032-3330001003121230-1132233211333022-1332122200121220-3121230122120222-1113332212033322-2213003020312212-1330111320023200"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.svm` property

Type: `"string"`. Computed.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2023303332033200-3230202031222033-3211200213300120-2133011200121013-2003130003031133-1213323202210112-2111330020213302-2000031223311021"></a>

<a id="canonical-2123232111202330-2002130311220003-0023000212320211-2310310220021131-1312032020001101-2330110010200322-3002331312200330-3231102030033012"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.trusted_ca_certificate` property

Type: `"string"`. Computed.

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

<a id="canonical-3300131302300022-3302103033201332-2222303322201123-3010020033220222-1300332321122311-3030112010203203-2031222223013202-2132032023303120"></a>

<a id="canonical-2311233100002112-3110012330130103-1130310113100232-3013120123231032-3112303130030001-0223231331323032-0223322211012001-1303320023222231"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.username` property

Type: `"string"`. Computed.

Username. Username to connect to the cluster/SVM.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](data-sources--fleet--reference--group-003.md#canonical-1031002301120210-1211113332013001-0221212030021230-1213323311322101-0103111130010301-1132312332210131-1123103220103002-1121031232101023): complete subsection reference.

<a id="canonical-3113120030020313-1132100122323130-1130312112201311-3233031212003201-2311312103330310-1001203133331302-1223332101303331-3202012131011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs

<a id="canonical-0022021131213113-0121021030013132-0221122233332210-2001200320013331-0212121130300121-3332023312033233-2220321021010203-3013013210223102"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3313102023221201-2122210233131100-2222202311012331-2321123100022330-3203011101303233-0223202220110230-0301001330133301-0110233121113332"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs`

<a id="canonical-0201223101123222-1202333220121203-3232303023011121-2122013033113003-2021032010030002-0133012131013210-1312022003110232-0133313110100110"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs.prefixes` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

<a id="canonical-3101033301333120-2110302232130213-0201321203213322-1123313312020232-2312120120131001-1213131231233120-3222221121100113-0203332022300002"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-0300021103101101-0020131211032110-0221220321311331-2312130031311312-0311031321211323-1230101210112331-3203021102232303-0112132033230011"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key`

- [blindfold_secret_info](data-sources--fleet--reference--group-002.md#canonical-3301232302303033-1230110300113201-1202000003222232-1202332230320113-2203001331333023-1233213001321013-2213302102102121-1231313100210302): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-002.md#canonical-2223120110212013-1333023221011333-2320332030300203-0020211013323313-1332233000122311-1212112012322013-1331203213320121-2312002122022220): complete subsection reference.

<a id="canonical-3301232302303033-1230110300113201-1202000003222232-1202332230320113-2203001331333023-1233213001321013-2213302102102121-1231313100210302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-002.md#canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info

<a id="canonical-3232222332110310-0311213202021101-0301110302122331-3033101031021033-1331101132212330-1122010322001030-2332123223120310-2121001123100112"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0331330021110120-3210220323332323-0003212223222300-0313301332301330-3031212331220132-1000323022012300-3213032130030223-0321030102200121"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info`

<a id="canonical-1113011332230003-0101021320122313-0332011133010102-0000332100002300-3013311023122312-1011131323001120-3021100233201120-1320003100322021"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0301002221330110-3302011030032031-2132313110130202-2103320123121211-0012332012220302-0221111222123200-2222131303213131-3032030020030022"></a>

<a id="canonical-3103303310113203-2302000312001311-1021210123110132-0233203012003333-3230331011311312-2220133300230310-2212203300021300-1321030201323001"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3100012231103323-3320312112013330-0130221220221231-3333103233121103-2013300123030021-2231012011201111-2203023333232320-2331321233023020"></a>

<a id="canonical-1211030121302233-3311111301021133-1020032222120320-1321300212003303-2310231123210103-0221331322123220-1011131033233331-2201231003033003"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-2223120110212013-1333023221011333-2320332030300203-0020211013323313-1332233000122311-1212112012322013-1331203213320121-2312002122022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-002.md#canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info

<a id="canonical-0110122013122001-3130103312010112-2310311132223230-1220030100313323-0223300111121000-1121113022220112-1300110003113121-2021133111233333"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1101201011310322-3230122120331212-3122210113012311-3300321121033213-1223131303210003-1013331210013010-0231031032322033-1102232310212003"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info`

<a id="canonical-2120012311222112-0020313220110200-2122130123030221-2132200101303003-2301212202013202-2203333310130312-1031102110033212-0010023203032311"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1202012022133132-3331023301021132-1130332333233021-3231331213303023-2022201310100312-1312000231322302-1133211110320222-2033103231323111"></a>

<a id="canonical-2103120333033113-2001330133310231-3331321032000202-1121221101030021-1000133303322323-1131011303012021-0323121110222032-3320213120133231"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2111022122333231-1020223333303033-3133230013103111-0100210323332213-1002022333132012-2031131203001201-3222003120130130-2112332322021200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password

<a id="canonical-0132202332301213-1211220100201111-3230100100111111-2020130003211121-2210003212200321-1032201202230332-0222201302310333-1022201101302333"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-2330311332130303-1323013323033221-1313013203030000-0030201003110203-1131233112220010-3302011001000003-1313130003301320-1200123333100222"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password`

- [blindfold_secret_info](data-sources--fleet--reference--group-002.md#canonical-0133033201021033-3323220103331213-2133010000231320-2321212112323001-1222222102230011-3003012022312013-2000310022310102-2122321232321121): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-002.md#canonical-0032221332011333-1012012111123231-0332010233101131-1200212030322313-2122331111212210-2033013012133301-3003333000211101-1032300121212131): complete subsection reference.

<a id="canonical-0133033201021033-3323220103331213-2133010000231320-2321212112323001-1222222102230011-3003012022312013-2000310022310102-2122321232321121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-002.md#canonical-2111022122333231-1020223333303033-3133230013103111-0100210323332213-1002022333132012-2031131203001201-3222003120130130-2112332322021200)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info

<a id="canonical-1200002210003303-2130111011102030-0212312132101323-2011331010232132-0101300121102223-1331322222301211-1000223300021112-3013202110201232"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3010220031132233-0331033303321223-1030312112213213-3003003111303132-3130323131200000-0021220211321211-3123002313102311-0122231021333330"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info`

<a id="canonical-0100113300020123-0100110013011301-3120213311111022-3133113232100213-1011311002121301-3200101130032130-3323113000131310-3012010110203310"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0210030313301320-2002111031101330-0323200101121213-3000203212210213-1010113023221202-0301321303123132-0021200201222002-3302303032212303"></a>

<a id="canonical-2213220003323233-3210311010000022-0230212102331212-3221300223222231-1321310100321021-1302031302201232-3033031230101022-1111010121312312"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2223110010330200-1100202321121202-3130012333133333-2223121303123032-2032300010310231-1022300121110210-3322301101312001-2201321013220302"></a>

<a id="canonical-0322031332111112-3213121122022311-0013312012100320-1312133000021130-0321231123222022-0232302231320131-2233311211131120-1131101223210220"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0032221332011333-1012012111123231-0332010233101131-1200212030322313-2122331111212210-2033013012133301-3003333000211101-1032300121212131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-002.md#canonical-2111022122333231-1020223333303033-3133230013103111-0100210323332213-1002022333132012-2031131203001201-3222003120130130-2112332322021200)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info

<a id="canonical-0111321322003023-2003101110132012-1020311222322113-0023233131322123-1033003033033120-2222312032031333-0020132101133010-1030112320321033"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3213022333033222-1221121133201110-1120021230201120-0030311302332022-0222020233313303-3112201030132012-2220110212031003-3303322210230230"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info`

<a id="canonical-3032211231103322-1122130200212330-1000102211010010-2223001221300020-1101122220201322-0331330203202022-3331312230002120-3300311023302221"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2203002033230233-2132033313220301-2121202301113222-2132233111133222-2232111130120130-2030321310011122-3321030120201101-0321331310132022"></a>

<a id="canonical-0300223323322211-2220301212023101-0120021031231103-2332103301013120-2011023030330312-2232231011312310-3011210022110211-2030313230332231"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2110120020013122-0302220310310103-1003013222131120-1301111323321323-2112112320103320-0131001012111321-0012200230303231-1111121213230133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

<a id="canonical-1022312031011321-1112330310300202-1221302332301303-1322221323332033-3120203020001121-3120320311023210-2031222121231101-3003213333303313"></a>

Type: `"list"`. Computed.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1202122020233303-2322302002233202-1101332233211011-1010122211000112-1132321113100023-0313111123200213-3103321200210331-2121021033201232"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage`

<a id="canonical-0101032211001230-0130131011011021-2111232232033300-2212322033102233-0321323331032310-2013021133221113-2321102210130021-0021033123133030"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels` property

Type: `["map", "string"]`. Computed.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](data-sources--fleet--reference--group-002.md#canonical-1211033321203322-2132121023112301-1231232232311122-2110303320220002-0122310010221301-3323203132223102-1121213222001002-1103112012220111): complete subsection reference.

<a id="canonical-2201031100230100-0021012102011032-3231200133001101-0333231311131001-0013002132101011-1020300301313120-0120000203321310-0021320030100002"></a>

<a id="canonical-2233003103332321-0013112110302332-1121331303323031-1301210330202130-2121212310132120-2011302221300010-0013112213021330-2310230213331130"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone` property

Type: `"string"`. Computed.

Virtual Pool Zone. Virtual Storage Pool zone definition.

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

<a id="canonical-1211033321203322-2132121023112301-1231232232311122-2110303320220002-0122310010221301-3323203132223102-1121213222001002-1103112012220111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults` properties

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-002.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-002.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-002.md#canonical-2110120020013122-0302220310310103-1003013222131120-1301111323321323-2112112320103320-0131001012111321-0012200230303231-1111121213230133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults

<a id="canonical-2222002033211122-2322213312112112-1212313101200013-1303033333133103-1103301233211120-3100031131002033-3100313211110333-2300111231100331"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

<a id="canonical-1103020130300030-0302100303112122-0331223022222303-1232112330133103-0113303102011231-1100203022113001-0332203122033200-1231011200232122"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults`

<a id="canonical-0121302320001102-2021300120212030-3232232211113020-3023120010331133-2122112321210032-0210110013230330-3323212202132323-2230010310120022"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1132133313322022-3001320213301011-3332122233212030-3110000010213223-2133030313031231-3313110123103231-3121020203001223-3220122111002033"></a>

<a id="canonical-0230203120002021-2223101203301121-1203033130200033-0311333210211011-0011031131030110-1200013100021013-3120202320113121-1312013223033312"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption` property

Type: `"bool"`. Computed.

Enable Encryption. Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100211330213101-0233201122303301-0130022000113311-3303012320122001-0001211130120130-0123311223001302-2112113322101212-1000301133002033"></a>

<a id="canonical-2100111231130103-0223330312220213-0011310113000120-0021132022223213-3110333132202203-2302221023223031-2120200120120131-1022331120001020"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Export policy to use.

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

- [no_qos](data-sources--fleet--reference--group-003.md#canonical-3113121113310310-2213300013020110-3121130113032000-0201301203122212-1202103110001302-2031103010102301-0200203322033132-3030200300000132): complete subsection reference.

<a id="canonical-1132012000333113-3321133321101222-1202123300330020-0310012332001202-1000310003122023-0211100030311113-1010221003131201-0013020212000023"></a>

<a id="canonical-2212310112131130-3320011102220202-2211301232320110-0220130312200221-1110120102120021-0302022230103212-1322000232002131-2003332312223022"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy` property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2020131220322220-1322011001002223-2313100103112230-1122130101120220-3132130003200012-3201232322020110-1110231111333032-0132021223020332"></a>
