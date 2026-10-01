---
page_title: "xcsh_route reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route reference."
---

# xcsh_route reference

<a id="canonical-0221223123303112-1132311121003011-1101113331001232-1133021113320022-3013323332321333-2303133030221000-1312331032102213-2010100302303203"></a>

## routes.route_destination.spdy_config — spdy_config / 031030323013 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-002.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.spdy_config

<a id="canonical-0030003212102300-1303002203222302-0012302330132121-3322003213202013-2302221310210203-3332130210330001-0312101133322231-0012213100003100"></a>

Type: `"object"`. single nested block, Optional.

Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'SPDY/3.1'
Configuration to allow UPGRADE of connection to SPDY and any additional tuning With configuration to
allow SPDY upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols
'Upgrade'..

Upstream description:

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

Terraform syntax:

```terraform
spdy_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332132111201103-0302031212220033-0300113300001222-3330012223233203-1130222210131233-3212310031122221-2311320023010230-2012210113330330"></a>

## Direct properties — spdy_config / 031030323013 / 3

<a id="canonical-3123311300030233-2123110021122301-2012210221223130-3012121032221111-3330210021300003-0021203222001111-3100001231322013-2303320203332323"></a>

<a id="canonical-2201101032110321-3130103110320202-1030322320101201-1101023300010222-0003132221311302-2233111212222133-0310032121232113-0231123223213033"></a>

## use_spdy property — spdy_config / 031030323013 / 4

Type: `"bool"`. Optional.

Specifies that the HTTP client connection to this route is allowed to upgrade to a SPDY connection.

Upstream description:

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

<a id="canonical-0201110322331011-0330010301033110-0231212302123332-1112330231222033-1110001201232213-0231311021331303-2011231011003102-1002021303022113"></a>

## Next pages — spdy_config / 031030323013 / 5

- [routes.route_destination](resources--route--reference--group-002.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-0223313132331320-3312320001233000-1122103111302032-2311223023000012-0002221013031321-2022311222102223-3033032231010300-1312132322100022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132301023323220-2232221010202320-0301233333220221-2310113330120323-1312220222020232-2013211313300000-1023131323102303-3010222022331002"></a>

## routes.route_destination.web_socket_config — web_socket_config / 300102202221 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-002.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.web_socket_config

<a id="canonical-3101002232312011-1211310201000221-2323130121302321-0210310032312112-0020010111312100-1232122231102203-2230232321313322-3312113132221333"></a>

Type: `"object"`. single nested block, Optional.

Configuration to allow Websocket Request headers of such upgrade looks like below 'connection',
'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce
following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'.

Upstream description:

Configuration to allow Websocket

Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'websocket'

With configuration to allow websocket upgrade, ADC will produce following response 'HTTP/1.1 101
Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'

Receipt-pinned upstream constraints:

```json
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
web_socket_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311201113021300-0121001023031003-0101130211213003-2220212120320121-2212001223312011-2131233303023312-0120303302021312-1112002011332011"></a>

## Direct properties — web_socket_config / 300102202221 / 3

<a id="canonical-1113111201133330-1023102023302213-3112232113230133-2211301220013013-3222101020120300-3322023031001122-3203323200221220-0120103221113331"></a>

<a id="canonical-2332033200112122-2222123130122210-1212211321321302-2221231031221202-2003111112111121-3132011213100223-2211302000230213-2131302323333211"></a>

## use_websocket property — web_socket_config / 300102202221 / 4

Type: `"bool"`. Optional.

Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket
connection.

Upstream description:

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

<a id="canonical-1033112123301313-1300202213233022-3000033111213222-1200102330313321-3202222000012021-2121001202022221-3212232121230102-2300030011232021"></a>

## Next pages — web_socket_config / 300102202221 / 5

