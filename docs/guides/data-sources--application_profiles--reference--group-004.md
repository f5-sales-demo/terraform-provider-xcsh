---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-1022331310002101-0221122112312111-2032231100313013-1022031020210202-3312201210321013-1211121020302023-3201110001133330-3003133303132200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.nat64.nat64_disable` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-3320001111230010-0321322131120330-1030210203213223-1310133111230310-1101021300130220-2333001002233033-0302123321201121-2333013113030122)
- virtual_server.nat64.nat64_disable

<a id="canonical-3330021000232332-3112120221301132-1202112021311233-1002102220223223-1323103112322321-3020330031020301-1010001112301333-3210122112321303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for nat64 disable.

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

<a id="canonical-2013103112220012-3012000211330010-0010213022332111-0130021120033031-2221020031113033-2333133200123101-0203122230130021-2130102133201003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.nat64.nat64_enable` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-3320001111230010-0321322131120330-1030210203213223-1310133111230310-1101021300130220-2333001002233033-0302123321201121-2333013113030122)
- virtual_server.nat64.nat64_enable

<a id="canonical-2323330300333113-0111210302021122-3222102210220312-3120003231033321-1223002012201122-2212212232011330-0022023201212313-2213003331011323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for nat64 enable.

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

<a id="canonical-1033032133123230-2203033131231333-2313120131311332-1121301201221211-1321320320100030-3030133220301332-1312323030311221-3002302202301220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.port_translation` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.port_translation

<a id="canonical-0230012221221331-3021110120120210-2301323100032210-0112002110033330-3020100332131003-3320320221322132-3320012303020331-0002233212311013"></a>

Type: `"single"`. Computed.

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance
connections to any service. The default is enabled.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_translation_choice": "[\"port_translation_disable\",\"port_translation_enable\"]"
}
```

<a id="canonical-1112213122012301-2030313331310310-1202112101230331-3213133201121031-2323131123132022-3020020213233300-1023310003202031-0132023003010220"></a>

### Direct properties for `virtual_server.port_translation`

- [port_translation_disable](data-sources--application_profiles--reference--group-004.md#canonical-0301002120332312-1203132031122302-3232230312232323-3302200330030232-1012113113033111-3222222100012201-1023031023301300-0312121120332223): complete subsection reference.

- [port_translation_enable](data-sources--application_profiles--reference--group-004.md#canonical-2133331011212010-0333112201230001-3033230312321133-1100122203213031-2132221222002332-0120313203223322-2221220003310211-3231122203030101): complete subsection reference.

<a id="canonical-0301002120332312-1203132031122302-3232230312232323-3302200330030232-1012113113033111-3222222100012201-1023031023301300-0312121120332223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.port_translation.port_translation_disable` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.port_translation](data-sources--application_profiles--reference--group-004.md#canonical-1033032133123230-2203033131231333-2313120131311332-1121301201221211-1321320320100030-3030133220301332-1312323030311221-3002302202301220)
- virtual_server.port_translation.port_translation_disable

<a id="canonical-1223333011230223-3003313223102211-3213332123113011-2011113321311333-1131201300001232-3222010100131121-3323120023222032-1210311122222011"></a>

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

<a id="canonical-2133331011212010-0333112201230001-3033230312321133-1100122203213031-2132221222002332-0120313203223322-2221220003310211-3231122203030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.port_translation.port_translation_enable` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.port_translation](data-sources--application_profiles--reference--group-004.md#canonical-1033032133123230-2203033131231333-2313120131311332-1121301201221211-1321320320100030-3030133220301332-1312323030311221-3002302202301220)
- virtual_server.port_translation.port_translation_enable

<a id="canonical-1123021020332223-3003002223232331-2000213010302201-0232230211330103-1311212121233320-2022302310202033-1333022301001012-0122310131120012"></a>

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

<a id="canonical-3232320301212213-2201233011223222-1121012211011023-1212323111231033-1300332330213333-1302020033323222-1223102122223230-1103311321323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.request_logging_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.request_logging_profile

<a id="canonical-1031110300301322-3011012111331131-3320322013310021-1331332303130013-0333231103331122-3012333020301221-1133212222013310-2220232130012002"></a>

Type: `"list"`. Computed.

Configuration parameter for request logging profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1303103030002221-3003031100220013-2130033012330213-1012201013020210-3333100231233011-0022133331213300-2232012231231232-3202031221133022"></a>

### Direct properties for `virtual_server.request_logging_profile`

<a id="canonical-3111011233010123-1322230210312302-3110220021311312-0123010212232222-1221021223020311-2233320201303112-3213221001323133-2222331001120331"></a>

#### `virtual_server.request_logging_profile.kind` property

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

<a id="canonical-0032131021213201-0032210012001322-1022033201203132-0221223233220321-1230311221001013-3311132303202130-1002223322112110-0211321030200212"></a>

<a id="canonical-1312211000000102-2221323220003130-3113312211031011-1133333232123013-1010211313203333-0003213210233300-2132112101122333-2100221222002000"></a>

#### `virtual_server.request_logging_profile.name` property

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

<a id="canonical-2333231231202222-3000120200101013-2333122312302300-1010300001320031-1312000333330232-3020130310023312-2300030113133222-1332021211122210"></a>

<a id="canonical-3112212121110112-2012313122300300-1113002220222020-0311223332330110-1111103121111132-2113313232110021-1222013311231121-1003333011233030"></a>

#### `virtual_server.request_logging_profile.namespace` property

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
  }
}
```

<a id="canonical-2321321001100300-2222002201310132-0312100012233022-1110131130102021-1202103112220223-1310120111103121-1312203321021203-0323333131013303"></a>

<a id="canonical-1120332101230301-2231132301200121-0030320001123010-3220101331223221-1310223112110222-1233203221223001-0123101311132221-2101220111001030"></a>

#### `virtual_server.request_logging_profile.tenant` property

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

<a id="canonical-1211021330212112-0113010233101110-3213130122212220-3122221003303033-3220003021232020-0230220003222200-0120231211301201-3100113311233200"></a>

<a id="canonical-1000200211132020-2133122022320230-0100032131032230-3300010310023131-0220210321230210-3121200221012312-3113030233031230-1202111233132131"></a>

#### `virtual_server.request_logging_profile.uid` property

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

<a id="canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.source_port` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.source_port

