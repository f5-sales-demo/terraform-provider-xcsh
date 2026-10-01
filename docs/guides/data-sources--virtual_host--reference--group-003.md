---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-0212013011210010-0001123033110331-3200022302101133-2002200133001302-3211010333321310-0121111001213223-2202021330022033-1233211111211012"></a>

## location property — blindfold_secret_info / 100101203321 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1102310020100012-3102100223203311-3110012320302010-0123100022202212-0033032111133122-0103313212300302-3202100320123303-2101331100023100"></a>

<a id="canonical-0121120322202303-1300022121232010-3123002302101100-0233120121212231-3232310331230220-1233323132302230-3323200012231220-1031330123112112"></a>

## store_provider property — blindfold_secret_info / 100101203321 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3032132213202133-1213012320030013-0032313123220222-1112210131023131-2210020101212211-3112020300032311-0322230102123101-1310013201213320"></a>

## Next pages — blindfold_secret_info / 100101203321 / 7

- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3101230321222032-3323212212032112-0000210100132300-2200202302113313-3202033223320201-3113222013230022-1120310312103300-3330002320302120)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0310123303110031-0303203333000323-2223131103332331-0200031312112020-1120223122110221-2312231113313233-0010021002321202-0100213233303303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030232320213122-3023111310003102-1131202220220003-2203301320031312-2221310213132110-0232030303023133-3101131002230231-0020213111133030"></a>

## response_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 223113322122 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [response_headers_to_add](data-sources--virtual_host--reference--group-002.md#canonical-1122301030300332-0033230331023011-1102233120003323-3100320132003322-0030000111220313-3302203220201033-0222213312322000-2123201002112113)
- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3101230321222032-3323212212032112-0000210100132300-2200202302113313-3202033223320201-3113222013230022-1120310312103300-3330002320302120)
- response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3130022021012020-2111013303010033-1333320203310213-1113300030300221-3010301303313032-0131000023311021-1311011231100022-0033220130331233"></a>

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

<a id="canonical-2021030021230321-3130013031300223-0112021201221331-0220011221032000-3032120322201103-2001122112012212-1032030132102101-3030021321112302"></a>

## Direct properties — clear_secret_info / 223113322122 / 3

<a id="canonical-3130311230122321-1302312201303300-2201333323000103-2221203213132300-1221320233200232-0202033103310201-1130033232123211-1020312201022013"></a>

<a id="canonical-3310322220221322-1312323223333031-1112313223232320-0111301321233011-2320113210003201-1110132333300213-3310023232113220-3222303120033330"></a>

## provider_ref property — clear_secret_info / 223113322122 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3021103000001322-1332201012301002-2323332233021010-3223020132221322-3321333002212312-0002112103120230-0002122102112030-2000112103213322"></a>

<a id="canonical-0210331303113301-0101233323031100-2212022233221201-3033120322022330-3012033211203113-0200001103012120-2321130232133002-2202123120330122"></a>

## URL property — clear_secret_info / 223113322122 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2122311310113001-1001233102301032-2320113002213003-3211032311221033-0321121222212003-2232223013031020-0311330332100231-1123320232231332"></a>

## Next pages — clear_secret_info / 223113322122 / 6

