---
page_title: "xcsh_app_api_group reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group reference."
---

# xcsh_app_api_group reference

<a id="canonical-3131003230120102-0133232130312100-3201332223003222-1312202223000112-1310030333311011-3211213212221220-1032213012023303-0012130201000232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010)
- Property reference

<a id="canonical-3212013020113213-3322201110203200-3132032123011303-2220021200323022-1311001202032303-0222100322331010-2312023121301331-3333033222313223"></a>

### Direct properties for `xcsh_app_api_group`

<a id="canonical-1220100132312311-2103330223022121-3303003220102103-2113112300102231-2100211022202202-0010222331110322-2112010121001330-1133002002230321"></a>

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

- [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-1213011112130221-2112210023332111-0122001122030111-3000113002222221-0222313333311230-0112200212111121-3203220023033102-3020320221222303): complete subsection reference.

- [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-1122213102230112-1112022003110130-3220331113222231-1102000331113300-0311030101111100-2010310330333033-3103121011001321-3320100122100111): complete subsection reference.

<a id="canonical-2121133123230210-3203002333110310-1311032131033113-0212301302303133-2011300002233232-2203120031311020-2200320000020123-0133032012310133"></a>

<a id="canonical-1202232112302320-2313300113312020-1111302301102012-2000333003023003-1103233031323102-3232303023232221-1021310130001310-0013001300221333"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the AppAPIGroup.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [elements](data-sources--app_api_group--reference--group-001.md#canonical-1202110301212310-3021113332133021-2031201131331320-0121110133320022-0312131300213201-2112323020001113-2232203201223111-3033200313231311): complete subsection reference.

- [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-1012330221310120-1213213021310010-2120320000001322-3233313300332303-2221221001101022-1013130323212102-1323232121331022-2210223203322322): complete subsection reference.

<a id="canonical-0320103320023210-3321323123221001-1102223310300011-1322022000123233-0330320303232120-0123101111133023-0300131320312313-3302131203003121"></a>

<a id="canonical-0310302031212101-2011233122032110-1223233202010222-0202303333200012-2130002032331031-2033102003331022-0332033021121312-3130133320301132"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0000023222221213-3030321100030331-0002103221102132-2013320223303313-1123302112221111-0230300233023211-2202022210133033-0031200203302202"></a>

<a id="canonical-1210210302222121-0102100331021201-3101232122201010-1222120330301311-1120122121001031-0102103020200213-2310312021110033-3301210330111002"></a>

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

<a id="canonical-2000323121311131-3101120103123022-2100312201323300-3113203013010330-2311220033010332-2111132032101023-0033211121333232-2213100203030020"></a>

<a id="canonical-1033301203303232-2022030031220002-0233233301211131-1313201102033201-3212301102132030-1201000222330213-2002033000302221-1101323323232233"></a>

#### `name` property

Type: `"string"`. Required.

Name of the AppAPIGroup.

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

<a id="canonical-2101031103133131-2100330010211222-1001032222212101-2322311301031031-2012301311223321-2003112213302323-2332020220333302-0212210020301211"></a>

<a id="canonical-2132122300321020-1102123230031122-0322032231111011-3130321200023202-1011030003110332-3100300021210202-2302022222031130-3313302012330221"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the AppAPIGroup exists.

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

<a id="canonical-1223200200010210-1230111133123310-0312332211021122-2031130110030233-2002111210103032-0130211120003030-2033312101023230-2033221330230131"></a>

### All schema paths for `xcsh_app_api_group`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--app_api_group--reference--group-001.md#canonical-1220100132312311-2103330223022121-3303003220102103-2113112300102231-2100211022202202-0010222331110322-2112010121001330-1133002002230321) |
| `bigip_virtual_server` | [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-1223020121013002-3231002112302121-2222001003103221-0010123020321013-3332302321230223-3311031223102200-3200002001013321-1311001033200032) |
| `bigip_virtual_server.bigip_virtual_server` | [bigip_virtual_server.bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-2320122302222322-0302131200003223-2010130122010112-2321030102210231-1230323131312331-0123100020222121-2131122231212022-0233323223233233) |
| `bigip_virtual_server.bigip_virtual_server.name` | [bigip_virtual_server.bigip_virtual_server.name](data-sources--app_api_group--reference--group-001.md#canonical-1220332313302220-0221001230302321-2300112202232010-2003102121221101-3003021311300112-2032331012033132-0213221123200013-3112121330001320) |
| `bigip_virtual_server.bigip_virtual_server.namespace` | [bigip_virtual_server.bigip_virtual_server.namespace](data-sources--app_api_group--reference--group-001.md#canonical-1031020311003221-2323321322221212-2332331202020131-3310021202111321-3103100133033333-3220211202222330-0210202321202200-2010322121102103) |
| `bigip_virtual_server.bigip_virtual_server.tenant` | [bigip_virtual_server.bigip_virtual_server.tenant](data-sources--app_api_group--reference--group-001.md#canonical-3001133120123303-1312022210300001-1212111323002011-2200331210120030-2020013002331120-0133322233130302-2121131133021300-3221321311101031) |
| `cdn_loadbalancer` | [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-0003010322012111-0031222322221101-0322200233230320-2012300303232121-0203310021102132-3000130013013221-2231202011120100-2210010100020323) |
| `cdn_loadbalancer.cdn_loadbalancer` | [cdn_loadbalancer.cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-0322010222211232-0133233303010030-1003020311131133-3301322132031131-1311030230220302-3213230102312330-2301003211000232-3001013020222202) |
| `cdn_loadbalancer.cdn_loadbalancer.name` | [cdn_loadbalancer.cdn_loadbalancer.name](data-sources--app_api_group--reference--group-001.md#canonical-0000133310133223-1110302113000332-3232303222313202-3022013111211203-3200303322302310-3332111130202110-0322200312300022-2313113012212200) |
| `cdn_loadbalancer.cdn_loadbalancer.namespace` | [cdn_loadbalancer.cdn_loadbalancer.namespace](data-sources--app_api_group--reference--group-001.md#canonical-0020210330130032-0120200013203101-1100203202210303-0003323110200213-1022312031331113-3321322320121102-1010332323230112-2302312123213310) |
| `cdn_loadbalancer.cdn_loadbalancer.tenant` | [cdn_loadbalancer.cdn_loadbalancer.tenant](data-sources--app_api_group--reference--group-001.md#canonical-2132201322213203-2302330323010023-1222011300031321-0103003223102302-1021233202133120-1013123233130033-0230012003022323-0013033222131033) |
| `description` | [description](data-sources--app_api_group--reference--group-001.md#canonical-2121133123230210-3203002333110310-1311032131033113-0212301302303133-2011300002233232-2203120031311020-2200320000020123-0133032012310133) |
| `elements` | [elements](data-sources--app_api_group--reference--group-001.md#canonical-2223232220012130-3103011132103021-0020132313221020-0332322001011203-1230103120211130-2010220310203123-0213201202232012-1021122202022202) |
| `elements.methods` | [elements.methods](data-sources--app_api_group--reference--group-001.md#canonical-1311320122323033-2221331303102300-1012232110220011-0012322101222013-1331221112232332-1012331211223221-3311221313110321-0013201003200013) |
| `elements.path_regex` | [elements.path_regex](data-sources--app_api_group--reference--group-001.md#canonical-0130230000213300-1000310131213223-2212231321301022-1301030120201321-1212321213133122-0122013022030332-3310030003120112-3011202102323312) |
| `http_loadbalancer` | [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-0110232220021222-2023123102302302-0123301223031232-1301311022022330-2213001022331012-3333332322310221-1200132031031021-3012210000121010) |
| `http_loadbalancer.http_loadbalancer` | [http_loadbalancer.http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-1310111031130231-1122102212312313-1130103102230103-2320133020222101-1310011011201032-3131213330332322-2123312310012222-2021003032200211) |
| `http_loadbalancer.http_loadbalancer.name` | [http_loadbalancer.http_loadbalancer.name](data-sources--app_api_group--reference--group-001.md#canonical-0033211102313033-1032111010213120-1312231023212310-1020222031022332-3012112210233323-3332311133122313-1110210103113223-0201320203012023) |
| `http_loadbalancer.http_loadbalancer.namespace` | [http_loadbalancer.http_loadbalancer.namespace](data-sources--app_api_group--reference--group-001.md#canonical-3202313221203301-2122020213310122-2122011022301301-0330012223120012-2222212332233323-3232103033223220-2010030112002202-2032230022033111) |
| `http_loadbalancer.http_loadbalancer.tenant` | [http_loadbalancer.http_loadbalancer.tenant](data-sources--app_api_group--reference--group-001.md#canonical-1123201312122013-2201121023233222-1110113131331122-2101210032210223-2221001113123113-2220302213320310-2322133311321132-1011333322212022) |
| `id` | [ID](data-sources--app_api_group--reference--group-001.md#canonical-0320103320023210-3321323123221001-1102223310300011-1322022000123233-0330320303232120-0123101111133023-0300131320312313-3302131203003121) |
| `labels` | [labels](data-sources--app_api_group--reference--group-001.md#canonical-0000023222221213-3030321100030331-0002103221102132-2013320223303313-1123302112221111-0230300233023211-2202022210133033-0031200203302202) |
| `name` | [name](data-sources--app_api_group--reference--group-001.md#canonical-2000323121311131-3101120103123022-2100312201323300-3113203013010330-2311220033010332-2111132032101023-0033211121333232-2213100203030020) |
| `namespace` | [namespace](data-sources--app_api_group--reference--group-001.md#canonical-2101031103133131-2100330010211222-1001032222212101-2322311301031031-2012301311223321-2003112213302323-2332020220333302-0212210020301211) |

<a id="canonical-1213011112130221-2112210023332111-0122001122030111-3000113002222221-0222313333311230-0112200212111121-3203220023033102-3020320221222303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bigip_virtual_server` properties

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-3131003230120102-0133232130312100-3201332223003222-1312202223000112-1310030333311011-3211213212221220-1032213012023303-0012130201000232)
- bigip_virtual_server

<a id="canonical-1223020121013002-3231002112302121-2222001003103221-0010123020321013-3332302321230223-3311031223102200-3200002001013321-1311001033200032"></a>

Type: `"single"`. Computed.

\[OneOf: bigip\_virtual\_server, cdn\_loadbalancer, http\_loadbalancer\] Set the scope of the API
Group to a specific BIG-IP Virtual Server.

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

- [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-1223020121013002-3231002112302121-2222001003103221-0010123020321013-3332302321230223-3311031223102200-3200002001013321-1311001033200032)
- [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-0003010322012111-0031222322221101-0322200233230320-2012300303232121-0203310021102132-3000130013013221-2231202011120100-2210010100020323)
- [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-0110232220021222-2023123102302302-0123301223031232-1301311022022330-2213001022331012-3333332322310221-1200132031031021-3012210000121010)

Select alternatives according to the provider validators above.

<a id="canonical-1332303320030133-2211121232012323-3213102022122112-2311131212113233-2112213023332032-0220330222002231-1302120302210332-0000023313203113"></a>

### Direct properties for `bigip_virtual_server`

- [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-2031003203311230-2022133313000313-2023011301332010-0112320002232100-3331102323032031-3120132331323232-0322220302001302-2111100123212030): complete subsection reference.

<a id="canonical-2031003203311230-2022133313000313-2023011301332010-0112320002232100-3331102323032031-3120132331323232-0322220302001302-2111100123212030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bigip_virtual_server.bigip_virtual_server` properties

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-3131003230120102-0133232130312100-3201332223003222-1312202223000112-1310030333311011-3211213212221220-1032213012023303-0012130201000232)
- [bigip_virtual_server](data-sources--app_api_group--reference--group-001.md#canonical-1213011112130221-2112210023332111-0122001122030111-3000113002222221-0222313333311230-0112200212111121-3203220023033102-3020320221222303)
- bigip_virtual_server.bigip_virtual_server

<a id="canonical-2320122302222322-0302131200003223-2010130122010112-2321030102210231-1230323131312331-0123100020222121-2131122231212022-0233323223233233"></a>

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

<a id="canonical-3031023010302000-3000221223213303-3120222020200020-0022221320333101-1010022121313033-3223210231323332-0100220120031232-3010112031130321"></a>

### Direct properties for `bigip_virtual_server.bigip_virtual_server`

<a id="canonical-1220332313302220-0221001230302321-2300112202232010-2003102121221101-3003021311300112-2032331012033132-0213221123200013-3112121330001320"></a>

#### `bigip_virtual_server.bigip_virtual_server.name` property

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

<a id="canonical-1031020311003221-2323321322221212-2332331202020131-3310021202111321-3103100133033333-3220211202222330-0210202321202200-2010322121102103"></a>

<a id="canonical-2110030110320012-3300121303201222-2133311012021200-3220120111230200-2302003101030331-3113220200100103-2011333201021000-3120031313302120"></a>

#### `bigip_virtual_server.bigip_virtual_server.namespace` property

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

<a id="canonical-3001133120123303-1312022210300001-1212111323002011-2200331210120030-2020013002331120-0133322233130302-2121131133021300-3221321311101031"></a>

<a id="canonical-1210012032023023-3123300032100113-0123122303103100-2123023313133013-2112001222011121-2212213200333101-2233020220102210-1121230002002000"></a>

#### `bigip_virtual_server.bigip_virtual_server.tenant` property

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

<a id="canonical-1122213102230112-1112022003110130-3220331113222231-1102000331113300-0311030101111100-2010310330333033-3103121011001321-3320100122100111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cdn_loadbalancer` properties

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-3131003230120102-0133232130312100-3201332223003222-1312202223000112-1310030333311011-3211213212221220-1032213012023303-0012130201000232)
- cdn_loadbalancer

<a id="canonical-0003010322012111-0031222322221101-0322200233230320-2012300303232121-0203310021102132-3000130013013221-2231202011120100-2210010100020323"></a>

Type: `"single"`. Computed.

Set the scope of the API Group to a specific CDN Loadbalancer.

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

<a id="canonical-2103132302330202-3222221232103013-2011010211313213-2200002122312132-0302223333320030-3302021033203232-1103001130221301-0200123013302220"></a>

### Direct properties for `cdn_loadbalancer`

- [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-3012223101332221-1012230031113201-3122111221000120-1111121020001131-2300230022312333-3211132020212331-3002313103120120-0303113013012031): complete subsection reference.

<a id="canonical-3012223101332221-1012230031113201-3122111221000120-1111121020001131-2300230022312333-3211132020212331-3002313103120120-0303113013012031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cdn_loadbalancer.cdn_loadbalancer` properties

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-3131003230120102-0133232130312100-3201332223003222-1312202223000112-1310030333311011-3211213212221220-1032213012023303-0012130201000232)
- [cdn_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-1122213102230112-1112022003110130-3220331113222231-1102000331113300-0311030101111100-2010310330333033-3103121011001321-3320100122100111)
- cdn_loadbalancer.cdn_loadbalancer

<a id="canonical-0322010222211232-0133233303010030-1003020311131133-3301322132031131-1311030230220302-3213230102312330-2301003211000232-3001013020222202"></a>

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

<a id="canonical-1000223032123203-2333300233000201-1123213023120223-0103103210002202-2000013302232123-0223002130312022-3030223010313031-0113221312123112"></a>

### Direct properties for `cdn_loadbalancer.cdn_loadbalancer`

<a id="canonical-0000133310133223-1110302113000332-3232303222313202-3022013111211203-3200303322302310-3332111130202110-0322200312300022-2313113012212200"></a>

#### `cdn_loadbalancer.cdn_loadbalancer.name` property

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

<a id="canonical-0020210330130032-0120200013203101-1100203202210303-0003323110200213-1022312031331113-3321322320121102-1010332323230112-2302312123213310"></a>

<a id="canonical-3021103323313202-0020313122010032-3321232333102211-1031122021032003-0233110030231303-0232223321313201-3220031022023300-3130022000022303"></a>

#### `cdn_loadbalancer.cdn_loadbalancer.namespace` property

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

<a id="canonical-2132201322213203-2302330323010023-1222011300031321-0103003223102302-1021233202133120-1013123233130033-0230012003022323-0013033222131033"></a>

<a id="canonical-0201201101211313-3202103223123001-1310103320331002-1211312130113011-1102023313232023-2311200031231203-1123022200100033-1101012121311213"></a>

#### `cdn_loadbalancer.cdn_loadbalancer.tenant` property

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

<a id="canonical-1202110301212310-3021113332133021-2031201131331320-0121110133320022-0312131300213201-2112323020001113-2232203201223111-3033200313231311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `elements` properties

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-3131003230120102-0133232130312100-3201332223003222-1312202223000112-1310030333311011-3211213212221220-1032213012023303-0012130201000232)
- elements

<a id="canonical-2223232220012130-3103011132103021-0020132313221020-0332322001011203-1230103120211130-2010220310203123-0213201202232012-1021122202022202"></a>

Type: `"list"`. Computed.

List of API group elements with methods and path regular expression for matching requests.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5000"
  }
}
```

<a id="canonical-3221131310110113-2011302212121322-0132210001030332-3312101022230232-3010021031331111-2220110233311113-0323030031102220-2203100132203012"></a>

### Direct properties for `elements`

<a id="canonical-1311320122323033-2221331303102300-1012232110220011-0012322101222013-1331221112232332-1012331211223221-3311221313110321-0013201003200013"></a>

#### `elements.methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of method values to
match the input request API method against. The match is considered to succeed if the input request
API method is a member of the list. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`,
\`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "0",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "0",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0130230000213300-1000310131213223-2212231321301022-1301030120201321-1212321213133122-0122013022030332-3310030003120112-3011202102323312"></a>

<a id="canonical-0013201120313111-3103212031023200-3101113121332123-1302211333233031-1221323123300102-0223122330310300-3011333331013321-1230010201312221"></a>

#### `elements.path_regex` property

Type: `"string"`. Computed.

Regular expression to match the input request API path against. The match is considered to succeed
if the input request API path matches the specified path regular expression.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1012330221310120-1213213021310010-2120320000001322-3233313300332303-2221221001101022-1013130323212102-1323232121331022-2210223203322322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_loadbalancer` properties

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-3131003230120102-0133232130312100-3201332223003222-1312202223000112-1310030333311011-3211213212221220-1032213012023303-0012130201000232)
- http_loadbalancer

<a id="canonical-0110232220021222-2023123102302302-0123301223031232-1301311022022330-2213001022331012-3333332322310221-1200132031031021-3012210000121010"></a>

Type: `"single"`. Computed.

Set the scope of the API Group to a specific HTTP Loadbalancer.

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

<a id="canonical-2311130030020230-3220012031130321-3021110112233013-0330212021233103-2313102212212312-0321113033000231-2332223121020323-2323203303323103"></a>

### Direct properties for `http_loadbalancer`

- [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-3111322020132302-2130122200112022-2202332332000003-1321021031213222-2331101000111212-2110112322231121-0123231130221132-1111320122330321): complete subsection reference.

<a id="canonical-3111322020132302-2130122200112022-2202332332000003-1321021031213222-2331101000111212-2110112322231121-0123231130221132-1111320122330321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_loadbalancer.http_loadbalancer` properties

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010)
- [Property reference](data-sources--app_api_group--reference--group-001.md#canonical-3131003230120102-0133232130312100-3201332223003222-1312202223000112-1310030333311011-3211213212221220-1032213012023303-0012130201000232)
- [http_loadbalancer](data-sources--app_api_group--reference--group-001.md#canonical-1012330221310120-1213213021310010-2120320000001322-3233313300332303-2221221001101022-1013130323212102-1323232121331022-2210223203322322)
- http_loadbalancer.http_loadbalancer

<a id="canonical-1310111031130231-1122102212312313-1130103102230103-2320133020222101-1310011011201032-3131213330332322-2123312310012222-2021003032200211"></a>

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

<a id="canonical-1301222222110110-0021111103033121-2112031022033020-2033313110130020-0132121222122000-0013111313130033-0033120213233132-0000232020111012"></a>

### Direct properties for `http_loadbalancer.http_loadbalancer`

<a id="canonical-0033211102313033-1032111010213120-1312231023212310-1020222031022332-3012112210233323-3332311133122313-1110210103113223-0201320203012023"></a>

#### `http_loadbalancer.http_loadbalancer.name` property

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

<a id="canonical-3202313221203301-2122020213310122-2122011022301301-0330012223120012-2222212332233323-3232103033223220-2010030112002202-2032230022033111"></a>

<a id="canonical-3321003120001233-3210213002103131-3332100302101301-1120033311332001-1212223112133111-0330123123232333-2330003322032313-2231012130000031"></a>

#### `http_loadbalancer.http_loadbalancer.namespace` property

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

<a id="canonical-1123201312122013-2201121023233222-1110113131331122-2101210032210223-2221001113123113-2220302213320310-2322133311321132-1011333322212022"></a>

<a id="canonical-0303022200321121-2203212202331131-0102203322333313-1211101322220030-0002000123012010-0110232230031232-1322220221021013-0032300002323230"></a>

#### `http_loadbalancer.http_loadbalancer.tenant` property

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
