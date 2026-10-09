---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3020133322310122-3120022303032302-1333013033231033-0203223021012000-1333112131131133-1101022331131132-0202323320200021-2220033102132032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.common_buffering` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.common_buffering

<a id="canonical-3201331033013213-1132133200002311-2213131321030011-2213222003212323-0000231132330102-0313011232100201-2323110323302011-1012311002020020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for common buffering.

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
common_buffering = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220202221213121-2003330033111010-3112031032321133-0331322330011032-2100203210203032-1312312203321022-3331311132203101-1323310311023321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.common_hash_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.common_hash_policy

<a id="canonical-0023222113102030-2133033113012302-0333322313321103-3221333321200100-1232100113213201-1011310221311332-1330002200013002-3032002111312231"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
common_hash_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222022011013211-0231322200331030-3003132300320323-0223133302201212-3011220203110302-1311320103313321-2022323330221202-0131123122131320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.cors_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.cors_policy

<a id="canonical-2023331023001023-2020203132023101-2331310220020032-0123020311022201-2110332123100303-3223032113133301-3030112112230200-3020302212332000"></a>

Type: `"object"`. single nested block, Optional.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel macOS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.HTML Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
macOS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8 Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive
Origin: http&#58;//foo.example Access-Control-Request-Method: POST Access-Control-Request-Headers:
X-PINGOTHER, Content-Type

HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http&#58;//foo.example Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary:
Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive.

Receipt-pinned upstream constraints:

```json
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
cors_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103133001033303-2311032132311131-2200232202330122-3131300230233322-0133233023020232-0021221101101111-1021103210222322-2121100212222323"></a>

### Direct properties for `routes.simple_route.advanced_options.cors_policy`

<a id="canonical-0222001133330130-3202233330130000-3303120213020120-1201312231203313-3322233010300130-0112031233222202-2110311230110102-1130333311230223"></a>

#### `routes.simple_route.advanced_options.cors_policy.allow_credentials` property

Type: `"bool"`. Optional.

Specifies whether the resource allows credentials.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3330102120203133-0130302222230202-2210300303011101-3130210131211322-1233011230103003-1212210133212021-3301303002131032-1020310002013111"></a>

<a id="canonical-3003310122030303-1011311030032320-0032222213302012-2112233100311332-2330213133303012-3132200322100313-1122221322112200-3311012113100200"></a>

#### `routes.simple_route.advanced_options.cors_policy.allow_headers` property

Type: `"string"`. Optional.

Specifies the content for the access-control-allow-headers header.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3121313321023303-0203101121221200-3232021031311223-1231202313101120-0200310023233100-1321200121210031-2110031032101102-0100233201122213"></a>

<a id="canonical-0230011133100001-1210011102221133-0133020200001203-3031212222330301-0312102301330012-0303221023303102-1023232101302021-3201011220003100"></a>

#### `routes.simple_route.advanced_options.cors_policy.allow_methods` property

Type: `"string"`. Optional.

Specifies the content for the access-control-allow-methods header.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-0200322113021211-2201113221130210-1102113331011023-3021013102111312-0120213023333232-0022011101200121-3133102110303010-3002323131233233"></a>

<a id="canonical-0203100221310222-0231223021021223-3130012122223110-2323103311103232-3030131012110302-1013231323320000-0312223021203131-3122020012212132"></a>

#### `routes.simple_route.advanced_options.cors_policy.allow_origin` property

Type: `["list", "string"]`. Optional.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0232110312022111-1331021301213022-0331131320102020-0011213311222311-1200120201230233-3220200201213200-1111312202000120-3300120030101313"></a>

<a id="canonical-3112322323332120-3201303113330320-0300330102113333-2031010230232212-1031102210022202-1122023213100320-3312000300300330-0321331133112003"></a>

#### `routes.simple_route.advanced_options.cors_policy.allow_origin_regex` property

Type: `["list", "string"]`. Optional.

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2321130231121322-0123311030333212-0210001130213311-0210110222102012-1111222002110302-2012203220332003-1120132301003310-1310021000002000"></a>

<a id="canonical-3220100113130021-1203001020003121-3120002131223030-1020223111211001-0231323230122013-3302302300203311-0030012233120131-2222213200022000"></a>

#### `routes.simple_route.advanced_options.cors_policy.disabled` property

Type: `"bool"`. Optional.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1231130321323201-0022232333300330-0300000312230221-0230322023133332-1100203332013130-1110031230231303-2202313012133100-3233210002300101"></a>

<a id="canonical-0231323002010212-0303022223211110-2313311100311312-0032213331322032-1220132202310022-1233201013011202-1202102100201033-2111332000213233"></a>

