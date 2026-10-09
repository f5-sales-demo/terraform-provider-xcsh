---
page_title: "xcsh_service_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule reference."
---

# xcsh_service_policy_rule reference

<a id="canonical-0012001101012313-2101111330322311-2001020120021103-2102102331311203-2222313122301232-2123002201333331-1010330323223020-1032023303003112"></a>

## `request_constraints.max_request_line_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_request\_line\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_line_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0000013323113031-2033013133320032-2010233210331233-2202002001102333-3222323012130033-2132022123121011-0231320022032213-0133013002000130): complete subsection reference.

<a id="canonical-3011003132011022-0130310200023122-2020210310200003-3111113130013010-1000222210303013-2302222331212011-1323123120232121-2303101010230003"></a>

<a id="canonical-2233111203233002-3113012332313333-1203012113132300-3321313223322101-3330203201123233-2200123022100301-1133230223211020-3222111220133222"></a>

## `request_constraints.max_request_size_exceeds` property

Type: `"number"`. Optional.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0211021211211232-0231130103220023-2211300110200302-1330123120002032-1233331131120002-0331131010010032-1001003322020101-2232031202032200): complete subsection reference.

<a id="canonical-3002332102130132-1023102213020233-3301321313011303-3130311311302021-1211031221320201-3312013130003120-3001020230310222-0111101030121022"></a>

<a id="canonical-0102031031003102-1311332313300303-2312322332212321-3031133102313113-3322320131010310-2020001220123223-3022130201030113-0000123320012120"></a>

## `request_constraints.max_url_size_exceeds` property

Type: `"number"`. Optional.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 128000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  }
}
```

- [max_url_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2330331322221000-2213212201103011-1120230100211311-3201011320111220-1123011221100112-1002130230011321-0213111023121131-3303130013023331): complete subsection reference.

<a id="canonical-1122022011000132-2310232120103311-1302322200300230-2231133020012123-0101113203221130-3103331101010130-0200032210220122-2333200330022310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_cookie_count_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_cookie_count_none

<a id="canonical-3330003322002303-3112011103313310-2310133300001032-0301112113002100-2231301330310022-1113310301123111-1211300233030201-0231330331021111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie count none.

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
max_cookie_count_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210311330313210-1030032003013232-2310012332113123-0200213031011110-3032031210102331-2131213220101131-3321201211111213-3111123300123122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_cookie_key_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_cookie_key_size_none

<a id="canonical-0122132010120230-1020232110103112-3221012123023201-1301101202303202-0301310003310113-3210303313321301-0313331323231233-1323000201213320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie key size none.

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
max_cookie_key_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210332330211102-3323302021220120-0333020011333130-3231233221330020-0311010101023031-1213120210302130-3032020321123303-3220013033000320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_cookie_value_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_cookie_value_size_none

<a id="canonical-0130202311032031-3021022220311213-2202111002300213-3112133122112130-3131100320201132-1230023003103202-2022301131222121-1130033202332102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie value size none.

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
max_cookie_value_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031333020012231-3100132320333220-1321002111211212-1121322112320113-1033112301122111-3221032333001212-3332000110133331-1031013122313000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_header_count_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_header_count_none

<a id="canonical-0231113303012011-0211220222302012-1200101001132312-3010010221132210-3233332300122222-0322021112300010-0300203121103023-3012210210021313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header count none.

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
max_header_count_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132101021313102-3332033311133003-2210133001000322-0010232331203032-2322133101011321-1202032320332312-3132033012032313-3332110002331110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_header_key_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_header_key_size_none

<a id="canonical-0203311333010103-1213002103233110-3332002301220100-0103103332020023-0001232020213231-0132101001323123-0021202133022020-3002032121223032"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header key size none.

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
max_header_key_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102120230210231-2100101103110120-1123203011201001-3031333211322330-1210231111011330-2310102033001102-1320213011013311-0021120132333103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_header_value_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_header_value_size_none

<a id="canonical-3332010032003122-2002310031220112-0200323132113003-0232210311322332-1102020313301231-0100201222233130-1113113121320301-3123210320132232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header value size none.

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
max_header_value_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221300301123333-0233121001221021-1311302003000011-2010203103033011-3100131213301022-0011232323122220-0300231113313310-1313310111333232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_parameter_count_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_parameter_count_none

<a id="canonical-1110333230123311-0312322321312122-1330211131123201-0010320132111230-0111232322111312-0021010103323333-1112221012202133-1303112320333201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max parameter count none.

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
max_parameter_count_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230202133131132-2331002310102321-2203332230123312-3120102201131031-2000222113110033-2122102213110001-3312010333210032-0032213011111222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_parameter_name_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_parameter_name_size_none

<a id="canonical-1223322223300233-0130113021331330-2323201312220132-1322022321310110-0202232103120301-1012123212220203-1011023221322021-3122033220301211"></a>

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
max_parameter_name_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023121102232230-3110132110130012-2003011223012323-0330131103322022-1233311030100300-0113123310232012-3233100301000331-3200002330112200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_parameter_value_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_parameter_value_size_none

<a id="canonical-1323313230133220-0222102003010130-2130000311203101-2221202332011303-2303321031302101-3311022303220021-0120321203333320-1221213122130020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max parameter value size none.

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
max_parameter_value_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100103003131210-2220030203003130-1213013113101012-2123313132020002-3211122013110230-0202020232331211-3111300202321122-1301331102130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_query_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_query_size_none

<a id="canonical-1203232220211220-0211100200300303-0022133022033311-1200031312211111-2322013231313110-0230232003011333-1100221300223112-3323222232010110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max query size none.

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
max_query_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000013323113031-2033013133320032-2010233210331233-2202002001102333-3222323012130033-2132022123121011-0231320022032213-0133013002000130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_request_line_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_request_line_size_none

<a id="canonical-3033011001110323-3122020112332330-0230210321000010-1213233000333111-0210010303213123-2203220130031212-0021101202312223-1002021313203221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max request line size none.

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
max_request_line_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211021211211232-0231130103220023-2211300110200302-1330123120002032-1233331131120002-0331131010010032-1001003322020101-2232031202032200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_request_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_request_size_none

<a id="canonical-1110021232200100-3021311022230001-3331231112100033-1332110300222123-0222022201021032-2032300302303102-3200320111320133-3030022013230313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max request size none.

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
max_request_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330331322221000-2213212201103011-1120230100211311-3201011320111220-1123011221100112-1002130230011321-0213111023121131-3303130013023331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_url_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012)
- request_constraints.max_url_size_none

<a id="canonical-3212210122123013-0021120011333112-2333023110102022-1033202113011330-1332300213030323-1122132301332232-2333002000233000-2323333110322112"></a>

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
max_url_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_policy` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- segment_policy