- [routes.route_destination](resources--route--reference--group-002.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-0033133232132220-1302111112032120-3000302211110332-3011333112200100-0032120221023020-1101332301231111-1032033302231223-1323030110311233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233121221012001-2310013102110301-3322233123110220-0210300132103321-3331202112323011-1323322231310320-2202123001312100-1120221202212230"></a>

## routes.route_direct_response — route_direct_response / 203013222033 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.route_direct_response

<a id="canonical-2302300123131210-3213103312320230-3310001122100100-2321123102113011-1202212130102111-2103211033323100-0000112122332101-2001013102102011"></a>

Type: `"object"`. single nested block, Optional.

Send this direct response in case of route match action is direct response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_code")}
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
route_direct_response {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020023222301203-2231120311123210-1212000122131021-1002133003020230-0103000222312322-2021100013013101-3110210131101032-2230231120321232"></a>

## Direct properties — route_direct_response / 203013222033 / 3

<a id="canonical-0222302303020322-3312111332232133-3113311202113321-0301303122100303-2101201311310001-3223302220120231-3233222213233331-1011112302133331"></a>

<a id="canonical-1103322300300202-2112202113031120-1130120000330030-0113112332303013-1333231321003333-1332120333313302-3010323012222320-0223300003103201"></a>

## response_body_encoded property — route_direct_response / 203013222033 / 4

Type: `"string"`. Optional.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML.

Upstream description:

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". Base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1121110031003012-0213231211323312-2123332020010221-0310030111300302-1300120221030231-0121002300230111-2030030021333322-0311203130333202"></a>

<a id="canonical-2301313230023001-1301231222310211-0102100120010232-2330033220122313-2312101013123320-1320002112111303-0100230032111301-0113223130001101"></a>

## response_code property — route_direct_response / 203013222033 / 5

Type: `"number"`. Optional.

Response Code. Response code to send.

Upstream description:

Response code to send.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(100, 599),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3011210313330133-1332232101123232-0031132301010011-1332212022211320-0331202032001220-3212102003220313-1110303321113320-1231003221022222"></a>

## Next pages — route_direct_response / 203013222033 / 6

- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-3030203133200300-2232032232300021-2021302312232013-2012023332102223-2002132031320330-1223213213033033-3233302022003331-1303231221020021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220002112122201-0031332233103023-1221320331313213-1023210213223020-2112101333123320-1032122011230321-2303220132002111-0332131201223310"></a>

## routes.route_redirect — route_redirect / 031233013023 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.route_redirect

<a id="canonical-2133332303131212-2113132023333213-3232122202313313-0300313003133233-3011011131333210-0010102301013122-2011110201201133-2223120002333222"></a>

Type: `"object"`. single nested block, Optional.

Route redirect parameters when match action is redirect.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path_redirect",
    "prefix_rewrite"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "replace_params"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "retain_all_params"),
  validators.ConflictingObjectAttributes("replace_params",
    "retain_all_params")}
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
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

Terraform syntax:

```terraform
route_redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123123103010122-3223003112220300-1020320001223022-3312221301113122-2212233022111133-1030131123002010-2302311032210010-0110211010200003"></a>

## Direct properties — route_redirect / 031233013023 / 3

<a id="canonical-0101010203031003-0230330123212111-0122102202320322-0030210232212021-1120120032332120-3030210302022322-3033101303213312-3031021030031331"></a>

<a id="canonical-0213222300302102-3200330113203122-2230032103201311-2211011322230130-3312112311000302-3300122113220220-1232031101011320-3311200232110211"></a>

## host_redirect property — route_redirect / 031233013023 / 4

Type: `"string"`. Optional.

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

<a id="canonical-2321122203200303-2303333300122111-0011200122023233-1303101211302100-2302021111020300-3203322033120320-1100123302120223-2112200202330333"></a>

<a id="canonical-2013311232013223-2222021232132002-2322322012000333-3321321102121222-3301323002033211-0001212301312112-3233312331310021-3321032022332101"></a>

## path_redirect property — route_redirect / 031233013023 / 5

Type: `"string"`. Optional.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Upstream description:

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

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

<a id="canonical-2133231130102301-0030333130120332-1203322110022230-2122132321233330-3010330213011322-2303233302103310-3303323200123100-0121211201310310"></a>

<a id="canonical-0013203212011000-2231200331301203-0100212213022331-2221232302133313-0132212332123303-0003103100321313-0121100300312301-3221001021301001"></a>

## prefix_rewrite property — route_redirect / 031233013023 / 6

Type: `"string"`. Optional.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Upstream description:

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

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

<a id="canonical-2111212001231200-3011023330002222-3100002230103332-3303130312113302-1131013220011203-3030101113312303-2333201203000013-2031123132031023"></a>

<a id="canonical-3333203330202322-1233131210131321-3233220120321311-1121131123232122-1100213133102011-3220211323222231-3020323120112330-3122100313303322"></a>

## proto_redirect property — route_redirect / 031233013023 / 7

Type: `"string"`. Optional.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Upstream description:

Swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or
HTTPS When incoming-proto option is specified, swapping of protocol is not done.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("incoming-proto",
    "http",
    "https"),
}
```

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
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](resources--route--reference--group-003.md#canonical-2112000102133131-3330233010323032-2330110232311113-1222013003331011-2113100211123113-0122010001101103-3012000133322130-0130312202323132): complete subsection reference.

<a id="canonical-1310012213020231-0021013011110001-0001123211220323-1210231031011330-0323002121101322-1120203232133303-0102213133121203-3011222130132001"></a>

<a id="canonical-0112321323121103-0300120023032020-1100100131032132-1211011300102313-1102333301020101-3102113233320030-1330223213321032-1020222320300103"></a>

## replace_params property — route_redirect / 031233013023 / 8

Type: `"string"`. Optional.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2130203220230202-0120301232033223-1203032010212222-1112301300010311-2012223123120031-3220330300000003-0303030302003330-0200322111102321"></a>

<a id="canonical-2020021010003100-1101000302333111-0222222123031212-2310131203230130-1313201221313332-3010311223333120-0311101122330221-3120130113030333"></a>

## response_code property — route_redirect / 031233013023 / 9

Type: `"number"`. Optional.

The HTTP status code to use in the redirect response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(599),
}
```

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
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](resources--route--reference--group-003.md#canonical-0312223003131122-3221203023023111-1111203122233310-1232321112210333-1002123301033232-0233302300300123-0101113332310031-1011301202103223): complete subsection reference.

<a id="canonical-3200201302021320-1231123311223203-1010110020013202-0023122301210022-3223023022221323-0212223231313021-3021032313323330-3201210102213023"></a>

## Next pages — route_redirect / 031233013023 / 10

- [routes.route_redirect.remove_all_params](resources--route--reference--group-003.md#canonical-2112000102133131-3330233010323032-2330110232311113-1222013003331011-2113100211123113-0122010001101103-3012000133322130-0130312202323132)
- [routes.route_redirect.retain_all_params](resources--route--reference--group-003.md#canonical-0312223003131122-3221203023023111-1111203122233310-1232321112210333-1002123301033232-0233302300300123-0101113332310031-1011301202103223)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-2112000102133131-3330233010323032-2330110232311113-1222013003331011-2113100211123113-0122010001101103-3012000133322130-0130312202323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122323022102112-2022203333032102-3001302312001003-2023322101111101-0230010312130003-3323213300210233-2321111122112221-3312331020131333"></a>

## routes.route_redirect.remove_all_params — remove_all_params / 103233032210 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_redirect](resources--route--reference--group-003.md#canonical-3030203133200300-2232032232300021-2021302312232013-2012023332102223-2002132031320330-1223213213033033-3233302022003331-1303231221020021)
- routes.route_redirect.remove_all_params

<a id="canonical-1130112110102132-0323112211230000-1232021131331331-3101130220131031-0111301330030320-0100103011021211-3213122000121013-1230222121201101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for remove all params.

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
remove_all_params = {}
```

<a id="canonical-3000002013001111-3302300221101203-2133001300221320-0310031211003003-1013030202213331-2112112002202210-0100320033013302-1031100001201031"></a>

## Direct properties — remove_all_params / 103233032210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103303322001033-0333233111030231-1020211131322321-3212120100210011-2031223221203113-1111120331033001-3103022310211231-2001112100221232"></a>

## Next pages — remove_all_params / 103233032210 / 4

