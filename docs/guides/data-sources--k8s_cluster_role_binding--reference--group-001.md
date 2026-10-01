---
page_title: "xcsh_k8s_cluster_role_binding reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role_binding reference."
---

# xcsh_k8s_cluster_role_binding reference

<a id="canonical-2330203132111311-2030223031030132-1113210123013110-2331032322332103-0220202113111221-2130320110300133-2230301113102200-0202032220222113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133331010211122-0000002310222320-1011023031102101-3011321202103101-3313323131001021-3311123132321100-2322231303012322-3233200322322333"></a>

## Property reference — Property reference / 213133331323 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-2021222210012102-1330130200013102-0220323113122312-2121230113313023-1123103021232230-1310303332323332-2100321301113012-3031120031031013)
- Property reference

<a id="canonical-2023213331221333-2301121102003022-2101333320113211-2233032201330000-0022233111110101-3310223013103302-1012023230301121-2033022102121001"></a>

## Direct properties — Property reference / 213133331323 / 3

<a id="canonical-1321313120202133-1300331130113031-1030020313131311-0103032333133220-2122003121133101-0031032210200332-1000023203303212-2302232112010012"></a>

<a id="canonical-2331323032320023-3301113222321331-0013211132322111-2122312202203132-0133120220130102-3032132033232122-3023200221302312-2302000032013013"></a>

## annotations property — Property reference / 213133331323 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-0313212213010311-3220332210112120-0213333221002202-3231210303332310-1222132003010000-3012202021210313-3133203133012300-2120331331011123"></a>

<a id="canonical-0021123012322110-0131123232121203-2022022012230323-1301033313320003-3220233000211112-3222201230002010-1301100302022023-0220113300003330"></a>

## description property — Property reference / 213133331323 / 5

Type: `"string"`. Computed.

Description of the K8SClusterRoleBinding.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0213131221200211-2013112033210033-1112121013020111-0301320023112223-2223313331022323-2210103300330012-3133323320302231-1022302033213313"></a>

<a id="canonical-2031101210110001-3212220103010032-1021332212311220-2322113011303231-3201320201321011-1202011200010131-1032002110211111-2333330300321311"></a>

## ID property — Property reference / 213133331323 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster_role](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-1301011003213113-0310111020223311-1320322331302332-3213321313333133-2010032200222001-2100000121133333-1322030233313033-0101322133020001): complete subsection reference.

<a id="canonical-3022102201012121-2212313213320111-0200333030003230-1130200233020122-1011130030233213-0232333012010211-1001300112031312-3312222311030213"></a>

<a id="canonical-3031212302203133-1233123330231020-2212232022020233-3011331311211010-1012003120331003-3011030002113330-3113113212303130-3022303300323012"></a>

## labels property — Property reference / 213133331323 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

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

<a id="canonical-1320112312312201-1031233202330101-1033112321120112-3031213013031121-2331031130302102-2003300123133022-2101312321001121-3112100130222222"></a>

<a id="canonical-1123122120200312-2230230302233203-1102103002221010-2120012113221132-0130210301112130-2111020132022330-3133322321012123-3223231001112010"></a>

## name property — Property reference / 213133331323 / 8

Type: `"string"`. Required.

Name of the K8SClusterRoleBinding.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2221132122020212-3200032102113202-0300013011121223-2031202123121231-1012323211330100-0120312132021132-0210032121100233-3222123330200310"></a>

<a id="canonical-3312232102011131-0022203222333031-1031100130010310-0122230212023013-1133231113001300-2130320021110311-3123332321200310-2111330322233311"></a>

## namespace property — Property reference / 213133331323 / 9

Type: `"string"`. Required.

Namespace where the K8SClusterRoleBinding exists.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [subjects](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-1203213300130313-1121302211223313-1022100131210132-3302123311203132-0000133220133122-1230322013121132-0322311121121123-1200020310023310): complete subsection reference.

<a id="canonical-2201123010231223-3202020001203113-3321121333012312-1302032031000022-1232010000022132-1011021223022121-0111230020110201-3320331310223101"></a>

