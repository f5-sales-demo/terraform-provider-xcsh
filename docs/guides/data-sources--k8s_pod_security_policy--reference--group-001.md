---
page_title: "xcsh_k8s_pod_security_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy reference."
---

# xcsh_k8s_pod_security_policy reference

<a id="canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110312223101220-1323100021322222-1010323112310020-1221112123021020-0001223022201130-2000202222220021-2333232023231101-2302232221120220"></a>

## Property reference — Property reference / 212122103232 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- Property reference

<a id="canonical-0001002312332210-3332221200001002-3101130303131003-2101001231310002-3120103302313122-3131123333202203-0221130322023303-2212311233013203"></a>

## Direct properties — Property reference / 212122103232 / 3

<a id="canonical-3111113123103011-1131033033113310-2102003122301330-3132210203233213-0222210200311232-1121002123321322-1131303000100231-1321330233021023"></a>

<a id="canonical-3022123112303120-2221332033321121-1023312111220010-3323103033012023-3301212031220020-2100230021001300-2213332201011322-1113100011121010"></a>

## annotations property — Property reference / 212122103232 / 4

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

<a id="canonical-2100223030223310-1102013121123331-1233023332232231-2121112211001033-1210032301332000-2310313323202102-2321222323122232-1113012101012232"></a>

<a id="canonical-1320212311021101-0101003310222031-3232332302020011-1313210231332030-1122321223031200-3310120001300003-3123130001023121-1210122320101212"></a>

## description property — Property reference / 212122103232 / 5

Type: `"string"`. Computed.

Description of the K8SPodSecurityPolicy.

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

<a id="canonical-2331000123100101-0103021302002031-1212203013103121-0122202131121101-1313102113313200-3010001211130011-2113321030020212-1321113021203012"></a>

<a id="canonical-3003122111212021-3210231200130202-0322032033033231-0303221302111022-1021322011120203-3131231222023020-2220232132223312-3031110310033200"></a>

## ID property — Property reference / 212122103232 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3003003102230022-2312202312001313-3210032132211201-3221232220333013-3203310312023102-1032100013231033-3321103001223122-0212222330223203"></a>

<a id="canonical-2102002103021131-2111201312222131-2132333313133311-0203022333131321-3123033023302332-3301301101011233-3231230230010013-0303202230220203"></a>

## labels property — Property reference / 212122103232 / 7

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

<a id="canonical-1301200132222322-1002111200003102-0131202302100300-3231320300112132-3320232002011200-0012111332121032-3000311001013210-0310203311233102"></a>

<a id="canonical-2033102100032112-1233100013221011-1213323113320222-3112230302320033-1000120120322022-0213332302220333-0101230031011000-0211202030133031"></a>

## name property — Property reference / 212122103232 / 8

Type: `"string"`. Required.

Name of the K8SPodSecurityPolicy.

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

<a id="canonical-1210033233222223-1231310001310020-0232112133200011-1223011100232200-1000232133330120-2000330010231021-0121313010202231-2033130320232301"></a>

<a id="canonical-2303013000013030-0120213003330113-0023331231030223-2003210002020113-3023002213213013-2123333321001101-0013212133312120-2220323120033030"></a>

## namespace property — Property reference / 212122103232 / 9

Type: `"string"`. Required.

Namespace where the K8SPodSecurityPolicy exists.

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

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033): complete subsection reference.

<a id="canonical-1130200223201022-2021310133101112-0023211131211011-1121322232122031-1003320001332320-3332123123011103-3132212123330103-3120021320322333"></a>

<a id="canonical-0023113323131131-3120021002101313-0003102003130212-3221011210001003-0011012302201001-3213310313110033-0220310312200201-1031031221102130"></a>

## YAML property — Property reference / 212122103232 / 10

Type: `"string"`. Computed.

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

Upstream description:

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2103012323233222-0103003201200011-1112333013230333-1220101312133233-2322223031330022-3123302021332001-1221101113013223-0313330022131033"></a>

