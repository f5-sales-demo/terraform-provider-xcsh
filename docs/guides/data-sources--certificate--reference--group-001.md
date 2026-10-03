---
page_title: "xcsh_certificate reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate reference."
---

# xcsh_certificate reference

<a id="canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002023323133223-0300011332221313-3321030100200023-0113103110301011-1333222022113222-2213111220030211-1111322001313113-2213101221111322"></a>

## Property reference — Property reference / 220132113323 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)
- Property reference

<a id="canonical-0031233202303322-2310030130121201-2102301301221321-0300101011300013-1200311111320301-3113203222301021-0002110221232132-3001020222012010"></a>

## Direct properties — Property reference / 220132113323 / 3

<a id="canonical-1031033320102310-0003103121010231-0201131230103313-1202030123022110-3222111223031312-3231323023131102-0331323110220321-2322100301313302"></a>

<a id="canonical-0031223103222220-3223233103202133-2200132201031103-3220102130033010-0101032002230311-2232132202213200-0201203300000231-1232102232130003"></a>

## annotations property — Property reference / 220132113323 / 4

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

- [certificate_chain](data-sources--certificate--reference--group-001.md#canonical-1201120301312222-1110032002030323-1311101123101112-2231320131101203-3033033031011131-0002122100230310-1313320132333010-2022233002033223): complete subsection reference.

<a id="canonical-1023002211020301-1013302112003011-2203311110210212-0103312201203010-2100032112113323-2002202013202120-2012212321122103-1210120000212001"></a>

<a id="canonical-2030101001311122-1020212220213010-3111013332233030-1133331200213113-2120132032322320-1213131331230121-1320021322010000-3210101110012221"></a>

## certificate_url property — Property reference / 220132113323 / 5

Type: `"string"`. Computed.

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](data-sources--certificate--reference--group-001.md#canonical-0303233230122130-1303313303031100-0320311101201112-1233011200123032-3233130001310213-0200202130000212-3132121010121103-3203232021031312): complete subsection reference.

<a id="canonical-1230123023323231-2332321302302000-2130033030201310-3320202131012023-0100103301011032-1010000200301230-3213013202010301-2020101231331132"></a>

<a id="canonical-3030232232003332-0000032121312212-2201213203201232-2203130222311321-3002132133233210-1201013330221110-3230102002111111-1302202231302201"></a>

## description property — Property reference / 220132113323 / 6

Type: `"string"`. Computed.

Description of the Certificate.

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

- [disable_ocsp_stapling](data-sources--certificate--reference--group-001.md#canonical-1303300001332333-2133003111102030-1211002312221203-1032210221332122-0332301200132133-2220203300032233-2102213320133233-2200000202032210): complete subsection reference.

<a id="canonical-2311220100131030-0131312132322210-2223301131020031-2303313210201223-0031322301222203-0123221132003323-3331203313331001-1030300131302220"></a>

<a id="canonical-2330110032121130-3101300011323231-0111013000211223-1022112103002011-0113312100100311-3110130103200201-1130333022123330-0303020022013210"></a>

## ID property — Property reference / 220132113323 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1322321031021320-2220301303032310-1030011213323001-0301130011130332-3000003210033011-3201021032300111-2001113012123121-0311301002201222"></a>

<a id="canonical-1231232223202202-2313023131210131-2132311123212010-2333311231230323-2233221200023303-2020323000130202-2110131110002002-2220320322323000"></a>

## labels property — Property reference / 220132113323 / 8

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

<a id="canonical-1013030003102222-0100203321003102-1120122020213320-0211331200011011-0120222022300231-2300312300330233-0100312121213123-0201300220033112"></a>

<a id="canonical-3130023013231002-2132220301233313-2000300302121110-3110033100321303-1313312110313312-3223001013202222-3313120103333030-3201310211132303"></a>

## name property — Property reference / 220132113323 / 9

Type: `"string"`. Required.

Name of the Certificate.

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

<a id="canonical-0330312230210003-0013203020300123-3332211311030012-1323123001133000-2223001112233323-2112333301110112-1302322311011133-2221233312223221"></a>

<a id="canonical-2033201113020102-3121202011223311-0201130112010110-1112000101221300-2121022102010320-0100202122100211-2203310321333100-2232103302123013"></a>

## namespace property — Property reference / 220132113323 / 10

Type: `"string"`. Required.

Namespace where the Certificate exists.

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

- [private_key](data-sources--certificate--reference--group-001.md#canonical-0303332301111032-1121312202010230-2010123001002110-1000200100332031-2022332113132133-1001003300221032-0013103033011303-1322002121022121): complete subsection reference.

- [use_system_defaults](data-sources--certificate--reference--group-001.md#canonical-0322002230033322-0010111113303313-1123023032313223-1232100302003111-3313122222131302-2120130313311330-1112120102030010-1300311202323230): complete subsection reference.

<a id="canonical-0100113222321033-0220000232133001-1300302332300221-1012023323211102-0012112222101231-0012002323320102-0321032232001132-1103232213301112"></a>

## All schema paths — Property reference / 220132113323 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--certificate--reference--group-001.md#canonical-1031033320102310-0003103121010231-0201131230103313-1202030123022110-3222111223031312-3231323023131102-0331323110220321-2322100301313302) |
| `certificate_chain` | [certificate_chain](data-sources--certificate--reference--group-001.md#canonical-3003212132011222-1300022320331311-2130322212233201-3132130010030333-2332113131212220-2010303310110033-2133133232031232-0013210233301132) |
| `certificate_chain.name` | [certificate_chain.name](data-sources--certificate--reference--group-001.md#canonical-2033313122310301-3313033122111113-1003002120200123-2332112121021021-2122301100033101-3030333301212020-2123311212230223-3230023222200003) |
| `certificate_chain.namespace` | [certificate_chain.namespace](data-sources--certificate--reference--group-001.md#canonical-2000023301211320-3303211022303111-1013331013313033-3002101023210011-1103220212032011-2012223001112002-3012332121332203-2203213113321102) |
| `certificate_chain.tenant` | [certificate_chain.tenant](data-sources--certificate--reference--group-001.md#canonical-0123223020020300-2023011000032222-2003000033113000-0110111101030321-3120123100002332-1001303202113220-1110301011223233-3111101130311203) |
| `certificate_url` | [certificate_url](data-sources--certificate--reference--group-001.md#canonical-1023002211020301-1013302112003011-2203311110210212-0103312201203010-2100032112113323-2002202013202120-2012212321122103-1210120000212001) |
| `custom_hash_algorithms` | [custom_hash_algorithms](data-sources--certificate--reference--group-001.md#canonical-3021320020200112-3312313110233023-3010032010202012-1100313232121230-1111011333323231-0302121303310330-2331220112302210-2333001132301233) |
| `custom_hash_algorithms.hash_algorithms` | [custom_hash_algorithms.hash_algorithms](data-sources--certificate--reference--group-001.md#canonical-2030223313212321-2311221131033223-1110122212103302-3011322203103010-2103311020012112-2221232102230020-3303001030200120-0223011101001233) |
| `description` | [description](data-sources--certificate--reference--group-001.md#canonical-1230123023323231-2332321302302000-2130033030201310-3320202131012023-0100103301011032-1010000200301230-3213013202010301-2020101231331132) |
| `disable_ocsp_stapling` | [disable_ocsp_stapling](data-sources--certificate--reference--group-001.md#canonical-3320302211031300-1232210333013122-3221332221330223-0212331200133323-0130333332310303-0003233222120103-2110133300220323-0023112100313011) |
| `id` | [ID](data-sources--certificate--reference--group-001.md#canonical-2311220100131030-0131312132322210-2223301131020031-2303313210201223-0031322301222203-0123221132003323-3331203313331001-1030300131302220) |
| `labels` | [labels](data-sources--certificate--reference--group-001.md#canonical-1322321031021320-2220301303032310-1030011213323001-0301130011130332-3000003210033011-3201021032300111-2001113012123121-0311301002201222) |
| `name` | [name](data-sources--certificate--reference--group-001.md#canonical-1013030003102222-0100203321003102-1120122020213320-0211331200011011-0120222022300231-2300312300330233-0100312121213123-0201300220033112) |
| `namespace` | [namespace](data-sources--certificate--reference--group-001.md#canonical-0330312230210003-0013203020300123-3332211311030012-1323123001133000-2223001112233323-2112333301110112-1302322311011133-2221233312223221) |
| `private_key` | [private_key](data-sources--certificate--reference--group-001.md#canonical-2312230320311013-3120021212122121-0133020121103221-0032021322322110-2002231323210020-0123310202130022-3010313210301003-1301203033123200) |
| `private_key.blindfold_secret_info` | [private_key.blindfold_secret_info](data-sources--certificate--reference--group-001.md#canonical-1012030302111020-0130330212100323-0312022303230033-0101213112122100-0203111320220110-0212011021012321-3312301330010203-1023232013312303) |
| `private_key.blindfold_secret_info.decryption_provider` | [private_key.blindfold_secret_info.decryption_provider](data-sources--certificate--reference--group-001.md#canonical-3210220112223310-2003120211021030-0032232333001013-1022113030110102-2033103013133200-2101212233013102-2001123232102022-2131301011112132) |
| `private_key.blindfold_secret_info.location` | [private_key.blindfold_secret_info.location](data-sources--certificate--reference--group-001.md#canonical-1020232200123212-0032310133001211-2202111020120320-0021011330101332-0201212122030011-0221123330211022-2203333031110032-1120332022003132) |
| `private_key.blindfold_secret_info.store_provider` | [private_key.blindfold_secret_info.store_provider](data-sources--certificate--reference--group-001.md#canonical-0201003212033120-2200223003212103-0320000302210032-2020332010000023-1232232320133011-1201101101323112-0200231330320131-0201311311101030) |
| `private_key.clear_secret_info` | [private_key.clear_secret_info](data-sources--certificate--reference--group-001.md#canonical-0232200202212232-2101331013023302-2133002211020020-3012221320203330-0211322311133232-2033011113200113-0132210332131102-2103313312032003) |
| `private_key.clear_secret_info.provider_ref` | [private_key.clear_secret_info.provider_ref](data-sources--certificate--reference--group-001.md#canonical-3002221210033232-3211002022310033-2120211203201113-1000121123123322-1202020300202313-0023121100213003-3032211201212202-2201233220032021) |
| `private_key.clear_secret_info.url` | [private_key.clear_secret_info.url](data-sources--certificate--reference--group-001.md#canonical-1300321132011223-2310110022233112-3322023213332212-1230331232113022-3133233120020312-2032110233221010-3103013013301201-3101011212330303) |
| `use_system_defaults` | [use_system_defaults](data-sources--certificate--reference--group-001.md#canonical-3013213321022321-0010331333133300-1231031313201232-1022213322320213-0131003122032330-1302123112100023-1121030212323011-3203013300302232) |

<a id="canonical-3113121113232112-0203210201012031-3030003230232331-0031101101202220-3022100321210212-0301002202101302-1303120010032010-2232103013311101"></a>

## Next pages — Property reference / 220132113323 / 12

- [certificate_chain](data-sources--certificate--reference--group-001.md#canonical-1201120301312222-1110032002030323-1311101123101112-2231320131101203-3033033031011131-0002122100230310-1313320132333010-2022233002033223)
- [custom_hash_algorithms](data-sources--certificate--reference--group-001.md#canonical-0303233230122130-1303313303031100-0320311101201112-1233011200123032-3233130001310213-0200202130000212-3132121010121103-3203232021031312)
- [disable_ocsp_stapling](data-sources--certificate--reference--group-001.md#canonical-1303300001332333-2133003111102030-1211002312221203-1032210221332122-0332301200132133-2220203300032233-2102213320133233-2200000202032210)
- [private_key](data-sources--certificate--reference--group-001.md#canonical-0303332301111032-1121312202010230-2010123001002110-1000200100332031-2022332113132133-1001003300221032-0013103033011303-1322002121022121)
- [use_system_defaults](data-sources--certificate--reference--group-001.md#canonical-0322002230033322-0010111113303313-1123023032313223-1232100302003111-3313122222131302-2120130313311330-1112120102030010-1300311202323230)
- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)

<a id="canonical-1201120301312222-1110032002030323-1311101123101112-2231320131101203-3033033031011131-0002122100230310-1313320132333010-2022233002033223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322322301030201-0112221131122313-0320311023330331-0030131211311030-2122320310201001-2110131322010030-2013132200301332-1201132112100101"></a>

## certificate_chain — certificate_chain / 020011330212 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- certificate_chain

<a id="canonical-3003212132011222-1300022320331311-2130322212233201-3132130010030333-2332113131212220-2010303310110033-2133133232031232-0013210233301132"></a>

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

<a id="canonical-2331312312202013-0233212002033213-2202002010000313-0333102202022132-3132331323111223-1001110222323121-3212213201202231-1010103113102110"></a>

## Direct properties — certificate_chain / 020011330212 / 3

<a id="canonical-2033313122310301-3313033122111113-1003002120200123-2332112121021021-2122301100033101-3030333301212020-2123311212230223-3230023222200003"></a>

<a id="canonical-0322300030301232-3310122030121030-1301300133010313-2200230200113013-2022013001013131-1013312222003110-2002231011320112-1000132312132031"></a>

## name property — certificate_chain / 020011330212 / 4

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

<a id="canonical-2000023301211320-3303211022303111-1013331013313033-3002101023210011-1103220212032011-2012223001112002-3012332121332203-2203213113321102"></a>

<a id="canonical-1031030202031030-1110102022100321-0200201120123201-0200321112203103-1110011220003213-1222101203120301-0032303132313311-3331200320121202"></a>

## namespace property — certificate_chain / 020011330212 / 5

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

<a id="canonical-0123223020020300-2023011000032222-2003000033113000-0110111101030321-3120123100002332-1001303202113220-1110301011223233-3111101130311203"></a>

<a id="canonical-0023123312332331-2101023013022231-1133323322113101-3220230223022302-1101023232022332-1121213113321112-3023203100313220-1022231210012200"></a>

## tenant property — certificate_chain / 020011330212 / 6

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

<a id="canonical-1132132031032123-0312123211223113-2013113303100113-2020122113002012-2213112320232310-3232110111320323-2223111011002232-3121303111121121"></a>

## Next pages — certificate_chain / 020011330212 / 7

- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)

<a id="canonical-0303233230122130-1303313303031100-0320311101201112-1233011200123032-3233130001310213-0200202130000212-3132121010121103-3203232021031312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233310330332032-3130121300232222-3003330021121332-2320112003322103-2231312301322033-2101020231000012-2022203033311032-3012310113211011"></a>

## custom_hash_algorithms — custom_hash_algorithms / 320033103231 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- custom_hash_algorithms

<a id="canonical-3021320020200112-3312313110233023-3010032010202012-1100313232121230-1111011333323231-0302121303310330-2331220112302210-2333001132301233"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_hash\_algorithms, disable\_ocsp\_stapling, use\_system\_defaults; Default:
use\_system\_defaults\] Specifies the hash algorithms to be used.

Upstream description:

Specifies the hash algorithms to be used.

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

- [custom_hash_algorithms](data-sources--certificate--reference--group-001.md#canonical-3021320020200112-3312313110233023-3010032010202012-1100313232121230-1111011333323231-0302121303310330-2331220112302210-2333001132301233)
- [disable_ocsp_stapling](data-sources--certificate--reference--group-001.md#canonical-3320302211031300-1232210333013122-3221332221330223-0212331200133323-0130333332310303-0003233222120103-2110133300220323-0023112100313011)
- [use_system_defaults](data-sources--certificate--reference--group-001.md#canonical-3013213321022321-0010331333133300-1231031313201232-1022213322320213-0131003122032330-1302123112100023-1121030212323011-3203013300302232)

Select alternatives according to the provider validators above.

<a id="canonical-2210213302010100-0202201033021002-2211211313100332-3223202322301311-3120020132011331-2202213122002200-3021321003233113-0020233212011123"></a>

## Direct properties — custom_hash_algorithms / 320033103231 / 3

<a id="canonical-2030223313212321-2311221131033223-1110122212103302-3011322203103010-2103311020012112-2221232102230020-3303001030200120-0223011101001233"></a>

<a id="canonical-1320131222231132-0133110211020202-3320321120323032-1232101130000120-0112131011322033-1133020011013032-3022022111331110-0130132011021012"></a>

## hash_algorithms property — custom_hash_algorithms / 320033103231 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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

<a id="canonical-0031213321322213-0332200131122130-1000122010001202-0323223111210213-2023120320212302-1131323131121303-2231131011302002-0020133113030211"></a>

## Next pages — custom_hash_algorithms / 320033103231 / 5

- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)

<a id="canonical-1303300001332333-2133003111102030-1211002312221203-1032210221332122-0332301200132133-2220203300032233-2102213320133233-2200000202032210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211231033220333-3112311221121002-2121300131022111-3220211113010122-0010031003022202-0232212031332121-3012223001323130-1301310122130212"></a>

## disable_ocsp_stapling — disable_ocsp_stapling / 320002311320 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- disable_ocsp_stapling

<a id="canonical-3320302211031300-1232210333013122-3221332221330223-0212331200133323-0130333332310303-0003233222120103-2110133300220323-0023112100313011"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3020012220133332-2333022301332223-0012321102011313-3221113212312230-1000212122032131-2202223130203033-1130002210301112-1210331302130301"></a>

## Direct properties — disable_ocsp_stapling / 320002311320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013111213122102-0111012331123221-2202001333201012-2020000012202123-3303002323020002-1102310132033100-2223010220312233-2022233103123231"></a>

## Next pages — disable_ocsp_stapling / 320002311320 / 4

- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)

<a id="canonical-0303332301111032-1121312202010230-2010123001002110-1000200100332031-2022332113132133-1001003300221032-0013103033011303-1322002121022121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311120233110133-2222030203032212-0321301330200320-1130131313111013-1203303003231222-0331023303021330-2300132330322020-2332100223030121"></a>

## private_key — private_key / 213110302200 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- private_key

<a id="canonical-2312230320311013-3120021212122121-0133020121103221-0032021322322110-2002231323210020-0123310202130022-3010313210301003-1301203033123200"></a>

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

<a id="canonical-1110102230010233-1033221012310002-3101012213230302-1021312011212010-1232220022103120-0333333131221012-0102120113033000-1333113200002301"></a>

## Direct properties — private_key / 213110302200 / 3

- [blindfold_secret_info](data-sources--certificate--reference--group-001.md#canonical-3032322222111030-1210310322211001-1302023313202232-1030121331103231-2221020002031332-1233220131331130-2131130231331113-2303112033021323): complete subsection reference.

- [clear_secret_info](data-sources--certificate--reference--group-001.md#canonical-0312212232033310-2233230322130022-1332030032003133-3132303102333130-1123121202123101-2130321121222323-3011222232331310-2213202123300230): complete subsection reference.

<a id="canonical-3321012003102031-3101122023223002-2323032233312001-0323023312131003-0101013111033022-3002321011220302-0133112111130002-1010003130000113"></a>

## Next pages — private_key / 213110302200 / 4

- [private_key.blindfold_secret_info](data-sources--certificate--reference--group-001.md#canonical-3032322222111030-1210310322211001-1302023313202232-1030121331103231-2221020002031332-1233220131331130-2131130231331113-2303112033021323)
- [private_key.clear_secret_info](data-sources--certificate--reference--group-001.md#canonical-0312212232033310-2233230322130022-1332030032003133-3132303102333130-1123121202123101-2130321121222323-3011222232331310-2213202123300230)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)

<a id="canonical-3032322222111030-1210310322211001-1302023313202232-1030121331103231-2221020002031332-1233220131331130-2131130231331113-2303112033021323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302233212033123-3111003130112013-0322123233122322-2102120003211213-3102311233200233-1210321000210330-3201210120332221-0102223202123131"></a>

## private_key.blindfold_secret_info — blindfold_secret_info / 210111131201 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- [private_key](data-sources--certificate--reference--group-001.md#canonical-0303332301111032-1121312202010230-2010123001002110-1000200100332031-2022332113132133-1001003300221032-0013103033011303-1322002121022121)
- private_key.blindfold_secret_info

<a id="canonical-1012030302111020-0130330212100323-0312022303230033-0101213112122100-0203111320220110-0212011021012321-3312301330010203-1023232013312303"></a>

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

<a id="canonical-0130102103133301-1032012300323012-3001030222131221-2021222033323100-3231033311201210-3111030230312111-2133132302311122-0003030021332123"></a>

## Direct properties — blindfold_secret_info / 210111131201 / 3

<a id="canonical-3210220112223310-2003120211021030-0032232333001013-1022113030110102-2033103013133200-2101212233013102-2001123232102022-2131301011112132"></a>

<a id="canonical-2031213201021032-2103223201112230-0023111320020003-3230011132333300-2313322320030232-0002301303103010-3001030300012102-1030123210313331"></a>

## decryption_provider property — blindfold_secret_info / 210111131201 / 4

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

<a id="canonical-1020232200123212-0032310133001211-2202111020120320-0021011330101332-0201212122030011-0221123330211022-2203333031110032-1120332022003132"></a>

<a id="canonical-0300222122212220-0200111110302003-1032022123032331-2102102013021232-1202123223003333-0102012231112011-3203311201133000-2001001320332011"></a>

## location property — blindfold_secret_info / 210111131201 / 5

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

<a id="canonical-0201003212033120-2200223003212103-0320000302210032-2020332010000023-1232232320133011-1201101101323112-0200231330320131-0201311311101030"></a>

<a id="canonical-3011202303013010-0312111003333320-1302322322011203-2013001101300032-3112322122020111-0230132023322213-2312312103112210-1231333303301031"></a>

## store_provider property — blindfold_secret_info / 210111131201 / 6

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

<a id="canonical-1200012102112121-1121010300112212-0200112300301230-3302312213122121-2022301331223132-1113020111222331-2003313010201130-2220311101032303"></a>

## Next pages — blindfold_secret_info / 210111131201 / 7

- [private_key](data-sources--certificate--reference--group-001.md#canonical-0303332301111032-1121312202010230-2010123001002110-1000200100332031-2022332113132133-1001003300221032-0013103033011303-1322002121022121)
- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)

<a id="canonical-0312212232033310-2233230322130022-1332030032003133-3132303102333130-1123121202123101-2130321121222323-3011222232331310-2213202123300230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120221233203012-1211333021321322-2213321121102012-3031013210320132-1100010111022230-2322132023031303-2233032020013011-2011000301013112"></a>

## private_key.clear_secret_info — clear_secret_info / 330020211232 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- [private_key](data-sources--certificate--reference--group-001.md#canonical-0303332301111032-1121312202010230-2010123001002110-1000200100332031-2022332113132133-1001003300221032-0013103033011303-1322002121022121)
- private_key.clear_secret_info

<a id="canonical-0232200202212232-2101331013023302-2133002211020020-3012221320203330-0211322311133232-2033011113200113-0132210332131102-2103313312032003"></a>

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

<a id="canonical-1231032123321003-2313312112210321-0113113112210130-0100310123213301-2330220321113321-0202130301013003-3303323023021211-1131223132101232"></a>

## Direct properties — clear_secret_info / 330020211232 / 3

<a id="canonical-3002221210033232-3211002022310033-2120211203201113-1000121123123322-1202020300202313-0023121100213003-3032211201212202-2201233220032021"></a>

<a id="canonical-3102112100100231-0003112013233201-2221110303023133-2231132123020312-2010330030012103-1132012333333320-0131112321133112-2010210221103122"></a>

## provider_ref property — clear_secret_info / 330020211232 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1300321132011223-2310110022233112-3322023213332212-1230331232113022-3133233120020312-2032110233221010-3103013013301201-3101011212330303"></a>

<a id="canonical-2211002223330020-2301032220101230-2131330033210321-0011212323323322-0102020022232320-2231313122101020-0320011103230213-3021112133010112"></a>

## URL property — clear_secret_info / 330020211232 / 5

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

<a id="canonical-3022031103223301-3113121121011200-0032201003333000-3303230100233010-1033312110223220-0333023003223132-0221110033333120-2130321310231322"></a>

## Next pages — clear_secret_info / 330020211232 / 6

- [private_key](data-sources--certificate--reference--group-001.md#canonical-0303332301111032-1121312202010230-2010123001002110-1000200100332031-2022332113132133-1001003300221032-0013103033011303-1322002121022121)
- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)

<a id="canonical-0322002230033322-0010111113303313-1123023032313223-1232100302003111-3313122222131302-2120130313311330-1112120102030010-1300311202323230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012031211303100-0233120210211002-0210122012333302-0000223132322203-2103321102110112-2011122220230322-2032021112011321-0100321312311320"></a>

## use_system_defaults — use_system_defaults / 332130232020 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- use_system_defaults

<a id="canonical-3013213321022321-0010331333133300-1231031313201232-1022213322320213-0131003122032330-1302123112100023-1121030212323011-3203013300302232"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3131122313003333-1133222200211102-1001232020330120-0130210031211232-2103002133310213-1021030101312202-1202332311221200-2102333212112020"></a>

## Direct properties — use_system_defaults / 332130232020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311000012103123-2312200111310000-1332303300110221-3330213000023231-3000202332323231-1132121031132330-2213230330002121-3322022102322021"></a>

## Next pages — use_system_defaults / 332130232020 / 4

- [Property reference](data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)
