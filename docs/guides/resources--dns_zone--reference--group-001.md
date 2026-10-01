---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231031313001232-2133233000012101-2310100312331322-2100231311001223-0011030020023311-2010002012133012-0001021302311300-3101021032323121"></a>

## Property reference — Property reference / 312030201022 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- Property reference

<a id="canonical-0330332023011331-2321002130323131-1122013032110222-0020000212220322-2202020220231023-0332032200012022-0120330130102200-1321323031322302"></a>

## Direct properties — Property reference / 312030201022 / 3

<a id="canonical-3022331302023203-3113031130203013-1233330202213031-2113331122233122-2120201313121001-2210302301113330-3102100030202022-2101001123002001"></a>

<a id="canonical-2321131101110301-0021120123202013-2232200111310132-2132202322201221-0202201201022203-0023102023311010-0131012303230332-2212010133320231"></a>

## annotations property — Property reference / 312030201022 / 4

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

<a id="canonical-1130132100211301-2202100213321022-3031213211012310-1332002232302000-3200020011231310-1110320312211132-3003211013123301-2022111022130202"></a>

<a id="canonical-0121020303200231-1323023102133200-1200113112231330-1213030213201200-0213331110120212-2311112312213031-3310000022311210-3123003213201130"></a>

## description property — Property reference / 312030201022 / 5

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

<a id="canonical-0202122222230130-3102312033321220-1002230010222030-3232032121302013-1101303200031322-3303331010000130-3100301200120232-3331132112002013"></a>

<a id="canonical-1101030313201232-2313233112212312-3320210202022003-2233223330032201-3213021302213212-1101101133113303-1012110101032123-0022133321033212"></a>

## disable property — Property reference / 312030201022 / 6

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

<a id="canonical-0021312232200332-3133103020333222-3333000023233112-3030123033220212-3311131102033102-3231211202321112-0320033000332201-3122203203311102"></a>

<a id="canonical-3021022321012031-2011213133203011-0331023200101030-1202022000213312-3301221303002202-3122303112313323-2313022100022013-3121310312002231"></a>

## ID property — Property reference / 312030201022 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0310021303113101-0121102012311221-0213313303023123-0323332322100131-0230300230030102-3300020010123223-3013113231002030-3103211123013210"></a>

<a id="canonical-0201022102313111-0023001023111232-0322302011000110-0310022010222121-0130020020002230-1220022011222031-3030003110131121-1131202000012013"></a>

## labels property — Property reference / 312030201022 / 8

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

<a id="canonical-0123300033112300-2333023023112121-2302003303321333-0011323002021203-0001222123211013-1122011133331331-3210213032120121-0132113031011120"></a>

<a id="canonical-1101103133023312-1213022332120303-3123002202301233-3031022202102210-0202200332210123-0000111323300200-0201012022100120-2221220021233311"></a>

## name property — Property reference / 312030201022 / 9

Type: `"string"`. Required.

Domain name for the DNS Zone (e.g., example.com). Must be a valid DNS domain name.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.DomainValidator(),
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

<a id="canonical-1032122023002133-1101103320111330-1212032121122012-0123102123300013-1102010223033300-0323212120230132-0303202223102323-2201323102101333"></a>

<a id="canonical-1001210011230323-0232302110020232-0331132330333330-2101102110332302-3121333110000322-2023130332011102-0221332100102320-0103311003120300"></a>

## namespace property — Property reference / 312030201022 / 10

Type: `"string"`. Optional, Computed.

Namespace for the DNS Zone. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312): complete subsection reference.

