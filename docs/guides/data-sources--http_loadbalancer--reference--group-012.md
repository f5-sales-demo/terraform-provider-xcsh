---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0130022123210311-2201101302112333-3223100112102232-2001232321020310-2103310302231321-1212133113020123-3210211111123300-2013132211103011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-011.md#canonical-2003210000020113-3233301301011330-1021101022302213-0230313020030322-3001200301020013-0103220011320230-0112202123312333-1033003002012021)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-011.md#canonical-0221321133322220-3013322000031310-1233102200200202-1200302200121121-0101231313122310-0031202001233001-2302312323221131-3120032113021321)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-1122220332211332-1311212103320002-1311303232012131-2122230133111303-3310120000131131-0312020332200233-3211312220202032-3112021131223021"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-0111120320012310-2130310010312300-3023220311132200-2203013332211220-2033123211013113-0132021112123302-2011303011313112-0313123021103023"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata`

<a id="canonical-2303230222333030-2112003203220200-3002002331333022-0330003110301310-0313001103213331-3120002103231231-3113332311102013-3323021300333110"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0031320123310230-2212001111330132-1013023230013113-2013203222313131-3203222110211112-3102221301313203-3233200200033112-3320231001210133"></a>

<a id="canonical-3210211301313302-3202203223033312-2321300321012310-2211012320112321-1121110301021300-1213113303020233-1220231110002302-0333320121002113"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-0013301332333223-3010223032132030-2031203110110223-0333002321133300-3231120030222112-1231232213022330-0203313330002112-0322323122233321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-011.md#canonical-2003210000020113-3233301301011330-1021101022302213-0230313020030322-3001200301020013-0103220011320230-0112202123312333-1033003002012021)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-011.md#canonical-0221321133322220-3013322000031310-1233102200200202-1200302200121121-0101231313122310-0031202001233001-2302312323221131-3120032113021321)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-3320213113201320-0110302232331300-3120320132311203-1213123312330311-2201311213132030-3033212333232331-0331000322221121-0020212120310110"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-0310313330332330-0022123020303113-1222222103313022-2312211212111033-1101320200121302-2112230113000330-1100032230231100-1331220220223002"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list.path`

<a id="canonical-2232301112230002-3212102312132111-1130102130303122-1110031001233320-1313330303101320-3323303303021132-2012333311133113-3320100302310121"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0103112321323213-1031010033203011-0232331211003230-1121132311012110-3203220213212113-2230330111303030-2201020121302212-2033332212033013"></a>

<a id="canonical-1333311223223222-3133202300001221-2231022222111321-3013311330301232-2111020221232312-3031210131210003-3112203123120023-3201202231121210"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1033312033130333-3230013013001113-1331031211230211-0313012232233200-1021000202131122-1223003330332303-0121023111101210-2210023213321222"></a>

<a id="canonical-0130033133012010-0022132301102323-0330120222310321-3012002022223101-1313012022211130-0113202202330122-0003302300133132-2020313031012021"></a>

#### `bot_defense.policy.js_insert_all_pages_except.exclude_list.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3112210113020320-0320111112203312-0330033032002301-0321012213011221-0121321013302000-2013213220000111-2201220330211321-1233000303031312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- bot_defense.policy.js_insertion_rules

<a id="canonical-3303203332022320-3113033301021023-2102000131010202-2030111002221331-3300033301022002-0311031332000321-1313222131112223-0002131320110301"></a>

Type: `"single"`. Computed.

This defines custom JavaScript insertion rules for Bot Defense Policy.

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

<a id="canonical-3201103222011033-2120230231212113-1211032133020200-1331201000032122-2122100212331132-0223212321222310-2220312212322102-1103201303032013"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules`

- [exclude_list](data-sources--http_loadbalancer--reference--group-012.md#canonical-1223210332032111-3213001101220132-2130012300023111-2003111033231131-3223113001311202-1002222113333232-2323122330210131-3112030110201030): complete subsection reference.

- [rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-0333331032203313-3033132210300301-0033202011331210-2203112033111301-3120130001003300-3233013131131031-1120031021003123-0030020320013102): complete subsection reference.

<a id="canonical-1223210332032111-3213001101220132-2130012300023111-2003111033231131-3223113001311202-1002222113333232-2323122330210131-3112030110201030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112210113020320-0320111112203312-0330033032002301-0321012213011221-0121321013302000-2013213220000111-2201220330211321-1233000303031312)
- bot_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-1023213102101332-2111320210321100-3101213101121113-2300213111031203-0032233210320313-1023103203030200-2200213223330030-0231001131112331"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3121002113301103-0212220032221101-0312011002211103-1031232303123032-1120133313000322-0233012103331231-0322101211022331-2122030310313232"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list`

- [any_domain](data-sources--http_loadbalancer--reference--group-012.md#canonical-0301330223323233-2211001230112213-1103311331233321-3012010111002332-1123333211002013-2031313002221310-0000001031233031-1002123332203313): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-012.md#canonical-1323311333101012-2200003121130011-3300000200022002-3111200022231310-3122120230103230-1112223322101121-1300021033202023-2033003102331211): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-012.md#canonical-1111101222313010-0113111021312112-2323301231122200-0222332103023202-2022000032130103-2312332332213113-2223211311031012-3011000221101320): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-012.md#canonical-2003011320333333-0302330021002313-0010320312300333-1320121212311203-0320311100212010-0322100132023220-0120323310210111-2001321223010210): complete subsection reference.

<a id="canonical-0301330223323233-2211001230112213-1103311331233321-3012010111002332-1123333211002013-2031313002221310-0000001031233031-1002123332203313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112210113020320-0320111112203312-0330033032002301-0321012213011221-0121321013302000-2013213220000111-2201220330211321-1233000303031312)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-012.md#canonical-1223210332032111-3213001101220132-2130012300023111-2003111033231131-3223113001311202-1002222113333232-2323122330210131-3112030110201030)
- bot_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-2130221110000331-3131012211100122-2321323101300332-1231231110012221-0222220123102222-1221003313323021-0200100020202130-2123210210131121"></a>

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

<a id="canonical-1323311333101012-2200003121130011-3300000200022002-3111200022231310-3122120230103230-1112223322101121-1300021033202023-2033003102331211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112210113020320-0320111112203312-0330033032002301-0321012213011221-0121321013302000-2013213220000111-2201220330211321-1233000303031312)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-012.md#canonical-1223210332032111-3213001101220132-2130012300023111-2003111033231131-3223113001311202-1002222113333232-2323122330210131-3112030110201030)
- bot_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-2103211301323221-2213110202330012-1210321001030013-0021121113200013-1312133030002201-0002301231331131-3323201230103101-0002032223222233"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-3302323202023222-0213222133322112-1032311120301031-0010102002312221-0010301103003331-1113100132300303-0012010220133300-0311020021132323"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.domain`

<a id="canonical-0011100021300111-3133100311031110-1312131011313012-2131313003100012-3023311220213212-3110121030212133-2332011211222012-1311303221213100"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3021111223200000-0102303202132312-3000112220022201-0202121122032332-3113100220210030-0001020020221322-1122131221211201-0303033320231112"></a>

<a id="canonical-2021100321121300-2312311111001300-2302100013001202-3213133223101333-0113030300221132-0033220002213313-0312021131130012-1020001312312001"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0311211220222110-3030333121000013-0031030331312201-1321011033103311-2003212032131323-2132011010003202-2322312110230033-1330231031330332"></a>

<a id="canonical-1203311000013001-3103213312103231-3310021231012002-3200311201012212-2303112021220103-3232111233032011-3332211023210030-1303201232013012"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.domain.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1111101222313010-0113111021312112-2323301231122200-0222332103023202-2022000032130103-2312332332213113-2223211311031012-3011000221101320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112210113020320-0320111112203312-0330033032002301-0321012213011221-0121321013302000-2013213220000111-2201220330211321-1233000303031312)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-012.md#canonical-1223210332032111-3213001101220132-2130012300023111-2003111033231131-3223113001311202-1002222113333232-2323122330210131-3112030110201030)
- bot_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-2212232312320221-1302010001302303-3011022301003321-2002310132202233-2111001320032131-1232201311030101-1332020123321013-1133110020210200"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-3232333310212203-2331102030012020-0320030323102233-3223001330010221-1210221312232331-3232211121023130-0303212212232300-0323311303122301"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.metadata`

<a id="canonical-1110233323221030-1130202133122101-1101203211311301-1311032001010122-1200113213231012-1221223120313302-1120111002032032-0311012303112221"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1330113230013331-0332332101132132-1321233113130113-2030103200220120-1103221023132320-2112331011221010-3202330012202023-1013202113300222"></a>

<a id="canonical-3311131231211332-1102333131112203-1232120013331032-2321033111320203-1030232020003110-2213020021312333-1001111132213221-3102221322300030"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-2003011320333333-0302330021002313-0010320312300333-1320121212311203-0320311100212010-0322100132023220-0120323310210111-2001321223010210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112210113020320-0320111112203312-0330033032002301-0321012213011221-0121321013302000-2013213220000111-2201220330211321-1233000303031312)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-012.md#canonical-1223210332032111-3213001101220132-2130012300023111-2003111033231131-3223113001311202-1002222113333232-2323122330210131-3112030110201030)
- bot_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-2310130322210213-2123222102300332-0300202112000012-2103310322303222-1020023123330120-0330221013221111-2203313012000020-3201000211322000"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-2330001323300113-1231010011232222-1303112310110221-3010120130030001-3000133113211202-3103123020000022-1330103232312101-2033123212110312"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.exclude_list.path`

<a id="canonical-2132213213303223-1231221010020001-2111211111130220-3013301233012312-3032120223123012-2201201313102310-2311001122010120-3231310031030233"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3221013223121300-1030222331332022-2132123011133111-2120112322120203-3320201022111111-0310113032230031-2032323101030321-0100213323110211"></a>

<a id="canonical-0222101111010013-1301333321332022-1000320111210302-1002310333110300-0031121211102310-3020203203022210-0320213123110000-3200211031331301"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2123220333303230-3113021203303310-2132331202111102-2231321122122012-2021231331200011-0303313221302133-2110103133103232-2201321202303331"></a>

<a id="canonical-1220012100222333-2012320032203223-2113000212202113-1322210121100022-0203330111331222-0003021113110021-1330111312322333-2233320203121311"></a>

#### `bot_defense.policy.js_insertion_rules.exclude_list.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0333331032203313-3033132210300301-0033202011331210-2203112033111301-3120130001003300-3233013131131031-1120031021003123-0030020320013102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112210113020320-0320111112203312-0330033032002301-0321012213011221-0121321013302000-2013213220000111-2201220330211321-1233000303031312)
- bot_defense.policy.js_insertion_rules.rules