<a id="canonical-2231031203231101-1320022022331310-2302331331320012-0013323110030013-0322230203211311-0231131003221110-2123031233033330-3020101331100032"></a>

Type: `"single"`. Computed.

Specifies whether the system preserves the source port of the connection. The default is Preserve.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-source_port_choice": "[\"source_port_change\",\"source_port_preserve\",\"source_port_preserve_strict\"]"
}
```

<a id="canonical-1030013022030230-1111202330120203-0213102313311010-0033113230031202-2230022213311211-1211003121021310-2022201211000113-1000120033010202"></a>

### Direct properties for `virtual_server.source_port`

- [source_port_change](data-sources--application_profiles--reference--group-004.md#canonical-1323123103122011-2030010212212212-3032020203020021-0312023101310131-0132200230200033-2130313320131300-3113332320113233-2220123122222023): complete subsection reference.

- [source_port_preserve](data-sources--application_profiles--reference--group-004.md#canonical-0330311022021121-0102221001112332-1012113013030313-1020320310333013-1123321003312213-3233332220100202-0120032020223221-0221012001023101): complete subsection reference.

- [source_port_preserve_strict](data-sources--application_profiles--reference--group-004.md#canonical-1122102321003221-3211213201020003-0010031201013323-2313110220301200-1300131012122331-0030322133003203-0011020313220021-0302002233210021): complete subsection reference.

<a id="canonical-1323123103122011-2030010212212212-3032020203020021-0312023101310131-0132200230200033-2130313320131300-3113332320113233-2220123122222023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.source_port.source_port_change` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.source_port](data-sources--application_profiles--reference--group-004.md#canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203)
- virtual_server.source_port.source_port_change

<a id="canonical-1033230100021110-1232321310030300-0013311113121012-3131323201213223-3131121020222213-2123022303233030-0200332113031001-0112011010112101"></a>

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

<a id="canonical-0330311022021121-0102221001112332-1012113013030313-1020320310333013-1123321003312213-3233332220100202-0120032020223221-0221012001023101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.source_port.source_port_preserve` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.source_port](data-sources--application_profiles--reference--group-004.md#canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203)
- virtual_server.source_port.source_port_preserve

<a id="canonical-3122210132111330-3213033313322121-2230323132323233-0202202102221001-2303103302013330-1203220001233003-0133322023132333-3231231023120323"></a>

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

<a id="canonical-1122102321003221-3211213201020003-0010031201013323-2313110220301200-1300131012122331-0030322133003203-0011020313220021-0302002233210021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.source_port.source_port_preserve_strict` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.source_port](data-sources--application_profiles--reference--group-004.md#canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203)
- virtual_server.source_port.source_port_preserve_strict

<a id="canonical-2103231131103021-2223110201303233-3123113133113121-2130101121221022-2030110102132300-0102210010213033-2201112112020000-0330100113133120"></a>

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

<a id="canonical-0230033123202003-3221010301120301-1321302111033133-3300013133310233-0333313030302000-0200200313230110-2223320101301232-1322122301001122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.statistics_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.statistics_profile

<a id="canonical-2030120322312202-1311203300012321-0321320031013201-3233002122323031-2033010022232303-2110113120103221-1303322002232120-0232013212222013"></a>

Type: `"list"`. Computed.

Configuration parameter for statistics profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1203011212210010-2313010312133211-2202021011203003-2320110013013112-2120031123003122-0130321200111203-0232010110310232-3103031223201022"></a>

### Direct properties for `virtual_server.statistics_profile`

<a id="canonical-2302010102032211-2101220322000120-1122102130000030-0331230231012103-2033003332311120-1211011002101230-2122113311020032-3233202230133301"></a>

#### `virtual_server.statistics_profile.kind` property

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

<a id="canonical-1210213022223212-3133101111210020-3322303133122300-0123123013113330-2003233011230003-3122232322221131-1113313133211010-2201303221233131"></a>

<a id="canonical-3210110331310221-1032203031003113-2332132321001123-3010020231333013-0130021201302230-3022010103321211-2220100333301232-3033110322012131"></a>

#### `virtual_server.statistics_profile.name` property

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

<a id="canonical-3233300113322310-1133113133232223-1230022323013133-1211132212130013-2113321221200310-3121033032102200-0220012013130010-0221021211333210"></a>

<a id="canonical-2233323200230031-1210000103202131-0130330231202303-3031331000330222-0111110301123303-3323233122110331-3200322132001121-0010301013131020"></a>

#### `virtual_server.statistics_profile.namespace` property

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
  }
}
```

<a id="canonical-3200032003300132-3303302103202001-2000030303131112-3100100220110121-3330100231021011-2333200103022023-0302221203020032-2031313331032301"></a>

<a id="canonical-0021301231322122-0032301323112330-2233112331000231-2223123011103132-3122331033210122-3232103101320333-1112323210333012-1112120222230332"></a>

#### `virtual_server.statistics_profile.tenant` property

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

<a id="canonical-3101021230321013-1312011302130111-3030332132013102-0111212322122021-3203222232012013-1112302322213300-3330331220232323-2233030123302330"></a>

<a id="canonical-3033221013013321-2220323103030311-0203321102112101-1100033002231301-1303203201001201-1130301111332103-2221333302001120-1202221321333310"></a>

#### `virtual_server.statistics_profile.uid` property

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

<a id="canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.tcp

<a id="canonical-3332310323031233-0313321330000232-3300010012123032-3321003311001223-1002000031011012-1221013001120331-2103101021011301-1010110123323100"></a>

Type: `"single"`. Computed.

TCP profiles.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3200332232302003-0202220332200232-3002022113000120-0303023020010233-1333321303012321-0221012013330232-1010323021102130-3223110122202211"></a>

### Direct properties for `virtual_server.tcp`

- [client_ssl_profile](data-sources--application_profiles--reference--group-004.md#canonical-0310001233311003-1033213103022230-2312022320203123-2031320231001100-0123321231030330-1010202003122200-1223230132221203-3220003013002122): complete subsection reference.

- [ocsp_profile](data-sources--application_profiles--reference--group-004.md#canonical-2213210301312212-3113123320032331-3111321222210312-2111022110101033-2213330323231020-3013133312232221-2021112223011032-0302110312212033): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-004.md#canonical-3132332111033112-3021230331103102-0000203200320323-3100311302321003-3203210332032103-3003230230222013-3033030213121211-1000123332303001): complete subsection reference.

- [tcp_client_profile](data-sources--application_profiles--reference--group-004.md#canonical-1121221121121022-3213323002333121-1322111332101300-0322022230313330-3212022232023001-2113120221231021-2013111100331123-0231030200223111): complete subsection reference.

- [tcp_server_profile](data-sources--application_profiles--reference--group-004.md#canonical-1231122033130320-3112100322132110-1122202030232302-3121120110303221-2321120012312233-0021322102110131-1211212012311111-2031032102321131): complete subsection reference.

<a id="canonical-0310001233311003-1033213103022230-2312022320203123-2031320231001100-0123321231030330-1010202003122200-1223230132221203-3220003013002122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp.client_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-004.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- virtual_server.tcp.client_ssl_profile

<a id="canonical-1223323031103032-1333301010223312-0323313113233102-2211030313223003-2112002133300232-3131120003201031-3300020231132200-2200333311102021"></a>

Type: `"list"`. Computed.

Client SSL Profile. Client-side configuration

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

<a id="canonical-2333233011120133-3022112332301232-3013102103020320-0012210320110212-3020031123203203-1033001120311110-2122311003213233-2020000011111322"></a>

### Direct properties for `virtual_server.tcp.client_ssl_profile`

<a id="canonical-1220010130112121-3220233320100101-3013123332101221-3221133232033011-0022001132301033-2210111002022001-1312132313210102-3332110020230311"></a>

#### `virtual_server.tcp.client_ssl_profile.kind` property

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

<a id="canonical-1201232331100312-2231133323321300-1101222002301001-0300120320023002-3312310012202030-2312332002222021-2213223111232022-1122310001101202"></a>

<a id="canonical-2123121121212131-2201303231210031-0113120031231101-3022220022311201-2312001332310212-2101312331103033-1000203232000103-2030022112303123"></a>

#### `virtual_server.tcp.client_ssl_profile.name` property

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

<a id="canonical-0122021320113200-3302220020122013-3130330102311221-3032220300022223-2333010320122220-2033103132200010-0323310333210302-2213121032322330"></a>

<a id="canonical-3022203011022012-0123301232003220-1121130100320320-2133003212120332-0121022001233012-1033011203101101-1020301033012130-0113210311131130"></a>

#### `virtual_server.tcp.client_ssl_profile.namespace` property

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
  }
}
```

