---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-0022023203222223-2233310313211013-3231110112311020-0302220001032013-2123020333021102-2010333100122021-3300132022010222-2001210102010202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-001.md#canonical-0033330122013310-3003120310012330-3211000223223221-0223111311312230-0022202003132033-0321011121003210-1210122112102020-2213303203301133)
- virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address

<a id="canonical-3220030303033011-3012301301201201-0101030020012131-1020011223121000-3211332310323201-1001300032010220-0220233121023231-1222131010302332"></a>

Type: `"single"`. Computed.

Destination Address Mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3223221032332200-1100313213112021-1322211001122123-0332103331210113-0231202232202310-2023233031323130-0113011123021322-1200202131131221"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address`

<a id="canonical-2010001302211202-1100010010013311-1233103233003303-1033120112232121-1103132203003312-3312101021011210-3301301231332313-2120331213111232"></a>

#### `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask` property

Type: `"number"`. Computed.

Configuration parameter for destination mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1203000232222321-0022300313030100-1233122111223210-1013001222300302-3030121231032212-1011223211131230-0322211311210331-2212223232333230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-001.md#canonical-0033330122013310-3003120310012330-3211000223223221-0223111311312230-0022202003132033-0321011121003210-1210122112102020-2213303203301133)
- virtual_server.connection_rate_limit_mode.per_virtual_server_source_address

<a id="canonical-2200310022132310-0022320033303201-1221121303212313-0023032311121021-0132213113023313-0203322021121123-0310121123003232-2203111130300323"></a>

Type: `"single"`. Computed.

Source Address Mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1223302133202201-1333022120303111-2300032312012111-2233302301230100-1001111300312221-1130211102323231-0201101310002002-0231303023032301"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address`

<a id="canonical-0210301212220123-3231221303212121-0200023030013231-2010121112031133-1210122320102032-0221312232012133-2032100003022101-3322203332220022"></a>

#### `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask` property

Type: `"number"`. Computed.

Configuration parameter for source mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1022121021032130-3303223311200321-3023112202000010-2312212012233102-1313112211301022-3100031320321011-1110221300212230-0022333003220131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-001.md#canonical-0033330122013310-3003120310012330-3211000223223221-0223111311312230-0022202003132033-0321011121003210-1210122112102020-2213303203301133)
- virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address

<a id="canonical-2101111302302031-3303033332212322-2221033113310323-2212330000212022-0031231213311210-1013012320030012-2123011313130230-2121022213132100"></a>

Type: `"single"`. Computed.

Destination and Source Address Mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2233103022000003-1221300123031001-1333232310322200-2101333002112013-1210022020200200-3110202232220301-3202131003301303-2323110201133300"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address`

<a id="canonical-3331131121102303-2113213123122300-2100102030220210-0100231313300302-0333313102223213-3203320303213332-1223313022313131-0033133122300323"></a>

#### `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask` property

Type: `"number"`. Computed.

Configuration parameter for destination mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-2201301222220233-0323210012201101-1011301021320331-1202310001012032-3201303223330020-3112323200233000-0000031021123301-3203230213203121"></a>

<a id="canonical-1202310020011030-0301101222231121-1120011200100023-1220000030330031-2121022311313122-3133231233301113-2121132330010031-1303332223303012"></a>

#### `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask` property

Type: `"number"`. Computed.

Configuration parameter for source mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-0003121231013131-0212113030120031-0011021022033310-1103303313001332-1210322112031030-2212031333233022-0102311320322312-2320101203011201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.default_persistence_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.default_persistence_profile

<a id="canonical-0132132103200311-2033001011211102-0100003330302202-0233230330203011-2233203102120101-2103030030021102-3123032023321232-2003112322313011"></a>

Type: `"list"`. Computed.

Configuration parameter for default persistence profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1301331313300231-0221330302011112-3012022301122113-2023300330131210-2013321313120320-2331202022100303-3300133231313223-1102321230211020"></a>

### Direct properties for `virtual_server.default_persistence_profile`

<a id="canonical-1003033222132011-1212223230221031-2122020121312023-3200220110320323-2202033120032100-2221200130121120-1111130032013303-1323321123211022"></a>

#### `virtual_server.default_persistence_profile.kind` property

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

<a id="canonical-0323022333203032-0212101021130100-2302201310301013-0222311202013221-0131201220112030-2000231322101020-3112210220032112-0010110110013033"></a>

<a id="canonical-3301202001320223-1022202210201331-3111123212302113-0130230130310211-2210232021001301-1212233213123021-3221133010213232-0203330010022020"></a>

#### `virtual_server.default_persistence_profile.name` property

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

<a id="canonical-2121032020200303-2110122310333003-0230110011331321-3021013132210133-2120003102102211-3111120101022303-0111322101212100-1123201200120132"></a>

<a id="canonical-1003310331101322-1233100023032232-2303233001320200-3121311311201132-2210223021333212-2013122323230000-0032320130021323-3130013200123313"></a>

#### `virtual_server.default_persistence_profile.namespace` property

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

<a id="canonical-3330133112002113-2311000132033033-2131202220220302-3320331331211023-1002010111320320-2323120222230320-2222203303200211-3121233133002222"></a>

<a id="canonical-0320313230313223-3311321322121332-0032201102300311-0023320032213132-2003113132230332-0310020212133231-1031033110031120-0032103010222110"></a>

#### `virtual_server.default_persistence_profile.tenant` property

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

<a id="canonical-2001112003110002-0113210100223020-0212001313320330-0201231332103133-3122033311100110-3012310202321213-0030000331123200-3203111200313103"></a>

<a id="canonical-3210020133020202-1133320020301012-2012223313223020-0211221130002213-0213110102102301-1130120003303102-0211223022012321-2301210212320222"></a>

#### `virtual_server.default_persistence_profile.uid` property

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

<a id="canonical-0333332101230033-1320131322211033-1021010222233010-0332010100110010-2322112110200221-2223211110003312-2100223212033220-1202302003303231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.default_pool` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.default_pool

<a id="canonical-0032003130213122-2132210021133011-3210323101123201-1331222302230120-0032201321121012-3110301121223230-0003122221011010-1032022120330101"></a>

Type: `"list"`. Computed.

Specifies the pool name that you want the virtual server to use as the default pool. A load
balancing virtual server sends traffic to this pool automatically, unless an iRule directs the
server to send the traffic to another pool instead.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1312122300113002-0201022001321323-3203200122312312-3133101032213002-1220202312213330-1121330203022132-0211121220113003-3002133123013230"></a>

### Direct properties for `virtual_server.default_pool`

<a id="canonical-2001322211031033-3012213203300203-0332021203222310-2022130233030313-0211030233313101-2311132012310131-1031331002122321-2211100200300020"></a>

#### `virtual_server.default_pool.kind` property

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

<a id="canonical-3032213320312323-1222330100023233-0203210311022012-2031300302130321-2113001100321111-1031330222223021-0022211120231000-3010311211301130"></a>

