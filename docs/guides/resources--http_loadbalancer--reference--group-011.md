---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1302322013132300-0121110311233002-1110210221000212-3113113230211323-0100000013020112-0031120011301131-1310321331022103-2021003311133110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3231023210233220-1131013113311231-3110223001111211-3003322301130101-0022023310120132-3322003100031313-0323213202311223-1220032202220300)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-3321313311023320-1002202022131322-2210021110000000-1121001331331323-1321111011210331-3210031032312101-3003203002210210-2223110230111033)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-1023130021300320-3221100012012033-0032023011211223-2311003013220200-0013230201030123-1232202222101231-0000203103013301-1311223221313121"></a>

Type: `"object"`. list nested block, Optional.

Custom Fall Through Rule List. Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111023302202213-0032000332322301-1233011331211121-3032113213210210-3232213200201023-0333011123203233-2200032313202111-3332120220023103"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`

- [action_block](resources--http_loadbalancer--reference--group-011.md#canonical-0122023311103310-0013111231213211-3311131012312210-3102020031220232-2320222322220112-2320222122003101-2001301322103222-0002003210321122): complete subsection reference.

- [action_report](resources--http_loadbalancer--reference--group-011.md#canonical-2213112120311323-3330333232322012-0202023333322100-2112301231102332-0223121131133301-0021030202131330-2322231233313300-2203013333100201): complete subsection reference.

- [action_skip](resources--http_loadbalancer--reference--group-011.md#canonical-2131001201203002-2121121220233310-1123330102303132-1203220222231322-3201112231023213-0123312030001310-3133320323023220-2221000213020000): complete subsection reference.

- [api_endpoint](resources--http_loadbalancer--reference--group-011.md#canonical-0010320122200311-2110103300130001-0110000111012010-2122030003230001-1210113003111303-1003303220132232-2121313100201112-0300000033120210): complete subsection reference.

<a id="canonical-0200113220200111-0301220123120111-2032333102200300-2121303032200030-0033101223231313-2101012101123020-2232102003111330-2113222231001211"></a>

<a id="canonical-3100023222131031-0220312130023223-3111212332233223-3230102301302031-3131211332120100-1223111331002010-1113333220220213-2233322231032321"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group` property

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0333233013302012-1321223320220101-3002203102222111-1311002310100210-2231030031100120-0022120112033320-0312321013201313-0233321212212212"></a>

<a id="canonical-0023030013132200-1032321323323330-1000230100122212-3030032332010103-3320020203110111-3233222120200101-1100220333002033-2201210011022023"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path` property

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-011.md#canonical-0012120101300020-3311220132210321-0221033232332120-3303132320332132-0321333303120231-1110202220212213-2020131213102012-0211012112303213): complete subsection reference.

<a id="canonical-0122023311103310-0013111231213211-3311131012312210-3102020031220232-2320222322220112-2320222122003101-2001301322103222-0002003210321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3231023210233220-1131013113311231-3110223001111211-3003322301130101-0022023310120132-3322003100031313-0323213202311223-1220032202220300)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-3321313311023320-1002202022131322-2210021110000000-1121001331331323-1321111011210331-3210031032312101-3003203002210210-2223110230111033)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-1302322013132300-0121110311233002-1110210221000212-3113113230211323-0100000013020112-0031120011301131-1310321331022103-2021003311133110)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-1023032033213113-1231001001020303-1220211203320310-1012221201113023-3002020112222113-1032000222102313-2033220021121022-2102300321133221"></a>

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
action_block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213112120311323-3330333232322012-0202023333322100-2112301231102332-0223121131133301-0021030202131330-2322231233313300-2203013333100201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3231023210233220-1131013113311231-3110223001111211-3003322301130101-0022023310120132-3322003100031313-0323213202311223-1220032202220300)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-3321313311023320-1002202022131322-2210021110000000-1121001331331323-1321111011210331-3210031032312101-3003203002210210-2223110230111033)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-1302322013132300-0121110311233002-1110210221000212-3113113230211323-0100000013020112-0031120011301131-1310321331022103-2021003311133110)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-3131322033330303-0030333233030232-3101231223113130-0221020311033312-1132303302321233-3021000000223223-0031010112222133-0121020121302233"></a>

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
action_report = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131001201203002-2121121220233310-1123330102303132-1203220222231322-3201112231023213-0123312030001310-3133320323023220-2221000213020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3231023210233220-1131013113311231-3110223001111211-3003322301130101-0022023310120132-3322003100031313-0323213202311223-1220032202220300)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-3321313311023320-1002202022131322-2210021110000000-1121001331331323-1321111011210331-3210031032312101-3003203002210210-2223110230111033)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-1302322013132300-0121110311233002-1110210221000212-3113113230211323-0100000013020112-0031120011301131-1310321331022103-2021003311133110)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-3200222313311123-0320101011233103-3001321221312303-3211030120132121-2020130322230012-2111031323021332-1002021113021313-2301103211202222"></a>

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
action_skip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010320122200311-2110103300130001-0110000111012010-2122030003230001-1210113003111303-1003303220132232-2121313100201112-0300000033120210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3231023210233220-1131013113311231-3110223001111211-3003322301130101-0022023310120132-3322003100031313-0323213202311223-1220032202220300)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-3321313311023320-1002202022131322-2210021110000000-1121001331331323-1321111011210331-3210031032312101-3003203002210210-2223110230111033)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-1302322013132300-0121110311233002-1110210221000212-3113113230211323-0100000013020112-0031120011301131-1310321331022103-2021003311133110)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-0311310003131331-1321213330313122-2333010122313030-1013230232222201-3201230002110330-3010231121033020-0231130121222112-1000221101320022"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303222030233333-2103212311013202-2133222232020312-0300211321003120-2022332012130021-2302132010010310-3301122323200132-2233112332112323"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint`

<a id="canonical-2203023132322330-2331031011212122-2013232302233320-1010200310203312-1200232002330200-1032302001323010-0322123101132013-1320202332202032"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1032011300011112-0101031000213022-1002021223002023-2001300203002230-2003213301031020-1231323010311231-0321331213010220-2221323030120320"></a>

<a id="canonical-3133200032332232-0113012232010113-2332303022300130-0012130133211023-1333100030032300-2211113123322103-0132203131110133-0030022220020200"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.path` property

Type: `"string"`. Optional.

Path. Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-0012120101300020-3311220132210321-0221033232332120-3303132320332132-0321333303120231-1110202220212213-2020131213102012-0211012112303213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3231023210233220-1131013113311231-3110223001111211-3003322301130101-0022023310120132-3322003100031313-0323213202311223-1220032202220300)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-3321313311023320-1002202022131322-2210021110000000-1121001331331323-1321111011210331-3210031032312101-3003203002210210-2223110230111033)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-1302322013132300-0121110311233002-1110210221000212-3113113230211323-0100000013020112-0031120011301131-1310321331022103-2021003311133110)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-2123320002323121-3301230313322330-2213323112302131-3020220232310213-2312111321232302-3213311321022110-2122202102210301-3110211303112110"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011211221002013-3021022232110110-1223300121011231-0331310003310022-0132200223032003-2031301310320310-3112130210232101-3121012203132020"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata`

<a id="canonical-1133120301121321-3320113332021110-3333323023102011-0022202001322010-3100133131330320-2013301231301131-2233121230130012-1313110003133013"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-1111200110201100-0000311323022112-2132312120221332-0321232320020213-1332032020003002-2023131033230031-0203203320300012-3021332223330310"></a>

<a id="canonical-2233012130231311-3030211220311333-1312213033221012-0230220022000030-2320230030333032-0323333322221210-3013110300311120-3130320333210111"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="canonical-2220013321223232-1030020022321033-0200333132333230-2003231202021213-2100212112211222-2333222122322013-1222231011112122-0303121110201013"></a>

Type: `"object"`. list nested block, Optional.

Validation List. Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223333000023101-2032220122011020-2012202112302010-0020133110313002-1200010013123203-2333230332312002-1000020230120323-3100111210010121"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules`

- [any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-0202321233031022-3122102033213200-3022120102111003-1003202332003013-2000233202332210-2312123330213323-3033111113021122-3310332233110033): complete subsection reference.