## All schema paths — Property reference / 213133331323 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-1321313120202133-1300331130113031-1030020313131311-0103032333133220-2122003121133101-0031032210200332-1000023203303212-2302232112010012) |
| `description` | [description](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-0313212213010311-3220332210112120-0213333221002202-3231210303332310-1222132003010000-3012202021210313-3133203133012300-2120331331011123) |
| `id` | [id](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-0213131221200211-2013112033210033-1112121013020111-0301320023112223-2223313331022323-2210103300330012-3133323320302231-1022302033213313) |
| `k8s_cluster_role` | [k8s_cluster_role](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-3312103123312203-3213201321123032-0212233100301322-3210333233212222-1111302033310232-0233102123123322-1223103323000031-3021020230132322) |
| `k8s_cluster_role.name` | [k8s_cluster_role.name](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2333001333100113-1320110012131311-2211011332131030-3010033023131021-1211202122121022-3321233211331000-0001021221133223-0103033300023102) |
| `k8s_cluster_role.namespace` | [k8s_cluster_role.namespace](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-0112110032110112-1222303001123311-1331122310120033-2102221322120121-0112111013100102-1111121211021010-2312010300021011-2232321333023120) |
| `k8s_cluster_role.tenant` | [k8s_cluster_role.tenant](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2323211300112302-1120122112313222-1213000121102212-3001202010221201-1231131031220200-0012032200003212-1302300123020311-1323013232323030) |
| `labels` | [labels](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-3022102201012121-2212313213320111-0200333030003230-1130200233020122-1011130030233213-0232333012010211-1001300112031312-3312222311030213) |
| `name` | [name](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-1320112312312201-1031233202330101-1033112321120112-3031213013031121-2331031130302102-2003300123133022-2101312321001121-3112100130222222) |
| `namespace` | [namespace](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2221132122020212-3200032102113202-0300013011121223-2031202123121231-1012323211330100-0120312132021132-0210032121100233-3222123330200310) |
| `subjects` | [subjects](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2310211110102212-3011101001112102-3311222111333211-1120033022133220-3123333323213003-1310300033230233-0231213312213322-0320333111022011) |
| `subjects.group` | [subjects.group](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-0002001033102000-0211320332003021-2103101100210323-3123130312030302-2020121003132213-0221333202213120-3231233223330323-3001202103032210) |
| `subjects.service_account` | [subjects.service_account](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-1020031133312030-0001222330120210-2021301320332033-1211333203122101-0132300321200110-3202223120101032-3220323313202223-0202331100301031) |
| `subjects.service_account.name` | [subjects.service_account.name](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2000221331123032-1223231101232301-0312111200013312-1020213012023310-3230012103000013-3322130030103120-1223303102030310-0311310321200300) |
| `subjects.service_account.namespace` | [subjects.service_account.namespace](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-3113011011220100-3310223022310311-3233103122123321-1010020000203331-1022013232233100-0000323032200001-1221010003230333-3102320221102132) |
| `subjects.user` | [subjects.user](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2212132102130233-0222022230132113-1010120303320131-2120213232222010-1331100103231222-1121010201313102-3020332203131220-0231101132021131) |

<a id="canonical-1223022132131133-2201220210022322-1003331330013120-3323223120010203-2202020101220232-1311312022323120-0123222210231221-1122130123212120"></a>

## Next pages — Property reference / 213133331323 / 11

- [k8s_cluster_role](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-1301011003213113-0310111020223311-1320322331302332-3213321313333133-2010032200222001-2100000121133333-1322030233313033-0101322133020001)
- [subjects](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-1203213300130313-1121302211223313-1022100131210132-3302123311203132-0000133220133122-1230322013121132-0322311121121123-1200020310023310)
- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-2021222210012102-1330130200013102-0220323113122312-2121230113313023-1123103021232230-1310303332323332-2100321301113012-3031120031031013)

<a id="canonical-1301011003213113-0310111020223311-1320322331302332-3213321313333133-2010032200222001-2100000121133333-1322030233313033-0101322133020001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211232002223222-1203032322022301-2212322222201101-3302030331000120-0220323213012010-1300133231321320-1030100200223113-0330333202300222"></a>