#### `routes.simple_route.advanced_options.cors_policy.expose_headers` property

Type: `"string"`. Optional.

Specifies the content for the access-control-expose-headers header.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311210313132213-3210132133203231-1032111210313312-3102310001303111-3133121130031012-1213201310133232-0103023321132210-3010313111030113"></a>

<a id="canonical-1011301312020212-1203111121310333-3331033003222231-3001223333320023-0311120231133200-0011032332100103-3103121311301313-0111211000311302"></a>

#### `routes.simple_route.advanced_options.cors_policy.maximum_age` property

Type: `"number"`. Optional.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

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
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": -1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  }
}
```

<a id="canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.csrf_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.csrf_policy

<a id="canonical-3210133010213111-1312130212033003-3022113113101331-1010102020220033-1010220233110232-3223132020131032-3310132330002033-2113303103211212"></a>

Type: `"object"`. single nested block, Optional.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302332133031330-3310210020103232-2203033023231331-3331100010012101-3213023113221102-0210332210200212-3120322231121132-3001200011110231"></a>

### Direct properties for `routes.simple_route.advanced_options.csrf_policy`

- [all_load_balancer_domains](resources--http_loadbalancer--reference--group-026.md#canonical-3102211303130002-0311302231123112-3021133021121320-2300112000101322-0022213122203220-3321203021313221-1030122323103302-1211130133103221): complete subsection reference.

- [custom_domain_list](resources--http_loadbalancer--reference--group-026.md#canonical-2232023010120200-1330112210123120-2122301123331132-3212302133120122-0002301100303000-2120003300300223-2211300033013032-0310121310232122): complete subsection reference.

- [disabled](resources--http_loadbalancer--reference--group-026.md#canonical-2213021333110233-0321023022220030-3221132311032333-0103000100120312-3322200031221311-2301022200220203-1302310322301221-1201100321332030): complete subsection reference.

<a id="canonical-3102211303130002-0311302231123112-3021133021121320-2300112000101322-0022213122203220-3321203021313221-1030122323103302-1211130133103221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains

<a id="canonical-3330010002123331-0022223323002002-3212032222132113-3133203012310321-3231113310301302-2003312122222010-0122012202020021-0030113232111333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all load balancer domains.

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
all_load_balancer_domains = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232023010120200-1330112210123120-2122301123331132-3212302133120122-0002301100303000-2120003300300223-2211300033013032-0310121310232122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.csrf_policy.custom_domain_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- routes.simple_route.advanced_options.csrf_policy.custom_domain_list

<a id="canonical-2111121030123202-0103312021020101-3012203203132300-0111131221313003-1323333033023330-0212311110023113-2011120333300333-3001022323033110"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

Receipt-pinned upstream constraints:

```json
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
custom_domain_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023230132321033-0101132302132103-0132322300300001-1302011001111303-1023311003231001-2123130210122232-3112110031123300-1223321232322103"></a>

### Direct properties for `routes.simple_route.advanced_options.csrf_policy.custom_domain_list`

<a id="canonical-2302133323021231-1023123012230012-3220033331021322-2033012111133121-3110212122013320-1320011012302003-3303101220133023-3201221033033110"></a>

#### `routes.simple_route.advanced_options.csrf_policy.custom_domain_list.domains` property

Type: `["list", "string"]`. Optional.

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2213021333110233-0321023022220030-3221132311032333-0103000100120312-3322200031221311-2301022200220203-1302310322301221-1201100321332030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.csrf_policy.disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- routes.simple_route.advanced_options.csrf_policy.disabled

<a id="canonical-2200132233200213-3323322130131233-2112201110103322-0331210302310203-0321233032120302-3311231133210010-0111212232201221-1300003232330211"></a>

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

<a id="canonical-3013002323211230-2000321010322031-0202220333302003-3320223120203023-3311023301310130-1123220212320222-2303233031330210-0222030031022311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.default_retry_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.default_retry_policy

