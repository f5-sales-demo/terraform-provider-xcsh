---
page_title: "xcsh_rate_limiter reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter reference."
---

# xcsh_rate_limiter reference

<a id="canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- Property reference

<a id="canonical-1011331010220203-3210220030320221-1202212301021000-0021001010111302-3100300223003323-1201102120313330-3220312301100123-0322122112030033"></a>

### Direct properties for `xcsh_rate_limiter`

<a id="canonical-2003010330033033-3132203300301323-2020012311211130-0011030032022032-1113322002303302-2023312223011203-2010213101313021-0030032322332003"></a>

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

<a id="canonical-2122203200003032-0220201222232332-2303232013330222-0123033010311000-0010023220223121-2312120121003011-0030220113223011-3332200213333110"></a>

<a id="canonical-1320301322022122-2131310210333033-2030203021320003-2212203001102332-2130302332210122-3112313222023313-2230330300123321-1130132212130312"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0132013330033110-2020122211232033-1101100001210012-0310031120002133-2022311122131101-0100311121220301-3232310201210312-3321300211301020"></a>

<a id="canonical-0010222101331233-3200222012001023-2120022001330123-2003032202322212-3302131303032200-3032103020012111-3320211132031322-1202030203133312"></a>

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

<a id="canonical-3133102211303121-2122211223323312-0232002202103111-0213022213323022-1300033030221231-2313231223222333-1203223221121321-3320310321032103"></a>

<a id="canonical-2003302210003230-2010321213010030-1100100220112100-0022122123133202-3113100200023210-1321313221120212-3330300230013012-2321222211130230"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2202032321033221-0332301330010013-1120002003120210-2130303222021111-1020112122123102-1302230322320022-3103203331121332-0010101333113221"></a>

<a id="canonical-2200223110032212-0132012002301033-2032232131312013-1200210011130103-1032323210220100-2120330101131102-1311031301113311-3010323003232320"></a>

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

