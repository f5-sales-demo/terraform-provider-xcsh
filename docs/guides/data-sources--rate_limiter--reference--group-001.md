---
page_title: "xcsh_rate_limiter reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter reference."
---

# xcsh_rate_limiter reference

<a id="canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- Property reference

<a id="canonical-3103220300230113-1200103313122203-1231001303003111-2323100103120221-1210202312220200-2333101101130223-2313001330123132-3213031130022010"></a>

### Direct properties for `xcsh_rate_limiter`

<a id="canonical-1203302103300313-3122233002112311-2001121310113033-1131321023101033-3120100311112032-3010230012300220-0303302110101112-3321201010320132"></a>

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

<a id="canonical-2020331020220023-0133212313132200-3120201220232032-0211202010320010-1332230002300000-2131103230000201-1020131033123313-2112223112121221"></a>

<a id="canonical-1030301302031033-3332222302233030-2121023201011333-0113103022032122-0130203311210132-3320233331132313-3133213003330202-1231102322321311"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the RateLimiter.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1332022000112100-3120311030130013-2120112031221200-0120313123020033-2220110321321031-2322002011323131-3211321323323322-0231322201133010"></a>

<a id="canonical-3130221213022113-3322302100112301-1232101122310202-3023121313023101-3200201111011230-0111233132000103-1130222210103120-1003031123200010"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0001033012023233-3333332131011110-1111313112330111-2020311333132133-2101020201310221-3101221103212331-3121131222332233-2201112031130221"></a>

<a id="canonical-2321112023032322-2130201200003001-1220003211013321-3130212100321210-2001320132003203-2300133323320032-1013222223221200-2221312012201212"></a>

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

- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-1210221011020222-1230201312322022-0013012231131330-1221200323202101-0311331002301321-2120032131302113-1312313103113202-2030231023103011): complete subsection reference.

<a id="canonical-3300032321030010-0221231322012101-1322100331330330-1313003311301331-2302131230103013-2003232201323103-0111003010002033-2313100233321302"></a>

<a id="canonical-2103320113031312-0221100223223122-3022300331323203-3102210013012012-0121121201221032-3222010113331012-0032331310230310-3002132013331013"></a>

#### `name` property

Type: `"string"`. Required.

Name of the RateLimiter.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0213203300223201-0000322310321323-0100103102312013-2012131213123223-1301113000132302-3211233110200123-0312102023111312-3320111102211021"></a>

