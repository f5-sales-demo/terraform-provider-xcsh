---
page_title: "xcsh_k8s_pod_security_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy reference."
---

# xcsh_k8s_pod_security_policy reference

<a id="canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200200331321011-2011211103233103-1313002010103220-3102012310031021-0232023120023213-3021132233301021-0330012100120011-2102202030100231"></a>

## Property reference — Property reference / 022030232230 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- Property reference

<a id="canonical-3003303310201202-2011020033032000-1312232031001210-1130003203212012-2223010103121303-1311131201132021-1203013112331223-0313202110133331"></a>

## Direct properties — Property reference / 022030232230 / 3

<a id="canonical-0301222223332322-3012002331320002-1231000132130122-0011201331102113-1323232312033103-3323122122301211-1012332333320303-1220031012312331"></a>

<a id="canonical-3213120010303101-3110113131003310-3220010122032303-3330231030000320-1010103110231121-0002323013130021-0331300030112203-0002213021233333"></a>

## annotations property — Property reference / 022030232230 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-3121221210302013-1110331321213002-2110102023100212-3303303131000003-2101230310033111-0200112300230112-2002021200320023-1200023303321221"></a>

<a id="canonical-3203322211000303-0033121331102213-2013120210110300-1123011311032211-2111112211001100-3330131201021021-0011203223313131-3302330003211010"></a>

## description property — Property reference / 022030232230 / 5

Type: `"string"`. Optional.

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

<a id="canonical-1332302312201112-0202331101012033-1202013033123211-3201122220020203-0200032020001233-3021132323300113-1210330310332132-3013122130022210"></a>

<a id="canonical-1331113130102120-0332033312213100-2300131301230003-0201221130300202-0111033031303201-3002122300100010-3130032130313333-2203030233320312"></a>

## disable property — Property reference / 022030232230 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="canonical-3123213021330230-0020030132331023-1131311130232212-3231031320333111-3113230310200100-3002132301301113-2331200223013233-0222003112002333"></a>

<a id="canonical-1103310332320113-2231123121121020-2011203330223203-2132200313012321-1120302023011133-1033011031133320-2301302020110323-3333000101010223"></a>

## ID property — Property reference / 022030232230 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1011213012032322-0022201333032311-3132201011100023-3132222000333112-0232113203022111-1112302102333212-2101121002000232-0222100322000311"></a>

<a id="canonical-0122122213033301-2331110033330303-0322221001101200-3322222210021322-1331000211012132-3331231312200310-1300111112102320-3102110121111321"></a>

## labels property — Property reference / 022030232230 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-3103131232222311-3233212302003233-3111322220002310-3013000110113321-0311000203200320-0023013102212111-0202321210302133-1121320222033312"></a>

<a id="canonical-2120201213311301-1103303323200010-0210110331003231-2020332303330112-3321303311121111-0112011001323310-2012002012030230-3022301221210332"></a>

## name property — Property reference / 022030232230 / 9

Type: `"string"`. Required.

Name of the K8S Pod Security Policy. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

<a id="canonical-1302211133031211-2121200101130001-1313033311203112-1002123013323203-3112112120203301-0202331030031230-3120322022113312-3010233323001202"></a>

<a id="canonical-1132230031210031-2123120131121312-3302210222212231-2223202322000203-3122201320222203-1012313131012211-1323333130212110-1201310121111131"></a>

## namespace property — Property reference / 022030232230 / 10

Type: `"string"`. Required.

Namespace where the K8S Pod Security Policy is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032): complete subsection reference.

- [timeouts](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0031002323110013-1322030133322223-2221322200102102-0133311200200011-2121212023100030-2232100302333013-3331312322021120-0032320331101121): complete subsection reference.

<a id="canonical-1011320231111120-2201133022033311-0201320102311310-3010330200000010-2000101322010200-2002311101130130-1033123120232120-0121000220102020"></a>

<a id="canonical-2332311231333212-1012201302302213-2222011002233203-1322203112211322-0331222202022113-1303232130023302-2131101023300201-1100313202003310"></a>

## yaml property — Property reference / 022030232230 / 11

Type: `"string"`. Optional, Computed.

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

Upstream description:

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1332012202220310-2011132320221113-3011300032313213-3102010313003111-0103233233320102-1113212322233201-0102030313132012-0332021231220310"></a>

