---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-0030012311113231-2132331032100200-0213202301120101-2322211313312320-1120200330113211-0132032212011312-2021211302302011-3311303303130131"></a>

#### `advanced_options.outlier_detection.consecutive_gateway_failure` property

Type: `"number"`. Computed.

If an upstream endpoint returns some number of consecutive “gateway errors” (502, 503 or 504 status
code), it will be ejected. Note that this includes events that would cause the HTTP router to return
one of these status codes on the upstream’s behalf (reset, connection failure, etc.).
Consecutive\_gateway\_failure indicates the number of consecutive gateway failures before a
consecutive gateway failure ejection occurs. Defaults to 5.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

<a id="canonical-0113320002320321-3020322133233101-2201201130231111-2230033232130112-2002122123201212-2000201211100322-0223320302102311-0203312201333100"></a>

<a id="canonical-1003333212002300-3030000011313003-1311110323002131-2130000202303213-3330032302323203-3310021231100302-3320210001222200-2323122202131032"></a>

#### `advanced_options.outlier_detection.interval` property

Type: `"number"`. Computed.

The time interval between ejection analysis sweeps. This can result in both new ejections as well as
endpoints being returned to service. Defaults to \`10000ms\`.

Additional upstream details:

Defaults to 10000ms or 10s. Specified in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3211300220032332-2220121022233030-2302332031013213-2322211311310211-2210211111312330-2032230032123301-1102201323103123-1300220220332201"></a>

<a id="canonical-3123013221320002-0010100231203031-3212002323121003-0303101230331013-2032231112131311-0313300010231310-2011101231013130-3132121112030320"></a>

#### `advanced_options.outlier_detection.max_ejection_percent` property

Type: `"number"`. Computed.

The maximum % of an upstream cluster that can be ejected due to outlier detection. but will eject at
least one host regardless of the value. Defaults to \`10%\`.

Additional upstream details:

The maximum % of an upstream cluster that can be ejected due to outlier detection. Defaults to 10%
but will eject at least one host regardless of the value.

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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-2112302233203232-2312130221113230-0302122201023112-3103002001321120-3131312111202222-2011003132202023-3012300321002230-1313002012030031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advanced_options.proxy_protocol_v1` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.proxy_protocol_v1

<a id="canonical-3131021033021122-3332300331321230-2213310130323231-0332020030322311-3310112321122301-3023232230203212-3210233003132203-2111213220202321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for proxy protocol v1.

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

<a id="canonical-1310101030312333-0010032203021230-0132232233023132-1230310123232110-0212203312131220-3121132023122030-0103100100313123-0301130111123320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advanced_options.proxy_protocol_v2` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [advanced_options](data-sources--origin_pool--reference--group-001.md#canonical-0213102330322032-1003131310312032-3100111012020311-3311313223000032-3030131320332203-2101113021210202-2333030123232123-3203300203331130)
- advanced_options.proxy_protocol_v2

<a id="canonical-3303032000000002-1211212303130132-3133011323011230-1210302310232032-2203102030322323-0211220011222302-1322211323103132-2010322102210113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for proxy protocol v2.

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

<a id="canonical-3232200323013011-3012121113003102-1101211132122231-3303101001222131-3132321221221112-0012322000210223-2010320011332333-2212232121313131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `automatic_port` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- automatic_port

<a id="canonical-3330030101213132-1023001113001113-2323003003322122-3003122331131110-1132010030223132-1131332022230111-3132311313331113-2130313202000312"></a>

Type: `["object", {}]`. Computed.

\[OneOf: automatic\_port, lb\_port, port\] Enable this option

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

