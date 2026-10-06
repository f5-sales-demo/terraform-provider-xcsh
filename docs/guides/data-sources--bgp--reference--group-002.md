---
page_title: "xcsh_bgp reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp reference."
---

# xcsh_bgp reference

<a id="canonical-0303101132100212-3332002022311120-0231032120223302-1131011102232211-3223100221311202-0120302320300312-0331032031111310-3303210120322103"></a>

## `peers.external.interface.namespace` property

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

<a id="canonical-2010302211213103-3013002221020001-3332112312311130-0130302200213211-3313300010320012-3321312021013003-3020320302330112-3122120320101211"></a>

<a id="canonical-3120110311230233-1012100230332300-0323322132002030-3301303321232002-0320132132302230-0112231210023330-0203332102321111-1031200322202123"></a>

## `peers.external.interface.tenant` property

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

<a id="canonical-0332112000321201-2012130032100313-3003023100112021-1221001212101031-3222213330302233-1032331013323132-1312222321313302-1202133211002223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.interface_list` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.interface_list

<a id="canonical-1131003102033211-0333221032220223-0023200032231100-2013300112301130-0031303020113131-2010322132133300-1012200303023331-3131102012332300"></a>

Type: `"single"`. Computed.

Interface List. List of network interfaces.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3002231032131311-0310213320100323-2101332303032233-3212232012132123-2012132130323122-3110332222003110-3003103012001011-0002022300221301"></a>

### Direct properties for `peers.external.interface_list`

