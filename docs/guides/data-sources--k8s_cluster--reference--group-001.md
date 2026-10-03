---
page_title: "xcsh_k8s_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster reference."
---

# xcsh_k8s_cluster reference

<a id="canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202011233001223-0231212103223330-1010032101212211-1032031113113203-0211302101320123-3232201103201013-3232130322320320-3330121233302232"></a>

## Property reference — Property reference / 131231220211 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- Property reference

<a id="canonical-1020030230321301-1003321212022221-2322103002311012-1332030333223000-3132003121331330-3332132313233321-1100332003133332-0332330332112322"></a>

## Direct properties — Property reference / 131231220211 / 3

<a id="canonical-2030303213222123-0312022131222011-2221230320013033-0231032002122012-2021321023120103-1201330300200011-1203331023300331-0230112213322030"></a>

<a id="canonical-1333222223223020-2201021011023301-0021033200120331-3033333001112003-0003333311103323-1120132022022330-1301312220320012-1113000320010310"></a>

## annotations property — Property reference / 131231220211 / 4

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

- [cluster_scoped_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-0120002331303020-2010212021110212-1002000103132231-3312311312112020-3222320120101120-3333021221032011-0010003213111232-1301303220301200): complete subsection reference.

- [cluster_scoped_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-1001213022130021-2301002303310100-0123321311121203-3133231111100330-2021121330000222-3200331033311311-3211110012111000-3233010132332202): complete subsection reference.

- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132): complete subsection reference.

<a id="canonical-0122033213110023-1003102330010021-3231112033131121-3110331211122220-1312113023133130-0311321120203330-2310020022220010-0112311233112023"></a>

<a id="canonical-2031112010210313-2031301031333030-2323331222330001-2011120000010131-1133133310121213-1123003312310031-1133000222223030-3211301002110132"></a>

## description property — Property reference / 131231220211 / 5

Type: `"string"`. Computed.

Description of the K8SCluster.

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

