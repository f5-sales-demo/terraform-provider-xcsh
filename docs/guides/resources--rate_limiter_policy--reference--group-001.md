---
page_title: "xcsh_rate_limiter_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy reference."
---

# xcsh_rate_limiter_policy reference

<a id="canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- Property reference

<a id="canonical-3333233001203321-0300313203223213-0022032331313103-3201020010211312-1023222230002333-0213111203030311-3311330330303000-0301310221201013"></a>

### Direct properties for `xcsh_rate_limiter_policy`

<a id="canonical-2313021003313013-0033232223303031-2330231330011200-1012322303102133-3123210220212203-2013030332233313-3113222212011101-0202231210230122"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

- [any_server](resources--rate_limiter_policy--reference--group-001.md#canonical-0211213222320023-0301120212121003-2221311230220112-0031320320033020-2032022001201303-0013020220032120-0021032212231130-1003333312100321): complete subsection reference.

<a id="canonical-1031012313201313-2332012230020321-3302030112112312-2120131112030133-1201022132123123-1312223133203022-3330211321200300-3302030030102122"></a>

<a id="canonical-3221332000212123-2121210201120000-0330110310311312-1011003300310121-0030023033032031-3222123032021111-3200013130003232-0110333013333232"></a>

#### `description` property

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

<a id="canonical-0322021130021133-0011323031002012-2312013202230321-3202120320132120-1031331010310201-2101022030330311-3230223012223231-3032203333322320"></a>

<a id="canonical-3212112011113201-3333330131122122-3100213310100003-2301200331201312-3013011010111321-0002300003032223-1333322231323303-3331301113213312"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

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

<a id="canonical-1102202220012203-1033331022233212-3302230331021323-0030021010020023-2132320002101301-0221122321130211-2302133233302211-2123101023133300"></a>

<a id="canonical-2302123322023121-1012110333031332-1122011210000022-2331323113031102-0311202030320110-0210313001212222-3322301312111323-2322012332313301"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3100203001120123-1002031001031032-0301132103323001-0012222323201120-0120131020200211-3102302020121000-1320002320102120-1033002012103233"></a>

<a id="canonical-0102233130310111-3211200300331221-0110122331331031-2300012203222021-1131030220122022-0131000233223002-0003201003221122-0202220203000211"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-1132320101331113-1210300110233210-2331011103003313-1022113300332330-0311031132200101-1111303320121030-1130210210101331-3112301220221212"></a>

<a id="canonical-3130132023312121-1212302023130331-1321313110310220-1101102231133133-2331122213220123-3330022020022212-3033113101110312-1203331002203200"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Rate Limiter Policy. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1310102202131323-3303001213301120-3301310321033000-1001203020331101-0221002122331303-2112311133020112-0121301022100303-2223333021010033"></a>

<a id="canonical-0320020011102132-1100222310030012-0223101322310002-3012003201311323-2313302102111303-0331220213210202-3031312022232200-1321201132111010"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Rate Limiter Policy is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303): complete subsection reference.

<a id="canonical-0100113011002011-0301110020000203-1332313311221231-3321302233233312-2111213221301013-1223133020120101-3032320003233020-1333310310333012"></a>

<a id="canonical-0231312112112310-3230021303003303-0110231112131113-3131230202310201-1111202320113313-2312011013112331-1013112211012003-0201221002301032"></a>

#### `server_name` property

Type: `"string"`. Optional, Computed.

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server. The actual names for the server are extracted from the HTTP Host header and the name of the
virtual\_host for the request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [server_name_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-0332131121330022-2320313020032211-0233232101323322-0302123232122021-0110301202203032-2012210121330301-1232012122013010-0312200313032203): complete subsection reference.