<a id="canonical-1012013322002021-2211211113222003-2032221301132030-3223122332033031-3311002131012113-1300301110330220-1032312100101300-0001322103310122"></a>

<a id="canonical-3010331202123211-1131223301221101-2100032333130212-0321011202023330-2223232311310330-0032110122021203-0133302220222113-0022300333233010"></a>

#### `virtual_server.tcp.client_ssl_profile.tenant` property

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

<a id="canonical-2121012201112312-0002001312031121-3201331313001310-2031032313001331-2023333220313133-2013320223201000-2101002012103321-0323010233232122"></a>

<a id="canonical-0102132332212001-0311333212020230-1333321020202223-1311202111121221-1231200030000133-3303110312100033-3012203202320012-2301002123301000"></a>

#### `virtual_server.tcp.client_ssl_profile.uid` property

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

<a id="canonical-2213210301312212-3113123320032331-3111321222210312-2111022110101033-2213330323231020-3013133312232221-2021112223011032-0302110312212033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp.ocsp_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-004.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- virtual_server.tcp.ocsp_profile

<a id="canonical-3222113202322322-0331222010032111-0231200332330131-1120010211000200-1320103023320021-3020032311122013-3012321300232301-1212233311202030"></a>

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1320013023203231-2210122121133012-1133020100223022-1210333320310310-3001102021322220-3001103120102032-3001100212301133-3203332330013110"></a>

