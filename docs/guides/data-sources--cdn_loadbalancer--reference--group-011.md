---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-1311023333203122-0300220001302131-2302022313302111-1211100313310110-2222132130133110-0012222132111103-0001111322332332-0223310211121122"></a>

#### `ddos_mitigation_rules.ip_prefix_list.ip_prefixes` property

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

<a id="canonical-2100220232201032-0220103110012110-0210030130100202-3031003111313030-0111111011112011-2100122010300210-2331230311310301-0313121133211311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- ddos_mitigation_rules.metadata

<a id="canonical-3133113000112203-3230210001020210-0011011132123320-0320032020121110-1333111223220120-3320012132111223-3231232103003232-1311331111033133"></a>

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

<a id="canonical-0332113331111303-1301010111001302-1312312322112001-0022111002231002-0012330230212213-0101121310021132-0311101132113120-0203203300201230"></a>

### Direct properties for `ddos_mitigation_rules.metadata`

<a id="canonical-1232021033022132-0203003001103303-1030120331000202-1100013220222233-0130330011010003-3302001323230111-0320102323310303-1200203132102112"></a>

#### `ddos_mitigation_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0302301302213332-0320232332133122-3211311132030012-2311322321310221-0313123122233103-1230010022211203-2113213230230010-2013112310301223"></a>

<a id="canonical-2032003131131111-2113013221212301-0230033100200312-0100122322102101-2301300313220231-1110202323232233-3133001030133101-2322300321331230"></a>

#### `ddos_mitigation_rules.metadata.name` property

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

<a id="canonical-1332133310300002-2012012123231201-2303220202332322-0013322233210021-2211322201103311-3322122203213210-1031203102223230-1301100022213013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_cache_action` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- default_cache_action

<a id="canonical-2030101302222322-1330132213002220-1131202312123021-2322002020000200-2211123110011331-1101102302311223-0012122223323231-3220130102213300"></a>

Type: `"single"`. Computed.

Default Cache Behaviour. This defines a Default Cache Action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

<a id="canonical-3012003101030210-2102101003312223-0022223303122303-0212213000332212-3131102322110133-3200100032320322-1200023002000111-2022223313320310"></a>

### Direct properties for `default_cache_action`

- [cache_disabled](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0121101321011330-1202101313121010-2321230230231101-1121120002103123-3001123113222001-1021202122301131-1312203331013203-3310023101032132): complete subsection reference.

<a id="canonical-2223201012222212-2000321033320123-3321011201203220-2031330220212013-2211111332331133-3320221031031203-3010020320131030-0032310123321101"></a>

<a id="canonical-0200111133333022-3003103300000101-1003210303122022-3110133000133132-1320200010211323-3203221321200120-3112111320302303-1210211330212002"></a>

#### `default_cache_action.cache_ttl_default` property

Type: `"string"`. Computed.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-2201230111011121-0300023133320132-2223213110320203-2233133202332302-2200000130112221-1032111321111220-0312300313001223-0331301133033212"></a>

<a id="canonical-1321300221001212-3312230320131032-2101120332231333-0110322203121120-1323131221122210-1310113120202223-2023301330102302-0211332203333232"></a>

#### `default_cache_action.cache_ttl_override` property

Type: `"string"`. Computed.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-0121101321011330-1202101313121010-2321230230231101-1121120002103123-3001123113222001-1021202122301131-1312203331013203-3310023101032132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_cache_action.cache_disabled` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [default_cache_action](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1332133310300002-2012012123231201-2303220202332322-0013322233210021-2211322201103311-3322122203213210-1031203102223230-1301100022213013)
- default_cache_action.cache_disabled

<a id="canonical-2131012120330121-3102201321320113-1113223212211030-0233223001030310-3113011120312213-3223302201300010-0323032311301001-1111212130310130"></a>

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

<a id="canonical-1010012002311312-2002113220132033-1221331121112232-1021133132132021-2130332101202113-0221322231220102-3102103213111300-3013223032223130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_sensitive_data_policy` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- default_sensitive_data_policy

<a id="canonical-2330133001212121-0330233001101130-0323020232230330-0330320330111311-0020031322121132-1031332002203112-2231100232133120-1121313001010103"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_sensitive\_data\_policy, sensitive\_data\_policy; Default:
default\_sensitive\_data\_policy\] Policy configuration for this feature.

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

- [default_sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2330133001212121-0330233001101130-0323020232230330-0330320330111311-0020031322121132-1031332002203112-2231100232133120-1121313001010103)
- [sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-2122233223233301-0013212123322233-3031133033203232-2102201333013333-1030211211110231-1012300301301312-3233203101222202-1212331200200122)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133121230022100-2022230330003230-3321032220303102-0222331220132321-1003003002112321-2113013000232022-2132122322130220-2202000001220303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_definition` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- disable_api_definition

<a id="canonical-0123123210010032-2000110022103021-1311000120001310-2022020021332330-3303230013303001-1302033001330103-1133323112122130-2022201130112011"></a>

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

<a id="canonical-2210310302113200-0331303012200133-2220232021130020-3021233123102120-2002110232001123-2011310312211231-3221320011121131-2310031201232130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_discovery` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- disable_api_discovery

<a id="canonical-2201000223210120-3010212111303012-0031202020003000-0301323201023023-2112212131222120-3122121213100233-1033122112001322-1003312022123101"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option

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

- [disable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2201000223210120-3010212111303012-0031202020003000-0301323201023023-2112212131222120-3122121213100233-1033122112001322-1003312022123101)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3000030200122030-0023023233032300-2302130231201112-3102133122112103-2101132123001120-3313103121133032-1220220123130310-1131311311231132)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310130022320333-2023113130232022-3001123220220101-1211301202220220-3211221311101222-0131032021311110-1330300000010302-1132112202231013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_client_side_defense` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- disable_client_side_defense

<a id="canonical-2220201123022201-0023320203013210-0103023102310013-3021121032321033-0322032301001210-1303311000012110-3313332333310202-3323311313300010"></a>

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

<a id="canonical-1230330223220032-2320012133303033-2132103221322133-1121210333222032-0231301133223331-2312100030231031-3133012220032010-0022222303103100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_ip_reputation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- disable_ip_reputation

<a id="canonical-0123331323112220-0331103211133121-3221123220010022-2230303012313022-3332211133112130-1120302111122103-0330013112002331-3000320211023010"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option

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

- [disable_ip_reputation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0123331323112220-0331103211133121-3221123220010022-2230303012313022-3332211133112130-1120302111122103-0330013112002331-3000320211023010)
- [enable_ip_reputation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0012233011212123-2221220022020201-0002231023133103-2113023310313233-3112330123000330-0213030121033133-0010020323331020-1120012311122200)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200122332202011-0321013102011031-2230312120233300-1320313121323121-0112202233132321-3110032333203110-1102302030023231-1110320100000231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- disable_malicious_user_detection

<a id="canonical-3200223301011022-2132031013210320-2122301303010130-2021133323320032-0311023033010332-2333113230131233-1121022000121320-3303113221032010"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.

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

- [disable_malicious_user_detection](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3200223301011022-2132031013210320-2122301303010130-2021133323320032-0311023033010332-2333113230131233-1121022000121320-3303113221032010)
- [enable_malicious_user_detection](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0302020030232132-1010023313332213-0103012013323223-2320302031211003-3201220220222322-1331111011020131-1321302331133010-1322032313131113)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103233312101321-1012103031202211-1313020110322221-3331212202131021-3223002230233033-1202020313103331-3110132313033031-0121221031132120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_rate_limit` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- disable_rate_limit

<a id="canonical-0230200132300230-0033310301032010-3311301312300002-1010120312120321-2303002232103233-3230310021030203-1302300133013231-0102323300011210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable rate limit.

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

<a id="canonical-2330031120223001-2010001102032332-1020100131030200-2020123121201321-3030200301231003-3000132310302113-1001322311010112-0120000032212121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_threat_mesh` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- disable_threat_mesh

<a id="canonical-0232231321312110-3012302202330333-3003333233111230-3103133220103212-2201032133332311-1112033320333111-3332000110102013-0123021232232313"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_threat\_mesh, enable\_threat\_mesh; Default: disable\_threat\_mesh\] Enable this
option

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