- [response_headers_to_add.secret_value](data-sources--virtual_host--reference--group-002.md#canonical-3101230321222032-3323212212032112-0000210100132300-2200202302113313-3202033223320201-3113222013230022-1120310312103300-3330002320302120)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1021031030330102-2100323023122311-3003021113211121-2312223333303303-2002222022211201-0131100013001013-2031022313222112-2203022023311021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310211322221223-3233311103233230-0132203200100030-3210121121113222-0312211323201121-1211031030203100-1033001331001123-3113303302011023"></a>

## retry_policy — retry_policy / 131033333332 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- retry_policy

<a id="canonical-3121300022323212-3032100030021200-3220123111011103-0022232031022323-2212331012001311-2303233320201201-1331103322001123-2010123121001101"></a>

Type: `"single"`. Computed.

Retry policy configuration for route destination.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3032221232321322-1122001213223321-0311123201322113-0030112003300333-1123122000203312-2010230121200323-0003311030323222-1312021220230301"></a>

## Direct properties — retry_policy / 131033333332 / 3

- [back_off](data-sources--virtual_host--reference--group-003.md#canonical-3213001230013032-0020221021322333-1332330301110002-0020303010112211-3321113230333322-3203333212321332-0203333021032210-3112130033232321): complete subsection reference.

<a id="canonical-1220312222312301-1032103101122112-0131120022103202-3313321111013110-0222113311201112-0130211121333200-1110212201330031-1102103301033132"></a>

<a id="canonical-2020303221110022-1112213013330221-0200330130112102-3211121313101031-3013020202132032-0100030023321020-0003112302210300-0121021311021213"></a>

## num_retries property — retry_policy / 131033333332 / 4

Type: `"number"`. Computed.

Specifies the allowed number of retries. Retries can be done any number of times. An exponential
back-off algorithm is used between each retry. Defaults to \`1\`.

Upstream description:

Specifies the allowed number of retries. Defaults to 1. Retries can be done any number of times. An
exponential back-off algorithm is used between each retry.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

<a id="canonical-0320211020202131-1110221323012302-0333100001031313-0322023302330120-0201222331000302-2030031122110001-1330003133012202-3121210000302323"></a>

<a id="canonical-3221231302313311-1320131002022202-0221212233020322-1023231103131333-2023223223003032-0011031231030132-3311022120111301-3200213230222323"></a>

## per_try_timeout property — retry_policy / 131033333332 / 5

Type: `"number"`. Computed.

Specifies a non-zero timeout per retry attempt. In milliseconds.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
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

<a id="canonical-2121022023022001-3000132120321230-2023030231120210-3023100021111221-1021123323210330-2002210323013301-3110301013223023-0133022312311100"></a>

<a id="canonical-0223322130220012-2202310323330121-3203302320001302-0111200020310102-3231113203333213-2013331031031103-3101111003032323-0232110210200220"></a>

## retriable_status_codes property — retry_policy / 131033333332 / 6

Type: `["list", "number"]`. Computed.

HTTP status codes that should trigger a retry in addition to those specified by retry\_on.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1330131133121001-0213001123313202-0333010330032300-2200023131011201-3032233303321102-0322311222232033-3011202221002323-1111222331132312"></a>

<a id="canonical-3020203321332320-0222230202102023-2313103112220100-0013310210131112-3001332231100001-0132101303131222-3321013323003230-1332132131001332"></a>

## retry_condition property — retry_policy / 131033333332 / 7

Type: `["list", "string"]`. Computed.

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if
the..

Upstream description:

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc

The possible values are

"5xx" : Retry will be done if the upstream server responds with any 5xx response code, or does not
respond at all (disconnect/reset/read timeout).

"gateway-error" : Retry will be done only if the upstream server responds with 502, 503 or 504
responses (Included in 5xx)

"connect-failure" : Retry will be done if the request fails because of a connection failure to the
upstream server (connect timeout, etc.). (Included in 5xx)

"refused-stream" : Retry is done if the upstream server resets the stream with a REFUSED\_STREAM
error code (Included in 5xx)

"retriable-4xx" : Retry is done if the upstream server responds with a retriable 4xx response code.
The only response code in this category is HTTP CONFLICT (409)

"retriable-status-codes" : Retry is done if the upstream server responds with any response code
matching one defined in retriable\_status\_codes field

"reset" : Retry is done if the upstream server does not respond at all (disconnect/reset/read
timeout.)

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 7,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333030010100333-2021211030131233-3323330130123213-0331323301123132-0022122212011312-1213023311133201-2230203333231032-1013211233123122"></a>

## Next pages — retry_policy / 131033333332 / 8

- [retry_policy.back_off](data-sources--virtual_host--reference--group-003.md#canonical-3213001230013032-0020221021322333-1332330301110002-0020303010112211-3321113230333322-3203333212321332-0203333021032210-3112130033232321)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3213001230013032-0020221021322333-1332330301110002-0020303010112211-3321113230333322-3203333212321332-0203333021032210-3112130033232321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000110230123330-1330333310310233-3312013201322233-2002030231320211-1123312310300003-1023320323223102-0303032322202000-3030311033223201"></a>

## retry_policy.back_off — back_off / 003033100003 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [retry_policy](data-sources--virtual_host--reference--group-003.md#canonical-1021031030330102-2100323023122311-3003021113211121-2312223333303303-2002222022211201-0131100013001013-2031022313222112-2203022023311021)
- retry_policy.back_off

<a id="canonical-2100010313230133-0131312130011120-3021201113200123-0111123210210102-2313130302102213-3321231130312130-2130210232001211-0132213311231233"></a>

Type: `"single"`. Computed.

Specifies parameters that control retry back off.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301320000002222-1232201300111020-3322113010030122-3320312110330131-1201011112233322-3002312003213322-2312302101212033-2311010100121120"></a>

## Direct properties — back_off / 003033100003 / 3

<a id="canonical-2203132023322013-0121130211022203-3232303101100320-3232300022223013-3121331123220320-2021300332201203-3201002020331213-1130003023311201"></a>

<a id="canonical-3311110331000202-1133231222210002-0121320301321203-3220332310122001-0310101023013013-2330201320133013-2131021132100221-3000333100132202"></a>

## base_interval property — back_off / 003033100003 / 4

Type: `"number"`. Computed.

Specifies the base interval between retries in milliseconds.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-2230122033122210-3301212221320032-2101022132012321-0313311303320300-2030013311022231-1011313330031023-0122310203312220-2101300302001102"></a>

<a id="canonical-2000201001020323-3332213132333311-2010301333201312-1000013221311203-3133200200202110-1011230220311131-2300302223233020-2021222013100311"></a>

## max_interval property — back_off / 003033100003 / 5

Type: `"number"`. Computed.

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The times the base\_interval. Defaults to
\`10\`.

Upstream description:

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The default is 10 times the base\_interval.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2322103202113000-0012311210202011-2022023022310312-0213223120333031-0003021031313131-2231022322121232-0133121111120313-2130111123210310"></a>

## Next pages — back_off / 003033100003 / 6

- [retry_policy](data-sources--virtual_host--reference--group-003.md#canonical-1021031030330102-2100323023122311-3003021113211121-2312223333303303-2002222022211201-0131100013001013-2031022313222112-2203022023311021)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3211011200123230-0330211200101113-1112022302023320-3120120232213300-3223122013100230-0003102201231003-1303322122132321-0212222131122110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230102101023212-3022333312221101-0030300303202311-2020110231221212-3113313303121221-1033112321330123-0133002210121201-2030131231330023"></a>

## routes — routes / 132113032033 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- routes

<a id="canonical-2002112323120121-1202312220122120-3322103020012232-0302122212212332-0222332110320333-1213010203011123-0000002223313331-0233310011312332"></a>

Type: `"list"`. Computed.

HTTP routing rules that match incoming requests based on path, headers, or query parameters and
forward them to appropriate backend origin pools.

Upstream description:

The list of routes that will be matched, in order, for incoming requests. The first route that
matches will be used. Currently route object is redundant in case of TCP proxy but required. For
TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts, the route object only specifies the
cluster/weighted-cluster as route destination without any match condition. In other words, match
condition in route object is ignored for TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts.
Routes used for TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts cannot have DirectResponse
or Redirect as actions.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-3203322102132213-1111322321310301-1100311202212200-1002331001023332-3100222113330012-1223221103301122-1022131022223002-2222130230202232"></a>

## Direct properties — routes / 132113032033 / 3

<a id="canonical-1223020302332210-2023312112223323-0221101131313303-3023200210310011-2031033021232212-1012021110032002-0130032212230020-3222220230002003"></a>

<a id="canonical-3322201321303032-1311131202310111-3322323222213033-0311210320303203-1211123332131203-1002020231320033-3031313232223331-2003200310131023"></a>

## kind property — routes / 132113032033 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2202300320101110-0113011022002232-2110123233001222-1131030303001313-3232302022013011-3223233110113210-2012021133120233-1111023032031320"></a>

<a id="canonical-1303313322012321-1013332220121211-0331310220223333-2333203303123000-2012031200222000-1121120013123030-1003101330312133-3121210013110300"></a>

## name property — routes / 132113032033 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0210112210122000-0322130213000103-2333231131123030-2030311131121213-3031222313311233-3123322202222310-2211113012232110-1123030213303122"></a>

<a id="canonical-3030320121131210-3312333323200112-3321322001233210-2203302323001123-2223011031013022-2013211310033033-2322123001111122-3030103312012213"></a>

## namespace property — routes / 132113032033 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3320303210103110-2113232323122122-2230323000030202-1112313303333200-1000223103110013-1213132102221031-2133102330130322-1031112033033212"></a>

<a id="canonical-3200220322230232-1201231021321113-2033032301121112-0203302133033122-2312001001133121-1213232222030110-1311232302310212-2120132203013300"></a>

## tenant property — routes / 132113032033 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3123331032100112-0230032013120120-0232012012302122-3120122102110232-3332303313212120-2022020010212033-1322322303230002-2123032322023220"></a>

<a id="canonical-1012013022020233-1321110011101213-2333312230331202-2000233303113303-1010113203131322-1201313030222322-1011323003302013-0120100122121031"></a>

## uid property — routes / 132113032033 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3303122001320221-2011020232021120-1002032210323111-2222011123313011-2033111222310113-2301031121133123-1202122300213103-1221300002312220"></a>

## Next pages — routes / 132113032033 / 9

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1111130000112110-1102321322100210-0013310223010100-3213130303222101-3131013201101301-1000021000133131-0023212331022302-0120103003321211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222233231020111-2022011320010212-1012001323201000-1132330011301331-3220031211121323-0220100011220012-1233320031013023-1123123312201310"></a>

## sensitive_data_policy — sensitive_data_policy / 033221102122 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- sensitive_data_policy

<a id="canonical-2113012213211213-0001300300111232-0000020003302202-0323031131213222-2010031320310213-2023110002103000-0233002333102213-0010303110201000"></a>

Type: `"list"`. Computed.

Policy configuration for this feature.

Upstream description:

References to sensitive\_data\_policy objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3221212123013022-2100211010103012-3013312231030201-2331233022013101-3020100323032100-3102102001212230-1321202101333121-3322311202030303"></a>

## Direct properties — sensitive_data_policy / 033221102122 / 3

<a id="canonical-1131233313023311-2110020033320000-0020311223322131-0211222333323230-0030120022330022-1210321330303210-2322023311132123-1021330022311000"></a>

<a id="canonical-1233020121002023-0221332232312312-2022010333031333-3331113332133113-1101233021323020-0332021320201223-3333111103133120-3120221112101122"></a>

## kind property — sensitive_data_policy / 033221102122 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3323132320003002-0020103200100221-1223211013120003-0123101131101131-0102111300011232-1020000320333220-0322022111123232-0233331300323131"></a>

<a id="canonical-2233032231023330-2110200032210123-3022301100230030-0000211321121332-0002211013203330-0322223203111201-3023211333000102-3032032323230011"></a>

## name property — sensitive_data_policy / 033221102122 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1230101132333321-0231110301231201-3311323123113322-0231331010203132-3310020311101021-2013203010312033-0300220112033322-2000021123301300"></a>

<a id="canonical-3222200200111320-3022111311220023-3333333333233111-3210001203101130-0020213323203133-1330323211301320-2013231132111132-0000103032110021"></a>

## namespace property — sensitive_data_policy / 033221102122 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1023212213333333-2233313120032322-2311231322131322-2113102110330011-1320113113102000-3031303012020202-2232220233102131-1131030333123211"></a>

<a id="canonical-2231031232333011-0231201322200020-0121332032320102-0023223030210030-1121110312230321-0313121103332210-1112001311131313-2030132032332331"></a>

## tenant property — sensitive_data_policy / 033221102122 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2212202010100222-2331323201130111-2313301113223320-2323302021000202-1120203311123113-3330231210332120-1023131010203020-0330312323130012"></a>

<a id="canonical-1021123323301131-0031133021133322-3222123203301101-2000333130032323-3003332322120303-1232112203300010-2332003030312222-2331322313103133"></a>

## uid property — sensitive_data_policy / 033221102122 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3332212221231330-2120021220013322-3220011111302122-3310023322311223-1310312032112221-2321221322003031-1231101021312233-2013332123132011"></a>

## Next pages — sensitive_data_policy / 033221102122 / 9

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2020302310011113-0110101030032202-1233233331203200-0321332220211010-1220330200213312-3321332320232003-2113021203322331-0330021133123211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301331131323101-2332321331302031-2001303103323101-3120032002302023-2013020302222030-0131230023022323-3013033300112120-3232023202313133"></a>

## slow_ddos_mitigation — slow_ddos_mitigation / 220121232312 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- slow_ddos_mitigation

<a id="canonical-1220330111031211-1120011223033003-2123202222131321-1201302331203110-1333233222300203-1200012010311001-1323233031232331-0100332123023100"></a>

Type: `"single"`. Computed.

'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

<a id="canonical-3120221223013210-1331033002013333-3232013103233331-1020220310303210-3012203322023131-1010203332130230-1032322313002000-1000220110213323"></a>

## Direct properties — slow_ddos_mitigation / 220121232312 / 3

- [disable_request_timeout](data-sources--virtual_host--reference--group-003.md#canonical-2332311122002120-0102020211112123-3101022133022100-2132210101231300-2123013103112213-2123200322222131-3202202312030120-1203310031330002): complete subsection reference.

<a id="canonical-2201022000222211-2233111121033301-2331123120010132-0323013203231112-2102022322113110-2313110030333310-0011100210121223-0210320030232231"></a>

<a id="canonical-0020301313200002-0121023030333301-1310032111003031-2122310033000320-2310123001233200-0030022302112232-2100230131103201-1323122230312202"></a>

## request_headers_timeout property — slow_ddos_mitigation / 220121232312 / 4

Type: `"number"`. Computed.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-3230302201331122-3001013330302310-3223123011023012-3101202230323030-3121033310132003-3001132133233333-3011121011033001-0311321300131312"></a>

<a id="canonical-1011022320100322-2221121220220031-3230010123012100-2333321310002332-1001000211211223-1123121221211132-1230321130201020-1022331333202220"></a>

## request_timeout property — slow_ddos_mitigation / 220121232312 / 5

Type: `"number"`. Computed.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-1111233000110111-3321032110320013-0213100213030311-3301013132112120-2133011013130200-0331021210331102-3133132001032012-3210012331332301"></a>

## Next pages — slow_ddos_mitigation / 220121232312 / 6

- [slow_ddos_mitigation.disable_request_timeout](data-sources--virtual_host--reference--group-003.md#canonical-2332311122002120-0102020211112123-3101022133022100-2132210101231300-2123013103112213-2123200322222131-3202202312030120-1203310031330002)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2332311122002120-0102020211112123-3101022133022100-2132210101231300-2123013103112213-2123200322222131-3202202312030120-1203310031330002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013032012021321-0321003313013033-1111312101013320-1132000102301011-1213223332231110-3302003233110103-0010220122121010-1313230303123221"></a>

## slow_ddos_mitigation.disable_request_timeout — disable_request_timeout / 123233022113 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [slow_ddos_mitigation](data-sources--virtual_host--reference--group-003.md#canonical-2020302310011113-0110101030032202-1233233331203200-0321332220211010-1220330200213312-3321332320232003-2113021203322331-0330021133123211)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-3202112331010000-2313202003002100-0032311010100012-3120320312113322-2121103100232313-1010200233110313-3312323002331313-2323013233203122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable request timeout.

Upstream description:

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

<a id="canonical-3121133312133320-1102001302302120-0312213331112133-0332300002032011-0232023303112031-0322230222210202-0132032030122123-1111333012022130"></a>

## Direct properties — disable_request_timeout / 123233022113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133310212000031-2112031031320013-0033202210201110-0223232231020332-2221132212122021-1200102210213001-0320030003113312-3203310210331001"></a>

## Next pages — disable_request_timeout / 123233022113 / 4

- [slow_ddos_mitigation](data-sources--virtual_host--reference--group-003.md#canonical-2020302310011113-0110101030032202-1233233331203200-0321332220211010-1220330200213312-3321332320232003-2113021203322331-0330021133123211)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032203020232210-3103120322031031-1323023031121331-3011331010033330-1203122030332300-3032333111032233-2301330211232002-2121310110301003"></a>

## tls_cert_params — tls_cert_params / 003030033333 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- tls_cert_params

<a id="canonical-3032031001222133-3201022301111333-2323000130331230-3210112110333013-1013322112101202-3313113011233020-3302022013222012-2000301331222210"></a>

Type: `"single"`. Computed.

\[OneOf: tls\_cert\_params, tls\_parameters\] Certificate Parameters for authentication, TLS
ciphers, and trust store.

Upstream description:

Certificate Parameters for authentication, TLS ciphers, and trust store.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_certificate_verify_choice": "[\"client_certificate_optional\",\"client_certificate_required\",\"no_client_certificate\"]"
}
```

OneOf alternatives in this subsection:

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-3032031001222133-3201022301111333-2323000130331230-3210112110333013-1013322112101202-3313113011233020-3302022013222012-2000301331222210)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-0013030212110331-2322313130021202-1331130021032333-0333100022023122-2301331212332300-0032103210001031-1030010233201111-2231131010212101)