- [secondary](resources--dns_zone--reference--group-003.md#canonical-2101113233122202-1230223131033031-2301310003322033-0133201323022132-1232110221020011-1332203122201231-3233232102003032-1103322203113021): complete subsection reference.

- [timeouts](resources--dns_zone--reference--group-003.md#canonical-1121023120323203-0300301033303311-3210213210020103-1313032000222330-2030312122211110-2210030103233231-0231221302231331-3010100002112333): complete subsection reference.

<a id="canonical-2001122332211031-2101203013333012-0320123333203010-2123320303332123-0002323213031222-1033221102102220-2103110212013113-1111122102232111"></a>

## All schema paths — Property reference / 312030201022 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dns_zone--reference--group-001.md#canonical-3022331302023203-3113031130203013-1233330202213031-2113331122233122-2120201313121001-2210302301113330-3102100030202022-2101001123002001) |
| `description` | [description](resources--dns_zone--reference--group-001.md#canonical-1130132100211301-2202100213321022-3031213211012310-1332002232302000-3200020011231310-1110320312211132-3003211013123301-2022111022130202) |
| `disable` | [disable](resources--dns_zone--reference--group-001.md#canonical-0202122222230130-3102312033321220-1002230010222030-3232032121302013-1101303200031322-3303331010000130-3100301200120232-3331132112002013) |
| `id` | [id](resources--dns_zone--reference--group-001.md#canonical-0021312232200332-3133103020333222-3333000023233112-3030123033220212-3311131102033102-3231211202321112-0320033000332201-3122203203311102) |
| `labels` | [labels](resources--dns_zone--reference--group-001.md#canonical-0310021303113101-0121102012311221-0213313303023123-0323332322100131-0230300230030102-3300020010123223-3013113231002030-3103211123013210) |
| `name` | [name](resources--dns_zone--reference--group-001.md#canonical-0123300033112300-2333023023112121-2302003303321333-0011323002021203-0001222123211013-1122011133331331-3210213032120121-0132113031011120) |
| `namespace` | [namespace](resources--dns_zone--reference--group-001.md#canonical-1032122023002133-1101103320111330-1212032121122012-0123102123300013-1102010223033300-0323212120230132-0303202223102323-2201323102101333) |
| `primary` | [primary](resources--dns_zone--reference--group-001.md#canonical-3300313313010111-0200031323101110-3311333102220103-0321130221302332-0222220031013301-1130233223202231-2200013222231201-1332310110100311) |
| `primary.allow_http_lb_managed_records` | [primary.allow_http_lb_managed_records](resources--dns_zone--reference--group-001.md#canonical-1120201111333222-0103222323322112-3012012212033121-2122223222121000-2202311002021122-3110102220111302-3001123312013103-1113202111210201) |
| `primary.default_rr_set_group` | [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-0212100320002333-0231332102121022-3122130003121010-2033020331213210-2100101301032130-1230002200300200-2122133210202203-2300130303131132) |
| `primary.default_rr_set_group.a_record` | [primary.default_rr_set_group.a_record](resources--dns_zone--reference--group-001.md#canonical-2330110111222012-3013233111132012-3010113113203222-1103221331220101-2131101120310020-3032311210302120-2302131010023023-3130112130330322) |
| `primary.default_rr_set_group.a_record.name` | [primary.default_rr_set_group.a_record.name](resources--dns_zone--reference--group-001.md#canonical-2120011132310033-2230100110101001-2122000133220203-2121110331330003-1030221322311103-0123110011233311-1211223133311103-3122111231302221) |
| `primary.default_rr_set_group.a_record.values` | [primary.default_rr_set_group.a_record.values](resources--dns_zone--reference--group-001.md#canonical-3330122023302310-2010022113313222-0333001310231331-1013132210032303-3020213031131001-2010231112123101-3022200331122220-2102111311232133) |
| `primary.default_rr_set_group.aaaa_record` | [primary.default_rr_set_group.aaaa_record](resources--dns_zone--reference--group-001.md#canonical-0232201120223002-2232233203100032-0031201302003112-0113012310031201-3232021031123113-1311033232013201-2232113012111301-2100032213311333) |
| `primary.default_rr_set_group.aaaa_record.name` | [primary.default_rr_set_group.aaaa_record.name](resources--dns_zone--reference--group-001.md#canonical-1220302033330220-1123320310213220-0311320300102313-1231222003123230-1210110333032033-3103101233302302-3130232212010003-2303002123232110) |
| `primary.default_rr_set_group.aaaa_record.values` | [primary.default_rr_set_group.aaaa_record.values](resources--dns_zone--reference--group-001.md#canonical-3201100002112022-1312232123211001-2001311111233032-2130110131013121-2202303333133002-3100122322332201-0300210320300122-2300201321121212) |
| `primary.default_rr_set_group.afsdb_record` | [primary.default_rr_set_group.afsdb_record](resources--dns_zone--reference--group-001.md#canonical-3332112233032323-1130123121203113-1300031201211222-2002330110323032-3331330112333210-0103212313223320-0301331111112020-2133313203211310) |
| `primary.default_rr_set_group.afsdb_record.name` | [primary.default_rr_set_group.afsdb_record.name](resources--dns_zone--reference--group-001.md#canonical-1210121323331023-2323212101132211-0110201001203313-3021231310320300-3021313131333301-1211330020200301-0020120211131023-0200312113233002) |
| `primary.default_rr_set_group.afsdb_record.values` | [primary.default_rr_set_group.afsdb_record.values](resources--dns_zone--reference--group-001.md#canonical-1110313211030132-1300320202200032-1011111132200023-1002020130222000-0003230203131023-1131230011212010-2321203100203020-0323303203102332) |
| `primary.default_rr_set_group.afsdb_record.values.hostname` | [primary.default_rr_set_group.afsdb_record.values.hostname](resources--dns_zone--reference--group-001.md#canonical-1211202222013210-3320211100213333-1331010123103222-0301000210111330-1222033333013100-2300110213003213-3300102231311233-2321231012312021) |
| `primary.default_rr_set_group.afsdb_record.values.subtype` | [primary.default_rr_set_group.afsdb_record.values.subtype](resources--dns_zone--reference--group-001.md#canonical-3132203033333121-3011120111133323-2200233101013102-0321121000211031-0232222002003220-1203012101021112-0002000132232333-2332133301331130) |
| `primary.default_rr_set_group.alias_record` | [primary.default_rr_set_group.alias_record](resources--dns_zone--reference--group-001.md#canonical-1122033020033331-3230212110322231-3312001030101011-3200131120200223-3202221221303020-3313132200130032-1103130221101212-0223020320123220) |
| `primary.default_rr_set_group.alias_record.value` | [primary.default_rr_set_group.alias_record.value](resources--dns_zone--reference--group-001.md#canonical-2112031333313310-1232102031201202-3102110310112321-1112231000120111-2000022313320022-1233011103223020-3211311023020031-3220110003312033) |
| `primary.default_rr_set_group.caa_record` | [primary.default_rr_set_group.caa_record](resources--dns_zone--reference--group-001.md#canonical-2320311010310232-2013131210222123-0230210012200013-0300323011131132-0010001331133232-1202021103121111-0201123020120231-0123001113221122) |
| `primary.default_rr_set_group.caa_record.name` | [primary.default_rr_set_group.caa_record.name](resources--dns_zone--reference--group-001.md#canonical-2313132303023023-3322003133323130-2331231301210210-0010100330122212-0222130212310031-3110332333011320-3321222003131223-3111112112313320) |
| `primary.default_rr_set_group.caa_record.values` | [primary.default_rr_set_group.caa_record.values](resources--dns_zone--reference--group-001.md#canonical-1210332311131222-2021220113323120-3202100220120002-2233110321013220-3033331030302103-1133300212103133-2322211103203010-3110011200210233) |
| `primary.default_rr_set_group.caa_record.values.flags` | [primary.default_rr_set_group.caa_record.values.flags](resources--dns_zone--reference--group-001.md#canonical-1133223100312020-2320121020310102-0313223122123123-0020313302312221-3222112222330331-0002300331301321-3332002200123310-1120221002311213) |
| `primary.default_rr_set_group.caa_record.values.tag` | [primary.default_rr_set_group.caa_record.values.tag](resources--dns_zone--reference--group-001.md#canonical-2132230321323312-2011201212233300-3301331300000023-2011122311300122-2120333123102112-3331011322030011-1211210221203011-1102013112130033) |
| `primary.default_rr_set_group.caa_record.values.value` | [primary.default_rr_set_group.caa_record.values.value](resources--dns_zone--reference--group-001.md#canonical-3111300130213212-1021201101023232-2101213201122101-2230112133002022-2020200102013031-1231032330132102-3233131332310021-2211032123212301) |
| `primary.default_rr_set_group.cds_record` | [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-0013211201021311-1311001010312112-2213333033202001-0132320223103001-3320233031011102-1301030001320120-3133132113100313-0320120323002313) |
| `primary.default_rr_set_group.cds_record.name` | [primary.default_rr_set_group.cds_record.name](resources--dns_zone--reference--group-001.md#canonical-2102220020323133-3000321130101002-1113002303030223-2223103122032120-3020202031231231-0011331313010322-2230032021213313-3123210013320201) |
| `primary.default_rr_set_group.cds_record.values` | [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-3020032231002122-1312133211211100-3201011133122012-1333000002231210-0112232302132132-0332331112301123-0020013313330321-3132222322321121) |
| `primary.default_rr_set_group.cds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.cds_record.values.ds_key_algorithm](resources--dns_zone--reference--group-001.md#canonical-3303333101110033-2133132031220220-1103213220100010-3233321021320112-3110123111313201-1321331211201202-3030123321030330-3120130310311122) |
| `primary.default_rr_set_group.cds_record.values.key_tag` | [primary.default_rr_set_group.cds_record.values.key_tag](resources--dns_zone--reference--group-001.md#canonical-1101300032130123-0333302133001233-1211230103001231-1033110300002301-0321312203020233-1311302233222010-2131310020130200-1221203221330320) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest](resources--dns_zone--reference--group-001.md#canonical-0321301303101121-2202032133032021-2310311222213113-3213302033300013-0312221013010200-0223312322201023-0312312002222130-3121013031121312) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest.digest](resources--dns_zone--reference--group-001.md#canonical-2321031021121300-2111210300001220-3201001133102223-0030230323002333-0013312102232010-1211010000313322-3103112111123022-0132101030200332) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest](resources--dns_zone--reference--group-001.md#canonical-3333311301201110-3031312030313020-1302301100210210-3320102220331201-2212322323201120-1031132102223001-1200110233221332-2023232220201011) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest.digest](resources--dns_zone--reference--group-001.md#canonical-0321011012130120-3122133113310101-3113101321212133-2102220021233010-3220231130133301-1120023030023123-0220110102032032-1212101333203222) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest](resources--dns_zone--reference--group-001.md#canonical-3013202332103313-0001221023330013-3220031300112200-0123033110121331-3011333120311230-0332102313313330-1331030121100332-2330310212302210) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest.digest](resources--dns_zone--reference--group-001.md#canonical-0310320203011011-1112012112001100-2012321100210310-0101231312300223-0020311003312203-2201023002110010-1203112000132210-2230323313221312) |
| `primary.default_rr_set_group.cert_record` | [primary.default_rr_set_group.cert_record](resources--dns_zone--reference--group-001.md#canonical-0311231110333010-2022331022011001-2200111021010022-0232133323223133-0231120022221010-3130133321311100-0133000100133310-0111111332003203) |
| `primary.default_rr_set_group.cert_record.name` | [primary.default_rr_set_group.cert_record.name](resources--dns_zone--reference--group-001.md#canonical-2230112122200022-2302233201001101-1011222133323100-0133011011010000-2031302323113013-2311102013313311-3213123233233300-2003120332022020) |
| `primary.default_rr_set_group.cert_record.values` | [primary.default_rr_set_group.cert_record.values](resources--dns_zone--reference--group-001.md#canonical-1200312002202202-0233012102010301-1232322223110133-0031013322033021-2132113022022101-1010323300111113-2222111121232022-2101003021033031) |
| `primary.default_rr_set_group.cert_record.values.algorithm` | [primary.default_rr_set_group.cert_record.values.algorithm](resources--dns_zone--reference--group-001.md#canonical-2122302332000202-1310102112310330-2230123231213121-1011222033320110-1132011323002110-3022303112123120-1311030220313322-0230112220023301) |
| `primary.default_rr_set_group.cert_record.values.cert_key_tag` | [primary.default_rr_set_group.cert_record.values.cert_key_tag](resources--dns_zone--reference--group-001.md#canonical-0010131002132313-1133322112133000-1221131222322212-1111323023000001-3012320011223213-2312003231001221-0201331030032231-1000020332222223) |
| `primary.default_rr_set_group.cert_record.values.cert_type` | [primary.default_rr_set_group.cert_record.values.cert_type](resources--dns_zone--reference--group-001.md#canonical-3022210110200323-1130212222200321-2310333303013323-0321211101333120-1302020330030333-2103012033211112-3030201302012022-3113022202101332) |
| `primary.default_rr_set_group.cert_record.values.certificate` | [primary.default_rr_set_group.cert_record.values.certificate](resources--dns_zone--reference--group-001.md#canonical-1130210313300023-2031001213111211-3112332303321012-3202212202003113-2121112231221202-2330232201132301-3200211233020032-0212320311300301) |
| `primary.default_rr_set_group.cname_record` | [primary.default_rr_set_group.cname_record](resources--dns_zone--reference--group-002.md#canonical-3213013112112213-2233222323123300-1300131223222322-3231230120300333-1003233213112103-1020331020111311-0120221020321123-3203313311231200) |
| `primary.default_rr_set_group.cname_record.name` | [primary.default_rr_set_group.cname_record.name](resources--dns_zone--reference--group-002.md#canonical-0133022313203033-2203213132101313-2012320213100012-3132220123330222-2000312122221010-0120121301001232-0112112212302123-2202213310202020) |
| `primary.default_rr_set_group.cname_record.value` | [primary.default_rr_set_group.cname_record.value](resources--dns_zone--reference--group-002.md#canonical-0031110333131103-0321200112203030-1131333320211332-3100303112030300-0313332122231113-0203112203112102-0223101233133010-3021121003102131) |
| `primary.default_rr_set_group.description_spec` | [primary.default_rr_set_group.description_spec](resources--dns_zone--reference--group-001.md#canonical-2303233302011210-2010121211233220-0230012323323223-1210320120031121-2331230210002213-3102022030102122-1023300200231221-2021222333200031) |
| `primary.default_rr_set_group.ds_record` | [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-002.md#canonical-2122200203213103-3312233120033010-1121300202323310-2012331030031320-0202011320020000-1031332320320322-1131032103120030-2100310212112121) |
| `primary.default_rr_set_group.ds_record.name` | [primary.default_rr_set_group.ds_record.name](resources--dns_zone--reference--group-002.md#canonical-0003003332132021-1232010212023013-1202202122212203-3321002101310211-2032030102010333-2332022330123010-0100200233021103-2101112313313201) |
| `primary.default_rr_set_group.ds_record.values` | [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-3203301320223220-3031220203300100-3132031200201333-2121030223322223-0012323033230132-2010022022203233-3000201203121101-0300300331031312) |
| `primary.default_rr_set_group.ds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.ds_record.values.ds_key_algorithm](resources--dns_zone--reference--group-002.md#canonical-3130200110322321-1033133210111000-3121132222003312-1310323332002003-3120112332202223-0320231203013310-0211302003333320-1312212222002301) |
| `primary.default_rr_set_group.ds_record.values.key_tag` | [primary.default_rr_set_group.ds_record.values.key_tag](resources--dns_zone--reference--group-002.md#canonical-2300310211211003-1123233032120003-3023032130011130-1301020002001301-0101122302120000-2113101202322213-1012000203212322-2213132110133321) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest](resources--dns_zone--reference--group-002.md#canonical-1212003013102330-2302112011102223-2031011033332313-0120232331313110-1011332213222321-2213313320031000-1112130122030203-0001301012233012) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest.digest](resources--dns_zone--reference--group-002.md#canonical-1133310021010122-1303322332310212-0331233303112333-1103321332321221-1003103332133312-3321102232021100-0010110123312033-1001322030132132) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest](resources--dns_zone--reference--group-002.md#canonical-2321133331330211-2300110231313303-1213311223110312-2000323201331311-1022333011031103-2112113320300001-1103323000212033-2331220200320120) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest.digest](resources--dns_zone--reference--group-002.md#canonical-3031003211221223-3302222231033331-3102021232010220-0303101130112303-1231231011111111-1333333323232331-3201222212222220-1131201122001211) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest](resources--dns_zone--reference--group-002.md#canonical-2030122233123310-0212122201203030-0330300330233113-1012010311212221-1032012133301211-2322013320000303-1330230301312032-2122102201210333) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest.digest](resources--dns_zone--reference--group-002.md#canonical-1101302331200220-2221110300012310-1122321000233133-3321121322212313-0110323200132312-2302203101220021-2323032100233233-0223111332330230) |
| `primary.default_rr_set_group.eui48_record` | [primary.default_rr_set_group.eui48_record](resources--dns_zone--reference--group-002.md#canonical-1231213302032020-1211213223312233-0233303332122300-3302330213013023-3010311130120300-1112221013032201-0001032011202133-0231203302203202) |
| `primary.default_rr_set_group.eui48_record.name` | [primary.default_rr_set_group.eui48_record.name](resources--dns_zone--reference--group-002.md#canonical-1012100330302012-0202201221210022-2221103202122300-3130110113030033-3021012203320133-0233232220033102-3310101030113001-0102131212231112) |
| `primary.default_rr_set_group.eui48_record.value` | [primary.default_rr_set_group.eui48_record.value](resources--dns_zone--reference--group-002.md#canonical-3201112200020032-3123032103131101-1033011212121330-1312202003131323-0320123031331112-0020031212003120-3300330303011223-3111022322202313) |
| `primary.default_rr_set_group.eui64_record` | [primary.default_rr_set_group.eui64_record](resources--dns_zone--reference--group-002.md#canonical-3223320120003011-3103320130031032-3311330233013221-1313223103010030-1113212203021121-0022012230030112-1113312132032223-1322223101102132) |
| `primary.default_rr_set_group.eui64_record.name` | [primary.default_rr_set_group.eui64_record.name](resources--dns_zone--reference--group-002.md#canonical-1110123123300000-2102133000131311-2223001311203032-1012132033120023-1113322230311331-3332032332102122-1302103333220210-3311012123110000) |
| `primary.default_rr_set_group.eui64_record.value` | [primary.default_rr_set_group.eui64_record.value](resources--dns_zone--reference--group-002.md#canonical-3030022232113003-1122010013003310-0022001100231103-3021222011200202-0130033322011032-2331223311112000-3231101330333020-1213331123313213) |
| `primary.default_rr_set_group.lb_record` | [primary.default_rr_set_group.lb_record](resources--dns_zone--reference--group-002.md#canonical-0123002111321001-2130320220120210-1321310003303003-2100203031031100-2233221131213003-3032110020220011-1013001123223221-1101033233232310) |
| `primary.default_rr_set_group.lb_record.name` | [primary.default_rr_set_group.lb_record.name](resources--dns_zone--reference--group-002.md#canonical-1202012023213130-2103102332233010-0230120330112132-1212223011111002-2022131301321010-1003300222312330-0103010213100322-1201001121033001) |
| `primary.default_rr_set_group.lb_record.value` | [primary.default_rr_set_group.lb_record.value](resources--dns_zone--reference--group-002.md#canonical-3323100013132323-2203130020313231-2231013033222222-0303012302232332-1033302010331212-3320121211101001-0313002232032000-2111101103302032) |
| `primary.default_rr_set_group.lb_record.value.name` | [primary.default_rr_set_group.lb_record.value.name](resources--dns_zone--reference--group-002.md#canonical-0210222223002130-2230001023300130-0101330120013331-2003001200212201-2330301013202032-1020331323121313-2330111010203310-2330210310031302) |
| `primary.default_rr_set_group.lb_record.value.namespace` | [primary.default_rr_set_group.lb_record.value.namespace](resources--dns_zone--reference--group-002.md#canonical-2131202322332212-3002101101303303-0113312022321123-0022323100111111-3222131232310203-3312003313113011-2211113303301103-1310021021213013) |
| `primary.default_rr_set_group.lb_record.value.tenant` | [primary.default_rr_set_group.lb_record.value.tenant](resources--dns_zone--reference--group-002.md#canonical-2311232221230032-1222033332001211-3213122203121101-1111312321031020-1230020213322220-0203213030132232-2020220311001121-3131322201223123) |
| `primary.default_rr_set_group.loc_record` | [primary.default_rr_set_group.loc_record](resources--dns_zone--reference--group-002.md#canonical-2132111113030021-3321322230122210-2330022320201032-0022111313131111-3112122321323003-3321113201120112-2331221333120033-0232020030113010) |
| `primary.default_rr_set_group.loc_record.name` | [primary.default_rr_set_group.loc_record.name](resources--dns_zone--reference--group-002.md#canonical-0012123332333200-2002113200301200-3333120212310101-1112202231311122-3313130131110103-3032330100233230-3031333110113100-2122330012023202) |
| `primary.default_rr_set_group.loc_record.values` | [primary.default_rr_set_group.loc_record.values](resources--dns_zone--reference--group-002.md#canonical-0000303313130103-3211013021131123-2212112102013332-2313032301323331-2113023333303021-2122000331232010-1033313323030300-0231003010101302) |
| `primary.default_rr_set_group.loc_record.values.altitude` | [primary.default_rr_set_group.loc_record.values.altitude](resources--dns_zone--reference--group-002.md#canonical-2131202112103221-2211310101233311-3013113321130321-3223213123102001-1223001232012103-2111100230303110-2011211320302332-0030000332201231) |
| `primary.default_rr_set_group.loc_record.values.horizontal_precision` | [primary.default_rr_set_group.loc_record.values.horizontal_precision](resources--dns_zone--reference--group-002.md#canonical-0223310311003231-2130112230301222-1200002110121021-3023210233223033-2013112222210132-3110100112113332-0022003203112030-0013132220111321) |
| `primary.default_rr_set_group.loc_record.values.latitude_degree` | [primary.default_rr_set_group.loc_record.values.latitude_degree](resources--dns_zone--reference--group-002.md#canonical-3101333230031331-3203232301103020-0123322112112103-3013313202211231-2002101312323213-2033001121131010-3313131200130212-1330321222233011) |
| `primary.default_rr_set_group.loc_record.values.latitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.latitude_hemisphere](resources--dns_zone--reference--group-002.md#canonical-3010322301212100-3221131321000232-2003311301111310-0231101220002001-1030003102211020-0301203102230320-3221023213233001-1002210230231013) |
| `primary.default_rr_set_group.loc_record.values.latitude_minute` | [primary.default_rr_set_group.loc_record.values.latitude_minute](resources--dns_zone--reference--group-002.md#canonical-3003200010121011-3122020100020232-3120300303303100-1300330210031023-1101122321033011-3320131033011310-1113300233220011-1130013003311130) |
| `primary.default_rr_set_group.loc_record.values.latitude_second` | [primary.default_rr_set_group.loc_record.values.latitude_second](resources--dns_zone--reference--group-002.md#canonical-1132311121131210-0032113300321100-2230002011302100-1312130031130022-0023131012023001-3011231130111301-1310210111302010-1112100100223123) |
| `primary.default_rr_set_group.loc_record.values.location_diameter` | [primary.default_rr_set_group.loc_record.values.location_diameter](resources--dns_zone--reference--group-002.md#canonical-2001302000003023-0110113301110220-3022310200322013-1322311000011013-2203323331133222-1332122010032033-1022302221123222-2310221233123231) |
| `primary.default_rr_set_group.loc_record.values.longitude_degree` | [primary.default_rr_set_group.loc_record.values.longitude_degree](resources--dns_zone--reference--group-002.md#canonical-0123032002122113-1101033300112321-3102333111012213-3111120331331330-0211113310100321-1201301111002221-2113201211123032-2231310312113203) |
| `primary.default_rr_set_group.loc_record.values.longitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.longitude_hemisphere](resources--dns_zone--reference--group-002.md#canonical-1033200313203132-3302301021222000-2012320320112210-3201220231010301-0013323120103203-1232023331020331-2301011211131303-2013000022210032) |
| `primary.default_rr_set_group.loc_record.values.longitude_minute` | [primary.default_rr_set_group.loc_record.values.longitude_minute](resources--dns_zone--reference--group-002.md#canonical-3101310132203301-1232323301110103-1023221213131013-0330301121231202-1033222223012023-0221300103211313-3200213130320330-0201323023322231) |
| `primary.default_rr_set_group.loc_record.values.longitude_second` | [primary.default_rr_set_group.loc_record.values.longitude_second](resources--dns_zone--reference--group-002.md#canonical-2023230311100231-3322110003113122-1033331333231311-0332023131230210-2233133003101311-2033101222223113-3110013203002202-0322120201123130) |
| `primary.default_rr_set_group.loc_record.values.vertical_precision` | [primary.default_rr_set_group.loc_record.values.vertical_precision](resources--dns_zone--reference--group-002.md#canonical-2303210211122011-3132033003000301-0102311002020001-0200122221213003-1021221132220102-1322301213122313-2220221303100310-1220013111321300) |
| `primary.default_rr_set_group.mx_record` | [primary.default_rr_set_group.mx_record](resources--dns_zone--reference--group-002.md#canonical-3031022012121010-1033000003212220-1333213133232001-2301203330223011-1202332101303120-0230200230301011-1311002130031131-0113212112300021) |
| `primary.default_rr_set_group.mx_record.name` | [primary.default_rr_set_group.mx_record.name](resources--dns_zone--reference--group-002.md#canonical-1321002202112310-2312002321302110-1130310200333323-3211330321331112-1031301011303103-1103133023220232-0302202132003202-0230330012010013) |
| `primary.default_rr_set_group.mx_record.values` | [primary.default_rr_set_group.mx_record.values](resources--dns_zone--reference--group-002.md#canonical-0003030310133033-1301312201120100-1102301001211011-1101203121013200-3232121203223313-3200201303223333-3212301101202311-2110102332330230) |
| `primary.default_rr_set_group.mx_record.values.domain` | [primary.default_rr_set_group.mx_record.values.domain](resources--dns_zone--reference--group-002.md#canonical-3212223332222301-2111210213031321-1302011002231203-2203220223311120-2233203000210201-2221000222013101-0013120102130102-2020000013302301) |
| `primary.default_rr_set_group.mx_record.values.priority` | [primary.default_rr_set_group.mx_record.values.priority](resources--dns_zone--reference--group-002.md#canonical-0301021110310111-1322301311102201-1231000121230232-1333010102002301-2202120233313002-3303212202312212-3030323011221232-2032301101201300) |
| `primary.default_rr_set_group.naptr_record` | [primary.default_rr_set_group.naptr_record](resources--dns_zone--reference--group-002.md#canonical-3302201020210130-3022011223202021-0033303101220201-0002010112333122-0031033012232221-3313032200103232-3102321021201211-1101130203111310) |
| `primary.default_rr_set_group.naptr_record.name` | [primary.default_rr_set_group.naptr_record.name](resources--dns_zone--reference--group-002.md#canonical-3303200222203311-3322303122022321-0310111001311232-1311233100222002-2000210202131011-3322313011110023-3133012122231120-2120030332332222) |
| `primary.default_rr_set_group.naptr_record.values` | [primary.default_rr_set_group.naptr_record.values](resources--dns_zone--reference--group-002.md#canonical-0012133211132101-0222112211233003-0202312222331133-1111020032121221-1311303333202031-3112101213011201-2030200321213312-3020100000132311) |
| `primary.default_rr_set_group.naptr_record.values.flags` | [primary.default_rr_set_group.naptr_record.values.flags](resources--dns_zone--reference--group-002.md#canonical-1033021233210122-3200232001131212-0130330200111101-2002003333123000-3310031000110031-3230321033232230-0130211012022123-3303002312103123) |
| `primary.default_rr_set_group.naptr_record.values.order` | [primary.default_rr_set_group.naptr_record.values.order](resources--dns_zone--reference--group-002.md#canonical-0003322302231222-1002002201313010-3201230211230220-1223201122320320-3311330200200203-0000121233330012-0311232032101331-3123313013332232) |
| `primary.default_rr_set_group.naptr_record.values.preference` | [primary.default_rr_set_group.naptr_record.values.preference](resources--dns_zone--reference--group-002.md#canonical-0203233120120121-1223303131330102-1223232310132002-3301221113103100-3312233033302311-0011020332332101-0011323233301132-3221221100312303) |
| `primary.default_rr_set_group.naptr_record.values.regexp` | [primary.default_rr_set_group.naptr_record.values.regexp](resources--dns_zone--reference--group-002.md#canonical-1133101022121030-3222111130232323-3003312223130130-3000111311111320-0102021122022012-3123102310331021-2331310210210022-1100132112013120) |
| `primary.default_rr_set_group.naptr_record.values.replacement` | [primary.default_rr_set_group.naptr_record.values.replacement](resources--dns_zone--reference--group-002.md#canonical-1303002220120012-3033311223202121-2210031032110320-0220031103301301-1110120230031311-0310212111211132-0101022331313302-1221303110022131) |
| `primary.default_rr_set_group.naptr_record.values.service` | [primary.default_rr_set_group.naptr_record.values.service](resources--dns_zone--reference--group-002.md#canonical-0100322000333330-1013302333103132-2202202223102312-1101112312031213-3310021222110212-1220033222200200-3132030223101200-0120130112301221) |
| `primary.default_rr_set_group.ns_record` | [primary.default_rr_set_group.ns_record](resources--dns_zone--reference--group-002.md#canonical-3133323100030001-3323320220032012-1023332123310223-2133233223111122-2110331310221011-3223131111103203-2022210101000332-0012033013311031) |
| `primary.default_rr_set_group.ns_record.name` | [primary.default_rr_set_group.ns_record.name](resources--dns_zone--reference--group-002.md#canonical-2033003111331232-2100321300133202-0112031132303032-0321230200103202-3333010122311223-3320231030313320-2300021021310111-2111212112023110) |
| `primary.default_rr_set_group.ns_record.values` | [primary.default_rr_set_group.ns_record.values](resources--dns_zone--reference--group-002.md#canonical-3113333313223202-3322202032213010-1202233000332200-3222103322113102-2222133303301233-0320203210302131-3231111130301002-2312123022210102) |
| `primary.default_rr_set_group.ptr_record` | [primary.default_rr_set_group.ptr_record](resources--dns_zone--reference--group-002.md#canonical-2000130221010012-1133232130332203-2023123012103123-1213333232011033-3021323331322132-0220320211113333-2311120321002201-0202020210201231) |
| `primary.default_rr_set_group.ptr_record.name` | [primary.default_rr_set_group.ptr_record.name](resources--dns_zone--reference--group-002.md#canonical-1102100003332030-3003322001321331-2232010102001313-3223110102112111-1320323213032321-1132212230311002-3130231300002110-0132033303303000) |
| `primary.default_rr_set_group.ptr_record.values` | [primary.default_rr_set_group.ptr_record.values](resources--dns_zone--reference--group-002.md#canonical-0212323031323213-3310023311233130-3013102103220130-2030220310203210-3221031003330200-3230100102132031-0110120031033122-3022123022013230) |
| `primary.default_rr_set_group.srv_record` | [primary.default_rr_set_group.srv_record](resources--dns_zone--reference--group-002.md#canonical-3232132023001031-0110132101031030-2312320323010000-0302113021120313-2200303232110120-0032332303330320-2200231003220012-3002111130330010) |
| `primary.default_rr_set_group.srv_record.name` | [primary.default_rr_set_group.srv_record.name](resources--dns_zone--reference--group-002.md#canonical-2021022302320110-0232311121203113-0131321030222332-3321300123301022-2202212013111310-3002012230220200-3010121102310002-1211021110021100) |
| `primary.default_rr_set_group.srv_record.values` | [primary.default_rr_set_group.srv_record.values](resources--dns_zone--reference--group-002.md#canonical-1233110110121133-0211332212112022-0220333223110020-0222101213000123-1223321332033030-1212201202333331-0222210032003303-0133220311010220) |
| `primary.default_rr_set_group.srv_record.values.port` | [primary.default_rr_set_group.srv_record.values.port](resources--dns_zone--reference--group-002.md#canonical-3020210212331312-2113033213301003-3332030210023002-1222331100023311-3213033310033211-3020101212033020-3002301000121233-0023302121210313) |
| `primary.default_rr_set_group.srv_record.values.priority` | [primary.default_rr_set_group.srv_record.values.priority](resources--dns_zone--reference--group-002.md#canonical-0212122213022002-1102130330230022-2113103310311202-3022213003321111-3031122031223210-2302032331122032-0033223320030102-3021033322001002) |
| `primary.default_rr_set_group.srv_record.values.target` | [primary.default_rr_set_group.srv_record.values.target](resources--dns_zone--reference--group-002.md#canonical-3223322100303321-0031100122012000-0231120301220323-3101323020211220-3321023031332030-2000233003301022-1012301232212321-0223131202102230) |
| `primary.default_rr_set_group.srv_record.values.weight` | [primary.default_rr_set_group.srv_record.values.weight](resources--dns_zone--reference--group-002.md#canonical-0223320020032303-2313320021010002-1100112012113203-2222112110223313-0220113331223131-1330112001302113-1102223203003133-2303132101301102) |
| `primary.default_rr_set_group.sshfp_record` | [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-0033310300233313-3001201311021021-3123210133213011-0313232002021102-0111202110303313-0230133031020323-2110230212303300-3320313013021301) |
| `primary.default_rr_set_group.sshfp_record.name` | [primary.default_rr_set_group.sshfp_record.name](resources--dns_zone--reference--group-002.md#canonical-1312323000120211-3122313122033001-2301100232310012-0010001313331121-1103323211133023-1111210113022110-0113201100200021-1323210100121211) |
| `primary.default_rr_set_group.sshfp_record.values` | [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-0323200310303211-3133202223300230-3333231133230112-0322212032230011-1120330010031221-3103010302313102-3310223021330202-0222321323003333) |
| `primary.default_rr_set_group.sshfp_record.values.algorithm` | [primary.default_rr_set_group.sshfp_record.values.algorithm](resources--dns_zone--reference--group-002.md#canonical-0100312300001122-1100322030101033-3013022100332131-0233211220102221-3132132323131321-3013000222230301-1021002323200020-1320101113003203) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint](resources--dns_zone--reference--group-002.md#canonical-3300230102310032-0013311210030213-3111312011020222-2122123002032112-2221302013330301-3202101323320302-3302012213021222-2320211131333010) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint](resources--dns_zone--reference--group-002.md#canonical-3230330302020020-1213132222131323-0002010120323031-3233011323002013-2201200301032200-0121322201011222-1313221131123210-2001212211321122) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint](resources--dns_zone--reference--group-002.md#canonical-0123302233213302-3311232213203031-2013230021213311-3000122000201211-0020103220202111-1121230212202312-2201032200130123-3303320103313102) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint](resources--dns_zone--reference--group-002.md#canonical-1302201230021103-1232233201132323-1030012333213200-1133320003102230-3323030011231103-1230022311102032-2000000122202223-3302112310002200) |
| `primary.default_rr_set_group.tlsa_record` | [primary.default_rr_set_group.tlsa_record](resources--dns_zone--reference--group-002.md#canonical-1320110233003023-3320023020133213-1311023002130001-3013232233301000-2313223031331311-1202330011123222-3011223300111332-1022132230100103) |
| `primary.default_rr_set_group.tlsa_record.name` | [primary.default_rr_set_group.tlsa_record.name](resources--dns_zone--reference--group-002.md#canonical-1223313213332221-2321231333010013-2102333210131220-2031033201102231-0113101310133310-1212302101010313-2333033303031031-1302331113230213) |
| `primary.default_rr_set_group.tlsa_record.values` | [primary.default_rr_set_group.tlsa_record.values](resources--dns_zone--reference--group-002.md#canonical-2033223312322302-1020120221212021-1021023130230232-3102322210033211-1133210131303000-2202231002002201-2100002112220313-2113311310333233) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_association_data` | [primary.default_rr_set_group.tlsa_record.values.certificate_association_data](resources--dns_zone--reference--group-002.md#canonical-2202311121123332-2023231031233301-3303012300022301-0020033213311212-2210302332102100-2311212113013020-3223323000212021-1313100321202233) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_usage` | [primary.default_rr_set_group.tlsa_record.values.certificate_usage](resources--dns_zone--reference--group-002.md#canonical-0113022033122333-3313311131112032-3030022320121300-2333302321230022-1112323100213131-1232222311222310-2222211013312320-0032120012312110) |
| `primary.default_rr_set_group.tlsa_record.values.matching_type` | [primary.default_rr_set_group.tlsa_record.values.matching_type](resources--dns_zone--reference--group-002.md#canonical-1023202102001012-3112230200010101-2321201132100033-0011113231330002-1323022102101133-1232311132322203-2003302222023300-3112032022303023) |
| `primary.default_rr_set_group.tlsa_record.values.selector` | [primary.default_rr_set_group.tlsa_record.values.selector](resources--dns_zone--reference--group-002.md#canonical-1203222200330322-1300312001022303-0100122323211131-2122203330333103-0300310121201012-2003000231312033-2322200333001333-1300020113311311) |
| `primary.default_rr_set_group.ttl` | [primary.default_rr_set_group.ttl](resources--dns_zone--reference--group-001.md#canonical-0031331110203210-2300321011113121-0032002133003202-2000013102013230-2310132031322131-0211231132303121-3301102001131021-0130311220013123) |
| `primary.default_rr_set_group.txt_record` | [primary.default_rr_set_group.txt_record](resources--dns_zone--reference--group-002.md#canonical-1101322310230312-3222303222113203-2322030321021101-3321113220132023-2012323110330200-0103131230030110-3300121021102002-1022211223103031) |
| `primary.default_rr_set_group.txt_record.name` | [primary.default_rr_set_group.txt_record.name](resources--dns_zone--reference--group-002.md#canonical-1021122013021120-1111100310231012-1003331323020003-1022023030313013-3203012320111313-0312211310312032-2000112201121211-3300232121200001) |
| `primary.default_rr_set_group.txt_record.values` | [primary.default_rr_set_group.txt_record.values](resources--dns_zone--reference--group-002.md#canonical-3203321001320331-0232131310021333-3320331222001200-1010030112223331-2002103203121112-0011011311332130-1010001210210030-1012311123120023) |
| `primary.default_soa_parameters` | [primary.default_soa_parameters](resources--dns_zone--reference--group-002.md#canonical-3200211123121101-0222032221231303-2032000002003211-0010311331231033-3032330230003031-1133110003111120-2222022011330333-3102013131201222) |
| `primary.dnssec_mode` | [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-2010020000212012-0312302321031122-0031232110200121-2021212131320032-1030321233203212-1122212300013010-0302322103301010-2301132300330100) |
| `primary.dnssec_mode.disable_spec` | [primary.dnssec_mode.disable_spec](resources--dns_zone--reference--group-002.md#canonical-0102011201220321-3223102212323133-1131213003012321-2300222021123323-2111313310131321-1033100302132121-0201310110333002-1330301122030221) |
| `primary.dnssec_mode.enable` | [primary.dnssec_mode.enable](resources--dns_zone--reference--group-002.md#canonical-2203011132121032-2323023131312322-1021021231331220-0131021102332033-3200331321131112-1130030220013101-2311023022113003-0200222031102223) |
| `primary.rr_set_group` | [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-2322001101312122-2303120132022113-0220120321000013-2022230220112311-0030333011013203-2232310232111211-1302303330103133-2221200303231123) |
| `primary.rr_set_group.metadata` | [primary.rr_set_group.metadata](resources--dns_zone--reference--group-002.md#canonical-1002010330023123-0133011030001021-0201203112331312-3203023122201120-3010100112102002-2211331230303303-3100201021201011-0301202201111300) |
| `primary.rr_set_group.metadata.description_spec` | [primary.rr_set_group.metadata.description_spec](resources--dns_zone--reference--group-002.md#canonical-1322010030310310-0013023022313121-1321130201203213-1020303010211132-1333232101323021-3001232033002200-3021331312303202-0200220301133203) |
| `primary.rr_set_group.metadata.name` | [primary.rr_set_group.metadata.name](resources--dns_zone--reference--group-002.md#canonical-2013210032112020-2012212313210202-3310000230131101-3102233100030210-0333311232132132-3021313113233010-0110022333233002-2300213020220012) |
| `primary.rr_set_group.rr_set` | [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-1321132121303313-2303003201012122-1132223021003222-3003122211220200-2032321003203202-2212100213021102-0323333113103201-2011103021121101) |
| `primary.rr_set_group.rr_set.a_record` | [primary.rr_set_group.rr_set.a_record](resources--dns_zone--reference--group-002.md#canonical-2031310110000111-1322322130201330-3323301001110212-3133310330133013-0210013033102131-2203230200200030-2213013220132213-1301102103210031) |
| `primary.rr_set_group.rr_set.a_record.name` | [primary.rr_set_group.rr_set.a_record.name](resources--dns_zone--reference--group-002.md#canonical-0220023032013002-1221230021322120-1213110113102212-2212322313110302-0021203310120013-0333222231333213-3301220311113111-3020010330101031) |
| `primary.rr_set_group.rr_set.a_record.values` | [primary.rr_set_group.rr_set.a_record.values](resources--dns_zone--reference--group-002.md#canonical-0300022002231102-2302030030032203-3030233303233321-0223032331312033-2323030002001021-3310200220101300-3022110323013323-0333223132133122) |
| `primary.rr_set_group.rr_set.aaaa_record` | [primary.rr_set_group.rr_set.aaaa_record](resources--dns_zone--reference--group-002.md#canonical-1323010102202331-0031123232222202-0211213021200023-3012223001120200-1132111003100012-2302000011113313-2120122120121020-0333031000330003) |
| `primary.rr_set_group.rr_set.aaaa_record.name` | [primary.rr_set_group.rr_set.aaaa_record.name](resources--dns_zone--reference--group-002.md#canonical-3223312000221212-3003122023110122-2003002023300333-0211011333112000-0012300333010312-2222110201231331-0203200221131130-0132121232031223) |
| `primary.rr_set_group.rr_set.aaaa_record.values` | [primary.rr_set_group.rr_set.aaaa_record.values](resources--dns_zone--reference--group-002.md#canonical-3020221301220010-1032211021132313-1123220110010130-0122113131112121-2231112232001330-3103313032100111-2232111032011301-1022233132013200) |
| `primary.rr_set_group.rr_set.afsdb_record` | [primary.rr_set_group.rr_set.afsdb_record](resources--dns_zone--reference--group-002.md#canonical-1212310322223223-2312001020221020-2230222312010033-2320000220311213-0003133123122130-2310303301101003-3111300121022223-2213303111133022) |
| `primary.rr_set_group.rr_set.afsdb_record.name` | [primary.rr_set_group.rr_set.afsdb_record.name](resources--dns_zone--reference--group-002.md#canonical-1213302211210031-2133320200203000-0111102230320010-1221332300120121-0223003031321222-0303132322330003-0030220030203103-1112223023220031) |
| `primary.rr_set_group.rr_set.afsdb_record.values` | [primary.rr_set_group.rr_set.afsdb_record.values](resources--dns_zone--reference--group-002.md#canonical-2302200222330321-1300110110323123-1303233000130322-3201222130200231-0003332331333121-3201222231130021-0232113000202130-0232132312031203) |
| `primary.rr_set_group.rr_set.afsdb_record.values.hostname` | [primary.rr_set_group.rr_set.afsdb_record.values.hostname](resources--dns_zone--reference--group-002.md#canonical-2102002201003330-2002101221012223-3001203121312002-3211211121013330-3111312122230022-1210303123021302-3111023333013132-2313221302033331) |
| `primary.rr_set_group.rr_set.afsdb_record.values.subtype` | [primary.rr_set_group.rr_set.afsdb_record.values.subtype](resources--dns_zone--reference--group-002.md#canonical-2000111133310013-2233120201210110-2220002303030010-2110000110303223-3010023113213023-3123100010310020-1201230020100323-3032301122200321) |
| `primary.rr_set_group.rr_set.alias_record` | [primary.rr_set_group.rr_set.alias_record](resources--dns_zone--reference--group-002.md#canonical-3023031303112002-1011220313330031-3300012333020203-1103120213122123-0332110100122323-0321332213233010-1003301332021321-3322231133202100) |
| `primary.rr_set_group.rr_set.alias_record.value` | [primary.rr_set_group.rr_set.alias_record.value](resources--dns_zone--reference--group-002.md#canonical-3203103303223000-2221133301210312-1313103330113003-1000031030201221-0122121222313102-1221020030310013-3021011011232033-3231211111121212) |
| `primary.rr_set_group.rr_set.caa_record` | [primary.rr_set_group.rr_set.caa_record](resources--dns_zone--reference--group-002.md#canonical-3203022223320231-1032100010021023-3102132232202011-3221031132212002-2332032001000210-3210130210033031-0230230202331030-0132232201123103) |
| `primary.rr_set_group.rr_set.caa_record.name` | [primary.rr_set_group.rr_set.caa_record.name](resources--dns_zone--reference--group-002.md#canonical-0202232310230231-0100333010221323-2311013202202210-1230220200223232-0130303212131312-1131222311113120-1033123120003232-3010203312311203) |
| `primary.rr_set_group.rr_set.caa_record.values` | [primary.rr_set_group.rr_set.caa_record.values](resources--dns_zone--reference--group-002.md#canonical-2113020310310122-0010131022220111-2010230133103022-0213103110113121-1002000132023232-2111320322131310-0112002102320313-1102302232201030) |
| `primary.rr_set_group.rr_set.caa_record.values.flags` | [primary.rr_set_group.rr_set.caa_record.values.flags](resources--dns_zone--reference--group-002.md#canonical-0301013201232232-3231012001032130-2102210023001223-3113003232330333-1230031010002122-3123133202333113-0212022022112022-2103100200100311) |
| `primary.rr_set_group.rr_set.caa_record.values.tag` | [primary.rr_set_group.rr_set.caa_record.values.tag](resources--dns_zone--reference--group-002.md#canonical-0330230220122010-2001122300032033-0202022222323200-0233331130103322-1131320300310111-1332321010302331-0203103301013320-3231111120230011) |
| `primary.rr_set_group.rr_set.caa_record.values.value` | [primary.rr_set_group.rr_set.caa_record.values.value](resources--dns_zone--reference--group-002.md#canonical-0000131020323322-3312101311333222-3310002101231232-2213233023311213-3002212311223300-0310121023110323-1301230212012130-0211303130103132) |
| `primary.rr_set_group.rr_set.cds_record` | [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-3012220321222313-2102100303223132-3320001303231312-1113220202122233-3233033202202311-3303302130320322-3130001301210120-3230110112300123) |
| `primary.rr_set_group.rr_set.cds_record.name` | [primary.rr_set_group.rr_set.cds_record.name](resources--dns_zone--reference--group-002.md#canonical-1202020213032001-1221001000120313-0200233210221010-1120001323311232-0231212323131302-1220021102102332-1111103201120221-0023103032201132) |
| `primary.rr_set_group.rr_set.cds_record.values` | [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-003.md#canonical-2321310003223320-1321332333030003-2322033021331223-1330322002030000-3221210101203310-1200302130113113-0210313223323033-2023023212230202) |
| `primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm](resources--dns_zone--reference--group-003.md#canonical-3221313023302011-1302331132022102-2102313232003220-1011111302030030-2132232010020200-3212131221320011-0101312021001313-2011310131011331) |
| `primary.rr_set_group.rr_set.cds_record.values.key_tag` | [primary.rr_set_group.rr_set.cds_record.values.key_tag](resources--dns_zone--reference--group-003.md#canonical-3231012122013300-0201102022233213-3230031002000002-1022233323000132-1132223033013323-1210102310333003-0201113120223330-1102111303211100) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest](resources--dns_zone--reference--group-003.md#canonical-3033132103122202-2320322021002103-2212203030021003-3313002010003333-0320010103312010-3332021022111200-0130221032312321-1230113302010031) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest](resources--dns_zone--reference--group-003.md#canonical-2132020012111003-1303102023010130-1010202313311110-2100201300203203-2201121330113311-0220010203302022-2223132130313211-3312311030322213) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest](resources--dns_zone--reference--group-003.md#canonical-0010233123223233-2012102031300132-3312323022113110-3031311102102121-2101311323203202-2131032213233033-1131311220010212-2221222002001002) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest](resources--dns_zone--reference--group-003.md#canonical-0221230202212112-0211131000331112-2321230213212000-3122310223312310-0210202320203302-2231231020101232-3230200313113132-3000021201202213) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest](resources--dns_zone--reference--group-003.md#canonical-3120330220013230-0303113122202021-2201232202113323-2022113002210013-0330202233323122-0233122303011112-1210231220322032-0101310120112201) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest](resources--dns_zone--reference--group-003.md#canonical-1130113223110200-0120022101323323-2031221211033022-0120123132011321-2021010302130322-0331110032010222-3300220222212311-0312312030112311) |
| `primary.rr_set_group.rr_set.cert_record` | [primary.rr_set_group.rr_set.cert_record](resources--dns_zone--reference--group-003.md#canonical-3012301320210302-1333331200213200-3010310033130210-1223002010211130-0022101013011131-0230313332121221-0332120101200133-1231213322232330) |
| `primary.rr_set_group.rr_set.cert_record.name` | [primary.rr_set_group.rr_set.cert_record.name](resources--dns_zone--reference--group-003.md#canonical-2020032320331031-0011123303232211-0321013130110322-0013222113013223-0111012212223311-1111032321332332-0222332031302331-0230333013301020) |
| `primary.rr_set_group.rr_set.cert_record.values` | [primary.rr_set_group.rr_set.cert_record.values](resources--dns_zone--reference--group-003.md#canonical-2321003200110130-1233110210113031-1223210301301013-2121223122301131-1323123030313000-3023232220110202-0033103212123212-1003302122000220) |
| `primary.rr_set_group.rr_set.cert_record.values.algorithm` | [primary.rr_set_group.rr_set.cert_record.values.algorithm](resources--dns_zone--reference--group-003.md#canonical-1211012122010310-0201120313120013-1032322211321110-3133121300311113-3210223231022230-1311321311232212-2301212110321012-1112220013310023) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_key_tag` | [primary.rr_set_group.rr_set.cert_record.values.cert_key_tag](resources--dns_zone--reference--group-003.md#canonical-1303010300120031-1323220122303030-2132211322123221-0022313212032000-0322201331020131-1120312020210320-2122311112312121-1212323022302322) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_type` | [primary.rr_set_group.rr_set.cert_record.values.cert_type](resources--dns_zone--reference--group-003.md#canonical-2131303313320230-3121011203223201-3112232213321203-2021210323002222-3133020003302001-3033100333133332-3230300002002000-2210211112033030) |
| `primary.rr_set_group.rr_set.cert_record.values.certificate` | [primary.rr_set_group.rr_set.cert_record.values.certificate](resources--dns_zone--reference--group-003.md#canonical-1030030023212121-1310122023033021-1222333022110320-2030112330130132-2021320300013223-3231113123013013-0221132013100002-0330200333032321) |
| `primary.rr_set_group.rr_set.cname_record` | [primary.rr_set_group.rr_set.cname_record](resources--dns_zone--reference--group-003.md#canonical-0310311121001122-2211200002312221-3310123210132123-1111003013102101-0220013300100011-3323021133223210-1203332212331011-2031212020103123) |
| `primary.rr_set_group.rr_set.cname_record.name` | [primary.rr_set_group.rr_set.cname_record.name](resources--dns_zone--reference--group-003.md#canonical-1220111321122301-3001030230300300-0313021211213223-3313012203222220-3233002233010220-3300121311110312-2112210113022301-3022113321331133) |
| `primary.rr_set_group.rr_set.cname_record.value` | [primary.rr_set_group.rr_set.cname_record.value](resources--dns_zone--reference--group-003.md#canonical-2120103131210132-3011322322030331-3011211302212322-0232133033012312-0020100212002222-2321013301033303-3231312232333112-3002112120212101) |
| `primary.rr_set_group.rr_set.description_spec` | [primary.rr_set_group.rr_set.description_spec](resources--dns_zone--reference--group-002.md#canonical-2133332032212210-0123213230222121-3123113232120122-1010332211233120-3100320331130120-0330130000220032-2113103331312030-1322231031232211) |
| `primary.rr_set_group.rr_set.ds_record` | [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-1312320202330202-3022222003331233-1020030132132200-1011023111332103-3330200202323012-1123022031311220-2221113301031332-2302121100021102) |
| `primary.rr_set_group.rr_set.ds_record.name` | [primary.rr_set_group.rr_set.ds_record.name](resources--dns_zone--reference--group-003.md#canonical-0032321310113023-1133320130332132-2121301000120130-1023313113323311-0330201003120100-2023003231303311-3101000322230333-1220332110021031) |
| `primary.rr_set_group.rr_set.ds_record.values` | [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-2100203210333023-2131000233103233-3110200332300003-0202103210212130-0311330213233211-3003301122200002-0210011133323020-3132332211132201) |
| `primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm](resources--dns_zone--reference--group-003.md#canonical-2323221310033131-0222331312210010-1322132223111200-0020031323310313-1203021201000031-0100113321133012-0303303100011131-1330230201310300) |
| `primary.rr_set_group.rr_set.ds_record.values.key_tag` | [primary.rr_set_group.rr_set.ds_record.values.key_tag](resources--dns_zone--reference--group-003.md#canonical-1022013312210330-3111132222231013-0131033022321103-1203322330311311-0233320321200221-2103232330233210-1131331232030031-1303232033110231) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest](resources--dns_zone--reference--group-003.md#canonical-0011133131321023-3130203301120032-0110112220131102-2111121011201021-2211331322120130-2002031220021330-2112231221322020-0223322123323200) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest](resources--dns_zone--reference--group-003.md#canonical-2232131303232033-0333312022022312-1111110010021003-2101222120032030-3330123100230122-2130222210233312-0022103013111032-0202231013211033) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest](resources--dns_zone--reference--group-003.md#canonical-3223300111103300-2131113003121202-0300202322223220-1212302211030213-1133121200002131-2300100110232331-0220122010230232-2213220212023120) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest](resources--dns_zone--reference--group-003.md#canonical-1102232200130203-3213023331321211-0313023210120322-2121210303232303-3020110200013132-3312310210123331-3231002011333230-1213030122123000) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest](resources--dns_zone--reference--group-003.md#canonical-2020321033011002-0102000003310030-2202320000020333-0010133010030020-0130313010330210-3311003231003101-0212100220121010-2220132120202300) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest](resources--dns_zone--reference--group-003.md#canonical-1320313210322131-2231002220303302-0332223110112201-2232123012321012-1221031322020202-0210103332300123-2223332212300301-1131002122200030) |
| `primary.rr_set_group.rr_set.eui48_record` | [primary.rr_set_group.rr_set.eui48_record](resources--dns_zone--reference--group-003.md#canonical-0201123302222103-2312013113132012-0122000120310320-0320022121033100-3013022323220030-3331223213221333-2321013100233321-3302330110121232) |
| `primary.rr_set_group.rr_set.eui48_record.name` | [primary.rr_set_group.rr_set.eui48_record.name](resources--dns_zone--reference--group-003.md#canonical-2002122332210301-1330301002010310-3110211312012330-2323213122132332-1333031220133212-1220032013012131-0133100103102002-1331021133023020) |
| `primary.rr_set_group.rr_set.eui48_record.value` | [primary.rr_set_group.rr_set.eui48_record.value](resources--dns_zone--reference--group-003.md#canonical-0002023132231011-0303112102101232-1021233301002103-3320333213213013-1212111020313131-1311033201013113-2110200100221133-2121213220021303) |
| `primary.rr_set_group.rr_set.eui64_record` | [primary.rr_set_group.rr_set.eui64_record](resources--dns_zone--reference--group-003.md#canonical-2013211131032121-0312023111000010-0122020330121311-3122211222020033-0030010323222303-0111203231222112-2211112301312120-1210203033031233) |
| `primary.rr_set_group.rr_set.eui64_record.name` | [primary.rr_set_group.rr_set.eui64_record.name](resources--dns_zone--reference--group-003.md#canonical-3003013232120122-2111212112102312-2002331003000031-0223123230200221-1022202212233002-3101121031230310-3230022312231202-2221012103233032) |
| `primary.rr_set_group.rr_set.eui64_record.value` | [primary.rr_set_group.rr_set.eui64_record.value](resources--dns_zone--reference--group-003.md#canonical-1102123323213102-3112220233132131-2232323223003302-0101010323201231-3130012311112111-3100222311132122-3113323010332001-1221321213220212) |
| `primary.rr_set_group.rr_set.lb_record` | [primary.rr_set_group.rr_set.lb_record](resources--dns_zone--reference--group-003.md#canonical-0021120310010231-0211130000112132-3000333101232313-3320012233310210-2300011022210313-2233120301332012-2023332030010130-0213202201302233) |
| `primary.rr_set_group.rr_set.lb_record.name` | [primary.rr_set_group.rr_set.lb_record.name](resources--dns_zone--reference--group-003.md#canonical-0021020300011012-2022131312323132-3001202303113200-2131321121113030-1312131110310332-0120032231133222-1210312011100312-3321113221002200) |
| `primary.rr_set_group.rr_set.lb_record.value` | [primary.rr_set_group.rr_set.lb_record.value](resources--dns_zone--reference--group-003.md#canonical-2030002130231132-2000012332213113-1213331012213220-2121101123302313-2103300331310021-3120010212222111-2022230201301023-1002333113300020) |
| `primary.rr_set_group.rr_set.lb_record.value.name` | [primary.rr_set_group.rr_set.lb_record.value.name](resources--dns_zone--reference--group-003.md#canonical-2130011331331021-0303222032023313-0222032030222131-1133331012200100-0032122011332321-3023221320120133-2130121033230223-0103230120303301) |
| `primary.rr_set_group.rr_set.lb_record.value.namespace` | [primary.rr_set_group.rr_set.lb_record.value.namespace](resources--dns_zone--reference--group-003.md#canonical-2112210102230222-2032201010213103-3312310231113002-3130221302303220-2312300113013012-3023233212002331-0230213132012331-1230101230212321) |
| `primary.rr_set_group.rr_set.lb_record.value.tenant` | [primary.rr_set_group.rr_set.lb_record.value.tenant](resources--dns_zone--reference--group-003.md#canonical-1322131212320031-0201330021321131-0313321131201233-1132132212021322-2000322002013013-1130310102311223-1033033120020102-1002010220201202) |
| `primary.rr_set_group.rr_set.loc_record` | [primary.rr_set_group.rr_set.loc_record](resources--dns_zone--reference--group-003.md#canonical-1203123000020330-2113200013101102-1113123013131013-1210020133113223-1223230311323133-1001123200011213-0103212102012311-0322023223102232) |
| `primary.rr_set_group.rr_set.loc_record.name` | [primary.rr_set_group.rr_set.loc_record.name](resources--dns_zone--reference--group-003.md#canonical-2201011333233103-3121332211033331-2231300010202123-0332212313332201-2113003330013321-2100003331321020-0332101220101331-3102323322321131) |
| `primary.rr_set_group.rr_set.loc_record.values` | [primary.rr_set_group.rr_set.loc_record.values](resources--dns_zone--reference--group-003.md#canonical-3313311200021130-1131002203100230-2031300301002002-3101101210022010-3311213223203013-2322301230200210-3020231212101013-2213131332202110) |
| `primary.rr_set_group.rr_set.loc_record.values.altitude` | [primary.rr_set_group.rr_set.loc_record.values.altitude](resources--dns_zone--reference--group-003.md#canonical-0013112313120332-3320132333222132-1331001033001122-0033001311322013-0302211032210011-3131023033300010-3000323211033002-0313030203101130) |
| `primary.rr_set_group.rr_set.loc_record.values.horizontal_precision` | [primary.rr_set_group.rr_set.loc_record.values.horizontal_precision](resources--dns_zone--reference--group-003.md#canonical-0302000232331101-3133201211201122-3132233113031112-2323313130202220-3321020231222211-2130033321232001-0012111322120301-0313020120221001) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.latitude_degree](resources--dns_zone--reference--group-003.md#canonical-3022310322123313-0311101020222012-1020230200300031-2322110031022310-2102101131200133-0132113323121131-3310021130310013-1123200031113211) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere](resources--dns_zone--reference--group-003.md#canonical-3310130223121023-2221330010131111-2001312003111220-0232232312132011-1310113100303011-1321212133120332-2030301033020023-3123133000211103) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.latitude_minute](resources--dns_zone--reference--group-003.md#canonical-2000212200110013-0000201320333332-0113331020020110-2102202223021330-2132110102021133-1310031021020032-3133313213032313-2330021201020213) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_second` | [primary.rr_set_group.rr_set.loc_record.values.latitude_second](resources--dns_zone--reference--group-003.md#canonical-0321123302220223-1213203131102300-1000112213311010-3302322030020202-0111130331232310-0002103001322223-1112231210023113-1223123312232322) |
| `primary.rr_set_group.rr_set.loc_record.values.location_diameter` | [primary.rr_set_group.rr_set.loc_record.values.location_diameter](resources--dns_zone--reference--group-003.md#canonical-0000130111101322-2233132231333001-2302120220310101-3312021312221011-0112012322212233-2003103030011131-0030222122223300-2132111303102301) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.longitude_degree](resources--dns_zone--reference--group-003.md#canonical-1112023302312232-1331003021013200-1021302200332101-2110023122323012-2123311211120002-0332122102330131-3121320123011311-3020232223323122) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere](resources--dns_zone--reference--group-003.md#canonical-3311213031211000-2203110013110303-3011131310313213-2102122010222302-0311321031303203-3303003111302301-1221312122122210-3213102101323012) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.longitude_minute](resources--dns_zone--reference--group-003.md#canonical-3333311133103112-1223122203322303-1101213130001103-2003122121110130-0121201213333112-0302012011133303-3302000002003231-0010030132101310) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_second` | [primary.rr_set_group.rr_set.loc_record.values.longitude_second](resources--dns_zone--reference--group-003.md#canonical-0003300311211321-3302033301013000-1303223131102001-0021310012301222-3002221110003121-3102100110322301-3301103010223123-0202202000321131) |
| `primary.rr_set_group.rr_set.loc_record.values.vertical_precision` | [primary.rr_set_group.rr_set.loc_record.values.vertical_precision](resources--dns_zone--reference--group-003.md#canonical-3023332133221031-2213200201011103-3213012023033333-0312213013012003-2200302200332113-0320122133201000-3000300330100222-0030211201320001) |
| `primary.rr_set_group.rr_set.mx_record` | [primary.rr_set_group.rr_set.mx_record](resources--dns_zone--reference--group-003.md#canonical-2132030313003332-3230003011110221-0130002031232112-3133020121102322-2223230102301312-2023030330230330-3131300100100313-3221132223101331) |
| `primary.rr_set_group.rr_set.mx_record.name` | [primary.rr_set_group.rr_set.mx_record.name](resources--dns_zone--reference--group-003.md#canonical-0102031200201131-1132303302003222-3210331133220110-3332322020300030-0033310330203230-2010332323122023-0032030122310211-1112003010201310) |
| `primary.rr_set_group.rr_set.mx_record.values` | [primary.rr_set_group.rr_set.mx_record.values](resources--dns_zone--reference--group-003.md#canonical-2202333202101131-1132102103013000-1112330121332320-3212223020221231-2023130231102331-3203303102210210-0033220002221200-0313132030222011) |
| `primary.rr_set_group.rr_set.mx_record.values.domain` | [primary.rr_set_group.rr_set.mx_record.values.domain](resources--dns_zone--reference--group-003.md#canonical-1133102322110311-3101323123323131-3101301203111231-3223212323211213-1331110102101010-3233312022332330-2030332033303201-1102021310032131) |
| `primary.rr_set_group.rr_set.mx_record.values.priority` | [primary.rr_set_group.rr_set.mx_record.values.priority](resources--dns_zone--reference--group-003.md#canonical-3111313022123021-1133213322112032-2203121002200232-1121301000133330-3000133312233103-1333132030312300-1102030223123323-0022112233333301) |
| `primary.rr_set_group.rr_set.naptr_record` | [primary.rr_set_group.rr_set.naptr_record](resources--dns_zone--reference--group-003.md#canonical-0113102211212002-1101222021022110-3101000310021221-3213123131122132-2222303310010210-3123021310123133-0312033331022113-3003012030033011) |
| `primary.rr_set_group.rr_set.naptr_record.name` | [primary.rr_set_group.rr_set.naptr_record.name](resources--dns_zone--reference--group-003.md#canonical-1200012232231001-1301100312311031-2133300113100220-1030101021010222-2032223233100320-0333201103130110-0123103332311012-3120002221030012) |
| `primary.rr_set_group.rr_set.naptr_record.values` | [primary.rr_set_group.rr_set.naptr_record.values](resources--dns_zone--reference--group-003.md#canonical-0222200233003323-0322031231313100-1201022110113322-1130023112121030-3103233101313121-1313030302302333-2322013132210231-0032111113103033) |
| `primary.rr_set_group.rr_set.naptr_record.values.flags` | [primary.rr_set_group.rr_set.naptr_record.values.flags](resources--dns_zone--reference--group-003.md#canonical-2000332132310023-1130303122313003-2101233300213131-2030002031201201-3111212121131033-1302032203222301-3013111030321123-0112300300330223) |
| `primary.rr_set_group.rr_set.naptr_record.values.order` | [primary.rr_set_group.rr_set.naptr_record.values.order](resources--dns_zone--reference--group-003.md#canonical-1131013031102231-1311302210303031-1031033032330023-0013330201133112-1300021322013100-2210010012012122-3132301203200123-0231312001211031) |
| `primary.rr_set_group.rr_set.naptr_record.values.preference` | [primary.rr_set_group.rr_set.naptr_record.values.preference](resources--dns_zone--reference--group-003.md#canonical-2011121233210300-0121201122303221-3300003023003222-0132232213132022-2310301231212110-1203320320310023-2113033110111312-1203321032011212) |
| `primary.rr_set_group.rr_set.naptr_record.values.regexp` | [primary.rr_set_group.rr_set.naptr_record.values.regexp](resources--dns_zone--reference--group-003.md#canonical-1022122030201233-0030103013313101-1220222121123303-1230222122133123-2112013030212201-3332121011100211-0330212323001202-3302211120230300) |
| `primary.rr_set_group.rr_set.naptr_record.values.replacement` | [primary.rr_set_group.rr_set.naptr_record.values.replacement](resources--dns_zone--reference--group-003.md#canonical-1202323213312222-3123130121131000-3331022313122333-2223222230203203-1020000320101032-3102220101303203-1212203303132201-1211223210113211) |
| `primary.rr_set_group.rr_set.naptr_record.values.service` | [primary.rr_set_group.rr_set.naptr_record.values.service](resources--dns_zone--reference--group-003.md#canonical-1030011011201021-2321133132211031-2303023321320202-1201311121123320-2320202102312212-1333111131112123-2131131301202323-1123331311213000) |
| `primary.rr_set_group.rr_set.ns_record` | [primary.rr_set_group.rr_set.ns_record](resources--dns_zone--reference--group-003.md#canonical-1231003122123303-2133013001222111-3132233130013001-3012311102221220-2103210112001212-3323000203323113-3120302331000232-2021211213312321) |
| `primary.rr_set_group.rr_set.ns_record.name` | [primary.rr_set_group.rr_set.ns_record.name](resources--dns_zone--reference--group-003.md#canonical-3320020101210010-1122130010031303-3301323332322213-2301330010001022-0121210231130222-1032122301130321-1320031201003011-2112231102302113) |
| `primary.rr_set_group.rr_set.ns_record.values` | [primary.rr_set_group.rr_set.ns_record.values](resources--dns_zone--reference--group-003.md#canonical-2013013321002000-1221302210011312-1303231000120231-3302010111320010-1331010332113222-2302333313221033-3033000111223302-1212201201230302) |
| `primary.rr_set_group.rr_set.ptr_record` | [primary.rr_set_group.rr_set.ptr_record](resources--dns_zone--reference--group-003.md#canonical-3013221023311022-0020211130201303-3131202122033003-1300210133223000-3222010023132110-1001102031230113-0020222130133000-1122303200203200) |
| `primary.rr_set_group.rr_set.ptr_record.name` | [primary.rr_set_group.rr_set.ptr_record.name](resources--dns_zone--reference--group-003.md#canonical-3100223311113131-0103210012013210-0331322110221133-0020032321022320-2022203200101222-2310003322230001-0213311230111011-1202012120212032) |
| `primary.rr_set_group.rr_set.ptr_record.values` | [primary.rr_set_group.rr_set.ptr_record.values](resources--dns_zone--reference--group-003.md#canonical-2023312301232211-1132111221331013-2102012300332121-0312000310311303-0030132032012211-3231013321332002-1203032112210230-0230333012212213) |
| `primary.rr_set_group.rr_set.srv_record` | [primary.rr_set_group.rr_set.srv_record](resources--dns_zone--reference--group-003.md#canonical-2001132300321330-0200133330021203-0313213322013100-0111121110201021-2012001100220230-1000022333221202-3112310302033103-1021221233100113) |
| `primary.rr_set_group.rr_set.srv_record.name` | [primary.rr_set_group.rr_set.srv_record.name](resources--dns_zone--reference--group-003.md#canonical-3022131012200113-3123121211100100-0221003010123321-0213303131121032-3122323030123030-0320131123002113-3101323220311231-2330310322133201) |
| `primary.rr_set_group.rr_set.srv_record.values` | [primary.rr_set_group.rr_set.srv_record.values](resources--dns_zone--reference--group-003.md#canonical-0330111210022221-0223330313120313-0123322001022203-3102121300120123-2000122203102101-1130011312000333-2023210331301111-1111113130321111) |
| `primary.rr_set_group.rr_set.srv_record.values.port` | [primary.rr_set_group.rr_set.srv_record.values.port](resources--dns_zone--reference--group-003.md#canonical-0320200121330033-2102220303010332-2331133001210202-0122123203111103-3221312301130130-1333222312220221-1110032222223301-3010223032320201) |
| `primary.rr_set_group.rr_set.srv_record.values.priority` | [primary.rr_set_group.rr_set.srv_record.values.priority](resources--dns_zone--reference--group-003.md#canonical-0330033232222012-2031220323210103-0232222212023222-1202232323320001-0113231211002301-3321331300333002-1132313110213332-2303222323230010) |
| `primary.rr_set_group.rr_set.srv_record.values.target` | [primary.rr_set_group.rr_set.srv_record.values.target](resources--dns_zone--reference--group-003.md#canonical-0230233002211220-0231113232232120-0113320001010100-3301022031010221-2213213123333233-3202101320301203-0310330333313131-2122120102222023) |
| `primary.rr_set_group.rr_set.srv_record.values.weight` | [primary.rr_set_group.rr_set.srv_record.values.weight](resources--dns_zone--reference--group-003.md#canonical-1111203100210131-0200111203203233-3202120003030231-0203121301223000-0202113011300021-3211321033111023-3301200001303300-1312311121000120) |
| `primary.rr_set_group.rr_set.sshfp_record` | [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-0121121322011211-1303301211230003-0201033312000112-0101223100102212-2102100201221222-1233012320013113-2021333323302131-0213331312123121) |
| `primary.rr_set_group.rr_set.sshfp_record.name` | [primary.rr_set_group.rr_set.sshfp_record.name](resources--dns_zone--reference--group-003.md#canonical-1213032112302330-2022300300122131-3030013113030212-0301021202311320-1331110100322013-3112223312103322-1201011233013001-1013022130133033) |
| `primary.rr_set_group.rr_set.sshfp_record.values` | [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-3300100220311322-3102300013023303-1002222232103113-1010021032220123-3021101232031310-0113110031110203-1030020300032022-2311333323323231) |
| `primary.rr_set_group.rr_set.sshfp_record.values.algorithm` | [primary.rr_set_group.rr_set.sshfp_record.values.algorithm](resources--dns_zone--reference--group-003.md#canonical-2002221222101033-0010212011212120-3123322310313233-3331213133311320-1023101333213032-3320100113110003-2022031113231031-1121332202330311) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint](resources--dns_zone--reference--group-003.md#canonical-1130201023023123-3003233001022022-1103002023221203-3120103301012110-2002112000301333-0001111201012033-0332131013022210-3321213103103321) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint](resources--dns_zone--reference--group-003.md#canonical-3012211102132302-0032303133312012-3112122000300033-3333100023133023-0333322332032103-0010100000331323-1101213233331000-0322311032202022) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint](resources--dns_zone--reference--group-003.md#canonical-2100332123103230-2213333330233032-1113231100012320-0211332122323122-0010033302130021-1112322003202312-0013033222301123-1012300030101212) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint](resources--dns_zone--reference--group-003.md#canonical-1203210033210132-2221003233232103-1230233003103323-3030131030010303-1010112211003312-0013220311321321-0112003033021310-1100203131022201) |
| `primary.rr_set_group.rr_set.tlsa_record` | [primary.rr_set_group.rr_set.tlsa_record](resources--dns_zone--reference--group-003.md#canonical-3310230100220110-2003230230011010-1023111321302213-2230131233321320-3001222102312131-3232032301310331-0320213001312302-2322332212011113) |
| `primary.rr_set_group.rr_set.tlsa_record.name` | [primary.rr_set_group.rr_set.tlsa_record.name](resources--dns_zone--reference--group-003.md#canonical-1311310012330022-1003131133330232-1230203212321112-0012021010333112-2320101211130013-0012211102012103-0022200321311220-0230020023123011) |
| `primary.rr_set_group.rr_set.tlsa_record.values` | [primary.rr_set_group.rr_set.tlsa_record.values](resources--dns_zone--reference--group-003.md#canonical-2211233033003132-3231130220321230-1111303101003111-3321000232232120-1212221101223233-3320200121020020-1232201123102121-0333003133110300) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data](resources--dns_zone--reference--group-003.md#canonical-2011321131202002-2100102112300020-3330121221020233-2132110032030031-3202131212312122-0100100011213313-3121100203120303-2020201311021230) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage](resources--dns_zone--reference--group-003.md#canonical-1011330200330030-2123312103302022-2113100323023323-1302211300022100-1130312210331310-1103233313330103-3322203320322103-2330211330322212) |
| `primary.rr_set_group.rr_set.tlsa_record.values.matching_type` | [primary.rr_set_group.rr_set.tlsa_record.values.matching_type](resources--dns_zone--reference--group-003.md#canonical-0120103313311223-3311212033320132-2313021132103012-3120322113032010-2210102330331300-1231213031030011-2233131013311220-2333023230333201) |
| `primary.rr_set_group.rr_set.tlsa_record.values.selector` | [primary.rr_set_group.rr_set.tlsa_record.values.selector](resources--dns_zone--reference--group-003.md#canonical-1231000011230302-1002130201101212-3023110333020223-3030001121300001-0200012003331321-2103210113120121-0311330322323123-3221133300203330) |
| `primary.rr_set_group.rr_set.ttl` | [primary.rr_set_group.rr_set.ttl](resources--dns_zone--reference--group-002.md#canonical-2212321033023303-0323001231312003-3210032310003003-0201001200102131-0131103232010111-2122312100213023-1303001300022011-0200102110001300) |
| `primary.rr_set_group.rr_set.txt_record` | [primary.rr_set_group.rr_set.txt_record](resources--dns_zone--reference--group-003.md#canonical-1100331110322030-0201111130213202-3110023110000301-0003213100033201-1112001001223102-1320003200303132-3000323323020202-0033331320301332) |
| `primary.rr_set_group.rr_set.txt_record.name` | [primary.rr_set_group.rr_set.txt_record.name](resources--dns_zone--reference--group-003.md#canonical-2030201221211031-0012002020130111-1310002001030002-0230223203110013-2302201111203001-2113210032213321-2332032321221200-2223333202010210) |
| `primary.rr_set_group.rr_set.txt_record.values` | [primary.rr_set_group.rr_set.txt_record.values](resources--dns_zone--reference--group-003.md#canonical-1213202100110122-1000332210102110-2233200221331210-3112130123122210-3131301222211012-2011200123210122-0203022103313330-2032101302022210) |
| `primary.soa_parameters` | [primary.soa_parameters](resources--dns_zone--reference--group-003.md#canonical-2121220321110032-3211203313030203-0103020311122201-0023320210202020-3121000202111212-1213011301133030-3111130103212201-0011312132222212) |
| `primary.soa_parameters.expire` | [primary.soa_parameters.expire](resources--dns_zone--reference--group-003.md#canonical-3230132300000033-3110101120232103-0102232131022221-0213122210211302-1321012320222201-0122020231231331-3331220311231102-3102322313020311) |
| `primary.soa_parameters.negative_ttl` | [primary.soa_parameters.negative_ttl](resources--dns_zone--reference--group-003.md#canonical-2321101100133013-2133313330023311-0230331332003001-3012220301332222-0102123320302303-1312030230112333-3112320201003201-1330311010030021) |
| `primary.soa_parameters.refresh` | [primary.soa_parameters.refresh](resources--dns_zone--reference--group-003.md#canonical-2312322201230120-2102133302301003-1020333031120112-3232213210311122-3322000112101202-3011312002030010-3001102333031203-2200333102220312) |
| `primary.soa_parameters.retry` | [primary.soa_parameters.retry](resources--dns_zone--reference--group-003.md#canonical-0132223313310300-3002321121310323-1322333201102312-3002131301301323-2310010230213030-1221131321220023-0222013132031032-0201101302133211) |
| `primary.soa_parameters.ttl` | [primary.soa_parameters.ttl](resources--dns_zone--reference--group-003.md#canonical-3031121330002222-1122010031033213-2002203113122001-2220320201322012-3303123323323323-0312201022120333-3312123322122323-2113131330202000) |
| `secondary` | [secondary](resources--dns_zone--reference--group-003.md#canonical-0223010131102331-0310121320000300-3131133100321303-1312200031031101-2101222213223300-3110301303132223-1023132100303123-3012223223200121) |
| `secondary.primary_servers` | [secondary.primary_servers](resources--dns_zone--reference--group-003.md#canonical-3022322020202002-2320120232202130-1332110110120132-3231131222233013-3023023031131313-2000321032320123-0222022323233201-0223032300003022) |
| `secondary.tsig_key_algorithm` | [secondary.tsig_key_algorithm](resources--dns_zone--reference--group-003.md#canonical-2131221012001030-0213220100322000-0121123330330200-3122302010233231-3322322030130132-1232101300233203-2022220033103233-2112112211323031) |
| `secondary.tsig_key_name` | [secondary.tsig_key_name](resources--dns_zone--reference--group-003.md#canonical-1312113303301123-0131313203312122-2000203020002221-2012011313232100-0302313003131321-3031110232312212-3300111032232002-3020122021101320) |
| `secondary.tsig_key_value` | [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-1133331300221133-0123021301132223-1331302203102230-0220012222331103-1201221303001002-3103010003311200-3002102213203003-3012013312111321) |
| `secondary.tsig_key_value.blindfold_secret_info` | [secondary.tsig_key_value.blindfold_secret_info](resources--dns_zone--reference--group-003.md#canonical-3301012122133333-0110202122221123-0220023330200102-2000211100012300-0130332230302202-1322101030213033-2113110313221230-1332020312120123) |
| `secondary.tsig_key_value.blindfold_secret_info.decryption_provider` | [secondary.tsig_key_value.blindfold_secret_info.decryption_provider](resources--dns_zone--reference--group-003.md#canonical-2033213322032312-0311201200230231-2001101223023201-1330222303203323-0301121110300103-1313212221133112-0100323011301232-3132213223101233) |
| `secondary.tsig_key_value.blindfold_secret_info.location` | [secondary.tsig_key_value.blindfold_secret_info.location](resources--dns_zone--reference--group-003.md#canonical-3023330002313330-0310233121031302-3013213121310201-3103011010012001-1331201301011322-0131330133132222-2031020230021022-1303333022323331) |
| `secondary.tsig_key_value.blindfold_secret_info.store_provider` | [secondary.tsig_key_value.blindfold_secret_info.store_provider](resources--dns_zone--reference--group-003.md#canonical-0231213330033102-3003110201202211-0202013213022303-0023000031003200-2101223332313210-0102312301311012-0013003223003233-0112100130310201) |
| `secondary.tsig_key_value.clear_secret_info` | [secondary.tsig_key_value.clear_secret_info](resources--dns_zone--reference--group-003.md#canonical-2023132111012003-0311221323013101-1122000230223032-2103123200220102-1023020011231121-3211122310021031-3232230201032312-0013130200310230) |
| `secondary.tsig_key_value.clear_secret_info.provider_ref` | [secondary.tsig_key_value.clear_secret_info.provider_ref](resources--dns_zone--reference--group-003.md#canonical-0313132220201333-2033222033331013-0030330303311112-0322211020220301-0102331021000110-3213030222322230-3021000030321003-0103110133023013) |
| `secondary.tsig_key_value.clear_secret_info.url` | [secondary.tsig_key_value.clear_secret_info.url](resources--dns_zone--reference--group-003.md#canonical-2222023222322022-1101121232133022-3110222320132113-3212123330200313-3031101201201313-0210120111300132-2003201020113203-2113012332131101) |
| `timeouts` | [timeouts](resources--dns_zone--reference--group-003.md#canonical-0300320321333023-0310332010031331-3031222113133311-3100030302022003-3212120202021302-0032200302202012-1100200031002133-3222120010110101) |
| `timeouts.create` | [timeouts.create](resources--dns_zone--reference--group-003.md#canonical-0300131231131222-2103011332202020-0301112203313120-0211010303001323-2310331101022332-1211200311113120-0121002013210222-0231323222020212) |
| `timeouts.delete` | [timeouts.delete](resources--dns_zone--reference--group-003.md#canonical-3321101022212312-1233211320310103-0013133032000303-0231010331321222-0213012223111311-0323213230221330-1133220321110110-2203021323101221) |
| `timeouts.read` | [timeouts.read](resources--dns_zone--reference--group-003.md#canonical-0321030200103001-1021031232022232-1030103103022221-2103022311101312-3111010301003211-1221312222231132-3121023010300312-3023202012113020) |
| `timeouts.update` | [timeouts.update](resources--dns_zone--reference--group-003.md#canonical-3323010230210303-0302232133303033-0311231221113031-1231331231312321-3203133220303021-3331110110321332-3232112303121310-3220332113032321) |

<a id="canonical-0323001130230330-0220112132203312-3330221323211022-1020331113222000-0103022210101223-0212301030333323-0203322130201001-2110331233130010"></a>

## Next pages — Property reference / 312030201022 / 12

- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-2101113233122202-1230223131033031-2301310003322033-0133201323022132-1232110221020011-1332203122201231-3233232102003032-1103322203113021)
- [timeouts](resources--dns_zone--reference--group-003.md#canonical-1121023120323203-0300301033303311-3210213210020103-1313032000222330-2030312122211110-2210030103233231-0231221302231331-3010100002112333)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211101230332131-0103320122123203-0021313231232100-3132013320113221-3030030001201001-0102202232331201-2100012221200200-1232221120210103"></a>

## primary — primary / 301112311202 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- primary

<a id="canonical-3300313313010111-0200031323101110-3311333102220103-0321130221302332-0222220031013301-1130233223202231-2200013222231201-1332310110100311"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: primary, secondary\] PrimaryDNSCreateSpecType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_soa_parameters",
    "soa_parameters")}
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
  "x-ves-oneof-field-soa_record_parameters_choice": "[\"default_soa_parameters\",\"soa_parameters\"]"
}
```

OneOf alternatives in this subsection:

- [primary](resources--dns_zone--reference--group-001.md#canonical-3300313313010111-0200031323101110-3311333102220103-0321130221302332-0222220031013301-1130233223202231-2200013222231201-1332310110100311)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-0223010131102331-0310121320000300-3131133100321303-1312200031031101-2101222213223300-3110301303132223-1023132100303123-3012223223200121)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
primary {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223310330220103-0331103120113133-0111313201330233-3330211232332020-1131033302322201-2301002312110302-3221212111302211-2001030131002101"></a>

## Direct properties — primary / 301112311202 / 3

<a id="canonical-1120201111333222-0103222323322112-3012012212033121-2122223222121000-2202311002021122-3110102220111302-3001123312013103-1113202111210201"></a>

<a id="canonical-3331031203022301-3311200311102103-0210200022121220-3333300012301103-3231302201230300-1232203312101112-3121031112322100-3212112133201302"></a>

## allow_http_lb_managed_records property — primary / 301112311202 / 4

Type: `"bool"`. Optional.

Option to allow user-created HTTP, TCP, and CDN load balancer related resource records to be
automatically managed in a protected RRset.

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

- [default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031): complete subsection reference.

- [default_soa_parameters](resources--dns_zone--reference--group-002.md#canonical-0001013031002232-2323220130321002-0110031123323021-1302303333101221-0212302022021323-3331121010231120-3131130311320003-3312131103003301): complete subsection reference.

- [dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102): complete subsection reference.

- [rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110): complete subsection reference.

- [soa_parameters](resources--dns_zone--reference--group-003.md#canonical-1020202022030312-2130233001011230-0123121101212023-0220103223030203-2021100102202011-2332311013003301-3000013203013302-0121111103000130): complete subsection reference.

<a id="canonical-0111003130212212-2131103311030301-3223111203000132-0102000022331333-3032300221101011-3113221120301033-2133030033322012-2123121211230310"></a>

## Next pages — primary / 301112311202 / 5

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_soa_parameters](resources--dns_zone--reference--group-002.md#canonical-0001013031002232-2323220130321002-0110031123323021-1302303333101221-0212302022021323-3331121010231120-3131130311320003-3312131103003301)
- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.soa_parameters](resources--dns_zone--reference--group-003.md#canonical-1020202022030312-2130233001011230-0123121101212023-0220103223030203-2021100102202011-2332311013003301-3000013203013302-0121111103000130)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012302022133131-2232300032011213-1230002121322021-1220013331223302-1122312313310330-0111103210022220-2301330013222331-3000110010121001"></a>

## primary.default_rr_set_group — default_rr_set_group / 111121301232 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- primary.default_rr_set_group

<a id="canonical-0212100320002333-0231332102121022-3122130003121010-2033020331213210-2100101301032130-1230002200300200-2122133210202203-2300130303131132"></a>

Type: `"object"`. list nested block, Optional.

Add and manage DNS resource record sets part of Default set group.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ttl"),
  validators.ConflictingListObjectAttributes("a_record",
    "aaaa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("tlsa_record",
    "txt_record")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50000,
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
    "ves.io.schema.rules.repeated.max_items": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  }
}
```

Terraform syntax:

```terraform
default_rr_set_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130031223100332-2333011222122112-3310112232112232-2233233200323032-3302112100213113-3211030330222311-3231103232212312-2302322110310000"></a>

## Direct properties — default_rr_set_group / 111121301232 / 3

- [a_record](resources--dns_zone--reference--group-001.md#canonical-3023010211102202-3313323303113130-0233003303333121-0201320312111120-2123031031303130-0312022032033333-2331202101220033-0131203310223123): complete subsection reference.

- [aaaa_record](resources--dns_zone--reference--group-001.md#canonical-1031330101033130-0302201202302020-2332001011130211-1020223031021101-1221322230230000-1221200331010310-3222300333231010-1232301202230330): complete subsection reference.

- [afsdb_record](resources--dns_zone--reference--group-001.md#canonical-2212212213003231-1312013213220131-2113100123320011-0233111201103103-3201203001121113-2022102303033101-2032211321000311-3102311032002231): complete subsection reference.

- [alias_record](resources--dns_zone--reference--group-001.md#canonical-1310110011113330-2230232211221303-3132001101231121-2112110123132002-1122320232200302-0203121123032310-3122312030112100-0330202102212001): complete subsection reference.

- [caa_record](resources--dns_zone--reference--group-001.md#canonical-0302331301302132-3312021202313230-0103122221113000-1331011210010120-3013121311301012-2300231121320230-3132203300001303-0133013331031331): complete subsection reference.

- [cds_record](resources--dns_zone--reference--group-001.md#canonical-1321103321010022-3321332102130111-3301230001300003-3110233320030032-2033002001201233-2232033102203003-3120321323002330-0111113031220313): complete subsection reference.

- [cert_record](resources--dns_zone--reference--group-001.md#canonical-3003321210210000-1031301202332201-2131021313332210-1112031010202333-1321123133123000-2330321312312210-0133321220130230-1113102030120101): complete subsection reference.

- [cname_record](resources--dns_zone--reference--group-001.md#canonical-2032330000021121-3311023021112132-1120113313102232-3121212313333300-1010212230003122-0321211321302330-0320223133020322-2302030220001130): complete subsection reference.

<a id="canonical-2303233302011210-2010121211233220-0230012323323223-1210320120031121-2331230210002213-3102022030102122-1023300200231221-2021222333200031"></a>

<a id="canonical-3023031130220201-2131311332301032-0330130000011222-1223030022132232-0032032121332233-0112112231313120-3002220131322120-3232312220213231"></a>

## description_spec property — default_rr_set_group / 111121301232 / 4

Type: `"string"`. Optional.

Comment. Human-readable description text

- [ds_record](resources--dns_zone--reference--group-002.md#canonical-2221112200013312-2222002032211003-3222010212133221-2312211212122130-3033033011333033-1322303331020103-1003331022130023-3312202311121303): complete subsection reference.

- [eui48_record](resources--dns_zone--reference--group-002.md#canonical-3100001201120203-0312033322023013-2332030222311021-2122332202011202-1210221203131033-3222103132311001-0030203223010032-3022320213132211): complete subsection reference.

- [eui64_record](resources--dns_zone--reference--group-002.md#canonical-0310232010230002-2032333311200131-2210000030122303-1120200331023020-3101110012303221-2302222310112000-2232130122311323-0203010122113311): complete subsection reference.

- [lb_record](resources--dns_zone--reference--group-002.md#canonical-2000023011031312-2003213123011132-0300331211333233-1003131020211202-0213013000200212-3132103320112330-2130111300321023-0221122033031030): complete subsection reference.

- [loc_record](resources--dns_zone--reference--group-002.md#canonical-0001331322032313-1331233300221231-2021032123331122-1313301111122120-1113123132222021-0013311333002210-3021113111120010-3121010323223333): complete subsection reference.

- [mx_record](resources--dns_zone--reference--group-002.md#canonical-1002110103301213-0230131110201331-3223012003111321-0300033220002210-3000130132020202-0011011200330101-0120220030322210-3213201300203312): complete subsection reference.

- [naptr_record](resources--dns_zone--reference--group-002.md#canonical-0221201201321311-3033210312022203-0010211233110120-0301032022010211-2322320001313111-2002020120203123-1333000323301120-3330303320333302): complete subsection reference.

- [ns_record](resources--dns_zone--reference--group-002.md#canonical-1320012000010221-1013101202011330-0222031102033131-0333201222231300-2310103201032220-3222020113330211-0230001303312231-1302222302130302): complete subsection reference.

- [ptr_record](resources--dns_zone--reference--group-002.md#canonical-3030100213211033-2113011011200322-2233222032221031-1220020031011201-1132132112023130-3211231011202332-2123323330002213-2200232102012230): complete subsection reference.

- [srv_record](resources--dns_zone--reference--group-002.md#canonical-0113001230120133-2303121012102031-1200311122002232-3032331232033323-1200010123122221-1212320123121123-0013213201323320-3313313211120233): complete subsection reference.

- [sshfp_record](resources--dns_zone--reference--group-002.md#canonical-0002133302013323-3312120311230021-3113212100312321-1000222020201120-3100200121220311-1202130100221130-3332302111201101-1102012023123221): complete subsection reference.

- [tlsa_record](resources--dns_zone--reference--group-002.md#canonical-1110210111012102-2031202010320123-1132320121232332-2002023031230032-2022213032232233-1212213101223203-1103321022220000-0330112323013122): complete subsection reference.

<a id="canonical-0031331110203210-2300321011113121-0032002133003202-2000013102013230-2310132031322131-0211231132303121-3301102001131021-0130311220013123"></a>

<a id="canonical-0311011303003203-1221112233102021-1231313222222200-0021302331121111-1312203101200121-2112131203232231-1322013231132332-0301122123000020"></a>

## ttl property — default_rr_set_group / 111121301232 / 5

Type: `"number"`. Optional.

Time to live. Time-to-live duration in seconds

Upstream description:

Time-to-live duration in seconds

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(60, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

- [txt_record](resources--dns_zone--reference--group-002.md#canonical-3331231312030132-3010311201030021-0330032020003003-1111312303203000-3232301002223130-3103200221113020-1202122130031211-1301211330201120): complete subsection reference.

<a id="canonical-3320211132310123-3030200213213101-0223021222122222-3332211202201011-3010223131122120-1021223200323320-2222000033020033-0211231103101323"></a>

## Next pages — default_rr_set_group / 111121301232 / 6

- [primary.default_rr_set_group.a_record](resources--dns_zone--reference--group-001.md#canonical-3023010211102202-3313323303113130-0233003303333121-0201320312111120-2123031031303130-0312022032033333-2331202101220033-0131203310223123)
- [primary.default_rr_set_group.aaaa_record](resources--dns_zone--reference--group-001.md#canonical-1031330101033130-0302201202302020-2332001011130211-1020223031021101-1221322230230000-1221200331010310-3222300333231010-1232301202230330)
- [primary.default_rr_set_group.afsdb_record](resources--dns_zone--reference--group-001.md#canonical-2212212213003231-1312013213220131-2113100123320011-0233111201103103-3201203001121113-2022102303033101-2032211321000311-3102311032002231)
- [primary.default_rr_set_group.alias_record](resources--dns_zone--reference--group-001.md#canonical-1310110011113330-2230232211221303-3132001101231121-2112110123132002-1122320232200302-0203121123032310-3122312030112100-0330202102212001)
- [primary.default_rr_set_group.caa_record](resources--dns_zone--reference--group-001.md#canonical-0302331301302132-3312021202313230-0103122221113000-1331011210010120-3013121311301012-2300231121320230-3132203300001303-0133013331031331)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-1321103321010022-3321332102130111-3301230001300003-3110233320030032-2033002001201233-2232033102203003-3120321323002330-0111113031220313)
- [primary.default_rr_set_group.cert_record](resources--dns_zone--reference--group-001.md#canonical-3003321210210000-1031301202332201-2131021313332210-1112031010202333-1321123133123000-2330321312312210-0133321220130230-1113102030120101)
- [primary.default_rr_set_group.cname_record](resources--dns_zone--reference--group-001.md#canonical-2032330000021121-3311023021112132-1120113313102232-3121212313333300-1010212230003122-0321211321302330-0320223133020322-2302030220001130)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-002.md#canonical-2221112200013312-2222002032211003-3222010212133221-2312211212122130-3033033011333033-1322303331020103-1003331022130023-3312202311121303)
- [primary.default_rr_set_group.eui48_record](resources--dns_zone--reference--group-002.md#canonical-3100001201120203-0312033322023013-2332030222311021-2122332202011202-1210221203131033-3222103132311001-0030203223010032-3022320213132211)
- [primary.default_rr_set_group.eui64_record](resources--dns_zone--reference--group-002.md#canonical-0310232010230002-2032333311200131-2210000030122303-1120200331023020-3101110012303221-2302222310112000-2232130122311323-0203010122113311)
- [primary.default_rr_set_group.lb_record](resources--dns_zone--reference--group-002.md#canonical-2000023011031312-2003213123011132-0300331211333233-1003131020211202-0213013000200212-3132103320112330-2130111300321023-0221122033031030)
- [primary.default_rr_set_group.loc_record](resources--dns_zone--reference--group-002.md#canonical-0001331322032313-1331233300221231-2021032123331122-1313301111122120-1113123132222021-0013311333002210-3021113111120010-3121010323223333)
- [primary.default_rr_set_group.mx_record](resources--dns_zone--reference--group-002.md#canonical-1002110103301213-0230131110201331-3223012003111321-0300033220002210-3000130132020202-0011011200330101-0120220030322210-3213201300203312)
- [primary.default_rr_set_group.naptr_record](resources--dns_zone--reference--group-002.md#canonical-0221201201321311-3033210312022203-0010211233110120-0301032022010211-2322320001313111-2002020120203123-1333000323301120-3330303320333302)
- [primary.default_rr_set_group.ns_record](resources--dns_zone--reference--group-002.md#canonical-1320012000010221-1013101202011330-0222031102033131-0333201222231300-2310103201032220-3222020113330211-0230001303312231-1302222302130302)
- [primary.default_rr_set_group.ptr_record](resources--dns_zone--reference--group-002.md#canonical-3030100213211033-2113011011200322-2233222032221031-1220020031011201-1132132112023130-3211231011202332-2123323330002213-2200232102012230)
- [primary.default_rr_set_group.srv_record](resources--dns_zone--reference--group-002.md#canonical-0113001230120133-2303121012102031-1200311122002232-3032331232033323-1200010123122221-1212320123121123-0013213201323320-3313313211120233)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-0002133302013323-3312120311230021-3113212100312321-1000222020201120-3100200121220311-1202130100221130-3332302111201101-1102012023123221)
- [primary.default_rr_set_group.tlsa_record](resources--dns_zone--reference--group-002.md#canonical-1110210111012102-2031202010320123-1132320121232332-2002023031230032-2022213032232233-1212213101223203-1103321022220000-0330112323013122)
- [primary.default_rr_set_group.txt_record](resources--dns_zone--reference--group-002.md#canonical-3331231312030132-3010311201030021-0330032020003003-1111312303203000-3232301002223130-3103200221113020-1202122130031211-1301211330201120)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3023010211102202-3313323303113130-0233003303333121-0201320312111120-2123031031303130-0312022032033333-2331202101220033-0131203310223123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001001301030221-0030103230133023-2101330302312300-2333133200311110-3222201203302332-3001210321331011-1110010303233101-1203130322000012"></a>

## primary.default_rr_set_group.a_record — a_record / 323202310010 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.a_record

<a id="canonical-2330110111222012-3013233111132012-3010113113203222-1103221331220101-2131101120310020-3032311210302120-2302131010023023-3130112130330322"></a>

Type: `"object"`. single nested block, Optional.

DNSAResourceRecord. A Records

Upstream description:

A Records

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
a_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130110332330311-2121130203011132-1331030101132323-1110023102001112-1011222001232021-1130321010230212-2200003122312100-3201123323011231"></a>

## Direct properties — a_record / 323202310010 / 3

<a id="canonical-2120011132310033-2230100110101001-2122000133220203-2121110331330003-1030221322311103-0123110011233311-1211223133311103-3122111231302221"></a>

<a id="canonical-0230201300113220-2322100130110303-3333131331133223-0032232013132130-3022230111330331-0123023330020002-0023201220110232-0303123330202030"></a>

## name property — a_record / 323202310010 / 4

Type: `"string"`. Optional.

Record name, please provide only the specific subdomain or record name without the base domain.

Upstream description:

A Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-3330122023302310-2010022113313222-0333001310231331-1013132210032303-3020213031131001-2010231112123101-3022200331122220-2102111311232133"></a>

<a id="canonical-2332301211303111-0021202013300310-1333303310331303-1323110333202222-0121021010333312-1201020132012223-0103010012301312-2331310123210030"></a>

## values property — a_record / 323202310010 / 5

Type: `["list", "string"]`. Optional.

IPv4 Addresses. A valid IPv4 address, for example: 192.0.2.242.

Upstream description:

A valid IPv4 address, for example: 192.0.2.242.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

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
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2302132013031212-2223100002011112-3002123123232322-3010000102223102-1110102232200032-2232311313011220-2132130002103121-0210130323331013"></a>

## Next pages — a_record / 323202310010 / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1031330101033130-0302201202302020-2332001011130211-1020223031021101-1221322230230000-1221200331010310-3222300333231010-1232301202230330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011000130023320-3133220321033012-3012330010112223-1313233031030223-1201331031231123-2011123001300330-0203132212321031-3113332031321103"></a>

## primary.default_rr_set_group.aaaa_record — aaaa_record / 131120001222 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.aaaa_record

<a id="canonical-0232201120223002-2232233203100032-0031201302003112-0113012310031201-3232021031123113-1311033232013201-2232113012111301-2100032213311333"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aaaa record.

Upstream description:

RecordSet for AAAA Records.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
aaaa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232302232212003-0321212323113001-2202132132033220-2010132003320320-3112021313002003-0023320210213102-2010022300320303-2111321120110332"></a>

## Direct properties — aaaa_record / 131120001222 / 3

<a id="canonical-1220302033330220-1123320310213220-0311320300102313-1231222003123230-1210110333032033-3103101233302302-3130232212010003-2303002123232110"></a>

<a id="canonical-2002311013311201-1323202300101332-0101313333303120-3313330001031000-3032021112331000-0312120113310132-3330123321320022-1010000112212313"></a>

## name property — aaaa_record / 131120001222 / 4

Type: `"string"`. Optional.

AAAA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-3201100002112022-1312232123211001-2001311111233032-2130110131013121-2202303333133002-3100122322332201-0300210320300122-2300201321121212"></a>

<a id="canonical-1111031203320220-2300232121032123-1311012230201321-3112210231202311-0022221220303202-0022131100230101-2011222100120303-1102020022330313"></a>

## values property — aaaa_record / 131120001222 / 5

Type: `["list", "string"]`. Optional.

IPv6 Addresses. A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Upstream description:

A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0131131223332203-2102221111213203-3230220012003011-3320230022112203-1212221121112012-2100000011001020-2300122131030021-3222230323003210"></a>

## Next pages — aaaa_record / 131120001222 / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2212212213003231-1312013213220131-2113100123320011-0233111201103103-3201203001121113-2022102303033101-2032211321000311-3102311032002231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021131210212200-2013022222310303-2122333021202103-3123110231213120-0022102233331311-3311211120000020-0213133230311031-1322031302323121"></a>

## primary.default_rr_set_group.afsdb_record — afsdb_record / 311220330300 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.afsdb_record

<a id="canonical-3332112233032323-1130123121203113-1300031201211222-2002330110323032-3331330112333210-0103212313223320-0301331111112020-2133313203211310"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for afsdb record.

Upstream description:

DNS AFSDB Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
afsdb_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100110030230200-0311113232123100-2112222032032120-3213311202232010-0133212311003020-3323223321321121-1233120031313230-2211211332321232"></a>

## Direct properties — afsdb_record / 311220330300 / 3

<a id="canonical-1210121323331023-2323212101132211-0110201001203313-3021231310320300-3021313131333301-1211330020200301-0020120211131023-0200312113233002"></a>

<a id="canonical-3031103310212120-3112112302021102-3021013100121221-3022012213311333-3112133113130233-3222033310012221-3330300033003223-2212301001302001"></a>

## name property — afsdb_record / 311220330300 / 4

Type: `"string"`. Optional.

AFSDB Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-001.md#canonical-3331130300030032-1310322011013310-3001312220210222-2001323131112230-0300231323022033-0010031111210201-0202000113010313-0012110301010110): complete subsection reference.

<a id="canonical-2231130112230302-0110210321222221-3210303330100010-1312002011123303-2223311310132301-1023121311003300-1130020102011212-3103133203323001"></a>

## Next pages — afsdb_record / 311220330300 / 5

- [primary.default_rr_set_group.afsdb_record.values](resources--dns_zone--reference--group-001.md#canonical-3331130300030032-1310322011013310-3001312220210222-2001323131112230-0300231323022033-0010031111210201-0202000113010313-0012110301010110)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3331130300030032-1310322011013310-3001312220210222-2001323131112230-0300231323022033-0010031111210201-0202000113010313-0012110301010110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000002202213101-2223310100221102-0011031311202221-0121312032130112-2130221032200030-0112321220333211-0121123003110202-0032331022010230"></a>

## primary.default_rr_set_group.afsdb_record.values — values / 331301110110 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.afsdb_record](resources--dns_zone--reference--group-001.md#canonical-2212212213003231-1312013213220131-2113100123320011-0233111201103103-3201203001121113-2022102303033101-2032211321000311-3102311032002231)
- primary.default_rr_set_group.afsdb_record.values

<a id="canonical-1110313211030132-1300320202200032-1011111132200023-1002020130222000-0003230203131023-1131230011212010-2321203100203020-0323303203102332"></a>

Type: `"object"`. list nested block, Optional.

AFSDB Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("hostname")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210012330033200-0232310131130321-2023212310033121-2301110032112311-3311032100023210-2013300021011333-3333130110012300-1333010210221103"></a>

## Direct properties — values / 331301110110 / 3

<a id="canonical-1211202222013210-3320211100213333-1331010123103222-0301000210111330-1222033333013100-2300110213003213-3300102231311233-2321231012312021"></a>

<a id="canonical-0020020031220033-3202100223330323-2111232113013130-3213122112333213-1220323132010110-3000200230001310-3100131231021111-1331222112131101"></a>

## hostname property — values / 331301110110 / 4

Type: `"string"`. Optional.

Server name of the AFS cell database server or the DCE name server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3132203033333121-3011120111133323-2200233101013102-0321121000211031-0232222002003220-1203012101021112-0002000132232333-2332133301331130"></a>

<a id="canonical-0331113100323010-2012221101113100-0201102332311032-1200222130121202-1231013230112023-1131000232121311-2020112301213232-0012011021002002"></a>

## subtype property — values / 331301110110 / 5

Type: `"string"`. Optional.

\[Enum: NONE|AFSVolumeLocationServer|DCEAuthenticationServer\] AFS Volume Location Server or DCE
Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server -
DCEAuthenticationServer: DCE Authentication Server. Possible values are \`NONE\`,
\`AFSVolumeLocationServer\`, \`DCEAuthenticationServer\`.

Upstream description:

AFS Volume Location Server or DCE Authentication Server.

&#8203;- NONE: NONE

&#8203;- AFSVolumeLocationServer: AFS Volume Location Server

&#8203;- DCEAuthenticationServer: DCE Authentication Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2301032232110112-2031212322003111-1003032312330210-1323003312213000-0332332233312000-1133230031110201-2010103003030222-2321322002121311"></a>

## Next pages — values / 331301110110 / 6

- [primary.default_rr_set_group.afsdb_record](resources--dns_zone--reference--group-001.md#canonical-2212212213003231-1312013213220131-2113100123320011-0233111201103103-3201203001121113-2022102303033101-2032211321000311-3102311032002231)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1310110011113330-2230232211221303-3132001101231121-2112110123132002-1122320232200302-0203121123032310-3122312030112100-0330202102212001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003102121122200-3233013023231023-3222000312221220-2110200132133111-2013211112010010-3123113212123112-1212222012013313-2001230331211300"></a>

## primary.default_rr_set_group.alias_record — alias_record / 210213012230 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.alias_record

<a id="canonical-1122033020033331-3230212110322231-3312001030101011-3200131120200223-3202221221303020-3313132200130032-1103130221101212-0223020320123220"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for alias record.

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
alias_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020322212032003-2222013030103210-1220332133331032-3213203131003202-2030021022011220-1302203132330221-1201100121012222-1030313001132000"></a>

## Direct properties — alias_record / 210213012230 / 3

<a id="canonical-2112031333313310-1232102031201202-3102110310112321-1112231000120111-2000022313320022-1233011103223020-3211311023020031-3220110003312033"></a>

<a id="canonical-0132133102131023-1310333130120300-1113020332013033-3012003000211023-3030311130020320-2010230222223303-2222320000101011-2000012022200132"></a>

## value property — alias_record / 210213012230 / 4

Type: `"string"`. Optional.

Domain. A valid domain name, for example: example.com.

Upstream description:

A valid domain name, for example: example.com.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-1303220130023202-0000211130311222-3212023222300302-1233013113003220-3210122313202111-2300233233120203-3103010013012002-3300021203330201"></a>

## Next pages — alias_record / 210213012230 / 5

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0302331301302132-3312021202313230-0103122221113000-1331011210010120-3013121311301012-2300231121320230-3132203300001303-0133013331031331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203302320123032-3233021321231323-1222032232230030-3211323111223030-2310331001003212-3020123031111202-2000033112003202-0300033201332111"></a>

## primary.default_rr_set_group.caa_record — caa_record / 323022322312 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.caa_record

<a id="canonical-2320311010310232-2013131210222123-0230210012200013-0300323011131132-0010001331133232-1202021103121111-0201123020120231-0123001113221122"></a>

Type: `"object"`. single nested block, Optional.

DNSCAAResourceRecord.

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
caa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213222112131001-1302322210210103-3131032232021200-1033203003102101-1031111003322121-3013133322222121-3201322030021110-0231022303300000"></a>

## Direct properties — caa_record / 323022322312 / 3

<a id="canonical-2313132303023023-3322003133323130-2331231301210210-0010100330122212-0222130212310031-3110332333011320-3321222003131223-3111112112313320"></a>

<a id="canonical-3213101112202121-1122032303012301-2302300320201313-0021211120111313-2023012310202000-2120103231330212-3203323123002321-0220032111212011"></a>

## name property — caa_record / 323022322312 / 4

Type: `"string"`. Optional.

CAA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-001.md#canonical-0323333130111320-0113011320030021-3130002132133023-0021030230003002-2310030310133132-2223211333030301-1023130220021312-2000013103033322): complete subsection reference.

<a id="canonical-2003023323121300-0013301123212212-1211031231210113-3302301323031102-3310333221313232-1200111120321330-3102130033333100-2132102321033203"></a>

## Next pages — caa_record / 323022322312 / 5

- [primary.default_rr_set_group.caa_record.values](resources--dns_zone--reference--group-001.md#canonical-0323333130111320-0113011320030021-3130002132133023-0021030230003002-2310030310133132-2223211333030301-1023130220021312-2000013103033322)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0323333130111320-0113011320030021-3130002132133023-0021030230003002-2310030310133132-2223211333030301-1023130220021312-2000013103033322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011300320312100-1202122330020031-0013322133023021-2223032020002201-3330300000210020-1020323231231303-3001312103223022-0021220000120122"></a>

## primary.default_rr_set_group.caa_record.values — values / 312013310122 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.caa_record](resources--dns_zone--reference--group-001.md#canonical-0302331301302132-3312021202313230-0103122221113000-1331011210010120-3013121311301012-2300231121320230-3132203300001303-0133013331031331)
- primary.default_rr_set_group.caa_record.values

<a id="canonical-1210332311131222-2021220113323120-3202100220120002-2233110321013220-3033331030302103-1133300212103133-2322211103203010-3110011200210233"></a>

Type: `"object"`. list nested block, Optional.

CAA Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321020222222011-0113230120332113-1032233231120313-0331213130131012-2231000121212312-2212333321100330-3100213220230311-3010021100130031"></a>

## Direct properties — values / 312013310122 / 3

<a id="canonical-1133223100312020-2320121020310102-0313223122123123-0020313302312221-3222112222330331-0002300331301321-3332002200123310-1120221002311213"></a>

<a id="canonical-2110212312023123-3020123331230100-2002310232321123-1031000121023033-2132201302220232-0013002323111022-0120022303032313-0233001312112103"></a>

## flags property — values / 312013310122 / 4

Type: `"number"`. Optional.

Flag should be an integer between 0 and 255.

Upstream description:

This flag should be an integer between 0 and 255.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-2132230321323312-2011201212233300-3301331300000023-2011122311300122-2120333123102112-3331011322030011-1211210221203011-1102013112130033"></a>

<a id="canonical-3022301333313033-0322011131012030-3101332003310123-2021302101000031-0022230003231033-1132111112313102-0023010030221213-2013131033112211"></a>

## tag property — values / 312013310122 / 5

Type: `"string"`. Optional.

\[Enum: issue|issuewild|iodef\] Tag. Tag for categorization and filtering. Possible values are
\`issue\`, \`issuewild\`, \`iodef\`.

Upstream description:

Tag for categorization and filtering

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("issue",
    "issuewild",
    "iodef"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "issue",
    "issuewild",
    "iodef"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  }
}
```

<a id="canonical-3111300130213212-1021201101023232-2101213201122101-2230112133002022-2020200102013031-1231032330132102-3233131332310021-2211032123212301"></a>

<a id="canonical-1013331222003121-3202230222120212-3231202032003022-0000223032021021-0303331202030000-0222133011103102-0321212022213123-2213102001223302"></a>

## value property — values / 312013310122 / 6

Type: `"string"`. Optional.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3120023201311121-3233332103333203-0122030333202132-1023300221013121-1020120033322023-1121100313332200-0020032332101113-1023012120332303"></a>

## Next pages — values / 312013310122 / 7

- [primary.default_rr_set_group.caa_record](resources--dns_zone--reference--group-001.md#canonical-0302331301302132-3312021202313230-0103122221113000-1331011210010120-3013121311301012-2300231121320230-3132203300001303-0133013331031331)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1321103321010022-3321332102130111-3301230001300003-3110233320030032-2033002001201233-2232033102203003-3120321323002330-0111113031220313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213111303123213-1033031330221121-1023300321231112-3100310333033333-1000222311311103-2231300332000021-1220033121001120-0232221232212312"></a>

## primary.default_rr_set_group.cds_record — cds_record / 013020231031 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.cds_record

<a id="canonical-0013211201021311-1311001010312112-2213333033202001-0132320223103001-3320233031011102-1301030001320120-3133132113100313-0320120323002313"></a>

Type: `"object"`. single nested block, Optional.

DNS CDS Record. DNS CDS Record.

Upstream description:

DNS CDS Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
cds_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023132231031321-1300313231101112-0002101223322302-0130303221022023-2220021130210101-2101222330000022-1031111221030211-0113312223301322"></a>

## Direct properties — cds_record / 013020231031 / 3

<a id="canonical-2102220020323133-3000321130101002-1113002303030223-2223103122032120-3020202031231231-0011331313010322-2230032021213313-3123210013320201"></a>

<a id="canonical-1310020011020200-1223111100131323-3030321330033133-1100010033300322-2331221320303101-1300202030212323-0231120102220333-1332333123210132"></a>

## name property — cds_record / 013020231031 / 4

Type: `"string"`. Optional.

CDS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-001.md#canonical-2201100232102321-3113213101202233-0231023310102230-2212203321131102-0100101230003111-1320323312013302-3102330102300203-3311231130101102): complete subsection reference.

<a id="canonical-0223323312032113-3330012111031222-3221232111131333-1033211000131332-3301330322013301-1320221320011022-3011202022022200-0200211223331002"></a>

## Next pages — cds_record / 013020231031 / 5

- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-2201100232102321-3113213101202233-0231023310102230-2212203321131102-0100101230003111-1320323312013302-3102330102300203-3311231130101102)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2201100232102321-3113213101202233-0231023310102230-2212203321131102-0100101230003111-1320323312013302-3102330102300203-3311231130101102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020021110212202-3312130302333311-0301232312130330-3002313001110333-2200202031202122-3003222003100103-1221100103133321-1213112132333302"></a>

## primary.default_rr_set_group.cds_record.values — values / 120002201130 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-1321103321010022-3321332102130111-3301230001300003-3110233320030032-2033002001201233-2232033102203003-3120321323002330-0111113031220313)
- primary.default_rr_set_group.cds_record.values

<a id="canonical-3020032231002122-1312133211211100-3201011133122012-1333000002231210-0112232302132132-0332331112301123-0020013313330321-3132222322321121"></a>

Type: `"object"`. list nested block, Optional.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key_tag"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha256_digest"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha384_digest"),
  validators.ConflictingListObjectAttributes("sha256_digest",
    "sha384_digest")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102112213011101-0133020102233020-0200222130111030-0023212231212311-2312133213202132-1113131113212232-2033020233203331-1020302120103103"></a>

## Direct properties — values / 120002201130 / 3

<a id="canonical-3303333101110033-2133132031220220-1103213220100010-3233321021320112-3110123111313201-1321331211201202-3030123321030330-3120130310311122"></a>

<a id="canonical-3121300303012303-0102313322102001-2233003110223100-1113030003333122-3233110113003211-1012102122311320-3103332210320220-0233123221310332"></a>

## ds_key_algorithm property — values / 120002201130 / 4

Type: `"string"`. Optional.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key-value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1101300032130123-0333302133001233-1211230103001231-1033110300002301-0321312203020233-1311302233222010-2131310020130200-1221203221330320"></a>

<a id="canonical-0311220321301333-1302213313223310-3131001110200011-0021201202221230-2110001112002000-3321331312023230-1203230023101111-2120223121122313"></a>

## key_tag property — values / 120002201130 / 5

Type: `"number"`. Optional.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

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

- [sha1_digest](resources--dns_zone--reference--group-001.md#canonical-0111101322302303-2303030123103220-0123132300330300-2233213222321122-2022110002300012-2313222202020213-2111122333202033-3031020133202321): complete subsection reference.

- [sha256_digest](resources--dns_zone--reference--group-001.md#canonical-3022132203233302-2131230303203032-1021201100330012-1220002130003110-3130311103021331-1110200021122003-2112203020002200-1033210013210121): complete subsection reference.

- [sha384_digest](resources--dns_zone--reference--group-001.md#canonical-2202003010201311-3121022231022232-0101220022222231-3023211213031220-0011030030320100-3312100322021330-1313321232031302-2130113330232312): complete subsection reference.

<a id="canonical-0320300200323222-0123030213221130-1201333112213123-3223033123232032-0000033013231112-1203121300101310-1320132111121321-3101330023002100"></a>

## Next pages — values / 120002201130 / 6

- [primary.default_rr_set_group.cds_record.values.sha1_digest](resources--dns_zone--reference--group-001.md#canonical-0111101322302303-2303030123103220-0123132300330300-2233213222321122-2022110002300012-2313222202020213-2111122333202033-3031020133202321)
- [primary.default_rr_set_group.cds_record.values.sha256_digest](resources--dns_zone--reference--group-001.md#canonical-3022132203233302-2131230303203032-1021201100330012-1220002130003110-3130311103021331-1110200021122003-2112203020002200-1033210013210121)
- [primary.default_rr_set_group.cds_record.values.sha384_digest](resources--dns_zone--reference--group-001.md#canonical-2202003010201311-3121022231022232-0101220022222231-3023211213031220-0011030030320100-3312100322021330-1313321232031302-2130113330232312)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-1321103321010022-3321332102130111-3301230001300003-3110233320030032-2033002001201233-2232033102203003-3120321323002330-0111113031220313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0111101322302303-2303030123103220-0123132300330300-2233213222321122-2022110002300012-2313222202020213-2111122333202033-3031020133202321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201113223001221-2223312213113023-3032233201133312-3001132112033310-2322022232102210-1312120033013020-3022212011121201-3021133322232202"></a>

## primary.default_rr_set_group.cds_record.values.sha1_digest — sha1_digest / 002021222200 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-1321103321010022-3321332102130111-3301230001300003-3110233320030032-2033002001201233-2232033102203003-3120321323002330-0111113031220313)
- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-2201100232102321-3113213101202233-0231023310102230-2212203321131102-0100101230003111-1320323312013302-3102330102300203-3311231130101102)
- primary.default_rr_set_group.cds_record.values.sha1_digest

<a id="canonical-0321301303101121-2202032133032021-2310311222213113-3213302033300013-0312221013010200-0223312322201023-0312312002222130-3121013031121312"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha1_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312232002302012-3133110112212221-0302303303031020-2133212121131320-0122322033123320-1032330023031012-2033332000213132-2321100210011210"></a>

## Direct properties — sha1_digest / 002021222200 / 3

<a id="canonical-2321031021121300-2111210300001220-3201001133102223-0030230323002333-0013312102232010-1211010000313322-3103112111123022-0132101030200332"></a>

<a id="canonical-3231023232000003-1113111120100002-1112010020011331-2331311232133010-1311013231331113-1323110003330030-2232200020302032-1023023201023321"></a>

## digest property — sha1_digest / 002021222200 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-3131122230112130-0312331033010122-3102220112311030-3302120002012121-2331203203202313-3332202121132313-0132200022131031-2301123121231200"></a>

## Next pages — sha1_digest / 002021222200 / 5

- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-2201100232102321-3113213101202233-0231023310102230-2212203321131102-0100101230003111-1320323312013302-3102330102300203-3311231130101102)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3022132203233302-2131230303203032-1021201100330012-1220002130003110-3130311103021331-1110200021122003-2112203020002200-1033210013210121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311021312222331-0302120211022232-2320030030333132-2123103233233101-2031100121103111-2021312200211131-2222113022121301-2003201020200333"></a>

## primary.default_rr_set_group.cds_record.values.sha256_digest — sha256_digest / 322212321122 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-1321103321010022-3321332102130111-3301230001300003-3110233320030032-2033002001201233-2232033102203003-3120321323002330-0111113031220313)
- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-2201100232102321-3113213101202233-0231023310102230-2212203321131102-0100101230003111-1320323312013302-3102330102300203-3311231130101102)
- primary.default_rr_set_group.cds_record.values.sha256_digest

<a id="canonical-3333311301201110-3031312030313020-1302301100210210-3320102220331201-2212322323201120-1031132102223001-1200110233221332-2023232220201011"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha256_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000220231113003-0320033300120012-1300031002201203-3100212012331010-3202310020233230-3323122311022322-2111122231133322-3002033121021011"></a>

## Direct properties — sha256_digest / 322212321122 / 3

<a id="canonical-0321011012130120-3122133113310101-3113101321212133-2102220021233010-3220231130133301-1120023030023123-0220110102032032-1212101333203222"></a>

<a id="canonical-1301302312311003-2212313202320102-0002133332220300-2231232010000003-3133231330231210-3200202211123223-3310101321001101-1303100031123332"></a>

## digest property — sha256_digest / 322212321122 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
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
    "minLength": 64
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
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-3000010310320210-0102033233022002-0332112122331313-0220121202332230-1213203023033210-2022202131112030-2322122201333303-1031122303321100"></a>

## Next pages — sha256_digest / 322212321122 / 5

- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-2201100232102321-3113213101202233-0231023310102230-2212203321131102-0100101230003111-1320323312013302-3102330102300203-3311231130101102)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2202003010201311-3121022231022232-0101220022222231-3023211213031220-0011030030320100-3312100322021330-1313321232031302-2130113330232312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232231131321000-3122322320023221-2113111330001013-1201131020033233-0211320022023313-2332012221302022-3230001100003010-0013202300231201"></a>

## primary.default_rr_set_group.cds_record.values.sha384_digest — sha384_digest / 031200322123 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.cds_record](resources--dns_zone--reference--group-001.md#canonical-1321103321010022-3321332102130111-3301230001300003-3110233320030032-2033002001201233-2232033102203003-3120321323002330-0111113031220313)
- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-2201100232102321-3113213101202233-0231023310102230-2212203321131102-0100101230003111-1320323312013302-3102330102300203-3311231130101102)
- primary.default_rr_set_group.cds_record.values.sha384_digest

<a id="canonical-3013202332103313-0001221023330013-3220031300112200-0123033110121331-3011333120311230-0332102313313330-1331030121100332-2330310212302210"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha384 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha384_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101100231120322-1212221220202222-0020203113301101-1321310202222330-0101131011033031-1130021231002103-3320002311131313-0211230020312032"></a>

## Direct properties — sha384_digest / 031200322123 / 3

<a id="canonical-0310320203011011-1112012112001100-2012321100210310-0101231312300223-0020311003312203-2201023002110010-1203112000132210-2230323313221312"></a>

<a id="canonical-2100120222222110-0122313132102112-1323230332330202-2320122031101223-0010013011132210-3200131020312302-0230311202000102-3330310312211013"></a>

## digest property — sha384_digest / 031200322123 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(96, 96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-2021302130112231-2331202123323232-1100010201203013-0021200002302013-2310010021032312-0303333222322020-1020030120210210-1010110111001131"></a>

## Next pages — sha384_digest / 031200322123 / 5

- [primary.default_rr_set_group.cds_record.values](resources--dns_zone--reference--group-001.md#canonical-2201100232102321-3113213101202233-0231023310102230-2212203321131102-0100101230003111-1320323312013302-3102330102300203-3311231130101102)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3003321210210000-1031301202332201-2131021313332210-1112031010202333-1321123133123000-2330321312312210-0133321220130230-1113102030120101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122300230310203-1132021110213123-1300200112232322-1011320001200231-0221221120303231-1232232221102232-3201102200012122-1023331121202231"></a>

## primary.default_rr_set_group.cert_record — cert_record / 221022313021 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- primary.default_rr_set_group.cert_record

<a id="canonical-0311231110333010-2022331022011001-2200111021010022-0232133323223133-0231120022221010-3130133321311100-0133000100133310-0111111332003203"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for cert record.

Upstream description:

DNS CERT Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
cert_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021121012123111-2132303320233202-0122211320222300-0100110112130312-0310233022100303-0101232120120223-2110132103032312-1133130213303122"></a>

## Direct properties — cert_record / 221022313021 / 3

<a id="canonical-2230112122200022-2302233201001101-1011222133323100-0133011011010000-2031302323113013-2311102013313311-3213123233233300-2003120332022020"></a>

<a id="canonical-2201130001212020-1002320131313102-2133000133330202-1123220202231130-1101000011021020-3110120010110032-1111332122221332-1302212301131123"></a>

## name property — cert_record / 221022313021 / 4

Type: `"string"`. Optional.

CERT Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-001.md#canonical-2210301201030001-3220132100011323-2300032123130231-0330021320200310-3101201002002300-3230310310312102-0101031212133331-2313003313003131): complete subsection reference.

<a id="canonical-2233300222012133-2122111312320203-2030231312022333-2111211031111220-3130110232231011-2300000023031223-2333210003223221-3000133303203012"></a>

## Next pages — cert_record / 221022313021 / 5

- [primary.default_rr_set_group.cert_record.values](resources--dns_zone--reference--group-001.md#canonical-2210301201030001-3220132100011323-2300032123130231-0330021320200310-3101201002002300-3230310310312102-0101031212133331-2313003313003131)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2210301201030001-3220132100011323-2300032123130231-0330021320200310-3101201002002300-3230310310312102-0101031212133331-2313003313003131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313000300102022-3021031103130322-1113123233112311-2023111122111101-1222303221303101-3030231113003323-0032200211133032-0311111302020332"></a>

## primary.default_rr_set_group.cert_record.values — values / 033212300133 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-2120123011321113-3203023110131203-2112300301013101-2301022120121332-3110310321100200-3222110202203023-3022313201030213-1202102120011031)
- [primary.default_rr_set_group.cert_record](resources--dns_zone--reference--group-001.md#canonical-3003321210210000-1031301202332201-2131021313332210-1112031010202333-1321123133123000-2330321312312210-0133321220130230-1113102030120101)
- primary.default_rr_set_group.cert_record.values

<a id="canonical-1200312002202202-0233012102010301-1232322223110133-0031013322033021-2132113022022101-1010323300111113-2222111121232022-2101003021033031"></a>

Type: `"object"`. list nested block, Optional.

CERT Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("cert_key_tag",
    "certificate")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222320212033233-3213300101123131-2030230133012320-3320013203000320-0201012220030231-1002021213222203-2101230310020112-1030221001002133"></a>

## Direct properties — values / 033212300133 / 3

<a id="canonical-2122302332000202-1310102112310330-2230123231213121-1011222033320110-1132011323002110-3022303112123120-1311030220313322-0230112220023301"></a>

<a id="canonical-2022313120230330-1210231311130132-0123232010321212-2202132032011321-2110001133201313-2021131323332333-3302223222232210-0110011203021023"></a>

## algorithm property — values / 033212300133 / 4

Type: `"string"`. Optional.

\[Enum: RESERVEDALGORITHM|RSAMD5|DH|DSASHA1|ECC|RSASHA1ALGORITHM|INDIRECT|PRIVATEDNS|PRIVATEOID\]
CERT algorithm value must be compatible with the specified algorithm. - RESERVEDALGORITHM:
RESERVEDALGORITHM - RSAMD5: RSAMD5 - DH: DH - DSASHA1: DSASHA1 - ECC: ECC - RSASHA1ALGORITHM:
RSA-SHA1 - INDIRECT: INDIRECT - PRIVATEDNS: PRIVATEDNS - PRIVATEOID: PRIVATEOID. Possible values are
\`RESERVEDALGORITHM\`, \`RSAMD5\`, \`DH\`, \`DSASHA1\`, \`ECC\`, \`RSASHA1ALGORITHM\`, \`INDIRECT\`,
\`PRIVATEDNS\`, \`PRIVATEOID\`. Defaults to \`RESERVEDALGORITHM\`.

Upstream description:

CERT algorithm value must be compatible with the specified algorithm.

&#8203;- RESERVEDALGORITHM: RESERVEDALGORITHM

&#8203;- RSAMD5: RSAMD5

&#8203;- DH: DH

&#8203;- DSASHA1: DSASHA1

&#8203;- ECC: ECC

&#8203;- RSASHA1ALGORITHM: RSA-SHA1

&#8203;- INDIRECT: INDIRECT

&#8203;- PRIVATEDNS: PRIVATEDNS

&#8203;- PRIVATEOID: PRIVATEOID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "RESERVEDALGORITHM",
  "enum": [
    "RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0010131002132313-1133322112133000-1221131222322212-1111323023000001-3012320011223213-2312003231001221-0201331030032231-1000020332222223"></a>

<a id="canonical-1323131231231112-3133011003210010-3320321232322000-2223223010231122-1120132023303323-3320332333000120-3211120322100022-0003222010030000"></a>

## cert_key_tag property — values / 033212300133 / 5

Type: `"number"`. Optional.

Key Tag. Tag for categorization and filtering

Upstream description:

Tag for categorization and filtering

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3022210110200323-1130212222200321-2310333303013323-0321211101333120-1302020330030333-2103012033211112-3030201302012022-3113022202101332"></a>

<a id="canonical-3200001301210110-3320133313321102-2323021212122201-0200110112110032-3103021111322010-1001023000323133-2210013301120133-3203302030000112"></a>

## cert_type property — values / 033212300133 / 6

Type: `"string"`. Optional.

\[Enum: INVALIDCERTTYPE|PKIX|SPKI|PGP|IPKIX|ISPKI|IPGP|ACPKIX|IACPKIX|URI\_|OID\] CERT type value
must be compatible with the specified types. - INVALIDCERTTYPE: INVALIDCERTTYPE - PKIX: PKIX - SPKI:
SPKI - PGP: PGP - IPKIX: IPKIX - ISPKI: ISPKI - IPGP: IPGP - ACPKIX: ACPKIX - IACPKIX: IACPKIX -
URI\_: URI - OID: OID. Possible values are \`INVALIDCERTTYPE\`, \`PKIX\`, \`SPKI\`, \`PGP\`,
\`IPKIX\`, \`ISPKI\`, \`IPGP\`, \`ACPKIX\`, \`IACPKIX\`, \`URI\_\`, \`OID\`. Defaults to
\`INVALIDCERTTYPE\`.

Upstream description:

CERT type value must be compatible with the specified types.

&#8203;- INVALIDCERTTYPE: INVALIDCERTTYPE

&#8203;- PKIX: PKIX

&#8203;- SPKI: SPKI

&#8203;- PGP: PGP

&#8203;- IPKIX: IPKIX

&#8203;- ISPKI: ISPKI

&#8203;- IPGP: IPGP

&#8203;- ACPKIX: ACPKIX

&#8203;- IACPKIX: IACPKIX

&#8203;- URI\_: URI

&#8203;- OID: OID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALIDCERTTYPE",
  "enum": [
    "INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1130210313300023-2031001213111211-3112332303321012-3202212202003113-2121112231221202-2330232201132301-3200211233020032-0212320311300301"></a>

<a id="canonical-3200212213011313-2303213323322101-0111130130003200-3133110013101020-3021303003010131-3121023032301002-1203232103230212-0020202230013030"></a>

## certificate property — values / 033212300133 / 7

Type: `"string"`. Optional.

Certificate. Certificate in base 64 format.

Upstream description:

Certificate in base 64 format.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2021013222121011-2320121131001021-2010000101200222-1103323230323310-1023202232311322-1023321231211103-3133321002210022-2002003133021112"></a>

## Next pages — values / 033212300133 / 8

- [primary.default_rr_set_group.cert_record](resources--dns_zone--reference--group-001.md#canonical-3003321210210000-1031301202332201-2131021313332210-1112031010202333-1321123133123000-2330321312312210-0133321220130230-1113102030120101)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2032330000021121-3311023021112132-1120113313102232-3121212313333300-1010212230003122-0321211321302330-0320223133020322-2302030220001130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