- [disable_threat_mesh](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0232231321312110-3012302202330333-3003333233111230-3103133220103212-2201032133332311-1112033320333111-3332000110102013-0123021232232313)
- [enable_threat_mesh](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1001133203101133-0332231303230311-0232110203300322-2201200302202303-3000232312132010-2020102113023300-1123023322022021-0321303022013012)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011233221222131-0010110333302333-0131112313222132-1210002303313123-1230101332030311-1003310103111003-1103032133222033-1011313002220211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_waf` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- disable_waf

<a id="canonical-2122210121011302-0110121231301020-0200000013333201-2333003210032130-3133020131132302-0002220100021211-1112113032300322-0202210321020123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable waf.

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

<a id="canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- enable_api_discovery

<a id="canonical-3000030200122030-0023023233032300-2302130231201112-3102133122112103-2101132123001120-3313103121133032-1220220123130310-1131311311231132"></a>

Type: `"single"`. Computed.

Specifies the settings used for API discovery.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

<a id="canonical-1122012121210121-2122102323312213-2122132033230102-1022233313130303-0223022023001322-2320120100033333-2213210301002331-3131232201200132"></a>

### Direct properties for `enable_api_discovery`

- [api_crawler](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0300012021302023-0230322002110233-1133222132031200-1322321331212113-3322121333331211-0122133331221303-1311110012222111-3322210220230132): complete subsection reference.

- [api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3130012221032110-1201130323101212-3030020021200222-3011210203313100-0313030313110313-2112110131233003-1102003221021013-2113130320202200): complete subsection reference.

- [custom_api_auth_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3120303222023021-2012311311330033-1012013330102033-1102212133223323-3013330220323313-2130103020333101-0223132211201203-3031221211130312): complete subsection reference.

- [default_api_auth_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102101211112103-1122022113010323-1301312112200322-3122112110211013-2131302132202221-1002230202200002-1013303331331030-0332112000232003): complete subsection reference.

- [disable_learn_from_redirect_traffic](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3100323111023201-2220120202012320-0013102130131001-1301013113003123-3223312300002222-0111233030223210-1201102331001032-1031330011132211): complete subsection reference.

- [discovered_api_settings](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2302020001133310-2202333002311232-1203010113012132-3110210300200013-1331110013310313-3330102023110330-3102023220123311-3002102322021321): complete subsection reference.

- [enable_learn_from_redirect_traffic](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1330102030002321-3330010110023132-1133221033012220-0211110123211101-2101021101321330-0300112330312021-1100323320102300-3110130021131300): complete subsection reference.

<a id="canonical-0300012021302023-0230322002110233-1133222132031200-1322321331212113-3322121333331211-0122133331221303-1311110012222111-3322210220230132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- enable_api_discovery.api_crawler

<a id="canonical-1032330130020110-2232012302212122-0230123323122112-0121031031201330-2123132112012122-0321311211310111-3310012012131333-3131220203231332"></a>

Type: `"single"`. Computed.

API Crawling. API Crawler message.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

<a id="canonical-2121331023023122-3033332330020130-1013212030121001-0330001311302120-2100112232202013-0322020302103310-3321301312213101-3233111212233120"></a>

### Direct properties for `enable_api_discovery.api_crawler`

- [api_crawler_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0232023031220030-2232221323300123-2203033010003000-3323123112201021-3133010230333212-0220123000312022-0302201202122302-2321030230333232): complete subsection reference.

- [disable_api_crawler](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0130301013113002-1230210303221010-2023131302013130-0331312032223333-3111203201110131-1100121312322211-1112023120233130-3202220113333213): complete subsection reference.

<a id="canonical-0232023031220030-2232221323300123-2203033010003000-3323123112201021-3133010230333212-0220123000312022-0302201202122302-2321030230333232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0300012021302023-0230322002110233-1133222132031200-1322321331212113-3322121333331211-0122133331221303-1311110012222111-3322210220230132)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="canonical-1121000130030303-2110233221122331-2031130230212012-0232003222032221-3032133310211313-1302112301213120-1112002200130103-1022330101033320"></a>

Type: `"single"`. Computed.

Crawler Configure.

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

<a id="canonical-1312231230322213-0210131001200200-3013011201001110-1022111101312331-0121030121221212-1222202332023110-1312102331023122-0110320230233332"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config`

- [domains](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2202201323202302-1032321320213201-0131003232123323-3012021221013122-1010300313010313-2033001032233210-0311030330031123-1310310233021021): complete subsection reference.

<a id="canonical-2202201323202302-1032321320213201-0131003232123323-3012021221013122-1010300313010313-2033001032233210-0311030330031123-1310310233021021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0300012021302023-0230322002110233-1133222132031200-1322321331212113-3322121333331211-0122133331221303-1311110012222111-3322210220230132)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0232023031220030-2232221323300123-2203033010003000-3323123112201021-3133010230333212-0220123000312022-0302201202122302-2321030230333232)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-0001022032311333-3021223121313101-3223120101100122-1102010330213020-0023012222102030-3011311003221113-0313311300123233-3011333312110310"></a>

Type: `"list"`. Computed.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-2201030303020021-1112312133211123-3011010020221201-2122101320000113-1301221310220032-0210013313020321-1033012231120232-1311323120012313"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains`

<a id="canonical-0233201211230100-1320213312122113-1013210203211113-2231123112323212-3132132130300322-2133223322000223-0121220200010222-1230132332021203"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.domain` property

Type: `"string"`. Computed.

Select the domain to execute API Crawling with given credentials.

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

- [simple_login](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1100331302331013-1303333312221230-3201221100230223-1133113103013222-1032303012011110-1331210020320201-1302222122102332-1312321323212231): complete subsection reference.

<a id="canonical-1100331302331013-1303333312221230-3201221100230223-1133113103013222-1032303012011110-1331210020320201-1302222122102332-1312321323212231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0300012021302023-0230322002110233-1133222132031200-1322321331212113-3322121333331211-0122133331221303-1311110012222111-3322210220230132)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0232023031220030-2232221323300123-2203033010003000-3323123112201021-3133010230333212-0220123000312022-0302201202122302-2321030230333232)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2202201323202302-1032321320213201-0131003232123323-3012021221013122-1010300313010313-2033001032233210-0311030330031123-1310310233021021)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-3110320122332003-3000000032013110-1301231131230120-0010221002220323-1120201232131000-2212022301031012-2011012223020031-3102113201200300"></a>

Type: `"single"`. Computed.

Configuration parameter for simple login.

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

<a id="canonical-3022210121332210-2001222233100201-3011202312113222-1300303211323311-2032103202130130-1003103220010012-3303212033023312-0302101010303100"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login`

- [password](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1002102223322121-3223101312320130-3012331213003003-1130201022323013-1201302322000312-3110002023113321-1123030101100010-0220020100120032): complete subsection reference.

<a id="canonical-2023122030030332-1313110033203020-3112100003132010-0333212200001031-2200213300211223-2330011033232012-1331020000232002-0130002232131310"></a>

<a id="canonical-1110120111023102-3132132031330203-1032310003132003-0300010113122323-1103320301223100-2203021322132012-0230330301231023-0321210031000022"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.user` property

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1002102223322121-3223101312320130-3012331213003003-1130201022323013-1201302322000312-3110002023113321-1123030101100010-0220020100120032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0300012021302023-0230322002110233-1133222132031200-1322321331212113-3322121333331211-0122133331221303-1311110012222111-3322210220230132)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0232023031220030-2232221323300123-2203033010003000-3323123112201021-3133010230333212-0220123000312022-0302201202122302-2321030230333232)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2202201323202302-1032321320213201-0131003232123323-3012021221013122-1010300313010313-2033001032233210-0311030330031123-1310310233021021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1100331302331013-1303333312221230-3201221100230223-1133113103013222-1032303012011110-1331210020320201-1302222122102332-1312321323212231)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-1323110012023321-0122111321011020-2202102222232003-1203301322101211-0303013321001022-1213200313021201-1300110131211113-0002010112320213"></a>

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

<a id="canonical-2012322133302212-2322110300221213-0000120313302021-2312110233132023-3202222012302121-0332202021111213-1113233013113202-1020111212333213"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password`

- [blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3201033032223011-0301123013220022-0231303010131201-0021323013310130-1232112311012202-1121132222220232-1311232033010313-0223331000103021): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0311311111001022-0303330201300322-1021222301230030-0231211102020211-3132331302220011-0013322333021022-1011122113200222-1132232202220023): complete subsection reference.

<a id="canonical-3201033032223011-0301123013220022-0231303010131201-0021323013310130-1232112311012202-1121132222220232-1311232033010313-0223331000103021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0300012021302023-0230322002110233-1133222132031200-1322321331212113-3322121333331211-0122133331221303-1311110012222111-3322210220230132)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0232023031220030-2232221323300123-2203033010003000-3323123112201021-3133010230333212-0220123000312022-0302201202122302-2321030230333232)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2202201323202302-1032321320213201-0131003232123323-3012021221013122-1010300313010313-2033001032233210-0311030330031123-1310310233021021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1100331302331013-1303333312221230-3201221100230223-1133113103013222-1032303012011110-1331210020320201-1302222122102332-1312321323212231)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1002102223322121-3223101312320130-3012331213003003-1130201022323013-1201302322000312-3110002023113321-1123030101100010-0220020100120032)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-2212221112111100-3311020100102113-1103233223212220-3220002020133000-1033113331323301-2330211313012330-1223002113313122-2222323000302133"></a>

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

<a id="canonical-0333320322330312-2010221303232022-2213122200100200-0313312203023313-3113032333131021-3113120122020113-2213011030110122-1033331313030231"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info`

<a id="canonical-2133313232330313-0303101013212120-2331333101132132-0001113300022221-2203311231102012-1202102330303002-3120302101210131-0332323122212303"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3020200103101110-2311311233101130-3322213001133100-1231201332133210-0112221311231010-0102013231013113-1220111023030221-3112322223013003"></a>

<a id="canonical-0333223101010203-0011220001233002-1032313130312112-1021123010033202-3000132001022220-1220211333331202-3320210030201222-1030202300133200"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.location` property

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

<a id="canonical-3132102320131111-1330303330012010-0120302113013211-0102322301033321-3002011223332333-1223213103033013-3021002303132232-3003122112313310"></a>

<a id="canonical-3010231213123121-0312232101331201-2211221222301100-0010332002311312-1120211210221231-2003002133313320-2203112010023232-3011332100110001"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.store_provider` property

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

<a id="canonical-0311311111001022-0303330201300322-1021222301230030-0231211102020211-3132331302220011-0013322333021022-1011122113200222-1132232202220023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0300012021302023-0230322002110233-1133222132031200-1322321331212113-3322121333331211-0122133331221303-1311110012222111-3322210220230132)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0232023031220030-2232221323300123-2203033010003000-3323123112201021-3133010230333212-0220123000312022-0302201202122302-2321030230333232)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2202201323202302-1032321320213201-0131003232123323-3012021221013122-1010300313010313-2033001032233210-0311030330031123-1310310233021021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1100331302331013-1303333312221230-3201221100230223-1133113103013222-1032303012011110-1331210020320201-1302222122102332-1312321323212231)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1002102223322121-3223101312320130-3012331213003003-1130201022323013-1201302322000312-3110002023113321-1123030101100010-0220020100120032)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-1112302101331030-2113133132302133-0030110312101211-2103230232321231-1010113111303011-3102220102110103-3032232303101222-3320113232130130"></a>

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

<a id="canonical-3020311201002332-1233211012132133-0312033321023222-3100010333331131-1001302101300102-2133021020322310-2320333230330123-0211133331132201"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info`

<a id="canonical-3332121203322012-3330221000211201-2303221332233103-3223122201323003-0023311212001010-2321323310112010-3233133313130033-1021212223213100"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1100021100232020-3233300230230020-0123223101322300-2230122312321113-0221310230232301-2333112100130010-3323101032233311-3100301132212311"></a>