Select alternatives according to the provider validators above.

<a id="canonical-3010320001323333-0200020303021300-1301223211302232-2302203323103111-3110131011120003-1222111012321010-3100212222100232-2003320021303012"></a>

## Direct properties — tls_cert_params / 003030033333 / 3

- [certificates](data-sources--virtual_host--reference--group-003.md#canonical-2330330032220120-0113103120120230-2302311102231101-2311310130322120-3110012121032233-3201221331111220-0301123210033033-2101230003011212): complete subsection reference.

<a id="canonical-3300013223123001-2110202103320120-1200121330023013-3332001132021311-0220123222132132-2203303231001013-0331100132323300-1331320322002033"></a>

<a id="canonical-0133302133233210-0023111101131330-3121000231210323-1213130221010101-3011022033323312-0122001003022131-1301120012103201-0023003031130101"></a>

## cipher_suites property — tls_cert_params / 003030033333 / 4

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-2122001113001222-2122230031211023-0223212321103110-3000231312023013-2003222030013130-2013010331200302-1212033102132002-2020300113132312): complete subsection reference.

- [client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-0001303310223111-1131201302232133-0110320020000022-1301012001003001-0200232012010012-1203002002323312-0012100122102031-3220031111002110): complete subsection reference.

<a id="canonical-1131110330012331-2221211112212321-3110313301221231-2330203202132300-3323300020312333-2122033010223003-3112221320102130-2323031130220211"></a>

<a id="canonical-0102002032302330-2300331321230202-0221222121310030-2012321323101233-3033330133010000-2102301030231013-2030233020302011-1303212302210033"></a>

## maximum_protocol_version property — tls_cert_params / 003030033333 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-0101030003010310-1113103230302113-3230113031002121-1100010200200232-2003020012321302-0223211012123202-1001212101020032-3201222102003001"></a>

<a id="canonical-1213321021131230-0311032021310310-3000131303023202-1121113221201110-0212233320213103-2003012023300032-3033010103033100-0320023222120310"></a>

## minimum_protocol_version property — tls_cert_params / 003030033333 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

- [no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-0320000122121123-2012030032231200-3120023321321110-2033031203310332-3133231211310321-2221301133132232-1231103323331112-2020300331231203): complete subsection reference.