- [server_selector](resources--rate_limiter_policy--reference--group-001.md#canonical-2213123230331302-2112033033112303-3303010013132210-2021022013123003-2001322313033232-2333120000201120-2312313120200303-1032200022123013): complete subsection reference.

- [timeouts](resources--rate_limiter_policy--reference--group-001.md#canonical-1121122303032230-0332030130132021-3012333310113223-0311011022311021-2203133033133302-0012001113101220-2001003203233201-1012131010323033): complete subsection reference.

<a id="canonical-1013031103030312-2110230201333033-1231221122332230-2011133223132322-1203213321103100-3132320311013001-3302122200313033-1100202322103003"></a>

### All schema paths for `xcsh_rate_limiter_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--rate_limiter_policy--reference--group-001.md#canonical-2313021003313013-0033232223303031-2330231330011200-1012322303102133-3123210220212203-2013030332233313-3113222212011101-0202231210230122) |
| `any_server` | [any_server](resources--rate_limiter_policy--reference--group-001.md#canonical-0213131333013020-3031100203003021-3112303110103113-2121121100002232-3022313102101310-3300001233312130-0103220220320030-1102332123211013) |
| `description` | [description](resources--rate_limiter_policy--reference--group-001.md#canonical-1031012313201313-2332012230020321-3302030112112312-2120131112030133-1201022132123123-1312223133203022-3330211321200300-3302030030102122) |
| `disable` | [disable](resources--rate_limiter_policy--reference--group-001.md#canonical-0322021130021133-0011323031002012-2312013202230321-3202120320132120-1031331010310201-2101022030330311-3230223012223231-3032203333322320) |
| `id` | [ID](resources--rate_limiter_policy--reference--group-001.md#canonical-1102202220012203-1033331022233212-3302230331021323-0030021010020023-2132320002101301-0221122321130211-2302133233302211-2123101023133300) |
| `labels` | [labels](resources--rate_limiter_policy--reference--group-001.md#canonical-3100203001120123-1002031001031032-0301132103323001-0012222323201120-0120131020200211-3102302020121000-1320002320102120-1033002012103233) |
| `name` | [name](resources--rate_limiter_policy--reference--group-001.md#canonical-1132320101331113-1210300110233210-2331011103003313-1022113300332330-0311031132200101-1111303320121030-1130210210101331-3112301220221212) |
| `namespace` | [namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-1310102202131323-3303001213301120-3301310321033000-1001203020331101-0221002122331303-2112311133020112-0121301022100303-2223333021010033) |
| `rules` | [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2212121222113012-2012103123200003-1233130300001122-0232000101330331-3213232131130123-3233212310223332-1112100333101032-2113013200123033) |
| `rules.metadata` | [rules.metadata](resources--rate_limiter_policy--reference--group-001.md#canonical-3032122222030011-2023203100132202-1100212021200001-3201033002310223-1021112200023112-0322113112010300-1332332210002210-1210302203210230) |
| `rules.metadata.description_spec` | [rules.metadata.description_spec](resources--rate_limiter_policy--reference--group-001.md#canonical-1111121002202121-2302331331100131-0132111132311210-1211212312110302-0113012300010233-3132223131323031-2020321033133310-0003302022013012) |
| `rules.metadata.name` | [rules.metadata.name](resources--rate_limiter_policy--reference--group-001.md#canonical-2221331011322101-1231130201030122-2031103223101220-2122313301013010-1223211130330213-0210312012201311-1200101331030122-3113333003030023) |
| `rules.spec` | [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-2133012232331210-3031131131232011-0210113022223120-2130210032332203-1230033222223133-0132010132202232-2102121113020123-0003232231132201) |
| `rules.spec.any_asn` | [rules.spec.any_asn](resources--rate_limiter_policy--reference--group-001.md#canonical-1013030322310323-1131301033112233-3202012233311011-3111220320111100-0012303123113030-3331130133213022-0330221112130002-1312032000222123) |
| `rules.spec.any_country` | [rules.spec.any_country](resources--rate_limiter_policy--reference--group-001.md#canonical-3013213001112330-3102122232330110-2111323032331013-2010002202133310-3102101110301010-1200021113231033-1101221233010132-0011301333312212) |
| `rules.spec.any_ip` | [rules.spec.any_ip](resources--rate_limiter_policy--reference--group-001.md#canonical-3002312003113100-3131303022320201-3332022030031322-3132001000011333-2032312132323010-3120213213033121-3212030120021132-3230011311230122) |
| `rules.spec.apply_rate_limiter` | [rules.spec.apply_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-3132220321033102-0020311131221310-0222133023033312-2031122302130001-2003213122133201-2123302120010111-3230122030130021-3022021012310302) |
| `rules.spec.asn_list` | [rules.spec.asn_list](resources--rate_limiter_policy--reference--group-001.md#canonical-3200121301010133-1311322013330202-3323030323323121-3122202012223111-0133330311012323-2222331332110213-0230021130222100-2020033331322102) |
| `rules.spec.asn_list.as_numbers` | [rules.spec.asn_list.as_numbers](resources--rate_limiter_policy--reference--group-001.md#canonical-2121203321112222-1233333023230300-2231021123022010-3301301131330013-2133020233211210-2132020300322203-3210033122320012-1023222220312223) |
| `rules.spec.asn_matcher` | [rules.spec.asn_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-1020113210302223-3303233033220111-1032231321303132-2012203110032021-0303021202233212-0203203112020300-0021130233123022-1023113112000302) |
| `rules.spec.asn_matcher.asn_sets` | [rules.spec.asn_matcher.asn_sets](resources--rate_limiter_policy--reference--group-001.md#canonical-0013021332331223-2001221023310231-3313203213330201-1000331231311132-3021213001313301-0000031023200323-2230222202232223-2120310001221331) |
| `rules.spec.asn_matcher.asn_sets.kind` | [rules.spec.asn_matcher.asn_sets.kind](resources--rate_limiter_policy--reference--group-001.md#canonical-3013122300032012-1031102232213020-0122020120021112-3023022223332203-2013221310232221-3212300131212123-2332000211303022-1212312112031021) |
| `rules.spec.asn_matcher.asn_sets.name` | [rules.spec.asn_matcher.asn_sets.name](resources--rate_limiter_policy--reference--group-001.md#canonical-2013303022312202-0311221330303120-2232212323001330-1333023211101123-0323000300000212-3231313302103212-0301100331333132-1200120203020221) |
| `rules.spec.asn_matcher.asn_sets.namespace` | [rules.spec.asn_matcher.asn_sets.namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-2121222113301033-0022100110120302-0012131123313031-0303222200001213-1130312002112031-2300022022233110-2002112201333103-3222311303011222) |
| `rules.spec.asn_matcher.asn_sets.tenant` | [rules.spec.asn_matcher.asn_sets.tenant](resources--rate_limiter_policy--reference--group-001.md#canonical-0312211112133101-3301120131203231-3123012200001221-0201212011013332-2331310001322311-0012313013021222-1101200223030131-0121301023221321) |
| `rules.spec.asn_matcher.asn_sets.uid` | [rules.spec.asn_matcher.asn_sets.uid](resources--rate_limiter_policy--reference--group-001.md#canonical-0203001320233221-0322022121222331-2331333231010310-3101233013333101-3132323121003211-3102001031332233-3311223103220320-0302230010232120) |
| `rules.spec.bypass_rate_limiter` | [rules.spec.bypass_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-2333232022012033-3203002233023322-3030222322333333-3222212000020312-3111333002122030-0221211012202112-1221323221102022-1013222300230220) |
| `rules.spec.country_list` | [rules.spec.country_list](resources--rate_limiter_policy--reference--group-001.md#canonical-0301133001110111-1232331231311332-2330120000332200-1032112012010330-1133230101130301-0032322332332210-3301022200332101-0313332232031201) |
| `rules.spec.country_list.country_codes` | [rules.spec.country_list.country_codes](resources--rate_limiter_policy--reference--group-001.md#canonical-1112202200110230-2221113113130013-1121200203130230-0121321312103213-3300121333103323-1011311110020300-0003300311102321-0000213133011200) |
| `rules.spec.country_list.invert_match` | [rules.spec.country_list.invert_match](resources--rate_limiter_policy--reference--group-001.md#canonical-2123203002232112-0020021213011111-3023323311310113-1123200322122321-1022110312312110-0321132201113021-2011231011323033-1123121021023212) |
| `rules.spec.custom_rate_limiter` | [rules.spec.custom_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-2232211100221132-2010233103301203-3033022222030221-0223231213031121-2222123232232211-3311223232100323-3030323010102002-2331232232210110) |
| `rules.spec.custom_rate_limiter.name` | [rules.spec.custom_rate_limiter.name](resources--rate_limiter_policy--reference--group-001.md#canonical-3332320021103031-0312212111331133-0320120111121212-1300303100021301-0331120330320113-3301023032120332-2113013210210320-1010233033020123) |
| `rules.spec.custom_rate_limiter.namespace` | [rules.spec.custom_rate_limiter.namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-2122103131322113-2331111030203110-0232331320030100-2310021131123223-0122013201100103-3231200002303212-3111020321330310-1331011231011333) |
| `rules.spec.custom_rate_limiter.tenant` | [rules.spec.custom_rate_limiter.tenant](resources--rate_limiter_policy--reference--group-001.md#canonical-2331101302220312-1123221322130300-3000310110112131-0311221221230231-2323133023132212-0323331013000220-0013111133000033-2122001203100000) |
| `rules.spec.domain_matcher` | [rules.spec.domain_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-3033230102310111-0233022112210313-2302233232122331-0322030000030213-3312011331123332-1213110103131303-2120201133230011-0030310332233001) |
| `rules.spec.domain_matcher.exact_values` | [rules.spec.domain_matcher.exact_values](resources--rate_limiter_policy--reference--group-001.md#canonical-0211012311122011-3330222202201121-1331323312210223-0002012222233113-0301033120310323-0023120033201302-2211220130211021-1012222231030021) |
| `rules.spec.domain_matcher.regex_values` | [rules.spec.domain_matcher.regex_values](resources--rate_limiter_policy--reference--group-001.md#canonical-1130121200021131-1102200001002200-2212122231230303-0303332030021003-2323010323002321-2113122223133301-2221333031120020-0322100333133123) |
| `rules.spec.headers` | [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-3220011130203200-2012022120312321-1312310113313330-1333331022211011-0312202201233012-2322130033020022-2132022211012111-3022323023023332) |
| `rules.spec.headers.check_not_present` | [rules.spec.headers.check_not_present](resources--rate_limiter_policy--reference--group-001.md#canonical-0211320222001103-1102312013133212-0023120022123031-2300132110301200-3032320000300300-2220020210130323-3303301320201323-2020310132112201) |
| `rules.spec.headers.check_present` | [rules.spec.headers.check_present](resources--rate_limiter_policy--reference--group-001.md#canonical-0233332111102312-3130313333312212-2200013111301010-0112210122201322-1013321002331010-0100121133022130-0333133002023011-2311031300022323) |
| `rules.spec.headers.invert_matcher` | [rules.spec.headers.invert_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-1222331133113032-3331033203301212-2301033211011212-0330222013022011-2112321100132330-2301210110102013-2223012003023000-0101031301202023) |
| `rules.spec.headers.item` | [rules.spec.headers.item](resources--rate_limiter_policy--reference--group-001.md#canonical-3220331330211333-2011331300232300-3012210302001010-1112330113030003-1213030332133103-2112030013330123-3302002232130222-2222120112002130) |
| `rules.spec.headers.item.exact_values` | [rules.spec.headers.item.exact_values](resources--rate_limiter_policy--reference--group-001.md#canonical-2232101002132231-3122013103033022-1202233221211111-2002100311222112-0031312031200201-3123321120120001-1222223032333221-0112110111312020) |
| `rules.spec.headers.item.regex_values` | [rules.spec.headers.item.regex_values](resources--rate_limiter_policy--reference--group-001.md#canonical-2010300301113121-2222211233132212-2333033332002310-3012022300100302-3112113220120232-3233202222331120-2120320103022310-3112333000020231) |
| `rules.spec.headers.item.transformers` | [rules.spec.headers.item.transformers](resources--rate_limiter_policy--reference--group-001.md#canonical-0320111313020210-3311011010232232-0232313132230011-0301330320301233-3032212232020023-3033103212111321-0102112002323111-2330001133111033) |
| `rules.spec.headers.name` | [rules.spec.headers.name](resources--rate_limiter_policy--reference--group-001.md#canonical-3221200021002213-0132222133002232-1030033102230001-3013133120222002-0133030332301021-2113001112213100-1303033232030221-0030101113302103) |
| `rules.spec.http_method` | [rules.spec.http_method](resources--rate_limiter_policy--reference--group-001.md#canonical-3331102020011001-2102030203000330-2313313210002033-3203210313233102-3120311201223212-2120033300001223-1333211011013022-3333332103301003) |
| `rules.spec.http_method.invert_matcher` | [rules.spec.http_method.invert_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-1232003103323100-3332031003110133-2131000011000102-3021220331330221-1223323132102133-3300222002023122-1300002003020131-1021100312312002) |
| `rules.spec.http_method.methods` | [rules.spec.http_method.methods](resources--rate_limiter_policy--reference--group-001.md#canonical-3223310022313111-1213313031020301-3111213213111100-3131301030222223-2122331331121011-1230001232101132-1200000212030212-2101031122010113) |
| `rules.spec.ip_matcher` | [rules.spec.ip_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-3131021000110101-1332023201102001-3103011112302121-0131230000023001-2131010011301101-3110122212110132-2201010012313030-2020312002230012) |
| `rules.spec.ip_matcher.invert_matcher` | [rules.spec.ip_matcher.invert_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-3230103003132120-3120110011332011-2233321222203130-3332230230010100-3232321301210320-1210231211203231-3021332331230321-3131323132202223) |
| `rules.spec.ip_matcher.prefix_sets` | [rules.spec.ip_matcher.prefix_sets](resources--rate_limiter_policy--reference--group-001.md#canonical-3002020023302010-3211003130001020-0122212033002221-1332130313322202-1111333120123233-3323320133123100-1310000023332120-1010333302212331) |
| `rules.spec.ip_matcher.prefix_sets.kind` | [rules.spec.ip_matcher.prefix_sets.kind](resources--rate_limiter_policy--reference--group-001.md#canonical-2221302223110223-1031011000333100-3111010021110100-2100130200330012-2121202231111300-2302220010122331-0213202323203222-3223112313313132) |
| `rules.spec.ip_matcher.prefix_sets.name` | [rules.spec.ip_matcher.prefix_sets.name](resources--rate_limiter_policy--reference--group-001.md#canonical-0310002022311120-0220223133323023-1330311322012113-2333330211203333-1303221130030113-1123303010201210-3310231212212030-0202203312222203) |
| `rules.spec.ip_matcher.prefix_sets.namespace` | [rules.spec.ip_matcher.prefix_sets.namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-2300010221022323-2123221221222223-1220112132031201-1123212200101221-1020322223130213-1312132022313202-1221330103222112-1223310212012030) |
| `rules.spec.ip_matcher.prefix_sets.tenant` | [rules.spec.ip_matcher.prefix_sets.tenant](resources--rate_limiter_policy--reference--group-001.md#canonical-2313201211303202-1120322223032012-1213200223110022-2210320121311322-1300003110213203-2223210003302130-3321220322000002-1213232123032213) |
| `rules.spec.ip_matcher.prefix_sets.uid` | [rules.spec.ip_matcher.prefix_sets.uid](resources--rate_limiter_policy--reference--group-001.md#canonical-0232210223033130-1221301002323223-1001212032211012-2223312213213020-3021302023313111-0030203301311312-2132002310200210-3011102232000031) |
| `rules.spec.ip_prefix_list` | [rules.spec.ip_prefix_list](resources--rate_limiter_policy--reference--group-001.md#canonical-3121112331203133-2111303013123102-1211022010003011-2322300123211123-2003020331313032-2312330320303111-3220231301331113-0231310333020132) |
| `rules.spec.ip_prefix_list.invert_match` | [rules.spec.ip_prefix_list.invert_match](resources--rate_limiter_policy--reference--group-001.md#canonical-1112000001101111-2221203112122321-2101132332310303-2131210121311022-3300021032333230-0121131111022011-1210221021220122-0213221130331101) |
| `rules.spec.ip_prefix_list.ip_prefixes` | [rules.spec.ip_prefix_list.ip_prefixes](resources--rate_limiter_policy--reference--group-001.md#canonical-3031110202220332-0222131332031311-2231130122302002-3233110133323020-0113333311022301-0011020031102210-3033103231302203-0222020101321113) |
| `rules.spec.path` | [rules.spec.path](resources--rate_limiter_policy--reference--group-001.md#canonical-3101112013200311-0331101233312222-1200322013303000-3211303310130031-2032100112312200-1020322123101232-3101012221000112-2133330232121030) |
| `rules.spec.path.encoded_path_matcher` | [rules.spec.path.encoded_path_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-2032302323200130-3123003032202112-3111313302032123-1122133210333312-3032000331222010-1312221011130030-2221323110223130-0100101203313100) |
| `rules.spec.path.exact_values` | [rules.spec.path.exact_values](resources--rate_limiter_policy--reference--group-001.md#canonical-0101112231212312-3003110202110200-2012020113130331-3201023202213111-2031200232033112-0122301312221302-3311011120010232-3321221000300321) |
| `rules.spec.path.invert_matcher` | [rules.spec.path.invert_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-0203013332122320-3023021321321130-2203300133332132-2200112231003221-1220232103032022-3113320130230302-3203220200221110-3220120013101222) |
| `rules.spec.path.prefix_values` | [rules.spec.path.prefix_values](resources--rate_limiter_policy--reference--group-001.md#canonical-2011122333301123-3303111013033002-3311100332113220-3000223010113032-0211131013032020-1130320120030102-3011022002201232-3010303323222121) |
| `rules.spec.path.regex_values` | [rules.spec.path.regex_values](resources--rate_limiter_policy--reference--group-001.md#canonical-0030121132223001-3131330000322300-3311003231021030-2030232013101311-1002033312003121-3110331000113113-2201211020213031-1200123101210123) |
| `rules.spec.path.suffix_values` | [rules.spec.path.suffix_values](resources--rate_limiter_policy--reference--group-001.md#canonical-1132210103212110-1101210232132000-3213030300311131-0100230030201011-0322113031312130-0023321102220311-2323030122300033-3021212213123203) |
| `rules.spec.path.transformers` | [rules.spec.path.transformers](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202000030202-0233002320021033-1213231022013332-1023331310113023-1301310301323322-0200220311310123-0212320012122033-3323320003220232) |
| `rules.spec.segment_policy` | [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-1320013331131320-3002302230231112-0103223330003231-1213113021101011-2123232311002011-0010303232323330-3101013001131322-2211123322031102) |
| `rules.spec.segment_policy.dst_any` | [rules.spec.segment_policy.dst_any](resources--rate_limiter_policy--reference--group-001.md#canonical-1303212221013131-0130212001102103-0131101132233130-0001020333330102-0131030213232012-1112331111010111-2012221021021302-3301012200122002) |
| `rules.spec.segment_policy.dst_segments` | [rules.spec.segment_policy.dst_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-0032131300210313-2101212120220002-3033300132301210-2332121213021212-2333233311010022-0210202221112220-0320030221133313-1220202022310002) |
| `rules.spec.segment_policy.dst_segments.segments` | [rules.spec.segment_policy.dst_segments.segments](resources--rate_limiter_policy--reference--group-001.md#canonical-2130113300012211-2212101131210133-3322132023123321-1202201210331221-0311222201122313-1010021110311333-2030023032003231-0031132123212322) |
| `rules.spec.segment_policy.dst_segments.segments.name` | [rules.spec.segment_policy.dst_segments.segments.name](resources--rate_limiter_policy--reference--group-001.md#canonical-2100110133021310-0112113103132231-3230100322022301-1220233033121233-3330011203233002-0200022130113302-3203021002300333-2120221023013102) |
| `rules.spec.segment_policy.dst_segments.segments.namespace` | [rules.spec.segment_policy.dst_segments.segments.namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-1220212120313301-2121132310123332-1210222220220121-1001021300001113-3301011003103010-2333112032322231-3031302312110312-0000322001013333) |
| `rules.spec.segment_policy.dst_segments.segments.tenant` | [rules.spec.segment_policy.dst_segments.segments.tenant](resources--rate_limiter_policy--reference--group-001.md#canonical-0203011200013120-1300102301302213-2112030012103230-3301301002003323-1213332021112021-1223013302103312-1002102010010100-1222010203223223) |
| `rules.spec.segment_policy.intra_segment` | [rules.spec.segment_policy.intra_segment](resources--rate_limiter_policy--reference--group-001.md#canonical-3133233132012333-3321202103131032-0011211323030230-1233213131330023-1000320022030232-3000033103033010-3302220232233133-0331002323210313) |
| `rules.spec.segment_policy.src_any` | [rules.spec.segment_policy.src_any](resources--rate_limiter_policy--reference--group-001.md#canonical-3211202133133103-1322321322200311-0000011101131313-0313233122220021-1023003121121000-1133332002200113-0212310032002310-3000022232332101) |
| `rules.spec.segment_policy.src_segments` | [rules.spec.segment_policy.src_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-0001102231230010-0332333302000022-1001010002031030-1132223020002332-3030012223321112-2113322112220010-2010202030022001-0133332132221221) |
| `rules.spec.segment_policy.src_segments.segments` | [rules.spec.segment_policy.src_segments.segments](resources--rate_limiter_policy--reference--group-001.md#canonical-3310022203231120-1321101111231030-0333212300133211-1201331010122230-3333231300330010-0201232120210223-0110203011111111-1012111233312322) |
| `rules.spec.segment_policy.src_segments.segments.name` | [rules.spec.segment_policy.src_segments.segments.name](resources--rate_limiter_policy--reference--group-001.md#canonical-0021221011122032-3223313013001203-3130223220312100-1211322221231023-1021333322123113-2102231223012321-3322201321121103-2102133230320213) |
| `rules.spec.segment_policy.src_segments.segments.namespace` | [rules.spec.segment_policy.src_segments.segments.namespace](resources--rate_limiter_policy--reference--group-001.md#canonical-2023101313133101-1323123223222032-1112211321101020-2100130331033021-1112223313300332-1103102000120001-1322332010202101-3021011031200320) |
| `rules.spec.segment_policy.src_segments.segments.tenant` | [rules.spec.segment_policy.src_segments.segments.tenant](resources--rate_limiter_policy--reference--group-001.md#canonical-0231101030213333-2100213303112230-3310033130132303-0033111131020302-0312121031230010-3012310210231033-3302211012023132-3211001002002230) |
| `server_name` | [server_name](resources--rate_limiter_policy--reference--group-001.md#canonical-0100113011002011-0301110020000203-1332313311221231-3321302233233312-2111213221301013-1223133020120101-3032320003233020-1333310310333012) |
| `server_name_matcher` | [server_name_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-0001101311201130-3312030322211210-3100011300130000-2100302121010320-2012130212301321-0013201312012022-2123122022002031-3033313212122332) |
| `server_name_matcher.exact_values` | [server_name_matcher.exact_values](resources--rate_limiter_policy--reference--group-001.md#canonical-0111011221011303-2223311331322200-2321220132112130-0102300130012112-0221033002121122-1233131002001302-3110013203232231-1112110222310333) |
| `server_name_matcher.regex_values` | [server_name_matcher.regex_values](resources--rate_limiter_policy--reference--group-001.md#canonical-3000302322312310-0302310112302221-2333321001033210-2123313001001112-3322310322230222-1113310102231121-2121322101103213-2221033303022100) |
| `server_selector` | [server_selector](resources--rate_limiter_policy--reference--group-001.md#canonical-3333222120133010-2212303113313110-1101120223022023-3033332033033312-0130133232220202-0321213221012311-3311301211321101-0220012223003333) |
| `server_selector.expressions` | [server_selector.expressions](resources--rate_limiter_policy--reference--group-001.md#canonical-1311030133210332-1102012030123233-1130333021103132-3000222032031121-1001302113023311-3212033313212023-0021203031333231-1231110322301130) |
| `timeouts` | [timeouts](resources--rate_limiter_policy--reference--group-001.md#canonical-3212103331233210-2000222302313032-0313013101020010-3110101321112111-3031023103002100-1202332200210132-3212322122100212-1101030313300101) |
| `timeouts.create` | [timeouts.create](resources--rate_limiter_policy--reference--group-001.md#canonical-3032230212001120-1211202232101202-2200012000122202-2100232320220213-1121213132231301-0000111312303030-0203312302222011-3132133312000020) |
| `timeouts.delete` | [timeouts.delete](resources--rate_limiter_policy--reference--group-001.md#canonical-0030331222233313-3011120131022010-0311013330110023-0200032313012010-0332021213230103-2003322311010331-2033222321030313-1123321031013200) |
| `timeouts.read` | [timeouts.read](resources--rate_limiter_policy--reference--group-001.md#canonical-2320100323003002-1212331113103103-2003203303100033-1233023002100233-3311013333300000-1330233120200103-3001311212023222-0213233010130001) |
| `timeouts.update` | [timeouts.update](resources--rate_limiter_policy--reference--group-001.md#canonical-2201030213200030-3311310030101113-3011112333130131-3320321220323232-2332332021113303-0332320022330122-3232111220103103-0002303301103122) |

<a id="canonical-0211213222320023-0301120212121003-2221311230220112-0031320320033020-2032022001201303-0013020220032120-0021032212231130-1003333312100321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `any_server` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- any_server

<a id="canonical-0213131333013020-3031100203003021-3112303110103113-2121121100002232-3022313102101310-3300001233312130-0103220220320030-1102332123211013"></a>

Type: `["object", {}]`. Optional.

\[OneOf: any\_server, server\_name, server\_name\_matcher, server\_selector\] Enable this option

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

OneOf alternatives in this subsection:

- [any_server](resources--rate_limiter_policy--reference--group-001.md#canonical-0213131333013020-3031100203003021-3112303110103113-2121121100002232-3022313102101310-3300001233312130-0103220220320030-1102332123211013)
- [server_name](resources--rate_limiter_policy--reference--group-001.md#canonical-0100113011002011-0301110020000203-1332313311221231-3321302233233312-2111213221301013-1223133020120101-3032320003233020-1333310310333012)
- [server_name_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-0001101311201130-3312030322211210-3100011300130000-2100302121010320-2012130212301321-0013201312012022-2123122022002031-3033313212122332)
- [server_selector](resources--rate_limiter_policy--reference--group-001.md#canonical-3333222120133010-2212303113313110-1101120223022023-3033332033033312-0130133232220202-0321213221012311-3311301211321101-0220012223003333)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_server = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- rules

<a id="canonical-2212121222113012-2012103123200003-1233130300001122-0232000101330331-3213232131130123-3233212310223332-1112100333101032-2113013200123033"></a>

Type: `"object"`. list nested block, Optional.

List of RateLimiterRules that are evaluated sequentially till a matching rule is identified.
Defaults to \`\[\]\`. Server applies default when omitted.

Additional upstream details:

A list of RateLimiterRules that are evaluated sequentially till a matching rule is identified.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112231311011021-0113111203002302-3003102322033001-3133312003020130-1102303330001233-0331330020202133-3331131021033202-1010303333022330"></a>

### Direct properties for `rules`

- [metadata](resources--rate_limiter_policy--reference--group-001.md#canonical-3023012010232323-1023103223030223-2220121223212031-0202030330120331-3212032012123331-0203223023100101-1002003331223120-0322213012210021): complete subsection reference.

- [spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031): complete subsection reference.

<a id="canonical-3023012010232323-1023103223030223-2220121223212031-0202030330120331-3212032012123331-0203223023100101-1002003331223120-0322213012210021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.metadata` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- rules.metadata

<a id="canonical-3032122222030011-2023203100132202-1100212021200001-3201033002310223-1021112200023112-0322113112010300-1332332210002210-1210302203210230"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102330031012211-0001231130301302-1322331100302112-1030203031313020-0100222031201312-1212311013110102-0201300212331030-2231231133110211"></a>

### Direct properties for `rules.metadata`

<a id="canonical-1111121002202121-2302331331100131-0132111132311210-1211212312110302-0113012300010233-3132223131323031-2020321033133310-0003302022013012"></a>

#### `rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2221331011322101-1231130201030122-2031103223101220-2122313301013010-1223211130330213-0210312012201311-1200101331030122-3113333003030023"></a>

<a id="canonical-3330210103131311-1022302303022100-0023210000200220-3233110103320122-2110233122303222-3212303200301001-2002111123331111-1222103231201301"></a>

#### `rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- rules.spec

<a id="canonical-2133012232331210-3031131131232011-0210113022223120-2130210032332203-1230033222223133-0132010132202232-2102121113020123-0003232231132201"></a>

Type: `"object"`. single nested block, Optional.

Rate Limiter Rule Specification. Shape of Rate Limiter Rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_country",
    "country_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("apply_rate_limiter",
    "bypass_rate_limiter"),
  validators.ConflictingObjectAttributes("apply_rate_limiter",
    "custom_rate_limiter"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("bypass_rate_limiter",
    "custom_rate_limiter"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-action_choice": "[\"apply_rate_limiter\",\"bypass_rate_limiter\",\"custom_rate_limiter\"]",
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-country_choice": "[\"any_country\",\"country_list\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212001011310231-1323121301221012-2111202232103332-0122010213010131-1203110210312000-2011331100102222-1101120020332332-0103331303111003"></a>

### Direct properties for `rules.spec`

- [any_asn](resources--rate_limiter_policy--reference--group-001.md#canonical-1013323220310111-0232222201001333-3133202000330200-1020020102312030-0300112110013211-1010111121330020-1313000013311301-2013230011030310): complete subsection reference.

- [any_country](resources--rate_limiter_policy--reference--group-001.md#canonical-3130220013132120-3231030013111330-2321300131000313-3223300330101230-1003230032103203-0011201222111123-0211333213223032-2001333223332020): complete subsection reference.

- [any_ip](resources--rate_limiter_policy--reference--group-001.md#canonical-1132231223032231-1022320133231322-2333210310222202-2201320103303031-2200112113231301-0212313312322233-3012023013002111-3111233313333121): complete subsection reference.

- [apply_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-3321310131202033-1221212002212021-0313200003110223-0100011023002103-3012212302313111-2132320202030101-1301113303131203-3022102210131311): complete subsection reference.

- [asn_list](resources--rate_limiter_policy--reference--group-001.md#canonical-3010103313212020-1101012202233132-1330330033300022-2311310200023033-3211203223232131-3320302101031231-0211130030203212-1231223121103332): complete subsection reference.

- [asn_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-1121002311123100-0122300203123223-3001131202202332-1011231120230130-0301100223033021-3113303010201322-1302112301323113-0212311213102221): complete subsection reference.

- [bypass_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-3022322213330233-0231202100023331-3200222121011103-2033231011132032-3233222020200330-1310313113302210-2333012103310303-1030113210010233): complete subsection reference.

- [country_list](resources--rate_limiter_policy--reference--group-001.md#canonical-1001100110131203-2022313122101031-3021020211203013-2011031120032103-2231210132312322-3012333221123303-3120110210323210-1332330233102031): complete subsection reference.

- [custom_rate_limiter](resources--rate_limiter_policy--reference--group-001.md#canonical-2130031031003201-3121021221313123-3212330332132031-3110303201002111-0022201121003120-3133032022030132-2210232020031333-2303020110311211): complete subsection reference.

- [domain_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-3300021332133333-1203323210333122-1332203202013001-2200321311220221-1112022320113120-3321332013312023-3310000020203121-3203013032022002): complete subsection reference.

- [headers](resources--rate_limiter_policy--reference--group-001.md#canonical-2221320321300321-2330321211030322-1123313203310202-0223000201333132-2300022030130203-1112330202302230-1331201200001010-2121212213033013): complete subsection reference.

- [http_method](resources--rate_limiter_policy--reference--group-001.md#canonical-2331130021113233-3222011222131023-1103132231122310-2111003103322212-0322230210113302-3111101213101230-0032132033102313-0202023221133112): complete subsection reference.

- [ip_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-3111131012320123-0102312201010201-0301133000221330-3213123012032312-1103133123133103-2230203102333021-2031120230010221-0130121221122012): complete subsection reference.

- [ip_prefix_list](resources--rate_limiter_policy--reference--group-001.md#canonical-1022003110312122-0131332020212300-2232023332132000-3020320211320300-0230101300230032-1301303223231100-0120210032032222-1100133300211300): complete subsection reference.

- [path](resources--rate_limiter_policy--reference--group-001.md#canonical-2131203301013120-3032013232130102-0212001020202111-1313312001121213-0121313301000233-0313201111113233-0323023020321211-3110123003332321): complete subsection reference.

- [segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-1003300311131110-2122022122210003-0233032231302312-0120003132000133-2322322113133003-2322331131222131-3112023232001233-0030221022001312): complete subsection reference.

<a id="canonical-1013323220310111-0232222201001333-3133202000330200-1020020102312030-0300112110013211-1010111121330020-1313000013311301-2013230011030310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.any_asn` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.any_asn

<a id="canonical-1013030322310323-1131301033112233-3202012233311011-3111220320111100-0012303123113030-3331130133213022-0330221112130002-1312032000222123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any_asn = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130220013132120-3231030013111330-2321300131000313-3223300330101230-1003230032103203-0011201222111123-0211333213223032-2001333223332020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.any_country` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.any_country

<a id="canonical-3013213001112330-3102122232330110-2111323032331013-2010002202133310-3102101110301010-1200021113231033-1101221233010132-0011301333312212"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for any country.

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

Terraform syntax:

```terraform
any_country = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132231223032231-1022320133231322-2333210310222202-2201320103303031-2200112113231301-0212313312322233-3012023013002111-3111233313333121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.any_ip` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.any_ip

<a id="canonical-3002312003113100-3131303022320201-3332022030031322-3132001000011333-2032312132323010-3120213213033121-3212030120021132-3230011311230122"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321310131202033-1221212002212021-0313200003110223-0100011023002103-3012212302313111-2132320202030101-1301113303131203-3022102210131311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.apply_rate_limiter` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.apply_rate_limiter

<a id="canonical-3132220321033102-0020311131221310-0222133023033312-2031122302130001-2003213122133201-2123302120010111-3230122030130021-3022021012310302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for apply rate limiter.

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

Terraform syntax:

```terraform
apply_rate_limiter = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010103313212020-1101012202233132-1330330033300022-2311310200023033-3211203223232131-3320302101031231-0211130030203212-1231223121103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.asn_list` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.asn_list

<a id="canonical-3200121301010133-1311322013330202-3323030323323121-3122202012223111-0133330311012323-2222331332110213-0230021130222100-2020033331322102"></a>

Type: `"object"`. single nested block, Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011113323010002-1123032032131231-1101232012302312-0002222332313220-2033120331323023-1220133223013222-0323200021220111-1021232121010130"></a>

### Direct properties for `rules.spec.asn_list`

<a id="canonical-2121203321112222-1233333023230300-2231021123022010-3301301131330013-2133020233211210-2132020300322203-3210033122320012-1023222220312223"></a>

#### `rules.spec.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

<a id="canonical-1121002311123100-0122300203123223-3001131202202332-1011231120230130-0301100223033021-3113303010201322-1302112301323113-0212311213102221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.asn_matcher` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.asn_matcher

<a id="canonical-1020113210302223-3303233033220111-1032231321303132-2012203110032021-0303021202233212-0203203112020300-0021130233123022-1023113112000302"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031201232222100-2030203010300011-0113323210310032-0200120202210003-2333203130010232-2310131123120212-3231101300010020-3032113231130131"></a>

### Direct properties for `rules.spec.asn_matcher`

- [asn_sets](resources--rate_limiter_policy--reference--group-001.md#canonical-3222300213221013-1200300102202312-3121113313322213-1120030220021013-3032230013322330-2320133213300330-0003131012323102-0323011223112201): complete subsection reference.

<a id="canonical-3222300213221013-1200300102202312-3121113313322213-1120030220021013-3032230013322330-2320133213300330-0003131012323102-0323011223112201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.asn_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-1121002311123100-0122300203123223-3001131202202332-1011231120230130-0301100223033021-3113303010201322-1302112301323113-0212311213102221)
- rules.spec.asn_matcher.asn_sets

<a id="canonical-0013021332331223-2001221023310231-3313203213330201-1000331231311132-3021213001313301-0000031023200323-2230222202232223-2120310001221331"></a>

Type: `"object"`. list nested block, Optional.

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113200001112111-1323322133312330-3011230032032110-1122123221233231-3300311101101002-2230003131300230-3231030301333210-3132033033303133"></a>

### Direct properties for `rules.spec.asn_matcher.asn_sets`

<a id="canonical-3013122300032012-1031102232213020-0122020120021112-3023022223332203-2013221310232221-3212300131212123-2332000211303022-1212312112031021"></a>

#### `rules.spec.asn_matcher.asn_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-2013303022312202-0311221330303120-2232212323001330-1333023211101123-0323000300000212-3231313302103212-0301100331333132-1200120203020221"></a>

<a id="canonical-2113110230001303-1301011201321210-3002222231222232-2102233311203200-1103322303223130-2031222310013012-3300123021122220-3132000203223023"></a>

#### `rules.spec.asn_matcher.asn_sets.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2121222113301033-0022100110120302-0012131123313031-0303222200001213-1130312002112031-2300022022233110-2002112201333103-3222311303011222"></a>

<a id="canonical-2030001010313310-0020021203223212-1010110020232111-1222113203013220-3231212222320220-3223002223312033-1303222313001012-3213110131010013"></a>

#### `rules.spec.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-0312211112133101-3301120131203231-3123012200001221-0201212011013332-2331310001322311-0012313013021222-1101200223030131-0121301023221321"></a>

<a id="canonical-0133220320012021-3232003211130023-0021012130331133-1233112232010100-2330213213300000-0121220130013203-3233130101233202-0301320301133130"></a>

#### `rules.spec.asn_matcher.asn_sets.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0203001320233221-0322022121222331-2331333231010310-3101233013333101-3132323121003211-3102001031332233-3311223103220320-0302230010232120"></a>

<a id="canonical-0300110322021230-0101133222310011-0322211210311310-1301211311300102-1211123201131212-0011012132210203-1203101113200301-0112133011330110"></a>

#### `rules.spec.asn_matcher.asn_sets.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-3022322213330233-0231202100023331-3200222121011103-2033231011132032-3233222020200330-1310313113302210-2333012103310303-1030113210010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.bypass_rate_limiter` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.bypass_rate_limiter

<a id="canonical-2333232022012033-3203002233023322-3030222322333333-3222212000020312-3111333002122030-0221211012202112-1221323221102022-1013222300230220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for bypass rate limiter.

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

Terraform syntax:

```terraform
bypass_rate_limiter = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001100110131203-2022313122101031-3021020211203013-2011031120032103-2231210132312322-3012333221123303-3120110210323210-1332330233102031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.country_list` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.country_list

<a id="canonical-0301133001110111-1232331231311332-2330120000332200-1032112012010330-1133230101130301-0032322332332210-3301022200332101-0313332232031201"></a>

Type: `"object"`. single nested block, Optional.

Country Codes List. List of Country Codes to match against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("country_codes")}
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
country_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011110132011300-0301133313231130-3221220230300113-0200110011101202-2013310321021333-2020332212303102-0131203203033002-1233332333221213"></a>

### Direct properties for `rules.spec.country_list`

<a id="canonical-1112202200110230-2221113113130013-1121200203130230-0121321312103213-3300121333103323-1011311110020300-0003300311102321-0000213133011200"></a>

#### `rules.spec.country_list.country_codes` property

Type: `["list", "string"]`. Optional.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Country Codes List. List of Country Codes. Possible values are \`COUNTRY\_NONE\`, \`COUNTRY\_AD\`,
\`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`, \`COUNTRY\_AI\`, \`COUNTRY\_AL\`,
\`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`, \`COUNTRY\_AQ\`, \`COUNTRY\_AR\`,
\`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`, \`COUNTRY\_AW\`, \`COUNTRY\_AX\`,
\`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`, \`COUNTRY\_BD\`, \`COUNTRY\_BE\`,
\`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`, \`COUNTRY\_BI\`, \`COUNTRY\_BJ\`,
\`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`, \`COUNTRY\_BO\`, \`COUNTRY\_BQ\`,
\`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`, \`COUNTRY\_BV\`, \`COUNTRY\_BW\`,
\`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`, \`COUNTRY\_CC\`, \`COUNTRY\_CD\`,
\`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`, \`COUNTRY\_CI\`, \`COUNTRY\_CK\`,
\`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`, \`COUNTRY\_CO\`, \`COUNTRY\_CR\`,
\`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`, \`COUNTRY\_CW\`, \`COUNTRY\_CX\`,
\`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`, \`COUNTRY\_DJ\`, \`COUNTRY\_DK\`,
\`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`, \`COUNTRY\_EC\`, \`COUNTRY\_EE\`,
\`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`, \`COUNTRY\_ES\`, \`COUNTRY\_ET\`,
\`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`, \`COUNTRY\_FM\`, \`COUNTRY\_FO\`,
\`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`, \`COUNTRY\_GD\`, \`COUNTRY\_GE\`,
\`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`, \`COUNTRY\_GI\`, \`COUNTRY\_GL\`,
\`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`, \`COUNTRY\_GQ\`, \`COUNTRY\_GR\`,
\`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`, \`COUNTRY\_GW\`, \`COUNTRY\_GY\`,
\`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`, \`COUNTRY\_HR\`, \`COUNTRY\_HT\`,
\`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`, \`COUNTRY\_IL\`, \`COUNTRY\_IM\`,
\`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`, \`COUNTRY\_IR\`, \`COUNTRY\_IS\`,
\`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`, \`COUNTRY\_JO\`, \`COUNTRY\_JP\`,
\`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`, \`COUNTRY\_KI\`, \`COUNTRY\_KM\`,
\`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`, \`COUNTRY\_KW\`, \`COUNTRY\_KY\`,
\`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`, \`COUNTRY\_LC\`, \`COUNTRY\_LI\`,
\`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`, \`COUNTRY\_LT\`, \`COUNTRY\_LU\`,
\`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`, \`COUNTRY\_MC\`, \`COUNTRY\_MD\`,
\`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`, \`COUNTRY\_MH\`, \`COUNTRY\_MK\`,
\`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`, \`COUNTRY\_MO\`, \`COUNTRY\_MP\`,
\`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`, \`COUNTRY\_MT\`, \`COUNTRY\_MU\`,
\`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`, \`COUNTRY\_MY\`, \`COUNTRY\_MZ\`,
\`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`, \`COUNTRY\_NF\`, \`COUNTRY\_NG\`,
\`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`, \`COUNTRY\_NP\`, \`COUNTRY\_NR\`,
\`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`, \`COUNTRY\_PA\`, \`COUNTRY\_PE\`,
\`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`, \`COUNTRY\_PK\`, \`COUNTRY\_PL\`,
\`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`, \`COUNTRY\_PS\`, \`COUNTRY\_PT\`,
\`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`, \`COUNTRY\_RE\`, \`COUNTRY\_RO\`,
\`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`, \`COUNTRY\_SA\`, \`COUNTRY\_SB\`,
\`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`, \`COUNTRY\_SG\`, \`COUNTRY\_SH\`,
\`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`, \`COUNTRY\_SL\`, \`COUNTRY\_SM\`,
\`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`, \`COUNTRY\_SS\`, \`COUNTRY\_ST\`,
\`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`, \`COUNTRY\_SZ\`, \`COUNTRY\_TC\`,
\`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`, \`COUNTRY\_TH\`, \`COUNTRY\_TJ\`,
\`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`, \`COUNTRY\_TN\`, \`COUNTRY\_TO\`,
\`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`, \`COUNTRY\_TW\`, \`COUNTRY\_TZ\`,
\`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`, \`COUNTRY\_US\`, \`COUNTRY\_UY\`,
\`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`, \`COUNTRY\_VE\`, \`COUNTRY\_VG\`,
\`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`, \`COUNTRY\_WF\`, \`COUNTRY\_WS\`,
\`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`, \`COUNTRY\_YT\`, \`COUNTRY\_ZA\`,
\`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2123203002232112-0020021213011111-3023323311310113-1123200322122321-1022110312312110-0321132201113021-2011231011323033-1123121021023212"></a>

<a id="canonical-3030113122322330-1121012221031303-1210210222100010-3323022101132232-3102333312230201-2012113213132132-3121021223120110-2330100112230102"></a>

#### `rules.spec.country_list.invert_match` property

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

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

<a id="canonical-2130031031003201-3121021221313123-3212330332132031-3110303201002111-0022201121003120-3133032022030132-2210232020031333-2303020110311211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.custom_rate_limiter` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.custom_rate_limiter

<a id="canonical-2232211100221132-2010233103301203-3033022222030221-0223231213031121-2222123232232211-3311223232100323-3030323010102002-2331232232210110"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
custom_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311112123121301-0032213221220020-2332002331130121-2131203110303223-2102321122133000-2323012213033221-2132200002102122-2131031023111120"></a>

### Direct properties for `rules.spec.custom_rate_limiter`

<a id="canonical-3332320021103031-0312212111331133-0320120111121212-1300303100021301-0331120330320113-3301023032120332-2113013210210320-1010233033020123"></a>

#### `rules.spec.custom_rate_limiter.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2122103131322113-2331111030203110-0232331320030100-2310021131123223-0122013201100103-3231200002303212-3111020321330310-1331011231011333"></a>

<a id="canonical-0031321223100220-3301221103203311-2220021331130110-3100020331000201-0210300211331020-2110202223330120-2303222121230122-3012011133121223"></a>

#### `rules.spec.custom_rate_limiter.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2331101302220312-1123221322130300-3000310110112131-0311221221230231-2323133023132212-0323331013000220-0013111133000033-2122001203100000"></a>

<a id="canonical-3333130123132132-0213012131222332-3202020311213210-1102113103203231-2303030023120002-1303321111111313-0110132020003032-2103310222122002"></a>

#### `rules.spec.custom_rate_limiter.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3300021332133333-1203323210333122-1332203202013001-2200321311220221-1112022320113120-3321332013312023-3310000020203121-3203013032022002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.domain_matcher` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.domain_matcher

<a id="canonical-3033230102310111-0233022112210313-2302233232122331-0322030000030213-3312011331123332-1213110103131303-2120201133230011-0030310332233001"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
domain_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003120022330013-1111223311313001-1020211120133311-3233113032032001-0013300030113023-1230320213322332-2022333002212323-3310002000033012"></a>

### Direct properties for `rules.spec.domain_matcher`

<a id="canonical-0211012311122011-3330222202201121-1331323312210223-0002012222233113-0301033120310323-0023120033201302-2211220130211021-1012222231030021"></a>

#### `rules.spec.domain_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1130121200021131-1102200001002200-2212122231230303-0303332030021003-2323010323002321-2113122223133301-2221333031120020-0322100333133123"></a>

<a id="canonical-1223212122310320-3313300210000112-0320312222033113-3102132022020321-1130022002320320-0300311122310202-3211320233210333-3302033213022101"></a>

#### `rules.spec.domain_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2221320321300321-2330321211030322-1123313203310202-0223000201333132-2300022030130203-1112330202302230-1331201200001010-2121212213033013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.headers` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.headers

<a id="canonical-3220011130203200-2012022120312321-1312310113313330-1333331022211011-0312202201233012-2322130033020022-2132022211012111-3022323023023332"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
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
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012303113132100-1233202232001233-2032222202113220-2203100322210301-1022002331011323-1033031210032111-1102100130201221-2110030221232030"></a>

### Direct properties for `rules.spec.headers`

- [check_not_present](resources--rate_limiter_policy--reference--group-001.md#canonical-0133223033233121-2010023303012120-0320100103321222-3220311322031003-1113222110010200-0130101032030312-1023102033303333-0121303130122122): complete subsection reference.

- [check_present](resources--rate_limiter_policy--reference--group-001.md#canonical-1110212301113113-2111313023331021-1011020130333033-3131312210030013-3022300033330223-1113310301311013-2103113033103232-2130301311322120): complete subsection reference.

<a id="canonical-1222331133113032-3331033203301212-2301033211011212-0330222013022011-2112321100132330-2301210110102013-2223012003023000-0101031301202023"></a>

<a id="canonical-3323222001101130-0130013230312332-1010020333021311-1123223023133013-1231323020001121-1021031122100030-0332230112033302-3001200300112010"></a>

#### `rules.spec.headers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

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

- [item](resources--rate_limiter_policy--reference--group-001.md#canonical-2323220313112011-3322000013022033-1021002100120221-2311101201131121-2000201120330132-3333013210320130-1203100112233232-1130331000222332): complete subsection reference.

<a id="canonical-3221200021002213-0132222133002232-1030033102230001-3013133120222002-0133030332301021-2113001112213100-1303033232030221-0030101113302103"></a>

<a id="canonical-2131101332010100-3022220231220001-0213320230212102-2300002223221132-2203010322223020-3011103002110212-3300123211333320-3033323121130132"></a>

#### `rules.spec.headers.name` property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0133223033233121-2010023303012120-0320100103321222-3220311322031003-1113222110010200-0130101032030312-1023102033303333-0121303130122122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-2221320321300321-2330321211030322-1123313203310202-0223000201333132-2300022030130203-1112330202302230-1331201200001010-2121212213033013)
- rules.spec.headers.check_not_present

<a id="canonical-0211320222001103-1102312013133212-0023120022123031-2300132110301200-3032320000300300-2220020210130323-3303301320201323-2020310132112201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110212301113113-2111313023331021-1011020130333033-3131312210030013-3022300033330223-1113310301311013-2103113033103232-2130301311322120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.headers.check_present` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-2221320321300321-2330321211030322-1123313203310202-0223000201333132-2300022030130203-1112330202302230-1331201200001010-2121212213033013)
- rules.spec.headers.check_present

<a id="canonical-0233332111102312-3130313333312212-2200013111301010-0112210122201322-1013321002331010-0100121133022130-0333133002023011-2311031300022323"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323220313112011-3322000013022033-1021002100120221-2311101201131121-2000201120330132-3333013210320130-1203100112233232-1130331000222332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.headers.item` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.headers](resources--rate_limiter_policy--reference--group-001.md#canonical-2221320321300321-2330321211030322-1123313203310202-0223000201333132-2300022030130203-1112330202302230-1331201200001010-2121212213033013)
- rules.spec.headers.item

<a id="canonical-3220331330211333-2011331300232300-3012210302001010-1112330113030003-1213030332133103-2112030013330123-3302002232130222-2222120112002130"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210221002103210-3330033331012222-0203001002110223-1130321110022220-1113303111302313-3302332220133232-3110210203001310-2002000132302121"></a>

### Direct properties for `rules.spec.headers.item`

<a id="canonical-2232101002132231-3122013103033022-1202233221211111-2002100311222112-0031312031200201-3123321120120001-1222223032333221-0112110111312020"></a>

#### `rules.spec.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2010300301113121-2222211233132212-2333033332002310-3012022300100302-3112113220120232-3233202222331120-2120320103022310-3112333000020231"></a>

<a id="canonical-3223230300232102-3032022201120033-0023033322021203-1201300201302110-0222313030133123-2302331312331231-3023100333330333-2230200001211302"></a>

#### `rules.spec.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0320111313020210-3311011010232232-0232313132230011-0301330320301233-3032212232020023-3033103212111321-0102112002323111-2330001133111033"></a>

<a id="canonical-1112210322221323-1000030203312103-1203300330001200-0131333222223310-1131332022113100-2332313220332031-1200301302021012-2210323122223320"></a>

#### `rules.spec.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2331130021113233-3222011222131023-1103132231122310-2111003103322212-0322230210113302-3111101213101230-0032132033102313-0202023221133112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.http_method` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.http_method

<a id="canonical-3331102020011001-2102030203000330-2313313210002033-3203210313233102-3120311201223212-2120033300001223-1333211011013022-3333332103301003"></a>

Type: `"object"`. single nested block, Optional.

An HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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
http_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301033201202032-2102100210130021-3301331220103102-0212303013032111-0330121313100101-0300100110111231-3321012010211121-3302231130230213"></a>

### Direct properties for `rules.spec.http_method`

<a id="canonical-1232003103323100-3332031003110133-2131000011000102-3021220331330221-1223323132102133-3300222002023122-1300002003020131-1021100312312002"></a>

#### `rules.spec.http_method.invert_matcher` property

Type: `"bool"`. Optional.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-3223310022313111-1213313031020301-3111213213111100-3131301030222223-2122331331121011-1230001232101132-1200000212030212-2101031122010113"></a>

<a id="canonical-1300323031220321-0101133203112013-2232332230033201-3203102133311221-0122001020311130-1333300211032311-1010231003002032-3301022022310212"></a>

#### `rules.spec.http_method.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3111131012320123-0102312201010201-0301133000221330-3213123012032312-1103133123133103-2230203102333021-2031120230010221-0130121221122012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.ip_matcher` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.ip_matcher

<a id="canonical-3131021000110101-1332023201102001-3103011112302121-0131230000023001-2131010011301101-3110122212110132-2201010012313030-2020312002230012"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001012223002102-0333330300211330-0302100202311123-2211122203120200-3211123033321221-3202230132230023-2031010320202323-0230302230120230"></a>

### Direct properties for `rules.spec.ip_matcher`

<a id="canonical-3230103003132120-3120110011332011-2233321222203130-3332230230010100-3232321301210320-1210231211203231-3021332331230321-3131323132202223"></a>

#### `rules.spec.ip_matcher.invert_matcher` property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](resources--rate_limiter_policy--reference--group-001.md#canonical-1122123220101132-0201033030320202-2332110100013011-0320121010333111-2232011333133233-0232200003023120-2120223011323023-1331230302111233): complete subsection reference.

<a id="canonical-1122123220101132-0201033030320202-2332110100013011-0320121010333111-2232011333133233-0232200003023120-2120223011323023-1331230302111233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.ip_matcher](resources--rate_limiter_policy--reference--group-001.md#canonical-3111131012320123-0102312201010201-0301133000221330-3213123012032312-1103133123133103-2230203102333021-2031120230010221-0130121221122012)
- rules.spec.ip_matcher.prefix_sets

<a id="canonical-3002020023302010-3211003130001020-0122212033002221-1332130313322202-1111333120123233-3323320133123100-1310000023332120-1010333302212331"></a>

Type: `"object"`. list nested block, Optional.

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032320022012333-0330003320223011-2110001023011333-3021333312023103-3012333112020122-1120010321330321-1030201220131133-2112320110000312"></a>

### Direct properties for `rules.spec.ip_matcher.prefix_sets`

<a id="canonical-2221302223110223-1031011000333100-3111010021110100-2100130200330012-2121202231111300-2302220010122331-0213202323203222-3223112313313132"></a>

#### `rules.spec.ip_matcher.prefix_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-0310002022311120-0220223133323023-1330311322012113-2333330211203333-1303221130030113-1123303010201210-3310231212212030-0202203312222203"></a>

<a id="canonical-1330323233101112-0030111303103311-1100101033321022-0322311103132112-3032313210231333-2230002211021132-3110113321100331-1012312302030121"></a>

#### `rules.spec.ip_matcher.prefix_sets.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2300010221022323-2123221221222223-1220112132031201-1123212200101221-1020322223130213-1312132022313202-1221330103222112-1223310212012030"></a>

<a id="canonical-0121031102100232-3300020032012300-1013101212312030-1131303213022203-1132101132001312-2102101203333322-3300020001231012-0223333333013031"></a>

#### `rules.spec.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-2313201211303202-1120322223032012-1213200223110022-2210320121311322-1300003110213203-2223210003302130-3321220322000002-1213232123032213"></a>

<a id="canonical-3011211121311233-2001323203112221-3010312131030230-0103113110010022-3111300312101121-1121322122232101-3133103010021120-2023100023200101"></a>

#### `rules.spec.ip_matcher.prefix_sets.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0232210223033130-1221301002323223-1001212032211012-2223312213213020-3021302023313111-0030203301311312-2132002310200210-3011102232000031"></a>

<a id="canonical-2202323202221113-1112222312112132-1110230133332321-2012010331033211-0221332120000100-2102032310212111-1133311310221311-0303300101233012"></a>

#### `rules.spec.ip_matcher.prefix_sets.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-1022003110312122-0131332020212300-2232023332132000-3020320211320300-0230101300230032-1301303223231100-0120210032032222-1100133300211300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.ip_prefix_list

<a id="canonical-3121112331203133-2111303013123102-1211022010003011-2322300123211123-2003020331313032-2312330320303111-3220231301331113-0231310333020132"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022230310320233-1321121202031030-1012112131030301-1323032032012332-3202310310100101-2020300033232313-0001122222120022-1302320212202233"></a>

### Direct properties for `rules.spec.ip_prefix_list`

<a id="canonical-1112000001101111-2221203112122321-2101132332310303-2131210121311022-3300021032333230-0121131111022011-1210221021220122-0213221130331101"></a>

#### `rules.spec.ip_prefix_list.invert_match` property

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

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

<a id="canonical-3031110202220332-0222131332031311-2231130122302002-3233110133323020-0113333311022301-0011020031102210-3033103231302203-0222020101321113"></a>

<a id="canonical-0121031301132333-0010133212313210-1113101013131232-3233323223113323-0303210231021222-3302122031033001-0030000010023000-3322202210121223"></a>

#### `rules.spec.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2131203301013120-3032013232130102-0212001020202111-1313312001121213-0121313301000233-0313201111113233-0323023020321211-3110123003332321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.path` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.path

<a id="canonical-3101112013200311-0331101233312222-1200322013303000-3211303310130031-2032100112312200-1020322123101232-3101012221000112-2133330232121030"></a>

Type: `"object"`. single nested block, Optional.

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

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
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001330003002222-1211111103122300-0002103023301131-0303333000022112-3013113202110011-1101303101210110-2102020303212123-2112133113310300"></a>

### Direct properties for `rules.spec.path`

<a id="canonical-2032302323200130-3123003032202112-3111313302032123-1122133210333312-3032000331222010-1312221011130030-2221323110223130-0100101203313100"></a>

#### `rules.spec.path.encoded_path_matcher` property

Type: `"bool"`. Optional.

Match against the encoded, escaped path.

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

<a id="canonical-0101112231212312-3003110202110200-2012020113130331-3201023202213111-2031200232033112-0122301312221302-3311011120010232-3321221000300321"></a>

<a id="canonical-3120003332012131-0203132222023223-1203223131233212-1322101131013311-0101002310211110-3121210232230123-2110130011331201-2312110323032203"></a>

#### `rules.spec.path.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact path values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0203013332122320-3023021321321130-2203300133332132-2200112231003221-1220232103032022-3113320130230302-3203220200221110-3220120013101222"></a>

<a id="canonical-0020011111113311-3220113023003330-2133030122112320-1222001220120001-3332302003310121-3023103210103021-3100012032031331-0111010233311120"></a>

#### `rules.spec.path.invert_matcher` property

Type: `"bool"`. Optional.

Invert Path Matcher. Invert the match result.

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

<a id="canonical-2011122333301123-3303111013033002-3311100332113220-3000223010113032-0211131013032020-1130320120030102-3011022002201232-3010303323222121"></a>

<a id="canonical-1133002120211231-2301003300002333-0313111223103110-2132123131321330-2211020102112033-2322032022213102-2001223231210130-0110311331323200"></a>

#### `rules.spec.path.prefix_values` property

Type: `["list", "string"]`. Optional.

A list of path prefix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0030121132223001-3131330000322300-3311003231021030-2030232013101311-1002033312003121-3110331000113113-2201211020213031-1200123101210123"></a>

<a id="canonical-2231223103200203-3101131211110003-2321101021123012-0331130333000303-2331221213231332-2030002112031023-3211032310101120-3221021323012012"></a>

#### `rules.spec.path.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1132210103212110-1101210232132000-3213030300311131-0100230030201011-0322113031312130-0023321102220311-2323030122300033-3021212213123203"></a>

<a id="canonical-1000323313233233-2032211110311113-3213011000301221-0112023011000320-2212302222102301-0300113133200132-3130102123330231-1100123001100200"></a>

#### `rules.spec.path.suffix_values` property

Type: `["list", "string"]`. Optional.

A list of path suffix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0132202000030202-0233002320021033-1213231022013332-1023331310113023-1301310301323322-0200220311310123-0212320012122033-3323320003220232"></a>

<a id="canonical-2010313313013133-1330233000333021-3133113012213011-3020000021121130-1211100032022202-0120101012010300-0313110131001302-1102020230101113"></a>

#### `rules.spec.path.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1003300311131110-2122022122210003-0233032231302312-0120003132000133-2322322113133003-2322331131222131-3112023232001233-0030221022001312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.segment_policy` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- rules.spec.segment_policy

<a id="canonical-1320013331131320-3002302230231112-0103223330003231-1213113021101011-2123232311002011-0010303232323330-3101013001131322-2211123322031102"></a>

Type: `"object"`. single nested block, Optional.

Configure source and destination segment for policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("dst_any",
    "dst_segments"),
  validators.ConflictingObjectAttributes("dst_any",
    "intra_segment"),
  validators.ConflictingObjectAttributes("dst_segments",
    "intra_segment"),
  validators.ConflictingObjectAttributes("src_any",
    "src_segments")}
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
  "x-ves-oneof-field-dst_segment_choice": "[\"dst_any\",\"dst_segments\",\"intra_segment\"]",
  "x-ves-oneof-field-src_segment_choice": "[\"src_any\",\"src_segments\"]"
}
```

Terraform syntax:

```terraform
segment_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000122011200001-1323010111323231-2020333011132033-1011233230033220-2120330313200200-2130213200311302-0021330213233102-2032002111103131"></a>

### Direct properties for `rules.spec.segment_policy`

- [dst_any](resources--rate_limiter_policy--reference--group-001.md#canonical-2221020101101013-2322303003003121-0222311101103230-2203322321313003-0231000101210230-1323100102311133-1321312120001110-0120202232013312): complete subsection reference.

- [dst_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-3322200110122100-0203202310213222-3302203020101330-2030302212120133-3300212213333233-1223033321020303-3002122332133101-1203212010210103): complete subsection reference.

- [intra_segment](resources--rate_limiter_policy--reference--group-001.md#canonical-1031111222233131-3030031320222313-1011000131211231-0212002101013303-2113213210020111-3011323310000203-0133311330023233-2212132203310000): complete subsection reference.

- [src_any](resources--rate_limiter_policy--reference--group-001.md#canonical-1320312221031301-0033011221303221-0111313210211221-0212120001322103-1302211211222002-1133133233133021-2310223023131111-3330233113322320): complete subsection reference.

- [src_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-0022310130211210-2222210220122323-2122321023133332-2031120230030333-1333330100313232-2002132113300110-0200120130020032-3101313230103323): complete subsection reference.

<a id="canonical-2221020101101013-2322303003003121-0222311101103230-2203322321313003-0231000101210230-1323100102311133-1321312120001110-0120202232013312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.segment_policy.dst_any` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-1003300311131110-2122022122210003-0233032231302312-0120003132000133-2322322113133003-2322331131222131-3112023232001233-0030221022001312)
- rules.spec.segment_policy.dst_any

<a id="canonical-1303212221013131-0130212001102103-0131101132233130-0001020333330102-0131030213232012-1112331111010111-2012221021021302-3301012200122002"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
dst_any = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322200110122100-0203202310213222-3302203020101330-2030302212120133-3300212213333233-1223033321020303-3002122332133101-1203212010210103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.segment_policy.dst_segments` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-1003300311131110-2122022122210003-0233032231302312-0120003132000133-2322322113133003-2322331131222131-3112023232001233-0030221022001312)
- rules.spec.segment_policy.dst_segments

<a id="canonical-0032131300210313-2101212120220002-3033300132301210-2332121213021212-2333233311010022-0210202221112220-0320030221133313-1220202022310002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dst segments.

Additional upstream details:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("segments")}
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
dst_segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133011021032321-0212013303303120-3023312101131012-1002110110323210-1123230313110102-3303232322100230-1122001200103001-0013202231212333"></a>

### Direct properties for `rules.spec.segment_policy.dst_segments`

- [segments](resources--rate_limiter_policy--reference--group-001.md#canonical-0133312222123001-0233120200121131-0111000101031312-1333311111120101-2121320123110300-3101122130311303-3001010030011233-0010310133100022): complete subsection reference.

<a id="canonical-0133312222123001-0233120200121131-0111000101031312-1333311111120101-2121320123110300-3101122130311303-3001010030011233-0010310133100022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.segment_policy.dst_segments.segments` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-1003300311131110-2122022122210003-0233032231302312-0120003132000133-2322322113133003-2322331131222131-3112023232001233-0030221022001312)
- [rules.spec.segment_policy.dst_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-3322200110122100-0203202310213222-3302203020101330-2030302212120133-3300212213333233-1223033321020303-3002122332133101-1203212010210103)
- rules.spec.segment_policy.dst_segments.segments

<a id="canonical-2130113300012211-2212101131210133-3322132023123321-1202201210331221-0311222201122313-1010021110311333-2030023032003231-0031132123212322"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

Terraform syntax:

```terraform
segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213130331233330-1301330333220021-1301023230222300-1231232230012232-1120222121330302-3012103333303002-1201213020122111-0231222113201213"></a>

### Direct properties for `rules.spec.segment_policy.dst_segments.segments`

<a id="canonical-2100110133021310-0112113103132231-3230100322022301-1220233033121233-3330011203233002-0200022130113302-3203021002300333-2120221023013102"></a>

#### `rules.spec.segment_policy.dst_segments.segments.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1220212120313301-2121132310123332-1210222220220121-1001021300001113-3301011003103010-2333112032322231-3031302312110312-0000322001013333"></a>

<a id="canonical-0022020220020010-1212330330133023-1313312013313003-2223212032001132-3002213211222121-0101030333222112-2100303230021001-3311222010310001"></a>

#### `rules.spec.segment_policy.dst_segments.segments.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0203011200013120-1300102301302213-2112030012103230-3301301002003323-1213332021112021-1223013302103312-1002102010010100-1222010203223223"></a>

<a id="canonical-2320023311333113-2332021300312320-1202231311002001-3133201022113021-2312022012301133-2320101321312023-3101020022003131-3022033020221231"></a>

#### `rules.spec.segment_policy.dst_segments.segments.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1031111222233131-3030031320222313-1011000131211231-0212002101013303-2113213210020111-3011323310000203-0133311330023233-2212132203310000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.segment_policy.intra_segment` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-1003300311131110-2122022122210003-0233032231302312-0120003132000133-2322322113133003-2322331131222131-3112023232001233-0030221022001312)
- rules.spec.segment_policy.intra_segment

<a id="canonical-3133233132012333-3321202103131032-0011211323030230-1233213131330023-1000320022030232-3000033103033010-3302220232233133-0331002323210313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for intra segment.

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

Terraform syntax:

```terraform
intra_segment = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320312221031301-0033011221303221-0111313210211221-0212120001322103-1302211211222002-1133133233133021-2310223023131111-3330233113322320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.segment_policy.src_any` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-1003300311131110-2122022122210003-0233032231302312-0120003132000133-2322322113133003-2322331131222131-3112023232001233-0030221022001312)
- rules.spec.segment_policy.src_any

<a id="canonical-3211202133133103-1322321322200311-0000011101131313-0313233122220021-1023003121121000-1133332002200113-0212310032002310-3000022232332101"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
src_any = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022310130211210-2222210220122323-2122321023133332-2031120230030333-1333330100313232-2002132113300110-0200120130020032-3101313230103323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.segment_policy.src_segments` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-1003300311131110-2122022122210003-0233032231302312-0120003132000133-2322322113133003-2322331131222131-3112023232001233-0030221022001312)
- rules.spec.segment_policy.src_segments

<a id="canonical-0001102231230010-0332333302000022-1001010002031030-1132223020002332-3030012223321112-2113322112220010-2010202030022001-0133332132221221"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for src segments.

Additional upstream details:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("segments")}
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
src_segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031013031001021-2320031131201023-3012302010002111-0221031110021210-0323100010002323-3230123200300210-2120030132321313-3101111331201221"></a>

### Direct properties for `rules.spec.segment_policy.src_segments`

- [segments](resources--rate_limiter_policy--reference--group-001.md#canonical-2333313230001310-1130023032101302-1220331102120203-1102212222203210-0022010103203220-0133110103013303-2301310221330000-1201222132333223): complete subsection reference.

<a id="canonical-2333313230001310-1130023032101302-1220331102120203-1102212222203210-0022010103203220-0133110103013303-2301310221330000-1201222132333223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.spec.segment_policy.src_segments.segments` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [rules](resources--rate_limiter_policy--reference--group-001.md#canonical-2222131300013222-2030020131223321-0301112221002023-2103133002001211-0202321013020301-0203133301000033-1033002002011112-3021313102033303)
- [rules.spec](resources--rate_limiter_policy--reference--group-001.md#canonical-0132202013211102-2330210311001320-0320321112002010-3322111012130323-0012003033121113-1211030313313322-2030022002333220-2033032132312031)
- [rules.spec.segment_policy](resources--rate_limiter_policy--reference--group-001.md#canonical-1003300311131110-2122022122210003-0233032231302312-0120003132000133-2322322113133003-2322331131222131-3112023232001233-0030221022001312)
- [rules.spec.segment_policy.src_segments](resources--rate_limiter_policy--reference--group-001.md#canonical-0022310130211210-2222210220122323-2122321023133332-2031120230030333-1333330100313232-2002132113300110-0200120130020032-3101313230103323)
- rules.spec.segment_policy.src_segments.segments

<a id="canonical-3310022203231120-1321101111231030-0333212300133211-1201331010122230-3333231300330010-0201232120210223-0110203011111111-1012111233312322"></a>

Type: `"object"`. list nested block, Optional.

Segments. Select list of segments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

Terraform syntax:

```terraform
segments {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130203003310200-3001230333022113-1321332233311133-0321330020332032-1200010332222132-0303112113231313-1032202001311211-2230010113133300"></a>

### Direct properties for `rules.spec.segment_policy.src_segments.segments`

<a id="canonical-0021221011122032-3223313013001203-3130223220312100-1211322221231023-1021333322123113-2102231223012321-3322201321121103-2102133230320213"></a>

#### `rules.spec.segment_policy.src_segments.segments.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2023101313133101-1323123223222032-1112211321101020-2100130331033021-1112223313300332-1103102000120001-1322332010202101-3021011031200320"></a>

<a id="canonical-2031333230331100-2222001000223331-3001222012332231-3312102232300332-2312023230110330-2102020203201012-0023032003012310-3311032301000200"></a>

#### `rules.spec.segment_policy.src_segments.segments.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0231101030213333-2100213303112230-3310033130132303-0033111131020302-0312121031230010-3012310210231033-3302211012023132-3211001002002230"></a>

<a id="canonical-1300200311032212-3020233002030121-2021133313311230-1212222023330320-0112212322003302-0332103333202020-2122101022213333-0322232320132021"></a>

#### `rules.spec.segment_policy.src_segments.segments.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0332131121330022-2320313020032211-0233232101323322-0302123232122021-0110301202203032-2012210121330301-1232012122013010-0312200313032203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `server_name_matcher` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- server_name_matcher

<a id="canonical-0001101311201130-3312030322211210-3100011300130000-2100302121010320-2012130212301321-0013201312012022-2123122022002031-3033313212122332"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
server_name_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013100131103120-3223010232333203-0010220003012033-1213223001001102-0113212033030213-1111012300023011-1033212230110033-0112222233203223"></a>

### Direct properties for `server_name_matcher`

<a id="canonical-0111011221011303-2223311331322200-2321220132112130-0102300130012112-0221033002121122-1233131002001302-3110013203232231-1112110222310333"></a>

#### `server_name_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3000302322312310-0302310112302221-2333321001033210-2123313001001112-3322310322230222-1113310102231121-2121322101103213-2221033303022100"></a>

<a id="canonical-2311010221101011-2223102232120201-2231300100113223-2030012022302032-1020211031323333-2302123133311210-2031112001231303-3232123321001303"></a>

#### `server_name_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2213123230331302-2112033033112303-3303010013132210-2021022013123003-2001322313033232-2333120000201120-2312313120200303-1032200022123013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `server_selector` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- server_selector

<a id="canonical-3333222120133010-2212303113313110-1101120223022023-3033332033033312-0130133232220202-0321213221012311-3311301211321101-0220012223003333"></a>

Type: `"object"`. single nested block, Optional.

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
server_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312020032223300-0322210000331300-3002233320112003-2012022032120102-0201233103132321-2133120033300203-2301202213320131-0110211002021302"></a>

### Direct properties for `server_selector`

<a id="canonical-1311030133210332-1102012030123233-1130333021103132-3000222032031121-1001302113023311-3212033313212023-0021203031333231-1231110322301130"></a>

#### `server_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1121122303032230-0332030130132021-3012333310113223-0311011022311021-2203133033133302-0012001113101220-2001003203233201-1012131010323033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Property reference](resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- timeouts

<a id="canonical-3212103331233210-2000222302313032-0313013101020010-3110101321112111-3031023103002100-1202332200210132-3212322122100212-1101030313300101"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032212221231011-1012330320021233-1303222023303303-3123110230133201-0321303232100012-1201303031221023-1212323210132110-1322311123002130"></a>

### Direct properties for `timeouts`

<a id="canonical-3032230212001120-1211202232101202-2200012000122202-2100232320220213-1121213132231301-0000111312303030-0203312302222011-3132133312000020"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0030331222233313-3011120131022010-0311013330110023-0200032313012010-0332021213230103-2003322311010331-2033222321030313-1123321031013200"></a>

<a id="canonical-3010013021022200-2110313301023013-2213022230333111-0333000133310210-2312223301131203-3001023020310131-1220100331032203-1121102231302100"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2320100323003002-1212331113103103-2003203303100033-1233023002100233-3311013333300000-1330233120200103-3001311212023222-0213233010130001"></a>

<a id="canonical-2031222321232010-3300310213333232-2031120222131001-3023330101013010-1120120000000212-3212110130121121-3130321311321331-2032132012121220"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2201030213200030-3311310030101113-3011112333130131-3320321220323232-2332332021113303-0332320022330122-3232111220103103-0002303301103122"></a>

<a id="canonical-0120032213320313-2301221323102231-0303023211010011-3303200310213001-2212301013321202-2011222121213331-3123010001223320-1331011100122311"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