<a id="canonical-1312033111110012-3330113103021310-1331323011201213-2032110200300013-3210133322222100-3100131013030103-3032122013001313-2011323021122132"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.url` property

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

<a id="canonical-0130301013113002-1230210303221010-2023131302013130-0331312032223333-3111203201110131-1100121312322211-1112023120233130-3202220113333213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.disable_api_crawler` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0300012021302023-0230322002110233-1133222132031200-1322321331212113-3322121333331211-0122133331221303-1311110012222111-3322210220230132)
- enable_api_discovery.api_crawler.disable_api_crawler

<a id="canonical-1120010133201031-1323231100012020-0011301000212212-3302110033032003-3012021121230230-3210322222033302-0230112110010132-1330303002211233"></a>

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

<a id="canonical-3130012221032110-1201130323101212-3030020021200222-3011210203313100-0313030313110313-2112110131233003-1102003221021013-2113130320202200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- enable_api_discovery.api_discovery_from_code_scan

<a id="canonical-2112113020012322-2130010322230131-1330210111333221-1231221232213102-3131231210131320-0313132112231131-0223201110312002-1333023032131212"></a>

Type: `"single"`. Computed.

Select codebase and Repositories.

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

<a id="canonical-3011131112011020-3313330101003300-2131210010132322-1211210203232023-2303330110021021-2013222223213201-2000300023010331-1331000012013020"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan`

- [code_base_integrations](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1011011310112303-0301002103310223-1232133010321031-2300111100333210-1102121133033230-0123101123001131-2123021320232222-0121021023212102): complete subsection reference.

<a id="canonical-1011011310112303-0301002103310223-1232133010321031-2300111100333210-1102121133033230-0123101123001131-2123021320232222-0121021023212102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3130012221032110-1201130323101212-3030020021200222-3011210203313100-0313030313110313-2112110131233003-1102003221021013-2113130320202200)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-3322222021332310-3330323101113130-0220212203111123-0032210003003221-2023201210011202-0010123310130223-3112100130323013-2231130130130210"></a>

Type: `"list"`. Computed.

Configuration parameter for codebase integrations.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0011122111333202-0021321022303310-2320122301131231-0111203131111212-0110331332102310-0222130133013100-1300201010001213-1000022213032222"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations`

- [all_repos](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1000201310031303-1200201310322323-2333000102112211-3022212222203022-1233221121233222-0110223132131310-3200332013313100-1030333131300130): complete subsection reference.

- [code_base_integration](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1330331101330300-0212110132232233-2330030212320302-3033131002001112-0131313201323120-1333200211132022-1103030033313113-0003232002233111): complete subsection reference.

- [selected_repos](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0300021133112232-2100131222231313-1000302303123022-0210033110033112-2011100310303033-3320333313320100-1031200203201303-1003010113310101): complete subsection reference.

<a id="canonical-1000201310031303-1200201310322323-2333000102112211-3022212222203022-1233221121233222-0110223132131310-3200332013313100-1030333131300130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3130012221032110-1201130323101212-3030020021200222-3011210203313100-0313030313110313-2112110131233003-1102003221021013-2113130320202200)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1011011310112303-0301002103310223-1232133010321031-2300111100333210-1102121133033230-0123101123001131-2123021320232222-0121021023212102)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-3333203331320013-2211113000320021-3113121201320300-2021020302130321-1112332231223020-1212313022012330-3123333020212001-0203133322221100"></a>

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

<a id="canonical-1330331101330300-0212110132232233-2330030212320302-3033131002001112-0131313201323120-1333200211132022-1103030033313113-0003232002233111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3130012221032110-1201130323101212-3030020021200222-3011210203313100-0313030313110313-2112110131233003-1102003221021013-2113130320202200)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1011011310112303-0301002103310223-1232133010321031-2300111100333210-1102121133033230-0123101123001131-2123021320232222-0121021023212102)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-2200303132300122-0013101310133322-0200113311210232-2323333011103210-2231032011220202-2311323302120012-3303312110000301-3200033333200202"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2302303013002101-2321010232010323-0322012330232233-1032203131320333-2330210022122331-3203332132101120-0233120003003210-2311132122210131"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration`

<a id="canonical-3131122202003302-0232230221012030-3330110011213131-0103220312033203-2112323200001300-2001210200032222-3003222213331202-0320021132313203"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.name` property

Type: `"string"`. Computed.

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

<a id="canonical-3203300110101130-3122210021231302-0310300020003120-2230301120012213-1311121301121012-0333321131023311-2020132201003331-0001201100322101"></a>

<a id="canonical-2031003013222000-0010100103000121-0101213030122031-1121230312322002-1030203021132130-3323013213210101-3132010203201100-2030000131321123"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-0223310012312030-0131031222113001-1122300213010302-0311313022231321-2232312332020212-0211130331223132-0130030200221122-3223120113122132"></a>

<a id="canonical-0223231122103030-0111230112122111-2230300033121310-2022133122031111-3312210032010231-0013131303102113-3330321230311312-1233010022010212"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.tenant` property

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

<a id="canonical-0300021133112232-2100131222231313-1000302303123022-0210033110033112-2011100310303033-3320333313320100-1031200203201303-1003010113310101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3130012221032110-1201130323101212-3030020021200222-3011210203313100-0313030313110313-2112110131233003-1102003221021013-2113130320202200)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1011011310112303-0301002103310223-1232133010321031-2300111100333210-1102121133033230-0123101123001131-2123021320232222-0121021023212102)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-0010023202022022-0013312322202230-2020010212332200-3130021223331010-3230101332132313-1022300120323212-2313022323202332-1022312022031312"></a>

Type: `"single"`. Computed.

Select which API repositories represent the LB applications.

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

<a id="canonical-0112330102322120-1233303101100023-1200102313322212-0120322231013312-3233332021003002-2131030323130222-0210020323301323-2313112311332303"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos`

<a id="canonical-1003313023302100-3222100231122001-2312101031102331-0220332013301311-3000122221233130-2023111311301323-3022210130023130-1103311201122033"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos.api_code_repo` property

Type: `["list", "string"]`. Computed.

Code repository which contain API endpoints.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3120303222023021-2012311311330033-1012013330102033-1102212133223323-3013330220323313-2130103020333101-0223132211201203-3031221211130312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.custom_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-1000331012303113-2332030233121110-3313310221120022-1333103202202300-0211221113310331-2330221131021012-2003013022002313-0120022131233233"></a>

Type: `"single"`. Computed.

API Discovery Advanced Settings. API Discovery Advanced settings.

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

<a id="canonical-1300212231310303-1220221012111301-1002002023232311-0101333110002132-3302302330331222-0211121020003120-2333232320021221-3132022302332113"></a>

### Direct properties for `enable_api_discovery.custom_api_auth_discovery`

- [api_discovery_ref](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3122232112023032-0030111201102002-0103023002210100-3031233233032131-3213210231332202-3300213203120322-1113311023202311-1223013320203222): complete subsection reference.

<a id="canonical-3122232112023032-0030111201102002-0103023002210100-3031233233032131-3213210231332202-3300213203120322-1113311023202311-1223013320203222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- [enable_api_discovery.custom_api_auth_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3120303222023021-2012311311330033-1012013330102033-1102212133223323-3013330220323313-2130103020333101-0223132211201203-3031221211130312)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-1130213322223311-2223331013121131-0323111102222021-1212013313213123-0221231231003013-3033011122230012-0001122102212222-2233210101122221"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2300001331201310-2232002332020012-2001133123303030-1021300033231123-3012010221023332-1330323033223213-0030232032011030-3211223103211312"></a>

### Direct properties for `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref`

<a id="canonical-1120331013323203-2323110001310130-0022113113121203-1011123013331220-0133212203002012-1013210110332121-3000311213121323-1101203220130320"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.name` property

Type: `"string"`. Computed.

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

<a id="canonical-2002210201313001-0223220130013120-2113321112222330-3110320322122303-1331030022130312-3013201013303212-1310032231100330-0030103101033102"></a>

<a id="canonical-2312130101331301-0013322230110013-0020033033203010-0232200023123310-3013231312010310-0223311131102031-2331320221300122-2220211100023200"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-2113310331121311-2210000333233303-1112311003101201-2133002102003001-0333100300333133-3000201000022310-1030131312311201-2020101322023223"></a>

<a id="canonical-2233001230123210-3100221311310030-1333120322133001-2100130002013310-1230101022303221-1003121102133212-1210321231203013-0313201130001313"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.tenant` property

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

<a id="canonical-2102101211112103-1122022113010323-1301312112200322-3122112110211013-2131302132202221-1002230202200002-1013303331331030-0332112000232003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.default_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-2113031213032001-0131233211322112-3110011100332021-2201320231110021-0121023200223033-3002000203331001-1322310333312133-0003113003102021"></a>

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

<a id="canonical-3100323111023201-2220120202012320-0013102130131001-1301013113003123-3223312300002222-0111233030223210-1201102331001032-1031330011132211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.disable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-1113120003332030-1000113201222220-1333011133312033-3013312210232002-2100112030101022-3223011302113312-0130201232001020-2020033200333322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable learn from redirect traffic.

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

<a id="canonical-2302020001133310-2202333002311232-1203010113012132-3110210300200013-1331110013310313-3330102023110330-3102023220123311-3002102322021321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.discovered_api_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- enable_api_discovery.discovered_api_settings

<a id="canonical-1333030300321333-0310223332030212-3313100100103221-0200120020330310-3022113212030021-3320200000112213-0222113201232122-1022330011300331"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

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

<a id="canonical-0022301132030220-0122021312202230-0232302033020312-0201000132133020-3112301333331333-3020332131033033-0110112221003133-3020031113102301"></a>

### Direct properties for `enable_api_discovery.discovered_api_settings`

<a id="canonical-2231310310221310-3221230013030130-0111302311223131-1230203321032202-0013110231213332-3232230030020312-3332122312023033-1231322213323000"></a>

#### `enable_api_discovery.discovered_api_settings.purge_duration_for_inactive_discovered_apis` property