- [interfaces](data-sources--bgp--reference--group-002.md#canonical-1330200232221112-0233331010013003-0133020022210001-0001002000131013-3122110233112112-0130321300320101-3213100313302113-3203002330022323): complete subsection reference.

<a id="canonical-1330200232221112-0233331010013003-0133020022210001-0001002000131013-3122110233112112-0130321300320101-3213100313302113-3203002330022323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.interface_list.interfaces` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [peers.external.interface_list](data-sources--bgp--reference--group-002.md#canonical-0332112000321201-2012130032100313-3003023100112021-1221001212101031-3222213330302233-1032331013323132-1312222321313302-1202133211002223)
- peers.external.interface_list.interfaces

<a id="canonical-0110222120101112-0312332110020303-1233001231211312-2011000133033220-1200203201231211-0133010013110113-1111210002300013-3133233300223320"></a>

Type: `"list"`. Computed.

Interface List. List of network interfaces.

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

<a id="canonical-2122102322121023-2101011021122303-2021200213312113-2300313202033300-1031020201101130-3122222210331120-0011130033311201-0322211122013113"></a>

### Direct properties for `peers.external.interface_list.interfaces`

<a id="canonical-1010031202002133-3313203100130222-0330201013102133-3031332223233202-3132211020110213-1210332120130013-0232322312021001-2323021112313303"></a>

#### `peers.external.interface_list.interfaces.name` property

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

<a id="canonical-2323113102032032-2221020100133310-0133010002111130-1222300222302212-3201123212100301-2332331203000113-3013322232132321-0020130332113212"></a>

<a id="canonical-2231212033202200-0310112200321321-0311310123022221-1002323310111032-2301332122232313-2010200201320310-3021200332121302-2310332020202120"></a>

#### `peers.external.interface_list.interfaces.namespace` property

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

<a id="canonical-2212312010113133-3333030010201100-2313310233320222-2103000213010102-1011203222231212-0333022010102102-3221223100331033-3011010022232222"></a>

<a id="canonical-1002232232330320-2121231332220221-0300121310300121-0110231001003102-0100111113031000-1333212010301033-0213120010322123-1013210032300110"></a>

#### `peers.external.interface_list.interfaces.tenant` property

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

<a id="canonical-1120030012320211-2222321211313311-2102010233321200-2013312131010001-1321330100100030-3300110231320011-3010023322032023-0111033230310112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.no_authentication` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.no_authentication

<a id="canonical-3010203300113023-1231010032331030-3213303023310212-2312220222312332-2012023013111300-1011231132301000-0211223031231302-1123031030321122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no authentication.

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

<a id="canonical-1210203212232330-2312201101103002-3021033131220003-3100323101002011-2333003323332203-2332113013012033-3302202221133313-1231022220113331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.metadata` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.metadata

<a id="canonical-0323110232012220-2323021101313302-1131100021310013-2300111000230232-0220333233310311-2210133023310010-2322001020321033-2000302001210002"></a>

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

<a id="canonical-0333001012233302-0111103302032010-2301003231013203-3211321012123013-1011300323331311-3222110130330130-3032102033102300-3000130312300312"></a>

### Direct properties for `peers.metadata`

<a id="canonical-2223311023010213-3121010210313020-2030220320232011-1200132130213032-0322313212333123-0320122132122203-1322103120101210-1031031231120232"></a>

#### `peers.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2302013311232220-1132332100102310-3323223320113000-2000310023001102-0300130112031333-2021013310233013-0130311133130201-1303102322103210"></a>

<a id="canonical-2133332301302101-2232320110110020-2001300110103222-1000210330201203-0210131101002132-3322003031011012-1130222223331312-0313310223221323"></a>

#### `peers.metadata.name` property

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

<a id="canonical-3330222003020013-3002312001203013-3102312030323313-0223103232321233-1203310122113212-3222022000002200-3300330030300020-0011203212322220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.passive_mode_disabled` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.passive_mode_disabled

<a id="canonical-2003303132030221-3021102113031322-3201232230221211-0232032133210320-2002211032030020-3020202121302320-0332212123132202-0113001321121222"></a>

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

<a id="canonical-3111202111113022-2232110120120100-0211312103003231-2022332313131333-2020120311312200-0310213133111102-0030021111133110-1110203031113002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.passive_mode_enabled` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.passive_mode_enabled

<a id="canonical-1120310213202313-2103130203003131-1220330133013100-0203002310202010-0302102131313121-2111323012030323-3212030222032012-1210112221231110"></a>

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

<a id="canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.routing_policies

<a id="canonical-0303322023121133-1100203002202001-0220122012203033-3113131201321322-0203132102321222-3232121332030323-3031221322011110-2303100013303102"></a>

Type: `"single"`. Computed.

List of rules which can be applied on all or particular nodes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0312230123101233-0231122220330203-3031000020110222-2310303021133011-1103012320002223-3012201203121133-2130223133133131-2222123131333332"></a>

### Direct properties for `peers.routing_policies`

- [route_policy](data-sources--bgp--reference--group-002.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213): complete subsection reference.

<a id="canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-002.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- peers.routing_policies.route_policy

<a id="canonical-1212222011200010-0211021023032113-0310303323230131-2022133121130021-2110310322233300-1233132303322002-3201202302003110-0331133131312320"></a>

Type: `"list"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Route policy to be applied.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-0131013011223233-2023123032323003-2332130031133020-0320001213303303-1332213111321300-1201101203132323-2112033111100131-1130100203302031"></a>

### Direct properties for `peers.routing_policies.route_policy`

- [all_nodes](data-sources--bgp--reference--group-002.md#canonical-1220302232032111-3231033031212122-2313203103110102-2021011101033322-2000233220133301-3223011201222103-3211020013233022-0232313021301111): complete subsection reference.

- [inbound](data-sources--bgp--reference--group-002.md#canonical-1232333103020021-1131222111101323-0002120202331110-0320123110230203-0302321323200120-2212113320302320-3213022323211110-0010220110221213): complete subsection reference.

- [node_name](data-sources--bgp--reference--group-002.md#canonical-3130312021232102-1102212310221310-0111003332211020-0112323103100110-0021103221220320-2202321201013013-1333301200000213-1303311003021121): complete subsection reference.

- [object_refs](data-sources--bgp--reference--group-002.md#canonical-1002032323113111-1211032132333310-1312111000010132-0321312013001231-2123000222012221-3002031100311132-2103312112331032-2010201313013100): complete subsection reference.

- [outbound](data-sources--bgp--reference--group-002.md#canonical-2332211331120011-1303021331222323-3132200032031301-0222302012100303-2113103131110121-0313230121332221-1023133020001100-1321221221331102): complete subsection reference.

<a id="canonical-1220302232032111-3231033031212122-2313203103110102-2021011101033322-2000233220133301-3223011201222103-3211020013233022-0232313021301111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy.all_nodes` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-002.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-002.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- peers.routing_policies.route_policy.all_nodes

<a id="canonical-1311112220200302-0101132330010231-3302020331130011-3000211222010120-3130222112113010-0312323131130201-3033310131333121-0020332300013131"></a>

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

<a id="canonical-1232333103020021-1131222111101323-0002120202331110-0320123110230203-0302321323200120-2212113320302320-3213022323211110-0010220110221213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy.inbound` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-002.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-002.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- peers.routing_policies.route_policy.inbound

<a id="canonical-1323200103330211-1213010211022210-3110123201112132-2133111300222332-3113002310011321-2121013101110202-1021221233122312-0331223331023231"></a>

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

<a id="canonical-3130312021232102-1102212310221310-0111003332211020-0112323103100110-0021103221220320-2202321201013013-1333301200000213-1303311003021121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy.node_name` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-002.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-002.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- peers.routing_policies.route_policy.node_name

<a id="canonical-2201013133111121-1020201010313021-3033132103120103-1021220020230210-2232303321111030-1323301310011212-0123231313213322-2001323303210311"></a>

Type: `"single"`. Computed.

List of nodes on which BGP routing policy has to be applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1233010232031010-0002031103213122-2011203102013311-1002332212300030-2020012300023300-1201112311221000-2220030202011033-1133030102223320"></a>

### Direct properties for `peers.routing_policies.route_policy.node_name`

<a id="canonical-2103021021302002-2111000313103212-2222233313312333-2020222330220001-2213021332112033-0022321033022200-1202013130103131-2030333321302010"></a>

#### `peers.routing_policies.route_policy.node_name.node` property

Type: `["list", "string"]`. Computed.

Select BGP Session on which policy will be applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1002032323113111-1211032132333310-1312111000010132-0321312013001231-2123000222012221-3002031100311132-2103312112331032-2010201313013100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy.object_refs` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-002.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-002.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- peers.routing_policies.route_policy.object_refs

<a id="canonical-2111012301313303-0112231230100100-0223033001211232-0022123311202023-0332330322221123-2033131122113120-0000010113121133-3222123313020332"></a>

Type: `"list"`. Computed.

BGP routing policy. Select route policy to apply.

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

<a id="canonical-2311002023320322-0130232001213113-3020303303012232-1102232101322302-1113023210200021-3311121320220320-0012031320213311-3332122010003330"></a>

### Direct properties for `peers.routing_policies.route_policy.object_refs`

<a id="canonical-0202332003120331-1113023102311312-0103212303121131-1012310313320202-1011132023012220-2120003210122002-1322102100121311-3022011022103120"></a>

#### `peers.routing_policies.route_policy.object_refs.kind` property

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

<a id="canonical-3301230133111313-3201102121033202-3201233321320301-1311310120112023-3131023033013331-2330021110003022-2101230112133220-0101331233032110"></a>

<a id="canonical-3303331303311200-0002010233222013-3311123121021110-0000102100200123-0032102331111333-1230100112123233-1112221010112011-1201323301032233"></a>

#### `peers.routing_policies.route_policy.object_refs.name` property

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

<a id="canonical-0212022032123033-1233310233300301-0120112213211032-3120221310033120-1310122012202333-2001312001101202-2332121020002011-3032222023000031"></a>

<a id="canonical-1032113013320301-1312301020232222-1233133012201112-0131322233112011-3211102112111030-0000320032203303-2020221322202033-2201130130332302"></a>

#### `peers.routing_policies.route_policy.object_refs.namespace` property

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

<a id="canonical-0020033130233233-0221110123131301-1211322223030033-1131031323131331-3202131313122310-1010031010110300-3100111301111113-1101122131302110"></a>

<a id="canonical-3012202102322210-0113121212321020-1022232033133021-1321323020320020-3202213022103213-2302332212211322-0122220031031133-1322211102300022"></a>

#### `peers.routing_policies.route_policy.object_refs.tenant` property

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

<a id="canonical-1203011331312213-2200211033202231-0310103203320222-1000133001210233-3220122332331200-0122123310211232-3020132321330032-1230331030312102"></a>

<a id="canonical-3032301302310101-2032310003010320-0311221213210012-1101322223212111-2221013213310313-0122202302311211-2212201030012031-2221113210332032"></a>

#### `peers.routing_policies.route_policy.object_refs.uid` property

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

<a id="canonical-2332211331120011-1303021331222323-3132200032031301-0222302012100303-2113103131110121-0313230121332221-1023133020001100-1321221221331102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.routing_policies.route_policy.outbound` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-002.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-002.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- peers.routing_policies.route_policy.outbound

<a id="canonical-2230131211311312-3310032303030110-2133222212333031-3322220311230132-0233201213210322-1213313033331032-2313100201003102-3121113333001002"></a>

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

<a id="canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- where

<a id="canonical-1032200123222212-3232232020300023-0212233033002010-1202033301101323-3132330001120313-2323310330203233-3001220010112023-3031110023033022"></a>

Type: `"single"`. Computed.

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-1003312111232110-2102333200023130-1233321230203212-1122201000200233-0011130331112132-1132221121310301-3203312220313032-3303310132201032"></a>

### Direct properties for `where`

- [site](data-sources--bgp--reference--group-002.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231): complete subsection reference.

- [virtual_site](data-sources--bgp--reference--group-002.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110): complete subsection reference.

<a id="canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-002.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- where.site

<a id="canonical-0130100031103010-0103100200123001-1333221132210021-1013113302101321-2111022303100303-1131223001303132-3300133301121000-2001012100223110"></a>

Type: `"single"`. Computed.

This specifies a direct reference to a site configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-0202100122002013-3230210302012031-3101320223333121-2313211230221101-2312221200012102-2221103321123300-0223132202230132-3023203000010203"></a>

### Direct properties for `where.site`

- [disable_internet_vip](data-sources--bgp--reference--group-002.md#canonical-3122023133003302-3023202213330003-0231200001213223-2231221223020101-3212031112123311-2313120302110201-2230313302210330-1012032020032320): complete subsection reference.

- [enable_internet_vip](data-sources--bgp--reference--group-002.md#canonical-1111201203213021-3301110111233310-2200022323201221-0102112120331310-0302300200123303-1210001110022232-1023002002230103-2200301321220133): complete subsection reference.

<a id="canonical-0013201030023013-3013120320223303-1303231311013213-0212223132021321-0130222020213212-0323010112233013-0102233123012011-0010313023313302"></a>

<a id="canonical-0020202212333003-2322101223210010-3122101333232112-3231203120301230-1102122313021122-3321121102223031-1300130023111303-0002313011130213"></a>

#### `where.site.network_type` property

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--bgp--reference--group-002.md#canonical-2032131210203002-0232030220020231-0311222200212231-2112302020010131-2132010132112002-3321332212213133-1313032231031213-0120311333001101): complete subsection reference.

<a id="canonical-3122023133003302-3023202213330003-0231200001213223-2231221223020101-3212031112123311-2313120302110201-2230313302210330-1012032020032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-002.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.site](data-sources--bgp--reference--group-002.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231)
- where.site.disable_internet_vip

<a id="canonical-2111130212202122-2010122132311213-1030211313221313-1020023231022331-0323221100113300-0100100113132122-0330332130122332-2131031100333222"></a>

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

<a id="canonical-1111201203213021-3301110111233310-2200022323201221-0102112120331310-0302300200123303-1210001110022232-1023002002230103-2200301321220133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-002.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.site](data-sources--bgp--reference--group-002.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231)
- where.site.enable_internet_vip

<a id="canonical-1223123112103332-2113010300103112-0321103331333231-1000022023300032-0032332112003312-0201011322212202-1203223301232121-1223103133003101"></a>

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

<a id="canonical-2032131210203002-0232030220020231-0311222200212231-2112302020010131-2132010132112002-3321332212213133-1313032231031213-0120311333001101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.ref` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-002.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.site](data-sources--bgp--reference--group-002.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231)
- where.site.ref

<a id="canonical-3213230032233033-1331131121320133-0210121110113120-1132322223032022-2222310300010111-2132033131221311-1233023021213132-3010032222031300"></a>

Type: `"list"`. Computed.

Reference. A site direct reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1031213310200033-3330300300020002-2200201130212013-2020201000230302-1303133123200132-3020111331030231-2020323332122203-3112210132223301"></a>

### Direct properties for `where.site.ref`

<a id="canonical-3310321001030322-0113113110022211-0233211231101120-2013221111322203-1103303200201230-3001230212302103-3221203101020231-1122000332003311"></a>

#### `where.site.ref.kind` property

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

<a id="canonical-1211120301233121-0322122020030002-3210301223221000-3210002033200200-0333220121303112-3023120310033320-0111210222003220-2301113123000133"></a>

<a id="canonical-0222131233102002-3230233210033313-0121102033132122-3333311103230303-3333132212011100-3000230223002331-1122323330111001-2032131002130000"></a>

#### `where.site.ref.name` property

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

<a id="canonical-1110310020320133-1023120201223012-0210002202303031-3030032022203021-2111301001031122-1331232133133333-2233313332112313-3200010302320312"></a>

<a id="canonical-1110213210030313-0033213230211033-0223311103003232-0022010200122330-0222003031111310-3122321310122001-0233212330013220-3201230322000103"></a>

#### `where.site.ref.namespace` property

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

<a id="canonical-3103132103131100-0323131023000023-1120220120133313-0121032322102310-1103011313330321-2020112132200023-3112133313300132-2111013303320020"></a>

<a id="canonical-2110223013310201-0112211013101133-0020112121013213-3112031303323102-1203003121200301-3101233013001010-0022221201130122-0300221213310111"></a>

#### `where.site.ref.tenant` property

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

<a id="canonical-3133002332002232-0322211230012010-3133023133200323-1103203232222332-3223211110123231-2100223303020201-3212010222011011-2332013232232323"></a>

<a id="canonical-1211310300320100-3002013330032112-0213203333333203-0002110130332122-1023301111113111-2113321233020030-2330031031220312-0112001111321220"></a>

#### `where.site.ref.uid` property

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

<a id="canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-002.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- where.virtual_site

<a id="canonical-3320101110123322-1011312002232332-3112322123131231-0301223333010302-0302133212210222-3033203023211103-1030203211121231-0313310101103321"></a>

Type: `"single"`. Computed.

Virtual Site. A reference to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-0222001202000120-0133233013213013-2322020302122311-3133030130030211-1302321302210213-0003222003201333-2131103032132000-0323032202312223"></a>

### Direct properties for `where.virtual_site`

- [disable_internet_vip](data-sources--bgp--reference--group-002.md#canonical-0003322330321310-0321121132103301-3001300220121020-0133002112133020-1323200033212220-0323331020002032-3202320112232122-0033333320130323): complete subsection reference.

- [enable_internet_vip](data-sources--bgp--reference--group-002.md#canonical-1031022120323101-1230331131031223-0010130110120102-3002031333203320-2231001012012212-1011332202111132-0030200133231123-3133131303023213): complete subsection reference.

<a id="canonical-2102131220223300-2101111232022231-1000223032013211-1003113123322001-1131111030113000-1321333310331003-3012221122103331-2103201030023011"></a>

<a id="canonical-0331310311233110-0312323013221031-1120311030131001-1331211202303112-0323032013320231-1201112032013230-0131303330223013-0120323310110231"></a>

#### `where.virtual_site.network_type` property

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--bgp--reference--group-002.md#canonical-2031111210002031-1022301331231023-3001212221103321-2200012000110022-2322013032303131-2033321130332102-1103231300011332-1310313221103001): complete subsection reference.

<a id="canonical-0003322330321310-0321121132103301-3001300220121020-0133002112133020-1323200033212220-0323331020002032-3202320112232122-0033333320130323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-002.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.virtual_site](data-sources--bgp--reference--group-002.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110)
- where.virtual_site.disable_internet_vip

<a id="canonical-3122313321211100-0331010002300012-1322002101211302-1000330021230021-1220322221321123-2312203011311023-3020010330202031-1233330030201001"></a>

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

<a id="canonical-1031022120323101-1230331131031223-0010130110120102-3002031333203320-2231001012012212-1011332202111132-0030200133231123-3133131303023213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-002.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.virtual_site](data-sources--bgp--reference--group-002.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110)
- where.virtual_site.enable_internet_vip

<a id="canonical-2211312233303220-2131111113212001-2312100331023211-0321120203022312-1231231132302000-3312013200222010-3320302311101100-0120020123021212"></a>

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

<a id="canonical-2031111210002031-1022301331231023-3001212221103321-2200012000110022-2322013032303131-2033321130332102-1103231300011332-1310313221103001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.ref` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-002.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.virtual_site](data-sources--bgp--reference--group-002.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110)
- where.virtual_site.ref

<a id="canonical-0110323222102130-2100212211222200-3013310010121220-2001201231032003-2123310330202102-2300100220333032-2012312321132010-3023331221311012"></a>

Type: `"list"`. Computed.

Reference. A virtual\_site direct reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3033111333220332-0122221332020201-1332332122123101-2033230002023111-3121303221123312-0231221321301220-2221023201223212-0223222311102022"></a>

### Direct properties for `where.virtual_site.ref`

<a id="canonical-3111030300032222-2332210030021322-1021310012201320-0231232003020011-3012031123130321-0321111103213302-3000120130311313-3223000022132111"></a>

#### `where.virtual_site.ref.kind` property

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

<a id="canonical-1313001302322031-3322100032323301-0031030220220230-0202333233021210-1323112312303223-0231013132022022-2302100113202312-2312000320210212"></a>

<a id="canonical-3312012213011031-1002322300311201-3333311033322312-2233003020212112-1313200232002322-1202031033320021-0132311211213011-2301100312232030"></a>

#### `where.virtual_site.ref.name` property

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

<a id="canonical-1320121200322031-0003231231001031-3231032202301230-3331201322111231-3222121133202033-0003213120101032-2012322001232001-3201332002212012"></a>

<a id="canonical-3332320000313201-1331213333123211-1022233133222220-1000323012130112-3212231110020331-1233232012203330-2110021323010021-3101331330021223"></a>

#### `where.virtual_site.ref.namespace` property

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

<a id="canonical-0033302331230021-1010030233100303-3032323132303033-3320233303000323-0201010303213210-3331330230211322-2312332013030101-0012103333110101"></a>

<a id="canonical-0321021033313221-3031223102231013-3010333013013312-2230211013030001-3020302021022100-2231133101203213-0320121300330303-1000031000330012"></a>

#### `where.virtual_site.ref.tenant` property

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

<a id="canonical-0021313313213030-2101211032000022-1301103021312311-0011201112303311-0003301213021213-0311102333131121-0013301021313322-3210231111111013"></a>

<a id="canonical-1030023323132123-0020032333132231-0332022333002111-0122311313100301-2321312123003002-3111031030012302-0100020023000121-2130133330201110"></a>

#### `where.virtual_site.ref.uid` property

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