- [validation_params](data-sources--virtual_host--reference--group-003.md#canonical-3213003322230222-1311331232300133-0222202123333301-3000203322302130-2210010031322323-3023210333303300-3333200301120023-3212003022203233): complete subsection reference.

<a id="canonical-3310302001020000-3312322010233310-2000321003231101-3322201300010222-2300321120032212-3012303220211021-0313022203333323-0001021210021231"></a>

<a id="canonical-3231213002210210-0132311203232013-1332011321021210-1130131010301313-3100122331322030-1123001010210110-3331132332200113-0331333330110133"></a>

## xfcc_header_elements property — tls_cert_params / 003030033333 / 7

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added. Possible values are \`XFCC\_NONE\`, \`XFCC\_CERT\`,
\`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to \`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0113323122113311-3212212212310010-0032110210222103-2031123213121332-1110221123120322-3010100132100330-1030010203031030-2203010213230131"></a>

## Next pages — tls_cert_params / 003030033333 / 8

- [tls_cert_params.certificates](data-sources--virtual_host--reference--group-003.md#canonical-2330330032220120-0113103120120230-2302311102231101-2311310130322120-3110012121032233-3201221331111220-0301123210033033-2101230003011212)
- [tls_cert_params.client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-2122001113001222-2122230031211023-0223212321103110-3000231312023013-2003222030013130-2013010331200302-1212033102132002-2020300113132312)
- [tls_cert_params.client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-0001303310223111-1131201302232133-0110320020000022-1301012001003001-0200232012010012-1203002002323312-0012100122102031-3220031111002110)
- [tls_cert_params.no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-0320000122121123-2012030032231200-3120023321321110-2033031203310332-3133231211310321-2221301133132232-1231103323331112-2020300331231203)
- [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-3213003322230222-1311331232300133-0222202123333301-3000203322302130-2210010031322323-3023210333303300-3333200301120023-3212003022203233)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2330330032220120-0113103120120230-2302311102231101-2311310130322120-3110012121032233-3201221331111220-0301123210033033-2101230003011212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132222001330203-0023103131102110-3331313113302123-2330231312230203-0233020101000123-0122120123011312-0020123323302213-1233222121020111"></a>

## tls_cert_params.certificates — certificates / 110220110102 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- tls_cert_params.certificates

<a id="canonical-3221233032310032-1212121233210100-0122010102321131-0033111303303211-1221232112012301-0301101013322103-0011110230221111-0213323331330132"></a>

Type: `"list"`. Computed.

Certificates. Set of certificates.

Upstream description:

Set of certificates.

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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3110300203303001-0111323332110230-1110333312210031-3301113132111203-2120001122203000-2303222322301331-1330222220012323-3202131112232203"></a>

## Direct properties — certificates / 110220110102 / 3

<a id="canonical-2331312322233001-3121000113203011-2233010221331311-2322002221201221-2110010232321003-2233300323021321-3302121113321032-1333213011131300"></a>

<a id="canonical-1021231123003122-3333003321332212-1111101121032230-3121130101102022-3320121230221202-3000323200213012-0131330102220121-0201232013001123"></a>

## kind property — certificates / 110220110102 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0011133212323013-2121300212212320-1130201002031120-2301132023313210-1112312113321130-1132330032332313-1013222223323111-3123320231112223"></a>

<a id="canonical-0210201112101232-1013022012320001-0100311000202202-2113231130120022-0032303203202330-0323030102300221-2213231210210100-1332022131110210"></a>

## name property — certificates / 110220110102 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2010313321322010-2113320011031333-0202230332321300-0312330122133123-1003202103330303-3020311211200213-1030130312102012-3030113303203323"></a>

<a id="canonical-3120323110202011-2133021023121201-0120023310312333-1020031303310301-3200223221011220-2232120203330332-1001013230000113-1002133331310303"></a>

## namespace property — certificates / 110220110102 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2302323213233122-0002011332320311-3000323113031301-3132302330131222-2311233210310120-1333323030112103-3113001012102333-3003003202101332"></a>

<a id="canonical-2123201210310310-2013312122301121-0223322122200103-3033001303030030-1030022112120311-2232332320212032-0031232202120320-0001102113123100"></a>

## tenant property — certificates / 110220110102 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3311121320100133-1131023023331100-1312203313001132-3101201202113022-3123230123323222-3011311301200213-1203022211211210-3003213132012223"></a>

<a id="canonical-3301320320113320-0303101010111222-3202011113130320-1011323023011313-3301232121231102-3113132002213030-3113331113000113-3320330033121111"></a>

## uid property — certificates / 110220110102 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2112021231033322-0023100233000131-0303233123211111-1101303210120120-2301030320331131-1032203121301332-2033031030313033-2231023203221010"></a>

## Next pages — certificates / 110220110102 / 9

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2122001113001222-2122230031211023-0223212321103110-3000231312023013-2003222030013130-2013010331200302-1212033102132002-2020300113132312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031133332321002-3331232320120333-0233001322131311-2210031222230313-3022303030222303-0210130213210121-3103220110202113-1230033330200013"></a>

## tls_cert_params.client_certificate_optional — client_certificate_optional / 022102022210 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- tls_cert_params.client_certificate_optional

<a id="canonical-0021020200323331-1101002130221300-3220121322010201-3023233303300022-0013223200023131-0221100021322011-1033001331022321-2131111100333022"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-3330231013231033-0001021233233201-0202311101203131-0123322010101013-3130333032100032-1300300221002302-0221022131130231-2111321121031113"></a>

## Direct properties — client_certificate_optional / 022102022210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312323130202202-0010221213002113-1223331222301300-2102031033133123-0213231320013002-1211020332021002-3220212323031211-2212110121320311"></a>

## Next pages — client_certificate_optional / 022102022210 / 4

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0001303310223111-1131201302232133-0110320020000022-1301012001003001-0200232012010012-1203002002323312-0012100122102031-3220031111002110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222033013312212-0101122133333230-2303202113331310-2302022121322303-1211032020303300-2030101032223130-1012233302133120-3002223121203112"></a>

## tls_cert_params.client_certificate_required — client_certificate_required / 103100121000 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- tls_cert_params.client_certificate_required

<a id="canonical-1310331110032112-3002212333220332-0011213300131032-1200013023013220-3212021223313301-0302122231011101-1321020103130121-2013232113023233"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2213113012022022-0301133100222232-0310123312303231-0300022200200301-0310212033223331-0313202121031233-0301120321121000-0011220123321103"></a>

## Direct properties — client_certificate_required / 103100121000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202033200030203-2303001200231323-1222232323230333-0212200102202333-2010203322031330-1133312101311010-3232311133320103-1121100331011333"></a>

## Next pages — client_certificate_required / 103100121000 / 4

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0320000122121123-2012030032231200-3120023321321110-2033031203310332-3133231211310321-2221301133132232-1231103323331112-2020300331231203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133331031310121-0320213313312111-0311103123031201-3322310032020213-1222221102200303-1032303121231321-2030123003100022-2223220210130332"></a>

## tls_cert_params.no_client_certificate — no_client_certificate / 333032020023 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- tls_cert_params.no_client_certificate

<a id="canonical-0210111112322023-2302233322300033-2211201100101112-3302123300112212-0022310311221102-1112120223020133-1232222230213123-2012112021333133"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-3202013112232202-3231121133000331-1333111202203332-0031213133333112-1111111011102223-2112213030030032-0020001113332221-1103103101033030"></a>

## Direct properties — no_client_certificate / 333032020023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221312231101312-2112003002013233-1322303210031212-0323022030121322-3133111132201201-3213112123103322-2022302112113133-1322112131211230"></a>

## Next pages — no_client_certificate / 333032020023 / 4

- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3213003322230222-1311331232300133-0222202123333301-3000203322302130-2210010031322323-3023210333303300-3333200301120023-3212003022203233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111212331113332-0233210310121303-2332232113233223-1221022331031110-2313101331021210-3121231122222122-2232131002202311-1311301322230203"></a>

## tls_cert_params.validation_params — validation_params / 321113132232 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- tls_cert_params.validation_params

<a id="canonical-2300302222301320-1013220320220122-2011112013302323-0003012133101110-2211222202003322-3003231101333301-1321221232232330-1330130311101123"></a>

Type: `"single"`. Computed.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-0210310022322000-1022103300003321-3102212221221221-3110221033131010-0112113011133031-2023200131302322-1112112013100003-0023231123010002"></a>

## Direct properties — validation_params / 321113132232 / 3

<a id="canonical-3320123120020312-2330230202230021-3222031121223112-0033231310023012-0121101333011131-3223211331322313-3003333221300120-1220202221101312"></a>

<a id="canonical-2101022131300022-3020032123310010-0332132331102321-1302003112011233-2302200221102303-2011221322233101-1023013331323220-3023321113013313"></a>

## skip_hostname_verification property — validation_params / 321113132232 / 4

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-1023311023232233-3220101100112021-1002203003331330-3333121303303220-2010202202032031-2331011122033221-1232000233230233-2221230013312223): complete subsection reference.

<a id="canonical-3121111203320320-3330323231303011-3031120310021332-2112002002112231-0030031211132121-3312023201203212-0233001313311131-0011311310103012"></a>

<a id="canonical-1120003120313033-1121310203030133-2022200032121220-1123033000223301-1031130211021201-2311221323203312-3211101031123003-1233003021130021"></a>

## trusted_ca_url property — validation_params / 321113132232 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-0202011000311301-2131110123021013-3132112300332303-2121210000202202-1221313221210323-3223123000000313-2122012321133201-3030321130313212"></a>

<a id="canonical-3112132223110330-1031113311122222-0033111010031232-0213000203330313-1222332011210203-2223322113132012-1233111123122022-0303210132123201"></a>

## verify_subject_alt_names property — validation_params / 321113132232 / 6

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3131320231202123-1200303113221310-1002330133103001-0202303311300220-2320033022200011-0030033210311200-2033131120200331-3033223220322033"></a>

## Next pages — validation_params / 321113132232 / 7

- [tls_cert_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-1023311023232233-3220101100112021-1002203003331330-3333121303303220-2010202202032031-2331011122033221-1232000233230233-2221230013312223)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1023311023232233-3220101100112021-1002203003331330-3333121303303220-2010202202032031-2331011122033221-1232000233230233-2221230013312223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311023033203213-0323331232233220-2320020112320222-0321132201323110-2032111212132021-3030000101012100-2223202012002012-1112133223131033"></a>

## tls_cert_params.validation_params.trusted_ca — trusted_ca / 221310332332 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-3213003322230222-1311331232300133-0222202123333301-3000203322302130-2210010031322323-3023210333303300-3333200301120023-3212003022203233)
- tls_cert_params.validation_params.trusted_ca

<a id="canonical-2120312133211320-0213010030203123-2212023213213332-3301130020302332-2300100101203032-2002001113213200-1120233211121010-0122112323012112"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3200210113011001-3031220321102321-2222310120301003-0323121111323213-2323130020323020-3031233301131033-2302331200013321-1300000222210323"></a>

## Direct properties — trusted_ca / 221310332332 / 3

- [trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-3123033001322311-3130310302032023-3021020200222003-3012011023123120-3210013130231123-2200312311130121-3212320230301331-3110222012020102): complete subsection reference.

<a id="canonical-0202210100202321-3122013332112330-1130123001131330-0101232000133312-0113231233021223-3031333023232020-0232312023102333-0120023300303131"></a>

## Next pages — trusted_ca / 221310332332 / 4

- [tls_cert_params.validation_params.trusted_ca.trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-3123033001322311-3130310302032023-3021020200222003-3012011023123120-3210013130231123-2200312311130121-3212320230301331-3110222012020102)
- [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-3213003322230222-1311331232300133-0222202123333301-3000203322302130-2210010031322323-3023210333303300-3333200301120023-3212003022203233)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3123033001322311-3130310302032023-3021020200222003-3012011023123120-3210013130231123-2200312311130121-3212320230301331-3110222012020102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010231220021020-1222323132123132-2000232110231222-1232102331221001-0031100130310022-0000012112111120-1123031332330333-1331122122321111"></a>

## tls_cert_params.validation_params.trusted_ca.trusted_ca_list — trusted_ca_list / 232002232323 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_cert_params](data-sources--virtual_host--reference--group-003.md#canonical-2321032002201002-2013003203020122-3031331300011230-0101031132103321-1031323221002200-2101332033001223-3033303303231132-3113100001111311)
- [tls_cert_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-3213003322230222-1311331232300133-0222202123333301-3000203322302130-2210010031322323-3023210333303300-3333200301120023-3212003022203233)
- [tls_cert_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-1023311023232233-3220101100112021-1002203003331330-3333121303303220-2010202202032031-2331011122033221-1232000233230233-2221230013312223)
- tls_cert_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-2322201223331011-1022202131013311-0012313030331113-1201110100110101-0000101131113030-2333112220310021-3010212122021220-0000300221301301"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3220122022302233-1131210313211220-1003313320002020-2321002122322231-2202021003200110-1031311112130231-1111001131033221-1213220132021000"></a>

## Direct properties — trusted_ca_list / 232002232323 / 3

<a id="canonical-2322133201223202-0211300020200333-3112010321130301-1023303210332031-3130203131022211-3133123200323202-1301003231003211-2103110113320300"></a>

<a id="canonical-2000122221333003-2130033021130200-0213333112312000-0231131230323112-0122201211310103-3212333010031001-0302311300230303-3001113301233000"></a>

## kind property — trusted_ca_list / 232002232323 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0203111133022121-0012200212331321-3333021230121123-2012003010301332-1033013301210213-0202330201300022-0231302103310221-1220120102102220"></a>

<a id="canonical-2110310002131022-2320120132201333-2021001330213010-1013222133021232-3312133022132200-2032132203132201-1331023100202113-1120230130312112"></a>

## name property — trusted_ca_list / 232002232323 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0020100332200221-2113321020230013-0120130211110000-2022223031332200-3103132231131033-2301330123211133-1010301332010300-1020313311200332"></a>

<a id="canonical-1301110310023213-1301213123110300-1301130331200331-2021122021320020-2212200103023111-3123200213322032-0233132210010233-0100221123221000"></a>

## namespace property — trusted_ca_list / 232002232323 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3210031221311021-0313012001202012-1220110122213101-2321102223232113-0223032000021133-1002301012003020-3311232103100303-0022010230332100"></a>

<a id="canonical-1100011110211200-2223231110030020-1333212110303002-0122123313301301-1132322201110010-0131330300013323-3010020012032131-2323213132203302"></a>

## tenant property — trusted_ca_list / 232002232323 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3301020111002020-3320012021321331-0131013011312100-1002220220322213-0212131123100033-0320001330223211-1230002230201103-1200233010300100"></a>

<a id="canonical-1320200121213213-3230003320023323-0000303113333132-0122031320333311-0113212121023313-3331332311023211-3020121001020001-3231113303220231"></a>

## uid property — trusted_ca_list / 232002232323 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3100131333130213-2200320211313211-0113031020312311-3200031021213130-0031030303020032-1330010102002103-3102232322030102-0001133212333332"></a>

## Next pages — trusted_ca_list / 232002232323 / 9

- [tls_cert_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-1023311023232233-3220101100112021-1002203003331330-3333121303303220-2010202202032031-2331011122033221-1232000233230233-2221230013312223)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322032232033230-3102300022321212-1003102323301111-3100012323013120-3020232132220132-2232232332300101-0030101210232133-2123112000030200"></a>

## tls_parameters — tls_parameters / 102323331313 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- tls_parameters

<a id="canonical-0013030212110331-2322313130021202-1331130021032333-0333100022023122-2301331212332300-0032103210001031-1030010233201111-2231131010212101"></a>

Type: `"single"`. Computed.

TLS configuration for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_certificate_verify_choice": "[\"client_certificate_optional\",\"client_certificate_required\",\"no_client_certificate\"]"
}
```

<a id="canonical-2333102111322322-0322111203333221-1202102333330122-1213031021332203-3130001012023013-1320111020331032-2023110000030300-0132012133333111"></a>

## Direct properties — tls_parameters / 102323331313 / 3

- [client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-0311213131212331-3033101301300130-0312230321201212-1030121302231320-2031103222211101-2130230222311110-1331212232232032-3232101311001120): complete subsection reference.

- [client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-3100311320103231-0130313003300212-2111010023132233-1002011031210200-1123131203311210-3013330113032213-1210102122213333-3201332021102312): complete subsection reference.

- [common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111): complete subsection reference.

- [no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-0003223203201211-1212131120130013-0331213103230031-2310001021211300-2022120121120211-0331000312213003-2223103133010011-2001303022303313): complete subsection reference.

<a id="canonical-3200332110312223-1011311221201220-1310300302322102-3213122331332330-1330232002303333-0010023311221110-3221003021203101-1030330303010122"></a>

<a id="canonical-3333012212232013-3003022023102222-3301310231230331-2003010321333323-2300201130022211-1130023012031002-1330210102213123-1030121000310122"></a>

## xfcc_header_elements property — tls_parameters / 102323331313 / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added. Possible values are \`XFCC\_NONE\`, \`XFCC\_CERT\`,
\`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to \`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3023303031032330-3111103100001320-0103203111030000-1311030020121211-3020212223212302-2222313033332032-2133321023022301-2330112203311203"></a>

## Next pages — tls_parameters / 102323331313 / 5

- [tls_parameters.client_certificate_optional](data-sources--virtual_host--reference--group-003.md#canonical-0311213131212331-3033101301300130-0312230321201212-1030121302231320-2031103222211101-2130230222311110-1331212232232032-3232101311001120)
- [tls_parameters.client_certificate_required](data-sources--virtual_host--reference--group-003.md#canonical-3100311320103231-0130313003300212-2111010023132233-1002011031210200-1123131203311210-3013330113032213-1210102122213333-3201332021102312)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.no_client_certificate](data-sources--virtual_host--reference--group-003.md#canonical-0003223203201211-1212131120130013-0331213103230031-2310001021211300-2022120121120211-0331000312213003-2223103133010011-2001303022303313)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0311213131212331-3033101301300130-0312230321201212-1030121302231320-2031103222211101-2130230222311110-1331212232232032-3232101311001120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131002202301112-0320233213221212-0330010323302223-0300202311013303-3122001310212132-1013322331220331-3100310320301003-0033131222113110"></a>

## tls_parameters.client_certificate_optional — client_certificate_optional / 011132302121 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- tls_parameters.client_certificate_optional

<a id="canonical-3221212102110032-2212333010322130-1110301031211112-0233300112310032-1220113321013000-3223321211103002-1021230021203211-3100311011212110"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-3000002323113212-3212330203103201-2000321222023231-0200132320311301-1220232103001011-3120032310331202-0203111022213103-0030012210133312"></a>

## Direct properties — client_certificate_optional / 011132302121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032233012132331-0223220231320131-0320200033001100-2132303030200221-3110100030222012-3310320210120321-3002002102320221-1221301023231023"></a>

## Next pages — client_certificate_optional / 011132302121 / 4

- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3100311320103231-0130313003300212-2111010023132233-1002011031210200-1123131203311210-3013330113032213-1210102122213333-3201332021102312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012200020023313-3113332223203102-1121201011020233-2132212232300120-0320003322211202-0113202112320103-1222333223132132-1310003212023123"></a>

## tls_parameters.client_certificate_required — client_certificate_required / 323331232332 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- tls_parameters.client_certificate_required

<a id="canonical-1322121022021001-1211013131003133-2233122232330113-2203121213300003-3132332233300311-2211003133210313-0312020003120012-2020002303032033"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-3312220321013002-1230211210033000-3303231231330020-0132003031112231-3301030110212020-2300323321320211-3321120333331031-2222213112120332"></a>

## Direct properties — client_certificate_required / 323331232332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311312303031113-1031201210310111-1210133230100223-3211233212232331-2013002101131122-1303222113201031-1002003111201113-1102220012322000"></a>

## Next pages — client_certificate_required / 323331232332 / 4

- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210302220330121-3112232323121230-0220021100300021-3001103311033112-0221103221332023-1020333100232310-3320011121320333-3223321121031012"></a>

## tls_parameters.common_params — common_params / 121020322111 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- tls_parameters.common_params

<a id="canonical-1122233013102130-2021122230220223-3002031003230030-3020321300000222-1322010202110003-2122231302323200-2203220302313131-2323311001200032"></a>

Type: `"single"`. Computed.

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

Upstream description:

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0122231313122002-1313333311300302-1223101123210031-1011110033201333-1332001032221220-0203031302311312-3300020100302022-3312311131120101"></a>

## Direct properties — common_params / 121020322111 / 3

<a id="canonical-0210103132203322-1000023122230330-1203201210212103-2301100121230113-0221302200221113-2212212100332303-0021112331001121-3310301321303211"></a>

<a id="canonical-0232132103112131-3011320300023121-2202130120121330-0000323231031131-1320020030002211-3212003212203033-0320333300230133-2231233132331303"></a>

## cipher_suites property — common_params / 121020322111 / 4

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3311002010103122-1013132310120330-0311130211112101-1233333323023213-2023300300221103-2123210233300320-1331211100323023-0123321210110212"></a>

<a id="canonical-0101121333122212-2313230211113210-3232122022011110-2320030010101101-2213031213003231-1233212121212231-3213302300010312-1112212030020012"></a>

## maximum_protocol_version property — common_params / 121020322111 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-2010113220100203-3122102002113033-1302302001303030-1203203331122233-1122311133000133-1033311000100001-0302011121113000-0322333003012230"></a>

<a id="canonical-3021012312330011-1301100200133323-1122210003111313-3010112213032000-2331303330121023-1002331101211120-3102311031103223-1003123130213110"></a>

## minimum_protocol_version property — common_params / 121020322111 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

- [tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312): complete subsection reference.

- [validation_params](data-sources--virtual_host--reference--group-003.md#canonical-1011200130001331-1320301220122001-3310301011012333-1301113202111223-0010221100002031-1203321002203131-1110303301010000-2102112302131230): complete subsection reference.

<a id="canonical-1002011120200300-2301031111031301-0021032023312313-3111013031222001-1231020020330002-3320312130212313-3012120131210233-2001100131333023"></a>

## Next pages — common_params / 121020322111 / 7

- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-1011200130001331-1320301220122001-3310301011012333-1301113202111223-0010221100002031-1203321002203131-1110303301010000-2102112302131230)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121222333010020-3123123020332321-3303133031312130-0033212232231303-3320232200332303-2231233333023111-0222031222323302-1211323021332203"></a>

## tls_parameters.common_params.tls_certificates — tls_certificates / 322320031000 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- tls_parameters.common_params.tls_certificates

<a id="canonical-3122211203131331-1110033333310310-1013132323031321-2023210123002111-2321322320000113-3321322313230001-2322021021212302-0321000021133001"></a>

Type: `"list"`. Computed.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1202211223012312-0200202033201113-0300331030031213-1001320200311032-2031020232223210-2102333131111310-3332303320301020-1030010123013130"></a>

## Direct properties — tls_certificates / 322320031000 / 3

<a id="canonical-0130131200233131-3102001332233101-3111311001312113-0021333333233323-3201133301032220-3131301230200012-1213301010211021-3300123211231210"></a>

<a id="canonical-1321131122022103-1023122231010101-1123312211222123-3302100313013223-3022310331100012-3303201112303032-0112123031131010-1231300032221230"></a>

## certificate_url property — tls_certificates / 322320031000 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [custom_hash_algorithms](data-sources--virtual_host--reference--group-003.md#canonical-2013221330301301-0200121011220020-0131100030133132-3111220100302013-0323301013010002-0032023213122022-1113030220331103-3321222302033332): complete subsection reference.

<a id="canonical-1112213221331013-1033133021011102-3230011200010333-1130220130001201-1201303333121131-0110032021011112-2132112133001020-0221123210030120"></a>

<a id="canonical-1120322112033013-0020321202201220-1202121301313100-0312223311221100-0322321313332021-3312010210313233-3233223033330102-2313201213121012"></a>

## description_spec property — tls_certificates / 322320031000 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--virtual_host--reference--group-003.md#canonical-1210223002300101-1110011013032331-1312303120101220-2322030232030110-3011330011000030-3022000300311013-3330122120300020-0310201103231212): complete subsection reference.

- [private_key](data-sources--virtual_host--reference--group-003.md#canonical-0032311131222331-3223213221002220-3130303232312133-0313321011032200-2101111120023012-2130010110101200-0311300221011321-3103120121101222): complete subsection reference.

- [use_system_defaults](data-sources--virtual_host--reference--group-003.md#canonical-2032000021233011-0333311310031021-0332130220312011-3322222321000123-1322322122132013-1131310310222201-0323023113221212-3333200213120130): complete subsection reference.

<a id="canonical-1020120313002313-2131323222100120-1021312210022101-1332011201022220-2112132302023222-2102201300022121-2211000023101010-2213332310120112"></a>

## Next pages — tls_certificates / 322320031000 / 6

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--virtual_host--reference--group-003.md#canonical-2013221330301301-0200121011220020-0131100030133132-3111220100302013-0323301013010002-0032023213122022-1113030220331103-3321222302033332)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--virtual_host--reference--group-003.md#canonical-1210223002300101-1110011013032331-1312303120101220-2322030232030110-3011330011000030-3022000300311013-3330122120300020-0310201103231212)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-0032311131222331-3223213221002220-3130303232312133-0313321011032200-2101111120023012-2130010110101200-0311300221011321-3103120121101222)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--virtual_host--reference--group-003.md#canonical-2032000021233011-0333311310031021-0332130220312011-3322222321000123-1322322122132013-1131310310222201-0323023113221212-3333200213120130)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2013221330301301-0200121011220020-0131100030133132-3111220100302013-0323301013010002-0032023213122022-1113030220331103-3321222302033332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011313203221101-2221300220022132-0322011002023203-2310110311103232-2030323202110033-3312223101332312-2321010013010031-2331002031132020"></a>

## tls_parameters.common_params.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 212213203100 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-1013100311311122-3133332231121321-0011003022010013-0022130200031322-1013213110331100-1302203032133203-2132003020022003-2203222120223211"></a>

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

<a id="canonical-1322012222302223-2033122102002133-2031232333333212-3333312000123332-2312322120212312-2010032222203202-1102001001233012-3133121303233003"></a>

## Direct properties — custom_hash_algorithms / 212213203100 / 3

<a id="canonical-0333323320231131-3112231210321031-1101203220011032-1030113210033133-0032022113023211-3302110133201303-1031332300002001-1221300003002300"></a>

<a id="canonical-2000301011030213-3013313230211332-2222201120332011-0311002012203002-3213011133030230-0022303320101101-3332022023333220-0133120113120221"></a>

## hash_algorithms property — custom_hash_algorithms / 212213203100 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3021102323313301-2322203120203032-3311212000123122-3001303213201323-2322010123033103-2132133211133122-0312213201311300-3120231211032011"></a>

## Next pages — custom_hash_algorithms / 212213203100 / 5

- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1210223002300101-1110011013032331-1312303120101220-2322030232030110-3011330011000030-3022000300311013-3330122120300020-0310201103231212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332313332313101-2211101011223113-0000023323213310-0022003331211133-2222323032200013-3113313213112012-3031230332201323-3330231231300103"></a>

## tls_parameters.common_params.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 133301311013 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-1032000020301032-1220002203030322-3011131031012231-2021333010023032-1333201221002002-2120120103230222-3022023201322203-2202221221002310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

Upstream description:

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

<a id="canonical-2120101333320212-1333222323301201-1310033222003003-0013100332021221-3021210103232011-1332220130301121-0332120233312213-1012131210321201"></a>

## Direct properties — disable_ocsp_stapling / 133301311013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030232031031323-2122333002212323-0211100330210200-2220111120122203-0113231301131232-3031030201122121-3220133123213303-2013123222331032"></a>

## Next pages — disable_ocsp_stapling / 133301311013 / 4

- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0032311131222331-3223213221002220-3130303232312133-0313321011032200-2101111120023012-2130010110101200-0311300221011321-3103120121101222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132311033130203-0113311301301332-2132110131310201-0203322132023203-0203110122302122-1001212200023030-1211320211112003-2113203211300310"></a>

## tls_parameters.common_params.tls_certificates.private_key — private_key / 331231101103 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-0132213332013300-3330132103202111-1123122112012003-2220302132012101-2032311122122211-2212023200322032-2310023213313333-3320300322330310"></a>

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

<a id="canonical-1032011212231210-2320202313103101-3101132312110220-0102322330000323-3200120311122100-1302012201121232-1312220230122222-2123033031301133"></a>

## Direct properties — private_key / 331231101103 / 3

- [blindfold_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-0212231131312301-3010010002131122-3333002010103203-0011203120313012-0133100230302122-3033100201321221-1331232010020232-2001310111231130): complete subsection reference.

- [clear_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-1030232231122130-3312120101201033-2312010133121330-3313001200230320-2020131003022200-1220101032132000-0313312030132203-3101322303003323): complete subsection reference.

<a id="canonical-3302312130200112-3012131001330030-3321230002303331-0300121203121213-2321111330202023-3122101202022122-0203032313321300-0022130130331031"></a>

## Next pages — private_key / 331231101103 / 4

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-0212231131312301-3010010002131122-3333002010103203-0011203120313012-0133100230302122-3033100201321221-1331232010020232-2001310111231130)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--virtual_host--reference--group-003.md#canonical-1030232231122130-3312120101201033-2312010133121330-3313001200230320-2020131003022200-1220101032132000-0313312030132203-3101322303003323)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0212231131312301-3010010002131122-3333002010103203-0011203120313012-0133100230302122-3033100201321221-1331232010020232-2001310111231130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300233323212200-1333322023302220-0322203211220230-3130013023110002-3211222203121121-3302011020132220-0030322100221111-0330232230210212"></a>

## tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 110303033031 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-0032311131222331-3223213221002220-3130303232312133-0313321011032200-2101111120023012-2130010110101200-0311300221011321-3103120121101222)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1122221313200001-1233111222132323-2102220222033021-3301220113113232-2223011220201030-0020102100120322-0331001311223303-3303202102232133"></a>

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

<a id="canonical-2320223233303203-0000231323022021-3300033301013312-0312330131321222-2022023122213032-3123023200033332-2301022223212322-1233210321002033"></a>

## Direct properties — blindfold_secret_info / 110303033031 / 3

<a id="canonical-2330012103000311-3301320210020113-1030130102020231-1300322021030020-1003010232212003-1002333231011310-2311231013322011-1322010331122310"></a>

<a id="canonical-2130130321023102-0010010331022232-1130013323331001-2313232301030033-3232022323322201-1213221211032103-1133201311322020-3130330131322132"></a>

## decryption_provider property — blindfold_secret_info / 110303033031 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0123303221223230-0110023130232133-1122313022202321-1321230303231133-2131013033210032-1100301301221121-1113300011322213-1131023312022123"></a>

<a id="canonical-0112122101333313-2303231023030023-1223220203003223-1132131211010213-3323221233323330-0333200213022302-2032123310003032-3311331200320011"></a>

## location property — blindfold_secret_info / 110303033031 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2231210311332313-1112022311211300-0200211232013110-1313032332000221-1211331230110022-3312311202120110-3022202310100322-3313201302021230"></a>

<a id="canonical-2102013302002321-2302301023020022-0222023133331310-2013203133101121-3033321112300211-0322230202121000-2230323330321311-0101100200221101"></a>

## store_provider property — blindfold_secret_info / 110303033031 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1312022111323031-0112212122032320-0222110232123113-3231031330321220-0001132213321222-3202112122010130-1010030111123102-1122220302031323"></a>

## Next pages — blindfold_secret_info / 110303033031 / 7

- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-0032311131222331-3223213221002220-3130303232312133-0313321011032200-2101111120023012-2130010110101200-0311300221011321-3103120121101222)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1030232231122130-3312120101201033-2312010133121330-3313001200230320-2020131003022200-1220101032132000-0313312030132203-3101322303003323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200102131112001-0222030130123232-2103122232012230-1200202220011311-2003010001322301-3230022320333303-3100113323031001-1303222300221002"></a>

## tls_parameters.common_params.tls_certificates.private_key.clear_secret_info — clear_secret_info / 003121312102 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-0032311131222331-3223213221002220-3130303232312133-0313321011032200-2101111120023012-2130010110101200-0311300221011321-3103120121101222)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-0102230201001020-3303302322012223-3222113230100120-0223231012231200-0303203121101011-1131310323021032-3030033333011100-3333000032000002"></a>

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

<a id="canonical-1202201003000003-2112233121100220-2103212210121312-3332000331323312-2011121221000012-0313011010221302-1332300131212111-0232000222112223"></a>

## Direct properties — clear_secret_info / 003121312102 / 3

<a id="canonical-3332122332302233-2300202203032313-0113122321220012-2122120032220312-2103212212222201-2001123100233031-1232222020320132-0331300103212101"></a>

<a id="canonical-0033120202100302-3102010022311002-1320221330021320-3333100333203221-1231120222101321-2003023110301313-0220112030101211-3103222302301131"></a>

## provider_ref property — clear_secret_info / 003121312102 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0323302110030333-1013121220021123-1301121113031233-3131020213232021-3102311233211130-2311211201200132-1100031320002011-3221332202201013"></a>

<a id="canonical-2321001300001303-2310210200101102-3220131122122213-3121210020223020-3312110011113332-3322211320300000-1023200200130231-1202221003233100"></a>

## URL property — clear_secret_info / 003121312102 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2232230112023320-2331002012331100-2211113132000303-0103303231011220-3233130133220030-0331023331010011-1000002012202031-0320201331322300"></a>

## Next pages — clear_secret_info / 003121312102 / 6

- [tls_parameters.common_params.tls_certificates.private_key](data-sources--virtual_host--reference--group-003.md#canonical-0032311131222331-3223213221002220-3130303232312133-0313321011032200-2101111120023012-2130010110101200-0311300221011321-3103120121101222)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2032000021233011-0333311310031021-0332130220312011-3322222321000123-1322322122132013-1131310310222201-0323023113221212-3333200213120130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223313010313132-1300012320001121-2222030121102322-3132032013230030-2020212233032230-2111010202320021-1033223330012003-3201322211103323"></a>

## tls_parameters.common_params.tls_certificates.use_system_defaults — use_system_defaults / 001301132303 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-2110303302121311-1330010201013331-1133321110131010-2033121213220212-3012002021330332-0133311012231022-2233202131232210-3203230332212322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

Upstream description:

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

<a id="canonical-0323020332211103-3001120021032302-3102213023033200-0022123223333223-1300322330103200-2223032030311010-1322100120232121-3222221031211120"></a>

## Direct properties — use_system_defaults / 001301132303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020123213323032-3312011121202213-1012113103111113-3030111320112333-1110312201302210-2203230030301110-3123312102023122-1012300013003130"></a>

## Next pages — use_system_defaults / 001301132303 / 4

- [tls_parameters.common_params.tls_certificates](data-sources--virtual_host--reference--group-003.md#canonical-1100031231013202-1322001023103001-3233302102202000-2112120202022113-1010330103023233-2131201330111302-2010203212233211-2003123332333312)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1011200130001331-1320301220122001-3310301011012333-1301113202111223-0010221100002031-1203321002203131-1110303301010000-2102112302131230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030133202311123-2110303220030213-2202021113323220-0110202123233120-3331222130121123-3010331022332022-2321332301333213-2100021112022231"></a>

## tls_parameters.common_params.validation_params — validation_params / 031130213012 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- tls_parameters.common_params.validation_params

<a id="canonical-2212012212101320-3112001032323123-2033201010100210-2133103003303313-3033022003211212-0213131112032203-2233331211102110-2032033120003310"></a>

Type: `"single"`. Computed.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-2013210020012333-1022113233222122-1331322312032132-1220311010001001-3223023332201023-2202203310322300-1300023333223011-1333312332132302"></a>

## Direct properties — validation_params / 031130213012 / 3

<a id="canonical-2030113320233213-0220133300211011-3122221331030231-1012202013002210-1200223021303102-3103213121121100-0323123213331301-0202123200321110"></a>

<a id="canonical-1222021223112312-0323330320132321-2023210202110003-3311231032002220-3231100011210322-3311120032002102-0212202323111003-0003301323000131"></a>

## skip_hostname_verification property — validation_params / 031130213012 / 4

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-3333331333311130-0131012003032211-0001030100130121-3211202103121133-2033233301023120-2201133131321102-3012333212020221-1322102322122012): complete subsection reference.

<a id="canonical-2221202123013033-0301222212222230-0002103322012211-1223210100330122-3332211023131310-0300223232202130-1310311203231300-2301201132223210"></a>

<a id="canonical-0301000122022223-1123313313102033-0200202023333113-1210313132330232-0102111232233320-2010211122223003-1321201203102103-0120200203110201"></a>

## trusted_ca_url property — validation_params / 031130213012 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2030332013012000-2133110033200303-2303032112332220-0031023223132301-2031331013200132-3121311113000130-3113133002323002-3230013220202200"></a>

<a id="canonical-1111123201222113-1113000313111320-3203023100013312-2100232113030131-0322030123031113-1112221121011001-3020230133112303-0322313230230332"></a>

## verify_subject_alt_names property — validation_params / 031130213012 / 6

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2211222021301303-1211122321202323-2102323312031221-0121012330312310-2330030220112303-3130322311131022-0213310222301220-0303123030201002"></a>

## Next pages — validation_params / 031130213012 / 7

- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-3333331333311130-0131012003032211-0001030100130121-3211202103121133-2033233301023120-2201133131321102-3012333212020221-1322102322122012)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3333331333311130-0131012003032211-0001030100130121-3211202103121133-2033233301023120-2201133131321102-3012333212020221-1322102322122012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021303302112113-0321213131122131-0331122302322310-1102213331122302-1232300000103100-3330001302032103-3221232022332120-3320203323012220"></a>

## tls_parameters.common_params.validation_params.trusted_ca — trusted_ca / 231213031011 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-1011200130001331-1320301220122001-3310301011012333-1301113202111223-0010221100002031-1203321002203131-1110303301010000-2102112302131230)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-0110222322332222-3321211030330203-0222013221030130-2231232312323102-1203032313103202-0303220312110330-0330122230200302-0230033112222122"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0213131333032103-2232233120103233-2131110032210310-2011111211210112-2012331131231222-0022102331220231-0100123020213311-1132131221203011"></a>

## Direct properties — trusted_ca / 231213031011 / 3

- [trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-2212123223300122-3011100310011320-1221112011132033-1101133222300003-0211321302002122-1121130121210322-0330000210323101-3321132332322301): complete subsection reference.

<a id="canonical-1113132013313220-3301030120001213-3120122311020113-0123303110230101-0021311033302231-0232020113300120-1003222202200310-3311130001332010"></a>

## Next pages — trusted_ca / 231213031011 / 4

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--virtual_host--reference--group-003.md#canonical-2212123223300122-3011100310011320-1221112011132033-1101133222300003-0211321302002122-1121130121210322-0330000210323101-3321132332322301)
- [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-1011200130001331-1320301220122001-3310301011012333-1301113202111223-0010221100002031-1203321002203131-1110303301010000-2102112302131230)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2212123223300122-3011100310011320-1221112011132033-1101133222300003-0211321302002122-1121130121210322-0330000210323101-3321132332322301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331333022120333-0223313300210322-1230122222022012-3321310030012233-3131230110233013-2133233232103102-2102311121001211-0221220122013313"></a>

## tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list — trusted_ca_list / 300202203232 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [tls_parameters.common_params](data-sources--virtual_host--reference--group-003.md#canonical-2012001013302103-2322123210310212-0133203100233331-1001131212022101-2100300232123221-2211121020013100-1121332020133311-3013113201100111)
- [tls_parameters.common_params.validation_params](data-sources--virtual_host--reference--group-003.md#canonical-1011200130001331-1320301220122001-3310301011012333-1301113202111223-0010221100002031-1203321002203131-1110303301010000-2102112302131230)
- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-3333331333311130-0131012003032211-0001030100130121-3211202103121133-2033233301023120-2201133131321102-3012333212020221-1322102322122012)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-1033000020332202-0023201321102003-0021322013210030-0121310110201111-0200220132013113-3211000033320120-1110013120222300-1321100223231013"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0222201130321302-1211002023112132-0032131110011121-0310112302321031-0111323202011212-3232201013121111-3103310033330013-1132202332002122"></a>

## Direct properties — trusted_ca_list / 300202203232 / 3

<a id="canonical-3213102212203012-1011323021323312-2133323212032212-1110122023202100-3202131112120123-3220230301330321-3033312132230303-3133002212001320"></a>

<a id="canonical-3211201030111121-1333331233222232-3232002001002110-1220301201313130-1123333200232312-3331020233312221-0223003112030132-0001321030312003"></a>

## kind property — trusted_ca_list / 300202203232 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3321030310313130-0123031000211330-3011131020210111-0112032230020130-0113212300003233-1203020021031122-1101202303111221-3022122322311321"></a>

<a id="canonical-1010200120311032-1113131212312230-1133331222322200-3032030200030320-2023233310220210-1333101031111221-1230023212012201-2301022221100123"></a>

## name property — trusted_ca_list / 300202203232 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3331123303122121-2031001311331303-2232012031321123-1012101132310002-3201200202011233-1201111103301222-0021112030303200-1333320310013212"></a>

<a id="canonical-2231130003222310-0111303133202102-1120110131210233-1321123311003023-0102213201302120-2002133221131113-1101003201313102-2333213131131213"></a>

## namespace property — trusted_ca_list / 300202203232 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2120122322022312-2330012223232210-0321010113112312-1203331221102222-0221032020003301-0132030133322001-1232000023320122-2202312001130002"></a>

<a id="canonical-3323121031203113-1011230013103123-1200003221323210-0112201120300030-3211022110123331-0311330013232002-1213231132230022-2120320212023023"></a>

## tenant property — trusted_ca_list / 300202203232 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1322130113222233-3221102020013002-2311301131232222-2231011212310112-2133022103213020-0022223130120321-1020223010310320-0213033333200123"></a>

<a id="canonical-3110330111322023-2101313021002232-0133202320130130-3320222130123200-2211233120223003-1111303013112032-1300022210330101-2203021022311133"></a>

## uid property — trusted_ca_list / 300202203232 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0213300232313010-2233101030233003-1330102033311130-1331033330120302-3311111003231202-3231113132122123-0033300210231032-3003132000332011"></a>

## Next pages — trusted_ca_list / 300202203232 / 9

- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--virtual_host--reference--group-003.md#canonical-3333331333311130-0131012003032211-0001030100130121-3211202103121133-2033233301023120-2201133131321102-3012333212020221-1322102322122012)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-0003223203201211-1212131120130013-0331213103230031-2310001021211300-2022120121120211-0331000312213003-2223103133010011-2001303022303313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123133203102230-1333201021233203-1232020221311203-3301233120103021-2130032321200011-1011321012231302-2112200021012121-0101211032332003"></a>

## tls_parameters.no_client_certificate — no_client_certificate / 321231331023 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- tls_parameters.no_client_certificate

<a id="canonical-2003121333303301-2111312111203022-1320320130233311-3033000220012333-0201022301310033-1312123300120330-3201231231111210-3213012001011200"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2130100332200311-1131232203200201-1330112003032230-2310113301111301-0222102200120202-2333112203001331-3033021003020222-2200222220130013"></a>

## Direct properties — no_client_certificate / 321231331023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023321022100121-1121010331001222-2332030010211121-0313333011311331-0032303233222022-0322203232020313-0211111023112003-3011031303123030"></a>

## Next pages — no_client_certificate / 321231331023 / 4

- [tls_parameters](data-sources--virtual_host--reference--group-003.md#canonical-1121021021313110-0023032010022012-3211322221232312-3332112102011331-0313200121313013-3111023011000333-1001310231100120-0312213222131001)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1002302233122000-1323022011111121-3132010100020022-1310010213200321-0312233032021223-1020313001300323-0020033031030302-1102203102302301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231111102103000-2203112333000221-0322033111120212-1201221013221332-2310303022031212-2011131120120313-1132220212230030-0113002120011122"></a>

## user_identification — user_identification / 313231233303 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- user_identification

<a id="canonical-3332310012322303-2301321002210033-3200311030002213-3100332131002300-3021122202111201-1333223133120302-0111201333302303-2100033020001110"></a>

Type: `"list"`. Computed.

Reference to user\_identification object. The rules in the user\_identification object are evaluated
to determine the user identifier to be rate limited.

Upstream description:

A reference to user\_identification object. The rules in the user\_identification object are
evaluated to determine the user identifier to be rate limited.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2201330131112031-3031033130000203-3101022211033212-0012131011000132-1232033333230133-1322211110213122-1233321212220030-0032130230322321"></a>

## Direct properties — user_identification / 313231233303 / 3

<a id="canonical-1211032333222130-1113032311112331-1322112023200233-2221210102302032-0111132202230233-0313311101113121-2231112101110013-2110120213032300"></a>

<a id="canonical-1023233111003003-2333132331232121-0102000032233000-1332033333210033-1022233221021311-3113002002300300-1323023100231201-0120303212032001"></a>

## kind property — user_identification / 313231233303 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1302023000122200-1233200213212012-0310002020132131-3212003200033311-1103203002212221-0012230220321313-0101102200232223-2111101323101303"></a>

<a id="canonical-3111332323111230-1232020330230002-2123330121110130-2000032302311331-1313221020000222-0211022202303021-0133321231021023-3211010021031000"></a>

## name property — user_identification / 313231233303 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1222313323133111-2012203111100120-2111302113221303-3130221231213012-0033000012131020-0301120010010010-1321330030230133-0220333320021122"></a>

<a id="canonical-0001230220010301-1310212332130303-1123133323021110-3003323022023101-0231331312220233-1230320033103311-1223032132011001-0012203130323010"></a>

## namespace property — user_identification / 313231233303 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1131031121132022-3200223321130120-3101323011322120-1220023033333233-3030131103211311-0001230112030013-0203121130311021-2311300302211020"></a>

<a id="canonical-1210130233012222-0222120110111202-0021203311122301-3302030122120210-1202310110221023-2022133321300201-1001100321221230-2013220100103100"></a>

## tenant property — user_identification / 313231233303 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0233310231303031-2221213101031123-0033311303301102-1023330030323230-2012313013130211-2322013030310211-2332122002033020-2220002133001131"></a>

<a id="canonical-2322213210313030-3230320002210010-2031332022322220-3310333132320022-3323321310233132-2030133013100121-2201032003031023-0103230122033033"></a>

## uid property — user_identification / 313231233303 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2131021133011120-2131333320332301-2221312000300122-2023301030001020-3001310230100113-2203121003130330-2120220221131313-0033110310202030"></a>

## Next pages — user_identification / 313231233303 / 9

- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331220111013021-1000033011133212-0000002130333012-2200101323102230-1212233033222220-3212101021123230-2011313022033200-0232111101122331"></a>

## waf_type — waf_type / 201330131033 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- waf_type

<a id="canonical-0212131312112320-3010013102333133-1003210023231021-0022011032133223-0310111333100300-0332102210001333-3113231121103003-0013001311203303"></a>

Type: `"single"`. Computed.

WAF instance will be pointing to an app\_firewall object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

<a id="canonical-2000232200002301-2311202112332212-0122120102110030-3313223323222201-1013003302233022-1333323010032222-2302120001103331-3230111320122133"></a>

## Direct properties — waf_type / 201330131033 / 3

- [app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-1333300322323113-1232131101000031-2233210011203232-3322023300230033-1233133023310221-1233001022120331-2003112011233330-2313020200000013): complete subsection reference.

- [disable_waf](data-sources--virtual_host--reference--group-003.md#canonical-2230010330302310-1312132323300100-3023320210221023-2030200010011000-1310013232033111-0133103323302222-1021312002012220-2103031000003013): complete subsection reference.

- [inherit_waf](data-sources--virtual_host--reference--group-003.md#canonical-3111320320023322-3033330033011220-2031221222111000-3211122202020221-2303103121031102-0213033333303222-0321322313130302-2022131333210103): complete subsection reference.

<a id="canonical-3213320133330322-3013031210030200-3122012110231123-2322013221130221-2011011132120132-3102131133130300-0033333302131113-3322333220300121"></a>

## Next pages — waf_type / 201330131033 / 4

- [waf_type.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-1333300322323113-1232131101000031-2233210011203232-3322023300230033-1233133023310221-1233001022120331-2003112011233330-2313020200000013)
- [waf_type.disable_waf](data-sources--virtual_host--reference--group-003.md#canonical-2230010330302310-1312132323300100-3023320210221023-2030200010011000-1310013232033111-0133103323302222-1021312002012220-2103031000003013)
- [waf_type.inherit_waf](data-sources--virtual_host--reference--group-003.md#canonical-3111320320023322-3033330033011220-2031221222111000-3211122202020221-2303103121031102-0213033333303222-0321322313130302-2022131333210103)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-1333300322323113-1232131101000031-2233210011203232-3322023300230033-1233133023310221-1233001022120331-2003112011233330-2313020200000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222022012102102-3310130230221003-1103130230102223-0203210130303230-3231301030101110-1122023023213323-1130311020203203-1103120221133101"></a>

## waf_type.app_firewall — app_firewall / 333320022102 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- waf_type.app_firewall

<a id="canonical-3000222221111110-0300100323231110-3213130110203022-3102100032020310-1302332210112312-3031031231222333-0210011003021121-0131313212222203"></a>

Type: `"single"`. Computed.

List of references to the app\_firewall configuration objects.

Upstream description:

A list of references to the app\_firewall configuration objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1033033203212313-2003321031301311-0133003022022030-1320311122302211-0300323211103213-2310333332201232-1012120113123232-3201131303103321"></a>

## Direct properties — app_firewall / 333320022102 / 3

- [app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-3300211200203031-0102030220211032-2101332013321022-2020233013011313-3111012122302233-3321322111201122-3301232102311203-1002222023222302): complete subsection reference.

<a id="canonical-1330133131122332-3211120023210002-0330121001222303-2000221301220030-1112312112203102-0300330221130300-0133032313223132-1113302003221310"></a>

## Next pages — app_firewall / 333320022102 / 4

- [waf_type.app_firewall.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-3300211200203031-0102030220211032-2101332013321022-2020233013011313-3111012122302233-3321322111201122-3301232102311203-1002222023222302)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3300211200203031-0102030220211032-2101332013321022-2020233013011313-3111012122302233-3321322111201122-3301232102311203-1002222023222302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313312001123103-0212312030231133-1133321222322032-2301313200203331-0000131313331022-2223130110311103-1300011232323033-2123133310103110"></a>

## waf_type.app_firewall.app_firewall — app_firewall / 120100300320 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- [waf_type.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-1333300322323113-1232131101000031-2233210011203232-3322023300230033-1233133023310221-1233001022120331-2003112011233330-2313020200000013)
- waf_type.app_firewall.app_firewall

<a id="canonical-3031320302232103-0313103102301121-3310310322030220-1331222233331322-3323312112300120-0132330303223100-3230033103103121-1302130133000100"></a>

Type: `"list"`. Computed.

References to an Application Firewall configuration object.

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
    "ves.io.schema.rules.repeated.num_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  }
}
```

<a id="canonical-2103012331212031-1333130101300122-1012310113022322-0201221003301311-3011000332212230-1001131232323223-3033101231302332-1212232213201300"></a>

## Direct properties — app_firewall / 120100300320 / 3

<a id="canonical-2122301313000122-3213122132202331-0303320313121310-1013233322231122-2033012022001310-1012311223231011-1122320130002211-0023012113203300"></a>

<a id="canonical-2330213213302112-2223011003232111-1321012213023020-2212203113200002-2122002320332113-1202123203023122-2132300003313031-0331011320132130"></a>

## kind property — app_firewall / 120100300320 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2210021022020201-1101003131010303-1033011311003033-3313112311133020-2212001121211021-1100213001032311-3233011301030120-1112313221311113"></a>

<a id="canonical-1310130333321133-0221032102000123-1332032020232033-2313123112111013-2231232323313122-3322303213023202-2213323330322002-3233230210012101"></a>

## name property — app_firewall / 120100300320 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3033002313132221-2232211031303012-3110113302303333-3021112021230110-3321323113210103-0223221211001323-0302100133213212-1021312212203133"></a>

<a id="canonical-0113012131012321-2023001120121203-1110300100120311-2213303031113133-3020123100232210-0022031112013202-1213212001230303-2220332231020130"></a>

## namespace property — app_firewall / 120100300320 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3331131032302210-1323121002333201-1230311011223103-0121133132223223-0130311111330201-0301032302331131-3303320103113200-0311132221113133"></a>

<a id="canonical-0230213200020303-1331112321313003-2002233322223310-1330032133130213-0003211211301301-3123303032310033-1211323121133002-2223221120123112"></a>

## tenant property — app_firewall / 120100300320 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1013211020013021-0011321100123322-0231222213313331-2021120130333111-3211303102212321-2133230011012132-0202012022103100-1012310300010120"></a>

<a id="canonical-1310220222113232-3203210201311133-0121320322003101-3130130111033200-3112130220132210-0121333210213101-2332112312220323-2200122320113111"></a>

## uid property — app_firewall / 120100300320 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1111123033230230-3021012012121100-0113020301201313-3301131323123221-0330131320322101-2303002321033311-0112122001130022-0313332232201112"></a>

## Next pages — app_firewall / 120100300320 / 9

- [waf_type.app_firewall](data-sources--virtual_host--reference--group-003.md#canonical-1333300322323113-1232131101000031-2233210011203232-3322023300230033-1233133023310221-1233001022120331-2003112011233330-2313020200000013)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-2230010330302310-1312132323300100-3023320210221023-2030200010011000-1310013232033111-0133103323302222-1021312002012220-2103031000003013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031313302021313-1023021213121220-1321301013130131-0222321321112113-3321203313212323-0122123302002032-3001303011203310-1201103323323332"></a>

## waf_type.disable_waf — disable_waf / 211011102233 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- waf_type.disable_waf

<a id="canonical-0013120130201131-1002011002120203-2223320033320023-3202121202131112-3200222111030331-3131000003002201-2333113123200223-2321301120031311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable waf.

Upstream description:

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

<a id="canonical-0102111123221331-2100232311310132-3102032310313131-0012032201130312-0011230213333031-0000230123322103-1032020221323100-0203130100100112"></a>

## Direct properties — disable_waf / 211011102233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022301202313232-3312033230001212-1330333010211313-1311120121000112-3222101221202002-2211303313322000-3132100300213322-1303322300333303"></a>

## Next pages — disable_waf / 211011102233 / 4

- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)

<a id="canonical-3111320320023322-3033330033011220-2031221222111000-3211122202020221-2303103121031102-0213033333303222-0321322313130302-2022131333210103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010111000003122-0133133210033110-3300313130332121-2002220311311003-0123023101101222-2023002231001113-0300131021312310-1003203103332223"></a>

## waf_type.inherit_waf — inherit_waf / 110010032201 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Property reference](data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- waf_type.inherit_waf

<a id="canonical-3010301221131302-2113032200130303-1223220212210300-0101211011121200-2020201320122112-3013320110313203-2213221213333033-2120022000130002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherit waf.

Upstream description:

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

<a id="canonical-1201322311133302-0332003210221323-3122233121023231-3021020212210321-0103123121210333-0101012302003113-1121330000202112-1311213032120130"></a>

## Direct properties — inherit_waf / 110010032201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001120323021003-1301311230031203-0011113201133020-1012033331022120-0011230100210130-3001333210132120-1131220101201222-0320213322011303"></a>

## Next pages — inherit_waf / 110010032201 / 4

- [waf_type](data-sources--virtual_host--reference--group-003.md#canonical-3313303333213033-3223023131200111-1302021310020211-2020230220202121-2013001333301310-3033012222123323-3123120210220011-2132111220233001)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