## All schema paths — Property reference / 022030232230 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0301222223332322-3012002331320002-1231000132130122-0011201331102113-1323232312033103-3323122122301211-1012332333320303-1220031012312331) |
| `description` | [description](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3121221210302013-1110331321213002-2110102023100212-3303303131000003-2101230310033111-0200112300230112-2002021200320023-1200023303321221) |
| `disable` | [disable](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1332302312201112-0202331101012033-1202013033123211-3201122220020203-0200032020001233-3021132323300113-1210330310332132-3013122130022210) |
| `id` | [ID](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3123213021330230-0020030132331023-1131311130232212-3231031320333111-3113230310200100-3002132301301113-2331200223013233-0222003112002333) |
| `labels` | [labels](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1011213012032322-0022201333032311-3132201011100023-3132222000333112-0232113203022111-1112302102333212-2101121002000232-0222100322000311) |
| `name` | [name](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3103131232222311-3233212302003233-3111322220002310-3013000110113321-0311000203200320-0023013102212111-0202321210302133-1121320222033312) |
| `namespace` | [namespace](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1302211133031211-2121200101130001-1313033311203112-1002123013323203-3112112120203301-0202331030031230-3120322022113312-3010233323001202) |
| `psp_spec` | [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2233323120020100-0132133001001111-3321212030232111-2232003300323302-2213301331121022-3120301230011223-0210022333011021-1313000102131213) |
| `psp_spec.allow_privilege_escalation` | [psp_spec.allow_privilege_escalation](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3031320230000302-1201323003130311-3211222002131020-0302003300012300-1302322112200212-0232003233100212-1330000003223110-0321312013000231) |
| `psp_spec.allowed_capabilities` | [psp_spec.allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2331333101333032-1123212232113122-1221202000102221-0021201122113022-0003030010201110-0232111101112133-0021300301211100-2120013100030231) |
| `psp_spec.allowed_capabilities.capabilities` | [psp_spec.allowed_capabilities.capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3212133031230033-2112230012232021-2013220202220010-1332130023102212-3321200021220111-0233231120002220-0132310223100233-0201002223233113) |
| `psp_spec.allowed_csi_drivers` | [psp_spec.allowed_csi_drivers](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3321130220302101-2111120103323331-1003011331200322-0013133302200301-3111311221013211-2213201303222122-1103321310213323-0332312203311133) |
| `psp_spec.allowed_flex_volumes` | [psp_spec.allowed_flex_volumes](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3321331320033023-3321012111212112-0100303201012223-3022010120200201-2010212032023113-3222002230003320-3122230223103002-2011032300130112) |
| `psp_spec.allowed_host_paths` | [psp_spec.allowed_host_paths](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3122030321110303-1320000032213302-1300220213111023-1101213222303012-1002023023021300-0203221302132120-3212330222300023-2211002010001310) |
| `psp_spec.allowed_host_paths.path_prefix` | [psp_spec.allowed_host_paths.path_prefix](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1310321023303000-0130212100030001-2131221133223213-3213130302301011-2203313033110233-3033321003121012-3330021131032321-1202231302213010) |
| `psp_spec.allowed_host_paths.read_only` | [psp_spec.allowed_host_paths.read_only](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1010132330200130-0201311220032202-2301200210213320-0101103320011230-0002321013013103-2033130120303220-2022331021033130-0010230101233223) |
| `psp_spec.allowed_proc_mounts` | [psp_spec.allowed_proc_mounts](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0233331230100211-3331022220111010-0032200110331023-0011323302131201-0012013200321133-1222113112203212-2333101031301032-3132001122113222) |
| `psp_spec.allowed_unsafe_sysctls` | [psp_spec.allowed_unsafe_sysctls](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0212023103200323-2331031223310113-1223332132333300-1201323131023330-3022223031232013-2110033030202133-2011121201212323-3220000231032030) |
| `psp_spec.default_allow_privilege_escalation` | [psp_spec.default_allow_privilege_escalation](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2220221333221312-0010023110232222-2330023030323031-0020100001330200-1033033121322031-0223000320110023-3321120032301302-0033221002303303) |
| `psp_spec.default_capabilities` | [psp_spec.default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2133120010120222-2133102230031312-1023212102200012-3333322213112221-0332120032021112-0013120223030230-1011230000132330-3320302003133020) |
| `psp_spec.default_capabilities.capabilities` | [psp_spec.default_capabilities.capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2320212321203020-0210322030012303-1120012300001030-2123212011020202-3100221102312310-1300030323300222-1100233030020311-1213311100100131) |
| `psp_spec.drop_capabilities` | [psp_spec.drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2033110210020121-1100012022222323-1133100013030013-0321013333112031-0302201233232101-0300130323003123-3213220030211121-1113323210000230) |
| `psp_spec.drop_capabilities.capabilities` | [psp_spec.drop_capabilities.capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2201332102333001-1331003332110112-2133313021120323-1021233210132311-1030101113303131-3012022112221223-0211120003121310-0020000302101312) |
| `psp_spec.forbidden_sysctls` | [psp_spec.forbidden_sysctls](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3203321123110000-3300013023320212-0101300303223322-1101123211132201-2003010320201122-1220011320210231-0102030220232020-3112323310003232) |
| `psp_spec.fs_group_strategy_options` | [psp_spec.fs_group_strategy_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1010113212302323-3221023332232222-2212203130001333-3231032012200301-3310121233211333-1322032011233131-2113010322313121-2212230133231212) |
| `psp_spec.fs_group_strategy_options.id_ranges` | [psp_spec.fs_group_strategy_options.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3223001232101023-0222120122102023-2122113202300303-3123332113113230-3313323120323223-1120312201330010-0003200103213110-3100013232022132) |
| `psp_spec.fs_group_strategy_options.id_ranges.max_id` | [psp_spec.fs_group_strategy_options.id_ranges.max_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1331020030122331-3121021332202303-1210331103001111-0000031101032222-3113233223011211-1311111322010003-0032120130223030-3330321022322210) |
| `psp_spec.fs_group_strategy_options.id_ranges.min_id` | [psp_spec.fs_group_strategy_options.id_ranges.min_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0332232301223132-1321333103311003-1011011031321002-1031133023120003-1010300200300121-1310300011132030-0031301213123130-0312330110113331) |
| `psp_spec.fs_group_strategy_options.rule` | [psp_spec.fs_group_strategy_options.rule](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2331033332012220-3301212310220213-3133212133120112-3210102000112211-1210223132133321-0020001020023010-1103212303001103-3301320320012333) |
| `psp_spec.host_ipc` | [psp_spec.host_ipc](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1333132021011301-0132313331310303-0300001002312323-0303131301023300-1330211121031030-2032312130203310-0011302121001331-1101132302200032) |
| `psp_spec.host_network` | [psp_spec.host_network](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1021331212332202-3201233112010330-2000121301301100-2131020132020332-1110202231022011-0210012332322233-3303131201031122-1330230010100033) |
| `psp_spec.host_pid` | [psp_spec.host_pid](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1030103120323001-3332112313230122-2121133133131322-3322120203123212-3003223221220132-1231123330223021-1211023310212313-3033220003210220) |
| `psp_spec.host_port_ranges` | [psp_spec.host_port_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3310202212211303-1001322011011132-0003031022231313-2330233303111333-2101111202301122-0013011203212011-2030120231110332-2001203230122222) |
| `psp_spec.no_allowed_capabilities` | [psp_spec.no_allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2223131311210133-2221101131210212-0331020310132312-0022113303211111-2010113233320122-3101003010220313-3021312203100002-2103102332032320) |
| `psp_spec.no_default_capabilities` | [psp_spec.no_default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0231021320213121-0302201000210130-2210202031023121-3112302323020320-2213000323212221-3102211002210120-3132322331212200-0320331103033233) |
| `psp_spec.no_drop_capabilities` | [psp_spec.no_drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1123130133101330-1032323212102021-2300211323013233-2133330103300311-1123303121330022-3201213200031220-2211232332122202-0132332121103001) |
| `psp_spec.no_fs_groups` | [psp_spec.no_fs_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2033033320023031-0332122012300103-2010120011023313-3012110331312101-1112333310010110-3131212311310200-3131021101312030-3310201221201032) |
| `psp_spec.no_run_as_group` | [psp_spec.no_run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2020230332103322-2102230203311110-2300010022031231-1101102232312322-2200103120003123-0131230330301103-0231203001010331-1222220120302312) |
| `psp_spec.no_run_as_user` | [psp_spec.no_run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1210021003320112-3221023132110030-0102121233210131-2130132300121223-2011103233031211-2301112000013233-1320103211220033-2011212330122032) |
| `psp_spec.no_runtime_class` | [psp_spec.no_runtime_class](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0223132013123120-2133000220022332-2223321022023301-1110110011300202-0231022122110212-0012003300303000-0032033312123310-0202100121330032) |
| `psp_spec.no_se_linux_options` | [psp_spec.no_se_linux_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2133203302020110-0000322120031120-0210101101101223-2223120121013313-2130131332032220-2121020231002003-1311211132320320-3230021310333033) |
| `psp_spec.no_supplemental_groups` | [psp_spec.no_supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1223022330011103-0302231320202003-1323101020312022-2230200030003003-0203201233030222-2132200320301110-0302003201330003-1323011320002321) |
| `psp_spec.privileged` | [psp_spec.privileged](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2020223121111021-0131202302202203-0311320121132210-2031211222122220-0130120323322032-1313303012300023-2112202133323222-1031131122131311) |
| `psp_spec.read_only_root_filesystem` | [psp_spec.read_only_root_filesystem](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2121332223123312-0111313122010323-0302101221023133-3031032031201001-2120211221000203-1010212002022021-2311103331302203-0010202322311012) |
| `psp_spec.run_as_group` | [psp_spec.run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0302301233020302-3110211020012321-3113022303002023-0031210102103303-2311112000221100-0333133002321220-3031032331100211-0230022211103331) |
| `psp_spec.run_as_group.id_ranges` | [psp_spec.run_as_group.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3133333023031201-3100211122320023-0300202312312220-3132032031210110-2300130301013223-3223301112223221-3130221022123033-2111113001121132) |
| `psp_spec.run_as_group.id_ranges.max_id` | [psp_spec.run_as_group.id_ranges.max_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1032331222103322-0121300131121100-2111011103320102-0023032011230213-3033023300201212-3120213300211021-3322331333121030-0230313313002332) |
| `psp_spec.run_as_group.id_ranges.min_id` | [psp_spec.run_as_group.id_ranges.min_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1001223203200002-2323302102130033-0333323331022231-1232013020331201-2231121133210222-3320322111233112-2111101002302223-2000132333200131) |
| `psp_spec.run_as_group.rule` | [psp_spec.run_as_group.rule](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3120233102100123-1333020022200210-1132133212001202-3333202101122023-1022110223130323-0313323231102230-2211210020031122-0102010231213021) |
| `psp_spec.run_as_user` | [psp_spec.run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1320301203002231-1212023221301203-2022300101100303-1203202120201002-2201201001330113-0130032000133330-1310112231010232-2223230113301021) |
| `psp_spec.run_as_user.id_ranges` | [psp_spec.run_as_user.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3112201212330321-2300203323311020-0131031311023201-0200333012001222-0133203212132111-3131311030323330-0203003023210222-2002103210302120) |
| `psp_spec.run_as_user.id_ranges.max_id` | [psp_spec.run_as_user.id_ranges.max_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1202213130133202-0203110231313121-3002121201333312-0323223310211203-1231030231001112-2231311000221311-1022220003321033-0211000323123011) |
| `psp_spec.run_as_user.id_ranges.min_id` | [psp_spec.run_as_user.id_ranges.min_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1312030233122202-3121131132211300-1012201021302310-2003011132033021-0130132322103123-0110313313310133-3200122331321030-1231033001011331) |
| `psp_spec.run_as_user.rule` | [psp_spec.run_as_user.rule](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2120031120022211-1223203023031123-1300003023203102-0212220323212001-0133233123111331-0312111221121212-0120203201320101-2303020221002332) |
| `psp_spec.supplemental_groups` | [psp_spec.supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1212300222201233-2031221103330210-3303121312232302-2332223110100230-3232121212320123-3300231233032113-2001203220202022-3011333102201003) |
| `psp_spec.supplemental_groups.id_ranges` | [psp_spec.supplemental_groups.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0320223023232200-1331123112323203-2102223021022023-1122222110221122-0223300211130013-1332112302132311-2132020102320000-3031030302301101) |
| `psp_spec.supplemental_groups.id_ranges.max_id` | [psp_spec.supplemental_groups.id_ranges.max_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3331223133302222-1213323321011012-0110322031320102-0310013331211230-2010131320121023-3113320202110033-3013022120212230-2331221302021321) |
| `psp_spec.supplemental_groups.id_ranges.min_id` | [psp_spec.supplemental_groups.id_ranges.min_id](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3121110213123323-0022112330031123-0311303320000113-2031220022020330-3223300311232233-2113031030301302-1231332020201211-1333200300223222) |
| `psp_spec.supplemental_groups.rule` | [psp_spec.supplemental_groups.rule](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1123320120313323-2320310122200020-1212201231023111-0002020101020312-2333323032333002-0113022002031200-1323213100301113-0320301200012031) |
| `psp_spec.volumes` | [psp_spec.volumes](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1120333133020032-1132121302320113-1302110313233312-1202102312002323-0102211020133301-2132221000030111-3003213010120131-0331031130030021) |
| `timeouts` | [timeouts](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0021230111132023-3022313010200233-1130230331003112-0301222132011302-1100302022200321-2333101322332131-3301002100130221-1230023211223222) |
| `timeouts.create` | [timeouts.create](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0233011001333330-2210333202220322-3332121330131302-3313303310320323-2022020132122011-1312332131333021-0233301023311101-1310022123332033) |
| `timeouts.delete` | [timeouts.delete](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1330212100221122-0030132221321230-2221102300303301-2312212213021232-0221010221213211-1031233311030003-2002231120222130-1302102122212123) |
| `timeouts.read` | [timeouts.read](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0122110121221331-2112002333200331-2010121100003231-1301113312120230-1233233003222310-0303111233001233-1031113010103211-2210320222313302) |
| `timeouts.update` | [timeouts.update](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1303111112330102-3121023110323210-3030100311030321-0230330013001200-0331200313312120-2110333121333030-0121300231111203-3232213322331103) |
| `yaml` | [yaml](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1011320231111120-2201133022033311-0201320102311310-3010330200000010-2000101322010200-2002311101130130-1033123120232120-0121000220102020) |