- [automatic_port](data-sources--origin_pool--reference--group-002.md#canonical-3330030101213132-1023001113001113-2323003003322122-3003122331131110-1132010030223132-1131332022230111-3132311313331113-2130313202000312)
- [lb_port](data-sources--origin_pool--reference--group-002.md#canonical-3332320330310202-3001331103323210-0100103311013120-0002100310001103-2000333310000320-3010022210312130-1113032332102000-2210230130200022)
- [port](data-sources--origin_pool--reference--group-001.md#canonical-1331233022212123-0232020012320101-0313013123210003-2023013132030210-1213332332113330-0220222032331222-1133021111301032-2011132211203233)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321301103022100-0002302311013132-3122331220210333-2233013212333133-1323332021333233-0320103330123010-3013120233010203-2201003110133022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `healthcheck` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- healthcheck

<a id="canonical-1013131122022033-3120111310102210-0122223231303100-1002203002313300-3112312220330112-0123221302122033-2111003202132002-2200220232301001"></a>

Type: `"list"`. Computed.

Reference to healthcheck configuration objects. Defaults to \`\[\]\`. Server applies default when
omitted.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-1111232223103301-0203322220312303-0003132301003133-1330301210203022-3010002311303311-0202202310011101-0122322222311203-2231131010120132"></a>

### Direct properties for `healthcheck`

<a id="canonical-1020202033220110-3030301031212233-1331211302231211-0022012012313121-1313000312202122-3213323223122230-0013220011313121-0320012130000011"></a>

#### `healthcheck.name` property

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

<a id="canonical-2132130030010021-2022302112231131-1233200313111231-3221201011001111-0220233100220123-2033332301330200-2313220132112001-0122122033122303"></a>

<a id="canonical-3311003322132300-1321110032320310-2230220323223210-3101221101103121-0323100002313002-1023323231202102-3013202111010130-0330221111201003"></a>

#### `healthcheck.namespace` property

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

<a id="canonical-3333122232012021-0202223110210203-1313302332213200-1122211113331230-3221012013200330-0102023300002010-3111010311122323-1033012331310313"></a>

<a id="canonical-0031301200230233-1311102332013221-0111232230033111-1210012030120010-0311110122211203-1212320213122020-1311003021210021-2222022110311213"></a>

#### `healthcheck.tenant` property

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

<a id="canonical-2321331312131123-3012123331102100-2201101123302212-1013031113132232-0023332220033133-3221300001120322-1131200220002203-2121000221033310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `lb_port` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- lb_port

<a id="canonical-3332320330310202-3001331103323210-0100103311013120-0002100310001103-2000333310000320-3010022210312130-1113032332102000-2210230130200022"></a>

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

<a id="canonical-3003002212002303-1310102203201220-1232202311101013-2320113213310102-0313212211221112-0022010223033133-1302122300202210-0233020303132231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_tls` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- no_tls

<a id="canonical-3003013113222220-2211123101222221-3301200210020030-1013302300212120-2102113320000233-3323230131232012-1312300231010103-1233101032222222"></a>

Type: `["object", {}]`. Computed.

\[OneOf: no\_tls, use\_tls; Default: no\_tls\] Enable this option. Defaults to \`map\[\]\`. Server
applies default when omitted.

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

- [no_tls](data-sources--origin_pool--reference--group-002.md#canonical-3003013113222220-2211123101222221-3301200210020030-1013302300212120-2102113320000233-3323230131232012-1312300231010103-1233101032222222)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-0002222103101132-0331331011030000-0221212201230213-1311030121013132-0223020023313332-2311112031102200-3132032023331103-2122132001320021)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- origin_servers

<a id="canonical-1202000233133133-3013023232001221-0213032112311332-2212331201233022-0011000211132213-3032110030112230-2120311230221030-2001112121001331"></a>

Type: `"list"`. Computed.

Origin Servers. List of origin servers in this pool.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0003112032023133-2320302323030301-2130300123231223-2011322102330210-1031030120021030-2202230111231303-3203011032011222-1021233003022232"></a>

### Direct properties for `origin_servers`

- [cbip_service](data-sources--origin_pool--reference--group-002.md#canonical-1131022223301121-0113110032012312-3110321130113110-1311223333131310-0030021302100133-1223302232123022-1321310233220231-3103231012021331): complete subsection reference.

- [consul_service](data-sources--origin_pool--reference--group-002.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101): complete subsection reference.

- [custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-3123002010230120-1211300002023332-3220012210300131-1320202221223001-3220302001323203-2022021333133032-1231223202322213-0303321213212231): complete subsection reference.

- [k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201): complete subsection reference.

<a id="canonical-1102110302003202-2233112110210002-0203310333030000-1232303012101122-3200332233322221-1200132132012213-0003110321322113-3113303101332101"></a>

<a id="canonical-1002320313212030-0221031330201123-0033032200200210-2333000303230110-2020030312303023-3103112332202000-2011302021123321-1202120120330302"></a>

#### `origin_servers.labels` property

Type: `["map", "string"]`. Computed.

Add Labels for this origin server, these labels can be used to form subset.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102): complete subsection reference.

- [private_name](data-sources--origin_pool--reference--group-003.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110): complete subsection reference.

- [public_ip](data-sources--origin_pool--reference--group-003.md#canonical-3233223332020121-3100021120202320-2031230003101211-3023100123103102-0220333130210323-2201111232211122-1033211313103113-3001033210113321): complete subsection reference.

- [public_name](data-sources--origin_pool--reference--group-003.md#canonical-3023121012300320-1113333033031232-1003130321112112-1203102212132000-1310200323133230-2031320320302202-3121312332032031-2323130012131111): complete subsection reference.

- [vn_private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1132312333312222-3332122302332200-2310202220131303-2312323300201122-2030322023323212-2010332131123003-0331231022031330-0313233113021202): complete subsection reference.

- [vn_private_name](data-sources--origin_pool--reference--group-003.md#canonical-3322112003203211-1300331030110100-3211113003210110-3003332110030121-2200102131333003-3323033022212213-0033103302220322-0021203022320330): complete subsection reference.

<a id="canonical-1131022223301121-0113110032012312-3110321130113110-1311223333131310-0030021302100133-1223302232123022-1321310233220231-3103231012021331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.cbip_service` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.cbip_service

<a id="canonical-0233032331300323-3222131020331213-1312130022210001-2212330232123302-1033120331010313-2201020010203311-0001110000200022-2101001211221212"></a>

Type: `"single"`. Computed.

Specify origin server with Classic BIG-IP Service (Virtual Server).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1021313333031100-3002312210320131-2000131002002311-1333031301332031-0213320300332322-1333022022123303-1033130320003022-0011222023100202"></a>

### Direct properties for `origin_servers.cbip_service`

<a id="canonical-3333012200100220-3211011203300311-0023133011123002-3000100012303212-1122110200000110-2330203112020310-2021320000210013-1232111223310102"></a>

#### `origin_servers.cbip_service.service_name` property

Type: `"string"`. Computed.

Name of the discovered Classic BIG-IP virtual server to be used as origin.

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

<a id="canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.consul_service` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.consul_service

<a id="canonical-0210310100023010-3001201121100323-2233111213111200-1110031130123110-1013203030001320-0312331000200223-2023121102221301-0122302120030023"></a>

Type: `"single"`. Computed.

Specify origin server with HashiCorp Consul service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]"
}
```

<a id="canonical-0012303011302210-3121102302031133-2323130310031201-0002001220032331-1322033331232203-0030330212011333-1021113122301301-1301130033201032"></a>

### Direct properties for `origin_servers.consul_service`

- [inside_network](data-sources--origin_pool--reference--group-002.md#canonical-1311221233101121-0021312202000313-3231111332102020-0021232221101013-3101223120103103-3130111102132330-3132232230311213-0110113113121122): complete subsection reference.

- [outside_network](data-sources--origin_pool--reference--group-002.md#canonical-1130022023121110-3213212022312101-3333221030031123-1321022032300221-0231103100232131-0200212110210030-1202311120300013-1232312120303113): complete subsection reference.

<a id="canonical-1300121122331122-1123001012133210-2000202031321320-1323001021210011-1211001310113113-1230021010000203-2312313322320211-1112211122322231"></a>

<a id="canonical-1102021123131220-3303000221322011-0233210020030231-3013003221001022-1003013013311120-2001200312022202-1110233330201303-0203131311131310"></a>

#### `origin_servers.consul_service.service_name` property

Type: `"string"`. Computed.

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

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

- [site_locator](data-sources--origin_pool--reference--group-002.md#canonical-0220232232031101-0021332122202322-0000010021303022-3120030132133221-3332012010321202-3213003213331110-3022311110231022-0030230300222112): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2320132102213312-1002323332032110-2303031210111010-2303200111311303-3103313332001302-2113122313133300-0233103020312222-0232102301003021): complete subsection reference.

<a id="canonical-1311221233101121-0021312202000313-3231111332102020-0021232221101013-3101223120103103-3130111102132330-3132232230311213-0110113113121122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.consul_service.inside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-002.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- origin_servers.consul_service.inside_network

<a id="canonical-1320231211131113-0312330100313120-0000133233102123-0232312210131300-1202121323203102-3332212312233131-0332223332023321-0113223122313111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-1130022023121110-3213212022312101-3333221030031123-1321022032300221-0231103100232131-0200212110210030-1202311120300013-1232312120303113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.consul_service.outside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-002.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- origin_servers.consul_service.outside_network

<a id="canonical-2001322200002122-1231301123332322-0300121011121113-2111202032020310-2100131213210122-0221000112331123-3302031003120303-1122102210222113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-0220232232031101-0021332122202322-0000010021303022-3120030132133221-3332012010321202-3213003213331110-3022311110231022-0030230300222112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.consul_service.site_locator` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-002.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- origin_servers.consul_service.site_locator

<a id="canonical-2201002203020211-2222200021030213-2303312333023310-1303230031112302-2123323031022112-3020032000231233-0020323030031323-2211133311111301"></a>

Type: `"single"`. Computed.

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-0121201133331112-3010203330030000-0111320022330120-0122233020200223-0211310210123031-1010303112201123-2333223130302201-2230332022300232"></a>

### Direct properties for `origin_servers.consul_service.site_locator`

- [site](data-sources--origin_pool--reference--group-002.md#canonical-3201110303231310-3011013210011132-1203330220301213-3020102100210210-1100000013303333-0323223030101023-3101331232320122-2200012221211323): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-0011311000232122-2211001023101233-2330112333230321-1011003130303122-0023333100100231-1121320123123211-2331100112332302-0232012030310231): complete subsection reference.

<a id="canonical-3201110303231310-3011013210011132-1203330220301213-3020102100210210-1100000013303333-0323223030101023-3101331232320122-2200012221211323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.consul_service.site_locator.site` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-002.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-0220232232031101-0021332122202322-0000010021303022-3120030132133221-3332012010321202-3213003213331110-3022311110231022-0030230300222112)
- origin_servers.consul_service.site_locator.site

<a id="canonical-1313003323320102-2303213312302230-3302110132132131-1110022102103022-1203310010303222-3230231330220332-3330203002223202-0232011311001021"></a>

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

<a id="canonical-3320201311133231-1110101302201221-1222302013300123-0301031033230232-1312103102203302-0001003312013212-0223013312122031-2311220130011213"></a>

### Direct properties for `origin_servers.consul_service.site_locator.site`

<a id="canonical-3000122233210221-0013332312311303-2001121210122023-1230331330203320-0111323012222032-3110013000332120-0103221311011303-2032122120322030"></a>

#### `origin_servers.consul_service.site_locator.site.name` property

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

<a id="canonical-0003010221112231-3301122210021031-2311333302322103-0223033330010021-3021221110202200-1101302222200231-0033200320333131-2011223230231030"></a>

<a id="canonical-1121303222103212-0000133001110231-3221130133212111-2213230003300311-0301303303223001-0103302001202010-3323101320331021-2011211121221020"></a>

#### `origin_servers.consul_service.site_locator.site.namespace` property

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

<a id="canonical-3211212202030121-3200311312212202-3232102102133030-3003213201311312-2301321112113100-2110002312012113-2003111200233312-0013002212013212"></a>

<a id="canonical-2310010130312321-1113210221113100-0333012230221103-2111120133310002-0211012233223030-1032231022030032-3312011210221320-3211000022011303"></a>

#### `origin_servers.consul_service.site_locator.site.tenant` property

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

<a id="canonical-0011311000232122-2211001023101233-2330112333230321-1011003130303122-0023333100100231-1121320123123211-2331100112332302-0232012030310231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.consul_service.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-002.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-0220232232031101-0021332122202322-0000010021303022-3120030132133221-3332012010321202-3213003213331110-3022311110231022-0030230300222112)
- origin_servers.consul_service.site_locator.virtual_site

<a id="canonical-0020302332213002-0211232210010332-0210230023321103-3012221020333233-3021321021313323-2331131012123100-0111302121230001-2322121202303023"></a>

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

<a id="canonical-1013322200102312-2210300131110031-2001113022223223-0002011332022031-0311200311231000-3200333130110103-3310102021032112-1001011111101330"></a>

### Direct properties for `origin_servers.consul_service.site_locator.virtual_site`

<a id="canonical-3311231130121310-1113010133301031-3110023330122112-0011201101203001-2212121300210130-1022032332122003-1223013001021303-0331233101011330"></a>

#### `origin_servers.consul_service.site_locator.virtual_site.name` property

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

<a id="canonical-0323311030133302-1020031231202313-2300133333321203-3012101332012111-0101001103021222-3302310211232223-0302200110133302-3012021120303010"></a>

<a id="canonical-2231120302302331-2232311200213102-1000202212302232-1013133020303201-3223031203133303-2003230132012031-2311113301110323-3310302000013220"></a>

#### `origin_servers.consul_service.site_locator.virtual_site.namespace` property

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

<a id="canonical-2031011002312332-0112122212112020-2021313002121220-1231113333220302-3002100213102213-2023133130021302-1113131023321112-2031000313113103"></a>

<a id="canonical-1133331023010111-2130002101010231-1012223312102212-0220122320010332-2131310003133213-2331033002023023-2120020310031211-2333120021203000"></a>

#### `origin_servers.consul_service.site_locator.virtual_site.tenant` property

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

<a id="canonical-2320132102213312-1002323332032110-2303031210111010-2303200111311303-3103313332001302-2113122313133300-0233103020312222-0232102301003021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.consul_service.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-002.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- origin_servers.consul_service.snat_pool

<a id="canonical-2002312100030302-3002022232311220-3301132331131212-0010310320111231-1113213010131230-3023023210130222-3222233301023132-3203022211022202"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-3201010013122030-3111333111321221-2132123221231311-1032303203023321-2002231002230112-0330320310111111-3033033023102110-1013010101323201"></a>

### Direct properties for `origin_servers.consul_service.snat_pool`

- [no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2212302023010121-3012311222032022-2331132001202020-1012020032233011-2122122310311010-3312003010312122-1030011201211321-1012203301322300): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1332111321010200-3023123230030001-0220322303300123-2032020220130332-2332222111130213-3203221213003320-0332211023323203-1220030013121033): complete subsection reference.

<a id="canonical-2212302023010121-3012311222032022-2331132001202020-1012020032233011-2122122310311010-3312003010312122-1030011201211321-1012203301322300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.consul_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-002.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2320132102213312-1002323332032110-2303031210111010-2303200111311303-3103313332001302-2113122313133300-0233103020312222-0232102301003021)
- origin_servers.consul_service.snat_pool.no_snat_pool

<a id="canonical-3202201023102303-2012010023330120-3312200223102300-3321033012123312-1000320132320311-2203022332211232-3230313130120111-3031211232132110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-1332111321010200-3023123230030001-0220322303300123-2032020220130332-2332222111130213-3203221213003320-0332211023323203-1220030013121033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.consul_service.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-002.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2320132102213312-1002323332032110-2303031210111010-2303200111311303-3103313332001302-2113122313133300-0233103020312222-0232102301003021)
- origin_servers.consul_service.snat_pool.snat_pool

<a id="canonical-3022220313021103-1130003021313320-1220333003113330-0332112323300013-1321022310011020-0200012202331312-2003202332120302-3203310202303332"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1310312232011011-1113113233120002-0213222132233300-0330333111331322-3122110213210300-0102113002313122-0222311200211320-1330031223100032"></a>

### Direct properties for `origin_servers.consul_service.snat_pool.snat_pool`

<a id="canonical-0210132323210231-2130202320012330-1103212300013211-1303200213300003-2011130211020101-2113031321300113-2113133111232103-3113201213020222"></a>

#### `origin_servers.consul_service.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3123002010230120-1211300002023332-3220012210300131-1320202221223001-3220302001323203-2022021333133032-1231223202322213-0303321213212231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.custom_endpoint_object` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.custom_endpoint_object

<a id="canonical-3333021302313233-2132310320013112-0113132311331331-2123210023300020-2030301130113303-3310200033000000-3230303211313120-3220310031310203"></a>

Type: `"single"`. Computed.

Specify origin server with a reference to endpoint object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3213212110323120-0303110023333310-0133030300113132-0303321031011330-3131130300031213-1312212202232120-2220300221031022-3113121101122202"></a>

### Direct properties for `origin_servers.custom_endpoint_object`

- [endpoint](data-sources--origin_pool--reference--group-002.md#canonical-2200222130022302-2303131210112110-0123312133022300-3033321312313320-1231231221303333-1111222311213303-1222113011223022-2131203212012131): complete subsection reference.

<a id="canonical-2200222130022302-2303131210112110-0123312133022300-3033321312313320-1231231221303333-1111222311213303-1222113011223022-2131203212012131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.custom_endpoint_object.endpoint` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-3123002010230120-1211300002023332-3220012210300131-1320202221223001-3220302001323203-2022021333133032-1231223202322213-0303321213212231)
- origin_servers.custom_endpoint_object.endpoint

<a id="canonical-1213110323211231-0333323132100100-0111033023230312-2120123301202333-1032110200313211-0223020233100211-2322033020203321-3333100211133201"></a>

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

<a id="canonical-2322111001211120-0011323303313313-0222031221011132-1130302233301131-1102113322102103-0332031110122221-2102101121221011-2123300332032032"></a>

### Direct properties for `origin_servers.custom_endpoint_object.endpoint`

<a id="canonical-3131322332321103-3233133102210301-3233120011202133-0002123111003220-2322230232312331-1201201301020111-0130232021332013-3232031000112323"></a>

#### `origin_servers.custom_endpoint_object.endpoint.name` property

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

<a id="canonical-0320122000303100-2211122300113100-1231000112232003-1100111233232010-0300033333232333-3120330230313123-2223100212332311-3320221232202122"></a>

<a id="canonical-1320233120322331-3002120100322131-1233011222220203-1332012223312202-0301032121310013-0122021312310101-0202022313002130-1212101130112321"></a>

#### `origin_servers.custom_endpoint_object.endpoint.namespace` property

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

<a id="canonical-3133030202000321-3301010003320331-3110322022313022-2001231320013112-2132002020322202-2301020022310311-3003201201021101-0311020122320122"></a>

<a id="canonical-2000122223033220-1110012203022121-1103322303302223-2103113222200112-1212100323211302-3123230321013202-3120120002000033-0032330132331120"></a>

#### `origin_servers.custom_endpoint_object.endpoint.tenant` property

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

<a id="canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.k8s_service

<a id="canonical-1223233020022101-3220211210102003-1322331302233310-3210312321210332-2022313212330223-3032211030303222-1023120103333200-1113332132013312"></a>

Type: `"single"`. Computed.

Specify origin server with K8s service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

<a id="canonical-3202032000113203-3311021032200010-0020022020201310-2011320211120033-2131122131103320-0021113100001111-2220103333122112-2031032110000133"></a>

### Direct properties for `origin_servers.k8s_service`

- [inside_network](data-sources--origin_pool--reference--group-002.md#canonical-2121000330003222-2030211333231213-1131222130230302-3113030331232101-3122113231312313-2202202111213213-1222131313231130-3111112122202202): complete subsection reference.

- [outside_network](data-sources--origin_pool--reference--group-002.md#canonical-2010233310313303-2232131212223021-1021133333220302-0031003331120131-1103201003023210-3231233320021013-3312300231122222-3032320221011203): complete subsection reference.

<a id="canonical-2033110102321231-2011330302322202-1233110020223330-2221122203002103-0300002033001031-3320102302330233-0332002101233021-3003323032022011"></a>

<a id="canonical-1302323302021001-2133133023303320-2210100303100202-3221203233122010-0330301210213133-2222112320121300-2330333103203122-2101221123330230"></a>

#### `origin_servers.k8s_service.protocol` property

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2203033320222001-2103023111123222-2123233202333022-1302301030303113-0032310111031133-3231203233112102-2313112330103011-0010231113212013"></a>

<a id="canonical-2123002203100223-1130133222120033-0102120331033003-0022103133311110-1123223030112113-1220331031013102-1010200213201202-2000133013201021"></a>

#### `origin_servers.k8s_service.service_name` property

Type: `"string"`. Computed.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Additional upstream details:

For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](data-sources--origin_pool--reference--group-002.md#canonical-3213302323123220-2200323333301123-0321023122332230-0200100310110120-2302312313312031-3230112131212203-3301021011100222-3000003312200123): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-1003201010222201-1101110320120012-1210031020121301-2311123310120002-3313110220322012-2200010133303101-3120221211111312-0310000310102010): complete subsection reference.

- [vk8s_networks](data-sources--origin_pool--reference--group-003.md#canonical-3003320220101232-2222122002000011-0122233212223100-0113213320301211-3321322300223222-1233131112321100-3101110202223112-2221113000223003): complete subsection reference.

<a id="canonical-2121000330003222-2030211333231213-1131222130230302-3113030331232101-3122113231312313-2202202111213213-1222131313231130-3111112122202202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.inside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- origin_servers.k8s_service.inside_network

<a id="canonical-2203311331100022-3222230321303230-2120012031212131-1200031230202303-0301032213011323-0233003220100333-0111101122312331-1300222230302132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-2010233310313303-2232131212223021-1021133333220302-0031003331120131-1103201003023210-3231233320021013-3312300231122222-3032320221011203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.outside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- origin_servers.k8s_service.outside_network

<a id="canonical-3012021223332031-2331031010321322-1321023001130222-1110130231311222-1030302202321020-0010220322233210-0200002000311313-0030121002132231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-3213302323123220-2200323333301123-0321023122332230-0200100310110120-2302312313312031-3230112131212203-3301021011100222-3000003312200123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.site_locator` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- origin_servers.k8s_service.site_locator

<a id="canonical-0022112310330022-0110101102111300-2022230203223313-1021300231102100-1320312212011221-0220311212012120-0021321021212233-1122023130231201"></a>

Type: `"single"`. Computed.

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-0232212300132300-1000231001231303-1030131000332223-0303220212302030-3013212120102020-1131020323333130-0322033233331321-3233301332122010"></a>

### Direct properties for `origin_servers.k8s_service.site_locator`

- [site](data-sources--origin_pool--reference--group-002.md#canonical-0023122133213320-0032120210322223-1233002230031131-1320321120020013-2032223202303301-1220121113011213-2120311220321203-3212310020202210): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-3032123210002103-2121113021330010-0033023103002023-0230300130130333-2031031021203312-0133011211103322-3002120330033330-3132122121131022): complete subsection reference.

<a id="canonical-0023122133213320-0032120210322223-1233002230031131-1320321120020013-2032223202303301-1220121113011213-2120311220321203-3212310020202210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.site_locator.site` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-3213302323123220-2200323333301123-0321023122332230-0200100310110120-2302312313312031-3230112131212203-3301021011100222-3000003312200123)
- origin_servers.k8s_service.site_locator.site

<a id="canonical-1132230003030301-1031133202102201-3111000112233130-2111001033320033-0323131211130233-2231310022112003-0220312223332322-0203131100022311"></a>

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

<a id="canonical-1300101013110220-0322232031130322-1002321021220202-1320033000102301-0201323302103002-1231113300310102-0003030331230023-2300123321000130"></a>

### Direct properties for `origin_servers.k8s_service.site_locator.site`

<a id="canonical-2020201212221111-3011330313010203-2021202000321221-1103133213210120-1123021302330132-0030122133300010-1112113130033231-0132212231001120"></a>

#### `origin_servers.k8s_service.site_locator.site.name` property

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

<a id="canonical-0202233210233213-0013230131322323-0121011223032222-3012120221000302-0120002202022021-0130001311111213-0123123020123322-1200003023220321"></a>

<a id="canonical-2231022101211132-1131020311231020-1312130302132301-3032201202123212-0230323001303032-0013013011003111-2232132231210102-0110121130031313"></a>

#### `origin_servers.k8s_service.site_locator.site.namespace` property

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

<a id="canonical-2222012313310210-0022032013100010-1211230330120201-1111302301322301-3113122201331221-3001023023003210-1031121200020003-1131223333333211"></a>

<a id="canonical-1321211131313112-0323120320231000-2133302231131300-2131333133311021-1102300310131220-1331232322102002-3103313333210022-3023030010303203"></a>

#### `origin_servers.k8s_service.site_locator.site.tenant` property

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

<a id="canonical-3032123210002103-2121113021330010-0033023103002023-0230300130130333-2031031021203312-0133011211103322-3002120330033330-3132122121131022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-3213302323123220-2200323333301123-0321023122332230-0200100310110120-2302312313312031-3230112131212203-3301021011100222-3000003312200123)
- origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-0310032322220213-0023320311333122-3300011313031132-3332003333030021-3202203330303110-1320232211123233-2113232133332122-2302112132103200"></a>

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

<a id="canonical-3232100330223100-0321100301330200-1030211312313011-3123203032003100-2212211230302103-1202320003131301-2012132213332133-2100031011333130"></a>

### Direct properties for `origin_servers.k8s_service.site_locator.virtual_site`

<a id="canonical-3121121032012310-3323221103112303-3033102120210112-1312011301331103-1011030030011113-3313210023300001-3311121031021013-3333031222000131"></a>

#### `origin_servers.k8s_service.site_locator.virtual_site.name` property

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

<a id="canonical-3100000122100212-0311032300110231-2102310101333300-3101201323130231-2330121223111103-2020301113203031-1122031231022020-1213022011303232"></a>
