---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0022220000230303-2320222123002122-0303112220202301-1203113230332223-0323010011023312-3013313203312231-0013100132210222-0230222332331121"></a>

## routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes — tag_attributes / 001321200211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-1103203303113331-2302123302001100-1121020032121231-2222122210210133-2002002122010003-2232003301121213-3201221312313321-0202203023112121)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags](resources--http_loadbalancer--reference--group-024.md#canonical-0230323012012233-1331313230201303-2013102231310211-2300022200010131-0013100032031013-1323313021311010-2301021330031011-3003103033102121)
- routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes

<a id="canonical-3121300113022303-1010023332312023-2330202003302102-2103331113311202-2102131223301021-0100232102101131-0122221010002021-3202020102013203"></a>

Type: `"object"`. list nested block, Optional.

Add the tag attributes you want to include in your JavaScript tag.

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tag_attributes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103332200311222-0010133221101320-1333313031011211-3012223301031302-0331120030031210-1211003103213200-3220100312312013-1011333333013032"></a>

## Direct properties — tag_attributes / 001321200211 / 3

<a id="canonical-0232100331022001-2332333203130221-0033203231203123-2213323100011233-0321101233331320-2133213032123012-3121221123210203-3010013030330332"></a>

<a id="canonical-2100333313222010-3012011220122102-1302100032110001-3023201011211312-3131233323111312-0323233131220011-1012023310300120-0100231103313013"></a>

## javascript_tag property — tag_attributes / 001321200211 / 4

Type: `"string"`. Optional.

\[Enum:
JS\_ATTR\_ID|JS\_ATTR\_CID|JS\_ATTR\_CN|JS\_ATTR\_API\_DOMAIN|JS\_ATTR\_API\_URL|JS\_ATTR\_API\_PATH|JS\_ATTR\_ASYNC|JS\_ATTR\_DEFER\]
Select from one of the predefined tag attributes. Possible values are \`JS\_ATTR\_ID\`,
\`JS\_ATTR\_CID\`, \`JS\_ATTR\_CN\`, \`JS\_ATTR\_API\_DOMAIN\`, \`JS\_ATTR\_API\_URL\`,
\`JS\_ATTR\_API\_PATH\`, \`JS\_ATTR\_ASYNC\`, \`JS\_ATTR\_DEFER\`. Defaults to \`JS\_ATTR\_ID\`.

Upstream description:

Select from one of the predefined tag attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "JS_ATTR_ID",
  "enum": [
    "JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3101202021202120-1022113031203323-3022020220313102-3030210203020111-1121331131331202-3222230210133333-3100111310120011-0022310133012310"></a>

<a id="canonical-0232120203332211-0201132100023033-0012313231131313-1211131321220013-1000123200100031-2312222012110032-0000000332023021-1223033102113010"></a>

## tag_value property — tag_attributes / 001321200211 / 5

Type: `"string"`. Optional.

Value. Add the tag attribute value.

Upstream description:

Add the tag attribute value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-3021132011020000-1023113220233023-0222210111202001-2331120302012113-3321203010200322-1020111002011332-2203122131112023-0302301101032033"></a>

## Next pages — tag_attributes / 001321200211 / 6

- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags](resources--http_loadbalancer--reference--group-024.md#canonical-0230323012012233-1331313230201303-2013102231310211-2300022200010131-0013100032031013-1323313021311010-2301021330031011-3003103033102121)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1211001130101123-0202021333221102-1013312002223231-0033202013223031-2322012302201213-0221112223001103-3230112102023130-3102232202213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133013003112221-2112101211021333-3332332112021231-1231330011030333-3310210023103300-0201133303223320-2302302103211133-0221221120002303"></a>

## routes.simple_route.advanced_options.buffer_policy — buffer_policy / 212011233213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.buffer_policy

<a id="canonical-0100003330320311-1130013313132330-1332211133101210-0112202221322202-2112023010123231-3132013111213022-1220001332310301-0223131121330331"></a>

Type: `"object"`. single nested block, Optional.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Upstream description:

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

Terraform syntax:

```terraform
buffer_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010012232101312-0212203313301302-0003020030111100-2302330201232231-0112330032030010-0310303120312232-2010233302310332-0112302212123332"></a>

## Direct properties — buffer_policy / 212011233213 / 3

<a id="canonical-1111222330201102-0010202011110122-0232230023100020-3203133201002332-0303223103122202-0032231023123021-2221231130023213-2232120232022021"></a>

<a id="canonical-3332123320111012-3022203202320102-3021213001000330-1131312001102322-0213222102231203-2230103021030101-2313311100032010-3032010122300220"></a>

## disabled property — buffer_policy / 212011233213 / 4

Type: `"bool"`. Optional.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

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

<a id="canonical-3210223022200001-0032230303021123-3201122300320012-1110002103233101-1300323231130023-3102100102033300-0331232303103001-0132232103233300"></a>

<a id="canonical-0101311120031101-1001231202323103-0001020212202322-1313120030112330-2302302220102310-1030111320021300-1020131311301113-3302130131223211"></a>

## max_request_bytes property — buffer_policy / 212011233213 / 5

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-3032221221222202-3203312301132210-3133211203002101-3110301002213011-3211113303013100-0301331201312131-0230111132231301-0032321103203103"></a>

## Next pages — buffer_policy / 212011233213 / 6

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3020133322310122-3120022303032302-1333013033231033-0203223021012000-1333112131131133-1101022331131132-0202323320200021-2220033102132032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101130212202302-1223233221000110-0310023101323113-0121201200210230-3202233111123123-2001100012203203-3202302312013320-2313311023123131"></a>

## routes.simple_route.advanced_options.common_buffering — common_buffering / 201011331221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.common_buffering

<a id="canonical-3201331033013213-1132133200002311-2213131321030011-2213222003212323-0000231132330102-0313011232100201-2323110323302011-1012311002020020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for common buffering.

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

Terraform syntax:

```terraform
common_buffering = {}
```

<a id="canonical-1222213120203020-0123112330121112-3213300020233310-0013201203211122-0022031331133032-0220130111213112-3301303331230012-2002333022220202"></a>

## Direct properties — common_buffering / 201011331221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030113003101101-0302111230320010-0122200110123220-0003331202102033-1121010002132213-2133101330201302-0313202101322033-1031131031020101"></a>

## Next pages — common_buffering / 201011331221 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1220202221213121-2003330033111010-3112031032321133-0331322330011032-2100203210203032-1312312203321022-3331311132203101-1323310311023321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121303102013311-2121120210232021-2023001311010101-1310031302230001-3301031223311120-2100023021101223-1211002322330213-1013210112030111"></a>

## routes.simple_route.advanced_options.common_hash_policy — common_hash_policy / 121010321321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.common_hash_policy

<a id="canonical-0023222113102030-2133033113012302-0333322313321103-3221333321200100-1232100113213201-1011310221311332-1330002200013002-3032002111312231"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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

Terraform syntax:

```terraform
common_hash_policy = {}
```

<a id="canonical-0032213120032210-3120013113003112-0332311013301103-3301212322132233-0231030100302201-0232210233333003-3001221310013311-0010132013021313"></a>

## Direct properties — common_hash_policy / 121010321321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200011233302112-3213323012113223-0211331133123220-0030102332133110-2233013311001311-3032201311112232-2311111323023133-2123121213221021"></a>

## Next pages — common_hash_policy / 121010321321 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3222022011013211-0231322200331030-3003132300320323-0223133302201212-3011220203110302-1311320103313321-2022323330221202-0131123122131320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103133001033303-2311032132311131-2200232202330122-3131300230233322-0133233023020232-0021221101101111-1021103210222322-2121100212222323"></a>

## routes.simple_route.advanced_options.cors_policy — cors_policy / 123031000013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.cors_policy

<a id="canonical-2023331023001023-2020203132023101-2331310220020032-0123020311022201-2110332123100303-3223032113133301-3030112112230200-3020302212332000"></a>

Type: `"object"`. single nested block, Optional.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence. An example of an Cross origin HTTP request GET
/resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel macOS
X 10.5..

Upstream description:

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

<a id="canonical-3003310122030303-1011311030032320-0032222213302012-2112233100311332-2330213133303012-3132200322100313-1122221322112200-3311012113100200"></a>

## Direct properties — cors_policy / 123031000013 / 3

<a id="canonical-0222001133330130-3202233330130000-3303120213020120-1201312231203313-3322233010300130-0112031233222202-2110311230110102-1130333311230223"></a>

<a id="canonical-0230011133100001-1210011102221133-0133020200001203-3031212222330301-0312102301330012-0303221023303102-1023232101302021-3201011220003100"></a>

## allow_credentials property — cors_policy / 123031000013 / 4

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

<a id="canonical-0203100221310222-0231223021021223-3130012122223110-2323103311103232-3030131012110302-1013231323320000-0312223021203131-3122020012212132"></a>

## allow_headers property — cors_policy / 123031000013 / 5

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

<a id="canonical-3121313321023303-0203101121221200-3232021031311223-1231202313101120-0200310023233100-1321200121210031-2110031032101102-0100233201122213"></a>

<a id="canonical-3112322323332120-3201303113330320-0300330102113333-2031010230232212-1031102210022202-1122023213100320-3312000300300330-0321331133112003"></a>

## allow_methods property — cors_policy / 123031000013 / 6

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
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-0200322113021211-2201113221130210-1102113331011023-3021013102111312-0120213023333232-0022011101200121-3133102110303010-3002323131233233"></a>

<a id="canonical-3220100113130021-1203001020003121-3120002131223030-1020223111211001-0231323230122013-3302302300203311-0030012233120131-2222213200022000"></a>

## allow_origin property — cors_policy / 123031000013 / 7

Type: `["list", "string"]`. Optional.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-0231323002010212-0303022223211110-2313311100311312-0032213331322032-1220132202310022-1233201013011202-1202102100201033-2111332000213233"></a>