### Direct properties for `virtual_server.tcp.ocsp_profile`

<a id="canonical-2133103031213311-1023200220030223-0300213312300320-3033203323222030-3003022112303203-3322032133130302-0203212000000210-3210320201220332"></a>

#### `virtual_server.tcp.ocsp_profile.kind` property

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

<a id="canonical-1303302102302030-0100000232300311-2210102233131213-3023320231311100-0230322133121211-0012123230110220-1211021233111031-3332133132300322"></a>

<a id="canonical-1000121100020120-3330302102031003-0211310200123112-1020002223003320-2220332231001233-1111003003131333-2223320121111012-1232103212213203"></a>

#### `virtual_server.tcp.ocsp_profile.name` property

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

<a id="canonical-3130230020320011-3103103201100321-1203031032010003-0133021210001121-0021231011103331-3233220302011033-3212223022020031-2202112331033011"></a>

<a id="canonical-1120000001102131-0230100022101302-1201103222130223-1001112102212330-0310320333222011-3031232113111130-3132330013232222-3223132303211123"></a>

#### `virtual_server.tcp.ocsp_profile.namespace` property

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
  }
}
```

<a id="canonical-0322021303013323-0202320023211102-3100101303010000-2210122222101322-2322233231000131-2303121110023200-1220332322133203-1222233023231233"></a>

<a id="canonical-2021120221212021-0101310120221110-2031312112112021-0103301011103203-2312322011030201-0003021310331321-1112313110030020-3001220120033011"></a>

#### `virtual_server.tcp.ocsp_profile.tenant` property

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

<a id="canonical-3233130311023002-2222211223202010-1211231210210222-1320312111330103-3231201020300111-3033322213122232-0100232132220223-1321323302211202"></a>

<a id="canonical-2032203110131313-0032323330300322-1110303221301200-3033030002321000-1032012002011230-2100103213002200-0303332203130312-0120302130003123"></a>

#### `virtual_server.tcp.ocsp_profile.uid` property

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

<a id="canonical-3132332111033112-3021230331103102-0000203200320323-3100311302321003-3203210332032103-3003230230222013-3033030213121211-1000123332303001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp.server_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-004.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- virtual_server.tcp.server_ssl_profile

<a id="canonical-1301210030113010-3001031000022023-1203202300132011-2320223221023320-0201113030211233-1113112203320131-0321103311130200-2313302302001132"></a>

Type: `"list"`. Computed.

Configuration parameter for server SSL profile.

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

<a id="canonical-1113000111313130-3303133222102223-0323311300210010-1301133022120032-1131100233013112-0032013112302212-0331022010210030-3331202020203111"></a>

### Direct properties for `virtual_server.tcp.server_ssl_profile`

<a id="canonical-3123220302001003-0221110312110212-3010023023031232-0100021213122233-0231101221102032-3010020223123303-3201231220101332-3233230022333100"></a>

#### `virtual_server.tcp.server_ssl_profile.kind` property

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

<a id="canonical-3220023021330010-1022331102300302-1130310302211113-2203331201322310-0233211102323232-1103213012203022-2011230332230113-0132221121220122"></a>

<a id="canonical-2100232120002203-3113211212122031-0100223133133001-2103213223110202-1013101223231302-2011010020312023-3301203123000123-2222330201120322"></a>

#### `virtual_server.tcp.server_ssl_profile.name` property

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

<a id="canonical-3111223102012123-0203023033221220-0332303002133231-2031010210000102-1020011001012231-1102000000113011-3002330230133202-1213201203331323"></a>

<a id="canonical-0332212233223102-0130032323012002-0112011121100220-2330302213323213-2203023213321103-0230310233222002-3010010130202211-1320100121003330"></a>

#### `virtual_server.tcp.server_ssl_profile.namespace` property

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
  }
}
```