<a id="canonical-2112102322301313-0210311322233203-1111313102201010-2033331312302133-3112121133011331-1023313101013202-1013202223101201-3202132313011201"></a>

#### `virtual_server.default_pool.name` property

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

<a id="canonical-2213113003303231-3012211330331322-2322232121021123-0301301300313033-2213011220003202-2202123002313023-1110221223322320-2201113301211122"></a>

<a id="canonical-2333022203013200-2230011301232320-0011213330103311-3100102110200322-3132333220012320-3331002321133200-0012033320301031-2110233032300310"></a>

#### `virtual_server.default_pool.namespace` property

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

<a id="canonical-2101231233003012-1022330031222033-3333232013003103-3000200131232022-3211121313031313-3210202200021021-1023111210130210-1330012202211330"></a>

<a id="canonical-2330230322311001-3112032102303120-3131220000101221-3231022020020231-1030213011113013-3130001101110321-3211123232113133-0200000220222130"></a>

#### `virtual_server.default_pool.tenant` property

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

<a id="canonical-1102111030133221-2033112123021013-0103232100123020-2101023303300120-2303133300032021-1132011233300002-1120020300312300-0203030222333133"></a>

<a id="canonical-0312002212320123-0112101122311220-1303221103133103-0301120123302213-3201311132201311-2320111103112133-2220302111032003-3301203021232112"></a>

#### `virtual_server.default_pool.uid` property

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

<a id="canonical-3313233021210032-2130033220330002-1330201211103120-1320001012332020-3303101002000220-0110120200003300-2322012100222132-0123332102020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.fallback_persistence_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.fallback_persistence_profile

<a id="canonical-3031222100321330-0030112200331121-0332312032101230-1311302311301132-3331230021310111-3211231210330021-0020121210000121-0002230321230031"></a>

Type: `"list"`. Computed.

Configuration parameter for fallback persistence profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1022112211113200-0031302000323033-0313012320123110-2001231013012331-3102212012232111-1311010132303021-1331312223220123-2220131022033011"></a>

### Direct properties for `virtual_server.fallback_persistence_profile`

<a id="canonical-0030302100303202-1212333230313111-0303102311120132-3221302203132232-0003230200203003-3113322120233110-0310222332331220-1013031233323303"></a>

#### `virtual_server.fallback_persistence_profile.kind` property

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

<a id="canonical-2011111330302121-0333221012103021-2300221212230222-1002311200222123-2120302201321022-0023113201310232-3222100312012030-1120032032031021"></a>

<a id="canonical-1302121322033122-2321021130200112-3112120001111012-0223313310032010-0103233220110320-0322223222121003-0032210313333012-2210211221233013"></a>

#### `virtual_server.fallback_persistence_profile.name` property

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

<a id="canonical-0303302310133032-3032201103100331-2010020001010120-3331030321032303-0012223230010003-3321023122310122-1003202200132331-3001130133211110"></a>

<a id="canonical-1113103300210231-1131212013103100-2102223033031202-3022021312032302-2313222230321013-1303210230313322-1131023111112131-1102013021221122"></a>

#### `virtual_server.fallback_persistence_profile.namespace` property

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

<a id="canonical-3201033001223330-3121021210100210-0300013013232022-2231132300230223-1312020323022223-1330123210130203-0320221230230233-2022203031303110"></a>

<a id="canonical-0012230122131302-1222223200131201-2123020300133133-0201223201221212-3101331232121102-1022301110012102-0231111321131032-3213121013332300"></a>

#### `virtual_server.fallback_persistence_profile.tenant` property

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

<a id="canonical-1233332121110312-2121230310203112-2122220322000230-0003010021220333-0113233222031100-1103103221322123-0021102002030130-0230131202201112"></a>

<a id="canonical-1232311311231031-1000211111303101-1332023302012020-1031211123200102-1312322001303120-3110013030310120-0032121030211313-0232303300200230"></a>

#### `virtual_server.fallback_persistence_profile.uid` property

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

<a id="canonical-3332013220003223-0301321210023212-0103101231222131-1313130322011033-3331230212003132-3203221202003221-3220112221012033-3121020033101103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.fix_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.fix_profile

<a id="canonical-2333300011301210-3213230203312113-1013022110101200-3032023200322121-0233300201201000-2231022121031032-0010213221331220-2022133112030130"></a>

Type: `"list"`. Computed.

Configuration parameter for fix profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3110301330010131-1223012122310221-3011200301103101-3321203320333132-2101033133231231-1131132322300301-2212021222211233-3113220322032320"></a>

### Direct properties for `virtual_server.fix_profile`

<a id="canonical-1122323333132002-1031133023202302-0020013310120130-2302230332230021-3113310310020301-1013332300031313-0321022323310031-1213321111323220"></a>

#### `virtual_server.fix_profile.kind` property

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

<a id="canonical-1202213102011031-0321201101321200-1230021110201302-2213102011121020-2301120021313320-3323131033231323-3322323101322211-1002113330323223"></a>

<a id="canonical-2011032002023231-2303121013311310-1123030221223011-3210000332231003-3121022302133310-0103312022313231-3121133121331010-3002310000032112"></a>

#### `virtual_server.fix_profile.name` property

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

<a id="canonical-1022213132133231-3232100032330111-3122311330023201-0302301220113330-3112201112001013-3003130000211031-3110213122111312-3232202203300201"></a>

<a id="canonical-3310020020202011-1030011201001111-1113111121300300-2333330110000233-1122220303110310-1103313110011132-3022322032030133-1101222322121320"></a>

#### `virtual_server.fix_profile.namespace` property

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

<a id="canonical-2210201132131102-1311223300223201-2300123300333022-0211333211131023-3220302112010123-0231103320230133-1322301332231230-1332032320230131"></a>

<a id="canonical-0322303031323102-2300001131321131-3031010101002301-2232210311232012-1313003303212311-0232123300302200-2301001300221021-3030301133101220"></a>

#### `virtual_server.fix_profile.tenant` property

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

<a id="canonical-0301101102332320-3101201312310111-3312202310302130-1030213111232002-0012022000032332-1201331030013301-2113312011323232-3122113333122331"></a>

<a id="canonical-1333130000002333-3212320201312100-0232322310123313-1223130222100231-0020311300012131-0301011302301122-0200111123100100-2121113023003230"></a>

#### `virtual_server.fix_profile.uid` property

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

<a id="canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.http

<a id="canonical-0332210133000202-2230023030102321-1003012303213311-3321201223311210-2203110311213223-0000100332023203-2000332111113203-0210213200130102"></a>

Type: `"single"`. Computed.

HTTP profiles.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0333130231020312-3132313301321210-3212211203031302-3222322120000113-1003332133110130-2120003302101001-0201232210031111-0131312211310033"></a>

### Direct properties for `virtual_server.http`