## allow_origin_regex property — cors_policy / 123031000013 / 8

Type: `["list", "string"]`. Optional.

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1011301312020212-1203111121310333-3331033003222231-3001223333320023-0311120231133200-0011032332100103-3103121311301313-0111211000311302"></a>

## disabled property — cors_policy / 123031000013 / 9

Type: `"bool"`. Optional.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

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

<a id="canonical-1333320202102121-2331202311110201-3130332021203020-0022322100331330-3022012320133013-1323200230311222-2021221223132113-0312120332211302"></a>

## expose_headers property — cors_policy / 123031000013 / 10

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

<a id="canonical-1311210313132213-3210132133203231-1032111210313312-3102310001303111-3133121130031012-1213201310133232-0103023321132210-3010313111030113"></a>

<a id="canonical-2120302111312333-1102101022111030-0202233311122221-3102321332011002-0121021111022210-0303220331301100-1130303022230211-2030002221231022"></a>

## maximum_age property — cors_policy / 123031000013 / 11

Type: `"number"`. Optional.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(-1, 86400),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0321103101303103-1310111021132023-0101222032001211-0111102301322330-2100312303302233-0231013201220233-1021320031331320-3322021200031323"></a>

## Next pages — cors_policy / 123031000013 / 12

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302332133031330-3310210020103232-2203033023231331-3331100010012101-3213023113221102-0210332210200212-3120322231121132-3001200011110231"></a>

## routes.simple_route.advanced_options.csrf_policy — csrf_policy / 311020101300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.csrf_policy

<a id="canonical-3210133010213111-1312130212033003-3022113113101331-1010102020220033-1010220233110232-3223132020131032-3310132330002033-2113303103211212"></a>

Type: `"object"`. single nested block, Optional.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host. 1. The origin that caused the user
agent..

Upstream description:

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "custom_domain_list"),
  validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "disabled"),
  validators.ConflictingObjectAttributes("custom_domain_list",
    "disabled")}
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
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330311103001232-2003032322200111-2221220213321223-3032031321312001-1122332113023233-3131010313333223-0132201120220113-3203003312211332"></a>

## Direct properties — csrf_policy / 311020101300 / 3