- [global_access_enable](data-sources--k8s_cluster--reference--group-001.md#canonical-3320222313103210-3123212313210032-0133030311211222-1010012023223121-3323132221202210-0121033131131221-1030202213312231-2212103112002220): complete subsection reference.

<a id="canonical-0202201133303221-0200100013313111-3132131133023011-1130233213111333-0110220303311333-1023303033203003-3020321332103330-1200210103110311"></a>

<a id="canonical-3233232000223112-3010300122223130-0132123223303121-0000030202300111-3031023022201013-2300001121230000-3302323101030233-2311221301032310"></a>

## ID property — Property reference / 131231220211 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [insecure_registry_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2000030331132132-2030211331113330-3220120113033231-2122210221011231-0213200233103132-2313012130101320-3322020233012010-0101023233133021): complete subsection reference.

<a id="canonical-0020132111001220-3201110200010100-1021102110301200-1032213031311122-1222221330022020-1030021033220120-2010032022311230-2112102312100003"></a>

<a id="canonical-1012111332000313-0321020331222032-1020302022110033-2303132212023032-3031030122003103-3333312212210022-1100312110213011-2130303103332222"></a>

## labels property — Property reference / 131231220211 / 7

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

- [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-1232130303333302-2123302020123001-1012302120120100-2031110103100133-1333002211210233-0210013221233333-2013211033200130-2122000103121331): complete subsection reference.

<a id="canonical-3200322033002013-1311103300031132-0131210100023310-2033203031131223-1013313300300300-2331030331313220-0030120232232130-2102332103133013"></a>

<a id="canonical-0011233330233313-1000101020122313-3333210121100332-0132201000003201-3232022010212232-2120222100122221-3302110312302321-3321330120300013"></a>

## name property — Property reference / 131231220211 / 8

Type: `"string"`. Required.

Name of the K8SCluster.

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

<a id="canonical-3001033013232323-3013301023320030-0212323020232121-2103313222101323-0103121231222332-2211331123100032-2212200220130231-3033301131010231"></a>

<a id="canonical-2232322300320101-0320102022212200-1211202020200213-0203030001232333-3123013212321002-3331310000221332-2322212320120023-0233003230302320"></a>

## namespace property — Property reference / 131231220211 / 9

Type: `"string"`. Optional, Computed.

Namespace where the K8SCluster exists.

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

- [no_cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-3030121200123230-2331212032333120-1222020201113023-2223333222111313-1111223011003323-2212200320132301-1013302010323310-1101133032011100): complete subsection reference.

- [no_global_access](data-sources--k8s_cluster--reference--group-001.md#canonical-0200200221310301-2230002322030021-1312111002133223-1002311321102011-2021031221111102-1103311003300332-3002323333213332-0332212203031002): complete subsection reference.

- [no_insecure_registries](data-sources--k8s_cluster--reference--group-001.md#canonical-0330033302000032-3231120203000223-3031000032022331-3013223231322231-0000021312123130-1030111211103110-3222212331102001-2032232111221130): complete subsection reference.

- [no_local_access](data-sources--k8s_cluster--reference--group-001.md#canonical-0302113023030331-3313211211313300-1121010212112131-0313100112232110-1210002211320123-1030321312223201-2032113121233310-2012223221032303): complete subsection reference.

- [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-3321211133132301-2331122022223321-2012231113301103-1020230201103020-0302233321331011-2321020130101031-2012122130121322-3131100020033332): complete subsection reference.

- [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2210312112133200-2120022110103011-3332312223021001-2132020112232033-3130120230202301-1002132302112011-3330122112333332-2312202300223020): complete subsection reference.

- [use_custom_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-3012230023323023-1200010011232313-2031133231312300-1101121323331103-0310202311002101-0323110101101321-1000130001232133-0102210223300320): complete subsection reference.

- [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-0222103111213102-1022030213200111-2301201033101101-2013300120212213-2320221103200012-2233132210011313-3020211000121131-1213103021100330): complete subsection reference.

- [use_default_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-3033213011211231-1311213031231322-0331302021011032-0000211222303012-1322100223221232-2302122333012033-3131332132222230-0030211100120123): complete subsection reference.

- [use_default_cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-1030120320023032-0121111102120012-2122221010210131-3023032202310122-3302021200233321-2123032002300221-3130110013130033-1212033311130030): complete subsection reference.

- [use_default_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-1101101122322220-1120213302132230-3212131032131321-0033133201001221-1011222030321331-3121013100102132-2131330220110032-3001313332022220): complete subsection reference.

- [use_default_psp](data-sources--k8s_cluster--reference--group-001.md#canonical-0331232133213200-1233022031333312-0230013230013120-1233232013212110-2030212322112002-2103210312230132-1012023100100202-0123100212303233): complete subsection reference.

- [vk8s_namespace_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-0002333010132133-0022011023203112-3231122111102303-0131030011100223-3233022321012030-3102031213223333-1323301112022032-3031200010100021): complete subsection reference.

- [vk8s_namespace_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-0030031301002303-0032103230211033-3333313322321331-3233011120002200-2020222121001033-3120301113302031-0013121112201100-2131301013001320): complete subsection reference.

<a id="canonical-3233021313023001-0332022213323232-1323012203220302-0121213131231313-1221020013122300-0323300221122012-2003010331103330-2100310012133330"></a>

## All schema paths — Property reference / 131231220211 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_cluster--reference--group-001.md#canonical-2030303213222123-0312022131222011-2221230320013033-0231032002122012-2021321023120103-1201330300200011-1203331023300331-0230112213322030) |
| `cluster_scoped_access_deny` | [cluster_scoped_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-1223010310201321-0021323111131231-1313113313002332-0230030320213032-3000102111133112-1101132103212333-1102310120331320-3201222203123122) |
| `cluster_scoped_access_permit` | [cluster_scoped_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-3320332310023201-3312021121102213-3230122120232112-2232313322003123-3321133322003112-3013232202013323-0311200200212130-2101033023323221) |
| `cluster_wide_app_list` | [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-0210223233021201-3011030020003221-1332211103002212-1013020101103030-3000113113001213-2032203000032203-0310312102131020-0111121303021311) |
| `cluster_wide_app_list.cluster_wide_apps` | [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-2320002113020121-0210113323132302-3231332100023110-2121003001120131-3200230323003132-3320013000303103-2101033230032003-2201003333221203) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd` | [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-1030031302211133-2121223320020233-3031120231003123-2203002212101012-3133110113110130-1001233101103230-1032323310222133-2212010133033232) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-0221132301002231-1132222303001001-2102330312103313-3303131203111333-1331112011113120-1122313103311213-1002010333131322-3003021221203200) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-2210113033323010-1123321133210020-3111302203113121-1100303331203201-2200313100022332-2221333123311210-3032331322202233-3111233131312033) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.local_domain` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-3202300211310002-3303221312133110-1321200122101331-1103333232103211-1123112002031121-2231101332010012-1121002012323111-0023211021132221) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-1030123311021300-1101231011333301-1000313021101123-0323122301320103-2220332203033311-1222023132101303-0323123203103032-3302233021211223) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-2333322301332203-2300023032132301-1212221012023112-2302011200021120-2320030121103133-0122233230033012-1233123203033122-0231320122133020) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.decryption_provider` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.decryption_provider](data-sources--k8s_cluster--reference--group-001.md#canonical-2201133102210001-1321103231200021-2301110202023012-3212031121300130-0013111331110131-1222210032333020-2122010211130132-1020233313303121) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.location` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.location](data-sources--k8s_cluster--reference--group-001.md#canonical-2320311011133321-1210113132111210-1100001020200032-1201313201210333-0030313113000202-0313230213221203-0301230222010033-1201022201020320) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.store_provider` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.store_provider](data-sources--k8s_cluster--reference--group-001.md#canonical-0113301120032303-3131212221232000-3031311220131123-0321221301113011-1111203320111303-0202230313003231-1122321013201101-2311022211022320) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-0213213011101002-2311102013000321-2331223213010320-2332001330022333-0010133123031233-0300130212311211-2102132033232010-1011121130331302) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.provider_ref` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.provider_ref](data-sources--k8s_cluster--reference--group-001.md#canonical-2312203211030321-2222230223012113-3112323020330020-3322200312210000-3210010110311011-2013302310010230-0332030311330222-3211220110230121) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.url` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.url](data-sources--k8s_cluster--reference--group-001.md#canonical-3303332231101113-2200310323000031-2221210210101122-3021211301211020-2320113030210230-2302232000112020-3013111311003030-2301033000313223) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.port` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.port](data-sources--k8s_cluster--reference--group-001.md#canonical-3122031332313003-2213233121200113-0131220101220030-3220133113202320-2022101031123012-0223011101101213-2302003233123312-0032130131122332) |
| `cluster_wide_app_list.cluster_wide_apps.dashboard` | [cluster_wide_app_list.cluster_wide_apps.dashboard](data-sources--k8s_cluster--reference--group-001.md#canonical-0302003332300321-2302233302000312-1003020112022322-3003031332300131-3133203132033002-3031002133122333-2020323310210120-2102101222022202) |
| `cluster_wide_app_list.cluster_wide_apps.metrics_server` | [cluster_wide_app_list.cluster_wide_apps.metrics_server](data-sources--k8s_cluster--reference--group-001.md#canonical-2302122132223022-0203212022131113-1012112003132030-3123013233120123-1220300232023312-3212331330032001-2012023121322033-1220020301113032) |
| `cluster_wide_app_list.cluster_wide_apps.prometheus` | [cluster_wide_app_list.cluster_wide_apps.prometheus](data-sources--k8s_cluster--reference--group-001.md#canonical-2321122102032213-0123230123323011-0332012023023132-2310123212100210-0301323102033012-3102003302030331-0232031113003022-0312223031102000) |
| `description` | [description](data-sources--k8s_cluster--reference--group-001.md#canonical-0122033213110023-1003102330010021-3231112033131121-3110331211122220-1312113023133130-0311321120203330-2310020022220010-0112311233112023) |
| `global_access_enable` | [global_access_enable](data-sources--k8s_cluster--reference--group-001.md#canonical-3023332120102213-3112021003210331-3113131231200301-1112120212001221-2211003000120203-2330233010302133-3121323323332122-2301012200120023) |
| `id` | [ID](data-sources--k8s_cluster--reference--group-001.md#canonical-0202201133303221-0200100013313111-3132131133023011-1130233213111333-0110220303311333-1023303033203003-3020321332103330-1200210103110311) |
| `insecure_registry_list` | [insecure_registry_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3033313331133230-1212122030330001-0033112013020112-1332331222101110-0013100013333132-0130210132212132-1302321302203102-0331202111123212) |
| `insecure_registry_list.insecure_registries` | [insecure_registry_list.insecure_registries](data-sources--k8s_cluster--reference--group-001.md#canonical-2323031033113222-0310201133313030-0212313312011122-3233101021112203-2303223221000202-3331200123221031-3303131300302233-2110311000100021) |
| `labels` | [labels](data-sources--k8s_cluster--reference--group-001.md#canonical-0020132111001220-3201110200010100-1021102110301200-1032213031311122-1222221330022020-1030021033220120-2010032022311230-2112102312100003) |
| `local_access_config` | [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-2122323233102321-3222222131003200-3012303112121122-1332101110200132-3230201133201033-2022232123013110-0033012231222210-2213302102032333) |
| `local_access_config.default_port` | [local_access_config.default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-0311330110113031-2033130021301223-3212311323130200-0122313302123323-1011123001020031-0002022002200100-1101001211321230-0330232200000303) |
| `local_access_config.local_domain` | [local_access_config.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-1031100012223030-1132002323231031-0010112333123220-1223010103012122-3000012313003013-0223003101211131-1013100200101032-2232321030310222) |
| `local_access_config.port` | [local_access_config.port](data-sources--k8s_cluster--reference--group-001.md#canonical-1023233113132023-3302231201212112-0122013113210132-2231331203101233-3302321231202203-1102100122320322-1120013310331332-3222211202031000) |
| `name` | [name](data-sources--k8s_cluster--reference--group-001.md#canonical-3200322033002013-1311103300031132-0131210100023310-2033203031131223-1013313300300300-2331030331313220-0030120232232130-2102332103133013) |
| `namespace` | [namespace](data-sources--k8s_cluster--reference--group-001.md#canonical-3001033013232323-3013301023320030-0212323020232121-2103313222101323-0103121231222332-2211331123100032-2212200220130231-3033301131010231) |
| `no_cluster_wide_apps` | [no_cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-2121002130233322-2131300330321301-0223201213012213-3213022102132001-0220213303001220-3100032322320313-0232020312010330-1020131130220203) |
| `no_global_access` | [no_global_access](data-sources--k8s_cluster--reference--group-001.md#canonical-3230320232203232-1223300012302210-2003112003211112-0113131023100013-2212101202031130-2130011333331333-1312320320213032-1311120130323102) |
| `no_insecure_registries` | [no_insecure_registries](data-sources--k8s_cluster--reference--group-001.md#canonical-1121203030110111-0223203223132331-0132102202000203-2020003200103333-0323132112130300-2331201123320110-1301121000311033-1032211210131023) |
| `no_local_access` | [no_local_access](data-sources--k8s_cluster--reference--group-001.md#canonical-0031011321200130-0233323111220302-1011020032330230-1230312112102033-2322032313113103-3031123333220133-1023022101301121-3312220032222023) |
| `use_custom_cluster_role_bindings` | [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-3333221130132112-1331103013220110-2313133011231003-2002132201011310-2200030123121203-1310022002102103-3120112310213200-3002222212311122) |
| `use_custom_cluster_role_bindings.cluster_role_bindings` | [use_custom_cluster_role_bindings.cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-2221031000001323-3211131211302333-2320223321013123-0300313312122303-1123320021111100-1313312112100230-0202013103302323-2023133132233100) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.name` | [use_custom_cluster_role_bindings.cluster_role_bindings.name](data-sources--k8s_cluster--reference--group-001.md#canonical-3032300103202301-1010313003031313-1033323103000112-1100031003130131-3220233332303302-3233132331221030-0023303011133320-1123131030313032) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.namespace` | [use_custom_cluster_role_bindings.cluster_role_bindings.namespace](data-sources--k8s_cluster--reference--group-001.md#canonical-3303131000111202-2022222002221332-0221100001021201-3333021311210331-0112222201130333-0132232112001322-1300033132332022-3213022123210303) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.tenant` | [use_custom_cluster_role_bindings.cluster_role_bindings.tenant](data-sources--k8s_cluster--reference--group-001.md#canonical-1021202311112200-3211003211023131-2331303021101310-2230210131330013-2321003133111221-3000013122101000-0030333310112210-1303213210311002) |
| `use_custom_cluster_role_list` | [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2212233112021010-3231210323301021-2130313323112021-1013021202001312-0301011111313112-3022322211303101-0032130300300203-2221101013332331) |
| `use_custom_cluster_role_list.cluster_roles` | [use_custom_cluster_role_list.cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-3311320011033310-1113323113323211-2323213202110232-2000023311101133-0230200332321003-0120000012202033-0013111112332122-0102311122011112) |
| `use_custom_cluster_role_list.cluster_roles.name` | [use_custom_cluster_role_list.cluster_roles.name](data-sources--k8s_cluster--reference--group-001.md#canonical-3323131202330000-1232013320013111-0113100210320012-3213302002230030-3203113332131230-1003332012320010-3022302022321302-1323002002030031) |
| `use_custom_cluster_role_list.cluster_roles.namespace` | [use_custom_cluster_role_list.cluster_roles.namespace](data-sources--k8s_cluster--reference--group-001.md#canonical-1131022012033031-2330322002213202-3010233030021111-2012201022010222-1131322323223002-1303130021310101-3002202112113131-1000303002033200) |
| `use_custom_cluster_role_list.cluster_roles.tenant` | [use_custom_cluster_role_list.cluster_roles.tenant](data-sources--k8s_cluster--reference--group-001.md#canonical-1023303122113212-2312222223110330-2300023120323120-2122132121202020-0111132332232012-3102122201000320-1102102313331331-2331121202301232) |
| `use_custom_pod_security_admission` | [use_custom_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-3013300322112212-2002321022121312-2123123230010000-2332313201102010-0320223012321001-0120311003102333-1022232320132131-3022001003103130) |
| `use_custom_pod_security_admission.name` | [use_custom_pod_security_admission.name](data-sources--k8s_cluster--reference--group-001.md#canonical-3332132322012021-0231321330232000-1232331132130020-0220123101011000-0233013211323100-0002300330011133-0030312133322021-0103122223303132) |
| `use_custom_pod_security_admission.namespace` | [use_custom_pod_security_admission.namespace](data-sources--k8s_cluster--reference--group-001.md#canonical-0300112133333000-1002122213031200-1101022233102333-1301132021231133-0123133232010003-2323212023230130-1001102210023220-1222302203220203) |
| `use_custom_pod_security_admission.tenant` | [use_custom_pod_security_admission.tenant](data-sources--k8s_cluster--reference--group-001.md#canonical-0202113122130100-3130022320200020-0023203330303100-3112100222133120-2300112212033010-3323330221321332-2321313211311111-3203001331222200) |
| `use_custom_psp_list` | [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3311023210120220-0322222023211012-2210300213333202-2103102103012311-0131232303212231-1012032300133021-0120201301223312-0302231101310102) |
| `use_custom_psp_list.pod_security_policies` | [use_custom_psp_list.pod_security_policies](data-sources--k8s_cluster--reference--group-001.md#canonical-0022000111322303-1133321310222211-1312013233021002-2213101331101312-1203220131222303-0032102313020110-3033130110230000-1032323120131033) |
| `use_custom_psp_list.pod_security_policies.name` | [use_custom_psp_list.pod_security_policies.name](data-sources--k8s_cluster--reference--group-001.md#canonical-2032231313222230-0032301011211003-2031123011120210-3201023100103302-0133210003332111-0113223012012133-2020333102100300-2332023120222301) |
| `use_custom_psp_list.pod_security_policies.namespace` | [use_custom_psp_list.pod_security_policies.namespace](data-sources--k8s_cluster--reference--group-001.md#canonical-0123112000103313-0121020102230312-0010013022132032-1033332032200201-2331212200003110-1211222332000023-2003210332320201-3133120010232002) |
| `use_custom_psp_list.pod_security_policies.tenant` | [use_custom_psp_list.pod_security_policies.tenant](data-sources--k8s_cluster--reference--group-001.md#canonical-1033323332000001-2203331321322213-1222010011311221-2002212001322122-3222323201212333-2230000321332220-1020233121101013-3310203110221120) |
| `use_default_cluster_role_bindings` | [use_default_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-2233130103302023-1310000211033111-1312311001121102-3221101201103030-3011312112103001-0010032001231222-3103003323120320-2033001210020133) |
| `use_default_cluster_roles` | [use_default_cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-1213301333310022-1000030000313220-1202033333103312-1100130121131203-3232103111112031-0330001231312102-2012021202023112-2032201223220012) |
| `use_default_pod_security_admission` | [use_default_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-2031132023103022-1303010320132012-3232033320320020-1030330320032210-0320100012001333-0110112323113031-0122132321220130-0203330302013200) |
| `use_default_psp` | [use_default_psp](data-sources--k8s_cluster--reference--group-001.md#canonical-2010021332111032-2131200311301021-3212213000010311-2023031202320012-3322010112023233-1312033102121133-0130230033031012-3021222132100320) |
| `vk8s_namespace_access_deny` | [vk8s_namespace_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-3030000233112133-3111323101120112-0310332201210202-1333321322332300-2312113230031011-3123002111031103-2201201021112323-2201013201223130) |
| `vk8s_namespace_access_permit` | [vk8s_namespace_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-0000323102132300-3112022320221332-1330003121122121-0211311010232213-2133222312020033-2110230010012021-1122102310120123-1113200220303033) |

<a id="canonical-1021002010301031-0221011101032321-1000113130223232-0313303321012311-3122023230230333-1233202322221221-0323221333332311-3013322000002131"></a>

## Next pages — Property reference / 131231220211 / 11

- [cluster_scoped_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-0120002331303020-2010212021110212-1002000103132231-3312311312112020-3222320120101120-3333021221032011-0010003213111232-1301303220301200)
- [cluster_scoped_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-1001213022130021-2301002303310100-0123321311121203-3133231111100330-2021121330000222-3200331033311311-3211110012111000-3233010132332202)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- [global_access_enable](data-sources--k8s_cluster--reference--group-001.md#canonical-3320222313103210-3123212313210032-0133030311211222-1010012023223121-3323132221202210-0121033131131221-1030202213312231-2212103112002220)
- [insecure_registry_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2000030331132132-2030211331113330-3220120113033231-2122210221011231-0213200233103132-2313012130101320-3322020233012010-0101023233133021)
- [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-1232130303333302-2123302020123001-1012302120120100-2031110103100133-1333002211210233-0210013221233333-2013211033200130-2122000103121331)
- [no_cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-3030121200123230-2331212032333120-1222020201113023-2223333222111313-1111223011003323-2212200320132301-1013302010323310-1101133032011100)
- [no_global_access](data-sources--k8s_cluster--reference--group-001.md#canonical-0200200221310301-2230002322030021-1312111002133223-1002311321102011-2021031221111102-1103311003300332-3002323333213332-0332212203031002)
- [no_insecure_registries](data-sources--k8s_cluster--reference--group-001.md#canonical-0330033302000032-3231120203000223-3031000032022331-3013223231322231-0000021312123130-1030111211103110-3222212331102001-2032232111221130)
- [no_local_access](data-sources--k8s_cluster--reference--group-001.md#canonical-0302113023030331-3313211211313300-1121010212112131-0313100112232110-1210002211320123-1030321312223201-2032113121233310-2012223221032303)
- [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-3321211133132301-2331122022223321-2012231113301103-1020230201103020-0302233321331011-2321020130101031-2012122130121322-3131100020033332)
- [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2210312112133200-2120022110103011-3332312223021001-2132020112232033-3130120230202301-1002132302112011-3330122112333332-2312202300223020)
- [use_custom_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-3012230023323023-1200010011232313-2031133231312300-1101121323331103-0310202311002101-0323110101101321-1000130001232133-0102210223300320)
- [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-0222103111213102-1022030213200111-2301201033101101-2013300120212213-2320221103200012-2233132210011313-3020211000121131-1213103021100330)
- [use_default_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-3033213011211231-1311213031231322-0331302021011032-0000211222303012-1322100223221232-2302122333012033-3131332132222230-0030211100120123)
- [use_default_cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-1030120320023032-0121111102120012-2122221010210131-3023032202310122-3302021200233321-2123032002300221-3130110013130033-1212033311130030)
- [use_default_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-1101101122322220-1120213302132230-3212131032131321-0033133201001221-1011222030321331-3121013100102132-2131330220110032-3001313332022220)
- [use_default_psp](data-sources--k8s_cluster--reference--group-001.md#canonical-0331232133213200-1233022031333312-0230013230013120-1233232013212110-2030212322112002-2103210312230132-1012023100100202-0123100212303233)
- [vk8s_namespace_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-0002333010132133-0022011023203112-3231122111102303-0131030011100223-3233022321012030-3102031213223333-1323301112022032-3031200010100021)
- [vk8s_namespace_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-0030031301002303-0032103230211033-3333313322321331-3233011120002200-2020222121001033-3120301113302031-0013121112201100-2131301013001320)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-0120002331303020-2010212021110212-1002000103132231-3312311312112020-3222320120101120-3333021221032011-0010003213111232-1301303220301200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331232110200312-0100113231030311-0000321233212112-0132132310101100-2322023121223013-0333310032132220-2322023132310231-3231123020030312"></a>

## cluster_scoped_access_deny — cluster_scoped_access_deny / 031223302311 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- cluster_scoped_access_deny

<a id="canonical-1223010310201321-0021323111131231-1313113313002332-0230030320213032-3000102111133112-1101132103212333-1102310120331320-3201222203123122"></a>

Type: `["object", {}]`. Computed.

\[OneOf: cluster\_scoped\_access\_deny, cluster\_scoped\_access\_permit\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [cluster_scoped_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-1223010310201321-0021323111131231-1313113313002332-0230030320213032-3000102111133112-1101132103212333-1102310120331320-3201222203123122)
- [cluster_scoped_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-3320332310023201-3312021121102213-3230122120232112-2232313322003123-3321133322003112-3013232202013323-0311200200212130-2101033023323221)

Select alternatives according to the provider validators above.

<a id="canonical-0332300020111012-0213302302311022-3312233003113323-1230310211221113-0133132303202121-0113332112113023-3021120203031021-0300333023003201"></a>

## Direct properties — cluster_scoped_access_deny / 031223302311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212132200222220-1110130321332013-2322331031033222-1311221222222002-0232201023323210-3021032302323123-3232311232303122-2132031331120020"></a>

## Next pages — cluster_scoped_access_deny / 031223302311 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-1001213022130021-2301002303310100-0123321311121203-3133231111100330-2021121330000222-3200331033311311-3211110012111000-3233010132332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132022330020103-1131200211113123-0101220020102122-0132221101313320-2222201200012321-2213322031010033-2330130333100112-0100231323101130"></a>

## cluster_scoped_access_permit — cluster_scoped_access_permit / 130111033020 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- cluster_scoped_access_permit

<a id="canonical-3320332310023201-3312021121102213-3230122120232112-2232313322003123-3321133322003112-3013232202013323-0311200200212130-2101033023323221"></a>

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

<a id="canonical-0202222133132020-3311010220031012-2112003323303002-0322030231012311-0030331230103132-0210201102231231-0120032313330333-0123000231230020"></a>

## Direct properties — cluster_scoped_access_permit / 130111033020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023011033303200-3310222120202000-1213303331233322-1223030200012220-2223010330103221-2010133200012021-0202320023123330-1100021200102022"></a>

## Next pages — cluster_scoped_access_permit / 130111033020 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200002333313321-3031102310230231-0102233310130221-3020330331200111-3010102001222301-2322112313313212-3133003221012031-2022013121201331"></a>

## cluster_wide_app_list — cluster_wide_app_list / 031102211012 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- cluster_wide_app_list

<a id="canonical-0210223233021201-3011030020003221-1332211103002212-1013020101103030-3000113113001213-2032203000032203-0310312102131020-0111121303021311"></a>

Type: `"single"`. Computed.

\[OneOf: cluster\_wide\_app\_list, no\_cluster\_wide\_apps; Default: no\_cluster\_wide\_apps\]
Cluster Wide Application List. List of cluster wide applications.

Upstream description:

List of cluster wide applications.

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

- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-0210223233021201-3011030020003221-1332211103002212-1013020101103030-3000113113001213-2032203000032203-0310312102131020-0111121303021311)
- [no_cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-2121002130233322-2131300330321301-0223201213012213-3213022102132001-0220213303001220-3100032322320313-0232020312010330-1020131130220203)

Select alternatives according to the provider validators above.

<a id="canonical-2221132000010110-0200323102011032-0200122232233232-2131313013102001-0331232223103021-0322132103133113-1203300110203212-2031232100333132"></a>

## Direct properties — cluster_wide_app_list / 031102211012 / 3

- [cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222): complete subsection reference.

<a id="canonical-0230112033101301-0332320131330130-2122222023231301-2203233310301330-1021222330313132-0030310330301003-3220120203311212-2003310312131010"></a>

## Next pages — cluster_wide_app_list / 031102211012 / 4

- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102030231302130-1321013322123311-2320123320120233-3120321221031103-1203013130120013-2120123001101301-0030201033320230-3030323223231100"></a>

## cluster_wide_app_list.cluster_wide_apps — cluster_wide_apps / 212132333123 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- cluster_wide_app_list.cluster_wide_apps

<a id="canonical-2320002113020121-0210113323132302-3231332100023110-2121003001120131-3200230323003132-3320013000303103-2101033230032003-2201003333221203"></a>

Type: `"list"`. Computed.

Cluster Wide Application List. List of cluster wide applications.

Upstream description:

List of cluster wide applications.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2133210031203130-2230320321220101-3220223100233120-2133202302300213-3033330133302023-0220131310033330-1332300312232110-1223132012201131"></a>

## Direct properties — cluster_wide_apps / 212132333123 / 3

- [argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-0133323210331201-3023333103300203-2121222213322202-3321111201013101-1033221030222322-3203133020133011-1332312230130023-2302130011030022): complete subsection reference.

- [dashboard](data-sources--k8s_cluster--reference--group-001.md#canonical-1232122232311011-2120233111220112-1010233123331223-3012230033300110-2110331223200011-3032221223130121-3312032303210022-0213212320233000): complete subsection reference.

- [metrics_server](data-sources--k8s_cluster--reference--group-001.md#canonical-3323332331301302-0231111010321312-1020222213230002-3121210311003110-1321031223233130-3323122111033331-3123231120130132-3102233022312001): complete subsection reference.

- [prometheus](data-sources--k8s_cluster--reference--group-001.md#canonical-1220200202300231-3022032303002030-1132133033320103-3211210131302202-1302112013133033-1333313310303303-0231103003331010-0303013010012113): complete subsection reference.

<a id="canonical-2201320002101001-0132221210332101-0020302331003120-1331001230212002-0211020010323020-0232330300100011-0203221031022011-1110321312312022"></a>

## Next pages — cluster_wide_apps / 212132333123 / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-0133323210331201-3023333103300203-2121222213322202-3321111201013101-1033221030222322-3203133020133011-1332312230130023-2302130011030022)
- [cluster_wide_app_list.cluster_wide_apps.dashboard](data-sources--k8s_cluster--reference--group-001.md#canonical-1232122232311011-2120233111220112-1010233123331223-3012230033300110-2110331223200011-3032221223130121-3312032303210022-0213212320233000)
- [cluster_wide_app_list.cluster_wide_apps.metrics_server](data-sources--k8s_cluster--reference--group-001.md#canonical-3323332331301302-0231111010321312-1020222213230002-3121210311003110-1321031223233130-3323122111033331-3123231120130132-3102233022312001)
- [cluster_wide_app_list.cluster_wide_apps.prometheus](data-sources--k8s_cluster--reference--group-001.md#canonical-1220200202300231-3022032303002030-1132133033320103-3211210131302202-1302112013133033-1333313310303303-0231103003331010-0303013010012113)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-0133323210331201-3023333103300203-2121222213322202-3321111201013101-1033221030222322-3203133020133011-1332312230130023-2302130011030022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131203031212133-2332230332031121-3310101012313321-0021131031031201-2120010331200320-3222003110210332-1310133201103321-2231102010130033"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd — argo_cd / 010221331000 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- cluster_wide_app_list.cluster_wide_apps.argo_cd

<a id="canonical-1030031302211133-2121223320020233-3031120231003123-2203002212101012-3133110113110130-1001233101103230-1032323310222133-2212010133033232"></a>

Type: `"single"`. Computed.

Description Parameters for Argo Continuous Deployment(CD) application.

Upstream description:

Description Parameters for Argo Continuous Deployment(CD) application.

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

<a id="canonical-2322120203022230-3221132110033231-1210333213232320-1020200320020101-3332113213223132-2331231221130321-1200032231012233-0211132302212102"></a>

## Direct properties — argo_cd / 010221331000 / 3

- [local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-0122232212231200-3012202233331230-2202011001323230-1120311323300112-2321133332001002-2301310012320121-2013222113203032-2200330001333023): complete subsection reference.

<a id="canonical-1323202221202331-2111211003331003-2120033120032222-1000002020121333-2313112032333023-2022300312212100-2310112212133203-2022031122313330"></a>

## Next pages — argo_cd / 010221331000 / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-0122232212231200-3012202233331230-2202011001323230-1120311323300112-2321133332001002-2301310012320121-2013222113203032-2200330001333023)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-0122232212231200-3012202233331230-2202011001323230-1120311323300112-2321133332001002-2301310012320121-2013222113203032-2200330001333023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010000101232321-1212113221002101-0123200211110023-2222020221220120-1120003010331123-0032021112332230-3311331330213103-2321200021222131"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain — local_domain / 302033131302 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-0133323210331201-3023333103300203-2121222213322202-3321111201013101-1033221030222322-3203133020133011-1332312230130023-2302130011030022)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain

<a id="canonical-0221132301002231-1132222303001001-2102330312103313-3303131203111333-1331112011113120-1122313103311213-1002010333131322-3003021221203200"></a>

Type: `"single"`. Computed.

Parameters required to enable local access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"default_port\",\"port\"]"
}
```

<a id="canonical-2101330013333121-1201122213120222-2210133022022132-3030121312332310-1231112103023021-2230332211210023-3300332211102221-2210132130121310"></a>

## Direct properties — local_domain / 302033131302 / 3

- [default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-2333302110002321-1111203231002130-2202231131211001-0011200112130322-2202230011210320-1131203332320031-2301232221220021-0000122200333111): complete subsection reference.

<a id="canonical-3202300211310002-3303221312133110-1321200122101331-1103333232103211-1123112002031121-2231101332010012-1121002012323111-0023211021132221"></a>

<a id="canonical-2213020003003322-3331031122333121-3223333331033231-2301113233012012-1101020220201233-3310102013310330-1113303032012013-1220131000223113"></a>

## local_domain property — local_domain / 302033131302 / 4

Type: `"string"`. Computed.

ArgoCD will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 192,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [password](data-sources--k8s_cluster--reference--group-001.md#canonical-1100011000321022-3202233233133220-2103202133320330-1011131020001201-3113301230031120-2213302313313000-2001033023033303-2333203212000012): complete subsection reference.

<a id="canonical-3122031332313003-2213233121200113-0131220101220030-3220133113202320-2022101031123012-0223011101101213-2302003233123312-0032130131122332"></a>

<a id="canonical-0321103302101232-2213033312003201-1330130303103203-2332002010130331-3030011202223123-1322021103133203-3012130120131230-3010131202030301"></a>

## port property — local_domain / 302033131302 / 5

Type: `"number"`. Computed.

Exclusive with \[default\_port\] Use custom ArgoCD port. Available port range is less than 65000
except reserved ports.

Upstream description:

Exclusive with \[default\_port\] Use custom ArgoCD port. Available port range is less than 65000
except reserved ports.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  }
}
```

<a id="canonical-1303120102121021-1223011030213122-3211113020221303-2112311211033113-3133030132312011-3222010130331322-1022102132301121-0032321003231211"></a>

## Next pages — local_domain / 302033131302 / 6

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-2333302110002321-1111203231002130-2202231131211001-0011200112130322-2202230011210320-1131203332320031-2301232221220021-0000122200333111)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-1100011000321022-3202233233133220-2103202133320330-1011131020001201-3113301230031120-2213302313313000-2001033023033303-2333203212000012)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-0133323210331201-3023333103300203-2121222213322202-3321111201013101-1033221030222322-3203133020133011-1332312230130023-2302130011030022)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-2333302110002321-1111203231002130-2202231131211001-0011200112130322-2202230011210320-1131203332320031-2301232221220021-0000122200333111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301213301031201-2312210020233131-3121013001033110-0032112311212012-0030323101333032-1232211013133031-0110112122011233-3233022310230333"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port — default_port / 312031230000 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-0133323210331201-3023333103300203-2121222213322202-3321111201013101-1033221030222322-3203133020133011-1332312230130023-2302130011030022)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-0122232212231200-3012202233331230-2202011001323230-1120311323300112-2321133332001002-2301310012320121-2013222113203032-2200330001333023)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port

<a id="canonical-2210113033323010-1123321133210020-3111302203113121-1100303331203201-2200313100022332-2221333123311210-3032331322202233-3111233131312033"></a>

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

<a id="canonical-1212022013121001-2210111203103221-3033023102302220-0031132322003112-2120221312202031-0331131131201223-3002320033210213-3103132303231012"></a>

## Direct properties — default_port / 312031230000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200231132021100-0021103010312011-3120213012003101-1102031301211131-0313012222200333-3102313303321123-1022213121122331-3111012202221112"></a>

## Next pages — default_port / 312031230000 / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-0122232212231200-3012202233331230-2202011001323230-1120311323300112-2321133332001002-2301310012320121-2013222113203032-2200330001333023)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-1100011000321022-3202233233133220-2103202133320330-1011131020001201-3113301230031120-2213302313313000-2001033023033303-2333203212000012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331330231032302-1210223121021030-3031133020020103-2303121133121320-2010102231112013-0122123313222112-1202230022311232-1130031323201333"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password — password / 230112301201 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-0133323210331201-3023333103300203-2121222213322202-3321111201013101-1033221030222322-3203133020133011-1332312230130023-2302130011030022)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-0122232212231200-3012202233331230-2202011001323230-1120311323300112-2321133332001002-2301310012320121-2013222113203032-2200330001333023)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password

<a id="canonical-1030123311021300-1101231011333301-1000313021101123-0323122301320103-2220332203033311-1222023132101303-0323123203103032-3302233021211223"></a>

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

<a id="canonical-0212320010120131-0033012011323221-2123132210213302-0010313121032211-3223001230310120-3010331213110101-1333333310211230-2102123222320123"></a>

## Direct properties — password / 230112301201 / 3

- [blindfold_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-1302213210033203-2112103130132200-3323230130023032-2110233032033233-1111230120331021-2013120002100030-0021130030112313-2313210133030301): complete subsection reference.

- [clear_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-2323320100132231-0003320103112230-2330022020231301-1102331303200120-3020321320133330-0131221101200320-0220033033321310-2013331010310212): complete subsection reference.

<a id="canonical-0303301333032231-3130122022020301-1130333102320030-2201121131331201-1230231312232223-3113321202101103-0211211123001333-3120303301310020"></a>

## Next pages — password / 230112301201 / 4

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-1302213210033203-2112103130132200-3323230130023032-2110233032033233-1111230120331021-2013120002100030-0021130030112313-2313210133030301)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info](data-sources--k8s_cluster--reference--group-001.md#canonical-2323320100132231-0003320103112230-2330022020231301-1102331303200120-3020321320133330-0131221101200320-0220033033321310-2013331010310212)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-0122232212231200-3012202233331230-2202011001323230-1120311323300112-2321133332001002-2301310012320121-2013222113203032-2200330001333023)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-1302213210033203-2112103130132200-3323230130023032-2110233032033233-1111230120331021-2013120002100030-0021130030112313-2313210133030301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100120233321021-0231212203013311-0313011031002033-1300221002003212-1323313320231310-3121011003223233-1113202121010221-3103013023113301"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info — blindfold_secret_info / 320200302130 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-0133323210331201-3023333103300203-2121222213322202-3321111201013101-1033221030222322-3203133020133011-1332312230130023-2302130011030022)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-0122232212231200-3012202233331230-2202011001323230-1120311323300112-2321133332001002-2301310012320121-2013222113203032-2200330001333023)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-1100011000321022-3202233233133220-2103202133320330-1011131020001201-3113301230031120-2213302313313000-2001033023033303-2333203212000012)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info

<a id="canonical-2333322301332203-2300023032132301-1212221012023112-2302011200021120-2320030121103133-0122233230033012-1233123203033122-0231320122133020"></a>

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

<a id="canonical-0113230022110032-3330313002331313-1201001212012010-3033211301003000-3211303121030102-3321202223132000-2221000311312102-3021323131322132"></a>

## Direct properties — blindfold_secret_info / 320200302130 / 3

<a id="canonical-2201133102210001-1321103231200021-2301110202023012-3212031121300130-0013111331110131-1222210032333020-2122010211130132-1020233313303121"></a>

<a id="canonical-0112031311220210-0330332130201333-3032313233032132-1033101320001020-3233000331023121-0230301301201223-0022210323022203-3320112033333222"></a>

## decryption_provider property — blindfold_secret_info / 320200302130 / 4

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

<a id="canonical-2320311011133321-1210113132111210-1100001020200032-1201313201210333-0030313113000202-0313230213221203-0301230222010033-1201022201020320"></a>

<a id="canonical-2333311022010213-0323022113303300-0332030221132230-3022023331001313-3031012123131231-2123212310203213-2220302300223030-0221112103210120"></a>

## location property — blindfold_secret_info / 320200302130 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0113301120032303-3131212221232000-3031311220131123-0321221301113011-1111203320111303-0202230313003231-1122321013201101-2311022211022320"></a>

<a id="canonical-2230300203333200-1111220113010130-2320012122003320-1023031331300230-0330100123313333-3203011320300213-0322310333233300-3012121323211132"></a>

## store_provider property — blindfold_secret_info / 320200302130 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-3100331101211323-2203322331320330-0332210300013232-1213222303220301-2130012112132213-0201001112203312-0330122323313220-2220032323020222"></a>

## Next pages — blindfold_secret_info / 320200302130 / 7

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-1100011000321022-3202233233133220-2103202133320330-1011131020001201-3113301230031120-2213302313313000-2001033023033303-2333203212000012)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-2323320100132231-0003320103112230-2330022020231301-1102331303200120-3020321320133330-0131221101200320-0220033033321310-2013331010310212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211013012322131-0201001010031131-2111101332323323-0320003121100212-1302020111323123-0122122323213131-2122301303112312-2023123220212321"></a>

## cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info — clear_secret_info / 201220223212 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](data-sources--k8s_cluster--reference--group-001.md#canonical-0133323210331201-3023333103300203-2121222213322202-3321111201013101-1033221030222322-3203133020133011-1332312230130023-2302130011030022)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](data-sources--k8s_cluster--reference--group-001.md#canonical-0122232212231200-3012202233331230-2202011001323230-1120311323300112-2321133332001002-2301310012320121-2013222113203032-2200330001333023)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-1100011000321022-3202233233133220-2103202133320330-1011131020001201-3113301230031120-2213302313313000-2001033023033303-2333203212000012)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info

<a id="canonical-0213213011101002-2311102013000321-2331223213010320-2332001330022333-0010133123031233-0300130212311211-2102132033232010-1011121130331302"></a>

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

<a id="canonical-2000120233322321-2033330110313230-0313030321301332-2022133331200321-0203001332203312-1111221000122032-3133000310311032-1020303002032231"></a>

## Direct properties — clear_secret_info / 201220223212 / 3

<a id="canonical-2312203211030321-2222230223012113-3112323020330020-3322200312210000-3210010110311011-2013302310010230-0332030311330222-3211220110230121"></a>

<a id="canonical-2331112333021212-3021313300031200-3122233033313210-2002213230130332-1212103320323022-2303100313012123-2220120211203200-0301223012130132"></a>

## provider_ref property — clear_secret_info / 201220223212 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3303332231101113-2200310323000031-2221210210101122-3021211301211020-2320113030210230-2302232000112020-3013111311003030-2301033000313223"></a>

<a id="canonical-2200002131202310-3003230022222201-2231331101032333-2013102333230132-2203002310303030-1121022222121011-1113001031211323-2333002031110022"></a>

## URL property — clear_secret_info / 201220223212 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0020033020012323-1130300313233220-2213002203101022-0132103033002331-0101003303331300-3310131122003321-2030102331032233-1320223232113131"></a>

## Next pages — clear_secret_info / 201220223212 / 6

- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](data-sources--k8s_cluster--reference--group-001.md#canonical-1100011000321022-3202233233133220-2103202133320330-1011131020001201-3113301230031120-2213302313313000-2001033023033303-2333203212000012)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-1232122232311011-2120233111220112-1010233123331223-3012230033300110-2110331223200011-3032221223130121-3312032303210022-0213212320233000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010001003032220-1310022301211003-0033020321200213-3230303100021033-1113202230111032-2102310003101300-3031003302330211-2212111302212031"></a>

## cluster_wide_app_list.cluster_wide_apps.dashboard — dashboard / 222003330020 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- cluster_wide_app_list.cluster_wide_apps.dashboard

<a id="canonical-0302003332300321-2302233302000312-1003020112022322-3003031332300131-3133203132033002-3031002133122333-2020323310210120-2102101222022202"></a>

Type: `["object", {}]`. Computed.

Description Parameters for K8s dashboard.

Upstream description:

Description Parameters for K8s dashboard.

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

<a id="canonical-1212112112323123-0110112000133002-0233010012113030-0001221120313132-3011010011231202-0121301303230031-3113221031011011-0103213211232102"></a>

## Direct properties — dashboard / 222003330020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313020202332213-3320301231011301-3003301121202013-3233120202003223-0003021121000313-0022120303032211-2223201322330321-0223312033102203"></a>

## Next pages — dashboard / 222003330020 / 4

- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-3323332331301302-0231111010321312-1020222213230002-3121210311003110-1321031223233130-3323122111033331-3123231120130132-3102233022312001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133311232013233-1031231033111203-3311333212022213-1323110330321113-2111220302222033-3001033223032112-0001122113123310-0010011213201233"></a>

## cluster_wide_app_list.cluster_wide_apps.metrics_server — metrics_server / 131113230023 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- cluster_wide_app_list.cluster_wide_apps.metrics_server

<a id="canonical-2302122132223022-0203212022131113-1012112003132030-3123013233120123-1220300232023312-3212331330032001-2012023121322033-1220020301113032"></a>

Type: `["object", {}]`. Computed.

Description Parameters for Kubernetes Metrics Server application.

Upstream description:

Description Parameters for Kubernetes Metrics Server application.

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

<a id="canonical-1300021303303303-3332110122131331-1121023322112202-2112232030231231-1003113131103032-0202232131311022-3030012312131030-3223132121120322"></a>

## Direct properties — metrics_server / 131113230023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210310230112230-3222122223113121-2020212211210133-3212011020013113-3303112131332230-0200113301232231-3203331002212102-0022302133103101"></a>

## Next pages — metrics_server / 131113230023 / 4

- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-1220200202300231-3022032303002030-1132133033320103-3211210131302202-1302112013133033-1333313310303303-0231103003331010-0303013010012113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111200111313331-1322011011210122-3131102333131330-0233101110000033-3300020211010231-1003212201203010-1120313122303233-1000032132013123"></a>

## cluster_wide_app_list.cluster_wide_apps.prometheus — prometheus / 112233022232 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [cluster_wide_app_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3012002212023212-0312020122110302-1031000322301012-1133301000011230-1113222032032232-2113012000333012-3031101320033120-2033112013002132)
- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- cluster_wide_app_list.cluster_wide_apps.prometheus

<a id="canonical-2321122102032213-0123230123323011-0332012023023132-2310123212100210-0301323102033012-3102003302030331-0232031113003022-0312223031102000"></a>

Type: `["object", {}]`. Computed.

Description Parameters for Prometheus server access.

Upstream description:

Description Parameters for Prometheus server access.

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

<a id="canonical-2122110230300311-3200311201112121-1113233333331332-0103012130130130-2301312331120213-0213313003301110-0121023212101211-3211000302012001"></a>

## Direct properties — prometheus / 112233022232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232220120022122-0021201020120332-2031213122322101-3200313320032332-2101331012303012-2101321213121113-3313223331213233-0102101120232130"></a>

## Next pages — prometheus / 112233022232 / 4

- [cluster_wide_app_list.cluster_wide_apps](data-sources--k8s_cluster--reference--group-001.md#canonical-1311132120113231-1220003233130312-3200110311213311-1200212223120103-2102323332111012-3311211121231222-2133100111131320-1333302222203222)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-3320222313103210-3123212313210032-0133030311211222-1010012023223121-3323132221202210-0121033131131221-1030202213312231-2212103112002220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031030212221201-0320103221033003-0002311021231122-2102231103223102-3202213021032230-3310100202123232-0013123113213022-2020113102331331"></a>

## global_access_enable — global_access_enable / 333021033021 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- global_access_enable

<a id="canonical-3023332120102213-3112021003210331-3113131231200301-1112120212001221-2211003000120203-2330233010302133-3121323323332122-2301012200120023"></a>

Type: `["object", {}]`. Computed.

\[OneOf: global\_access\_enable, no\_global\_access; Default: no\_global\_access\] Configuration
parameter for global access enable.

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

OneOf alternatives in this subsection:

- [global_access_enable](data-sources--k8s_cluster--reference--group-001.md#canonical-3023332120102213-3112021003210331-3113131231200301-1112120212001221-2211003000120203-2330233010302133-3121323323332122-2301012200120023)
- [no_global_access](data-sources--k8s_cluster--reference--group-001.md#canonical-3230320232203232-1223300012302210-2003112003211112-0113131023100013-2212101202031130-2130011333331333-1312320320213032-1311120130323102)

Select alternatives according to the provider validators above.

<a id="canonical-1301102223121122-1321121310311101-0210320313213033-0222222003303200-1122113011231121-3231010112032221-2200330321323311-2101000221113200"></a>

## Direct properties — global_access_enable / 333021033021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220232210302030-3121232102323013-1100320313130032-2220222333323031-3300110200201220-3113212212333132-0212010023003100-0033220102131022"></a>

## Next pages — global_access_enable / 333021033021 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-2000030331132132-2030211331113330-3220120113033231-2122210221011231-0213200233103132-2313012130101320-3322020233012010-0101023233133021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320130231001331-1330102002301300-3001203022100211-3101033212303110-3201203123032133-2123210011322010-0133302312031230-1322033320000000"></a>

## insecure_registry_list — insecure_registry_list / 201230223020 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- insecure_registry_list

<a id="canonical-3033313331133230-1212122030330001-0033112013020112-1332331222101110-0013100013333132-0130210132212132-1302321302203102-0331202111123212"></a>

Type: `"single"`. Computed.

\[OneOf: insecure\_registry\_list, no\_insecure\_registries; Default: no\_insecure\_registries\]
Docker Insecure Registry List. List of Docker insecure registries.

Upstream description:

List of Docker insecure registries.

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

- [insecure_registry_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3033313331133230-1212122030330001-0033112013020112-1332331222101110-0013100013333132-0130210132212132-1302321302203102-0331202111123212)
- [no_insecure_registries](data-sources--k8s_cluster--reference--group-001.md#canonical-1121203030110111-0223203223132331-0132102202000203-2020003200103333-0323132112130300-2331201123320110-1301121000311033-1032211210131023)

Select alternatives according to the provider validators above.

<a id="canonical-0000121200101221-0133012111020013-3213222323123233-2013100112200302-1233230310332013-1233302300332323-3200322213233122-2013012221102213"></a>

## Direct properties — insecure_registry_list / 201230223020 / 3

<a id="canonical-2323031033113222-0310201133313030-0212313312011122-3233101021112203-2303223221000202-3331200123221031-3303131300302233-2110311000100021"></a>

<a id="canonical-0300013122330211-1003101222333032-3301320113320230-3332232200100321-0313311103233112-3033330330222313-0102121211031130-2120120210032201"></a>

## insecure_registries property — insecure_registry_list / 201230223020 / 4

Type: `["list", "string"]`. Computed.

List of Docker insecure registries in format 'example.com:5000'.

Upstream description:

List of Docker insecure registries in format "example.com:5000"

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
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0201301033133211-1102302031011322-0233031012002032-1030322132323232-3331231120320012-2220021212003111-2333300022022212-1300220310131233"></a>

## Next pages — insecure_registry_list / 201230223020 / 5

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-1232130303333302-2123302020123001-1012302120120100-2031110103100133-1333002211210233-0210013221233333-2013211033200130-2122000103121331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010032212330333-0200313213321131-1110232221312102-2211022310213310-2312122301211010-0121200303221132-2012232100230111-3023013020131333"></a>

## local_access_config — local_access_config / 203332021210 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- local_access_config

<a id="canonical-2122323233102321-3222222131003200-3012303112121122-1332101110200132-3230201133201033-2022232123013110-0033012231222210-2213302102032333"></a>

Type: `"single"`. Computed.

\[OneOf: local\_access\_config, no\_local\_access; Default: no\_local\_access\] Parameters required
to enable local access.

Upstream description:

Parameters required to enable local access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"default_port\",\"port\"]"
}
```

OneOf alternatives in this subsection:

- [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-2122323233102321-3222222131003200-3012303112121122-1332101110200132-3230201133201033-2022232123013110-0033012231222210-2213302102032333)
- [no_local_access](data-sources--k8s_cluster--reference--group-001.md#canonical-0031011321200130-0233323111220302-1011020032330230-1230312112102033-2322032313113103-3031123333220133-1023022101301121-3312220032222023)

Select alternatives according to the provider validators above.

<a id="canonical-3222123313002232-2231330030330112-1030200120120111-3303321221210233-1031121321300103-0210101220020201-0231311110330100-2100133013302230"></a>

## Direct properties — local_access_config / 203332021210 / 3

- [default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-0222020131003301-1323001311331011-0002213031210122-3311300003003213-0313203230212223-0332003121313013-1230011321233111-0022030132003121): complete subsection reference.

<a id="canonical-1031100012223030-1132002323231031-0010112333123220-1223010103012122-3000012313003013-0223003101211131-1013100200101032-2232321030310222"></a>

<a id="canonical-1312001230210303-1300330133023233-1103203220331030-2233322220311012-3223211022112210-1013130301233120-0033120322312000-3230321100100010"></a>

## local_domain property — local_access_config / 203332021210 / 4

Type: `"string"`. Computed.

Local K8s API server will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 192,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1023233113132023-3302231201212112-0122013113210132-2231331203101233-3302321231202203-1102100122320322-1120013310331332-3222211202031000"></a>

<a id="canonical-2010001020223031-3310223232233023-2311121303110121-0001102033202201-1313212332121000-1210213330022202-1133103233203011-3301312231330303"></a>

## port property — local_access_config / 203332021210 / 5

Type: `"number"`. Computed.

Exclusive with \[default\_port\] Use custom K8s port for API server. Available port range is less
than 65000 except reserved ports.

Upstream description:

Exclusive with \[default\_port\] Use custom K8s port for API server. Available port range is less
than 65000 except reserved ports.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  }
}
```

<a id="canonical-3302331323000111-0100112013100313-2121311103133200-0301303302121222-0232213120200231-3120331100233303-0113102230320231-0123130213012010"></a>

## Next pages — local_access_config / 203332021210 / 6

- [local_access_config.default_port](data-sources--k8s_cluster--reference--group-001.md#canonical-0222020131003301-1323001311331011-0002213031210122-3311300003003213-0313203230212223-0332003121313013-1230011321233111-0022030132003121)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-0222020131003301-1323001311331011-0002213031210122-3311300003003213-0313203230212223-0332003121313013-1230011321233111-0022030132003121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231131331001121-3101132333313031-1312130102122123-1120110123303001-2301012302001321-1221301322233121-0032102130212200-1031213231110133"></a>

## local_access_config.default_port — default_port / 310331012213 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-1232130303333302-2123302020123001-1012302120120100-2031110103100133-1333002211210233-0210013221233333-2013211033200130-2122000103121331)
- local_access_config.default_port

<a id="canonical-0311330110113031-2033130021301223-3212311323130200-0122313302123323-1011123001020031-0002022002200100-1101001211321230-0330232200000303"></a>

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

<a id="canonical-0030333201121101-0123201003301012-0310320321311311-1022011010212112-3220010312201111-1001322120110103-0231112102302310-3012120310131310"></a>

## Direct properties — default_port / 310331012213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000312220002110-3301302011230110-1333113033133031-0210313312312020-2232100232221133-3301120331133222-2130021001000331-3330002302303303"></a>

## Next pages — default_port / 310331012213 / 4

- [local_access_config](data-sources--k8s_cluster--reference--group-001.md#canonical-1232130303333302-2123302020123001-1012302120120100-2031110103100133-1333002211210233-0210013221233333-2013211033200130-2122000103121331)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-3030121200123230-2331212032333120-1222020201113023-2223333222111313-1111223011003323-2212200320132301-1013302010323310-1101133032011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012210001121132-2033331013333112-0213101222203100-3302312231320313-1010312300231021-0201121321132011-3022013113202232-0102223213010132"></a>

## no_cluster_wide_apps — no_cluster_wide_apps / 022203222222 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- no_cluster_wide_apps

<a id="canonical-2121002130233322-2131300330321301-0223201213012213-3213022102132001-0220213303001220-3100032322320313-0232020312010330-1020131130220203"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-1011000121030113-1102001112202131-1222323312200320-3212210130012330-1322102013022003-3312100011311223-0133032101102311-1013011331103310"></a>

## Direct properties — no_cluster_wide_apps / 022203222222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203323201132101-3310102010013221-1122031103010001-0102030231103102-1322113002233110-2223102122222210-3110313202111022-2110132122010210"></a>

## Next pages — no_cluster_wide_apps / 022203222222 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-0200200221310301-2230002322030021-1312111002133223-1002311321102011-2021031221111102-1103311003300332-3002323333213332-0332212203031002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323302012020021-0022313310302203-3311322012002311-1301210111032012-0211030103003122-0100220130103311-2113331203132032-1323120323001222"></a>

## no_global_access — no_global_access / 230000221010 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- no_global_access

<a id="canonical-3230320232203232-1223300012302210-2003112003211112-0113131023100013-2212101202031130-2130011333331333-1312320320213032-1311120130323102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no global access. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-0211313302012312-2011213120333300-2131201133120033-1230003123122313-2013220112133200-0232010012203222-0230210000210302-2012223120102201"></a>

## Direct properties — no_global_access / 230000221010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211131110102033-0310100220032323-1020111321002000-1003312312100100-0021002002223012-2203211221322111-2202300133303233-1113031303221310"></a>

## Next pages — no_global_access / 230000221010 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-0330033302000032-3231120203000223-3031000032022331-3013223231322231-0000021312123130-1030111211103110-3222212331102001-2032232111221130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012112020130212-1312333333332333-0322010111203000-3111211302211023-0300331313222000-0121000310202333-0323130231020302-1223130103122100"></a>

## no_insecure_registries — no_insecure_registries / 302121301013 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- no_insecure_registries

<a id="canonical-1121203030110111-0223203223132331-0132102202000203-2020003200103333-0323132112130300-2331201123320110-1301121000311033-1032211210131023"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-0033133023002212-0122013002012023-0231302000200130-3213302323210233-3120200323010022-2302102222020130-3112210033131121-1212133202112202"></a>

## Direct properties — no_insecure_registries / 302121301013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203132300010312-3211212221011301-1222022331213231-3320331111212303-0300121030333031-3202100000232332-1122312033101032-0332123212031310"></a>

## Next pages — no_insecure_registries / 302121301013 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-0302113023030331-3313211211313300-1121010212112131-0313100112232110-1210002211320123-1030321312223201-2032113121233310-2012223221032303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000321200100033-2020031302312232-2132112202113003-3002030221111302-0330121232303010-3222230230112202-3020102203010033-2030003022310322"></a>

## no_local_access — no_local_access / 100203113103 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- no_local_access

<a id="canonical-0031011321200130-0233323111220302-1011020032330230-1230312112102033-2322032313113103-3031123333220133-1023022101301121-3312220032222023"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no local access. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-0030022031023120-1112112122310203-2112110112202331-3202203112003212-2021012220101021-2030013311330220-3030103322030212-2220232130313120"></a>

## Direct properties — no_local_access / 100203113103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202021032313223-0102122031210001-2333010200122122-2102210200332202-1133202211103033-0201310212221202-2200121000023302-1103200320333212"></a>

## Next pages — no_local_access / 100203113103 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-3321211133132301-2331122022223321-2012231113301103-1020230201103020-0302233321331011-2321020130101031-2012122130121322-3131100020033332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200012231013232-3132222330121320-3133130132100021-1202330330233301-2002212201102230-3000303001301211-1120100123231130-2313013013202202"></a>

## use_custom_cluster_role_bindings — use_custom_cluster_role_bindings / 320330123003 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- use_custom_cluster_role_bindings

<a id="canonical-3333221130132112-1331103013220110-2313133011231003-2002132201011310-2200030123121203-1310022002102103-3120112310213200-3002222212311122"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_cluster\_role\_bindings, use\_default\_cluster\_role\_bindings; Default:
use\_default\_cluster\_role\_bindings\] List of active cluster role binding list for a K8s cluster.

Upstream description:

List of active cluster role binding list for a K8s cluster.

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

- [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-3333221130132112-1331103013220110-2313133011231003-2002132201011310-2200030123121203-1310022002102103-3120112310213200-3002222212311122)
- [use_default_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-2233130103302023-1310000211033111-1312311001121102-3221101201103030-3011312112103001-0010032001231222-3103003323120320-2033001210020133)

Select alternatives according to the provider validators above.

<a id="canonical-0130100123131002-1032212301203011-1213333101031333-0102232010132003-2033200322233031-0103321211313311-2000331023300320-2211100100302200"></a>

## Direct properties — use_custom_cluster_role_bindings / 320330123003 / 3

- [cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-2132212101120032-1212222212032202-1202102000030113-1331203103111032-1313332033011223-2300101132120011-0223322031111001-0210201032301123): complete subsection reference.

<a id="canonical-0113221300220303-2121233200113300-3122012323103233-2233121120213322-3233222201301012-1223103032113220-0222030101332311-2312122011211210"></a>

## Next pages — use_custom_cluster_role_bindings / 320330123003 / 4

- [use_custom_cluster_role_bindings.cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-2132212101120032-1212222212032202-1202102000030113-1331203103111032-1313332033011223-2300101132120011-0223322031111001-0210201032301123)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-2132212101120032-1212222212032202-1202102000030113-1331203103111032-1313332033011223-2300101132120011-0223322031111001-0210201032301123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010010022330222-3322301122213211-0030203323212221-0200230221033323-3030111113331112-2120112130210323-2121311023013110-2231220111313222"></a>

## use_custom_cluster_role_bindings.cluster_role_bindings — cluster_role_bindings / 303223020203 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-3321211133132301-2331122022223321-2012231113301103-1020230201103020-0302233321331011-2321020130101031-2012122130121322-3131100020033332)
- use_custom_cluster_role_bindings.cluster_role_bindings

<a id="canonical-2221031000001323-3211131211302333-2320223321013123-0300313312122303-1123320021111100-1313312112100230-0202013103302323-2023133132233100"></a>

Type: `"list"`. Computed.

List of active cluster role binding list for a K8s cluster.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3302012321230011-2123003110322333-0323312302002221-1133100301233231-3221013120321231-1023120233012012-3003201311103003-1223300030010030"></a>

## Direct properties — cluster_role_bindings / 303223020203 / 3

<a id="canonical-3032300103202301-1010313003031313-1033323103000112-1100031003130131-3220233332303302-3233132331221030-0023303011133320-1123131030313032"></a>

<a id="canonical-3102030100200321-1132231112112000-3313200303310323-0303202010002211-1301223121120303-0132003323002101-2021200202111022-3110000333022121"></a>

## name property — cluster_role_bindings / 303223020203 / 4

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

<a id="canonical-3303131000111202-2022222002221332-0221100001021201-3333021311210331-0112222201130333-0132232112001322-1300033132332022-3213022123210303"></a>

<a id="canonical-0130232023333000-0232313111300301-2023302221103012-0212230013301310-0020200021020300-1003021010001112-3102321022021022-2222033112321231"></a>

## namespace property — cluster_role_bindings / 303223020203 / 5

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

<a id="canonical-1021202311112200-3211003211023131-2331303021101310-2230210131330013-2321003133111221-3000013122101000-0030333310112210-1303213210311002"></a>

<a id="canonical-2320120223211110-1121300133133221-0101132200303303-1220012210011223-3123110223103322-2131020132213231-2130222311011133-1133131031211320"></a>

## tenant property — cluster_role_bindings / 303223020203 / 6

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

<a id="canonical-2130222331013003-3110211032213220-0022121112030123-3231121333130202-1212203310013312-1210222010023310-1310121231231013-0321122323303120"></a>

## Next pages — cluster_role_bindings / 303223020203 / 7

- [use_custom_cluster_role_bindings](data-sources--k8s_cluster--reference--group-001.md#canonical-3321211133132301-2331122022223321-2012231113301103-1020230201103020-0302233321331011-2321020130101031-2012122130121322-3131100020033332)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-2210312112133200-2120022110103011-3332312223021001-2132020112232033-3130120230202301-1002132302112011-3330122112333332-2312202300223020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033330030203321-3132103231301033-3012312001013332-2230030303333010-3202212231012333-2231102200133302-3330111313213103-2000222321222302"></a>

## use_custom_cluster_role_list — use_custom_cluster_role_list / 322102301003 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- use_custom_cluster_role_list

<a id="canonical-2212233112021010-3231210323301021-2130313323112021-1013021202001312-0301011111313112-3022322211303101-0032130300300203-2221101013332331"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_cluster\_role\_list, use\_default\_cluster\_roles; Default:
use\_default\_cluster\_roles\] List of active cluster role list for a K8s cluster.

Upstream description:

List of active cluster role list for a K8s cluster.

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

- [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2212233112021010-3231210323301021-2130313323112021-1013021202001312-0301011111313112-3022322211303101-0032130300300203-2221101013332331)
- [use_default_cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-1213301333310022-1000030000313220-1202033333103312-1100130121131203-3232103111112031-0330001231312102-2012021202023112-2032201223220012)

Select alternatives according to the provider validators above.

<a id="canonical-1322132220311001-3110320020131313-2303030102120030-0311102210031202-2232300213321313-1211203123103223-0032211211323203-2111001310322122"></a>

## Direct properties — use_custom_cluster_role_list / 322102301003 / 3

- [cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-1103213301200022-0312322121221201-3231133331003003-3121311020231031-1023110320031120-3311323213121102-1332301300023001-3133002330221330): complete subsection reference.

<a id="canonical-3002002112003023-1302132230011212-0212000210332102-2320311330331103-1302030331301333-2100223132200122-2203212021330021-3322321322210013"></a>

## Next pages — use_custom_cluster_role_list / 322102301003 / 4

- [use_custom_cluster_role_list.cluster_roles](data-sources--k8s_cluster--reference--group-001.md#canonical-1103213301200022-0312322121221201-3231133331003003-3121311020231031-1023110320031120-3311323213121102-1332301300023001-3133002330221330)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-1103213301200022-0312322121221201-3231133331003003-3121311020231031-1023110320031120-3311323213121102-1332301300023001-3133002330221330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330022121133313-2313111110112131-0331012320203201-3001200301201220-3112332323112233-1132303231230013-1102330003201100-1233301211211020"></a>

## use_custom_cluster_role_list.cluster_roles — cluster_roles / 101131120131 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2210312112133200-2120022110103011-3332312223021001-2132020112232033-3130120230202301-1002132302112011-3330122112333332-2312202300223020)
- use_custom_cluster_role_list.cluster_roles

<a id="canonical-3311320011033310-1113323113323211-2323213202110232-2000023311101133-0230200332321003-0120000012202033-0013111112332122-0102311122011112"></a>

Type: `"list"`. Computed.

List of active cluster role list for a K8s cluster.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0012323101021230-3012110012231011-2030221310332221-1211211233103112-2021310000032223-2313110110221300-3211000022101130-1333231202202200"></a>

## Direct properties — cluster_roles / 101131120131 / 3

<a id="canonical-3323131202330000-1232013320013111-0113100210320012-3213302002230030-3203113332131230-1003332012320010-3022302022321302-1323002002030031"></a>

<a id="canonical-2332003223330031-1302011021213202-0313113121131320-3003113101200133-3103133213322102-1120211221013012-2331003232210321-1010131302213200"></a>

## name property — cluster_roles / 101131120131 / 4

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

<a id="canonical-1131022012033031-2330322002213202-3010233030021111-2012201022010222-1131322323223002-1303130021310101-3002202112113131-1000303002033200"></a>

<a id="canonical-0000220103002002-3003120212030031-0111100331102133-0331131101013333-3320132332023030-3100122333103021-3312213031112320-3300333111010202"></a>

## namespace property — cluster_roles / 101131120131 / 5

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

<a id="canonical-1023303122113212-2312222223110330-2300023120323120-2122132121202020-0111132332232012-3102122201000320-1102102313331331-2331121202301232"></a>

<a id="canonical-3030212233211220-1022331220102123-2021023022333323-0031110300221121-2100203220213210-2013301330310111-0121321033103330-0110321331001022"></a>

## tenant property — cluster_roles / 101131120131 / 6

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

<a id="canonical-2122203122130032-1031210212202320-2300100330011222-0203002332231312-2233313132002202-1100000032231210-0201032123323131-3232021112303232"></a>

## Next pages — cluster_roles / 101131120131 / 7

- [use_custom_cluster_role_list](data-sources--k8s_cluster--reference--group-001.md#canonical-2210312112133200-2120022110103011-3332312223021001-2132020112232033-3130120230202301-1002132302112011-3330122112333332-2312202300223020)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-3012230023323023-1200010011232313-2031133231312300-1101121323331103-0310202311002101-0323110101101321-1000130001232133-0102210223300320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123102001100000-2312233201102020-3010110012320301-3321313333032111-1213211230213021-1213212332210131-1333001122030300-0313110102323213"></a>

## use_custom_pod_security_admission — use_custom_pod_security_admission / 020101013130 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- use_custom_pod_security_admission

<a id="canonical-3013300322112212-2002321022121312-2123123230010000-2332313201102010-0320223012321001-0120311003102333-1022232320132131-3022001003103130"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_pod\_security\_admission, use\_default\_pod\_security\_admission; Default:
use\_default\_pod\_security\_admission\] Type establishes a direct reference from one object(the
referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [use_custom_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-3013300322112212-2002321022121312-2123123230010000-2332313201102010-0320223012321001-0120311003102333-1022232320132131-3022001003103130)
- [use_default_pod_security_admission](data-sources--k8s_cluster--reference--group-001.md#canonical-2031132023103022-1303010320132012-3232033320320020-1030330320032210-0320100012001333-0110112323113031-0122132321220130-0203330302013200)

Select alternatives according to the provider validators above.

<a id="canonical-1103013121102110-2321202032221320-3010131212033110-1311112030033031-1031011220033313-1012300021312221-3021303322231012-2211033230121231"></a>

## Direct properties — use_custom_pod_security_admission / 020101013130 / 3

<a id="canonical-3332132322012021-0231321330232000-1232331132130020-0220123101011000-0233013211323100-0002300330011133-0030312133322021-0103122223303132"></a>

<a id="canonical-1320011121131313-2222303013120213-2121303101323111-1133111123000220-1220222202233102-3021003310311211-2201102211031001-3010322213110201"></a>

## name property — use_custom_pod_security_admission / 020101013130 / 4

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

<a id="canonical-0300112133333000-1002122213031200-1101022233102333-1301132021231133-0123133232010003-2323212023230130-1001102210023220-1222302203220203"></a>

<a id="canonical-2112220011311231-2032231320320121-1330222211010220-3321020030303201-1231221232233000-0231330200112212-2211312311202132-0302110011201201"></a>

## namespace property — use_custom_pod_security_admission / 020101013130 / 5

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

<a id="canonical-0202113122130100-3130022320200020-0023203330303100-3112100222133120-2300112212033010-3323330221321332-2321313211311111-3203001331222200"></a>

<a id="canonical-0222313321110103-2211031002300210-1002001001113201-2120021331301323-1012331131200113-2323032313223032-0211320231231021-0021322323132123"></a>

## tenant property — use_custom_pod_security_admission / 020101013130 / 6

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

<a id="canonical-0133020312322110-2000333110322202-0323303213200100-0323210322121031-1123021232232132-1112300210000000-3130320130230202-2321103311102020"></a>

## Next pages — use_custom_pod_security_admission / 020101013130 / 7

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-0222103111213102-1022030213200111-2301201033101101-2013300120212213-2320221103200012-2233132210011313-3020211000121131-1213103021100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101103020201220-2323230302210202-3101231331230310-1133231213332311-0001222132111030-1333202121121113-2311020012132222-2231331223001223"></a>

## use_custom_psp_list — use_custom_psp_list / 002203221022 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- use_custom_psp_list

<a id="canonical-3311023210120220-0322222023211012-2210300213333202-2103102103012311-0131232303212231-1012032300133021-0120201301223312-0302231101310102"></a>

Type: `"single"`. Computed.

\[OneOf: use\_custom\_psp\_list, use\_default\_psp; Default: use\_default\_psp\] List of active Pod
security policies for a K8s cluster.

Upstream description:

List of active Pod security policies for a K8s cluster.

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

- [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-3311023210120220-0322222023211012-2210300213333202-2103102103012311-0131232303212231-1012032300133021-0120201301223312-0302231101310102)
- [use_default_psp](data-sources--k8s_cluster--reference--group-001.md#canonical-2010021332111032-2131200311301021-3212213000010311-2023031202320012-3322010112023233-1312033102121133-0130230033031012-3021222132100320)

Select alternatives according to the provider validators above.

<a id="canonical-3211300132001112-2012010103303103-1013031000302110-0320012120201200-2111102133111101-0132111211233320-1021011223110301-3100020101133013"></a>

## Direct properties — use_custom_psp_list / 002203221022 / 3

- [pod_security_policies](data-sources--k8s_cluster--reference--group-001.md#canonical-3233222111231321-2021102131232003-2310001031312033-3010123111010233-1221330101122310-0202102220130102-3010110332210330-0030012101010123): complete subsection reference.

<a id="canonical-0231312003333332-1110103221212221-3302201330311302-3103013010320130-0303013311301211-2221210312000301-1322312022303300-0121302323130301"></a>

## Next pages — use_custom_psp_list / 002203221022 / 4

- [use_custom_psp_list.pod_security_policies](data-sources--k8s_cluster--reference--group-001.md#canonical-3233222111231321-2021102131232003-2310001031312033-3010123111010233-1221330101122310-0202102220130102-3010110332210330-0030012101010123)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-3233222111231321-2021102131232003-2310001031312033-3010123111010233-1221330101122310-0202102220130102-3010110332210330-0030012101010123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012130121001200-0121033121021120-3310010133013122-0133020310211021-2330201213211320-0301223033222223-2012223201330221-3212131112003222"></a>

## use_custom_psp_list.pod_security_policies — pod_security_policies / 312311223233 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-0222103111213102-1022030213200111-2301201033101101-2013300120212213-2320221103200012-2233132210011313-3020211000121131-1213103021100330)
- use_custom_psp_list.pod_security_policies

<a id="canonical-0022000111322303-1133321310222211-1312013233021002-2213101331101312-1203220131222303-0032102313020110-3033130110230000-1032323120131033"></a>

Type: `"list"`. Computed.

List of active Pod security policies for a K8s cluster.

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

<a id="canonical-3210020000220312-3331113221001112-3033323010333323-2123231300033131-1322310312202112-2330101023202210-3333133221311110-2112032230310222"></a>

## Direct properties — pod_security_policies / 312311223233 / 3

<a id="canonical-2032231313222230-0032301011211003-2031123011120210-3201023100103302-0133210003332111-0113223012012133-2020333102100300-2332023120222301"></a>

<a id="canonical-1312220031222312-2213222103031111-3231032021113300-1011111221111210-2210131011300212-2110030030311310-3022222312333121-0330200130200322"></a>

## name property — pod_security_policies / 312311223233 / 4

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

<a id="canonical-0123112000103313-0121020102230312-0010013022132032-1033332032200201-2331212200003110-1211222332000023-2003210332320201-3133120010232002"></a>

<a id="canonical-3230223121213310-3301211001011211-0212332210000033-2122122122122231-2100302123322201-3112213030120011-0323313230123121-2122331101110122"></a>

## namespace property — pod_security_policies / 312311223233 / 5

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

<a id="canonical-1033323332000001-2203331321322213-1222010011311221-2002212001322122-3222323201212333-2230000321332220-1020233121101013-3310203110221120"></a>

<a id="canonical-2232301103033003-2112322223212321-3022232010011120-0320230301021201-1201113132303011-1323122033223000-2200103222213231-3011010110032303"></a>

## tenant property — pod_security_policies / 312311223233 / 6

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

<a id="canonical-3300132002112123-1321321022130131-2222033312221003-0032301203022131-2310000220321213-2203122001030201-3011031211330230-3010211301333312"></a>

## Next pages — pod_security_policies / 312311223233 / 7

- [use_custom_psp_list](data-sources--k8s_cluster--reference--group-001.md#canonical-0222103111213102-1022030213200111-2301201033101101-2013300120212213-2320221103200012-2233132210011313-3020211000121131-1213103021100330)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-3033213011211231-1311213031231322-0331302021011032-0000211222303012-1322100223221232-2302122333012033-3131332132222230-0030211100120123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123013312102020-1213202321100231-0311031313230212-0002102121112113-3320030112002002-3320112133213021-1101320300030021-0002033203021110"></a>

## use_default_cluster_role_bindings — use_default_cluster_role_bindings / 130332213301 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- use_default_cluster_role_bindings

<a id="canonical-2233130103302023-1310000211033111-1312311001121102-3221101201103030-3011312112103001-0010032001231222-3103003323120320-2033001210020133"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-3313210123323001-3320131120232132-2312323100101010-2002223123121110-1303121323323222-1020221022031211-0022101322012310-2110222313202233"></a>

## Direct properties — use_default_cluster_role_bindings / 130332213301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032113230330001-0201010001033330-0111202321121123-1112122111022311-0102200313312123-2310011213113333-0222003233021121-1302301012002123"></a>

## Next pages — use_default_cluster_role_bindings / 130332213301 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-1030120320023032-0121111102120012-2122221010210131-3023032202310122-3302021200233321-2123032002300221-3130110013130033-1212033311130030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030223133301102-1303233110130122-3200332111210122-0221322310112203-3231220312130333-3032231032132033-2203230030221110-0033331201002300"></a>

## use_default_cluster_roles — use_default_cluster_roles / 313311200130 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- use_default_cluster_roles

<a id="canonical-1213301333310022-1000030000313220-1202033333103312-1100130121131203-3232103111112031-0330001231312102-2012021202023112-2032201223220012"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-2013102200120023-3230032010221231-2320221332322312-2133331102223233-1221101020323011-1212010312020303-2113023331010331-1331220321200210"></a>

## Direct properties — use_default_cluster_roles / 313311200130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200332121311332-2300132221221002-1232321211323133-2333133333230333-2212201313103222-2002322110210301-3021010203131211-1103101302000301"></a>

## Next pages — use_default_cluster_roles / 313311200130 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-1101101122322220-1120213302132230-3212131032131321-0033133201001221-1011222030321331-3121013100102132-2131330220110032-3001313332022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230232202032113-3113222122330331-2033330201320002-2111322121010033-2000203133311331-0131123321321001-3020122103230031-2012033222022101"></a>

## use_default_pod_security_admission — use_default_pod_security_admission / 310033332202 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- use_default_pod_security_admission

<a id="canonical-2031132023103022-1303010320132012-3232033320320020-1030330320032210-0320100012001333-0110112323113031-0122132321220130-0203330302013200"></a>

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

<a id="canonical-2331220113002132-2133212110021320-3031313300020312-0022002301121003-3330233221233001-0021311033203330-2313023203230313-3211033023033332"></a>

## Direct properties — use_default_pod_security_admission / 310033332202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132033023322132-1032133120210233-2222021300013231-1220302300232021-2232213300101113-3110122110232132-1231033200231331-3211233113201001"></a>

## Next pages — use_default_pod_security_admission / 310033332202 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-0331232133213200-1233022031333312-0230013230013120-1233232013212110-2030212322112002-2103210312230132-1012023100100202-0123100212303233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013000232033123-1321013230332010-0001132102112300-3232322012300000-3222033121323333-1033121021120023-2231323131323132-1010003102220120"></a>

## use_default_psp — use_default_psp / 213011111032 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- use_default_psp

<a id="canonical-2010021332111032-2131200311301021-3212213000010311-2023031202320012-3322010112023233-1312033102121133-0130230033031012-3021222132100320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use default psp. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-2012330010020120-1012103113221211-3022020030331203-3230033301323200-3231113333331100-2303301133302230-3312313131032333-3311300322331321"></a>

## Direct properties — use_default_psp / 213011111032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211133022002001-2301113220220133-2303221020130012-3203202133311322-3323131130311202-0131021332100201-3233033233121212-3122023000222023"></a>

## Next pages — use_default_psp / 213011111032 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-0002333010132133-0022011023203112-3231122111102303-0131030011100223-3233022321012030-3102031213223333-1323301112022032-3031200010100021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011332022121222-0133110333021113-1121123310213221-0033330113031202-0330013120102330-3023211332031102-1223111000230010-3221133111303232"></a>

## vk8s_namespace_access_deny — vk8s_namespace_access_deny / 110013133121 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- vk8s_namespace_access_deny

<a id="canonical-3030000233112133-3111323101120112-0310332201210202-1333321322332300-2312113230031011-3123002111031103-2201201021112323-2201013201223130"></a>

Type: `["object", {}]`. Computed.

\[OneOf: vk8s\_namespace\_access\_deny, vk8s\_namespace\_access\_permit\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [vk8s_namespace_access_deny](data-sources--k8s_cluster--reference--group-001.md#canonical-3030000233112133-3111323101120112-0310332201210202-1333321322332300-2312113230031011-3123002111031103-2201201021112323-2201013201223130)
- [vk8s_namespace_access_permit](data-sources--k8s_cluster--reference--group-001.md#canonical-0000323102132300-3112022320221332-1330003121122121-0211311010232213-2133222312020033-2110230010012021-1122102310120123-1113200220303033)

Select alternatives according to the provider validators above.

<a id="canonical-3023011303330002-2120022021233002-1122131332323313-0001202230213200-2223000020101212-1021103033322211-1123201013320323-1231230010330100"></a>

## Direct properties — vk8s_namespace_access_deny / 110013133121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202331312102200-2210301300303311-2133000132021102-1100122130021300-3222310033313203-0120211121330110-2011232233111021-2333202111130032"></a>

## Next pages — vk8s_namespace_access_deny / 110013133121 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)

<a id="canonical-0030031301002303-0032103230211033-3333313322321331-3233011120002200-2020222121001033-3120301113302031-0013121112201100-2131301013001320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303201110230122-0002133311331310-2013221021003101-1123112103203330-0031333130230110-2201102233311110-2220032021023112-0123223112003210"></a>

## vk8s_namespace_access_permit — vk8s_namespace_access_permit / 110100213323 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- vk8s_namespace_access_permit

<a id="canonical-0000323102132300-3112022320221332-1330003121122121-0211311010232213-2133222312020033-2110230010012021-1122102310120123-1113200220303033"></a>

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

<a id="canonical-3121333321133132-0302221313231222-1332202113130100-0321002330202222-3122230211221221-2131113110123102-1022231203332302-0333330130231031"></a>

## Direct properties — vk8s_namespace_access_permit / 110100213323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231021000323013-0210211321223330-3031202312211001-1232130330001021-3200203332132220-2112122222231232-1221320322000120-3202001222212033"></a>

## Next pages — vk8s_namespace_access_permit / 110100213323 / 4

- [Property reference](data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
