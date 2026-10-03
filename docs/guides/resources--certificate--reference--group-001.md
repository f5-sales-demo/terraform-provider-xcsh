---
page_title: "xcsh_certificate reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate reference."
---

# xcsh_certificate reference

<a id="canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032333300302201-3220332313112200-3003102333023113-0030202022003231-2102001101333333-1332121201202201-3033300012210132-1311321222133201"></a>

## Property reference — Property reference / 102233003113 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- Property reference

<a id="canonical-1320122003230231-2020100331123012-1003002022320313-3221013022023311-3203010210121023-2300113220110111-2111132022232301-2131210031111312"></a>

## Direct properties — Property reference / 102233003113 / 3

<a id="canonical-0320220311103002-1100000320010211-0023313322031110-0211201123101220-0102012203201200-2032211233103030-1330132322012212-3322131010100202"></a>

<a id="canonical-3210120012232312-2121330203222203-3112221010030003-2102130223200030-3223201113303133-2112203002310012-3333313113302203-1100010132330211"></a>

## annotations property — Property reference / 102233003113 / 4

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

- [certificate_chain](resources--certificate--reference--group-001.md#canonical-1030320000233232-2232302213110323-3310002333011103-3020003203301303-3332202000210013-1213222102122030-2002010110130123-2123301311203302): complete subsection reference.

<a id="canonical-0120232130213103-0001203133221331-2100132201133130-3333330022201311-3300233010333110-1031203021302101-2331022012031000-3303231032102113"></a>

<a id="canonical-0303112333001220-3301203022322322-3103030222102130-1300222301231111-2210022323303013-0221103023111221-3032302111301000-2122000333020022"></a>

## certificate_url property — Property reference / 102233003113 / 5

Type: `"string"`. Required.

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--certificate--reference--group-001.md#canonical-3131031131220113-3330221330321201-1023222220022220-0312222333120011-2110302312132303-2122013133211112-3110020100032323-3230021122333211): complete subsection reference.

<a id="canonical-2100113220220032-0013033300222122-1101030123020132-2012311302022211-0311012202232331-3022232331120112-0210101101313312-0100320213313020"></a>

<a id="canonical-3133302111000022-0232011011000000-1223321113333130-0122332210030301-1202202321013031-0010102021011133-3323331210123220-0100220310221301"></a>

## description property — Property reference / 102233003113 / 6

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

<a id="canonical-3302011201332201-1011130333020332-2000332013012220-3200032331010300-1330013110302301-1001133331223120-2023103331221011-0023313132321032"></a>

<a id="canonical-3221012323011333-2233222123232133-2231301323220010-0331121231212032-1000010120223211-0203033210021022-0113200010221331-1212003330313021"></a>

## disable property — Property reference / 102233003113 / 7

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

- [disable_ocsp_stapling](resources--certificate--reference--group-001.md#canonical-3232212033302331-1222133011022110-3212321322101213-0313001110013132-0330213212211122-0302330121213011-1313201320102231-0233302020223133): complete subsection reference.

<a id="canonical-0323000032312003-0203202331330222-3020123330112102-2113333223131110-1032223011103231-3220223110332311-2301003103222200-0122211113222131"></a>

<a id="canonical-0223332303002030-2231323322022331-1211303322023132-0202222111000123-3212130310012211-2002102320121231-1133020212223200-3023233311003300"></a>

## ID property — Property reference / 102233003113 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1031133332223333-1310130312111213-0332203121120233-2332033020233003-3100123323330031-2001322013023210-1023320203311003-2210322121100310"></a>

<a id="canonical-1200222320122220-3103312202211220-0103322303110220-0200221230301300-2112033301300302-0301010200320203-2113321302123211-0032131230021022"></a>

## labels property — Property reference / 102233003113 / 9

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

<a id="canonical-0133301113113232-1303121313102300-1010102233023201-0020120302311111-2332322122021030-2100103001331231-3331010222033103-0011100132212202"></a>

<a id="canonical-1201220031312011-2002202102010023-1213020031013303-1322200220320103-2323120203031110-2211130121101202-0020210320003122-2321332311022110"></a>

## name property — Property reference / 102233003113 / 10

Type: `"string"`. Required.

Name of the Certificate. Must be unique within the namespace.

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

<a id="canonical-0002231220003012-0111032313012131-0322312332223333-0022203321203111-3031323303133223-1132231231211330-0301020233320132-0332303332001313"></a>

<a id="canonical-2232102101030213-3032212100103112-1202021213220000-1331221221123020-1101302202302032-3312321202211112-3300001012310331-3121220310101202"></a>

## namespace property — Property reference / 102233003113 / 11

Type: `"string"`. Required.

Namespace where the Certificate is created.

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

- [private_key](resources--certificate--reference--group-001.md#canonical-2032012232223302-3313110001002321-0100312230130202-0202320002311022-3233121322320013-0213333220230021-1223231110210112-2132120202111322): complete subsection reference.

- [timeouts](resources--certificate--reference--group-001.md#canonical-0022303111210311-0212221001121021-2111110133110332-1103310031101303-1020011013221332-0130113102021020-1110332113310011-1103211100200222): complete subsection reference.

- [use_system_defaults](resources--certificate--reference--group-001.md#canonical-1011132030030013-2320300233313232-2321133121131212-2123120301103000-1003011033212130-0213210031100131-0330120313310313-1100200323301322): complete subsection reference.

<a id="canonical-2110233301200011-1032012001110212-3303023130310302-2002300001310131-2033032202303203-2230301312132023-3102010310222123-3032300123130122"></a>

## All schema paths — Property reference / 102233003113 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--certificate--reference--group-001.md#canonical-0320220311103002-1100000320010211-0023313322031110-0211201123101220-0102012203201200-2032211233103030-1330132322012212-3322131010100202) |
| `certificate_chain` | [certificate_chain](resources--certificate--reference--group-001.md#canonical-1330301200113122-3202132011031313-2233013030020112-1010133233032132-1131231131000232-0131233310311330-1213323111312102-1011302321023122) |
| `certificate_chain.name` | [certificate_chain.name](resources--certificate--reference--group-001.md#canonical-3100102310320022-2001321103213123-0332112213002323-2221231200312220-0202301212320121-1230121322331213-1232132021203012-0121013303100313) |
| `certificate_chain.namespace` | [certificate_chain.namespace](resources--certificate--reference--group-001.md#canonical-3303011203122322-1221332120333011-0312101111221311-2002220331020300-0201203012303303-2112130020110000-1111013200311313-1111031212033212) |
| `certificate_chain.tenant` | [certificate_chain.tenant](resources--certificate--reference--group-001.md#canonical-3121330331022102-2013220212222130-3332322302230102-1103233332120203-1331031100021020-3330200011030302-0310300020331320-0203230220201332) |
| `certificate_url` | [certificate_url](resources--certificate--reference--group-001.md#canonical-0120232130213103-0001203133221331-2100132201133130-3333330022201311-3300233010333110-1031203021302101-2331022012031000-3303231032102113) |
| `custom_hash_algorithms` | [custom_hash_algorithms](resources--certificate--reference--group-001.md#canonical-0332010030333003-3020030330020200-0323221203003331-1013103210232131-0222132321131330-1310101323302011-1300122110213332-2223123113131111) |
| `custom_hash_algorithms.hash_algorithms` | [custom_hash_algorithms.hash_algorithms](resources--certificate--reference--group-001.md#canonical-3110300333231130-3310010200131301-3023003021121333-3333230133221303-3131031032220301-3010012221131003-2013120100300213-0321023300310313) |
| `description` | [description](resources--certificate--reference--group-001.md#canonical-2100113220220032-0013033300222122-1101030123020132-2012311302022211-0311012202232331-3022232331120112-0210101101313312-0100320213313020) |
| `disable` | [disable](resources--certificate--reference--group-001.md#canonical-3302011201332201-1011130333020332-2000332013012220-3200032331010300-1330013110302301-1001133331223120-2023103331221011-0023313132321032) |
| `disable_ocsp_stapling` | [disable_ocsp_stapling](resources--certificate--reference--group-001.md#canonical-3112210023202301-0330301120233313-3300023132021120-3202331002133131-1010200312002010-3302112331202331-1120331210032120-1030211100333130) |
| `id` | [ID](resources--certificate--reference--group-001.md#canonical-0323000032312003-0203202331330222-3020123330112102-2113333223131110-1032223011103231-3220223110332311-2301003103222200-0122211113222131) |
| `labels` | [labels](resources--certificate--reference--group-001.md#canonical-1031133332223333-1310130312111213-0332203121120233-2332033020233003-3100123323330031-2001322013023210-1023320203311003-2210322121100310) |
| `name` | [name](resources--certificate--reference--group-001.md#canonical-0133301113113232-1303121313102300-1010102233023201-0020120302311111-2332322122021030-2100103001331231-3331010222033103-0011100132212202) |
| `namespace` | [namespace](resources--certificate--reference--group-001.md#canonical-0002231220003012-0111032313012131-0322312332223333-0022203321203111-3031323303133223-1132231231211330-0301020233320132-0332303332001313) |
| `private_key` | [private_key](resources--certificate--reference--group-001.md#canonical-0100130131110220-0011113301322303-1310022010313231-1132001300000030-2032203200130133-0102112301032301-1130003031303201-2001103211102333) |
| `private_key.blindfold_secret_info` | [private_key.blindfold_secret_info](resources--certificate--reference--group-001.md#canonical-2231111122010112-0121323323011231-2320123113013123-2320233213210120-2030020301002303-1112232200021020-2221102232222130-2120021011011330) |
| `private_key.blindfold_secret_info.decryption_provider` | [private_key.blindfold_secret_info.decryption_provider](resources--certificate--reference--group-001.md#canonical-2230002020113101-0010223013223001-0023222311313032-1001021123200303-0132202013121202-1031103013232231-2233211123202223-0231331231301021) |
| `private_key.blindfold_secret_info.location` | [private_key.blindfold_secret_info.location](resources--certificate--reference--group-001.md#canonical-1031022210300020-0331210101210101-0113301301311002-1320102133002130-3230212211213312-3221000011220122-3023312133112102-2313233121203232) |
| `private_key.blindfold_secret_info.store_provider` | [private_key.blindfold_secret_info.store_provider](resources--certificate--reference--group-001.md#canonical-1022312313300303-3330120320200312-1223210113301132-0010202320302300-1232311333132031-0001111002311220-3022333202321222-0020300021110202) |
| `private_key.clear_secret_info` | [private_key.clear_secret_info](resources--certificate--reference--group-001.md#canonical-0232100033021131-0200222301033031-1100332131030221-1133110310203222-1122211202113033-0303220221133302-0200212021332222-0210312222120202) |
| `private_key.clear_secret_info.provider_ref` | [private_key.clear_secret_info.provider_ref](resources--certificate--reference--group-001.md#canonical-1322302211223031-0333203031303322-1313121200221133-1121100002222013-3333002122121013-3231101302103311-0220323110133302-1311013332020313) |
| `private_key.clear_secret_info.url` | [private_key.clear_secret_info.url](resources--certificate--reference--group-001.md#canonical-3032023123012320-0111332113111321-2002131103110320-1213210132200100-0321200110221232-0212000300110103-0233111221211333-0002022300022330) |
| `timeouts` | [timeouts](resources--certificate--reference--group-001.md#canonical-0000303133321230-1012002213111110-0031031201201332-2222102332200222-3132331013210231-2301313033030213-0210202323230022-0130230330123212) |
| `timeouts.create` | [timeouts.create](resources--certificate--reference--group-001.md#canonical-2103123102303330-0320110110203203-0010002101020110-0122223020320211-0100313010001123-3122202321233220-0031021030311012-0022033101332000) |
| `timeouts.delete` | [timeouts.delete](resources--certificate--reference--group-001.md#canonical-1002220213101223-3320020313122203-2030320231313200-0303133310200030-0232223001021120-3010020033032301-3023333120220321-0310112011103333) |
| `timeouts.read` | [timeouts.read](resources--certificate--reference--group-001.md#canonical-1133311000020021-0311113100113302-1232321120310122-2330113330313303-2210311313202012-3232311300021212-3111222301201022-3323030232232123) |
| `timeouts.update` | [timeouts.update](resources--certificate--reference--group-001.md#canonical-1120011233003201-2000212210010112-0033011100022110-3231030033323133-3230321202231000-3233032332223033-2102210201101110-0231222103012330) |
| `use_system_defaults` | [use_system_defaults](resources--certificate--reference--group-001.md#canonical-3101230030230231-1320130120122112-2020012023230002-2003032110301202-1121311100130003-0001321111002231-2003020200001100-2312032130010111) |

<a id="canonical-0221123031222230-3000312201320122-0133103112302311-2002000010101211-2020120102321302-3133230111033000-2101101231021322-2210133302131012"></a>

## Next pages — Property reference / 102233003113 / 13

- [certificate_chain](resources--certificate--reference--group-001.md#canonical-1030320000233232-2232302213110323-3310002333011103-3020003203301303-3332202000210013-1213222102122030-2002010110130123-2123301311203302)
- [custom_hash_algorithms](resources--certificate--reference--group-001.md#canonical-3131031131220113-3330221330321201-1023222220022220-0312222333120011-2110302312132303-2122013133211112-3110020100032323-3230021122333211)
- [disable_ocsp_stapling](resources--certificate--reference--group-001.md#canonical-3232212033302331-1222133011022110-3212321322101213-0313001110013132-0330213212211122-0302330121213011-1313201320102231-0233302020223133)
- [private_key](resources--certificate--reference--group-001.md#canonical-2032012232223302-3313110001002321-0100312230130202-0202320002311022-3233121322320013-0213333220230021-1223231110210112-2132120202111322)
- [timeouts](resources--certificate--reference--group-001.md#canonical-0022303111210311-0212221001121021-2111110133110332-1103310031101303-1020011013221332-0130113102021020-1110332113310011-1103211100200222)
- [use_system_defaults](resources--certificate--reference--group-001.md#canonical-1011132030030013-2320300233313232-2321133121131212-2123120301103000-1003011033212130-0213210031100131-0330120313310313-1100200323301322)
- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)

<a id="canonical-1030320000233232-2232302213110323-3310002333011103-3020003203301303-3332202000210013-1213222102122030-2002010110130123-2123301311203302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010012321132012-1201100030230020-1201320202121111-3111111021122300-0022221110312023-2321031330200203-1132122123201213-0331000000023300"></a>

## certificate_chain — certificate_chain / 032123200230 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- certificate_chain

<a id="canonical-1330301200113122-3202132011031313-2233013030020112-1010133233032132-1131231131000232-0131233310311330-1213323111312102-1011302321023122"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
certificate_chain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330022201323033-1011132333211021-1221021321101302-0313001221211311-3133133331130210-0322231231231111-1030101102233112-2003002021302332"></a>

## Direct properties — certificate_chain / 032123200230 / 3

<a id="canonical-3100102310320022-2001321103213123-0332112213002323-2221231200312220-0202301212320121-1230121322331213-1232132021203012-0121013303100313"></a>

<a id="canonical-0321132231303210-1322303130323020-2112033201320133-1002022330012101-1232012112300320-0302323100233231-1222303322002130-3300303221203102"></a>

## name property — certificate_chain / 032123200230 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3303011203122322-1221332120333011-0312101111221311-2002220331020300-0201203012303303-2112130020110000-1111013200311313-1111031212033212"></a>

<a id="canonical-0313223232331211-1322211012220102-2012330001101323-3322312122003303-2003302010303121-2003120222013331-1303010232333221-0022321201132022"></a>

## namespace property — certificate_chain / 032123200230 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3121330331022102-2013220212222130-3332322302230102-1103233332120203-1331031100021020-3330200011030302-0310300020331320-0203230220201332"></a>

<a id="canonical-3221100331202111-0032111210303102-0102332000112120-2213202132232221-3230211001320232-3022031013131200-3310130331132001-0103012231203320"></a>

## tenant property — certificate_chain / 032123200230 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3332011211233110-3331322221211023-0132000121201103-0313330200313222-3110132323311110-1001110130231300-0210003030013012-2103103301012332"></a>

## Next pages — certificate_chain / 032123200230 / 7

- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)

<a id="canonical-3131031131220113-3330221330321201-1023222220022220-0312222333120011-2110302312132303-2122013133211112-3110020100032323-3230021122333211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130122011202323-1101200223301122-1302232231202330-3311302201102113-2012101121333332-2332110213001201-3032003023002023-3100322222102230"></a>

## custom_hash_algorithms — custom_hash_algorithms / 302020012312 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- custom_hash_algorithms

<a id="canonical-0332010030333003-3020030330020200-0323221203003331-1013103210232131-0222132321131330-1310101323302011-1300122110213332-2223123113131111"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_hash\_algorithms, disable\_ocsp\_stapling, use\_system\_defaults; Default:
use\_system\_defaults\] Specifies the hash algorithms to be used.

Upstream description:

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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

OneOf alternatives in this subsection:

- [custom_hash_algorithms](resources--certificate--reference--group-001.md#canonical-0332010030333003-3020030330020200-0323221203003331-1013103210232131-0222132321131330-1310101323302011-1300122110213332-2223123113131111)
- [disable_ocsp_stapling](resources--certificate--reference--group-001.md#canonical-3112210023202301-0330301120233313-3300023132021120-3202331002133131-1010200312002010-3302112331202331-1120331210032120-1030211100333130)
- [use_system_defaults](resources--certificate--reference--group-001.md#canonical-3101230030230231-1320130120122112-2020012023230002-2003032110301202-1121311100130003-0001321111002231-2003020200001100-2312032130010111)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113303032233311-2323100031233331-3301111330120123-2203301322120321-0013130023302213-0123003023112030-2210000223220133-3101013321011013"></a>

## Direct properties — custom_hash_algorithms / 302020012312 / 3

<a id="canonical-3110300333231130-3310010200131301-3023003021121333-3333230133221303-3131031032220301-3010012221131003-2013120100300213-0321023300310313"></a>

<a id="canonical-2021021021122122-2320132323010033-1231111122302331-0122110130030133-0202120302100030-3123101211332200-3212332311100332-2003310133020031"></a>

## hash_algorithms property — custom_hash_algorithms / 302020012312 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3001112330303200-2000033113313332-1131103102031300-3333221203100331-0121131020102100-1310221012132323-2201332330220202-3110123131211023"></a>

## Next pages — custom_hash_algorithms / 302020012312 / 5

- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)

<a id="canonical-3232212033302331-1222133011022110-3212321322101213-0313001110013132-0330213212211122-0302330121213011-1313201320102231-0233302020223133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203212333312020-2222330021311310-1010132001200003-0131033003102121-0100112223103201-3031302023302302-0023112211311300-2210033303203113"></a>

## disable_ocsp_stapling — disable_ocsp_stapling / 322013200201 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- disable_ocsp_stapling

<a id="canonical-3112210023202301-0330301120233313-3300023132021120-3202331002133131-1010200312002010-3302112331202331-1120331210032120-1030211100333130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-2230023312021330-2030223123231220-1212312100211012-2130220010212313-1103001023200130-1021003103122111-2213132303223030-2202022133110230"></a>

## Direct properties — disable_ocsp_stapling / 322013200201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001212000011132-2012300033302233-0212002122000121-0330112332303320-1121011303102301-1132312113202113-0002222211010100-1330120303111000"></a>

## Next pages — disable_ocsp_stapling / 322013200201 / 4

- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)

<a id="canonical-2032012232223302-3313110001002321-0100312230130202-0202320002311022-3233121322320013-0213333220230021-1223231110210112-2132120202111322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323221230032120-1132011311110302-0131310032220310-2321102000211103-2221333120222211-2120211333031313-1233322323303012-3310233203031023"></a>

## private_key — private_key / 020213210320 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- private_key

<a id="canonical-0100130131110220-0011113301322303-1310022010313231-1132001300000030-2032203200130133-0102112301032301-1130003031303201-2001103211102333"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111212322232331-3332100230210221-0222130233021313-1323012003132203-2110222100122311-2022011220320002-2123000222213122-0123102223300133"></a>

## Direct properties — private_key / 020213210320 / 3

- [blindfold_secret_info](resources--certificate--reference--group-001.md#canonical-1212130133113120-3213313012221220-1110001201213313-2131112331202231-2020230313231212-0330310130022011-1120030320133221-3001003323101233): complete subsection reference.

- [clear_secret_info](resources--certificate--reference--group-001.md#canonical-3002130021002013-0003231011201212-2110203232203110-2331222133010022-1020311000320201-2113111203202220-1313300012200102-2201233122303102): complete subsection reference.

<a id="canonical-3332131122133332-3323320331001222-0101333310201200-2121203330232132-2230210023302032-3020032221311203-0213101303212013-0320221310032023"></a>

## Next pages — private_key / 020213210320 / 4

- [private_key.blindfold_secret_info](resources--certificate--reference--group-001.md#canonical-1212130133113120-3213313012221220-1110001201213313-2131112331202231-2020230313231212-0330310130022011-1120030320133221-3001003323101233)
- [private_key.clear_secret_info](resources--certificate--reference--group-001.md#canonical-3002130021002013-0003231011201212-2110203232203110-2331222133010022-1020311000320201-2113111203202220-1313300012200102-2201233122303102)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)

<a id="canonical-1212130133113120-3213313012221220-1110001201213313-2131112331202231-2020230313231212-0330310130022011-1120030320133221-3001003323101233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210330131331010-1021111132121230-1312013010021002-2233111231132021-3233013012311302-1110001203311010-2213103110012212-3022311232312032"></a>

## private_key.blindfold_secret_info — blindfold_secret_info / 023031011100 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- [private_key](resources--certificate--reference--group-001.md#canonical-2032012232223302-3313110001002321-0100312230130202-0202320002311022-3233121322320013-0213333220230021-1223231110210112-2132120202111322)
- private_key.blindfold_secret_info

<a id="canonical-2231111122010112-0121323323011231-2320123113013123-2320233213210120-2030020301002303-1112232200021020-2221102232222130-2120021011011330"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203233002013110-1210101013321003-1200232122111201-1001133030130122-0013302000201221-2122131023020231-2331223131022100-0110310032310021"></a>

## Direct properties — blindfold_secret_info / 023031011100 / 3

<a id="canonical-2230002020113101-0010223013223001-0023222311313032-1001021123200303-0132202013121202-1031103013232231-2233211123202223-0231331231301021"></a>

<a id="canonical-2120312133210223-2132210312210310-1020330010221213-2320011100102000-2132120021111130-2102003222012103-0310213000331300-3210102233330320"></a>

## decryption_provider property — blindfold_secret_info / 023031011100 / 4

Type: `"string"`. Optional.

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

<a id="canonical-1031022210300020-0331210101210101-0113301301311002-1320102133002130-3230212211213312-3221000011220122-3023312133112102-2313233121203232"></a>

<a id="canonical-0121113101103200-2321023213331120-1200033132300112-1333301213003322-3012210121330211-3113123001122120-3111003111322011-0033022331233101"></a>

## location property — blindfold_secret_info / 023031011100 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-1022312313300303-3330120320200312-1223210113301132-0010202320302300-1232311333132031-0001111002311220-3022333202321222-0020300021110202"></a>

<a id="canonical-2203211223332230-3113302102231032-3130301111311333-1201322332023100-0111232303112120-2122321112033111-1332021302211231-3000230031202112"></a>

## store_provider property — blindfold_secret_info / 023031011100 / 6

Type: `"string"`. Optional.

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

<a id="canonical-2030203300333113-3332333101131222-1202231321200333-3332133313333201-1333013222211331-0321003010211132-1211121322311111-3203022321021102"></a>

## Next pages — blindfold_secret_info / 023031011100 / 7

- [private_key](resources--certificate--reference--group-001.md#canonical-2032012232223302-3313110001002321-0100312230130202-0202320002311022-3233121322320013-0213333220230021-1223231110210112-2132120202111322)
- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)

<a id="canonical-3002130021002013-0003231011201212-2110203232203110-2331222133010022-1020311000320201-2113111203202220-1313300012200102-2201233122303102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333321220201021-0131022001131101-0232203112211033-2111223320123220-0330132031033301-2223121232210023-3031202033001313-3000300012333230"></a>

## private_key.clear_secret_info — clear_secret_info / 333223020222 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- [private_key](resources--certificate--reference--group-001.md#canonical-2032012232223302-3313110001002321-0100312230130202-0202320002311022-3233121322320013-0213333220230021-1223231110210112-2132120202111322)
- private_key.clear_secret_info

<a id="canonical-0232100033021131-0200222301033031-1100332131030221-1133110310203222-1122211202113033-0303220221133302-0200212021332222-0210312222120202"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211023231131213-1102101103303311-3020231130330030-0102023200003120-2103333122123102-3013123012230013-1211012323010332-1300320113111221"></a>

## Direct properties — clear_secret_info / 333223020222 / 3

<a id="canonical-1322302211223031-0333203031303322-1313121200221133-1121100002222013-3333002122121013-3231101302103311-0220323110133302-1311013332020313"></a>

<a id="canonical-1033121003113233-1223203031223303-0113121321212131-3200122213303023-0123102120213331-1101030023131200-3132103312001102-0012210001021201"></a>

## provider_ref property — clear_secret_info / 333223020222 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3032023123012320-0111332113111321-2002131103110320-1213210132200100-0321200110221232-0212000300110103-0233111221211333-0002022300022330"></a>

<a id="canonical-2312223300210212-2120110000120221-1333220233300332-0332023002003032-0103203112213001-0001222113221011-0022110002202120-1102321201303310"></a>

## URL property — clear_secret_info / 333223020222 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-1122221132300312-2130313333303222-0313212230021122-3002103310020230-2223001130312120-0003232003222231-3321111313233221-1112100002321233"></a>

## Next pages — clear_secret_info / 333223020222 / 6

- [private_key](resources--certificate--reference--group-001.md#canonical-2032012232223302-3313110001002321-0100312230130202-0202320002311022-3233121322320013-0213333220230021-1223231110210112-2132120202111322)
- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)

<a id="canonical-0022303111210311-0212221001121021-2111110133110332-1103310031101303-1020011013221332-0130113102021020-1110332113310011-1103211100200222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233000100202132-1102032113000320-2031311002130031-3320320230002003-1211001233121001-0333230121101203-1000230121202110-0101211330001003"></a>

## timeouts — timeouts / 303230122200 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- timeouts

<a id="canonical-0000303133321230-1012002213111110-0031031201201332-2222102332200222-3132331013210231-2301313033030213-0210202323230022-0130230330123212"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310131222131102-0220332022111332-3123220303212123-2130131310011222-0020303000113110-2310330023100203-3132331321200010-2111110130311313"></a>

## Direct properties — timeouts / 303230122200 / 3

<a id="canonical-2103123102303330-0320110110203203-0010002101020110-0122223020320211-0100313010001123-3122202321233220-0031021030311012-0022033101332000"></a>

<a id="canonical-2133012230311210-0203133312033231-3213312231011023-3221213230222132-1000313333223130-2012001002330111-3102321223300310-3130113123111232"></a>

## create property — timeouts / 303230122200 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1002220213101223-3320020313122203-2030320231313200-0303133310200030-0232223001021120-3010020033032301-3023333120220321-0310112011103333"></a>

<a id="canonical-3031111003133313-0330102333012101-3323232223210323-2213301320200302-1210120211102203-1020203222020023-3001302232332313-2212030133012312"></a>

## delete property — timeouts / 303230122200 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1133311000020021-0311113100113302-1232321120310122-2330113330313303-2210311313202012-3232311300021212-3111222301201022-3323030232232123"></a>

<a id="canonical-3123103230313330-3210212131201001-3121302132300200-3320311311222201-3322122032010201-0003211112131312-1201023222003311-0010010122033131"></a>

## read property — timeouts / 303230122200 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1120011233003201-2000212210010112-0033011100022110-3231030033323133-3230321202231000-3233032332223033-2102210201101110-0231222103012330"></a>

<a id="canonical-0033032122031103-0321130300113012-3330210203032033-2120001100222333-0202121321320223-1231320213312000-0120112213320302-1133012322303311"></a>

## update property — timeouts / 303230122200 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1333103032132001-3010103210201021-1133000321103222-2101213331213203-1330121332130122-3310122302022103-3112033301013010-1223013322230312"></a>

## Next pages — timeouts / 303230122200 / 8

- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)

<a id="canonical-1011132030030013-2320300233313232-2321133121131212-2123120301103000-1003011033212130-0213210031100131-0330120313310313-1100200323301322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212213010010010-1323300300222220-3123013003223202-0221102103311223-0000023221032133-0311313000132312-2002212012231233-1123133100213200"></a>

## use_system_defaults — use_system_defaults / 111020330231 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- use_system_defaults

<a id="canonical-3101230030230231-1320130120122112-2020012023230002-2003032110301202-1121311100130003-0001321111002231-2003020200001100-2312032130010111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-3331223312002030-3033313203322321-0322300112132211-3320212221101003-3002012103303332-3222331332131302-1201023010130301-0332012001213123"></a>

## Direct properties — use_system_defaults / 111020330231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322300020330010-1332120323033333-0231302310321332-1032221233223002-2332032213331123-2112100220130022-2021301302200102-1303102002103220"></a>

## Next pages — use_system_defaults / 111020330231 / 4

- [Property reference](resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