Type: `"number"`. Computed.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-1330102030002321-3330010110023132-1133221033012220-0211110123211101-2101021101321330-0300112330312021-1100323320102300-3110130021131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.enable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-0312023031311332-1000222221323103-0110231311313321-1211220212311201-0130201010032100-0011331203100002-1210221212230302-3222112032102211"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable learn from redirect traffic.

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

<a id="canonical-3100203021021010-2201131230132213-0121112120112123-1013013330103303-0200011323230211-3000331323222323-3133302303313021-3321010120100122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- enable_challenge

<a id="canonical-0212033300102210-2322130010322222-2030202300113231-1111303300313333-1032231121102300-0101020111100333-3330020033210300-1322203202210323"></a>

Type: `"single"`. Computed.

Configure auto mitigation i.e risk based challenges for malicious users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

<a id="canonical-3111101001002110-0013322221320133-1200131303303333-3002223302220113-1130230032212330-0122023110203011-3231101331102212-0202022101233022"></a>

### Direct properties for `enable_challenge`

- [captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1130331020232012-3312313121000101-1202212031323000-2212103221121211-0333223002223021-1103221211311230-2133131313112320-2223210103202100): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0212203122000231-1013120203032331-3202101120121031-3121112010103031-2323233212211303-3230130232322222-3032102121011333-1310202113331232): complete subsection reference.

- [default_js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1011203213133022-3210001013232312-0300302030221033-1002322322013313-1011102211021132-2012323000223223-1011122322111011-2120320100331021): complete subsection reference.

- [default_mitigation_settings](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0332303130122331-2320113230333122-3120102330113012-3302300120223003-1231312201110323-3121113031300103-0222010323113013-3323211111330322): complete subsection reference.

- [js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1303323212011011-3123132002222113-3133220323100322-2131030133102020-1212323233023211-3001301001002202-1132330001220110-0330321230031300): complete subsection reference.

- [malicious_user_mitigation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1310230021220030-2031231231012310-1003231330020201-1232121330300010-2222121001321201-3332330221203323-0101301001322333-3332001332032331): complete subsection reference.

<a id="canonical-1130331020232012-3312313121000101-1202212031323000-2212103221121211-0333223002223021-1103221211311230-2133131313112320-2223210103202100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3100203021021010-2201131230132213-0121112120112123-1013013330103303-0200011323230211-3000331323222323-3133302303313021-3321010120100122)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-2023010300033323-1330102101202233-2232331003132032-2113322003022232-1201212132030331-2331132130230333-1112301200102011-3020023102230103"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

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

<a id="canonical-3330020332223023-1012010313220331-1010333212303310-1232232132020103-0131333130310320-1312213201132120-0130302011122303-2201132000101331"></a>

### Direct properties for `enable_challenge.captcha_challenge_parameters`

<a id="canonical-1021110212123011-3123031323213332-1002213200002103-1222122123031233-2202032301330220-2001210111313120-0210000230331021-3211223020033200"></a>

#### `enable_challenge.captcha_challenge_parameters.cookie_expiry` property

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0112031203010020-2113110301002133-3212233201122331-2221202200322023-3032022330110230-3130332030202330-0331030101301112-2222213100311121"></a>

<a id="canonical-0133110100321333-1213013113130321-3300213121212122-3132222302310323-0211001223212121-2330023302200133-1013111313323110-2303031133203330"></a>

#### `enable_challenge.captcha_challenge_parameters.custom_page` property

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0212203122000231-1013120203032331-3202101120121031-3121112010103031-2323233212211303-3230130232322222-3032102121011333-1310202113331232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3100203021021010-2201131230132213-0121112120112123-1013013330103303-0200011323230211-3000331323222323-3133302303313021-3321010120100122)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-1223300012112003-3130033131201223-3203231123211303-3023313023310311-3231220233022302-3120303331223111-0001212112121223-0331030320002332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default captcha challenge parameters.

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

<a id="canonical-1011203213133022-3210001013232312-0300302030221033-1002322322013313-1011102211021132-2012323000223223-1011122322111011-2120320100331021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3100203021021010-2201131230132213-0121112120112123-1013013330103303-0200011323230211-3000331323222323-3133302303313021-3321010120100122)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-3320031021002020-2301110103311331-3210220013130103-0102020303301232-0301131202300313-3012123000312122-2022131122133002-0310311011223011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default js challenge parameters.

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

<a id="canonical-0332303130122331-2320113230333122-3120102330113012-3302300120223003-1231312201110323-3121113031300103-0222010323113013-3323211111330322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_mitigation_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3100203021021010-2201131230132213-0121112120112123-1013013330103303-0200011323230211-3000331323222323-3133302303313021-3321010120100122)
- enable_challenge.default_mitigation_settings

<a id="canonical-1010022100132111-2231030001131100-2222001232222312-3231321003201113-2200130130210200-1321313303230223-0311022211310211-0221110320103023"></a>

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

<a id="canonical-1303323212011011-3123132002222113-3133220323100322-2131030133102020-1212323233023211-3001301001002202-1132330001220110-0330321230031300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3100203021021010-2201131230132213-0121112120112123-1013013330103303-0200011323230211-3000331323222323-3133302303313021-3321010120100122)
- enable_challenge.js_challenge_parameters

<a id="canonical-1001301002211312-2322300033020222-1231112123310322-2033112013221112-2322211313220031-3000122301132222-2001303312211133-3222103301020031"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

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

<a id="canonical-1233332333130113-3001302033123211-0212211321301112-1313002301123223-3211122302101020-3220103322212300-2000131212321220-2311330232130311"></a>

### Direct properties for `enable_challenge.js_challenge_parameters`

<a id="canonical-2001131102232111-3000033132021123-2033320003322231-1332110012230212-3303303313313100-2300330313331312-3010032200201202-0231001131321312"></a>

#### `enable_challenge.js_challenge_parameters.cookie_expiry` property

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0200300021332113-1100022001121001-1222200010113032-3131310212221213-3330330133310102-3001210101001120-1222310322023022-0010101232232002"></a>

<a id="canonical-0301331033110323-0220000202000300-2033132202322222-2023000122302133-1321323020132000-1203322130222211-3300132102201021-3010020010020120"></a>

#### `enable_challenge.js_challenge_parameters.custom_page` property

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1223202200012121-3202321203302302-0033003021331112-2130113120023023-3312130210321330-1130203302022200-2230023121213031-3131113331210132"></a>

<a id="canonical-0133302311222313-2120321200230232-1230233312232221-0312200200133011-1113131133330321-1222001121020030-2322002312001313-0003132102122300"></a>

#### `enable_challenge.js_challenge_parameters.js_script_delay` property

Type: `"number"`. Computed.

Delay introduced by JavaScript, in milliseconds.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-1310230021220030-2031231231012310-1003231330020201-1232121330300010-2222121001321201-3332330221203323-0101301001322333-3332001332032331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.malicious_user_mitigation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3100203021021010-2201131230132213-0121112120112123-1013013330103303-0200011323230211-3000331323222323-3133302303313021-3321010120100122)
- enable_challenge.malicious_user_mitigation

<a id="canonical-1230211211031103-3202021012003200-0233312202310000-3110220100030201-0230303213323323-0223202221202200-1303022022333321-1123013132102333"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1230221101220033-0222302133121333-0000013023101110-0102331322021331-3321011212131000-1102113303120113-2010121303311331-1331003102303230"></a>

### Direct properties for `enable_challenge.malicious_user_mitigation`

<a id="canonical-1003120222331222-1313311021001002-0232023000332133-0002131201000231-0322031020030200-3103122301322031-1020000201131112-2032010232320123"></a>

#### `enable_challenge.malicious_user_mitigation.name` property

Type: `"string"`. Computed.

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

<a id="canonical-1321322331231303-3233012310202332-2031222302323002-0231000311330123-1223132010232200-3030212003132002-1321122110012321-0232313211202312"></a>

<a id="canonical-3102200022323210-0330123112101320-3131210121213230-0320233003200123-0033122130332133-1111030222112003-3111220223232333-0223021312323003"></a>

#### `enable_challenge.malicious_user_mitigation.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-1030330300231100-2030332001111233-2212133203331012-0302221211300132-3101131233331223-0220012030121322-3123022112222222-1101223130033221"></a>

<a id="canonical-3001000310103003-1203111131312130-0003111322323323-1321013000130320-0320202200003111-0230212232310230-3100210302113011-3031000003212221"></a>

#### `enable_challenge.malicious_user_mitigation.tenant` property

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

<a id="canonical-3313210111131130-3013223323200003-3222320132312223-1223312001120130-3200303232330111-0021013210202313-3232133003332212-2303232210322300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ip_reputation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- enable_ip_reputation

<a id="canonical-0012233011212123-2221220022020201-0002231023133103-2113023310313233-3112330123000330-0213030121033133-0010020323331020-1120012311122200"></a>

Type: `"single"`. Computed.

IP Threat Category List. List of IP threat categories.

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

<a id="canonical-1303200213123200-1123221110123233-1000002222121313-2222203311203131-3211131220233121-3022132311300011-3300230011333333-2233232202133013"></a>

### Direct properties for `enable_ip_reputation`

<a id="canonical-3321300231001203-0303310220030322-2001330233331213-1202113012333033-0002012320030103-2332033012133001-0233100321121212-3311302213310200"></a>

#### `enable_ip_reputation.ip_threat_categories` property

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1203313010102222-3212102003230230-2302023211203303-0222300301333121-1211131120202323-0002103332313100-2010320313200122-1201232211313203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- enable_malicious_user_detection

<a id="canonical-0302020030232132-1010023313332213-0103012013323223-2320302031211003-3201220220222322-1331111011020131-1321302331133010-1322032313131113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable malicious user detection.

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

<a id="canonical-3300213101122201-3232231121230323-2211130201313010-2120031103123202-3222331012020323-0110112112321033-1223200100130021-3333301222320233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_threat_mesh` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- enable_threat_mesh

<a id="canonical-1001133203101133-0332231303230311-0232110203300322-2201200302202303-3000232312132010-2020102113023300-1123023322022021-0321303022013012"></a>

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

<a id="canonical-3301232011322201-3322122212320230-0212320203323203-1233020330031223-0013321230233213-1003032122130202-3223331312211021-0020303030213011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- graphql_rules