<a id="canonical-2022110203301233-2233230031101202-1300223210332300-3321000123200220-1131320332213333-2202311133113200-2133233021110212-0032013030303331"></a>

<a id="canonical-1112120111323233-3022312233310010-3010133011321311-2203021312030112-0302003032301201-0122120231110231-2021223221200303-2030211111132331"></a>

#### `virtual_server.tcp.server_ssl_profile.tenant` property

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

<a id="canonical-0103200032003121-2303201330033021-0032320131330211-3131000310203303-3233202112120102-1231321133321301-3102221230101311-1003122022310022"></a>

<a id="canonical-2013102031300303-2221321202330131-2020231121231211-3013332302000100-3201303003330022-3303003232233011-0323032222321003-2000231020223331"></a>

#### `virtual_server.tcp.server_ssl_profile.uid` property

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

<a id="canonical-1121221121121022-3213323002333121-1322111332101300-0322022230313330-3212022232023001-2113120221231021-2013111100331123-0231030200223111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp.tcp_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-004.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- virtual_server.tcp.tcp_client_profile

<a id="canonical-0233130021010220-1121023203223312-3032230023001021-2100113110302023-2010233012010211-1320211021101331-2021110312323310-1030121232022110"></a>

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1210023332310012-2320320030000120-2022232132312122-3331331211000022-3022322223203023-3121321031122303-0332013202220133-2221323112200231"></a>

### Direct properties for `virtual_server.tcp.tcp_client_profile`

<a id="canonical-3032220312231213-2300232121210230-2212123331131322-1031321232021130-3023223320313113-2021331121331212-1331312010331211-2032022210121023"></a>

#### `virtual_server.tcp.tcp_client_profile.kind` property

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

<a id="canonical-0303213032221112-1100121300200200-1320222322310231-3123002211221131-1230311310232123-3131301212131303-1102221213030213-2031300131331031"></a>

<a id="canonical-2002203012221300-1321222022122023-3102230201333021-2210012130320013-1300302002213110-3221031322212202-0323202122111103-1212131231331031"></a>

#### `virtual_server.tcp.tcp_client_profile.name` property

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

<a id="canonical-2001200320213233-1120030103320132-1032112022111031-3211100010100330-2302333212033031-2133231220001010-2233111313032112-3212032333310201"></a>

<a id="canonical-0113122302112201-3313202013301001-1101202320002200-3322201222112233-0033132320213220-0302310321301123-3312100301311231-3132133020203102"></a>

#### `virtual_server.tcp.tcp_client_profile.namespace` property

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
  }
}
```

<a id="canonical-1212100302003313-0331210323332030-1201303123012230-0030221332203201-3001113132301312-3101032331100222-3233223310013113-1110332333101122"></a>

<a id="canonical-0320300122122031-0301003211223103-1010021102310201-1101321311322223-2123010110001202-0301032100132033-1212232313221032-0120032232000222"></a>

#### `virtual_server.tcp.tcp_client_profile.tenant` property

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

<a id="canonical-0233221223033011-1212023121031111-0321033132333302-0003331011310333-1113023121012311-1020210031000302-1221023202000030-0220012210312210"></a>

<a id="canonical-0022000203301301-2310012313332210-0330101003100332-3033330321020121-2201031330203011-0231213030010222-1312131100020320-3313323212030200"></a>

#### `virtual_server.tcp.tcp_client_profile.uid` property

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

<a id="canonical-1231122033130320-3112100322132110-1122202030232302-3121120110303221-2321120012312233-0021322102110131-1211212012311111-2031032102321131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp.tcp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.tcp](data-sources--application_profiles--reference--group-004.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101)
- virtual_server.tcp.tcp_server_profile

<a id="canonical-2223101012002002-0211221222230013-2203113222222221-2000010023113232-0022101110310031-3312130211212133-3101130033321322-2112213130333221"></a>

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1313022020310122-2122312111210020-0022303233213203-2303221010302213-1230112133022333-0331101113012203-0133032121031200-1320122202323310"></a>

### Direct properties for `virtual_server.tcp.tcp_server_profile`

<a id="canonical-2331002323321112-1113033321031011-1321232120031023-1023222002103233-2021021130003312-0301112010210103-0120110110023201-2002121230011020"></a>

#### `virtual_server.tcp.tcp_server_profile.kind` property

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

<a id="canonical-2003031013322331-0100100330122133-0303211313113032-2110233133231212-3012313200011313-2031020020322202-1020332111121023-1223121200213100"></a>

<a id="canonical-0103010000310203-2100032001231000-2102311101212301-0303230130233302-2231223030000332-0012320232112131-3100300003213111-0233111011112230"></a>

#### `virtual_server.tcp.tcp_server_profile.name` property

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

<a id="canonical-3322220201010210-0203220112133211-1120123201023000-0231313101033100-1332111013202002-0310230131120331-2321312000120302-2102230102110023"></a>

<a id="canonical-2110330032020220-3213113011013131-0321231333300213-3133022002111121-2111131322120023-3232210030320001-2001013113222110-3313133130101230"></a>

#### `virtual_server.tcp.tcp_server_profile.namespace` property

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
  }
}
```