<a id="canonical-1323333320300332-2212101201303333-2302231122103122-0002232103333002-3302331110100133-3101321311103300-3031223031011020-3121200311101001"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1221031213221110-1132002103010031-3012230000220301-3200030011012010-3331312002112232-2022123130122320-2100313000212020-1101123200321332"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules`

- [any_domain](data-sources--http_loadbalancer--reference--group-012.md#canonical-1332332321023030-3030332131330131-3223032223202323-1211323122000102-2112232312212022-0331201320133220-1233010030112033-1012003220102100): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-012.md#canonical-1222313002213011-2011211222100112-2301121221010003-3312133203213200-0100303233020332-2330023200013222-3333022232231301-0102023203001110): complete subsection reference.

<a id="canonical-0301120101011300-2122111102222222-1210331120021330-1201311010132210-0132100021112022-0031023310133221-0021123022030231-3010003120122203"></a>

<a id="canonical-3232202110122323-0102301302103103-1030002213123132-0132323330013332-2202301122212013-3231300310301022-3331322213311233-2302032013201300"></a>

#### `bot_defense.policy.js_insertion_rules.rules.javascript_location` property

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-012.md#canonical-0112001021211111-0320120112003120-2213130103101121-2313122001333311-2311303332031231-3102121103232120-0033322330011103-2310332033313332): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-012.md#canonical-0233112023220202-0011002202130323-2301321210212133-1031021002132110-1221022013232331-2232210031032121-0313300003031011-3203011013232330): complete subsection reference.

<a id="canonical-1332332321023030-3030332131330131-3223032223202323-1211323122000102-2112232312212022-0331201320133220-1233010030112033-1012003220102100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112210113020320-0320111112203312-0330033032002301-0321012213011221-0121321013302000-2013213220000111-2201220330211321-1233000303031312)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-0333331032203313-3033132210300301-0033202011331210-2203112033111301-3120130001003300-3233013131131031-1120031021003123-0030020320013102)
- bot_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-1212133213000202-0322330132031330-1112303331102300-1220112020110103-1112001203311230-3103332210102202-1002033331312000-1200013310012102"></a>

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

<a id="canonical-1222313002213011-2011211222100112-2301121221010003-3312133203213200-0100303233020332-2330023200013222-3333022232231301-0102023203001110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112210113020320-0320111112203312-0330033032002301-0321012213011221-0121321013302000-2013213220000111-2201220330211321-1233000303031312)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-0333331032203313-3033132210300301-0033202011331210-2203112033111301-3120130001003300-3233013131131031-1120031021003123-0030020320013102)
- bot_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-2200331130012313-3000021020102122-2230110022201233-2330012120100002-0131333110133300-2310100300222321-0101030230123111-0130222103300213"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-0233111230011221-3021132201210200-2021230131132300-0300021030120012-3132121132112011-1323322331300000-1111122211010103-3213311231021113"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.domain`

<a id="canonical-2231210200133122-2321132203023211-1222011010033321-1021303030022200-3311300233230322-0133122101030331-2031322123310100-2133121330013201"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2222203322313123-0313221302130322-1212032031222020-2322223010113212-2211001333111200-1331233103303121-1013321101211032-3112331011332321"></a>

<a id="canonical-1323023212002123-2033033100013121-0222102020101223-0122032222010210-2032133230132111-3333212201310330-2210102311133022-1211132211123203"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3031323303033022-0100203220131012-1312100130201033-2231220033233100-1200213303301103-3203223321132110-0212000333213303-2020312331023203"></a>

<a id="canonical-0000200111300311-3021011111021220-2302102023220030-0233120021211130-2110321210113213-0032110023330231-2023312220210101-2330023211132103"></a>

#### `bot_defense.policy.js_insertion_rules.rules.domain.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0112001021211111-0320120112003120-2213130103101121-2313122001333311-2311303332031231-3102121103232120-0033322330011103-2310332033313332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112210113020320-0320111112203312-0330033032002301-0321012213011221-0121321013302000-2013213220000111-2201220330211321-1233000303031312)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-0333331032203313-3033132210300301-0033202011331210-2203112033111301-3120130001003300-3233013131131031-1120031021003123-0030020320013102)
- bot_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-1112021212102130-2302331122121030-0123101202230110-0320121032112002-2130323220033300-1223220311011321-3223132331030112-3212022123110201"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-0331120331200321-3232213110311033-0011223201321322-3312203001011111-3032013013110323-1321101311031111-1011312220313200-3111320212303300"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.metadata`

<a id="canonical-2013221030012302-3003303232321332-3012323112321200-3011023320233303-1312211022101210-1201222232301121-3130230131010321-0212221132120300"></a>

#### `bot_defense.policy.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0121301203101102-3321121323113231-0100312131131221-1322012031121131-2130222331003201-1001330020310310-3232313232233011-1003312010011123"></a>

<a id="canonical-1020001000333131-3023023130321233-0230203020130112-0330230310210202-2022303013233113-0202220023102231-3311001332213222-3321100132200012"></a>

#### `bot_defense.policy.js_insertion_rules.rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-0233112023220202-0011002202130323-2301321210212133-1031021002132110-1221022013232331-2232210031032121-0313300003031011-3203011013232330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insertion_rules.rules.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112210113020320-0320111112203312-0330033032002301-0321012213011221-0121321013302000-2013213220000111-2201220330211321-1233000303031312)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-012.md#canonical-0333331032203313-3033132210300301-0033202011331210-2203112033111301-3120130001003300-3233013131131031-1120031021003123-0030020320013102)
- bot_defense.policy.js_insertion_rules.rules.path