## k8s_cluster_role — k8s_cluster_role / 300021120112 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-2021222210012102-1330130200013102-0220323113122312-2121230113313023-1123103021232230-1310303332323332-2100321301113012-3031120031031013)
- [Property reference](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2330203132111311-2030223031030132-1113210123013110-2331032322332103-0220202113111221-2130320110300133-2230301113102200-0202032220222113)
- k8s_cluster_role

<a id="canonical-3312103123312203-3213201321123032-0212233100301322-3210333233212222-1111302033310232-0233102123123322-1223103323000031-3021020230132322"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-3220332201123112-1023110300223123-3002321311102212-2330331223001113-0203000300320333-0232012113320121-3032201202130230-0102201312222201"></a>

## Direct properties — k8s_cluster_role / 300021120112 / 3

<a id="canonical-2333001333100113-1320110012131311-2211011332131030-3010033023131021-1211202122121022-3321233211331000-0001021221133223-0103033300023102"></a>

<a id="canonical-2300332031002200-1132330032202112-2101330000130010-0020232312102111-2233301103220021-0231112022311103-3010310032032230-3132222212023310"></a>

## name property — k8s_cluster_role / 300021120112 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0112110032110112-1222303001123311-1331122310120033-2102221322120121-0112111013100102-1111121211021010-2312010300021011-2232321333023120"></a>

<a id="canonical-2121132303133101-3012110300013223-1101303333210122-0101322202300211-1102033021230202-0321232311103132-2033300033103110-1312310110100002"></a>

## namespace property — k8s_cluster_role / 300021120112 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2323211300112302-1120122112313222-1213000121102212-3001202010221201-1231131031220200-0012032200003212-1302300123020311-1323013232323030"></a>

<a id="canonical-0332003330022313-0010111323213132-0301110102330002-2003111132112001-3022003122000333-1010030120320233-1232121103323331-2100302310311002"></a>

## tenant property — k8s_cluster_role / 300021120112 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3011002232323031-3322120031120120-1201022103121303-2022013120100122-2102120212101002-1010202001130030-2313122320000332-2131120123022232"></a>

## Next pages — k8s_cluster_role / 300021120112 / 7

- [Property reference](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2330203132111311-2030223031030132-1113210123013110-2331032322332103-0220202113111221-2130320110300133-2230301113102200-0202032220222113)
- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-2021222210012102-1330130200013102-0220323113122312-2121230113313023-1123103021232230-1310303332323332-2100321301113012-3031120031031013)

<a id="canonical-1203213300130313-1121302211223313-1022100131210132-3302123311203132-0000133220133122-1230322013121132-0322311121121123-1200020310023310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121332102300231-1113212213101000-3231021123320002-1232302120221230-0012122112013202-3311230323322311-2312331102100333-3300102023002112"></a>

## subjects — subjects / 210023111300 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-2021222210012102-1330130200013102-0220323113122312-2121230113313023-1123103021232230-1310303332323332-2100321301113012-3031120031031013)
- [Property reference](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2330203132111311-2030223031030132-1113210123013110-2331032322332103-0220202113111221-2130320110300133-2230301113102200-0202032220222113)
- subjects

<a id="canonical-2310211110102212-3011101001112102-3311222111333211-1120033022133220-3123333323213003-1310300033230233-0231213312213322-0320333111022011"></a>

Type: `"list"`. Computed.

List of subjects (user, group or service account) to which this role is bound.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0102030123111030-0213032021032113-0302230112001110-2303223013213121-2303211102011122-0330010001221003-2013112023022221-1223310231233010"></a>

## Direct properties — subjects / 210023111300 / 3

<a id="canonical-0002001033102000-0211320332003021-2103101100210323-3123130312030302-2020121003132213-0221333202213120-3231233223330323-3001202103032210"></a>

<a id="canonical-2130133112033110-1320123113332223-2322123211010000-3113131131122010-0301031011330130-2302020002211020-2212020102031131-2213230100111000"></a>

## group property — subjects / 210023111300 / 4

Type: `"string"`. Computed.

Exclusive with \[service\_account user\] Group ID of the user group.

Upstream description:

Exclusive with \[service\_account user\] Group ID of the user group.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [service_account](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-0302320213122131-2132210213022232-2313232022002210-2212003010331030-0332012002312021-2102222313221011-3212202210223320-0301303332311332): complete subsection reference.

