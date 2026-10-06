---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0120231131230303-2233111220301332-0232011211110101-2021222130233232-2030211122323220-2030302123022331-1013200212022223-2330003101302300"></a>

## `more_option.response_cookies_to_add.max_age_value` property

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-0112013303113300-1102301123010323-0020102330112213-1003313331031230-1131200231110012-0203032313110313-0212120133031311-1201210213323010"></a>

<a id="canonical-2020033200112223-3301102130232003-1313031010302003-1103331013222001-0100003231111203-2223110132111000-2202323031222231-2123133201101020"></a>

## `more_option.response_cookies_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1330320033311122-1203310312122300-2120132031222212-3023011232013222-3032220000322222-1221233102203133-2200302132302020-1213330323301123"></a>

<a id="canonical-0232131133010032-1111032131020311-2202010202312002-1001220122122022-1001213232222202-1301012200133213-0012120310210002-3230220213022301"></a>

## `more_option.response_cookies_to_add.overwrite` property

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

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

- [samesite_lax](data-sources--http_loadbalancer--reference--group-021.md#canonical-2213213033300102-1001312000303233-2322030203332322-2001210201311300-0230221021321222-1102121031331202-0203231312121202-3021220200223030): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-021.md#canonical-3033010232001333-3312013123310011-2122310320111100-2112121201231211-2223200010133310-2310133112233130-1113122300002022-0212302132121131): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-021.md#canonical-1013210113102030-0000002220023203-2120301333021130-3233130200233111-3000132120010233-3313013030021312-1003112120021100-2133312110113211): complete subsection reference.

- [secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-0302230030012322-3321031112030212-2012330212323332-3012132212011311-0011120231202020-3203220032012223-3212032111310010-2023320221011122): complete subsection reference.

<a id="canonical-0321330303012123-1101331211313221-2200103112200211-3113322301101333-2211013222012200-1112201123333133-0310212333321032-3133101031120230"></a>

<a id="canonical-0110332131302103-2111012331122221-0111212100311222-0302303023220021-1200212332103132-0003233133123310-1101203232301232-3102231113232301"></a>

## `more_option.response_cookies_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0300223001322130-2030032331032033-1003212011020133-2233320201331011-0030003131232030-2312222132220200-0001123123322032-2313020031031221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.add_httponly

<a id="canonical-2210010000020033-0101102003330202-2013322110320320-1310021201333112-2311111331132102-0311322203120022-3320030101013302-1031023323102110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-2330103311303032-3130031021333331-1202200000210223-1010020332113233-3213302010103030-1032313211131110-1222033112230303-2200112031020203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.add_partitioned

<a id="canonical-3212132312101300-1032332332010301-0031323103313131-0021302002002222-1003301203313223-0131222133020203-1201123330303002-0122201303031103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add partitioned.

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

<a id="canonical-1120032113222132-2213222211032003-1322223121013023-2312012011220113-2000301120202213-0212000001233232-0103230330300101-1111330203321121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.add_secure

<a id="canonical-0223113220332011-1101231133103130-1121030012013001-2330020323013132-0130110011223033-1232231012120011-2201321002330031-2122121121202231"></a>

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

<a id="canonical-0312220213101021-3030212313010302-3202012101032100-2111031101023331-3120133123002032-2223302102000101-2330220300103220-0111233100003032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_domain

<a id="canonical-0212331123303031-2311101123010131-2332322120301021-3122301122130112-3203021301201131-3122330312303331-2103220132231322-3030323201113100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore domain.

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

<a id="canonical-2022131312122120-0313230201310123-1101001012033031-3210112120013101-3230021320030132-3013030010002033-3312223102022221-3000132002223012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_expiry` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-2010321213022011-2021003102023110-3112132321220201-3002300001112310-1210030210320033-0332112313232320-1232312101202222-2311231230212100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore expiry.

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

<a id="canonical-2330101032030000-0303210323310033-3012130033223300-0002321010133212-1002323223130123-3212131223020023-3310023220121103-3323022002320332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-2132001130220023-2123200300303001-3302202213222001-1221102012002320-2320302123010323-2012102333333311-1110131003102320-2200003322123303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-1020012011300002-0212210002102001-2332232100130203-1310112320331220-0321213103033111-3321203221331232-3233103020213010-2231210113023131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_max_age` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-0123110211333212-0223212133021233-2000031303313232-2300300033110303-2003022231221310-2033221232121130-1322302101011300-3022120101131213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore max age.

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

<a id="canonical-1112032110322101-0323332202012031-1111232121233213-0100112000110132-1312301300330230-0203012013200310-2202212102303321-0201113021011112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_partitioned` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-0021230111112231-3323322032231231-1222211031223232-1212310313101033-0112230022112322-2200212321011012-0122000102313002-3002332232203032"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore partitioned.

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

<a id="canonical-1123231131001201-3301320030020311-0021020221303001-1213311000213321-1133011030330202-1122013213211230-2033103300113223-2230232203220113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_path

<a id="canonical-3212210333221011-2033303231212210-3223302312222110-1001203020332203-2100212000300013-0122323220211203-1311002013202131-0132031121111101"></a>

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

<a id="canonical-3310323001033201-2132300030302321-0131212102122102-2311300130333121-1021223233022010-0110233321213022-3003311000212200-1100013303321211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-2303301021020021-3021100310002011-0213120312212231-3221203310313020-0012200023001321-3100211301233233-3030313230332233-1211322222002131"></a>

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

<a id="canonical-1032301223020021-1112010033311012-1021301212331103-2133132101232310-2120231003021210-1022132222130213-1003021013201232-2031020120332230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_secure

<a id="canonical-0111333201233001-1210201310023113-1313022230312320-2210002221032302-3131233002223302-3321031120200103-0021010321123203-1200003023331101"></a>

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

<a id="canonical-3323210223311103-2002213230221021-1310310300003112-1331323212212201-2232002302031111-3003312203213313-0020013010022100-1310301302011320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_value

<a id="canonical-1213003230331300-1210022201122012-0223322131110033-1013332003303100-1221223232101031-2233232201311301-1030121222310022-0330131121123332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore value.

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

<a id="canonical-2213213033300102-1001312000303233-2322030203332322-2001210201311300-0230221021321222-1102121031331202-0203231312121202-3021220200223030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.samesite_lax

<a id="canonical-3202200003013012-3232001111333212-1332212002202121-0312000000033102-1210133320023222-3002330212312132-1013102233333331-1313130320110120"></a>

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

<a id="canonical-3033010232001333-3312013123310011-2122310320111100-2112121201231211-2223200010133310-2310133112233130-1113122300002022-0212302132121131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.samesite_none

<a id="canonical-2132321113321230-1122320201302110-0112133101012130-3222103001120103-2003312312133101-0030003303301200-0112222111322010-0300232013330321"></a>

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

<a id="canonical-1013210113102030-0000002220023203-2120301333021130-3233130200233111-3000132120010233-3313013030021312-1003112120021100-2133312110113211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.samesite_strict

<a id="canonical-1133020122101211-0002222210031222-0110222101113133-2130010111302303-0033021130311032-1003300313113323-1222131133331302-0000101111121033"></a>

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

<a id="canonical-0302230030012322-3321031112030212-2012330212323332-3012132212011311-0011120231202020-3203220032012223-3212032111310010-2023320221011122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.secret_value

<a id="canonical-3022202032232212-2130011311110303-2030201002012220-3200111200322300-1130120011210233-0231223220023312-2000203012102132-0002031313011133"></a>

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

<a id="canonical-0100100033033332-0202002322133300-2212321103103311-1231023201113220-3321131031011122-0330132112123310-2211210202132223-0331122010001033"></a>

### Direct properties for `more_option.response_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-3221030132320202-3330201112031213-2012330013120223-2021202223202021-0130212303210312-1232101010033313-0002202330201133-0131031300020113): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-2223031102111001-0010323200311031-0021330231133232-1220122101032213-0302030312303310-1200122123232132-2210220312130033-2131020131213301): complete subsection reference.

<a id="canonical-3221030132320202-3330201112031213-2012330013120223-2021202223202021-0130212303210312-1232101010033313-0002202330201133-0131031300020113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [more_option.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-0302230030012322-3321031112030212-2012330212323332-3012132212011311-0011120231202020-3203220032012223-3212032111310010-2023320221011122)
- more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-2323322311100301-2023212100122202-1233211110221220-3332023210000132-1203230031303132-0220303002212302-1131233132101220-2030111132003323"></a>

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

<a id="canonical-2012113300012231-0212313011000101-3231003120121113-1212020013320100-3132033100310203-3311201120021221-1311322223321101-1322021021010002"></a>

### Direct properties for `more_option.response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0211011303230133-3203001120322130-1202331322333221-0201023130331000-1120213103111323-3320302231010301-3021233010212313-0331100331303313"></a>

#### `more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1023312210333200-1213313101302330-2212313311302003-2332133213303200-2231333122302033-2203301120202003-3320112010302012-2031322031321022"></a>

<a id="canonical-0013313002320302-2002110301113200-1203103311330211-2232111303203301-0322030331131022-2210031023011003-3302123322000002-3320030212233223"></a>

#### `more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0011110321232320-0100111113300113-3213113311322112-1311203132321203-2113203203313323-3011120023301032-3000312300332232-0113100232121010"></a>

<a id="canonical-3322010321201321-2223211221003320-3223230210220011-3310022132002102-3130032212012030-2132320122303211-3101233020212021-0323232133231002"></a>

#### `more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2223031102111001-0010323200311031-0021330231133232-1220122101032213-0302030312303310-1200122123232132-2210220312130033-2131020131213301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [more_option.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-0302230030012322-3321031112030212-2012330212323332-3012132212011311-0011120231202020-3203220032012223-3212032111310010-2023320221011122)
- more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-0131231201022222-1333320222323103-0010121310321311-2212113303322011-2133133222333103-3112223013113120-0112320102222133-1222213130101012"></a>

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

<a id="canonical-3231112111223213-2300222111323221-0300203223112321-2100132031002012-1133230200210133-1313223230331313-1121300020212223-3223221201111321"></a>

### Direct properties for `more_option.response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-3031012323202111-1033312233230033-3310230320022030-0311201100120221-0301001331111301-0020120133302312-0331302102122223-2233320200233323"></a>

#### `more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1233331323213123-1120332333300210-3212213310001101-2201321111011111-0311200002212012-3221030103033323-1110012322212112-2002030201122130"></a>

<a id="canonical-0111231020323321-1030333132121021-3020223313303321-3201200133003210-3221233321230132-3222123001332023-0000100231020120-0223221103310131"></a>

#### `more_option.response_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1002330210300303-2230220311120303-1212313221313221-2322232302001122-2301211303310033-3003133330013332-2202232123020003-3010130310012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_headers_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- more_option.response_headers_to_add

<a id="canonical-3111223013313030-2310133003221013-0030203200203220-3321132313210023-2302201203232323-2132200011303312-0131200132233310-0122322122332200"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

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

<a id="canonical-2132130223220231-2000232002231310-2011311232001222-3212113133211113-2232222111011201-3031113223201103-1001222201100312-2021000120022031"></a>

### Direct properties for `more_option.response_headers_to_add`

<a id="canonical-1331131101121321-2100303111213210-1211030110110130-0023003230123112-2222132002011012-2001331130220011-1112122012003333-0220033220130011"></a>

#### `more_option.response_headers_to_add.append` property

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

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

<a id="canonical-2300221123233033-2212121033131301-3311312333300131-2300233213031302-1232001323203203-3031100311231230-3022331123003111-0003022031310113"></a>

<a id="canonical-3032210011023301-3003023330000320-3301010331220202-3221133030212300-3120010011232011-3223033300113132-2120132333310012-1121030233113032"></a>

#### `more_option.response_headers_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3212022231010330-0203331033101323-1001132002301230-2012330123131230-2110021112302030-2102103312212011-0310313023132302-0221201132030000): complete subsection reference.

<a id="canonical-1000202123102321-0310332211023122-1123011033213003-3223202222102023-3121311312013100-3002130213320120-0033031120002100-1301202310300132"></a>

<a id="canonical-0332301311123032-1322223210203130-3011213213320321-1100213033210121-3230202233110100-2112200330031302-0003200030333013-0302111302230020"></a>

#### `more_option.response_headers_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3212022231010330-0203331033101323-1001132002301230-2012330123131230-2110021112302030-2102103312212011-0310313023132302-0221201132030000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_headers_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-1002330210300303-2230220311120303-1212313221313221-2322232302001122-2301211303310033-3003133330013332-2202232123020003-3010130310012121)
- more_option.response_headers_to_add.secret_value

<a id="canonical-1231112121103033-1202232131311230-3203100001110310-0323220003313312-3130221102023320-2210222120123103-2331303230313333-2131222303332023"></a>

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

<a id="canonical-1100311021010313-2232311013030101-2330322323310321-3111001132003000-1020120101211110-3310112311133212-3220123020201020-1133301102211302"></a>

### Direct properties for `more_option.response_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-0210300123133110-2331100013321200-1031202323100310-1123013031321012-1021100031121211-2313132101122002-3110101102302301-2323102022030001): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-2303032030300103-0313202112202213-1331301030302200-0303313310021033-0313032233023013-0232112131331202-2110032133213001-0233112131102203): complete subsection reference.

<a id="canonical-0210300123133110-2331100013321200-1031202323100310-1123013031321012-1021100031121211-2313132101122002-3110101102302301-2323102022030001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_headers_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-1002330210300303-2230220311120303-1212313221313221-2322232302001122-2301211303310033-3003133330013332-2202232123020003-3010130310012121)
- [more_option.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3212022231010330-0203331033101323-1001132002301230-2012330123131230-2110021112302030-2102103312212011-0310313023132302-0221201132030000)
- more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3023312322233333-1300020231332312-1232113302011210-2013022230203010-1103333000103211-3101212131210021-3020031231213301-3200103332303223"></a>

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

<a id="canonical-0121101023331131-2022211020212212-3132320023311321-1203233003120210-3001003232303330-1110212203311131-2322313312233303-3232033231230133"></a>

### Direct properties for `more_option.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0130300122010313-0112020323030302-2233220033112221-2331003010121333-2213030022331332-0031002122130231-2220322232303112-1333010112001200"></a>

#### `more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2321202012112312-2121210021103303-3221202332221123-1332131031021010-0203301200110230-0033300211131302-2331331133013013-0332010322110211"></a>

<a id="canonical-3032011310232012-0030223112000131-3123201301322212-1012310223111002-1022121011312102-0003102212130232-2223030312233331-0300311111300321"></a>

#### `more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0110201101200312-2023100222311020-2010003133303210-2302030003321031-0001201032131003-2221120320230332-1023120303222133-1231223120203331"></a>

<a id="canonical-3203311312230011-1022320201003002-1023313210010030-2121031303013322-1122203013012230-1132031321131303-0002302012302203-0011220002230112"></a>

#### `more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2303032030300103-0313202112202213-1331301030302200-0303313310021033-0313032233023013-0232112131331202-2110032133213001-0233112131102203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_headers_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-1002330210300303-2230220311120303-1212313221313221-2322232302001122-2301211303310033-3003133330013332-2202232123020003-3010130310012121)
- [more_option.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3212022231010330-0203331033101323-1001132002301230-2012330123131230-2110021112302030-2102103312212011-0310313023132302-0221201132030000)
- more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2022331010030303-0023102311110131-2123321313311110-2300302130002123-3110310201212123-2010022022032031-3030302311230112-2302312033131130"></a>

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

<a id="canonical-2012333113210231-2331100012013331-1311023212302123-0231311033102032-1201311001222003-3013030220320331-0130031021133133-0021111301123000"></a>

### Direct properties for `more_option.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-0222321202020133-2101010112302122-2002331202212003-1221022113010222-0003030122213211-0213112011320220-3120011011130320-3111001120002221"></a>

#### `more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1103021303130222-2230210032223221-0313131103200031-0002310011200322-3111333312201020-0001002203331023-3333121110201003-3300010022133023"></a>

<a id="canonical-0103132202011123-0331222113302303-2110201101121313-3310200011333113-0212202230233020-2100302321211202-1020102300333311-3010002230213230"></a>

#### `more_option.response_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3312033113300312-1002132313113311-3233122303230311-2103222022222003-3000233023022122-2023233330322332-1333021223220210-1021232333203310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `multi_lb_app` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- multi_lb_app

<a id="canonical-3303101211012301-0112301312033011-3101222111133102-2112323301002031-3333132011001210-1303100102132022-2302112333113013-3120233030123231"></a>

Type: `["object", {}]`. Computed.

\[OneOf: multi\_lb\_app, single\_lb\_app\] Configuration parameter for multi lb app.

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

- [multi_lb_app](data-sources--http_loadbalancer--reference--group-021.md#canonical-3303101211012301-0112301312033011-3101222111133102-2112323301002031-3333132011001210-1303100102132022-2302112333113013-3120233030123231)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-0031200011131311-2233112323100033-0111001011133120-2221230201121322-2021123130030211-2222030012320123-0323231202112012-3300012012003232)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033322013033333-1012021230311113-3030032222123110-2130202012323311-2130011103202221-3200200101001031-2230202203012132-2012120233023302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- no_challenge

<a id="canonical-0121222330302201-1313000022222103-0000332201021130-0021331132102001-3010230020131302-3021233321102212-0331230031230303-0000030023131011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-1012220231323000-0002101033211030-3111020332233130-0011331233023201-2102321003221301-1230013230021133-3230200102213220-0022221221200321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_service_policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- no_service_policies

<a id="canonical-2210312233210311-0002123133331020-1220122121222001-0132233321112131-1121331112303320-2113021232203331-1202011022320022-1211301213031113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no service policies.

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

<a id="canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- origin_server_subset_rule_list

<a id="canonical-0303030233300122-3332220002211232-0333301002233300-1211031003230112-0301311012023111-1131321021222333-3023203303301301-3202001221001112"></a>

Type: `"single"`. Computed.

Origin Server Subset Rule List Type. List of Origin Pools.

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

<a id="canonical-1213121323112302-0202100331103131-1101123211331323-3001020221001100-1002323322011021-3212300113210122-1311303211312300-0001022001100120"></a>

### Direct properties for `origin_server_subset_rule_list`

- [origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113): complete subsection reference.

<a id="canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- origin_server_subset_rule_list.origin_server_subset_rules

<a id="canonical-0002300201201300-0100200220100312-1111220332220113-1301331212323213-3312131311330033-0330313200012320-0310112101221300-3221112020123100"></a>

Type: `"list"`. Computed.

Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN,
Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin
Server Subset is a sequential engine where rules are evaluated one after the other. It's important
to define the correct order for Origin Server Subset to GET the intended result, rules are evaluated
from top to bottom in the list. When an Origin server subset rule is matched, then this selection
rule takes effect and no more rules are evaluated.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-3101122010233301-2100121131311320-0331012311012031-2100103013012201-3310320301013301-3331022030111033-0132332013032122-0321003131001223"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules`

- [any_asn](data-sources--http_loadbalancer--reference--group-021.md#canonical-2032000020020010-1123302032300201-2232220322231131-1012023101220201-0302130002001223-3333300112021112-2220300201332001-0211130212230322): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-021.md#canonical-3230013232200221-3212322310012222-2211001322100202-3333010132022233-1113223202303220-1030321113212203-2201011013213133-2132220200022021): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-2313113113021300-2203100123203023-0230212303310111-2310301301023231-1330302032312322-0121213323110311-3033331203021112-1331330030120322): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-2220012312230103-1201130010233021-3102111031311130-0010033203103323-0230113231130031-2111131202211032-2013212123010202-3220003023311113): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-021.md#canonical-2103123233211101-1110032203230202-2002021201211231-0022011033332321-3321332221312222-3333023200110202-0103113000313133-2232303110110310): complete subsection reference.

<a id="canonical-1031031132022321-1131002303133210-0020121223323233-0223031220032212-1232312123221133-1201130012222312-1323311221210300-2012131223201220"></a>

<a id="canonical-0302023003211131-0001103033300022-0022323233313030-3100121221102312-0203032302320220-3023112022213012-2231011010120131-1211310002213033"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.country_codes` property

Type: `["list", "string"]`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-1303030032032131-1112131301012021-1211031323120231-1112220312122233-1330203201231213-0133101222312300-0200032222021123-2031211213323002): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3303312131012321-3302333210230231-2123131120233002-0011002032332003-3331302301210111-0311032201303101-0323212202312030-3230123212332022): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-021.md#canonical-3310311213320212-1221113303321301-3312133113302311-2203200311323223-1033032331103013-3312230022130121-0310110333322033-1132230223230331): complete subsection reference.

- [none](data-sources--http_loadbalancer--reference--group-022.md#canonical-0111233133120120-0323011111302022-2230023031221212-1230110330012322-2211321231001201-2002321333213021-1021202022222123-0122221101002230): complete subsection reference.

<a id="canonical-2112101032222020-2002233112211323-2000320210020201-0131120200012111-1202313023311220-2330312111213221-0010033321121112-3320011103201232"></a>

<a id="canonical-3122002002012221-3011212003022312-2323311021022003-0122032003320333-2010231221332323-2202012203311302-2103010303311000-1130232220332133"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.origin_server_subsets_action` property

Type: `["map", "string"]`. Computed.

Add labels to select one or more origin servers. Note: The pre-requisite settings to be configured
in the origin pool are: &#8203;1. Add labels to origin servers &#8203;2. Enable subset load
balancing in the Origin Server Subsets section and configure keys in origin server subsets classes.

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
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1101221301132003-3003201321003130-2333302203011230-1311131331023220-0230300222023130-0102302332002323-1330311000220021-1000202130113113"></a>

<a id="canonical-0111001230212122-3233230230000303-1232303031310300-3232123132113333-2102200111110331-2222130303303101-3322022123320130-0323132322223032"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.re_name_list` property

Type: `["list", "string"]`. Computed.

RE Names. List of RE names for match.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2032000020020010-1123302032300201-2232220322231131-1012023101220201-0302130002001223-3333300112021112-2220300201332001-0211130212230322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.any_asn` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.any_asn

<a id="canonical-2323022213231322-1200212323010120-3121132003010100-3323023011323023-2313122223303121-2032230112031132-1203032322102001-2203231100121110"></a>

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

<a id="canonical-3230013232200221-3212322310012222-2211001322100202-3333010132022233-1113223202303220-1030321113212203-2201011013213133-2132220200022021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.any_ip

<a id="canonical-1230033312100101-3021031120123003-1332300130112103-1230202333220221-0210301032310103-0212002202203220-3313201302021302-1321131003001223"></a>

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

<a id="canonical-2313113113021300-2203100123203023-0230212303310111-2310301301023231-1330302032312322-0121213323110311-3033331203021112-1331330030120322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_list

<a id="canonical-3132333230033232-3032331310223000-3002012213023232-3022233020022233-2203211320312301-2330012201103111-1221132311003033-3222332022000130"></a>

Type: `"single"`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-0122202023332102-3013322003002031-2320011032123000-1222212202030220-3233121333203011-1103211113302301-0003122130021333-0221132212311231"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.asn_list`

<a id="canonical-2111022031310330-3223112120332001-1220303301120203-1033322331232102-2020311020320032-1302030013301323-0113121032302110-3132012233110331"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_list.as_numbers` property

Type: `["list", "number"]`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-2220012312230103-1201130010233021-3102111031311130-0010033203103323-0230113231130031-2111131202211032-2013212123010202-3220003023311113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher

<a id="canonical-3211120030102301-3013320101331223-3232121322113031-1301222002020320-2323031103130211-0311110201021013-3302000113001130-1202000010333100"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-0301211101222033-2301120011233131-2133323230201221-3033200313323102-2323301032202233-0210033211221033-0122321233110303-1313323023031110"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher`

- [asn_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-1332131120323011-2223232132012303-2133332011330200-0230332230031002-1100111232122322-0112232222002300-1223203131110001-1100321320230323): complete subsection reference.

<a id="canonical-1332131120323011-2223232132012303-2133332011330200-0230332230031002-1100111232122322-0112232222002300-1223203131110001-1100321320230323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-2220012312230103-1201130010233021-3102111031311130-0010033203103323-0230113231130031-2111131202211032-2013212123010202-3220003023311113)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets

<a id="canonical-1002231233131302-2103201001220320-0033300110000312-0203123201311231-2133103103212012-1331311033303223-1111011332212001-2101020111223033"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0110022222021133-0323210123232111-3002313200331111-2010102013003100-0113300130121230-3123320322122230-3312111313220312-1013011130031212"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets`

<a id="canonical-1123311103230003-3032132231223102-0100220221332103-3213031332032113-2000030032200200-0130322110200032-1333112302202332-0003200301122011"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets.kind` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1330310200032312-1113031031023030-1233302020333301-0300031311220213-1210200131220010-1230201210021011-1232022103111332-3310031230100133"></a>

<a id="canonical-1120320031201132-1212111122132132-3011331211303331-0211231200120133-1233321030213011-2131301011131230-0232130332100103-0002202031331231"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3331021030223322-1122302112303210-0111332130232022-2110333231333003-2121012233231132-0130020323033120-1300031312032101-3112131222332210"></a>

<a id="canonical-2210003222310321-0133202300121320-2133121110032121-0010200201302122-3021332103030111-2301230313330220-3210311223300330-1331231110321022"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets.namespace` property

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

<a id="canonical-2200032103002103-1013123232313102-0111300132312133-2100301132101202-3231312332110232-1110102220200202-1022121212203111-0330023030301331"></a>

<a id="canonical-2213203211202112-0101331221120201-1300110333033222-2220023333200033-3000002202230223-2003120003110031-2313333100032112-2313101203031303"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2211033002332221-3000213122311130-0200021103020320-3112101022222100-2223300333231202-1123122210022113-3003103003231212-0213020320130232"></a>

<a id="canonical-0121231121103001-0212303200011101-2220003333031202-0232222002201120-0313330201300233-0301201101033032-3201220123033222-0033033030231030"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets.uid` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2103123233211101-1110032203230202-2002021201211231-0022011033332321-3321332221312222-3333023200110202-0103113000313133-2232303110110310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.client_selector

<a id="canonical-0013310020221101-2000013030322220-2312313301110132-3303133002010113-2113210311203302-2200112011312013-2002323211023003-2301020201122203"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1110021303320012-3100033032200212-1010322332233012-0300133310333100-3010113331110221-1310210321331323-3021210221222032-0212313122230333"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.client_selector`

<a id="canonical-1203101203332301-0231113132313230-1320111130211102-2311002330103031-3001131133330032-2303111002310202-0222000013103020-0021012103320121"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.client_selector.expressions` property

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1303030032032131-1112131301012021-1211031323120231-1112220312122233-1330203201231213-0133101222312300-0200032222021123-2031211213323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher

<a id="canonical-2021203220112022-2301120121233022-0010132303113231-2131202131003110-1112112112021100-0132303122102301-3131332220231123-2132013030130231"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="canonical-3230110223020300-1302013320232312-0322030333323011-2112211303332332-1100211332022000-1332022320323121-2332330120211033-3033130033223021"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher`

<a id="canonical-3322220000020103-2012222131020003-2122223033310221-3003211002130100-1120131003123100-3202122011311321-1223103122023012-2232220223030130"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.invert_matcher` property

Type: `"bool"`. Computed.

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-0133331230123333-2203000212003303-3302122222033113-3131002322322033-2121033010211002-1221011313010111-3102230202231223-0131311131122131): complete subsection reference.

<a id="canonical-0133331230123333-2203000212003303-3302122222033113-3131002322322033-2121033010211002-1221011313010111-3102230202231223-0131311131122131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-1303030032032131-1112131301012021-1211031323120231-1112220312122233-1330203201231213-0133101222312300-0200032222021123-2031211213323002)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets

<a id="canonical-2003132012113010-2213111022120321-0213102320332201-3321233032303000-1013322312103333-1033002130232210-3201230011310200-2022113200033213"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3312320131100301-1020233101222131-1303303013211230-0330231013130330-0120223303221122-1123322213313302-1002302131030122-0100112201012110"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets`

<a id="canonical-2121301102320312-2132121311032011-3112021031120233-0123303312132101-3223011230003000-1030320232013121-2333320101223231-3211310102223332"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets.kind` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1233311020001130-2210003301131013-3323321121220022-0102000010012111-3320133101210321-0321213332203213-0121220011011112-1233002221233132"></a>

<a id="canonical-1013312302032230-0221310023202122-1131121220310000-0302102333320023-2011111131020302-2230110030010023-1121021120303331-1331222221220031"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2100302321132222-2300002220222230-3121300113002133-1103113311121231-2312312101031011-2232123322101100-2013202313331011-3012033322203003"></a>

<a id="canonical-2102032001230211-0331112131012322-0113202032323101-3021110333211310-3020301103120113-3332020013021201-3233313311113033-1333001113321133"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets.namespace` property

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

<a id="canonical-1100011113321013-0032113230011302-0203122002300021-0101023203112233-0210010111210102-2332331131002111-1032211130323333-0321231012021021"></a>

<a id="canonical-1231111013332133-0023030100102322-1202010310203123-1331210120010220-1233023300210132-0122320003311103-3032303102120303-2221211020100213"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2203233111030212-0201320112010200-1302223300233211-1312032110222223-1000113033200123-3110213001031101-1122312312122103-1103122302032323"></a>

<a id="canonical-0103020121333032-3221101312013013-2133033303030100-2232303332322310-1002121122323220-2013121102102133-2303200202120103-0013211210033001"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets.uid` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3303312131012321-3302333210230231-2123131120233002-0011002032332003-3331302301210111-0311032201303101-0323212202312030-3230123212332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list

<a id="canonical-1231120132111110-0302201332013110-3011313003213022-0101220021011203-0113001023210111-0313111021021100-0300232300220010-1111313132213303"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1102002010002302-2320000302201231-2333231022111303-1131320231301301-2312102213121123-1101132002031020-1223022012222331-3223310120002311"></a>

### Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list`

<a id="canonical-1030213202333023-3023200133032220-2230300031222211-1121122130032013-2210013013011112-1220311003013211-1101310010000212-0202321203222000"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list.invert_match` property

Type: `"bool"`. Computed.

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

<a id="canonical-1113232330210222-3320102230110200-1132232020031021-2130132110122222-0301011013321110-0103211310020230-1312032230012201-3131210312220100"></a>

<a id="canonical-2211331302322113-3202132031222113-2332231300213320-2210033210231110-0120313100003122-3222331012122230-3003303133003010-3101132103011322"></a>

#### `origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

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

<a id="canonical-3310311213320212-1221113303321301-3312133113302311-2203200311323223-1033032331103013-3312230022130121-0310110333322033-1132230223230331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.metadata

<a id="canonical-2300211303021310-2223312113033301-0022231122301233-2023300110020200-0122332131121231-0000211113301033-0010001121133230-0031100011130200"></a>

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