## All schema paths — Property reference / 212122103232 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3111113123103011-1131033033113310-2102003122301330-3132210203233213-0222210200311232-1121002123321322-1131303000100231-1321330233021023) |
| `description` | [description](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2100223030223310-1102013121123331-1233023332232231-2121112211001033-1210032301332000-2310313323202102-2321222323122232-1113012101012232) |
| `id` | [ID](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2331000123100101-0103021302002031-1212203013103121-0122202131121101-1313102113313200-3010001211130011-2113321030020212-1321113021203012) |
| `labels` | [labels](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3003003102230022-2312202312001313-3210032132211201-3221232220333013-3203310312023102-1032100013231033-3321103001223122-0212222330223203) |
| `name` | [name](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1301200132222322-1002111200003102-0131202302100300-3231320300112132-3320232002011200-0012111332121032-3000311001013210-0310203311233102) |
| `namespace` | [namespace](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1210033233222223-1231310001310020-0232112133200011-1223011100232200-1000232133330120-2000330010231021-0121313010202231-2033130320232301) |
| `psp_spec` | [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0310321221222330-0323013203323212-2221232321023210-3033312222221030-1001230130002223-1103123322032322-0221000013201120-0133303222223110) |
| `psp_spec.allow_privilege_escalation` | [psp_spec.allow_privilege_escalation](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2012001230002012-2031101331211231-1120203022301132-3213311101312333-0110012200123300-2130203333311303-1212203300000122-2323201323210032) |
| `psp_spec.allowed_capabilities` | [psp_spec.allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3331331100211301-0030330120032231-1123103101322030-3032013321323032-2122202020023302-2320232203232320-2201103031310310-3213101321322013) |
| `psp_spec.allowed_capabilities.capabilities` | [psp_spec.allowed_capabilities.capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3112223020122010-2100322313221301-2002021022013010-2002212323213210-2230323111133211-0223330130010103-0323001123321322-3312100121132030) |
| `psp_spec.allowed_csi_drivers` | [psp_spec.allowed_csi_drivers](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0321233320032021-2011211112333011-3001331220210232-2012323303102000-0333102313010030-3210021202012300-1333101131001123-1312310211232320) |
| `psp_spec.allowed_flex_volumes` | [psp_spec.allowed_flex_volumes](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2130000022231030-1120032001130020-1022223121320210-1322132233210300-3013022033331001-3111111333121302-0303032100130313-2110232200021233) |
| `psp_spec.allowed_host_paths` | [psp_spec.allowed_host_paths](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2133022132210000-0312333111300111-2023130303333313-0322303131011130-0212211102323102-3201200322323022-0232001022031311-0132022133203320) |
| `psp_spec.allowed_host_paths.path_prefix` | [psp_spec.allowed_host_paths.path_prefix](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0333031030302231-3200330223012210-3330323303131320-0313132320300030-3211103210230002-1311030010011131-3233102323131321-1023000021130233) |
| `psp_spec.allowed_host_paths.read_only` | [psp_spec.allowed_host_paths.read_only](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0302201113221131-1200313302222302-0203112123233231-2200022210310322-3303112112230032-0232000122303200-1133230313021231-1100213211033203) |
| `psp_spec.allowed_proc_mounts` | [psp_spec.allowed_proc_mounts](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1303330132110313-3301012231122223-2210122102022201-2020323002001222-2222233322120210-3113030001223310-2032333013330231-3301321331023021) |
| `psp_spec.allowed_unsafe_sysctls` | [psp_spec.allowed_unsafe_sysctls](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0320200313101322-2122023313212223-3011200212010032-3313013000321200-0321201110133333-3111212232213230-3022211200233203-3301102322330102) |
| `psp_spec.default_allow_privilege_escalation` | [psp_spec.default_allow_privilege_escalation](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1013020130223230-0101020300213112-0310232022311330-3030312010121303-3032231300310123-1330221233221221-1021301123002113-0221022121203231) |
| `psp_spec.default_capabilities` | [psp_spec.default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1220002000233002-2302232021102030-2132300123203132-0013330310313111-0103000220033002-1101020032033313-2031311031211123-0132323012030031) |
| `psp_spec.default_capabilities.capabilities` | [psp_spec.default_capabilities.capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1311322023003222-3032311211010201-1302221033232100-3021030033003030-0022010302023111-0103000102010312-0311202300131021-0301223330331213) |
| `psp_spec.drop_capabilities` | [psp_spec.drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3021230301022111-2101103331233011-2023303123332131-0112030312123320-3330133103321232-0003303132311022-3030232203110000-3111300322202000) |
| `psp_spec.drop_capabilities.capabilities` | [psp_spec.drop_capabilities.capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2111212330031323-3133202312101210-0113013121230110-3020021013021202-3300102022222101-2332332222212123-2112220200010013-0211130130332031) |
| `psp_spec.forbidden_sysctls` | [psp_spec.forbidden_sysctls](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0323213130021003-1301123201233031-1300312232333210-2321322230311131-2120202301123032-3012133233313123-1100330101233231-1332221330323210) |
| `psp_spec.fs_group_strategy_options` | [psp_spec.fs_group_strategy_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0123321301101333-3120122010031211-0113222211310123-3002213323122013-0210110130312210-0212101021310013-2211111132133222-3121210122102313) |
| `psp_spec.fs_group_strategy_options.id_ranges` | [psp_spec.fs_group_strategy_options.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1323120111332320-0330113133100100-2031202111003201-1300303130311230-2100131203131110-1223103112120003-1301203130110200-0333210330300320) |
| `psp_spec.fs_group_strategy_options.id_ranges.max_id` | [psp_spec.fs_group_strategy_options.id_ranges.max_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1211222310022320-2120322020002102-0332112021333331-1221023210112100-2311331232113120-3312001323220130-0121012011122022-3122313100301211) |
| `psp_spec.fs_group_strategy_options.id_ranges.min_id` | [psp_spec.fs_group_strategy_options.id_ranges.min_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2203121231300003-3301330003232011-0023200110030323-3230021130032311-3000012332133321-1303220311220333-0100322331313103-0133110131023321) |
| `psp_spec.fs_group_strategy_options.rule` | [psp_spec.fs_group_strategy_options.rule](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3130222112122122-1310223112333231-0330222201301332-1103003302303212-0330130230230310-1013232023100211-0211010133221132-2311322132130222) |
| `psp_spec.host_ipc` | [psp_spec.host_ipc](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2101110101221311-3133312220133031-0213012331323313-3022211133231301-3033333220010212-1202032130022123-2203002023010021-2132133320320233) |
| `psp_spec.host_network` | [psp_spec.host_network](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0123001031033020-3332112111013122-3303211333023013-1003013000232021-0310331011000030-0122013002331111-1222002123120203-1011220122330333) |
| `psp_spec.host_pid` | [psp_spec.host_pid](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0033301131212012-0123203211120330-3102321131001102-2330112300103023-2013202133021231-1301231331130301-1320031001232110-0310121133313010) |
| `psp_spec.host_port_ranges` | [psp_spec.host_port_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1312210020100121-2321213101101210-1002321221212031-1300322303021100-2230312231210130-3012303302022111-2000033301011131-0120031033220300) |
| `psp_spec.no_allowed_capabilities` | [psp_spec.no_allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0102331032131022-3201332322300120-0001113133223033-3313333323122301-3022331100310021-1012130131033101-0103032332110101-2023222131132033) |
| `psp_spec.no_default_capabilities` | [psp_spec.no_default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2330101030022112-2311010233032230-2120033102012332-1031130220031310-1102033211010333-0302112331103302-0032310130110003-3320010131131331) |
| `psp_spec.no_drop_capabilities` | [psp_spec.no_drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1103001030213130-2032230231301012-1330312130132213-3022203330010112-2111322011230002-2221323130030111-1000233322011101-1123211323223023) |
| `psp_spec.no_fs_groups` | [psp_spec.no_fs_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2001033320020223-0113110031110333-0030013133030223-3032000100203121-0021331033222322-3232100222333303-3011211102203323-2223003223310210) |
| `psp_spec.no_run_as_group` | [psp_spec.no_run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2000010000103221-1021310123202103-2122121122021123-0022211122200020-2320002103232201-0311203031303002-1130001313333223-3032211131022220) |
| `psp_spec.no_run_as_user` | [psp_spec.no_run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1231203103200000-1223022000203013-2100110030201132-2320103332303011-1313211221231302-3112302112303030-1101312302030101-2010121102313023) |
| `psp_spec.no_runtime_class` | [psp_spec.no_runtime_class](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3202101213303123-0000321002331331-3011000131333112-0032331330013213-0211033212121201-0012020221113203-0311202201220313-2320022333003302) |
| `psp_spec.no_se_linux_options` | [psp_spec.no_se_linux_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0330132130213201-3121120100130222-1132313213202022-2330232313020221-1311002112110031-1121123312200122-1023003223113223-2323013210010232) |
| `psp_spec.no_supplemental_groups` | [psp_spec.no_supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2103002300000201-2331231310331100-0300230202132102-1202230123122313-1110022100231123-1113002100320333-1201102030003301-1020212133322303) |
| `psp_spec.privileged` | [psp_spec.privileged](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0323112020302321-2002321122223320-0023103100200001-3101232230302323-0230302132302103-1230102222023220-1202032300001222-2232012221310303) |
| `psp_spec.read_only_root_filesystem` | [psp_spec.read_only_root_filesystem](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0101321331102001-1320120103302321-2330131001222120-0223311033311302-3112320021031113-3000010323321210-0111131312001232-3212033011011230) |
| `psp_spec.run_as_group` | [psp_spec.run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1003002133000312-3222113302132232-1303201313133120-3012233300233322-2010123302121131-1320123213203013-2232003233330220-0030203000323011) |
| `psp_spec.run_as_group.id_ranges` | [psp_spec.run_as_group.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0002233013030230-2111230012312010-3013131311303023-1123213001001010-0201322123310030-0003021310312102-0310210022000023-2200210232100022) |
| `psp_spec.run_as_group.id_ranges.max_id` | [psp_spec.run_as_group.id_ranges.max_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3013312301001131-1000120333000302-2301131020223302-1131101202020300-1313133310201022-1230230203023303-0220223331021032-3330310111203222) |
| `psp_spec.run_as_group.id_ranges.min_id` | [psp_spec.run_as_group.id_ranges.min_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1212022223130010-3001330322113023-2320332013103203-1221321002323212-3123322020003231-1323303022020031-0120230111003120-1330111321000302) |
| `psp_spec.run_as_group.rule` | [psp_spec.run_as_group.rule](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1030133133123220-3331213330131310-1030202303222220-0000020220111231-2123230232130030-2023010232322032-1102130301331111-3113033011131300) |
| `psp_spec.run_as_user` | [psp_spec.run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0121123133323102-0310312122222221-1103211102103231-2202201020121120-0113111312202023-1110331021212023-2202013022231203-2212102101321111) |
| `psp_spec.run_as_user.id_ranges` | [psp_spec.run_as_user.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2213101120321200-2231103130300300-1100112023000000-3200002013023000-2211202022120133-2010210022322000-1221002202212103-1310221130212212) |
| `psp_spec.run_as_user.id_ranges.max_id` | [psp_spec.run_as_user.id_ranges.max_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2203133321111030-3110323123131321-1202310103232200-1232121000032203-1130011110200000-2213302322302103-1001233222110103-3003201112312111) |
| `psp_spec.run_as_user.id_ranges.min_id` | [psp_spec.run_as_user.id_ranges.min_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0331311311223031-3312033121030031-3200313133310311-3231321312103213-1001231320113310-3310312012330220-2213120003023232-2013211212011331) |
| `psp_spec.run_as_user.rule` | [psp_spec.run_as_user.rule](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3110302321122132-1113001210012033-0120003121033130-0033311223212322-1121300120330300-2301300122223102-2312133212103310-0010030332231031) |
| `psp_spec.supplemental_groups` | [psp_spec.supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3003312122032133-3202130210011120-0210231032003223-3131223311333132-1323121021333122-2312212111310112-2003020212333022-1212100203213131) |
| `psp_spec.supplemental_groups.id_ranges` | [psp_spec.supplemental_groups.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2001010103033331-2023022030133323-1113211013113212-2210320200110211-1222013323022030-1310030333200202-2231201220121313-3001211030123111) |
| `psp_spec.supplemental_groups.id_ranges.max_id` | [psp_spec.supplemental_groups.id_ranges.max_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0010300312313320-1123100311322020-0212210113113233-3311112023223302-2330202101122301-2111332130110111-0032120220210331-2231302320320310) |
| `psp_spec.supplemental_groups.id_ranges.min_id` | [psp_spec.supplemental_groups.id_ranges.min_id](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0012232323310302-0200332332331311-0322313303112321-3001313233133033-2231030223212330-1301131032023232-1330330002110302-0300233100333322) |
| `psp_spec.supplemental_groups.rule` | [psp_spec.supplemental_groups.rule](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1130032303323213-2133320113301132-2320230213012121-1332100310100233-2103131101313010-2212003321011020-1020123320001010-0321233010133310) |
| `psp_spec.volumes` | [psp_spec.volumes](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0203231212211020-2031100201100302-3001031132133201-0020121010100112-2321010222203320-1001200330230202-1033022023110333-0122222113323320) |
| `yaml` | [YAML](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1130200223201022-2021310133101112-0023211131211011-1121322232122031-1003320001332320-3332123123011103-3132212123330103-3120021320322333) |