<a id="canonical-3002333113320112-0230020210102123-0012120223313320-2113232311113212-1232330322311020-3233311322333003-0231113200210202-0000221012101220"></a>

Type: `"list"`. Computed.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

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

<a id="canonical-2011002100120321-3230022031302103-1001302103133002-0033120331132320-1101321112033031-0232203123331123-1131133223233320-1012101210321000"></a>

### Direct properties for `graphql_rules`

- [any_domain](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3122123330200010-3221021031230202-2022223233111302-0331010223321233-3211131302111200-0203100020121200-3101013132120110-0312232012113221): complete subsection reference.

<a id="canonical-2033330122330102-3001333313001001-1101111323020101-2101203120223211-3313123132213001-0110233220133011-2121322211002200-3313031311222322"></a>

<a id="canonical-3302233203132323-2330230023310213-3000213022232030-3221333133013210-2200311031021122-2000203101212211-1300331313312122-3301221210230301"></a>

#### `graphql_rules.exact_path` property

Type: `"string"`. Computed.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Additional upstream details:

Default value is /GraphQL.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1033130132111130-1131102131210021-2221230321131021-3033010322103331-1301332211322313-3111232230311111-0023330112113211-3332023020102201"></a>

<a id="canonical-2321133111032100-3111221332113013-2300110112213033-2000212013023031-0011220122210001-2110331301322320-2120330002013331-0110103331012102"></a>

#### `graphql_rules.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [graphql_settings](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0133031100210221-1323313022121310-3033102310033213-0203322113001232-2310323131112130-0100031120121320-3200031210201120-0002020120000011): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3031332133020023-1012223131330320-3213321102112313-0023203102000220-3210011311100102-0030000023202210-2223303313202332-1220102233031031): complete subsection reference.

- [method_get](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3121322313302002-0000113322110302-1200010200311302-3202123222100100-3301213132221000-0320220203200331-0222020222321022-1221232232303020): complete subsection reference.

- [method_post](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2033130003220333-0221001233022330-3023102331213101-3332330012100221-0001122211302312-2000302210110133-2131003002321231-2101120330323213): complete subsection reference.

<a id="canonical-3320300012002010-3132020022220113-1013323123102122-1000301212021121-2111223231202333-1220020212322122-2112003203230123-0100212313331301"></a>

<a id="canonical-0122223323230323-3302110210122203-1211032310201221-0303021300220231-3323112132123232-3020132301120000-3012101120113121-0102003203031211"></a>

#### `graphql_rules.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3122123330200010-3221021031230202-2022223233111302-0331010223321233-3211131302111200-0203100020121200-3101013132120110-0312232012113221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3301232011322201-3322122212320230-0212320203323203-1233020330031223-0013321230233213-1003032122130202-3223331312211021-0020303030213011)
- graphql_rules.any_domain

<a id="canonical-1313203110112033-0220230313000231-0313230103113001-0302233230122131-2032321121101331-3232110000022130-1310220211211022-0003323121023313"></a>

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

<a id="canonical-0133031100210221-1323313022121310-3033102310033213-0203322113001232-2310323131112130-0100031120121320-3200031210201120-0002020120000011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3301232011322201-3322122212320230-0212320203323203-1233020330031223-0013321230233213-1003032122130202-3223331312211021-0020303030213011)
- graphql_rules.graphql_settings

<a id="canonical-1210123203203212-0130322210013233-2003001121321302-0133303200020012-0003300130321002-0312023320203310-3003112213110133-2220100212132220"></a>

Type: `"single"`. Computed.

Configuration parameter for GraphQL settings.

Additional upstream details:

GraphQL configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allow_introspection_queries_choice": "[\"disable_introspection\",\"enable_introspection\"]"
}
```

<a id="canonical-0231213221101310-1033223300330012-2120030113122111-0130233333031130-1013220121133221-2213120030033030-1102230223202011-1210201212012013"></a>

### Direct properties for `graphql_rules.graphql_settings`

- [disable_introspection](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1102101103100110-1033113013022203-0333322013112001-1331332320212301-2120211002300001-3321330212220023-2223212131210130-0302102322213003): complete subsection reference.

- [enable_introspection](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3213213230203211-2222231212030110-0133000120003020-2131103210121022-0020100211320131-2301222011003201-1233101333302102-0323100212330321): complete subsection reference.

<a id="canonical-0221003301212000-0103033302101323-1203003000333311-3201201021122113-0121000323310231-3223332113100223-3303221232303322-0211332222122221"></a>

<a id="canonical-2221323110033033-0010020001002230-1010330010031021-0201013302222231-0200232312213213-3033320100303013-2033201002301213-1001200203222003"></a>

#### `graphql_rules.graphql_settings.max_batched_queries` property

Type: `"number"`. Computed.

Specify maximum number of queries in a single batched request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-2001111021022233-3121220022121200-0312223001311022-2223011321000030-0002332132300032-2133223113021333-2002011302023320-1003003113103213"></a>

<a id="canonical-3213002033000020-2300102322202120-2002201131232111-0311113330122101-1002113310302212-1321112030213332-3202120301011212-2230101121002130"></a>

#### `graphql_rules.graphql_settings.max_depth` property

Type: `"number"`. Computed.

Specify maximum depth for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-3313213313232033-2302023321333213-2103301233233002-2213132222023332-3321013101100030-3222030010022031-3202102322120202-0312222200200023"></a>

<a id="canonical-0303311033332221-0031032201300020-0111001013031010-1122002022003322-1301110032101013-1223301221323222-1121020122210101-3000011110201003"></a>

#### `graphql_rules.graphql_settings.max_total_length` property

Type: `"number"`. Computed.

Specify maximum length in bytes for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16386,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16386"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  }
}
```

<a id="canonical-1102101103100110-1033113013022203-0333322013112001-1331332320212301-2120211002300001-3321330212220023-2223212131210130-0302102322213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings.disable_introspection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3301232011322201-3322122212320230-0212320203323203-1233020330031223-0013321230233213-1003032122130202-3223331312211021-0020303030213011)
- [graphql_rules.graphql_settings](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0133031100210221-1323313022121310-3033102310033213-0203322113001232-2310323131112130-0100031120121320-3200031210201120-0002020120000011)
- graphql_rules.graphql_settings.disable_introspection

<a id="canonical-3312332103231303-1302230113013332-3112203320210333-0023031313000213-2312133003200000-1232103110011221-1232233002030210-2100032221200020"></a>

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

<a id="canonical-3213213230203211-2222231212030110-0133000120003020-2131103210121022-0020100211320131-2301222011003201-1233101333302102-0323100212330321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings.enable_introspection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3301232011322201-3322122212320230-0212320203323203-1233020330031223-0013321230233213-1003032122130202-3223331312211021-0020303030213011)
- [graphql_rules.graphql_settings](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0133031100210221-1323313022121310-3033102310033213-0203322113001232-2310323131112130-0100031120121320-3200031210201120-0002020120000011)
- graphql_rules.graphql_settings.enable_introspection

<a id="canonical-1300322301103131-0012020233201020-2000111122113121-3323102101213220-0023022010023323-0003130333333202-1310310031033220-2023333223000220"></a>

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

<a id="canonical-3031332133020023-1012223131330320-3213321102112313-0023203102000220-3210011311100102-0030000023202210-2223303313202332-1220102233031031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3301232011322201-3322122212320230-0212320203323203-1233020330031223-0013321230233213-1003032122130202-3223331312211021-0020303030213011)
- graphql_rules.metadata

<a id="canonical-0220230131132130-1001232112111231-2020232220133130-0203023103023311-2010231003120302-3001131223213100-0233213320110121-2012301201122032"></a>

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

<a id="canonical-0323022011313200-2331231122310202-1113121321000101-1021121121000300-3210133213203221-0013130031222211-3210121332223223-3211031322221123"></a>

### Direct properties for `graphql_rules.metadata`

<a id="canonical-3221221230303233-0211313233321211-3311330000222300-1110332211331221-1321310100110212-1300110312010232-3012133110133223-3312120020112212"></a>

#### `graphql_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1010331202233122-1000130302300301-1102010230323101-1023011222030011-0132212031032123-3002311003302230-2023030001301110-1232112220311133"></a>

<a id="canonical-2320333230232201-3323330030101020-3103110232033031-3030233211323012-0030033031201210-3310012202113023-0311330111323113-0011003100031023"></a>

#### `graphql_rules.metadata.name` property

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

<a id="canonical-3121322313302002-0000113322110302-1200010200311302-3202123222100100-3301213132221000-0320220203200331-0222020222321022-1221232232303020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.method_get` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3301232011322201-3322122212320230-0212320203323203-1233020330031223-0013321230233213-1003032122130202-3223331312211021-0020303030213011)
- graphql_rules.method_get

<a id="canonical-3133220201333033-1000301003022020-3302203130002230-2023021023001211-1021201312302011-2113312200033121-3030333232302201-2010010212213332"></a>

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

<a id="canonical-2033130003220333-0221001233022330-3023102331213101-3332330012100221-0001122211302312-2000302210110133-2131003002321231-2101120330323213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.method_post` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3301232011322201-3322122212320230-0212320203323203-1233020330031223-0013321230233213-1003032122130202-3223331312211021-0020303030213011)
- graphql_rules.method_post

<a id="canonical-0013302332312311-0201102122110302-1000220333033222-0013110001132213-2230021011213210-0211321332123022-0103230102222033-3221222233230230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for method post.

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

<a id="canonical-0000202303000030-1313011110300110-2010231120333333-2012111302020010-2232013230200230-3211311320303033-1310131202202210-0323010122200301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- http

<a id="canonical-3231011100222011-2021033110010331-0133011013121213-0200330020220303-2020300012110230-0210201023013011-3322222232011323-1100003020003301"></a>

Type: `"single"`. Computed.

\[OneOf: http, https, https\_auto\_cert; Default: https\_auto\_cert\] HTTP Choice. Choice for
selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

OneOf alternatives in this subsection:

- [http](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3231011100222011-2021033110010331-0133011013121213-0200330020220303-2020300012110230-0210201023013011-3322222232011323-1100003020003301)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2332222233303113-3031230002222300-1012101102102000-2100300212232203-2021132123112230-0130132120023233-3102221102112300-1230003103212203)
- [https_auto_cert](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1133023022021021-3123031201112323-2013213230333010-2121333120032213-2302232102103320-3300321310201302-0211002132112210-0203032111320311)

Select alternatives according to the provider validators above.

<a id="canonical-3101000230321320-1331123102132121-3012203032301232-2203133303202021-2102331131333120-3023331123332231-3330311231210231-2101112003131330"></a>

### Direct properties for `http`

<a id="canonical-1300030210331220-2100210232111331-1030003020103231-0210130303221310-2300021113221310-3120311223300223-2031102223320111-2031301223111333"></a>

#### `http.dns_volterra_managed` property

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

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

<a id="canonical-2330110332002132-0303020311331311-1000021123000011-2300120311120113-3331212330212031-2102121130323112-2002121230210310-2303320333303021"></a>

<a id="canonical-1322201230203003-2111102122312221-2221023132213102-0122300331123331-0222032221022311-3233100103310332-2212012020331113-3321022300321233"></a>

#### `http.port` property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1230013311223110-3011301133100322-0012002012123013-3230332211200230-2311031300013211-1212101031132102-2122032313111121-3331032021222302"></a>

<a id="canonical-0113300102213310-0011030013012213-3213333213101331-1323301231332310-0310011331330211-2231231201313202-1303321120330121-3123002011233121"></a>

#### `http.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- https

<a id="canonical-2332222233303113-3031230002222300-1012101102102000-2100300212232203-2021132123112230-0130132120023233-3102221102112300-1230003103212203"></a>

Type: `"single"`. Computed.

Choice for selecting CDN Distribution with bring your own certificates.

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

<a id="canonical-1321221013313331-2011332331012310-2022202321211233-1301223311030110-0131213320203213-2212202112121131-2333021122312232-3233222130023232"></a>

### Direct properties for `https`

<a id="canonical-2313232321212203-1000330122232111-0210130112223303-1001001103012322-0213330313020222-2122221000013012-1200022300221031-3301223131002030"></a>

#### `https.add_hsts` property

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-1132132013012033-3020233301110120-3233302120102120-0033032200210223-3011032202000220-1323101322233120-2320310231212222-3032223133221220"></a>

<a id="canonical-3012201322020302-2210202122222321-2010302101033302-2022110032132122-0223102231130022-2001230222021302-0030000000032132-2200103132220323"></a>

#### `https.http_redirect` property

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

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

- [tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211): complete subsection reference.

<a id="canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- https.tls_cert_options

<a id="canonical-1311131101031113-2300033010320110-0130122302023221-2213321030022112-0120320001120001-0301323113003130-3211011322332130-2312222221120222"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert options.

Additional upstream details:

TLS Certificate OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_inline_params\"]"
}
```

<a id="canonical-0132000233000201-3301110010132022-2001320311311112-0333320100232003-3030122001330323-2103312223030122-1330001331000222-0203220230112132"></a>

### Direct properties for `https.tls_cert_options`

- [tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032): complete subsection reference.

- [tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010): complete subsection reference.

<a id="canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- https.tls_cert_options.tls_cert_params

<a id="canonical-3200222033222002-2311021101101100-2120301233132222-0232112310201331-2312313303301332-1320111000331312-0313221011203302-2333102032030021"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-0030213200302021-2311131023010203-1220121111201323-0223122101020333-0210131010102122-2313131323131331-3013212310212112-2110030202010212"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params`

- [certificates](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2121122300022221-0200130311231330-0213133012121323-1021331300113312-1201333213313100-0023101013330321-1010310011313002-0333321131022330): complete subsection reference.

- [no_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2233213322023302-0201231112111023-2022102130320131-3303002312332013-0130322200302312-1030302101023021-2322022221103323-1000013130032230): complete subsection reference.

- [tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2320332103223010-3111133032133213-0032102303222331-3333110021133013-0000222330030301-0013103023201012-1210101021002101-3021320112213300): complete subsection reference.

- [use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0230022233122110-2020333321222312-0001321132011133-3003223021323310-3120231222211320-3232123200011202-0212003121201201-1011001220020113): complete subsection reference.

<a id="canonical-2121122300022221-0200130311231330-0213133012121323-1021331300113312-1201333213313100-0023101013330321-1010310011313002-0333321131022330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- https.tls_cert_options.tls_cert_params.certificates

<a id="canonical-0032103223003331-3133321221030230-0012131233203112-2100222330302312-0122021333303003-0220311231020002-2212232032203323-3010123001331200"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1310121111230121-0203101031201120-1201300112121103-1021130311131131-2033201121301202-2133002322220013-2112201131220032-2211320202230320"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.certificates`

<a id="canonical-2331133002300021-1213101130203033-1021022011223301-3033020121002231-1011320233232223-0110021202001033-3111120323100120-1103330103123032"></a>

#### `https.tls_cert_options.tls_cert_params.certificates.name` property

Type: `"string"`. Computed.

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

<a id="canonical-1201232223100320-2311102301122000-0331131013122313-3233202300001301-0130121012132012-2130021102131201-2110113233332000-1201033222211110"></a>

<a id="canonical-3022003203231210-1003000313233022-3102003320022102-1300232020023130-3223320301021313-0033322301210223-2130032332231120-3320022003130001"></a>

#### `https.tls_cert_options.tls_cert_params.certificates.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-3121312230222331-2223113302312210-3331202302230210-3133111022001113-3002002003310003-0320133112310032-2121132022303223-0113222100011302"></a>

<a id="canonical-0301213031212323-2220012200220023-3102331303121020-1130130303023222-1311120030200121-0310313121233103-1333001023210330-1313111203011330"></a>

#### `https.tls_cert_options.tls_cert_params.certificates.tenant` property

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

<a id="canonical-2233213322023302-0201231112111023-2022102130320131-3303002312332013-0130322200302312-1030302101023021-2322022221103323-1000013130032230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- https.tls_cert_options.tls_cert_params.no_mtls

<a id="canonical-0330313312321201-1210302111010120-3033312203333200-3122133010303311-0223011333313000-2120220202000333-3231323100320221-3132000220003101"></a>

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

<a id="canonical-2320332103223010-3111133032133213-0032102303222331-3333110021133013-0000222330030301-0013103023201012-1210101021002101-3021320112213300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- https.tls_cert_options.tls_cert_params.tls_config

<a id="canonical-1201021133333300-1031303332332201-1232100211332330-1323301310301332-0003110301233211-3220330002233120-2131002232203301-1120120020110133"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-0312223013133203-2132001311112221-3230110201131010-3030031210212003-0000303211020000-1210330230111312-0110010033121201-0013100100102213"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.tls_config`

- [custom_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0223002302121223-0102323200232021-2021113233130003-1230211331312012-0011020301100102-1322121221000001-2123322202111133-2330213112332110): complete subsection reference.

- [default_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2303103212210122-3302220112101331-1331112022010311-0232021030110003-0211323022312310-3013320021311231-1120231322113233-0223122110211021): complete subsection reference.

- [low_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0033103100100333-2130213103101130-3103301122001333-2012223301112111-1012212121130020-1232213212132221-3110000321103112-2102102332101202): complete subsection reference.

- [medium_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1232322311030230-0030310101210203-0303002120203310-2000133013312001-0231033211010202-2311213001023100-3232003113003212-2000033312212301): complete subsection reference.

<a id="canonical-0223002302121223-0102323200232021-2021113233130003-1230211331312012-0011020301100102-1322121221000001-2123322202111133-2330213112332110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2320332103223010-3111133032133213-0032102303222331-3333110021133013-0000222330030301-0013103023201012-1210101021002101-3021320112213300)
- https.tls_cert_options.tls_cert_params.tls_config.custom_security

<a id="canonical-2320322202321320-0020101110123310-2330130111312130-0321230033103331-0033222030030323-3221032331230100-3120112321130131-0201120012122000"></a>

Type: `"single"`. Computed.

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-3120233030010230-0202332032333110-3001301231132220-2201210031101132-3112101300130120-1323001120122210-3210332322121121-1121223202120123"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.tls_config.custom_security`

<a id="canonical-0022212020121002-3212200302123201-3001102333112000-1330000113333223-0302010222312201-2000303221213122-1030311323300122-2113323303032110"></a>

#### `https.tls_cert_options.tls_cert_params.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0301301021003203-2132331231222300-1220033010003111-3031011322203303-2030123022222101-3233020020221133-0330330032033221-1113022113213120"></a>

<a id="canonical-0211000313330111-3311302112331131-0120110320132213-2132021323321130-1021113121002022-0320301202111321-3123231333133100-1220303231111213"></a>

#### `https.tls_cert_options.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2333011331323110-3032122312101010-2023011200121000-3101310031012212-2302131133223231-1231213322030110-1311120222202210-3021122021223130"></a>

<a id="canonical-1221311100330120-0302001133031020-3212201021100130-2300132011032320-0312132123133013-0212020032112203-2112210013113020-1213233301121300"></a>

#### `https.tls_cert_options.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2303103212210122-3302220112101331-1331112022010311-0232021030110003-0211323022312310-3013320021311231-1120231322113233-0223122110211021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2320332103223010-3111133032133213-0032102303222331-3333110021133013-0000222330030301-0013103023201012-1210101021002101-3021320112213300)
- https.tls_cert_options.tls_cert_params.tls_config.default_security

<a id="canonical-3330232013120332-1210130011013323-1101300311100312-2222030123213333-0311012022223100-2130110133221112-1000022233123012-3100100000220200"></a>

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

<a id="canonical-0033103100100333-2130213103101130-3103301122001333-2012223301112111-1012212121130020-1232213212132221-3110000321103112-2102102332101202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2320332103223010-3111133032133213-0032102303222331-3333110021133013-0000222330030301-0013103023201012-1210101021002101-3021320112213300)
- https.tls_cert_options.tls_cert_params.tls_config.low_security