<a id="canonical-0023121110212323-1232033130231212-1033112100323111-0033303032002021-0031120210200312-1120223030332212-0330221013103030-2130323000232212"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-0031232022330232-0132120010231223-2003111133221203-0233130112300003-2222103322213013-1201212101203313-3312210220021000-2132202321312312"></a>

### Direct properties for `bot_defense.policy.js_insertion_rules.rules.path`

<a id="canonical-3213303122022113-0020310302011023-1212122201302331-3333132001233022-3233111112001213-1100123010122333-3301002330232313-2322231032210222"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3110130323231320-3301112111313131-0201113123123231-2032221100122303-0232301020213001-3010230332031013-2321303330202201-1111112002330330"></a>

<a id="canonical-0013121123311303-0122221011301303-3202232030101120-3022222333322210-3233213202130112-0003332302020021-1032210021213011-2311102010212130"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1023101233110323-1213330231122132-3023121001313211-0023000320332131-0103230022210303-3032203001120300-0020312100111303-0022220213022301"></a>

<a id="canonical-2000123023311013-0233313211112133-0030200201232103-2222232210000003-0000210312312201-0012331121213132-3032330003111300-3020223212110231"></a>

#### `bot_defense.policy.js_insertion_rules.rules.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1020010112011222-1203033213323302-3102202132222212-1233320101300013-1001211221013333-1123121322230200-3301110301123030-2322013122221223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- bot_defense.policy.mobile_sdk_config

<a id="canonical-1313021333220230-3321201201133011-3231021230133321-0320111032101013-3200323232102121-0031212220130330-1133110110002232-3012133120000212"></a>

Type: `"single"`. Computed.

Mobile SDK Configuration. Mobile SDK configuration.

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

<a id="canonical-2122212313101210-3131230102123030-3032221121100202-2032031110100330-0121132332021330-1300231120220113-2121303320121012-0001320003221220"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config`

- [mobile_identifier](data-sources--http_loadbalancer--reference--group-012.md#canonical-3023130302320000-0112122023101330-2000131211111302-1130211032012131-1113011010030113-2131210333211310-0200113031121103-3010021121130003): complete subsection reference.

<a id="canonical-3023130302320000-0112122023101330-2000131211111302-1130211032012131-1113011010030113-2131210333211310-0200113031121103-3010021121130003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-012.md#canonical-1020010112011222-1203033213323302-3102202132222212-1233320101300013-1001211221013333-1123121322230200-3301110301123030-2322013122221223)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="canonical-2312321011233302-0131130201112211-1303021013232121-2320322301103202-3030110231131203-0231221013313020-2010311232300303-0112133300220022"></a>

Type: `"single"`. Computed.

Mobile Traffic Identifier. Mobile traffic identifier type.

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

<a id="canonical-3211302112133011-2030112100011230-2222302213330220-1013203301001122-2333103033202231-2121333333123101-1110001332200302-2212221110213030"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier`