- [client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-0001322323303133-1333133021112302-1002230001003010-1101232133321100-2321311303023022-2210002003111103-3000302113211331-2122123130110023): complete subsection reference.

- [http2_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-2221201120100212-2122011231330310-2001222123103101-2002103323100303-1131303312031130-2332201210122333-3002002110202002-1312230010113033): complete subsection reference.

- [http2_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-1211301201001113-3230033010111111-1321001022132330-3200013303310101-1323132133122333-2301031102303312-3201122203201320-2230221322031221): complete subsection reference.

- [http_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-0002203130231201-2231310003011120-3032103022003312-0022312232033011-2212010301230323-2100132120103211-3320022311300332-0023010102321032): complete subsection reference.

- [http_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-2113110223002213-2212301222102222-0033301103333223-2022323121102003-0310330202000233-2131311133211121-1223132213233312-2323020202323010): complete subsection reference.

- [ocsp_profile](data-sources--application_profiles--reference--group-002.md#canonical-2112012213112112-3101111330121313-0001001132012010-2111302203031223-3313103021103201-0011311330001003-2120120120323103-0112123212202203): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-1011200300103030-0110101233023121-1133301320122122-3322213123110332-2002112030013303-0001022210230030-2303320330203331-0022102012131003): complete subsection reference.

- [stream_profile](data-sources--application_profiles--reference--group-002.md#canonical-0123230200023201-2132000020310100-1111312322311233-2023222212322111-1333231221330300-1211033002101103-3031311120112113-2313032102130102): complete subsection reference.

- [tcp_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-2222233310322030-3130301110121031-3311131123110102-2101322013110330-1013333113321210-1122223131221313-0233232232213123-3231210122131302): complete subsection reference.

- [tcp_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-3132223203012130-1010101120231200-2220012111301303-2023302011121212-2201323032110230-2202013320230313-0000110302021211-2322101122330013): complete subsection reference.

- [websocket_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-2203212220330201-2211030130032001-2311212010311313-1301011121033212-1103323033323000-2231211010013022-0310112133313231-0101330123321022): complete subsection reference.

- [websocket_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-0000013133110310-1202210210001023-2310112231201200-3221210302130202-0000111312112212-3300323001032133-2202301013123233-3013100021021302): complete subsection reference.

<a id="canonical-0001322323303133-1333133021112302-1002230001003010-1101232133321100-2321311303023022-2210002003111103-3000302113211331-2122123130110023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.client_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.client_ssl_profile

<a id="canonical-0123212130303322-2310313033030130-0210020200201023-0021321231300010-2132332222301330-1101211203203322-3312033201032210-1130021131011132"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2030313000222301-1110230321023230-0202231312023102-3302222231000031-3032211122312020-1322332310331131-2113220231002032-1130123120103230"></a>

### Direct properties for `virtual_server.http.client_ssl_profile`

<a id="canonical-0330212301111223-1331320121103011-0212221011211302-3303331313023020-1200330111331303-0102303120033001-3213310120312331-3332321033233232"></a>

#### `virtual_server.http.client_ssl_profile.kind` property

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

<a id="canonical-1301013003203323-3121333113231220-3122313222231100-0211133220321300-2130023302300010-3033212131012010-1101331330003003-3320203132312211"></a>

<a id="canonical-1212222033230023-3002003033231023-1002110122001202-0132302311120312-1000230012130030-3223110321301000-2020010232003031-3323301021131221"></a>

#### `virtual_server.http.client_ssl_profile.name` property

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

<a id="canonical-3202100111313031-1033032122032101-1221112321211023-3220002113000231-1021032033321032-1110102212212031-3010331220333030-0222131333330322"></a>

<a id="canonical-0300311231320302-1332102213300222-2212323013323000-2130200101120302-2133201301132033-1310313131121310-1302112232311032-1303231311021002"></a>

#### `virtual_server.http.client_ssl_profile.namespace` property

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

<a id="canonical-0011122300002311-0022113012033013-0100020111021212-3203312231221023-1331210321103232-0332220223030030-0213210003023121-0112002313221222"></a>

<a id="canonical-1013132301011111-3002032000101033-3110200322313130-3132123123330133-1231131122110011-1100233300033123-0012100211330101-1021200030323220"></a>

#### `virtual_server.http.client_ssl_profile.tenant` property

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

<a id="canonical-3003121022321333-2131310000203333-2000032211011121-0003102030023003-1211123100111121-1213121300012333-0230132301332200-1013033301331001"></a>

<a id="canonical-2012023212122312-1012223031310233-2013101000111322-2200202121320002-2201010203123103-0011002203223312-2002312223310211-2011023002101212"></a>

#### `virtual_server.http.client_ssl_profile.uid` property

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

<a id="canonical-2221201120100212-2122011231330310-2001222123103101-2002103323100303-1131303312031130-2332201210122333-3002002110202002-1312230010113033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.http2_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.http2_client_profile

<a id="canonical-0012301231330132-1110003010313111-1012101002322020-2033103122023020-2112120112323313-2021302032211212-3123123233130132-1102123322120112"></a>

Type: `"list"`. Computed.

HTTP/2 Profile Client. Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0130100310020113-0130101331120312-0120322022013112-0211223211002210-1323233332331202-2212333230322211-0120113031023031-2232123210200313"></a>

### Direct properties for `virtual_server.http.http2_client_profile`

<a id="canonical-1121320112113233-2321000023103231-3232010021233333-0102102331321232-1031220113320312-0200000101222203-2312002111011032-1331333213011222"></a>

#### `virtual_server.http.http2_client_profile.kind` property

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

<a id="canonical-0333113220111131-0200313300010222-1033201120321030-0333010222010011-1133132202332130-1223013300012111-1320233011230332-1332311102301102"></a>

<a id="canonical-2313031233220123-1122110133322131-3232001322023011-2223332222113220-3222320203232013-1222033231203112-0003231332032020-0020022323133302"></a>

#### `virtual_server.http.http2_client_profile.name` property

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

<a id="canonical-3011122323210220-2300230321021320-3030310102122103-1222031033020212-3000110301030203-3123311202021100-1112020220302313-0130303101003130"></a>

<a id="canonical-3130013012322031-0233221122121011-2003201031121213-1100102131212013-0200122211002101-3021103303003233-0003111112101012-2123031221300132"></a>

#### `virtual_server.http.http2_client_profile.namespace` property

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

<a id="canonical-0023132203130003-3100033101022030-2233220303000121-3333102120213030-0332021010310331-0320301323221122-1023123002231333-1102301202213102"></a>

<a id="canonical-3302121200033100-3201313322030101-1030012331230002-2120200131011303-3230031333331120-1123231230231110-3122030203221033-2222032223333221"></a>

#### `virtual_server.http.http2_client_profile.tenant` property

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

<a id="canonical-3310202122033310-0333103032303210-0111200102021023-0011231000220302-1322130233221310-0331301312021102-3122101102332021-0323123001232021"></a>

<a id="canonical-0003121013113113-1221101030101013-0113021021112033-3133333313233332-2322231203320123-3210101331300122-0003120120211010-1122110000232011"></a>

#### `virtual_server.http.http2_client_profile.uid` property

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

<a id="canonical-1211301201001113-3230033010111111-1321001022132330-3200013303310101-1323132133122333-2301031102303312-3201122203201320-2230221322031221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.http2_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.http2_server_profile

<a id="canonical-3100221000003202-0011121223033221-0100212112312130-2332013321322131-2011132332002032-0103012333032331-0033303331120310-2322313033002211"></a>

Type: `"list"`. Computed.

Configuration parameter for http2 server profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1303001002111222-3313031303331331-1000322333132101-2220322210200333-0300230301223102-0001033013102122-2321121002020131-0100220111323030"></a>

### Direct properties for `virtual_server.http.http2_server_profile`

<a id="canonical-3003201302123102-0131323320121100-0312023002330311-0213200212301013-3200323032133303-0320131331310023-3011232121133130-0333312223212110"></a>

#### `virtual_server.http.http2_server_profile.kind` property

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

<a id="canonical-0323000333332100-0212123233220221-1013023332202330-1302132132022331-3310133030130320-2003110323303230-3111001213100333-1223200020210133"></a>

<a id="canonical-0212120212331320-0033133331112212-3202232211301112-0132313202212102-0023133030313120-0211132022003110-3301021310133232-0102300023013221"></a>

#### `virtual_server.http.http2_server_profile.name` property

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

<a id="canonical-2023133010231332-2322113012220303-3223103233230002-2220312133100221-3210111012001213-1003311203113022-3232213213102311-1112122030231232"></a>

<a id="canonical-1200213131221313-2310300221320311-1323112202013222-0231113301231323-3223130113220210-1310222013333010-0030100233112002-3020031123312220"></a>

#### `virtual_server.http.http2_server_profile.namespace` property

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

<a id="canonical-2230201011113020-0303123301033333-0333232122013021-1312132130100211-1333322202132023-3230002122003302-0331302021010310-3333030123012312"></a>

<a id="canonical-2030102202120323-1213100322030231-3130002112211311-0123123200101201-3323330310331313-0233323030230302-2302312200002033-0002000211223200"></a>

#### `virtual_server.http.http2_server_profile.tenant` property

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

<a id="canonical-0221011001313102-2201313332032313-1213311303203310-1103222111020332-1131000202123010-2131012000332021-2022220302221003-2110223233200011"></a>

<a id="canonical-3033311223232012-3301133113102321-3121130331002130-3320130212000123-3223231203222120-0220223320333021-0222301122013332-2231312222032131"></a>

#### `virtual_server.http.http2_server_profile.uid` property

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

<a id="canonical-0002203130231201-2231310003011120-3032103022003312-0022312232033011-2212010301230323-2100132120103211-3320022311300332-0023010102321032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.http_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.http_client_profile

<a id="canonical-0122213013203311-0011121202023332-1112322313000120-3233320001311220-1000111203212220-3103101321130103-3221111101110100-2110113111101021"></a>

Type: `"list"`. Computed.

HTTP Profile (Client). Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3022210313331020-0323130112113230-2302122102212300-2000223013101233-1131002011200303-0000312231120312-0212333323223002-2220233222110302"></a>

### Direct properties for `virtual_server.http.http_client_profile`

<a id="canonical-3130223010330110-3111311310000132-3130133113011032-0003311000320310-2201012310213331-0013203313230313-3220010003322101-2102202312303120"></a>

#### `virtual_server.http.http_client_profile.kind` property

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

<a id="canonical-2030010032032303-0223103012303101-1321230032221332-0331332221202220-0130331323221231-0232212311303331-2021203323123203-0101323320330013"></a>

<a id="canonical-2310000302302233-1332221313000023-3223202113320011-2003201302021030-1013201223300120-1312031023120211-1132001011222212-3303032001001202"></a>

#### `virtual_server.http.http_client_profile.name` property

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

<a id="canonical-3302333001011032-2331022302123000-1002232303031221-0313332232300122-0131210102123221-1101301131232013-1212001201210220-3223303202021302"></a>

<a id="canonical-2320013012323302-1121010331011200-0320210113103211-2221322322302212-2323200021113013-3212221020223203-2131220030133300-2030120303220310"></a>

#### `virtual_server.http.http_client_profile.namespace` property

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

<a id="canonical-2103131031223013-3213122033332101-2203321320001313-1221011301300011-3301123011020222-1202103112032001-1212233221012303-2023202102012103"></a>

<a id="canonical-0223020133231302-2210103122322130-1130023332123023-1302113013301203-2002021300120100-1331002130333011-2113100012322222-1202032121031122"></a>

#### `virtual_server.http.http_client_profile.tenant` property

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

<a id="canonical-0022030222030032-1221331331300322-1313031203311011-2223010123220202-3330002332112323-3031130300322331-0332113301120201-0121330030002222"></a>

<a id="canonical-2333320112002001-0203132103302223-1213333002011120-0322120320031323-3300112223022121-2003203202230102-1111320301012330-2121123120122220"></a>

#### `virtual_server.http.http_client_profile.uid` property

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

<a id="canonical-2113110223002213-2212301222102222-0033301103333223-2022323121102003-0310330202000233-2131311133211121-1223132213233312-2323020202323010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.http_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.http_server_profile

<a id="canonical-2222000223211031-2223113121002311-0222131313111320-3033332030100010-2312312323103010-2303030133110132-2300013102311330-2312310013203000"></a>

Type: `"list"`. Computed.

Configuration parameter for http server profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1120222112123311-2221011132021002-2333020023101212-1323312022231220-3133020231112210-2232300113102310-0031133010032031-1013320330113301"></a>

### Direct properties for `virtual_server.http.http_server_profile`

<a id="canonical-0220020120332021-1320223210302022-2012100001300212-0210310322201231-3131122303321211-3031230200122331-0313230023322333-2003201111210133"></a>

#### `virtual_server.http.http_server_profile.kind` property

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

<a id="canonical-0330022211333111-2133203201302210-2023133103232103-0003112023321033-1202010012013121-0120300310022303-0022101201332130-2311001330212123"></a>

<a id="canonical-1130110003002113-1233012222131112-0111101112023300-3130320000303131-2302031020211220-1022133120213131-3221003333221011-1200330213221203"></a>

#### `virtual_server.http.http_server_profile.name` property

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

<a id="canonical-0221113310311100-2303211132011110-2132010031333232-3212122120332330-3132330122031230-2213233101002201-0021100300201332-2300032210022022"></a>

<a id="canonical-0322012212202200-0022212101122311-3230200320113111-1310301130233022-1131202322220002-2313231002131202-1322203302031301-3021021003302122"></a>

#### `virtual_server.http.http_server_profile.namespace` property

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

<a id="canonical-1123230331322200-0032200020310200-3300311113320211-3333330011121301-0032131300123023-3012111131010021-1031303003012033-3002031313020300"></a>

<a id="canonical-1302013123001003-3002030312120112-0113330003123120-0221332301313113-1000112123000123-0222222101032220-3110320023213011-2221001000031233"></a>

#### `virtual_server.http.http_server_profile.tenant` property

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

<a id="canonical-2111302002103033-0130233300201033-2311032012120212-2002332233333122-3201023331132122-1231302011033020-0021232302301031-1323113020231331"></a>

<a id="canonical-1031120231022200-3132122330100302-0213120211311133-1101322121001301-3033311012112233-0233310031110210-0210132111010003-1001222110110220"></a>

#### `virtual_server.http.http_server_profile.uid` property

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

<a id="canonical-2112012213112112-3101111330121313-0001001132012010-2111302203031223-3313103021103201-0011311330001003-2120120120323103-0112123212202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.ocsp_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.ocsp_profile

<a id="canonical-0312201123200303-2231230312032301-0320013132131331-2110201112323330-0303012111100303-2113311232223102-1132201032212311-0223113002312122"></a>

Type: `"list"`. Computed.

Configuration parameter for ocsp profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1012333022202131-3333231203033113-2111302111130023-0301131021312121-0011110201220111-2032131220130223-2103232102202200-2222213312330010"></a>

### Direct properties for `virtual_server.http.ocsp_profile`

<a id="canonical-1230321100122233-1330030032331102-1011321132311202-1230201333001112-3020131021003002-2313032102101122-2033020111301233-3310222001333002"></a>

#### `virtual_server.http.ocsp_profile.kind` property

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

<a id="canonical-2300120023011101-1003011311103222-1021331130111103-0110312333233121-2210110233332230-3323203223021213-3011323230123332-1113121003220113"></a>

<a id="canonical-3032023321231010-0322302312211330-1233211000013310-0023131221332200-2310102200001003-1023120121031103-1203211020331103-3233121202200223"></a>

#### `virtual_server.http.ocsp_profile.name` property

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

<a id="canonical-1013230020021112-2103023303323201-1313311132102120-1213302210323321-1231013202311330-2221011000321322-0233032101210112-2111100112223332"></a>

<a id="canonical-3201323310022212-1033203211112023-2123312233011103-0120210300200123-0311030111133200-0303331330200103-2020332130110332-2010022200201313"></a>

#### `virtual_server.http.ocsp_profile.namespace` property

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

<a id="canonical-0100022002001322-1021321332112223-3311221331010022-3020313303313020-0002010112210202-1021000120223003-3212212132212021-2303310033001010"></a>

<a id="canonical-1031110123220110-3300020202310111-1012231121111210-0302033133223133-2220221023103100-0022023313020100-2230313213321012-2331333231313032"></a>

#### `virtual_server.http.ocsp_profile.tenant` property

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

<a id="canonical-0012103220212103-3133322111302033-3111110101321033-2133221330012002-0021020122032223-2111033232303201-2033022300101231-3322220221110322"></a>

<a id="canonical-3000133031211220-3033332112300313-0103110101213133-1123100230202102-0022031101302130-0013302230123130-2011133022202220-1003030211100323"></a>

#### `virtual_server.http.ocsp_profile.uid` property

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

<a id="canonical-1011200300103030-0110101233023121-1133301320122122-3322213123110332-2002112030013303-0001022210230030-2303320330203331-0022102012131003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.server_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.server_ssl_profile

<a id="canonical-1002021031003123-3301122331321210-2221332022112120-0113010120102001-2113123001232222-2032211313333321-0210030103211211-2311113023013102"></a>

Type: `"list"`. Computed.

Configuration parameter for server SSL profile.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0303332320230313-2021120011220303-1213031032020311-1102111320323013-3333111000033223-3100210111322220-1201311320000330-2300223111311113"></a>

### Direct properties for `virtual_server.http.server_ssl_profile`

<a id="canonical-3311012123223100-0110231323110230-0211311200221320-3221103133333001-0212032323233010-2130021033312022-3030030020121301-3002132121032210"></a>

#### `virtual_server.http.server_ssl_profile.kind` property

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

<a id="canonical-1030021331210101-2200002322120211-2323013310202120-1203021323310130-1210331100232330-3120031033132003-2331022130003202-0122030011013203"></a>

<a id="canonical-2033231102211023-1310322011112331-2102332100110301-2021031221230200-0320112030012003-0222210211233232-1321001001303003-2130010010200303"></a>

#### `virtual_server.http.server_ssl_profile.name` property

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

<a id="canonical-3113120212023101-2300332332201020-0101203020000121-2201310223112231-0001201222020302-1300033002210211-0022321321223200-3122103120231000"></a>

<a id="canonical-2302010033323001-2220221303301122-0200310020210102-3211212300031110-3313201010112120-3220320212132030-0311121320203022-1301101033001302"></a>

#### `virtual_server.http.server_ssl_profile.namespace` property

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

<a id="canonical-3031231301112020-0231332321123113-0300221303333200-1301012232120202-0021010332013201-3121003000301320-3121130131313122-0201320020303223"></a>

<a id="canonical-3302100221121210-1011320311332030-2211032002003002-3321323020323202-0013220330010022-0203223210221321-2110320132001033-3000030200123130"></a>

#### `virtual_server.http.server_ssl_profile.tenant` property

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

<a id="canonical-1122012112311002-2201110321212122-1321110010312023-0210330310203121-1020231213111131-0022303111300131-2231123123033200-0130200022003103"></a>

<a id="canonical-3220012323021033-0313132123310333-3123130201033333-3331012311013310-3233110221132133-3231213222010103-0312120220203320-3122000323203033"></a>

#### `virtual_server.http.server_ssl_profile.uid` property

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

<a id="canonical-0123230200023201-2132000020310100-1111312322311233-2023222212322111-1333231221330300-1211033002101103-3031311120112113-2313032102130102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.stream_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.stream_profile

<a id="canonical-3132213201233220-2021321102011100-2220012312310301-2111100133020122-0113222003302321-0210230000211023-0212123303121012-0330303213211020"></a>

Type: `"list"`. Computed.

Configuration parameter for stream profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2213303111300312-3012022221113023-2230010002022312-0110313212100022-2113003330003000-0022211102313200-1210123023302002-3022221030333302"></a>

### Direct properties for `virtual_server.http.stream_profile`

<a id="canonical-0230000210332311-3222330121211231-0313100023201302-3121323333230211-1230211122223110-0010112302220312-2313032022222302-3113110312202011"></a>

#### `virtual_server.http.stream_profile.kind` property

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

<a id="canonical-2313322103002010-3131132321133021-2322122220300121-2333301230313310-2033221132021321-1022002030010210-2211220030230310-2031100202031213"></a>

<a id="canonical-2323002332033020-2210210121113312-3203200023212102-2132011331113202-2200022021133101-1303012131013022-0001212210131013-1021312002121023"></a>

#### `virtual_server.http.stream_profile.name` property

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

<a id="canonical-0112132222301202-1200112033033102-2321311030201032-1120322323111123-0000110320100111-1031102032000001-3221213112001131-3000133211321112"></a>

<a id="canonical-3212110321220221-1203000332103313-1321212331020110-3232323233123233-3032023311111230-1232302111323101-1211031002012020-1233223221130233"></a>

#### `virtual_server.http.stream_profile.namespace` property

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

<a id="canonical-0211303002301213-3133131212013203-2010030331002032-1321123011301201-2010003013300133-1203311212211200-1322322013130321-1311333002021333"></a>

<a id="canonical-1032120330203030-0113322232131233-1213213113333230-3003130133030022-0032031033202333-3231021202001300-2100102131130000-0002123331330112"></a>

#### `virtual_server.http.stream_profile.tenant` property

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

<a id="canonical-0113231310200301-1311112021202013-3120232101113123-3013332122201322-0123131300000300-2202313030200323-3112032231221310-2210002310210331"></a>

<a id="canonical-1231123022222113-0031021003132320-1000220233011122-0131310211121322-0333131103331021-1131033332223003-3310211133011203-3321023333213101"></a>

#### `virtual_server.http.stream_profile.uid` property

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

<a id="canonical-2222233310322030-3130301110121031-3311131123110102-2101322013110330-1013333113321210-1122223131221313-0233232232213123-3231210122131302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.tcp_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.tcp_client_profile

<a id="canonical-3011012310321011-2333310101023002-0011200302330111-2000010233110013-3320100100233230-2301020320100133-0132033001200123-0113100311321001"></a>

Type: `"list"`. Computed.

Protocol Profile (Client). Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3120112213333131-2331332300320321-2222103231010211-2103000232233101-0112200322333230-3332301003121323-3230212311001001-3013020130122031"></a>

### Direct properties for `virtual_server.http.tcp_client_profile`

<a id="canonical-1212311020300120-3223311303210313-3002012312212301-1022303013313113-3100021220010112-3222130302123310-2010120322132300-0323212023320012"></a>

#### `virtual_server.http.tcp_client_profile.kind` property

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

<a id="canonical-3332132222033033-2301003133213321-1213111300233311-1212232002231210-2110011220113012-0323003320000212-3111300220210113-0022001320313130"></a>

<a id="canonical-1103023020211030-3211011230131203-3010003222131110-3321302002302012-3010032303332300-1131022211102113-3010113032312202-3312131110133220"></a>

#### `virtual_server.http.tcp_client_profile.name` property

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

<a id="canonical-3220331132322333-3230233201232213-3303322201112300-0221213302032333-2102021112002120-2301330330101033-2131311012203003-1203120030202132"></a>

<a id="canonical-3022321221022133-1032321032103032-1311113233201232-2211301212032201-0320021121332213-3231313121321320-1130310100200102-2012133131212211"></a>

#### `virtual_server.http.tcp_client_profile.namespace` property

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

<a id="canonical-1133313331130332-2322311123201010-3320121211210312-3020312130211013-2323012013220031-3032311310222013-3022322131322301-2310200333131302"></a>

<a id="canonical-2233320331303213-1011100121223022-3310130110022222-1100032301130121-2200223111211332-3020113303123010-0230203202310022-1022332210302231"></a>

#### `virtual_server.http.tcp_client_profile.tenant` property

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

<a id="canonical-1033210200222021-1022310333120121-2330230122033303-2031002132312230-1121022000001310-2213323200123123-1211320023200301-3111011021330233"></a>

<a id="canonical-1030111102011213-1212112000312100-0332232202223323-2230120311301311-2303332223001020-1030302110120321-2222020031332122-2223331323103313"></a>

#### `virtual_server.http.tcp_client_profile.uid` property

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

<a id="canonical-3132223203012130-1010101120231200-2220012111301303-2023302011121212-2201323032110230-2202013320230313-0000110302021211-2322101122330013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.tcp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.tcp_server_profile

<a id="canonical-1031331212321311-2121010201011133-3302302120310303-1102130320220322-3003012010022312-1023032032222021-1022233221133022-2330033333122101"></a>

Type: `"list"`. Computed.

Configuration parameter for tcp server profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1100001133301233-1110111110101220-2120323030023123-1001201012330022-0222320011221022-2221132222030100-3323321001221122-3201223230000020"></a>

### Direct properties for `virtual_server.http.tcp_server_profile`

<a id="canonical-1323303322331223-1320320020221330-3021330232013222-2333003012232321-2231000113111220-2003132302303211-2232213011203200-1110021211323301"></a>

#### `virtual_server.http.tcp_server_profile.kind` property

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

<a id="canonical-3300330202103222-0100203132332121-3232202230203232-2131001120123333-3022320322211211-2233200120311222-3020100220131332-1121013213120102"></a>

<a id="canonical-0001200033223032-3312300033000100-1221231032202212-2320103223202303-2223130200111030-3021123233330201-0213030330310021-2023210332331322"></a>

#### `virtual_server.http.tcp_server_profile.name` property

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

<a id="canonical-0320310322010003-3223133131303000-1133212111311322-2200213301033300-3312123331221330-3322033122131302-3012013231211333-0103122012211312"></a>

<a id="canonical-2213312013000122-0220133231220130-3210122331022313-0112302202331132-1032323130330211-3222011110021212-1030230032130202-0131013112220000"></a>

#### `virtual_server.http.tcp_server_profile.namespace` property

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

<a id="canonical-3123302210032010-0232313203233113-1013211031331112-1333202330031020-1132302220112131-2000332303021013-0020310203102020-1232233011132113"></a>

<a id="canonical-3033102113221102-0300222323113000-0301002020123202-1002122310001312-0131122320012003-3220033103231023-0011320011213213-3120103331332311"></a>

#### `virtual_server.http.tcp_server_profile.tenant` property

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

<a id="canonical-0110223330032031-3001323230312213-0031231320211103-2320031033002210-1020300230111321-0223210330021113-0311332323313323-2021202311213012"></a>

<a id="canonical-0303031001231312-2220333100002112-1311211031331200-3231333320012021-0012312123313030-2033323300010201-3330033321122011-3013300103310302"></a>

#### `virtual_server.http.tcp_server_profile.uid` property

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

<a id="canonical-2203212220330201-2211030130032001-2311212010311313-1301011121033212-1103323033323000-2231211010013022-0310112133313231-0101330123321022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.websocket_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.websocket_client_profile

<a id="canonical-2121221311233212-2120123101131321-2110032213122311-3123023313202112-2020133303031211-1032112331010320-3131021133220001-3113233032033111"></a>

Type: `"list"`. Computed.

WebSocket Profile Client. Web-related configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0033221000013111-0111301201233210-2302100133103303-2132101302010012-0210313012322031-0201130101121103-3130013203322322-1331321112233011"></a>

### Direct properties for `virtual_server.http.websocket_client_profile`

<a id="canonical-0222222103110001-1002002221103301-3202200233023103-1310232302223113-3023320111221312-2212112111132131-1223332200310333-1113210310212102"></a>

#### `virtual_server.http.websocket_client_profile.kind` property

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

<a id="canonical-1311012201312322-3202230111100332-0020212011322213-1233102122013231-3303303210030131-0011023300132032-1111013302222222-1211322013233232"></a>

<a id="canonical-2133322120023201-1113320221012122-0032323320331223-1222020122103212-0322111123003203-2223201232203113-3303033001013110-1333221010323233"></a>

#### `virtual_server.http.websocket_client_profile.name` property

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

<a id="canonical-3202101013120120-1333202000102333-2313323001303201-1221213201032333-3300220033232132-0311211233201113-0031210302211023-2133100300023003"></a>

<a id="canonical-0202122330001013-2002311132211011-3103121301020013-3300233333232322-1310010323332103-0122020100320331-0223303222022322-2103012103321031"></a>

#### `virtual_server.http.websocket_client_profile.namespace` property

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

<a id="canonical-2313300013011033-2233210100003122-2232111002302210-0103030023021332-2112211101320221-3300233330011203-0103320023130202-0010122201310320"></a>

<a id="canonical-0112131132201301-2100113231000300-2210000003232021-0221320312302101-3020231222202302-3202101222220331-2231301133120211-1311023123333031"></a>

#### `virtual_server.http.websocket_client_profile.tenant` property

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

<a id="canonical-1001222312200023-1333330130232011-3020300133322313-2003110012100022-1213310030132202-0323303301321010-3223023103022213-2302120100230311"></a>

<a id="canonical-3110011213022011-3321322131012132-3112023020320130-0122032002021203-1330210310001201-2333132020313012-3032012032222230-2310110200031312"></a>

#### `virtual_server.http.websocket_client_profile.uid` property

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

<a id="canonical-0000013133110310-1202210210001023-2310112231201200-3221210302130202-0000111312112212-3300323001032133-2202301013123233-3013100021021302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http.websocket_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031)
- virtual_server.http.websocket_server_profile

<a id="canonical-2201322310330012-0030203000221221-0021002000020131-2001313000332122-1212102021103111-1202323011122211-0022203212120322-2021130023213312"></a>

Type: `"list"`. Computed.

WebSocket Profile Server. Web-related configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0222203033002001-0332203131022331-1303022310132132-2210201321311223-1311210221001103-3212230022200130-3112312100322001-0211233022113122"></a>

### Direct properties for `virtual_server.http.websocket_server_profile`

<a id="canonical-2113303302210122-1331100322212320-3002313112332101-0111033133110202-0013132123301012-2211230311200211-0021232021323020-3013020023022323"></a>

#### `virtual_server.http.websocket_server_profile.kind` property

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

<a id="canonical-1220330222021131-2101022210032013-2232112322030021-1003313011323211-1321333223100132-0200110333130223-2301330013113200-3023133222313032"></a>

<a id="canonical-0132021002213233-3213333112121113-1220132023000001-1000022122302032-3223230030130123-1101101320221332-2100031201313011-1103110201332203"></a>

#### `virtual_server.http.websocket_server_profile.name` property

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

<a id="canonical-1013011332013003-0132001011101221-2102320030130101-1200201121202313-3202122102202223-2323113121102300-2121020030232022-0331322121201310"></a>

<a id="canonical-2231022033020231-0302132303103001-0033233310233112-2133100220133203-1020030303023123-2220023222201100-1011200313320102-3220202322322103"></a>

#### `virtual_server.http.websocket_server_profile.namespace` property

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

<a id="canonical-1122001220100113-3013222333021331-3132110100200130-0333021013321101-0023321213021310-1222131313112121-1221212033032302-1011002133300211"></a>

<a id="canonical-2003101103030332-2011110211211003-1020311312010312-1233020321231120-3220320302001311-1310111013131310-3312203113032112-0002011031132101"></a>

#### `virtual_server.http.websocket_server_profile.tenant` property

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

<a id="canonical-3013331233210030-1111331300310222-2123230102000022-1331223232103330-1223113213322022-1021101222011022-2000012012102120-2132213313022113"></a>

<a id="canonical-0132030233022111-3313101121223331-0321030211200203-2032202221211233-3000302232001301-0121233233001212-0220111220310202-3001303220221232"></a>

#### `virtual_server.http.websocket_server_profile.uid` property

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

<a id="canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.http3

<a id="canonical-1121301310330311-2011010320101113-1233100013121330-1212222102110332-0233311301302301-3101011232131032-2201031113231103-1111233100033111"></a>

Type: `"single"`. Computed.

HTTP/3 profiles.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1201112323032011-3322301222311233-1313310020020310-1103311222011121-1012223232001310-0233020233031210-1323321300100303-0113031102312122"></a>

### Direct properties for `virtual_server.http3`

- [client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-3223130030220301-1333230231001230-1333301123031023-1103303013333031-3132303110023301-3321011223002121-2001110033011001-1333123012101131): complete subsection reference.

- [http3_profile](data-sources--application_profiles--reference--group-002.md#canonical-3321122110312100-1133113231112132-1012330200102131-3002110310120321-2030132110330203-0210132101333030-3110331101333303-3133213210012003): complete subsection reference.

- [http_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-2312123203221210-1233302333132313-0230223320333200-1133012233012001-2321022311103200-0102122130200133-0122113230123122-3230221100210022): complete subsection reference.

- [http_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-2132210212333322-0102002020003113-0313200000111313-3111331333013320-2330110023112311-3001003223132132-2202330323331111-3203001132221202): complete subsection reference.

- [quic_profile](data-sources--application_profiles--reference--group-003.md#canonical-0100013232022300-0111113200101122-0031003230203222-1131311321302302-2321000312232033-2202310303231100-0321032330201012-2211321031201111): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-0033200333002111-3332200110001031-0202212132212132-2013332121113223-0031132032002310-3123011022020321-2122311103211212-3130200323302122): complete subsection reference.

- [tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-0213332302322013-0132020312101330-1222131011201110-0332001323001021-2221123031210203-1102100013030330-3113113211011122-1120001133003210): complete subsection reference.

- [udp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-1201123121201311-2131112003333311-3011331232130201-0103122032230311-2013030111222030-1310213111333000-3320122110022121-3121022321232010): complete subsection reference.

- [udp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-3211232310201103-2321212210101101-1222112010313233-0230303023333102-0300133202221120-3311302022122102-1000103222013012-2031210202330200): complete subsection reference.

<a id="canonical-3223130030220301-1333230231001230-1333301123031023-1103303013333031-3132303110023301-3321011223002121-2001110033011001-1333123012101131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.client_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213)
- virtual_server.http3.client_ssl_profile

<a id="canonical-0231222012023300-3000220213330321-1201232010123230-2122300030212300-3323232103310130-2020222202310003-2032120013231110-1103203210211223"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2321222320012211-2023212131210301-1100011212001320-3112312123111111-2303332302332312-1320211313220000-0333130000010333-2311013003303320"></a>

### Direct properties for `virtual_server.http3.client_ssl_profile`

<a id="canonical-3023210330100002-0301132022221320-1210111112032033-2323010021100302-1221111220121021-2013012310011322-0132212333030201-3023331310002100"></a>

#### `virtual_server.http3.client_ssl_profile.kind` property

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

<a id="canonical-3102011121030132-2022113223311310-2021312011221123-1321102032013232-0001133212123102-3231120011011211-1203000133011022-3203132301322111"></a>

<a id="canonical-1210221323103320-2100111200102111-3012221302022110-1330012023033000-0022012130212003-0113011013201320-2333221323320301-1201023232203302"></a>

#### `virtual_server.http3.client_ssl_profile.name` property

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

<a id="canonical-2200103300303200-0032203021033103-1220023131000030-0113130133232110-2230031210213110-3222010131020210-2010221031310013-3110202013331212"></a>

<a id="canonical-1133100303330123-3121121132323000-0123033201201030-2123230212033123-1302220100303133-2320330233112221-0310211323300230-2012203033131321"></a>

#### `virtual_server.http3.client_ssl_profile.namespace` property

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

<a id="canonical-3013210231013221-2032101203312023-2020332233133332-1031301303312231-0200223320202002-3212100031100202-0220233232230203-1102132230232333"></a>

<a id="canonical-2232223232210321-1131001201023120-0310103213003021-2212200300303313-1012030023031302-1331011013323101-1312332122110221-2030313232210113"></a>

#### `virtual_server.http3.client_ssl_profile.tenant` property

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

<a id="canonical-0211001030212123-1133333311333100-3201021301100023-3010002033101321-1111311221012113-0000132323310323-1223212133233113-3313201313310101"></a>

<a id="canonical-3303202231030300-0100302011213133-3223233130101313-2133320323033303-2012010220200321-2131121111331001-2222032302020000-3322221212011220"></a>

#### `virtual_server.http3.client_ssl_profile.uid` property

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

<a id="canonical-3321122110312100-1133113231112132-1012330200102131-3002110310120321-2030132110330203-0210132101333030-3110331101333303-3133213210012003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.http3_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213)
- virtual_server.http3.http3_profile

<a id="canonical-3220112303100231-3223101201022303-2333230121221013-3100012311121031-2123220033020201-2112013333213311-2220331222111132-1231020322012332"></a>

Type: `"list"`. Computed.

Configuration parameter for http3 profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0300012021122311-0121313111321231-2010003003323120-2310111130001302-1313023202302313-2303023211020020-2013033111103021-1213310123131330"></a>

### Direct properties for `virtual_server.http3.http3_profile`

<a id="canonical-2212211331103123-1222030223132231-3212202301323310-0333202003032001-1330111230323301-0111003313132013-2322132033200203-1230321003122103"></a>

#### `virtual_server.http3.http3_profile.kind` property

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

<a id="canonical-1320331030103202-3222003032300311-1023223332000212-2121000221320301-2023213301210031-3112302333020303-3222130021023113-0003131122012022"></a>

<a id="canonical-1110122331023103-3003211303011101-2130201322200120-1023133310022010-1011213102101000-2111300333233002-3101023103331311-0310011330032003"></a>

#### `virtual_server.http3.http3_profile.name` property

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

<a id="canonical-1220031310231311-1323022030121301-3031000223321303-0221013200110131-2221232021012213-2231301102231013-1003022122312310-0302200000312013"></a>

<a id="canonical-0323130122230032-0122003333101112-0222001023212333-3330033100201023-0131333030331230-0010000121000023-0031133122213111-2123130010212210"></a>

#### `virtual_server.http3.http3_profile.namespace` property

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

<a id="canonical-0000311231313331-3002032321122220-1031230122133331-1233111201131021-3033112200202330-2030303232000313-3321033111212131-3322320000100201"></a>

<a id="canonical-1201200112023303-2103032322133322-3230101201300211-3201310011121121-0311123021222211-1230112103033010-0103322212333011-2112302231321123"></a>

#### `virtual_server.http3.http3_profile.tenant` property

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

<a id="canonical-1020010203221023-0322101300123233-2303221333012232-2320233011032131-3101122132123030-3123211213322113-0111130210113032-2110002311121321"></a>

<a id="canonical-3201002022301230-1130331223132102-2012013210233101-2130312101323101-2010210020010230-3110223231032311-3110330020113300-0023232211311021"></a>

#### `virtual_server.http3.http3_profile.uid` property

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

<a id="canonical-2312123203221210-1233302333132313-0230223320333200-1133012233012001-2321022311103200-0102122130200133-0122113230123122-3230221100210022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.http_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213)
- virtual_server.http3.http_client_profile

<a id="canonical-0111303033120332-3212222230103002-1130222323000111-1122332302102222-1201303031332111-0001211231200012-0101300313323110-2011320002122311"></a>

Type: `"list"`. Computed.

HTTP Profile (Client). Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0222111021000110-1113012233222103-3330313000310000-0033121312210020-3122102111023333-1002203030003132-1111010311133031-1213001012300100"></a>

### Direct properties for `virtual_server.http3.http_client_profile`

<a id="canonical-2200202211102132-3330231311121203-0031101230223300-1112023020223203-3120232213213200-3022331010033130-0131321303031111-2210303121203000"></a>

#### `virtual_server.http3.http_client_profile.kind` property

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

<a id="canonical-1221301021233131-2123220302230000-1003133311333212-0230013231000302-2032022023213211-2302201121132103-0232112103303030-1200301302220303"></a>

<a id="canonical-3233033120121102-1013132132100223-2130112110320303-3121112212110113-1212223310302102-1130012322031203-1212002312021011-2330133300021303"></a>

#### `virtual_server.http3.http_client_profile.name` property

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

<a id="canonical-1320213213032310-0031223033202201-2030311113223331-0212230001330133-3033003332220010-0003300001302120-0023021022000313-3233122021020312"></a>

<a id="canonical-2232102033200303-2011233212311303-3210102010110031-3103012210010132-2031322303111303-2130022022301321-0003211313031131-3312320112102313"></a>

#### `virtual_server.http3.http_client_profile.namespace` property

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

<a id="canonical-2212113133011312-2110111302222332-2120213221033022-2031203211000323-1023113003030331-0322300321332121-0021000232113332-3021232022231002"></a>

<a id="canonical-0310332202330210-3031011311300313-3103111313301003-3310110121010302-1103023203103211-3222002321320302-3122202230102230-3312002033120300"></a>

#### `virtual_server.http3.http_client_profile.tenant` property

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

<a id="canonical-1232303301322000-1231030012103001-1131120020012023-0030100113120313-2230303110102000-2020220231300302-0103300133302130-1211122220000310"></a>