<a id="canonical-3203132122033312-3131323302102322-2321131000111123-3102002003013301-3002010322313302-3010121320220333-0133211100122311-0022311022323002"></a>

## Next pages — Property reference / 212122103232 / 12

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013302113302013-1100231301101123-0320101011202201-3010013012230322-2230201102021021-2310022233203322-2113131223312201-0312023333302012"></a>

## psp_spec — psp_spec / 223302101030 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- psp_spec

<a id="canonical-0310321221222330-0323013203323212-2221232321023210-3033312222221030-1001230130002223-1103123322032322-0221000013201120-0133303222223110"></a>

Type: `"single"`. Computed.

\[OneOf: psp\_spec, YAML\] Pod Security Policy Specification. Form based pod security specification.

Upstream description:

Form based pod security specification.

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

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0310321221222330-0323013203323212-2221232321023210-3033312222221030-1001230130002223-1103123322032322-0221000013201120-0133303222223110)
- [YAML](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1130200223201022-2021310133101112-0023211131211011-1121322232122031-1003320001332320-3332123123011103-3132212123330103-3120021320322333)

Select alternatives according to the provider validators above.

<a id="canonical-1203100331101300-0132301021322001-3100112222003113-2033022322213033-3201333012200111-0321103311010302-3001012323023120-1022012023320023"></a>

## Direct properties — psp_spec / 223302101030 / 3

<a id="canonical-2012001230002012-2031101331211231-1120203022301132-3213311101312333-0110012200123300-2130203333311303-1212203300000122-2323201323210032"></a>

<a id="canonical-0311222003300231-1130000312031313-1230301301112333-3131010101103030-0203122203002321-1130001010031202-0003332320123001-1102133101002311"></a>

## allow_privilege_escalation property — psp_spec / 223302101030 / 4

Type: `"bool"`. Computed.

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

- [allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2120301132322033-3220133213213031-3022230331303232-2030022211010133-1223013202210223-1322012213102021-1123121021010100-2200231031320302): complete subsection reference.

<a id="canonical-0321233320032021-2011211112333011-3001331220210232-2012323303102000-0333102313010030-3210021202012300-1333101131001123-1312310211232320"></a>

<a id="canonical-0131100011131323-2322210010003021-2320202232330130-0122312002222323-2121313200003120-0102110212120333-2122212322232332-0103202110223212"></a>

## allowed_csi_drivers property — psp_spec / 223302101030 / 5

Type: `["list", "string"]`. Computed.

Restrict the available CSI drivers for POD, default all drivers are available.

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

<a id="canonical-2130000022231030-1120032001130020-1022223121320210-1322132233210300-3013022033331001-3111111333121302-0303032100130313-2110232200021233"></a>

<a id="canonical-3031233220013202-1033330301233213-2302331220302110-3310230002220211-1103133321330100-2132201201100312-1112020222012323-1012033023331030"></a>

## allowed_flex_volumes property — psp_spec / 223302101030 / 6

Type: `["list", "string"]`. Computed.

Restrict list of Flex volumes, default all volumes are allowed.

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

- [allowed_host_paths](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0101033331100131-2202222132303330-1000030301221232-1201012100231322-2330013133031322-3011110101120113-2320013210130323-3210210103131231): complete subsection reference.

<a id="canonical-1303330132110313-3301012231122223-2210122102022201-2020323002001222-2222233322120210-3113030001223310-2032333013330231-3301321331023021"></a>

<a id="canonical-3010220331011120-3201313301133221-1013102020311102-0001222031020200-2021033130203101-0000313101103030-2001001102303312-3013030233210302"></a>

## allowed_proc_mounts property — psp_spec / 223302101030 / 7

Type: `["list", "string"]`. Computed.

Allowed list of proc mounts, empty list allows default proc mounts.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0320200313101322-2122023313212223-3011200212010032-3313013000321200-0321201110133333-3111212232213230-3022211200233203-3301102322330102"></a>

<a id="canonical-0013312033113013-2220001212230312-2113202010222131-0320133312030013-3333201223033213-1123322122232332-1132231200231202-3232032101102033"></a>

## allowed_unsafe_sysctls property — psp_spec / 223302101030 / 8

Type: `["list", "string"]`. Computed.

Allowed list of unsafe sysctls, empty list allows none. Supports prefix reg-ex.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1013020130223230-0101020300213112-0310232022311330-3030312010121303-3032231300310123-1330221233221221-1021301123002113-0221022121203231"></a>

<a id="canonical-2110303211002103-0332030200022230-3020313323002001-0031313131101031-0210001312133311-1202223020013003-1032303123001230-3010321130111001"></a>

## default_allow_privilege_escalation property — psp_spec / 223302101030 / 9

Type: `"bool"`. Computed.

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

- [default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3102110120302123-3302032323132310-1230323222210021-2100012211120023-2123030331002122-3023133330000222-2233110332303122-0201112332313302): complete subsection reference.

- [drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0300011102223101-3030213220010330-3220230101032201-0032331022322000-3020201202311233-1121013210030111-2303020101303201-0301120300011230): complete subsection reference.

<a id="canonical-0323213130021003-1301123201233031-1300312232333210-2321322230311131-2120202301123032-3012133233313123-1100330101233231-1332221330323210"></a>

<a id="canonical-1111210313013230-3100021333111200-3130232120032221-2000211222301001-0022213213023002-1012230032213002-3013101011122030-3311222303211113"></a>

## forbidden_sysctls property — psp_spec / 223302101030 / 10

Type: `["list", "string"]`. Computed.

Forbidden list of sysctls, empty list forbids none. Supports prefix reg-ex.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [fs_group_strategy_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1101333110023110-2311233020333110-0032311033310122-3103011323131313-1022123333302101-0200100132232213-1033101102031102-0022303221320231): complete subsection reference.

<a id="canonical-2101110101221311-3133312220133031-0213012331323313-3022211133231301-3033333220010212-1202032130022123-2203002023010021-2132133320320233"></a>

<a id="canonical-3122012233301102-1100123021130331-3113230032120310-2201023200113213-1030320333300311-3122010113233033-1213302302320112-0022331313120113"></a>

## host_ipc property — psp_spec / 223302101030 / 11

Type: `"bool"`. Computed.

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

<a id="canonical-0123001031033020-3332112111013122-3303211333023013-1003013000232021-0310331011000030-0122013002331111-1222002123120203-1011220122330333"></a>

<a id="canonical-2331221130021131-2031030310011313-1232322010221131-2032133312322030-1232030123311010-2110022210103001-1332210011010210-2202321302300330"></a>