<a id="canonical-0110323201000310-3132212213302023-1320001113330210-3133233322311322-2313200230300120-0323020211213311-2302301330121112-2122220110231032"></a>

## Next pages — Property reference / 022030232230 / 13

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [timeouts](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0031002323110013-1322030133322223-2221322200102102-0133311200200011-2121212023100030-2232100302333013-3331312322021120-0032320331101121)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213011100301131-1213312111133100-3212232203000320-2010201012212023-2221012120132202-2112012202302132-1231030213201113-0320203220330323"></a>

## psp_spec — psp_spec / 123010312130 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- psp_spec

<a id="canonical-2233323120020100-0132133001001111-3321212030232111-2232003300323302-2213301331121022-3120301230011223-0210022333011021-1313000102131213"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: psp\_spec, yaml\] Pod Security Policy Specification. Form based pod security specification.

Upstream description:

Form based pod security specification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allowed_capabilities",
    "no_allowed_capabilities"),
  validators.ConflictingObjectAttributes("default_capabilities",
    "no_default_capabilities"),
  validators.ConflictingObjectAttributes("drop_capabilities",
    "no_drop_capabilities"),
  validators.ConflictingObjectAttributes("fs_group_strategy_options",
    "no_fs_groups"),
  validators.ConflictingObjectAttributes("no_run_as_group",
    "run_as_group"),
  validators.ConflictingObjectAttributes("no_run_as_user",
    "run_as_user"),
  validators.ConflictingObjectAttributes("no_supplemental_groups",
    "supplemental_groups")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_capabilities_choice": "[\"allowed_capabilities\",\"no_allowed_capabilities\"]",
  "x-ves-oneof-field-default_capabilities_choice": "[\"default_capabilities\",\"no_default_capabilities\"]",
  "x-ves-oneof-field-drop_capabilities_choice": "[\"drop_capabilities\",\"no_drop_capabilities\"]",
  "x-ves-oneof-field-fs_group_choice": "[\"fs_group_strategy_options\",\"no_fs_groups\"]",
  "x-ves-oneof-field-group_choice": "[\"no_run_as_group\",\"run_as_group\"]",
  "x-ves-oneof-field-runtime_class_choice": "[\"no_runtime_class\"]",
  "x-ves-oneof-field-se_linux_choice": "[\"no_se_linux_options\"]",
  "x-ves-oneof-field-supplemental_group_choice": "[\"no_supplemental_groups\",\"supplemental_groups\"]",
  "x-ves-oneof-field-user_choice": "[\"no_run_as_user\",\"run_as_user\"]"
}
```

OneOf alternatives in this subsection:

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2233323120020100-0132133001001111-3321212030232111-2232003300323302-2213301331121022-3120301230011223-0210022333011021-1313000102131213)
- [yaml](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1011320231111120-2201133022033311-0201320102311310-3010330200000010-2000101322010200-2002311101130130-1033123120232120-0121000220102020)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
psp_spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210111130221211-2332330201233213-1230031230330221-3221010003101311-2131130101322232-0000220223210233-2023013030010111-3133122320213130"></a>

## Direct properties — psp_spec / 123010312130 / 3

<a id="canonical-3031320230000302-1201323003130311-3211222002131020-0302003300012300-1302322112200212-0232003233100212-1330000003223110-0321312013000231"></a>

<a id="canonical-1112003233222213-3010113101110303-1123200233030222-2132222100120330-0312323200001211-0112001032210320-3131210221013020-1111022200021001"></a>

## allow_privilege_escalation property — psp_spec / 123010312130 / 4

Type: `"bool"`. Optional.

Pod can request to privilege escalation.

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

- [allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2010232122200033-3131102120101000-2230012232013333-0030312320331011-0222312113203301-2111131111321223-3330321021211233-2330333310103110): complete subsection reference.

<a id="canonical-3321130220302101-2111120103323331-1003011331200322-0013133302200301-3111311221013211-2213201303222122-1103321310213323-0332312203311133"></a>

<a id="canonical-3103233032231021-3222233003010220-0312022313121031-0202022333100120-1220021120013113-1030120232233013-1010311123223002-3032331200033103"></a>

## allowed_csi_drivers property — psp_spec / 123010312130 / 5

Type: `["list", "string"]`. Optional.

Restrict the available CSI drivers for POD, default all drivers are available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3321331320033023-3321012111212112-0100303201012223-3022010120200201-2010212032023113-3222002230003320-3122230223103002-2011032300130112"></a>

<a id="canonical-1100022200223010-1333213322211333-2211120213313013-1112303332321300-0303233211113100-3233231123221112-0222000100312002-0112213121130311"></a>

## allowed_flex_volumes property — psp_spec / 123010312130 / 6

Type: `["list", "string"]`. Optional.

Restrict list of Flex volumes, default all volumes are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [allowed_host_paths](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2333000103132100-2333300322022233-3102303231332032-1323133112130033-1003320033111122-1311000322012202-2312012021200312-3102122310133210): complete subsection reference.

<a id="canonical-0233331230100211-3331022220111010-0032200110331023-0011323302131201-0012013200321133-1222113112203212-2333101031301032-3132001122113222"></a>

<a id="canonical-1120030030112203-2011020322212131-0223211332132223-3211030111130101-3010110310000320-0120223112322021-1113030103002301-2302022030233011"></a>

## allowed_proc_mounts property — psp_spec / 123010312130 / 7

Type: `["list", "string"]`. Optional.

Allowed list of proc mounts, empty list allows default proc mounts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0212023103200323-2331031223310113-1223332132333300-1201323131023330-3022223031232013-2110033030202133-2011121201212323-3220000231032030"></a>

<a id="canonical-0200123100112003-2330323223132121-2112003322032003-1010201223130133-0110321101100003-1300002013230003-0332011011220113-1111233123201130"></a>

## allowed_unsafe_sysctls property — psp_spec / 123010312130 / 8

Type: `["list", "string"]`. Optional.

Allowed list of unsafe sysctls, empty list allows none. Supports prefix reg-ex.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2220221333221312-0010023110232222-2330023030323031-0020100001330200-1033033121322031-0223000320110023-3321120032301302-0033221002303303"></a>

<a id="canonical-1131020310111310-0031120100102203-3030200222011133-2320231230102211-0210130202232231-3332120322102003-1233102301132131-2100322010132120"></a>

## default_allow_privilege_escalation property — psp_spec / 123010312130 / 9

Type: `"bool"`. Optional.

Pod has permission for privilege escalation by default.

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

- [default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0320121200123032-3300223023003031-3312110132302322-2011111221211331-0121231010130122-2123330302022113-3223122013011331-0202032133232031): complete subsection reference.

- [drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0013012311000133-0230030322003010-2212223103133133-3020103022310121-3310221212330010-3303130332110211-3020133310030001-1131321110022132): complete subsection reference.

<a id="canonical-3203321123110000-3300013023320212-0101300303223322-1101123211132201-2003010320201122-1220011320210231-0102030220232020-3112323310003232"></a>

<a id="canonical-0201031001301121-0320320331332221-3103120320122002-2030220010131022-3220120110002303-0030322131230130-3103311000103233-1230320112123123"></a>

## forbidden_sysctls property — psp_spec / 123010312130 / 10

Type: `["list", "string"]`. Optional.

Forbidden list of sysctls, empty list forbids none. Supports prefix reg-ex.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [fs_group_strategy_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2202103223103213-0332112021012020-3021123103001202-2322203103031130-2212121011111303-3011023200220000-3201232010202000-1000101323210301): complete subsection reference.

<a id="canonical-1333132021011301-0132313331310303-0300001002312323-0303131301023300-1330211121031030-2032312130203310-0011302121001331-1101132302200032"></a>

<a id="canonical-2022313120202313-3001312102232113-2133330123222210-3200013303201032-0313301010101201-0023303000003203-2010213000221132-0230332233013213"></a>

## host_ipc property — psp_spec / 123010312130 / 11

Type: `"bool"`. Optional.

Host IPC determines if the policy allows the use of host IPC in the pod spec.

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

<a id="canonical-1021331212332202-3201233112010330-2000121301301100-2131020132020332-1110202231022011-0210012332322233-3303131201031122-1330230010100033"></a>

<a id="canonical-2322130213121302-0132011020331320-0011232222103100-3102323232221311-2320303202111121-0212013310000200-2102102202321002-2212313001230310"></a>

## host_network property — psp_spec / 123010312130 / 12

Type: `"bool"`. Optional.

Host Network determines if the policy allows the use of host network in the pod spec.

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

<a id="canonical-1030103120323001-3332112313230122-2121133133131322-3322120203123212-3003223221220132-1231123330223021-1211023310212313-3033220003210220"></a>

<a id="canonical-1312020032222002-2331001200320130-1031233101030313-1000221130113200-2303320210300200-2323231111131200-0313323231011203-3002110331332120"></a>

## host_pid property — psp_spec / 123010312130 / 13

Type: `"bool"`. Optional.

Host PID determines if the policy allows the use of host PID in the pod spec.

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

<a id="canonical-3310202212211303-1001322011011132-0003031022231313-2330233303111333-2101111202301122-0013011203212011-2030120231110332-2001203230122222"></a>

<a id="canonical-0120223012332302-1333320231011030-2131023303110211-0132022333002210-2122321210103130-0230300021012223-3010112321330212-0010303021322222"></a>

## host_port_ranges property — psp_spec / 123010312130 / 14

Type: `"string"`. Optional.

Host port ranges determines which ports ranges are allowed to be exposed.

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
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

- [no_allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0302200003012332-0333223331131203-1001032011223130-1201302313020200-1330012011221003-2321301131211303-1303321133213011-0331302101321220): complete subsection reference.

- [no_default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2013301201110122-0122201311121120-2202301211333130-3310213212123112-2033030032033330-2021122211203213-0131202331011102-1103120220002100): complete subsection reference.

- [no_drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3332301003130032-0033030313113032-2313133021322300-0330030130300121-3300000323022212-1020000121022203-2100132122013221-1333223031030220): complete subsection reference.

- [no_fs_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3303333110121231-1032022001222131-2013100100313320-1232011121001120-1201233321002320-2110021110202323-3001322323213202-2223213100220223): complete subsection reference.

- [no_run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0133101123231121-1022302101213010-2131233100121310-0020033112321320-1233023300313001-2210110301020100-3013233023123310-3112333010223210): complete subsection reference.

- [no_run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3031131231301130-3310003011010303-2302320021032020-1232213122002300-0321120231333232-2102013303133132-0031022203320001-3033002201331313): complete subsection reference.

- [no_runtime_class](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1301222013302213-3220110221303233-1220001130013130-0201131131110303-1301320001031300-3232103313220031-0311203100101133-3123010122213033): complete subsection reference.

- [no_se_linux_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3130313121011310-0123010210201133-3131302201133111-3133223221313323-1300212110323010-2113233103232320-1300011103111310-1100130020123210): complete subsection reference.

- [no_supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2221200010101222-2323303120020313-2013122021202203-2012001211213022-0033133221122031-1022002312331021-1113121110231313-0132033330031120): complete subsection reference.

<a id="canonical-2020223121111021-0131202302202203-0311320121132210-2031211222122220-0130120323322032-1313303012300023-2112202133323222-1031131122131311"></a>

<a id="canonical-0003212032010033-3132222031030012-1130123323320231-3233213210110221-0102011331121102-0130330110010200-0203232303222211-1030322122103003"></a>

## privileged property — psp_spec / 123010312130 / 15

Type: `"bool"`. Optional.

Privileged determines if a pod can request to be run as privileged.

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

<a id="canonical-2121332223123312-0111313122010323-0302101221023133-3031032031201001-2120211221000203-1010212002022021-2311103331302203-0010202322311012"></a>

<a id="canonical-0223233133223121-3030301131013122-3113223223132333-1122121223000321-2321123022202013-0200200313330303-3230211111011211-3321232012323122"></a>

## read_only_root_filesystem property — psp_spec / 123010312130 / 16

Type: `"bool"`. Optional.

Containers can only run with read only root filesystem.

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

- [run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0303231030223211-2203013023100112-0312011300120231-3121030333202320-1003002031332121-3211020330320023-0121002131022312-0133133232122332): complete subsection reference.

- [run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3322031322303132-2233300101211301-1033020232212012-3020303310332020-3232213002213301-1001212310130123-1203202133013320-0022212231230221): complete subsection reference.

- [supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3212320333313232-2233312031023110-2223231010003002-2023103303022112-2302230232033013-1011011211221023-1010000011313231-2102321103331011): complete subsection reference.

<a id="canonical-1120333133020032-1132121302320113-1302110313233312-1202102312002323-0102211020133301-2132221000030111-3003213010120131-0331031130030021"></a>

<a id="canonical-3023112201220031-3312202330202200-0031003302331131-1102301013000113-3023223312103032-0302312300111021-1123233112330223-1011300111132103"></a>

## volumes property — psp_spec / 123010312130 / 17

Type: `["list", "string"]`. Optional.

Allow List of volume plugins. Empty no volumes are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3212331010001023-0301303323201211-0311331202222310-2130313210232330-0312101121232201-1322131100122122-1211133013232322-3330212123233312"></a>

## Next pages — psp_spec / 123010312130 / 18

- [psp_spec.allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2010232122200033-3131102120101000-2230012232013333-0030312320331011-0222312113203301-2111131111321223-3330321021211233-2330333310103110)
- [psp_spec.allowed_host_paths](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2333000103132100-2333300322022233-3102303231332032-1323133112130033-1003320033111122-1311000322012202-2312012021200312-3102122310133210)
- [psp_spec.default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0320121200123032-3300223023003031-3312110132302322-2011111221211331-0121231010130122-2123330302022113-3223122013011331-0202032133232031)
- [psp_spec.drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0013012311000133-0230030322003010-2212223103133133-3020103022310121-3310221212330010-3303130332110211-3020133310030001-1131321110022132)
- [psp_spec.fs_group_strategy_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2202103223103213-0332112021012020-3021123103001202-2322203103031130-2212121011111303-3011023200220000-3201232010202000-1000101323210301)
- [psp_spec.no_allowed_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0302200003012332-0333223331131203-1001032011223130-1201302313020200-1330012011221003-2321301131211303-1303321133213011-0331302101321220)
- [psp_spec.no_default_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2013301201110122-0122201311121120-2202301211333130-3310213212123112-2033030032033330-2021122211203213-0131202331011102-1103120220002100)
- [psp_spec.no_drop_capabilities](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3332301003130032-0033030313113032-2313133021322300-0330030130300121-3300000323022212-1020000121022203-2100132122013221-1333223031030220)
- [psp_spec.no_fs_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3303333110121231-1032022001222131-2013100100313320-1232011121001120-1201233321002320-2110021110202323-3001322323213202-2223213100220223)
- [psp_spec.no_run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0133101123231121-1022302101213010-2131233100121310-0020033112321320-1233023300313001-2210110301020100-3013233023123310-3112333010223210)
- [psp_spec.no_run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3031131231301130-3310003011010303-2302320021032020-1232213122002300-0321120231333232-2102013303133132-0031022203320001-3033002201331313)
- [psp_spec.no_runtime_class](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1301222013302213-3220110221303233-1220001130013130-0201131131110303-1301320001031300-3232103313220031-0311203100101133-3123010122213033)
- [psp_spec.no_se_linux_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3130313121011310-0123010210201133-3131302201133111-3133223221313323-1300212110323010-2113233103232320-1300011103111310-1100130020123210)
- [psp_spec.no_supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2221200010101222-2323303120020313-2013122021202203-2012001211213022-0033133221122031-1022002312331021-1113121110231313-0132033330031120)
- [psp_spec.run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0303231030223211-2203013023100112-0312011300120231-3121030333202320-1003002031332121-3211020330320023-0121002131022312-0133133232122332)
- [psp_spec.run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3322031322303132-2233300101211301-1033020232212012-3020303310332020-3232213002213301-1001212310130123-1203202133013320-0022212231230221)
- [psp_spec.supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3212320333313232-2233312031023110-2223231010003002-2023103303022112-2302230232033013-1011011211221023-1010000011313231-2102321103331011)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-2010232122200033-3131102120101000-2230012232013333-0030312320331011-0222312113203301-2111131111321223-3330321021211233-2330333310103110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000101221211322-2301111201210011-0233201112221213-2130001122101221-2301012331233012-3312303132021200-0102122113331022-3330323222103102"></a>

## psp_spec.allowed_capabilities — allowed_capabilities / 013121210013 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.allowed_capabilities

<a id="canonical-2331333101333032-1123212232113122-1221202000102221-0021201122113022-0003030010201110-0232111101112133-0021300301211100-2120013100030231"></a>

Type: `"object"`. single nested block, Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("capabilities")}
```

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

