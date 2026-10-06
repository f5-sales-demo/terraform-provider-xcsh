---
page_title: "xcsh_healthcheck reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_healthcheck reference."
---

# xcsh_healthcheck reference

<a id="canonical-1221303312230102-1010122132313112-0112020000032300-3233231221210023-0111020022032002-3313323101102012-2003313311113022-2311311312333213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- Property reference

<a id="canonical-2131300333311323-1223132032211132-3012211321331210-2321333010033321-1331122001312200-2123120011232210-1032131211133121-1300000120030011"></a>

### Direct properties for `xcsh_healthcheck`

<a id="canonical-2203101331123201-2120121303223323-0103221320112202-1323332201002303-2010001132302221-3131002202101200-3320120111011313-2010132112120123"></a>

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

- [default_jitter](resources--healthcheck--reference--group-001.md#canonical-0032121011311030-1110312231031233-0202202123002100-2121330323303003-2222200101211120-2333213303331031-2332113211331332-0103331322333130): complete subsection reference.

<a id="canonical-3013132123023221-3133230132123101-1113220133020112-3303012232212011-3103222312033320-0011103311110100-0230022120110031-3231311332201330"></a>

<a id="canonical-3020120220031210-3231023301031101-2003121313121213-0313232230030100-2002112323130211-3211310111232232-3021022120103333-1230100330213120"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3213001201013001-0322010112321233-3102321110032210-1310310101011012-3121100023020331-1222323031010332-1021100220223000-2130302120323303"></a>

<a id="canonical-0112200330321331-0010012332311332-3031320001113132-2000113220030110-1001001032221100-2020322301310320-1132101110002122-3121303323123030"></a>

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

<a id="canonical-1031012320312201-0111230102112201-3102132123310301-0320111021322210-2301011111012013-2232101133013331-0100130312312220-2322112123223113"></a>

<a id="canonical-1331302030112013-0211202123103033-2133311313101203-0002003222203323-3331231011113220-3121312311020013-1123321310133203-2010020220323202"></a>

#### `healthy_threshold` property

Type: `"number"`. Required.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy. Recommended: \`3\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "threshold",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "category": "threshold",
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](resources--healthcheck--reference--group-001.md#canonical-0212103111002213-0133110113331310-1231103331113202-3131112210303202-2022303102133021-2222120211122220-2222303221233210-0103203232220230): complete subsection reference.

<a id="canonical-2222130130220201-2120132100223220-3331210011010012-2122223221020221-3011212030100311-1112001313330202-2203212113201211-3301101323122002"></a>

<a id="canonical-0201010231312002-1123132012303231-3213213122102301-2112013223132021-2301230030223132-2300312221022230-3001221220033323-3013031130000333"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0003230220211112-1210231132321001-0002133130321220-2210201233322020-2332032203313102-1210121010010020-0022131021111013-2013002100201130"></a>

<a id="canonical-3211332100333131-0311123310033132-0310123003220301-1323131030220301-2232331023110301-2120001102010332-2031322002032011-1002030031330231"></a>

#### `interval` property

Type: `"number"`. Required.

Time interval in seconds between two healthcheck requests. Recommended: \`15\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-3000001222021000-2201312333201211-0023103120321122-2201002003310320-1203020313101122-3131300100120232-3220211203311211-0211112232122010"></a>

<a id="canonical-3222230132131331-0011113033201011-2033200233311123-3110331113120233-3332011003322310-3201033133030110-3130002231102313-3022202320322023"></a>

#### `jitter_percent` property

Type: `"number"`. Optional, Computed.

Exclusive with \[default\_jitter\] Specify a custom jitter value as a percentage of the health check
interval. Valid values are 0 (to disable jitter) and 10 to 50. Server applies default when omitted.
Recommended: \`30\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 50},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "Non-contiguous: {0} union [10, 50] — values 1-9 rejected by API",
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "multipleOf": 1,
    "ranges": [
      {
        "maximum": 0,
        "minimum": 0
      },
      {
        "maximum": 50,
        "minimum": 10
      }
    ]
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-50"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-50"
  }
}
```

<a id="canonical-1312222101230303-3021312110133001-3110010022213013-3331333310011233-1232332023100232-2212100021313133-0122311331000303-0311230313030033"></a>

<a id="canonical-2130303102321303-1213121010022330-0232133202331303-1323132303000103-2301133210310232-0312222003120201-2212320223332220-0210313011023223"></a>

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

<a id="canonical-3102220313032110-1003213202332221-1231010122121001-1212112110213023-3001300302120333-3133300001202022-2330223111030213-0102210321100002"></a>

<a id="canonical-0201113220000121-2211133230030321-0023202023123130-1300211232110012-1321231000002023-0032031101120313-2001323330100210-1111323200303210"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Healthcheck. Must be unique within the namespace.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3210010312333322-0233333000322211-0221303010331211-3313132322330220-0023211113303111-0332031012330112-2321033130231333-3323323100221123"></a>

<a id="canonical-0230201312303002-2333122123010330-1231332330101021-3231030133331103-3213032331001220-0123312121032012-1223133331100100-2113132010012030"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Healthcheck is created.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [tcp_health_check](resources--healthcheck--reference--group-001.md#canonical-3103201213110121-0000211233303022-2112302331133111-1001213022123321-3021132031330212-3210101201133230-2323313331201022-0113112312002002): complete subsection reference.

<a id="canonical-0032101320311333-2022331203112121-2020022310312110-0211020003230331-2311130320202323-1103111330202111-1132232311213301-1213003332131102"></a>

<a id="canonical-0122110033001033-3002230031123220-2222112202112201-0201313223332131-2331013331221112-0330310330100321-1003311230021210-2222123201231333"></a>

#### `timeout` property

Type: `"number"`. Required.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure. Recommended: \`3\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "API rejects timeout > 600 for healthchecks (global pattern says 3600)",
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [timeouts](resources--healthcheck--reference--group-001.md#canonical-1013101312032330-0120233023300220-1310333032113231-0132003013113332-1032210331013111-2100310230131132-2130033023103312-1010020313220122): complete subsection reference.

- [udp_icmp_health_check](resources--healthcheck--reference--group-001.md#canonical-2113132323233313-0012332121130131-0221300000220210-3133200021330003-1022130321212233-1211101310120213-3323202003303021-0302001212112303): complete subsection reference.

<a id="canonical-1233323330321210-0232012111200330-0201130333202021-1212102121231231-0013113310202102-0230231033031230-2211121220002022-2132120313330223"></a>

<a id="canonical-1332202032223301-3003033000331221-0210210313301200-2301122012202321-2210121232100212-1133001221100221-1201200023210001-1330121212033320"></a>

#### `unhealthy_threshold` property

Type: `"number"`. Required.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately. Recommended: \`1\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "threshold",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "category": "threshold",
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-0002002311200000-1323120213323113-3110233313112032-2123223100030012-2133303131323033-2320100322131031-1330032032211123-2332010131010201"></a>

### All schema paths for `xcsh_healthcheck`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--healthcheck--reference--group-001.md#canonical-2203101331123201-2120121303223323-0103221320112202-1323332201002303-2010001132302221-3131002202101200-3320120111011313-2010132112120123) |
| `default_jitter` | [default_jitter](resources--healthcheck--reference--group-001.md#canonical-1331331322323332-2100313331022020-2212112020232103-2323311330200113-3201221313200211-3122310222223011-3330020033021332-1122212200033110) |
| `description` | [description](resources--healthcheck--reference--group-001.md#canonical-3013132123023221-3133230132123101-1113220133020112-3303012232212011-3103222312033320-0011103311110100-0230022120110031-3231311332201330) |
| `disable` | [disable](resources--healthcheck--reference--group-001.md#canonical-3213001201013001-0322010112321233-3102321110032210-1310310101011012-3121100023020331-1222323031010332-1021100220223000-2130302120323303) |
| `healthy_threshold` | [healthy_threshold](resources--healthcheck--reference--group-001.md#canonical-1031012320312201-0111230102112201-3102132123310301-0320111021322210-2301011111012013-2232101133013331-0100130312312220-2322112123223113) |
| `http_health_check` | [http_health_check](resources--healthcheck--reference--group-001.md#canonical-3113122122303310-3101033110201203-0213011102020012-0220211013201122-1032312331100200-0211223031300133-0313033320112103-2213112120321101) |
| `http_health_check.expected_response` | [http_health_check.expected_response](resources--healthcheck--reference--group-001.md#canonical-0333301210131320-2303211100322333-0230302212021023-0020321020011233-3131312320030220-2200001200032332-0003120121322332-1323112231111301) |
| `http_health_check.expected_status_codes` | [http_health_check.expected_status_codes](resources--healthcheck--reference--group-001.md#canonical-0212120000203023-0311300310302012-2320231100223213-0332101002110101-2333220332301102-2313233030211331-3333100121012332-2021332113123023) |
| `http_health_check.headers` | [http_health_check.headers](resources--healthcheck--reference--group-001.md#canonical-2212011132321101-1113131132300120-1303023111321111-2020033321030123-3132110310010132-3321102201330021-0320011302303321-1101303231112000) |
| `http_health_check.host_header` | [http_health_check.host_header](resources--healthcheck--reference--group-001.md#canonical-0022303311211021-3332110203030120-0033013312130323-3121002031022020-2230133031212121-3223212111113022-3202120202101001-3101122231310203) |
| `http_health_check.path` | [http_health_check.path](resources--healthcheck--reference--group-001.md#canonical-1322121201313220-1012103203333220-1222130312333132-0322100331231231-0122223200322032-1111130030133123-0313312030000311-3230330302120022) |
| `http_health_check.request_headers_to_remove` | [http_health_check.request_headers_to_remove](resources--healthcheck--reference--group-001.md#canonical-2111112133121033-1002313220002030-3121302033033103-0300032023123211-1122321310321330-3100103133111101-2321111303122313-3101123230223112) |
| `http_health_check.use_http2` | [http_health_check.use_http2](resources--healthcheck--reference--group-001.md#canonical-3101303220221333-0222002031202030-2130311002203232-0003313130031121-1030003112301122-2211201033011211-0210110202012313-3333203113133122) |
| `http_health_check.use_origin_server_name` | [http_health_check.use_origin_server_name](resources--healthcheck--reference--group-001.md#canonical-3303220132033131-2330200330211010-0010123033232221-1300313312331323-0321110013332202-0303232120201122-1002022120101232-3120030100211230) |
| `id` | [ID](resources--healthcheck--reference--group-001.md#canonical-2222130130220201-2120132100223220-3331210011010012-2122223221020221-3011212030100311-1112001313330202-2203212113201211-3301101323122002) |
| `interval` | [interval](resources--healthcheck--reference--group-001.md#canonical-0003230220211112-1210231132321001-0002133130321220-2210201233322020-2332032203313102-1210121010010020-0022131021111013-2013002100201130) |
| `jitter_percent` | [jitter_percent](resources--healthcheck--reference--group-001.md#canonical-3000001222021000-2201312333201211-0023103120321122-2201002003310320-1203020313101122-3131300100120232-3220211203311211-0211112232122010) |
| `labels` | [labels](resources--healthcheck--reference--group-001.md#canonical-1312222101230303-3021312110133001-3110010022213013-3331333310011233-1232332023100232-2212100021313133-0122311331000303-0311230313030033) |
| `name` | [name](resources--healthcheck--reference--group-001.md#canonical-3102220313032110-1003213202332221-1231010122121001-1212112110213023-3001300302120333-3133300001202022-2330223111030213-0102210321100002) |
| `namespace` | [namespace](resources--healthcheck--reference--group-001.md#canonical-3210010312333322-0233333000322211-0221303010331211-3313132322330220-0023211113303111-0332031012330112-2321033130231333-3323323100221123) |
| `tcp_health_check` | [tcp_health_check](resources--healthcheck--reference--group-001.md#canonical-1102322130010123-0330003020001103-3210220030211230-0320111200201102-2222101102000102-0023212002233302-2302031223011023-2312323020023202) |
| `tcp_health_check.expected_response` | [tcp_health_check.expected_response](resources--healthcheck--reference--group-001.md#canonical-1111003332210203-2331133212331120-1322202302130312-1030132120231030-1112233023022322-3012313330133111-3030122331213010-1123202233200200) |
| `tcp_health_check.send_payload` | [tcp_health_check.send_payload](resources--healthcheck--reference--group-001.md#canonical-3011321332113000-1301201122120232-1322133230112233-0310123323130210-3100110033231123-2311220320221303-2103022013203021-3313300322331320) |
| `timeout` | [timeout](resources--healthcheck--reference--group-001.md#canonical-0032101320311333-2022331203112121-2020022310312110-0211020003230331-2311130320202323-1103111330202111-1132232311213301-1213003332131102) |
| `timeouts` | [timeouts](resources--healthcheck--reference--group-001.md#canonical-0201231321030211-2000122012322323-1232013101201321-3032103030200221-1010130131200223-2313301000300330-2020021033030301-3230101322022212) |
| `timeouts.create` | [timeouts.create](resources--healthcheck--reference--group-001.md#canonical-3303030301230002-0210332301000213-0132021231130113-0011300111032002-1013012332120312-2123233303020323-3110013301231232-3212123013012311) |
| `timeouts.delete` | [timeouts.delete](resources--healthcheck--reference--group-001.md#canonical-0120013301310031-2030123031222030-1112203003102100-1113131100112033-1130321213232122-2130021231222000-1010031011021301-1321330103000312) |
| `timeouts.read` | [timeouts.read](resources--healthcheck--reference--group-001.md#canonical-2203311002212023-0200101002202333-1021322310211313-0012112311331022-3230032100211130-1101001230022212-3222201023002132-2221312212132131) |
| `timeouts.update` | [timeouts.update](resources--healthcheck--reference--group-001.md#canonical-3311331203230203-3201301010200030-2331110333330301-1100121332323333-1313203322030222-1111123321100103-2032131131110100-2003120330130002) |
| `udp_icmp_health_check` | [udp_icmp_health_check](resources--healthcheck--reference--group-001.md#canonical-0130020323011213-3113002311300330-0023131320312023-1201330200333001-2200033021212210-3111012111121130-3302002023201121-2311123301333321) |
| `unhealthy_threshold` | [unhealthy_threshold](resources--healthcheck--reference--group-001.md#canonical-1233323330321210-0232012111200330-0201130333202021-1212102121231231-0013113310202102-0230231033031230-2211121220002022-2132120313330223) |

<a id="canonical-0032121011311030-1110312231031233-0202202123002100-2121330323303003-2222200101211120-2333213303331031-2332113211331332-0103331322333130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_jitter` properties

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-1221303312230102-1010122132313112-0112020000032300-3233231221210023-0111020022032002-3313323101102012-2003313311113022-2311311312333213)
- default_jitter

<a id="canonical-1331331322323332-2100313331022020-2212112020232103-2323311330200113-3201221313200211-3122310222223011-3330020033021332-1122212200033110"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_jitter, jitter\_percent; Default: default\_jitter\] Configuration parameter for
default jitter.

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