<a id="canonical-3103011220122302-3321311323030103-1032101303211301-3103011110101031-0300022110300232-0103123021013002-3201132121210023-2302110330111320"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
default_retry_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223210220132003-1032031311131022-0221333010313232-3300232111321333-3020023100102303-3000311333203311-2323022130322120-1301313300201232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.disable_mirroring` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_mirroring

<a id="canonical-1021031031132110-2033021030110230-0231111133321222-3131103331311021-3332130220020323-0222201323200231-0312010103333220-0200212111112210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable mirroring.

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
disable_mirroring = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101222121331231-0002101210102120-1000231102102220-2110231230123212-2111001231312132-2121211202233201-1300220103223003-2023201222120123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.disable_prefix_rewrite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_prefix_rewrite

<a id="canonical-1222222120001310-3020213133313220-0021000331131010-3113133202110330-1221201321122132-2331133300020222-2220221223203102-2221233223111201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable prefix rewrite.

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
disable_prefix_rewrite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333112331031332-3110230102112221-1123102212221222-1320333331030110-2102300032332222-3010102120200200-0200212202113120-3221312122220012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.disable_spdy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_spdy

<a id="canonical-0301033111111300-1200133121222020-0233012113003210-0010312030213210-0330310232113003-0103132321001201-1102321013201122-1022131111003311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable spdy.

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
disable_spdy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012322210102001-1000202101032101-0301012012020133-2222012131120302-2033130323220023-0131111302321232-1103220102022322-0130330011322111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.disable_waf` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_waf

<a id="canonical-3303013011100012-0200023313012302-0002023003301032-3111022310002230-0112220121131100-1213103021322322-2210322122321100-3322103003022132"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_waf = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123013100222203-3012313301032221-2333300120200121-2031023303323312-2330102330202121-3313212031211202-1300121232122232-3230133031002030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.disable_web_socket_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_web_socket_config

<a id="canonical-0013020202203110-3312230002300003-2330231002010231-1111020021200000-1032312100120321-0201310310023232-3223000123310102-0213233200030033"></a>

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
disable_web_socket_config = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013100112201010-3320010213003210-2110103020232131-1233302023023333-1320111303223010-1220201100133030-2322330330102010-1111303301333333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.do_not_retract_cluster` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.do_not_retract_cluster

<a id="canonical-3032010210001000-0132301310112113-3231312222323112-0320100231221110-1313022301112231-3320211323123121-3202333203021331-1011211031002103"></a>

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
do_not_retract_cluster = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220221103013030-1033023211020212-2313000220302030-0112212123203222-2223110003210120-1330110301122333-0120013212320200-2011130012212031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.enable_spdy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.enable_spdy

<a id="canonical-1012101022332202-2310312301302112-1301110311103223-0312211212223123-0320331111303232-3230130213121033-0020320312112021-2022113200301210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable spdy.

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
enable_spdy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123231212310000-0130220212030233-2332222131111223-3330003103011023-0211122003332012-0200233023011101-3232021110002213-0302122311123131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.endpoint_subsets

<a id="canonical-1221121322322213-0221232101101003-3131212031322320-0000323232312133-3201011032032323-3200100001302133-1113231122011333-1332133101210333"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101301002211102-3101130203012322-2333001122113100-1010021103332111-0203220210230210-1230033102313122-1102112111223121-2210120101222000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection

<a id="canonical-1122201311111012-3332331300213133-0331001223212103-2210100300201122-0021211020120201-0011310112302132-0031001122302210-3233322011310332"></a>

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
inherited_bot_defense_javascript_injection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213201113012322-1102122130221322-3332120321331122-3312320023133031-1013323333130130-0033111303102122-3133232312300201-0023120200310323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.inherited_waf` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.inherited_waf

<a id="canonical-1213202313123210-3132132322133331-1121203130032333-1001231202200300-2211032020131300-0311123000323133-0312302203222331-0301233211313322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherited waf.

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
inherited_waf = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002012001013000-3310311312101011-3313033113321211-2131033121023113-1002001131311300-2001101231211002-3100031020332322-2132012212030231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.inherited_waf_exclusion` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.inherited_waf_exclusion

<a id="canonical-3223112102232310-3321103232021123-3331233121022220-0013021212303220-3230121322333000-2030221101321013-1031220232231203-0200203001323331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherited waf exclusion.

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
inherited_waf_exclusion = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.mirror_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.mirror_policy

<a id="canonical-3203310310232132-1110311230130323-3100032331000310-1332003132203113-2013321032131232-1221311332102131-1210021022223002-1210211131322011"></a>

Type: `"object"`. single nested block, Optional.

MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is
'fire and forget', meaning it will not wait for the shadow origin pool to respond before returning
the response from the primary origin pool. All normal statistics are collected for the shadow
origin..

Additional upstream details:

The approach used is "fire and forget", meaning it will not wait for the shadow origin pool to
respond before returning the response from the primary origin pool. All normal statistics are
collected for the shadow origin pool making this feature useful for testing and troubleshooting.

Receipt-pinned upstream constraints:

```json
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
mirror_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100332221110030-3001201303011101-2231212231032300-0232132330301212-2001322222130131-3333201200212012-3131303302000201-1233321032303130"></a>

