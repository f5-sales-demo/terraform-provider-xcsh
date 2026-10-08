---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- job.deploy_options

<a id="canonical-2201123312301322-3002213010120110-0101310130032110-1222000033232121-3323010221320132-2031123210021311-2201300222000213-1333213122220313"></a>

Type: `"single"`. Computed.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-deploy_choice": "[\"all_res\",\"default_virtual_sites\",\"deploy_ce_sites\",\"deploy_ce_virtual_sites\",\"deploy_re_sites\",\"deploy_re_virtual_sites\"]"
}
```

<a id="canonical-2003313022100213-1202211003010203-3033213333110112-0013201201202031-1100231213031333-0121230022030003-3021330002232220-0333130103021313"></a>

### Direct properties for `job.deploy_options`

- [all_res](data-sources--workload--reference--group-005.md#canonical-3022333212212303-1321111310001310-2200103030102023-0111233103012013-3001010023022230-0332203321032201-1311030210111222-3333332311110332): complete subsection reference.

- [default_virtual_sites](data-sources--workload--reference--group-005.md#canonical-2330033102203320-3100220210201032-2312320030013333-2113003020011002-0321030131003332-0032002200101202-2023332331223130-1132210133013232): complete subsection reference.

- [deploy_ce_sites](data-sources--workload--reference--group-005.md#canonical-1301300130132210-2233200331133331-1133231310033011-3223200201212310-1030110112222002-3131123301231130-1310120302333320-2323023111323030): complete subsection reference.

- [deploy_ce_virtual_sites](data-sources--workload--reference--group-005.md#canonical-1211013222123021-1112130112123202-1213020033031330-3330223311120002-1000010231103301-1221032002020001-0100301311003321-0213202223331213): complete subsection reference.

- [deploy_re_sites](data-sources--workload--reference--group-005.md#canonical-1202320102200012-1131132202020311-1033032131123113-1220132013322033-1122003012002200-0100312333012111-1211201333200120-1301030321103033): complete subsection reference.

- [deploy_re_virtual_sites](data-sources--workload--reference--group-005.md#canonical-1023013222020311-2332112300123231-1132323202230033-2312310303310110-1311210102231320-2303202003020301-3011111300300013-2031030103021011): complete subsection reference.

<a id="canonical-3022333212212303-1321111310001310-2200103030102023-0111233103012013-3001010023022230-0332203321032201-1311030210111222-3333332311110332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.all_res` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.deploy_options](data-sources--workload--reference--group-005.md#canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100)
- job.deploy_options.all_res

<a id="canonical-0332330230333012-1001310303100211-1101013002212120-1010220213012231-2102211231102111-0002232011021022-1300213233232120-0233011120300221"></a>

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

<a id="canonical-2330033102203320-3100220210201032-2312320030013333-2113003020011002-0321030131003332-0032002200101202-2023332331223130-1132210133013232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.default_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.deploy_options](data-sources--workload--reference--group-005.md#canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100)
- job.deploy_options.default_virtual_sites

<a id="canonical-3000311130223033-0303132013011031-2210200211130011-1300233003203220-0222023100133212-2102100301311223-0103333311333012-1201001000310103"></a>

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

<a id="canonical-1301300130132210-2233200331133331-1133231310033011-3223200201212310-1030110112222002-3131123301231130-1310120302333320-2323023111323030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_ce_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.deploy_options](data-sources--workload--reference--group-005.md#canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100)
- job.deploy_options.deploy_ce_sites

<a id="canonical-1332323223010223-3002123232331331-3123221210130223-2131123230023012-0033010333223003-3303231002312030-1221220131120022-3203022213233302"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Customer sites.

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

<a id="canonical-3321023211110201-0231010101002232-3332301320211303-1132100020212031-3300123201311033-2310132033130110-0302300110032212-0110331011001332"></a>

### Direct properties for `job.deploy_options.deploy_ce_sites`

- [site](data-sources--workload--reference--group-005.md#canonical-0011310132211003-2021222110132333-2011013213032020-3110030211131113-3100030221320022-2210023221301002-1112210332020332-1220200213310131): complete subsection reference.

<a id="canonical-0011310132211003-2021222110132333-2011013213032020-3110030211131113-3100030221320022-2210023221301002-1112210332020332-1220200213310131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_ce_sites.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.deploy_options](data-sources--workload--reference--group-005.md#canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100)
- [job.deploy_options.deploy_ce_sites](data-sources--workload--reference--group-005.md#canonical-1301300130132210-2233200331133331-1133231310033011-3223200201212310-1030110112222002-3131123301231130-1310120302333320-2323023111323030)
- job.deploy_options.deploy_ce_sites.site

<a id="canonical-1321232000313032-3003113133013300-0133213011330123-3231130301110330-1231100323112001-3200212103010201-2210202010110103-1010203133111211"></a>

Type: `"list"`. Computed.

Which customer sites should this workload be deployed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3330121123223101-1231321230212113-1203111200120200-0102202211222233-2300003231232303-1003112003331311-3223221133330030-3112021201020032"></a>

### Direct properties for `job.deploy_options.deploy_ce_sites.site`

<a id="canonical-0132311133022103-2312222011322121-1010323223013300-0100210002011210-1232301332200200-3021000311031100-3013321110321102-1313212320133131"></a>

#### `job.deploy_options.deploy_ce_sites.site.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3021123032002022-0230310233120113-0011222202031120-0221130333310231-0001003202110323-2100100011023230-0300212221111010-3000303022220320"></a>

<a id="canonical-1120311222323311-2033313010321122-1323030313022101-2310102222332111-1302232002232023-1310331201021112-2311112231222232-3110010321200313"></a>

#### `job.deploy_options.deploy_ce_sites.site.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0120311223200213-1003123211300020-2121222001303121-1321201002131222-3313133032033310-0010213223221131-0013200132232112-3120203211013013"></a>

<a id="canonical-1031111223000203-2212220313120032-3333210032320101-2333110030203023-1202022111120032-1033113122233210-1222100011110003-0001011211031100"></a>

#### `job.deploy_options.deploy_ce_sites.site.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1211013222123021-1112130112123202-1213020033031330-3330223311120002-1000010231103301-1221032002020001-0100301311003321-0213202223331213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_ce_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.deploy_options](data-sources--workload--reference--group-005.md#canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100)
- job.deploy_options.deploy_ce_virtual_sites

<a id="canonical-1331032012232112-0032331021120320-0031302223201001-2033231112330211-3113032200012233-3132030203330233-2131321112332302-3223212230332022"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Customer virtual sites.

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

<a id="canonical-3222131220210001-3331221130100223-2221133020113132-2232321213303002-1221132300133023-0323311130003022-2022123301122112-3302223101122131"></a>

### Direct properties for `job.deploy_options.deploy_ce_virtual_sites`

- [virtual_site](data-sources--workload--reference--group-005.md#canonical-3030003033201332-0223322021321322-1213030200000231-2211220001002212-1012202100130133-0313121232322213-3322202110021121-3200121323032010): complete subsection reference.

<a id="canonical-3030003033201332-0223322021321322-1213030200000231-2211220001002212-1012202100130133-0313121232322213-3322202110021121-3200121323032010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_ce_virtual_sites.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.deploy_options](data-sources--workload--reference--group-005.md#canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100)
- [job.deploy_options.deploy_ce_virtual_sites](data-sources--workload--reference--group-005.md#canonical-1211013222123021-1112130112123202-1213020033031330-3330223311120002-1000010231103301-1221032002020001-0100301311003321-0213202223331213)
- job.deploy_options.deploy_ce_virtual_sites.virtual_site

<a id="canonical-0022010112302101-1222213030211001-1320232331113110-3123113210332301-3313330212102120-1100301000122332-1233320310323102-3203020123302102"></a>

Type: `"list"`. Computed.

Which customer virtual sites should this workload be deployed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1330132003333110-1013303311331301-3023003331110130-2221313311222222-3223202122021233-0303223110002203-3111112013203110-3130332010113232"></a>

### Direct properties for `job.deploy_options.deploy_ce_virtual_sites.virtual_site`

<a id="canonical-0010320001201122-0003223011203100-0200213113000201-0223003301001331-3320323112101111-1131330101131013-0303132031321201-0231310303021331"></a>

#### `job.deploy_options.deploy_ce_virtual_sites.virtual_site.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2111101331123311-3121232313010023-1223232331203023-2213221102211202-3031201132303020-0322223321202323-1100211211101112-3211021103202332"></a>

<a id="canonical-1131110311131023-3021323332303220-2213302002210320-0130120030103213-0200031020232330-0300023232223201-1021111012232212-3003300000010120"></a>

#### `job.deploy_options.deploy_ce_virtual_sites.virtual_site.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3320231130213021-3011201213121231-3122020233023310-1201000113012111-0113320232130322-1112200111123103-1332203320100030-0220002312012323"></a>

<a id="canonical-3131011212013211-0221232311003200-0301130000022233-1133101000113113-3330300201320000-1011203322320123-2132210103013000-3130200032321130"></a>

#### `job.deploy_options.deploy_ce_virtual_sites.virtual_site.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1202320102200012-1131132202020311-1033032131123113-1220132013322033-1122003012002200-0100312333012111-1211201333200120-1301030321103033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_re_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.deploy_options](data-sources--workload--reference--group-005.md#canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100)
- job.deploy_options.deploy_re_sites

<a id="canonical-0121010203332011-3233132113313020-0200022323123133-1222013022222010-1102212010220000-3103311310121223-2303313023223313-0021030201202012"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Regional Edge sites.

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

<a id="canonical-0033233120000111-2210002101022322-3010301211213323-1031301331313322-3032121110223323-0031302202113033-2131102113010222-3012330310213010"></a>

### Direct properties for `job.deploy_options.deploy_re_sites`

- [site](data-sources--workload--reference--group-005.md#canonical-3032232110111102-1002001332310332-0132220031332211-2102202330122101-3022021331323132-3310030023232300-0231313111320103-1133120312332201): complete subsection reference.

<a id="canonical-3032232110111102-1002001332310332-0132220031332211-2102202330122101-3022021331323132-3310030023232300-0231313111320103-1133120312332201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_re_sites.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.deploy_options](data-sources--workload--reference--group-005.md#canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100)
- [job.deploy_options.deploy_re_sites](data-sources--workload--reference--group-005.md#canonical-1202320102200012-1131132202020311-1033032131123113-1220132013322033-1122003012002200-0100312333012111-1211201333200120-1301030321103033)
- job.deploy_options.deploy_re_sites.site

<a id="canonical-1303101002313201-1120103230221210-3103112101022100-0010032330001313-2001111030133131-2210313023100000-0123010331031102-1201323203303332"></a>

Type: `"list"`. Computed.

Which regional edge sites should this workload be deployed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0103100120030202-0301033111002011-2323113000231300-3003000210013310-1230211023233320-2323322122012003-3112320130133312-1101200233102021"></a>

### Direct properties for `job.deploy_options.deploy_re_sites.site`

<a id="canonical-0112330100020323-3022112110232121-1110223310220032-0302222323133130-0201203122301322-0130312320322132-1102133020122320-0012301230221032"></a>

#### `job.deploy_options.deploy_re_sites.site.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1201031231020012-1011221033301113-1311123201000113-3031032133202202-1212110001132231-0223111213003022-1322032303233133-3203213223110333"></a>

<a id="canonical-1101330130211021-0131131011203301-3312030031301021-0223112123222330-0211110031013302-0111112111332211-1032033303131000-1332013332303033"></a>

#### `job.deploy_options.deploy_re_sites.site.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3032021012310110-2010300123222021-3120212332232323-1033101211230323-2312220110123202-2233001223223103-2132020300311011-3322222000320322"></a>

<a id="canonical-1103200120010332-2310011031010102-0331201101313133-0310013123203110-2010131222020031-2312230003100121-2212302003013122-0210102210210130"></a>

#### `job.deploy_options.deploy_re_sites.site.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1023013222020311-2332112300123231-1132323202230033-2312310303310110-1311210102231320-2303202003020301-3011111300300013-2031030103021011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_re_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.deploy_options](data-sources--workload--reference--group-005.md#canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100)
- job.deploy_options.deploy_re_virtual_sites

<a id="canonical-2112313310323131-2133323013100023-1122321221332122-1310321000211101-1321031311112100-2001100321233223-3303130331100300-3232131131302302"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Regional Edge virtual sites.

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

<a id="canonical-0101012033102031-0312023021312123-2023112122113301-1200230113111222-0200033102223313-0031112032003303-0222321201103101-2311322011200312"></a>

### Direct properties for `job.deploy_options.deploy_re_virtual_sites`

- [virtual_site](data-sources--workload--reference--group-005.md#canonical-2133321113003323-0121133130113100-1022312020300223-2211310011311302-3213130311100023-0002103030003100-1300233030201100-2020332002000122): complete subsection reference.

<a id="canonical-2133321113003323-0121133130113100-1022312020300223-2211310011311302-3213130311100023-0002103030003100-1300233030201100-2020332002000122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_re_virtual_sites.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.deploy_options](data-sources--workload--reference--group-005.md#canonical-1021013221100131-2310031003132222-1301022330312301-0200301110312230-2121132222333110-2302011012310013-1021020103223211-3032100212011100)
- [job.deploy_options.deploy_re_virtual_sites](data-sources--workload--reference--group-005.md#canonical-1023013222020311-2332112300123231-1132323202230033-2312310303310110-1311210102231320-2303202003020301-3011111300300013-2031030103021011)
- job.deploy_options.deploy_re_virtual_sites.virtual_site

<a id="canonical-0323110211120011-3201202300102110-0000102010222133-0200222003100131-2012032223132230-1123230333302100-2020130333301112-1111101000032031"></a>

Type: `"list"`. Computed.

Which regional edge virtual sites should this workload be deployed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2201203310321221-2310321320130202-1310313132102010-1211330013323111-0311012023131101-3221230200332303-1000223123012021-3201312333102220"></a>

### Direct properties for `job.deploy_options.deploy_re_virtual_sites.virtual_site`

<a id="canonical-3101001313223003-0022230210303233-0222110311010132-1331313310010130-3333211032112313-1300013322010303-0023121213120312-0310110322232133"></a>

#### `job.deploy_options.deploy_re_virtual_sites.virtual_site.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2000113333221213-1323120213012111-0020212100230333-1121100022333123-0232121330212021-3100211132320003-3100330332222112-2000322202200032"></a>

<a id="canonical-3222321332322211-3330310302011331-1233323133003003-0320220112122120-1021213113110312-2312121131220223-2000230133011113-3121201221232313"></a>

#### `job.deploy_options.deploy_re_virtual_sites.virtual_site.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1102211003021130-2320032201212320-1230102331320312-3200302201021313-3310302322030313-1323122130111031-1313103210212202-0312001033233333"></a>

<a id="canonical-1212012101213332-0013210032321322-2200033101120023-3113121331112132-3003233210100323-3013121100203303-0220201321101001-2123123223000012"></a>

#### `job.deploy_options.deploy_re_virtual_sites.virtual_site.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- job.volumes

<a id="canonical-1033130023211123-2012222333010201-0330213011230302-2120102303131130-0331011011022111-2210010313021230-2001123122232322-3133211213312202"></a>

Type: `"list"`. Computed.

Volumes. Volumes for the job.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2101113133120003-3003202322121223-3023301132320100-1332121223022221-2110110103233303-2100021133200312-0213122232320121-1132331013010222"></a>

### Direct properties for `job.volumes`

- [empty_dir](data-sources--workload--reference--group-005.md#canonical-2130120111010100-0110033232023301-3123203333112202-2023301113120022-2031322212021222-1301300101103333-2131320002113230-3100111121000003): complete subsection reference.

- [host_path](data-sources--workload--reference--group-005.md#canonical-1101123330022122-0312023230320121-2333332030310233-1331312111301103-0120032000312031-2200113123130032-2333123300102003-0212211023011122): complete subsection reference.

<a id="canonical-0321010210201100-3301003113223022-1300021101110003-3010320203102120-2120113100032001-3301032131022121-0132131120203020-0211202303213211"></a>

<a id="canonical-1322333113003001-3311000220023202-2033313003030100-1220031210131310-1000121023300111-2333113133113023-3102032310120030-2130013222300020"></a>

#### `job.volumes.name` property

Type: `"string"`. Computed.

Name. Name of the volume.

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

- [persistent_volume](data-sources--workload--reference--group-005.md#canonical-3011233111030203-0033211202203110-3130100112113111-2300013120101211-1010133312301103-0331203321122300-2120131033331030-1130111020230122): complete subsection reference.

<a id="canonical-2130120111010100-0110033232023301-3123203333112202-2023301113120022-2031322212021222-1301300101103333-2131320002113230-3100111121000003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.empty_dir` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-005.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- job.volumes.empty_dir

<a id="canonical-0310032230133032-0023010001213022-1133303022000033-1333322022002223-3003321031301121-2112001232120103-0213111013230010-3011332003011330"></a>

Type: `"single"`. Computed.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

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

<a id="canonical-3032221200301202-1232210131113233-2100300013122332-0113202333302023-2102130201121101-3012233222111212-3223113021311103-1001313032220221"></a>

### Direct properties for `job.volumes.empty_dir`

- [mount](data-sources--workload--reference--group-005.md#canonical-0101212202101110-3211221011321012-1321002221302231-0011232010322023-0213312120301331-2131211330303312-0021201120020221-3023211232310311): complete subsection reference.

<a id="canonical-0201223211031312-1031002311011112-1313131011133102-0110010000112033-2320022231312112-1232123230220210-1223321231030210-2220120020030011"></a>

<a id="canonical-0003220233130012-3122101311231010-3020222210312302-1310131000101333-2333231010333231-2001013103002213-3132202100333002-3032331310222203"></a>

#### `job.volumes.empty_dir.size_limit` property

Type: `"number"`. Computed.

Size Limit (in GiB). Configuration parameter for size limit

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
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0101212202101110-3211221011321012-1321002221302231-0011232010322023-0213312120301331-2131211330303312-0021201120020221-3023211232310311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.empty_dir.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-005.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- [job.volumes.empty_dir](data-sources--workload--reference--group-005.md#canonical-2130120111010100-0110033232023301-3123203333112202-2023301113120022-2031322212021222-1301300101103333-2131320002113230-3100111121000003)
- job.volumes.empty_dir.mount

<a id="canonical-0332021122031212-0002123312321303-2332222001332012-1122232330203313-1303301030233333-1001003100022020-1013011323233333-1210330300323220"></a>

Type: `"single"`. Computed.

Volume mount describes how volume is mounted inside a workload.

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

<a id="canonical-1021011220123132-2303202230001222-0120032201300320-0003023331300230-0233211220122131-1203202123203303-2120003103132333-3132332100022303"></a>

### Direct properties for `job.volumes.empty_dir.mount`

<a id="canonical-1212313131211332-0230320213220330-0000131332130203-1230212331033111-1101211332102110-2323332122232310-3313020220201103-0321120100223300"></a>

#### `job.volumes.empty_dir.mount.mode` property

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0032310221112310-1021012113333232-2321111301322033-3222112230113113-3320121221012001-1010103230113330-0211010113012223-2223122010112000"></a>

<a id="canonical-2002010121033112-3200102223232111-0011032010113333-3033103312003103-2213221122033102-0302133132201130-3332101332012030-0222333222103033"></a>

#### `job.volumes.empty_dir.mount.mount_path` property

Type: `"string"`. Computed.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^[^:]*$"
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
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-3113011202103333-0311112133002231-0031313322120303-2320232001112310-0111120101021233-3312202220132133-2222110331231033-0202020222000332"></a>

<a id="canonical-3322222132111132-3131033032033020-1113013113223020-0312321232331021-2132010333032310-2321332000033230-3331333221232223-3013122211031210"></a>

#### `job.volumes.empty_dir.mount.sub_path` property

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Additional upstream details:

Defaults to "" (volume's root).

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1101123330022122-0312023230320121-2333332030310233-1331312111301103-0120032000312031-2200113123130032-2333123300102003-0212211023011122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.host_path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-005.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- job.volumes.host_path

<a id="canonical-3032121221233110-3030310122102033-2030132313213021-3132230121020002-3232202223132212-1001200103123221-1320321202312112-1332010002123113"></a>

Type: `"single"`. Computed.

Volume containing a host mapped path into the workload.

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

<a id="canonical-1131132030121101-1230330312001132-3121131331003211-0131220202101303-3220002231230321-2010131203330331-0020213203202100-1110203022110023"></a>

### Direct properties for `job.volumes.host_path`

- [mount](data-sources--workload--reference--group-005.md#canonical-0022221320310321-2112221112300113-0011001001322222-0322103111312303-3301123301230120-2110322111300310-1131013123130113-2032020111303012): complete subsection reference.

<a id="canonical-3231122111330220-2010302321312021-3032310003111130-1002103133101122-3203223331012213-2120213033020203-3002223030210030-1002013232020311"></a>

<a id="canonical-2020210030133023-1111230302013113-0222330031032001-3321230121332222-2002323222012112-2023322132230103-1011300322330111-1022200010101003"></a>

#### `job.volumes.host_path.path` property

Type: `"string"`. Computed.

Path. Path of the directory on the host.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "[^\\\\0]+"
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
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  }
}
```

<a id="canonical-0022221320310321-2112221112300113-0011001001322222-0322103111312303-3301123301230120-2110322111300310-1131013123130113-2032020111303012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.host_path.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-005.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- [job.volumes.host_path](data-sources--workload--reference--group-005.md#canonical-1101123330022122-0312023230320121-2333332030310233-1331312111301103-0120032000312031-2200113123130032-2333123300102003-0212211023011122)
- job.volumes.host_path.mount

<a id="canonical-2333332310101300-2123303210202313-2330123121300020-0030003303230020-2211012303131320-1231100223030202-0201311200132221-2333310302020130"></a>

Type: `"single"`. Computed.

Volume mount describes how volume is mounted inside a workload.

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

<a id="canonical-3102202012221213-2323012232130120-2231121213210320-0030121020321322-3231003202322202-1200313303310213-2221111131220222-0312133113321112"></a>

### Direct properties for `job.volumes.host_path.mount`

<a id="canonical-1202111233111003-2012021012123312-2333121312023103-0131110030230022-0310310000323300-2100131020211233-0202312110011232-0222222132113323"></a>

#### `job.volumes.host_path.mount.mode` property

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3322330232021103-1211030331111320-0031010220303232-2031010333020000-2201101121111201-0313222213211130-0031210033221130-3230302003233323"></a>

<a id="canonical-2001221201223331-0132231331300033-0200210021233110-1101022212212200-3020010312323102-1323101211210203-2112221302301201-2121301301121220"></a>

#### `job.volumes.host_path.mount.mount_path` property

Type: `"string"`. Computed.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^[^:]*$"
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
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-3312223013002300-2200103213313212-2020230133110102-2200033301201031-3222131200020031-1223030212203011-3002120303113302-2112200010233213"></a>

<a id="canonical-3031101120121103-0322012312232210-1211212202311230-3110003330321030-2031323113033003-1121030330031030-2213103233123201-0101133023131333"></a>

#### `job.volumes.host_path.mount.sub_path` property

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Additional upstream details:

Defaults to "" (volume's root).

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3011233111030203-0033211202203110-3130100112113111-2300013120101211-1010133312301103-0331203321122300-2120131033331030-1130111020230122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.persistent_volume` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-005.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- job.volumes.persistent_volume

<a id="canonical-0233313130310321-0131232032110223-1003303130023333-0113210030202110-1333330331002000-0211200300312120-2202002303300002-0021023112233221"></a>

Type: `"single"`. Computed.

Volume containing the Persistent Storage for the workload.

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

<a id="canonical-1300231213311111-2312022103113030-0302030303223203-1233210000103011-2102323123010320-3313000230313131-2001033303123122-3221210313230330"></a>

### Direct properties for `job.volumes.persistent_volume`

- [mount](data-sources--workload--reference--group-005.md#canonical-2010021302120122-2302212220000131-0331121210311130-1132100012321020-2300301213211223-1103022231030103-3223102120121032-3230310011031333): complete subsection reference.

- [storage](data-sources--workload--reference--group-005.md#canonical-2213202331120023-2231110111323223-1301320301300323-2103321202200000-2113211210103200-1220133202221033-2101020223110213-3203011223010331): complete subsection reference.

<a id="canonical-2010021302120122-2302212220000131-0331121210311130-1132100012321020-2300301213211223-1103022231030103-3223102120121032-3230310011031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.persistent_volume.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-005.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- [job.volumes.persistent_volume](data-sources--workload--reference--group-005.md#canonical-3011233111030203-0033211202203110-3130100112113111-2300013120101211-1010133312301103-0331203321122300-2120131033331030-1130111020230122)
- job.volumes.persistent_volume.mount

<a id="canonical-2301102032332330-0323131113112130-1023320001030031-0023113120002002-0003003303231020-2000333023311123-0003330122321232-2023122103113002"></a>

Type: `"single"`. Computed.

Volume mount describes how volume is mounted inside a workload.

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

<a id="canonical-2003123031221212-2011122320322011-2323011310030311-0011031230122122-0023302322303112-0313123200233112-1111210102222202-2203202221103000"></a>

### Direct properties for `job.volumes.persistent_volume.mount`

<a id="canonical-1220022010120222-3130130220133103-0333112311020013-1210223021133331-1100302002331010-2010133320212122-0111001130210330-2213020333020230"></a>

#### `job.volumes.persistent_volume.mount.mode` property

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3201313312213302-3311003103200220-2102223212300011-1211000313310120-1223202313212223-0031223231133003-3020211020323123-1321231331011000"></a>

<a id="canonical-2012213233202333-0233303120011021-0333330133220230-3303130212123011-3331023121310210-2313221132002033-3021220333000221-1311132302022122"></a>

#### `job.volumes.persistent_volume.mount.mount_path` property

Type: `"string"`. Computed.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^[^:]*$"
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
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-0202311201231023-2103331210112320-3103300211313330-2323122112023223-1132313033030112-0031302300121302-3200200033120233-1230332212112212"></a>

<a id="canonical-1303303200033023-1002313123313022-3111122033113220-0333220003200220-3322032213111222-1233212212320133-1022032331203322-0031031220120033"></a>

#### `job.volumes.persistent_volume.mount.sub_path` property

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Additional upstream details:

Defaults to "" (volume's root).

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2213202331120023-2231110111323223-1301320301300323-2103321202200000-2113211210103200-1220133202221033-2101020223110213-3203011223010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.persistent_volume.storage` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-005.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- [job.volumes.persistent_volume](data-sources--workload--reference--group-005.md#canonical-3011233111030203-0033211202203110-3130100112113111-2300013120101211-1010133312301103-0331203321122300-2120131033331030-1130111020230122)
- job.volumes.persistent_volume.storage

<a id="canonical-1133313113333331-0113021203200232-0323101023131132-2222323202131120-0001232203231303-0032031301233302-2330331332213222-0123131112012311"></a>

Type: `"single"`. Computed.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

<a id="canonical-2301231112323020-0113230100000212-2212020111010033-1120232312201323-2312300221021101-2102223203031230-1031001103103030-1213103232001130"></a>

### Direct properties for `job.volumes.persistent_volume.storage`

<a id="canonical-3031230232020210-3332201001000002-1213031313223022-0312220031000013-3023320012320222-2003330222022320-1031033132323323-0011210033111310"></a>

#### `job.volumes.persistent_volume.storage.access_mode` property

Type: `"string"`. Computed.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Additional upstream details:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

Receipt-pinned upstream constraints:

```json
{
  "default": "ACCESS_MODE_READ_WRITE_ONCE",
  "enum": [
    "ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1030321013112223-0012322202200231-3200300030221000-2002030202020103-3011010021131003-3100321320120302-1212333212220022-0132022221203023"></a>

<a id="canonical-0021012032223202-1212011213303021-1213311321131023-2001103211000230-0003122100311030-2322323330310130-2233201233123121-0201122112003131"></a>

#### `job.volumes.persistent_volume.storage.class_name` property

Type: `"string"`. Computed.

Exclusive with \[default\] Use the specified class name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [default](data-sources--workload--reference--group-005.md#canonical-3031333332222000-0002301011111123-2020200223121303-0302222312002201-2011203021303011-2011031030112301-1112303033132013-2102231120000012): complete subsection reference.

<a id="canonical-3032030303333030-1131201333223310-0333133101133311-3001120202030013-0110211200203303-3102211022112202-1112232223030103-3023021302212320"></a>

<a id="canonical-0230333132122310-3302120120000213-1021300110313310-1101133231232123-3221202303303223-0221003232031303-2023230112002220-2120232130120032"></a>

#### `job.volumes.persistent_volume.storage.storage_size` property

Type: `"number"`. Computed.

Size (in GiB). Size in GiB of the persistent storage.

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
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3031333332222000-0002301011111123-2020200223121303-0302222312002201-2011203021303011-2011031030112301-1112303033132013-2102231120000012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.persistent_volume.storage.default` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [job](data-sources--workload--reference--group-004.md#canonical-1303110033111110-3332103120123120-2010122031013021-3300123211220120-0230203021233123-3120213311303022-1232311010202130-0212201103321131)
- [job.volumes](data-sources--workload--reference--group-005.md#canonical-3021110112330331-0022212121010102-1001033022212233-0103130310010213-1220331221322222-1121220303110310-2120010230203232-3103100103322011)
- [job.volumes.persistent_volume](data-sources--workload--reference--group-005.md#canonical-3011233111030203-0033211202203110-3130100112113111-2300013120101211-1010133312301103-0331203321122300-2120131033331030-1130111020230122)
- [job.volumes.persistent_volume.storage](data-sources--workload--reference--group-005.md#canonical-2213202331120023-2231110111323223-1301320301300323-2103321202200000-2113211210103200-1220133202221033-2101020223110213-3203011223010331)
- job.volumes.persistent_volume.storage.default

<a id="canonical-2331303313212213-3033311322132033-3213230131121102-2002321003221021-3033232321311103-1101101232101230-3332033321123302-3323300212022132"></a>

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

<a id="canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- service

<a id="canonical-3033313320232002-0000103223202222-2330001311332233-1012320230331333-0131031100203120-3310301322123012-3111111013032000-3120131111111003"></a>

Type: `"single"`. Computed.

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers,
traditional SQL databases, etc.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-scaling_choice": "[\"num_replicas\",\"scale_to_zero\"]"
}
```

<a id="canonical-3213321311112322-0123030330020013-1033031333333121-0122333201200230-1010002000203300-3132202123130323-2110311312211020-1022032002221212"></a>

### Direct properties for `service`

- [advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333): complete subsection reference.

- [configuration](data-sources--workload--reference--group-013.md#canonical-3123002002331230-0201233111020120-1333221333320003-3110223112030011-2003333010313113-2333211001212031-3333112223203131-0322131002310210): complete subsection reference.

- [containers](data-sources--workload--reference--group-014.md#canonical-3303321131200101-3301301201213103-2323012233313212-1210201321103331-0012010102322220-2030122213101120-3013200302130003-0123211020123033): complete subsection reference.

- [deploy_options](data-sources--workload--reference--group-014.md#canonical-0121032001011031-2321102332010202-3221321233030232-3111022330331330-3220203323221010-1002033122322321-2321230330230311-1011232330212320): complete subsection reference.

<a id="canonical-3033322032211333-1002012201102210-3311003332200001-2230320021303332-1231101322120211-0020012313011102-1312312032131100-3011312122022012"></a>

<a id="canonical-0333113322000120-1031210112210330-1312013301011320-0112130331103100-2203123113112330-3102112213002333-1200021000113011-3333203233311130"></a>

#### `service.num_replicas` property

Type: `"number"`. Computed.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  }
}
```

- [scale_to_zero](data-sources--workload--reference--group-014.md#canonical-1232232223133202-3202330012213121-2321010021332212-2131311121001321-2011131113311031-0312303012321112-3333213331000032-1000100101021221): complete subsection reference.

- [volumes](data-sources--workload--reference--group-015.md#canonical-2012110033111201-2323122231313121-3301031330103210-2312132322110233-3323213233311103-0320232202220320-3231001222133133-1333230032132310): complete subsection reference.

<a id="canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- service.advertise_options

<a id="canonical-2111112122123010-3312212303211300-3011110031232000-2122013331121023-0302003021221312-0100321310010132-2120030223023101-2323320131131121"></a>

Type: `"single"`. Computed.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

<a id="canonical-3333022032223101-3022322211323203-1330201330032003-1213103133202112-3002121132230101-2223222331021110-1003101203201201-0223213211322211"></a>

### Direct properties for `service.advertise_options`

- [advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202): complete subsection reference.

- [advertise_in_cluster](data-sources--workload--reference--group-008.md#canonical-2131020211211101-0213303111310032-0232312121101303-0303013213123313-0112311302213302-2012200033230321-1231231303121203-0121103301113123): complete subsection reference.

- [advertise_on_public](data-sources--workload--reference--group-008.md#canonical-1310110010122033-2210320003110123-0220310131330230-1101011120230112-0132203111212320-2110123230130221-3300230123133301-1202011310100130): complete subsection reference.

- [do_not_advertise](data-sources--workload--reference--group-013.md#canonical-3303332030303022-2112310112031201-1232033023113213-2310022310310212-0130333030130000-2200322102123323-3201131322011030-3120330230033000): complete subsection reference.

<a id="canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- service.advertise_options.advertise_custom

<a id="canonical-0320002303303233-1133333031302200-3112000002132103-1002313233333012-0030203231021001-3212311010020303-1222231120010030-2223221301211033"></a>

Type: `"single"`. Computed.

Advertise this workload via loadbalancer on specific sites.

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

<a id="canonical-3232001112121033-1213130020020113-0131213222310210-1001301323103233-2033312221010112-2233201210100222-1310022330320100-0010300300131330"></a>

### Direct properties for `service.advertise_options.advertise_custom`

- [advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120): complete subsection reference.

- [ports](data-sources--workload--reference--group-006.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102): complete subsection reference.

<a id="canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- service.advertise_options.advertise_custom.advertise_where

<a id="canonical-0132031021332232-3310223001033103-2201030130333220-1302301213231102-3102321210131300-3020223133233323-3212323130333320-3122012330333112"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2200132002321331-0220011230311103-1212332110133031-1311330123311230-0001212212021301-1020103013200032-3230321201202102-1330013301121220"></a>

### Direct properties for `service.advertise_options.advertise_custom.advertise_where`

- [site](data-sources--workload--reference--group-005.md#canonical-0210013023032210-1303200212213212-0113330010223002-3121320332310012-2133211001112103-2302323102210031-3211012011131201-0112121032223032): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-005.md#canonical-2112123100111000-3331023313331020-0233223333000211-0132023311103233-2201330132013110-1212313303103212-1033013200123000-3033112112202011): complete subsection reference.

- [vk8s_service](data-sources--workload--reference--group-005.md#canonical-2232032323312210-1320210021123102-2323123132332020-2021233311111231-1003000211331102-3032322003100330-1011330303201322-1303131022011320): complete subsection reference.

<a id="canonical-0210013023032210-1303200212213212-0113330010223002-3121320332310012-2133211001112103-2302323102210031-3211012011131201-0112121032223032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-2013031300033220-2222211301013030-0222301231211010-1233101000331333-0221103132121002-1013131033111112-0112030331331011-0012220021221210"></a>

Type: `"single"`. Computed.

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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

<a id="canonical-0211011101013302-3121100301211103-0022131122020311-0022232232002221-2330202301323212-3222013032312002-3232231103322323-1302023211013130"></a>

### Direct properties for `service.advertise_options.advertise_custom.advertise_where.site`

<a id="canonical-3103122022111023-0130030220320133-1233310022331212-1012010231100031-1013121113130301-0300330003333202-2132332120022320-2011111100011100"></a>

#### `service.advertise_options.advertise_custom.advertise_where.site.ip` property

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1213300211031021-2230312332223021-2210220121001332-1322012213203003-2131202230220001-0230233102012232-1001222030003131-3201030313002012"></a>

<a id="canonical-0210122303333321-2230033001312211-0232112212220103-0121030012303120-0311201033311112-2302203333130113-3300311212310112-0200020312111301"></a>

#### `service.advertise_options.advertise_custom.advertise_where.site.network` property

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](data-sources--workload--reference--group-005.md#canonical-0130221011231112-0333020231020303-0033321013223220-3100120222201331-2320332021131131-1002223001132122-2230012033013322-3030303323010032): complete subsection reference.

<a id="canonical-0130221011231112-0333020231020303-0033321013223220-3100120222201331-2320332021131131-1002223001132122-2230012033013322-3030303323010032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- [service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-005.md#canonical-0210013023032210-1303200212213212-0113330010223002-3121320332310012-2133211001112103-2302323102210031-3211012011131201-0112121032223032)
- service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-1303030110021132-0212122032100033-3002303022100012-0212311233031211-0200300332112121-2311331233321022-2013112003011130-3102303133011123"></a>

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

<a id="canonical-1031310103331320-3022022333110121-0132100213031102-1102311302111323-3200210032322322-3322201131212301-3132033302020033-0200310311302231"></a>

### Direct properties for `service.advertise_options.advertise_custom.advertise_where.site.site`

<a id="canonical-2020102011302220-3201203331211113-0022210131133110-0103113131210033-3231300223311120-1103323230223303-0321232203330131-2012120110302113"></a>

#### `service.advertise_options.advertise_custom.advertise_where.site.site.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0131032220022101-2333113332130103-3012203222020021-0212201332120233-0103210333031233-1023112301322231-0203123030303032-3333200013201113"></a>

<a id="canonical-1113010131323012-1302031222133033-3123200023120332-3020023020102030-0032023032200210-1302312321132023-3110133032221333-1120023320111031"></a>

#### `service.advertise_options.advertise_custom.advertise_where.site.site.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1010331112001102-3321023021023113-0323223201321113-3230032023010312-3213110000033000-1132003322231330-2111023322200331-3123133031133102"></a>

<a id="canonical-0011000100213202-0302022032000312-1032020032221313-2111002002100132-1320302102110001-1302320333111131-3313013310322332-0100103322213121"></a>

#### `service.advertise_options.advertise_custom.advertise_where.site.site.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2112123100111000-3331023313331020-0233223333000211-0132023311103233-2201330132013110-1212313303103212-1033013200123000-3033112112202011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-1232231102031020-2031301000002113-3121033311113200-0123213122023222-0102301312110333-1321333012003012-2112121232320320-0133231131310003"></a>

Type: `"single"`. Computed.

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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

<a id="canonical-3201133030002202-0010211122100311-0031121100012030-1000002212013113-0130120210003030-3000011111100322-1332003020010313-0202333122202303"></a>

### Direct properties for `service.advertise_options.advertise_custom.advertise_where.virtual_site`

<a id="canonical-1113101220321022-3013313331113112-1222211131213230-0132210333301103-1000003213032210-0101010033030321-1111023032021213-1010302131000301"></a>

#### `service.advertise_options.advertise_custom.advertise_where.virtual_site.network` property

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--workload--reference--group-005.md#canonical-1100031201330203-1100112112130230-0030033303022010-3323302131320331-2313203032003003-0120033213031110-2300123303211331-3332102012011213): complete subsection reference.

<a id="canonical-1100031201330203-1100112112130230-0030033303022010-3323302131320331-2313203032003003-0120033213031110-2300123303211331-3332102012011213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- [service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-005.md#canonical-2112123100111000-3331023313331020-0233223333000211-0132023311103233-2201330132013110-1212313303103212-1033013200123000-3033112112202011)
- service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-0231221030213202-0102021130103102-0121120330223312-3103231100223130-1222211310221003-2121101320311023-0032100300022231-1313023200202000"></a>

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

<a id="canonical-1231202123330300-0303331010200020-0321120133010220-3002211220212330-0113202022000120-2213332223030023-3220130211320030-1021232211120303"></a>

### Direct properties for `service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site`

<a id="canonical-3023100231333232-3201322112312011-1033120210033033-0211220013003311-1211220320131322-1110002023100202-2303322021233030-3211100333311222"></a>

#### `service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3002101231011210-0213013100202020-2123123331011321-2221233031302111-1323020213113233-3322020213013021-2130211111020132-1221100303023021"></a>

<a id="canonical-2301112300102200-3203132113121033-1000233021031202-0121021021110033-1132020303323202-2101130102030123-1220303103210221-1003221211233213"></a>

#### `service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1013131322221321-2212330131033133-0003201211101313-3132332022320322-1201130132320201-3222300112311221-2323123213313220-3321321321223100"></a>

<a id="canonical-1330103013202223-0030323330300100-0020220333112230-0321230032213223-2031022121100032-0012211323113023-3202332210102031-0231233233130023"></a>

#### `service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2232032323312210-1320210021123102-2323123132332020-2021233311111231-1003000211331102-3032322003100330-1011330303201322-1303131022011320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.advertise_where.vk8s_service` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="canonical-2220321101212303-3121021023003030-0013131231220312-0032003002331331-1230130301333031-0030302312303301-2222100110332333-1112100001203002"></a>

Type: `"single"`. Computed.

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

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

<a id="canonical-2003303222030103-2013002220220230-3201003323232211-2203030312031323-0102320132033300-2123133323023101-1002111111030233-0200022021231231"></a>

### Direct properties for `service.advertise_options.advertise_custom.advertise_where.vk8s_service`

- [site](data-sources--workload--reference--group-005.md#canonical-0323310031021201-3000011021131322-0303113302302030-3213110111310120-1100202322111223-2233303021030113-3011200312221033-1303310313222331): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-006.md#canonical-1332131311232133-2211310100013233-3200002302333013-3300200321100312-1212031232231123-0120222202220200-2021332303332132-1020033010323130): complete subsection reference.

<a id="canonical-0323310031021201-3000011021131322-0303113302302030-3213110111310120-1100202322111223-2233303021030113-3011200312221033-1303310313222331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.advertise_where.vk8s_service.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-005.md#canonical-3020022032332103-2113232112010012-1113012311221332-1020310023132012-2121001113110122-0133032022102032-2233121212023331-3101110033310120)
- [service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-005.md#canonical-2232032323312210-1320210021123102-2323123132332020-2021233311111231-1003000211331102-3032322003100330-1011330303201322-1303131022011320)
- service.advertise_options.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-2000021112203331-0030201103032320-2120310311211233-0322011321211300-1121231333202212-3130122022011303-2002022312221102-1231111033230210"></a>

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