Terraform syntax:

```terraform
allowed_capabilities {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122113110112230-3111223132303113-0030103103220300-0331313020200030-0012022301303213-2011111122302330-2322133232221303-1003302022133032"></a>

## Direct properties — allowed_capabilities / 013121210013 / 3

<a id="canonical-3212133031230033-2112230012232021-2013220202220010-1332130023102212-3321200021220111-0233231120002220-0132310223100233-0201002223233113"></a>

<a id="canonical-1033111312030322-0332323320221103-0210201310322220-2032100201101212-3311130300113032-1302220221002031-0312321020000323-1211322103302103"></a>

## capabilities property — allowed_capabilities / 013121210013 / 4

Type: `["list", "string"]`. Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3310300221100110-1012201322122120-1200003300312112-3231100033331033-0330003100021231-0200323032131122-3011323211211232-0310222103222000"></a>

## Next pages — allowed_capabilities / 013121210013 / 5

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-2333000103132100-2333300322022233-3102303231332032-1323133112130033-1003320033111122-1311000322012202-2312012021200312-3102122310133210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123303102101323-2003032213331320-3011332332211012-2312212310313032-0312233311031233-3232221310230100-0311120312131130-3100011010321031"></a>

## psp_spec.allowed_host_paths — allowed_host_paths / 032330031313 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.allowed_host_paths

<a id="canonical-3122030321110303-1320000032213302-1300220213111023-1101213222303012-1002023023021300-0203221302132120-3212330222300023-2211002010001310"></a>

Type: `"object"`. list nested block, Optional.

Restrict list of host paths, default all host paths are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("path_prefix")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
allowed_host_paths {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103311013330223-3202020202120021-2210100001012120-0031112212213033-2222022123130331-1120231102311130-2012211133221203-1032132002031230"></a>

## Direct properties — allowed_host_paths / 032330031313 / 3

<a id="canonical-1310321023303000-0130212100030001-2131221133223213-3213130302301011-2203313033110233-3033321003121012-3330021131032321-1202231302213010"></a>

<a id="canonical-0210002212122321-3001010101131123-1121000110220032-1221001212011302-1021332233311320-1332020122022211-1000122323231211-2212220100231310"></a>

## path_prefix property — allowed_host_paths / 032330031313 / 4

Type: `"string"`. Optional.

Host path prefix is the path prefix that the host volume must match. It does not support \*.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1010132330200130-0201311220032202-2301200210213320-0101103320011230-0002321013013103-2033130120303220-2022331021033130-0010230101233223"></a>

<a id="canonical-3133311213220113-3102202200132222-3201122322222001-0221232112100203-0033302223132313-1303321000220211-1302202313002200-1111003331113123"></a>

## read_only property — allowed_host_paths / 032330031313 / 5

Type: `"bool"`. Optional.

Volume will be allowed to mount read only.

Upstream description:

This volume will be allowed to mount read only.

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

<a id="canonical-1331223213120313-2311333132112000-2230012113032330-2320103203302330-1322132013112310-2013203000012000-0123311301000310-1103011321222221"></a>

## Next pages — allowed_host_paths / 032330031313 / 6

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-0320121200123032-3300223023003031-3312110132302322-2011111221211331-0121231010130122-2123330302022113-3223122013011331-0202032133232031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202033302132011-3332102002012203-3333002131221131-3210001110113312-0032022202133022-3221102223211331-3203232030330130-2220313221332101"></a>

## psp_spec.default_capabilities — default_capabilities / 320113023000 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.default_capabilities

<a id="canonical-2133120010120222-2133102230031312-1023212102200012-3333322213112221-0332120032021112-0013120223030230-1011230000132330-3320302003133020"></a>

Type: `"object"`. single nested block, Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("capabilities")}
```

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