## host_network property — psp_spec / 223302101030 / 12

Type: `"bool"`. Computed.

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

<a id="canonical-0033301131212012-0123203211120330-3102321131001102-2330112300103023-2013202133021231-1301231331130301-1320031001232110-0310121133313010"></a>

<a id="canonical-3100211100310301-3220202102312211-3132001002112022-3312210212020113-2331322301222222-0221221333103103-1220033201310232-2211111321032133"></a>

## host_pid property — psp_spec / 223302101030 / 13

Type: `"bool"`. Computed.

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

<a id="canonical-1312210020100121-2321213101101210-1002321221212031-1300322303021100-2230312231210130-3012303302022111-2000033301011131-0120031033220300"></a>

<a id="canonical-2323113303110230-2320320221100130-0221003101233123-3112103032221230-0310033320133022-2231331312201112-3322030230032121-1023013322202333"></a>

## host_port_ranges property — psp_spec / 223302101030 / 14

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

- [no_allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2010013123031310-3211213220010312-0310000023233332-3100323323010221-0332101332122222-3220102332203132-1223121003033110-3311320222302103): complete subsection reference.

- [no_default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3013311031321132-3311012331112203-2121021111012313-0321121303333132-2323103313331233-1221132113121010-3001122111233000-2323310220100200): complete subsection reference.

- [no_drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0022102100203333-0233121320103112-3223201023013302-2123102300020312-2323102030020203-1100200333332203-1121003121233010-0220120101203303): complete subsection reference.

- [no_fs_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2120121233213121-2213211103022200-2021230003320022-0330301021323223-2103031103113323-2030300030033120-2112210212121112-2022212130123323): complete subsection reference.

- [no_run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1001302122030322-1233303133231222-2311133230020122-1213211121231122-2321302013031030-0022310003231232-1313223232003111-2220103033103123): complete subsection reference.

- [no_run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0231310033201230-0230120223002112-0321312331201101-2313111012202332-0113313123013320-2233011213310021-0020203330133032-3100210130120311): complete subsection reference.

- [no_runtime_class](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2320003122022032-1200110113102333-2030212232102323-0103303130233201-1302320313211031-0221323112321330-0330303133130102-1102022012300101): complete subsection reference.

- [no_se_linux_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3033121311022123-1233232113312211-1000103330023133-3200031030302220-1121111322213232-0303122111223313-2110122133332120-0100001023031123): complete subsection reference.

- [no_supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1000112100012022-3030310223032102-2331000103320032-0031112102321123-2100210202111230-0313033312101103-2003313131212103-2321220133001330): complete subsection reference.

<a id="canonical-0323112020302321-2002321122223320-0023103100200001-3101232230302323-0230302132302103-1230102222023220-1202032300001222-2232012221310303"></a>

<a id="canonical-3020330021000321-0132003322333023-1213201101010330-3113131010200022-0000112222032223-2302101113002312-1111203222300321-0231213021333333"></a>

## privileged property — psp_spec / 223302101030 / 15

Type: `"bool"`. Computed.

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

<a id="canonical-0101321331102001-1320120103302321-2330131001222120-0223311033311302-3112320021031113-3000010323321210-0111131312001232-3212033011011230"></a>

<a id="canonical-0332000210122200-0010102331230122-1223233302013331-3111231111133320-0333230120102212-2111210103233211-0011130110212301-2313133103020132"></a>

## read_only_root_filesystem property — psp_spec / 223302101030 / 16

Type: `"bool"`. Computed.

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

- [run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1000330010132213-0013201221322322-3310132230301122-0231103202203321-3203000002330032-0023030100013121-3203212122101103-2331010222302333): complete subsection reference.

- [run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2211301322001300-1211102230212212-3222112002333031-0321030003031212-2312222233032220-2020331210031013-2231123302321131-2103033201020311): complete subsection reference.

- [supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0312231101331321-2213322320201323-3232313110332322-1030122210300331-0322302003312220-2311303203230011-1023102100212132-1232002030111120): complete subsection reference.

<a id="canonical-0203231212211020-2031100201100302-3001031132133201-0020121010100112-2321010222203320-1001200330230202-1033022023110333-0122222113323320"></a>

<a id="canonical-0222013010030011-0233203302311111-0013230332331233-0231020302131022-3033212120312001-2233231302000301-0310211120212130-1223300333203100"></a>

## volumes property — psp_spec / 223302101030 / 17

Type: `["list", "string"]`. Computed.

Allow List of volume plugins. Empty no volumes are allowed.

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

<a id="canonical-1320202311101120-3213233200121312-3101130030133321-3210211203012201-1013212302302133-3112232210122100-2100312321001331-2323310233112003"></a>

## Next pages — psp_spec / 223302101030 / 18

- [psp_spec.allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2120301132322033-3220133213213031-3022230331303232-2030022211010133-1223013202210223-1322012213102021-1123121021010100-2200231031320302)
- [psp_spec.allowed_host_paths](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0101033331100131-2202222132303330-1000030301221232-1201012100231322-2330013133031322-3011110101120113-2320013210130323-3210210103131231)
- [psp_spec.default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3102110120302123-3302032323132310-1230323222210021-2100012211120023-2123030331002122-3023133330000222-2233110332303122-0201112332313302)
- [psp_spec.drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0300011102223101-3030213220010330-3220230101032201-0032331022322000-3020201202311233-1121013210030111-2303020101303201-0301120300011230)
- [psp_spec.fs_group_strategy_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1101333110023110-2311233020333110-0032311033310122-3103011323131313-1022123333302101-0200100132232213-1033101102031102-0022303221320231)
- [psp_spec.no_allowed_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2010013123031310-3211213220010312-0310000023233332-3100323323010221-0332101332122222-3220102332203132-1223121003033110-3311320222302103)
- [psp_spec.no_default_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3013311031321132-3311012331112203-2121021111012313-0321121303333132-2323103313331233-1221132113121010-3001122111233000-2323310220100200)
- [psp_spec.no_drop_capabilities](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0022102100203333-0233121320103112-3223201023013302-2123102300020312-2323102030020203-1100200333332203-1121003121233010-0220120101203303)
- [psp_spec.no_fs_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2120121233213121-2213211103022200-2021230003320022-0330301021323223-2103031103113323-2030300030033120-2112210212121112-2022212130123323)
- [psp_spec.no_run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1001302122030322-1233303133231222-2311133230020122-1213211121231122-2321302013031030-0022310003231232-1313223232003111-2220103033103123)
- [psp_spec.no_run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0231310033201230-0230120223002112-0321312331201101-2313111012202332-0113313123013320-2233011213310021-0020203330133032-3100210130120311)
- [psp_spec.no_runtime_class](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2320003122022032-1200110113102333-2030212232102323-0103303130233201-1302320313211031-0221323112321330-0330303133130102-1102022012300101)
- [psp_spec.no_se_linux_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3033121311022123-1233232113312211-1000103330023133-3200031030302220-1121111322213232-0303122111223313-2110122133332120-0100001023031123)
- [psp_spec.no_supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1000112100012022-3030310223032102-2331000103320032-0031112102321123-2100210202111230-0313033312101103-2003313131212103-2321220133001330)
- [psp_spec.run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1000330010132213-0013201221322322-3310132230301122-0231103202203321-3203000002330032-0023030100013121-3203212122101103-2331010222302333)
- [psp_spec.run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2211301322001300-1211102230212212-3222112002333031-0321030003031212-2312222233032220-2020331210031013-2231123302321131-2103033201020311)
- [psp_spec.supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0312231101331321-2213322320201323-3232313110332322-1030122210300331-0322302003312220-2311303203230011-1023102100212132-1232002030111120)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-2120301132322033-3220133213213031-3022230331303232-2030022211010133-1223013202210223-1322012213102021-1123121021010100-2200231031320302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120211003212211-0001021032012201-1022313330331013-2311333003322111-1003312203033021-2033232021112322-1031133132130200-2033003131102331"></a>

## psp_spec.allowed_capabilities — allowed_capabilities / 100130312001 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.allowed_capabilities