<a id="canonical-0321123000011210-2010112330021131-1331101313300303-2310210120100000-1323223202001223-1130031103303331-1222032231310200-2101120201110231"></a>

<a id="canonical-2013021302131321-3210010103330022-2010320221321210-3333101023221123-2331033030332220-0111202321312113-3122002222001033-3131122222200033"></a>

#### `virtual_server.tcp.tcp_server_profile.tenant` property

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

<a id="canonical-3112202321320130-0302202312113323-0330303200001100-3323010223012021-2212010221311202-3223032212230232-1231023121312230-0110130011003100"></a>

<a id="canonical-2301302113133312-1311033221033102-2223331000333001-3111132020133300-3122233022301120-3102020232102211-3223031330131103-2001210100121303"></a>

#### `virtual_server.tcp.tcp_server_profile.uid` property

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

<a id="canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.udp` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.udp

<a id="canonical-3230102230310120-1010023200321221-3133023233101012-0120203133222230-3233210233210103-1332031102321121-2231002312110123-2021321301012232"></a>

Type: `"single"`. Computed.

UDP profiles.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1123102213120331-0211003211120013-1223210122230113-0011130230302213-2313212302033223-2023030002103203-0321313210132103-2321301330330330"></a>

### Direct properties for `virtual_server.udp`