- [all_load_balancer_domains](resources--http_loadbalancer--reference--group-025.md#canonical-3102211303130002-0311302231123112-3021133021121320-2300112000101322-0022213122203220-3321203021313221-1030122323103302-1211130133103221): complete subsection reference.

- [custom_domain_list](resources--http_loadbalancer--reference--group-025.md#canonical-2232023010120200-1330112210123120-2122301123331132-3212302133120122-0002301100303000-2120003300300223-2211300033013032-0310121310232122): complete subsection reference.

- [disabled](resources--http_loadbalancer--reference--group-025.md#canonical-2213021333110233-0321023022220030-3221132311032333-0103000100120312-3322200031221311-2301022200220203-1302310322301221-1201100321332030): complete subsection reference.

<a id="canonical-3000022100212032-2212212330322301-3331223113221230-1103101213223013-3233220000321231-3032100333322010-3222310132321002-1220221312122020"></a>

## Next pages — csrf_policy / 311020101300 / 4

- [routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains](resources--http_loadbalancer--reference--group-025.md#canonical-3102211303130002-0311302231123112-3021133021121320-2300112000101322-0022213122203220-3321203021313221-1030122323103302-1211130133103221)
- [routes.simple_route.advanced_options.csrf_policy.custom_domain_list](resources--http_loadbalancer--reference--group-025.md#canonical-2232023010120200-1330112210123120-2122301123331132-3212302133120122-0002301100303000-2120003300300223-2211300033013032-0310121310232122)
- [routes.simple_route.advanced_options.csrf_policy.disabled](resources--http_loadbalancer--reference--group-025.md#canonical-2213021333110233-0321023022220030-3221132311032333-0103000100120312-3322200031221311-2301022200220203-1302310322301221-1201100321332030)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3102211303130002-0311302231123112-3021133021121320-2300112000101322-0022213122203220-3321203021313221-1030122323103302-1211130133103221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110312212021003-0101021022130122-1203323300012303-2021132123223112-0312100203012033-1213003003111230-1103123013001133-3002121230123122"></a>

## routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains — all_load_balancer_domains / 330000331121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-025.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains

<a id="canonical-3330010002123331-0022223323002002-3212032222132113-3133203012310321-3231113310301302-2003312122222010-0122012202020021-0030113232111333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all load balancer domains.

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

Terraform syntax:

```terraform
all_load_balancer_domains = {}
```

<a id="canonical-1232113332333111-3012321121211110-0211221321310111-2200303320321031-3133101202320200-3120303103211333-3212120120232121-0212103113102123"></a>

## Direct properties — all_load_balancer_domains / 330000331121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221000023103212-1001320130302130-0030113333203101-0032330313103020-3230033212321222-0211323220212011-3331122320122111-2103112232221333"></a>

## Next pages — all_load_balancer_domains / 330000331121 / 4

- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-025.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2232023010120200-1330112210123120-2122301123331132-3212302133120122-0002301100303000-2120003300300223-2211300033013032-0310121310232122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023230132321033-0101132302132103-0132322300300001-1302011001111303-1023311003231001-2123130210122232-3112110031123300-1223321232322103"></a>

## routes.simple_route.advanced_options.csrf_policy.custom_domain_list — custom_domain_list / 222232001101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-025.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- routes.simple_route.advanced_options.csrf_policy.custom_domain_list

<a id="canonical-2111121030123202-0103312021020101-3012203203132300-0111131221313003-1323333033023330-0212311110023113-2011120333300333-3001022323033110"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
custom_domain_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1022230123322111-1102133022121232-1130020003103023-0013210030330011-0000213332311232-1303303303003021-2120110121300200-3322321010201223"></a>

## Direct properties — custom_domain_list / 222232001101 / 3

<a id="canonical-2302133323021231-1023123012230012-3220033331021322-2033012111133121-3110212122013320-1320011012302003-3303101220133023-3201221033033110"></a>

<a id="canonical-0212221023102300-3131100303133312-3230310121103032-2023031311031320-3300232002223031-3012110033020231-1012131023213201-2031222133030032"></a>

## domains property — custom_domain_list / 222232001101 / 4

Type: `["list", "string"]`. Optional.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

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

<a id="canonical-2330002203230013-1012033021222210-3120013232303033-3322210002113013-2210201010103303-2210200212030001-2111232020003220-2300030010012003"></a>

## Next pages — custom_domain_list / 222232001101 / 5

- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-025.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2213021333110233-0321023022220030-3221132311032333-0103000100120312-3322200031221311-2301022200220203-1302310322301221-1201100321332030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233330323301022-0330222333001300-3022013202302321-2323222203210102-2101210210103120-0323221310300213-0313221103100221-0230320222322223"></a>

## routes.simple_route.advanced_options.csrf_policy.disabled — disabled / 033221222121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-025.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- routes.simple_route.advanced_options.csrf_policy.disabled

<a id="canonical-2200132233200213-3323322130131233-2112201110103322-0331210302310203-0321233032120302-3311231133210010-0111212232201221-1300003232330211"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disabled = {}
```

<a id="canonical-3321111322312020-2103011230122330-3000013012111212-2020022010111133-1311300113021303-1313032330222332-2303320200021112-1020330112002023"></a>

## Direct properties — disabled / 033221222121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120201013201313-2200322301100313-0113123120200211-0211110102103202-3313122001123121-0220301210301131-1220320212312100-1032322322120032"></a>

## Next pages — disabled / 033221222121 / 4

- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-025.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3013002323211230-2000321010322031-0202220333302003-3320223120203023-3311023301310130-1123220212320222-2303233031330210-0222030031022311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202220302232210-3203100103213110-1111223323303302-2330322020020100-2130320113101021-1303133311203203-1001102322203123-0113111230201333"></a>

## routes.simple_route.advanced_options.default_retry_policy — default_retry_policy / 212132012302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.default_retry_policy

<a id="canonical-3103011220122302-3321311323030103-1032101303211301-3103011110101031-0300022110300232-0103123021013002-3201132121210023-2302110330111320"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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

Terraform syntax:

```terraform
default_retry_policy = {}
```

<a id="canonical-1012210231003100-3320113320100323-1033213232113121-1001312110321331-2113022203200311-0201122200202130-0022132013000112-2010021122001011"></a>

## Direct properties — default_retry_policy / 212132012302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113220132000030-1102003033230002-2102301221312021-3000122020212003-2121000323231320-3322312100323330-1001212030010300-0333111312113133"></a>

## Next pages — default_retry_policy / 212132012302 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1223210220132003-1032031311131022-0221333010313232-3300232111321333-3020023100102303-3000311333203311-2323022130322120-1301313300201232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002330222322001-1110021222210211-3230123303220320-0012010310120321-1301231100113112-2322302121003011-0031200301211330-3320331023002312"></a>

## routes.simple_route.advanced_options.disable_mirroring — disable_mirroring / 320331332022 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_mirroring

<a id="canonical-1021031031132110-2033021030110230-0231111133321222-3131103331311021-3332130220020323-0222201323200231-0312010103333220-0200212111112210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable mirroring.

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

Terraform syntax:

```terraform
disable_mirroring = {}
```

<a id="canonical-2010312233230202-2201333320323003-1301301132313200-0332231302330103-1301233211003030-3013321001222010-2101021202233230-2022003321233123"></a>

## Direct properties — disable_mirroring / 320331332022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031202311231213-0022233013101200-1301300200210312-0130201230102211-2210013231122233-2003120310300312-3230222311212231-3232030112123010"></a>

## Next pages — disable_mirroring / 320331332022 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3101222121331231-0002101210102120-1000231102102220-2110231230123212-2111001231312132-2121211202233201-1300220103223003-2023201222120123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131002111302001-3201122321223132-2320010133002100-3331132102222122-1230303013233121-0310103332200003-2333001000110303-1221023310030321"></a>

## routes.simple_route.advanced_options.disable_prefix_rewrite — disable_prefix_rewrite / 301212201013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_prefix_rewrite

<a id="canonical-1222222120001310-3020213133313220-0021000331131010-3113133202110330-1221201321122132-2331133300020222-2220221223203102-2221233223111201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable prefix rewrite.

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

Terraform syntax:

```terraform
disable_prefix_rewrite = {}
```

<a id="canonical-3022133202012322-1101023001301133-2130132320020033-3033021333100322-0001300201111323-1230210330233131-0023001222130222-1023320323121211"></a>

## Direct properties — disable_prefix_rewrite / 301212201013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220010100231311-3201131133131333-2312320332302012-3222220232120023-0231233031211202-3223012321320021-2322331232202103-0123201200120321"></a>

## Next pages — disable_prefix_rewrite / 301212201013 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1333112331031332-3110230102112221-1123102212221222-1320333331030110-2102300032332222-3010102120200200-0200212202113120-3221312122220012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200013213323023-1322223223132031-2320011121222200-2202133200020003-1232121200302111-2102000032002123-1331301110210210-2230012020200003"></a>

## routes.simple_route.advanced_options.disable_spdy — disable_spdy / 012200020100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_spdy

<a id="canonical-0301033111111300-1200133121222020-0233012113003210-0010312030213210-0330310232113003-0103132321001201-1102321013201122-1022131111003311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable spdy.

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

Terraform syntax:

```terraform
disable_spdy = {}
```

<a id="canonical-1333111133211100-0121030131023010-0013102133321102-2323310031102300-1230122232223202-3100320110302303-1023223121321122-3212120123210131"></a>

## Direct properties — disable_spdy / 012200020100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320232131111331-3220201222301202-3231233131133001-2213001320030333-0301133321302303-1301220122200203-2010221122020132-0212122203220033"></a>

## Next pages — disable_spdy / 012200020100 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2012322210102001-1000202101032101-0301012012020133-2222012131120302-2033130323220023-0131111302321232-1103220102022322-0130330011322111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201300303213203-3220030130312133-3020032321332112-1133303111300322-1331131201331220-3120313122032033-3000321333032320-3123132013000323"></a>

## routes.simple_route.advanced_options.disable_waf — disable_waf / 132122122322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_waf

<a id="canonical-3303013011100012-0200023313012302-0002023003301032-3111022310002230-0112220121131100-1213103021322322-2210322122321100-3322103003022132"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_waf = {}
```

<a id="canonical-2010022303211211-1132221102013313-0111031000221122-1031000013020223-3213330021030331-2321310221310210-2032031130123103-1202331321023001"></a>

## Direct properties — disable_waf / 132122122322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310222001211011-1303323100102130-0310122330202230-2030221231222031-0332100111230033-3002313320031033-2121130232233310-0311321133311210"></a>

## Next pages — disable_waf / 132122122322 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1123013100222203-3012313301032221-2333300120200121-2031023303323312-2330102330202121-3313212031211202-1300121232122232-3230133031002030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132230333022111-0303200331130220-0232012032211103-1330213331001131-3013103020102131-0301110203230130-1222010331211101-3120310300232213"></a>

## routes.simple_route.advanced_options.disable_web_socket_config — disable_web_socket_config / 030203013013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.disable_web_socket_config

<a id="canonical-0013020202203110-3312230002300003-2330231002010231-1111020021200000-1032312100120321-0201310310023232-3223000123310102-0213233200030033"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_web_socket_config = {}
```

<a id="canonical-2101033012312320-3123233102321300-0013213232201013-2321323231321133-2101203132322031-1301331011223102-0122300331333012-3001101120112322"></a>

## Direct properties — disable_web_socket_config / 030203013013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102003320220202-3233310112033313-3113202321203312-2303121303221231-2220010232023000-0231032201120202-1120332231020122-0211320002223230"></a>

## Next pages — disable_web_socket_config / 030203013013 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3013100112201010-3320010213003210-2110103020232131-1233302023023333-1320111303223010-1220201100133030-2322330330102010-1111303301333333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300123200223233-0102122312101312-3012130130203131-3022123103001011-1012131130110133-0233212223212223-0203131103212202-2321122113321303"></a>

## routes.simple_route.advanced_options.do_not_retract_cluster — do_not_retract_cluster / 133213002020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.do_not_retract_cluster

<a id="canonical-3032010210001000-0132301310112113-3231312222323112-0320100231221110-1313022301112231-3320211323123121-3202333203021331-1011211031002103"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
do_not_retract_cluster = {}
```

<a id="canonical-2203021003002320-3320322330331120-3223331232010121-0001132222230231-1312332122030111-2123112023100010-0212322012101313-0103200322300001"></a>

## Direct properties — do_not_retract_cluster / 133213002020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123121312131021-2213220032330130-1023121122103320-3320001320002112-1013111003220131-0321332202130121-3132230222300303-2133302032103302"></a>

## Next pages — do_not_retract_cluster / 133213002020 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3220221103013030-1033023211020212-2313000220302030-0112212123203222-2223110003210120-1330110301122333-0120013212320200-2011130012212031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113113120300210-2222030331302022-2032123220302021-1122212303212113-2110022010322321-1122211312310332-3100130012223132-2232231302331101"></a>

## routes.simple_route.advanced_options.enable_spdy — enable_spdy / 133200311012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.enable_spdy

<a id="canonical-1012101022332202-2310312301302112-1301110311103223-0312211212223123-0320331111303232-3230130213121033-0020320312112021-2022113200301210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable spdy.

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

Terraform syntax:

```terraform
enable_spdy = {}
```

<a id="canonical-3000123121101220-3103022122200123-3120220311312233-1023120211203311-3133033100331113-3021313103112100-0203120332321213-1110010022113332"></a>

## Direct properties — enable_spdy / 133200311012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323101301302310-3031111213221100-1020112213231110-0131122220301300-3300121130323230-1232030110102100-0312102231211113-3330101222222033"></a>

## Next pages — enable_spdy / 133200311012 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2123231212310000-0130220212030233-2332222131111223-3330003103011023-0211122003332012-0200233023011101-3232021110002213-0302122311123131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112012320031233-2333201013000023-1010302301131000-0001023001013322-1131132122323020-0200022312131222-1111312002120021-1020222132301333"></a>

## routes.simple_route.advanced_options.endpoint_subsets — endpoint_subsets / 310130221023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.endpoint_subsets

<a id="canonical-1221121322322213-0221232101101003-3131212031322320-0000323232312133-3201011032032323-3200100001302133-1113231122011333-1332133101210333"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

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

<a id="canonical-0323122223121331-3223111302132320-0212031022100230-0132120211202021-2310321102023122-2033030310321313-0102130200311302-3212030001002223"></a>

## Direct properties — endpoint_subsets / 310130221023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130003103101301-3131331332232320-1200220230122000-2032312023232022-1220211232101220-3030210132002120-3322023002303331-3122111033322023"></a>

## Next pages — endpoint_subsets / 310130221023 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3101301002211102-3101130203012322-2333001122113100-1010021103332111-0203220210230210-1230033102313122-1102112111223121-2210120101222000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030302002031023-3111203111322310-2222202123121011-1223023233133130-3202001132222202-3101220013100000-3201323003232310-1020100130102223"></a>

## routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection — inherited_bot_defense_javascript_injection / 323330101121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection

<a id="canonical-1122201311111012-3332331300213133-0331001223212103-2210100300201122-0021211020120201-0011310112302132-0031001122302210-3233322011310332"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
inherited_bot_defense_javascript_injection = {}
```

<a id="canonical-3321100020332013-3033000322131111-1230121010031033-1312123301201203-3302301220021022-3003102311011231-1113232012233022-0212313020310112"></a>

## Direct properties — inherited_bot_defense_javascript_injection / 323330101121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320330012133112-2012032100220301-2001031330102122-1212302122221013-0232003221331331-2001231032302033-3311302110133013-2211233230220000"></a>

## Next pages — inherited_bot_defense_javascript_injection / 323330101121 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0213201113012322-1102122130221322-3332120321331122-3312320023133031-1013323333130130-0033111303102122-3133232312300201-0023120200310323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202020111012102-3201211321030010-3302001100023301-3300023223122010-3220122302313110-3121022123013312-0023232331231210-3301303122022220"></a>

## routes.simple_route.advanced_options.inherited_waf — inherited_waf / 312130011220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.inherited_waf

<a id="canonical-1213202313123210-3132132322133331-1121203130032333-1001231202200300-2211032020131300-0311123000323133-0312302203222331-0301233211313322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherited waf.

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

Terraform syntax:

```terraform
inherited_waf = {}
```

<a id="canonical-2321301221020000-1031113132320221-1120202033110222-3321232022311311-2230121000020022-3012032223303322-1030230322113113-1233030210211230"></a>

## Direct properties — inherited_waf / 312130011220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100022330110101-3020010333003313-0222030300321012-1110010200110320-1131212120312010-2023220222131230-1111113202233013-3222231330233323"></a>

## Next pages — inherited_waf / 312130011220 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2002012001013000-3310311312101011-3313033113321211-2131033121023113-1002001131311300-2001101231211002-3100031020332322-2132012212030231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201022320022312-2223200223122313-3032032111033223-1021323130232331-3112122030202103-0101123123033123-0300330200011310-2330201333323202"></a>

## routes.simple_route.advanced_options.inherited_waf_exclusion — inherited_waf_exclusion / 331002322310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.inherited_waf_exclusion

<a id="canonical-3223112102232310-3321103232021123-3331233121022220-0013021212303220-3230121322333000-2030221101321013-1031220232231203-0200203001323331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherited waf exclusion.

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

Terraform syntax:

```terraform
inherited_waf_exclusion = {}
```

<a id="canonical-0131223110333030-3030220302223002-1211021213232130-0013022031330220-0303133131220211-0033101232221023-0210000310322102-1000023302222021"></a>

## Direct properties — inherited_waf_exclusion / 331002322310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121223020201323-3300100000010033-1121101313123003-1010020333200223-2111003203102022-1112230011233130-3210002310021131-1213202000233330"></a>

## Next pages — inherited_waf_exclusion / 331002322310 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100332221110030-3001201303011101-2231212231032300-0232132330301212-2001322222130131-3333201200212012-3131303302000201-1233321032303130"></a>

## routes.simple_route.advanced_options.mirror_policy — mirror_policy / 111112311131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.mirror_policy

<a id="canonical-3203310310232132-1110311230130323-3100032331000310-1332003132203113-2013321032131232-1221311332102131-1210021022223002-1210211131322011"></a>

Type: `"object"`. single nested block, Optional.

MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is
'fire and forget', meaning it will not wait for the shadow origin pool to respond before returning
the response from the primary origin pool. All normal statistics are collected for the shadow
origin..

Upstream description:

MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is
"fire and forget", meaning it will not wait for the shadow origin pool to respond before returning
the response from the primary origin pool. All normal statistics are collected for the shadow origin
pool making this feature useful for testing and troubleshooting.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3211023202321323-2113130003023220-2330212300223233-0011201313230201-3031111311021012-2332022031220301-0201213023020031-0221121323311002"></a>

## Direct properties — mirror_policy / 111112311131 / 3

- [origin_pool](resources--http_loadbalancer--reference--group-025.md#canonical-1103223210103233-3323023333223103-3202033211110230-3123111220212303-1220320130301312-2303110030230323-3323221001100011-2221131001203201): complete subsection reference.

- [percent](resources--http_loadbalancer--reference--group-025.md#canonical-2313121201002301-1112103031031030-2213123011002321-0330220012031022-3302303212211102-3121203313102011-2121121230313002-3102200032303031): complete subsection reference.

<a id="canonical-1112123203020102-0211202302013122-2112211201011331-3231110222030212-2302101111000313-0221303200330203-3212301332330033-1210303330103313"></a>

## Next pages — mirror_policy / 111112311131 / 4

- [routes.simple_route.advanced_options.mirror_policy.origin_pool](resources--http_loadbalancer--reference--group-025.md#canonical-1103223210103233-3323023333223103-3202033211110230-3123111220212303-1220320130301312-2303110030230323-3323221001100011-2221131001203201)
- [routes.simple_route.advanced_options.mirror_policy.percent](resources--http_loadbalancer--reference--group-025.md#canonical-2313121201002301-1112103031031030-2213123011002321-0330220012031022-3302303212211102-3121203313102011-2121121230313002-3102200032303031)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1103223210103233-3323023333223103-3202033211110230-3123111220212303-1220320130301312-2303110030230323-3323221001100011-2221131001203201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310301012321302-3012000020211213-3001323230130101-0023110001010020-1020100213112230-0330310133132321-3110231003032231-0020233233010001"></a>

## routes.simple_route.advanced_options.mirror_policy.origin_pool — origin_pool / 102020012033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-025.md#canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023)
- routes.simple_route.advanced_options.mirror_policy.origin_pool

<a id="canonical-0330312013233303-3302331312302211-3130203301310003-3323332322132021-3113122133301332-3100000312203131-1220023302033213-1312100321202331"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
origin_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233013230300010-0231122103112002-2032232212110212-2330322210032221-2203001021201130-0120111310001002-0301100221220311-0300003210020201"></a>

## Direct properties — origin_pool / 102020012033 / 3

<a id="canonical-3211113123010223-1023002120123110-2303113110200102-2013031032121301-0323023200310200-3013001323231213-0011303123330111-1222312130011333"></a>

<a id="canonical-0302003331102210-1031302132212133-0210320322331310-1203223233300232-1130200132102133-0323020301210110-0333333123020333-1113013302132012"></a>

## name property — origin_pool / 102020012033 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0302032003200200-2201131213112220-0101002322123021-1302123202333210-2021101303031222-1203023133003212-1213132231302233-3132113003202213"></a>

## namespace property — origin_pool / 102020012033 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0002010030122003-2030003123030200-0001323321023001-3221000102031223-0010122213021322-0102210220312202-3111220010133323-3021222203312202"></a>

## tenant property — origin_pool / 102020012033 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3332003310202310-3231111322012122-3332211012131102-0021020313022222-0231113303300313-1012013303132211-1000030313301020-2033222311210211"></a>

## Next pages — origin_pool / 102020012033 / 7

- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-025.md#canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2313121201002301-1112103031031030-2213123011002321-0330220012031022-3302303212211102-3121203313102011-2121121230313002-3102200032303031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211221302013022-0010303233000013-3213203023312110-2331011030110300-0122231003123101-0221311031203020-2213013002103310-0220300220313203"></a>

## routes.simple_route.advanced_options.mirror_policy.percent — percent / 323131023111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-025.md#canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023)
- routes.simple_route.advanced_options.mirror_policy.percent

<a id="canonical-3333100001321201-2123213121321233-0212200010330223-1300330210212023-2031031221302133-2331330023102203-0332000030211210-2000032223012111"></a>

Type: `"object"`. single nested block, Optional.

Fraction used where sampling percentages are needed. Example sampled requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("numerator")}
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
percent {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022203002213123-0303331021233320-2312120103223133-2132120323211000-0130211003032312-1330313002023310-0010323321110310-3230133212323313"></a>

## Direct properties — percent / 323131023111 / 3

<a id="canonical-1002101303031213-0323213000330212-3331000001300302-2212110102211110-0123023102130113-2021212232100200-3322212213202130-1020301213033210"></a>

<a id="canonical-3120023220112133-0303113023203323-2332012333132233-1203211122011230-3032210130332101-0200102111312023-1312213332323101-1003100012331333"></a>

## denominator property — percent / 323131023111 / 4

Type: `"string"`. Optional.

\[Enum: HUNDRED|TEN\_THOUSAND|MILLION\] Denominator used in fraction where sampling percentages are
needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use
million as denominator. Possible values are \`HUNDRED\`, \`TEN\_THOUSAND\`, \`MILLION\`. Defaults to
\`HUNDRED\`.

Upstream description:

Denominator used in fraction where sampling percentages are needed. Example sampled requests

Use hundred as denominator Use ten thousand as denominator Use million as denominator.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HUNDRED",
    "TEN_THOUSAND",
    "MILLION"),
}
```

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

<a id="canonical-1303123103001030-0123020131211003-2120013011113232-0122001202322220-2201021110313022-2030221300203222-2323031221301103-0310331303213101"></a>

## numerator property — percent / 323131023111 / 5

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

<a id="canonical-0132323023313130-3210310332220222-1012333333310322-0233221101012233-2200230331121200-2103033000302211-2112131333311233-3333313312322133"></a>

## Next pages — percent / 323131023111 / 6

- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-025.md#canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2220333121303102-3132022333332033-2111322310103130-0230222232313212-1102113200030202-3202032303023333-2120330013220013-0120312220300111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213013122111111-2201300322221313-3301221330311301-0110330210133113-3120202110112120-1122022330320001-3211111223133113-3311211302122101"></a>

## routes.simple_route.advanced_options.no_retry_policy — no_retry_policy / 011222221020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.no_retry_policy

<a id="canonical-3223110100331210-0311003333323320-0032321010031313-2301320120121302-0333220301202313-2203311000223033-3333032123130023-3022213322223331"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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

Terraform syntax:

```terraform
no_retry_policy = {}
```

<a id="canonical-0122102001123003-0023222200003212-3232232303030302-3010133222331322-1303021001111010-2302321231001111-2302312300321233-2112011023230102"></a>

## Direct properties — no_retry_policy / 011222221020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021332321123101-2300222032021302-2003031133221111-3111132023321021-0103330221111023-2030133001011321-2223332102121120-1012321301112111"></a>

## Next pages — no_retry_policy / 011222221020 / 4

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0320010111202212-1023320033013103-0030311122023131-1323130311311223-3130123101023222-0210202202123330-0003303022231213-3303032312010332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310131301033321-3100012113110200-0203100220010203-0103101010121210-3300130331322132-1020021213013303-3123223111231110-2022303212131032"></a>

## routes.simple_route.advanced_options.regex_rewrite — regex_rewrite / 301002023313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.regex_rewrite

<a id="canonical-1332213201131233-1311103113232313-0003110113230122-0032003101331030-1322220113200113-0122021121110230-1031201222131010-3132223232111312"></a>

Type: `"object"`. single nested block, Optional.

RegexMatchRewrite describes how to match a string and then produce a new string using a regular
expression and a substitution string.

Upstream description:

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

<a id="canonical-1103023313300101-1133133011221330-3121121322111110-2100302312000330-3030232303331310-1103110231033312-3120213322120031-2203230220303133"></a>

## Direct properties — regex_rewrite / 301002023313 / 3

<a id="canonical-2312210120331321-1032020300102000-3332010231222123-0100121301312303-3013033311120332-0102013213132310-0313331222302330-2023212230212100"></a>

<a id="canonical-2223122022122333-3331201021321100-3331033133310223-0321300001312031-3123231300210323-0103313123211311-3311200201322120-3333223020012223"></a>

## pattern property — regex_rewrite / 301002023313 / 4

Type: `"string"`. Optional.

The regular expression used to find portions of a string that should be replaced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1122012311112000-2230023323312003-3121113201213201-0100301002223303-2323103122132001-2101130300202313-2013321102231210-0102020302333221"></a>

## substitution property — regex_rewrite / 301002023313 / 5

Type: `"string"`. Optional.

The string that should be substituted into matching portions of the subject string during a
substitution operation to produce a new string.

Upstream description:

The string that should be substituted into matching portions of the subject string during a
substitution operation to produce a new string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2021133020103000-0333230221001201-2112323120102111-1322222303231300-2231120012023001-3320020103333332-0323211102001312-2203031202221300"></a>

## Next pages — regex_rewrite / 301002023313 / 6

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022200000101330-2030003120211121-0101301210330201-0233102013101111-3330323233323310-3313001111230230-3132222101302102-2202230331200110"></a>

## routes.simple_route.advanced_options.request_cookies_to_add — request_cookies_to_add / 101133133233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.request_cookies_to_add

<a id="canonical-2023123122330012-1011112221023022-2310012330221032-2230313222233031-0033302110203301-3331032233211112-2302113000113003-3330003113232023"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

<a id="canonical-0221022211010232-0232210033000222-2112133100332110-2200312303020133-0100201223120102-3111111031022302-3210132121000110-3003000301020233"></a>

## Direct properties — request_cookies_to_add / 101133133233 / 3

<a id="canonical-1131032303302203-1213301003023000-2111120020000022-0210222233211112-2310210113113112-0112122003003220-1032131033121202-1211101022032010"></a>

<a id="canonical-1133100103113331-3301200103200002-2101033012101123-3002330033101020-0023313210212113-1010101300102212-1001333321021233-3122302312022213"></a>

## name property — request_cookies_to_add / 101133133233 / 4

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2122321013132100-1203203120022131-0212200030022100-1122100313003031-3223021030110012-2210311221002003-0030310222211332-1222331220023311"></a>

## overwrite property — request_cookies_to_add / 101133133233 / 5

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121): complete subsection reference.

<a id="canonical-2012233030000330-3220233123110232-2333031310332021-1133012323021311-1301222201131211-1301211302100110-3313123010113220-1103211303333312"></a>

<a id="canonical-3131000021020300-3121130222133002-2310023302112322-3113010310321032-2133220220222021-3203300321223122-2203023031110201-0332113332332201"></a>

## value property — request_cookies_to_add / 101133133233 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0201323130322033-3323213211030030-0031022023231223-0123233032211033-2300002221023022-2011013110020331-2220212320122002-3120112133312212"></a>

## Next pages — request_cookies_to_add / 101133133233 / 7

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120133230213313-0101030300023122-0121011332003121-3221213100211133-2322330001202303-1020313230020312-1322211310013310-3001031033220331"></a>

## routes.simple_route.advanced_options.request_cookies_to_add.secret_value — secret_value / 110310201232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value

<a id="canonical-2221112133101200-2321000221313010-3233000123000132-3333220220332012-2220301120332123-0303030011102322-0221103012033122-3313033011201220"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113131102210313-3201220011122210-2311210203012333-3110003330201321-2111320333201123-3011313223302033-3112113001332323-1113102321223130"></a>

## Direct properties — secret_value / 110310201232 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-025.md#canonical-0212231213230312-3102112110221211-1100022003031232-3320033210021133-3132120212202021-1320121122013221-2330103113101111-0102012112312132): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-025.md#canonical-2100100321020011-0120311233112012-3121323322113103-2310100301231321-3030210112013001-2301002311012321-0210103011301002-1211023132203232): complete subsection reference.

<a id="canonical-0002300032321002-1031000233112031-2220213301202303-3202100030210122-2323113031120101-1113112020211302-1112212103030230-1331110202020232"></a>

## Next pages — secret_value / 110310201232 / 4

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-025.md#canonical-0212231213230312-3102112110221211-1100022003031232-3320033210021133-3132120212202021-1320121122013221-2330103113101111-0102012112312132)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-025.md#canonical-2100100321020011-0120311233112012-3121323322113103-2310100301231321-3030210112013001-2301002311012321-0210103011301002-1211023132203232)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0212231213230312-3102112110221211-1100022003031232-3320033210021133-3132120212202021-1320121122013221-2330103113101111-0102012112312132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200202220202113-1233210102200032-0303032123211230-0202111233000330-3233000010032113-2202001232222102-2310112122300132-2203201132101302"></a>

## routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 301330311031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-2321123312320130-2111000211032220-1333121322132313-1131131203330213-2101113330320112-0233030331133212-3213000133011220-1311302102031003"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000000100231011-2301212013033322-3103002001323233-0121332300321130-3233211120021030-0133200101222223-3312130210323111-2010022220200232"></a>

## Direct properties — blindfold_secret_info / 301330311031 / 3

<a id="canonical-1302213220323032-2213302113221312-1100321203333312-0211030210031110-1133330101223231-1001331110210202-0232121203212031-1331123031220033"></a>

<a id="canonical-0012223231311133-0333210333132131-1120323223221312-2113220023112201-0113233223033021-2302022002001110-2112130202121022-0211310221230221"></a>

## decryption_provider property — blindfold_secret_info / 301330311031 / 4

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

<a id="canonical-3031231230232321-2202211213301132-2010033002031122-1121212201212003-0321300232233112-2103202301121321-3231211211211130-3211031200233223"></a>

<a id="canonical-0312122221231202-3131233210013133-0003322213210333-0021301020312002-2211221212320030-3020313022000201-3332332033232321-2220102130020311"></a>

## location property — blindfold_secret_info / 301330311031 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-2231012030000121-1030312220211210-0102130001311301-2120021100022123-2023000313121213-2003320333302013-2203312112021223-2213121000121211"></a>

<a id="canonical-2230312123122300-2211011220220130-2130020331130320-1030323302003103-3113010202203002-2311031212102100-3223313232122001-2013203102233322"></a>

## store_provider property — blindfold_secret_info / 301330311031 / 6

Type: `"string"`. Optional.

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

<a id="canonical-0210030312313022-3210012110330112-0233201110033100-0200223110232232-1011220021033201-3331110022020331-1323300002200032-2103212231221212"></a>

## Next pages — blindfold_secret_info / 301330311031 / 7

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2100100321020011-0120311233112012-3121323322113103-2310100301231321-3030210112013001-2301002311012321-0210103011301002-1211023132203232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300110310302021-1310323313121210-0120011111012211-1332132311312110-3233030302033030-1010003002131212-0303030201203101-1010103032310320"></a>

## routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 323210023033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2110221023333101-2333331023011323-3300301203303210-1212033110112011-1020130110030311-1033313332230122-2311312010322022-0013101011333132"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111102033300320-2001111311230313-0021030220122302-3020222222002003-2201100300312200-2102312211023323-3303232030310333-1000122123233322"></a>

## Direct properties — clear_secret_info / 323210023033 / 3

<a id="canonical-1203202002331030-1302322301331032-1232033122211020-3313102011212222-0130033230303220-1200233231100122-2101310011003312-0122111123003003"></a>

<a id="canonical-1231211323132132-3110310101301033-3312101200303032-3230101101321332-3021310132011233-3110112310323331-1132300020103202-1121223120220132"></a>

## provider_ref property — clear_secret_info / 323210023033 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0320030231133333-1122012032301220-1000332032033100-1030203132012230-1212130021001100-3110022110103210-3203210323000102-0003130022123020"></a>

<a id="canonical-2213211201322110-2011121120303220-3013311123211333-0121110331212220-0001102122333101-3320020003113221-0310010233323313-1031312003323023"></a>

## URL property — clear_secret_info / 323210023033 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-0112233200121333-1123310210311200-3013320031122003-1301011012110133-2230030120331311-1020021123331211-0021220211303101-0330010200332310"></a>

## Next pages — clear_secret_info / 323210023033 / 6

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3232131122101331-1302212131220203-2030233131233200-2223210010302311-0103202032210231-3032031202222333-0322102313210230-3220202311121121)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301023222200221-1213303102301331-3110203022303312-3013221121210030-1322012123013133-0120221102313122-1303003213002132-0220121300321021"></a>

## routes.simple_route.advanced_options.request_headers_to_add — request_headers_to_add / 332201300020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.request_headers_to_add

<a id="canonical-2022011310031200-0013303231220310-0010311312321030-2121221330003233-2132210232231103-0223031231323331-0303001302312200-3001132120123023"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

<a id="canonical-1300302012033230-3300001333212232-3213302012011221-0231111020002122-1322321122020023-2313020112200010-0001130322130030-2022330231302003"></a>

## Direct properties — request_headers_to_add / 332201300020 / 3

<a id="canonical-2223222320001202-0311023120022130-3202111032220231-0010132132011203-0132331000121303-3301001213331220-2211111301330213-3220222211200023"></a>

<a id="canonical-1322323303122211-1111330021310333-1320313002003232-0310323023021133-2321302312022231-2312233231003222-2310303312000323-2100323220331030"></a>

## append property — request_headers_to_add / 332201300020 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0313112313223000-2332111113001113-1113222332211330-1011001220330131-2102201123033231-1012321203222210-3313311203232023-1032031333332220"></a>

## name property — request_headers_to_add / 332201300020 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120): complete subsection reference.

<a id="canonical-2311012021002332-2130231220030200-1113301213210230-3202032333030013-2111111303033303-0232310320232303-2113012011312131-0010212130033231"></a>

<a id="canonical-1313113000320220-0220330102102211-2131210002233212-1102231122103023-2033110030021203-1103100233100233-1313311132230232-2302201312102211"></a>

## value property — request_headers_to_add / 332201300020 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0331312302310332-2110111030123330-3210000232100333-1232002023122011-3212023333132330-3001022121002203-1022322130101133-1032313203313222"></a>

## Next pages — request_headers_to_add / 332201300020 / 7

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101133230113233-0302320303220001-1210002210013132-2321123100302033-2102113301222213-2200323210030200-0133323132231102-2333300103101013"></a>

## routes.simple_route.advanced_options.request_headers_to_add.secret_value — secret_value / 101332302001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value

<a id="canonical-0103233013032131-3102300312132100-3232201300122033-0013003323321312-0223302212101113-1333111120303012-2231331233310322-1301201022002320"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002120002213101-0113131323103020-1312313232231300-2321211203113011-3222033113132322-2012130303233020-0333323120121120-2112021020333101"></a>

## Direct properties — secret_value / 101332302001 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-025.md#canonical-2010300000023322-0213320200330302-1131301130213121-2121133311202221-1220022210031330-0132020101323311-0330010132222303-1023120132222101): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-025.md#canonical-3333203031110332-1013301011031232-1122112311011033-3030310211111231-2320012202112313-1233233131230221-0102310210020010-1113102130122120): complete subsection reference.

<a id="canonical-1222003330021012-2010210323012111-1222100230303323-2312122021100223-0230012223322201-1211010102110301-3331310213022303-1101220203311203"></a>

## Next pages — secret_value / 101332302001 / 4

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-025.md#canonical-2010300000023322-0213320200330302-1131301130213121-2121133311202221-1220022210031330-0132020101323311-0330010132222303-1023120132222101)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-025.md#canonical-3333203031110332-1013301011031232-1122112311011033-3030310211111231-2320012202112313-1233233131230221-0102310210020010-1113102130122120)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2010300000023322-0213320200330302-1131301130213121-2121133311202221-1220022210031330-0132020101323311-0330010132222303-1023120132222101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203231113230310-1331231300211100-3111033032310323-2030300130322012-0301033112202002-2332131332033301-2313100230231023-0222332311001123"></a>

## routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 001102333323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0331110020333233-1120231003031120-1001123130102012-1012232131300200-0212201213113003-2220123333233300-3323333101010322-2022232123233221"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111100202110020-1202202111210001-0232201330121200-1011110232010222-1312300033333213-0333310231123111-1113013201130123-2221302113202022"></a>

## Direct properties — blindfold_secret_info / 001102333323 / 3

<a id="canonical-0012123013230032-1331120113202120-2230313323302000-3321110220111020-1101012322133111-0030131201030303-1113132132321122-1203223202313010"></a>

<a id="canonical-0321210121020233-1220231303131033-0302022220232022-0110332021122320-1200212032012203-0230001323221213-2233201200222021-0023310222030021"></a>

## decryption_provider property — blindfold_secret_info / 001102333323 / 4

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

<a id="canonical-0033333121002323-3000030022322320-3221033131123210-1230123230120210-0301013312123322-3130300033323003-3000330032310201-0321203011111010"></a>

<a id="canonical-3123310322330101-0232323213322001-0003023033323100-2133312230332321-1122220022033012-2122023212023121-3322331121230110-3221221332333123"></a>

## location property — blindfold_secret_info / 001102333323 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-1210311213320121-1332222211221220-0130313311221332-0131023303121203-2200133203202300-3031101120302302-1012000022211133-0113003331232011"></a>

<a id="canonical-2003110200221001-2330022120312111-1220332230230201-2020312001120221-0200003112223323-1311110001010113-3133120302032330-1103331221203113"></a>

## store_provider property — blindfold_secret_info / 001102333323 / 6

Type: `"string"`. Optional.

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

<a id="canonical-0011220321021232-2001320033300113-1130121303300020-0001011031221223-0312001303202010-2103102011231133-2300202013301111-1221103023002310"></a>

## Next pages — blindfold_secret_info / 001102333323 / 7

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3333203031110332-1013301011031232-1122112311011033-3030310211111231-2320012202112313-1233233131230221-0102310210020010-1113102130122120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012121223223110-3220132031031200-3233321011112031-0300023223103003-1232100032201223-3210322331330110-2113212001011003-3313331113332022"></a>

## routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 232020112233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-1101111300232320-0212013002313233-3301002021322303-0320032200233011-0212202323331121-2031332022330031-1333131002310033-0110031110103331"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221323101033333-3311301013113113-0122002111310200-3330033320223033-0322020230032133-0223230113002232-3321101331230120-3003320112003102"></a>

## Direct properties — clear_secret_info / 232020112233 / 3

<a id="canonical-1310303211021213-1331123110201200-2003332020310022-0300000312033110-1013020023332312-2101320002111002-0003001231231221-1332322230323133"></a>

<a id="canonical-0323130221110231-1131032101123022-2133223102230101-1123113222313030-0103111333232322-3020121103232230-1103212000120312-0013130120323310"></a>

## provider_ref property — clear_secret_info / 232020112233 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2230012223313332-1201321331230010-3030332322300321-2313111022202022-2002222333331103-3112033210322113-0010023022003013-3033021003132002"></a>

<a id="canonical-3032313033222002-3232100121102100-3330020010012312-2033333311110230-2332331233223300-3032011213103223-0300033120131013-3330021021022033"></a>

## URL property — clear_secret_info / 232020112233 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-2000000113231032-2100033003131320-0102302321000300-3022323020100201-2020232123311011-1231210000301222-1220012111201323-1122222332330211"></a>

## Next pages — clear_secret_info / 232020112233 / 6

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-025.md#canonical-3033132120210012-2302112221301321-0110020012012002-2111110021330131-1110112113102200-1321101120121202-1233300210201120-1311112013232120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113103010301100-2233121001213031-1023011100311220-1110021210202332-2322032220031100-3001030100303301-0312031112232311-0102311101313102"></a>

## routes.simple_route.advanced_options.response_cookies_to_add — response_cookies_to_add / 003002030131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.response_cookies_to_add

<a id="canonical-2202302222030131-2310030001022313-0132123230331302-2333213311020302-1102132333301132-1301021330220120-2122002133211030-1221303312023122"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_domain",
    "ignore_domain"),
  validators.ConflictingListObjectAttributes("add_expiry",
    "ignore_expiry"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_partitioned",
    "ignore_partitioned"),
  validators.ConflictingListObjectAttributes("add_path",
    "ignore_path"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "secret_value"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "value"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

<a id="canonical-2122100303213132-0122332032230320-1310331203101202-0312033233313103-3122231313310233-1120233102220322-1021221120211201-0100122301003122"></a>

## Direct properties — response_cookies_to_add / 003002030131 / 3

<a id="canonical-2223300310031313-2033311311202321-3131200332101212-1233110331233031-3021122100332331-3120011311333201-3123133322300223-2231330233332020"></a>

<a id="canonical-0201133003000312-3201210031122103-1131212311112333-0120210233222032-0132222100110131-3221330103002301-2302302213200323-2101332312221312"></a>

## add_domain property — response_cookies_to_add / 003002030131 / 4

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0223222311201000-0333333111113200-1023331221133033-0102201031012330-0313310010310201-2301000000012330-2003003033301233-1101311231102222"></a>

## add_expiry property — response_cookies_to_add / 003002030131 / 5

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](resources--http_loadbalancer--reference--group-025.md#canonical-2313030222311331-0310301120323200-3212110122231121-2111000013022300-2100300222033231-2110122110010311-3021221230021320-0223003232223323): complete subsection reference.

- [add_partitioned](resources--http_loadbalancer--reference--group-025.md#canonical-1323232012020231-1323232033333100-0022200303021233-1132201023301022-1311100300203202-1032201201220220-2212131302003003-2302012011131123): complete subsection reference.

<a id="canonical-2122310233111221-0023021131031103-3202000210110231-2000103032001030-1022202202213222-0220130330120031-0022100101323231-2330213030323230"></a>

<a id="canonical-2313100113221130-0112330031101310-3233333010220102-0101232220013003-3031113011101301-3202313121210312-0132003223023220-2302001110310212"></a>

## add_path property — response_cookies_to_add / 003002030131 / 6

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](resources--http_loadbalancer--reference--group-025.md#canonical-0213002321303031-3200322021113130-3120313220233100-0313331231203332-0301220311230132-0033000312233230-0221310333001103-3012223000332102): complete subsection reference.

- [ignore_domain](resources--http_loadbalancer--reference--group-025.md#canonical-1231230310231311-3031012003130120-0203331001131021-3303121233230332-0000320021103021-2122031210201132-0221212102002000-1020101120000013): complete subsection reference.

- [ignore_expiry](resources--http_loadbalancer--reference--group-025.md#canonical-3000330231110221-3010013022021133-0111331201233132-2023302321123223-1133033122331331-1310300003223121-1301302023322230-0130102220120122): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-025.md#canonical-3033003021303110-2221320232111022-3210213120013012-1202002313131101-0111030110111110-3121230031330010-1320203302233301-2101101020133230): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--reference--group-025.md#canonical-3332100132012110-1201031310330221-3300001231012323-2123023123313110-1320002032122101-2110030113013200-3000320302313022-0313132123300010): complete subsection reference.

- [ignore_partitioned](resources--http_loadbalancer--reference--group-025.md#canonical-0011012023201321-2203022033203003-2130131111312003-1202010321023003-2031320331333121-1133130221220202-3130312111111113-3112011213003301): complete subsection reference.

- [ignore_path](resources--http_loadbalancer--reference--group-025.md#canonical-3130303020100023-1032211010221013-0011013101110003-2100102032021202-2002232131231101-1112031231033121-3231013030113202-1311212333001023): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-025.md#canonical-1113003101311201-1030323301021312-3320212231220300-2131010001120201-1300010123101211-1021011203220320-2223002133203022-3330320022311103): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-025.md#canonical-0121100310301220-3302022132310000-0002023211103222-1331230311320323-0101322131033012-3212012122033212-0102321221032303-1231203220020120): complete subsection reference.

- [ignore_value](resources--http_loadbalancer--reference--group-025.md#canonical-0330032312013200-2213310211110302-3311111113310303-1020202111102302-2331212002020232-1032313002232013-1211000230200012-1203301033112132): complete subsection reference.

<a id="canonical-2331330333332022-3313133111012021-2221303223013301-2013323220323210-3101023033300033-2022202103103233-3120012221103003-3333110313110321"></a>

<a id="canonical-0212222210031222-1212231220011123-2201322302023120-3321230213303303-1100232031030221-1222322103203311-3100300010310112-1122011101232313"></a>

## max_age_value property — response_cookies_to_add / 003002030131 / 7

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-0022013303312003-0001333102103022-0033131321220310-2321221001202003-1131023330011210-3133003132313012-3111210123001223-3232103332003212"></a>

<a id="canonical-3120323232301130-3203221230333300-1301122110112203-3313102032202233-1320110321201232-3101320322232210-1130230332010301-2232022000131113"></a>

## name property — response_cookies_to_add / 003002030131 / 8

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3213310201000011-1103322322033033-0031020213032331-3030322110212131-2102123321331320-0323001001113113-2330102320321333-1022113311230120"></a>

## overwrite property — response_cookies_to_add / 003002030131 / 9

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](resources--http_loadbalancer--reference--group-025.md#canonical-2122100130200302-2010033202110221-2101223332201321-1312310321132322-1313033110030110-2323000311312221-3030331110000103-0032112303312122): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-025.md#canonical-3213021033222202-3012330311220203-1130031023201103-1302120001111331-3030131223120232-2212010231121301-0013330311320112-1233330332212232): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-025.md#canonical-0111123301032121-3302113132300132-1031023023123223-1322221302020311-0103222130111210-2302010031223231-2032223221033320-2033012013022313): complete subsection reference.

- [secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-0203212013131210-3112233202311012-1211021022312223-2301111332221021-3300232131111021-2003332002013101-3123100331013113-1132013121310321): complete subsection reference.

<a id="canonical-2021312331201033-3233002311300333-0122132332121013-2302323033100220-0302020022001300-1203033020002030-0312031321222123-2131021211120333"></a>

<a id="canonical-1310312322331132-1023301032120231-0302320320131333-0122100210030020-3313023111203113-3010213102213231-1222022032213310-0330202122123133"></a>

## value property — response_cookies_to_add / 003002030131 / 10

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0103003212222300-0023213112131001-2100112010300001-2312332100212121-1132033101010112-0003123100021310-3132023103213200-0223030201321122"></a>

## Next pages — response_cookies_to_add / 003002030131 / 11

- [routes.simple_route.advanced_options.response_cookies_to_add.add_httponly](resources--http_loadbalancer--reference--group-025.md#canonical-2313030222311331-0310301120323200-3212110122231121-2111000013022300-2100300222033231-2110122110010311-3021221230021320-0223003232223323)
- [routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned](resources--http_loadbalancer--reference--group-025.md#canonical-1323232012020231-1323232033333100-0022200303021233-1132201023301022-1311100300203202-1032201201220220-2212131302003003-2302012011131123)
- [routes.simple_route.advanced_options.response_cookies_to_add.add_secure](resources--http_loadbalancer--reference--group-025.md#canonical-0213002321303031-3200322021113130-3120313220233100-0313331231203332-0301220311230132-0033000312233230-0221310333001103-3012223000332102)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain](resources--http_loadbalancer--reference--group-025.md#canonical-1231230310231311-3031012003130120-0203331001131021-3303121233230332-0000320021103021-2122031210201132-0221212102002000-1020101120000013)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry](resources--http_loadbalancer--reference--group-025.md#canonical-3000330231110221-3010013022021133-0111331201233132-2023302321123223-1133033122331331-1310300003223121-1301302023322230-0130102220120122)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly](resources--http_loadbalancer--reference--group-025.md#canonical-3033003021303110-2221320232111022-3210213120013012-1202002313131101-0111030110111110-3121230031330010-1320203302233301-2101101020133230)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age](resources--http_loadbalancer--reference--group-025.md#canonical-3332100132012110-1201031310330221-3300001231012323-2123023123313110-1320002032122101-2110030113013200-3000320302313022-0313132123300010)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned](resources--http_loadbalancer--reference--group-025.md#canonical-0011012023201321-2203022033203003-2130131111312003-1202010321023003-2031320331333121-1133130221220202-3130312111111113-3112011213003301)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_path](resources--http_loadbalancer--reference--group-025.md#canonical-3130303020100023-1032211010221013-0011013101110003-2100102032021202-2002232131231101-1112031231033121-3231013030113202-1311212333001023)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite](resources--http_loadbalancer--reference--group-025.md#canonical-1113003101311201-1030323301021312-3320212231220300-2131010001120201-1300010123101211-1021011203220320-2223002133203022-3330320022311103)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure](resources--http_loadbalancer--reference--group-025.md#canonical-0121100310301220-3302022132310000-0002023211103222-1331230311320323-0101322131033012-3212012122033212-0102321221032303-1231203220020120)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_value](resources--http_loadbalancer--reference--group-025.md#canonical-0330032312013200-2213310211110302-3311111113310303-1020202111102302-2331212002020232-1032313002232013-1211000230200012-1203301033112132)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax](resources--http_loadbalancer--reference--group-025.md#canonical-2122100130200302-2010033202110221-2101223332201321-1312310321132322-1313033110030110-2323000311312221-3030331110000103-0032112303312122)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_none](resources--http_loadbalancer--reference--group-025.md#canonical-3213021033222202-3012330311220203-1130031023201103-1302120001111331-3030131223120232-2212010231121301-0013330311320112-1233330332212232)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict](resources--http_loadbalancer--reference--group-025.md#canonical-0111123301032121-3302113132300132-1031023023123223-1322221302020311-0103222130111210-2302010031223231-2032223221033320-2033012013022313)
- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-026.md#canonical-0203212013131210-3112233202311012-1211021022312223-2301111332221021-3300232131111021-2003332002013101-3123100331013113-1132013121310321)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2313030222311331-0310301120323200-3212110122231121-2111000013022300-2100300222033231-2110122110010311-3021221230021320-0223003232223323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123023200200310-0001112221032220-2233020011110200-3033130312131112-1212000000311312-3303303100003232-1330000301212122-1203030121200223"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.add_httponly — add_httponly / 132101230200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.add_httponly

<a id="canonical-1213203130131311-0320303331203300-1100201221200310-1121213002001330-1220321233222002-3033311021113201-2102200302022331-2310223231010123"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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

Terraform syntax:

```terraform
add_httponly = {}
```

<a id="canonical-1322201130232002-2330122111320233-1210321100112220-0302310132101230-2303111200331101-1211300111003011-2102303112120312-1030210201230010"></a>

## Direct properties — add_httponly / 132101230200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333123033223101-1100211002133102-2201330132213331-2111233023000232-2202013130003200-2113033113302010-2300102211302220-3232002332322023"></a>

## Next pages — add_httponly / 132101230200 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1323232012020231-1323232033333100-0022200303021233-1132201023301022-1311100300203202-1032201201220220-2212131302003003-2302012011131123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331201232223332-3030320311033113-1310311012113322-1101123211032003-2302010110023232-0333120032332022-1032330011300121-3231321223331330"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned — add_partitioned / 312200323320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned

<a id="canonical-0331010332023100-1232113001202213-0310220321001301-3121003003230311-1113112012213221-2331021120003033-3023232020302212-3121000320022130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add partitioned.

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

Terraform syntax:

```terraform
add_partitioned = {}
```

<a id="canonical-1121321112233333-1101233202200233-0331122223300010-0331233121032101-0103030212003203-3113222203331223-0133230323121213-0113122001101203"></a>

## Direct properties — add_partitioned / 312200323320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020131113113113-1222211203321123-2303102310220012-1323110002311011-2323211301012011-3130023123313230-2100322113032301-0322032112233330"></a>

## Next pages — add_partitioned / 312200323320 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0213002321303031-3200322021113130-3120313220233100-0313331231203332-0301220311230132-0033000312233230-0221310333001103-3012223000332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313301322313321-0021113013103133-2223213302231200-0120031020112011-1000300011001111-1201310312012112-0122301111022311-0212133133003001"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.add_secure — add_secure / 312121103120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.add_secure

<a id="canonical-2210203101310221-1121133030010130-2131212201233021-0233312101110333-2112131320112130-2000032233312100-2300013011212310-0121232132031011"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
add_secure = {}
```

<a id="canonical-1110121012222311-0201131010113221-2003203110331133-0302120302131232-0331110222213012-3212111101230123-3322201110001201-3312322122320300"></a>

## Direct properties — add_secure / 312121103120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322213013311013-3011231331103332-1002121322221202-0111031203232120-1111120110233020-3201110333011330-1322233132013202-0313101211030302"></a>

## Next pages — add_secure / 312121103120 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1231230310231311-3031012003130120-0203331001131021-3303121233230332-0000320021103021-2122031210201132-0221212102002000-1020101120000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303111232312321-2112231120231011-3133022113222300-3122330300233213-1321203210031030-2121112201100011-0132003212322200-0102100111321213"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain — ignore_domain / 312110221332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain

<a id="canonical-1321232010010302-0301323230010302-0331121101130230-3330011201212113-2123000011311232-2211313311311320-3212122022100031-0110201222002321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore domain.

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

Terraform syntax:

```terraform
ignore_domain = {}
```

<a id="canonical-3320001121233013-1133201200012122-2012331111323310-3003123231212233-1213133301301002-3313203221122333-3112331010213321-0121110013001200"></a>

## Direct properties — ignore_domain / 312110221332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211133122211023-1300033231033223-2312223021210302-2322020312130023-0323321103112323-3011102213031120-1333002120201023-3030001110322123"></a>

## Next pages — ignore_domain / 312110221332 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3000330231110221-3010013022021133-0111331201233132-2023302321123223-1133033122331331-1310300003223121-1301302023322230-0130102220120122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302030331331230-0211300311200221-3103201110203013-0031100230210201-1221000031120332-3300032221112002-0212011000002202-0032221212112333"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry — ignore_expiry / 023000123002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry

<a id="canonical-0201022101233001-1333131213020323-3333312100013310-3311113312220202-3231020310103120-2331311302010033-0023302123323102-0223031313221300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore expiry.

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

Terraform syntax:

```terraform
ignore_expiry = {}
```

<a id="canonical-2212320331201000-1333102333323121-1100201133122313-2100331102113030-1302102102220111-3100231000030331-2300013220033210-0332131321201212"></a>

## Direct properties — ignore_expiry / 023000123002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303310000112232-1131320110201320-0232111012223232-2333233000313101-1223210222123332-0003023000031003-0212221310233111-0333010323213201"></a>

## Next pages — ignore_expiry / 023000123002 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3033003021303110-2221320232111022-3210213120013012-1202002313131101-0111030110111110-3121230031330010-1320203302233301-2101101020133230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112210120203133-2231012200113130-0223320110213112-2223312333030003-2010101101123002-0103323131120220-2312233233330222-2001032233001220"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly — ignore_httponly / 123311301113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly

<a id="canonical-2013012112312012-0010012333132023-0330113013023311-2121211230123303-0213212332021002-2102030013022211-1001310000312030-0133011101313321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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

Terraform syntax:

```terraform
ignore_httponly = {}
```

<a id="canonical-0120312333100120-3022301213320121-2013022102021201-0001322020022123-0223300200001113-1021021010322103-2210131101200302-2220323311020200"></a>

## Direct properties — ignore_httponly / 123311301113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013000003122030-3303031301323132-2222010221300102-1222323010320202-2131221032101003-3231012122231030-0211100112300331-0002310013221320"></a>

## Next pages — ignore_httponly / 123311301113 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3332100132012110-1201031310330221-3300001231012323-2123023123313110-1320002032122101-2110030113013200-3000320302313022-0313132123300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203101220110021-1313030021231301-0110101333313103-3133010123131231-3000301131011203-2002313010322000-2320232003201013-0011011111010221"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age — ignore_max_age / 230300310310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age

<a id="canonical-0232310221211123-2322202222123001-0022002313313131-0010323032030030-0010330320202320-1321033132220303-1330301302020211-3332203003130301"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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

Terraform syntax:

```terraform
ignore_max_age = {}
```

<a id="canonical-3322331233102133-0120310310021100-2012210110300230-0130021232321133-2311233310310231-1232000011303101-0220123313203032-1333132132201221"></a>

## Direct properties — ignore_max_age / 230300310310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301212113232000-1113221330333022-3101012122013111-0213101110003002-2032323121202211-3202302321120220-3232202202022231-3300212130002021"></a>

## Next pages — ignore_max_age / 230300310310 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0011012023201321-2203022033203003-2130131111312003-1202010321023003-2031320331333121-1133130221220202-3130312111111113-3112011213003301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331300213011113-0110330202112020-2221102011333023-0002000110013220-2320213330311000-1112211023211213-3123200123331101-0301200033030303"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned — ignore_partitioned / 231231022310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned

<a id="canonical-0321022301331033-0010222123332202-2202200320110023-2313003312020222-1123102233300003-2321023200123232-0131032011131222-1013033011202002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore partitioned.

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

Terraform syntax:

```terraform
ignore_partitioned = {}
```

<a id="canonical-0132133321011203-3033100101100020-3303131203323113-1121022333130030-3201220200202211-0200101023212210-2022113100231100-0001210023302113"></a>

## Direct properties — ignore_partitioned / 231231022310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101123321300110-0122112313201313-0330233322332132-1202021333321122-1313112121031100-1231010120220120-3102022210233022-2301123210211113"></a>

## Next pages — ignore_partitioned / 231231022310 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3130303020100023-1032211010221013-0011013101110003-2100102032021202-2002232131231101-1112031231033121-3231013030113202-1311212333001023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320322213211321-3231233021133131-1132003000223320-3302221320131133-1322320011133010-3030323010221003-0000321001030211-0131312303210310"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_path — ignore_path / 303011333201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_path

<a id="canonical-1132221220132031-2203322002022021-3203032023232331-0132313021132102-3200300221100132-3332311220311223-1300303210201330-1202013202100223"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_path = {}
```

<a id="canonical-2022000022121212-0012111221311132-0130122310022211-0303223000031332-3013230323111131-2320221020133223-3333131121110000-2230233112010211"></a>

## Direct properties — ignore_path / 303011333201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023030121232310-0211020320221313-3301301013030220-1023331231011001-1101313233210132-0220202320221022-2030213302330332-3111300330202211"></a>

## Next pages — ignore_path / 303011333201 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1113003101311201-1030323301021312-3320212231220300-2131010001120201-1300010123101211-1021011203220320-2223002133203022-3330320022311103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230311022123332-2030233031300233-2002031202001023-2222212321021220-3310201303132010-1033112012033110-2301101102102033-2230212233102221"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite — ignore_samesite / 311321100312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite

<a id="canonical-2020232020222332-1010030113130321-2131222102010120-1022220300223032-2031000121001333-0212331221331111-2031301022223211-2131013232121302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_samesite = {}
```

<a id="canonical-1023103002121001-2333001123032302-2323112233233200-2330000332002031-3310131100011302-2233112121012330-2113320123323320-2103013121131311"></a>

## Direct properties — ignore_samesite / 311321100312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010331310122002-1301101131111331-2000032001130300-0211302102322111-3213223121110211-3310221102111031-2021010211320233-0312313023012303"></a>

## Next pages — ignore_samesite / 311321100312 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0121100310301220-3302022132310000-0002023211103222-1331230311320323-0101322131033012-3212012122033212-0102321221032303-1231203220020120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013112330121311-2003212113102200-0121131321130033-2302321110023201-1001312233102133-0133211120313111-1020112302333003-0332001203331111"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure — ignore_secure / 313111001130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure

<a id="canonical-1103212002112230-3033121222223213-1212110202321302-0021121323023122-0232230030021202-2233122001103300-1031011202210212-0303203020131013"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_secure = {}
```

<a id="canonical-1023332133210032-0232013330201301-3110222303032113-1122001321012010-0321321230111313-3112023002300323-2122122102022122-2303310313331201"></a>

## Direct properties — ignore_secure / 313111001130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203013033230133-0003100003002110-1023112012033030-0200010011122333-2133102022320312-3331313101312211-0012013103132001-2212100223013102"></a>

## Next pages — ignore_secure / 313111001130 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0330032312013200-2213310211110302-3311111113310303-1020202111102302-2331212002020232-1032313002232013-1211000230200012-1203301033112132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230202103202030-3133033230020123-1330103031123300-0103300032220230-3210232010220131-1221101020103303-0302103132131311-3323120310322223"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_value — ignore_value / 203311211113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_value

<a id="canonical-3021113133012033-0100300103233230-2332312233222012-2133313311111320-1031131103121021-2133020300103221-1200000021331021-0313112123011321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore value.

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

Terraform syntax:

```terraform
ignore_value = {}
```

<a id="canonical-2303003202222112-0310130333032113-3031322131121210-2021211312223123-1113312221030233-3003133101021112-2231130030303113-1001223320021132"></a>

## Direct properties — ignore_value / 203311211113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203212333103320-3020012213130300-2311003230231020-2233001222230212-0331000200030100-1310220233321133-1022332212300233-0103200011001333"></a>

## Next pages — ignore_value / 203311211113 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2122100130200302-2010033202110221-2101223332201321-1312310321132322-1313033110030110-2323000311312221-3030331110000103-0032112303312122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011301100010112-1010110313020030-1012321313321111-1313213033233200-1202101301000103-2331311323030113-0313220101123102-2000001030013312"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax — samesite_lax / 112033200221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax

<a id="canonical-2102130012003111-2212232132223221-1021310112321301-1121223001321102-1202102322213013-2230000303011330-3220301030113311-0232200212301331"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
samesite_lax = {}
```

<a id="canonical-2321113113011021-3221030210012323-2233031303031300-2022221003302233-2310311002332310-3233232003322000-1203013123320103-3230202310003333"></a>

## Direct properties — samesite_lax / 112033200221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111002231333311-2302112313211001-0320221112321133-0013013000330103-3323001023030201-2012223201310012-0311203112132012-1010031102313100"></a>

## Next pages — samesite_lax / 112033200221 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3213021033222202-3012330311220203-1130031023201103-1302120001111331-3030131223120232-2212010231121301-0013330311320112-1233330332212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223211013330113-0122333333303213-2131101101303112-3302012220303011-3221023133120033-3313221300301132-1032032310131120-3301100121302221"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.samesite_none — samesite_none / 312122313131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.samesite_none

<a id="canonical-2301332010021220-2222111301221210-1021303213323233-3311033311022031-1332301110032123-2113323100111013-0113011031311213-0232021012200022"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
samesite_none = {}
```

<a id="canonical-0312023102202032-2303323100213220-2203302012133222-2121002010232313-2301011213002231-2011232321323312-0231310212222312-3220320021321310"></a>

## Direct properties — samesite_none / 312122313131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103300333110012-1313221121022333-2011010232211213-0020033332122333-3210200122113323-1202300300032330-0313210101301120-3223122101123321"></a>

## Next pages — samesite_none / 312122313131 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0111123301032121-3302113132300132-1031023023123223-1322221302020311-0103222130111210-2302010031223231-2032223221033320-2033012013022313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310232233211222-2201300201210312-3333210331011131-1311313001311212-3310211330130001-0223211202323230-1312033230123213-2031030223210231"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict — samesite_strict / 303211311000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict

<a id="canonical-3020200210230222-1131210121203310-0303213223322102-0201312333033303-1333032031133201-0310220310222021-3331122001002131-2111231330123232"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
samesite_strict = {}
```

<a id="canonical-2332213030110002-1130111123002031-2312331000233311-1310200220121103-2210121330033231-2000032000120322-0103233321120032-0302220312030220"></a>

## Direct properties — samesite_strict / 303211311000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031021020333011-0300130013313123-3123320333221132-1220323302300123-3111112103122300-2103300020113300-3200030132223000-1111121202333121"></a>

## Next pages — samesite_strict / 303211311000 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