<a id="canonical-3331331100211301-0030330120032231-1123103101322030-3032013321323032-2122202020023302-2320232203232320-2201103031310310-3213101321322013"></a>

Type: `"single"`. Computed.

List of capabilities that Docker container has.

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

<a id="canonical-1210120311213012-0320013211201321-3102213323223321-0303200211023313-1031231112131330-1313333120330312-1233010100210111-0321232023132000"></a>

## Direct properties — allowed_capabilities / 100130312001 / 3

<a id="canonical-3112223020122010-2100322313221301-2002021022013010-2002212323213210-2230323111133211-0223330130010103-0323001123321322-3312100121132030"></a>

<a id="canonical-2301100122333120-2132032201022010-3112231113302331-1301133013300203-3012130112100221-1102312110113121-1013222132013200-1032303010031301"></a>

## capabilities property — allowed_capabilities / 100130312001 / 4

Type: `["list", "string"]`. Computed.

List of capabilities that Docker container has.

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

<a id="canonical-3332112021331123-0302111220022303-1232032232222123-2311110013112310-0033201300211023-3130031220003323-2201320303112323-2203100331310001"></a>

## Next pages — allowed_capabilities / 100130312001 / 5

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-0101033331100131-2202222132303330-1000030301221232-1201012100231322-2330013133031322-3011110101120113-2320013210130323-3210210103131231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320110113232300-1132320130012103-2312031103013111-0222213300110300-2333301311222231-1310231301311301-1132231131132003-1231331031220230"></a>

## psp_spec.allowed_host_paths — allowed_host_paths / 102233213123 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.allowed_host_paths

<a id="canonical-2133022132210000-0312333111300111-2023130303333313-0322303131011130-0212211102323102-3201200322323022-0232001022031311-0132022133203320"></a>

Type: `"list"`. Computed.

Restrict list of host paths, default all host paths are allowed.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1332200112111310-2301110102111320-2121102132131211-3320131112331030-0211032333113322-0313032101202220-1022333122223102-1313301310202320"></a>

## Direct properties — allowed_host_paths / 102233213123 / 3

<a id="canonical-0333031030302231-3200330223012210-3330323303131320-0313132320300030-3211103210230002-1311030010011131-3233102323131321-1023000021130233"></a>

<a id="canonical-3230001320231003-1203032333021012-2333332202333201-2013313300332213-0001212132012321-2230031012111013-1012221031320111-2231113100101232"></a>

## path_prefix property — allowed_host_paths / 102233213123 / 4

Type: `"string"`. Computed.

Host path prefix is the path prefix that the host volume must match. It does not support \*.

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

<a id="canonical-0302201113221131-1200313302222302-0203112123233231-2200022210310322-3303112112230032-0232000122303200-1133230313021231-1100213211033203"></a>

<a id="canonical-1030012121322330-1330201020113201-2110030310113200-0332303011010212-1101300212213300-2302002030032100-3030301321122022-0222222110330322"></a>

## read_only property — allowed_host_paths / 102233213123 / 5

Type: `"bool"`. Computed.

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

<a id="canonical-2210233023222131-2232213003121113-2002001313130202-2203321000013232-2133320001021132-3221010131000221-0032022111220011-0013300102032130"></a>

## Next pages — allowed_host_paths / 102233213123 / 6

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-3102110120302123-3302032323132310-1230323222210021-2100012211120023-2123030331002122-3023133330000222-2233110332303122-0201112332313302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300033012321322-0203233033333102-3121010023111002-0332331201200110-2221033112123311-3213013301132110-1003100133331210-1221323312100330"></a>

## psp_spec.default_capabilities — default_capabilities / 031313000022 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.default_capabilities

<a id="canonical-1220002000233002-2302232021102030-2132300123203132-0013330310313111-0103000220033002-1101020032033313-2031311031211123-0132323012030031"></a>

Type: `"single"`. Computed.

List of capabilities that Docker container has.

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

<a id="canonical-0013211233032222-0120211112112311-1112133023103012-0322103110000133-2013033122200201-2313223231032023-3101231211131100-1003301122103133"></a>

## Direct properties — default_capabilities / 031313000022 / 3

<a id="canonical-1311322023003222-3032311211010201-1302221033232100-3021030033003030-0022010302023111-0103000102010312-0311202300131021-0301223330331213"></a>

<a id="canonical-3003320210123033-1201120013133233-2230200001222112-3232310121202123-3331323231202331-2031130110000023-3032021003230302-3123312013313101"></a>

## capabilities property — default_capabilities / 031313000022 / 4

Type: `["list", "string"]`. Computed.

List of capabilities that Docker container has.

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

<a id="canonical-1303123032031100-0221001030232222-1122121301201322-3332333132322232-3230110303011300-3121101011020023-2022332012023123-3100100100120310"></a>

## Next pages — default_capabilities / 031313000022 / 5

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-0300011102223101-3030213220010330-3220230101032201-0032331022322000-3020201202311233-1121013210030111-2303020101303201-0301120300011230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213311333312231-2110200011221230-0220322113110302-2320130313333111-3312130022200310-1121213320033101-2031200210003223-2223302033212233"></a>

## psp_spec.drop_capabilities — drop_capabilities / 202110133112 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.drop_capabilities

<a id="canonical-3021230301022111-2101103331233011-2023303123332131-0112030312123320-3330133103321232-0003303132311022-3030232203110000-3111300322202000"></a>

Type: `"single"`. Computed.

List of capabilities that Docker container has.

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

<a id="canonical-3203131121031333-3120012100331323-1132203303320330-3331323333301233-3133311320122001-1212100311311131-1020122300012130-2011231132130112"></a>

## Direct properties — drop_capabilities / 202110133112 / 3

<a id="canonical-2111212330031323-3133202312101210-0113013121230110-3020021013021202-3300102022222101-2332332222212123-2112220200010013-0211130130332031"></a>

<a id="canonical-3021023113021321-2103333030323102-1333132101202201-2122220211323032-2110210212013022-2021032130313310-1311011033222312-1011011232002131"></a>

## capabilities property — drop_capabilities / 202110133112 / 4

Type: `["list", "string"]`. Computed.

List of capabilities that Docker container has.

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

<a id="canonical-1320313011010132-2203313122031320-0000232001311110-1101311303310333-1302312313212332-1123010322223102-0003100133333000-1031002123221031"></a>

## Next pages — drop_capabilities / 202110133112 / 5

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-1101333110023110-2311233020333110-0032311033310122-3103011323131313-1022123333302101-0200100132232213-1033101102031102-0022303221320231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233201302101330-1300120333231000-3232011333012100-1022323101212121-1330330322010231-3321201322002301-1333102230313301-1103331001021233"></a>

## psp_spec.fs_group_strategy_options — fs_group_strategy_options / 033101020130 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.fs_group_strategy_options

<a id="canonical-0123321301101333-3120122010031211-0113222211310123-3002213323122013-0210110130312210-0212101021310013-2211111132133222-3121210122102313"></a>

Type: `"single"`. Computed.

Configuration parameter for fs group strategy options.

Upstream description:

ID ranges and rules.

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

<a id="canonical-1002202102220120-2210331222021033-0003122133100131-2301312202000001-3132302232133021-1013321220312131-1312031011220031-1011012011213331"></a>

## Direct properties — fs_group_strategy_options / 033101020130 / 3

- [id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0021220102113331-0103212121213202-3123302032011133-2321210301121321-1303222122330023-2311023002110012-0102100320211033-0303113232200302): complete subsection reference.

<a id="canonical-3130222112122122-1310223112333231-0330222201301332-1103003302303212-0330130230230310-1013232023100211-0211010133221132-2311322132130222"></a>

<a id="canonical-3213321233033320-1201020332031130-0031120112102012-3303322221023131-2021003000013112-1012331110223123-3211001122012012-0203212121113023"></a>

## rule property — fs_group_strategy_options / 033101020130 / 4

Type: `"string"`. Computed.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3211001033322323-0103230000230031-1212130112123100-0332333222021212-3212312222323030-3033030032320003-2020020312301211-1110303122121131"></a>