<a id="canonical-2133310212232123-2011222322103211-2230030230223111-1000000031211010-2323101123222220-0302320312032033-0203000032103223-3123120022312310"></a>

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

<a id="canonical-1232322311030230-0030310101210203-0303002120203310-2000133013312001-0231033211010202-2311213001023100-3232003113003212-2000033312212301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2320332103223010-3111133032133213-0032102303222331-3333110021133013-0000222330030301-0013103023201012-1210101021002101-3021320112213300)
- https.tls_cert_options.tls_cert_params.tls_config.medium_security

<a id="canonical-3200013033210320-3210013212121000-3212232212132300-3220222232312312-3011022322113303-1021101012213122-0000113313321033-1222012013202331"></a>

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

<a id="canonical-0230022233122110-2020333321222312-0001321132011133-3003223021323310-3120231222211320-3232123200011202-0212003121201201-1011001220020113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- https.tls_cert_options.tls_cert_params.use_mtls

<a id="canonical-2310001322300231-1233003310101203-3103221220103333-2023131011231013-3213021322222001-2103310323020123-0333002123123333-1320102332301210"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-3011201132012213-1031102131103031-0331331312023002-2133101231230220-0123332231313002-3021222231002311-2333212103102331-3020013100200301"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.use_mtls`

<a id="canonical-2221332011221302-2021221212030021-3230121130121230-2323233122302221-2012110203111330-1130102331313311-2111330022320032-0123032212133321"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2322210031003101-3301330312111132-2301333223022111-0131001012313212-3222301323322121-2121023023322310-0332002110120031-0313012120222013): complete subsection reference.

- [no_crl](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3202101313332103-1100310100331210-1001012213013013-3033002013030123-3003030221213001-1333220221001100-2221113013132023-2211332030310323): complete subsection reference.

- [trusted_ca](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0313031200323100-0122111133001211-3012031311313202-3113110333321013-3300210302122132-1312132233021100-2111101312122022-2023023120001013): complete subsection reference.

<a id="canonical-2033020222133331-1131213200323013-2213101101112312-0221003221113223-3022302333200120-0122120310032203-1133233032323003-0220220233123330"></a>

<a id="canonical-3301313002003100-2011113221202301-3001133130133303-1002232333002023-3022002313311113-3202022213312203-2331112201200101-3101130121033313"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3301021113103010-1033213223133122-0323131320221131-0012232010220310-3010212030030202-1300120023123032-3200211102212211-2210113031330233): complete subsection reference.

- [xfcc_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2303223030323230-3233330232301333-3022102322033232-1030013321201033-0233332131223323-1331330000210132-0122023103221030-2010301231321223): complete subsection reference.

<a id="canonical-2322210031003101-3301330312111132-2301333223022111-0131001012313212-3222301323322121-2121023023322310-0332002110120031-0313012120222013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0230022233122110-2020333321222312-0001321132011133-3003223021323310-3120231222211320-3232123200011202-0212003121201201-1011001220020113)
- https.tls_cert_options.tls_cert_params.use_mtls.crl

<a id="canonical-0310132312233210-2030000210232132-0103220122103302-2133003322300031-1013113032121021-3233212230230010-3030021000313232-3101021213230201"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2101330110320111-1101103011222202-0011021123110311-2032130132213102-2130320211212132-0002303330211113-3031323333203222-2132223230213131"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.use_mtls.crl`

<a id="canonical-1201302221113131-1130330000313231-1312021022100200-1133103231111303-2203113113021032-0312230003121022-0320003222223001-1230102213122130"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.crl.name` property

Type: `"string"`. Computed.

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

<a id="canonical-3321220102013202-0022130220301012-1212322311030131-0030031213223200-1022311302222110-3003202101102103-0133131322103301-3132212123130202"></a>

<a id="canonical-3210333302311302-1300233020000003-0312212121312012-0130203022330111-3323203212230333-1202022121132321-1332023300213330-2102303320113223"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.crl.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-1003231202122212-3010321112033203-3113232213330211-3221030321202123-1221232133211212-3231210030212321-2013320232301200-3103322101210313"></a>

<a id="canonical-1321001023220201-3302230100230322-0301313133023132-1030221023111212-0332212023020121-1330301103122231-3302203310321033-2202332332030202"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.crl.tenant` property

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

<a id="canonical-3202101313332103-1100310100331210-1001012213013013-3033002013030123-3003030221213001-1333220221001100-2221113013132023-2211332030310323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0230022233122110-2020333321222312-0001321132011133-3003223021323310-3120231222211320-3232123200011202-0212003121201201-1011001220020113)
- https.tls_cert_options.tls_cert_params.use_mtls.no_crl

<a id="canonical-2121020330230033-1110112200023320-0030032220212332-3013211012332213-1111102212331223-2330023130211323-3311312232303300-2103122133300122"></a>

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

<a id="canonical-0313031200323100-0122111133001211-3012031311313202-3113110333321013-3300210302122132-1312132233021100-2111101312122022-2023023120001013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0230022233122110-2020333321222312-0001321132011133-3003223021323310-3120231222211320-3232123200011202-0212003121201201-1011001220020113)
- https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-3110311101013321-1101311030023322-2333211323312132-0032212330020330-1032333222331312-0211313211200302-0332033230310001-2300321122113002"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0220321002120010-0112002303203003-2201103011230023-1203322230121011-0010202212301120-3322021000201123-0221313220213200-3233321333032333"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-1113301312131011-0013331223111020-2101031110323302-2133010131022330-0031121012130120-2221000200010112-0222020132023111-0203313101231131"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca.name` property

Type: `"string"`. Computed.

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

<a id="canonical-3002112111132312-3033022322133213-2212310021133120-2320031220302311-1301023302020231-1330212312222031-2120131102121110-0001211211012030"></a>

<a id="canonical-1022313231131022-1010201203320333-3030101333011022-1112121303231001-2321330300122102-1202202021000301-3001132103220030-2203201110221313"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-0321222311021003-2013002331203021-2321323101212222-1022200032201332-1213000001213231-1011223022201113-1201203122221133-1200300220120001"></a>

<a id="canonical-3122011230301213-1033213232021231-0032013210223203-1220202220111132-0123123202323323-0033223333020221-3101302021131100-1021022310222020"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-3301021113103010-1033213223133122-0323131320221131-0012232010220310-3010212030030202-1300120023123032-3200211102212211-2210113031330233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0230022233122110-2020333321222312-0001321132011133-3003223021323310-3120231222211320-3232123200011202-0212003121201201-1011001220020113)
- https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-2330200222112033-3300211203121103-0101313221002131-1300000122222122-1121233113012203-1012123230200022-1301020211113211-0223001000222322"></a>

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

<a id="canonical-2303223030323230-3233330232301333-3022102322033232-1030013321201033-0233332131223323-1331330000210132-0122023103221030-2010301231321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2102213030323200-0011230110200231-0212313130013000-0223212220103300-2133100002311113-2123312123122300-2130112200001323-1022310011132032)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0230022233122110-2020333321222312-0001321132011133-3003223021323310-3120231222211320-3232123200011202-0212003121201201-1011001220020113)
- https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-1330213203302000-3210020111301000-3302212333220033-0233312311121102-2220121221212121-2302120301210002-1323121121103223-2122112002031000"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-3331221100121020-0112112121110213-2203303221110133-1033221110232122-3100333331223333-0101223313002111-2100332032200133-3203133322223200"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-1312100021113011-0023330323210032-1103220222000201-1311102023331313-3312222113001013-1030112120002310-0232020123231333-1032320003103120"></a>

#### `https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- https.tls_cert_options.tls_inline_params

<a id="canonical-1022302201303222-0130323001311111-3312013000003232-2020201011002120-0310022101011230-3002001033100310-3212023233311100-0113133201232301"></a>

Type: `"single"`. Computed.

Configuration parameter for tls inline params.

Additional upstream details:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-3000333200101230-2031131022032112-0203032112110030-1231011000233230-2210313330111313-3310202201002220-3321000301211202-1123100111333330"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params`

- [no_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3322200222332111-0001331222323232-0032303332311133-3202112311312222-2321303001332011-3022311103300311-1132002111201331-1023202212322211): complete subsection reference.

- [tls_certificates](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0031120333223322-2313123203320101-3023130220033130-3313312120330220-1021111020212220-3033121120011112-3320300302312222-1211221012123301): complete subsection reference.

- [tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1113100120020201-0123033213020223-2302310033123110-2201313231300232-0101133330302212-2012122210031301-1112131221311003-0212331320023121): complete subsection reference.

- [use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1020121201233003-1212133021201001-2331021023212310-0111220113103030-3222211102322103-2331113010121022-0102011311113033-3111232102002113): complete subsection reference.

<a id="canonical-3322200222332111-0001331222323232-0032303332311133-3202112311312222-2321303001332011-3022311103300311-1132002111201331-1023202212322211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.no_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- https.tls_cert_options.tls_inline_params.no_mtls

<a id="canonical-3102122331012122-0333002301033131-0100232203102002-3222102101331000-0202030303013111-3313120112010301-1223123211300011-0210133230230101"></a>

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

<a id="canonical-0031120333223322-2313123203320101-3023130220033130-3313312120330220-1021111020212220-3033121120011112-3320300302312222-1211221012123301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- https.tls_cert_options.tls_inline_params.tls_certificates

<a id="canonical-2323202223323133-3230021012010112-3011200010200123-1212322230010320-0002312222013012-3330020003031000-1032212321121020-0002213111032322"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2200213233130011-0102310001132310-3201132001232020-3333101123211323-3021302223213201-0200011220030321-2112112113212002-1221200032120010"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_certificates`