- [headers](data-sources--http_loadbalancer--reference--group-012.md#canonical-3200102213000012-0131210312123312-0102011032301030-2303202332302032-0112032133212221-1021213033121120-2011002211130022-1133130133321223): complete subsection reference.

<a id="canonical-3200102213000012-0131210312123312-0102011032301030-2303202332302032-0112032133212221-1021213033121120-2011002211130022-1133130133321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-012.md#canonical-1020010112011222-1203033213323302-3102202132222212-1233320101300013-1001211221013333-1123121322230200-3301110301123030-2322013122221223)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-012.md#canonical-3023130302320000-0112122023101330-2000131211111302-1130211032012131-1113011010030113-2131210333211310-0200113031121103-3010021121130003)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-3301210201030232-2302131110321032-0213021210310013-1132012201122030-0021233002113131-0302222102220211-2013312221131210-2120332032202322"></a>

Type: `"list"`. Computed.

Headers that can be used to identify mobile traffic.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2233231120301313-0022233230100201-2310310233120212-2302221232201233-0202000132312022-1011302212231111-0220132322331303-0303101122011222"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-012.md#canonical-3000331322123222-1013001323310133-0222120030131201-1323123303303112-1303213123010211-1102000112033002-3301311013020221-0131222223022312): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-012.md#canonical-1302310101331121-0300023113001221-3232202301213311-2103230110133232-1023232203312010-2120222230130110-1031122133213300-1023123220233203): complete subsection reference.

- [item](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213120033310011-0312001222203112-3100013112202022-2312213112012033-3232022113320013-1221200303311110-3123122330330323-1021332210211301): complete subsection reference.

<a id="canonical-3001303231321232-2100232100220120-1101011122232333-1033110301010203-2100301302311102-2020020112122331-0331322122211221-2202102212211001"></a>

<a id="canonical-2021123012313001-2123012032302312-3310022002302022-3212211002220122-1103311002230022-3032113030133002-0222113102010133-2021133000000231"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

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

<a id="canonical-3000331322123222-1013001323310133-0222120030131201-1323123303303112-1303213123010211-1102000112033002-3301311013020221-0131222223022312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-012.md#canonical-1020010112011222-1203033213323302-3102202132222212-1233320101300013-1001211221013333-1123121322230200-3301110301123030-2322013122221223)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-012.md#canonical-3023130302320000-0112122023101330-2000131211111302-1130211032012131-1113011010030113-2131210333211310-0200113031121103-3010021121130003)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-012.md#canonical-3200102213000012-0131210312123312-0102011032301030-2303202332302032-0112032133212221-1021213033121120-2011002211130022-1133130133321223)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-2131213231331321-3112132110232123-2333030231033300-0222300213120200-3211233213031013-3130200331332311-3212020312111101-1322333321213320"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302310101331121-0300023113001221-3232202301213311-2103230110133232-1023232203312010-2120222230130110-1031122133213300-1023123220233203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-012.md#canonical-1020010112011222-1203033213323302-3102202132222212-1233320101300013-1001211221013333-1123121322230200-3301110301123030-2322013122221223)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-012.md#canonical-3023130302320000-0112122023101330-2000131211111302-1130211032012131-1113011010030113-2131210333211310-0200113031121103-3010021121130003)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-012.md#canonical-3200102213000012-0131210312123312-0102011032301030-2303202332302032-0112032133212221-1021213033121120-2011002211130022-1133130133321223)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-0210323211222311-3011111302023021-0033220132210133-0003113023120321-2130223003030320-0132322102123113-2102121131003331-1303110112331301"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213120033310011-0312001222203112-3100013112202022-2312213112012033-3232022113320013-1221200303311110-3123122330330323-1021332210211301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.mobile_sdk_config](data-sources--http_loadbalancer--reference--group-012.md#canonical-1020010112011222-1203033213323302-3102202132222212-1233320101300013-1001211221013333-1123121322230200-3301110301123030-2322013122221223)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--http_loadbalancer--reference--group-012.md#canonical-3023130302320000-0112122023101330-2000131211111302-1130211032012131-1113011010030113-2131210333211310-0200113031121103-3010021121130003)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--http_loadbalancer--reference--group-012.md#canonical-3200102213000012-0131210312123312-0102011032301030-2303202332302032-0112032133212221-1021213033121120-2011002211130022-1133130133321223)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-1202032210212330-2330311203110301-1332213213133332-2132210000221013-3313032002011330-0303013303111213-2303002021020021-1220133231030033"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2020210310011100-2212023120020031-0122330013132103-0300301023300311-2331122203002011-1101122032131330-1122313311301031-3021110220232030"></a>

### Direct properties for `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item`

<a id="canonical-1112331232321021-3312022310000233-3013222111322130-2220322220222203-1201003120021130-3102302123012323-0231013011020003-0222020320002231"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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

<a id="canonical-3122231021011101-2321000103132001-2231111102110121-0213032010030113-1211130023233011-1233230303300221-3323222211332102-1022103323302331"></a>

<a id="canonical-1013030300012110-0033030122332131-3122200013311212-0203222310321330-3221323122203131-3121000010323000-1232201301212232-1122323002321302"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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

<a id="canonical-2230311121310333-2133000232202321-3333121213332113-0102321033213300-1032031120200233-1023122230223322-3213201213112221-3003013222100120"></a>

<a id="canonical-2230212102103200-3323321123301300-0113102202220031-0023011123330313-1101302213102122-0130001313103232-3331100221220323-1331023030232032"></a>

#### `bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- bot_defense.policy.protected_app_endpoints

<a id="canonical-1003331131130002-3002310202210022-0322113212213211-2330001331330210-2132321200131332-2223100112311022-2231330032332120-3002010200103102"></a>

Type: `"list"`. Computed.

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32
endpoints per LB' after 4 LBs.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2320011301122112-3100121312013212-2313001212111200-3001122230220312-3011331130322122-0311321233223103-2313010121230031-3001100331101122"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints`

- [allow_good_bots](data-sources--http_loadbalancer--reference--group-012.md#canonical-3213301230322202-1111023200223001-2213312210223202-1002322310223102-1210213133302331-2302302311311103-3031100212131232-2020113032203303): complete subsection reference.

- [any_domain](data-sources--http_loadbalancer--reference--group-012.md#canonical-2023201120102131-1212110300032333-0022003223331322-3320231000030112-3000231030302100-1210011332311232-2030300200202123-1010213120101320): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-012.md#canonical-0132113300231101-0023213122011022-1212331133022333-2321132131122302-2032320010222233-1201123123233130-3200303301121000-0020323121203320): complete subsection reference.

- [flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-013.md#canonical-1011120012122232-0132300313310233-3303103100123011-2133130110032023-0322303330000031-3213301213000300-3103103102220320-1012213013230200): complete subsection reference.

<a id="canonical-3031132322001231-3111001312002020-1021313202211100-2112011213311123-2332000013330222-0310332320312212-1232101112112000-2012102021301213"></a>

<a id="canonical-3100321002202222-0200313202331003-0232013233032210-1220033121123132-1303221330110030-2213300023122121-1233133120323211-2201313321021311"></a>

#### `bot_defense.policy.protected_app_endpoints.http_methods` property

Type: `["list", "string"]`. Computed.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-013.md#canonical-0231121312101122-3300332003101222-3123312033323213-2002220233221030-3033222311320011-1333102010130320-0310113302113312-3301101303233333): complete subsection reference.

- [mitigate_good_bots](data-sources--http_loadbalancer--reference--group-013.md#canonical-0220301223033033-1321302030220300-1012011022230102-0230322301113331-3323300111112022-3112123120333223-0103331300203001-1101213011100221): complete subsection reference.

- [mitigation](data-sources--http_loadbalancer--reference--group-013.md#canonical-3113121001322020-2023101302120333-2312320133003033-3131032033220132-1003311333020002-1031111020332022-0101120021023330-1032333020221200): complete subsection reference.

- [mobile](data-sources--http_loadbalancer--reference--group-013.md#canonical-1330122130032220-2130032113023113-1203321233113122-2322113130203210-3133301320022000-1000110032320221-3023202213002030-1320021232113122): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-013.md#canonical-2330222321121202-3232131223011232-0133333021120113-0021202023030201-2033130201031100-3300303203201202-1121231122031110-1301321011110232): complete subsection reference.

<a id="canonical-1030321311220313-2200121111300200-1302030133333323-1330010112000303-3302133100130010-2131032012333003-1100313321030033-0312302203001312"></a>

<a id="canonical-2002333201023203-2000203312233112-2222000220013311-2122021321211002-0100002212121012-3210021311022213-3033002131222232-3131231110213231"></a>

#### `bot_defense.policy.protected_app_endpoints.protocol` property

Type: `"string"`. Computed.

\[Enum: BOTH|HTTP|HTTPS\] SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for
HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.
Possible values are \`BOTH\`, \`HTTP\`, \`HTTPS\`. Defaults to \`BOTH\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "BOTH",
  "enum": [
    "BOTH",
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](data-sources--http_loadbalancer--reference--group-013.md#canonical-0132133021110010-0321123323101321-2211322032021100-0210233111121203-1233000110332322-0110320212010333-3330033122311332-3133010130210313): complete subsection reference.

- [undefined_flow_label](data-sources--http_loadbalancer--reference--group-013.md#canonical-0012333121330033-0013212031331012-1212213230302121-3332301031220311-1223123102033211-1232302330022321-2100320113010222-1013120311223200): complete subsection reference.

- [web](data-sources--http_loadbalancer--reference--group-013.md#canonical-2232300101302232-0203112120223220-2310123012320320-2312113333230332-3122220101033102-1132123323313023-2202231133300220-1222312113130330): complete subsection reference.

- [web_mobile](data-sources--http_loadbalancer--reference--group-013.md#canonical-3123231102203023-1030300113000102-1203011010223112-2023322130101302-3102220232033111-0303100013030102-2000120111321332-0012210012001203): complete subsection reference.

<a id="canonical-3213301230322202-1111023200223001-2213312210223202-1002322310223102-1210213133302331-2302302311311103-3031100212131232-2020113032203303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.allow_good_bots` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- bot_defense.policy.protected_app_endpoints.allow_good_bots

<a id="canonical-3302212103312122-3212231232120222-0033332230012221-1011311311210001-3110203103333201-1320213303012232-1202201101232332-3311210011203103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for allow good bots.

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

<a id="canonical-2023201120102131-1212110300032333-0022003223331322-3320231000030112-3000231030302100-1210011332311232-2030300200202123-1010213120101320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- bot_defense.policy.protected_app_endpoints.any_domain

<a id="canonical-0202003130120230-3221333101030001-3213210310020003-1202233020232032-1310310100300111-2323323112202110-1331103210030013-1212233300001113"></a>

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

<a id="canonical-0132113300231101-0023213122011022-1212331133022333-2321132131122302-2032320010222233-1201123123233130-3200303301121000-0020323121203320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- bot_defense.policy.protected_app_endpoints.domain

<a id="canonical-0331233130332113-0021020012212103-3231120133032011-1100201220300230-3101011201031303-1311321133230303-3233300322320121-0021311002112223"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-1130000303201213-1332130212202133-0330202012232333-2112013102032210-0320132121033311-1332302311132333-1331122031232112-3212323301333031"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.domain`

<a id="canonical-0310203332100020-3233202313102030-1202302203222132-0003122033102233-3233112321033311-1332330001300112-1120332233232113-1002222132022000"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3223310213012002-3020321200110333-2212012123033311-0303003113301303-1210030210230300-0100330210132230-0332232010212002-1131301222121030"></a>

<a id="canonical-3200311223103131-1010230033033333-3300230111133030-0302233312023020-0220323101013013-0320003222302133-3323032123312033-2022322033110111"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1322123023330101-1201002031122201-2100330302202333-1323020230000300-3321122103200113-2331012232032011-0333213210312030-0302331013111112"></a>

<a id="canonical-0101200201120130-3121200113330121-2313210211201210-1121331301320221-2210031132331221-3002202132322131-0030221222333123-2210312212320112"></a>

#### `bot_defense.policy.protected_app_endpoints.domain.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="canonical-2111031213010232-3111100233121230-0321213212332203-0223312111113223-3200220301103130-3313300021202013-1320213331121313-1231013132033021"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

<a id="canonical-0203013030221002-3303303332111002-3230111310203323-1213032021221310-0221130310232012-3210322130203311-2302233023002013-0132012211030011"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label`

- [account_management](data-sources--http_loadbalancer--reference--group-012.md#canonical-0222003330120203-0023033212212033-0001102031020013-2202103300233003-1320322202210012-0223300222003003-1221013010300201-0013211001201210): complete subsection reference.

- [authentication](data-sources--http_loadbalancer--reference--group-012.md#canonical-2033012130013012-3031300120120321-2131033131123132-3110220110332121-1233101233210213-2332130001223300-1303330323011233-2313201312010222): complete subsection reference.

- [financial_services](data-sources--http_loadbalancer--reference--group-012.md#canonical-1302200031100303-2002010011121010-1121011120131123-2313102202010222-1233121213022000-3230003011103203-0321300123220310-1133131311230332): complete subsection reference.

- [flight](data-sources--http_loadbalancer--reference--group-012.md#canonical-1022111001302011-0121311202121023-1130000303031323-2312133133112133-1302012103321310-0320213322001312-2102333332111230-1223203113322111): complete subsection reference.

- [profile_management](data-sources--http_loadbalancer--reference--group-012.md#canonical-3310221020203222-3010103000301002-3011303021002132-3212323101201031-0302022323211300-2333332003330033-2201133300123330-1113311213002131): complete subsection reference.

- [search](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112133023312102-1032003322110323-3131220313222003-3302313012123211-2102303331023320-2003223332333322-2123021130111101-2333021021212123): complete subsection reference.

- [shopping_gift_cards](data-sources--http_loadbalancer--reference--group-012.md#canonical-2330220103221230-1311132331101310-0111111002023331-2131030132332211-0120321320000201-0233122230033323-1103023322132110-3323203130022223): complete subsection reference.

<a id="canonical-0222003330120203-0023033212212033-0001102031020013-2202103300233003-1320322202210012-0223300222003003-1221013010300201-0013211001201210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="canonical-1302203203111122-2111200131122013-2320121110131012-1212010033032330-3030233030312003-2002120101300220-0001001133101121-1020112012021033"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Account Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

<a id="canonical-0101023003000232-1010032332130112-0020301032101321-2102211122303221-1121013012223010-0111011013313201-2131220101110023-1233031103300221"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.account_management`

- [create](data-sources--http_loadbalancer--reference--group-012.md#canonical-2123002122322123-3223213133033211-0330013021003132-2223313310133232-0003000200010111-1132320101332231-0320012112131212-1321020210002121): complete subsection reference.

- [password_reset](data-sources--http_loadbalancer--reference--group-012.md#canonical-2110330110320002-1022321303102001-2101022301312231-1003000103230012-1310232001332030-0100310200013303-0300331331002313-0102031030011210): complete subsection reference.

<a id="canonical-2123002122322123-3223213133033211-0330013021003132-2223313310133232-0003000200010111-1132320101332231-0320012112131212-1321020210002121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management.create` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--http_loadbalancer--reference--group-012.md#canonical-0222003330120203-0023033212212033-0001102031020013-2202103300233003-1320322202210012-0223300222003003-1221013010300201-0013211001201210)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

<a id="canonical-1032201033002020-1003022231111113-3001022321132103-3130133133003321-1113210322112002-2101013131001102-2023033231123031-1322223120322110"></a>

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

<a id="canonical-2110330110320002-1022321303102001-2101022301312231-1003000103230012-1310232001332030-0100310200013303-0300331331002313-0102031030011210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--http_loadbalancer--reference--group-012.md#canonical-0222003330120203-0023033212212033-0001102031020013-2202103300233003-1320322202210012-0223300222003003-1221013010300201-0013211001201210)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset

<a id="canonical-0001111320212222-3002222031203032-0103313303233022-3203203020102100-2203022202323101-2121212102322023-2032230120332200-1211202021233020"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for password reset.

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

<a id="canonical-2033012130013012-3031300120120321-2131033131123132-3110220110332121-1233101233210213-2332130001223300-1303330323011233-2313201312010222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="canonical-3000022113203303-0301010302020301-3332210202302223-1030023122203003-2013203330331312-1301300100330233-2032231103123312-0212210020103113"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Authentication Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

<a id="canonical-0021321223020100-2122310333203210-2203002011001002-1332213320103323-0100303103231303-3310212201203301-0230120101013123-1101021200330201"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication`

- [login](data-sources--http_loadbalancer--reference--group-012.md#canonical-3311213121201213-1131133010203223-0130032000302112-2130333333302122-2120102131021020-3033222103210211-1211103213101030-2113102112231012): complete subsection reference.

- [login_mfa](data-sources--http_loadbalancer--reference--group-012.md#canonical-3220031102133110-3003200222130202-2221212223001133-2331103222200321-0022011132133232-2120032320120221-0223332322021113-0031112231001012): complete subsection reference.

- [login_partner](data-sources--http_loadbalancer--reference--group-012.md#canonical-3010113221023303-2310221001211011-1310131202012121-1000220121330322-2100020201311213-2202203103322012-2201023131221130-1333100033302030): complete subsection reference.

- [logout](data-sources--http_loadbalancer--reference--group-012.md#canonical-3331311022101202-2132030323233211-3012332312312032-0011122322103121-2032130003111101-2310201211220122-0331210200133130-2303201121321321): complete subsection reference.

- [token_refresh](data-sources--http_loadbalancer--reference--group-012.md#canonical-3313020202121330-1201222233131120-2002203020333131-1322123301220122-0331010210033112-0322200310222130-3213131301220233-0123000030120333): complete subsection reference.

<a id="canonical-3311213121201213-1131133010203223-0130032000302112-2130333333302122-2120102131021020-3033222103210211-1211103213101030-2113102112231012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-012.md#canonical-2033012130013012-3031300120120321-2131033131123132-3110220110332121-1233101233210213-2332130001223300-1303330323011233-2313201312010222)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="canonical-3333030101103332-0013321003321013-3333232220010001-2120311020323133-0020330002203323-3132001210322222-3223222111103300-0032013022212020"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

<a id="canonical-1131223311113000-2121202130331013-0203022033221023-1013312322033223-3012301132201201-2203223112010023-3220321002231121-3120203133330222"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login`

- [disable_transaction_result](data-sources--http_loadbalancer--reference--group-012.md#canonical-0312023000002223-2001211301323312-2212322132322323-0103001310233321-3113231300120200-1120103232122332-3231201331332111-2331222123332021): complete subsection reference.

- [transaction_result](data-sources--http_loadbalancer--reference--group-012.md#canonical-1222303031131220-1321101200312031-0111113211303003-1320202131121302-3102302211331111-2200213202112130-3322211311110202-0011321102212010): complete subsection reference.

<a id="canonical-0312023000002223-2001211301323312-2212322132322323-0103001310233321-3113231300120200-1120103232122332-3231201331332111-2331222123332021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-012.md#canonical-2033012130013012-3031300120120321-2131033131123132-3110220110332121-1233101233210213-2332130001223300-1303330323011233-2313201312010222)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--http_loadbalancer--reference--group-012.md#canonical-3311213121201213-1131133010203223-0130032000302112-2130333333302122-2120102131021020-3033222103210211-1211103213101030-2113102112231012)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-0212032211230303-2203010232331102-2120300211111203-1331120032100112-0221300023113213-2032030101021333-1222300303010001-2223223303133310"></a>

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

<a id="canonical-1222303031131220-1321101200312031-0111113211303003-1320202131121302-3102302211331111-2200213202112130-3322211311110202-0011321102212010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-012.md#canonical-2033012130013012-3031300120120321-2131033131123132-3110220110332121-1233101233210213-2332130001223300-1303330323011233-2313201312010222)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--http_loadbalancer--reference--group-012.md#canonical-3311213121201213-1131133010203223-0130032000302112-2130333333302122-2120102131021020-3033222103210211-1211103213101030-2113102112231012)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-3021323102132303-1230331311023210-0330103112130223-1211211121212012-2303131220032201-1000210201133222-3121011211300120-3212020021013110"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

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

<a id="canonical-3013221102132202-2202021111003312-2333211301300233-3002122233232212-3111231133312033-2300033121131200-1120103220031212-0320003022012300"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result`

- [failure_conditions](data-sources--http_loadbalancer--reference--group-012.md#canonical-2202210222033323-2003200120122001-1323023232000323-3210212021010302-1200032302012111-0210032223321001-0022331011103311-3233331101311202): complete subsection reference.

- [success_conditions](data-sources--http_loadbalancer--reference--group-012.md#canonical-2332221212322310-3310122302113222-2231023103003320-2303331133030223-0032333023130333-0330323231031112-1122211123010133-3112120220132232): complete subsection reference.

<a id="canonical-2202210222033323-2003200120122001-1323023232000323-3210212021010302-1200032302012111-0210032223321001-0022331011103311-3233331101311202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-012.md#canonical-2033012130013012-3031300120120321-2131033131123132-3110220110332121-1233101233210213-2332130001223300-1303330323011233-2313201312010222)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--http_loadbalancer--reference--group-012.md#canonical-3311213121201213-1131133010203223-0130032000302112-2130333333302122-2120102131021020-3033222103210211-1211103213101030-2113102112231012)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--http_loadbalancer--reference--group-012.md#canonical-1222303031131220-1321101200312031-0111113211303003-1320202131121302-3102302211331111-2200213202112130-3322211311110202-0011321102212010)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-1312122230212012-0213301332310212-3113123110101211-1011322312133002-0113302031310310-0030032232120021-1230002210303131-3133001311301331"></a>

Type: `"list"`. Computed.

Failure Conditions. Failure Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1323302312233122-3012002030302203-2031330111330133-2012231203013113-3301012302103321-0023313023111223-0100310320033222-0311231121210322"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions`

<a id="canonical-3010231110113310-3320331132331333-0210010010312201-2023122031020133-3032122100212001-2001030013212022-3001312322301320-0012020220230330"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2222220300121133-3110201201222303-2113231202031130-3231221321111101-1330213213133311-0100100213010333-3022122012333122-1311103302220121"></a>

<a id="canonical-0313003100321022-2332011300233302-0001112112313200-3003103321003131-0121320123230111-2033102120103102-3033113121002002-2330312003022233"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2312031130122102-2010120111120130-0231113231230111-2230223200310010-3222213023130100-0011121201030320-1202220322020310-1122133021120003"></a>

<a id="canonical-1311132103101330-2232233202322202-1232232032321331-2002230011030022-2332011211213333-0302010312303232-1131331002031033-2131133311021033"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status` property

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Additional upstream details:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2332221212322310-3310122302113222-2231023103003320-2303331133030223-0032333023130333-0330323231031112-1122211123010133-3112120220132232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-012.md#canonical-2033012130013012-3031300120120321-2131033131123132-3110220110332121-1233101233210213-2332130001223300-1303330323011233-2313201312010222)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--http_loadbalancer--reference--group-012.md#canonical-3311213121201213-1131133010203223-0130032000302112-2130333333302122-2120102131021020-3033222103210211-1211103213101030-2113102112231012)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--http_loadbalancer--reference--group-012.md#canonical-1222303031131220-1321101200312031-0111113211303003-1320202131121302-3102302211331111-2200213202112130-3322211311110202-0011321102212010)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-0132312012332031-3310020200111023-1123113013321230-1303232121111131-1222213302212101-0303223000113312-1021210201132233-0320032032202332"></a>

Type: `"list"`. Computed.

Success Conditions. Success Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3313212201033032-2021312213001022-0202222231131130-3013032030213000-0320333011120011-1310032030131022-0011023310002012-2213320213211230"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions`

<a id="canonical-2201113322230110-2330000021013031-1012033230301223-3033030231220213-2201110221022213-1230133220131233-0311130020332203-3212023311133332"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1010202212311133-3223010002220303-0223223212011300-0010132203100123-0111301003113322-0002220111230213-1332001020301023-2000022101330300"></a>

<a id="canonical-2203113301112331-1032332222031202-2112032130311333-3012320102012203-0321201320033313-3122311120301231-2121221111132032-2111111233321120"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0320101031010011-2131133122101022-1231330220211211-3202322213032101-0102123011111002-2023003133201023-2101132033213212-1310301003332223"></a>

<a id="canonical-3232002202003130-0101110322202002-0021031310321000-0233111101223322-1222110303212020-2223311100211001-3331122323231111-3303300333230022"></a>

#### `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status` property

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Additional upstream details:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3220031102133110-3003200222130202-2221212223001133-2331103222200321-0022011132133232-2120032320120221-0223332322021113-0031112231001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-012.md#canonical-2033012130013012-3031300120120321-2131033131123132-3110220110332121-1233101233210213-2332130001223300-1303330323011233-2313201312010222)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa

<a id="canonical-0020121301232323-0002331330210302-2201310223123300-0120312202203212-0023230301130112-0222133120103321-3001011202112312-1312323012013122"></a>

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

<a id="canonical-3010113221023303-2310221001211011-1310131202012121-1000220121330322-2100020201311213-2202203103322012-2201023131221130-1333100033302030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-012.md#canonical-2033012130013012-3031300120120321-2131033131123132-3110220110332121-1233101233210213-2332130001223300-1303330323011233-2313201312010222)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner

<a id="canonical-1011202211032310-1311121231333310-1300212033301232-1201211301203023-3121131031100112-3032003230220113-2032111132120202-3322010010120103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for login partner.

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

<a id="canonical-3331311022101202-2132030323233211-3012332312312032-0011122322103121-2032130003111101-2310201211220122-0331210200133130-2303201121321321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-012.md#canonical-2033012130013012-3031300120120321-2131033131123132-3110220110332121-1233101233210213-2332130001223300-1303330323011233-2313201312010222)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout

<a id="canonical-0202011011113001-1211023102001210-2232100100322200-2333303001223010-2012031331331211-0302222013001212-3323332122130300-2100031010132110"></a>

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

<a id="canonical-3313020202121330-1201222233131120-2002203020333131-1322123301220122-0331010210033112-0322200310222130-3213131301220233-0123000030120333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--reference--group-012.md#canonical-2033012130013012-3031300120120321-2131033131123132-3110220110332121-1233101233210213-2332130001223300-1303330323011233-2313201312010222)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh

<a id="canonical-3002030221330002-3012333002310100-3221202300032030-3012003110121202-0210033222302331-0220110121033203-3130323320333301-3111002201031210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for token refresh.

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

<a id="canonical-1302200031100303-2002010011121010-1121011120131123-2313102202010222-1233121213022000-3230003011103203-0321300123220310-1133131311230332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services

<a id="canonical-0301332302301200-2021330220133221-3313132213132231-3112003303233120-0331133122002212-3011020121203020-3123121321322231-2323100332221212"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Financial Services Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

<a id="canonical-3011312101222231-2002213002013131-0022322030113023-3113130302021103-2220320220313310-0302330012013223-3001111323303200-3311212300012020"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.financial_services`

- [apply](data-sources--http_loadbalancer--reference--group-012.md#canonical-2233301233331222-0001020213113232-2222201332033313-3301330103021303-1103120113310320-2322332032110112-0300030212032032-3112311133102113): complete subsection reference.

- [money_transfer](data-sources--http_loadbalancer--reference--group-012.md#canonical-3000230230333001-1322021011303221-0332010332002200-0020131333121232-2300221131213310-2312022011223231-0012211333300120-3112300021103310): complete subsection reference.

<a id="canonical-2233301233331222-0001020213113232-2222201332033313-3301330103021303-1103120113310320-2322332032110112-0300030212032032-3112311133102113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--http_loadbalancer--reference--group-012.md#canonical-1302200031100303-2002010011121010-1121011120131123-2313102202010222-1233121213022000-3230003011103203-0321300123220310-1133131311230332)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply

<a id="canonical-2131122220303100-1123333233302113-2313133321033122-0221313103213202-2133211132022113-0300123021021112-2210321101231022-2233310120013012"></a>

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

<a id="canonical-3000230230333001-1322021011303221-0332010332002200-0020131333121232-2300221131213310-2312022011223231-0012211333300120-3112300021103310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--http_loadbalancer--reference--group-012.md#canonical-1302200031100303-2002010011121010-1121011120131123-2313102202010222-1233121213022000-3230003011103203-0321300123220310-1133131311230332)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-3203021313333000-0230132032211320-2301002220203212-0023230131232201-3012233310003103-3230303322131332-1332231122320020-1000010200033032"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for money transfer.

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

<a id="canonical-1022111001302011-0121311202121023-1130000303031323-2312133133112133-1302012103321310-0320213322001312-2102333332111230-1223203113322111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.flight` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- bot_defense.policy.protected_app_endpoints.flow_label.flight

<a id="canonical-2203112320202032-0220211221133023-0113123013111223-3112312111003031-3312311102013102-2001032031013132-3332203133001333-2231322310321112"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

<a id="canonical-0122110313200212-1223311321201321-2311110130033203-1230101003012001-2033123223022132-0013000300203331-1332311310133321-2210230011211231"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.flight`

- [checkin](data-sources--http_loadbalancer--reference--group-012.md#canonical-1222323103223332-3021331231022133-1032121130213211-1321022200010321-1012211012030300-0333200111021023-0122210113021320-1013332203222111): complete subsection reference.

<a id="canonical-1222323103223332-3021331231022133-1032121130213211-1321022200010321-1012211012030300-0333200111021023-0122210113021320-1013332203222111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](data-sources--http_loadbalancer--reference--group-012.md#canonical-1022111001302011-0121311202121023-1130000303031323-2312133133112133-1302012103321310-0320213322001312-2102333332111230-1223203113322111)
- bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin

<a id="canonical-2031200030121033-3113223332102320-2302303320002012-2323011222102230-0020200010003012-3030231032121101-0033331102213021-2221332113212011"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3310221020203222-3010103000301002-3011303021002132-3212323101201031-0302022323211300-2333332003330033-2201133300123330-1113311213002131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management

<a id="canonical-1213011310330120-2333111112010310-3333333331330220-1131012103022222-1212031030130212-2230130323110133-0102312002221310-3121102220202201"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Profile Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

<a id="canonical-0222333013330101-0000113321031322-2021212313233122-3032322133221303-2023203102030120-1202220200330032-0202130020120222-0110200121200301"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.profile_management`

- [create](data-sources--http_loadbalancer--reference--group-012.md#canonical-2000131103212102-2003232130320231-2213133211132332-2313201013020230-1323231000313102-1032030233301103-1132333223123131-0211100200033212): complete subsection reference.

- [update](data-sources--http_loadbalancer--reference--group-012.md#canonical-2230312010123022-3033333333010101-3001211030012113-2120121301023300-2013100231120033-2302111203211222-2303121030113303-2210303011330213): complete subsection reference.

- [view](data-sources--http_loadbalancer--reference--group-012.md#canonical-1112212002303123-3121202013103031-0303212203130311-1123333110222301-3313222022113011-0212210232322323-1102003132220320-3302121133232231): complete subsection reference.

<a id="canonical-2000131103212102-2003232130320231-2213133211132332-2313201013020230-1323231000313102-1032030233301103-1132333223123131-0211100200033212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--http_loadbalancer--reference--group-012.md#canonical-3310221020203222-3010103000301002-3011303021002132-3212323101201031-0302022323211300-2333332003330033-2201133300123330-1113311213002131)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create

<a id="canonical-1320213012113330-1330331311032102-2101322233301212-0201310203002021-1213130202110032-0323201313222212-2331330303012301-0302301133001320"></a>

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

<a id="canonical-2230312010123022-3033333333010101-3001211030012113-2120121301023300-2013100231120033-2302111203211222-2303121030113303-2210303011330213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--http_loadbalancer--reference--group-012.md#canonical-3310221020203222-3010103000301002-3011303021002132-3212323101201031-0302022323211300-2333332003330033-2201133300123330-1113311213002131)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update

<a id="canonical-3221310121101322-2201312121201203-3311023210000120-0212000200100233-2130233322132130-3101200301131020-3120132220201220-1031203103303010"></a>

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

<a id="canonical-1112212002303123-3121202013103031-0303212203130311-1123333110222301-3313222022113011-0212210232322323-1102003132220320-3302121133232231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--http_loadbalancer--reference--group-012.md#canonical-3310221020203222-3010103000301002-3011303021002132-3212323101201031-0302022323211300-2333332003330033-2201133300123330-1113311213002131)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view

<a id="canonical-0212221002310222-1122010332322202-2223121112023333-0230230203201201-3231302312110221-0310230202313112-2032212322331011-0232013330321111"></a>

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

<a id="canonical-3112133023312102-1032003322110323-3131220313222003-3302313012123211-2102303331023320-2003223332333322-2123021130111101-2333021021212123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- bot_defense.policy.protected_app_endpoints.flow_label.search

<a id="canonical-0131230223323233-2120331112033321-0300022132123110-2100221331131302-1232010032321121-1113230213323031-3331020130203032-2120300312201100"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```

<a id="canonical-2320203133201110-0222310131201330-3200120133302123-0303323313021100-2230120300031200-1023013233331010-0212002302030121-0211202033111121"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.search`

- [flight_search](data-sources--http_loadbalancer--reference--group-012.md#canonical-1001010212331200-1011001332300322-1333322031131313-0322123032131203-2100312320113330-2030130021231130-0211013313013103-2220322231113202): complete subsection reference.

- [product_search](data-sources--http_loadbalancer--reference--group-012.md#canonical-1020223211113001-1012011020102213-0330012031013023-1032322030010321-0210212122021120-0001331331310002-0111101301310120-1000303323201220): complete subsection reference.

- [reservation_search](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100320102313221-0033033013223312-0132221323130002-3321220203212223-1021130310132201-0020013002223122-2112200212022301-3221210232311212): complete subsection reference.

- [room_search](data-sources--http_loadbalancer--reference--group-012.md#canonical-3212313023200300-2320203003330023-2111013300012113-2210232132110220-1132202221330013-1201012021113001-1330213001212212-3123310010020102): complete subsection reference.

<a id="canonical-1001010212331200-1011001332300322-1333322031131313-0322123032131203-2100312320113330-2030130021231130-0211013313013103-2220322231113202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112133023312102-1032003322110323-3131220313222003-3302313012123211-2102303331023320-2003223332333322-2123021130111101-2333021021212123)
- bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search

<a id="canonical-3331231021021031-1130113213302022-2211202220320101-3213322012102132-0102311331002102-0131023313030100-2210132332202321-0331300231303000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for flight search.

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

<a id="canonical-1020223211113001-1012011020102213-0330012031013023-1032322030010321-0210212122021120-0001331331310002-0111101301310120-1000303323201220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.product_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112133023312102-1032003322110323-3131220313222003-3302313012123211-2102303331023320-2003223332333322-2123021130111101-2333021021212123)
- bot_defense.policy.protected_app_endpoints.flow_label.search.product_search

<a id="canonical-0021201223203222-2303002311100102-1313102220303230-1233300102332132-0210220012110012-0322200121130333-2223312102130030-1313333201110120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for product search.

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

<a id="canonical-3100320102313221-0033033013223312-0132221323130002-3321220203212223-1021130310132201-0020013002223122-2112200212022301-3221210232311212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112133023312102-1032003322110323-3131220313222003-3302313012123211-2102303331023320-2003223332333322-2123021130111101-2333021021212123)
- bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search

<a id="canonical-3230021020311320-2221001211323022-2211203123132301-3133131000013130-2233330212133221-1013013003002223-1231121000101102-2012231102000011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reservation search.

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

<a id="canonical-3212313023200300-2320203003330023-2111013300012113-2210232132110220-1132202221330013-1201012021113001-1330213001212212-3123310010020102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.room_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--reference--group-012.md#canonical-3112133023312102-1032003322110323-3131220313222003-3302313012123211-2102303331023320-2003223332333322-2123021130111101-2333021021212123)
- bot_defense.policy.protected_app_endpoints.flow_label.search.room_search

<a id="canonical-3101103121102330-3330201321112332-0101210010333302-1123101220311330-1310103113123102-2310023233311231-3132000113102111-0300333103312323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for room search.

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

<a id="canonical-2330220103221230-1311132331101310-0111111002023331-2131030132332211-0120321320000201-0233122230033323-1103023322132110-3323203130022223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

<a id="canonical-2133130313333322-2301110212202233-1111100101310201-1003130232210321-2320213322233221-2230310120313013-1333233202333130-0013022201231002"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"gift_card_make_purchase_with_gift_card\",\"gift_card_validation\",\"shop_add_to_cart\",\"shop_checkout\",\"shop_choose_seat\",\"shop_enter_drawing_submission\",\"shop_make_payment\",\"shop_order\",\"shop_price_inquiry\",\"shop_promo_code_validation\",\"shop_purchase_gift_card\",\"shop_update_quantity\"]"
}
```

<a id="canonical-3230103002000003-2023000321121223-2100323200331210-1231310101221013-0110010122330210-0210100112200333-1310302001103130-3130332332321332"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards`

- [gift_card_make_purchase_with_gift_card](data-sources--http_loadbalancer--reference--group-012.md#canonical-0302332022202330-1312233023333211-0120111112230300-0213113123212113-0210030231200221-0200200032203132-0311021323111222-0032002121231021): complete subsection reference.

- [gift_card_validation](data-sources--http_loadbalancer--reference--group-012.md#canonical-0122333320122032-3313112331030300-3020101322223223-0101001031100221-2223313212330223-1313201121331131-1310222202213101-2230220222122022): complete subsection reference.

- [shop_add_to_cart](data-sources--http_loadbalancer--reference--group-012.md#canonical-0333312123113300-3111232002220331-0303333120303103-2020132310111111-3112202002230230-0123022033003303-1320213131322201-3331112223213233): complete subsection reference.

- [shop_checkout](data-sources--http_loadbalancer--reference--group-012.md#canonical-2321323211201122-0220121222131300-1002212213002331-1313223130133013-3002010011033020-2311211032331133-3311231302302032-0323002332021212): complete subsection reference.

- [shop_choose_seat](data-sources--http_loadbalancer--reference--group-012.md#canonical-2000033032000012-2110221311021220-2300000301320323-3133113322112113-0101030001233021-2212100013220300-0222021130100003-2132320012323223): complete subsection reference.

- [shop_enter_drawing_submission](data-sources--http_loadbalancer--reference--group-012.md#canonical-2313313111103231-1222302303033101-0333013033222330-0331223201000332-0200233131012201-3030113302002100-3211222111200022-0303221321232310): complete subsection reference.

- [shop_make_payment](data-sources--http_loadbalancer--reference--group-012.md#canonical-3220313221002212-3211220222022322-0202101213331111-2123322322122002-1003122123213333-2311130101120001-2101102221322302-0210302132231002): complete subsection reference.

- [shop_order](data-sources--http_loadbalancer--reference--group-012.md#canonical-0100011122022222-1121112312010321-1301103013310023-2203021112333302-2003120101331332-3121202201013123-1302100001231223-3010333031021022): complete subsection reference.

- [shop_price_inquiry](data-sources--http_loadbalancer--reference--group-012.md#canonical-3321222310011311-2310302030223113-0223120002321122-0112220232323023-1013101102112123-1311223301223111-0222231022120022-3022201210331022): complete subsection reference.

- [shop_promo_code_validation](data-sources--http_loadbalancer--reference--group-013.md#canonical-2002122001221013-3233123211020200-3032033030122232-2130020032311013-3002113001103133-3033113100030111-2310122010213220-1312320022220110): complete subsection reference.

- [shop_purchase_gift_card](data-sources--http_loadbalancer--reference--group-013.md#canonical-3213303331203201-3012200133203011-3303131220010212-1011112223110010-1333102103211012-0210031231022313-3332203212202012-1110200310320331): complete subsection reference.

- [shop_update_quantity](data-sources--http_loadbalancer--reference--group-013.md#canonical-3030332203100320-0233123131033003-1331310100112332-2231110010230121-1012303122212310-2201233023031230-3301021033203210-1130110232100123): complete subsection reference.

<a id="canonical-0302332022202330-1312233023333211-0120111112230300-0213113123212113-0210030231200221-0200200032203132-0311021323111222-0032002121231021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--http_loadbalancer--reference--group-012.md#canonical-2330220103221230-1311132331101310-0111111002023331-2131030132332211-0120321320000201-0233122230033323-1103023322132110-3323203130022223)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-0200030230232010-3103313221132320-0330222202222131-0212103301213222-1131131113110003-0202303310013012-3300232302311021-3232233130300203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for gift card make purchase with gift card.

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

<a id="canonical-0122333320122032-3313112331030300-3020101322223223-0101001031100221-2223313212330223-1313201121331131-1310222202213101-2230220222122022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--http_loadbalancer--reference--group-012.md#canonical-2330220103221230-1311132331101310-0111111002023331-2131030132332211-0120321320000201-0233122230033323-1103023322132110-3323203130022223)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-2010330320231302-2331322011032012-2032021023132302-3212221220101020-2021223301313332-0101330311303311-1103023121222302-2321223100232132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for gift card validation.

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

<a id="canonical-0333312123113300-3111232002220331-0303333120303103-2020132310111111-3112202002230230-0123022033003303-1320213131322201-3331112223213233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--http_loadbalancer--reference--group-012.md#canonical-2330220103221230-1311132331101310-0111111002023331-2131030132332211-0120321320000201-0233122230033323-1103023322132110-3323203130022223)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-0210322300203213-2112003003213220-0120123311120320-1331112233323030-2000031201231202-2101313121223332-0030331230323122-0312200030233102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop add to cart.

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

<a id="canonical-2321323211201122-0220121222131300-1002212213002331-1313223130133013-3002010011033020-2311211032331133-3311231302302032-0323002332021212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--http_loadbalancer--reference--group-012.md#canonical-2330220103221230-1311132331101310-0111111002023331-2131030132332211-0120321320000201-0233122230033323-1103023322132110-3323203130022223)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-0313233122221011-0311333331020021-3031220132220201-0021120322222302-0311103033333201-3313110033232020-2321131123333133-0302130220220200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop checkout.

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

<a id="canonical-2000033032000012-2110221311021220-2300000301320323-3133113322112113-0101030001233021-2212100013220300-0222021130100003-2132320012323223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--http_loadbalancer--reference--group-012.md#canonical-2330220103221230-1311132331101310-0111111002023331-2131030132332211-0120321320000201-0233122230033323-1103023322132110-3323203130022223)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat

<a id="canonical-0110222030031313-3322131322020300-1001032121300230-1221011312131111-2210011113033233-3310133032213000-2333122200211123-3312323100020101"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop choose seat.

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

<a id="canonical-2313313111103231-1222302303033101-0333013033222330-0331223201000332-0200233131012201-3030113302002100-3211222111200022-0303221321232310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--http_loadbalancer--reference--group-012.md#canonical-2330220103221230-1311132331101310-0111111002023331-2131030132332211-0120321320000201-0233122230033323-1103023322132110-3323203130022223)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission

<a id="canonical-3033321100331220-2023002231023320-1221331032001033-2011001030211220-3232023002130110-3212220332212231-1103131020323313-2232201220112223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop enter drawing submission.

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

<a id="canonical-3220313221002212-3211220222022322-0202101213331111-2123322322122002-1003122123213333-2311130101120001-2101102221322302-0210302132231002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--http_loadbalancer--reference--group-012.md#canonical-2330220103221230-1311132331101310-0111111002023331-2131030132332211-0120321320000201-0233122230033323-1103023322132110-3323203130022223)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment

<a id="canonical-1013201310110201-2012323111212100-0231123001212300-0002011230201032-1123031323120032-3333031300130220-0000320113211220-0330111031221000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop make payment.

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

<a id="canonical-0100011122022222-1121112312010321-1301103013310023-2203021112333302-2003120101331332-3121202201013123-1302100001231223-3010333031021022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--http_loadbalancer--reference--group-012.md#canonical-2330220103221230-1311132331101310-0111111002023331-2131030132332211-0120321320000201-0233122230033323-1103023322132110-3323203130022223)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order

<a id="canonical-0310031103221110-3233312230032121-0333222222233021-1022022220120333-0100120100111322-2301230100230013-1121023102112012-2300011210122303"></a>

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

<a id="canonical-3321222310011311-2310302030223113-0223120002321122-0112220232323023-1013101102112123-1311223301223111-0222231022120022-3022201210331022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense](data-sources--http_loadbalancer--reference--group-011.md#canonical-2121310311020030-0310032123301133-3330001203321213-0221223012112322-0233021322233301-3330113230111321-0230133323233300-0130013023302100)
- [bot_defense.policy](data-sources--http_loadbalancer--reference--group-011.md#canonical-1002101132312223-3030322032232111-3012200010011330-2232130302103003-0233031122310200-1233122103231212-2012321101330030-1212221333201322)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--reference--group-012.md#canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--reference--group-012.md#canonical-3100102132210230-1101303330011031-3020023132130022-2332100213001033-1000102321331303-2023311301213131-3033013200023232-1030122012202112)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--http_loadbalancer--reference--group-012.md#canonical-2330220103221230-1311132331101310-0111111002023331-2131030132332211-0120321320000201-0233122230033323-1103023322132110-3323203130022223)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry

<a id="canonical-3113102132223211-0213201000101201-0132013130100011-0311220101200310-1320020112322230-2212031032333323-1300033001112302-1200031122112230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop price inquiry.

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