## Next pages — fs_group_strategy_options / 033101020130 / 5

- [psp_spec.fs_group_strategy_options.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0021220102113331-0103212121213202-3123302032011133-2321210301121321-1303222122330023-2311023002110012-0102100320211033-0303113232200302)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-0021220102113331-0103212121213202-3123302032011133-2321210301121321-1303222122330023-2311023002110012-0102100320211033-0303113232200302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330131332313200-2031303303333230-2120130212033031-1120033133112231-2122132210111100-2332032230200100-3030320233001023-1310102312011022"></a>

## psp_spec.fs_group_strategy_options.id_ranges — id_ranges / 302131011021 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [psp_spec.fs_group_strategy_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1101333110023110-2311233020333110-0032311033310122-3103011323131313-1022123333302101-0200100132232213-1033101102031102-0022303221320231)
- psp_spec.fs_group_strategy_options.id_ranges

<a id="canonical-1323120111332320-0330113133100100-2031202111003201-1300303130311230-2100131203131110-1223103112120003-1301203130110200-0333210330300320"></a>

Type: `"list"`. Computed.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1010302120023000-1331111222012133-2032001332112310-1333123233122303-3020111303033211-1200311120001011-2102212011233202-3203321031031012"></a>

## Direct properties — id_ranges / 302131011021 / 3

<a id="canonical-1211222310022320-2120322020002102-0332112021333331-1221023210112100-2311331232113120-3312001323220130-0121012011122022-3122313100301211"></a>

<a id="canonical-1320031311033311-2022312301232130-2120020100230002-1012103221202003-0232012202320131-3213111212322121-1020321313313121-2002303213331002"></a>

## max_id property — id_ranges / 302131011021 / 4

Type: `"number"`. Computed.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2203121231300003-3301330003232011-0023200110030323-3230021130032311-3000012332133321-1303220311220333-0100322331313103-0133110131023321"></a>

<a id="canonical-1201011113112313-1110113310011013-1320112221230232-3031330333023302-1221121130132103-2230131020300001-3102320302023202-0121200113002113"></a>

## min_id property — id_ranges / 302131011021 / 5

Type: `"number"`. Computed.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3032303220100301-3011331001012331-3213132201010323-0303310233101313-0301010111232233-0120203010213012-3020220320313211-2120333332112321"></a>

## Next pages — id_ranges / 302131011021 / 6

- [psp_spec.fs_group_strategy_options](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1101333110023110-2311233020333110-0032311033310122-3103011323131313-1022123333302101-0200100132232213-1033101102031102-0022303221320231)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-2010013123031310-3211213220010312-0310000023233332-3100323323010221-0332101332122222-3220102332203132-1223121003033110-3311320222302103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201232311323112-2321111233203011-2003222312332320-2023330323022103-3000202012011231-1213201032301131-1203310301101032-1311202020211333"></a>

## psp_spec.no_allowed_capabilities — no_allowed_capabilities / 002123320211 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.no_allowed_capabilities

<a id="canonical-0102331032131022-3201332322300120-0001113133223033-3313333323122301-3022331100310021-1012130131033101-0103032332110101-2023222131132033"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1300033312330320-0010123203122222-3202012202310010-0123222022032302-0313312210203133-2120012300000233-3333120203300012-1323031220311123"></a>

## Direct properties — no_allowed_capabilities / 002123320211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133020012023012-3020013103300132-2300113122013310-1010031332112113-1012213211303031-0311331203113020-0231010210320032-0113313133303302"></a>

## Next pages — no_allowed_capabilities / 002123320211 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-3013311031321132-3311012331112203-2121021111012313-0321121303333132-2323103313331233-1221132113121010-3001122111233000-2323310220100200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200213022232300-1031211033112331-0230230121231131-3123201333321221-1033221100220010-0323123230300222-1213100030213320-3310212013133113"></a>

## psp_spec.no_default_capabilities — no_default_capabilities / 000313201002 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.no_default_capabilities

<a id="canonical-2330101030022112-2311010233032230-2120033102012332-1031130220031310-1102033211010333-0302112331103302-0032310130110003-3320010131131331"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1122001120121303-2321021033030002-3310330130220230-2003020322033320-3231110003321103-0020332322110212-0132120103301220-2133020302211113"></a>

## Direct properties — no_default_capabilities / 000313201002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233330010300333-0112120200311223-2021022320020100-2223022110210111-3221021210023013-0222021213230330-2032212102022110-0202101113030033"></a>

## Next pages — no_default_capabilities / 000313201002 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-0022102100203333-0233121320103112-3223201023013302-2123102300020312-2323102030020203-1100200333332203-1121003121233010-0220120101203303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300011301202113-2210022100321203-2113320311131102-1300200311303313-1011232322213020-0013332100310111-3013000032331200-1132123012200112"></a>

## psp_spec.no_drop_capabilities — no_drop_capabilities / 303122210310 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.no_drop_capabilities

<a id="canonical-1103001030213130-2032230231301012-1330312130132213-3022203330010112-2111322011230002-2221323130030111-1000233322011101-1123211323223023"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1200100311201221-0212120223220133-3201011133121321-2020212320103213-2330312033102112-3000323133111003-2310023122220330-3031212301222130"></a>

## Direct properties — no_drop_capabilities / 303122210310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223022322130303-2001233133230032-3223000331032222-1103001122331033-3121220230322131-2303000031032112-3331000123111203-2210032313112310"></a>

## Next pages — no_drop_capabilities / 303122210310 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-2120121233213121-2213211103022200-2021230003320022-0330301021323223-2103031103113323-2030300030033120-2112210212121112-2022212130123323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201000323313323-3012023303211112-1313012013103321-0131131231103210-1112202203301212-0330221100013120-1011302122231103-0012232131320320"></a>

## psp_spec.no_fs_groups — no_fs_groups / 220002230201 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.no_fs_groups

<a id="canonical-2001033320020223-0113110031110333-0030013133030223-3032000100203121-0021331033222322-3232100222333303-3011211102203323-2223003223310210"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3223320130233331-2003223321032222-1021210300010320-1013332132201130-2321113210222321-2011321223230322-0003001203302013-3330323023210123"></a>

## Direct properties — no_fs_groups / 220002230201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023120113011032-3022200121231221-3113301313032233-2232211330032300-2133022130310201-0100122323103103-3202102123133333-3121212100001202"></a>

## Next pages — no_fs_groups / 220002230201 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-1001302122030322-1233303133231222-2311133230020122-1213211121231122-2321302013031030-0022310003231232-1313223232003111-2220103033103123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300312311110033-3202113022011320-0213221230132200-3112332102133212-0131110002110331-2020303303202222-2130213000013311-3131111002230023"></a>

## psp_spec.no_run_as_group — no_run_as_group / 202302023233 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.no_run_as_group

<a id="canonical-2000010000103221-1021310123202103-2122121122021123-0022211122200020-2320002103232201-0311203031303002-1130001313333223-3032211131022220"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0001033321201130-1122303020322123-3023331313010202-0133203010303133-0133332330300013-1112132230232101-2030203120303022-2233120220001002"></a>

## Direct properties — no_run_as_group / 202302023233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030223113012132-3302313332032033-3233122233120331-2021110212321102-3122302003333230-2331010030322313-1221322211131213-0120330330302202"></a>

## Next pages — no_run_as_group / 202302023233 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-0231310033201230-0230120223002112-0321312331201101-2313111012202332-0113313123013320-2233011213310021-0020203330133032-3100210130120311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222223213101323-3212312010102203-3303021102222010-0223233012101211-1333022133211032-0113213210220110-3122030000132330-2223231030000032"></a>

## psp_spec.no_run_as_user — no_run_as_user / 001131201031 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.no_run_as_user

<a id="canonical-1231203103200000-1223022000203013-2100110030201132-2320103332303011-1313211221231302-3112302112303030-1101312302030101-2010121102313023"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3020012011122121-1010033320333330-2213003332103122-1310021030022231-3130203000323300-0032322220312103-2303021030002231-3201002331312012"></a>