- [api_endpoint](resources--http_loadbalancer--reference--group-011.md#canonical-0303302112230122-0202101133320122-3123200110211220-3122002213103300-3230232001321302-2221300220113232-3320233322312032-0202231233213300): complete subsection reference.

<a id="canonical-1231212030233002-3002212312122033-2221231213303120-2112212213330310-1302112012100322-1001011011033232-3300202030230203-1110122233310321"></a>

<a id="canonical-2022132002322311-1032102101112020-0103101102233010-2212331322102100-0221302211113120-3200322011211233-2130303022021302-3302101323302331"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.api_group` property

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2022032013010030-3322002303333223-1133302212111332-1201121113231123-2112120101212312-3103200010103203-0320230013011310-0332133003123301"></a>

<a id="canonical-2302213110220212-1313301233130320-3031312130330022-0010222123212101-2213031322011020-1312212133103003-0022010003312302-0300121103022011"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.base_path` property

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-011.md#canonical-3121002130331012-2113131002333231-3332221023230010-1110030220303322-1030332022301132-0133103000212000-0123331011333133-2312321033303231): complete subsection reference.

<a id="canonical-1020230113212111-1113023330320333-1332320321233122-1233012122132321-2200112023110223-3131312100213322-3213320303220102-2302012313221123"></a>

<a id="canonical-3233111030313330-1303320122210302-2321131030032133-1033332012213230-3101323023222333-3230100003100010-1022332033323012-3102023010210320"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.specific_domain` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [validation_mode](resources--http_loadbalancer--reference--group-011.md#canonical-1300210222332112-3202232332012200-0000013303132120-2201221030000121-1230122120233211-1001120020312010-0300212200032021-0203310121102333): complete subsection reference.

<a id="canonical-0202321233031022-3122102033213200-3022120102111003-1003202332003013-2000233202332210-2312123330213323-3033111113021122-3310332233110033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- api_specification.validation_custom_list.open_api_validation_rules.any_domain

<a id="canonical-2032222010323100-1131101311312230-2331222113202031-2331313111032121-0003212021201122-0331101023311010-0202021220212030-2110013101023131"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303302112230122-0202101133320122-3123200110211220-3122002213103300-3230232001321302-2221300220113232-3320233322312032-0202231233213300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

<a id="canonical-3212101023010232-0031310232223311-0220311122003101-1103122100001210-1331331323003330-1321132221222210-1112330233333021-3133111212102002"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033001021203101-1230231200122101-0133223222222231-2033102100013322-0210223303320120-1221223201020302-3011101103133203-2033102123333123"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint`

<a id="canonical-3010320102000222-3020111303211102-2332001221001231-3211012301221130-1200002201021320-3033002132332100-1123322013013223-3210221111333213"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2320312110232323-1211231113131200-1332300311202223-3222132111120322-2002002103022232-1223230323000332-3202233002010120-1112132222011302"></a>

<a id="canonical-3233101131130211-1130200033030320-3132131010302000-3101122011010232-2323013033213130-1001113113112221-3031020330303313-0211000112031111"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint.path` property

Type: `"string"`. Optional.

Path. Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-3121002130331012-2113131002333231-3332221023230010-1110030220303322-1030332022301132-0133103000212000-0123331011333133-2312321033303231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- api_specification.validation_custom_list.open_api_validation_rules.metadata

<a id="canonical-3221010310332010-2030011300111002-1010330130101100-0100130321102111-3030130031002300-3233111212121032-1300301031230331-0202332230013123"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013303112122110-0033010303013211-1301231022212212-2333033221211223-2102232131011030-3200222231132011-0300331232330230-3221203310233103"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.metadata`

<a id="canonical-1201023220300020-0131301202331023-2200223330103101-3012200032011212-2232310332003212-0112222112121010-2313301301330203-0132202322102010"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-3013320231113101-0313112302220302-1210120203010303-2110030231303302-1332100310220332-1022222322002212-2002012110200303-1112323030003030"></a>

<a id="canonical-3032013230001120-3303211202313231-2001003003101230-2102331230210032-3000232332213230-0230201311120233-1302111032303013-3121001303300111"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1300210222332112-3202232332012200-0000013303132120-2201221030000121-1230122120233211-1001120020312010-0300212200032021-0203310121102333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode

<a id="canonical-0210323213000220-0321000230222210-0123113031311203-0332233100100122-1030020323233301-1122020012231232-1232133110330113-2201132120000323"></a>

Type: `"object"`. single nested block, Optional.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

Terraform syntax:

```terraform
validation_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330003103020311-3000111210000103-2101130313133101-0111030231000232-1113333033001121-3303213012101001-2110020222123321-2021311202001021"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.validation_mode`

- [response_validation_mode_active](resources--http_loadbalancer--reference--group-011.md#canonical-1311300210023220-1321022203102121-1333312301220232-1000021322233033-0133221112020113-0320201221210332-2222312100022110-3332201100030323): complete subsection reference.

- [skip_response_validation](resources--http_loadbalancer--reference--group-011.md#canonical-3310013100312312-2023322302323313-2321021113112113-2301310323300102-1310331120301033-1322010123030200-2031220233021333-1030231311203200): complete subsection reference.

- [skip_validation](resources--http_loadbalancer--reference--group-011.md#canonical-0133133221023030-0320011133101202-0022112122322030-0210210011331310-0131213033111103-1301031012122121-1110300010213120-0102132020313000): complete subsection reference.

- [validation_mode_active](resources--http_loadbalancer--reference--group-011.md#canonical-3130122203320120-2301300100003113-1120303332111031-2120212003232230-2101012122121033-3101022300111302-1103301132101300-2121103321313120): complete subsection reference.

<a id="canonical-1311300210023220-1321022203102121-1333312301220232-1000021322233033-0133221112020113-0320201221210332-2222312100022110-3332201100030323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-011.md#canonical-1300210222332112-3202232332012200-0000013303132120-2201221030000121-1230122120233211-1001120020312010-0300212200032021-0203310121102333)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="canonical-0303332300231321-0201100202223010-0231213230121000-3201031200001210-1020201112232031-1102022121222311-3101011112131103-0000110203133330"></a>

Type: `"object"`. single nested block, Optional.

Open API Validation Mode Active. Validation mode properties of response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
response_validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120200013102320-2101222033302302-3001201100102200-0101130302212230-2103222133223131-2203123210200003-1033010320302332-1201123110003300"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active`

- [enforcement_block](resources--http_loadbalancer--reference--group-011.md#canonical-3233321032030222-0301022100000302-2011323000033020-3101123102300013-0112120012131311-1333211213010033-3133120120213022-1120032123220001): complete subsection reference.

- [enforcement_report](resources--http_loadbalancer--reference--group-011.md#canonical-1312133322013323-3223120021322101-1210321300322123-0200333333033221-1131303122211033-2303332300223113-2331103233112332-0231222001010320): complete subsection reference.

<a id="canonical-1131233113123010-2320001133310320-3233203002011322-2233322202132001-3211100003211130-2001232223302333-3103220030212123-1220333132232022"></a>

<a id="canonical-0003321133313101-3202323101033221-3322230110001302-2131101020221312-3331020120132113-3023030102133110-3221022033300200-1202223023303100"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.response_validation_properties` property

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3233321032030222-0301022100000302-2011323000033020-3101123102300013-0112120012131311-1333211213010033-3133120120213022-1120032123220001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-011.md#canonical-1300210222332112-3202232332012200-0000013303132120-2201221030000121-1230122120233211-1001120020312010-0300212200032021-0203310121102333)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-011.md#canonical-1311300210023220-1321022203102121-1333312301220232-1000021322233033-0133221112020113-0320201221210332-2222312100022110-3332201100030323)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-3333003322200011-0010122233121311-1013312000330030-0213132033002033-0022032011333131-0313102213102021-0310132022331022-3200132322020223"></a>

Type: `["object", {}]`. Optional.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

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
enforcement_block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312133322013323-3223120021322101-1210321300322123-0200333333033221-1131303122211033-2303332300223113-2331103233112332-0231222001010320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-011.md#canonical-1300210222332112-3202232332012200-0000013303132120-2201221030000121-1230122120233211-1001120020312010-0300212200032021-0203310121102333)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-011.md#canonical-1311300210023220-1321022203102121-1333312301220232-1000021322233033-0133221112020113-0320201221210332-2222312100022110-3332201100030323)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-0120211222223310-2020223313002111-2020320331121132-1130130123131220-1332032101113202-0213221313013032-1121233011200312-0023032313232232"></a>

Type: `["object", {}]`. Optional.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

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
enforcement_report = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310013100312312-2023322302323313-2321021113112113-2301310323300102-1310331120301033-1322010123030200-2031220233021333-1030231311203200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-011.md#canonical-1300210222332112-3202232332012200-0000013303132120-2201221030000121-1230122120233211-1001120020312010-0300212200032021-0203310121102333)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation

<a id="canonical-2022022112023230-2201011021101311-2330210212222003-3032001321312223-3200010232112000-1103032303321300-0303010002131213-3110333001012211"></a>

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
skip_response_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133133221023030-0320011133101202-0022112122322030-0210210011331310-0131213033111103-1301031012122121-1110300010213120-0102132020313000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-011.md#canonical-1300210222332112-3202232332012200-0000013303132120-2201221030000121-1230122120233211-1001120020312010-0300212200032021-0203310121102333)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation

<a id="canonical-3021130102101230-1310333110011120-1311231032131301-2330100312003123-2313001302022300-0230323221233123-1331011211101221-0301030321222102"></a>

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
skip_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130122203320120-2301300100003113-1120303332111031-2120212003232230-2101012122121033-3101022300111302-1103301132101300-2121103321313120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-011.md#canonical-1300210222332112-3202232332012200-0000013303132120-2201221030000121-1230122120233211-1001120020312010-0300212200032021-0203310121102333)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active

<a id="canonical-3313322203033232-0021210333312033-2121202231222023-2003132200003222-3133320211223311-1100203132121003-1232200100101110-0201120313122112"></a>

Type: `"object"`. single nested block, Optional.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031112322332131-1110112023320002-1001002313211210-0220223133033210-1332102202210133-1231231320002030-3311222202302031-3032120202302201"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active`

- [enforcement_block](resources--http_loadbalancer--reference--group-011.md#canonical-3203110020211132-1002032203111131-0312131123223303-0003113113013221-3212001101020002-3121221303002100-2030130301021323-0302222133223233): complete subsection reference.

- [enforcement_report](resources--http_loadbalancer--reference--group-011.md#canonical-3311300202030203-1313203103021011-3210232223032300-2212323132131031-0213203200010111-1320313112322210-3203102130322133-3120200020000200): complete subsection reference.

<a id="canonical-0030111111201102-2101010113203323-1113112112133012-2333210333320122-3330001110221011-0120320300012203-0203131031320001-1003321133231310"></a>

<a id="canonical-1031211011012213-1313233131112130-0310230110000211-1110331022212012-2131101130310320-1302102112103323-1123121121000302-2333300023212223"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.request_validation_properties` property

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3203110020211132-1002032203111131-0312131123223303-0003113113013221-3212001101020002-3121221303002100-2030130301021323-0302222133223233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-011.md#canonical-1300210222332112-3202232332012200-0000013303132120-2201221030000121-1230122120233211-1001120020312010-0300212200032021-0203310121102333)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-011.md#canonical-3130122203320120-2301300100003113-1120303332111031-2120212003232230-2101012122121033-3101022300111302-1103301132101300-2121103321313120)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-2102030002200000-0301132221120210-1112110102111231-0110010131230200-3132132330331221-1303202031312132-3033230332222031-2033100121322232"></a>

Type: `["object", {}]`. Optional.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

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
enforcement_block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311300202030203-1313203103021011-3210232223032300-2212323132131031-0213203200010111-1320313112322210-3203102130322133-3120200020000200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-011.md#canonical-1300210222332112-3202232332012200-0000013303132120-2201221030000121-1230122120233211-1001120020312010-0300212200032021-0203310121102333)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-011.md#canonical-3130122203320120-2301300100003113-1120303332111031-2120212003232230-2101012122121033-3101022300111302-1103301132101300-2121103321313120)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-3030222331202031-3222300123323313-0322333133123013-3331102310121013-1110013131222322-1300012213101102-0011031200213112-0330322203033223"></a>

Type: `["object", {}]`. Optional.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

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
enforcement_report = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120033000331323-2023121101030022-1120013320021212-2032223212012232-2012300302330330-2321131313020133-0332322003002322-2231033320320223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- api_specification.validation_custom_list.settings

<a id="canonical-0101031200213302-0223322310230020-2203112301311020-3202323031111100-3032111322323212-3022002033333123-1103111300121211-0022311132130212"></a>

Type: `"object"`. single nested block, Optional.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Additional upstream details:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

Terraform syntax:

```terraform
settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202120021123102-0132321322010113-2110110002212321-3121013113101232-1322002011301203-1322220310202103-2220113202211300-2231312023201132"></a>

### Direct properties for `api_specification.validation_custom_list.settings`

- [oversized_body_fail_validation](resources--http_loadbalancer--reference--group-011.md#canonical-3113101131302132-2021230011110203-1032121102300301-0120031130300012-1310020320201200-0111130213332030-0330220330102131-0212131331133300): complete subsection reference.

- [oversized_body_skip_validation](resources--http_loadbalancer--reference--group-011.md#canonical-0203330231032200-0130312221120333-0222301312113211-2213013113103210-2110302032233320-1222320221302333-3122030030133110-2200033201100221): complete subsection reference.

- [property_validation_settings_custom](resources--http_loadbalancer--reference--group-011.md#canonical-1020003303011023-0303010032123322-1000320033300202-0001221320210101-1131301230232330-3230303021121101-1332103202230103-2013232231303212): complete subsection reference.

- [property_validation_settings_default](resources--http_loadbalancer--reference--group-011.md#canonical-3203131121031232-3120013032202303-3300322030100202-2030302120033210-1313320021310332-2012231123013010-2220033321013202-0323013203003231): complete subsection reference.

<a id="canonical-3113101131302132-2021230011110203-1032121102300301-0120031130300012-1310020320201200-0111130213332030-0330220330102131-0212131331133300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.oversized_body_fail_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-011.md#canonical-2120033000331323-2023121101030022-1120013320021212-2032223212012232-2012300302330330-2321131313020133-0332322003002322-2231033320320223)
- api_specification.validation_custom_list.settings.oversized_body_fail_validation

<a id="canonical-1323302013030302-3013003220103111-2010100322131220-3210123010010232-1232122002332230-1212021233023212-0103301223300023-3220313222000222"></a>

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
oversized_body_fail_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203330231032200-0130312221120333-0222301312113211-2213013113103210-2110302032233320-1222320221302333-3122030030133110-2200033201100221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.oversized_body_skip_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-011.md#canonical-2120033000331323-2023121101030022-1120013320021212-2032223212012232-2012300302330330-2321131313020133-0332322003002322-2231033320320223)
- api_specification.validation_custom_list.settings.oversized_body_skip_validation

<a id="canonical-3213130020323100-0223233232210111-3213210230103101-1103221112210020-1111033201111110-3022130331232223-0102201023031121-0213211213001321"></a>

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
oversized_body_skip_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020003303011023-0303010032123322-1000320033300202-0001221320210101-1131301230232330-3230303021121101-1332103202230103-2013232231303212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-011.md#canonical-2120033000331323-2023121101030022-1120013320021212-2032223212012232-2012300302330330-2321131313020133-0332322003002322-2231033320320223)
- api_specification.validation_custom_list.settings.property_validation_settings_custom

<a id="canonical-2102113122320010-0312002103311103-0000311113332103-3231002231333301-1003131110121320-1213201323102312-2020110010230311-3011113333003221"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for property validation settings custom.

Additional upstream details:

Custom property validation settings.

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
property_validation_settings_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311020323003221-3010020133120110-1132001102303300-0311221110302220-3120012212211323-2030130111101001-2323032111303200-0002221210013003"></a>

### Direct properties for `api_specification.validation_custom_list.settings.property_validation_settings_custom`

- [query_parameters](resources--http_loadbalancer--reference--group-011.md#canonical-3021030320203033-0223013023100133-2213202322202003-2203303231023101-3203013330020331-1030233301212013-3231230321302033-2333002103031122): complete subsection reference.

<a id="canonical-3021030320203033-0223013023100133-2213202322202003-2203303231023101-3203013330020331-1030233301212013-3231230321302033-2333002103031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-011.md#canonical-2120033000331323-2023121101030022-1120013320021212-2032223212012232-2012300302330330-2321131313020133-0332322003002322-2231033320320223)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-011.md#canonical-1020003303011023-0303010032123322-1000320033300202-0001221320210101-1131301230232330-3230303021121101-1332103202230103-2013232231303212)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters

<a id="canonical-0331212322100103-0111331203021210-2332233201122221-0310212333200211-3312303131033120-3032211302030313-3022123113301232-3202111312332202"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202231330010223-2100312211101213-2213211230323312-1021021131021330-3002011222013133-2001330112110120-1202223210332101-1113110033212212"></a>

### Direct properties for `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters`

- [allow_additional_parameters](resources--http_loadbalancer--reference--group-011.md#canonical-2111113112313201-1011231220101133-3222210113100302-0213031213221111-3333312113122330-3330130010211030-1222333213012223-2320220203322001): complete subsection reference.

- [disallow_additional_parameters](resources--http_loadbalancer--reference--group-011.md#canonical-3222300222001200-1021012323212003-3303212230220112-1032300310112223-2133231110323321-3212220313200301-1333103233003233-0033313031232333): complete subsection reference.

<a id="canonical-2111113112313201-1011231220101133-3222210113100302-0213031213221111-3333312113122330-3330130010211030-1222333213012223-2320220203322001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-011.md#canonical-2120033000331323-2023121101030022-1120013320021212-2032223212012232-2012300302330330-2321131313020133-0332322003002322-2231033320320223)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-011.md#canonical-1020003303011023-0303010032123322-1000320033300202-0001221320210101-1131301230232330-3230303021121101-1332103202230103-2013232231303212)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-011.md#canonical-3021030320203033-0223013023100133-2213202322202003-2203303231023101-3203013330020331-1030233301212013-3231230321302033-2333002103031122)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-1323332113001302-1022121332000010-3000331030303303-1001213012221211-3011212002131220-1220111221301020-3233210032332322-0003121133201121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow additional parameters.

Terraform syntax:

```terraform
allow_additional_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222300222001200-1021012323212003-3303212230220112-1032300310112223-2133231110323321-3212220313200301-1333103233003233-0033313031232333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-011.md#canonical-2120033000331323-2023121101030022-1120013320021212-2032223212012232-2012300302330330-2321131313020133-0332322003002322-2231033320320223)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-011.md#canonical-1020003303011023-0303010032123322-1000320033300202-0001221320210101-1131301230232330-3230303021121101-1332103202230103-2013232231303212)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-011.md#canonical-3021030320203033-0223013023100133-2213202322202003-2203303231023101-3203013330020331-1030233301212013-3231230321302033-2333002103031122)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-1310101231010321-1101311103230323-1231300323033203-0010101110013033-0211230131032312-3021011102322012-1313012133110110-0311222311203002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disallow additional parameters.

Terraform syntax:

```terraform
disallow_additional_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203131121031232-3120013032202303-3300322030100202-2030302120033210-1313320021310332-2012231123013010-2220033321013202-0323013203003231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_default` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-011.md#canonical-2120033000331323-2023121101030022-1120013320021212-2032223212012232-2012300302330330-2321131313020133-0332322003002322-2231033320320223)
- api_specification.validation_custom_list.settings.property_validation_settings_default

<a id="canonical-2202102321123310-1222212211323013-1333210132132010-0102323031213000-0001310211131223-1013132011230301-1201312333001320-3131322110103300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for property validation settings default.

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
property_validation_settings_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313322031101011-2030333022323320-1122212003130032-1133211032021000-3320233120012101-2233031012300132-0112202012322000-0031131021323032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- api_specification.validation_disabled

<a id="canonical-2123122102131131-0000032230010112-0022102333010221-2313001010232320-0312131023133002-3311221112321231-1123023301012222-3013323003310211"></a>

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
validation_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- api_testing

<a id="canonical-3232323311203001-3210130000000231-3232333100201001-0130312001022131-3303101331113220-2213330231331122-0232302010032112-3320123123122023"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_testing, disable\_api\_testing; Default: disable\_api\_testing\] API Testing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-frequency_choice": "[\"every_day\",\"every_month\",\"every_week\"]"
}
```

OneOf alternatives in this subsection:

- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-3232323311203001-3210130000000231-3232333100201001-0130312001022131-3303101331113220-2213330231331122-0232302010032112-3320123123122023)
- [disable_api_testing](resources--http_loadbalancer--reference--group-017.md#canonical-2102333133332200-1121021111010220-1122302223121000-2100012203322001-3000221313331130-1113021300210303-2123021303330312-2201022000213330)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_testing {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103201211223333-2202003303021303-1003032300332332-3002331203001120-1212330302213320-2130032332103033-3101113000231201-1332312031330221"></a>

### Direct properties for `api_testing`

<a id="canonical-1003133112223113-0133332023000332-1011020003210112-2222313003131213-2202330223102011-1101022022103213-1230201201313000-0301132120003102"></a>

#### `api_testing.custom_header_value` property

Type: `"string"`. Optional.

Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022): complete subsection reference.

- [every_day](resources--http_loadbalancer--reference--group-011.md#canonical-1233112201020322-2222030330013333-1023011223312313-0102332131132022-3310303122320110-1030121211022030-0233310032132321-1133031012110031): complete subsection reference.

- [every_month](resources--http_loadbalancer--reference--group-011.md#canonical-1331332012233330-1220333203011120-1230312030310000-1331012303202002-0001133000011230-3032330101333221-0121211313220102-0212301130222100): complete subsection reference.

- [every_week](resources--http_loadbalancer--reference--group-011.md#canonical-0032022203110230-3332223021213130-2030021102321203-1200320200112001-2133021102333131-3110131220033213-0010303232311313-2100300200333312): complete subsection reference.

<a id="canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- api_testing.domains

<a id="canonical-3122233123223300-3133132333323230-0311133222010212-3211303112230023-2110302210133003-0212330320111110-3011102310300220-3221312001023020"></a>

Type: `"object"`. list nested block, Optional.

Add and configure testing domains and credentials.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303210120012210-1332102210231002-2221312020031312-2103312333111112-3232220103211223-3031102022201223-3111321020222302-1022002011301030"></a>

### Direct properties for `api_testing.domains`

<a id="canonical-3132133213231320-0300311321130011-3221300123312230-1221000333323321-2212312011310031-3103100203331001-2123201011033132-1210232321023331"></a>

#### `api_testing.domains.allow_destructive_methods` property

Type: `"bool"`. Optional.

Enable to allow API Testing to execute against destructive methods. Use with caution as these may
modify or DELETE data.

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

- [credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312): complete subsection reference.

<a id="canonical-1331210000031011-1233321101113030-1010200310121031-2103303010101232-3300123301020112-3030201123311120-0112332212231002-0311111321311102"></a>

<a id="canonical-1323101123110330-0200300022021133-3033112211130113-3122213131113032-0123111000100102-3001223021121212-3231312131210101-2230100202101133"></a>

#### `api_testing.domains.domain` property

Type: `"string"`. Optional.

Add your testing environment domain. Be aware that running tests on a production domain can impact
live applications, as API testing cannot distinguish between production and testing environments.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- api_testing.domains.credentials

<a id="canonical-1230211121313200-0220003211012021-3031221203222300-3230210003122330-3331312022322001-0112013310012002-0121200300000033-2023030031030213"></a>

Type: `"object"`. list nested block, Optional.

Add credentials for API testing to use in the selected environment.

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
credentials {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210011321323111-3022033013132200-1233303201100310-3331110100130122-1222310120200000-2210021123101023-0100033231023122-1011130303302100"></a>

### Direct properties for `api_testing.domains.credentials`

- [admin](resources--http_loadbalancer--reference--group-011.md#canonical-1322011302331332-2223020220013000-0220232333032030-2010120133020302-1333232300133103-1010303111233302-1101301013203131-0001110302321011): complete subsection reference.

- [api_key](resources--http_loadbalancer--reference--group-011.md#canonical-0103031311032213-0332300131230220-2021203001321003-1232212031313201-0212233013102021-3101303022303103-1032221301110012-3110131021231130): complete subsection reference.

- [basic_auth](resources--http_loadbalancer--reference--group-011.md#canonical-2322122232020022-3212320201032301-0230021210211012-2023311320030303-2111131003303120-2300132100013223-1020330112130333-3200103000313120): complete subsection reference.

- [bearer_token](resources--http_loadbalancer--reference--group-011.md#canonical-1212030332010323-2120332003311231-1203023001233020-2301122132222113-2030003021021212-2121120303332001-0323203211110113-3103130011231130): complete subsection reference.

<a id="canonical-3130130220212222-1131313223320031-2330012122213000-0122303112323312-1011023212010322-0231311100101020-3301100103222010-1301311211020132"></a>

<a id="canonical-0111203323332333-3023002311123000-1310310021022200-2013001211210320-1103211031313223-0333311312310323-3322232032202022-0232200112232030"></a>

#### `api_testing.domains.credentials.credential_name` property

Type: `"string"`. Optional.

Enter a unique name for the credentials used in API testing.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [login_endpoint](resources--http_loadbalancer--reference--group-011.md#canonical-1130011202302313-0022131322001022-1000230030131212-1222102303330221-2122122200200111-2012122330022110-0103231113102030-2013233123311101): complete subsection reference.

- [standard](resources--http_loadbalancer--reference--group-011.md#canonical-0002022321301120-1333223102000123-0203112001030121-0002110032101120-1221113310302033-2220213223321221-1010221221211200-3123223032322002): complete subsection reference.

<a id="canonical-1322011302331332-2223020220013000-0220232333032030-2010120133020302-1333232300133103-1010303111233302-1101301013203131-0001110302321011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.admin` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- api_testing.domains.credentials.admin

<a id="canonical-2012102302121031-2001221302020101-2003011230033302-1131231020332232-0101022032323230-1302000130032002-3013331222000312-0023212202322033"></a>

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
admin = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103031311032213-0332300131230220-2021203001321003-1232212031313201-0212233013102021-3101303022303103-1032221301110012-3110131021231130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.api_key` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- api_testing.domains.credentials.api_key

<a id="canonical-2333101033232030-0130132001221101-1102313330130310-3232003003031113-0322230122331002-0233122313201102-0012233332133003-0330003222201300"></a>

Type: `"object"`. single nested block, Optional.

API Key

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
api_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121331002323330-1221022102133030-1121313302012223-2220210300021131-2212321123213303-3330020111012011-2010111113020121-2310030201001233"></a>

### Direct properties for `api_testing.domains.credentials.api_key`

<a id="canonical-2302023033303231-2320330223032000-1020132223232001-3211131231030011-3331000303120001-0021130202311102-0212223211213112-2023203300211303"></a>

#### `api_testing.domains.credentials.api_key.key` property

Type: `"string"`. Optional.

Key. Cryptographic key material

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [value](resources--http_loadbalancer--reference--group-011.md#canonical-2322332300011132-0300122113000302-2231310231010220-0031023031010332-2203011200100131-3112121023121102-3203301102030011-2132313022222321): complete subsection reference.

<a id="canonical-2322332300011132-0300122113000302-2231310231010220-0031023031010332-2203011200100131-3112121023121102-3203301102030011-2132313022222321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.api_key.value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.api_key](resources--http_loadbalancer--reference--group-011.md#canonical-0103031311032213-0332300131230220-2021203001321003-1232212031313201-0212233013102021-3101303022303103-1032221301110012-3110131021231130)
- api_testing.domains.credentials.api_key.value

<a id="canonical-1013123100102031-2313200013033130-3313112033233020-1023030121232300-3201221211213011-2102322130231333-0133120003132102-1130233103310030"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
value {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310031313010230-1131010322101301-0113001211231010-0131313200033010-3300101021102213-0010100011220020-2200230112123212-0002100322221200"></a>

### Direct properties for `api_testing.domains.credentials.api_key.value`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-011.md#canonical-3133203113202101-2203223323022010-1132102213013023-1330313331102211-1003011011330311-2013220013031123-2233001200230003-2010213011301312): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-011.md#canonical-3321200123111033-0121230003131213-3122231033023002-3033020331023120-2121033001312303-2011231323233020-2110013110223101-2133003303222232): complete subsection reference.

<a id="canonical-3133203113202101-2203223323022010-1132102213013023-1330313331102211-1003011011330311-2013220013031123-2233001200230003-2010213011301312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.api_key.value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.api_key](resources--http_loadbalancer--reference--group-011.md#canonical-0103031311032213-0332300131230220-2021203001321003-1232212031313201-0212233013102021-3101303022303103-1032221301110012-3110131021231130)
- [api_testing.domains.credentials.api_key.value](resources--http_loadbalancer--reference--group-011.md#canonical-2322332300011132-0300122113000302-2231310231010220-0031023031010332-2203011200100131-3112121023121102-3203301102030011-2132313022222321)
- api_testing.domains.credentials.api_key.value.blindfold_secret_info

<a id="canonical-1321233023033233-1000310310211113-2203021201033230-1132211021000103-2330132012201102-1100011312331020-2301132322202030-3323121200313011"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002300100320012-3030211113211112-1020112300303011-1012100110231311-2112101033120123-2312333311030220-1122313313101312-3031110201221112"></a>

### Direct properties for `api_testing.domains.credentials.api_key.value.blindfold_secret_info`

<a id="canonical-0031333203213331-0312323003210223-1330110310201023-1221132132330110-0021230131330013-3103100002031210-0312002320130003-2201010021231103"></a>

#### `api_testing.domains.credentials.api_key.value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1310013012233223-1332123210133022-3322320212110102-3232122003222102-0102020312021113-0203001233121021-1112201313223120-0022303302033211"></a>

<a id="canonical-3313330123302220-1000021210303220-1220232203202321-1133001010023311-3212100103003322-0301303210130202-0320332033230003-1112201000322013"></a>

#### `api_testing.domains.credentials.api_key.value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1211030020303022-2231121312203022-0103322123331313-2230220203201230-0301232003121302-0212131221011330-2003211201012302-2322232313323301"></a>

<a id="canonical-3220012210001303-2320302301022130-0112111103103301-2010010010213012-3013002030321313-3002100032221122-3022102000112013-1213002022023001"></a>

#### `api_testing.domains.credentials.api_key.value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3321200123111033-0121230003131213-3122231033023002-3033020331023120-2121033001312303-2011231323233020-2110013110223101-2133003303222232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.api_key.value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.api_key](resources--http_loadbalancer--reference--group-011.md#canonical-0103031311032213-0332300131230220-2021203001321003-1232212031313201-0212233013102021-3101303022303103-1032221301110012-3110131021231130)
- [api_testing.domains.credentials.api_key.value](resources--http_loadbalancer--reference--group-011.md#canonical-2322332300011132-0300122113000302-2231310231010220-0031023031010332-2203011200100131-3112121023121102-3203301102030011-2132313022222321)
- api_testing.domains.credentials.api_key.value.clear_secret_info

<a id="canonical-0213222113132222-2300321211001030-1320021303203210-1112122023332331-1313321133313223-0031110311101332-3013200012122232-3020031130200221"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132002110232112-2313110302220232-0130103131002031-1231212303111103-3312210310011212-2310031032320031-1101312022230231-0120121002110222"></a>

### Direct properties for `api_testing.domains.credentials.api_key.value.clear_secret_info`

<a id="canonical-1022221303230331-2011230021332130-0131001331003133-3302001310012133-3333101033103322-0311201320213302-3203130331121321-2222320231001001"></a>

#### `api_testing.domains.credentials.api_key.value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1221133232033232-0203020212133322-1332030001331032-2002130121032123-2021220321112102-3331223212032312-0320003320110112-0302122000101101"></a>

<a id="canonical-2122021002201230-1210121311200302-2132311133012031-2031011221013200-0023111113213320-3021323331101223-1232132033321323-1301331002213323"></a>

#### `api_testing.domains.credentials.api_key.value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2322122232020022-3212320201032301-0230021210211012-2023311320030303-2111131003303120-2300132100013223-1020330112130333-3200103000313120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.basic_auth` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- api_testing.domains.credentials.basic_auth

<a id="canonical-3221230232100212-1003220212022110-2021331333302021-1121330331133301-3321232320220203-2100012210010331-0130200332010000-3232212301121201"></a>

Type: `"object"`. single nested block, Optional.

Basic Authentication.

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
basic_auth {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332111033300302-1200202022212203-3202031130213020-0210311033021120-2203302323333201-2030013001011122-0011000233101010-1213031102212232"></a>

### Direct properties for `api_testing.domains.credentials.basic_auth`

- [password](resources--http_loadbalancer--reference--group-011.md#canonical-1302233020313203-2330010203220112-1331103101130332-0011020212000222-0102333132211212-2331102330132111-0331032201202330-0331301300210301): complete subsection reference.

<a id="canonical-1100112020002232-0332102121110030-1031210001120303-1003131332220031-0012023301010202-0200132012031213-2012330200333003-1310321100301211"></a>

<a id="canonical-1101031323231111-3030012212301313-3020311002200120-2313232202133222-0330120122132233-0301311103103032-3130230023000203-3313223310333130"></a>

#### `api_testing.domains.credentials.basic_auth.user` property

Type: `"string"`. Optional.

User. Configuration parameter for user

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1302233020313203-2330010203220112-1331103101130332-0011020212000222-0102333132211212-2331102330132111-0331032201202330-0331301300210301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.basic_auth.password` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.basic_auth](resources--http_loadbalancer--reference--group-011.md#canonical-2322122232020022-3212320201032301-0230021210211012-2023311320030303-2111131003303120-2300132100013223-1020330112130333-3200103000313120)
- api_testing.domains.credentials.basic_auth.password

<a id="canonical-1102322033102100-1110332013231330-2010001010101111-3331123330303132-0230103313322003-2213223300102002-1200012033212322-3210332111033300"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032120323103203-3133013011132100-0213220220212202-0313021312012101-2330302222200122-1000100233130010-0100233202223001-2300311213220222"></a>

### Direct properties for `api_testing.domains.credentials.basic_auth.password`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-011.md#canonical-1201220103111011-3310100112001302-3332113102233112-2131033012013021-0011031312200030-3233110322131311-0302210131033031-2120000110213033): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-011.md#canonical-0223023202001333-1201000013202033-0210213020213323-2010001313311103-0322121031100200-3213221001323313-3233201003321332-2100023220022110): complete subsection reference.

<a id="canonical-1201220103111011-3310100112001302-3332113102233112-2131033012013021-0011031312200030-3233110322131311-0302210131033031-2120000110213033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.basic_auth.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.basic_auth](resources--http_loadbalancer--reference--group-011.md#canonical-2322122232020022-3212320201032301-0230021210211012-2023311320030303-2111131003303120-2300132100013223-1020330112130333-3200103000313120)
- [api_testing.domains.credentials.basic_auth.password](resources--http_loadbalancer--reference--group-011.md#canonical-1302233020313203-2330010203220112-1331103101130332-0011020212000222-0102333132211212-2331102330132111-0331032201202330-0331301300210301)
- api_testing.domains.credentials.basic_auth.password.blindfold_secret_info

<a id="canonical-0311331322332332-3121020302300000-1333132310110231-0110033110232323-2320222211030231-0332210130003110-0112201312232033-1020120021002103"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311321333300100-1202323130333031-1323000100122030-3333202020003203-1203310103200013-2312123223010112-0102230233232120-2111030301133320"></a>

### Direct properties for `api_testing.domains.credentials.basic_auth.password.blindfold_secret_info`

<a id="canonical-3221000202102032-1300120122022330-0120312333211112-0321023000233331-1120100111132300-3023023211100311-0110311103103112-1311121201012101"></a>

#### `api_testing.domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0211232000100120-2002223120322100-0331111112202331-3021100001003030-3032230311213002-0231201102012303-3230321200100203-2323033300021022"></a>

<a id="canonical-3222031313033333-0223213321002100-3033312121212302-3331233333332132-0211022001013121-1221303313210111-2313120210120123-0030220013002301"></a>

#### `api_testing.domains.credentials.basic_auth.password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3130001300221333-3010002312311030-2111101221311101-0110033000200221-0303030020322020-0100011021201002-3313313020323130-3123303210320103"></a>

<a id="canonical-3231031122003320-2100301233303330-2002021233320001-0113201331033203-3331223132300323-1232210010320122-1312322220220120-2032100320102333"></a>

#### `api_testing.domains.credentials.basic_auth.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0223023202001333-1201000013202033-0210213020213323-2010001313311103-0322121031100200-3213221001323313-3233201003321332-2100023220022110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.basic_auth.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.basic_auth](resources--http_loadbalancer--reference--group-011.md#canonical-2322122232020022-3212320201032301-0230021210211012-2023311320030303-2111131003303120-2300132100013223-1020330112130333-3200103000313120)
- [api_testing.domains.credentials.basic_auth.password](resources--http_loadbalancer--reference--group-011.md#canonical-1302233020313203-2330010203220112-1331103101130332-0011020212000222-0102333132211212-2331102330132111-0331032201202330-0331301300210301)
- api_testing.domains.credentials.basic_auth.password.clear_secret_info

<a id="canonical-0130032301010233-0120223110101021-3121030311021030-2211312002301103-1122222033233322-2300020130100212-0231322121202031-2322000221223033"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311002110332213-2023111130033102-1010132233133031-2000030203232121-2320230221210020-2002202313033020-0032300303211201-3322010333131021"></a>

### Direct properties for `api_testing.domains.credentials.basic_auth.password.clear_secret_info`

<a id="canonical-3222131103023213-1102321033302120-3103220021210022-2122333221130313-3321211233013012-3302231312001003-2123100323021030-2233313130221323"></a>

#### `api_testing.domains.credentials.basic_auth.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2203111313110222-1122030212332121-0302131022122200-3223212312330320-3120031000021212-2120130323131321-2200331321313021-2310200233211331"></a>

<a id="canonical-0132303320122103-2303130000302202-3331132131322030-1213221103300213-2111130003223023-3303131210320302-1311321332212032-1230201232020113"></a>

#### `api_testing.domains.credentials.basic_auth.password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1212030332010323-2120332003311231-1203023001233020-2301122132222113-2030003021021212-2121120303332001-0323203211110113-3103130011231130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.bearer_token` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- api_testing.domains.credentials.bearer_token

<a id="canonical-3033131331230030-2013101031310310-1031231231212001-2132322213201323-2230023222321110-2132011121233032-0221100130331322-3323220300000210"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bearer token.

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
bearer_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131011231320213-3010322100212113-2211212031010211-3102320103102112-2322123201101020-3032022133231110-1133132310030301-1112031221310101"></a>

### Direct properties for `api_testing.domains.credentials.bearer_token`

- [token](resources--http_loadbalancer--reference--group-011.md#canonical-3223203032111230-2332221333133011-0012311033133223-3010031123130122-2111202011233232-0303301122313233-3020123022030000-3220203310330102): complete subsection reference.

<a id="canonical-3223203032111230-2332221333133011-0012311033133223-3010031123130122-2111202011233232-0303301122313233-3020123022030000-3220203310330102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.bearer_token.token` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.bearer_token](resources--http_loadbalancer--reference--group-011.md#canonical-1212030332010323-2120332003311231-1203023001233020-2301122132222113-2030003021021212-2121120303332001-0323203211110113-3103130011231130)
- api_testing.domains.credentials.bearer_token.token

<a id="canonical-0333302122002303-3221120103331133-1121103331110221-3200133313320031-3221100301332220-1011200301203113-1003323123031002-3000201111111000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
token {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002001220220012-3333012200013220-2133221222100310-3100021032212213-3002122321202313-0001032310332223-2030003331112013-1012123103003302"></a>

### Direct properties for `api_testing.domains.credentials.bearer_token.token`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-011.md#canonical-1230123223000203-0311220022302323-1021122001001103-0033333031232013-0333102021100121-2201231033113121-3021113330033221-2011333013030331): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-011.md#canonical-2023200201111102-3203311212112230-1230200222302223-3321203233111200-0001133323232120-0210111020330312-0111003211332220-0132132212133120): complete subsection reference.

<a id="canonical-1230123223000203-0311220022302323-1021122001001103-0033333031232013-0333102021100121-2201231033113121-3021113330033221-2011333013030331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.bearer_token.token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.bearer_token](resources--http_loadbalancer--reference--group-011.md#canonical-1212030332010323-2120332003311231-1203023001233020-2301122132222113-2030003021021212-2121120303332001-0323203211110113-3103130011231130)
- [api_testing.domains.credentials.bearer_token.token](resources--http_loadbalancer--reference--group-011.md#canonical-3223203032111230-2332221333133011-0012311033133223-3010031123130122-2111202011233232-0303301122313233-3020123022030000-3220203310330102)
- api_testing.domains.credentials.bearer_token.token.blindfold_secret_info

<a id="canonical-1122020231230303-2232102102101333-2300311231333230-1311103323130203-1110300201002223-2213101201020112-1100131031301201-0233033203222100"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200300130222221-3013030301030111-2123312310300112-1022323221030102-3230033030032011-1332321333213222-3000233311303203-3013321120010222"></a>

### Direct properties for `api_testing.domains.credentials.bearer_token.token.blindfold_secret_info`

<a id="canonical-1213101203330321-0103201323202023-0011322212012002-2003202110001132-3031321021221011-0333130131331102-2211101323310203-1021300313332323"></a>

#### `api_testing.domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2220132020233332-1022102121120321-0000113232221332-0322331230122021-1123003321320231-2013312210312011-2102310303213333-1001201002330113"></a>

<a id="canonical-3212212303000222-2003103233231020-1112233102230312-1311002230333313-2112201032202001-3331212120103010-2310320332333001-2010003030213121"></a>

#### `api_testing.domains.credentials.bearer_token.token.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3033202221302132-3102032300310121-1322132301001321-0033123232120231-3122000211333022-3231222030111030-1012232321032012-0220312110013233"></a>

<a id="canonical-1011332000023222-1220013310022333-2333201102130013-1103222210030023-0322011011322112-0312222200332301-2313210333220100-1220331100113120"></a>

#### `api_testing.domains.credentials.bearer_token.token.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2023200201111102-3203311212112230-1230200222302223-3321203233111200-0001133323232120-0210111020330312-0111003211332220-0132132212133120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.bearer_token.token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.bearer_token](resources--http_loadbalancer--reference--group-011.md#canonical-1212030332010323-2120332003311231-1203023001233020-2301122132222113-2030003021021212-2121120303332001-0323203211110113-3103130011231130)
- [api_testing.domains.credentials.bearer_token.token](resources--http_loadbalancer--reference--group-011.md#canonical-3223203032111230-2332221333133011-0012311033133223-3010031123130122-2111202011233232-0303301122313233-3020123022030000-3220203310330102)
- api_testing.domains.credentials.bearer_token.token.clear_secret_info

<a id="canonical-0133023012001010-1321313000313131-3302133021212202-3000232312130200-3021322023203000-2300012202302111-3323012102000112-0303110322230032"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121321301130210-2201013333120103-1013221111012320-0110003203310332-0231220223223023-0101112120202132-1033010003012230-0221230113321032"></a>

### Direct properties for `api_testing.domains.credentials.bearer_token.token.clear_secret_info`

<a id="canonical-1023203222220220-3100133103321113-1012002121100000-3320233223102132-3112112031233102-1010020022131023-1312302100202203-2313022031031222"></a>

#### `api_testing.domains.credentials.bearer_token.token.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2122101111022212-3212220232220303-2133230320332320-2103312203320313-3101110013020003-2011321020002023-3213200112101103-3313001230210123"></a>

<a id="canonical-3110220211010313-1021203121010101-3112210322230023-3101021311120212-2313212012231332-0002100120303003-3123023201100310-2221300303121121"></a>

#### `api_testing.domains.credentials.bearer_token.token.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1130011202302313-0022131322001022-1000230030131212-1222102303330221-2122122200200111-2012122330022110-0103231113102030-2013233123311101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.login_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- api_testing.domains.credentials.login_endpoint

<a id="canonical-1232011003221003-3221200303011231-3212020313220113-3203010322011112-2201331202331001-3011300000231132-1012232200011323-1211132331312032"></a>

Type: `"object"`. single nested block, Optional.

Login Endpoint.

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
login_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200100220203021-0000311322301233-1212203023302123-0223102111133332-3101323112101023-3020333313001212-3233221333123310-3313213322130022"></a>

### Direct properties for `api_testing.domains.credentials.login_endpoint`

- [json_payload](resources--http_loadbalancer--reference--group-011.md#canonical-3221130131220320-2123002333113211-3221300230021010-2033011110202030-1212202322200031-2201321010023033-0012320112101213-0212220311112303): complete subsection reference.

<a id="canonical-3223333132221010-1123302331001201-2303320003111101-3003112011010312-3200001302201321-0002000331020030-3023232202323231-3123030023102000"></a>

<a id="canonical-1322000220322301-2032102333331112-2310331130201002-2233103330332020-0030321012233121-0330332212301132-3333332222122122-2102330333133332"></a>

#### `api_testing.domains.credentials.login_endpoint.method` property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3332110303213233-2303123011310020-3330130312030333-0333301003033301-3121222330213103-1032322223110122-2333020310123012-2101312013221212"></a>

<a id="canonical-2122021120213221-2323013021013212-1130122003112122-1031211333000103-1121311301303322-3311030323131200-2221033300330302-3023211121302001"></a>

#### `api_testing.domains.credentials.login_endpoint.path` property

Type: `"string"`. Optional.

Path. URL path for the endpoint

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  }
}
```

<a id="canonical-2110322131122320-1221120231311320-0030133123131223-2031200033012311-3331332022113323-2332030313130330-1113020321132320-1133031010302013"></a>

<a id="canonical-2122200003311220-1330120101322320-0110012331212203-1311031221220103-2031233320310331-3121122212320222-2022120102200001-0121020332301220"></a>

#### `api_testing.domains.credentials.login_endpoint.token_response_key` property

Type: `"string"`. Optional.

Specifies the key name used to extract the authentication token from the login response, such as
token or access\_token.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3221130131220320-2123002333113211-3221300230021010-2033011110202030-1212202322200031-2201321010023033-0012320112101213-0212220311112303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.login_endpoint.json_payload` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.login_endpoint](resources--http_loadbalancer--reference--group-011.md#canonical-1130011202302313-0022131322001022-1000230030131212-1222102303330221-2122122200200111-2012122330022110-0103231113102030-2013233123311101)
- api_testing.domains.credentials.login_endpoint.json_payload

<a id="canonical-2303113122310311-2321021230330121-3132013301330023-3010223320011111-3103330313222220-1110030102301030-1311303332013230-0001202230001200"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
json_payload {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101023032110323-3001030200001100-2130110132133212-2001223021010101-3202221332301100-0232003132300033-0031113111133222-0332113133020201"></a>

### Direct properties for `api_testing.domains.credentials.login_endpoint.json_payload`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-011.md#canonical-1330101011222221-1231012132302333-1221132133102121-0131001330133231-2133133112320113-1303212101020123-2020322303122200-0113230101210132): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-011.md#canonical-2322200123330322-3100122333331210-1031202030021233-2122221022120313-2002210001133222-0323330102102201-0122210213331021-0200131102023301): complete subsection reference.

<a id="canonical-1330101011222221-1231012132302333-1221132133102121-0131001330133231-2133133112320113-1303212101020123-2020322303122200-0113230101210132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.login_endpoint](resources--http_loadbalancer--reference--group-011.md#canonical-1130011202302313-0022131322001022-1000230030131212-1222102303330221-2122122200200111-2012122330022110-0103231113102030-2013233123311101)
- [api_testing.domains.credentials.login_endpoint.json_payload](resources--http_loadbalancer--reference--group-011.md#canonical-3221130131220320-2123002333113211-3221300230021010-2033011110202030-1212202322200031-2201321010023033-0012320112101213-0212220311112303)
- api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info

<a id="canonical-1110203133303030-2122123133022110-3122112303202010-0301011120230012-3220111303111001-2102322123132221-0020121032003233-2212332100000013"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232130112012110-0201021311331111-0230230020302223-0021321233003011-0231020022203300-2201231302202231-2321030333330010-3323312113123203"></a>

### Direct properties for `api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info`

<a id="canonical-3002023200320130-1111222300221111-0020320131313331-2003012130023201-1133120110123312-0112131201130103-0113201210031321-3321201000223231"></a>

#### `api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3323202001112221-1203221000230123-1032202320013303-2231021310123030-2222110200023231-0022200333031230-1002320201321200-2121331002313001"></a>

<a id="canonical-0221003013311022-2100210021232002-3011100321221101-1111212201332201-2032010202200321-0012321332111312-3201110320122230-2033331033330000"></a>

#### `api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1203331233122332-1323310103303331-2120310310300112-3233313122122202-2311012221101331-0221003013213220-1221321121323312-2133010102133213"></a>

<a id="canonical-1122003223223120-1231111012031100-2231110122120232-1311302220302031-1113222111302332-0033333112132022-3300331303033000-0123021200233131"></a>

#### `api_testing.domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2322200123330322-3100122333331210-1031202030021233-2122221022120313-2002210001133222-0323330102102201-0122210213331021-0200131102023301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- [api_testing.domains.credentials.login_endpoint](resources--http_loadbalancer--reference--group-011.md#canonical-1130011202302313-0022131322001022-1000230030131212-1222102303330221-2122122200200111-2012122330022110-0103231113102030-2013233123311101)
- [api_testing.domains.credentials.login_endpoint.json_payload](resources--http_loadbalancer--reference--group-011.md#canonical-3221130131220320-2123002333113211-3221300230021010-2033011110202030-1212202322200031-2201321010023033-0012320112101213-0212220311112303)
- api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info

<a id="canonical-0112230121022030-0123033123213202-3332301112332300-2012110032213320-2231230300033112-1212323012230302-1021221302232311-3133322311131131"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111020002033203-1021000133132122-3121202303211200-0303020333122013-2122001202033010-1110020212232012-3200232312230333-3320211212023112"></a>

### Direct properties for `api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info`

<a id="canonical-1122131223111110-1232021211332203-0033010230113330-2202120022002132-2010032213311020-1030132021320233-2003022110233311-3103230212022013"></a>

#### `api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3133321021333121-3023133221323121-3310102100030010-1002203213303003-3233201032101300-3203333233221212-0021023022332303-1321103322100302"></a>

<a id="canonical-3021003231020120-0032123110121002-0111023332201111-3123202333203331-0222020230220033-2000323321113100-0102300212022022-3313310132032012"></a>

#### `api_testing.domains.credentials.login_endpoint.json_payload.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0002022321301120-1333223102000123-0203112001030121-0002110032101120-1221113310302033-2220213223321221-1010221221211200-3123223032322002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.domains.credentials.standard` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- [api_testing.domains](resources--http_loadbalancer--reference--group-011.md#canonical-1220132300110212-3102001121032231-0000033120101103-3333323030113103-2302300311022213-3222330212301132-3011230011323203-2230301111033022)
- [api_testing.domains.credentials](resources--http_loadbalancer--reference--group-011.md#canonical-2312010002010331-0002233102233010-2003120031010323-1221302211302101-1021202031232032-2323230013133033-2201323133112203-1021001121330312)
- api_testing.domains.credentials.standard

<a id="canonical-0123333303030010-1133012222020322-3321000013132123-1122310133013322-0102222003323321-1232212111203002-1112323103202213-3301023101322001"></a>

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
standard = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233112201020322-2222030330013333-1023011223312313-0102332131132022-3310303122320110-1030121211022030-0233310032132321-1133031012110031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.every_day` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- api_testing.every_day

<a id="canonical-3301031011033033-3323320210032013-1231321121302010-1103120222201320-1102133113223303-0030220033213301-0131313331323032-1230200323312131"></a>

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
every_day = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331332012233330-1220333203011120-1230312030310000-1331012303202002-0001133000011230-3032330101333221-0121211313220102-0212301130222100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.every_month` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- api_testing.every_month

<a id="canonical-1003313302133220-1111133012133022-1231303102222002-2130113131320022-3012211030311212-3313323210032032-0012223332213001-0031110033101313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for every month.

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
every_month = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032022203110230-3332223021213130-2030021102321203-1200320200112001-2133021102333131-3110131220033213-0010303232311313-2100300200333312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_testing.every_week` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_testing](resources--http_loadbalancer--reference--group-011.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301)
- api_testing.every_week

<a id="canonical-2113020221302122-2111021230231210-2320221230033323-3033013121032311-2212331211200011-0202311100221030-2322233132223100-0112232112201122"></a>

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
every_week = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333332121111031-2232311201113212-1330231222032103-0232233320112212-0011002003013022-0302021111130330-0013131230131031-3111322301332033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_firewall` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- app_firewall

<a id="canonical-1033022001011122-1002131312313033-0023001220031130-1213230211331231-1323233310333010-0130133011112221-1030321213020202-3011222200231101"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: app\_firewall, disable\_waf; Default: disable\_waf\] Type establishes a direct reference
from one object(the referrer) to another(the referred). Such a reference is in form of
tenant/namespace/name.

Additional upstream details:

This type establishes a direct reference from one object(the referrer) to another(the referred).

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

- [app_firewall](resources--http_loadbalancer--reference--group-011.md#canonical-1033022001011122-1002131312313033-0023001220031130-1213230211331231-1323233310333010-0130133011112221-1030321213020202-3011222200231101)
- [disable_waf](resources--http_loadbalancer--reference--group-017.md#canonical-0021311132330031-3013103020211321-1223133002212101-2010122201101202-3101313331110313-2013021300302031-2213220033313203-3000113332031001)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212001012001300-1000203031213011-3311312213210113-3103130222100200-2222322212120103-1013211031301221-2123110220232331-1030303030131223"></a>

### Direct properties for `app_firewall`

<a id="canonical-3301022032121001-3113000023032013-3031223100232100-1020121301230300-2302303313002011-0122123223130131-3310130011022100-3020223122312332"></a>

#### `app_firewall.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3022230212310333-0023123132312233-2010212223102133-1110220310023223-3033332130121122-2210123011212132-1132232200203021-0202131002013121"></a>

<a id="canonical-3221320211232211-0020132332030233-3110021202321200-0132011313012013-1021010202021132-1301030022120103-1221120201123321-3201322020300311"></a>

#### `app_firewall.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3002021311212212-0031003203333203-0213332213303201-3120322221203322-1333033131101033-0103113313300030-1230210213321322-3331303003101123"></a>

<a id="canonical-2230312102312213-2120002322211032-3113230301211213-1313230322231131-1003311113111010-1313213213323210-3212133311011022-3131123233031023"></a>

#### `app_firewall.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- blocked_clients

<a id="canonical-2102010120100111-3021303003302301-1023210222200021-3022123022113121-3130121311232322-0030330010221221-2231112202130201-3013220013220220"></a>

Type: `"object"`. list nested block, Optional.

Define rules to block IP Prefixes or AS numbers.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
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
blocked_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100322010103020-3313321020233303-3323003011220101-1330221113332133-3002302003232023-2203012131120133-3000213023322102-3321103231123130"></a>

### Direct properties for `blocked_clients`

<a id="canonical-0031321111230022-1201303211023020-2000200102103122-1000033331132110-0133013131203301-2022323031033220-1202303303212200-2322021230230102"></a>

#### `blocked_clients.actions` property

Type: `["list", "string"]`. Optional.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0002311001101132-3332300111301311-1310320112020230-2332123023200101-3023120200121331-2112003331133301-1113122301233013-3313113001312033"></a>

<a id="canonical-1232102103031231-1222131233300221-3022302112222302-2031231212110131-2332032322211223-1320021211021133-3221002111301322-0123112230331220"></a>

#### `blocked_clients.as_number` property

Type: `"number"`. Optional.

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](resources--http_loadbalancer--reference--group-011.md#canonical-2111320002233121-1332032232333101-0220121330111113-0102332012121311-0012310101321021-3313202101231323-0023331030210211-0123313012123213): complete subsection reference.

<a id="canonical-2320201130332120-1300221301131312-3223032122123122-0020322320003131-0011033312010120-3303032321132100-1013120213222320-1322310333332130"></a>

<a id="canonical-1210330032303110-2220231020221200-0221112010223233-1311110112113100-2222202233332031-0130103001132332-1122213123133130-0313320210031130"></a>

#### `blocked_clients.expiration_timestamp` property

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](resources--http_loadbalancer--reference--group-011.md#canonical-0230221313302212-2300233200220021-0311223130100202-3122001333022211-1330120103113220-2311323302123200-2002021232133220-1003100000033000): complete subsection reference.

<a id="canonical-1331023203001203-0310210033002123-2321323332100232-0111021002302130-3212010121220311-0213313123232312-0122133321321312-3031121211133000"></a>

<a id="canonical-2032212323222033-0012020011331202-2211002022232232-0030032330300121-1301020102301103-2002220220112230-1323120032023031-3020202102231211"></a>

#### `blocked_clients.ip_prefix` property

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header IPv6\_prefix user\_identifier\] IPv4 prefix string.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0202131103133311-0332020302311110-0102031102132111-3303303020200012-1330331321301112-1012220201021231-0323002201103030-0222201321011332"></a>

<a id="canonical-0030032322001233-2322300112303210-2021222131310031-1020102100010110-2120002003333201-3033211211212121-2113120123321332-1013311313200333"></a>

#### `blocked_clients.ipv6_prefix` property

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-011.md#canonical-3110213322201011-0103123011100223-3112113313330113-1023002030300111-0033330301202031-3332132022000213-2222021202322010-1231013331301002): complete subsection reference.

- [skip_processing](resources--http_loadbalancer--reference--group-011.md#canonical-0130313120300133-3201120201321012-2121223223133003-2030221323001131-1003233221121103-3330032100110222-3012322122302023-2012002122023312): complete subsection reference.

<a id="canonical-2111221000002302-2110020301103111-2013212132320303-2321110211002210-1213301032032000-1312110302011121-1022323121210303-3103221122003122"></a>

<a id="canonical-1300221322103020-1201202121011032-3030311230203220-0233203110012102-3121020001220023-1001102101223200-0321102030111330-1300313201331203"></a>

#### `blocked_clients.user_identifier` property

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](resources--http_loadbalancer--reference--group-011.md#canonical-3101023120211000-3213013112130202-0320210320110301-0232020131122102-0122130233002131-3213203302322003-2100201130323002-0231223230232131): complete subsection reference.

<a id="canonical-2111320002233121-1332032232333101-0220121330111113-0102332012121311-0012310101321021-3313202101231323-0023331030210211-0123313012123213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.bot_skip_processing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-011.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- blocked_clients.bot_skip_processing

<a id="canonical-0112123000020201-3102131223033300-1102221113123201-3101001231320322-2330300301120322-0123313112231033-0212133202301011-2221221312212221"></a>

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
bot_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230221313302212-2300233200220021-0311223130100202-3122001333022211-1330120103113220-2311323302123200-2002021232133220-1003100000033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.http_header` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-011.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- blocked_clients.http_header

<a id="canonical-0001333302002210-0321313332331200-2131103011012020-2231330302110030-2130130312320132-0002111200321013-3230310022300021-0121022310200021"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Additional upstream details:

Request header name and value pairs.

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
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323030031203202-3103102333013122-2312123201231112-3102230100310002-0321303331000212-1031200122212100-0222331333022322-3011120102303023"></a>

### Direct properties for `blocked_clients.http_header`

- [headers](resources--http_loadbalancer--reference--group-011.md#canonical-3223333312331330-0133131100312023-1222202302023323-1020102321223011-3013101320311223-1210230303100011-2023201132302000-2100112133012232): complete subsection reference.

<a id="canonical-3223333312331330-0133131100312023-1222202302023323-1020102321223011-3013101320311223-1210230303100011-2023201132302000-2100112133012232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.http_header.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-011.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- [blocked_clients.http_header](resources--http_loadbalancer--reference--group-011.md#canonical-0230221313302212-2300233200220021-0311223130100202-3122001333022211-1330120103113220-2311323302123200-2002021232133220-1003100000033000)
- blocked_clients.http_header.headers

<a id="canonical-1033231112201131-3023130010111222-3033230301100322-1102310203002221-1012123100010111-2233021302121332-1312201302313112-3202233231022203"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
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

<a id="canonical-0033212103022000-1120221032312321-0313213333020202-1303030201003323-2210133331033003-1102231221333323-2132110331311001-1001033011230230"></a>

### Direct properties for `blocked_clients.http_header.headers`

<a id="canonical-0300131010210022-2010131123001330-2012021020311333-1102213213122203-2000233320120203-1313010111322311-1102233333123013-3301200233333100"></a>

#### `blocked_clients.http_header.headers.exact` property

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2020333103333221-0123330211231200-3203110330312313-2003310133033030-2321110220012000-2120130102011033-3003312003203310-1221021203000001"></a>

<a id="canonical-1201112303321033-0112110230322001-2323332202101333-1001123013001032-0102300330001003-3031031231031300-0103113330132102-0113230223203301"></a>

#### `blocked_clients.http_header.headers.invert_match` property

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-0331020132023011-2222010131203121-1122102012030230-0101310023032231-3211332213220320-0223223232113103-2011223232132323-3120023130113101"></a>

<a id="canonical-1301203133211032-3012220101012030-2311013230331003-1210022231122222-0200110323132020-3133113331112203-2323000032313232-1303223003202230"></a>

#### `blocked_clients.http_header.headers.name` property

Type: `"string"`. Optional.

Name. Name of the header.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2313231023220131-0200012201101321-2110032120121211-3031200221211322-1131333122012112-0002033222120123-2222313333020123-3033330321030202"></a>

<a id="canonical-3100100331120121-3022200330022023-2231230212122102-0121211320021233-1131001100233321-0113031130110222-2022201300003200-1213230121102301"></a>

#### `blocked_clients.http_header.headers.presence` property

Type: `"bool"`. Optional.

Exclusive with \[exact regular expression\] If true, check for presence of header.

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

<a id="canonical-2132323130130023-0302031110030033-1130200233002132-0203100321022203-0321003030122210-1230312112120123-1203302123011230-1012323312220011"></a>

<a id="canonical-2232202002210033-1130002313311330-1331233022301130-3302300103121023-2012002000212132-1203010113232221-3321221302311332-0111211222120202"></a>

#### `blocked_clients.http_header.headers.regex` property

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3110213322201011-0103123011100223-3112113313330113-1023002030300111-0033330301202031-3332132022000213-2222021202322010-1231013331301002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-011.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- blocked_clients.metadata

<a id="canonical-2313331102302312-2122131303312022-1010122223131202-3011102203102201-1101202033331220-1212213020101300-3033321021120021-0332100220010332"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203021111103121-3211322001302101-3122110333212030-0111202123300012-3110333320122130-2023303102123013-0230230320003001-1120011333003133"></a>

### Direct properties for `blocked_clients.metadata`

<a id="canonical-0220033212100310-0232203020102000-0110000131133103-3212323133130032-3320310130300330-0203322210333021-2222333112022232-3101332212103111"></a>

#### `blocked_clients.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-2103221223323010-0132112020001010-0131111200221332-0112123302112331-1000010112113321-2313210203333110-3330212132331332-1202003022002000"></a>

<a id="canonical-1300032231301020-2131120111021000-1131030300012122-1113301333101131-2200013302311102-1023313122011331-2022320033203233-2300301200033220"></a>

#### `blocked_clients.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0130313120300133-3201120201321012-2121223223133003-2030221323001131-1003233221121103-3330032100110222-3012322122302023-2012002122023312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.skip_processing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-011.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- blocked_clients.skip_processing

<a id="canonical-2021111323123203-2213312311102223-1113123022031013-0230022030121122-1033320103203010-1121032220222311-2213210300333102-1033211001210323"></a>

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
skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101023120211000-3213013112130202-0320210320110301-0232020131122102-0122130233002131-3213203302322003-2100201130323002-0231223230232131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [blocked_clients](resources--http_loadbalancer--reference--group-011.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112)
- blocked_clients.waf_skip_processing

<a id="canonical-0120122023220132-1130132311332022-0322100113100322-0030002213331320-2231232322222013-0232023323323011-1203100100032232-3100120233032300"></a>

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

<a id="canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- bot_defense

<a id="canonical-0313032320202312-3130032212302312-3120330303330321-1012133211302100-3231033302020103-2202000220123011-2231131012233312-2310002330302100"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bot\_defense, bot\_defense\_advanced\_protection, disable\_bot\_defense; Default:
disable\_bot\_defense\] Defines various configuration OPTIONS for Bot Defense Policy.

Additional upstream details:

This defines various configuration OPTIONS for Bot Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cors_support_choice": "[\"disable_cors_support\",\"enable_cors_support\"]"
}
```

OneOf alternatives in this subsection:

- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0313032320202312-3130032212302312-3120330303330321-1012133211302100-3231033302020103-2202000220123011-2231131012233312-2310002330302100)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-1331023323322330-0120310010111332-1123202023223302-3231330202031222-2112233201102033-2111320023203321-1120030130231113-3212030313203022)
- [disable_bot_defense](resources--http_loadbalancer--reference--group-017.md#canonical-0031322221201123-0210020223332001-1212312133213033-3030023312222333-1332200121023202-1033110032330120-1010322121131033-3331133222011133)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bot_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120032000222200-0010021300110020-2222231300231020-1303010302232303-2221233310311100-3010231221103130-2020300133301311-1031111113333330"></a>

### Direct properties for `bot_defense`

- [disable_cors_support](resources--http_loadbalancer--reference--group-011.md#canonical-1020123332231113-2020103222100113-2122232221332121-1211111212021302-3002223112023123-2130122130110110-3233211000012303-1000032320312103): complete subsection reference.

- [enable_cors_support](resources--http_loadbalancer--reference--group-011.md#canonical-2000023030332210-2100002021230202-3310100211301003-0000113131311231-0130023331301033-3223323212221100-3032133113102211-2211111313333003): complete subsection reference.

- [policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013): complete subsection reference.

<a id="canonical-0203323233332000-0321101112201023-3000100113032122-0003310023133130-3233231002310032-2100202033311131-0133332112223103-3110010120131320"></a>

<a id="canonical-1331321031323210-3103221200002103-1003021100022311-0010331021131100-1303101211231221-1330121302133332-3211323313001111-2330010302203233"></a>

#### `bot_defense.regional_endpoint` property

Type: `"string"`. Optional.

\[Enum: AUTO|US|EU|ASIA\] Defines a selection for Bot Defense region - AUTO: AUTO Automatic
selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA
Asia region. Possible values are \`AUTO\`, \`US\`, \`EU\`, \`ASIA\`. Defaults to \`AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "AUTO",
  "enum": [
    "AUTO",
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2320230333301021-0232013112223132-0103103201012023-1130313001100210-3231111030031332-2131030200101323-1132122123011223-3320232102122221"></a>

<a id="canonical-1132323332111010-1302231301011312-0110313332103002-1332033001213001-0312101210132013-0032323210120111-3113111100320133-0231231000200100"></a>

#### `bot_defense.timeout` property

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-1020123332231113-2020103222100113-2122232221332121-1211111212021302-3002223112023123-2130122130110110-3233211000012303-1000032320312103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.disable_cors_support` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- bot_defense.disable_cors_support

<a id="canonical-1020012310020232-3331230311031220-2223113120230001-2113320023302220-0200302113323310-2131222132211230-1321131011020000-2200032001311310"></a>

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
disable_cors_support = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000023030332210-2100002021230202-3310100211301003-0000113131311231-0130023331301033-3223323212221100-3032133113102211-2211111313333003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.enable_cors_support` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- bot_defense.enable_cors_support

<a id="canonical-2122011110300220-3121100330333311-0102112231122120-2312301000013210-2313310000132122-3002220001232101-1221203332110231-1123320202223133"></a>

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
enable_cors_support = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- bot_defense.policy

<a id="canonical-2222131100233100-2220002021300311-0313312030210120-3100202301102120-3010031021220201-0332113213232312-2022121122101212-1212001320211323"></a>

Type: `"object"`. single nested block, Optional.

This defines various configuration OPTIONS for Bot Defense policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001033102330210-3032320032202312-0101123110010100-3111122002032003-1111301100220132-1220220200333023-2320332223031021-1202001321022223"></a>

### Direct properties for `bot_defense.policy`

- [disable_js_insert](resources--http_loadbalancer--reference--group-011.md#canonical-1023120000231211-1120322321201221-1310010000102023-0310221102300332-0010113223211330-1021003232220330-2222333021111203-1003303112012220): complete subsection reference.

- [disable_mobile_sdk](resources--http_loadbalancer--reference--group-011.md#canonical-2230330011120300-3101212101220222-1220232201312210-1112101223121321-0200330122010231-1323221201321200-0331012223032322-3200102123011303): complete subsection reference.

<a id="canonical-1010212102203023-0013311013313220-0233300311131010-1133132301300230-1102121221202231-0111110133331133-0333110320303130-2311021202311012"></a>

<a id="canonical-1220013011230330-2023100102302120-1010120332302122-0213100001221001-2023001301001112-0112130133021320-0213001113113112-1011202213102101"></a>

#### `bot_defense.policy.javascript_mode` property

Type: `"string"`. Optional.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Additional upstream details:

Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is cacheable.

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1122030331022310-0210132211023230-2030303021110302-2111103131213323-2033012333021012-2311123000001212-0133011133011312-3132031230021201"></a>

<a id="canonical-2201322302000112-0000333310303100-0200022312020233-1112013033202210-1200233102220030-1233203012112230-0210212213121211-3110330231002302"></a>

#### `bot_defense.policy.js_download_path` property

Type: `"string"`. Optional.

Customize Bot Defense Client JavaScript path. If not specified, default \`/common.js\`

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [js_insert_all_pages](resources--http_loadbalancer--reference--group-011.md#canonical-3012330211022301-3011201012333021-2322122100300313-0123223010221231-1200011130111223-2330022223013320-3221323030122131-3213132322333022): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--reference--group-012.md#canonical-3010233111031320-1000010120220212-2010131002212030-3010123010010223-0130012300221033-2121300300312120-3332330320111012-2300030113131323): complete subsection reference.

- [mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-2123101112110001-2230223232113311-3200210313122201-3023210012031221-2033032323201332-0233120301211233-3302003023321332-3031131301202030): complete subsection reference.

- [protected_app_endpoints](resources--http_loadbalancer--reference--group-012.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021): complete subsection reference.

<a id="canonical-1023120000231211-1120322321201221-1310010000102023-0310221102300332-0010113223211330-1021003232220330-2222333021111203-1003303112012220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.disable_js_insert` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.disable_js_insert

<a id="canonical-2201222023013111-2022320122113322-0103123122033301-2302031210111222-2310331330201102-1321233301201213-0332021112320223-1320110103000020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230330011120300-3101212101220222-1220232201312210-1112101223121321-0200330122010231-1323221201321200-0331012223032322-3200102123011303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.disable_mobile_sdk` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.disable_mobile_sdk

<a id="canonical-0020211222310220-2300231103123101-3312301120101133-3330222101330110-1020332021101123-2310002201201221-0323021013332312-3112120222013030"></a>

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
disable_mobile_sdk = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012330211022301-3011201012333021-2322122100300313-0123223010221231-1200011130111223-2330022223013320-3221323030122131-3213132322333022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.js_insert_all_pages

<a id="canonical-3010223110121330-2211012121113201-1232233233103003-1021310112023321-2300120321123023-2133303300101013-2112322030010012-3221213020331111"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages.

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
js_insert_all_pages {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102131011022123-0001331101030211-2102133212033232-2232211121020033-3111013010221101-2121013300003200-1002101122211310-2312103202222031"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages`

<a id="canonical-0301223223131311-0210322021030100-3022133133101103-2321222132232133-0023132221001210-0003132011103011-3331122022310130-0223120102223100"></a>

#### `bot_defense.policy.js_insert_all_pages.javascript_location` property

Type: `"string"`. Optional.

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

<a id="canonical-0003112301221022-3203021200330313-3102213123032103-2122020001032201-2313120131131020-1031130310021132-3233031303221021-1201302320113211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- bot_defense.policy.js_insert_all_pages_except

<a id="canonical-0130201302011113-3300313233220221-0312202130112001-0033302221211032-1110113211123303-0211323010120203-0001012000120310-0200203031330112"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages with the exceptions.

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
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202313123113310-1033203112020131-2210100322312020-2213002320203330-1123200101310202-1232210101012212-1301320320223301-0120131323202123"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except`

- [exclude_list](resources--http_loadbalancer--reference--group-012.md#canonical-0220210010200300-0333332012212012-3030202323033020-0011020010332333-3022031112121301-0311322221231113-2132223111233132-1312332011100330): complete subsection reference.

<a id="canonical-3023001301210112-0102331321212123-1132010121330303-2132113010301220-3313231311112203-2200211021130130-1130301100013311-2323222111033332"></a>

<a id="canonical-0121331021020203-3112313000000203-2110110211222202-0223311231031100-2020002020020002-3230322211030020-1322320333010323-1021111033331212"></a>

#### `bot_defense.policy.js_insert_all_pages_except.javascript_location` property

Type: `"string"`. Optional.

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