<a id="canonical-2212132102130233-0222022230132113-1010120303320131-2120213232222010-1331100103231222-1121010201313102-3020332203131220-0231101132021131"></a>

<a id="canonical-0131022031002213-0032021230130330-2302321113320023-1030322103113022-1003330103123333-1230013210110313-0011111121232230-1023101111223223"></a>

## user property — subjects / 210023111300 / 5

Type: `"string"`. Computed.

Exclusive with \[group service\_account\] User ID of the user.

Upstream description:

Exclusive with \[group service\_account\] User ID of the user.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1002300322230123-1102122303012111-1120013231332122-1101033202221221-2000130323201232-1013323311321231-0312131220301203-3012123202311330"></a>

## Next pages — subjects / 210023111300 / 6

- [subjects.service_account](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-0302320213122131-2132210213022232-2313232022002210-2212003010331030-0332012002312021-2102222313221011-3212202210223320-0301303332311332)
- [Property reference](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2330203132111311-2030223031030132-1113210123013110-2331032322332103-0220202113111221-2130320110300133-2230301113102200-0202032220222113)
- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-2021222210012102-1330130200013102-0220323113122312-2121230113313023-1123103021232230-1310303332323332-2100321301113012-3031120031031013)

<a id="canonical-0302320213122131-2132210213022232-2313232022002210-2212003010331030-0332012002312021-2102222313221011-3212202210223320-0301303332311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003030220200323-0213013302002131-2320012030033221-2012123010310032-2201210231220321-0333010130132032-3322111222113021-0110002001221031"></a>

## subjects.service_account — service_account / 112132211033 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-2021222210012102-1330130200013102-0220323113122312-2121230113313023-1123103021232230-1310303332323332-2100321301113012-3031120031031013)
- [Property reference](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2330203132111311-2030223031030132-1113210123013110-2331032322332103-0220202113111221-2130320110300133-2230301113102200-0202032220222113)
- [subjects](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-1203213300130313-1121302211223313-1022100131210132-3302123311203132-0000133220133122-1230322013121132-0322311121121123-1200020310023310)
- subjects.service_account

<a id="canonical-1020031133312030-0001222330120210-2021301320332033-1211333203122101-0132300321200110-3202223120101032-3220323313202223-0202331100301031"></a>

Type: `"single"`. Computed.

ServiceAccountType.

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

<a id="canonical-2210302120233203-1330321233232332-3022102122332131-0231200300030033-0313332123123112-3112332001003203-3203233211123322-1103300330032102"></a>

## Direct properties — service_account / 112132211033 / 3

<a id="canonical-2000221331123032-1223231101232301-0312111200013312-1020213012023310-3230012103000013-3322130030103120-1223303102030310-0311310321200300"></a>

<a id="canonical-2321221202120230-0301012313011320-2111121212212233-2002003000100232-0031221321130213-1001333233220032-3120112221101033-1301131003120033"></a>

## name property — service_account / 112132211033 / 4

Type: `"string"`. Computed.

Name. Name of the service account.

Upstream description:

Name of the service account.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3113011011220100-3310223022310311-3233103122123321-1010020000203331-1022013232233100-0000323032200001-1221010003230333-3102320221102132"></a>

<a id="canonical-3131223111020211-2021212100110103-0020203000123122-3223033110001031-0302220331010033-2011102112303123-1303321112232103-3221322122130210"></a>

## namespace property — service_account / 112132211033 / 5

Type: `"string"`. Computed.

Namespace. Namespace of the service account.

Upstream description:

Namespace of the service account.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2222202312122202-2221201012013120-3210022120233321-2322231023012121-1100232102103113-1221230331110022-1022201101130031-0303132023310221"></a>

## Next pages — service_account / 112132211033 / 6

- [subjects](data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-1203213300130313-1121302211223313-1022100131210132-3302123311203132-0000133220133122-1230322013121132-0322311121121123-1200020310023310)
- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-2021222210012102-1330130200013102-0220323113122312-2121230113313023-1123103021232230-1310303332323332-2100321301113012-3031120031031013)