## Direct properties — no_run_as_user / 001131201031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222131101301123-3231100031310203-3233120232000220-1301033320111213-0023303221222001-0002002232132312-0101121212223033-1301213311310201"></a>

## Next pages — no_run_as_user / 001131201031 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-2320003122022032-1200110113102333-2030212232102323-0103303130233201-1302320313211031-0221323112321330-0330303133130102-1102022012300101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303220211301133-0033233002201221-2120300333020211-3011213200221120-0203033112133100-0033111031122210-0300213213320111-2203101311003313"></a>

## psp_spec.no_runtime_class — no_runtime_class / 220310110211 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.no_runtime_class

<a id="canonical-3202101213303123-0000321002331331-3011000131333112-0032331330013213-0211033212121201-0012020221113203-0311202201220313-2320022333003302"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0310222213013113-3230212003233002-1100333112003101-2211032200300302-1322213211010101-2232113030002300-0300111033122201-3011222010303130"></a>

## Direct properties — no_runtime_class / 220310110211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010213310333102-3320112120003101-3311320003001213-0230110222010023-0112300102112022-3020320033110011-1331323132102010-3310112300220120"></a>

## Next pages — no_runtime_class / 220310110211 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-3033121311022123-1233232113312211-1000103330023133-3200031030302220-1121111322213232-0303122111223313-2110122133332120-0100001023031123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213110000201232-1013101130223033-1321010032310320-1123013213301131-3211121221322331-0023111130020201-0303210032120202-2220121122233131"></a>

## psp_spec.no_se_linux_options — no_se_linux_options / 303113022013 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.no_se_linux_options

<a id="canonical-0330132130213201-3121120100130222-1132313213202022-2330232313020221-1311002112110031-1121123312200122-1023003223113223-2323013210010232"></a>

Type: `"single"`. Computed.

Configuration parameter for no se Linux options.

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

<a id="canonical-1021202303123220-1020100130213221-2221122222212112-3223201201100102-3003212000232311-1032321301120332-0021121301302012-0311001111312322"></a>

## Direct properties — no_se_linux_options / 303113022013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022130320313130-2032113331302020-3013200303113210-2130222122311222-2000031222020013-2032323013013212-0231031232220300-3321312133210313"></a>

## Next pages — no_se_linux_options / 303113022013 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-1000112100012022-3030310223032102-2331000103320032-0031112102321123-2100210202111230-0313033312101103-2003313131212103-2321220133001330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210303122233230-2121111312332011-1231200133131103-2331020221111330-3303003303012013-0333012123313231-0212112022301202-0323202312113302"></a>

## psp_spec.no_supplemental_groups — no_supplemental_groups / 312020332010 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.no_supplemental_groups

<a id="canonical-2103002300000201-2331231310331100-0300230202132102-1202230123122313-1110022100231123-1113002100320333-1201102030003301-1020212133322303"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3023000120002303-1131211323100031-0310302221212333-3122102012121302-2320303222201221-1021020032123312-0100301010320321-1303003013121002"></a>

## Direct properties — no_supplemental_groups / 312020332010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231213221220022-0332133121202131-2322011310231303-0210212021300001-1212132131213313-3100030231110210-1011320122101320-1203101231020000"></a>

## Next pages — no_supplemental_groups / 312020332010 / 4

- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-1000330010132213-0013201221322322-3310132230301122-0231103202203321-3203000002330032-0023030100013121-3203212122101103-2331010222302333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222030131312313-3300120233322031-2330301002331333-1033203021031000-2200323220121333-1012023222023110-1312330322020330-2230303100301012"></a>

## psp_spec.run_as_group — run_as_group / 300230222030 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.run_as_group

<a id="canonical-1003002133000312-3222113302132232-1303201313133120-3012233300233322-2010123302121131-1320123213203013-2232003233330220-0030203000323011"></a>

Type: `"single"`. Computed.

Configuration parameter for run as group.

Upstream description:

ID ranges and rules.

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

<a id="canonical-1310101021302103-0213323202330203-3310012212230200-1201122010112300-3000211000301212-0031303312100322-0211003133012223-0101202330001002"></a>

## Direct properties — run_as_group / 300230222030 / 3

- [id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1202031332233301-0013323213221003-2020011001203300-3012112132323222-3032300132310001-3313123303312322-2211331131120231-3323032001132211): complete subsection reference.

<a id="canonical-1030133133123220-3331213330131310-1030202303222220-0000020220111231-2123230232130030-2023010232322032-1102130301331111-3113033011131300"></a>

<a id="canonical-2120202201010322-2123111100100112-1303023213333221-3120103213321301-3230023210102001-3202010021300122-0311203130333120-2233012220131202"></a>

## rule property — run_as_group / 300230222030 / 4

Type: `"string"`. Computed.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1102202012301333-0332101111133233-2101023213322300-1210120113210231-3120031020022222-0310302321213313-3103211201211212-1332330130231033"></a>

## Next pages — run_as_group / 300230222030 / 5

- [psp_spec.run_as_group.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1202031332233301-0013323213221003-2020011001203300-3012112132323222-3032300132310001-3313123303312322-2211331131120231-3323032001132211)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-1202031332233301-0013323213221003-2020011001203300-3012112132323222-3032300132310001-3313123303312322-2211331131120231-3323032001132211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000013012311012-2022000100033130-3333110223330011-2202023321323300-0020011202031002-1023110121201031-0101030320000211-2110203121021020"></a>

## psp_spec.run_as_group.id_ranges — id_ranges / 201203233213 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [psp_spec.run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1000330010132213-0013201221322322-3310132230301122-0231103202203321-3203000002330032-0023030100013121-3203212122101103-2331010222302333)
- psp_spec.run_as_group.id_ranges

<a id="canonical-0002233013030230-2111230012312010-3013131311303023-1123213001001010-0201322123310030-0003021310312102-0310210022000023-2200210232100022"></a>

Type: `"list"`. Computed.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2012032003002322-2302220203302112-1132101002333310-3330023020100101-0322023200212120-0100220302220032-0011031220101100-2012130012003301"></a>

## Direct properties — id_ranges / 201203233213 / 3

<a id="canonical-3013312301001131-1000120333000302-2301131020223302-1131101202020300-1313133310201022-1230230203023303-0220223331021032-3330310111203222"></a>

<a id="canonical-0303011310021211-2302330010001300-0220210212012122-1333122330323101-3010210231201222-1103210023111013-2111122202102230-2021010121300123"></a>

## max_id property — id_ranges / 201203233213 / 4

Type: `"number"`. Computed.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1212022223130010-3001330322113023-2320332013103203-1221321002323212-3123322020003231-1323303022020031-0120230111003120-1330111321000302"></a>

<a id="canonical-1001100323210221-1033000320221202-2100023022230200-3221302212123212-3313321312101213-0100222303103003-0302302230102131-2022200232221020"></a>

## min_id property — id_ranges / 201203233213 / 5

Type: `"number"`. Computed.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0001133123130011-2012100321222222-1131102212020313-1322233211301120-0011002322232211-3003200120112011-1102023302333020-3333213200222112"></a>

## Next pages — id_ranges / 201203233213 / 6

- [psp_spec.run_as_group](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-1000330010132213-0013201221322322-3310132230301122-0231103202203321-3203000002330032-0023030100013121-3203212122101103-2331010222302333)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-2211301322001300-1211102230212212-3222112002333031-0321030003031212-2312222233032220-2020331210031013-2231123302321131-2103033201020311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220222000100001-0123222223310011-3200333221202211-0302333331220131-3311003022213130-0132320312023202-1301221122213221-0233012010120103"></a>

## psp_spec.run_as_user — run_as_user / 001013201013 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.run_as_user

<a id="canonical-0121123133323102-0310312122222221-1103211102103231-2202201020121120-0113111312202023-1110331021212023-2202013022231203-2212102101321111"></a>

Type: `"single"`. Computed.

Configuration parameter for run as user.

Upstream description:

ID ranges and rules.

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