<a id="canonical-3131012202003223-3222000113133203-1303132202002011-1231133023332332-2011110231220022-3233003200200112-1332323223332232-1033202303231331"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.certificate_url` property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3130213200003112-3311111233100201-3200022232203233-3021101023312303-1131231110110233-3001021333222233-1233222123002320-2112121200212233): complete subsection reference.

<a id="canonical-3300313320113223-1101122333121322-3021031103313221-3122133021220020-2130320322321020-1130111031222331-0231300210202100-2321120030220322"></a>

<a id="canonical-2332311033022301-2213100132132231-0310102003211102-1000300200010313-1210323212210002-3111012222233333-0231320330232323-1001313200030232"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1213111311321322-0201122101222233-0020031222212023-0032223230010213-1120320110203132-1210203121333220-2130131101203322-2013331000200202): complete subsection reference.

- [private_key](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0331111221200031-0110011132222030-2003213230111332-3100312122311003-3112100333012122-0223022233022313-2320110220002232-1301133113112023): complete subsection reference.

- [use_system_defaults](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0022003333102111-2113231200021311-3323313330022120-1212210123303301-3302101332010122-1100320300131121-1300100233333113-2113022002120311): complete subsection reference.

<a id="canonical-3130213200003112-3311111233100201-3200022232203233-3021101023312303-1131231110110233-3001021333222233-1233222123002320-2112121200212233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0031120333223322-2313123203320101-3023130220033130-3313312120330220-1021111020212220-3033121120011112-3320300302312222-1211221012123301)
- https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms

<a id="canonical-3030110223012230-2300233011032003-1302213233011030-3210121322330221-1321031311110303-1231003023023231-0031302021213301-3011102012113332"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2213023100310210-3333212310231000-3112101210000322-3300013212301313-1331110013301002-1302121102112013-1231111231032233-3201221002211200"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms`

<a id="canonical-3321100201110322-2201320102123220-2310311221011122-0132022320103002-1113120213332313-2011023132032222-2321112230311133-0013202201211113"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

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

<a id="canonical-1213111311321322-0201122101222233-0020031222212023-0032223230010213-1120320110203132-1210203121333220-2130131101203322-2013331000200202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0031120333223322-2313123203320101-3023130220033130-3313312120330220-1021111020212220-3033121120011112-3320300302312222-1211221012123301)
- https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-2010100202102001-2001020101131311-3130312300203120-0202000003231013-2111130013211032-2112003313112230-1011231133133231-2213303000030310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-0331111221200031-0110011132222030-2003213230111332-3100312122311003-3112100333012122-0223022233022313-2320110220002232-1301133113112023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0031120333223322-2313123203320101-3023130220033130-3313312120330220-1021111020212220-3033121120011112-3320300302312222-1211221012123301)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key

<a id="canonical-1131110021023002-1323222020132320-1312322120303112-0020111300000012-1211133000232133-2131101310200133-0301232021303232-0133131133111031"></a>

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

<a id="canonical-1201321033330001-0323113020202012-3122333202032000-2030222111211021-1000032111333221-1133033330310311-1332330010230031-1100302202300020"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0121222030120303-0233101333323220-2020330112011131-1301332002103333-2303232103201201-1303212102003110-1102133020131223-0322303201331310): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3001010231221012-3112310030312122-3032110202131202-3120331221113131-1011031301102110-3101322133012113-0211323011232123-1120232110103111): complete subsection reference.

<a id="canonical-0121222030120303-0233101333323220-2020330112011131-1301332002103333-2303232103201201-1303212102003110-1102133020131223-0322303201331310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0031120333223322-2313123203320101-3023130220033130-3313312120330220-1021111020212220-3033121120011112-3320300302312222-1211221012123301)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0331111221200031-0110011132222030-2003213230111332-3100312122311003-3112100333012122-0223022233022313-2320110220002232-1301133113112023)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2031311111033131-3213133320303002-3110020021311310-3100010110223033-1012210030203131-2101033330330310-2003303020000030-3311203203110113"></a>

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

<a id="canonical-2032033021101022-0221111132111133-2222102220110001-3323133330121311-0122120322301033-0112321123023230-0122113221300312-0321033110202130"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-1010212002303210-3310120121032200-2303320112002332-0121023120133020-2331210222232132-2221133102223132-1223313121211320-0021213012111130"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2302320310130112-0310212311001311-2021202121332221-2031203223212233-3200311102012313-3111321320312213-0300103123302231-1210031131020133"></a>

<a id="canonical-2010132000202133-3210000000002133-2321002333021023-3322021021321013-3101321020113013-1201313032311313-0231203302021132-2010101031202301"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info.location` property

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

<a id="canonical-0230301111100331-2210001310233202-2013303013202233-3201313001323210-1111210232032120-1333312230010100-2221113000201201-2213001323220033"></a>

<a id="canonical-0322001030003311-3010333101103300-2303202210300321-3213102112212230-3312011303212120-2131023012123100-0211202301330022-2300331132203110"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-3001010231221012-3112310030312122-3032110202131202-3120331221113131-1011031301102110-3101322133012113-0211323011232123-1120232110103111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0031120333223322-2313123203320101-3023130220033130-3313312120330220-1021111020212220-3033121120011112-3320300302312222-1211221012123301)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0331111221200031-0110011132222030-2003213230111332-3100312122311003-3112100333012122-0223022233022313-2320110220002232-1301133113112023)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-0120232230302303-3230033211133203-0202013202011001-0113002103001333-0232203121032330-2232133300120203-1212120331323133-1020012030230203"></a>

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

<a id="canonical-0211002032221100-0230030133012220-3132321212002113-0333301212111331-2213311013002020-1011210213030202-3300031100211011-2210300320100201"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0122132211112022-0112233012121102-1212312110210122-2331000201110333-3233100333002221-3323111013302133-2330222233101120-0110231033201003"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0013001321303303-2233201211000113-0321121331212322-2231110012121221-2300223133112311-2123300032312031-1221002123210233-2103202312102033"></a>

<a id="canonical-1330212211021321-2331032031023233-2023203103300130-1012102103112103-3332212332133033-1313010033012210-1031232110231300-3121231213202310"></a>

#### `https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info.url` property

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

<a id="canonical-0022003333102111-2113231200021311-3323313330022120-1212210123303301-3302101332010122-1100320300131121-1300100233333113-2113022002120311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0031120333223322-2313123203320101-3023130220033130-3313312120330220-1021111020212220-3033121120011112-3320300302312222-1211221012123301)
- https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults

<a id="canonical-3023033020002013-3222031002022332-0132002102200330-2121113201010300-1132323100021133-1212132221333321-1120113010332230-2221131101331120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-1113100120020201-0123033213020223-2302310033123110-2201313231300232-0101133330302212-2012122210031301-1112131221311003-0212331320023121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- https.tls_cert_options.tls_inline_params.tls_config

<a id="canonical-3301002210211033-1332233213221323-0032132323122120-3020000110122021-1002020130302202-3322223211012311-2331120311300032-2222223112311101"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-2112231103000110-1122300112021131-3033103230232113-1311212022130301-0020212203101130-2302132101331033-0212313011112113-2232130122110032"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_config`

- [custom_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0013021312023032-3002033233213003-2110322312221133-1331302002303023-0223102333222310-3021102202231013-1320201331331113-1311323221011301): complete subsection reference.

- [default_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1133323030113120-1202111320021313-2200111020332122-3030312302133200-1120230133030331-1313010112003210-1232100102310123-1113320103011233): complete subsection reference.

- [low_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3122022332010232-3030233023211321-0130101130110113-1023333123121130-2321001200013231-2013001231211011-3201133100130200-3000221202011232): complete subsection reference.

- [medium_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1113003211013103-2102201110023000-0220203221113302-1200123023113023-2021000001322313-2123020102321113-2220013010330212-1203210121020300): complete subsection reference.

<a id="canonical-0013021312023032-3002033233213003-2110322312221133-1331302002303023-0223102333222310-3021102202231013-1320201331331113-1311323221011301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1113100120020201-0123033213020223-2302310033123110-2201313231300232-0101133330302212-2012122210031301-1112131221311003-0212331320023121)
- https.tls_cert_options.tls_inline_params.tls_config.custom_security

<a id="canonical-3123011102213011-0223203132021111-2200200011213213-1001332310001031-1301310223130212-3313032232130212-3101123013001301-3333320322033330"></a>

Type: `"single"`. Computed.

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-1323121012211212-2111233113032200-3222123102233200-0012212100032103-2320012232033231-1003021320231033-1033100313002333-3130332202333222"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.tls_config.custom_security`

<a id="canonical-3211330100322131-0200223020221202-3233001110101010-2210022333103323-2230003220333212-1002223000101303-0223033032322322-0130002232203031"></a>

#### `https.tls_cert_options.tls_inline_params.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3122300321232303-1203303011121032-0312130100323130-2310212310213312-0000212003101331-1212011231222031-2201001022132310-1333031331021211"></a>

<a id="canonical-1331003123202130-0101311033001013-1221322111121131-1131003330301131-1022321332121121-0210021220022113-3020300230230003-1321032313302123"></a>

#### `https.tls_cert_options.tls_inline_params.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1202022211223310-1010221233211300-0112001123220212-3130231110000103-3321332130102113-0100301322230311-2210133022311332-2011330313020000"></a>

<a id="canonical-0101130122212212-3232313210333203-1101212303220011-0122010312110332-2032303121033231-2102213221110030-2033032113330331-1112023203102101"></a>

#### `https.tls_cert_options.tls_inline_params.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1133323030113120-1202111320021313-2200111020332122-3030312302133200-1120230133030331-1313010112003210-1232100102310123-1113320103011233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1113100120020201-0123033213020223-2302310033123110-2201313231300232-0101133330302212-2012122210031301-1112131221311003-0212331320023121)
- https.tls_cert_options.tls_inline_params.tls_config.default_security

<a id="canonical-0102131133302210-0121323213101001-1222133310232131-0121332220220112-2130220103221221-2201230133332320-0001222030330031-1231310210301302"></a>

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

<a id="canonical-3122022332010232-3030233023211321-0130101130110113-1023333123121130-2321001200013231-2013001231211011-3201133100130200-3000221202011232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1113100120020201-0123033213020223-2302310033123110-2201313231300232-0101133330302212-2012122210031301-1112131221311003-0212331320023121)
- https.tls_cert_options.tls_inline_params.tls_config.low_security

<a id="canonical-1030121120320022-1332310022032213-1010001023033112-3230010310122323-1130101010001010-2232313010003221-0100112303100333-1333011221320312"></a>

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