- [client_ssl_profile](data-sources--application_profiles--reference--group-004.md#canonical-1130220011130010-2223332122330201-2233130011332112-1222012131322323-2032233002322102-1330122332023010-2310031123300223-3001301201213031): complete subsection reference.

- [server_ssl_profile](data-sources--application_profiles--reference--group-004.md#canonical-1233221132021202-2231130020101220-2012310103223310-1012001123132022-2002120023221313-0323030301120010-0302122033300222-0013002113111320): complete subsection reference.

- [udp_client_profile](data-sources--application_profiles--reference--group-004.md#canonical-1333123302300223-1003212003320120-3011131021320211-3121211002322222-0101202302010313-3232321330230032-2023323301311101-0003010123121122): complete subsection reference.

- [udp_server_profile](data-sources--application_profiles--reference--group-004.md#canonical-0223131020311101-3120313230320201-0310233233131121-2012011133331201-0212100200033333-0303313300113123-1231122103120022-2211230303302203): complete subsection reference.

<a id="canonical-1130220011130010-2223332122330201-2233130011332112-1222012131322323-2032233002322102-1330122332023010-2310031123300223-3001301201213031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.udp.client_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.udp](data-sources--application_profiles--reference--group-004.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- virtual_server.udp.client_ssl_profile

<a id="canonical-2223111031211133-0101333002323221-1330302200031101-2132112223123222-2201122023213030-1020321313300011-1010133331202031-3230300310130112"></a>

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2301002120223123-1131111310032203-0203121333111332-3011100332231011-1012031122121101-2112230312111001-1121212212322131-2320011131021133"></a>

### Direct properties for `virtual_server.udp.client_ssl_profile`

<a id="canonical-0322301012103310-2211210202321302-2220100202033221-3123231121221030-2222333013100200-2223332112112210-1131021123131111-3231210003232133"></a>

#### `virtual_server.udp.client_ssl_profile.kind` property

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

<a id="canonical-0132103230323301-2310210101023022-3221100322330230-2020223323320212-0212212311121321-0113221020133331-2303003101010200-2122303332230211"></a>

<a id="canonical-0222211331100211-0111032221102221-0213021112131032-2322011133021230-2310013311122200-3000213001131032-0230230102022320-3030322030323211"></a>

#### `virtual_server.udp.client_ssl_profile.name` property

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

<a id="canonical-1210103233123230-0011200330100300-3122332312110310-0233010331020013-1122002212112101-1212203332320202-1321202211323003-1102301221013010"></a>

<a id="canonical-3200212111130220-3223231002011210-2010213022133131-0330223300020211-2311000200231001-1230131322232230-1331202311113332-0023113033000200"></a>

#### `virtual_server.udp.client_ssl_profile.namespace` property

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
  }
}
```

<a id="canonical-0211311122032003-3222130111113002-3101000323101010-2021113023000331-2111000102130221-2111120203222011-1100013033011023-0203132122212300"></a>

<a id="canonical-0102320313100000-2310023020033120-1302223323202103-1021113221031320-3100312210300122-3012202310000003-0031123030023101-3120003320323010"></a>

#### `virtual_server.udp.client_ssl_profile.tenant` property

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

<a id="canonical-2211100301020302-2320012111132033-1201303120022021-1220213100122220-2012021300310132-0233030333313032-2011213112313003-0231302123123110"></a>

<a id="canonical-0021203102322203-3011332213312113-0211320022033223-1122023333333031-3121132102320021-3302331022133133-1103011222310223-2213011002310313"></a>

#### `virtual_server.udp.client_ssl_profile.uid` property

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

<a id="canonical-1233221132021202-2231130020101220-2012310103223310-1012001123132022-2002120023221313-0323030301120010-0302122033300222-0013002113111320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.udp.server_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.udp](data-sources--application_profiles--reference--group-004.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- virtual_server.udp.server_ssl_profile

<a id="canonical-0101033012301333-2330013230023300-1030200021120003-3111233103113030-1102102100332031-2131121031322222-2213131312112330-2301020232113100"></a>

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2032222111023030-0131112222003030-1011223321001313-3003230222200331-2113112312301210-0231201213103332-0231121023122112-3112031232322201"></a>

### Direct properties for `virtual_server.udp.server_ssl_profile`

<a id="canonical-2332323020202202-2301010013113311-3201203231333312-0131221313130303-2301001121121100-0303113122120330-0132303133110002-2020321130031112"></a>

#### `virtual_server.udp.server_ssl_profile.kind` property

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

<a id="canonical-2221303121101223-0133032033012012-3313000223203101-3331221001133121-2010132220010230-0231313311333303-1323002202323311-1301012013102131"></a>

<a id="canonical-2131022310222203-3232022200111110-3122331110231232-1311233121301123-3203200200203312-0321132303113131-1130020010011203-2020131033133303"></a>

#### `virtual_server.udp.server_ssl_profile.name` property

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

<a id="canonical-1100203111120031-2120323223132010-2301303001222002-3012020030133313-1213312200202111-0210120101322223-2301011312021011-2000111012011011"></a>

<a id="canonical-0032210201011002-3312100100012123-0330300123100001-0130221120230012-0120210030101331-3320221330210020-2331210002111323-0233233211020033"></a>

#### `virtual_server.udp.server_ssl_profile.namespace` property

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
  }
}
```

<a id="canonical-1223122120321102-2000203333123003-1322122023312212-3000031011123100-3212313310230210-2312232300021021-2013231031020232-1302021022310010"></a>

<a id="canonical-0022021230201103-2333133132012010-1133322300321313-2322333330310202-1030220313231021-1012233030121012-2033003130312100-3313200100112023"></a>

#### `virtual_server.udp.server_ssl_profile.tenant` property

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

<a id="canonical-2231130232100022-3123001103220031-2033212310210310-2000023222321131-0000302302133013-2212322120103331-0210122331000112-0102011101103103"></a>

<a id="canonical-0010120010023033-2022000221031010-0032012232132100-2202110221111020-2101111131322111-2231303333003023-3331222213020030-2132201233132203"></a>

#### `virtual_server.udp.server_ssl_profile.uid` property

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

<a id="canonical-1333123302300223-1003212003320120-3011131021320211-3121211002322222-0101202302010313-3232321330230032-2023323301311101-0003010123121122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.udp.udp_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.udp](data-sources--application_profiles--reference--group-004.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- virtual_server.udp.udp_client_profile

<a id="canonical-1022022013310211-3131023311012011-0332223013233330-0310123313002302-3003030100013303-1331103101100223-0002131121201323-2223000021033100"></a>

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0301210302022132-0020130322110121-3212332231103000-2220010231321101-1121210123212000-3331231200313121-2313212213032213-2111102110313201"></a>

### Direct properties for `virtual_server.udp.udp_client_profile`

<a id="canonical-3331321213300032-2002032122201223-3132221303332332-3112013011102232-2303100021103021-1210120010303210-3203033000020233-0323212333002033"></a>

#### `virtual_server.udp.udp_client_profile.kind` property

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

<a id="canonical-1320223331022233-0211201230233120-0223002213100200-3210120200222330-2133220001131202-2031210132300313-2232233032130111-3320131033220202"></a>

<a id="canonical-3320111222213100-3231320333013320-1323221302220200-3100321131210321-0103231030033221-1002333030233011-1231030110210213-0230130031033300"></a>

#### `virtual_server.udp.udp_client_profile.name` property

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

<a id="canonical-2303121332220020-2200102003102322-3210233320133012-3021031022032321-2130102310302211-3320012122133032-1213132002202031-3321312330233131"></a>

<a id="canonical-2230010210211302-3022101312021012-0313010120022311-1202121201221210-0313201121001130-0132132131332323-3100022200200113-3003001230213130"></a>

#### `virtual_server.udp.udp_client_profile.namespace` property

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
  }
}
```

<a id="canonical-1111313003330213-2112303201023030-0332303020331322-0221001122220130-3311311030311033-3231130330213022-3301032100122122-3323200303203223"></a>

<a id="canonical-3201000211322032-2111322311331122-3323221103332121-3030313023322300-1202110211203313-0333312100103222-1210321132212031-1301102220332223"></a>

#### `virtual_server.udp.udp_client_profile.tenant` property

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

<a id="canonical-3123323202130100-3132220203121313-3210312130122012-1311201213301201-2212033102311122-2000111122023232-1101032320213121-3232013020101233"></a>

<a id="canonical-3122130230320101-3203103031110002-0122200103112033-1112133222132323-3331232303100131-3311221320003331-1203020212121003-3030020301312012"></a>

#### `virtual_server.udp.udp_client_profile.uid` property

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

<a id="canonical-0223131020311101-3120313230320201-0310233233131121-2012011133331201-0212100200033333-0303313300113123-1231122103120022-2211230303302203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.udp.udp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.udp](data-sources--application_profiles--reference--group-004.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102)
- virtual_server.udp.udp_server_profile

<a id="canonical-1232210232213322-3200033101023210-1022100231121012-3311323303321202-0130032002133020-2301003122033333-3021010230323203-3032301322031133"></a>

Type: `"list"`. Computed.

Configuration parameter for udp server profile.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1032330312030011-0031331201200120-0223011323013211-2023313113232321-0223213313011220-3202100000302102-3311203320000323-3130310110123323"></a>

### Direct properties for `virtual_server.udp.udp_server_profile`

<a id="canonical-0301232220333320-0212023310223321-2233100220312112-2300331223010303-2011223202200333-1133121101223323-2123120033132120-0231132012112322"></a>

#### `virtual_server.udp.udp_server_profile.kind` property

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

<a id="canonical-1301113121300332-2011130120100301-2332131131321130-2110210213232321-3331311222230133-2321020223101121-1033201330333101-3311123323211313"></a>

<a id="canonical-3012022333322231-2221313102110200-0312033122212223-1031103012333032-1213111332021213-1111013200112310-2100202201233103-1130011332330333"></a>

#### `virtual_server.udp.udp_server_profile.name` property

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

<a id="canonical-1100321230312010-3023132212020002-3332131310320113-0331330002112322-1132021322313012-1211330321022323-0202220332323020-1300313033013323"></a>

<a id="canonical-1013330012330101-3022133131323022-0230231332301333-2313120311102201-0131031311230030-3210233322203202-3322331111221223-1333212130033002"></a>

#### `virtual_server.udp.udp_server_profile.namespace` property

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
  }
}
```