<a id="canonical-3232332031312201-0201331300300032-2202023202311031-3333100331310031-2301313121232330-3031111222220113-3210333310320203-1311302123311321"></a>

## Direct properties — run_as_user / 001013201013 / 3

- [id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3202000213223201-2101132131203101-0011330011003113-0001121103011211-0021022132131100-1203300131121230-1233013322010010-1012030021102103): complete subsection reference.

<a id="canonical-3110302321122132-1113001210012033-0120003121033130-0033311223212322-1121300120330300-2301300122223102-2312133212103310-0010030332231031"></a>

<a id="canonical-2113222323020311-0311323233011320-0201313230100201-0113000310323220-1032223110211101-1103311311101011-0331323212130021-1230132002133212"></a>

## rule property — run_as_user / 001013201013 / 4

Type: `"string"`. Computed.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3211103313130032-2132130121333032-0001333122231322-1112222210322103-1023020032133303-1321102233132330-2231233213211112-0330231121201333"></a>

## Next pages — run_as_user / 001013201013 / 5

- [psp_spec.run_as_user.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3202000213223201-2101132131203101-0011330011003113-0001121103011211-0021022132131100-1203300131121230-1233013322010010-1012030021102103)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-3202000213223201-2101132131203101-0011330011003113-0001121103011211-0021022132131100-1203300131121230-1233013322010010-1012030021102103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122301120111000-1013332232112202-1030000231330310-0003033323113203-1000010030212233-0201131200033033-2231232012300233-3213212301333032"></a>

## psp_spec.run_as_user.id_ranges — id_ranges / 122011012031 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [psp_spec.run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2211301322001300-1211102230212212-3222112002333031-0321030003031212-2312222233032220-2020331210031013-2231123302321131-2103033201020311)
- psp_spec.run_as_user.id_ranges

<a id="canonical-2213101120321200-2231103130300300-1100112023000000-3200002013023000-2211202022120133-2010210022322000-1221002202212103-1310221130212212"></a>

Type: `"list"`. Computed.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1330222212212200-1210211322022223-2211100130211230-3312302302321030-1211202201312230-1222203300031301-2123032000003010-1233313113232321"></a>

## Direct properties — id_ranges / 122011012031 / 3

<a id="canonical-2203133321111030-3110323123131321-1202310103232200-1232121000032203-1130011110200000-2213302322302103-1001233222110103-3003201112312111"></a>

<a id="canonical-1102332201211031-1300213020121331-1221232020032001-2010000333130001-3212020202100301-2023322003333221-1321000312330013-0020333322211130"></a>

## max_id property — id_ranges / 122011012031 / 4

Type: `"number"`. Computed.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0331311311223031-3312033121030031-3200313133310311-3231321312103213-1001231320113310-3310312012330220-2213120003023232-2013211212011331"></a>

<a id="canonical-2210010220112032-1133300030023330-2031120310101032-1201002310011331-3301022133330121-1330301323010011-1130212333012212-2213021110213131"></a>

## min_id property — id_ranges / 122011012031 / 5

Type: `"number"`. Computed.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1021222003021101-0300330003330133-3131112012200121-3021213023010032-3222022001222021-2130312011033012-0202031201320301-0230320010230021"></a>

## Next pages — id_ranges / 122011012031 / 6

- [psp_spec.run_as_user](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2211301322001300-1211102230212212-3222112002333031-0321030003031212-2312222233032220-2020331210031013-2231123302321131-2103033201020311)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-0312231101331321-2213322320201323-3232313110332322-1030122210300331-0322302003312220-2311303203230011-1023102100212132-1232002030111120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213222231232200-3000303330020003-2323210230101103-0133111122101231-2213112331333003-0220001131002212-0000022210233033-2023321002110112"></a>

## psp_spec.supplemental_groups — supplemental_groups / 222101011212 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- psp_spec.supplemental_groups

<a id="canonical-3003312122032133-3202130210011120-0210231032003223-3131223311333132-1323121021333122-2312212111310112-2003020212333022-1212100203213131"></a>

Type: `"single"`. Computed.

ID(User,Group,FSGroup) Strategy. ID ranges and rules.

Upstream description:

ID ranges and rules.

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

<a id="canonical-1130112212223300-3133020120021313-2201023113232100-2003012210003203-0230233121011133-2201103130012311-0123023230022001-0121103210110323"></a>

## Direct properties — supplemental_groups / 222101011212 / 3

- [id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3022130322331311-0030123221213010-3003100132303010-2321033012030202-0122222323213221-1300130300013033-2212220331301110-2330112301333321): complete subsection reference.

<a id="canonical-1130032303323213-2133320113301132-2320230213012121-1332100310100233-2103131101313010-2212003321011020-1020123320001010-0321233010133310"></a>

<a id="canonical-3121032232332013-0232030110020333-1212001103101222-0220111212302303-0330130313202023-1102230323323010-3101233211211323-0002013212101002"></a>

## rule property — supplemental_groups / 222101011212 / 4

Type: `"string"`. Computed.

Rule indicated how the FS group ID range is used.

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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0012013231032332-3113223301203232-1210223112300300-0210122030311310-2020130222010121-0322310123021223-1001320111202221-0023020211211321"></a>

## Next pages — supplemental_groups / 222101011212 / 5

- [psp_spec.supplemental_groups.id_ranges](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-3022130322331311-0030123221213010-3003100132303010-2321033012030202-0122222323213221-1300130300013033-2212220331301110-2330112301333321)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)

<a id="canonical-3022130322331311-0030123221213010-3003100132303010-2321033012030202-0122222323213221-1300130300013033-2212220331301110-2330112301333321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211120001203022-0103333022310210-0312200202030202-0232221321023012-3223121012231022-1110113313001333-3110212230231312-1021333001120221"></a>

## psp_spec.supplemental_groups.id_ranges — id_ranges / 210111030130 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Property reference](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [psp_spec](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033)
- [psp_spec.supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0312231101331321-2213322320201323-3232313110332322-1030122210300331-0322302003312220-2311303203230011-1023102100212132-1232002030111120)
- psp_spec.supplemental_groups.id_ranges

<a id="canonical-2001010103033331-2023022030133323-1113211013113212-2210320200110211-1222013323022030-1310030333200202-2231201220121313-3001211030123111"></a>

Type: `"list"`. Computed.

ID Ranges. List of range of ID(s)

Upstream description:

List of range of ID(s)

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1223012331220120-3130011231321022-1302223023001031-1010112010022333-2111301131310300-0223320022230333-1132232331003010-3020223101020033"></a>

## Direct properties — id_ranges / 210111030130 / 3

<a id="canonical-0010300312313320-1123100311322020-0212210113113233-3311112023223302-2330202101122301-2111332130110111-0032120220210331-2231302320320310"></a>

<a id="canonical-0223332211030020-0321032003133232-0210331310031212-1130120122010001-2333330213331310-0303133012232131-3010030130113022-0101130100100212"></a>

## max_id property — id_ranges / 210111030130 / 4

Type: `"number"`. Computed.

Ending ID. Ending(maximum) ID for for ID range.

Upstream description:

Ending(maximum) ID for for ID range.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0012232323310302-0200332332331311-0322313303112321-3001313233133033-2231030223212330-1301131032023232-1330330002110302-0300233100333322"></a>

<a id="canonical-1103011002100102-0113133201023302-2300333032201112-3013201300103113-1302213101123112-3113130233321012-3210203103123301-2221012131122032"></a>

## min_id property — id_ranges / 210111030130 / 5

Type: `"number"`. Computed.

Starting ID. Starting(minimum) ID for for ID range.

Upstream description:

Starting(minimum) ID for for ID range.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1102102030312013-1300100113201122-0323331320101212-3313332301130011-0002332010102230-2220202221003113-0310220303332320-0330330003020232"></a>

## Next pages — id_ranges / 210111030130 / 6

- [psp_spec.supplemental_groups](data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0312231101331321-2213322320201323-3232313110332322-1030122210300331-0322302003312220-2311303203230011-1023102100212132-1232002030111120)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