- [default_jitter](resources--healthcheck--reference--group-001.md#canonical-1331331322323332-2100313331022020-2212112020232103-2323311330200113-3201221313200211-3122310222223011-3330020033021332-1122212200033110)
- [jitter_percent](resources--healthcheck--reference--group-001.md#canonical-3000001222021000-2201312333201211-0023103120321122-2201002003310320-1203020313101122-3131300100120232-3220211203311211-0211112232122010)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_jitter = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212103111002213-0133110113331310-1231103331113202-3131112210303202-2022303102133021-2222120211122220-2222303221233210-0103203232220230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_health_check` properties

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-1221303312230102-1010122132313112-0112020000032300-3233231221210023-0111020022032002-3313323101102012-2003313311113022-2311311312333213)
- http_health_check

<a id="canonical-3113122122303310-3101033110201203-0213011102020012-0220211013201122-1032312331100200-0211223031300133-0313033320112103-2213112120321101"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: http\_health\_check, tcp\_health\_check, udp\_icmp\_health\_check\] Healthy if 'GET' method
on URL 'HTTP(s)://&lt;host&gt;/&lt;path&gt;' with optional '&lt;header&gt;' returns success. 'host'
is not used for DNS resolution. It is used as HTTP Header in the request.

Additional upstream details:

Healthy if "GET" method on URL "HTTP(s)://&lt;host&gt;/&lt;path&gt;" with optional "&lt;header&gt;"
returns success. "host" is not used for DNS resolution.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("path"),
  validators.ConflictingObjectAttributes("host_header",
    "use_origin_server_name")}
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
  "x-ves-oneof-field-host_header_choice": "[\"host_header\",\"use_origin_server_name\"]"
}
```

OneOf alternatives in this subsection:

- [http_health_check](resources--healthcheck--reference--group-001.md#canonical-3113122122303310-3101033110201203-0213011102020012-0220211013201122-1032312331100200-0211223031300133-0313033320112103-2213112120321101)
- [tcp_health_check](resources--healthcheck--reference--group-001.md#canonical-1102322130010123-0330003020001103-3210220030211230-0320111200201102-2222101102000102-0023212002233302-2302031223011023-2312323020023202)
- [udp_icmp_health_check](resources--healthcheck--reference--group-001.md#canonical-0130020323011213-3113002311300330-0023131320312023-1201330200333001-2200033021212210-3111012111121130-3302002023201121-2311123301333321)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332112223302300-3133013233332023-0212133213210322-2202233221120122-0311333003110300-3132311330001222-0320332020220020-0112312102121113"></a>

### Direct properties for `http_health_check`

<a id="canonical-0333301210131320-2303211100322333-0230302212021023-0020321020011233-3131312320030220-2200001200032332-0003120121322332-1323112231111301"></a>

#### `http_health_check.expected_response` property

Type: `"string"`. Optional, Computed.

Raw bytes expected in the response of HTTP health check. Input is to be given in Hex encoded format.
If left empty, then response body is not considered for evaluating health check status. Server
applies default when omitted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-0212120000203023-0311300310302012-2320231100223213-0332101002110101-2333220332301102-2313233030211331-3333100121012332-2021332113123023"></a>

<a id="canonical-2200002221120221-0101203223131013-0132203001100123-0022220210023112-2323030030233312-2002211323330113-0020303122303333-1123012301122013"></a>

#### `http_health_check.expected_status_codes` property

Type: `["list", "string"]`. Optional, Computed.

Specifies a list of HTTP response status codes considered healthy. To treat default HTTP expected
status code 200 as healthy, user has to configure it explicitly. This is a list of strings, each of
which is single HTTP status code or a range with start and end values separated by '-'. Defaults to
\`\[\]\`. Server applies default when omitted.

Additional upstream details:

This is a list of strings, each of which is single HTTP status code or a range with start and end
values separated by "-".

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.http_status_range": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "10",
    "ves.io.schema.rules.repeated.items.string.min_len": "3",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_status_range": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "10",
    "ves.io.schema.rules.repeated.items.string.min_len": "3",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2212011132321101-1113131132300120-1303023111321111-2020033321030123-3132110310010132-3321102201330021-0320011302303321-1101303231112000"></a>

<a id="canonical-2130010122110123-2213231120133322-3003323002203103-2002113233113211-2333023012120023-2232220011023103-0231222033110111-3132012132110120"></a>

#### `http_health_check.headers` property

Type: `["map", "string"]`. Optional, Computed.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked cluster. This is a list of key-value pairs. Defaults to \`map\[\]\`. Server applies default
when omitted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":256,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"256\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"2048\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":2048,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "256",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "2048",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 2048,
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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0022303311211021-3332110203030120-0033013312130323-3121002031022020-2230133031212121-3223212111113022-3202120202101001-3101122231310203"></a>

<a id="canonical-0212121032301031-1012331223012010-3121103310223220-2323101221312112-0032012201130312-1122230112022100-1011130221013101-3111032101330130"></a>

#### `http_health_check.host_header` property

Type: `"string"`. Optional.

Exclusive with \[use\_origin\_server\_name\] The value of the host header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-1322121201313220-1012103203333220-1222130312333132-0322100331231231-0122223200322032-1111130030133123-0313312030000311-3230330302120022"></a>

<a id="canonical-2221011232213301-2330201321301301-3120322012323201-1321113333320132-3130221322203323-2103100332312131-3030002203212331-2103303302222233"></a>

#### `http_health_check.path` property

Type: `"string"`. Optional.

Specifies the HTTP path that will be requested during health checking. Recommended: \`/\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-2111112133121033-1002313220002030-3121302033033103-0300032023123211-1122321310321330-3100103133111101-2321111303122313-3101123230223112"></a>

<a id="canonical-1202230210133022-0310221123322322-2020023323323012-0310011021333223-1021321230102310-1112002201132101-2223203320201001-3002001323231012"></a>

#### `http_health_check.request_headers_to_remove` property

Type: `["list", "string"]`. Optional, Computed.

Specifies a list of HTTP headers that should be removed from each request that is sent to the health
checked cluster. This is a list of keys of headers. Defaults to \`\[\]\`. Server applies default
when omitted.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-3101303220221333-0222002031202030-2130311002203232-0003313130031121-1030003112301122-2211201033011211-0210110202012313-3333203113133122"></a>

<a id="canonical-0022310022230231-1130200123100202-2030203210203313-3110113301131003-1001120001331232-1011231333323120-1131012303223302-2003100103203231"></a>

#### `http_health_check.use_http2` property

Type: `"bool"`. Optional, Computed.

If set, health checks will be made using HTTP/2. Defaults to \`false\`. Server applies default when
omitted. Recommended: \`false\`.

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

- [use_origin_server_name](resources--healthcheck--reference--group-001.md#canonical-2132133223202302-3212113232122201-0322200101231130-3023020122113221-1213112133312113-2313100320200122-0211103112310313-0300123102313030): complete subsection reference.

<a id="canonical-2132133223202302-3212113232122201-0322200101231130-3023020122113221-1213112133312113-2313100320200122-0211103112310313-0300123102313030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_health_check.use_origin_server_name` properties

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-1221303312230102-1010122132313112-0112020000032300-3233231221210023-0111020022032002-3313323101102012-2003313311113022-2311311312333213)
- [http_health_check](resources--healthcheck--reference--group-001.md#canonical-0212103111002213-0133110113331310-1231103331113202-3131112210303202-2022303102133021-2222120211122220-2222303221233210-0103203232220230)
- http_health_check.use_origin_server_name

<a id="canonical-3303220132033131-2330200330211010-0010123033232221-1300313312331323-0321110013332202-0303232120201122-1002022120101232-3120030100211230"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
use_origin_server_name = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103201213110121-0000211233303022-2112302331133111-1001213022123321-3021132031330212-3210101201133230-2323313331201022-0113112312002002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tcp_health_check` properties

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-1221303312230102-1010122132313112-0112020000032300-3233231221210023-0111020022032002-3313323101102012-2003313311113022-2311311312333213)
- tcp_health_check

<a id="canonical-1102322130010123-0330003020001103-3210220030211230-0320111200201102-2222101102000102-0023212002233302-2302031223011023-2312323020023202"></a>

Type: `"object"`. single nested block, Optional.

Healthy if TCP connection is successful and response payload matches &lt;expected\_response&gt;.

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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232303221132003-3101300122300301-3332203202122020-1230301103320311-0122032202000330-3323200232122220-0310322301233212-1331111003011131"></a>

### Direct properties for `tcp_health_check`

<a id="canonical-1111003332210203-2331133212331120-1322202302130312-1030132120231030-1112233023022322-3012313330133111-3030122331213010-1123202233200200"></a>

#### `tcp_health_check.expected_response` property

Type: `"string"`. Optional.

Raw bytes expected in the request. Describes the encoding of the payload bytes in the payload. Hex
encoded payload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-3011321332113000-1301201122120232-1322133230112233-0310123323130210-3100110033231123-2311220320221303-2103022013203021-3313300322331320"></a>

<a id="canonical-2201312300022111-3002231011212210-1301230330332221-3021112013133323-1023102130011311-0220032210133220-1103201131103202-3320323122103213"></a>

#### `tcp_health_check.send_payload` property

Type: `"string"`. Optional.

Raw bytes sent in the request. Empty payloads imply a connect-only health check. Describes the
encoding of the payload bytes in the payload. Hex encoded payload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-1013101312032330-0120233023300220-1310333032113231-0132003013113332-1032210331013111-2100310230131132-2130033023103312-1010020313220122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-1221303312230102-1010122132313112-0112020000032300-3233231221210023-0111020022032002-3313323101102012-2003313311113022-2311311312333213)
- timeouts

<a id="canonical-0201231321030211-2000122012322323-1232013101201321-3032103030200221-1010130131200223-2313301000300330-2020021033030301-3230101322022212"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033022132122102-2023113012203210-2111013221130022-3001132202212203-3221330300011230-2013312311230323-2221201132033333-3303311223221300"></a>

### Direct properties for `timeouts`

<a id="canonical-3303030301230002-0210332301000213-0132021231130113-0011300111032002-1013012332120312-2123233303020323-3110013301231232-3212123013012311"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0120013301310031-2030123031222030-1112203003102100-1113131100112033-1130321213232122-2130021231222000-1010031011021301-1321330103000312"></a>

<a id="canonical-0212031211021011-0301001303002003-2102100303132110-1322030320231212-3033222213311023-3233330022033031-2013210301111022-3101233023013123"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2203311002212023-0200101002202333-1021322310211313-0012112311331022-3230032100211130-1101001230022212-3222201023002132-2221312212132131"></a>

<a id="canonical-3203221111133103-0211022001202111-0203011033023001-0211033312233020-3100033200002102-1001033201322032-1203300030132101-3021312233300003"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3311331203230203-3201301010200030-2331110333330301-1100121332323333-1313203322030222-1111123321100103-2032131131110100-2003120330130002"></a>

<a id="canonical-0311230031230123-1232022000111033-1012300120311310-3302222130100030-0023301303013010-0003332330130212-1030311323320001-0203230103220121"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2113132323233313-0012332121130131-0221300000220210-3133200021330003-1022130321212233-1211101310120213-3323202003303021-0302001212112303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `udp_icmp_health_check` properties

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Property reference](resources--healthcheck--reference--group-001.md#canonical-1221303312230102-1010122132313112-0112020000032300-3233231221210023-0111020022032002-3313323101102012-2003313311113022-2311311312333213)
- udp_icmp_health_check

<a id="canonical-0130020323011213-3113002311300330-0023131320312023-1201330200333001-2200033021212210-3111012111121130-3302002023201121-2311123301333321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for udp icmp health check.

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
udp_icmp_health_check = {}
```

This is an empty object or choice marker. It has no direct properties.