Terraform syntax:

```terraform
default_capabilities {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212023001200330-2133120132101210-3101331321200030-3331132021210302-0203320330330000-2303100333123212-3211203123310022-3312030020032310"></a>

## Direct properties — default_capabilities / 320113023000 / 3

<a id="canonical-2320212321203020-0210322030012303-1120012300001030-2123212011020202-3100221102312310-1300030323300222-1100233030020311-1213311100100131"></a>

<a id="canonical-0302001313002211-1310203300122102-2313332000220011-0213032333322013-3213101213112213-0132233231322001-1110320111100301-2013130101113121"></a>

## capabilities property — default_capabilities / 320113023000 / 4

Type: `["list", "string"]`. Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1231101113003031-1232320312230201-3232031100221310-0222332310101123-2211011100331333-3102222132012320-2103120322113233-2021122232322012"></a>

## Next pages — default_capabilities / 320113023000 / 5

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-0013012311000133-0230030322003010-2212223103133133-3020103022310121-3310221212330010-3303130332110211-3020133310030001-1131321110022132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130211232123312-1303012122010013-0300101022033032-0113030131132122-3100333111310222-2232333301312022-2211222321023323-3332030301322011"></a>

## psp_spec.drop_capabilities — drop_capabilities / 231231221100 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.drop_capabilities

<a id="canonical-2033110210020121-1100012022222323-1133100013030013-0321013333112031-0302201233232101-0300130323003123-3213220030211121-1113323210000230"></a>

Type: `"object"`. single nested block, Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("capabilities")}
```

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

Terraform syntax:

```terraform
drop_capabilities {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121003322211130-2211001032012021-0132130221313210-2110301020310211-3031222332221030-3201021323130130-3212231011101131-0321213220130001"></a>

## Direct properties — drop_capabilities / 231231221100 / 3

<a id="canonical-2201332102333001-1331003332110112-2133313021120323-1021233210132311-1030101113303131-3012022112221223-0211120003121310-0020000302101312"></a>

<a id="canonical-1210101102302300-1231003212123201-2210123310102131-2021210021112102-2333101102030301-1313230213100013-0210310313312332-3200011023102200"></a>

## capabilities property — drop_capabilities / 231231221100 / 4

Type: `["list", "string"]`. Optional.

List of capabilities that docker container has.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1032313031222213-0333332110001022-2311323303132130-1132111302102200-2010220132120233-3123031211111033-3100222201022313-2132133000101122"></a>

## Next pages — drop_capabilities / 231231221100 / 5

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-2202103223103213-0332112021012020-3021123103001202-2322203103031130-2212121011111303-3011023200220000-3201232010202000-1000101323210301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103133303211303-2120013303321001-2332033200332320-1011320120000320-0203021311311303-3211301222003212-1203202131312003-3210113210223331"></a>

## psp_spec.fs_group_strategy_options — fs_group_strategy_options / 020233310310 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.fs_group_strategy_options

<a id="canonical-1010113212302323-3221023332232222-2212203130001333-3231032012200301-3310121233211333-1322032011233131-2113010322313121-2212230133231212"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fs group strategy options.

Upstream description:

ID ranges and rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule")}
```

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

Terraform syntax:

```terraform
fs_group_strategy_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011110013132010-1132201212231121-2220002331131122-0331311223121233-3301101300123311-3221320221113203-0320130102233231-1111013221332320"></a>

## Direct properties — fs_group_strategy_options / 020233310310 / 3

- [id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1213222223311331-2131213310331112-1121113031221200-0333023102323233-3333312231232231-2210203032110022-3032230000333300-1230101333213020): complete subsection reference.

<a id="canonical-2331033332012220-3301212310220213-3133212133120112-3210102000112211-1210223132133321-0020001020023010-1103212303001103-3301320320012333"></a>

<a id="canonical-3323213312311330-3311023331320020-0022011303031310-1023212321000333-1220303211033201-1032131212322013-0122023323212010-0220223213111012"></a>

## rule property — fs_group_strategy_options / 020233310310 / 4

Type: `"string"`. Optional.

Rule indicated how the FS group ID range is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0200112021333002-0030112320021301-1211130113131213-3103130111032013-2300322110212211-1232201110300133-2022201223121000-1231013300300323"></a>

## Next pages — fs_group_strategy_options / 020233310310 / 5

- [psp_spec.fs_group_strategy_options.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1213222223311331-2131213310331112-1121113031221200-0333023102323233-3333312231232231-2210203032110022-3032230000333300-1230101333213020)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-1213222223311331-2131213310331112-1121113031221200-0333023102323233-3333312231232231-2210203032110022-3032230000333300-1230101333213020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101322302131203-2220202303013322-0032103222322030-0011133301010323-2301020232331330-1320033003212120-0213020111101302-3221020301131113"></a>

## psp_spec.fs_group_strategy_options.id_ranges — id_ranges / 200122122230 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [psp_spec.fs_group_strategy_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2202103223103213-0332112021012020-3021123103001202-2322203103031130-2212121011111303-3011023200220000-3201232010202000-1000101323210301)
- psp_spec.fs_group_strategy_options.id_ranges

<a id="canonical-3223001232101023-0222120122102023-2122113202300303-3123332113113230-3313323120323223-1120312201330010-0003200103213110-3100013232022132"></a>

Type: `"object"`. list nested block, Optional.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("max_id",
    "min_id")}
```

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
id_ranges {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032302020120300-2322000012221322-0110110300222310-3110021101002330-2010113330003020-2310300300320011-0213310000120011-1120133122030023"></a>

## Direct properties — id_ranges / 200122122230 / 3

<a id="canonical-1331020030122331-3121021332202303-1210331103001111-0000031101032222-3113233223011211-1311111322010003-0032120130223030-3330321022322210"></a>

<a id="canonical-0310202023131302-3331032302230311-2132131311113303-2303111123303200-1002202303321320-3332111302022103-2303030311232131-2010021033330331"></a>

## max_id property — id_ranges / 200122122230 / 4

Type: `"number"`. Optional.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0332232301223132-1321333103311003-1011011031321002-1031133023120003-1010300200300121-1310300011132030-0031301213123130-0312330110113331"></a>

<a id="canonical-1021211303123322-1013220331102223-2021231113201222-0232232003310002-1012123332011122-3320300232102012-3303333213133111-3101111110200212"></a>

## min_id property — id_ranges / 200122122230 / 5

Type: `"number"`. Optional.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1303122322132121-2320230030221103-2221202022300233-1301122301210120-3010210110021001-1201200120201233-0002212000032313-3023120132301031"></a>

## Next pages — id_ranges / 200122122230 / 6

- [psp_spec.fs_group_strategy_options](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2202103223103213-0332112021012020-3021123103001202-2322203103031130-2212121011111303-3011023200220000-3201232010202000-1000101323210301)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-0302200003012332-0333223331131203-1001032011223130-1201302313020200-1330012011221003-2321301131211303-1303321133213011-0331302101321220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030122323233213-3311013333330332-3222321231311122-2312210213002231-3110331312120033-2220112201022300-2101320203332002-3331010032311230"></a>

## psp_spec.no_allowed_capabilities — no_allowed_capabilities / 320120230211 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.no_allowed_capabilities

<a id="canonical-2223131311210133-2221101131210212-0331020310132312-0022113303211111-2010113233320122-3101003010220313-3021312203100002-2103102332032320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no allowed capabilities.

Upstream description:

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

Terraform syntax:

```terraform
no_allowed_capabilities = {}
```

<a id="canonical-1031033013003113-2122322033200330-1000031002202101-0333220022023010-1031213020111023-0010223210220123-0121232121012303-1103033002011322"></a>

## Direct properties — no_allowed_capabilities / 320120230211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010222002320211-2132033122131312-2302220032002031-1032103102213301-2320323221301333-3102330221321113-1221110202221013-2011220310303203"></a>

## Next pages — no_allowed_capabilities / 320120230211 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-2013301201110122-0122201311121120-2202301211333130-3310213212123112-2033030032033330-2021122211203213-0131202331011102-1103120220002100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000201201131103-1320113201223132-0033032210213012-3332222322011023-0202301220010002-0213330101232121-2123021122333331-3230031312303221"></a>

## psp_spec.no_default_capabilities — no_default_capabilities / 323302022231 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.no_default_capabilities

<a id="canonical-0231021320213121-0302201000210130-2210202031023121-3112302323020320-2213000323212221-3102211002210120-3132322331212200-0320331103033233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no default capabilities.

Upstream description:

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

Terraform syntax:

```terraform
no_default_capabilities = {}
```

<a id="canonical-0221103303333313-2202030101331212-0211203011133232-1110323231221032-1310013100313122-2221101031212322-0002022102112113-2203301302330112"></a>

## Direct properties — no_default_capabilities / 323302022231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113112021003122-2000231300330211-0002312012221303-1220112201220131-0122311001023210-2312301120232013-3301130112132331-1022031020330330"></a>

## Next pages — no_default_capabilities / 323302022231 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-3332301003130032-0033030313113032-2313133021322300-0330030130300121-3300000323022212-1020000121022203-2100132122013221-1333223031030220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233032300201033-2202233321301322-2313320000312232-0301213211300120-2323310020320003-1113000130030201-2321032022100203-1231321202231200"></a>

## psp_spec.no_drop_capabilities — no_drop_capabilities / 332222121113 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.no_drop_capabilities

<a id="canonical-1123130133101330-1032323212102021-2300211323013233-2133330103300311-1123303121330022-3201213200031220-2211232332122202-0132332121103001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no drop capabilities.

Upstream description:

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

Terraform syntax:

```terraform
no_drop_capabilities = {}
```

<a id="canonical-3313330121313010-1201333131101321-3023232130210232-2313211002102213-0111220111223110-2211320032301212-1312312322000311-2313022221300213"></a>

## Direct properties — no_drop_capabilities / 332222121113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231212201111211-3210133313022333-2200210213331133-3122232311013231-1121230330301012-1120223032130032-2000011231131103-1222110303113131"></a>

## Next pages — no_drop_capabilities / 332222121113 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-3303333110121231-1032022001222131-2013100100313320-1232011121001120-1201233321002320-2110021110202323-3001322323213202-2223213100220223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110323332210000-0233003223213311-0033212230100311-2301100331203023-3213311223332310-3100210021203202-1111131130230031-2332111010311303"></a>

## psp_spec.no_fs_groups — no_fs_groups / 232203031212 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.no_fs_groups

<a id="canonical-2033033320023031-0332122012300103-2010120011023313-3012110331312101-1112333310010110-3131212311310200-3131021101312030-3310201221201032"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

Terraform syntax:

```terraform
no_fs_groups = {}
```

<a id="canonical-1203100200302121-3231311330301001-2310233001031222-1000123022201200-2101332102220002-1222223223310332-1021300202212111-3130312231032213"></a>

## Direct properties — no_fs_groups / 232203031212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322320021212321-2202302320303020-3001201003102223-2012313110121313-3013033112030222-1022211020330220-0030221312313112-2113301200300101"></a>

## Next pages — no_fs_groups / 232203031212 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-0133101123231121-1022302101213010-2131233100121310-0020033112321320-1233023300313001-2210110301020100-3013233023123310-3112333010223210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122010232210031-0313302330112211-3113302322231201-0303111231330213-2202302122021332-0132302322202121-0310232333000031-3012120223211232"></a>

## psp_spec.no_run_as_group — no_run_as_group / 311200212322 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.no_run_as_group

<a id="canonical-2020230332103322-2102230203311110-2300010022031231-1101102232312322-2200103120003123-0131230330301103-0231203001010331-1222220120302312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no run as group.

Upstream description:

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

Terraform syntax:

```terraform
no_run_as_group = {}
```

<a id="canonical-2132120220211213-1122112133021030-0311203302020200-1320302020002312-2123103220212031-0230210222103000-2303221000301332-0102100032332123"></a>

## Direct properties — no_run_as_group / 311200212322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200113022212021-1311211031223023-3300313032012020-1313001033230020-1010330012023101-2302010312331210-0003030031221332-0032222020102213"></a>

## Next pages — no_run_as_group / 311200212322 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-3031131231301130-3310003011010303-2302320021032020-1232213122002300-0321120231333232-2102013303133132-0031022203320001-3033002201331313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300103033223102-2232200110003030-0301133001123211-0221003120333100-1332333130030300-1303333302110002-3202302100032331-0013311122002313"></a>

## psp_spec.no_run_as_user — no_run_as_user / 031301031020 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.no_run_as_user

<a id="canonical-1210021003320112-3221023132110030-0102121233210131-2130132300121223-2011103233031211-2301112000013233-1320103211220033-2011212330122032"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no run as user.

Upstream description:

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

Terraform syntax:

```terraform
no_run_as_user = {}
```

<a id="canonical-1322210030112321-2013123221212013-3003212311130121-1201303231213203-0000321102113013-1303201122203023-3020222203112333-1220002003310033"></a>

## Direct properties — no_run_as_user / 031301031020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102312120323220-0011212221101110-2221300021133011-1313013222030102-3020302033222111-1020110123233103-2100333221101110-1202021030331303"></a>

## Next pages — no_run_as_user / 031301031020 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-1301222013302213-3220110221303233-1220001130013130-0201131131110303-1301320001031300-3232103313220031-0311203100101133-3123010122213033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021010303033202-0032223113330230-2120111031021312-1212110300031021-3001231033130231-0020223112000030-0300301212132102-0210032230223130"></a>

## psp_spec.no_runtime_class — no_runtime_class / 001133033302 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.no_runtime_class

<a id="canonical-0223132013123120-2133000220022332-2223321022023301-1110110011300202-0231022122110212-0012003300303000-0032033312123310-0202100121330032"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for no runtime class.

Upstream description:

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

Terraform syntax:

```terraform
no_runtime_class {}
```

<a id="canonical-0301212213113211-2103122032123030-1301213010100211-0003013300323121-3322001001130232-3220122031020120-0133111203132003-2332203103023031"></a>

## Direct properties — no_runtime_class / 001133033302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000123202012103-3000312233232223-0121122133033032-2002311010130222-0222132010212020-3201010012333100-2212000122121303-1303211133001230"></a>

## Next pages — no_runtime_class / 001133033302 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-3130313121011310-0123010210201133-3131302201133111-3133223221313323-1300212110323010-2113233103232320-1300011103111310-1100130020123210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311223123131321-2123220103030230-2330223100321133-3221230333013230-2122211331033300-2310331131223222-3312020022113112-2001330031032122"></a>

## psp_spec.no_se_linux_options — no_se_linux_options / 003302101232 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.no_se_linux_options

<a id="canonical-2133203302020110-0000322120031120-0210101101101223-2223120121013313-2130131332032220-2121020231002003-1311211132320320-3230021310333033"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for no se linux options.

Upstream description:

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

Terraform syntax:

```terraform
no_se_linux_options {}
```

<a id="canonical-2211122222101100-3211101311012132-0310112102202023-1310303232022131-1223022230233102-1203202321023023-2230211102230310-3200220000030133"></a>

## Direct properties — no_se_linux_options / 003302101232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221021310230220-2323320332201120-1011033322113330-1021300210223130-2111003012001323-0310002110130120-1001031020322300-2131201122112111"></a>

## Next pages — no_se_linux_options / 003302101232 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-2221200010101222-2323303120020313-2013122021202203-2012001211213022-0033133221122031-1022002312331021-1113121110231313-0132033330031120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111133212030300-0020222003003110-0310222000301302-2123300201022033-1220212023123322-1310301132010113-3131211303333101-3302102322011211"></a>

## psp_spec.no_supplemental_groups — no_supplemental_groups / 212012013120 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.no_supplemental_groups

<a id="canonical-1223022330011103-0302231320202003-1323101020312022-2230200030003003-0203201233030222-2132200320301110-0302003201330003-1323011320002321"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

Terraform syntax:

```terraform
no_supplemental_groups = {}
```

<a id="canonical-1331322132121300-0103111103020130-0110033213001230-2033201202331332-2211002013122210-0002131102010232-2030311321120321-1000122102321122"></a>

## Direct properties — no_supplemental_groups / 212012013120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311321201313202-3331111300232131-1103322212202303-0121011230213331-2110302210202330-2031322010112131-1311203132233302-2013020302330113"></a>

## Next pages — no_supplemental_groups / 212012013120 / 4

- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-0303231030223211-2203013023100112-0312011300120231-3121030333202320-1003002031332121-3211020330320023-0121002131022312-0133133232122332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000002203003031-2231211303122110-0122232122031320-1223013012233303-1213331221130120-1000333121333303-1211231223313100-1110201230033033"></a>

## psp_spec.run_as_group — run_as_group / 203313031221 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.run_as_group

<a id="canonical-0302301233020302-3110211020012321-3113022303002023-0031210102103303-2311112000221100-0333133002321220-3031032331100211-0230022211103331"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for run as group.

Upstream description:

ID ranges and rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule")}
```

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

Terraform syntax:

```terraform
run_as_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303020233002011-3002202131111302-1201230333110211-2200203311333203-3032223222110322-3000313230303320-3021213320232232-3231010230110023"></a>

## Direct properties — run_as_group / 203313031221 / 3

- [id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2333233330121101-3002122021112322-1003022201300012-2032110011113132-2232203132331300-3111321113333321-3113333211032310-3302200203203130): complete subsection reference.

<a id="canonical-3120233102100123-1333020022200210-1132133212001202-3333202101122023-1022110223130323-0313323231102230-2211210020031122-0102010231213021"></a>

<a id="canonical-3332000211223101-1102211100022032-0301023301132133-0323201302131000-3013212121210001-0321033130203201-0202101103201013-0320112323121331"></a>

## rule property — run_as_group / 203313031221 / 4

Type: `"string"`. Optional.

Rule indicated how the FS group ID range is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1311113323003012-1212210301302223-1011210000311232-3120100212321031-0300021131301023-3201210233133331-1032213012103210-3232031132212311"></a>

## Next pages — run_as_group / 203313031221 / 5

- [psp_spec.run_as_group.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-2333233330121101-3002122021112322-1003022201300012-2032110011113132-2232203132331300-3111321113333321-3113333211032310-3302200203203130)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-2333233330121101-3002122021112322-1003022201300012-2032110011113132-2232203132331300-3111321113333321-3113333211032310-3302200203203130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221201001211112-3223320003213232-2231231221203222-2311123333311212-2112320333220000-2330002201013332-1201230323022312-0200000232211000"></a>

## psp_spec.run_as_group.id_ranges — id_ranges / 021322211233 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [psp_spec.run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0303231030223211-2203013023100112-0312011300120231-3121030333202320-1003002031332121-3211020330320023-0121002131022312-0133133232122332)
- psp_spec.run_as_group.id_ranges

<a id="canonical-3133333023031201-3100211122320023-0300202312312220-3132032031210110-2300130301013223-3223301112223221-3130221022123033-2111113001121132"></a>

Type: `"object"`. list nested block, Optional.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("max_id",
    "min_id")}