- [routes.route_redirect](resources--route--reference--group-003.md#canonical-3030203133200300-2232032232300021-2021302312232013-2012023332102223-2002132031320330-1223213213033033-3233302022003331-1303231221020021)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-0312223003131122-3221203023023111-1111203122233310-1232321112210333-1002123301033232-0233302300300123-0101113332310031-1011301202103223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020302112321200-2333231023130112-0220303333033201-0100310312210230-0003010113232001-3121011313230322-2021231132113310-2202333131102011"></a>

## routes.route_redirect.retain_all_params — retain_all_params / 203012120231 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_redirect](resources--route--reference--group-003.md#canonical-3030203133200300-2232032232300021-2021302312232013-2012023332102223-2002132031320330-1223213213033033-3233302022003331-1303231221020021)
- routes.route_redirect.retain_all_params

<a id="canonical-3233111231320202-3103213010231200-1000333130330113-0120203100231100-0322122110033111-3021201220220321-3011313123021331-3033333032201200"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

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
retain_all_params = {}
```

<a id="canonical-1113231012231020-0002203223101333-3323031023113000-0110032032200123-1133123321132222-1021111331321220-3130020302322211-1223233300300023"></a>

## Direct properties — retain_all_params / 203012120231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230131322023000-0231321323021000-1001231212302123-1000231223222132-0223023303322211-3210223220113311-2303131311111303-2221321112321110"></a>

## Next pages — retain_all_params / 203012120231 / 4

- [routes.route_redirect](resources--route--reference--group-003.md#canonical-3030203133200300-2232032232300021-2021302312232013-2012023332102223-2002132031320330-1223213213033033-3233302022003331-1303231221020021)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-0223310311031311-1102132201322103-0030022030223013-3213023033100232-1022023332031231-1310033332022211-3110312333301203-2212302001232030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332321232321313-1301131200303003-1002012310110033-0203002010203012-3211310213021030-3222333223210100-1211130021322230-0002122210102120"></a>

## routes.service_policy — service_policy / 321123221102 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.service_policy

<a id="canonical-2023323222222023-1122101000020101-0133022030032131-3110002210200313-0221223013131021-1331312110320202-3111210010012031-2010022313202003"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
service_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020200130333331-1221010301103120-1321200111003300-1322311310132032-3333220122033020-0032202310123332-3000232231202312-3022001333100123"></a>

## Direct properties — service_policy / 321123221102 / 3

<a id="canonical-0212233320321030-1201131112030001-3332030223133030-2022210001322213-0213221132002203-1020301233333320-2032231013032230-2120121200230223"></a>

<a id="canonical-1120003200303330-1320002130121320-1313011103022311-0020303023020200-3231120210312011-3033223332110003-1023233313013300-0213131121230222"></a>

## disable_spec property — service_policy / 321123221102 / 4

Type: `"bool"`. Optional.

Exclusive with \[\] disable service policy at route level, if it is configured at virtual-host
level.

<a id="canonical-1111012111111221-2101321020233332-0301200012100000-1102122002032022-3002011321132023-1313111013112313-0110110111120102-1000031301232302"></a>

## Next pages — service_policy / 321123221102 / 5

- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-3113313331301130-2022203212231123-3330200013112321-1303112232000301-0332001232221002-3313202122033003-1323230113203023-3011123020002001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313001200301230-2302032333211113-2303131022201123-2223222223103020-3310133202211020-3113110121213032-2233203203113333-3022202112303011"></a>

## routes.waf_exclusion_policy — waf_exclusion_policy / 201033211233 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.waf_exclusion_policy

<a id="canonical-2002013313001200-2211230001233100-3213002023010302-0003113113213113-1220101021221323-2310310033302131-2033311313033122-3303301021200331"></a>

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
waf_exclusion_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023321002330111-0232212132312212-3003013310101032-3030232032323133-0022233212031202-3310031201001000-0131231113032102-0102210123133121"></a>

## Direct properties — waf_exclusion_policy / 201033211233 / 3

<a id="canonical-0323333110321132-1321231002100031-0332113312313302-2121031002221313-1132233200120001-2311323000001222-0120011311300232-2131030120330001"></a>

<a id="canonical-3111003221131220-1310023313203103-2111331031313332-2303310021131000-2233123020132113-0311000101100023-0311301331112123-1133021120200203"></a>

## name property — waf_exclusion_policy / 201033211233 / 4

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

<a id="canonical-3320223201311312-1230332200221132-2023211032030213-2311120001232102-1330201233101010-3312101333222000-0032101301312212-3231202330011003"></a>

<a id="canonical-0313111112232110-3231010200013213-3333221122301002-2321320102312312-2110220101231302-3000131213123101-1000213131233010-3203310023200331"></a>

## namespace property — waf_exclusion_policy / 201033211233 / 5

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

<a id="canonical-0112303323221002-0112010132321100-2020213300131313-2313332203100003-1110211220313310-3011301330230102-0320120110211120-0220022210011311"></a>

<a id="canonical-3220312123000012-2302112311302200-2031121132300133-2130003322312321-1121012101313202-1112312322323132-2111320030000231-2122202320233222"></a>

## tenant property — waf_exclusion_policy / 201033211233 / 6

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

<a id="canonical-2332322012102102-1320030332031001-1010110322211102-1321302021030333-1322022012032312-0220330110101100-1312322130301020-3100020113110120"></a>

## Next pages — waf_exclusion_policy / 201033211233 / 7

- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211223001102212-3303201100103210-3330303002132210-1100203222113213-1221031130233202-2333321331122320-0102303122011233-2101033321321320"></a>

## routes.waf_type — waf_type / 200011220212 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.waf_type

<a id="canonical-2001321113131313-1313201110111212-3122331232022222-2331202122321020-0221120200230002-3200330210012331-0322102302101033-3110033121012002"></a>

Type: `"object"`. single nested block, Optional.

WAF instance will be pointing to an app\_firewall object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall",
    "disable_waf"),
  validators.ConflictingObjectAttributes("app_firewall",
    "inherit_waf"),
  validators.ConflictingObjectAttributes("disable_waf",
    "inherit_waf")}
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
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

Terraform syntax:

```terraform
waf_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132000222002212-2010320103111330-2302320123013110-0030111003200211-2000323001121100-0210023123121032-3300332230203301-0031233000022201"></a>

## Direct properties — waf_type / 200011220212 / 3

- [app_firewall](resources--route--reference--group-003.md#canonical-0200132232312312-1323232312210200-0011013333320012-3210230132301211-2013022232102130-2010311100203312-1233212201021313-2302110120000122): complete subsection reference.

- [disable_waf](resources--route--reference--group-003.md#canonical-2312222012021023-0130011200022021-2111231323310220-0123320233110003-3320301100010120-3210220010033100-2100321203330010-0023002113330003): complete subsection reference.

- [inherit_waf](resources--route--reference--group-003.md#canonical-3311301022131220-1131221221133322-2320231022201202-2033021133023220-3000103220310303-2210232232001201-2121022131100322-2012102321223011): complete subsection reference.

<a id="canonical-2201221111210231-3323000123003012-1232210312333222-2311331120333302-0313101233021223-2131322032311021-1220200320331132-0103002320332133"></a>

## Next pages — waf_type / 200011220212 / 4

- [routes.waf_type.app_firewall](resources--route--reference--group-003.md#canonical-0200132232312312-1323232312210200-0011013333320012-3210230132301211-2013022232102130-2010311100203312-1233212201021313-2302110120000122)
- [routes.waf_type.disable_waf](resources--route--reference--group-003.md#canonical-2312222012021023-0130011200022021-2111231323310220-0123320233110003-3320301100010120-3210220010033100-2100321203330010-0023002113330003)
- [routes.waf_type.inherit_waf](resources--route--reference--group-003.md#canonical-3311301022131220-1131221221133322-2320231022201202-2033021133023220-3000103220310303-2210232232001201-2121022131100322-2012102321223011)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-0200132232312312-1323232312210200-0011013333320012-3210230132301211-2013022232102130-2010311100203312-1233212201021313-2302110120000122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333221203222030-0323031111233021-3000303232212103-3003211330133230-1300302220220101-1230321132302032-1031200111312102-1221210110133232"></a>

## routes.waf_type.app_firewall — app_firewall / 031213310133 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301)
- routes.waf_type.app_firewall

<a id="canonical-2231021231220033-0233233122220013-1111123320222301-3022211031003332-0300100210121302-1331002312111300-3302022222231002-0021032131203131"></a>

Type: `"object"`. single nested block, Optional.

List of references to the app\_firewall configuration objects.

Upstream description:

A list of references to the app\_firewall configuration objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("app_firewall")}
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
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332332011310101-2102212330131023-2112211102002122-3322312231100123-1031102213133212-1111300210221020-3302033102122002-2323203022210221"></a>

## Direct properties — app_firewall / 031213310133 / 3

- [app_firewall](resources--route--reference--group-003.md#canonical-2311111213330130-1003101232130330-3221320120310311-3233210300103003-3102313130023232-2302313001002221-0100030100023323-0310113023320310): complete subsection reference.

<a id="canonical-0220323221220313-0001112023203100-0122132213122203-2231313230020230-2320110011110001-0111201211112311-2101223301023110-0211312130032221"></a>

## Next pages — app_firewall / 031213310133 / 4

- [routes.waf_type.app_firewall.app_firewall](resources--route--reference--group-003.md#canonical-2311111213330130-1003101232130330-3221320120310311-3233210300103003-3102313130023232-2302313001002221-0100030100023323-0310113023320310)
- [routes.waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-2311111213330130-1003101232130330-3221320120310311-3233210300103003-3102313130023232-2302313001002221-0100030100023323-0310113023320310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011110123321200-1202220213311320-3021301001133110-3310132121312002-2000231313322001-3233003001012223-2011200113321223-0301311330011311"></a>

## routes.waf_type.app_firewall.app_firewall — app_firewall / 130222222120 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301)
- [routes.waf_type.app_firewall](resources--route--reference--group-003.md#canonical-0200132232312312-1323232312210200-0011013333320012-3210230132301211-2013022232102130-2010311100203312-1233212201021313-2302110120000122)
- routes.waf_type.app_firewall.app_firewall

<a id="canonical-2230210213222132-1012121003022303-0323000302220103-3321133300210232-3211231012303323-1310203122333131-3311113033023200-0213002021322231"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123021320203232-3110102011301310-3300302310323133-1213223012302022-0033212202320330-3120003303101331-1110032303221111-1302323113203331"></a>

## Direct properties — app_firewall / 130222222120 / 3

<a id="canonical-0221200322002023-1111100102000210-0022012102232211-2131103030222311-2003300313323301-1231311300220312-0010332013321010-1212012202001003"></a>

<a id="canonical-3323232321022001-2223131020102222-2112003121013212-3232222003110021-2210322111133203-0133231123322213-2010031030111113-3333322332313200"></a>

## kind property — app_firewall / 130222222120 / 4

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

<a id="canonical-2113332133021120-1220103233200303-2230331332001002-0023322013133132-1333132123222221-1021002333312120-2033020003322030-1333321132312220"></a>

<a id="canonical-0002031022101030-0121030000331013-0321221220332023-1120330232223120-1132031123230321-2020110130133313-1012220121012011-0203013321031101"></a>

## name property — app_firewall / 130222222120 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3023230311231311-0210103201300133-3023212300121331-2221021202312331-3302003200221011-1332131113221120-2022033100100023-0032333322132230"></a>

<a id="canonical-0111002201202321-3311123300300333-0122101231310013-0130032121312130-2210332120120211-3001121330020101-0200122233310121-0000132201103310"></a>

## namespace property — app_firewall / 130222222120 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-1002113013001210-3323003100130111-3313321303111123-2221132302330032-1323233312030212-3313311302122331-0112123023103011-1201311213223302"></a>

<a id="canonical-2210233210202300-3203001131233011-3030230030000022-2213123102120021-0013321301111032-2233010230003102-1031130020301231-0113210330203000"></a>

## tenant property — app_firewall / 130222222120 / 7

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

<a id="canonical-2121302121131022-3021111201022332-2320202231233313-0322330330202130-1202113003123103-0121310301331031-2210223120012030-2333122130310120"></a>

<a id="canonical-3200220012312322-1121003210002122-2122231210112112-3111033123331310-0010310322310322-2001231011302203-3333211313321022-3011033000312123"></a>

## uid property — app_firewall / 130222222120 / 8

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

<a id="canonical-2310230231012113-0311231110131320-2110333102010200-3103331132303000-2232200213302001-3021300223330123-2130120312130013-3110211322202003"></a>

## Next pages — app_firewall / 130222222120 / 9

- [routes.waf_type.app_firewall](resources--route--reference--group-003.md#canonical-0200132232312312-1323232312210200-0011013333320012-3210230132301211-2013022232102130-2010311100203312-1233212201021313-2302110120000122)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-2312222012021023-0130011200022021-2111231323310220-0123320233110003-3320301100010120-3210220010033100-2100321203330010-0023002113330003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032133102100123-2132313222302121-3203123200202212-1331302220113200-0023311303100311-3323220120131332-1110331223210323-2200311121013331"></a>

## routes.waf_type.disable_waf — disable_waf / 203211210131 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301)
- routes.waf_type.disable_waf

<a id="canonical-3011003032123110-2333031310103322-1333213021123113-0113330312130021-1111112331131121-0333010000212002-3120112133321320-2201031223002331"></a>

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

<a id="canonical-3013300212232112-0102323123331201-2231131111212110-1010221330021031-3323233100011231-2132122311002001-3311313112232233-1133021330002330"></a>

## Direct properties — disable_waf / 203211210131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021331110033012-2110121011331232-3012003231300330-0113101003310212-3023032103032221-0311130300310013-3322200202301032-2301022032102133"></a>

## Next pages — disable_waf / 203211210131 / 4

- [routes.waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-3311301022131220-1131221221133322-2320231022201202-2033021133023220-3000103220310303-2210232232001201-2121022131100322-2012102321223011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023100102200031-0123133211120121-1203101313210230-0022222213331003-1302013203203002-2133033023211010-3032230123200020-0200110230323333"></a>

## routes.waf_type.inherit_waf — inherit_waf / 123230223222 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301)
- routes.waf_type.inherit_waf

<a id="canonical-1231213121221303-0023310201031321-1202211302132031-0232203213331212-2012120322121332-0132103323132111-0233231003033032-2121200121120001"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
inherit_waf = {}
```

<a id="canonical-2003110303133313-1122221320232201-3131232201320003-0132131121312003-1310320110013313-1000113023133222-3000320031013111-0131121230123012"></a>

## Direct properties — inherit_waf / 123230223222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123033011013200-0020121310203222-2010210230333112-0221333033102231-3110000233203323-0311300032003302-2130013123000013-0033102233302222"></a>

## Next pages — inherit_waf / 123230223222 / 4

- [routes.waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)

<a id="canonical-2000122211223303-0323203133120331-3213333010022223-1120331312123010-0233002030003331-1012023100032301-3200122303322132-3110023202032230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100233003010002-0030312120330200-3312030232010311-3323210211132213-3221203012303110-0322031322111300-0302020113103331-1122100103321210"></a>

## timeouts — timeouts / 210321212011 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- timeouts

<a id="canonical-3111321022202013-2032120201220322-1301231002330300-0031231130003330-2001303332311331-2230212213301331-2030333231301120-3130212320122332"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232022323332102-1002123232311003-1001112113123122-1101301203131013-3333233202020000-3312132132212310-2033021331002121-3130323133000112"></a>

## Direct properties — timeouts / 210321212011 / 3

<a id="canonical-1130001230123010-1221321323133311-0332223011301032-1222031132033112-0312023111332000-1030220102100201-3331011330003112-1100012311023001"></a>

<a id="canonical-3130220211302101-2302100311331130-1231203023322111-3033122030321111-2133121332032320-0302102201200001-3032331113201321-1233332132330221"></a>

## create property — timeouts / 210321212011 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0323312333201213-3130331323123302-0301112112123302-3300000201101000-2112201012020120-0302111212132300-1013131300210210-2222020332310212"></a>

<a id="canonical-0031033100220110-2000131103202232-0133333313013320-3130231032203020-2302113000012230-1123021223332202-2231230310120321-1323232313113210"></a>

## delete property — timeouts / 210321212011 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1101103120000211-1211330033202310-1212223212233002-1111001232202201-0121320303030010-1032322300320121-3022222303233011-3223030111021110"></a>

<a id="canonical-0022120001020123-3031131023311021-0302233332211023-3030201322232121-0301010213032030-3201123112032311-2220023321210303-1102000002203013"></a>

## read property — timeouts / 210321212011 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0123103313202120-1031213332201320-2333113113000030-1103221112203330-0103121100203123-3132223020113212-0000203132130333-1332013103312230"></a>

<a id="canonical-0112113323212110-3010101201203223-1001010203220233-2013023112003031-1312230010002312-0111002021121231-1032032022120111-2330233331133013"></a>

## update property — timeouts / 210321212011 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3322020110100232-1112030132202221-2322331323003122-0103131332300231-1100010113202220-2012101201313222-0003311121213322-0123212012201000"></a>

## Next pages — timeouts / 210321212011 / 8

- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
