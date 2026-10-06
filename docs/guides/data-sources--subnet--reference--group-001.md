---
page_title: "xcsh_subnet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet reference."
---

# xcsh_subnet reference

<a id="canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- Property reference

<a id="canonical-2011321311320301-0121013032130023-1010001100321210-1321000222000200-3233001001220002-1322232230300231-1303301013001002-0021001321231303"></a>

### Direct properties for `xcsh_subnet`

<a id="canonical-0100321102211332-1132222123020222-2012230232113001-3300102003102123-2313300013333223-3210130311222100-2222322032021202-0221201020122230"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [connect_to_layer2](data-sources--subnet--reference--group-001.md#canonical-3130121321030111-2013113023133031-1213132113021010-2313220033022320-0122011323332121-3021003231200303-0213001233000220-0202302100030301): complete subsection reference.

- [connect_to_slo](data-sources--subnet--reference--group-001.md#canonical-1012300201312111-2313203113003101-3132331220322101-0023000203030112-0122320223110323-3330232223100323-1230223312031103-0333021321332002): complete subsection reference.

<a id="canonical-0120331112023212-1323321311213113-1201020230100112-1333001021112201-1311231032230202-3333330032203000-2211213213032023-2203322331131032"></a>

<a id="canonical-3022311233100302-0020020333002323-2130123122000131-2112023020200122-3210123002032212-0222030010111133-1103111111202231-2203011302111310"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Subnet.

Additional upstream details:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-0100102201231020-3132032311212203-1332212212103013-3032033330121033-1000132000012011-0220221310302003-0001200301012332-3212212122301221"></a>

<a id="canonical-0201133031031032-3223101033130130-3100223333211211-0113032331021202-0011013002311000-3122331213310001-0022232133111131-0213322221302123"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [isolated_nw](data-sources--subnet--reference--group-001.md#canonical-1102221302223332-3311312000120021-2332120301312013-3031323122310132-1201033003230231-3313320312210131-3130022322010033-1200213032110230): complete subsection reference.

<a id="canonical-1023320103220121-0213331201121020-1102022111123101-2002230311000110-2032211022302300-1110020201123202-1131131201010300-0101131020331233"></a>

<a id="canonical-0120002302230130-0331033112210102-0012011101223232-1030123211320010-2200013302103223-2210120112011210-3012121333320020-1212003020323120"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="canonical-3220122013030311-3320033032332320-2320210112011123-2021210232233310-3021203120001321-3213233320012323-0322303203020121-2031213313321231"></a>

<a id="canonical-3011301033200113-0103232111103331-2032220013013130-3133300311122211-1120211031233101-0031002012030200-0232013323321220-2330023230033032"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Subnet.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0103210311221213-3211132221003133-2131303021030301-0020133203310132-2010333223132130-1032232022012011-2302130121302211-1230110222110113"></a>

<a id="canonical-3323022201010323-3022011300023211-3330022200122233-1001210121232111-0110220011103123-3022002210230321-1101022212301232-2321201032030030"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Subnet exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-3013003112313223-3230212001032101-0110030102302020-1201202301310330-3121320032220330-1233021100322013-3321021102221013-3132002132213003): complete subsection reference.

<a id="canonical-0231022203201310-3220101031320020-0303302302002011-1133320012130103-0333211331101103-1121133030212120-3132210020122203-0100212331100223"></a>

### All schema paths for `xcsh_subnet`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--subnet--reference--group-001.md#canonical-0100321102211332-1132222123020222-2012230232113001-3300102003102123-2313300013333223-3210130311222100-2222322032021202-0221201020122230) |
| `connect_to_layer2` | [connect_to_layer2](data-sources--subnet--reference--group-001.md#canonical-0212332133031033-0100212001110111-1101213001010011-0102133302112323-2031022311203332-2332020210023133-0133300323320231-3002221132321311) |
| `connect_to_layer2.layer2_intf_ref` | [connect_to_layer2.layer2_intf_ref](data-sources--subnet--reference--group-001.md#canonical-3133012203233313-1231132232210032-0332022211100033-3312021130112001-2033032313212223-3221220010220302-2002120322212123-3211103022123323) |
| `connect_to_layer2.layer2_intf_ref.name` | [connect_to_layer2.layer2_intf_ref.name](data-sources--subnet--reference--group-001.md#canonical-2100332310132311-1020310213000120-3230312003112311-1022320300123312-2133312010333232-2032131212010333-0133231300132213-3021232233031030) |
| `connect_to_layer2.layer2_intf_ref.namespace` | [connect_to_layer2.layer2_intf_ref.namespace](data-sources--subnet--reference--group-001.md#canonical-3021011023220000-2100323133332322-2313020131011303-0302220133101033-2202132033200113-2032131200220131-0321011230301301-0233122301313223) |
| `connect_to_layer2.layer2_intf_ref.tenant` | [connect_to_layer2.layer2_intf_ref.tenant](data-sources--subnet--reference--group-001.md#canonical-2310121200111110-1132301113000021-2212321212303310-0020202010110013-2110300013201031-2212220223003233-0123002003133001-0113023222103013) |
| `connect_to_slo` | [connect_to_slo](data-sources--subnet--reference--group-001.md#canonical-3301020232002333-0331232233111203-2101320201330100-1102111013120132-2120101110233023-0123113321012133-1012103121310122-0011110031132313) |
| `description` | [description](data-sources--subnet--reference--group-001.md#canonical-0120331112023212-1323321311213113-1201020230100112-1333001021112201-1311231032230202-3333330032203000-2211213213032023-2203322331131032) |
| `id` | [ID](data-sources--subnet--reference--group-001.md#canonical-0100102201231020-3132032311212203-1332212212103013-3032033330121033-1000132000012011-0220221310302003-0001200301012332-3212212122301221) |
| `isolated_nw` | [isolated_nw](data-sources--subnet--reference--group-001.md#canonical-2132231233230111-2223232321331220-2212220020132312-1113011013133012-1122131223121303-0223323131132101-3220232231103332-1001001220113220) |
| `labels` | [labels](data-sources--subnet--reference--group-001.md#canonical-1023320103220121-0213331201121020-1102022111123101-2002230311000110-2032211022302300-1110020201123202-1131131201010300-0101131020331233) |
| `name` | [name](data-sources--subnet--reference--group-001.md#canonical-3220122013030311-3320033032332320-2320210112011123-2021210232233310-3021203120001321-3213233320012323-0322303203020121-2031213313321231) |
| `namespace` | [namespace](data-sources--subnet--reference--group-001.md#canonical-0103210311221213-3211132221003133-2131303021030301-0020133203310132-2010333223132130-1032232022012011-2302130121302211-1230110222110113) |
| `site_subnet_params` | [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-3030203301023330-1220023031110031-0101210123000031-3133002122021033-3200003121003200-3332013303011303-1310203100002231-3131233233230302) |
| `site_subnet_params.dhcp` | [site_subnet_params.dhcp](data-sources--subnet--reference--group-001.md#canonical-0010123111312301-2101111103333202-0113220213000223-2320311000030011-1100230231101321-1312210332111120-1011102312333300-2123333331201322) |
| `site_subnet_params.site` | [site_subnet_params.site](data-sources--subnet--reference--group-001.md#canonical-0003310222133001-2130122211102221-3101032020303312-1221022033020332-1210200232210133-2201011113332011-1231021010221332-1331301311313102) |
| `site_subnet_params.site.name` | [site_subnet_params.site.name](data-sources--subnet--reference--group-001.md#canonical-0101121313322000-0120221032000121-3012330332113032-1332211000110330-0323200222212303-1113202222201300-1333222103201221-0233233123202200) |
| `site_subnet_params.site.namespace` | [site_subnet_params.site.namespace](data-sources--subnet--reference--group-001.md#canonical-0111021233120021-1332023032311323-1330002113210021-2123003230120133-2102103023123012-0103001310030303-1320130203000323-1331011012302202) |
| `site_subnet_params.site.tenant` | [site_subnet_params.site.tenant](data-sources--subnet--reference--group-001.md#canonical-1121201320122133-0112222211330002-0112121232022223-1020232011012013-1002003211213302-1111132130232110-1031133222312033-2032031120303331) |
| `site_subnet_params.static_ip` | [site_subnet_params.static_ip](data-sources--subnet--reference--group-001.md#canonical-3230031101213122-0312113002311012-3202121220310023-3031231300200211-1031221311120313-1322312320200101-0131002021020121-1113330032302331) |
| `site_subnet_params.subnet_dhcp_server_params` | [site_subnet_params.subnet_dhcp_server_params](data-sources--subnet--reference--group-001.md#canonical-1131210323330203-3333203110220213-1333112031020003-0030103131030320-0000221231111332-2331020022310203-0223322030033101-3110230213110011) |
| `site_subnet_params.subnet_dhcp_server_params.dhcp_networks` | [site_subnet_params.subnet_dhcp_server_params.dhcp_networks](data-sources--subnet--reference--group-001.md#canonical-0013013011022013-2020232232223010-0013301232212133-1012232031002112-2330221020012233-0300333232300320-3112302120130213-0100233323330013) |
| `site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix` | [site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix](data-sources--subnet--reference--group-001.md#canonical-1021312230323301-2000021332032212-3030011302211131-2301302131023001-0323212031100223-2122300013201011-3003003210302301-3322321023110203) |

<a id="canonical-3130121321030111-2013113023133031-1213132113021010-2313220033022320-0122011323332121-3021003231200303-0213001233000220-0202302100030301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `connect_to_layer2` properties

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200)
- connect_to_layer2

<a id="canonical-0212332133031033-0100212001110111-1101213001010011-0102133302112323-2031022311203332-2332020210023133-0133300323320231-3002221132321311"></a>

Type: `"single"`. Computed.

\[OneOf: connect\_to\_layer2, connect\_to\_slo, isolated\_nw\] Configuration parameter for connect
to layer2.

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

- [connect_to_layer2](data-sources--subnet--reference--group-001.md#canonical-0212332133031033-0100212001110111-1101213001010011-0102133302112323-2031022311203332-2332020210023133-0133300323320231-3002221132321311)
- [connect_to_slo](data-sources--subnet--reference--group-001.md#canonical-3301020232002333-0331232233111203-2101320201330100-1102111013120132-2120101110233023-0123113321012133-1012103121310122-0011110031132313)
- [isolated_nw](data-sources--subnet--reference--group-001.md#canonical-2132231233230111-2223232321331220-2212220020132312-1113011013133012-1122131223121303-0223323131132101-3220232231103332-1001001220113220)

Select alternatives according to the provider validators above.

<a id="canonical-0100203302002021-2303102010132133-1322111322132020-3131111002210201-0320310101020011-0201310120022232-0101233121310112-1030301021303012"></a>

### Direct properties for `connect_to_layer2`

- [layer2_intf_ref](data-sources--subnet--reference--group-001.md#canonical-1303220020303222-3133310001313133-2202222323220201-1223213312303001-1023303301321120-3201303021213002-0230220012320201-2331021311031122): complete subsection reference.

<a id="canonical-1303220020303222-3133310001313133-2202222323220201-1223213312303001-1023303301321120-3201303021213002-0230220012320201-2331021311031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `connect_to_layer2.layer2_intf_ref` properties

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200)
- [connect_to_layer2](data-sources--subnet--reference--group-001.md#canonical-3130121321030111-2013113023133031-1213132113021010-2313220033022320-0122011323332121-3021003231200303-0213001233000220-0202302100030301)
- connect_to_layer2.layer2_intf_ref

<a id="canonical-3133012203233313-1231132232210032-0332022211100033-3312021130112001-2033032313212223-3221220010220302-2002120322212123-3211103022123323"></a>

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

<a id="canonical-3010111030000121-3311011310121021-1310201011210322-3223032223103311-2221032301301123-2111010121220002-0110003133203302-0331002210330312"></a>

### Direct properties for `connect_to_layer2.layer2_intf_ref`

<a id="canonical-2100332310132311-1020310213000120-3230312003112311-1022320300123312-2133312010333232-2032131212010333-0133231300132213-3021232233031030"></a>

#### `connect_to_layer2.layer2_intf_ref.name` property

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

<a id="canonical-3021011023220000-2100323133332322-2313020131011303-0302220133101033-2202132033200113-2032131200220131-0321011230301301-0233122301313223"></a>

<a id="canonical-0203122102230010-2001231232112011-0110120201110231-0003122121322312-2012201222000332-2323201032320011-3020303221132333-1210322110321123"></a>

#### `connect_to_layer2.layer2_intf_ref.namespace` property

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

<a id="canonical-2310121200111110-1132301113000021-2212321212303310-0020202010110013-2110300013201031-2212220223003233-0123002003133001-0113023222103013"></a>

<a id="canonical-1332203303211322-1121213210101200-2312031122122102-0330301213032021-1313212330123032-1113001311203013-1301012031010033-0322231311000303"></a>

#### `connect_to_layer2.layer2_intf_ref.tenant` property

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

<a id="canonical-1012300201312111-2313203113003101-3132331220322101-0023000203030112-0122320223110323-3330232223100323-1230223312031103-0333021321332002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `connect_to_slo` properties

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200)
- connect_to_slo

<a id="canonical-3301020232002333-0331232233111203-2101320201330100-1102111013120132-2120101110233023-0123113321012133-1012103121310122-0011110031132313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for connect to slo.

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

<a id="canonical-1102221302223332-3311312000120021-2332120301312013-3031323122310132-1201033003230231-3313320312210131-3130022322010033-1200213032110230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `isolated_nw` properties

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200)
- isolated_nw

<a id="canonical-2132231233230111-2223232321331220-2212220020132312-1113011013133012-1122131223121303-0223323131132101-3220232231103332-1001001220113220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for isolated nw.

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

<a id="canonical-3013003112313223-3230212001032101-0110030102302020-1201202301310330-3121320032220330-1233021100322013-3321021102221013-3132002132213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params` properties

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200)
- site_subnet_params

<a id="canonical-3030203301023330-1220023031110031-0101210123000031-3133002122021033-3200003121003200-3332013303011303-1310203100002231-3131233233230302"></a>

Type: `"list"`. Computed.

Site Subnet Parameters. Configure subnet parameters per site.

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

<a id="canonical-1033131232333311-1033312000100003-2331320223300012-1321123313012122-1112313022000222-0222002010212221-3121131003013322-2202333001233221"></a>

### Direct properties for `site_subnet_params`

- [dhcp](data-sources--subnet--reference--group-001.md#canonical-1201201102010321-3220010000003132-1020330030323302-1122321310310013-2223323011033120-3010332012211331-1213211132223231-0112321011031233): complete subsection reference.

- [site](data-sources--subnet--reference--group-001.md#canonical-1102300303022033-3030130331121330-2003030211011001-0203121220022121-2323133003002011-1011313300322232-1323100112311122-1331230302311203): complete subsection reference.

- [static_ip](data-sources--subnet--reference--group-001.md#canonical-2333020322313301-0121103002033132-0221313132103331-1013010000211203-0310021233203213-3301003131311310-1233211001303010-1122011322012003): complete subsection reference.

- [subnet_dhcp_server_params](data-sources--subnet--reference--group-001.md#canonical-2233201121022323-2002333311211310-3012010103312210-2213220232202332-1313320111021112-1332102200123102-1303013222113103-1110310333100110): complete subsection reference.

<a id="canonical-1201201102010321-3220010000003132-1020330030323302-1122321310310013-2223323011033120-3010332012211331-1213211132223231-0112321011031233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params.dhcp` properties

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-3013003112313223-3230212001032101-0110030102302020-1201202301310330-3121320032220330-1233021100322013-3321021102221013-3132002132213003)
- site_subnet_params.dhcp

<a id="canonical-0010123111312301-2101111103333202-0113220213000223-2320311000030011-1100230231101321-1312210332111120-1011102312333300-2123333331201322"></a>

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

<a id="canonical-1102300303022033-3030130331121330-2003030211011001-0203121220022121-2323133003002011-1011313300322232-1323100112311122-1331230302311203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params.site` properties

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-3013003112313223-3230212001032101-0110030102302020-1201202301310330-3121320032220330-1233021100322013-3321021102221013-3132002132213003)
- site_subnet_params.site

<a id="canonical-0003310222133001-2130122211102221-3101032020303312-1221022033020332-1210200232210133-2201011113332011-1231021010221332-1331301311313102"></a>

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

<a id="canonical-2201122103200221-2112312121030303-2111102122023130-3110231212111232-0122223130313211-1200130120011112-2010011133333112-1131223333221323"></a>

### Direct properties for `site_subnet_params.site`

<a id="canonical-0101121313322000-0120221032000121-3012330332113032-1332211000110330-0323200222212303-1113202222201300-1333222103201221-0233233123202200"></a>

#### `site_subnet_params.site.name` property

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

<a id="canonical-0111021233120021-1332023032311323-1330002113210021-2123003230120133-2102103023123012-0103001310030303-1320130203000323-1331011012302202"></a>

<a id="canonical-1210101313310110-1031320310133110-2020333132123133-0002323122201212-3133112222120110-3002002200213230-3022102232101002-0210103300133220"></a>

#### `site_subnet_params.site.namespace` property

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

<a id="canonical-1121201320122133-0112222211330002-0112121232022223-1020232011012013-1002003211213302-1111132130232110-1031133222312033-2032031120303331"></a>

<a id="canonical-0211331013010030-3221200101121032-2133033100302102-1032121312110222-2123212102000203-3033201002231103-2123121302113131-2212032130303120"></a>

#### `site_subnet_params.site.tenant` property

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

<a id="canonical-2333020322313301-0121103002033132-0221313132103331-1013010000211203-0310021233203213-3301003131311310-1233211001303010-1122011322012003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params.static_ip` properties

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-3013003112313223-3230212001032101-0110030102302020-1201202301310330-3121320032220330-1233021100322013-3321021102221013-3132002132213003)
- site_subnet_params.static_ip

<a id="canonical-3230031101213122-0312113002311012-3202121220310023-3031231300200211-1031221311120313-1322312320200101-0131002021020121-1113330032302331"></a>

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

<a id="canonical-2233201121022323-2002333311211310-3012010103312210-2213220232202332-1313320111021112-1332102200123102-1303013222113103-1110310333100110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params.subnet_dhcp_server_params` properties

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-3013003112313223-3230212001032101-0110030102302020-1201202301310330-3121320032220330-1233021100322013-3321021102221013-3132002132213003)
- site_subnet_params.subnet_dhcp_server_params

<a id="canonical-1131210323330203-3333203110220213-1333112031020003-0030103131030320-0000221231111332-2331020022310203-0223322030033101-3110230213110011"></a>

Type: `"single"`. Computed.

Subnet DHCP parameters will be a subset of network\_interface.dhcpserverparameterstype as all
features in network\_interface.dhcpserverparameterstype may not be supported in a subnet.

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

<a id="canonical-3312200203001021-1122233123320202-1113323322201031-1020000311021103-1303201212030100-3220102330332121-2320331300211023-2013211122133333"></a>

### Direct properties for `site_subnet_params.subnet_dhcp_server_params`

- [dhcp_networks](data-sources--subnet--reference--group-001.md#canonical-2031201330311312-2001113032012130-3001233233212132-0133221111022111-0102222310323013-1103012220321303-0103012321103020-1323030333030211): complete subsection reference.

<a id="canonical-2031201330311312-2001113032012130-3001233233212132-0133221111022111-0102222310323013-1103012220321303-0103012321103020-1323030333030211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_subnet_params.subnet_dhcp_server_params.dhcp_networks` properties

Breadcrumbs:

- [xcsh_subnet](../data-sources/subnet.md#canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210)
- [Property reference](data-sources--subnet--reference--group-001.md#canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200)
- [site_subnet_params](data-sources--subnet--reference--group-001.md#canonical-3013003112313223-3230212001032101-0110030102302020-1201202301310330-3121320032220330-1233021100322013-3321021102221013-3132002132213003)
- [site_subnet_params.subnet_dhcp_server_params](data-sources--subnet--reference--group-001.md#canonical-2233201121022323-2002333311211310-3012010103312210-2213220232202332-1313320111021112-1332102200123102-1303013222113103-1110310333100110)
- site_subnet_params.subnet_dhcp_server_params.dhcp_networks

<a id="canonical-0013013011022013-2020232232223010-0013301232212133-1012232031002112-2330221020012233-0300333232300320-3112302120130213-0100233323330013"></a>

Type: `"list"`. Computed.

List of networks from which DHCP server can allocate IP addresses.

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

<a id="canonical-0310332021210200-2323203223132330-1122031031031332-1323112311020322-2230103230002010-1230020300301022-2011201232312322-0100102213021210"></a>

### Direct properties for `site_subnet_params.subnet_dhcp_server_params.dhcp_networks`

<a id="canonical-1021312230323301-2000021332032212-3030011302211131-2301302131023001-0323212031100223-2122300013201011-3003003210302301-3322321023110203"></a>

#### `site_subnet_params.subnet_dhcp_server_params.dhcp_networks.network_prefix` property

Type: `"string"`. Computed.

Exclusive with \[\] Network prefix for subnet.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```