- [limits](resources--rate_limiter--reference--group-001.md#canonical-3310131000101102-2223023103233001-2320233313020222-3232300220331020-0300101132023133-1111032010320111-3101031132110011-3311313100200222): complete subsection reference.

<a id="canonical-0310201323223200-2312020130103130-0330132112300013-0133320311022303-3101210233200310-2321030030121031-2101000033331102-1232110220302302"></a>

<a id="canonical-1311030230231323-0322032131231011-1131023011303130-0001303011302220-3322131312021301-1023021203301313-1202202120223312-0123033222203113"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Rate Limiter. Must be unique within the namespace.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0023323203010113-2313311223210103-0321120230303313-1001111122132033-0213223213110321-2200221002000032-3203133300002233-1113320200300022"></a>

<a id="canonical-2122203221331112-2231023333023312-2110011002233022-0301113210223233-0233313132130110-0023331131310103-2221220032333231-2303223100302230"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Rate Limiter is created.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [timeouts](resources--rate_limiter--reference--group-001.md#canonical-1301331213220103-0121001211213312-1211021321130203-0131002331132330-3130102221110321-2330223303001322-2211131311120322-0120301200121313): complete subsection reference.

- [user_identification](resources--rate_limiter--reference--group-001.md#canonical-3013012011131013-2130113331323002-1022320231321331-2301301331133002-3322313112300321-1232322020332131-2203031111131013-0132220031100223): complete subsection reference.

<a id="canonical-1100210010011201-0133202302022323-0122320213012011-1113301023303321-0002323222131020-3231322120110113-1233001301203223-0321311330231100"></a>

### All schema paths for `xcsh_rate_limiter`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--rate_limiter--reference--group-001.md#canonical-2003010330033033-3132203300301323-2020012311211130-0011030032022032-1113322002303302-2023312223011203-2010213101313021-0030032322332003) |
| `description` | [description](resources--rate_limiter--reference--group-001.md#canonical-2122203200003032-0220201222232332-2303232013330222-0123033010311000-0010023220223121-2312120121003011-0030220113223011-3332200213333110) |
| `disable` | [disable](resources--rate_limiter--reference--group-001.md#canonical-0132013330033110-2020122211232033-1101100001210012-0310031120002133-2022311122131101-0100311121220301-3232310201210312-3321300211301020) |
| `id` | [ID](resources--rate_limiter--reference--group-001.md#canonical-3133102211303121-2122211223323312-0232002202103111-0213022213323022-1300033030221231-2313231223222333-1203223221121321-3320310321032103) |
| `labels` | [labels](resources--rate_limiter--reference--group-001.md#canonical-2202032321033221-0332301330010013-1120002003120210-2130303222021111-1020112122123102-1302230322320022-3103203331121332-0010101333113221) |
| `limits` | [limits](resources--rate_limiter--reference--group-001.md#canonical-2313333233020312-0200121021121220-0220201130302010-2000121333002011-0221212233330333-2020030010032211-3012010312033311-3033311220310123) |
| `limits.action_block` | [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-0001000331023210-3211233312131301-3321312031000020-2333323331223122-0220200102023032-3000031230112012-3221130311312032-1012113132112130) |
| `limits.action_block.hours` | [limits.action_block.hours](resources--rate_limiter--reference--group-001.md#canonical-1030322102213113-2022332012112113-0133110200300121-3233322133203103-2011322002201101-1023310023202300-3120023212302002-1021012301103322) |
| `limits.action_block.hours.duration` | [limits.action_block.hours.duration](resources--rate_limiter--reference--group-001.md#canonical-1010322313212313-2021203202003101-2103001021000101-3201323303102021-1123322100201203-3030010110013220-0033123132002030-2201330311011220) |
| `limits.action_block.minutes` | [limits.action_block.minutes](resources--rate_limiter--reference--group-001.md#canonical-0102301133222213-2101232231010223-2302123230203221-1111310321010110-3100300120211300-0132123201312230-1320212320210123-0121300330212222) |
| `limits.action_block.minutes.duration` | [limits.action_block.minutes.duration](resources--rate_limiter--reference--group-001.md#canonical-0122200312223320-0332222331321011-1312132303113123-0222332101230322-2121101212302013-1112221012330021-0101033323303102-0012130202300032) |
| `limits.action_block.seconds` | [limits.action_block.seconds](resources--rate_limiter--reference--group-001.md#canonical-0213110132203321-3101032021131300-2020211331321113-3013302211320011-1323120211113200-2011133013000000-1230131112200120-3303301223311110) |
| `limits.action_block.seconds.duration` | [limits.action_block.seconds.duration](resources--rate_limiter--reference--group-001.md#canonical-0221313001121121-2013320003201213-3123311102202000-2001303013311213-3010030020332120-2112132131011033-2312020113320012-3230013123200203) |
| `limits.burst_multiplier` | [limits.burst_multiplier](resources--rate_limiter--reference--group-001.md#canonical-2120330323302231-1110220330132222-3101321123203032-0313321301201330-3302311312031230-2231230130123100-3013300223301133-2022221100203300) |
| `limits.disabled` | [limits.disabled](resources--rate_limiter--reference--group-001.md#canonical-1223321211233201-0222023232100210-3313323012101012-3313112003132030-3232222030333102-3011120220112330-0333212031022321-0113311233111221) |
| `limits.leaky_bucket` | [limits.leaky_bucket](resources--rate_limiter--reference--group-001.md#canonical-0110330011221322-2202020332222103-0211133032331301-0303113020330013-0222202220100103-2122320230102003-0211032320022002-0232233133103012) |
| `limits.period_multiplier` | [limits.period_multiplier](resources--rate_limiter--reference--group-001.md#canonical-1130230313033233-0223132032102301-1303122132023300-3201333020130212-2100232313312010-2220012232002130-1013120231000102-2030211100110322) |
| `limits.token_bucket` | [limits.token_bucket](resources--rate_limiter--reference--group-001.md#canonical-1103001113333302-0121332230211300-3331311130302012-1021032123131103-1321333330111122-2333221131020302-1201202203133111-0222112222202311) |
| `limits.total_number` | [limits.total_number](resources--rate_limiter--reference--group-001.md#canonical-3303231222022300-2110123223222300-2133130331320012-2023021201013022-3130120223012213-2313011201313303-2001030100110131-3223232310203023) |
| `limits.unit` | [limits.unit](resources--rate_limiter--reference--group-001.md#canonical-0122132232203311-2210313323003210-0001231320330010-3203021221030131-1211332000320030-0222020200222223-0133103201323230-3021223322323301) |
| `name` | [name](resources--rate_limiter--reference--group-001.md#canonical-0310201323223200-2312020130103130-0330132112300013-0133320311022303-3101210233200310-2321030030121031-2101000033331102-1232110220302302) |
| `namespace` | [namespace](resources--rate_limiter--reference--group-001.md#canonical-0023323203010113-2313311223210103-0321120230303313-1001111122132033-0213223213110321-2200221002000032-3203133300002233-1113320200300022) |
| `timeouts` | [timeouts](resources--rate_limiter--reference--group-001.md#canonical-2021211103122023-0102222002121031-1131130323032111-3332320333330102-0002331233322000-0100302323310011-1131330311113331-0311320030131233) |
| `timeouts.create` | [timeouts.create](resources--rate_limiter--reference--group-001.md#canonical-0303001310012132-0230031032322131-1331103323310010-3112213320312012-2230111003203111-1033312221200200-0023200030121131-1222300001032300) |
| `timeouts.delete` | [timeouts.delete](resources--rate_limiter--reference--group-001.md#canonical-3301032122121002-2100230120003323-3233102320222202-1200212101203203-2222323020113130-0312212112022102-2023302111310321-0303212010023221) |
| `timeouts.read` | [timeouts.read](resources--rate_limiter--reference--group-001.md#canonical-1020101311020230-1322300300320221-3120012223032110-2012100130121112-1303222133201312-3111133301313011-0300232311302123-2321120331000012) |
| `timeouts.update` | [timeouts.update](resources--rate_limiter--reference--group-001.md#canonical-2020012310320110-0223202133010212-2011121022100212-0113130100001310-1221113310130222-2312033130013211-0102321231311333-3221220011221022) |
| `user_identification` | [user_identification](resources--rate_limiter--reference--group-001.md#canonical-3312003012232321-0022113332020233-2232101312320321-3010022233330103-2011031031301023-2223023002130311-2121232121330020-1011313320110022) |
| `user_identification.kind` | [user_identification.kind](resources--rate_limiter--reference--group-001.md#canonical-0000112211113112-3120300203101213-3223212003312011-3200230030110221-3311133132131111-1012113123303302-3011033113211323-2033013031132221) |
| `user_identification.name` | [user_identification.name](resources--rate_limiter--reference--group-001.md#canonical-0202031022120002-1131112120330322-2210202101003320-2102030203121323-3023322110302013-2322110101331200-1031310221333233-1300230211230123) |
| `user_identification.namespace` | [user_identification.namespace](resources--rate_limiter--reference--group-001.md#canonical-3300312220322302-0103231130030300-2003210030122100-1110122103222313-3331321102023033-0301020121331233-1102210120133132-2310111112201332) |
| `user_identification.tenant` | [user_identification.tenant](resources--rate_limiter--reference--group-001.md#canonical-3313210233321332-2312032211100321-3322113232230323-3003200321112133-3220231200032101-1221033001300302-1002033103012103-0233310311311310) |
| `user_identification.uid` | [user_identification.uid](resources--rate_limiter--reference--group-001.md#canonical-1201001003121013-3101331213200313-2313201020321231-0120310013222333-1121002120200211-1010013233310023-3133100302202113-0023301222220030) |

<a id="canonical-3310131000101102-2223023103233001-2320233313020222-3232300220331020-0300101132023133-1111032010320111-3101031132110011-3311313100200222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits` properties

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022)
- limits

<a id="canonical-2313333233020312-0200121021121220-0220201130302010-2000121333002011-0221212233330333-2020030010032211-3012010312033311-3033311220310123"></a>

Type: `"object"`. list nested block, Optional.

A list of RateLimitValues that specifies the total number of allowed requests for each specified
period.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("total_number"),
  validators.ConflictingListObjectAttributes("action_block",
    "disabled"),
  validators.ConflictingListObjectAttributes("leaky_bucket",
    "token_bucket")}
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
limits {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020030131013303-3200123032013300-3123200213111310-1313222130130132-0003213110020031-1101101320022003-1033122332133331-3102320001232200"></a>

### Direct properties for `limits`

- [action_block](resources--rate_limiter--reference--group-001.md#canonical-3330231013002132-1002333201033023-3203301313302131-2201003132020220-0001323103320111-3330113203231300-1211302013131322-2012323210131033): complete subsection reference.

<a id="canonical-2120330323302231-1110220330132222-3101321123203032-0313321301201330-3302311312031230-2231230130123100-3013300223301133-2022221100203300"></a>

<a id="canonical-0311213311300113-0012011210231211-3330312130023212-3003223311300302-1133312221233121-1231121201301011-2211010200100112-1320101200233230"></a>

#### `limits.burst_multiplier` property

Type: `"number"`. Optional.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [disabled](resources--rate_limiter--reference--group-001.md#canonical-2100333132122033-3311323310131120-0232112103120101-0232121211203013-2132200113113312-1323303211112203-3100011121020032-2231102302303300): complete subsection reference.

- [leaky_bucket](resources--rate_limiter--reference--group-001.md#canonical-0033210111313212-1023232320021302-3210302120301120-0233313213120000-1310003221131223-0201333030032230-3312022033230322-1002000100302233): complete subsection reference.

<a id="canonical-1130230313033233-0223132032102301-1303122132023300-3201333020130212-2100232313312010-2220012232002130-1013120231000102-2030211100110322"></a>

<a id="canonical-1222022012131033-1131333022202131-3022000231001211-2002031221123203-3332113103131020-3313211200123223-0012233020212130-0122212021110002"></a>

#### `limits.period_multiplier` property

Type: `"number"`. Optional.

This setting, combined with Per Period units, provides a duration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [token_bucket](resources--rate_limiter--reference--group-001.md#canonical-1302232130321021-2133323322200100-3130021303303201-1032330201320202-3312123312011213-2223302122030112-0223210212111123-3011213120132033): complete subsection reference.

<a id="canonical-3303231222022300-2110123223222300-2133130331320012-2023021201013022-3130120223012213-2313011201313303-2001030100110131-3223232310203023"></a>

<a id="canonical-3320313131010220-3201023100313213-0200131010000031-1203231110012313-1321021223121033-0310200331233102-1112032031201102-0131210012330011"></a>

#### `limits.total_number` property

Type: `"number"`. Optional.

The total number of allowed requests per rate-limiting period.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0122132232203311-2210313323003210-0001231320330010-3203021221030131-1211332000320030-0222020200222223-0133103201323230-3021223322323301"></a>

<a id="canonical-3303131310023001-0112202230330213-2133120031331310-2223333002012013-2232002202003102-3230210120300213-1001101213231110-0030312012000231"></a>

#### `limits.unit` property

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["HOUR","MINUTE","SECOND"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

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

<a id="canonical-3330231013002132-1002333201033023-3203301313302131-2201003132020220-0001323103320111-3330113203231300-1211302013131322-2012323210131033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.action_block` properties

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-3310131000101102-2223023103233001-2320233313020222-3232300220331020-0300101132023133-1111032010320111-3101031132110011-3311313100200222)
- limits.action_block

<a id="canonical-0001000331023210-3211233312131301-3321312031000020-2333323331223122-0220200102023032-3000031230112012-3221130311312032-1012113132112130"></a>

Type: `"object"`. single nested block, Optional.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes"),
  validators.ConflictingObjectAttributes("hours",
    "seconds"),
  validators.ConflictingObjectAttributes("minutes",
    "seconds")}
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
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

Terraform syntax:

```terraform
action_block {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320112112030103-2130100112103132-0221200003203222-2230223330111300-2001331212011010-1121212032323202-3030032133223332-1311002231012330"></a>

### Direct properties for `limits.action_block`

- [hours](resources--rate_limiter--reference--group-001.md#canonical-2313020202132300-2301332222033330-0232002033301121-1231203021100103-0223232003310333-3013321313011131-0323210021133130-0112312002301101): complete subsection reference.

- [minutes](resources--rate_limiter--reference--group-001.md#canonical-0012101312323230-3233310323121010-1103321221230031-3320231101133311-2120002131223101-3230033100333013-3022112012123113-2103231111332100): complete subsection reference.

- [seconds](resources--rate_limiter--reference--group-001.md#canonical-0011212230131232-3023123121100112-1012333313030131-3113310220030101-1300110330113001-0103310120122031-0123112032301031-1300300231301132): complete subsection reference.

<a id="canonical-2313020202132300-2301332222033330-0232002033301121-1231203021100103-0223232003310333-3013321313011131-0323210021133130-0112312002301101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.action_block.hours` properties

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-3310131000101102-2223023103233001-2320233313020222-3232300220331020-0300101132023133-1111032010320111-3101031132110011-3311313100200222)
- [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-3330231013002132-1002333201033023-3203301313302131-2201003132020220-0001323103320111-3330113203231300-1211302013131322-2012323210131033)
- limits.action_block.hours

<a id="canonical-1030322102213113-2022332012112113-0133110200300121-3233322133203103-2011322002201101-1023310023202300-3120023212302002-1021012301103322"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233222312232112-1211132212200311-2303303303310103-2221000111001300-2223122200032332-2321130113230212-2233001221302102-2032023302210230"></a>

### Direct properties for `limits.action_block.hours`

<a id="canonical-1010322313212313-2021203202003101-2103001021000101-3201323303102021-1123322100201203-3030010110013220-0033123132002030-2201330311011220"></a>

#### `limits.action_block.hours.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 48),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0012101312323230-3233310323121010-1103321221230031-3320231101133311-2120002131223101-3230033100333013-3022112012123113-2103231111332100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.action_block.minutes` properties

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-3310131000101102-2223023103233001-2320233313020222-3232300220331020-0300101132023133-1111032010320111-3101031132110011-3311313100200222)
- [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-3330231013002132-1002333201033023-3203301313302131-2201003132020220-0001323103320111-3330113203231300-1211302013131322-2012323210131033)
- limits.action_block.minutes

<a id="canonical-0102301133222213-2101232231010223-2302123230203221-1111310321010110-3100300120211300-0132123201312230-1320212320210123-0121300330212222"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311201030300232-3120213321323031-3310013103203223-1103212333121001-2010202120332122-3003332032201300-0123221312300223-0111003031201003"></a>

### Direct properties for `limits.action_block.minutes`

<a id="canonical-0122200312223320-0332222331321011-1312132303113123-0222332101230322-2121101212302013-1112221012330021-0101033323303102-0012130202300032"></a>

#### `limits.action_block.minutes.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 60),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0011212230131232-3023123121100112-1012333313030131-3113310220030101-1300110330113001-0103310120122031-0123112032301031-1300300231301132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.action_block.seconds` properties

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-3310131000101102-2223023103233001-2320233313020222-3232300220331020-0300101132023133-1111032010320111-3101031132110011-3311313100200222)
- [limits.action_block](resources--rate_limiter--reference--group-001.md#canonical-3330231013002132-1002333201033023-3203301313302131-2201003132020220-0001323103320111-3330113203231300-1211302013131322-2012323210131033)
- limits.action_block.seconds

<a id="canonical-0213110132203321-3101032021131300-2020211331321113-3013302211320011-1323120211113200-2011133013000000-1230131112200120-3303301223311110"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
seconds {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222123022101112-3021000310001201-3032121110232301-0300120332310030-3312013103301020-3020131130221330-0201003131313221-3122003020233323"></a>

### Direct properties for `limits.action_block.seconds`

<a id="canonical-0221313001121121-2013320003201213-3123311102202000-2001303013311213-3010030020332120-2112132131011033-2312020113320012-3230013123200203"></a>

#### `limits.action_block.seconds.duration` property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 300),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2100333132122033-3311323310131120-0232112103120101-0232121211203013-2132200113113312-1323303211112203-3100011121020032-2231102302303300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.disabled` properties

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-3310131000101102-2223023103233001-2320233313020222-3232300220331020-0300101132023133-1111032010320111-3101031132110011-3311313100200222)
- limits.disabled

<a id="canonical-1223321211233201-0222023232100210-3313323012101012-3313112003132030-3232222030333102-3011120220112330-0333212031022321-0113311233111221"></a>

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
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033210111313212-1023232320021302-3210302120301120-0233313213120000-1310003221131223-0201333030032230-3312022033230322-1002000100302233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.leaky_bucket` properties

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-3310131000101102-2223023103233001-2320233313020222-3232300220331020-0300101132023133-1111032010320111-3101031132110011-3311313100200222)
- limits.leaky_bucket

<a id="canonical-0110330011221322-2202020332222103-0211133032331301-0303113020330013-0222202220100103-2122320230102003-0211032320022002-0232233133103012"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
leaky_bucket = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302232130321021-2133323322200100-3130021303303201-1032330201320202-3312123312011213-2223302122030112-0223210212111123-3011213120132033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `limits.token_bucket` properties

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022)
- [limits](resources--rate_limiter--reference--group-001.md#canonical-3310131000101102-2223023103233001-2320233313020222-3232300220331020-0300101132023133-1111032010320111-3101031132110011-3311313100200222)
- limits.token_bucket

<a id="canonical-1103001113333302-0121332230211300-3331311130302012-1021032123131103-1321333330111122-2333221131020302-1201202203133111-0222112222202311"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
token_bucket = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301331213220103-0121001211213312-1211021321130203-0131002331132330-3130102221110321-2330223303001322-2211131311120322-0120301200121313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022)
- timeouts

<a id="canonical-2021211103122023-0102222002121031-1131130323032111-3332320333330102-0002331233322000-0100302323310011-1131330311113331-0311320030131233"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012010231332302-2132200321102211-0210121100213311-1121320123003101-2102120113021211-3110301000002022-3210001221332103-2201120101103023"></a>

### Direct properties for `timeouts`

<a id="canonical-0303001310012132-0230031032322131-1331103323310010-3112213320312012-2230111003203111-1033312221200200-0023200030121131-1222300001032300"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3301032122121002-2100230120003323-3233102320222202-1200212101203203-2222323020113130-0312212112022102-2023302111310321-0303212010023221"></a>

<a id="canonical-2001222330321212-0122113303022220-1103320320320132-1222013212201132-2211020301331200-2031123211132201-2103230312330203-1001332203000133"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1020101311020230-1322300300320221-3120012223032110-2012100130121112-1303222133201312-3111133301313011-0300232311302123-2321120331000012"></a>

<a id="canonical-1021102310001110-3020231313030102-0123201101111130-0113331331201122-2230233013100031-1331322222122320-3303033122331301-0020323033232212"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2020012310320110-0223202133010212-2011121022100212-0113130100001310-1221113310130222-2312033130013211-0102321231311333-3221220011221022"></a>

<a id="canonical-2230321111201133-2300223113131022-0001311130200103-1033331312102320-0010332320203012-0232222030031102-2123003322220223-3231111221333123"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3013012011131013-2130113331323002-1022320231321331-2301301331133002-3322313112300321-1232322020332131-2203031111131013-0132220031100223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_identification` properties

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Property reference](resources--rate_limiter--reference--group-001.md#canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022)
- user_identification

<a id="canonical-3312003012232321-0022113332020233-2232101312320321-3010022233330103-2011031031301023-2223023002130311-2121232121330020-1011313320110022"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
user_identification {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113001130132132-0330033010002130-1110222033231123-2322221210120000-0310120022220312-0312302313233121-3012301201013202-2121302010031032"></a>

### Direct properties for `user_identification`

<a id="canonical-0000112211113112-3120300203101213-3223212003312011-3200230030110221-3311133132131111-1012113123303302-3011033113211323-2033013031132221"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0202031022120002-1131112120330322-2210202101003320-2102030203121323-3023322110302013-2322110101331200-1031310221333233-1300230211230123"></a>

<a id="canonical-3000210003312213-2001321021302300-0022220211312121-1022220021232310-2212301000233200-1300232110100121-3100111322123033-2000232232003111"></a>

#### `user_identification.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3300312220322302-0103231130030300-2003210030122100-1110122103222313-3331321102023033-0301020121331233-1102210120133132-2310111112201332"></a>

<a id="canonical-1333210213110033-2201210132220101-1120111011303230-3323303122233222-1132303322110332-0112330112011303-2303011133003030-0111312121130202"></a>

#### `user_identification.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3313210233321332-2312032211100321-3322113232230323-3003200321112133-3220231200032101-1221033001300302-1002033103012103-0233310311311310"></a>

<a id="canonical-0101210132113102-2022021211001302-2002131012001013-3233112322311133-2132202320010222-1100221121133100-3122132020231221-2031203000012013"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1201001003121013-3101331213200313-2313201020321231-0120310013222333-1121002120200211-1010013233310023-3133100302202113-0023301222220030"></a>

<a id="canonical-2033100210333310-2112133112232300-1013100303302201-1033231303011202-1200202201232323-2120311111013022-0103332030113122-1133332002132302"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