<a id="canonical-3232010202313120-3101311220210330-2211230102022122-0012130130220101-3112021001300212-3102202011310311-0313302131102211-1113131101021110"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the RateLimiter exists.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [user_identification](data-sources--rate_limiter--reference--group-001.md#canonical-0232132011012110-3032113110213030-1122101233132033-1212003223102123-1132201133332233-0201331031010211-0100030312303132-1032321033203102): complete subsection reference.

<a id="canonical-1121330003133021-0030320112113033-1303312221031030-3213021310122300-2003230330111330-1022010131011310-3132102120302310-3230102210111332"></a>

### All schema paths for `xcsh_rate_limiter`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--rate_limiter--reference--group-001.md#canonical-1203302103300313-3122233002112311-2001121310113033-1131321023101033-3120100311112032-3010230012300220-0303302110101112-3321201010320132) |
| `description` | [description](data-sources--rate_limiter--reference--group-001.md#canonical-2020331020220023-0133212313132200-3120201220232032-0211202010320010-1332230002300000-2131103230000201-1020131033123313-2112223112121221) |
| `id` | [ID](data-sources--rate_limiter--reference--group-001.md#canonical-1332022000112100-3120311030130013-2120112031221200-0120313123020033-2220110321321031-2322002011323131-3211321323323322-0231322201133010) |
| `labels` | [labels](data-sources--rate_limiter--reference--group-001.md#canonical-0001033012023233-3333332131011110-1111313112330111-2020311333132133-2101020201310221-3101221103212331-3121131222332233-2201112031130221) |
| `limits` | [limits](data-sources--rate_limiter--reference--group-001.md#canonical-2113123322222122-3030132111030323-1221233113001031-2022200331010020-0102011123120021-1132333212313120-1002011231112222-1021132311100303) |
| `limits.action_block` | [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-1232112300321011-1111232010123300-3222133330203301-1203030333032233-0332310010000022-0321301221110321-0203101123212231-1302122030000300) |
| `limits.action_block.hours` | [limits.action_block.hours](data-sources--rate_limiter--reference--group-001.md#canonical-1111122210303303-1231033230313231-1311003333103213-0311022000121112-0330312210211032-2123001113033203-3031231310212332-0331331003310030) |
| `limits.action_block.hours.duration` | [limits.action_block.hours.duration](data-sources--rate_limiter--reference--group-001.md#canonical-3220223133000323-1233210001130322-1222010100203232-0112333333312002-2202213020013101-2113211213103332-2210213202320323-1310121133333200) |
| `limits.action_block.minutes` | [limits.action_block.minutes](data-sources--rate_limiter--reference--group-001.md#canonical-3032021111132122-2230021210033121-1322121312323012-1313010310033033-1131231303022212-1201111333300212-1012301011232100-3112233231233022) |
| `limits.action_block.minutes.duration` | [limits.action_block.minutes.duration](data-sources--rate_limiter--reference--group-001.md#canonical-0102002113303201-3313123120010111-2202112001233330-2232031320023123-2022302300233021-2223020212321032-3311023031202032-2213013231221313) |
| `limits.action_block.seconds` | [limits.action_block.seconds](data-sources--rate_limiter--reference--group-001.md#canonical-0010332110012310-3213212021011002-3212331232222103-3131221132332303-2301223112330100-3220210201112112-3100332000222023-0232212021120002) |
| `limits.action_block.seconds.duration` | [limits.action_block.seconds.duration](data-sources--rate_limiter--reference--group-001.md#canonical-0331023013031110-0201200211131021-0103311010112323-2121302323003313-1123123221313331-1230131332322221-0030002031330031-1212202323113001) |
| `limits.burst_multiplier` | [limits.burst_multiplier](data-sources--rate_limiter--reference--group-001.md#canonical-2313113300033330-1221010121022301-0333312323221103-2120323311221203-3220202100303102-2212023113132032-1221133121221031-2032033331101212) |
| `limits.disabled` | [limits.disabled](data-sources--rate_limiter--reference--group-001.md#canonical-2323332301212212-3212230103230010-2011203111001211-3203210120131120-1102331110200332-3113321000213003-1321203100212330-3331013120110332) |
| `limits.leaky_bucket` | [limits.leaky_bucket](data-sources--rate_limiter--reference--group-001.md#canonical-3212131130020213-0332130112031013-2310111123221010-3133021302302200-0003331101122210-3302100001202202-2021033322110313-3131012013230313) |
| `limits.period_multiplier` | [limits.period_multiplier](data-sources--rate_limiter--reference--group-001.md#canonical-1322123020002230-3301200302113232-0230300212120101-0013031003231233-2011320103113111-3300122112203201-2220321000021222-3033232001132220) |
| `limits.token_bucket` | [limits.token_bucket](data-sources--rate_limiter--reference--group-001.md#canonical-2132313003123103-3203012330020201-3113331212113003-3000002031100111-1303332310102321-3200022130022001-1001120220322212-0022223301300300) |
| `limits.total_number` | [limits.total_number](data-sources--rate_limiter--reference--group-001.md#canonical-2312121001000013-3320113101111100-0220313131320320-3311121030132321-3301322033220320-2130211133003031-2311023332303110-2330030331033003) |
| `limits.unit` | [limits.unit](data-sources--rate_limiter--reference--group-001.md#canonical-2010010103310001-1132000133102320-2121033222301300-1113330130120010-2033031222301213-0323101322013223-3030022231031333-2313102211322332) |
| `name` | [name](data-sources--rate_limiter--reference--group-001.md#canonical-3300032321030010-0221231322012101-1322100331330330-1313003311301331-2302131230103013-2003232201323103-0111003010002033-2313100233321302) |
| `namespace` | [namespace](data-sources--rate_limiter--reference--group-001.md#canonical-0213203300223201-0000322310321323-0100103102312013-2012131213123223-1301113000132302-3211233110200123-0312102023111312-3320111102211021) |
| `user_identification` | [user_identification](data-sources--rate_limiter--reference--group-001.md#canonical-2233301200303121-0033313003323033-1011132121010113-1321200121223221-3130320021200313-1321120133331230-0013230223310220-0320020230120203) |
| `user_identification.kind` | [user_identification.kind](data-sources--rate_limiter--reference--group-001.md#canonical-1121132232021110-1033330012131033-1020023033332103-1210233122101310-2233212203302133-0333300300311332-2312120220130112-1222313122003111) |
| `user_identification.name` | [user_identification.name](data-sources--rate_limiter--reference--group-001.md#canonical-3232232032120210-0013230032030012-3233320312102233-0110012013013303-0000132012220011-1221111021330032-0310230200032303-2131223223223212) |
| `user_identification.namespace` | [user_identification.namespace](data-sources--rate_limiter--reference--group-001.md#canonical-2313022001221132-0321200120132021-2003131233000303-3223232332033021-0201233313030120-0131132122230313-3311333022112300-3222132213010123) |
| `user_identification.tenant` | [user_identification.tenant](data-sources--rate_limiter--reference--group-001.md#canonical-0123200202100112-3113232213333312-3232320332113220-2213310032210110-0230310010101333-1313121000030030-1330022330022100-1021313313022021) |
| `user_identification.uid` | [user_identification.uid](data-sources--rate_limiter--reference--group-001.md#canonical-2331323231331000-1332231001101133-1132212111131310-0301110111133131-0213231021111130-1220011121322322-3301001103312021-2121130001013223) |

<a id="canonical-1210221011020222-1230201312322022-0013012231131330-1221200323202101-0311331002301321-2120032131302113-1312313103113202-2030231023103011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits` properties

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223)
- limits

<a id="canonical-2113123322222122-3030132111030323-1221233113001031-2022200331010020-0102011123120021-1132333212313120-1002011231112222-1021132311100303"></a>

Type: `"list"`. Computed.

A list of RateLimitValues that specifies the total number of allowed requests for each specified
period.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1232102123112331-3201200100111213-3002331103011210-3121012113021001-3100010112123210-1010310020123210-2001012332231222-2120202110312333"></a>

### Direct properties for `limits`

- [action_block](data-sources--rate_limiter--reference--group-001.md#canonical-3221130221020103-0312010202000323-2021222110132133-0013331101311010-2012013120020130-0323321210003020-2032322033000210-3211121201221012): complete subsection reference.

<a id="canonical-2313113300033330-1221010121022301-0333312323221103-2120323311221203-3220202100303102-2212023113132032-1221133121221031-2032033331101212"></a>

<a id="canonical-1131101000010312-1223302030310122-3303220022133021-3211120223311330-0032033022022232-3321303031211120-3303131131312231-1221013220330033"></a>

#### `limits.burst_multiplier` property

Type: `"number"`. Computed.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](data-sources--rate_limiter--reference--group-001.md#canonical-1223022203202202-1230121110333132-3002133003103103-0213121333323003-2010211221300101-1122303302121023-3311022223130030-3010303331223310): complete subsection reference.

- [leaky_bucket](data-sources--rate_limiter--reference--group-001.md#canonical-0112310023313020-1232320113031100-2013312330221001-1033201111103222-0313311032101323-2030302122211033-3133220213122100-3313230032200103): complete subsection reference.

<a id="canonical-1322123020002230-3301200302113232-0230300212120101-0013031003231233-2011320103113111-3300122112203201-2220321000021222-3033232001132220"></a>

<a id="canonical-3103332302323200-0000033211202232-3012301030200132-1221132201110220-3323030003030111-1001021000301113-1121033130302113-2333032123100122"></a>

#### `limits.period_multiplier` property

Type: `"number"`. Computed.

This setting, combined with Per Period units, provides a duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

- [token_bucket](data-sources--rate_limiter--reference--group-001.md#canonical-2212113022203132-2321331333203133-3102222103200021-0310003111032310-3320200113203333-1303022001101013-2122211002201100-3302033133201321): complete subsection reference.

<a id="canonical-2312121001000013-3320113101111100-0220313131320320-3311121030132321-3301322033220320-2130211133003031-2311023332303110-2330030331033003"></a>

<a id="canonical-0002220111322102-3010131012230031-1300013132200321-2321302031101320-2223122201022313-2131312223032331-0210223231021331-0011003023231313"></a>

#### `limits.total_number` property

Type: `"number"`. Computed.

The total number of allowed requests per rate-limiting period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-2010010103310001-1132000133102320-2121033222301300-1113330130120010-2033031222301213-0323101322013223-3030022231031333-2313102211322332"></a>

<a id="canonical-1233020321333302-3110231332201321-1021000313022013-2203110333121233-1221233202112030-1232302232023232-3112023132233302-1221320100211222"></a>

#### `limits.unit` property

Type: `"string"`. Computed.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3221130221020103-0312010202000323-2021222110132133-0013331101311010-2012013120020130-0323321210003020-2032322033000210-3211121201221012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.action_block` properties

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-1210221011020222-1230201312322022-0013012231131330-1221200323202101-0311331002301321-2120032131302113-1312313103113202-2030231023103011)
- limits.action_block

<a id="canonical-1232112300321011-1111232010123300-3222133330203301-1203030333032233-0332310010000022-0321301221110321-0203101123212231-1302122030000300"></a>

Type: `"single"`. Computed.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

<a id="canonical-2303313232123223-2113023303223200-0123122001212223-0031001101220330-0302311312010103-3033030111113210-1220323213330300-1103322031320123"></a>

### Direct properties for `limits.action_block`

- [hours](data-sources--rate_limiter--reference--group-001.md#canonical-1132312121130203-3220112001102221-3003211201330132-1031212331101010-1133321320123003-2110103322221010-2110033021202011-1133203221101133): complete subsection reference.

- [minutes](data-sources--rate_limiter--reference--group-001.md#canonical-1021302332110220-0031003201110210-1223110320122223-1021123123222303-3011130330232230-3301131123231033-3013003000121130-3332130300001322): complete subsection reference.

- [seconds](data-sources--rate_limiter--reference--group-001.md#canonical-0122001301103022-3321210303100100-2323003322132223-1010113331111022-2023132130133030-2100130133013021-2023020230110133-3132112030012122): complete subsection reference.

<a id="canonical-1132312121130203-3220112001102221-3003211201330132-1031212331101010-1133321320123003-2110103322221010-2110033021202011-1133203221101133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.action_block.hours` properties

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-1210221011020222-1230201312322022-0013012231131330-1221200323202101-0311331002301321-2120032131302113-1312313103113202-2030231023103011)
- [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-3221130221020103-0312010202000323-2021222110132133-0013331101311010-2012013120020130-0323321210003020-2032322033000210-3211121201221012)
- limits.action_block.hours

<a id="canonical-1111122210303303-1231033230313231-1311003333103213-0311022000121112-0330312210211032-2123001113033203-3031231310212332-0331331003310030"></a>

Type: `"single"`. Computed.

Hours. Input Duration Hours.

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

<a id="canonical-2231103312300121-0301112301300203-2021322301020300-3013100212110212-3121121302132010-3310311230102121-1102332202021233-0003021013223212"></a>

### Direct properties for `limits.action_block.hours`

<a id="canonical-3220223133000323-1233210001130322-1222010100203232-0112333333312002-2202213020013101-2113211213103332-2210213202320323-1310121133333200"></a>

#### `limits.action_block.hours.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-1021302332110220-0031003201110210-1223110320122223-1021123123222303-3011130330232230-3301131123231033-3013003000121130-3332130300001322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.action_block.minutes` properties

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-1210221011020222-1230201312322022-0013012231131330-1221200323202101-0311331002301321-2120032131302113-1312313103113202-2030231023103011)
- [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-3221130221020103-0312010202000323-2021222110132133-0013331101311010-2012013120020130-0323321210003020-2032322033000210-3211121201221012)
- limits.action_block.minutes

<a id="canonical-3032021111132122-2230021210033121-1322121312323012-1313010310033033-1131231303022212-1201111333300212-1012301011232100-3112233231233022"></a>

Type: `"single"`. Computed.

Minutes. Input Duration Minutes.

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

<a id="canonical-0213001032001122-1113131102133021-2132013030123003-1202110130213210-2331333113211033-0333120102311020-0313133201132021-1320220100132013"></a>

### Direct properties for `limits.action_block.minutes`

<a id="canonical-0102002113303201-3313123120010111-2202112001233330-2232031320023123-2022302300233021-2223020212321032-3311023031202032-2213013231221313"></a>

#### `limits.action_block.minutes.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-0122001301103022-3321210303100100-2323003322132223-1010113331111022-2023132130133030-2100130133013021-2023020230110133-3132112030012122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.action_block.seconds` properties

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-1210221011020222-1230201312322022-0013012231131330-1221200323202101-0311331002301321-2120032131302113-1312313103113202-2030231023103011)
- [limits.action_block](data-sources--rate_limiter--reference--group-001.md#canonical-3221130221020103-0312010202000323-2021222110132133-0013331101311010-2012013120020130-0323321210003020-2032322033000210-3211121201221012)
- limits.action_block.seconds

<a id="canonical-0010332110012310-3213212021011002-3212331232222103-3131221132332303-2301223112330100-3220210201112112-3100332000222023-0232212021120002"></a>

Type: `"single"`. Computed.

Seconds. Input Duration Seconds.

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

<a id="canonical-3023310121120003-0213113231001130-0202120001131112-1310102101012013-1323032321032112-0133002123131203-3232222213332330-3303202130012233"></a>

### Direct properties for `limits.action_block.seconds`

<a id="canonical-0331023013031110-0201200211131021-0103311010112323-2121302323003313-1123123221313331-1230131332322221-0030002031330031-1212202323113001"></a>

#### `limits.action_block.seconds.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-1223022203202202-1230121110333132-3002133003103103-0213121333323003-2010211221300101-1122303302121023-3311022223130030-3010303331223310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.disabled` properties

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-1210221011020222-1230201312322022-0013012231131330-1221200323202101-0311331002301321-2120032131302113-1312313103113202-2030231023103011)
- limits.disabled

<a id="canonical-2323332301212212-3212230103230010-2011203111001211-3203210120131120-1102331110200332-3113321000213003-1321203100212330-3331013120110332"></a>

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

<a id="canonical-0112310023313020-1232320113031100-2013312330221001-1033201111103222-0313311032101323-2030302122211033-3133220213122100-3313230032200103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.leaky_bucket` properties

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-1210221011020222-1230201312322022-0013012231131330-1221200323202101-0311331002301321-2120032131302113-1312313103113202-2030231023103011)
- limits.leaky_bucket

<a id="canonical-3212131130020213-0332130112031013-2310111123221010-3133021302302200-0003331101122210-3302100001202202-2021033322110313-3131012013230313"></a>

Type: `["object", {}]`. Computed.

Leaky-Bucket is the default rate limiter algorithm for F5.

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

<a id="canonical-2212113022203132-2321331333203133-3102222103200021-0310003111032310-3320200113203333-1303022001101013-2122211002201100-3302033133201321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.token_bucket` properties

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223)
- [limits](data-sources--rate_limiter--reference--group-001.md#canonical-1210221011020222-1230201312322022-0013012231131330-1221200323202101-0311331002301321-2120032131302113-1312313103113202-2030231023103011)
- limits.token_bucket

<a id="canonical-2132313003123103-3203012330020201-3113331212113003-3000002031100111-1303332310102321-3200022130022001-1001120220322212-0022223301300300"></a>

Type: `["object", {}]`. Computed.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

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

<a id="canonical-0232132011012110-3032113110213030-1122101233132033-1212003223102123-1132201133332233-0201331031010211-0100030312303132-1032321033203102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_identification` properties

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- [Property reference](data-sources--rate_limiter--reference--group-001.md#canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223)
- user_identification

<a id="canonical-2233301200303121-0033313003323033-1011132121010113-1321200121223221-3130320021200313-1321120133331230-0013230223310220-0320020230120203"></a>

Type: `"list"`. Computed.

Reference to user\_identification object. The rules in the user\_identification object are evaluated
to determine the user identifier to be rate limited. Defaults to \`\[\]\`. Server applies default
when omitted.

Additional upstream details:

A reference to user\_identification object.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0213203220202000-3131012232300031-0321303200111303-0311332030000300-3113301302203133-1032122101221021-0333102000203021-1032231130223120"></a>

### Direct properties for `user_identification`

<a id="canonical-1121132232021110-1033330012131033-1020023033332103-1210233122101310-2233212203302133-0333300300311332-2312120220130112-1222313122003111"></a>

#### `user_identification.kind` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3232232032120210-0013230032030012-3233320312102233-0110012013013303-0000132012220011-1221111021330032-0310230200032303-2131223223223212"></a>

<a id="canonical-2122333020003102-2223112233230312-2113221200303312-1222121011201223-0312221000100332-1133111312123221-0101231310101210-1121333323230031"></a>

#### `user_identification.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2313022001221132-0321200120132021-2003131233000303-3223232332033021-0201233313030120-0131132122230313-3311333022112300-3222132213010123"></a>

<a id="canonical-3112103101101301-1102103001331113-2011323231303031-3233102113310301-3220303030110310-0200033033212210-0112002100233113-1022222212002111"></a>

#### `user_identification.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0123200202100112-3113232213333312-3232320332113220-2213310032210110-0230310010101333-1313121000030030-1330022330022100-1021313313022021"></a>

<a id="canonical-2022331220312110-0133130122322011-1333030321021030-1022120330010331-0322030032312112-3000320231021231-0213113323103310-3300330331111220"></a>

#### `user_identification.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2331323231331000-1332231001101133-1132212111131310-0301110111133131-0213231021111130-1220011121322322-3301001103312021-2121130001013223"></a>

<a id="canonical-0101113100111121-1113232132313121-3313010231331102-0332223333220122-2221102000333032-2222003221201223-2313011310013211-3303112201010032"></a>

#### `user_identification.uid` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