### Direct properties for `routes.simple_route.advanced_options.mirror_policy`

- [origin_pool](resources--http_loadbalancer--reference--group-026.md#canonical-1103223210103233-3323023333223103-3202033211110230-3123111220212303-1220320130301312-2303110030230323-3323221001100011-2221131001203201): complete subsection reference.

- [percent](resources--http_loadbalancer--reference--group-026.md#canonical-2313121201002301-1112103031031030-2213123011002321-0330220012031022-3302303212211102-3121203313102011-2121121230313002-3102200032303031): complete subsection reference.

<a id="canonical-1103223210103233-3323023333223103-3202033211110230-3123111220212303-1220320130301312-2303110030230323-3323221001100011-2221131001203201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.mirror_policy.origin_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-026.md#canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023)
- routes.simple_route.advanced_options.mirror_policy.origin_pool

<a id="canonical-0330312013233303-3302331312302211-3130203301310003-3323332322132021-3113122133301332-3100000312203131-1220023302033213-1312100321202331"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
origin_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310301012321302-3012000020211213-3001323230130101-0023110001010020-1020100213112230-0330310133132321-3110231003032231-0020233233010001"></a>

### Direct properties for `routes.simple_route.advanced_options.mirror_policy.origin_pool`

<a id="canonical-3211113123010223-1023002120123110-2303113110200102-2013031032121301-0323023200310200-3013001323231213-0011303123330111-1222312130011333"></a>

#### `routes.simple_route.advanced_options.mirror_policy.origin_pool.name` property

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

<a id="canonical-1101100332220010-1331313023021002-2302102000023023-1130311310012022-0033212210212011-0201013130300232-3233203310222213-3011030320301022"></a>

<a id="canonical-3233013230300010-0231122103112002-2032232212110212-2330322210032221-2203001021201130-0120111310001002-0301100221220311-0300003210020201"></a>

#### `routes.simple_route.advanced_options.mirror_policy.origin_pool.namespace` property

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

<a id="canonical-3100000232102232-1213203100330213-1030310333213110-2223112310100221-2332311033311321-2133220321111002-3133233001111313-0232003011123032"></a>

<a id="canonical-0302003331102210-1031302132212133-0210320322331310-1203223233300232-1130200132102133-0323020301210110-0333333123020333-1113013302132012"></a>

#### `routes.simple_route.advanced_options.mirror_policy.origin_pool.tenant` property

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

<a id="canonical-2313121201002301-1112103031031030-2213123011002321-0330220012031022-3302303212211102-3121203313102011-2121121230313002-3102200032303031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.mirror_policy.percent` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-026.md#canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023)
- routes.simple_route.advanced_options.mirror_policy.percent

<a id="canonical-3333100001321201-2123213121321233-0212200010330223-1300330210212023-2031031221302133-2331330023102203-0332000030211210-2000032223012111"></a>

Type: `"object"`. single nested block, Optional.

Fraction used where sampling percentages are needed. Example sampled requests.

Receipt-pinned upstream constraints:

```json
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
percent {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211221302013022-0010303233000013-3213203023312110-2331011030110300-0122231003123101-0221311031203020-2213013002103310-0220300220313203"></a>

### Direct properties for `routes.simple_route.advanced_options.mirror_policy.percent`

<a id="canonical-1002101303031213-0323213000330212-3331000001300302-2212110102211110-0123023102130113-2021212232100200-3322212213202130-1020301213033210"></a>

#### `routes.simple_route.advanced_options.mirror_policy.percent.denominator` property

Type: `"string"`. Optional.

\[Enum: HUNDRED|TEN\_THOUSAND|MILLION\] Denominator used in fraction where sampling percentages are
needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use
million as denominator. Possible values are \`HUNDRED\`, \`TEN\_THOUSAND\`, \`MILLION\`. Defaults to
\`HUNDRED\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "HUNDRED",
  "enum": [
    "HUNDRED",
    "TEN_THOUSAND",
    "MILLION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3011323121203331-1131130302322021-3101022232210103-0111322310012121-1201100121201102-2031000312000101-0232010002110301-1102212200002221"></a>

<a id="canonical-3022203002213123-0303331021233320-2312120103223133-2132120323211000-0130211003032312-1330313002023310-0010323321110310-3230133212323313"></a>

#### `routes.simple_route.advanced_options.mirror_policy.percent.numerator` property

Type: `"number"`. Optional.

Sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000.

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

<a id="canonical-2220333121303102-3132022333332033-2111322310103130-0230222232313212-1102113200030202-3202032303023333-2120330013220013-0120312220300111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.no_retry_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.no_retry_policy

<a id="canonical-3223110100331210-0311003333323320-0032321010031313-2301320120121302-0333220301202313-2203311000223033-3333032123130023-3022213322223331"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_retry_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320010111202212-1023320033013103-0030311122023131-1323130311311223-3130123101023222-0210202202123330-0003303022231213-3303032312010332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.regex_rewrite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.regex_rewrite

<a id="canonical-1332213201131233-1311103113232313-0003110113230122-0032003101331030-1322220113200113-0122021121110230-1031201222131010-3132223232111312"></a>

Type: `"object"`. single nested block, Optional.

RegexMatchRewrite describes how to match a string and then produce a new string using a regular
expression and a substitution string.

Receipt-pinned upstream constraints:

```json
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
regex_rewrite {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310131301033321-3100012113110200-0203100220010203-0103101010121210-3300130331322132-1020021213013303-3123223111231110-2022303212131032"></a>

### Direct properties for `routes.simple_route.advanced_options.regex_rewrite`

<a id="canonical-2312210120331321-1032020300102000-3332010231222123-0100121301312303-3013033311120332-0102013213132310-0313331222302330-2023212230212100"></a>

#### `routes.simple_route.advanced_options.regex_rewrite.pattern` property

Type: `"string"`. Optional.

The regular expression used to find portions of a string that should be replaced.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3100331213020210-3122202001023210-0132103301222210-2000123200322211-0231021313202230-1131000122220200-1332122311032222-0303122322211300"></a>

<a id="canonical-1103023313300101-1133133011221330-3121121322111110-2100302312000330-3030232303331310-1103110231033312-3120213322120031-2203230220303133"></a>

#### `routes.simple_route.advanced_options.regex_rewrite.substitution` property

Type: `"string"`. Optional.

The string that should be substituted into matching portions of the subject string during a
substitution operation to produce a new string.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.request_cookies_to_add

<a id="canonical-2023123122330012-1011112221023022-2310012330221032-2230313222233031-0033302110203301-3331032233211112-2302113000113003-3330003113232023"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022200000101330-2030003120211121-0101301210330201-0233102013101111-3330323233323310-3313001111230230-3132222101302102-2202230331200110"></a>

### Direct properties for `routes.simple_route.advanced_options.request_cookies_to_add`

<a id="canonical-1131032303302203-1213301003023000-2111120020000022-0210222233211112-2310210113113112-0112122003003220-1032131033121202-1211101022032010"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0012300203220231-3111002121230131-3311031230030213-2022110323132220-2110100103123333-1030332010102111-3101212200322201-1113032212003000"></a>

<a id="canonical-0221022211010232-0232210033000222-2112133100332110-2200312303020133-0100201223120102-3111111031022302-3210132121000110-3003000301020233"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.overwrite` property

Type: `"bool"`. Optional.

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

- [secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121): complete subsection reference.

<a id="canonical-2012233030000330-3220233123110232-2333031310332021-1133012323021311-1301222201131211-1301211302100110-3313123010113220-1103211303333312"></a>

<a id="canonical-1133100103113331-3301200103200002-2101033012101123-3002330033101020-0023313210212113-1010101300102212-1001333321021233-3122302312022213"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value

<a id="canonical-2221112133101200-2321000221313010-3233000123000132-3333220220332012-2220301120332123-0303030011102322-0221103012033122-3313033011201220"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120133230213313-0101030300023122-0121011332003121-3221213100211133-2322330001202303-1020313230020312-1322211310013310-3001031033220331"></a>

### Direct properties for `routes.simple_route.advanced_options.request_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-0212231213230312-3102112110221211-1100022003031232-3320033210021133-3132120212202021-1320121122013221-2330103113101111-0102012112312132): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-2100100321020011-0120311233112012-3121323322113103-2310100301231321-3030210112013001-2301002311012321-0210103011301002-1211023132203232): complete subsection reference.

<a id="canonical-0212231213230312-3102112110221211-1100022003031232-3320033210021133-3132120212202021-1320121122013221-2330103113101111-0102012112312132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-2321123312320130-2111000211032220-1333121322132313-1131131203330213-2101113330320112-0233030331133212-3213000133011220-1311302102031003"></a>

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

<a id="canonical-3200202220202113-1233210102200032-0303032123211230-0202111233000330-3233000010032113-2202001232222102-2310112122300132-2203201132101302"></a>

### Direct properties for `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1302213220323032-2213302113221312-1100321203333312-0211030210031110-1133330101223231-1001331110210202-0232121203212031-1331123031220033"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3031231230232321-2202211213301132-2010033002031122-1121212201212003-0321300232233112-2103202301121321-3231211211211130-3211031200233223"></a>

<a id="canonical-0000000100231011-2301212013033322-3103002001323233-0121332300321130-3233211120021030-0133200101222223-3312130210323111-2010022220200232"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2231012030000121-1030312220211210-0102130001311301-2120021100022123-2023000313121213-2003320333302013-2203312112021223-2213121000121211"></a>

<a id="canonical-0012223231311133-0333210333132131-1120323223221312-2113220023112201-0113233223033021-2302022002001110-2112130202121022-0211310221230221"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2100100321020011-0120311233112012-3121323322113103-2310100301231321-3030210112013001-2301002311012321-0210103011301002-1211023132203232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2110221023333101-2333331023011323-3300301203303210-1212033110112011-1020130110030311-1033313332230122-2311312010322022-0013101011333132"></a>

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

<a id="canonical-0300110310302021-1310323313121210-0120011111012211-1332132311312110-3233030302033030-1010003002131212-0303030201203101-1010103032310320"></a>

### Direct properties for `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-1203202002331030-1302322301331032-1232033122211020-3313102011212222-0130033230303220-1200233231100122-2101310011003312-0122111123003003"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0320030231133333-1122012032301220-1000332032033100-1030203132012230-1212130021001100-3110022110103210-3203210323000102-0003130022123020"></a>

<a id="canonical-0111102033300320-2001111311230313-0021030220122302-3020222222002003-2201100300312200-2102312211023323-3303232030310333-1000122123233322"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info.url` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_headers_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.request_headers_to_add

<a id="canonical-2022011310031200-0013303231220310-0010311312321030-2121221330003233-2132210232231103-0223031231323331-0303001302312200-3001132120123023"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301023222200221-1213303102301331-3110203022303312-3013221121210030-1322012123013133-0120221102313122-1303003213002132-0220121300321021"></a>

### Direct properties for `routes.simple_route.advanced_options.request_headers_to_add`

<a id="canonical-2223222320001202-0311023120022130-3202111032220231-0010132132011203-0132331000121303-3301001213331220-2211111301330213-3220222211200023"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.append` property

Type: `"bool"`. Optional.

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

<a id="canonical-3210020222301333-3302221031223113-1232211003232012-3102323322000220-3332021022213113-0221303313110200-2002022111210112-3110322113303120"></a>

<a id="canonical-1300302012033230-3300001333212232-3213302012011221-0231111020002122-1322321122020023-2313020112200010-0001130322130030-2022330231302003"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.name` property

Type: `"string"`. Optional.

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

- [secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120): complete subsection reference.

<a id="canonical-2311012021002332-2130231220030200-1113301213210230-3202032333030013-2111111303033303-0232310320232303-2113012011312131-0010212130033231"></a>

<a id="canonical-1322323303122211-1111330021310333-1320313002003232-0310323023021133-2321302312022231-2312233231003222-2310303312000323-2100323220331030"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.value` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value

<a id="canonical-0103233013032131-3102300312132100-3232201300122033-0013003323321312-0223302212101113-1333111120303012-2231331233310322-1301201022002320"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101133230113233-0302320303220001-1210002210013132-2321123100302033-2102113301222213-2200323210030200-0133323132231102-2333300103101013"></a>

### Direct properties for `routes.simple_route.advanced_options.request_headers_to_add.secret_value`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-2010300000023322-0213320200330302-1131301130213121-2121133311202221-1220022210031330-0132020101323311-0330010132222303-1023120132222101): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-026.md#canonical-3333203031110332-1013301011031232-1122112311011033-3030310211111231-2320012202112313-1233233131230221-0102310210020010-1113102130122120): complete subsection reference.

<a id="canonical-2010300000023322-0213320200330302-1131301130213121-2121133311202221-1220022210031330-0132020101323311-0330010132222303-1023120132222101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0331110020333233-1120231003031120-1001123130102012-1012232131300200-0212201213113003-2220123333233300-3323333101010322-2022232123233221"></a>

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

<a id="canonical-0203231113230310-1331231300211100-3111033032310323-2030300130322012-0301033112202002-2332131332033301-2313100230231023-0222332311001123"></a>

### Direct properties for `routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0012123013230032-1331120113202120-2230313323302000-3321110220111020-1101012322133111-0030131201030303-1113132132321122-1203223202313010"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0033333121002323-3000030022322320-3221033131123210-1230123230120210-0301013312123322-3130300033323003-3000330032310201-0321203011111010"></a>

<a id="canonical-1111100202110020-1202202111210001-0232201330121200-1011110232010222-1312300033333213-0333310231123111-1113013201130123-2221302113202022"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1210311213320121-1332222211221220-0130313311221332-0131023303121203-2200133203202300-3031101120302302-1012000022211133-0113003331232011"></a>

<a id="canonical-0321210121020233-1220231303131033-0302022220232022-0110332021122320-1200212032012203-0230001323221213-2233201200222021-0023310222030021"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3333203031110332-1013301011031232-1122112311011033-3030310211111231-2320012202112313-1233233131230221-0102310210020010-1113102130122120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-1101111300232320-0212013002313233-3301002021322303-0320032200233011-0212202323331121-2031332022330031-1333131002310033-0110031110103331"></a>

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

<a id="canonical-2012121223223110-3220132031031200-3233321011112031-0300023223103003-1232100032201223-3210322331330110-2113212001011003-3313331113332022"></a>

### Direct properties for `routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1310303211021213-1331123110201200-2003332020310022-0300000312033110-1013020023332312-2101320002111002-0003001231231221-1332322230323133"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2230012223313332-1201321331230010-3030332322300321-2313111022202022-2002222333331103-3112033210322113-0010023022003013-3033021003132002"></a>

<a id="canonical-2221323101033333-3311301013113113-0122002111310200-3330033320223033-0322020230032133-0223230113002232-3321101331230120-3003320112003102"></a>

#### `routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info.url` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.response_cookies_to_add

<a id="canonical-2202302222030131-2310030001022313-0132123230331302-2333213311020302-1102132333301132-1301021330220120-2122002133211030-1221303312023122"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113103010301100-2233121001213031-1023011100311220-1110021210202332-2322032220031100-3001030100303301-0312031112232311-0102311101313102"></a>

### Direct properties for `routes.simple_route.advanced_options.response_cookies_to_add`

<a id="canonical-2223300310031313-2033311311202321-3131200332101212-1233110331233031-3021122100332331-3120011311333201-3123133322300223-2231330233332020"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.add_domain` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2311130213121103-2330332121000222-3123021003013100-2021233200323003-3031303223312220-0322212320021001-2021211031211121-1003001212331003"></a>

<a id="canonical-2122100303213132-0122332032230320-1310331203101202-0312033233313103-3122231313310233-1120233102220322-1021221120211201-0100122301003122"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.add_expiry` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](resources--http_loadbalancer--reference--group-026.md#canonical-2313030222311331-0310301120323200-3212110122231121-2111000013022300-2100300222033231-2110122110010311-3021221230021320-0223003232223323): complete subsection reference.

- [add_partitioned](resources--http_loadbalancer--reference--group-026.md#canonical-1323232012020231-1323232033333100-0022200303021233-1132201023301022-1311100300203202-1032201201220220-2212131302003003-2302012011131123): complete subsection reference.

<a id="canonical-2122310233111221-0023021131031103-3202000210110231-2000103032001030-1022202202213222-0220130330120031-0022100101323231-2330213030323230"></a>

<a id="canonical-0201133003000312-3201210031122103-1131212311112333-0120210233222032-0132222100110131-3221330103002301-2302302213200323-2101332312221312"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.add_path` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](resources--http_loadbalancer--reference--group-026.md#canonical-0213002321303031-3200322021113130-3120313220233100-0313331231203332-0301220311230132-0033000312233230-0221310333001103-3012223000332102): complete subsection reference.

- [ignore_domain](resources--http_loadbalancer--reference--group-026.md#canonical-1231230310231311-3031012003130120-0203331001131021-3303121233230332-0000320021103021-2122031210201132-0221212102002000-1020101120000013): complete subsection reference.

- [ignore_expiry](resources--http_loadbalancer--reference--group-027.md#canonical-3000330231110221-3010013022021133-0111331201233132-2023302321123223-1133033122331331-1310300003223121-1301302023322230-0130102220120122): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-027.md#canonical-3033003021303110-2221320232111022-3210213120013012-1202002313131101-0111030110111110-3121230031330010-1320203302233301-2101101020133230): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--reference--group-027.md#canonical-3332100132012110-1201031310330221-3300001231012323-2123023123313110-1320002032122101-2110030113013200-3000320302313022-0313132123300010): complete subsection reference.

- [ignore_partitioned](resources--http_loadbalancer--reference--group-027.md#canonical-0011012023201321-2203022033203003-2130131111312003-1202010321023003-2031320331333121-1133130221220202-3130312111111113-3112011213003301): complete subsection reference.

- [ignore_path](resources--http_loadbalancer--reference--group-027.md#canonical-3130303020100023-1032211010221013-0011013101110003-2100102032021202-2002232131231101-1112031231033121-3231013030113202-1311212333001023): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-027.md#canonical-1113003101311201-1030323301021312-3320212231220300-2131010001120201-1300010123101211-1021011203220320-2223002133203022-3330320022311103): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-027.md#canonical-0121100310301220-3302022132310000-0002023211103222-1331230311320323-0101322131033012-3212012122033212-0102321221032303-1231203220020120): complete subsection reference.

- [ignore_value](resources--http_loadbalancer--reference--group-027.md#canonical-0330032312013200-2213310211110302-3311111113310303-1020202111102302-2331212002020232-1032313002232013-1211000230200012-1203301033112132): complete subsection reference.

<a id="canonical-2331330333332022-3313133111012021-2221303223013301-2013323220323210-3101023033300033-2022202103103233-3120012221103003-3333110313110321"></a>

<a id="canonical-0223222311201000-0333333111113200-1023331221133033-0102201031012330-0313310010310201-2301000000012330-2003003033301233-1101311231102222"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.max_age_value` property

Type: `"number"`. Optional.

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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-0022013303312003-0001333102103022-0033131321220310-2321221001202003-1131023330011210-3133003132313012-3111210123001223-3232103332003212"></a>

<a id="canonical-2313100113221130-0112330031101310-3233333010220102-0101232220013003-3031113011101301-3202313121210312-0132003223023220-2302001110310212"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3323321121121000-0220301213210022-0331212230231233-1203113332232120-1312312323231101-2310231303113030-0321310211202210-2321122012002332"></a>

<a id="canonical-0212222210031222-1212231220011123-2201322302023120-3321230213303303-1100232031030221-1222322103203311-3100300010310112-1122011101232313"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.overwrite` property

Type: `"bool"`. Optional.

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

- [samesite_lax](resources--http_loadbalancer--reference--group-027.md#canonical-2122100130200302-2010033202110221-2101223332201321-1312310321132322-1313033110030110-2323000311312221-3030331110000103-0032112303312122): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-027.md#canonical-3213021033222202-3012330311220203-1130031023201103-1302120001111331-3030131223120232-2212010231121301-0013330311320112-1233330332212232): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-027.md#canonical-0111123301032121-3302113132300132-1031023023123223-1322221302020311-0103222130111210-2302010031223231-2032223221033320-2033012013022313): complete subsection reference.

- [secret_value](resources--http_loadbalancer--reference--group-027.md#canonical-0203212013131210-3112233202311012-1211021022312223-2301111332221021-3300232131111021-2003332002013101-3123100331013113-1132013121310321): complete subsection reference.

<a id="canonical-2021312331201033-3233002311300333-0122132332121013-2302323033100220-0302020022001300-1203033020002030-0312031321222123-2131021211120333"></a>

<a id="canonical-3120323232301130-3203221230333300-1301122110112203-3313102032202233-1320110321201232-3101320322232210-1130230332010301-2232022000131113"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_add.value` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-2313030222311331-0310301120323200-3212110122231121-2111000013022300-2100300222033231-2110122110010311-3021221230021320-0223003232223323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.add_httponly

<a id="canonical-1213203130131311-0320303331203300-1100201221200310-1121213002001330-1220321233222002-3033311021113201-2102200302022331-2310223231010123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
add_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323232012020231-1323232033333100-0022200303021233-1132201023301022-1311100300203202-1032201201220220-2212131302003003-2302012011131123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned

<a id="canonical-0331010332023100-1232113001202213-0310220321001301-3121003003230311-1113112012213221-2331021120003033-3023232020302212-3121000320022130"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
add_partitioned = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213002321303031-3200322021113130-3120313220233100-0313331231203332-0301220311230132-0033000312233230-0221310333001103-3012223000332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.add_secure

<a id="canonical-2210203101310221-1121133030010130-2131212201233021-0233312101110333-2112131320112130-2000032233312100-2300013011212310-0121232132031011"></a>

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
add_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231230310231311-3031012003130120-0203331001131021-3303121233230332-0000320021103021-2122031210201132-0221212102002000-1020101120000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain

<a id="canonical-1321232010010302-0301323230010302-0331121101130230-3330011201212113-2123000011311232-2211313311311320-3212122022100031-0110201222002321"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_domain = {}
```

This is an empty object or choice marker. It has no direct properties.