```

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
id_ranges {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303033300310121-0311333020220123-1012112322222323-3300003311302112-1321320313321000-2131202322010331-3331300223310313-1300203201011032"></a>

## Direct properties — id_ranges / 021322211233 / 3

<a id="canonical-1032331222103322-0121300131121100-2111011103320102-0023032011230213-3033023300201212-3120213300211021-3322331333121030-0230313313002332"></a>

<a id="canonical-0130233220230033-1333103211032210-2313203003333211-2012102123023120-1001131322133113-0120023200031330-3312102131331123-1200023321330302"></a>

## max_id property — id_ranges / 021322211233 / 4

Type: `"number"`. Optional.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1001223203200002-2323302102130033-0333323331022231-1232013020331201-2231121133210222-3320322111233112-2111101002302223-2000132333200131"></a>

<a id="canonical-2000303210233013-0313230331112201-2313201301210103-1032130123212201-1102330111223022-1210211232011020-0323133101120011-0222222311010220"></a>

## min_id property — id_ranges / 021322211233 / 5

Type: `"number"`. Optional.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2013121103321023-1311213131121303-3023231013011100-3212321211221300-0012320231331122-0212203012232030-0221130300231223-2032111122133102"></a>

## Next pages — id_ranges / 021322211233 / 6

- [psp_spec.run_as_group](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0303231030223211-2203013023100112-0312011300120231-3121030333202320-1003002031332121-3211020330320023-0121002131022312-0133133232122332)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-3322031322303132-2233300101211301-1033020232212012-3020303310332020-3232213002213301-1001212310130123-1203202133013320-0022212231230221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231202031331231-3101103110210022-3203010300123130-1020311002130321-0322232133201101-3200233212000031-3223001232313213-2210311303303303"></a>

## psp_spec.run_as_user — run_as_user / 032100331312 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.run_as_user

<a id="canonical-1320301203002231-1212023221301203-2022300101100303-1203202120201002-2201201001330113-0130032000133330-1310112231010232-2223230113301021"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for run as user.

Upstream description:

ID ranges and rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule")}
```

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

Terraform syntax:

```terraform
run_as_user {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333132302113123-3110010122101232-1311221011101310-3332111211023230-2002003030021012-0332321332203010-0303331023323320-0111123332220331"></a>

## Direct properties — run_as_user / 032100331312 / 3

- [id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1232013232033210-0321313213303012-2211320012213301-0330112212220012-3102033031001321-1330002000131200-0000200213012122-2220023010301121): complete subsection reference.

<a id="canonical-2120031120022211-1223203023031123-1300003023203102-0212220323212001-0133233123111331-0312111221121212-0120203201320101-2303020221002332"></a>

<a id="canonical-1222312131123211-2321123020000232-1230333203232112-2130220212120323-1023002332112220-2021202112133012-0002123213032131-0223331220300013"></a>

## rule property — run_as_user / 032100331312 / 4

Type: `"string"`. Optional.

Rule indicated how the FS group ID range is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2031303000103213-2121031133000122-2301211121101232-3100023230113101-1212112110002021-3331222113022202-2322001220001023-3300220331010002"></a>

## Next pages — run_as_user / 032100331312 / 5

- [psp_spec.run_as_user.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-1232013232033210-0321313213303012-2211320012213301-0330112212220012-3102033031001321-1330002000131200-0000200213012122-2220023010301121)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-1232013232033210-0321313213303012-2211320012213301-0330112212220012-3102033031001321-1330002000131200-0000200213012122-2220023010301121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113112012101202-2200310302132220-3100222110230222-3230000120123321-2112213003213323-3223101212303300-3223013200103122-0122102022120121"></a>

## psp_spec.run_as_user.id_ranges — id_ranges / 301313322201 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [psp_spec.run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3322031322303132-2233300101211301-1033020232212012-3020303310332020-3232213002213301-1001212310130123-1203202133013320-0022212231230221)
- psp_spec.run_as_user.id_ranges

<a id="canonical-3112201212330321-2300203323311020-0131031311023201-0200333012001222-0133203212132111-3131311030323330-0203003023210222-2002103210302120"></a>

Type: `"object"`. list nested block, Optional.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("max_id",
    "min_id")}
```

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
id_ranges {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322310213020132-1320010323011031-0212022103330033-1300003121203000-0033201331120322-1102321330002232-1220230022200223-3003010020202303"></a>

## Direct properties — id_ranges / 301313322201 / 3

<a id="canonical-1202213130133202-0203110231313121-3002121201333312-0323223310211203-1231030231001112-2231311000221311-1022220003321033-0211000323123011"></a>

<a id="canonical-1231021123203100-3323101111213232-1233222222030322-3331010101202031-1303203010023120-1021033333121123-0002203133221322-3103221111123131"></a>

## max_id property — id_ranges / 301313322201 / 4

Type: `"number"`. Optional.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1312030233122202-3121131132211300-1012201021302310-2003011132033021-0130132322103123-0110313313310133-3200122331321030-1231033001011331"></a>

<a id="canonical-3233120011213200-2221231013322111-3332302022111302-2220100223330001-3201120102221030-1300220100232112-0130000112202212-2230001133232331"></a>

## min_id property — id_ranges / 301313322201 / 5

Type: `"number"`. Optional.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2013122213232313-0011120303101121-3221300113310133-0113211113310223-3001121200101322-0332231010213300-2033320130101223-0033130223102303"></a>

## Next pages — id_ranges / 301313322201 / 6

- [psp_spec.run_as_user](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3322031322303132-2233300101211301-1033020232212012-3020303310332020-3232213002213301-1001212310130123-1203202133013320-0022212231230221)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-3212320333313232-2233312031023110-2223231010003002-2023103303022112-2302230232033013-1011011211221023-1010000011313231-2102321103331011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310221331223130-2213331221320200-3110001202311231-1113322033223312-3032013022020000-3102000220123011-0223230030202331-3233313003312332"></a>

## psp_spec.supplemental_groups — supplemental_groups / 110213022311 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- psp_spec.supplemental_groups

<a id="canonical-1212300222201233-2031221103330210-3303121312232302-2332223110100230-3232121212320123-3300231233032113-2001203220202022-3011333102201003"></a>

Type: `"object"`. single nested block, Optional.

ID(User,Group,FSGroup) Strategy. ID ranges and rules.

Upstream description:

ID ranges and rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule")}
```

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

Terraform syntax:

```terraform
supplemental_groups {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302021102133033-0113013220332023-0030311010322230-2320102022220003-3113021111012313-1130021322302121-3122203011122303-0110102233212222"></a>

## Direct properties — supplemental_groups / 110213022311 / 3

- [id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3201111331020301-0321203222222222-3030211031010121-0131023303302002-1003333330320321-1131220111023023-2321122030120121-0033230301231332): complete subsection reference.

<a id="canonical-1123320120313323-2320310122200020-1212201231023111-0002020101020312-2333323032333002-0113022002031200-1323213100301113-0320301200012031"></a>

<a id="canonical-0223021012311230-2313322223021123-2131002320301221-2311202321002033-0320302001312130-0223221221133033-0000231322111003-2333101202012111"></a>

## rule property — supplemental_groups / 110213022311 / 4

Type: `"string"`. Optional.

Rule indicated how the FS group ID range is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2233313303022121-0010212001330110-0020101300310301-1122300303230200-2111231033211130-0213102121220302-3033012312231012-2233211330000100"></a>

## Next pages — supplemental_groups / 110213022311 / 5

- [psp_spec.supplemental_groups.id_ranges](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3201111331020301-0321203222222222-3030211031010121-0131023303302002-1003333330320321-1131220111023023-2321122030120121-0033230301231332)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-3201111331020301-0321203222222222-3030211031010121-0131023303302002-1003333330320321-1131220111023023-2321122030120121-0033230301231332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223120221333130-2020002212202121-2322210312132030-2031220321330100-0001201300322203-3123003210012310-1330302121031033-1123211220332230"></a>

## psp_spec.supplemental_groups.id_ranges — id_ranges / 132032310320 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [psp_spec](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0311110030320132-0303100100311023-3331303332310330-3213130303020000-1213221002210031-3321231331031122-2102131331313231-3332303002031032)
- [psp_spec.supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3212320333313232-2233312031023110-2223231010003002-2023103303022112-2302230232033013-1011011211221023-1010000011313231-2102321103331011)
- psp_spec.supplemental_groups.id_ranges

<a id="canonical-0320223023232200-1331123112323203-2102223021022023-1122222110221122-0223300211130013-1332112302132311-2132020102320000-3031030302301101"></a>

Type: `"object"`. list nested block, Optional.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("max_id",
    "min_id")}
```

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
id_ranges {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212200030032312-3233311132202213-2203211221323123-2312012223003120-1310212302320332-3221230123202002-2132303022320330-1113113311332303"></a>

## Direct properties — id_ranges / 132032310320 / 3

<a id="canonical-3331223133302222-1213323321011012-0110322031320102-0310013331211230-2010131320121023-3113320202110033-3013022120212230-2331221302021321"></a>

<a id="canonical-3203020133321100-2211331312222221-0002223110132103-0111203300332221-0121230102020000-1001130032112000-1200303002232223-0031132033120000"></a>

## max_id property — id_ranges / 132032310320 / 4

Type: `"number"`. Optional.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3121110213123323-0022112330031123-0311303320000113-2031220022020330-3223300311232233-2113031030301302-1231332020201211-1333200300223222"></a>

<a id="canonical-0322220032111022-1023133131212221-3023123023213211-3002012031200032-2030301022121120-2120322131003212-1200030202221313-0300333022132012"></a>

## min_id property — id_ranges / 132032310320 / 5

Type: `"number"`. Optional.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2002100132032333-3113013300221023-2220232131320013-2010131102203203-0311222002332200-2133321102132011-2032102200100203-1310303223023323"></a>

## Next pages — id_ranges / 132032310320 / 6

- [psp_spec.supplemental_groups](resources--k8s_pod_security_policy--reference--group-001.md#canonical-3212320333313232-2233312031023110-2223231010003002-2023103303022112-2302230232033013-1011011211221023-1010000011313231-2102321103331011)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)

<a id="canonical-0031002323110013-1322030133322223-2221322200102102-0133311200200011-2121212023100030-2232100302333013-3331312322021120-0032320331101121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211323220031201-0200303103221002-0231303110101302-2012032330002202-1001231133003311-0021323302223203-1300212021313201-1202200103302331"></a>

## timeouts — timeouts / 332303333223 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- timeouts

<a id="canonical-0021230111132023-3022313010200233-1130230331003112-0301222132011302-1100302022200321-2333101322332131-3301002100130221-1230023211223222"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000002233113112-0302233200003003-2222201020111213-1033110322100321-2322122033203031-1110300033221313-3303300133220030-3013103320101010"></a>

## Direct properties — timeouts / 332303333223 / 3

<a id="canonical-0233011001333330-2210333202220322-3332121330131302-3313303310320323-2022020132122011-1312332131333021-0233301023311101-1310022123332033"></a>

<a id="canonical-0020211303321113-3322210313200102-1202213312331023-1023001100311121-0313031131232311-3232130132233100-0030233200023010-0033023112233001"></a>

## create property — timeouts / 332303333223 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1330212100221122-0030132221321230-2221102300303301-2312212213021232-0221010221213211-1031233311030003-2002231120222130-1302102122212123"></a>

<a id="canonical-2311231330320303-2000102210131210-3122313221211302-2103212001323322-3120011313133322-3130232000333310-1012203011202120-3300020223301323"></a>

## delete property — timeouts / 332303333223 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0122110121221331-2112002333200331-2010121100003231-1301113312120230-1233233003222310-0303111233001233-1031113010103211-2210320222313302"></a>

<a id="canonical-3312133201022113-2313110230001131-3110213023230103-1013231001330112-1101011121323232-3012301020333102-2220003323023212-3123133310320202"></a>

## read property — timeouts / 332303333223 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1303111112330102-3121023110323210-3030100311030321-0230330013001200-0331200313312120-2110333121333030-0121300231111203-3232213322331103"></a>

<a id="canonical-0000203003110323-1220102211211112-1030230330110111-2002202233000220-3123300031220130-2312100002020221-1310221032110010-2212222301222133"></a>

## update property — timeouts / 332303333223 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0321322020331000-1210001332200311-3320131331003020-3120130112213203-2103103322030121-0303003123220023-3030313110103222-3113031111132122"></a>

## Next pages — timeouts / 332303333223 / 8

- [Property reference](resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
