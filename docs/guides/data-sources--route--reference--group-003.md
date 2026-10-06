---
page_title: "xcsh_route reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route reference."
---

# xcsh_route reference

<a id="canonical-3102101002203110-3133012232102233-0103000213200132-3012233033202030-1030303023012011-3332031213233002-2111110122113313-0123120100103300"></a>

## Direct properties for `routes.route_destination`

<a id="canonical-3320103021021333-2320302212111331-0333023310023013-2021311320021223-1102233012033200-3001223012121022-0032120021023131-2331312331131202"></a>

### `routes.route_destination.auto_host_rewrite` property

Type: `"bool"`. Computed.

Exclusive with \[host\_rewrite\] Indicates that during forwarding, the host header will be swapped
with the hostname of the upstream host chosen by the cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [buffer_policy](data-sources--route--reference--group-003.md#canonical-2103102023331100-0133321201022320-1222313312300110-1131311133202030-1302332312131033-2321021222020322-3133223311131013-0100021223021211): complete subsection reference.

- [cors_policy](data-sources--route--reference--group-003.md#canonical-0201000211112222-3333031213200101-3213222110323103-3321222303012132-0231301210223203-0130303132302021-2221211112123311-2330112332213030): complete subsection reference.

- [csrf_policy](data-sources--route--reference--group-003.md#canonical-3022102001211100-2313000310321230-1003303001311110-2021313112020011-0200313211131120-1330122223032323-1321033100202233-3221011110001113): complete subsection reference.

- [destinations](data-sources--route--reference--group-003.md#canonical-0030003212020121-1223121010102011-2322312132003222-2210220313332211-1221133210032101-1103010232121021-2321313022100233-2111112303331310): complete subsection reference.

- [do_not_retract_cluster](data-sources--route--reference--group-003.md#canonical-2100002123202302-1001312313130222-1330132031031321-1032102030233001-3312231122101202-3323222210331211-0032010132120022-1311033101330202): complete subsection reference.

- [endpoint_subsets](data-sources--route--reference--group-003.md#canonical-2013022110111230-1123202020123012-1102201000122011-1313332231301300-0203301231132222-3220230221020003-1003012020313203-3102133133012031): complete subsection reference.

- [hash_policy](data-sources--route--reference--group-003.md#canonical-0011101332313032-0313133311230002-0012013103021233-0323013133332212-2021120323022222-3130333332220130-3001323200020210-3320100123102320): complete subsection reference.

<a id="canonical-2133310012230300-3121322112101113-1301130310223033-3213232333331021-2033000013121000-0321223033222022-2121013031130022-1203022101010203"></a>

<a id="canonical-3101313212203303-0132230211130303-2322231212033031-0121120302133002-1121202101211302-3323003033100323-3311112011101210-3302010302111313"></a>

### `routes.route_destination.host_rewrite` property

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite\] Indicates that during forwarding, the host header will be
swapped with this value.

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [mirror_policy](data-sources--route--reference--group-003.md#canonical-3102020312000110-0133212131123210-2001013313121121-0310130120100121-2021031311110033-1002330012133302-0123002120221110-1203302331112303): complete subsection reference.

<a id="canonical-3323002322120103-1100303323310211-1001322230333213-2303112211320333-2200131301120012-1222320003002222-2031003021122002-1113333110203322"></a>

<a id="canonical-3113320031112131-3033202100210210-2201011001120002-3132103331120032-0320113233222311-2001012323012133-0013001013023200-0220211101312122"></a>

### `routes.route_destination.prefix_rewrite` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_rewrite\] prefix\_rewrite indicates that during forwarding, the matched
prefix (or path) should be swapped with its value. When using regular expression path matching, the entire path
(not including the query string) will be swapped with this value. This option allows application
URLs to be rooted at a different path from those exposed at the reverse proxy layer.

Example : gcSpec: routes: &#8203;- match: &#8203;- headers: \[\] path: prefix : /register/
query\_params: \[\] &#8203;- headers: \[\] path: prefix: /register query\_params: \[\]
routeDestination: prefixRewrite: "/" destinations: &#8203;- cluster: &#8203;- kind: cluster.object
uid: cluster-1

Having above entries in the config, requests to /register will be stripped to /, while requests to
/register/public will be stripped to /public.

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

<a id="canonical-3033021010220123-0333220312323132-2132112322230112-0322221030111300-2220001112223212-0212333211012212-3130203320220013-0110333332131131"></a>

<a id="canonical-1211211323200210-2323330301101120-2203012121301310-3120221122123100-0131120200003022-2302112213333223-0331222012202311-2011010000113301"></a>

### `routes.route_destination.priority` property

Type: `"string"`. Computed.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Additional upstream details:

Priority routing for each request. Default routing mechanism High-Priority routing mechanism.

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](data-sources--route--reference--group-003.md#canonical-3322110030332211-3202003113030013-3113001312131212-3012021332111102-2300201312101321-1100011310013321-1323001321020110-1023331130200203): complete subsection reference.

- [regex_rewrite](data-sources--route--reference--group-003.md#canonical-1320330021331332-2111003023320303-0233313332222330-1002220332210022-0332130010123020-0021220321311000-2110022123211121-0220012110133301): complete subsection reference.

- [retract_cluster](data-sources--route--reference--group-003.md#canonical-3230120011210302-2102003303223110-3332103110331023-2033311110132002-2022330313021210-1321212023122331-1101031131301232-3203333300023131): complete subsection reference.

- [retry_policy](data-sources--route--reference--group-003.md#canonical-3113132130203033-3133232023221222-1012202110032132-2201200132220010-1230031332210222-3010022000320131-3323310030230202-1033112331331111): complete subsection reference.

- [spdy_config](data-sources--route--reference--group-003.md#canonical-2033321310103033-0133103213233310-2321121102310112-0031320012321002-3003013021300212-0111211013213310-0120322130212131-2220332212000120): complete subsection reference.

<a id="canonical-1022102013100110-0100022022001233-1021301310030120-0300100021032012-0032131220012321-0110003122310132-2123022131202200-3300130330010310"></a>

<a id="canonical-0223313301122020-0032212302213132-0333310100320023-0022332010102132-2001132100100323-0112322121313131-3302032020301113-1133232200202022"></a>

### `routes.route_destination.timeout` property

Type: `"number"`. Computed.

Specifies the timeout for the route in milliseconds. This timeout includes all retries. For server
side streaming, configure this field with higher value or leave it un-configured for infinite
timeout.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

- [web_socket_config](data-sources--route--reference--group-003.md#canonical-2220023021332222-0321110302230020-3013022312112230-1000032111332113-1232203003003021-1212100001333133-2322110231321100-1312103000111302): complete subsection reference.

<a id="canonical-2103102023331100-0133321201022320-1222313312300110-1131311133202030-1302332312131033-2321021222020322-3133223311131013-0100021223021211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.buffer_policy` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.buffer_policy

<a id="canonical-1302233012213012-2332130203230332-0022033321122321-0133100101210112-1312310002210131-1113101103331313-1232122213110233-2321302000032132"></a>

Type: `"single"`. Computed.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3331101103032022-0310012133120102-1203102110120131-2222221211112213-2232031010311101-1211033122010033-0230131101032222-0113112021231223"></a>

### Direct properties for `routes.route_destination.buffer_policy`

<a id="canonical-0002220032322303-1310330103010313-2313010133021202-1200032201132202-2321101203320221-1022100321030230-1332003311230310-0333311223011321"></a>

#### `routes.route_destination.buffer_policy.disabled` property

Type: `"bool"`. Computed.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3112021200022010-1000103231203003-2000321330011133-2231312333020033-2022001202122221-3210232312211001-2113022212320313-0103323002000111"></a>

<a id="canonical-1122233211320300-2312301210302120-1303323311000033-3331033233223202-2110010113222312-1231130111223120-2121020031103320-3113020230201301"></a>

#### `routes.route_destination.buffer_policy.max_request_bytes` property

Type: `"number"`. Computed.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-0201000211112222-3333031213200101-3213222110323103-3321222303012132-0231301210223203-0130303132302021-2221211112123311-2330112332213030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.cors_policy` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.cors_policy

<a id="canonical-1222330011322002-0121211303010100-3332332302301010-2230013201321332-1100220320003030-3310212210030220-3013310222103221-3023003132010010"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3223012020033200-0231201013103322-1023013012330011-0221212102011103-1023103100011003-3230001331322303-0111130312333220-2331313230332022"></a>

### Direct properties for `routes.route_destination.cors_policy`

<a id="canonical-2000320200211022-2200333213120312-1101101031300230-0033313323013312-2021212213011303-3130322330332330-3210301122102123-0101211320312030"></a>

#### `routes.route_destination.cors_policy.allow_credentials` property

Type: `"bool"`. Computed.

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

<a id="canonical-1232032130332330-2322331323030120-1223220211331111-0213010132302221-0032101022023323-2222231330331211-0320003320131213-3323102201220132"></a>

<a id="canonical-0033312300303112-0120030022100320-2033203032131212-3313213132211323-1111020222223132-3310121011131332-3201313222123302-3221322231030111"></a>

#### `routes.route_destination.cors_policy.allow_headers` property

Type: `"string"`. Computed.

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

<a id="canonical-1200200212313201-3312003232031101-0323120333002110-1002213223002023-1123121120002022-1300222010102300-1112302103133010-2133321102300203"></a>

<a id="canonical-3121201101201013-0023023131230332-0233233220203312-3222003011121330-1222223302113110-0120201232332313-3300210131200012-3203013332031201"></a>

#### `routes.route_destination.cors_policy.allow_methods` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-1331003012231210-1332013312023323-3310312012031111-0313202302002112-3133110332320223-2321002132122022-2132220020313330-3311100300232132"></a>

<a id="canonical-2110221222023010-1011212323023303-3311010320020320-0220303320213313-1312023110132232-1333233031032122-3020101311131012-1003230111232031"></a>

#### `routes.route_destination.cors_policy.allow_origin` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2022023130102200-0323232032112102-0312000030311021-2313130302321231-2121112311003300-2013111230211301-1011300231332132-1130201232330321"></a>

<a id="canonical-0131321033033023-1031021130202001-1033321010030113-0202320121301111-0213331210212220-0001020133320202-2203231232132303-3212230332221233"></a>

#### `routes.route_destination.cors_policy.allow_origin_regex` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0130220300201313-3120312323303032-0220202022310132-0113023010220230-3221313101102111-1232222030333232-2203020313200120-3210320330120300"></a>

<a id="canonical-3003133131322213-3221300230302010-2032211031322203-3112311121011223-3223301030311100-3031303123310132-1303213310133202-3131132131223223"></a>

#### `routes.route_destination.cors_policy.disabled` property

Type: `"bool"`. Computed.

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

<a id="canonical-1202201232130132-2103221113011010-0130002022203300-3002222232001003-3312222321231032-1332012002313322-2100133303123301-1100222022330131"></a>

<a id="canonical-3330020110303302-1202100333020322-0232130301320220-0122123033321333-0020032220231012-3011313331133230-1301331131012320-2112103213010031"></a>

#### `routes.route_destination.cors_policy.expose_headers` property

Type: `"string"`. Computed.

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

<a id="canonical-3100202122313133-1230200133310222-1100322311323303-1011221101323331-0101300021101333-2001133330020123-3300332332113331-0200021300210112"></a>

<a id="canonical-0222121002033210-2101010121312022-1103031321023201-1330012212030132-2100203023331203-0120332331012132-3033301332011213-1311222103132111"></a>

#### `routes.route_destination.cors_policy.maximum_age` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3022102001211100-2313000310321230-1003303001311110-2021313112020011-0200313211131120-1330122223032323-1321033100202233-3221011110001113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.csrf_policy` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.csrf_policy

<a id="canonical-0012221000302210-2120130031101110-2012202323101031-2313320121113330-2200000231213221-1313213202031322-1011110302101130-1333231101330023"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1123130202133020-0012133201213031-3132031211311230-1032313231010223-3311003221201113-0122231003011203-1002221200113202-1203030312132313"></a>

### Direct properties for `routes.route_destination.csrf_policy`

- [all_load_balancer_domains](data-sources--route--reference--group-003.md#canonical-3031030203000011-0302032033023211-0322022113010301-0113030322221212-2001101312021103-0311330221100012-3022003011322302-2231313232300133): complete subsection reference.

- [custom_domain_list](data-sources--route--reference--group-003.md#canonical-1133000222330111-2230111330210321-0331311220112133-2211133101103300-3330003213200321-2131103130033111-2230021010213203-3020030120332331): complete subsection reference.

- [disabled](data-sources--route--reference--group-003.md#canonical-2230130112002001-3000102303210012-0331233023230202-3023111112322312-0303201003101223-0010222201130030-3302333310333232-2132202331131223): complete subsection reference.

<a id="canonical-3031030203000011-0302032033023211-0322022113010301-0113030322221212-2001101312021103-0311330221100012-3022003011322302-2231313232300133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.csrf_policy.all_load_balancer_domains` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.csrf_policy](data-sources--route--reference--group-003.md#canonical-3022102001211100-2313000310321230-1003303001311110-2021313112020011-0200313211131120-1330122223032323-1321033100202233-3221011110001113)
- routes.route_destination.csrf_policy.all_load_balancer_domains

<a id="canonical-1032000232021032-2203220130120331-0131301313131311-1113122112330013-1303030130220023-3213321221033130-1233133133222002-2310000300110123"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133000222330111-2230111330210321-0331311220112133-2211133101103300-3330003213200321-2131103130033111-2230021010213203-3020030120332331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.csrf_policy.custom_domain_list` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.csrf_policy](data-sources--route--reference--group-003.md#canonical-3022102001211100-2313000310321230-1003303001311110-2021313112020011-0200313211131120-1330122223032323-1321033100202233-3221011110001113)
- routes.route_destination.csrf_policy.custom_domain_list

<a id="canonical-1310230003302211-0320032131312003-0302033020300113-1032310211003213-3020121102221123-3111233313101101-1022002112100212-0110310320332300"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1220322131220120-2311330302020033-3001333112313213-1000020102220122-2033132300033311-1033133200313003-1023022000103131-3220321300331211"></a>

### Direct properties for `routes.route_destination.csrf_policy.custom_domain_list`

<a id="canonical-2213030300110021-1211211230230101-0012111231313220-1123333213120122-0011132033230101-1301213333331200-0002213210303013-0121310321000213"></a>

#### `routes.route_destination.csrf_policy.custom_domain_list.domains` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2230130112002001-3000102303210012-0331233023230202-3023111112322312-0303201003101223-0010222201130030-3302333310333232-2132202331131223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.csrf_policy.disabled` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.csrf_policy](data-sources--route--reference--group-003.md#canonical-3022102001211100-2313000310321230-1003303001311110-2021313112020011-0200313211131120-1330122223032323-1321033100202233-3221011110001113)
- routes.route_destination.csrf_policy.disabled

<a id="canonical-2301122332221223-2232322311213222-1122330322022022-2221330002100111-1302230213111031-3232231211221220-1131132101110002-1121311021311332"></a>

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

<a id="canonical-0030003212020121-1223121010102011-2322312132003222-2210220313332211-1221133210032101-1103010232121021-2321313022100233-2111112303331310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.destinations` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.destinations

<a id="canonical-0123103300201300-0133220001112102-3123313021033101-2011331230332103-2003131110320020-0120031313112313-2032310300213120-1220333031112033"></a>

Type: `"list"`. Computed.

When requests have to distributed among multiple upstream clusters, multiple destinations are
configured, each having its own cluster and weight. Traffic is distributed among clusters based on
the weight configured.

Example: destinations: &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-1 weight: 20 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-2 weight: 30 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-3 weight: 50

This indicates that out of every 100 requests, 50 goes to cluster-3, 30 to cluster-2 and 20 to
cluster-1

When single destination is configured, weight is ignored. All the requests are sent to the cluster
specified in the destination.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2112023313031110-3012213110301101-0311012332030332-3021200230131221-2131111120211302-0310211013331130-0033130203120132-2132212023321323"></a>

### Direct properties for `routes.route_destination.destinations`

- [cluster](data-sources--route--reference--group-003.md#canonical-2133112322322321-1020122202303100-0222323122300210-1100320200322330-3123031102133302-1232023333133303-1212033032300201-3122121212223210): complete subsection reference.

- [endpoint_subsets](data-sources--route--reference--group-003.md#canonical-1312330300302211-0331122033121010-0110312211112133-3130111000103233-1303311322212122-3331120300002312-0102002021301102-0321122312100103): complete subsection reference.

<a id="canonical-1103023203302233-1103002001301022-0222322021130231-1132132130320201-1310323121113320-2221333201321103-0331130022003132-3312332022233033"></a>

<a id="canonical-3302021011222101-2221330032022112-0302103021111211-2121320100020302-1131122110231002-3113012212000332-2333122322121233-3302132331031103"></a>

#### `routes.route_destination.destinations.priority` property

Type: `"number"`. Computed.

Priority of this cluster, valid only with multiple destinations are configured. Value of 0 will make
the cluster as lowest priority upstream cluster Priority of 1 means highest priority and is
considered active. When active cluster is not available, lower priority clusters are made active as
per the increasing priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-2022211031111231-2212022313310213-0013233222122113-0212131131202200-1033030331020132-1212312011300011-2131120020221311-1320021331131102"></a>

<a id="canonical-1031322310332001-3132003031310110-0331003332310102-3010121301310122-2223233122320023-2012310223312331-1210013001232122-2033010320122001"></a>

#### `routes.route_destination.destinations.weight` property

Type: `"number"`. Computed.

When requests have to distributed among multiple upstream clusters, multiple destinations are
configured, each having its own cluster and weight. Traffic is distributed among clusters based on
the weight configured.

Example: destinations: &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-1 weight: 20 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-2 weight: 30 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-3 weight: 10

This indicates that out of every 60 requests, 10 goes to cluster-3, 30 to cluster-2 and 20 to
cluster-1

When single destination is configured, weight is ignored. All the requests are sent to the cluster
specified in the destination.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2133112322322321-1020122202303100-0222323122300210-1100320200322330-3123031102133302-1232023333133303-1212033032300201-3122121212223210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.destinations.cluster` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.destinations](data-sources--route--reference--group-003.md#canonical-0030003212020121-1223121010102011-2322312132003222-2210220313332211-1221133210032101-1103010232121021-2321313022100233-2111112303331310)
- routes.route_destination.destinations.cluster

<a id="canonical-3231230112300012-3211213301201003-2110332000213333-0101013113222110-3201203101110031-0232021023213013-2203003301020202-0203200201230120"></a>

Type: `"list"`. Computed.

Indicates the upstream cluster to which the request should be sent. If the cluster does not exist
ServiceUnavailable response will be sent.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0203121332223333-2310012121030003-0223030120232112-0112221112033011-1311231321131312-3122210201032313-0311211310301122-1312101112121133"></a>

### Direct properties for `routes.route_destination.destinations.cluster`

<a id="canonical-2023203023322221-1321230112011222-3011221132032203-3302010331000120-0320332303303131-2003102120210212-2101330023112130-0100302203123323"></a>

#### `routes.route_destination.destinations.cluster.kind` property

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

<a id="canonical-3112311221211210-0313322000331322-3221221300132210-2232132021013313-1220221000112000-2333233001121333-2213300321023020-2002022213302313"></a>

<a id="canonical-2003123000302231-3001322313233202-0223203212233201-3031030103031130-0323330332310002-0201100033232221-3211132312231021-2311131312231021"></a>

#### `routes.route_destination.destinations.cluster.name` property

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

<a id="canonical-1132332131312022-0012121102011013-0303013013202103-0003020321232012-1030301332303130-0103231313233132-2201032303312133-2130120022312231"></a>

<a id="canonical-0032333000312102-2133120131212203-0033120120312313-2010130022101310-0313311300023210-0020012121323313-3133330313010311-3013112130031310"></a>

#### `routes.route_destination.destinations.cluster.namespace` property

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
  }
}
```

<a id="canonical-1312033031332331-2333123300111032-3312101222223021-0203001102102000-0213003332030022-1000132313313211-1133131312313020-0203333113111120"></a>

<a id="canonical-1311332103231102-2123110131012212-0223332302312030-3230032023111310-1122131011111221-1003000020320122-2101003023303011-1301323103103230"></a>

#### `routes.route_destination.destinations.cluster.tenant` property

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

<a id="canonical-3100310020103031-3322011000230033-0321313032113210-1231022103033223-0033200311230123-2321003021313322-3033102131111333-2002011131010313"></a>

<a id="canonical-0111301131001023-1310301313131302-2333110323033320-2313323203131000-3330113221303321-2201223001321011-0121231003301200-2021001033232202"></a>

#### `routes.route_destination.destinations.cluster.uid` property

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

<a id="canonical-1312330300302211-0331122033121010-0110312211112133-3130111000103233-1303311322212122-3331120300002312-0102002021301102-0321122312100103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.destinations.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.destinations](data-sources--route--reference--group-003.md#canonical-0030003212020121-1223121010102011-2322312132003222-2210220313332211-1221133210032101-1103010232121021-2321313022100233-2111112303331310)
- routes.route_destination.destinations.endpoint_subsets

<a id="canonical-3221313010300123-3230231303320312-1033130310323111-3113200111331201-1130201213122301-0000220101113010-0203222023210212-2203212230320121"></a>

Type: `"single"`. Computed.

Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached
to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be
selected by the load balancer

Labels field of endpoint object's metadata is used for subset matching. For endpoints which are
discovered in K8s or Consul cluster, the label of the service is merged with endpoint's labels. In
case of Consul, the label is derived from the "Tag" field. For labels that are common between
configured endpoint and discovered service, labels from discovered service takes precedence.

List of key-value pairs that will be used as matching metadata. Only those endpoints of upstream
cluster which match this metadata will be selected for load balancing.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100002123202302-1001312313130222-1330132031031321-1032102030233001-3312231122101202-3323222210331211-0032010132120022-1311033101330202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.do_not_retract_cluster` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.do_not_retract_cluster

<a id="canonical-3300002313111001-3022013222331312-3300201203021323-1120332212313202-2030231331113330-0331132000121002-2312231120211321-1211203021322210"></a>

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

<a id="canonical-2013022110111230-1123202020123012-1102201000122011-1313332231301300-0203301231132222-3220230221020003-1003012020313203-3102133133012031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.endpoint_subsets

<a id="canonical-0111220131233220-2011102323133313-1112213011000231-1212031321203230-2101000321020020-0113220131200313-2211002023110200-3210022323112030"></a>

Type: `"single"`. Computed.

Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached
to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be
selected by the load balancer

Labels field of endpoint object's metadata is used for subset matching. For endpoint's which are
discovered in K8s or Consul cluster, the label of the service is merged with endpoint's labels. In
case of Consul, the label is derived from the "Tag" field. For labels that are common between
configured endpoint and discovered service, labels from discovered service takes precedence.

List of key-value pairs that will be used as matching metadata. Only those endpoints of upstream
cluster which match this metadata will be selected for load balancing.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011101332313032-0313133311230002-0012013103021233-0323013133332212-2021120323022222-3130333332220130-3001323200020210-3320100123102320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.hash_policy

<a id="canonical-0101203232313012-0132111110103110-1210323003221112-0231201021123032-0131030302333120-2031100233303130-1203221210132131-0010013120110303"></a>

Type: `"list"`. Computed.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-3201202300113030-1313101212321131-0322000200120013-1021220222003020-3333120212302000-0121123301213221-3032310332011212-3003332333101000"></a>

### Direct properties for `routes.route_destination.hash_policy`

- [cookie](data-sources--route--reference--group-003.md#canonical-3300222200100202-2221333320332130-2220132213300033-3113023022203311-1103220320113033-3021320101300123-2103000333122303-1032032230202231): complete subsection reference.

<a id="canonical-3331100322001333-1103323120000230-2313220003312222-1303013313123111-3233301112322001-0012033330101111-2321313103233322-0200031001133103"></a>

<a id="canonical-2002212213222033-0302100123020003-2032210233003012-1311232221100222-0323111013312302-0312033200122021-2230310130223001-2330210330312221"></a>

#### `routes.route_destination.hash_policy.header_name` property

Type: `"string"`. Computed.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0212101201012011-1230332123331333-0202321200213020-3021233331332113-2232303131323312-0200130112201132-1122020111330132-1231221320233020"></a>

<a id="canonical-0021001311211110-1220231011001122-0030121220332113-2031230331333023-3122023330330301-3132012033031001-2111020102223330-1312133112220120"></a>

#### `routes.route_destination.hash_policy.source_ip` property

Type: `"bool"`. Computed.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1112023313102231-1202210123321230-0122133132110022-1222210102011110-0133222213021121-2020011220321210-2200022230011222-0331211031120112"></a>

<a id="canonical-3122202222112330-1302330313102323-1322130131213111-1223011212033000-0130131111032310-0313032232303231-2023322100010101-2030000112020213"></a>

#### `routes.route_destination.hash_policy.terminal` property

Type: `"bool"`. Computed.

Terminal. Specify if its a terminal policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3300222200100202-2221333320332130-2220132213300033-3113023022203311-1103220320113033-3021320101300123-2103000333122303-1032032230202231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.hash_policy](data-sources--route--reference--group-003.md#canonical-0011101332313032-0313133311230002-0012013103021233-0323013133332212-2021120323022222-3130333332220130-3001323200020210-3320100123102320)
- routes.route_destination.hash_policy.cookie

<a id="canonical-1330312103300201-0203133022010202-0102130111003330-3333121220300030-0233032132221322-1310332022201012-1012311331011031-1010011023002321"></a>

Type: `"single"`. Computed.

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

<a id="canonical-1323010120113110-3333113032212110-3031320211303110-2320333212022312-1222130121003220-3210321212030010-2213303123213131-0002233323302201"></a>

### Direct properties for `routes.route_destination.hash_policy.cookie`

- [add_httponly](data-sources--route--reference--group-003.md#canonical-2210031120100201-1123003101233220-3012030030123323-1030111321211022-2120022222112201-3211120031102202-3312110131100023-3113010021222013): complete subsection reference.

- [add_secure](data-sources--route--reference--group-003.md#canonical-3310202200022303-1223213133123021-0313300220212020-3023211223102202-0220022030223313-2312213011103313-1120333121123113-0321231310200122): complete subsection reference.

- [ignore_httponly](data-sources--route--reference--group-003.md#canonical-1303023312321011-3231003113210112-2030010023112332-2332222213211120-1300012332302013-1110011012132002-2231321103121000-3321313120103123): complete subsection reference.

- [ignore_samesite](data-sources--route--reference--group-003.md#canonical-3303032121231010-3330103001112203-3322012020233202-2010230033121312-2001110022110302-2030301333021310-3003333203232201-0003212311222111): complete subsection reference.

- [ignore_secure](data-sources--route--reference--group-003.md#canonical-0102221013313302-0300302121323301-0323201021100223-3020011121333013-1001031232202010-2103311103020132-1131002111330000-0101202030120322): complete subsection reference.

<a id="canonical-1132232000111222-2303302022220233-3220220322013210-0210332111012131-3331310201301301-0103232033033301-0100011232301232-2130211213323102"></a>

<a id="canonical-2031101103032201-3230030131000001-2033120122312000-1202320021000102-1030212231032113-1200022233022110-2022033332330211-0120113330220011"></a>

#### `routes.route_destination.hash_policy.cookie.name` property

Type: `"string"`. Computed.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3302313302130112-1001330323221022-3222023021331021-3223211221001020-0113121120212231-1332013201232000-3201330321222012-0133021032003223"></a>

<a id="canonical-0003200030331333-1002121113310112-0213311330113322-2033030230210313-3111031303130113-1011232020203301-1232021130201010-3021033012021001"></a>

#### `routes.route_destination.hash_policy.cookie.path` property

Type: `"string"`. Computed.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](data-sources--route--reference--group-003.md#canonical-3003001102112303-0130003230103331-1322300120200110-1311012323131203-2321130320302313-1222202021223110-0310030022332110-2221303100020233): complete subsection reference.

- [samesite_none](data-sources--route--reference--group-003.md#canonical-0123000331001201-0301000122102120-2231332102112002-3023120321101322-0223023002113320-3323030203233312-0201120302020300-3102333120130030): complete subsection reference.

- [samesite_strict](data-sources--route--reference--group-003.md#canonical-0102231000000212-1310032220210123-1230321000003123-3133312001210101-0310300033132311-3121020023323212-0011003002221310-0102312032032201): complete subsection reference.

<a id="canonical-3233023033100300-3030032022322130-2121003320033303-3202301333130030-2202033010010003-3111330230102323-1133232011203332-1113221110000213"></a>

<a id="canonical-2201011323021302-3312203311213212-3212310002312030-1010302213320213-3130223233331011-1020311312311031-2101201013220121-2313233203030122"></a>

#### `routes.route_destination.hash_policy.cookie.ttl` property

Type: `"number"`. Computed.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2210031120100201-1123003101233220-3012030030123323-1030111321211022-2120022222112201-3211120031102202-3312110131100023-3113010021222013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.add_httponly` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.hash_policy](data-sources--route--reference--group-003.md#canonical-0011101332313032-0313133311230002-0012013103021233-0323013133332212-2021120323022222-3130333332220130-3001323200020210-3320100123102320)
- [routes.route_destination.hash_policy.cookie](data-sources--route--reference--group-003.md#canonical-3300222200100202-2221333320332130-2220132213300033-3113023022203311-1103220320113033-3021320101300123-2103000333122303-1032032230202231)
- routes.route_destination.hash_policy.cookie.add_httponly

<a id="canonical-2212202030102202-2300032213100111-0133023110320230-3130220102103131-2321111133123221-2202303313113203-3133320213311320-0221003103021112"></a>

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

<a id="canonical-3310202200022303-1223213133123021-0313300220212020-3023211223102202-0220022030223313-2312213011103313-1120333121123113-0321231310200122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.add_secure` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.hash_policy](data-sources--route--reference--group-003.md#canonical-0011101332313032-0313133311230002-0012013103021233-0323013133332212-2021120323022222-3130333332220130-3001323200020210-3320100123102320)
- [routes.route_destination.hash_policy.cookie](data-sources--route--reference--group-003.md#canonical-3300222200100202-2221333320332130-2220132213300033-3113023022203311-1103220320113033-3021320101300123-2103000333122303-1032032230202231)
- routes.route_destination.hash_policy.cookie.add_secure

<a id="canonical-0312100323202112-3130232033222231-3022333013221230-3100310231021111-1223022301123111-0113021121122201-2102120222222011-3321133012303010"></a>

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

<a id="canonical-1303023312321011-3231003113210112-2030010023112332-2332222213211120-1300012332302013-1110011012132002-2231321103121000-3321313120103123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.ignore_httponly` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.hash_policy](data-sources--route--reference--group-003.md#canonical-0011101332313032-0313133311230002-0012013103021233-0323013133332212-2021120323022222-3130333332220130-3001323200020210-3320100123102320)
- [routes.route_destination.hash_policy.cookie](data-sources--route--reference--group-003.md#canonical-3300222200100202-2221333320332130-2220132213300033-3113023022203311-1103220320113033-3021320101300123-2103000333122303-1032032230202231)
- routes.route_destination.hash_policy.cookie.ignore_httponly

<a id="canonical-0121200123102130-2010232111210121-2320220033300111-2123003301332311-2030323003303203-2222113313203132-0000121303001112-2033010302213101"></a>

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

<a id="canonical-3303032121231010-3330103001112203-3322012020233202-2010230033121312-2001110022110302-2030301333021310-3003333203232201-0003212311222111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.ignore_samesite` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.hash_policy](data-sources--route--reference--group-003.md#canonical-0011101332313032-0313133311230002-0012013103021233-0323013133332212-2021120323022222-3130333332220130-3001323200020210-3320100123102320)
- [routes.route_destination.hash_policy.cookie](data-sources--route--reference--group-003.md#canonical-3300222200100202-2221333320332130-2220132213300033-3113023022203311-1103220320113033-3021320101300123-2103000333122303-1032032230202231)
- routes.route_destination.hash_policy.cookie.ignore_samesite

<a id="canonical-1113021333132020-1312013331300031-1313100321010012-1012233213112311-1021211022023020-1012221010321113-1013223213033323-3320232033111331"></a>

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

<a id="canonical-0102221013313302-0300302121323301-0323201021100223-3020011121333013-1001031232202010-2103311103020132-1131002111330000-0101202030120322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.ignore_secure` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.hash_policy](data-sources--route--reference--group-003.md#canonical-0011101332313032-0313133311230002-0012013103021233-0323013133332212-2021120323022222-3130333332220130-3001323200020210-3320100123102320)
- [routes.route_destination.hash_policy.cookie](data-sources--route--reference--group-003.md#canonical-3300222200100202-2221333320332130-2220132213300033-3113023022203311-1103220320113033-3021320101300123-2103000333122303-1032032230202231)
- routes.route_destination.hash_policy.cookie.ignore_secure

<a id="canonical-1200220122311232-1100133010122202-3010312200212223-2230213033113302-2211003220132203-3223021103101232-3210030223021222-1202112013003322"></a>

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

<a id="canonical-3003001102112303-0130003230103331-1322300120200110-1311012323131203-2321130320302313-1222202021223110-0310030022332110-2221303100020233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.samesite_lax` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.hash_policy](data-sources--route--reference--group-003.md#canonical-0011101332313032-0313133311230002-0012013103021233-0323013133332212-2021120323022222-3130333332220130-3001323200020210-3320100123102320)
- [routes.route_destination.hash_policy.cookie](data-sources--route--reference--group-003.md#canonical-3300222200100202-2221333320332130-2220132213300033-3113023022203311-1103220320113033-3021320101300123-2103000333122303-1032032230202231)
- routes.route_destination.hash_policy.cookie.samesite_lax

<a id="canonical-1101030203132110-0030302300002311-3200031203113002-2220100020130211-0123333231031310-1212001300113030-0132202103212210-2101103331311332"></a>

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

<a id="canonical-0123000331001201-0301000122102120-2231332102112002-3023120321101322-0223023002113320-3323030203233312-0201120302020300-3102333120130030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.samesite_none` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.hash_policy](data-sources--route--reference--group-003.md#canonical-0011101332313032-0313133311230002-0012013103021233-0323013133332212-2021120323022222-3130333332220130-3001323200020210-3320100123102320)
- [routes.route_destination.hash_policy.cookie](data-sources--route--reference--group-003.md#canonical-3300222200100202-2221333320332130-2220132213300033-3113023022203311-1103220320113033-3021320101300123-2103000333122303-1032032230202231)
- routes.route_destination.hash_policy.cookie.samesite_none

<a id="canonical-2020310323332301-0111033321003303-3303032020132310-2023000202310013-1120222211233133-1110211113102203-1233133332301211-1012230110133210"></a>

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

<a id="canonical-0102231000000212-1310032220210123-1230321000003123-3133312001210101-0310300033132311-3121020023323212-0011003002221310-0102312032032201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.samesite_strict` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.hash_policy](data-sources--route--reference--group-003.md#canonical-0011101332313032-0313133311230002-0012013103021233-0323013133332212-2021120323022222-3130333332220130-3001323200020210-3320100123102320)
- [routes.route_destination.hash_policy.cookie](data-sources--route--reference--group-003.md#canonical-3300222200100202-2221333320332130-2220132213300033-3113023022203311-1103220320113033-3021320101300123-2103000333122303-1032032230202231)
- routes.route_destination.hash_policy.cookie.samesite_strict

<a id="canonical-3222201312131122-2200132013000131-0220203130232302-1330230132131011-1020201000323001-2321323313303221-1101003200020202-1110030012302011"></a>

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

<a id="canonical-3102020312000110-0133212131123210-2001013313121121-0310130120100121-2021031311110033-1002330012133302-0123002120221110-1203302331112303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.mirror_policy` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.mirror_policy

<a id="canonical-3012212223031212-2310112013100223-0103302112000003-1321112330111213-2022221331131203-2320313013030310-1221323010120003-3000300130110112"></a>

Type: `"single"`. Computed.

MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is 'fire
and forget', meaning it will not wait for the shadow cluster to respond before returning the
response from the primary cluster. All normal statistics are collected for the shadow cluster
making..

Additional upstream details:

The approach used is "fire and forget", meaning it will not wait for the shadow cluster to respond
before returning the response from the primary cluster. All normal statistics are collected for the
shadow cluster making this feature useful for testing and troubleshooting.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1203002312231302-2121202113200302-2332013131131220-0033300221031232-0003101202322213-2301231201213110-1212302231020310-2101313013202211"></a>

### Direct properties for `routes.route_destination.mirror_policy`

- [cluster](data-sources--route--reference--group-003.md#canonical-2130110203200021-1301131010012213-3300021201303223-1301300320233211-3302110001022322-2302133100121120-1201102223121220-0021102033201011): complete subsection reference.

- [percent](data-sources--route--reference--group-003.md#canonical-2033102300222201-1130002321303102-0113123131212000-3012011131223313-0112000033320223-3120011322233132-2113120310210211-1113023103311022): complete subsection reference.

<a id="canonical-2130110203200021-1301131010012213-3300021201303223-1301300320233211-3302110001022322-2302133100121120-1201102223121220-0021102033201011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.mirror_policy.cluster` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.mirror_policy](data-sources--route--reference--group-003.md#canonical-3102020312000110-0133212131123210-2001013313121121-0310130120100121-2021031311110033-1002330012133302-0123002120221110-1203302331112303)
- routes.route_destination.mirror_policy.cluster

<a id="canonical-2333032232110233-1231100021031203-2133320101002202-1023200020210300-3313031000032200-3313031103001111-2201032122321320-1333032232332012"></a>

Type: `"list"`. Computed.

Specifies the cluster to which the requests will be mirrored. The cluster object referred here must
be present.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0103220231221311-2103300232300321-2021200011013022-0123231113120030-0020320022313203-3201012012322210-1130322221200211-1313010100321133"></a>

### Direct properties for `routes.route_destination.mirror_policy.cluster`

<a id="canonical-1333330032203031-0311112100110322-0320200012203021-3232122000001310-3033131200133132-3221021331202122-3331112333331100-0113333112102002"></a>

#### `routes.route_destination.mirror_policy.cluster.kind` property

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

<a id="canonical-0311212211213111-0002033132023332-3000322320123311-0130122023200220-0233200123210023-0003210331102100-1000332101213131-1221113310213321"></a>

<a id="canonical-2111313213211030-3310323223312002-0133220010202220-2231212102320202-2111112231303321-2332202210213323-1233012212303010-3103211212200023"></a>

#### `routes.route_destination.mirror_policy.cluster.name` property

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

<a id="canonical-3212000122103123-2000213232202201-0232333323212122-2210202300201320-2131330332011231-3030321223332230-3103231003220133-3120213103223130"></a>

<a id="canonical-1201203311110203-1210210231113310-3123122023232313-3330113323022202-3132021021333303-0302211323301301-2103011221333332-3030303112200103"></a>

#### `routes.route_destination.mirror_policy.cluster.namespace` property

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
  }
}
```

<a id="canonical-0233202213221102-1123313031222303-2232033020320011-3130323123021131-1111220023012030-0102131101232300-0103300013222013-3101022230132321"></a>

<a id="canonical-1211013021231233-2212031322231013-1003221201100231-1131122130021202-3121132130100203-3101013120302311-2232123103201333-3313111102013113"></a>

#### `routes.route_destination.mirror_policy.cluster.tenant` property

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

<a id="canonical-2202302113211221-3232020113200333-1220100223101002-0120003333201002-2111000013213132-1030121011132323-0201110133033203-3101131110103121"></a>

<a id="canonical-1100001313102111-1132213112320323-1301110332211201-3211323101311132-1312212201300020-2203023300030112-3323313001321023-3013101301222301"></a>

#### `routes.route_destination.mirror_policy.cluster.uid` property

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

<a id="canonical-2033102300222201-1130002321303102-0113123131212000-3012011131223313-0112000033320223-3120011322233132-2113120310210211-1113023103311022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.mirror_policy.percent` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.mirror_policy](data-sources--route--reference--group-003.md#canonical-3102020312000110-0133212131123210-2001013313121121-0310130120100121-2021031311110033-1002330012133302-0123002120221110-1203302331112303)
- routes.route_destination.mirror_policy.percent

<a id="canonical-0012112212323330-0312130300111032-1211320130330222-1132223200230113-2323320222211032-3003323203003223-1113022031201121-1101200120102323"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0122102231312313-0032232032123032-1222331022020022-0231223030132001-2311130030220203-2331210222323020-3313333232031213-3211001032010020"></a>

### Direct properties for `routes.route_destination.mirror_policy.percent`

<a id="canonical-3031032003122003-0220033330032222-2113031220000311-1330022322202213-2030101100303013-3003000130000210-2332203302200113-2113130103231300"></a>

#### `routes.route_destination.mirror_policy.percent.denominator` property

Type: `"string"`. Computed.

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

<a id="canonical-0102022313121132-2120020311213101-0212302221012011-2220232001001032-2020302202100333-1330221313010230-3321001002202302-2110312210230001"></a>

<a id="canonical-0321222300322221-2203300121132331-3330313030030321-1012202020310203-1312333301133010-3201022030203001-1023023200023232-1220220003333021"></a>

#### `routes.route_destination.mirror_policy.percent.numerator` property

Type: `"number"`. Computed.

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

<a id="canonical-3322110030332211-3202003113030013-3113001312131212-3012021332111102-2300201312101321-1100011310013321-1323001321020110-1023331130200203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.query_params` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.query_params

<a id="canonical-3321110120313220-2013020002030030-3011001131132323-1231023120231312-1013333001131123-0333102222303113-1323321101023310-3033020133300023"></a>

Type: `"single"`. Computed.

Handling of incoming query parameters in simple route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]"
}
```

<a id="canonical-2211332121203100-1303030330110310-2100333111121130-3202230121032101-2030210100230212-0231012023101003-2212310202201021-3001101110132212"></a>

### Direct properties for `routes.route_destination.query_params`

- [remove_all_params](data-sources--route--reference--group-003.md#canonical-2012021100101312-0311203230112022-3323123003000302-1220000122002233-0332302020132302-0223302122113012-1112032032132131-1213012230111020): complete subsection reference.

<a id="canonical-0203010002012011-0333230002222010-2023232102311120-2211330233113322-1331212322311002-1103013310322120-2022231122013220-1013322101231230"></a>

<a id="canonical-0330320132003112-2320123103232212-1200222030222111-2230330320133330-3113120012020202-2200323102220313-0031020221121200-3130103132110303"></a>

#### `routes.route_destination.query_params.replace_params` property

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [retain_all_params](data-sources--route--reference--group-003.md#canonical-2310031100113112-2202002021110101-2203032230232112-0020233202030223-1223003300311131-0103111101222033-3110222120333300-2103122222020010): complete subsection reference.

<a id="canonical-2012021100101312-0311203230112022-3323123003000302-1220000122002233-0332302020132302-0223302122113012-1112032032132131-1213012230111020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.query_params.remove_all_params` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.query_params](data-sources--route--reference--group-003.md#canonical-3322110030332211-3202003113030013-3113001312131212-3012021332111102-2300201312101321-1100011310013321-1323001321020110-1023331130200203)
- routes.route_destination.query_params.remove_all_params

<a id="canonical-2101121003203123-0223021020032113-2011330012200012-1112330210122201-0123310011221321-1021000212230312-1300311210320321-3113000211110202"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for remove all params.

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

<a id="canonical-2310031100113112-2202002021110101-2203032230232112-0020233202030223-1223003300311131-0103111101222033-3110222120333300-2103122222020010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.query_params.retain_all_params` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.query_params](data-sources--route--reference--group-003.md#canonical-3322110030332211-3202003113030013-3113001312131212-3012021332111102-2300201312101321-1100011310013321-1323001321020110-1023331130200203)
- routes.route_destination.query_params.retain_all_params

<a id="canonical-3110131320031132-3132130210012030-3101203030232302-0011131012011223-2303022312202101-3213312211231331-1021213213223021-0102223332230102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for retain all params.

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

<a id="canonical-1320330021331332-2111003023320303-0233313332222330-1002220332210022-0332130010123020-0021220321311000-2110022123211121-0220012110133301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.regex_rewrite` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.regex_rewrite

<a id="canonical-1113333100322133-0311223100133020-1312111323022322-3112013221213100-0132012310210201-0121122001313302-2311000331013232-1322213121333000"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3302200300021311-3200211020133001-2312121303313030-0123100323022310-0320201010110313-3021120332001300-2110003300210111-2033123032220300"></a>

### Direct properties for `routes.route_destination.regex_rewrite`

<a id="canonical-0221130301220333-3203200010020132-0012131302201110-0313131221100312-1231103120310021-2022103212312202-2301120220220012-2113001333130132"></a>

#### `routes.route_destination.regex_rewrite.pattern` property

Type: `"string"`. Computed.

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

<a id="canonical-1200223033022111-3331121302202130-3320122212220003-2323300002020100-0113231123331033-3200123033103003-0230201033222133-1021321310023201"></a>

<a id="canonical-3122013030322320-2201132312312312-3011223320101012-3313000300231323-0222032010221212-0232210333112330-1033301013303202-3330103222303030"></a>

#### `routes.route_destination.regex_rewrite.substitution` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3230120011210302-2102003303223110-3332103110331023-2033311110132002-2022330313021210-1321212023122331-1101031131301232-3203333300023131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.retract_cluster` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.retract_cluster

<a id="canonical-3001033031301301-2302130030020123-1022212302202131-0113301002112223-1223321322333122-1311001322132100-3321100203130133-2321133102001113"></a>

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

<a id="canonical-3113132130203033-3133232023221222-1012202110032132-2201200132220010-1230031332210222-3010022000320131-3323310030230202-1033112331331111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.retry_policy` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.retry_policy

<a id="canonical-3100123200200022-0301001002022331-1320200312130020-2310213003202101-1322001202303011-1023021322132313-1330113100311330-1303130112011230"></a>

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

<a id="canonical-0223023120133000-2312112130232030-3322232333203210-2103031200021112-0011111032000123-2331220110133103-1323110201200312-2311021230001012"></a>

### Direct properties for `routes.route_destination.retry_policy`

- [back_off](data-sources--route--reference--group-003.md#canonical-3102121022302111-0213310320213223-3302022031001103-0002012330323031-3203310302312010-0223223000221300-0112101003313012-3310120012220312): complete subsection reference.

<a id="canonical-0212213312211330-0201013320113230-3011031203221021-2102321122023302-1013022030011203-3303332312132332-0332113120100002-0303021023002030"></a>

<a id="canonical-3002220013020033-0131021121121330-1020122223111100-3233022230321201-0333131132003120-1210013211003031-2300230202301021-0210023201313003"></a>

#### `routes.route_destination.retry_policy.num_retries` property

Type: `"number"`. Computed.

Specifies the allowed number of retries. Retries can be done any number of times. An exponential
back-off algorithm is used between each retry. Defaults to \`1\`.

Additional upstream details:

Defaults to 1.

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
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

<a id="canonical-2102011133012302-0130330312323333-2200322110200233-1323320330000233-2032013220221302-2311100131313032-0323302013332131-2221012011113211"></a>

<a id="canonical-2120130212200002-0030230000222010-0023220231011030-3300200211002222-2311223010123121-1113100221222011-0102203321133223-1213110200102122"></a>

#### `routes.route_destination.retry_policy.per_try_timeout` property

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-0101101120322203-3300200013002221-3023012203033301-3220012031012032-3221003120003210-1102133211030212-2330003033312130-2311220321320122"></a>

<a id="canonical-0302333022322122-3020120023211133-0202023111320212-1323300110112000-2203333033300332-0332021200202310-0000133100130213-0022010211312310"></a>

#### `routes.route_destination.retry_policy.retriable_status_codes` property

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2231110123331132-3203021031101033-0131112023322130-0303303000320130-2301012201123310-3120222200321303-0131110022000300-0221102311202020"></a>

<a id="canonical-2122101231130111-3112022021020130-2220220323101121-1200122133100030-3301321132133021-2222020033300130-1130303101131121-2123303110020223"></a>

#### `routes.route_destination.retry_policy.retry_condition` property

Type: `["list", "string"]`. Computed.

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if
the..

Additional upstream details:

For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc

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

<a id="canonical-3102121022302111-0213310320213223-3302022031001103-0002012330323031-3203310302312010-0223223000221300-0112101003313012-3310120012220312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.retry_policy.back_off` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- [routes.route_destination.retry_policy](data-sources--route--reference--group-003.md#canonical-3113132130203033-3133232023221222-1012202110032132-2201200132220010-1230031332210222-3010022000320131-3323310030230202-1033112331331111)
- routes.route_destination.retry_policy.back_off

<a id="canonical-1233313102211331-3003211003132213-1031031220002231-3032021330100013-0131133100210112-2323312123000110-1130113221302320-3132321112331111"></a>

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

<a id="canonical-0333021023203223-3310000200331020-2131032221102011-2311021002111300-3120310023012301-0230302112332010-2031300300122000-0310023232013211"></a>

### Direct properties for `routes.route_destination.retry_policy.back_off`

<a id="canonical-0300121220113300-1032120312211100-0010210002303111-2301220310131312-3113112100030023-3311302230222210-2010323331131312-0022111013101020"></a>

#### `routes.route_destination.retry_policy.back_off.base_interval` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1320130323132312-3322222223222320-0000332110302322-3332120321021021-3200222131200300-3202020033210030-3331020000303330-3203032323021302"></a>

<a id="canonical-0123213222330301-0212121031033210-0032023111022002-2012220330223301-1213110030212133-1313122111310301-1103021312212011-2223023213023023"></a>

#### `routes.route_destination.retry_policy.back_off.max_interval` property

Type: `"number"`. Computed.

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The times the base\_interval. Defaults to
\`10\`.

Additional upstream details:

The default is 10 times the base\_interval.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2033321310103033-0133103213233310-2321121102310112-0031320012321002-3003013021300212-0111211013213310-0120322130212131-2220332212000120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.spdy_config` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.spdy_config

<a id="canonical-2220333321120322-1012311301331321-1222310330312303-0211211333230002-1211312321113233-2121331210102023-2211301303232213-0333222022011000"></a>

Type: `"single"`. Computed.

Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'SPDY/3.1'

Configuration to allow UPGRADE of connection to SPDY and any additional tuning With configuration to
allow SPDY upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade':
'SPDY/3.1' 'Connection': 'Upgrade'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3121321210330101-1120131123101311-0322030233322023-0312010233003211-2313331103010020-2022031233230132-1323233310122100-1332130233202100"></a>

### Direct properties for `routes.route_destination.spdy_config`

<a id="canonical-2300000302331011-2100120201023231-0330310011312311-3302233202003231-0031322101013331-1022030122131200-3010120231013300-0310201300003323"></a>

#### `routes.route_destination.spdy_config.use_spdy` property

Type: `"bool"`. Computed.

Specifies that the HTTP client connection to this route is allowed to upgrade to a SPDY connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2220023021332222-0321110302230020-3013022312112230-1000032111332113-1232203003003021-1212100001333133-2322110231321100-1312103000111302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.web_socket_config` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230)
- routes.route_destination.web_socket_config

<a id="canonical-2111330201102302-2111201100010002-2010103013301033-3332110123332101-1001013130200012-0302033232323010-1010103000323230-1023130113102023"></a>

Type: `"single"`. Computed.

Configuration to allow Websocket Request headers of such upgrade looks like below 'connection',
'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce
following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0000231313100100-2230031223123023-0120230000030313-1030020011213302-0102020002000333-3001223033021220-0211300212121333-0230100231302001"></a>

### Direct properties for `routes.route_destination.web_socket_config`

<a id="canonical-1213211200100330-1201003203223231-1301100011132011-3332333010001231-1222202321000103-0133312301202333-1002002233333331-2230211303311112"></a>

#### `routes.route_destination.web_socket_config.use_websocket` property

Type: `"bool"`. Computed.

Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket
connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1120103010111221-1131011202033302-3012012120230213-0222131120131102-2101233332131203-3211033100232203-0013202020220303-2233021120011122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_direct_response` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.route_direct_response

<a id="canonical-1212322131103131-3331322002112210-3302312013303101-2022332233000021-2331331310211102-2012232111300232-1232100300131230-3102300330203330"></a>

Type: `"single"`. Computed.

Send this direct response in case of route match action is direct response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0312122022002003-1220101111313220-0030023102303013-2332313201002011-1012021212010303-3121320113112331-0132321231233322-2100311011101210"></a>

### Direct properties for `routes.route_direct_response`

<a id="canonical-1122333030303213-2302103111311320-0330203100031100-2211211113110203-2030010322221212-1322231213011213-1011202022323233-0011313231001112"></a>

#### `routes.route_direct_response.response_body_encoded` property

Type: `"string"`. Computed.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 65536
    },
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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1331003122112003-2233330230331213-0213233122301301-0103212310112113-3030300132232231-2012203111312333-3231122332311232-0030210200011012"></a>

<a id="canonical-0301032131321130-2102221332021301-1232112122112233-3321330330001320-0213112230331003-0230032233212021-0203030111111112-0003233000220120"></a>

#### `routes.route_direct_response.response_code` property

Type: `"number"`. Computed.

Response Code. Response code to send.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

<a id="canonical-3031003311233213-2113111021311003-2111031332300130-1111311130312100-2120221111012231-3331330300223001-2023020031123100-0201231213001103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_redirect` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.route_redirect

<a id="canonical-2102022120213232-1030011300302311-0311321010220223-3200101231120130-2100103001102211-2012101031333022-3332232323210112-3300121301223120"></a>

Type: `"single"`. Computed.

Route redirect parameters when match action is redirect.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

<a id="canonical-2312210110011202-0201310020312033-3301232321103022-1331310101311033-0000122312020122-0030330001112321-1332021230321230-3000200032213310"></a>

### Direct properties for `routes.route_redirect`

<a id="canonical-1333112031333303-2102222201203013-1320032333212121-2103101132312333-1320303130132010-3101212013012312-1332113113101102-0100121310012130"></a>

#### `routes.route_redirect.host_redirect` property

Type: `"string"`. Computed.

Swap host part of incoming URL in redirect URL.

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

<a id="canonical-0203021000131202-0301230010213332-2133320320301333-0123332320131101-3132213300232012-0302201202023130-0033120012000311-1200130121130130"></a>

<a id="canonical-1202323300123213-0311010001103023-1102313211103313-0233202232221020-1333201012111332-2322002110022220-2130130333200322-0123010131000221"></a>

#### `routes.route_redirect.path_redirect` property

Type: `"string"`. Computed.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

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

<a id="canonical-0223302113020133-3311020313211320-1023233220001201-3301213001122110-0032031001220231-1300010032031003-3032002220021330-1323213102301232"></a>

<a id="canonical-2102230013320232-1033203123120003-0203221322322201-2300322000230021-2311300313330111-3022103312312303-0003130012002230-3132001312102012"></a>

#### `routes.route_redirect.prefix_rewrite` property

Type: `"string"`. Computed.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

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

<a id="canonical-0301210131233012-2311101233232210-1222320003203320-3230132112222223-2003310120202031-0132232130322320-2210021103001221-2001332211010213"></a>

<a id="canonical-3321000011031301-3211023003020111-1201012231232030-3203313312020023-1000012021002120-0123133200112330-0111121300123303-1303023221130122"></a>

#### `routes.route_redirect.proto_redirect` property

Type: `"string"`. Computed.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "incoming-proto",
    "http",
    "https"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](data-sources--route--reference--group-003.md#canonical-3212312301222212-3230213212111021-2210000103300211-2132001001121130-0122030121113210-2220232312012230-0232002211300323-3033133132030320): complete subsection reference.

<a id="canonical-1110232033212233-2302201030122003-3010120323123111-1323212221231220-0220223211022233-3231330233133030-1012211130301111-2133021130002133"></a>

<a id="canonical-3220303000120031-3320333001131112-1232002232210000-1220332211111202-2020110023210130-2012012013112221-2210022232010300-3101211003031131"></a>

#### `routes.route_redirect.replace_params` property

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2102333210131021-1012211112331000-0212111313222313-0122033201002133-1332113322023102-3102030000211222-3223233301203022-2121222332120001"></a>

<a id="canonical-0312032320100330-1230323332232103-2011211101102213-3101031031133131-1111110001300203-1023222101000310-1233202012102000-0223022001020212"></a>

#### `routes.route_redirect.response_code` property

Type: `"number"`. Computed.

The HTTP status code to use in the redirect response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
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
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](data-sources--route--reference--group-003.md#canonical-3001133033010312-3112110313130000-0103031123012010-2010002101212033-1201320120320022-3223103111033311-0023212131113320-3031130220030310): complete subsection reference.

<a id="canonical-3212312301222212-3230213212111021-2210000103300211-2132001001121130-0122030121113210-2220232312012230-0232002211300323-3033133132030320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_redirect.remove_all_params` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_redirect](data-sources--route--reference--group-003.md#canonical-3031003311233213-2113111021311003-2111031332300130-1111311130312100-2120221111012231-3331330300223001-2023020031123100-0201231213001103)
- routes.route_redirect.remove_all_params

<a id="canonical-2032120322132110-1000112332320112-0310031222110120-2202130021203033-2122032200112200-2233012000301212-2303221033021232-1011122110332110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for remove all params.

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

<a id="canonical-3001133033010312-3112110313130000-0103031123012010-2010002101212033-1201320120320022-3223103111033311-0023212131113320-3031130220030310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_redirect.retain_all_params` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_redirect](data-sources--route--reference--group-003.md#canonical-3031003311233213-2113111021311003-2111031332300130-1111311130312100-2120221111012231-3331330300223001-2023020031123100-0201231213001103)
- routes.route_redirect.retain_all_params

<a id="canonical-2033321200101322-1032113232103113-3000112022122030-3230131201303010-1010030033011233-1121222123100323-1133313020021332-3231311103130002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for retain all params.

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

<a id="canonical-0122003031122022-2312121320313011-3232203000312111-2300213223121003-3333300031112101-3031102112131312-1230313311203332-3230301311131320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.service_policy` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.service_policy

<a id="canonical-0120211110302033-0112203233032231-3310003010012113-3123123220311133-2311331120231020-2023310002002100-1322033322100030-2211011113112032"></a>

Type: `"single"`. Computed.

ServicePolicy configuration details at route level.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-service_policy_choice": "[\"disable\"]"
}
```

<a id="canonical-3323133310312121-3322323123322013-3201212310020031-0323300022202200-3222123201032003-0201230103231213-2003202012000331-3312303132121011"></a>

### Direct properties for `routes.service_policy`

<a id="canonical-0103020030002230-3110133023123232-1123201313310303-1223212303110230-0012032210122132-3301011300031101-1201321010100311-2022202212311012"></a>

#### `routes.service_policy.disable_spec` property

Type: `"bool"`. Computed.

Exclusive with \[\] disable service policy at route level, if it is configured at virtual-host
level.

<a id="canonical-2130023122111302-3312131310223330-3111030320001222-1300131123332221-1300033010300132-3331301213032123-3013023033221232-0223231010031303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_exclusion_policy` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.waf_exclusion_policy

<a id="canonical-0102021320210111-2132312231202321-3013023220302010-0003020210112022-2020200210101302-2333300233103113-1131321320110012-3110003001331232"></a>

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

<a id="canonical-2323213331032302-2113131003213323-2003312221021201-3101022231113230-3220023100320030-1200312331120310-1100203331333333-2203022031113102"></a>

### Direct properties for `routes.waf_exclusion_policy`

<a id="canonical-1330111020203003-3313231120300213-0133202300130220-1233231202100320-1133331011211112-3122010203130312-2113030030332303-0002123321223002"></a>

#### `routes.waf_exclusion_policy.name` property

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

<a id="canonical-1230223032103331-2131013110202030-1010203010332101-3321203033111002-3113222200013311-1210031301322010-1233201322223302-1330022223222001"></a>

<a id="canonical-1321130301212031-3110233303012101-1311232222112121-2012203233111212-0322001231300122-3022112002032303-1000103221301122-1013310103020313"></a>

#### `routes.waf_exclusion_policy.namespace` property

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

<a id="canonical-0001332303212101-0002322132021222-2310330202221010-1221121012011120-1103321233231200-1120033111330313-2231210122233322-1123101111220111"></a>

<a id="canonical-2211321001323232-0312033310211133-0022320321013110-3232320221010031-0232123102200310-2322230132300001-1021000232310311-3011203031210210"></a>

#### `routes.waf_exclusion_policy.tenant` property

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

<a id="canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_type` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.waf_type

<a id="canonical-0000203232210111-1312212321222112-1133232323000121-0223302032303213-3021033121313200-1330222203301112-2110020213030231-1033301211332320"></a>

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

<a id="canonical-2232122022211110-0333010202033102-0220223213100022-0231030012102130-3030110211320002-1021130203002300-1033003000211211-1102232313232232"></a>

### Direct properties for `routes.waf_type`

- [app_firewall](data-sources--route--reference--group-003.md#canonical-2020000300332202-2313022220012333-0031313002211002-1203003321210130-2133210011020131-2022331332130112-3122033002213300-3020310233103332): complete subsection reference.

- [disable_waf](data-sources--route--reference--group-003.md#canonical-1320201023113001-3121120311102013-3030312122111212-1202223212111100-3103200212220332-1223330020120011-3320011033130113-2103131322333023): complete subsection reference.

- [inherit_waf](data-sources--route--reference--group-003.md#canonical-2101320330322032-3311333110132211-3223101103110030-0020132323000112-0110211120000303-3223313302010332-2211030221221303-2131323020032203): complete subsection reference.

<a id="canonical-2020000300332202-2313022220012333-0031313002211002-1203003321210130-2133210011020131-2022331332130112-3122033002213300-3020310233103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_type.app_firewall` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120)
- routes.waf_type.app_firewall

<a id="canonical-0020231331220323-3031331033100200-1120333103130222-1211103132113331-3010320123323121-0313201322312212-2010031112310211-0330130310103330"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0211223333032003-1222213131103301-0000000133000333-1021112230220030-2032321103223302-2311231311311032-2113023301123112-2000331013122322"></a>

### Direct properties for `routes.waf_type.app_firewall`

- [app_firewall](data-sources--route--reference--group-003.md#canonical-1131313111311232-3120310130010230-3312112323200010-3011033332233100-0032331002023211-0101101233000110-0120203210121320-0222031013231001): complete subsection reference.

<a id="canonical-1131313111311232-3120310130010230-3312112323200010-3011033332233100-0032331002023211-0101101233000110-0120203210121320-0222031013231001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_type.app_firewall.app_firewall` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120)
- [routes.waf_type.app_firewall](data-sources--route--reference--group-003.md#canonical-2020000300332202-2313022220012333-0031313002211002-1203003321210130-2133210011020131-2022331332130112-3122033002213300-3020310233103332)
- routes.waf_type.app_firewall.app_firewall

<a id="canonical-3333133223101013-2213111111313231-0012020331031330-0223031321232131-3303022200200131-0330201332023012-1100001222032202-2202330011023331"></a>

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

<a id="canonical-2122222012302220-1012223020231323-1112302002323003-2121131200320303-2111003301333023-1130130221003001-2013202333203300-0032001221100110"></a>

### Direct properties for `routes.waf_type.app_firewall.app_firewall`

<a id="canonical-2121213330030331-2200332111100100-3121131311211211-3010133023032333-1101300010220200-3232333101212131-1302231301032323-1131031110231221"></a>

#### `routes.waf_type.app_firewall.app_firewall.kind` property

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

<a id="canonical-3110120212000211-0221110303020000-0330132003200110-1302313121021130-1011330120000230-2121122333010111-1021321310322203-1120333002303213"></a>

<a id="canonical-0123103021321102-1332022201111122-2231131200322022-3031302000001201-0323031233332122-2111101321122033-1321310010111010-1112301200103221"></a>

#### `routes.waf_type.app_firewall.app_firewall.name` property

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

<a id="canonical-2112212222101111-2302011323133332-1030330011000032-3312203323222110-0322230312033001-1333032110211100-1132322323130120-2132302230320110"></a>

<a id="canonical-2232313020012230-2102013101013321-0200023202130220-3230012113321033-2103331012000322-3323223220131221-0110003013112110-3011222112233111"></a>

#### `routes.waf_type.app_firewall.app_firewall.namespace` property

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
  }
}
```

<a id="canonical-0232130132102221-1230300033232101-1300023002221300-3223300100330332-0230020201323120-0331232021101021-1102021222311321-3211122332022312"></a>

<a id="canonical-1310222233010113-0323321222003312-0222200133033110-2223211203111223-3121010202123021-3303233200111322-0232021123202300-0122033031022320"></a>

#### `routes.waf_type.app_firewall.app_firewall.tenant` property

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

<a id="canonical-2300302123331103-0220313323130323-2210230032300213-3123223001320032-1213003303311121-1132023111130221-2200123202333313-2320300033301002"></a>

<a id="canonical-2301212020301102-3323321000213332-3112123010002213-3211201213333211-2312330110201132-0100020331031231-1321302000200331-1330020000132023"></a>

#### `routes.waf_type.app_firewall.app_firewall.uid` property

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

<a id="canonical-1320201023113001-3121120311102013-3030312122111212-1202223212111100-3103200212220332-1223330020120011-3320011033130113-2103131322333023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_type.disable_waf` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120)
- routes.waf_type.disable_waf

<a id="canonical-1311301330230310-3322000023323223-0231222322120003-0133003210322311-1030011031130121-2212310322320133-3203333030201113-1313023003101221"></a>

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

<a id="canonical-2101320330322032-3311333110132211-3223101103110030-0020132323000112-0110211120000303-3223313302010332-2211030221221303-2131323020032203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_type.inherit_waf` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120)
- routes.waf_type.inherit_waf

<a id="canonical-0113313123212030-3112020123111331-0123333101202012-1330102303203232-1011233123333203-2330313020102010-0120222220011000-0213232212330100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherit waf.

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