<a id="canonical-0000310300000020-3030312333212300-1033312300110000-2012001210311021-3333123312322130-3100112131120320-3322132313111200-1332233120101131"></a>

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

<a id="canonical-1102212311203021-2320303011122133-2302233322132200-1220230233322203-3321121233032001-2023101331322102-1121121231331030-3213012032111231"></a>

### Direct properties for `segment_policy`

- [dst_any](resources--service_policy_rule--reference--group-002.md#canonical-2113201130221021-1032210031323120-3302033001311130-3010031030233213-2223213113211003-0220220301230130-0113013133110302-0033310222310123): complete subsection reference.

- [dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-2022332223212133-1222113102221310-0111310321311112-0032220010012112-0311310031133322-1221210033230122-3132203110031011-1312332032002013): complete subsection reference.

- [intra_segment](resources--service_policy_rule--reference--group-002.md#canonical-2323132121321022-3120112231121021-1221330022210220-2023333321111201-0123312221120313-2323021022122033-0103131123103200-2201010312333323): complete subsection reference.

- [src_any](resources--service_policy_rule--reference--group-002.md#canonical-1021033100031322-0210332232032110-0310012210003102-3310330313201232-1013122333100201-1010021111102100-0221022030003001-1011013331111111): complete subsection reference.

- [src_segments](resources--service_policy_rule--reference--group-002.md#canonical-0101332320113122-2200203113113133-1130203002313322-0202221012003110-2321113003133011-3130122110120132-1322213110010211-0010122011230310): complete subsection reference.

<a id="canonical-2113201130221021-1032210031323120-3302033001311130-3010031030233213-2223213113211003-0220220301230130-0113013133110302-0033310222310123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_policy.dst_any` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- segment_policy.dst_any

<a id="canonical-2332321232122233-1303323222111121-1220010021332231-2031211320212330-2001012101311130-0310102211102210-0322223312220030-2221230311322313"></a>

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

<a id="canonical-2022332223212133-1222113102221310-0111310321311112-0032220010012112-0311310031133322-1221210033230122-3132203110031011-1312332032002013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_policy.dst_segments` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- segment_policy.dst_segments

<a id="canonical-3113011300132303-2321300122002120-2020323213211011-1031022001212230-2122011332330022-3111201320100133-3002320222223022-0201022330113323"></a>

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

<a id="canonical-0111301333333310-2323130002200130-1002122133222231-2013033112233200-2303322123200022-0021321322021301-2323322333330123-1102311303233213"></a>

### Direct properties for `segment_policy.dst_segments`

- [segments](resources--service_policy_rule--reference--group-002.md#canonical-0102102011212003-2000031103020230-2021320331231313-1313122001333300-3002001311202123-3202320310031110-1202101220012321-0033220113231320): complete subsection reference.

<a id="canonical-0102102011212003-2000031103020230-2021320331231313-1313122001333300-3002001311202123-3202320310031110-1202101220012321-0033220113231320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_policy.dst_segments.segments` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- [segment_policy.dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-2022332223212133-1222113102221310-0111310321311112-0032220010012112-0311310031133322-1221210033230122-3132203110031011-1312332032002013)
- segment_policy.dst_segments.segments

<a id="canonical-2202013210133013-0011231011230330-3111001023132023-1000201001113100-1231320031223030-1101010200302311-2021021111110302-0130312101001000"></a>

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

<a id="canonical-0130003331032303-0310232303023220-3021111131120022-0312202300222330-1013100221233112-2023231232030331-2323213212313003-1020322330331331"></a>

### Direct properties for `segment_policy.dst_segments.segments`

<a id="canonical-0031001033010112-3333100222122113-2333202111213332-1332030103113201-1113013303003120-0300211113000120-3223302103102220-0313301033231133"></a>

#### `segment_policy.dst_segments.segments.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0213322130102120-3323021203320203-1220022200203100-1011323023311213-0332212223030110-0131231122113313-0110220312132221-3301321231031301"></a>

<a id="canonical-3103321030302111-0231100110321332-0201333332211212-2223333133201200-0110031323132033-2331102103000112-3323202020300110-3320013230112110"></a>

#### `segment_policy.dst_segments.segments.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2221200313302300-2313130133003030-0001223103030203-0101131201221012-0011032323103103-1132112021123020-2231333133322031-2112333223333000"></a>

<a id="canonical-3313120030101312-1330012330113231-2112333301223103-0120030001101310-1210100313103122-1322302301000002-2303002302313330-2303310223000132"></a>

#### `segment_policy.dst_segments.segments.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2323132121321022-3120112231121021-1221330022210220-2023333321111201-0123312221120313-2323021022122033-0103131123103200-2201010312333323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_policy.intra_segment` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- segment_policy.intra_segment

<a id="canonical-0230132313131123-0123032033333211-1203232110211110-3133123130013103-0333333132312203-3313310313311102-0133021021001210-3220000033031211"></a>

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

<a id="canonical-1021033100031322-0210332232032110-0310012210003102-3310330313201232-1013122333100201-1010021111102100-0221022030003001-1011013331111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_policy.src_any` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- segment_policy.src_any

<a id="canonical-2133300232030230-3302313230113313-2230112133321120-1332211003133132-3132203311230011-0321232132203331-2303133222030121-3223131223223100"></a>

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

<a id="canonical-0101332320113122-2200203113113133-1130203002313322-0202221012003110-2321113003133011-3130122110120132-1322213110010211-0010122011230310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_policy.src_segments` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- segment_policy.src_segments

<a id="canonical-1010311123002311-3011223222303031-0010001002300303-1012200110130003-3312203003320212-0132220003222002-3010330210130202-2023310332233331"></a>

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

<a id="canonical-0120210303010322-2212022123012031-1022311020030101-2003023320312012-0111333230132013-1231331203221323-0123302220221000-1223220221011233"></a>

### Direct properties for `segment_policy.src_segments`

- [segments](resources--service_policy_rule--reference--group-002.md#canonical-0222002030202103-3002001110223120-1111302113013331-2122323332211302-0131202200010210-1103103302211023-3232001301123333-3132310231302323): complete subsection reference.

<a id="canonical-0222002030202103-3002001110223120-1111302113013331-2122323332211302-0131202200010210-1103103302211023-3232001301123333-3132310231302323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_policy.src_segments.segments` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221)
- [segment_policy.src_segments](resources--service_policy_rule--reference--group-002.md#canonical-0101332320113122-2200203113113133-1130203002313322-0202221012003110-2321113003133011-3130122110120132-1322213110010211-0010122011230310)
- segment_policy.src_segments.segments

<a id="canonical-3033220131033301-0221220332120201-0230121221221110-1300033010331320-3231120233003320-1233310312232321-1213032210033323-0332010221113213"></a>

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

<a id="canonical-2311121331223301-1303232210013222-3122221100311312-1301120133031230-0030333001103111-3333210313102222-3013233021012111-3232110122223233"></a>

### Direct properties for `segment_policy.src_segments.segments`

<a id="canonical-1021300030021222-1320103122013202-2022123003021031-0220021310110031-1312020031222001-0013203312301302-0332031320013132-0031333323032213"></a>

#### `segment_policy.src_segments.segments.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2313332313212032-1010011002013002-0003111333002113-0330032321023103-1321120322322300-0033020202201312-2331230000232031-1203010320231331"></a>

<a id="canonical-0312013220031000-2311122001020212-1112322202110121-3300310320110102-3231323203233213-0213333001210332-2100033322130113-3132303020232113"></a>

#### `segment_policy.src_segments.segments.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0301231033320131-0303133320301200-0021113210022300-0303123011102101-3022032301321332-1221021321211323-1330103010102103-1111212220302321"></a>

<a id="canonical-0133220133230331-2200003230212211-1022032311210110-3031003230232122-0102321131303212-3301103020122101-0232133320221200-3300121100211212"></a>

#### `segment_policy.src_segments.segments.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1021310131303231-3111121233212022-1000112321100121-1023301303021202-3313011013101213-2203330203232131-3231103133122222-2233231203321311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- timeouts

<a id="canonical-3220311022103123-2030031323233001-0321211311322230-1333222313202203-0233320303323213-3133330130021203-3233123132133020-0320211203311000"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210022313331330-3322323331201030-1303303103012300-2233022210312011-0330302132321011-2332331133030311-0300333030031013-1101003111211021"></a>

### Direct properties for `timeouts`

<a id="canonical-0222333130210210-2122032222313301-0321310123301310-3231302222101003-1030221121110102-2131020211020112-3112300131113212-2030203130111300"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3310021130020030-3212212120032333-1100120210100100-3313113223313122-0222013210030010-0211003202300320-1211010013102103-1222023100022233"></a>

<a id="canonical-0323312321110010-1010022101001021-0120232123030010-0322100102202331-0013313110213320-1100011031113313-0130002101101131-1010030001301303"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3013131233001012-2230120010033112-3331333322011332-3031030002030231-1222232132220233-0023310310001220-0212211202200112-2023303010113101"></a>

<a id="canonical-1102302231122322-3331033133301031-1230313200031220-2221033200333103-1211130003032233-2131303331210203-0210021131131313-3021310033333103"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1120022113001331-2130320113031203-2010200322130313-2302311201332123-2312301110123333-1123322112031032-1122000320011102-2330131200103321"></a>

<a id="canonical-2010122231022000-3010103101033213-0032310030133001-2332323111010221-1202000301012100-1302001111223212-1131010110322003-1131321301300313"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2200000321123222-1121231032320123-3112320132222010-3101030202331210-2103131131000030-1333130312003211-1330202233211000-0302103210011223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- tls_fingerprint_matcher

<a id="canonical-2300023332310032-3012332003021312-3311001110322002-3132320121320002-3212132123303111-2220301013203202-0001020200302013-3212122333321223"></a>

Type: `"object"`. single nested block, Optional.

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013200113333233-1232302213301102-2300303020121311-2202020222010210-2321120333013131-2131120223222012-2303131103322120-0320133203020112"></a>

### Direct properties for `tls_fingerprint_matcher`

<a id="canonical-2213113001321332-2230010132132201-3131200121012322-3110112223120201-1022203210311023-1112312302231023-1032132201320101-1032321103021011"></a>

#### `tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0311230130200231-1033100312020100-1322132333033120-0201003033310310-1023101322130221-0331223123301332-1332011031310202-1311100322001031"></a>

<a id="canonical-3310200013222202-0003221022223113-1101021230003111-1032333303023333-2100300300023013-3100322201120100-2331323311330201-0031030022200310"></a>

#### `tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3120011321131203-0012123131322323-1101203311102121-3213221221210010-0302310333001031-1102100220102211-0000021211231302-1033300200332323"></a>

<a id="canonical-0031311102201111-0030331032133003-2231111121202310-2002020312213110-2033102312100013-0233333313230030-3003323103233233-0221220313120031"></a>

#### `tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_action` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- waf_action

<a id="canonical-0112112222221031-1331310203113201-2000213313323030-0130211021321221-2010210010030031-3223302012121302-1312233021121013-0000331203202102"></a>

Type: `"object"`. single nested block, Optional.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "none"),
  validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingObjectAttributes("none",
    "waf_skip_processing")}
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
  "x-ves-oneof-field-action_type": "[\"app_firewall_detection_control\",\"none\",\"waf_skip_processing\"]"
}
```

Terraform syntax:

```terraform
waf_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203011233200203-1310000123233023-1322133103312130-0011013101200312-2321121332112013-3113223200000003-2311102213211333-0230221212312300"></a>

### Direct properties for `waf_action`

- [app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100): complete subsection reference.

- [none](resources--service_policy_rule--reference--group-002.md#canonical-0122121012021132-1000022311102111-3131301323221132-0220303102210311-3223101332203121-2201303111022201-2320030331311003-3211100200333212): complete subsection reference.

- [waf_skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-2300310201333330-3133013011222002-2102212230202020-3113323220303320-0031002131200313-1213322102302103-0221203123001033-2303330302222312): complete subsection reference.

<a id="canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_action.app_firewall_detection_control` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- waf_action.app_firewall_detection_control

<a id="canonical-2302020223330013-0320000111031132-0032321313011101-1220221302020001-1023310022311120-3220322131302302-2112320233110232-1201003221033023"></a>

Type: `"object"`. single nested block, Optional.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223330200331303-0332201102300301-3322222123303010-3233030313032211-2302113220301111-1111213012211323-0300001213032321-3133313033222020"></a>

### Direct properties for `waf_action.app_firewall_detection_control`

- [exclude_attack_type_contexts](resources--service_policy_rule--reference--group-002.md#canonical-0221033123313311-2131133301001003-2221110233020201-2110000023011300-0302133300313312-2320022333022111-1133323323131212-3132220121231210): complete subsection reference.

- [exclude_bot_name_contexts](resources--service_policy_rule--reference--group-002.md#canonical-0321130232020121-0010023111030323-3023100322011320-3222001110203201-1212221212323003-0013202010231232-0131203313132331-2023221130323302): complete subsection reference.

- [exclude_signature_contexts](resources--service_policy_rule--reference--group-002.md#canonical-2023301201222231-0311313312020102-3232203322212121-1323122133301123-3321220330323000-3202033033333020-2121210022301213-3021102030120201): complete subsection reference.

- [exclude_violation_contexts](resources--service_policy_rule--reference--group-002.md#canonical-2323023101030330-0311033202012133-3120010032230210-2311100301223311-0221003311113000-2301222003331101-0133202202221320-0222121311320033): complete subsection reference.

<a id="canonical-0221033123313311-2131133301001003-2221110233020201-2110000023011300-0302133300313312-2320022333022111-1133323323131212-3132220121231210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_action.app_firewall_detection_control.exclude_attack_type_contexts` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- waf_action.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-0001100132222033-3233203111303132-1302101100302133-0001120212102001-0013332011013322-3332101122303200-2221332012000110-1332301022301122"></a>

Type: `"object"`. list nested block, Optional.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_attack_type_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003002230330111-2012220333233130-0100211032221223-2030033233131201-3301320212012001-3030332333122223-1302110110023122-3231332233003303"></a>

### Direct properties for `waf_action.app_firewall_detection_control.exclude_attack_type_contexts`

<a id="canonical-1112112313312322-3303313323212300-1000023203011031-2113002332231213-3101113000020001-2033311222112200-1300131112311013-3331021112301001"></a>

#### `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` property

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1002102131033021-1002323121230322-3221012131011021-2310003221323200-2122221201033021-0021231103230130-0330311130120232-2110300102211033"></a>

<a id="canonical-0232003221121303-2203223132322311-1212200232132020-2022133222002001-2311313122131111-1030332101331113-1312031233113211-3112020221303230"></a>

#### `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` property

Type: `"string"`. Optional.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1013333333330111-3301332333313001-0332010122001102-2333102303321222-2120201302231111-1302302312201221-2232022023200030-1001113202120131"></a>

<a id="canonical-1002313022100210-0010020102002021-2301122200330102-0333101132203002-0113213222311303-0132131331300131-2003332030320021-3031301331312203"></a>

#### `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` property

Type: `"string"`. Optional.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Additional upstream details:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY","ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS","ATTACK_TYPE_BUFFER_OVERFLOW","ATTACK_TYPE_COMMAND_EXECUTION","ATTACK_TYPE_CROSS_SITE_SCRIPTING","ATTACK_TYPE_DENIAL_OF_SERVICE","ATTACK_TYPE_DETECTION_EVASION","ATTACK_TYPE_DIRECTORY_INDEXING","ATTACK_TYPE_FORCEFUL_BROWSING","ATTACK_TYPE_GRAPHQL_PARSER_ATTACK","ATTACK_TYPE_HTTP_PARSER_ATTACK","ATTACK_TYPE_HTTP_RESPONSE_SPLITTING","ATTACK_TYPE_INFORMATION_LEAKAGE","ATTACK_TYPE_LDAP_INJECTION","ATTACK_TYPE_MALICIOUS_FILE_UPLOAD","ATTACK_TYPE_NONE","ATTACK_TYPE_NON_BROWSER_CLIENT","ATTACK_TYPE_OTHER_APPLICATION_ATTACKS","ATTACK_TYPE_PATH_TRAVERSAL","ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION","ATTACK_TYPE_REMOTE_FILE_INCLUDE","ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION","ATTACK_TYPE_SESSION_HIJACKING","ATTACK_TYPE_SQL_INJECTION","ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE","ATTACK_TYPE_VULNERABILITY_SCAN","ATTACK_TYPE_XPATH_INJECTION"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0321130232020121-0010023111030323-3023100322011320-3222001110203201-1212221212323003-0013202010231232-0131203313132331-2023221130323302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_action.app_firewall_detection_control.exclude_bot_name_contexts` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- waf_action.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-2231121312223131-1320132030312333-2213010200222123-3220300221010031-1332021211103133-3322232020002310-1330100201031010-3023130323130302"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("bot_name")}
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_bot_name_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0233322021121021-1020101030332310-1202103223100132-2131012001322332-1232122023333000-0210103031110231-0331231210011003-0222301201212322"></a>

### Direct properties for `waf_action.app_firewall_detection_control.exclude_bot_name_contexts`

<a id="canonical-1321301102011012-0310203102012333-2113223321021332-3113032022330013-3211001120101012-2333013030133200-1330031210330003-1010332331213121"></a>

#### `waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` property

Type: `"string"`. Optional.

Bot Name. Human-readable name for the resource

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2023301201222231-0311313312020102-3232203322212121-1323122133301123-3321220330323000-3202033033333020-2121210022301213-3021102030120201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_action.app_firewall_detection_control.exclude_signature_contexts` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- waf_action.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-2002232121300211-3232321313232120-2123203121030122-1110010100222330-3220123313310231-1203132200202311-2202013020330233-0330220012112210"></a>

Type: `"object"`. list nested block, Optional.

Signature IDs to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("signature_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_signature_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023232013102103-0310202330011213-1112122123033131-3200101330233312-3221032000032222-0113113200222103-2223021033100013-2120312122132222"></a>

### Direct properties for `waf_action.app_firewall_detection_control.exclude_signature_contexts`

<a id="canonical-0302331300021002-2202300131013122-3312102321333212-1331330313212033-0203031212000012-3112230120120031-3303100213201232-1000032133303030"></a>

#### `waf_action.app_firewall_detection_control.exclude_signature_contexts.context` property

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0211210033323210-1303120212112233-0220101030201301-1001122232200022-0330130023310100-2323332112012113-3021323233203322-3101001013023222"></a>

<a id="canonical-1221032131303311-1232313133202112-1201000313332332-3123210213011311-1211130103102320-2220100332212103-0131002033002020-2330123302022220"></a>

#### `waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` property

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1303233330233001-2100200201011001-3133220013222203-0131100220332012-3113211013321200-2031322310321033-3212003110011311-0113123121223023"></a>

<a id="canonical-3320210010102320-0303012330033322-1022122023020100-3122121022123100-0112022211330031-2112320122110331-3211021200331313-3200223122121003"></a>

#### `waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` property

Type: `"number"`. Optional.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 299999999),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-2323023101030330-0311033202012133-3120010032230210-2311100301223311-0221003311113000-2301222003331101-0133202202221320-0222121311320033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_action.app_firewall_detection_control.exclude_violation_contexts` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100)
- waf_action.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-2002213110202130-2303030331103001-2313312323000001-1321231022313032-2011322002310132-2110121132211122-2230111031233322-3100013033001321"></a>

Type: `"object"`. list nested block, Optional.

Violations to be excluded for the defined match criteria.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_violation_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130131102022212-3223010121321022-2101123101233002-2012000001220333-3332132120320220-2132132233033301-2220311313123310-3310313010202303"></a>

### Direct properties for `waf_action.app_firewall_detection_control.exclude_violation_contexts`

<a id="canonical-2213132312321330-2111101320303132-2310111300031003-3311110023111200-2013133303030133-0000211013301221-1113132131211113-0121201011120111"></a>

#### `waf_action.app_firewall_detection_control.exclude_violation_contexts.context` property

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CONTEXT_ANY","CONTEXT_BODY","CONTEXT_COOKIE","CONTEXT_HEADER","CONTEXT_PARAMETER","CONTEXT_REQUEST","CONTEXT_RESPONSE","CONTEXT_URI","CONTEXT_URL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0312113122213230-1031201020012332-2011002031322133-2322321323001012-3331200211033320-0011331021302131-0233301011003231-3213212101100231"></a>

<a id="canonical-1322232212201130-1320022300032032-3013000031222223-2122200312201203-2230011131120330-2122123120131121-1101132112110100-0301202233000000"></a>

#### `waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` property

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0333110202323322-1120213031231030-3332332221131011-0302102232020300-2012033000201110-1130100133333112-1210023010302103-3201000102230023"></a>

<a id="canonical-3233302200232200-0213032023132331-0323130321322201-2021300021230210-0033301132320312-3122033320331022-1012132311110021-3111012021323113"></a>

#### `waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` property

Type: `"string"`. Optional.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Additional upstream details:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIOL_ASM_COOKIE_MODIFIED","VIOL_COOKIE_MALFORMED","VIOL_COOKIE_MODIFIED","VIOL_DATA_GUARD","VIOL_ENCODING","VIOL_EVASION_APACHE_WHITESPACE","VIOL_EVASION_BAD_UNESCAPE","VIOL_EVASION_BARE_BYTE_DECODING","VIOL_EVASION_DIRECTORY_TRAVERSALS","VIOL_EVASION_IIS_BACKSLASHES","VIOL_EVASION_IIS_UNICODE_CODEPOINTS","VIOL_EVASION_MULTIPLE_DECODING","VIOL_EVASION_PERCENT_U_DECODING","VIOL_FILETYPE","VIOL_FILE_UPLOAD","VIOL_FILE_UPLOAD_IN_BODY","VIOL_GRAPHQL_FORMAT","VIOL_GRAPHQL_INTROSPECTION_QUERY","VIOL_GRAPHQL_MALFORMED","VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE","VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION","VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST","VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS","VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST","VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS","VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT","VIOL_HTTP_RESPONSE_STATUS","VIOL_JSON_MALFORMED","VIOL_MALFORMED_REQUEST","VIOL_MANDATORY_HEADER","VIOL_METHOD","VIOL_NONE","VIOL_REQUEST_MAX_LENGTH","VIOL_XML_MALFORMED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0122121012021132-1000022311102111-3131301323221132-0220303102210311-3223101332203121-2201303111022201-2320030331311003-3211100200333212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_action.none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- waf_action.none

<a id="canonical-0203313012102132-1232220131100013-2102313200231111-2202023011200331-0321000311321331-2112212300301323-2303313202121100-0203313013320102"></a>

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
none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300310201333330-3133013011222002-2102212230202020-3113323220303320-0031002131200313-1213322102302103-0221203123001033-2303330302222312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_action.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312)
- waf_action.waf_skip_processing

<a id="canonical-3313002212212011-1012223132001111-2033223121201032-0333003111331031-1323210230303010-3113320221320021-2300200010122121-0103203202012311"></a>

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
waf_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.