<a id="canonical-0020233331120310-1320122000011301-1300121301323333-1233233113202320-3223012111320202-0231232222032033-2103203000222311-1012123333332121"></a>

<a id="canonical-2111321131320132-0230330323312120-2331032221030312-0300022022320122-2221130122211023-2220302002111122-1132030232213111-3301332323210033"></a>

#### `virtual_server.udp.udp_server_profile.tenant` property

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

<a id="canonical-3302132310200220-1112211012111031-0123230123211311-2131012311210121-3310321130120201-2201123213011210-1123332101010032-1222202211310232"></a>

<a id="canonical-2301120023020030-1302230102012010-3302022112130110-3301131130121012-3012322301301131-1312200131313331-2123121332131231-1033132022233212"></a>

#### `virtual_server.udp.udp_server_profile.uid` property

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

<a id="canonical-0300001203202301-3231100101123312-2010220203023303-2030320120111312-3201010231210333-1312201331313222-1321111320032321-3200132220130130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.virtual_server_state` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.virtual_server_state

<a id="canonical-2021230121300320-1123230031330120-2201000331133031-3132210011103321-1233130023200123-0213223330301201-0132022232032021-1202010313232100"></a>

Type: `"single"`. Computed.

Displays the current state on the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-state_choice": "[\"state_disabled\",\"state_enabled\"]"
}
```

<a id="canonical-1333033230000332-0021310100001130-0021023213233010-3332231111100332-1230032033003120-0003312231011111-0200201331200011-3100010301130332"></a>

### Direct properties for `virtual_server.virtual_server_state`

- [state_disabled](data-sources--application_profiles--reference--group-004.md#canonical-2212232221003212-1302330323113311-3333332222133201-0120022103313000-3321101301200111-3002031100221102-3323101001203133-0111000121101200): complete subsection reference.

- [state_enabled](data-sources--application_profiles--reference--group-004.md#canonical-3203221331230032-2132020121120220-2200010100021202-3030323231022121-2210002223020313-3132333233332103-3321200032123011-2330111003302323): complete subsection reference.

<a id="canonical-2212232221003212-1302330323113311-3333332222133201-0120022103313000-3321101301200111-3002031100221102-3323101001203133-0111000121101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.virtual_server_state.state_disabled` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.virtual_server_state](data-sources--application_profiles--reference--group-004.md#canonical-0300001203202301-3231100101123312-2010220203023303-2030320120111312-3201010231210333-1312201331313222-1321111320032321-3200132220130130)
- virtual_server.virtual_server_state.state_disabled

<a id="canonical-2333230033233103-1022032110001213-1211312103321332-2213010220333300-1321201012201122-0222101313111331-2101220020032332-2311032233010201"></a>

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

<a id="canonical-3203221331230032-2132020121120220-2200010100021202-3030323231022121-2210002223020313-3132333233332103-3321200032123011-2330111003302323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.virtual_server_state.state_enabled` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.virtual_server_state](data-sources--application_profiles--reference--group-004.md#canonical-0300001203202301-3231100101123312-2010220203023303-2030320120111312-3201010231210333-1312201331313222-1321111320032321-3200132220130130)
- virtual_server.virtual_server_state.state_enabled

<a id="canonical-2113010333112121-3313300103110012-2033000201111131-3131131322212022-2332033122133012-0222030133022310-0331101012023000-1102133031030230"></a>

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
